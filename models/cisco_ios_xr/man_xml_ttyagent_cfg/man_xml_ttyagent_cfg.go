// This module contains a collection of YANG definitions
// for Cisco IOS-XR man-xml-ttyagent package configuration.
// 
// This module contains definitions
// for the following management objects:
//   xr-xml: XML
//   netconf: netconf
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package man_xml_ttyagent_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package man_xml_ttyagent_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-man-xml-ttyagent-cfg xr-xml}", reflect.TypeOf(XrXml{}))
    ydk.RegisterEntity("Cisco-IOS-XR-man-xml-ttyagent-cfg:xr-xml", reflect.TypeOf(XrXml{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-man-xml-ttyagent-cfg netconf}", reflect.TypeOf(Netconf{}))
    ydk.RegisterEntity("Cisco-IOS-XR-man-xml-ttyagent-cfg:netconf", reflect.TypeOf(Netconf{}))
}

// XrXml
// XML
type XrXml struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // XML agent.
    Agent XrXml_Agent
}

func (xrXml *XrXml) GetEntityData() *types.CommonEntityData {
    xrXml.EntityData.YFilter = xrXml.YFilter
    xrXml.EntityData.YangName = "xr-xml"
    xrXml.EntityData.BundleName = "cisco_ios_xr"
    xrXml.EntityData.ParentYangName = "Cisco-IOS-XR-man-xml-ttyagent-cfg"
    xrXml.EntityData.SegmentPath = "Cisco-IOS-XR-man-xml-ttyagent-cfg:xr-xml"
    xrXml.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    xrXml.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    xrXml.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    xrXml.EntityData.Children = types.NewOrderedMap()
    xrXml.EntityData.Children.Append("agent", types.YChild{"Agent", &xrXml.Agent})
    xrXml.EntityData.Leafs = types.NewOrderedMap()

    xrXml.EntityData.YListKeys = []string {}

    return &(xrXml.EntityData)
}

// XrXml_Agent
// XML agent
type XrXml_Agent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // XML default dedicated agent.
    Default XrXml_Agent_Default

    // XML TTY agent.
    Tty XrXml_Agent_Tty

    // XML SSL agent.
    Ssl XrXml_Agent_Ssl
}

func (agent *XrXml_Agent) GetEntityData() *types.CommonEntityData {
    agent.EntityData.YFilter = agent.YFilter
    agent.EntityData.YangName = "agent"
    agent.EntityData.BundleName = "cisco_ios_xr"
    agent.EntityData.ParentYangName = "xr-xml"
    agent.EntityData.SegmentPath = "agent"
    agent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    agent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    agent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    agent.EntityData.Children = types.NewOrderedMap()
    agent.EntityData.Children.Append("default", types.YChild{"Default", &agent.Default})
    agent.EntityData.Children.Append("tty", types.YChild{"Tty", &agent.Tty})
    agent.EntityData.Children.Append("ssl", types.YChild{"Ssl", &agent.Ssl})
    agent.EntityData.Leafs = types.NewOrderedMap()

    agent.EntityData.YListKeys = []string {}

    return &(agent.EntityData)
}

// XrXml_Agent_Default
// XML default dedicated agent
type XrXml_Agent_Default struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 Transport State. The type is bool.
    Ipv6Enable interface{}

    // TRUE to disable IPV4. The type is bool.
    Ipv4Disable interface{}

    // Iterator size, in KBytes, of the XML response. Specify 0 to turn off the
    // XML response iterator. The type is interface{} with range: 0..100000. Units
    // are kilobyte. The default value is 48.
    IterationSize interface{}

    // Enable specified XML agent. The type is interface{}.
    Enable interface{}

    // Streaming size, in KBytes, of the XML response. The type is interface{}
    // with range: 1..100000. Units are kilobyte.
    StreamingSize interface{}

    // Session attributes.
    Session XrXml_Agent_Default_Session

    // XML agent throttling.
    Throttle XrXml_Agent_Default_Throttle

    // List of VRFs.
    Vrfs XrXml_Agent_Default_Vrfs
}

func (self *XrXml_Agent_Default) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "default"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "agent"
    self.EntityData.SegmentPath = "default"
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("session", types.YChild{"Session", &self.Session})
    self.EntityData.Children.Append("throttle", types.YChild{"Throttle", &self.Throttle})
    self.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &self.Vrfs})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("ipv6-enable", types.YLeaf{"Ipv6Enable", self.Ipv6Enable})
    self.EntityData.Leafs.Append("ipv4-disable", types.YLeaf{"Ipv4Disable", self.Ipv4Disable})
    self.EntityData.Leafs.Append("iteration-size", types.YLeaf{"IterationSize", self.IterationSize})
    self.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", self.Enable})
    self.EntityData.Leafs.Append("streaming-size", types.YLeaf{"StreamingSize", self.StreamingSize})

    self.EntityData.YListKeys = []string {}

    return &(self.EntityData)
}

// XrXml_Agent_Default_Session
// Session attributes
type XrXml_Agent_Default_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timeout in minutes. The type is interface{} with range: 1..1440. Units are
    // minute.
    Timeout interface{}
}

func (session *XrXml_Agent_Default_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "default"
    session.EntityData.SegmentPath = "session"
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", session.Timeout})

    session.EntityData.YListKeys = []string {}

    return &(session.EntityData)
}

// XrXml_Agent_Default_Throttle
// XML agent throttling
type XrXml_Agent_Default_Throttle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process rate in number of XML tags per second. The type is interface{} with
    // range: 1000..30000.
    ProcessRate interface{}

    // Size of memory usage, in MBytes, per session. The type is interface{} with
    // range: 100..1024. Units are megabyte. The default value is 300.
    Memory interface{}
}

func (throttle *XrXml_Agent_Default_Throttle) GetEntityData() *types.CommonEntityData {
    throttle.EntityData.YFilter = throttle.YFilter
    throttle.EntityData.YangName = "throttle"
    throttle.EntityData.BundleName = "cisco_ios_xr"
    throttle.EntityData.ParentYangName = "default"
    throttle.EntityData.SegmentPath = "throttle"
    throttle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    throttle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    throttle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    throttle.EntityData.Children = types.NewOrderedMap()
    throttle.EntityData.Leafs = types.NewOrderedMap()
    throttle.EntityData.Leafs.Append("process-rate", types.YLeaf{"ProcessRate", throttle.ProcessRate})
    throttle.EntityData.Leafs.Append("memory", types.YLeaf{"Memory", throttle.Memory})

    throttle.EntityData.YListKeys = []string {}

    return &(throttle.EntityData)
}

// XrXml_Agent_Default_Vrfs
// List of VRFs
type XrXml_Agent_Default_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A specific VRF. The type is slice of XrXml_Agent_Default_Vrfs_Vrf.
    Vrf []*XrXml_Agent_Default_Vrfs_Vrf
}

func (vrfs *XrXml_Agent_Default_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "default"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = types.NewOrderedMap()
    vrfs.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children.Append(types.GetSegmentPath(vrfs.Vrf[i]), types.YChild{"Vrf", vrfs.Vrf[i]})
    }
    vrfs.EntityData.Leafs = types.NewOrderedMap()

    vrfs.EntityData.YListKeys = []string {}

    return &(vrfs.EntityData)
}

// XrXml_Agent_Default_Vrfs_Vrf
// A specific VRF
type XrXml_Agent_Default_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with length: 1..32.
    VrfName interface{}

    // IPv6 Transport Access list for VRF. The type is string with length: 1..32.
    Ipv6AccessList interface{}

    // IPv4 Transport Access list for VRF. The type is string with length: 1..32.
    Ipv4AccessList interface{}

    // Access list for XML agent. The type is string with length: 1..32.
    AccessList interface{}

    // Shutdown default VRF. This is applicable only for VRF default. The type is
    // interface{}.
    Shutdown interface{}
}

func (vrf *XrXml_Agent_Default_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})
    vrf.EntityData.Leafs.Append("ipv6-access-list", types.YLeaf{"Ipv6AccessList", vrf.Ipv6AccessList})
    vrf.EntityData.Leafs.Append("ipv4-access-list", types.YLeaf{"Ipv4AccessList", vrf.Ipv4AccessList})
    vrf.EntityData.Leafs.Append("access-list", types.YLeaf{"AccessList", vrf.AccessList})
    vrf.EntityData.Leafs.Append("shutdown", types.YLeaf{"Shutdown", vrf.Shutdown})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// XrXml_Agent_Tty
// XML TTY agent
type XrXml_Agent_Tty struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Iterator size, in KBytes, of the XML response. Specify 0 to turn off the
    // XML response iterator. The type is interface{} with range: 0..100000. Units
    // are kilobyte. The default value is 48.
    IterationSize interface{}

    // Enable specified XML agent. The type is interface{}.
    Enable interface{}

    // Streaming size, in KBytes, of the XML response. The type is interface{}
    // with range: 1..100000. Units are kilobyte.
    StreamingSize interface{}

    // Session attributes.
    Session XrXml_Agent_Tty_Session

    // XML agent throttling.
    Throttle XrXml_Agent_Tty_Throttle
}

func (tty *XrXml_Agent_Tty) GetEntityData() *types.CommonEntityData {
    tty.EntityData.YFilter = tty.YFilter
    tty.EntityData.YangName = "tty"
    tty.EntityData.BundleName = "cisco_ios_xr"
    tty.EntityData.ParentYangName = "agent"
    tty.EntityData.SegmentPath = "tty"
    tty.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tty.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tty.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tty.EntityData.Children = types.NewOrderedMap()
    tty.EntityData.Children.Append("session", types.YChild{"Session", &tty.Session})
    tty.EntityData.Children.Append("throttle", types.YChild{"Throttle", &tty.Throttle})
    tty.EntityData.Leafs = types.NewOrderedMap()
    tty.EntityData.Leafs.Append("iteration-size", types.YLeaf{"IterationSize", tty.IterationSize})
    tty.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", tty.Enable})
    tty.EntityData.Leafs.Append("streaming-size", types.YLeaf{"StreamingSize", tty.StreamingSize})

    tty.EntityData.YListKeys = []string {}

    return &(tty.EntityData)
}

// XrXml_Agent_Tty_Session
// Session attributes
type XrXml_Agent_Tty_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timeout in minutes. The type is interface{} with range: 1..1440. Units are
    // minute.
    Timeout interface{}
}

func (session *XrXml_Agent_Tty_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "tty"
    session.EntityData.SegmentPath = "session"
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", session.Timeout})

    session.EntityData.YListKeys = []string {}

    return &(session.EntityData)
}

// XrXml_Agent_Tty_Throttle
// XML agent throttling
type XrXml_Agent_Tty_Throttle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process rate in number of XML tags per second. The type is interface{} with
    // range: 1000..30000.
    ProcessRate interface{}

    // Size of memory usage, in MBytes, per session. The type is interface{} with
    // range: 100..1024. Units are megabyte. The default value is 300.
    Memory interface{}
}

func (throttle *XrXml_Agent_Tty_Throttle) GetEntityData() *types.CommonEntityData {
    throttle.EntityData.YFilter = throttle.YFilter
    throttle.EntityData.YangName = "throttle"
    throttle.EntityData.BundleName = "cisco_ios_xr"
    throttle.EntityData.ParentYangName = "tty"
    throttle.EntityData.SegmentPath = "throttle"
    throttle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    throttle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    throttle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    throttle.EntityData.Children = types.NewOrderedMap()
    throttle.EntityData.Leafs = types.NewOrderedMap()
    throttle.EntityData.Leafs.Append("process-rate", types.YLeaf{"ProcessRate", throttle.ProcessRate})
    throttle.EntityData.Leafs.Append("memory", types.YLeaf{"Memory", throttle.Memory})

    throttle.EntityData.YListKeys = []string {}

    return &(throttle.EntityData)
}

// XrXml_Agent_Ssl
// XML SSL agent
type XrXml_Agent_Ssl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Iterator size, in KBytes, of the XML response. Specify 0 to turn off the
    // XML response iterator. The type is interface{} with range: 0..100000. Units
    // are kilobyte. The default value is 48.
    IterationSize interface{}

    // Enable specified XML agent. The type is interface{}.
    Enable interface{}

    // Streaming size, in KBytes, of the XML response. The type is interface{}
    // with range: 1..100000. Units are kilobyte.
    StreamingSize interface{}

    // Session attributes.
    Session XrXml_Agent_Ssl_Session

    // XML agent throttling.
    Throttle XrXml_Agent_Ssl_Throttle

    // List of VRFs.
    Vrfs XrXml_Agent_Ssl_Vrfs
}

func (ssl *XrXml_Agent_Ssl) GetEntityData() *types.CommonEntityData {
    ssl.EntityData.YFilter = ssl.YFilter
    ssl.EntityData.YangName = "ssl"
    ssl.EntityData.BundleName = "cisco_ios_xr"
    ssl.EntityData.ParentYangName = "agent"
    ssl.EntityData.SegmentPath = "ssl"
    ssl.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ssl.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ssl.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ssl.EntityData.Children = types.NewOrderedMap()
    ssl.EntityData.Children.Append("session", types.YChild{"Session", &ssl.Session})
    ssl.EntityData.Children.Append("throttle", types.YChild{"Throttle", &ssl.Throttle})
    ssl.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &ssl.Vrfs})
    ssl.EntityData.Leafs = types.NewOrderedMap()
    ssl.EntityData.Leafs.Append("iteration-size", types.YLeaf{"IterationSize", ssl.IterationSize})
    ssl.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", ssl.Enable})
    ssl.EntityData.Leafs.Append("streaming-size", types.YLeaf{"StreamingSize", ssl.StreamingSize})

    ssl.EntityData.YListKeys = []string {}

    return &(ssl.EntityData)
}

// XrXml_Agent_Ssl_Session
// Session attributes
type XrXml_Agent_Ssl_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timeout in minutes. The type is interface{} with range: 1..1440. Units are
    // minute.
    Timeout interface{}
}

func (session *XrXml_Agent_Ssl_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "ssl"
    session.EntityData.SegmentPath = "session"
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", session.Timeout})

    session.EntityData.YListKeys = []string {}

    return &(session.EntityData)
}

// XrXml_Agent_Ssl_Throttle
// XML agent throttling
type XrXml_Agent_Ssl_Throttle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process rate in number of XML tags per second. The type is interface{} with
    // range: 1000..30000.
    ProcessRate interface{}

    // Size of memory usage, in MBytes, per session. The type is interface{} with
    // range: 100..1024. Units are megabyte. The default value is 300.
    Memory interface{}
}

func (throttle *XrXml_Agent_Ssl_Throttle) GetEntityData() *types.CommonEntityData {
    throttle.EntityData.YFilter = throttle.YFilter
    throttle.EntityData.YangName = "throttle"
    throttle.EntityData.BundleName = "cisco_ios_xr"
    throttle.EntityData.ParentYangName = "ssl"
    throttle.EntityData.SegmentPath = "throttle"
    throttle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    throttle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    throttle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    throttle.EntityData.Children = types.NewOrderedMap()
    throttle.EntityData.Leafs = types.NewOrderedMap()
    throttle.EntityData.Leafs.Append("process-rate", types.YLeaf{"ProcessRate", throttle.ProcessRate})
    throttle.EntityData.Leafs.Append("memory", types.YLeaf{"Memory", throttle.Memory})

    throttle.EntityData.YListKeys = []string {}

    return &(throttle.EntityData)
}

// XrXml_Agent_Ssl_Vrfs
// List of VRFs
type XrXml_Agent_Ssl_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A specific VRF. The type is slice of XrXml_Agent_Ssl_Vrfs_Vrf.
    Vrf []*XrXml_Agent_Ssl_Vrfs_Vrf
}

func (vrfs *XrXml_Agent_Ssl_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "ssl"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = types.NewOrderedMap()
    vrfs.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children.Append(types.GetSegmentPath(vrfs.Vrf[i]), types.YChild{"Vrf", vrfs.Vrf[i]})
    }
    vrfs.EntityData.Leafs = types.NewOrderedMap()

    vrfs.EntityData.YListKeys = []string {}

    return &(vrfs.EntityData)
}

// XrXml_Agent_Ssl_Vrfs_Vrf
// A specific VRF
type XrXml_Agent_Ssl_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with length: 1..32.
    VrfName interface{}

    // IPv6 Transport Access list for VRF. The type is string with length: 1..32.
    Ipv6AccessList interface{}

    // IPv4 Transport Access list for VRF. The type is string with length: 1..32.
    Ipv4AccessList interface{}

    // Access list for XML agent. The type is string with length: 1..32.
    AccessList interface{}

    // Shutdown default VRF. This is applicable only for VRF default. The type is
    // interface{}.
    Shutdown interface{}
}

func (vrf *XrXml_Agent_Ssl_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})
    vrf.EntityData.Leafs.Append("ipv6-access-list", types.YLeaf{"Ipv6AccessList", vrf.Ipv6AccessList})
    vrf.EntityData.Leafs.Append("ipv4-access-list", types.YLeaf{"Ipv4AccessList", vrf.Ipv4AccessList})
    vrf.EntityData.Leafs.Append("access-list", types.YLeaf{"AccessList", vrf.AccessList})
    vrf.EntityData.Leafs.Append("shutdown", types.YLeaf{"Shutdown", vrf.Shutdown})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// Netconf
// netconf
type Netconf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // XML agent.
    Agent Netconf_Agent
}

func (netconf *Netconf) GetEntityData() *types.CommonEntityData {
    netconf.EntityData.YFilter = netconf.YFilter
    netconf.EntityData.YangName = "netconf"
    netconf.EntityData.BundleName = "cisco_ios_xr"
    netconf.EntityData.ParentYangName = "Cisco-IOS-XR-man-xml-ttyagent-cfg"
    netconf.EntityData.SegmentPath = "Cisco-IOS-XR-man-xml-ttyagent-cfg:netconf"
    netconf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    netconf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    netconf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    netconf.EntityData.Children = types.NewOrderedMap()
    netconf.EntityData.Children.Append("agent", types.YChild{"Agent", &netconf.Agent})
    netconf.EntityData.Leafs = types.NewOrderedMap()

    netconf.EntityData.YListKeys = []string {}

    return &(netconf.EntityData)
}

// Netconf_Agent
// XML agent
type Netconf_Agent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NETCONF agent over TTY.
    Tty Netconf_Agent_Tty
}

func (agent *Netconf_Agent) GetEntityData() *types.CommonEntityData {
    agent.EntityData.YFilter = agent.YFilter
    agent.EntityData.YangName = "agent"
    agent.EntityData.BundleName = "cisco_ios_xr"
    agent.EntityData.ParentYangName = "netconf"
    agent.EntityData.SegmentPath = "agent"
    agent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    agent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    agent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    agent.EntityData.Children = types.NewOrderedMap()
    agent.EntityData.Children.Append("tty", types.YChild{"Tty", &agent.Tty})
    agent.EntityData.Leafs = types.NewOrderedMap()

    agent.EntityData.YListKeys = []string {}

    return &(agent.EntityData)
}

// Netconf_Agent_Tty
// NETCONF agent over TTY
type Netconf_Agent_Tty struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable specified NETCONF agent. The type is interface{}.
    Enable interface{}

    // NETCONF agent throttling.
    Throttle Netconf_Agent_Tty_Throttle

    // Session attributes.
    Session Netconf_Agent_Tty_Session
}

func (tty *Netconf_Agent_Tty) GetEntityData() *types.CommonEntityData {
    tty.EntityData.YFilter = tty.YFilter
    tty.EntityData.YangName = "tty"
    tty.EntityData.BundleName = "cisco_ios_xr"
    tty.EntityData.ParentYangName = "agent"
    tty.EntityData.SegmentPath = "tty"
    tty.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tty.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tty.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tty.EntityData.Children = types.NewOrderedMap()
    tty.EntityData.Children.Append("throttle", types.YChild{"Throttle", &tty.Throttle})
    tty.EntityData.Children.Append("session", types.YChild{"Session", &tty.Session})
    tty.EntityData.Leafs = types.NewOrderedMap()
    tty.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", tty.Enable})

    tty.EntityData.YListKeys = []string {}

    return &(tty.EntityData)
}

// Netconf_Agent_Tty_Throttle
// NETCONF agent throttling
type Netconf_Agent_Tty_Throttle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Size of memory usage, in MBytes, per session. The type is interface{} with
    // range: 100..1024. Units are megabyte. The default value is 300.
    Memory interface{}

    // Size of memory usage, in MBytes, per session. The type is interface{} with
    // range: 0..12000. Units are megabyte. The default value is 0.
    OffloadMemory interface{}

    // Process rate in number of XML tags per second. The type is interface{} with
    // range: 1000..30000.
    ProcessRate interface{}
}

func (throttle *Netconf_Agent_Tty_Throttle) GetEntityData() *types.CommonEntityData {
    throttle.EntityData.YFilter = throttle.YFilter
    throttle.EntityData.YangName = "throttle"
    throttle.EntityData.BundleName = "cisco_ios_xr"
    throttle.EntityData.ParentYangName = "tty"
    throttle.EntityData.SegmentPath = "throttle"
    throttle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    throttle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    throttle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    throttle.EntityData.Children = types.NewOrderedMap()
    throttle.EntityData.Leafs = types.NewOrderedMap()
    throttle.EntityData.Leafs.Append("memory", types.YLeaf{"Memory", throttle.Memory})
    throttle.EntityData.Leafs.Append("offload-memory", types.YLeaf{"OffloadMemory", throttle.OffloadMemory})
    throttle.EntityData.Leafs.Append("process-rate", types.YLeaf{"ProcessRate", throttle.ProcessRate})

    throttle.EntityData.YListKeys = []string {}

    return &(throttle.EntityData)
}

// Netconf_Agent_Tty_Session
// Session attributes
type Netconf_Agent_Tty_Session struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Timeout in minutes. The type is interface{} with range: 1..1440. Units are
    // minute.
    Timeout interface{}
}

func (session *Netconf_Agent_Tty_Session) GetEntityData() *types.CommonEntityData {
    session.EntityData.YFilter = session.YFilter
    session.EntityData.YangName = "session"
    session.EntityData.BundleName = "cisco_ios_xr"
    session.EntityData.ParentYangName = "tty"
    session.EntityData.SegmentPath = "session"
    session.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    session.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    session.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    session.EntityData.Children = types.NewOrderedMap()
    session.EntityData.Leafs = types.NewOrderedMap()
    session.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", session.Timeout})

    session.EntityData.YListKeys = []string {}

    return &(session.EntityData)
}

