// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-hsrp package configuration.
// 
// This module contains definitions
// for the following management objects:
//   hsrp: HSRP configuration
// 
// This YANG module augments the
//   Cisco-IOS-XR-snmp-agent-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_hsrp_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_hsrp_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-hsrp-cfg hsrp}", reflect.TypeOf(Hsrp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-hsrp-cfg:hsrp", reflect.TypeOf(Hsrp{}))
}

// HsrpLinklocal represents Hsrp linklocal
type HsrpLinklocal string

const (
    // Manual Linklocal address configuration
    HsrpLinklocal_manual HsrpLinklocal = "manual"

    // Automatic Linklocal address configuration
    HsrpLinklocal_auto HsrpLinklocal = "auto"

    // Automatic legacy-compatible Linklocal address
    // configuration
    HsrpLinklocal_legacy HsrpLinklocal = "legacy"
)

// Hsrp
// HSRP configuration
type Hsrp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface Table for HSRP configuration.
    Interfaces Hsrp_Interfaces

    // HSRP logging options.
    Logging Hsrp_Logging
}

func (hsrp *Hsrp) GetEntityData() *types.CommonEntityData {
    hsrp.EntityData.YFilter = hsrp.YFilter
    hsrp.EntityData.YangName = "hsrp"
    hsrp.EntityData.BundleName = "cisco_ios_xr"
    hsrp.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-hsrp-cfg"
    hsrp.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-hsrp-cfg:hsrp"
    hsrp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    hsrp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    hsrp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    hsrp.EntityData.Children = make(map[string]types.YChild)
    hsrp.EntityData.Children["interfaces"] = types.YChild{"Interfaces", &hsrp.Interfaces}
    hsrp.EntityData.Children["logging"] = types.YChild{"Logging", &hsrp.Logging}
    hsrp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(hsrp.EntityData)
}

// Hsrp_Interfaces
// Interface Table for HSRP configuration
type Hsrp_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Per-interface HSRP configuration. The type is slice of
    // Hsrp_Interfaces_Interface_.
    Interface_ []Hsrp_Interfaces_Interface
}

func (interfaces *Hsrp_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "hsrp"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = make(map[string]types.YChild)
    interfaces.EntityData.Children["interface"] = types.YChild{"Interface_", nil}
    for i := range interfaces.Interface_ {
        interfaces.EntityData.Children[types.GetSegmentPath(&interfaces.Interface_[i])] = types.YChild{"Interface_", &interfaces.Interface_[i]}
    }
    interfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(interfaces.EntityData)
}

// Hsrp_Interfaces_Interface
// Per-interface HSRP configuration
type Hsrp_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface name. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // HSRP MGO slave MAC refresh rate. The type is interface{} with range:
    // 0..10000. The default value is 60.
    MacRefresh interface{}

    // Use burned-in address. The type is interface{}.
    UseBia interface{}

    // Disable HSRP filtered ICMP redirects. The type is interface{}.
    RedirectsDisable interface{}

    // IPv6 HSRP configuration.
    Ipv6 Hsrp_Interfaces_Interface_Ipv6

    // BFD configuration.
    Bfd Hsrp_Interfaces_Interface_Bfd

    // Minimum and Reload Delay.
    Delay Hsrp_Interfaces_Interface_Delay

    // IPv4 HSRP configuration.
    Ipv4 Hsrp_Interfaces_Interface_Ipv4
}

func (self *Hsrp_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + "[interface-name='" + fmt.Sprintf("%v", self.InterfaceName) + "']"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = make(map[string]types.YChild)
    self.EntityData.Children["ipv6"] = types.YChild{"Ipv6", &self.Ipv6}
    self.EntityData.Children["bfd"] = types.YChild{"Bfd", &self.Bfd}
    self.EntityData.Children["delay"] = types.YChild{"Delay", &self.Delay}
    self.EntityData.Children["ipv4"] = types.YChild{"Ipv4", &self.Ipv4}
    self.EntityData.Leafs = make(map[string]types.YLeaf)
    self.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", self.InterfaceName}
    self.EntityData.Leafs["mac-refresh"] = types.YLeaf{"MacRefresh", self.MacRefresh}
    self.EntityData.Leafs["use-bia"] = types.YLeaf{"UseBia", self.UseBia}
    self.EntityData.Leafs["redirects-disable"] = types.YLeaf{"RedirectsDisable", self.RedirectsDisable}
    return &(self.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv6
// IPv6 HSRP configuration
type Hsrp_Interfaces_Interface_Ipv6 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Version 2 HSRP configuration.
    Version2 Hsrp_Interfaces_Interface_Ipv6_Version2

    // The HSRP slave group configuration table.
    SlaveGroups Hsrp_Interfaces_Interface_Ipv6_SlaveGroups
}

func (ipv6 *Hsrp_Interfaces_Interface_Ipv6) GetEntityData() *types.CommonEntityData {
    ipv6.EntityData.YFilter = ipv6.YFilter
    ipv6.EntityData.YangName = "ipv6"
    ipv6.EntityData.BundleName = "cisco_ios_xr"
    ipv6.EntityData.ParentYangName = "interface"
    ipv6.EntityData.SegmentPath = "ipv6"
    ipv6.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv6.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv6.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv6.EntityData.Children = make(map[string]types.YChild)
    ipv6.EntityData.Children["version2"] = types.YChild{"Version2", &ipv6.Version2}
    ipv6.EntityData.Children["slave-groups"] = types.YChild{"SlaveGroups", &ipv6.SlaveGroups}
    ipv6.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv6.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv6_Version2
// Version 2 HSRP configuration
type Hsrp_Interfaces_Interface_Ipv6_Version2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The HSRP group configuration table.
    Groups Hsrp_Interfaces_Interface_Ipv6_Version2_Groups
}

func (version2 *Hsrp_Interfaces_Interface_Ipv6_Version2) GetEntityData() *types.CommonEntityData {
    version2.EntityData.YFilter = version2.YFilter
    version2.EntityData.YangName = "version2"
    version2.EntityData.BundleName = "cisco_ios_xr"
    version2.EntityData.ParentYangName = "ipv6"
    version2.EntityData.SegmentPath = "version2"
    version2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    version2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    version2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    version2.EntityData.Children = make(map[string]types.YChild)
    version2.EntityData.Children["groups"] = types.YChild{"Groups", &version2.Groups}
    version2.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(version2.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv6_Version2_Groups
// The HSRP group configuration table
type Hsrp_Interfaces_Interface_Ipv6_Version2_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The HSRP group being configured. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group.
    Group []Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group
}

func (groups *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "cisco_ios_xr"
    groups.EntityData.ParentYangName = "version2"
    groups.EntityData.SegmentPath = "groups"
    groups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groups.EntityData.Children = make(map[string]types.YChild)
    groups.EntityData.Children["group"] = types.YChild{"Group", nil}
    for i := range groups.Group {
        groups.EntityData.Children[types.GetSegmentPath(&groups.Group[i])] = types.YChild{"Group", &groups.Group[i]}
    }
    groups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(groups.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group
// The HSRP group being configured
type Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. HSRP group number. The type is interface{} with
    // range: 0..4095.
    GroupNumber interface{}

    // Priority value. The type is interface{} with range: 0..255. The default
    // value is 100.
    Priority interface{}

    // Force active if higher priority. The type is interface{} with range:
    // -2147483648..2147483647. The default value is 0.
    Preempt interface{}

    // HSRP Session name (for MGO). The type is string with length: 1..16.
    SessionName interface{}

    // HSRP MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    VirtualMacAddress interface{}

    // Enable use of Bidirectional Forwarding Detection.
    Bfd Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Bfd

    // The HSRP tracked interface configuration table.
    TrackedInterfaces Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces

    // The HSRP tracked interface configuration table.
    TrackedObjects Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects

    // Hello and hold timers.
    Timers Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Timers

    // The HSRP IPv6 virtual linklocal address.
    LinkLocalIpv6Address Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_LinkLocalIpv6Address

    // The table of HSRP virtual global IPv6 addresses.
    GlobalIpv6Addresses Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses
}

func (group *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + "[group-number='" + fmt.Sprintf("%v", group.GroupNumber) + "']"
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = make(map[string]types.YChild)
    group.EntityData.Children["bfd"] = types.YChild{"Bfd", &group.Bfd}
    group.EntityData.Children["tracked-interfaces"] = types.YChild{"TrackedInterfaces", &group.TrackedInterfaces}
    group.EntityData.Children["tracked-objects"] = types.YChild{"TrackedObjects", &group.TrackedObjects}
    group.EntityData.Children["timers"] = types.YChild{"Timers", &group.Timers}
    group.EntityData.Children["link-local-ipv6-address"] = types.YChild{"LinkLocalIpv6Address", &group.LinkLocalIpv6Address}
    group.EntityData.Children["global-ipv6-addresses"] = types.YChild{"GlobalIpv6Addresses", &group.GlobalIpv6Addresses}
    group.EntityData.Leafs = make(map[string]types.YLeaf)
    group.EntityData.Leafs["group-number"] = types.YLeaf{"GroupNumber", group.GroupNumber}
    group.EntityData.Leafs["priority"] = types.YLeaf{"Priority", group.Priority}
    group.EntityData.Leafs["preempt"] = types.YLeaf{"Preempt", group.Preempt}
    group.EntityData.Leafs["session-name"] = types.YLeaf{"SessionName", group.SessionName}
    group.EntityData.Leafs["virtual-mac-address"] = types.YLeaf{"VirtualMacAddress", group.VirtualMacAddress}
    return &(group.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Bfd
// Enable use of Bidirectional Forwarding
// Detection
type Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable BFD for this remote IP. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Interface name to run BFD. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}
}

func (bfd *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "group"
    bfd.EntityData.SegmentPath = "bfd"
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = make(map[string]types.YChild)
    bfd.EntityData.Leafs = make(map[string]types.YLeaf)
    bfd.EntityData.Leafs["address"] = types.YLeaf{"Address", bfd.Address}
    bfd.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", bfd.InterfaceName}
    return &(bfd.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces
// The HSRP tracked interface configuration
// table
type Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface being tracked. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces_TrackedInterface.
    TrackedInterface []Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces_TrackedInterface
}

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces) GetEntityData() *types.CommonEntityData {
    trackedInterfaces.EntityData.YFilter = trackedInterfaces.YFilter
    trackedInterfaces.EntityData.YangName = "tracked-interfaces"
    trackedInterfaces.EntityData.BundleName = "cisco_ios_xr"
    trackedInterfaces.EntityData.ParentYangName = "group"
    trackedInterfaces.EntityData.SegmentPath = "tracked-interfaces"
    trackedInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackedInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackedInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackedInterfaces.EntityData.Children = make(map[string]types.YChild)
    trackedInterfaces.EntityData.Children["tracked-interface"] = types.YChild{"TrackedInterface", nil}
    for i := range trackedInterfaces.TrackedInterface {
        trackedInterfaces.EntityData.Children[types.GetSegmentPath(&trackedInterfaces.TrackedInterface[i])] = types.YChild{"TrackedInterface", &trackedInterfaces.TrackedInterface[i]}
    }
    trackedInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trackedInterfaces.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces_TrackedInterface
// Interface being tracked
type Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces_TrackedInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface being tracked. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Priority decrement. The type is interface{} with range: 1..255. This
    // attribute is mandatory.
    PriorityDecrement interface{}
}

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetEntityData() *types.CommonEntityData {
    trackedInterface.EntityData.YFilter = trackedInterface.YFilter
    trackedInterface.EntityData.YangName = "tracked-interface"
    trackedInterface.EntityData.BundleName = "cisco_ios_xr"
    trackedInterface.EntityData.ParentYangName = "tracked-interfaces"
    trackedInterface.EntityData.SegmentPath = "tracked-interface" + "[interface-name='" + fmt.Sprintf("%v", trackedInterface.InterfaceName) + "']"
    trackedInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackedInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackedInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackedInterface.EntityData.Children = make(map[string]types.YChild)
    trackedInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    trackedInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", trackedInterface.InterfaceName}
    trackedInterface.EntityData.Leafs["priority-decrement"] = types.YLeaf{"PriorityDecrement", trackedInterface.PriorityDecrement}
    return &(trackedInterface.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects
// The HSRP tracked interface configuration
// table
type Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Object being tracked. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects_TrackedObject.
    TrackedObject []Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects_TrackedObject
}

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects) GetEntityData() *types.CommonEntityData {
    trackedObjects.EntityData.YFilter = trackedObjects.YFilter
    trackedObjects.EntityData.YangName = "tracked-objects"
    trackedObjects.EntityData.BundleName = "cisco_ios_xr"
    trackedObjects.EntityData.ParentYangName = "group"
    trackedObjects.EntityData.SegmentPath = "tracked-objects"
    trackedObjects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackedObjects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackedObjects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackedObjects.EntityData.Children = make(map[string]types.YChild)
    trackedObjects.EntityData.Children["tracked-object"] = types.YChild{"TrackedObject", nil}
    for i := range trackedObjects.TrackedObject {
        trackedObjects.EntityData.Children[types.GetSegmentPath(&trackedObjects.TrackedObject[i])] = types.YChild{"TrackedObject", &trackedObjects.TrackedObject[i]}
    }
    trackedObjects.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trackedObjects.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects_TrackedObject
// Object being tracked
type Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects_TrackedObject struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface being tracked. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ObjectName interface{}

    // Priority decrement. The type is interface{} with range: 1..255. This
    // attribute is mandatory.
    PriorityDecrement interface{}
}

func (trackedObject *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_TrackedObjects_TrackedObject) GetEntityData() *types.CommonEntityData {
    trackedObject.EntityData.YFilter = trackedObject.YFilter
    trackedObject.EntityData.YangName = "tracked-object"
    trackedObject.EntityData.BundleName = "cisco_ios_xr"
    trackedObject.EntityData.ParentYangName = "tracked-objects"
    trackedObject.EntityData.SegmentPath = "tracked-object" + "[object-name='" + fmt.Sprintf("%v", trackedObject.ObjectName) + "']"
    trackedObject.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackedObject.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackedObject.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackedObject.EntityData.Children = make(map[string]types.YChild)
    trackedObject.EntityData.Leafs = make(map[string]types.YLeaf)
    trackedObject.EntityData.Leafs["object-name"] = types.YLeaf{"ObjectName", trackedObject.ObjectName}
    trackedObject.EntityData.Leafs["priority-decrement"] = types.YLeaf{"PriorityDecrement", trackedObject.PriorityDecrement}
    return &(trackedObject.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Timers
// Hello and hold timers
type Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Timers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TRUE - Hello time configured in milliseconds, FALSE - Hello time configured
    // in seconds. The type is bool. The default value is false.
    HelloMsecFlag interface{}

    // Hello time in msecs. The type is interface{} with range: 100..3000. Units
    // are millisecond.
    HelloMsec interface{}

    // Hello time in seconds. The type is interface{} with range: 1..255. Units
    // are second. The default value is 3.
    HelloSec interface{}

    // TRUE - Hold time configured in milliseconds, FALSE - Hold time configured
    // in seconds. The type is bool. The default value is false.
    HoldMsecFlag interface{}

    // Hold time in msecs. The type is interface{} with range: 100..3000. Units
    // are millisecond.
    HoldMsec interface{}

    // Hold time in seconds. The type is interface{} with range: 1..255. Units are
    // second. The default value is 10.
    HoldSec interface{}
}

func (timers *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_Timers) GetEntityData() *types.CommonEntityData {
    timers.EntityData.YFilter = timers.YFilter
    timers.EntityData.YangName = "timers"
    timers.EntityData.BundleName = "cisco_ios_xr"
    timers.EntityData.ParentYangName = "group"
    timers.EntityData.SegmentPath = "timers"
    timers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timers.EntityData.Children = make(map[string]types.YChild)
    timers.EntityData.Leafs = make(map[string]types.YLeaf)
    timers.EntityData.Leafs["hello-msec-flag"] = types.YLeaf{"HelloMsecFlag", timers.HelloMsecFlag}
    timers.EntityData.Leafs["hello-msec"] = types.YLeaf{"HelloMsec", timers.HelloMsec}
    timers.EntityData.Leafs["hello-sec"] = types.YLeaf{"HelloSec", timers.HelloSec}
    timers.EntityData.Leafs["hold-msec-flag"] = types.YLeaf{"HoldMsecFlag", timers.HoldMsecFlag}
    timers.EntityData.Leafs["hold-msec"] = types.YLeaf{"HoldMsec", timers.HoldMsec}
    timers.EntityData.Leafs["hold-sec"] = types.YLeaf{"HoldSec", timers.HoldSec}
    return &(timers.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_LinkLocalIpv6Address
// The HSRP IPv6 virtual linklocal address
type Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_LinkLocalIpv6Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // HSRP IPv6 virtual linklocal address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Linklocal Configuration Type. The type is HsrpLinklocal. The default value
    // is manual.
    AutoConfigure interface{}
}

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_LinkLocalIpv6Address) GetEntityData() *types.CommonEntityData {
    linkLocalIpv6Address.EntityData.YFilter = linkLocalIpv6Address.YFilter
    linkLocalIpv6Address.EntityData.YangName = "link-local-ipv6-address"
    linkLocalIpv6Address.EntityData.BundleName = "cisco_ios_xr"
    linkLocalIpv6Address.EntityData.ParentYangName = "group"
    linkLocalIpv6Address.EntityData.SegmentPath = "link-local-ipv6-address"
    linkLocalIpv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkLocalIpv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkLocalIpv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkLocalIpv6Address.EntityData.Children = make(map[string]types.YChild)
    linkLocalIpv6Address.EntityData.Leafs = make(map[string]types.YLeaf)
    linkLocalIpv6Address.EntityData.Leafs["address"] = types.YLeaf{"Address", linkLocalIpv6Address.Address}
    linkLocalIpv6Address.EntityData.Leafs["auto-configure"] = types.YLeaf{"AutoConfigure", linkLocalIpv6Address.AutoConfigure}
    return &(linkLocalIpv6Address.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses
// The table of HSRP virtual global IPv6
// addresses
type Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A HSRP virtual global IPv6 IP address. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses_GlobalIpv6Address.
    GlobalIpv6Address []Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses_GlobalIpv6Address
}

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses) GetEntityData() *types.CommonEntityData {
    globalIpv6Addresses.EntityData.YFilter = globalIpv6Addresses.YFilter
    globalIpv6Addresses.EntityData.YangName = "global-ipv6-addresses"
    globalIpv6Addresses.EntityData.BundleName = "cisco_ios_xr"
    globalIpv6Addresses.EntityData.ParentYangName = "group"
    globalIpv6Addresses.EntityData.SegmentPath = "global-ipv6-addresses"
    globalIpv6Addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalIpv6Addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalIpv6Addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalIpv6Addresses.EntityData.Children = make(map[string]types.YChild)
    globalIpv6Addresses.EntityData.Children["global-ipv6-address"] = types.YChild{"GlobalIpv6Address", nil}
    for i := range globalIpv6Addresses.GlobalIpv6Address {
        globalIpv6Addresses.EntityData.Children[types.GetSegmentPath(&globalIpv6Addresses.GlobalIpv6Address[i])] = types.YChild{"GlobalIpv6Address", &globalIpv6Addresses.GlobalIpv6Address[i]}
    }
    globalIpv6Addresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(globalIpv6Addresses.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses_GlobalIpv6Address
// A HSRP virtual global IPv6 IP address
type Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses_GlobalIpv6Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. HSRP virtual global IPv6 address. The type is
    // string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}
}

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_Version2_Groups_Group_GlobalIpv6Addresses_GlobalIpv6Address) GetEntityData() *types.CommonEntityData {
    globalIpv6Address.EntityData.YFilter = globalIpv6Address.YFilter
    globalIpv6Address.EntityData.YangName = "global-ipv6-address"
    globalIpv6Address.EntityData.BundleName = "cisco_ios_xr"
    globalIpv6Address.EntityData.ParentYangName = "global-ipv6-addresses"
    globalIpv6Address.EntityData.SegmentPath = "global-ipv6-address" + "[address='" + fmt.Sprintf("%v", globalIpv6Address.Address) + "']"
    globalIpv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalIpv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalIpv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalIpv6Address.EntityData.Children = make(map[string]types.YChild)
    globalIpv6Address.EntityData.Leafs = make(map[string]types.YLeaf)
    globalIpv6Address.EntityData.Leafs["address"] = types.YLeaf{"Address", globalIpv6Address.Address}
    return &(globalIpv6Address.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv6_SlaveGroups
// The HSRP slave group configuration table
type Hsrp_Interfaces_Interface_Ipv6_SlaveGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The HSRP slave group being configured. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup.
    SlaveGroup []Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup
}

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups) GetEntityData() *types.CommonEntityData {
    slaveGroups.EntityData.YFilter = slaveGroups.YFilter
    slaveGroups.EntityData.YangName = "slave-groups"
    slaveGroups.EntityData.BundleName = "cisco_ios_xr"
    slaveGroups.EntityData.ParentYangName = "ipv6"
    slaveGroups.EntityData.SegmentPath = "slave-groups"
    slaveGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slaveGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slaveGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slaveGroups.EntityData.Children = make(map[string]types.YChild)
    slaveGroups.EntityData.Children["slave-group"] = types.YChild{"SlaveGroup", nil}
    for i := range slaveGroups.SlaveGroup {
        slaveGroups.EntityData.Children[types.GetSegmentPath(&slaveGroups.SlaveGroup[i])] = types.YChild{"SlaveGroup", &slaveGroups.SlaveGroup[i]}
    }
    slaveGroups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(slaveGroups.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup
// The HSRP slave group being configured
type Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. HSRP group number. The type is interface{} with
    // range: 0..4095.
    SlaveGroupNumber interface{}

    // HSRP Group name for this slave to follow. The type is string.
    Follow interface{}

    // HSRP MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    VirtualMacAddress interface{}

    // The HSRP IPv6 virtual linklocal address.
    LinkLocalIpv6Address Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_LinkLocalIpv6Address

    // The table of HSRP virtual global IPv6 addresses.
    GlobalIpv6Addresses Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses
}

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup) GetEntityData() *types.CommonEntityData {
    slaveGroup.EntityData.YFilter = slaveGroup.YFilter
    slaveGroup.EntityData.YangName = "slave-group"
    slaveGroup.EntityData.BundleName = "cisco_ios_xr"
    slaveGroup.EntityData.ParentYangName = "slave-groups"
    slaveGroup.EntityData.SegmentPath = "slave-group" + "[slave-group-number='" + fmt.Sprintf("%v", slaveGroup.SlaveGroupNumber) + "']"
    slaveGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slaveGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slaveGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slaveGroup.EntityData.Children = make(map[string]types.YChild)
    slaveGroup.EntityData.Children["link-local-ipv6-address"] = types.YChild{"LinkLocalIpv6Address", &slaveGroup.LinkLocalIpv6Address}
    slaveGroup.EntityData.Children["global-ipv6-addresses"] = types.YChild{"GlobalIpv6Addresses", &slaveGroup.GlobalIpv6Addresses}
    slaveGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    slaveGroup.EntityData.Leafs["slave-group-number"] = types.YLeaf{"SlaveGroupNumber", slaveGroup.SlaveGroupNumber}
    slaveGroup.EntityData.Leafs["follow"] = types.YLeaf{"Follow", slaveGroup.Follow}
    slaveGroup.EntityData.Leafs["virtual-mac-address"] = types.YLeaf{"VirtualMacAddress", slaveGroup.VirtualMacAddress}
    return &(slaveGroup.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_LinkLocalIpv6Address
// The HSRP IPv6 virtual linklocal address
type Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_LinkLocalIpv6Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // HSRP IPv6 virtual linklocal address. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Linklocal Configuration Type. The type is HsrpLinklocal. The default value
    // is manual.
    AutoConfigure interface{}
}

func (linkLocalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_LinkLocalIpv6Address) GetEntityData() *types.CommonEntityData {
    linkLocalIpv6Address.EntityData.YFilter = linkLocalIpv6Address.YFilter
    linkLocalIpv6Address.EntityData.YangName = "link-local-ipv6-address"
    linkLocalIpv6Address.EntityData.BundleName = "cisco_ios_xr"
    linkLocalIpv6Address.EntityData.ParentYangName = "slave-group"
    linkLocalIpv6Address.EntityData.SegmentPath = "link-local-ipv6-address"
    linkLocalIpv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    linkLocalIpv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    linkLocalIpv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    linkLocalIpv6Address.EntityData.Children = make(map[string]types.YChild)
    linkLocalIpv6Address.EntityData.Leafs = make(map[string]types.YLeaf)
    linkLocalIpv6Address.EntityData.Leafs["address"] = types.YLeaf{"Address", linkLocalIpv6Address.Address}
    linkLocalIpv6Address.EntityData.Leafs["auto-configure"] = types.YLeaf{"AutoConfigure", linkLocalIpv6Address.AutoConfigure}
    return &(linkLocalIpv6Address.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses
// The table of HSRP virtual global IPv6
// addresses
type Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A HSRP virtual global IPv6 IP address. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses_GlobalIpv6Address.
    GlobalIpv6Address []Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses_GlobalIpv6Address
}

func (globalIpv6Addresses *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses) GetEntityData() *types.CommonEntityData {
    globalIpv6Addresses.EntityData.YFilter = globalIpv6Addresses.YFilter
    globalIpv6Addresses.EntityData.YangName = "global-ipv6-addresses"
    globalIpv6Addresses.EntityData.BundleName = "cisco_ios_xr"
    globalIpv6Addresses.EntityData.ParentYangName = "slave-group"
    globalIpv6Addresses.EntityData.SegmentPath = "global-ipv6-addresses"
    globalIpv6Addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalIpv6Addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalIpv6Addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalIpv6Addresses.EntityData.Children = make(map[string]types.YChild)
    globalIpv6Addresses.EntityData.Children["global-ipv6-address"] = types.YChild{"GlobalIpv6Address", nil}
    for i := range globalIpv6Addresses.GlobalIpv6Address {
        globalIpv6Addresses.EntityData.Children[types.GetSegmentPath(&globalIpv6Addresses.GlobalIpv6Address[i])] = types.YChild{"GlobalIpv6Address", &globalIpv6Addresses.GlobalIpv6Address[i]}
    }
    globalIpv6Addresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(globalIpv6Addresses.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses_GlobalIpv6Address
// A HSRP virtual global IPv6 IP address
type Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses_GlobalIpv6Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. HSRP virtual global IPv6 address. The type is
    // string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Address interface{}
}

func (globalIpv6Address *Hsrp_Interfaces_Interface_Ipv6_SlaveGroups_SlaveGroup_GlobalIpv6Addresses_GlobalIpv6Address) GetEntityData() *types.CommonEntityData {
    globalIpv6Address.EntityData.YFilter = globalIpv6Address.YFilter
    globalIpv6Address.EntityData.YangName = "global-ipv6-address"
    globalIpv6Address.EntityData.BundleName = "cisco_ios_xr"
    globalIpv6Address.EntityData.ParentYangName = "global-ipv6-addresses"
    globalIpv6Address.EntityData.SegmentPath = "global-ipv6-address" + "[address='" + fmt.Sprintf("%v", globalIpv6Address.Address) + "']"
    globalIpv6Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    globalIpv6Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    globalIpv6Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    globalIpv6Address.EntityData.Children = make(map[string]types.YChild)
    globalIpv6Address.EntityData.Leafs = make(map[string]types.YLeaf)
    globalIpv6Address.EntityData.Leafs["address"] = types.YLeaf{"Address", globalIpv6Address.Address}
    return &(globalIpv6Address.EntityData)
}

// Hsrp_Interfaces_Interface_Bfd
// BFD configuration
type Hsrp_Interfaces_Interface_Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detection multiplier for BFD sessions created by hsrp. The type is
    // interface{} with range: 2..50.
    DetectionMultiplier interface{}

    // Hello interval for BFD sessions created by hsrp. The type is interface{}
    // with range: 3..30000. Units are millisecond.
    Interval interface{}
}

func (bfd *Hsrp_Interfaces_Interface_Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "interface"
    bfd.EntityData.SegmentPath = "bfd"
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = make(map[string]types.YChild)
    bfd.EntityData.Leafs = make(map[string]types.YLeaf)
    bfd.EntityData.Leafs["detection-multiplier"] = types.YLeaf{"DetectionMultiplier", bfd.DetectionMultiplier}
    bfd.EntityData.Leafs["interval"] = types.YLeaf{"Interval", bfd.Interval}
    return &(bfd.EntityData)
}

// Hsrp_Interfaces_Interface_Delay
// Minimum and Reload Delay
// This type is a presence type.
type Hsrp_Interfaces_Interface_Delay struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum delay in seconds. The type is interface{} with range: 0..10000.
    // This attribute is mandatory. Units are second.
    MinimumDelay interface{}

    // Reload delay in seconds. The type is interface{} with range: 0..10000. This
    // attribute is mandatory. Units are second.
    ReloadDelay interface{}
}

func (delay *Hsrp_Interfaces_Interface_Delay) GetEntityData() *types.CommonEntityData {
    delay.EntityData.YFilter = delay.YFilter
    delay.EntityData.YangName = "delay"
    delay.EntityData.BundleName = "cisco_ios_xr"
    delay.EntityData.ParentYangName = "interface"
    delay.EntityData.SegmentPath = "delay"
    delay.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    delay.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    delay.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    delay.EntityData.Children = make(map[string]types.YChild)
    delay.EntityData.Leafs = make(map[string]types.YLeaf)
    delay.EntityData.Leafs["minimum-delay"] = types.YLeaf{"MinimumDelay", delay.MinimumDelay}
    delay.EntityData.Leafs["reload-delay"] = types.YLeaf{"ReloadDelay", delay.ReloadDelay}
    return &(delay.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4
// IPv4 HSRP configuration
type Hsrp_Interfaces_Interface_Ipv4 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The HSRP slave group configuration table.
    SlaveGroups Hsrp_Interfaces_Interface_Ipv4_SlaveGroups

    // Version 1 HSRP configuration.
    Version1 Hsrp_Interfaces_Interface_Ipv4_Version1

    // Version 2 HSRP configuration.
    Version2 Hsrp_Interfaces_Interface_Ipv4_Version2
}

func (ipv4 *Hsrp_Interfaces_Interface_Ipv4) GetEntityData() *types.CommonEntityData {
    ipv4.EntityData.YFilter = ipv4.YFilter
    ipv4.EntityData.YangName = "ipv4"
    ipv4.EntityData.BundleName = "cisco_ios_xr"
    ipv4.EntityData.ParentYangName = "interface"
    ipv4.EntityData.SegmentPath = "ipv4"
    ipv4.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipv4.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipv4.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipv4.EntityData.Children = make(map[string]types.YChild)
    ipv4.EntityData.Children["slave-groups"] = types.YChild{"SlaveGroups", &ipv4.SlaveGroups}
    ipv4.EntityData.Children["version1"] = types.YChild{"Version1", &ipv4.Version1}
    ipv4.EntityData.Children["version2"] = types.YChild{"Version2", &ipv4.Version2}
    ipv4.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipv4.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_SlaveGroups
// The HSRP slave group configuration table
type Hsrp_Interfaces_Interface_Ipv4_SlaveGroups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The HSRP slave group being configured. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup.
    SlaveGroup []Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup
}

func (slaveGroups *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups) GetEntityData() *types.CommonEntityData {
    slaveGroups.EntityData.YFilter = slaveGroups.YFilter
    slaveGroups.EntityData.YangName = "slave-groups"
    slaveGroups.EntityData.BundleName = "cisco_ios_xr"
    slaveGroups.EntityData.ParentYangName = "ipv4"
    slaveGroups.EntityData.SegmentPath = "slave-groups"
    slaveGroups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slaveGroups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slaveGroups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slaveGroups.EntityData.Children = make(map[string]types.YChild)
    slaveGroups.EntityData.Children["slave-group"] = types.YChild{"SlaveGroup", nil}
    for i := range slaveGroups.SlaveGroup {
        slaveGroups.EntityData.Children[types.GetSegmentPath(&slaveGroups.SlaveGroup[i])] = types.YChild{"SlaveGroup", &slaveGroups.SlaveGroup[i]}
    }
    slaveGroups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(slaveGroups.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup
// The HSRP slave group being configured
type Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. HSRP group number. The type is interface{} with
    // range: 0..4095.
    SlaveGroupNumber interface{}

    // HSRP Group name for this slave to follow. The type is string.
    Follow interface{}

    // HSRP MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    VirtualMacAddress interface{}

    // Primary HSRP IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    PrimaryIpv4Address interface{}

    // Secondary HSRP IP address Table.
    SecondaryIpv4Addresses Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses
}

func (slaveGroup *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup) GetEntityData() *types.CommonEntityData {
    slaveGroup.EntityData.YFilter = slaveGroup.YFilter
    slaveGroup.EntityData.YangName = "slave-group"
    slaveGroup.EntityData.BundleName = "cisco_ios_xr"
    slaveGroup.EntityData.ParentYangName = "slave-groups"
    slaveGroup.EntityData.SegmentPath = "slave-group" + "[slave-group-number='" + fmt.Sprintf("%v", slaveGroup.SlaveGroupNumber) + "']"
    slaveGroup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    slaveGroup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    slaveGroup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    slaveGroup.EntityData.Children = make(map[string]types.YChild)
    slaveGroup.EntityData.Children["secondary-ipv4-addresses"] = types.YChild{"SecondaryIpv4Addresses", &slaveGroup.SecondaryIpv4Addresses}
    slaveGroup.EntityData.Leafs = make(map[string]types.YLeaf)
    slaveGroup.EntityData.Leafs["slave-group-number"] = types.YLeaf{"SlaveGroupNumber", slaveGroup.SlaveGroupNumber}
    slaveGroup.EntityData.Leafs["follow"] = types.YLeaf{"Follow", slaveGroup.Follow}
    slaveGroup.EntityData.Leafs["virtual-mac-address"] = types.YLeaf{"VirtualMacAddress", slaveGroup.VirtualMacAddress}
    slaveGroup.EntityData.Leafs["primary-ipv4-address"] = types.YLeaf{"PrimaryIpv4Address", slaveGroup.PrimaryIpv4Address}
    return &(slaveGroup.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses
// Secondary HSRP IP address Table
type Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Secondary HSRP IP address. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses_SecondaryIpv4Address.
    SecondaryIpv4Address []Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses_SecondaryIpv4Address
}

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses) GetEntityData() *types.CommonEntityData {
    secondaryIpv4Addresses.EntityData.YFilter = secondaryIpv4Addresses.YFilter
    secondaryIpv4Addresses.EntityData.YangName = "secondary-ipv4-addresses"
    secondaryIpv4Addresses.EntityData.BundleName = "cisco_ios_xr"
    secondaryIpv4Addresses.EntityData.ParentYangName = "slave-group"
    secondaryIpv4Addresses.EntityData.SegmentPath = "secondary-ipv4-addresses"
    secondaryIpv4Addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secondaryIpv4Addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secondaryIpv4Addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secondaryIpv4Addresses.EntityData.Children = make(map[string]types.YChild)
    secondaryIpv4Addresses.EntityData.Children["secondary-ipv4-address"] = types.YChild{"SecondaryIpv4Address", nil}
    for i := range secondaryIpv4Addresses.SecondaryIpv4Address {
        secondaryIpv4Addresses.EntityData.Children[types.GetSegmentPath(&secondaryIpv4Addresses.SecondaryIpv4Address[i])] = types.YChild{"SecondaryIpv4Address", &secondaryIpv4Addresses.SecondaryIpv4Address[i]}
    }
    secondaryIpv4Addresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(secondaryIpv4Addresses.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses_SecondaryIpv4Address
// Secondary HSRP IP address
type Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses_SecondaryIpv4Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. HSRP IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}
}

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_SlaveGroups_SlaveGroup_SecondaryIpv4Addresses_SecondaryIpv4Address) GetEntityData() *types.CommonEntityData {
    secondaryIpv4Address.EntityData.YFilter = secondaryIpv4Address.YFilter
    secondaryIpv4Address.EntityData.YangName = "secondary-ipv4-address"
    secondaryIpv4Address.EntityData.BundleName = "cisco_ios_xr"
    secondaryIpv4Address.EntityData.ParentYangName = "secondary-ipv4-addresses"
    secondaryIpv4Address.EntityData.SegmentPath = "secondary-ipv4-address" + "[address='" + fmt.Sprintf("%v", secondaryIpv4Address.Address) + "']"
    secondaryIpv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secondaryIpv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secondaryIpv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secondaryIpv4Address.EntityData.Children = make(map[string]types.YChild)
    secondaryIpv4Address.EntityData.Leafs = make(map[string]types.YLeaf)
    secondaryIpv4Address.EntityData.Leafs["address"] = types.YLeaf{"Address", secondaryIpv4Address.Address}
    return &(secondaryIpv4Address.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_Version1
// Version 1 HSRP configuration
type Hsrp_Interfaces_Interface_Ipv4_Version1 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The HSRP group configuration table.
    Groups Hsrp_Interfaces_Interface_Ipv4_Version1_Groups
}

func (version1 *Hsrp_Interfaces_Interface_Ipv4_Version1) GetEntityData() *types.CommonEntityData {
    version1.EntityData.YFilter = version1.YFilter
    version1.EntityData.YangName = "version1"
    version1.EntityData.BundleName = "cisco_ios_xr"
    version1.EntityData.ParentYangName = "ipv4"
    version1.EntityData.SegmentPath = "version1"
    version1.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    version1.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    version1.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    version1.EntityData.Children = make(map[string]types.YChild)
    version1.EntityData.Children["groups"] = types.YChild{"Groups", &version1.Groups}
    version1.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(version1.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_Version1_Groups
// The HSRP group configuration table
type Hsrp_Interfaces_Interface_Ipv4_Version1_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The HSRP group being configured. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group.
    Group []Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group
}

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "cisco_ios_xr"
    groups.EntityData.ParentYangName = "version1"
    groups.EntityData.SegmentPath = "groups"
    groups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groups.EntityData.Children = make(map[string]types.YChild)
    groups.EntityData.Children["group"] = types.YChild{"Group", nil}
    for i := range groups.Group {
        groups.EntityData.Children[types.GetSegmentPath(&groups.Group[i])] = types.YChild{"Group", &groups.Group[i]}
    }
    groups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(groups.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group
// The HSRP group being configured
type Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. HSRP group number. The type is interface{} with
    // range: 0..255.
    GroupNumber interface{}

    // Authentication string. The type is string with length: 1..8. The default
    // value is cisco.
    Authentication interface{}

    // HSRP Session name (for MGO). The type is string with length: 1..16.
    SessionName interface{}

    // Priority value. The type is interface{} with range: 0..255. The default
    // value is 100.
    Priority interface{}

    // Force active if higher priority. The type is interface{} with range:
    // -2147483648..2147483647. The default value is 0.
    Preempt interface{}

    // HSRP MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    VirtualMacAddress interface{}

    // The HSRP tracked interface configuration table.
    TrackedInterfaces Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces

    // Enable use of Bidirectional Forwarding Detection.
    Bfd Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Bfd

    // The HSRP tracked interface configuration table.
    TrackedObjects Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects

    // Hello and hold timers.
    Timers Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Timers

    // Primary HSRP IP address.
    PrimaryIpv4Address Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_PrimaryIpv4Address

    // Secondary HSRP IP address Table.
    SecondaryIpv4Addresses Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses
}

func (group *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + "[group-number='" + fmt.Sprintf("%v", group.GroupNumber) + "']"
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = make(map[string]types.YChild)
    group.EntityData.Children["tracked-interfaces"] = types.YChild{"TrackedInterfaces", &group.TrackedInterfaces}
    group.EntityData.Children["bfd"] = types.YChild{"Bfd", &group.Bfd}
    group.EntityData.Children["tracked-objects"] = types.YChild{"TrackedObjects", &group.TrackedObjects}
    group.EntityData.Children["timers"] = types.YChild{"Timers", &group.Timers}
    group.EntityData.Children["primary-ipv4-address"] = types.YChild{"PrimaryIpv4Address", &group.PrimaryIpv4Address}
    group.EntityData.Children["secondary-ipv4-addresses"] = types.YChild{"SecondaryIpv4Addresses", &group.SecondaryIpv4Addresses}
    group.EntityData.Leafs = make(map[string]types.YLeaf)
    group.EntityData.Leafs["group-number"] = types.YLeaf{"GroupNumber", group.GroupNumber}
    group.EntityData.Leafs["authentication"] = types.YLeaf{"Authentication", group.Authentication}
    group.EntityData.Leafs["session-name"] = types.YLeaf{"SessionName", group.SessionName}
    group.EntityData.Leafs["priority"] = types.YLeaf{"Priority", group.Priority}
    group.EntityData.Leafs["preempt"] = types.YLeaf{"Preempt", group.Preempt}
    group.EntityData.Leafs["virtual-mac-address"] = types.YLeaf{"VirtualMacAddress", group.VirtualMacAddress}
    return &(group.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces
// The HSRP tracked interface configuration
// table
type Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface being tracked. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces_TrackedInterface.
    TrackedInterface []Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces_TrackedInterface
}

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces) GetEntityData() *types.CommonEntityData {
    trackedInterfaces.EntityData.YFilter = trackedInterfaces.YFilter
    trackedInterfaces.EntityData.YangName = "tracked-interfaces"
    trackedInterfaces.EntityData.BundleName = "cisco_ios_xr"
    trackedInterfaces.EntityData.ParentYangName = "group"
    trackedInterfaces.EntityData.SegmentPath = "tracked-interfaces"
    trackedInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackedInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackedInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackedInterfaces.EntityData.Children = make(map[string]types.YChild)
    trackedInterfaces.EntityData.Children["tracked-interface"] = types.YChild{"TrackedInterface", nil}
    for i := range trackedInterfaces.TrackedInterface {
        trackedInterfaces.EntityData.Children[types.GetSegmentPath(&trackedInterfaces.TrackedInterface[i])] = types.YChild{"TrackedInterface", &trackedInterfaces.TrackedInterface[i]}
    }
    trackedInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trackedInterfaces.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces_TrackedInterface
// Interface being tracked
type Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces_TrackedInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface being tracked. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Priority decrement. The type is interface{} with range: 1..255. This
    // attribute is mandatory.
    PriorityDecrement interface{}
}

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedInterfaces_TrackedInterface) GetEntityData() *types.CommonEntityData {
    trackedInterface.EntityData.YFilter = trackedInterface.YFilter
    trackedInterface.EntityData.YangName = "tracked-interface"
    trackedInterface.EntityData.BundleName = "cisco_ios_xr"
    trackedInterface.EntityData.ParentYangName = "tracked-interfaces"
    trackedInterface.EntityData.SegmentPath = "tracked-interface" + "[interface-name='" + fmt.Sprintf("%v", trackedInterface.InterfaceName) + "']"
    trackedInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackedInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackedInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackedInterface.EntityData.Children = make(map[string]types.YChild)
    trackedInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    trackedInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", trackedInterface.InterfaceName}
    trackedInterface.EntityData.Leafs["priority-decrement"] = types.YLeaf{"PriorityDecrement", trackedInterface.PriorityDecrement}
    return &(trackedInterface.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Bfd
// Enable use of Bidirectional Forwarding
// Detection
type Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable BFD for this remote IP. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Interface name to run BFD. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}
}

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "group"
    bfd.EntityData.SegmentPath = "bfd"
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = make(map[string]types.YChild)
    bfd.EntityData.Leafs = make(map[string]types.YLeaf)
    bfd.EntityData.Leafs["address"] = types.YLeaf{"Address", bfd.Address}
    bfd.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", bfd.InterfaceName}
    return &(bfd.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects
// The HSRP tracked interface configuration
// table
type Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Object being tracked. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects_TrackedObject.
    TrackedObject []Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects_TrackedObject
}

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects) GetEntityData() *types.CommonEntityData {
    trackedObjects.EntityData.YFilter = trackedObjects.YFilter
    trackedObjects.EntityData.YangName = "tracked-objects"
    trackedObjects.EntityData.BundleName = "cisco_ios_xr"
    trackedObjects.EntityData.ParentYangName = "group"
    trackedObjects.EntityData.SegmentPath = "tracked-objects"
    trackedObjects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackedObjects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackedObjects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackedObjects.EntityData.Children = make(map[string]types.YChild)
    trackedObjects.EntityData.Children["tracked-object"] = types.YChild{"TrackedObject", nil}
    for i := range trackedObjects.TrackedObject {
        trackedObjects.EntityData.Children[types.GetSegmentPath(&trackedObjects.TrackedObject[i])] = types.YChild{"TrackedObject", &trackedObjects.TrackedObject[i]}
    }
    trackedObjects.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trackedObjects.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects_TrackedObject
// Object being tracked
type Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects_TrackedObject struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface being tracked. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ObjectName interface{}

    // Priority decrement. The type is interface{} with range: 1..255. This
    // attribute is mandatory.
    PriorityDecrement interface{}
}

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_TrackedObjects_TrackedObject) GetEntityData() *types.CommonEntityData {
    trackedObject.EntityData.YFilter = trackedObject.YFilter
    trackedObject.EntityData.YangName = "tracked-object"
    trackedObject.EntityData.BundleName = "cisco_ios_xr"
    trackedObject.EntityData.ParentYangName = "tracked-objects"
    trackedObject.EntityData.SegmentPath = "tracked-object" + "[object-name='" + fmt.Sprintf("%v", trackedObject.ObjectName) + "']"
    trackedObject.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackedObject.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackedObject.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackedObject.EntityData.Children = make(map[string]types.YChild)
    trackedObject.EntityData.Leafs = make(map[string]types.YLeaf)
    trackedObject.EntityData.Leafs["object-name"] = types.YLeaf{"ObjectName", trackedObject.ObjectName}
    trackedObject.EntityData.Leafs["priority-decrement"] = types.YLeaf{"PriorityDecrement", trackedObject.PriorityDecrement}
    return &(trackedObject.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Timers
// Hello and hold timers
type Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Timers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TRUE - Hello time configured in milliseconds, FALSE - Hello time configured
    // in seconds. The type is bool. The default value is false.
    HelloMsecFlag interface{}

    // Hello time in msecs. The type is interface{} with range: 100..3000. Units
    // are millisecond.
    HelloMsec interface{}

    // Hello time in seconds. The type is interface{} with range: 1..255. Units
    // are second. The default value is 3.
    HelloSec interface{}

    // TRUE - Hold time configured in milliseconds, FALSE - Hold time configured
    // in seconds. The type is bool. The default value is false.
    HoldMsecFlag interface{}

    // Hold time in msecs. The type is interface{} with range: 100..3000. Units
    // are millisecond.
    HoldMsec interface{}

    // Hold time in seconds. The type is interface{} with range: 1..255. Units are
    // second. The default value is 10.
    HoldSec interface{}
}

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_Timers) GetEntityData() *types.CommonEntityData {
    timers.EntityData.YFilter = timers.YFilter
    timers.EntityData.YangName = "timers"
    timers.EntityData.BundleName = "cisco_ios_xr"
    timers.EntityData.ParentYangName = "group"
    timers.EntityData.SegmentPath = "timers"
    timers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timers.EntityData.Children = make(map[string]types.YChild)
    timers.EntityData.Leafs = make(map[string]types.YLeaf)
    timers.EntityData.Leafs["hello-msec-flag"] = types.YLeaf{"HelloMsecFlag", timers.HelloMsecFlag}
    timers.EntityData.Leafs["hello-msec"] = types.YLeaf{"HelloMsec", timers.HelloMsec}
    timers.EntityData.Leafs["hello-sec"] = types.YLeaf{"HelloSec", timers.HelloSec}
    timers.EntityData.Leafs["hold-msec-flag"] = types.YLeaf{"HoldMsecFlag", timers.HoldMsecFlag}
    timers.EntityData.Leafs["hold-msec"] = types.YLeaf{"HoldMsec", timers.HoldMsec}
    timers.EntityData.Leafs["hold-sec"] = types.YLeaf{"HoldSec", timers.HoldSec}
    return &(timers.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_PrimaryIpv4Address
// Primary HSRP IP address
type Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_PrimaryIpv4Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TRUE if the HSRP protocol is to learn the virtual IP address it is to use.
    // The type is bool.
    VirtualIpLearn interface{}

    // HSRP IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}
}

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_PrimaryIpv4Address) GetEntityData() *types.CommonEntityData {
    primaryIpv4Address.EntityData.YFilter = primaryIpv4Address.YFilter
    primaryIpv4Address.EntityData.YangName = "primary-ipv4-address"
    primaryIpv4Address.EntityData.BundleName = "cisco_ios_xr"
    primaryIpv4Address.EntityData.ParentYangName = "group"
    primaryIpv4Address.EntityData.SegmentPath = "primary-ipv4-address"
    primaryIpv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    primaryIpv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    primaryIpv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    primaryIpv4Address.EntityData.Children = make(map[string]types.YChild)
    primaryIpv4Address.EntityData.Leafs = make(map[string]types.YLeaf)
    primaryIpv4Address.EntityData.Leafs["virtual-ip-learn"] = types.YLeaf{"VirtualIpLearn", primaryIpv4Address.VirtualIpLearn}
    primaryIpv4Address.EntityData.Leafs["address"] = types.YLeaf{"Address", primaryIpv4Address.Address}
    return &(primaryIpv4Address.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses
// Secondary HSRP IP address Table
type Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Secondary HSRP IP address. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address.
    SecondaryIpv4Address []Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address
}

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses) GetEntityData() *types.CommonEntityData {
    secondaryIpv4Addresses.EntityData.YFilter = secondaryIpv4Addresses.YFilter
    secondaryIpv4Addresses.EntityData.YangName = "secondary-ipv4-addresses"
    secondaryIpv4Addresses.EntityData.BundleName = "cisco_ios_xr"
    secondaryIpv4Addresses.EntityData.ParentYangName = "group"
    secondaryIpv4Addresses.EntityData.SegmentPath = "secondary-ipv4-addresses"
    secondaryIpv4Addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secondaryIpv4Addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secondaryIpv4Addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secondaryIpv4Addresses.EntityData.Children = make(map[string]types.YChild)
    secondaryIpv4Addresses.EntityData.Children["secondary-ipv4-address"] = types.YChild{"SecondaryIpv4Address", nil}
    for i := range secondaryIpv4Addresses.SecondaryIpv4Address {
        secondaryIpv4Addresses.EntityData.Children[types.GetSegmentPath(&secondaryIpv4Addresses.SecondaryIpv4Address[i])] = types.YChild{"SecondaryIpv4Address", &secondaryIpv4Addresses.SecondaryIpv4Address[i]}
    }
    secondaryIpv4Addresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(secondaryIpv4Addresses.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address
// Secondary HSRP IP address
type Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. HSRP IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}
}

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version1_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetEntityData() *types.CommonEntityData {
    secondaryIpv4Address.EntityData.YFilter = secondaryIpv4Address.YFilter
    secondaryIpv4Address.EntityData.YangName = "secondary-ipv4-address"
    secondaryIpv4Address.EntityData.BundleName = "cisco_ios_xr"
    secondaryIpv4Address.EntityData.ParentYangName = "secondary-ipv4-addresses"
    secondaryIpv4Address.EntityData.SegmentPath = "secondary-ipv4-address" + "[address='" + fmt.Sprintf("%v", secondaryIpv4Address.Address) + "']"
    secondaryIpv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secondaryIpv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secondaryIpv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secondaryIpv4Address.EntityData.Children = make(map[string]types.YChild)
    secondaryIpv4Address.EntityData.Leafs = make(map[string]types.YLeaf)
    secondaryIpv4Address.EntityData.Leafs["address"] = types.YLeaf{"Address", secondaryIpv4Address.Address}
    return &(secondaryIpv4Address.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_Version2
// Version 2 HSRP configuration
type Hsrp_Interfaces_Interface_Ipv4_Version2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The HSRP group configuration table.
    Groups Hsrp_Interfaces_Interface_Ipv4_Version2_Groups
}

func (version2 *Hsrp_Interfaces_Interface_Ipv4_Version2) GetEntityData() *types.CommonEntityData {
    version2.EntityData.YFilter = version2.YFilter
    version2.EntityData.YangName = "version2"
    version2.EntityData.BundleName = "cisco_ios_xr"
    version2.EntityData.ParentYangName = "ipv4"
    version2.EntityData.SegmentPath = "version2"
    version2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    version2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    version2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    version2.EntityData.Children = make(map[string]types.YChild)
    version2.EntityData.Children["groups"] = types.YChild{"Groups", &version2.Groups}
    version2.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(version2.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_Version2_Groups
// The HSRP group configuration table
type Hsrp_Interfaces_Interface_Ipv4_Version2_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The HSRP group being configured. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group.
    Group []Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group
}

func (groups *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "cisco_ios_xr"
    groups.EntityData.ParentYangName = "version2"
    groups.EntityData.SegmentPath = "groups"
    groups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groups.EntityData.Children = make(map[string]types.YChild)
    groups.EntityData.Children["group"] = types.YChild{"Group", nil}
    for i := range groups.Group {
        groups.EntityData.Children[types.GetSegmentPath(&groups.Group[i])] = types.YChild{"Group", &groups.Group[i]}
    }
    groups.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(groups.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group
// The HSRP group being configured
type Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. HSRP group number. The type is interface{} with
    // range: 0..4095.
    GroupNumber interface{}

    // Force active if higher priority. The type is interface{} with range:
    // -2147483648..2147483647. The default value is 0.
    Preempt interface{}

    // Priority value. The type is interface{} with range: 0..255. The default
    // value is 100.
    Priority interface{}

    // HSRP MAC address. The type is string with pattern:
    // b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    VirtualMacAddress interface{}

    // HSRP Session name (for MGO). The type is string with length: 1..16.
    SessionName interface{}

    // Secondary HSRP IP address Table.
    SecondaryIpv4Addresses Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses

    // Enable use of Bidirectional Forwarding Detection.
    Bfd Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Bfd

    // Primary HSRP IP address.
    PrimaryIpv4Address Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_PrimaryIpv4Address

    // The HSRP tracked interface configuration table.
    TrackedObjects Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects

    // The HSRP tracked interface configuration table.
    TrackedInterfaces Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces

    // Hello and hold timers.
    Timers Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Timers
}

func (group *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group) GetEntityData() *types.CommonEntityData {
    group.EntityData.YFilter = group.YFilter
    group.EntityData.YangName = "group"
    group.EntityData.BundleName = "cisco_ios_xr"
    group.EntityData.ParentYangName = "groups"
    group.EntityData.SegmentPath = "group" + "[group-number='" + fmt.Sprintf("%v", group.GroupNumber) + "']"
    group.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    group.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    group.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    group.EntityData.Children = make(map[string]types.YChild)
    group.EntityData.Children["secondary-ipv4-addresses"] = types.YChild{"SecondaryIpv4Addresses", &group.SecondaryIpv4Addresses}
    group.EntityData.Children["bfd"] = types.YChild{"Bfd", &group.Bfd}
    group.EntityData.Children["primary-ipv4-address"] = types.YChild{"PrimaryIpv4Address", &group.PrimaryIpv4Address}
    group.EntityData.Children["tracked-objects"] = types.YChild{"TrackedObjects", &group.TrackedObjects}
    group.EntityData.Children["tracked-interfaces"] = types.YChild{"TrackedInterfaces", &group.TrackedInterfaces}
    group.EntityData.Children["timers"] = types.YChild{"Timers", &group.Timers}
    group.EntityData.Leafs = make(map[string]types.YLeaf)
    group.EntityData.Leafs["group-number"] = types.YLeaf{"GroupNumber", group.GroupNumber}
    group.EntityData.Leafs["preempt"] = types.YLeaf{"Preempt", group.Preempt}
    group.EntityData.Leafs["priority"] = types.YLeaf{"Priority", group.Priority}
    group.EntityData.Leafs["virtual-mac-address"] = types.YLeaf{"VirtualMacAddress", group.VirtualMacAddress}
    group.EntityData.Leafs["session-name"] = types.YLeaf{"SessionName", group.SessionName}
    return &(group.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses
// Secondary HSRP IP address Table
type Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Secondary HSRP IP address. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address.
    SecondaryIpv4Address []Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address
}

func (secondaryIpv4Addresses *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses) GetEntityData() *types.CommonEntityData {
    secondaryIpv4Addresses.EntityData.YFilter = secondaryIpv4Addresses.YFilter
    secondaryIpv4Addresses.EntityData.YangName = "secondary-ipv4-addresses"
    secondaryIpv4Addresses.EntityData.BundleName = "cisco_ios_xr"
    secondaryIpv4Addresses.EntityData.ParentYangName = "group"
    secondaryIpv4Addresses.EntityData.SegmentPath = "secondary-ipv4-addresses"
    secondaryIpv4Addresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secondaryIpv4Addresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secondaryIpv4Addresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secondaryIpv4Addresses.EntityData.Children = make(map[string]types.YChild)
    secondaryIpv4Addresses.EntityData.Children["secondary-ipv4-address"] = types.YChild{"SecondaryIpv4Address", nil}
    for i := range secondaryIpv4Addresses.SecondaryIpv4Address {
        secondaryIpv4Addresses.EntityData.Children[types.GetSegmentPath(&secondaryIpv4Addresses.SecondaryIpv4Address[i])] = types.YChild{"SecondaryIpv4Address", &secondaryIpv4Addresses.SecondaryIpv4Address[i]}
    }
    secondaryIpv4Addresses.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(secondaryIpv4Addresses.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address
// Secondary HSRP IP address
type Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. HSRP IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}
}

func (secondaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_SecondaryIpv4Addresses_SecondaryIpv4Address) GetEntityData() *types.CommonEntityData {
    secondaryIpv4Address.EntityData.YFilter = secondaryIpv4Address.YFilter
    secondaryIpv4Address.EntityData.YangName = "secondary-ipv4-address"
    secondaryIpv4Address.EntityData.BundleName = "cisco_ios_xr"
    secondaryIpv4Address.EntityData.ParentYangName = "secondary-ipv4-addresses"
    secondaryIpv4Address.EntityData.SegmentPath = "secondary-ipv4-address" + "[address='" + fmt.Sprintf("%v", secondaryIpv4Address.Address) + "']"
    secondaryIpv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secondaryIpv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secondaryIpv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secondaryIpv4Address.EntityData.Children = make(map[string]types.YChild)
    secondaryIpv4Address.EntityData.Leafs = make(map[string]types.YLeaf)
    secondaryIpv4Address.EntityData.Leafs["address"] = types.YLeaf{"Address", secondaryIpv4Address.Address}
    return &(secondaryIpv4Address.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Bfd
// Enable use of Bidirectional Forwarding
// Detection
type Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable BFD for this remote IP. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Interface name to run BFD. The type is string with pattern:
    // b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}
}

func (bfd *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "group"
    bfd.EntityData.SegmentPath = "bfd"
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = make(map[string]types.YChild)
    bfd.EntityData.Leafs = make(map[string]types.YLeaf)
    bfd.EntityData.Leafs["address"] = types.YLeaf{"Address", bfd.Address}
    bfd.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", bfd.InterfaceName}
    return &(bfd.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_PrimaryIpv4Address
// Primary HSRP IP address
type Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_PrimaryIpv4Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TRUE if the HSRP protocol is to learn the virtual IP address it is to use.
    // The type is bool.
    VirtualIpLearn interface{}

    // HSRP IP address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}
}

func (primaryIpv4Address *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_PrimaryIpv4Address) GetEntityData() *types.CommonEntityData {
    primaryIpv4Address.EntityData.YFilter = primaryIpv4Address.YFilter
    primaryIpv4Address.EntityData.YangName = "primary-ipv4-address"
    primaryIpv4Address.EntityData.BundleName = "cisco_ios_xr"
    primaryIpv4Address.EntityData.ParentYangName = "group"
    primaryIpv4Address.EntityData.SegmentPath = "primary-ipv4-address"
    primaryIpv4Address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    primaryIpv4Address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    primaryIpv4Address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    primaryIpv4Address.EntityData.Children = make(map[string]types.YChild)
    primaryIpv4Address.EntityData.Leafs = make(map[string]types.YLeaf)
    primaryIpv4Address.EntityData.Leafs["virtual-ip-learn"] = types.YLeaf{"VirtualIpLearn", primaryIpv4Address.VirtualIpLearn}
    primaryIpv4Address.EntityData.Leafs["address"] = types.YLeaf{"Address", primaryIpv4Address.Address}
    return &(primaryIpv4Address.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects
// The HSRP tracked interface configuration
// table
type Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Object being tracked. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects_TrackedObject.
    TrackedObject []Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects_TrackedObject
}

func (trackedObjects *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects) GetEntityData() *types.CommonEntityData {
    trackedObjects.EntityData.YFilter = trackedObjects.YFilter
    trackedObjects.EntityData.YangName = "tracked-objects"
    trackedObjects.EntityData.BundleName = "cisco_ios_xr"
    trackedObjects.EntityData.ParentYangName = "group"
    trackedObjects.EntityData.SegmentPath = "tracked-objects"
    trackedObjects.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackedObjects.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackedObjects.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackedObjects.EntityData.Children = make(map[string]types.YChild)
    trackedObjects.EntityData.Children["tracked-object"] = types.YChild{"TrackedObject", nil}
    for i := range trackedObjects.TrackedObject {
        trackedObjects.EntityData.Children[types.GetSegmentPath(&trackedObjects.TrackedObject[i])] = types.YChild{"TrackedObject", &trackedObjects.TrackedObject[i]}
    }
    trackedObjects.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trackedObjects.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects_TrackedObject
// Object being tracked
type Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects_TrackedObject struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface being tracked. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ObjectName interface{}

    // Priority decrement. The type is interface{} with range: 1..255. This
    // attribute is mandatory.
    PriorityDecrement interface{}
}

func (trackedObject *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedObjects_TrackedObject) GetEntityData() *types.CommonEntityData {
    trackedObject.EntityData.YFilter = trackedObject.YFilter
    trackedObject.EntityData.YangName = "tracked-object"
    trackedObject.EntityData.BundleName = "cisco_ios_xr"
    trackedObject.EntityData.ParentYangName = "tracked-objects"
    trackedObject.EntityData.SegmentPath = "tracked-object" + "[object-name='" + fmt.Sprintf("%v", trackedObject.ObjectName) + "']"
    trackedObject.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackedObject.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackedObject.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackedObject.EntityData.Children = make(map[string]types.YChild)
    trackedObject.EntityData.Leafs = make(map[string]types.YLeaf)
    trackedObject.EntityData.Leafs["object-name"] = types.YLeaf{"ObjectName", trackedObject.ObjectName}
    trackedObject.EntityData.Leafs["priority-decrement"] = types.YLeaf{"PriorityDecrement", trackedObject.PriorityDecrement}
    return &(trackedObject.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces
// The HSRP tracked interface configuration
// table
type Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface being tracked. The type is slice of
    // Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces_TrackedInterface.
    TrackedInterface []Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces_TrackedInterface
}

func (trackedInterfaces *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces) GetEntityData() *types.CommonEntityData {
    trackedInterfaces.EntityData.YFilter = trackedInterfaces.YFilter
    trackedInterfaces.EntityData.YangName = "tracked-interfaces"
    trackedInterfaces.EntityData.BundleName = "cisco_ios_xr"
    trackedInterfaces.EntityData.ParentYangName = "group"
    trackedInterfaces.EntityData.SegmentPath = "tracked-interfaces"
    trackedInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackedInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackedInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackedInterfaces.EntityData.Children = make(map[string]types.YChild)
    trackedInterfaces.EntityData.Children["tracked-interface"] = types.YChild{"TrackedInterface", nil}
    for i := range trackedInterfaces.TrackedInterface {
        trackedInterfaces.EntityData.Children[types.GetSegmentPath(&trackedInterfaces.TrackedInterface[i])] = types.YChild{"TrackedInterface", &trackedInterfaces.TrackedInterface[i]}
    }
    trackedInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(trackedInterfaces.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces_TrackedInterface
// Interface being tracked
type Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces_TrackedInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Interface being tracked. The type is string with
    // pattern: b'[a-zA-Z0-9./-]+'.
    InterfaceName interface{}

    // Priority decrement. The type is interface{} with range: 1..255. This
    // attribute is mandatory.
    PriorityDecrement interface{}
}

func (trackedInterface *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_TrackedInterfaces_TrackedInterface) GetEntityData() *types.CommonEntityData {
    trackedInterface.EntityData.YFilter = trackedInterface.YFilter
    trackedInterface.EntityData.YangName = "tracked-interface"
    trackedInterface.EntityData.BundleName = "cisco_ios_xr"
    trackedInterface.EntityData.ParentYangName = "tracked-interfaces"
    trackedInterface.EntityData.SegmentPath = "tracked-interface" + "[interface-name='" + fmt.Sprintf("%v", trackedInterface.InterfaceName) + "']"
    trackedInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trackedInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trackedInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trackedInterface.EntityData.Children = make(map[string]types.YChild)
    trackedInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    trackedInterface.EntityData.Leafs["interface-name"] = types.YLeaf{"InterfaceName", trackedInterface.InterfaceName}
    trackedInterface.EntityData.Leafs["priority-decrement"] = types.YLeaf{"PriorityDecrement", trackedInterface.PriorityDecrement}
    return &(trackedInterface.EntityData)
}

// Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Timers
// Hello and hold timers
type Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Timers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TRUE - Hello time configured in milliseconds, FALSE - Hello time configured
    // in seconds. The type is bool. The default value is false.
    HelloMsecFlag interface{}

    // Hello time in msecs. The type is interface{} with range: 100..3000. Units
    // are millisecond.
    HelloMsec interface{}

    // Hello time in seconds. The type is interface{} with range: 1..255. Units
    // are second. The default value is 3.
    HelloSec interface{}

    // TRUE - Hold time configured in milliseconds, FALSE - Hold time configured
    // in seconds. The type is bool. The default value is false.
    HoldMsecFlag interface{}

    // Hold time in msecs. The type is interface{} with range: 100..3000. Units
    // are millisecond.
    HoldMsec interface{}

    // Hold time in seconds. The type is interface{} with range: 1..255. Units are
    // second. The default value is 10.
    HoldSec interface{}
}

func (timers *Hsrp_Interfaces_Interface_Ipv4_Version2_Groups_Group_Timers) GetEntityData() *types.CommonEntityData {
    timers.EntityData.YFilter = timers.YFilter
    timers.EntityData.YangName = "timers"
    timers.EntityData.BundleName = "cisco_ios_xr"
    timers.EntityData.ParentYangName = "group"
    timers.EntityData.SegmentPath = "timers"
    timers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timers.EntityData.Children = make(map[string]types.YChild)
    timers.EntityData.Leafs = make(map[string]types.YLeaf)
    timers.EntityData.Leafs["hello-msec-flag"] = types.YLeaf{"HelloMsecFlag", timers.HelloMsecFlag}
    timers.EntityData.Leafs["hello-msec"] = types.YLeaf{"HelloMsec", timers.HelloMsec}
    timers.EntityData.Leafs["hello-sec"] = types.YLeaf{"HelloSec", timers.HelloSec}
    timers.EntityData.Leafs["hold-msec-flag"] = types.YLeaf{"HoldMsecFlag", timers.HoldMsecFlag}
    timers.EntityData.Leafs["hold-msec"] = types.YLeaf{"HoldMsec", timers.HoldMsec}
    timers.EntityData.Leafs["hold-sec"] = types.YLeaf{"HoldSec", timers.HoldSec}
    return &(timers.EntityData)
}

// Hsrp_Logging
// HSRP logging options
type Hsrp_Logging struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // HSRP state change IOS messages disable. The type is interface{}.
    StateChangeDisable interface{}
}

func (logging *Hsrp_Logging) GetEntityData() *types.CommonEntityData {
    logging.EntityData.YFilter = logging.YFilter
    logging.EntityData.YangName = "logging"
    logging.EntityData.BundleName = "cisco_ios_xr"
    logging.EntityData.ParentYangName = "hsrp"
    logging.EntityData.SegmentPath = "logging"
    logging.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    logging.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    logging.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    logging.EntityData.Children = make(map[string]types.YChild)
    logging.EntityData.Leafs = make(map[string]types.YLeaf)
    logging.EntityData.Leafs["state-change-disable"] = types.YLeaf{"StateChangeDisable", logging.StateChangeDisable}
    return &(logging.EntityData)
}

