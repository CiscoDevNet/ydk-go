// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv6-ospfv3 package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ospfv3: OSPFv3 configuration
// 
// This YANG module augments the
//   Cisco-IOS-XR-snmp-agent-cfg
// module with configuration data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ipv6_ospfv3_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv6_ospfv3_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv6-ospfv3-cfg ospfv3}", reflect.TypeOf(Ospfv3{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3", reflect.TypeOf(Ospfv3{}))
}

// Ospfv3Protocol represents Ospfv3 protocol
type Ospfv3Protocol string

const (
    // All
    Ospfv3Protocol_all Ospfv3Protocol = "all"

    // Connected
    Ospfv3Protocol_connected Ospfv3Protocol = "connected"

    // Static
    Ospfv3Protocol_static Ospfv3Protocol = "static"

    // BGP
    Ospfv3Protocol_bgp Ospfv3Protocol = "bgp"

    // ISIS
    Ospfv3Protocol_isis Ospfv3Protocol = "isis"

    // OSPFv3
    Ospfv3Protocol_ospfv3 Ospfv3Protocol = "ospfv3"

    // EIGRP
    Ospfv3Protocol_eigrp Ospfv3Protocol = "eigrp"
)

// Ospfv3FastRerouteTiebreakers represents Ospfv3 fast reroute tiebreakers
type Ospfv3FastRerouteTiebreakers string

const (
    // Downstream
    Ospfv3FastRerouteTiebreakers_downstream Ospfv3FastRerouteTiebreakers = "downstream"

    // LC Disjoint
    Ospfv3FastRerouteTiebreakers_line_card_disjoint Ospfv3FastRerouteTiebreakers = "line-card-disjoint"

    // Lowest metric
    Ospfv3FastRerouteTiebreakers_lowest_metric Ospfv3FastRerouteTiebreakers = "lowest-metric"

    // Node protection
    Ospfv3FastRerouteTiebreakers_node_protect Ospfv3FastRerouteTiebreakers = "node-protect"

    // Primary path
    Ospfv3FastRerouteTiebreakers_primary_path Ospfv3FastRerouteTiebreakers = "primary-path"

    // Secondar path
    Ospfv3FastRerouteTiebreakers_secondary_path Ospfv3FastRerouteTiebreakers = "secondary-path"

    // SRLG
    Ospfv3FastRerouteTiebreakers_srlg_disjoint Ospfv3FastRerouteTiebreakers = "srlg-disjoint"
)

// Ospfv3isisRoute represents Ospfv3isis route
type Ospfv3isisRoute string

const (
    // IS-IS level-1 routes
    Ospfv3isisRoute_level1 Ospfv3isisRoute = "level1"

    // IS-IS level-2 routes
    Ospfv3isisRoute_level2 Ospfv3isisRoute = "level2"

    // IS-IS level-1 and level-2 routes
    Ospfv3isisRoute_level1_and2 Ospfv3isisRoute = "level1-and2"
)

// Ospfv3bfdEnableMode represents Ospfv3bfd enable mode
type Ospfv3bfdEnableMode string

const (
    // Disable Mode - Prevent inheritance
    Ospfv3bfdEnableMode_disable Ospfv3bfdEnableMode = "disable"

    // Default Mode - Default BFD behavior
    Ospfv3bfdEnableMode_default_ Ospfv3bfdEnableMode = "default"

    // Strict Mode - Hold down adj until BFD sesion up
    Ospfv3bfdEnableMode_strict Ospfv3bfdEnableMode = "strict"
)

// Ospfv3LogAdj represents Ospfv3 log adj
type Ospfv3LogAdj string

const (
    // No output
    Ospfv3LogAdj_suppress Ospfv3LogAdj = "suppress"

    // Limited output
    Ospfv3LogAdj_brief Ospfv3LogAdj = "brief"

    // Verbose output
    Ospfv3LogAdj_detail Ospfv3LogAdj = "detail"
)

// Ospfv3ProtocolType2 represents Ospfv3 protocol type2
type Ospfv3ProtocolType2 string

const (
    // Connected
    Ospfv3ProtocolType2_connected Ospfv3ProtocolType2 = "connected"

    // Static
    Ospfv3ProtocolType2_static Ospfv3ProtocolType2 = "static"

    // BGP
    Ospfv3ProtocolType2_bgp Ospfv3ProtocolType2 = "bgp"

    // ISIS
    Ospfv3ProtocolType2_isis Ospfv3ProtocolType2 = "isis"

    // OSPFv3
    Ospfv3ProtocolType2_ospfv3 Ospfv3ProtocolType2 = "ospfv3"

    // EIGRP
    Ospfv3ProtocolType2_eigrp Ospfv3ProtocolType2 = "eigrp"

    // Subscriber
    Ospfv3ProtocolType2_subscriber Ospfv3ProtocolType2 = "subscriber"

    // Application
    Ospfv3ProtocolType2_application Ospfv3ProtocolType2 = "application"

    // Mobile
    Ospfv3ProtocolType2_mobile Ospfv3ProtocolType2 = "mobile"
)

// Ospfv3Metric represents Ospfv3 metric
type Ospfv3Metric string

const (
    // OSPFv3 external type 1 metrics
    Ospfv3Metric_type1 Ospfv3Metric = "type1"

    // OSPFv3 external type 2 metrics
    Ospfv3Metric_type2 Ospfv3Metric = "type2"
)

// Ospfv3TraceBufSize represents Ospfv3 trace buf size
type Ospfv3TraceBufSize string

const (
    // Disable trace
    Ospfv3TraceBufSize_size0 Ospfv3TraceBufSize = "size0"

    // trace buffer size 256
    Ospfv3TraceBufSize_size256 Ospfv3TraceBufSize = "size256"

    // trace buffer size 512
    Ospfv3TraceBufSize_size512 Ospfv3TraceBufSize = "size512"

    // trace buffer size 1024
    Ospfv3TraceBufSize_size1024 Ospfv3TraceBufSize = "size1024"

    // trace buffer size 2048
    Ospfv3TraceBufSize_size2048 Ospfv3TraceBufSize = "size2048"

    // trace buffer size 4096
    Ospfv3TraceBufSize_size4096 Ospfv3TraceBufSize = "size4096"

    // trace buffer size 8192
    Ospfv3TraceBufSize_size8192 Ospfv3TraceBufSize = "size8192"

    // trace buffer size 16384
    Ospfv3TraceBufSize_size16384 Ospfv3TraceBufSize = "size16384"

    // trace buffer size 32768
    Ospfv3TraceBufSize_size32768 Ospfv3TraceBufSize = "size32768"

    // trace buffer size 65536
    Ospfv3TraceBufSize_size65536 Ospfv3TraceBufSize = "size65536"
)

// Ospfv3ExternalRoute represents Ospfv3 external route
type Ospfv3ExternalRoute string

const (
    // External type 1 routes
    Ospfv3ExternalRoute_external1 Ospfv3ExternalRoute = "external1"

    // External type 2 routes
    Ospfv3ExternalRoute_external2 Ospfv3ExternalRoute = "external2"

    // External (type 1 and 2) routes
    Ospfv3ExternalRoute_external Ospfv3ExternalRoute = "external"
)

// Ospfv3SubsequentAddressFamily represents Ospfv3 subsequent address family
type Ospfv3SubsequentAddressFamily string

const (
    // Unicast subsequent address family
    Ospfv3SubsequentAddressFamily_unicast Ospfv3SubsequentAddressFamily = "unicast"
)

// Ospfv3InternalRoute represents Ospfv3 internal route
type Ospfv3InternalRoute string

const (
    // OSPFv3 internal routes
    Ospfv3InternalRoute_internal Ospfv3InternalRoute = "internal"
)

// Ospfv3FastReroute represents Ospfv3 fast reroute
type Ospfv3FastReroute string

const (
    // Disable
    Ospfv3FastReroute_none Ospfv3FastReroute = "none"

    // Per link
    Ospfv3FastReroute_per_link Ospfv3FastReroute = "per-link"

    // Per prefix
    Ospfv3FastReroute_per_prefix Ospfv3FastReroute = "per-prefix"
)

// Ospfv3DomainId represents Ospfv3 domain id
type Ospfv3DomainId string

const (
    // Type 0x0005
    Ospfv3DomainId_type0005 Ospfv3DomainId = "type0005"

    // Type 0x0105
    Ospfv3DomainId_type0105 Ospfv3DomainId = "type0105"

    // Type 0x0205
    Ospfv3DomainId_type0205 Ospfv3DomainId = "type0205"

    // Type 0x8005
    Ospfv3DomainId_type8005 Ospfv3DomainId = "type8005"
)

// Ospfv3AuthenticationType2 represents Ospfv3 authentication type2
type Ospfv3AuthenticationType2 string

const (
    // NULL authentication
    Ospfv3AuthenticationType2_null Ospfv3AuthenticationType2 = "null"

    // MD5 algorithm
    Ospfv3AuthenticationType2_md5 Ospfv3AuthenticationType2 = "md5"

    // SHA1 algorithm
    Ospfv3AuthenticationType2_sha1 Ospfv3AuthenticationType2 = "sha1"
)

// Ospfv3AddressFamily represents Ospfv3 address family
type Ospfv3AddressFamily string

const (
    // IPv6 address family
    Ospfv3AddressFamily_ipv6 Ospfv3AddressFamily = "ipv6"
)

// Ospfv3Authentication represents Ospfv3 authentication
type Ospfv3Authentication string

const (
    // MD5 algorithm
    Ospfv3Authentication_md5 Ospfv3Authentication = "md5"

    // SHA1 algorithm
    Ospfv3Authentication_sha1 Ospfv3Authentication = "sha1"
)

// Ospfv3EncryptionAlgorithm represents Ospfv3 encryption algorithm
type Ospfv3EncryptionAlgorithm string

const (
    // Use NULL encryption
    Ospfv3EncryptionAlgorithm_null Ospfv3EncryptionAlgorithm = "null"

    // Use the DES algorithm
    Ospfv3EncryptionAlgorithm_des Ospfv3EncryptionAlgorithm = "des"

    // Use the triple DES algorithm
    Ospfv3EncryptionAlgorithm_Y_3des Ospfv3EncryptionAlgorithm = "3des"

    // Use the AES algorithm
    Ospfv3EncryptionAlgorithm_aes Ospfv3EncryptionAlgorithm = "aes"

    // Use the 192-bit AES algorithm
    Ospfv3EncryptionAlgorithm_aes192 Ospfv3EncryptionAlgorithm = "aes192"

    // Use the 256-bit AES algorithm
    Ospfv3EncryptionAlgorithm_aes256 Ospfv3EncryptionAlgorithm = "aes256"
)

// Ospfv3nsr represents Ospfv3nsr
type Ospfv3nsr string

const (
    // Enable non-stop routing
    Ospfv3nsr_true_ Ospfv3nsr = "true"

    // Disable non-stop routing
    Ospfv3nsr_false_ Ospfv3nsr = "false"
)

// Ospfv3nssaExternalRoute represents Ospfv3nssa external route
type Ospfv3nssaExternalRoute string

const (
    // NSSA external type 1 routes
    Ospfv3nssaExternalRoute_external1 Ospfv3nssaExternalRoute = "external1"

    // NSSA external type 2 routes
    Ospfv3nssaExternalRoute_external2 Ospfv3nssaExternalRoute = "external2"

    // NSSA external (type 1 and 2) routes
    Ospfv3nssaExternalRoute_external Ospfv3nssaExternalRoute = "external"
)

// Ospfv3EigrpRoute represents Ospfv3 eigrp route
type Ospfv3EigrpRoute string

const (
    // EIGRP internal routes
    Ospfv3EigrpRoute_internal Ospfv3EigrpRoute = "internal"

    // EIGRP external routes
    Ospfv3EigrpRoute_external Ospfv3EigrpRoute = "external"
)

// Ospfv3FastReroutePriority represents Ospfv3 fast reroute priority
type Ospfv3FastReroutePriority string

const (
    // Critical
    Ospfv3FastReroutePriority_critical Ospfv3FastReroutePriority = "critical"

    // High
    Ospfv3FastReroutePriority_high Ospfv3FastReroutePriority = "high"

    // Medium
    Ospfv3FastReroutePriority_medium Ospfv3FastReroutePriority = "medium"

    // Low
    Ospfv3FastReroutePriority_low Ospfv3FastReroutePriority = "low"
)

// Ospfv3Network represents Ospfv3 network
type Ospfv3Network string

const (
    // Broadcast multi-access network
    Ospfv3Network_broadcast Ospfv3Network = "broadcast"

    // Non-broadcast multi-access network
    Ospfv3Network_non_broadcast Ospfv3Network = "non-broadcast"

    // Point-to-point network
    Ospfv3Network_point_to_point Ospfv3Network = "point-to-point"

    // Point-to-multipoint network
    Ospfv3Network_point_to_multipoint Ospfv3Network = "point-to-multipoint"

    // Non-broadcast point-to-multipoint network
    Ospfv3Network_non_broadcast_point_to_multipoint Ospfv3Network = "non-broadcast-point-to-multipoint"
)

// Ospfv3
// OSPFv3 configuration
type Ospfv3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable OSPFv3 router IDs as DNS names. The type is interface{}.
    DnsNameLookup interface{}

    // OSPFv3 processes.
    Processes Ospfv3_Processes
}

func (ospfv3 *Ospfv3) GetEntityData() *types.CommonEntityData {
    ospfv3.EntityData.YFilter = ospfv3.YFilter
    ospfv3.EntityData.YangName = "ospfv3"
    ospfv3.EntityData.BundleName = "cisco_ios_xr"
    ospfv3.EntityData.ParentYangName = "Cisco-IOS-XR-ipv6-ospfv3-cfg"
    ospfv3.EntityData.SegmentPath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3"
    ospfv3.EntityData.AbsolutePath = ospfv3.EntityData.SegmentPath
    ospfv3.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv3.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv3.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv3.EntityData.Children = types.NewOrderedMap()
    ospfv3.EntityData.Children.Append("processes", types.YChild{"Processes", &ospfv3.Processes})
    ospfv3.EntityData.Leafs = types.NewOrderedMap()
    ospfv3.EntityData.Leafs.Append("dns-name-lookup", types.YLeaf{"DnsNameLookup", ospfv3.DnsNameLookup})

    ospfv3.EntityData.YListKeys = []string {}

    return &(ospfv3.EntityData)
}

// Ospfv3_Processes
// OSPFv3 processes
type Ospfv3_Processes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An OSPFv3 process. The type is slice of Ospfv3_Processes_Process.
    Process []*Ospfv3_Processes_Process
}

func (processes *Ospfv3_Processes) GetEntityData() *types.CommonEntityData {
    processes.EntityData.YFilter = processes.YFilter
    processes.EntityData.YangName = "processes"
    processes.EntityData.BundleName = "cisco_ios_xr"
    processes.EntityData.ParentYangName = "ospfv3"
    processes.EntityData.SegmentPath = "processes"
    processes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/" + processes.EntityData.SegmentPath
    processes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processes.EntityData.Children = types.NewOrderedMap()
    processes.EntityData.Children.Append("process", types.YChild{"Process", nil})
    for i := range processes.Process {
        processes.EntityData.Children.Append(types.GetSegmentPath(processes.Process[i]), types.YChild{"Process", processes.Process[i]})
    }
    processes.EntityData.Leafs = types.NewOrderedMap()

    processes.EntityData.YListKeys = []string {}

    return &(processes.EntityData)
}

// Ospfv3_Processes_Process
// An OSPFv3 process
type Ospfv3_Processes_Process struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. OSPFv3 process name. The type is string with
    // length: 1..32.
    ProcessName interface{}

    // Enable non-stop routing. The type is Ospfv3nsr. The default value is true.
    Nsr interface{}

    // Enable protocol shutdown. The type is interface{}.
    ProtocolShutdown interface{}

    // Default VRF related configuration.
    DefaultVrf Ospfv3_Processes_Process_DefaultVrf

    // VRF related configuration.
    Vrfs Ospfv3_Processes_Process_Vrfs

    // Address Family (AF).
    Af Ospfv3_Processes_Process_Af

    // Configuration to change size of trace buffer.
    TraceBufs Ospfv3_Processes_Process_TraceBufs
}

func (process *Ospfv3_Processes_Process) GetEntityData() *types.CommonEntityData {
    process.EntityData.YFilter = process.YFilter
    process.EntityData.YangName = "process"
    process.EntityData.BundleName = "cisco_ios_xr"
    process.EntityData.ParentYangName = "processes"
    process.EntityData.SegmentPath = "process" + types.AddKeyToken(process.ProcessName, "process-name")
    process.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/" + process.EntityData.SegmentPath
    process.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    process.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    process.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    process.EntityData.Children = types.NewOrderedMap()
    process.EntityData.Children.Append("default-vrf", types.YChild{"DefaultVrf", &process.DefaultVrf})
    process.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &process.Vrfs})
    process.EntityData.Children.Append("af", types.YChild{"Af", &process.Af})
    process.EntityData.Children.Append("trace-bufs", types.YChild{"TraceBufs", &process.TraceBufs})
    process.EntityData.Leafs = types.NewOrderedMap()
    process.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", process.ProcessName})
    process.EntityData.Leafs.Append("nsr", types.YLeaf{"Nsr", process.Nsr})
    process.EntityData.Leafs.Append("protocol-shutdown", types.YLeaf{"ProtocolShutdown", process.ProtocolShutdown})

    process.EntityData.YListKeys = []string {"ProcessName"}

    return &(process.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf
// Default VRF related configuration
type Ospfv3_Processes_Process_DefaultVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable MPLS LDP sync. The type is bool.
    LdpSync interface{}

    // Enable/disable prefix suppression on an interface. The type is bool.
    PrefixSuppression interface{}

    // SPF prefix prioritization disabled. The type is interface{}.
    SpfPrefixPriorityDisable interface{}

    // Specify the transmit interval in seconds. The type is interface{} with
    // range: 1..65535. Units are second.
    RetransmitInterval interface{}

    // Enable/disable routing updates on an interface. The type is bool.
    Passive interface{}

    // Set metric of redistributed routes. The type is interface{} with range:
    // 1..16777214.
    DefaultMetric interface{}

    // Enable/disable flood reduction. The type is bool.
    FloodReduction interface{}

    // Time between HELLO packets. The type is interface{} with range: 1..65535.
    // Units are second.
    HelloInterval interface{}

    // Specify router priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // Interface cost. The type is interface{} with range: 1..65535.
    Cost interface{}

    // Interval after which a neighbor is declared dead (in seconds). The type is
    // interface{} with range: 1..65535. Units are second.
    DeadInterval interface{}

    // Limit size of OSPFv3 packets. The type is interface{} with range:
    // 256..10000.
    PacketSize interface{}

    // Instance ID. The type is interface{} with range: 0..255.
    Instance interface{}

    // Route policy for SPF prefix prioritization. The type is string.
    SpfPrefixPriorityPolicy interface{}

    // Specify the router ID for this OSPFv3 process in IPv4 address format. The
    // type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Specify network type. The type is Ospfv3Network.
    Network interface{}

    // Enable/disable ignoring of MTU in DBD packets. The type is bool.
    MtuIgnore interface{}

    // Log changes in adjacency state. The type is Ospfv3LogAdj.
    LogAdjacencyChanges interface{}

    // Enable/disable demand circuit operation. The type is bool.
    DemandCircuit interface{}

    // Specify the transmit delay in seconds. The type is interface{} with range:
    // 1..65535. Units are second.
    TransmitDelay interface{}

    // Area configuration.
    AreaAddresses Ospfv3_Processes_Process_DefaultVrf_AreaAddresses

    // Adjust routing timers.
    Timers Ospfv3_Processes_Process_DefaultVrf_Timers

    // Summarize redistributed routes matching prefix/length.
    SummaryPrefixes Ospfv3_Processes_Process_DefaultVrf_SummaryPrefixes

    // SNMP configuration.
    Snmp Ospfv3_Processes_Process_DefaultVrf_Snmp

    // Fast-reroute instance scoped parameters.
    FastReroute Ospfv3_Processes_Process_DefaultVrf_FastReroute

    // Define an administrative distance.
    Distance Ospfv3_Processes_Process_DefaultVrf_Distance

    // Set OSPFv3 limits.
    Maximum Ospfv3_Processes_Process_DefaultVrf_Maximum

    // Redistribute information from another routing protocol.
    Redistributes Ospfv3_Processes_Process_DefaultVrf_Redistributes

    // Do not complain about a specified event.
    Ignore Ospfv3_Processes_Process_DefaultVrf_Ignore

    // Filter prefixes from RIB .
    DistributeListOut Ospfv3_Processes_Process_DefaultVrf_DistributeListOut

    // Filter prefixes to/from RIB.
    DistributeList Ospfv3_Processes_Process_DefaultVrf_DistributeList

    // Stub router configuration.
    StubRouter Ospfv3_Processes_Process_DefaultVrf_StubRouter

    // Configure BFD parameters.
    Bfd Ospfv3_Processes_Process_DefaultVrf_Bfd

    // Database filter.
    DatabaseFilter Ospfv3_Processes_Process_DefaultVrf_DatabaseFilter

    // OSPFv3 Capability.
    Capability Ospfv3_Processes_Process_DefaultVrf_Capability

    // Authenticate OSPFv3 packets.
    Authentication Ospfv3_Processes_Process_DefaultVrf_Authentication

    // Graceful restart configuration.
    GracefulRestart Ospfv3_Processes_Process_DefaultVrf_GracefulRestart

    // Control distribution of default information.
    DefaultInformation Ospfv3_Processes_Process_DefaultVrf_DefaultInformation

    // Process scope configuration.
    ProcessScope Ospfv3_Processes_Process_DefaultVrf_ProcessScope

    // Encrypt and authenticate OSPFv3 packets.
    Encryption Ospfv3_Processes_Process_DefaultVrf_Encryption

    // Calculate interface cost according to bandwidth.
    AutoCost Ospfv3_Processes_Process_DefaultVrf_AutoCost
}

func (defaultVrf *Ospfv3_Processes_Process_DefaultVrf) GetEntityData() *types.CommonEntityData {
    defaultVrf.EntityData.YFilter = defaultVrf.YFilter
    defaultVrf.EntityData.YangName = "default-vrf"
    defaultVrf.EntityData.BundleName = "cisco_ios_xr"
    defaultVrf.EntityData.ParentYangName = "process"
    defaultVrf.EntityData.SegmentPath = "default-vrf"
    defaultVrf.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/" + defaultVrf.EntityData.SegmentPath
    defaultVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultVrf.EntityData.Children = types.NewOrderedMap()
    defaultVrf.EntityData.Children.Append("area-addresses", types.YChild{"AreaAddresses", &defaultVrf.AreaAddresses})
    defaultVrf.EntityData.Children.Append("timers", types.YChild{"Timers", &defaultVrf.Timers})
    defaultVrf.EntityData.Children.Append("summary-prefixes", types.YChild{"SummaryPrefixes", &defaultVrf.SummaryPrefixes})
    defaultVrf.EntityData.Children.Append("snmp", types.YChild{"Snmp", &defaultVrf.Snmp})
    defaultVrf.EntityData.Children.Append("fast-reroute", types.YChild{"FastReroute", &defaultVrf.FastReroute})
    defaultVrf.EntityData.Children.Append("distance", types.YChild{"Distance", &defaultVrf.Distance})
    defaultVrf.EntityData.Children.Append("maximum", types.YChild{"Maximum", &defaultVrf.Maximum})
    defaultVrf.EntityData.Children.Append("redistributes", types.YChild{"Redistributes", &defaultVrf.Redistributes})
    defaultVrf.EntityData.Children.Append("ignore", types.YChild{"Ignore", &defaultVrf.Ignore})
    defaultVrf.EntityData.Children.Append("distribute-list-out", types.YChild{"DistributeListOut", &defaultVrf.DistributeListOut})
    defaultVrf.EntityData.Children.Append("distribute-list", types.YChild{"DistributeList", &defaultVrf.DistributeList})
    defaultVrf.EntityData.Children.Append("stub-router", types.YChild{"StubRouter", &defaultVrf.StubRouter})
    defaultVrf.EntityData.Children.Append("bfd", types.YChild{"Bfd", &defaultVrf.Bfd})
    defaultVrf.EntityData.Children.Append("database-filter", types.YChild{"DatabaseFilter", &defaultVrf.DatabaseFilter})
    defaultVrf.EntityData.Children.Append("capability", types.YChild{"Capability", &defaultVrf.Capability})
    defaultVrf.EntityData.Children.Append("authentication", types.YChild{"Authentication", &defaultVrf.Authentication})
    defaultVrf.EntityData.Children.Append("graceful-restart", types.YChild{"GracefulRestart", &defaultVrf.GracefulRestart})
    defaultVrf.EntityData.Children.Append("default-information", types.YChild{"DefaultInformation", &defaultVrf.DefaultInformation})
    defaultVrf.EntityData.Children.Append("process-scope", types.YChild{"ProcessScope", &defaultVrf.ProcessScope})
    defaultVrf.EntityData.Children.Append("encryption", types.YChild{"Encryption", &defaultVrf.Encryption})
    defaultVrf.EntityData.Children.Append("auto-cost", types.YChild{"AutoCost", &defaultVrf.AutoCost})
    defaultVrf.EntityData.Leafs = types.NewOrderedMap()
    defaultVrf.EntityData.Leafs.Append("ldp-sync", types.YLeaf{"LdpSync", defaultVrf.LdpSync})
    defaultVrf.EntityData.Leafs.Append("prefix-suppression", types.YLeaf{"PrefixSuppression", defaultVrf.PrefixSuppression})
    defaultVrf.EntityData.Leafs.Append("spf-prefix-priority-disable", types.YLeaf{"SpfPrefixPriorityDisable", defaultVrf.SpfPrefixPriorityDisable})
    defaultVrf.EntityData.Leafs.Append("retransmit-interval", types.YLeaf{"RetransmitInterval", defaultVrf.RetransmitInterval})
    defaultVrf.EntityData.Leafs.Append("passive", types.YLeaf{"Passive", defaultVrf.Passive})
    defaultVrf.EntityData.Leafs.Append("default-metric", types.YLeaf{"DefaultMetric", defaultVrf.DefaultMetric})
    defaultVrf.EntityData.Leafs.Append("flood-reduction", types.YLeaf{"FloodReduction", defaultVrf.FloodReduction})
    defaultVrf.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", defaultVrf.HelloInterval})
    defaultVrf.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", defaultVrf.Priority})
    defaultVrf.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", defaultVrf.Cost})
    defaultVrf.EntityData.Leafs.Append("dead-interval", types.YLeaf{"DeadInterval", defaultVrf.DeadInterval})
    defaultVrf.EntityData.Leafs.Append("packet-size", types.YLeaf{"PacketSize", defaultVrf.PacketSize})
    defaultVrf.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", defaultVrf.Instance})
    defaultVrf.EntityData.Leafs.Append("spf-prefix-priority-policy", types.YLeaf{"SpfPrefixPriorityPolicy", defaultVrf.SpfPrefixPriorityPolicy})
    defaultVrf.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", defaultVrf.RouterId})
    defaultVrf.EntityData.Leafs.Append("network", types.YLeaf{"Network", defaultVrf.Network})
    defaultVrf.EntityData.Leafs.Append("mtu-ignore", types.YLeaf{"MtuIgnore", defaultVrf.MtuIgnore})
    defaultVrf.EntityData.Leafs.Append("log-adjacency-changes", types.YLeaf{"LogAdjacencyChanges", defaultVrf.LogAdjacencyChanges})
    defaultVrf.EntityData.Leafs.Append("demand-circuit", types.YLeaf{"DemandCircuit", defaultVrf.DemandCircuit})
    defaultVrf.EntityData.Leafs.Append("transmit-delay", types.YLeaf{"TransmitDelay", defaultVrf.TransmitDelay})

    defaultVrf.EntityData.YListKeys = []string {}

    return &(defaultVrf.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses
// Area configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for a particular area. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress.
    AreaAddress []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress

    // Configuration for a particular area. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId.
    AreaAreaId []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId
}

func (areaAddresses *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses) GetEntityData() *types.CommonEntityData {
    areaAddresses.EntityData.YFilter = areaAddresses.YFilter
    areaAddresses.EntityData.YangName = "area-addresses"
    areaAddresses.EntityData.BundleName = "cisco_ios_xr"
    areaAddresses.EntityData.ParentYangName = "default-vrf"
    areaAddresses.EntityData.SegmentPath = "area-addresses"
    areaAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/" + areaAddresses.EntityData.SegmentPath
    areaAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    areaAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    areaAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    areaAddresses.EntityData.Children = types.NewOrderedMap()
    areaAddresses.EntityData.Children.Append("area-address", types.YChild{"AreaAddress", nil})
    for i := range areaAddresses.AreaAddress {
        areaAddresses.EntityData.Children.Append(types.GetSegmentPath(areaAddresses.AreaAddress[i]), types.YChild{"AreaAddress", areaAddresses.AreaAddress[i]})
    }
    areaAddresses.EntityData.Children.Append("area-area-id", types.YChild{"AreaAreaId", nil})
    for i := range areaAddresses.AreaAreaId {
        areaAddresses.EntityData.Children.Append(types.GetSegmentPath(areaAddresses.AreaAreaId[i]), types.YChild{"AreaAreaId", areaAddresses.AreaAreaId[i]})
    }
    areaAddresses.EntityData.Leafs = types.NewOrderedMap()

    areaAddresses.EntityData.YListKeys = []string {}

    return &(areaAddresses.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress
// Configuration for a particular area
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Area ID if in IP address format. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Specify area as a stub area.  Allowed only in non-backbone areas. The type
    // is bool.
    Stub interface{}

    // Limit size of OSPFv3 packets. The type is interface{} with range:
    // 256..10000.
    PacketSize interface{}

    // Instance ID. The type is interface{} with range: 0..255.
    Instance interface{}

    // Enable/disable demand circuit operation. The type is bool.
    DemandCircuit interface{}

    // Specify router priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // Translate Type 7 to Type 5, even if not elected NSSA translator. The type
    // is bool.
    Type7TranslateAlways interface{}

    // Enable/disable prefix suppression on an interface. The type is bool.
    PrefixSuppression interface{}

    // Enable OSPFv3 area. The type is interface{}.
    Enable interface{}

    // Enable/disable ignoring of MTU in DBD packets. The type is bool.
    MtuIgnore interface{}

    // Enable/disable routing updates on an interface. The type is bool.
    Passive interface{}

    // Time between HELLO packets. The type is interface{} with range: 1..65535.
    // Units are second.
    HelloInterval interface{}

    // Interval after which a neighbor is declared dead (in seconds). The type is
    // interface{} with range: 1..65535. Units are second.
    DeadInterval interface{}

    // Set the summary default-cost of a NSSA/stub area. The type is interface{}
    // with range: 0..16777215.
    DefaultCost interface{}

    // Enable/disable flood reduction. The type is bool.
    FloodReduction interface{}

    // Specify the transmit interval in seconds. The type is interface{} with
    // range: 1..65535. Units are second.
    RetransmitInterval interface{}

    // Enable/Disable MPLS LDP sync. The type is bool.
    LdpSync interface{}

    // Specify network type. The type is Ospfv3Network.
    Network interface{}

    // Specify the transmit delay in seconds. The type is interface{} with range:
    // 1..65535. Units are second.
    TransmitDelay interface{}

    // Interface cost. The type is interface{} with range: 1..65535.
    Cost interface{}

    // Authenticate OSPFv3 packets.
    Authentication Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Authentication

    // Configure BFD parameters.
    Bfd Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Bfd

    // Range configuration.
    Ranges Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Ranges

    // Encrypt and authenticate OSPFv3 packets.
    Encryption Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Encryption

    // Specify area as a NSSA area.  Allowed only in non-backbone areas.
    Nssa Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Nssa

    // Database filter.
    DatabaseFilter Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_DatabaseFilter

    // Filter prefixes to/from RIB.
    DistributeList Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_DistributeList

    // OSPFv3 interfaces.
    Interfaces Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces

    // Area Scope Configuration.
    AreaScope Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope

    // Sham Link sub-mode.
    ShamLinks Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_ShamLinks

    // Virtual link sub-mode.
    VirtualLinks Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_VirtualLinks
}

func (areaAddress *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress) GetEntityData() *types.CommonEntityData {
    areaAddress.EntityData.YFilter = areaAddress.YFilter
    areaAddress.EntityData.YangName = "area-address"
    areaAddress.EntityData.BundleName = "cisco_ios_xr"
    areaAddress.EntityData.ParentYangName = "area-addresses"
    areaAddress.EntityData.SegmentPath = "area-address" + types.AddKeyToken(areaAddress.Address, "address")
    areaAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/" + areaAddress.EntityData.SegmentPath
    areaAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    areaAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    areaAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    areaAddress.EntityData.Children = types.NewOrderedMap()
    areaAddress.EntityData.Children.Append("authentication", types.YChild{"Authentication", &areaAddress.Authentication})
    areaAddress.EntityData.Children.Append("bfd", types.YChild{"Bfd", &areaAddress.Bfd})
    areaAddress.EntityData.Children.Append("ranges", types.YChild{"Ranges", &areaAddress.Ranges})
    areaAddress.EntityData.Children.Append("encryption", types.YChild{"Encryption", &areaAddress.Encryption})
    areaAddress.EntityData.Children.Append("nssa", types.YChild{"Nssa", &areaAddress.Nssa})
    areaAddress.EntityData.Children.Append("database-filter", types.YChild{"DatabaseFilter", &areaAddress.DatabaseFilter})
    areaAddress.EntityData.Children.Append("distribute-list", types.YChild{"DistributeList", &areaAddress.DistributeList})
    areaAddress.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &areaAddress.Interfaces})
    areaAddress.EntityData.Children.Append("area-scope", types.YChild{"AreaScope", &areaAddress.AreaScope})
    areaAddress.EntityData.Children.Append("sham-links", types.YChild{"ShamLinks", &areaAddress.ShamLinks})
    areaAddress.EntityData.Children.Append("virtual-links", types.YChild{"VirtualLinks", &areaAddress.VirtualLinks})
    areaAddress.EntityData.Leafs = types.NewOrderedMap()
    areaAddress.EntityData.Leafs.Append("address", types.YLeaf{"Address", areaAddress.Address})
    areaAddress.EntityData.Leafs.Append("stub", types.YLeaf{"Stub", areaAddress.Stub})
    areaAddress.EntityData.Leafs.Append("packet-size", types.YLeaf{"PacketSize", areaAddress.PacketSize})
    areaAddress.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", areaAddress.Instance})
    areaAddress.EntityData.Leafs.Append("demand-circuit", types.YLeaf{"DemandCircuit", areaAddress.DemandCircuit})
    areaAddress.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", areaAddress.Priority})
    areaAddress.EntityData.Leafs.Append("type7-translate-always", types.YLeaf{"Type7TranslateAlways", areaAddress.Type7TranslateAlways})
    areaAddress.EntityData.Leafs.Append("prefix-suppression", types.YLeaf{"PrefixSuppression", areaAddress.PrefixSuppression})
    areaAddress.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", areaAddress.Enable})
    areaAddress.EntityData.Leafs.Append("mtu-ignore", types.YLeaf{"MtuIgnore", areaAddress.MtuIgnore})
    areaAddress.EntityData.Leafs.Append("passive", types.YLeaf{"Passive", areaAddress.Passive})
    areaAddress.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", areaAddress.HelloInterval})
    areaAddress.EntityData.Leafs.Append("dead-interval", types.YLeaf{"DeadInterval", areaAddress.DeadInterval})
    areaAddress.EntityData.Leafs.Append("default-cost", types.YLeaf{"DefaultCost", areaAddress.DefaultCost})
    areaAddress.EntityData.Leafs.Append("flood-reduction", types.YLeaf{"FloodReduction", areaAddress.FloodReduction})
    areaAddress.EntityData.Leafs.Append("retransmit-interval", types.YLeaf{"RetransmitInterval", areaAddress.RetransmitInterval})
    areaAddress.EntityData.Leafs.Append("ldp-sync", types.YLeaf{"LdpSync", areaAddress.LdpSync})
    areaAddress.EntityData.Leafs.Append("network", types.YLeaf{"Network", areaAddress.Network})
    areaAddress.EntityData.Leafs.Append("transmit-delay", types.YLeaf{"TransmitDelay", areaAddress.TransmitDelay})
    areaAddress.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", areaAddress.Cost})

    areaAddress.EntityData.YListKeys = []string {"Address"}

    return &(areaAddress.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Authentication
// Authenticate OSPFv3 packets
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec AH authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Use the MD5 or SHA1 algorithm. The type is Ospfv3Authentication.
    Algorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    Password interface{}
}

func (authentication *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "area-address"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", authentication.Enable})
    authentication.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", authentication.Spi})
    authentication.EntityData.Leafs.Append("algorithm", types.YLeaf{"Algorithm", authentication.Algorithm})
    authentication.EntityData.Leafs.Append("password", types.YLeaf{"Password", authentication.Password})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Bfd
// Configure BFD parameters
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detect multiplier. The type is interface{} with range: 2..50.
    DetectionMultiplier interface{}

    // Hello interval in milli-seconds. The type is interface{} with range:
    // 3..30000. Units are millisecond.
    Interval interface{}

    // Enable or disable BFD fast detection. The type is Ospfv3bfdEnableMode.
    FastDetectMode interface{}
}

func (bfd *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "area-address"
    bfd.EntityData.SegmentPath = "bfd"
    bfd.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/" + bfd.EntityData.SegmentPath
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = types.NewOrderedMap()
    bfd.EntityData.Leafs = types.NewOrderedMap()
    bfd.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", bfd.DetectionMultiplier})
    bfd.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", bfd.Interval})
    bfd.EntityData.Leafs.Append("fast-detect-mode", types.YLeaf{"FastDetectMode", bfd.FastDetectMode})

    bfd.EntityData.YListKeys = []string {}

    return &(bfd.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Ranges
// Range configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Ranges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summarize inter-area routes matching prefix/length. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Ranges_Range.
    Range []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Ranges_Range
}

func (ranges *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Ranges) GetEntityData() *types.CommonEntityData {
    ranges.EntityData.YFilter = ranges.YFilter
    ranges.EntityData.YangName = "ranges"
    ranges.EntityData.BundleName = "cisco_ios_xr"
    ranges.EntityData.ParentYangName = "area-address"
    ranges.EntityData.SegmentPath = "ranges"
    ranges.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/" + ranges.EntityData.SegmentPath
    ranges.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ranges.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ranges.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ranges.EntityData.Children = types.NewOrderedMap()
    ranges.EntityData.Children.Append("range", types.YChild{"Range", nil})
    for i := range ranges.Range {
        ranges.EntityData.Children.Append(types.GetSegmentPath(ranges.Range[i]), types.YChild{"Range", ranges.Range[i]})
    }
    ranges.EntityData.Leafs = types.NewOrderedMap()

    ranges.EntityData.YListKeys = []string {}

    return &(ranges.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Ranges_Range
// Summarize inter-area routes matching
// prefix/length
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Ranges_Range struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IPv6 prefix format. The type is string with
    // pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Prefix interface{}

    // This attribute is a key. IPV6 prefix length. The type is interface{} with
    // range: 0..128.
    PrefixLength interface{}

    // Do not advertise address range. The type is bool. The default value is
    // false.
    NotAdvertise interface{}

    // Specified metric for this range. The type is interface{} with range:
    // 1..16777214.
    Cost interface{}
}

func (self *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Ranges_Range) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "range"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "ranges"
    self.EntityData.SegmentPath = "range" + types.AddKeyToken(self.Prefix, "prefix") + types.AddKeyToken(self.PrefixLength, "prefix-length")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/ranges/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", self.Prefix})
    self.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", self.PrefixLength})
    self.EntityData.Leafs.Append("not-advertise", types.YLeaf{"NotAdvertise", self.NotAdvertise})
    self.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", self.Cost})

    self.EntityData.YListKeys = []string {"Prefix", "PrefixLength"}

    return &(self.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Encryption
// Encrypt and authenticate OSPFv3 packets
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Encryption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec ESP authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Specify the encryption algorithm. The type is Ospfv3EncryptionAlgorithm.
    EncryptionAlgorithm interface{}

    // Encryption password. The type is string with pattern: b'(!.+)|([^!].+)'.
    EncryptionPassword interface{}

    // Use the NULL, MD5 or SHA1 algorithm. The type is Ospfv3AuthenticationType2.
    AuthenticationAlgorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    AuthenticationPassword interface{}
}

func (encryption *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Encryption) GetEntityData() *types.CommonEntityData {
    encryption.EntityData.YFilter = encryption.YFilter
    encryption.EntityData.YangName = "encryption"
    encryption.EntityData.BundleName = "cisco_ios_xr"
    encryption.EntityData.ParentYangName = "area-address"
    encryption.EntityData.SegmentPath = "encryption"
    encryption.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/" + encryption.EntityData.SegmentPath
    encryption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryption.EntityData.Children = types.NewOrderedMap()
    encryption.EntityData.Leafs = types.NewOrderedMap()
    encryption.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", encryption.Enable})
    encryption.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", encryption.Spi})
    encryption.EntityData.Leafs.Append("encryption-algorithm", types.YLeaf{"EncryptionAlgorithm", encryption.EncryptionAlgorithm})
    encryption.EntityData.Leafs.Append("encryption-password", types.YLeaf{"EncryptionPassword", encryption.EncryptionPassword})
    encryption.EntityData.Leafs.Append("authentication-algorithm", types.YLeaf{"AuthenticationAlgorithm", encryption.AuthenticationAlgorithm})
    encryption.EntityData.Leafs.Append("authentication-password", types.YLeaf{"AuthenticationPassword", encryption.AuthenticationPassword})

    encryption.EntityData.YListKeys = []string {}

    return &(encryption.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Nssa
// Specify area as a NSSA area.  Allowed only in
// non-backbone areas
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Nssa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // No redistribution into this NSSA area. The type is bool. The default value
    // is false.
    NoRedistribution interface{}

    // Originate Type 7 default into NSSA area. The type is bool. The default
    // value is false.
    DefaultInfoOriginate interface{}

    // Only valid with DefaultInfoOriginate. The type is interface{} with range:
    // 0..16777214.
    Metric interface{}

    // Only valid with DefaultInfoOriginate. The type is Ospfv3Metric.
    MetricType interface{}

    // Do not send summary LSA into NSSA. The type is interface{}.
    NoSummary interface{}
}

func (nssa *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Nssa) GetEntityData() *types.CommonEntityData {
    nssa.EntityData.YFilter = nssa.YFilter
    nssa.EntityData.YangName = "nssa"
    nssa.EntityData.BundleName = "cisco_ios_xr"
    nssa.EntityData.ParentYangName = "area-address"
    nssa.EntityData.SegmentPath = "nssa"
    nssa.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/" + nssa.EntityData.SegmentPath
    nssa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nssa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nssa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nssa.EntityData.Children = types.NewOrderedMap()
    nssa.EntityData.Leafs = types.NewOrderedMap()
    nssa.EntityData.Leafs.Append("no-redistribution", types.YLeaf{"NoRedistribution", nssa.NoRedistribution})
    nssa.EntityData.Leafs.Append("default-info-originate", types.YLeaf{"DefaultInfoOriginate", nssa.DefaultInfoOriginate})
    nssa.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", nssa.Metric})
    nssa.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", nssa.MetricType})
    nssa.EntityData.Leafs.Append("no-summary", types.YLeaf{"NoSummary", nssa.NoSummary})

    nssa.EntityData.YListKeys = []string {}

    return &(nssa.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_DatabaseFilter
// Database filter
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_DatabaseFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All.
    All Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_DatabaseFilter_All
}

func (databaseFilter *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_DatabaseFilter) GetEntityData() *types.CommonEntityData {
    databaseFilter.EntityData.YFilter = databaseFilter.YFilter
    databaseFilter.EntityData.YangName = "database-filter"
    databaseFilter.EntityData.BundleName = "cisco_ios_xr"
    databaseFilter.EntityData.ParentYangName = "area-address"
    databaseFilter.EntityData.SegmentPath = "database-filter"
    databaseFilter.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/" + databaseFilter.EntityData.SegmentPath
    databaseFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    databaseFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    databaseFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    databaseFilter.EntityData.Children = types.NewOrderedMap()
    databaseFilter.EntityData.Children.Append("all", types.YChild{"All", &databaseFilter.All})
    databaseFilter.EntityData.Leafs = types.NewOrderedMap()

    databaseFilter.EntityData.YListKeys = []string {}

    return &(databaseFilter.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_DatabaseFilter_All
// All
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_DatabaseFilter_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable or disable database-filter. The type is bool.
    Out interface{}
}

func (all *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_DatabaseFilter_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "database-filter"
    all.EntityData.SegmentPath = "all"
    all.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/database-filter/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Leafs = types.NewOrderedMap()
    all.EntityData.Leafs.Append("out", types.YLeaf{"Out", all.Out})

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_DistributeList
// Filter prefixes to/from RIB
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_DistributeList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter prefixes installed to RIB.
    In Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_DistributeList_In
}

func (distributeList *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_DistributeList) GetEntityData() *types.CommonEntityData {
    distributeList.EntityData.YFilter = distributeList.YFilter
    distributeList.EntityData.YangName = "distribute-list"
    distributeList.EntityData.BundleName = "cisco_ios_xr"
    distributeList.EntityData.ParentYangName = "area-address"
    distributeList.EntityData.SegmentPath = "distribute-list"
    distributeList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/" + distributeList.EntityData.SegmentPath
    distributeList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    distributeList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    distributeList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    distributeList.EntityData.Children = types.NewOrderedMap()
    distributeList.EntityData.Children.Append("in", types.YChild{"In", &distributeList.In})
    distributeList.EntityData.Leafs = types.NewOrderedMap()

    distributeList.EntityData.YListKeys = []string {}

    return &(distributeList.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_DistributeList_In
// Filter prefixes installed to RIB
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_DistributeList_In struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter prefixes based on an IPv6 prefix-list. The type is string.
    PrefixList interface{}
}

func (in *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_DistributeList_In) GetEntityData() *types.CommonEntityData {
    in.EntityData.YFilter = in.YFilter
    in.EntityData.YangName = "in"
    in.EntityData.BundleName = "cisco_ios_xr"
    in.EntityData.ParentYangName = "distribute-list"
    in.EntityData.SegmentPath = "in"
    in.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/distribute-list/" + in.EntityData.SegmentPath
    in.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    in.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    in.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    in.EntityData.Children = types.NewOrderedMap()
    in.EntityData.Leafs = types.NewOrderedMap()
    in.EntityData.Leafs.Append("prefix-list", types.YLeaf{"PrefixList", in.PrefixList})

    in.EntityData.YListKeys = []string {}

    return &(in.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces
// OSPFv3 interfaces
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPFv3 interface. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface.
    Interface []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface
}

func (interfaces *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "area-address"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/" + interfaces.EntityData.SegmentPath
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface
// OSPFv3 interface
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface to configure. The type is string with
    // pattern: b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Enable OSPFv3 interface. The type is interface{}.
    Enable interface{}

    // Interval after which a neighbor is declared dead (in seconds). The type is
    // interface{} with range: 1..65535. Units are second.
    DeadInterval interface{}

    // Enable/disable flood reduction. The type is bool.
    FloodReduction interface{}

    // Interface cost. The type is interface{} with range: 1..65535.
    Cost interface{}

    // Specify the transmit delay in seconds. The type is interface{} with range:
    // 1..65535. Units are second.
    TransmitDelay interface{}

    // Instance ID. The type is interface{} with range: 0..255.
    Instance interface{}

    // Enable/Disable MPLS LDP sync. The type is bool.
    LdpSync interface{}

    // Enable/disable ignoring of MTU in DBD packets. The type is bool.
    MtuIgnore interface{}

    // Specify the transmit interval in seconds. The type is interface{} with
    // range: 1..65535. Units are second.
    RetransmitInterval interface{}

    // Time between HELLO packets. The type is interface{} with range: 1..65535.
    // Units are second.
    HelloInterval interface{}

    // Enable/disable routing updates on an interface. The type is bool.
    Passive interface{}

    // Limit size of OSPFv3 packets. The type is interface{} with range:
    // 256..10000.
    PacketSize interface{}

    // Enable/disable prefix suppression on an interface. The type is bool.
    PrefixSuppression interface{}

    // Specify router priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // Specify network type. The type is Ospfv3Network.
    Network interface{}

    // Enable/disable demand circuit operation. The type is bool.
    DemandCircuit interface{}

    // Authenticate OSPFv3 packets.
    Authentication Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_Authentication

    // Specify a neighbor router.
    Neighbors Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_Neighbors

    // Encrypt and authenticate OSPFv3 packets.
    Encryption Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_Encryption

    // Configure BFD parameters.
    Bfd Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_Bfd

    // Database filter.
    DatabaseFilter Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_DatabaseFilter

    // Filter prefixes to/from RIB.
    DistributeList Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_DistributeList

    // Fast-reroute configuration.
    FastReroute Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute
}

func (self *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("authentication", types.YChild{"Authentication", &self.Authentication})
    self.EntityData.Children.Append("neighbors", types.YChild{"Neighbors", &self.Neighbors})
    self.EntityData.Children.Append("encryption", types.YChild{"Encryption", &self.Encryption})
    self.EntityData.Children.Append("bfd", types.YChild{"Bfd", &self.Bfd})
    self.EntityData.Children.Append("database-filter", types.YChild{"DatabaseFilter", &self.DatabaseFilter})
    self.EntityData.Children.Append("distribute-list", types.YChild{"DistributeList", &self.DistributeList})
    self.EntityData.Children.Append("fast-reroute", types.YChild{"FastReroute", &self.FastReroute})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", self.Enable})
    self.EntityData.Leafs.Append("dead-interval", types.YLeaf{"DeadInterval", self.DeadInterval})
    self.EntityData.Leafs.Append("flood-reduction", types.YLeaf{"FloodReduction", self.FloodReduction})
    self.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", self.Cost})
    self.EntityData.Leafs.Append("transmit-delay", types.YLeaf{"TransmitDelay", self.TransmitDelay})
    self.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", self.Instance})
    self.EntityData.Leafs.Append("ldp-sync", types.YLeaf{"LdpSync", self.LdpSync})
    self.EntityData.Leafs.Append("mtu-ignore", types.YLeaf{"MtuIgnore", self.MtuIgnore})
    self.EntityData.Leafs.Append("retransmit-interval", types.YLeaf{"RetransmitInterval", self.RetransmitInterval})
    self.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", self.HelloInterval})
    self.EntityData.Leafs.Append("passive", types.YLeaf{"Passive", self.Passive})
    self.EntityData.Leafs.Append("packet-size", types.YLeaf{"PacketSize", self.PacketSize})
    self.EntityData.Leafs.Append("prefix-suppression", types.YLeaf{"PrefixSuppression", self.PrefixSuppression})
    self.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", self.Priority})
    self.EntityData.Leafs.Append("network", types.YLeaf{"Network", self.Network})
    self.EntityData.Leafs.Append("demand-circuit", types.YLeaf{"DemandCircuit", self.DemandCircuit})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_Authentication
// Authenticate OSPFv3 packets
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec AH authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Use the MD5 or SHA1 algorithm. The type is Ospfv3Authentication.
    Algorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    Password interface{}
}

func (authentication *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "interface"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/interfaces/interface/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", authentication.Enable})
    authentication.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", authentication.Spi})
    authentication.EntityData.Leafs.Append("algorithm", types.YLeaf{"Algorithm", authentication.Algorithm})
    authentication.EntityData.Leafs.Append("password", types.YLeaf{"Password", authentication.Password})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_Neighbors
// Specify a neighbor router
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 address. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_Neighbors_Neighbor.
    Neighbor []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_Neighbors_Neighbor
}

func (neighbors *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "cisco_ios_xr"
    neighbors.EntityData.ParentYangName = "interface"
    neighbors.EntityData.SegmentPath = "neighbors"
    neighbors.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/interfaces/interface/" + neighbors.EntityData.SegmentPath
    neighbors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbors.EntityData.Children = types.NewOrderedMap()
    neighbors.EntityData.Children.Append("neighbor", types.YChild{"Neighbor", nil})
    for i := range neighbors.Neighbor {
        neighbors.EntityData.Children.Append(types.GetSegmentPath(neighbors.Neighbor[i]), types.YChild{"Neighbor", neighbors.Neighbor[i]})
    }
    neighbors.EntityData.Leafs = types.NewOrderedMap()

    neighbors.EntityData.YListKeys = []string {}

    return &(neighbors.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_Neighbors_Neighbor
// IPv6 address
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IPV6 address. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    NeighborAddress interface{}

    // OSPFv3 priority of non-broadcast neighbor. The type is interface{} with
    // range: 0..255.
    Priority interface{}

    // OSPFv3 dead-router polling interval (in seconds). The type is interface{}
    // with range: 0..65535. Units are second.
    PollInterval interface{}

    // OSPFv3 cost for point-to-multipoint neighbor. The type is interface{} with
    // range: 1..65535.
    Cost interface{}

    // Filter OSPFv3 LSA during synchronization and flooding for
    // point-to-multipoint neighbor. The type is bool.
    DatabaseFilter interface{}

    // Zone. The type is string.
    Zone interface{}
}

func (neighbor *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_Neighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "neighbors"
    neighbor.EntityData.SegmentPath = "neighbor" + types.AddKeyToken(neighbor.NeighborAddress, "neighbor-address")
    neighbor.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/interfaces/interface/neighbors/" + neighbor.EntityData.SegmentPath
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", neighbor.NeighborAddress})
    neighbor.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", neighbor.Priority})
    neighbor.EntityData.Leafs.Append("poll-interval", types.YLeaf{"PollInterval", neighbor.PollInterval})
    neighbor.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", neighbor.Cost})
    neighbor.EntityData.Leafs.Append("database-filter", types.YLeaf{"DatabaseFilter", neighbor.DatabaseFilter})
    neighbor.EntityData.Leafs.Append("zone", types.YLeaf{"Zone", neighbor.Zone})

    neighbor.EntityData.YListKeys = []string {"NeighborAddress"}

    return &(neighbor.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_Encryption
// Encrypt and authenticate OSPFv3 packets
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_Encryption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec ESP authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Specify the encryption algorithm. The type is Ospfv3EncryptionAlgorithm.
    EncryptionAlgorithm interface{}

    // Encryption password. The type is string with pattern: b'(!.+)|([^!].+)'.
    EncryptionPassword interface{}

    // Use the NULL, MD5 or SHA1 algorithm. The type is Ospfv3AuthenticationType2.
    AuthenticationAlgorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    AuthenticationPassword interface{}
}

func (encryption *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_Encryption) GetEntityData() *types.CommonEntityData {
    encryption.EntityData.YFilter = encryption.YFilter
    encryption.EntityData.YangName = "encryption"
    encryption.EntityData.BundleName = "cisco_ios_xr"
    encryption.EntityData.ParentYangName = "interface"
    encryption.EntityData.SegmentPath = "encryption"
    encryption.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/interfaces/interface/" + encryption.EntityData.SegmentPath
    encryption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryption.EntityData.Children = types.NewOrderedMap()
    encryption.EntityData.Leafs = types.NewOrderedMap()
    encryption.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", encryption.Enable})
    encryption.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", encryption.Spi})
    encryption.EntityData.Leafs.Append("encryption-algorithm", types.YLeaf{"EncryptionAlgorithm", encryption.EncryptionAlgorithm})
    encryption.EntityData.Leafs.Append("encryption-password", types.YLeaf{"EncryptionPassword", encryption.EncryptionPassword})
    encryption.EntityData.Leafs.Append("authentication-algorithm", types.YLeaf{"AuthenticationAlgorithm", encryption.AuthenticationAlgorithm})
    encryption.EntityData.Leafs.Append("authentication-password", types.YLeaf{"AuthenticationPassword", encryption.AuthenticationPassword})

    encryption.EntityData.YListKeys = []string {}

    return &(encryption.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_Bfd
// Configure BFD parameters
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Hello interval in milli-seconds. The type is interface{} with range:
    // 3..30000. Units are millisecond.
    Interval interface{}

    // Enable or disable BFD fast detection. The type is Ospfv3bfdEnableMode.
    FastDetectMode interface{}

    // Detect multiplier. The type is interface{} with range: 2..50.
    DetectionMultiplier interface{}
}

func (bfd *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "interface"
    bfd.EntityData.SegmentPath = "bfd"
    bfd.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/interfaces/interface/" + bfd.EntityData.SegmentPath
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = types.NewOrderedMap()
    bfd.EntityData.Leafs = types.NewOrderedMap()
    bfd.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", bfd.Interval})
    bfd.EntityData.Leafs.Append("fast-detect-mode", types.YLeaf{"FastDetectMode", bfd.FastDetectMode})
    bfd.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", bfd.DetectionMultiplier})

    bfd.EntityData.YListKeys = []string {}

    return &(bfd.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_DatabaseFilter
// Database filter
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_DatabaseFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All.
    All Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_DatabaseFilter_All
}

func (databaseFilter *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_DatabaseFilter) GetEntityData() *types.CommonEntityData {
    databaseFilter.EntityData.YFilter = databaseFilter.YFilter
    databaseFilter.EntityData.YangName = "database-filter"
    databaseFilter.EntityData.BundleName = "cisco_ios_xr"
    databaseFilter.EntityData.ParentYangName = "interface"
    databaseFilter.EntityData.SegmentPath = "database-filter"
    databaseFilter.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/interfaces/interface/" + databaseFilter.EntityData.SegmentPath
    databaseFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    databaseFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    databaseFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    databaseFilter.EntityData.Children = types.NewOrderedMap()
    databaseFilter.EntityData.Children.Append("all", types.YChild{"All", &databaseFilter.All})
    databaseFilter.EntityData.Leafs = types.NewOrderedMap()

    databaseFilter.EntityData.YListKeys = []string {}

    return &(databaseFilter.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_DatabaseFilter_All
// All
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_DatabaseFilter_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable or disable database-filter. The type is bool.
    Out interface{}
}

func (all *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_DatabaseFilter_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "database-filter"
    all.EntityData.SegmentPath = "all"
    all.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/interfaces/interface/database-filter/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Leafs = types.NewOrderedMap()
    all.EntityData.Leafs.Append("out", types.YLeaf{"Out", all.Out})

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_DistributeList
// Filter prefixes to/from RIB
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_DistributeList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter prefixes installed to RIB.
    In Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_DistributeList_In
}

func (distributeList *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_DistributeList) GetEntityData() *types.CommonEntityData {
    distributeList.EntityData.YFilter = distributeList.YFilter
    distributeList.EntityData.YangName = "distribute-list"
    distributeList.EntityData.BundleName = "cisco_ios_xr"
    distributeList.EntityData.ParentYangName = "interface"
    distributeList.EntityData.SegmentPath = "distribute-list"
    distributeList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/interfaces/interface/" + distributeList.EntityData.SegmentPath
    distributeList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    distributeList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    distributeList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    distributeList.EntityData.Children = types.NewOrderedMap()
    distributeList.EntityData.Children.Append("in", types.YChild{"In", &distributeList.In})
    distributeList.EntityData.Leafs = types.NewOrderedMap()

    distributeList.EntityData.YListKeys = []string {}

    return &(distributeList.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_DistributeList_In
// Filter prefixes installed to RIB
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_DistributeList_In struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter prefixes based on an IPv6 prefix-list. The type is string.
    PrefixList interface{}
}

func (in *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_DistributeList_In) GetEntityData() *types.CommonEntityData {
    in.EntityData.YFilter = in.YFilter
    in.EntityData.YangName = "in"
    in.EntityData.BundleName = "cisco_ios_xr"
    in.EntityData.ParentYangName = "distribute-list"
    in.EntityData.SegmentPath = "in"
    in.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/interfaces/interface/distribute-list/" + in.EntityData.SegmentPath
    in.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    in.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    in.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    in.EntityData.Children = types.NewOrderedMap()
    in.EntityData.Leafs = types.NewOrderedMap()
    in.EntityData.Leafs.Append("prefix-list", types.YLeaf{"PrefixList", in.PrefixList})

    in.EntityData.YListKeys = []string {}

    return &(in.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute
// Fast-reroute configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable Fast-reroute per-link or per-prefix. The type is
    // Ospfv3FastReroute.
    FastRerouteEnable interface{}

    // Fast-reroute per-link configuration.
    PerLink Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink

    // Fast-reroute per-link configuration.
    PerPrefix Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix
}

func (fastReroute *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute) GetEntityData() *types.CommonEntityData {
    fastReroute.EntityData.YFilter = fastReroute.YFilter
    fastReroute.EntityData.YangName = "fast-reroute"
    fastReroute.EntityData.BundleName = "cisco_ios_xr"
    fastReroute.EntityData.ParentYangName = "interface"
    fastReroute.EntityData.SegmentPath = "fast-reroute"
    fastReroute.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/interfaces/interface/" + fastReroute.EntityData.SegmentPath
    fastReroute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fastReroute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fastReroute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fastReroute.EntityData.Children = types.NewOrderedMap()
    fastReroute.EntityData.Children.Append("per-link", types.YChild{"PerLink", &fastReroute.PerLink})
    fastReroute.EntityData.Children.Append("per-prefix", types.YChild{"PerPrefix", &fastReroute.PerPrefix})
    fastReroute.EntityData.Leafs = types.NewOrderedMap()
    fastReroute.EntityData.Leafs.Append("fast-reroute-enable", types.YLeaf{"FastRerouteEnable", fastReroute.FastRerouteEnable})

    fastReroute.EntityData.YListKeys = []string {}

    return &(fastReroute.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink
// Fast-reroute per-link configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Use only interfaces on the candidate list as a backup path. The type is
    // bool. The default value is false.
    FastRerouteUseCandidateOnly interface{}

    // Fast-reroute per-link/per-prefix candidate interface configuration.
    CandidateInterfaces Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces

    // Fast-reroute per-link/per-prefix exclude interface configuration.
    ExcludeInterfaces Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces
}

func (perLink *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink) GetEntityData() *types.CommonEntityData {
    perLink.EntityData.YFilter = perLink.YFilter
    perLink.EntityData.YangName = "per-link"
    perLink.EntityData.BundleName = "cisco_ios_xr"
    perLink.EntityData.ParentYangName = "fast-reroute"
    perLink.EntityData.SegmentPath = "per-link"
    perLink.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/interfaces/interface/fast-reroute/" + perLink.EntityData.SegmentPath
    perLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    perLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    perLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    perLink.EntityData.Children = types.NewOrderedMap()
    perLink.EntityData.Children.Append("candidate-interfaces", types.YChild{"CandidateInterfaces", &perLink.CandidateInterfaces})
    perLink.EntityData.Children.Append("exclude-interfaces", types.YChild{"ExcludeInterfaces", &perLink.ExcludeInterfaces})
    perLink.EntityData.Leafs = types.NewOrderedMap()
    perLink.EntityData.Leafs.Append("fast-reroute-use-candidate-only", types.YLeaf{"FastRerouteUseCandidateOnly", perLink.FastRerouteUseCandidateOnly})

    perLink.EntityData.YListKeys = []string {}

    return &(perLink.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces
// Fast-reroute per-link/per-prefix candidate
// interface configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Candidate backup interface. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces_CandidateInterface.
    CandidateInterface []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces_CandidateInterface
}

func (candidateInterfaces *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces) GetEntityData() *types.CommonEntityData {
    candidateInterfaces.EntityData.YFilter = candidateInterfaces.YFilter
    candidateInterfaces.EntityData.YangName = "candidate-interfaces"
    candidateInterfaces.EntityData.BundleName = "cisco_ios_xr"
    candidateInterfaces.EntityData.ParentYangName = "per-link"
    candidateInterfaces.EntityData.SegmentPath = "candidate-interfaces"
    candidateInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/interfaces/interface/fast-reroute/per-link/" + candidateInterfaces.EntityData.SegmentPath
    candidateInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterfaces.EntityData.Children = types.NewOrderedMap()
    candidateInterfaces.EntityData.Children.Append("candidate-interface", types.YChild{"CandidateInterface", nil})
    for i := range candidateInterfaces.CandidateInterface {
        candidateInterfaces.EntityData.Children.Append(types.GetSegmentPath(candidateInterfaces.CandidateInterface[i]), types.YChild{"CandidateInterface", candidateInterfaces.CandidateInterface[i]})
    }
    candidateInterfaces.EntityData.Leafs = types.NewOrderedMap()

    candidateInterfaces.EntityData.YListKeys = []string {}

    return &(candidateInterfaces.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces_CandidateInterface
// Candidate backup interface
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces_CandidateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (candidateInterface *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces_CandidateInterface) GetEntityData() *types.CommonEntityData {
    candidateInterface.EntityData.YFilter = candidateInterface.YFilter
    candidateInterface.EntityData.YangName = "candidate-interface"
    candidateInterface.EntityData.BundleName = "cisco_ios_xr"
    candidateInterface.EntityData.ParentYangName = "candidate-interfaces"
    candidateInterface.EntityData.SegmentPath = "candidate-interface" + types.AddKeyToken(candidateInterface.InterfaceName, "interface-name")
    candidateInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/interfaces/interface/fast-reroute/per-link/candidate-interfaces/" + candidateInterface.EntityData.SegmentPath
    candidateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterface.EntityData.Children = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", candidateInterface.InterfaceName})

    candidateInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(candidateInterface.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces
// Fast-reroute per-link/per-prefix exclude
// interface configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exclude an interface from becoming a backup. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface.
    ExcludeInterface []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface
}

func (excludeInterfaces *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces) GetEntityData() *types.CommonEntityData {
    excludeInterfaces.EntityData.YFilter = excludeInterfaces.YFilter
    excludeInterfaces.EntityData.YangName = "exclude-interfaces"
    excludeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    excludeInterfaces.EntityData.ParentYangName = "per-link"
    excludeInterfaces.EntityData.SegmentPath = "exclude-interfaces"
    excludeInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/interfaces/interface/fast-reroute/per-link/" + excludeInterfaces.EntityData.SegmentPath
    excludeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterfaces.EntityData.Children = types.NewOrderedMap()
    excludeInterfaces.EntityData.Children.Append("exclude-interface", types.YChild{"ExcludeInterface", nil})
    for i := range excludeInterfaces.ExcludeInterface {
        excludeInterfaces.EntityData.Children.Append(types.GetSegmentPath(excludeInterfaces.ExcludeInterface[i]), types.YChild{"ExcludeInterface", excludeInterfaces.ExcludeInterface[i]})
    }
    excludeInterfaces.EntityData.Leafs = types.NewOrderedMap()

    excludeInterfaces.EntityData.YListKeys = []string {}

    return &(excludeInterfaces.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface
// Exclude an interface from becoming a backup
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (excludeInterface *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface) GetEntityData() *types.CommonEntityData {
    excludeInterface.EntityData.YFilter = excludeInterface.YFilter
    excludeInterface.EntityData.YangName = "exclude-interface"
    excludeInterface.EntityData.BundleName = "cisco_ios_xr"
    excludeInterface.EntityData.ParentYangName = "exclude-interfaces"
    excludeInterface.EntityData.SegmentPath = "exclude-interface" + types.AddKeyToken(excludeInterface.InterfaceName, "interface-name")
    excludeInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/interfaces/interface/fast-reroute/per-link/exclude-interfaces/" + excludeInterface.EntityData.SegmentPath
    excludeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterface.EntityData.Children = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", excludeInterface.InterfaceName})

    excludeInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(excludeInterface.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix
// Fast-reroute per-link configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Use only interfaces on the candidate list as a backup path. The type is
    // bool. The default value is false.
    FastRerouteUseCandidateOnly interface{}

    // Fast-reroute per-link/per-prefix candidate interface configuration.
    CandidateInterfaces Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces

    // Fast-reroute per-link/per-prefix exclude interface configuration.
    ExcludeInterfaces Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces
}

func (perPrefix *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix) GetEntityData() *types.CommonEntityData {
    perPrefix.EntityData.YFilter = perPrefix.YFilter
    perPrefix.EntityData.YangName = "per-prefix"
    perPrefix.EntityData.BundleName = "cisco_ios_xr"
    perPrefix.EntityData.ParentYangName = "fast-reroute"
    perPrefix.EntityData.SegmentPath = "per-prefix"
    perPrefix.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/interfaces/interface/fast-reroute/" + perPrefix.EntityData.SegmentPath
    perPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    perPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    perPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    perPrefix.EntityData.Children = types.NewOrderedMap()
    perPrefix.EntityData.Children.Append("candidate-interfaces", types.YChild{"CandidateInterfaces", &perPrefix.CandidateInterfaces})
    perPrefix.EntityData.Children.Append("exclude-interfaces", types.YChild{"ExcludeInterfaces", &perPrefix.ExcludeInterfaces})
    perPrefix.EntityData.Leafs = types.NewOrderedMap()
    perPrefix.EntityData.Leafs.Append("fast-reroute-use-candidate-only", types.YLeaf{"FastRerouteUseCandidateOnly", perPrefix.FastRerouteUseCandidateOnly})

    perPrefix.EntityData.YListKeys = []string {}

    return &(perPrefix.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces
// Fast-reroute per-link/per-prefix candidate
// interface configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Candidate backup interface. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface.
    CandidateInterface []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface
}

func (candidateInterfaces *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces) GetEntityData() *types.CommonEntityData {
    candidateInterfaces.EntityData.YFilter = candidateInterfaces.YFilter
    candidateInterfaces.EntityData.YangName = "candidate-interfaces"
    candidateInterfaces.EntityData.BundleName = "cisco_ios_xr"
    candidateInterfaces.EntityData.ParentYangName = "per-prefix"
    candidateInterfaces.EntityData.SegmentPath = "candidate-interfaces"
    candidateInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/interfaces/interface/fast-reroute/per-prefix/" + candidateInterfaces.EntityData.SegmentPath
    candidateInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterfaces.EntityData.Children = types.NewOrderedMap()
    candidateInterfaces.EntityData.Children.Append("candidate-interface", types.YChild{"CandidateInterface", nil})
    for i := range candidateInterfaces.CandidateInterface {
        candidateInterfaces.EntityData.Children.Append(types.GetSegmentPath(candidateInterfaces.CandidateInterface[i]), types.YChild{"CandidateInterface", candidateInterfaces.CandidateInterface[i]})
    }
    candidateInterfaces.EntityData.Leafs = types.NewOrderedMap()

    candidateInterfaces.EntityData.YListKeys = []string {}

    return &(candidateInterfaces.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface
// Candidate backup interface
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (candidateInterface *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface) GetEntityData() *types.CommonEntityData {
    candidateInterface.EntityData.YFilter = candidateInterface.YFilter
    candidateInterface.EntityData.YangName = "candidate-interface"
    candidateInterface.EntityData.BundleName = "cisco_ios_xr"
    candidateInterface.EntityData.ParentYangName = "candidate-interfaces"
    candidateInterface.EntityData.SegmentPath = "candidate-interface" + types.AddKeyToken(candidateInterface.InterfaceName, "interface-name")
    candidateInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/interfaces/interface/fast-reroute/per-prefix/candidate-interfaces/" + candidateInterface.EntityData.SegmentPath
    candidateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterface.EntityData.Children = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", candidateInterface.InterfaceName})

    candidateInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(candidateInterface.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces
// Fast-reroute per-link/per-prefix exclude
// interface configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exclude an interface from becoming a backup. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface.
    ExcludeInterface []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface
}

func (excludeInterfaces *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces) GetEntityData() *types.CommonEntityData {
    excludeInterfaces.EntityData.YFilter = excludeInterfaces.YFilter
    excludeInterfaces.EntityData.YangName = "exclude-interfaces"
    excludeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    excludeInterfaces.EntityData.ParentYangName = "per-prefix"
    excludeInterfaces.EntityData.SegmentPath = "exclude-interfaces"
    excludeInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/interfaces/interface/fast-reroute/per-prefix/" + excludeInterfaces.EntityData.SegmentPath
    excludeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterfaces.EntityData.Children = types.NewOrderedMap()
    excludeInterfaces.EntityData.Children.Append("exclude-interface", types.YChild{"ExcludeInterface", nil})
    for i := range excludeInterfaces.ExcludeInterface {
        excludeInterfaces.EntityData.Children.Append(types.GetSegmentPath(excludeInterfaces.ExcludeInterface[i]), types.YChild{"ExcludeInterface", excludeInterfaces.ExcludeInterface[i]})
    }
    excludeInterfaces.EntityData.Leafs = types.NewOrderedMap()

    excludeInterfaces.EntityData.YListKeys = []string {}

    return &(excludeInterfaces.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface
// Exclude an interface from becoming a backup
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (excludeInterface *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface) GetEntityData() *types.CommonEntityData {
    excludeInterface.EntityData.YFilter = excludeInterface.YFilter
    excludeInterface.EntityData.YangName = "exclude-interface"
    excludeInterface.EntityData.BundleName = "cisco_ios_xr"
    excludeInterface.EntityData.ParentYangName = "exclude-interfaces"
    excludeInterface.EntityData.SegmentPath = "exclude-interface" + types.AddKeyToken(excludeInterface.InterfaceName, "interface-name")
    excludeInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/interfaces/interface/fast-reroute/per-prefix/exclude-interfaces/" + excludeInterface.EntityData.SegmentPath
    excludeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterface.EntityData.Children = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", excludeInterface.InterfaceName})

    excludeInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(excludeInterface.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope
// Area Scope Configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fast-reroute configuration.
    FastReroute Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute
}

func (areaScope *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope) GetEntityData() *types.CommonEntityData {
    areaScope.EntityData.YFilter = areaScope.YFilter
    areaScope.EntityData.YangName = "area-scope"
    areaScope.EntityData.BundleName = "cisco_ios_xr"
    areaScope.EntityData.ParentYangName = "area-address"
    areaScope.EntityData.SegmentPath = "area-scope"
    areaScope.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/" + areaScope.EntityData.SegmentPath
    areaScope.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    areaScope.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    areaScope.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    areaScope.EntityData.Children = types.NewOrderedMap()
    areaScope.EntityData.Children.Append("fast-reroute", types.YChild{"FastReroute", &areaScope.FastReroute})
    areaScope.EntityData.Leafs = types.NewOrderedMap()

    areaScope.EntityData.YListKeys = []string {}

    return &(areaScope.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute
// Fast-reroute configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable Fast-reroute per-link or per-prefix. The type is
    // Ospfv3FastReroute.
    FastRerouteEnable interface{}

    // Fast-reroute per-link configuration.
    PerLink Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink

    // Fast-reroute per-link configuration.
    PerPrefix Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix
}

func (fastReroute *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute) GetEntityData() *types.CommonEntityData {
    fastReroute.EntityData.YFilter = fastReroute.YFilter
    fastReroute.EntityData.YangName = "fast-reroute"
    fastReroute.EntityData.BundleName = "cisco_ios_xr"
    fastReroute.EntityData.ParentYangName = "area-scope"
    fastReroute.EntityData.SegmentPath = "fast-reroute"
    fastReroute.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/area-scope/" + fastReroute.EntityData.SegmentPath
    fastReroute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fastReroute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fastReroute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fastReroute.EntityData.Children = types.NewOrderedMap()
    fastReroute.EntityData.Children.Append("per-link", types.YChild{"PerLink", &fastReroute.PerLink})
    fastReroute.EntityData.Children.Append("per-prefix", types.YChild{"PerPrefix", &fastReroute.PerPrefix})
    fastReroute.EntityData.Leafs = types.NewOrderedMap()
    fastReroute.EntityData.Leafs.Append("fast-reroute-enable", types.YLeaf{"FastRerouteEnable", fastReroute.FastRerouteEnable})

    fastReroute.EntityData.YListKeys = []string {}

    return &(fastReroute.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink
// Fast-reroute per-link configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Use only interfaces on the candidate list as a backup path. The type is
    // bool. The default value is false.
    FastRerouteUseCandidateOnly interface{}

    // Fast-reroute per-link/per-prefix candidate interface configuration.
    CandidateInterfaces Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_CandidateInterfaces

    // Fast-reroute per-link/per-prefix exclude interface configuration.
    ExcludeInterfaces Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_ExcludeInterfaces
}

func (perLink *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink) GetEntityData() *types.CommonEntityData {
    perLink.EntityData.YFilter = perLink.YFilter
    perLink.EntityData.YangName = "per-link"
    perLink.EntityData.BundleName = "cisco_ios_xr"
    perLink.EntityData.ParentYangName = "fast-reroute"
    perLink.EntityData.SegmentPath = "per-link"
    perLink.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/area-scope/fast-reroute/" + perLink.EntityData.SegmentPath
    perLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    perLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    perLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    perLink.EntityData.Children = types.NewOrderedMap()
    perLink.EntityData.Children.Append("candidate-interfaces", types.YChild{"CandidateInterfaces", &perLink.CandidateInterfaces})
    perLink.EntityData.Children.Append("exclude-interfaces", types.YChild{"ExcludeInterfaces", &perLink.ExcludeInterfaces})
    perLink.EntityData.Leafs = types.NewOrderedMap()
    perLink.EntityData.Leafs.Append("fast-reroute-use-candidate-only", types.YLeaf{"FastRerouteUseCandidateOnly", perLink.FastRerouteUseCandidateOnly})

    perLink.EntityData.YListKeys = []string {}

    return &(perLink.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_CandidateInterfaces
// Fast-reroute per-link/per-prefix candidate
// interface configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_CandidateInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Candidate backup interface. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface.
    CandidateInterface []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface
}

func (candidateInterfaces *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_CandidateInterfaces) GetEntityData() *types.CommonEntityData {
    candidateInterfaces.EntityData.YFilter = candidateInterfaces.YFilter
    candidateInterfaces.EntityData.YangName = "candidate-interfaces"
    candidateInterfaces.EntityData.BundleName = "cisco_ios_xr"
    candidateInterfaces.EntityData.ParentYangName = "per-link"
    candidateInterfaces.EntityData.SegmentPath = "candidate-interfaces"
    candidateInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/area-scope/fast-reroute/per-link/" + candidateInterfaces.EntityData.SegmentPath
    candidateInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterfaces.EntityData.Children = types.NewOrderedMap()
    candidateInterfaces.EntityData.Children.Append("candidate-interface", types.YChild{"CandidateInterface", nil})
    for i := range candidateInterfaces.CandidateInterface {
        candidateInterfaces.EntityData.Children.Append(types.GetSegmentPath(candidateInterfaces.CandidateInterface[i]), types.YChild{"CandidateInterface", candidateInterfaces.CandidateInterface[i]})
    }
    candidateInterfaces.EntityData.Leafs = types.NewOrderedMap()

    candidateInterfaces.EntityData.YListKeys = []string {}

    return &(candidateInterfaces.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface
// Candidate backup interface
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (candidateInterface *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface) GetEntityData() *types.CommonEntityData {
    candidateInterface.EntityData.YFilter = candidateInterface.YFilter
    candidateInterface.EntityData.YangName = "candidate-interface"
    candidateInterface.EntityData.BundleName = "cisco_ios_xr"
    candidateInterface.EntityData.ParentYangName = "candidate-interfaces"
    candidateInterface.EntityData.SegmentPath = "candidate-interface" + types.AddKeyToken(candidateInterface.InterfaceName, "interface-name")
    candidateInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/area-scope/fast-reroute/per-link/candidate-interfaces/" + candidateInterface.EntityData.SegmentPath
    candidateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterface.EntityData.Children = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", candidateInterface.InterfaceName})

    candidateInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(candidateInterface.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_ExcludeInterfaces
// Fast-reroute per-link/per-prefix exclude
// interface configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_ExcludeInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exclude an interface from becoming a backup. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface.
    ExcludeInterface []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface
}

func (excludeInterfaces *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_ExcludeInterfaces) GetEntityData() *types.CommonEntityData {
    excludeInterfaces.EntityData.YFilter = excludeInterfaces.YFilter
    excludeInterfaces.EntityData.YangName = "exclude-interfaces"
    excludeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    excludeInterfaces.EntityData.ParentYangName = "per-link"
    excludeInterfaces.EntityData.SegmentPath = "exclude-interfaces"
    excludeInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/area-scope/fast-reroute/per-link/" + excludeInterfaces.EntityData.SegmentPath
    excludeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterfaces.EntityData.Children = types.NewOrderedMap()
    excludeInterfaces.EntityData.Children.Append("exclude-interface", types.YChild{"ExcludeInterface", nil})
    for i := range excludeInterfaces.ExcludeInterface {
        excludeInterfaces.EntityData.Children.Append(types.GetSegmentPath(excludeInterfaces.ExcludeInterface[i]), types.YChild{"ExcludeInterface", excludeInterfaces.ExcludeInterface[i]})
    }
    excludeInterfaces.EntityData.Leafs = types.NewOrderedMap()

    excludeInterfaces.EntityData.YListKeys = []string {}

    return &(excludeInterfaces.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface
// Exclude an interface from becoming a backup
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (excludeInterface *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface) GetEntityData() *types.CommonEntityData {
    excludeInterface.EntityData.YFilter = excludeInterface.YFilter
    excludeInterface.EntityData.YangName = "exclude-interface"
    excludeInterface.EntityData.BundleName = "cisco_ios_xr"
    excludeInterface.EntityData.ParentYangName = "exclude-interfaces"
    excludeInterface.EntityData.SegmentPath = "exclude-interface" + types.AddKeyToken(excludeInterface.InterfaceName, "interface-name")
    excludeInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/area-scope/fast-reroute/per-link/exclude-interfaces/" + excludeInterface.EntityData.SegmentPath
    excludeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterface.EntityData.Children = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", excludeInterface.InterfaceName})

    excludeInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(excludeInterface.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix
// Fast-reroute per-link configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Use only interfaces on the candidate list as a backup path. The type is
    // bool. The default value is false.
    FastRerouteUseCandidateOnly interface{}

    // Fast-reroute per-link/per-prefix candidate interface configuration.
    CandidateInterfaces Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_CandidateInterfaces

    // Fast-reroute per-link/per-prefix exclude interface configuration.
    ExcludeInterfaces Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces
}

func (perPrefix *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix) GetEntityData() *types.CommonEntityData {
    perPrefix.EntityData.YFilter = perPrefix.YFilter
    perPrefix.EntityData.YangName = "per-prefix"
    perPrefix.EntityData.BundleName = "cisco_ios_xr"
    perPrefix.EntityData.ParentYangName = "fast-reroute"
    perPrefix.EntityData.SegmentPath = "per-prefix"
    perPrefix.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/area-scope/fast-reroute/" + perPrefix.EntityData.SegmentPath
    perPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    perPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    perPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    perPrefix.EntityData.Children = types.NewOrderedMap()
    perPrefix.EntityData.Children.Append("candidate-interfaces", types.YChild{"CandidateInterfaces", &perPrefix.CandidateInterfaces})
    perPrefix.EntityData.Children.Append("exclude-interfaces", types.YChild{"ExcludeInterfaces", &perPrefix.ExcludeInterfaces})
    perPrefix.EntityData.Leafs = types.NewOrderedMap()
    perPrefix.EntityData.Leafs.Append("fast-reroute-use-candidate-only", types.YLeaf{"FastRerouteUseCandidateOnly", perPrefix.FastRerouteUseCandidateOnly})

    perPrefix.EntityData.YListKeys = []string {}

    return &(perPrefix.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_CandidateInterfaces
// Fast-reroute per-link/per-prefix candidate
// interface configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_CandidateInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Candidate backup interface. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface.
    CandidateInterface []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface
}

func (candidateInterfaces *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_CandidateInterfaces) GetEntityData() *types.CommonEntityData {
    candidateInterfaces.EntityData.YFilter = candidateInterfaces.YFilter
    candidateInterfaces.EntityData.YangName = "candidate-interfaces"
    candidateInterfaces.EntityData.BundleName = "cisco_ios_xr"
    candidateInterfaces.EntityData.ParentYangName = "per-prefix"
    candidateInterfaces.EntityData.SegmentPath = "candidate-interfaces"
    candidateInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/area-scope/fast-reroute/per-prefix/" + candidateInterfaces.EntityData.SegmentPath
    candidateInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterfaces.EntityData.Children = types.NewOrderedMap()
    candidateInterfaces.EntityData.Children.Append("candidate-interface", types.YChild{"CandidateInterface", nil})
    for i := range candidateInterfaces.CandidateInterface {
        candidateInterfaces.EntityData.Children.Append(types.GetSegmentPath(candidateInterfaces.CandidateInterface[i]), types.YChild{"CandidateInterface", candidateInterfaces.CandidateInterface[i]})
    }
    candidateInterfaces.EntityData.Leafs = types.NewOrderedMap()

    candidateInterfaces.EntityData.YListKeys = []string {}

    return &(candidateInterfaces.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface
// Candidate backup interface
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (candidateInterface *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface) GetEntityData() *types.CommonEntityData {
    candidateInterface.EntityData.YFilter = candidateInterface.YFilter
    candidateInterface.EntityData.YangName = "candidate-interface"
    candidateInterface.EntityData.BundleName = "cisco_ios_xr"
    candidateInterface.EntityData.ParentYangName = "candidate-interfaces"
    candidateInterface.EntityData.SegmentPath = "candidate-interface" + types.AddKeyToken(candidateInterface.InterfaceName, "interface-name")
    candidateInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/area-scope/fast-reroute/per-prefix/candidate-interfaces/" + candidateInterface.EntityData.SegmentPath
    candidateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterface.EntityData.Children = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", candidateInterface.InterfaceName})

    candidateInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(candidateInterface.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces
// Fast-reroute per-link/per-prefix exclude
// interface configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exclude an interface from becoming a backup. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface.
    ExcludeInterface []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface
}

func (excludeInterfaces *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces) GetEntityData() *types.CommonEntityData {
    excludeInterfaces.EntityData.YFilter = excludeInterfaces.YFilter
    excludeInterfaces.EntityData.YangName = "exclude-interfaces"
    excludeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    excludeInterfaces.EntityData.ParentYangName = "per-prefix"
    excludeInterfaces.EntityData.SegmentPath = "exclude-interfaces"
    excludeInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/area-scope/fast-reroute/per-prefix/" + excludeInterfaces.EntityData.SegmentPath
    excludeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterfaces.EntityData.Children = types.NewOrderedMap()
    excludeInterfaces.EntityData.Children.Append("exclude-interface", types.YChild{"ExcludeInterface", nil})
    for i := range excludeInterfaces.ExcludeInterface {
        excludeInterfaces.EntityData.Children.Append(types.GetSegmentPath(excludeInterfaces.ExcludeInterface[i]), types.YChild{"ExcludeInterface", excludeInterfaces.ExcludeInterface[i]})
    }
    excludeInterfaces.EntityData.Leafs = types.NewOrderedMap()

    excludeInterfaces.EntityData.YListKeys = []string {}

    return &(excludeInterfaces.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface
// Exclude an interface from becoming a backup
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (excludeInterface *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface) GetEntityData() *types.CommonEntityData {
    excludeInterface.EntityData.YFilter = excludeInterface.YFilter
    excludeInterface.EntityData.YangName = "exclude-interface"
    excludeInterface.EntityData.BundleName = "cisco_ios_xr"
    excludeInterface.EntityData.ParentYangName = "exclude-interfaces"
    excludeInterface.EntityData.SegmentPath = "exclude-interface" + types.AddKeyToken(excludeInterface.InterfaceName, "interface-name")
    excludeInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/area-scope/fast-reroute/per-prefix/exclude-interfaces/" + excludeInterface.EntityData.SegmentPath
    excludeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterface.EntityData.Children = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", excludeInterface.InterfaceName})

    excludeInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(excludeInterface.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_ShamLinks
// Sham Link sub-mode
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_ShamLinks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ShamLink local and remote endpoints. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_ShamLinks_ShamLink.
    ShamLink []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_ShamLinks_ShamLink
}

func (shamLinks *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_ShamLinks) GetEntityData() *types.CommonEntityData {
    shamLinks.EntityData.YFilter = shamLinks.YFilter
    shamLinks.EntityData.YangName = "sham-links"
    shamLinks.EntityData.BundleName = "cisco_ios_xr"
    shamLinks.EntityData.ParentYangName = "area-address"
    shamLinks.EntityData.SegmentPath = "sham-links"
    shamLinks.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/" + shamLinks.EntityData.SegmentPath
    shamLinks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shamLinks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shamLinks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shamLinks.EntityData.Children = types.NewOrderedMap()
    shamLinks.EntityData.Children.Append("sham-link", types.YChild{"ShamLink", nil})
    for i := range shamLinks.ShamLink {
        shamLinks.EntityData.Children.Append(types.GetSegmentPath(shamLinks.ShamLink[i]), types.YChild{"ShamLink", shamLinks.ShamLink[i]})
    }
    shamLinks.EntityData.Leafs = types.NewOrderedMap()

    shamLinks.EntityData.YListKeys = []string {}

    return &(shamLinks.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_ShamLinks_ShamLink
// ShamLink local and remote endpoints
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_ShamLinks_ShamLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Local sham-link endpoint. The type is string with
    // pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // This attribute is a key. Remote sham-link endpoint. The type is string with
    // pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    DestinationAddress interface{}

    // Enable sham link. The type is interface{}.
    Enable interface{}

    // Time between HELLO packets. The type is interface{} with range: 1..65535.
    // Units are second.
    HelloInterval interface{}

    // Interval after which a neighbor is declared dead (in seconds). The type is
    // interface{} with range: 1..65535. Units are second.
    DeadInterval interface{}

    // Specify the transmit interval in seconds. The type is interface{} with
    // range: 1..65535. Units are second.
    RetransmitInterval interface{}

    // Specify the transmit delay in seconds. The type is interface{} with range:
    // 1..65535. Units are second.
    TransmitDelay interface{}

    // Authenticate OSPFv3 packets.
    Authentication Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_ShamLinks_ShamLink_Authentication

    // Encrypt and authenticate OSPFv3 packets.
    Encryption Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_ShamLinks_ShamLink_Encryption
}

func (shamLink *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_ShamLinks_ShamLink) GetEntityData() *types.CommonEntityData {
    shamLink.EntityData.YFilter = shamLink.YFilter
    shamLink.EntityData.YangName = "sham-link"
    shamLink.EntityData.BundleName = "cisco_ios_xr"
    shamLink.EntityData.ParentYangName = "sham-links"
    shamLink.EntityData.SegmentPath = "sham-link" + types.AddKeyToken(shamLink.SourceAddress, "source-address") + types.AddKeyToken(shamLink.DestinationAddress, "destination-address")
    shamLink.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/sham-links/" + shamLink.EntityData.SegmentPath
    shamLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shamLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shamLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shamLink.EntityData.Children = types.NewOrderedMap()
    shamLink.EntityData.Children.Append("authentication", types.YChild{"Authentication", &shamLink.Authentication})
    shamLink.EntityData.Children.Append("encryption", types.YChild{"Encryption", &shamLink.Encryption})
    shamLink.EntityData.Leafs = types.NewOrderedMap()
    shamLink.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", shamLink.SourceAddress})
    shamLink.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", shamLink.DestinationAddress})
    shamLink.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", shamLink.Enable})
    shamLink.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", shamLink.HelloInterval})
    shamLink.EntityData.Leafs.Append("dead-interval", types.YLeaf{"DeadInterval", shamLink.DeadInterval})
    shamLink.EntityData.Leafs.Append("retransmit-interval", types.YLeaf{"RetransmitInterval", shamLink.RetransmitInterval})
    shamLink.EntityData.Leafs.Append("transmit-delay", types.YLeaf{"TransmitDelay", shamLink.TransmitDelay})

    shamLink.EntityData.YListKeys = []string {"SourceAddress", "DestinationAddress"}

    return &(shamLink.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_ShamLinks_ShamLink_Authentication
// Authenticate OSPFv3 packets
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_ShamLinks_ShamLink_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec AH authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Use the MD5 or SHA1 algorithm. The type is Ospfv3Authentication.
    Algorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    Password interface{}
}

func (authentication *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_ShamLinks_ShamLink_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "sham-link"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/sham-links/sham-link/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", authentication.Enable})
    authentication.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", authentication.Spi})
    authentication.EntityData.Leafs.Append("algorithm", types.YLeaf{"Algorithm", authentication.Algorithm})
    authentication.EntityData.Leafs.Append("password", types.YLeaf{"Password", authentication.Password})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_ShamLinks_ShamLink_Encryption
// Encrypt and authenticate OSPFv3 packets
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_ShamLinks_ShamLink_Encryption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec ESP authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Specify the encryption algorithm. The type is Ospfv3EncryptionAlgorithm.
    EncryptionAlgorithm interface{}

    // Encryption password. The type is string with pattern: b'(!.+)|([^!].+)'.
    EncryptionPassword interface{}

    // Use the NULL, MD5 or SHA1 algorithm. The type is Ospfv3AuthenticationType2.
    AuthenticationAlgorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    AuthenticationPassword interface{}
}

func (encryption *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_ShamLinks_ShamLink_Encryption) GetEntityData() *types.CommonEntityData {
    encryption.EntityData.YFilter = encryption.YFilter
    encryption.EntityData.YangName = "encryption"
    encryption.EntityData.BundleName = "cisco_ios_xr"
    encryption.EntityData.ParentYangName = "sham-link"
    encryption.EntityData.SegmentPath = "encryption"
    encryption.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/sham-links/sham-link/" + encryption.EntityData.SegmentPath
    encryption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryption.EntityData.Children = types.NewOrderedMap()
    encryption.EntityData.Leafs = types.NewOrderedMap()
    encryption.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", encryption.Enable})
    encryption.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", encryption.Spi})
    encryption.EntityData.Leafs.Append("encryption-algorithm", types.YLeaf{"EncryptionAlgorithm", encryption.EncryptionAlgorithm})
    encryption.EntityData.Leafs.Append("encryption-password", types.YLeaf{"EncryptionPassword", encryption.EncryptionPassword})
    encryption.EntityData.Leafs.Append("authentication-algorithm", types.YLeaf{"AuthenticationAlgorithm", encryption.AuthenticationAlgorithm})
    encryption.EntityData.Leafs.Append("authentication-password", types.YLeaf{"AuthenticationPassword", encryption.AuthenticationPassword})

    encryption.EntityData.YListKeys = []string {}

    return &(encryption.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_VirtualLinks
// Virtual link sub-mode
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_VirtualLinks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router ID of virtual link neighbor. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_VirtualLinks_VirtualLink.
    VirtualLink []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_VirtualLinks_VirtualLink
}

func (virtualLinks *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_VirtualLinks) GetEntityData() *types.CommonEntityData {
    virtualLinks.EntityData.YFilter = virtualLinks.YFilter
    virtualLinks.EntityData.YangName = "virtual-links"
    virtualLinks.EntityData.BundleName = "cisco_ios_xr"
    virtualLinks.EntityData.ParentYangName = "area-address"
    virtualLinks.EntityData.SegmentPath = "virtual-links"
    virtualLinks.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/" + virtualLinks.EntityData.SegmentPath
    virtualLinks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    virtualLinks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    virtualLinks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    virtualLinks.EntityData.Children = types.NewOrderedMap()
    virtualLinks.EntityData.Children.Append("virtual-link", types.YChild{"VirtualLink", nil})
    for i := range virtualLinks.VirtualLink {
        virtualLinks.EntityData.Children.Append(types.GetSegmentPath(virtualLinks.VirtualLink[i]), types.YChild{"VirtualLink", virtualLinks.VirtualLink[i]})
    }
    virtualLinks.EntityData.Leafs = types.NewOrderedMap()

    virtualLinks.EntityData.YListKeys = []string {}

    return &(virtualLinks.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_VirtualLinks_VirtualLink
// Router ID of virtual link neighbor
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_VirtualLinks_VirtualLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Router ID of virtual link neighbor. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    VirtualLinkAddress interface{}

    // Enabled virtual link. The type is interface{}.
    Enable interface{}

    // Time between HELLO packets. The type is interface{} with range: 1..65535.
    // Units are second.
    HelloInterval interface{}

    // Interval after which a neighbor is declared dead (in seconds). The type is
    // interface{} with range: 1..65535. Units are second.
    DeadInterval interface{}

    // Specify the transmit interval in seconds. The type is interface{} with
    // range: 1..65535. Units are second.
    RetransmitInterval interface{}

    // Specify the transmit delay in seconds. The type is interface{} with range:
    // 1..65535. Units are second.
    TransmitDelay interface{}

    // Authenticate OSPFv3 packets.
    Authentication Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_VirtualLinks_VirtualLink_Authentication

    // Encrypt and authenticate OSPFv3 packets.
    Encryption Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_VirtualLinks_VirtualLink_Encryption
}

func (virtualLink *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_VirtualLinks_VirtualLink) GetEntityData() *types.CommonEntityData {
    virtualLink.EntityData.YFilter = virtualLink.YFilter
    virtualLink.EntityData.YangName = "virtual-link"
    virtualLink.EntityData.BundleName = "cisco_ios_xr"
    virtualLink.EntityData.ParentYangName = "virtual-links"
    virtualLink.EntityData.SegmentPath = "virtual-link" + types.AddKeyToken(virtualLink.VirtualLinkAddress, "virtual-link-address")
    virtualLink.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/virtual-links/" + virtualLink.EntityData.SegmentPath
    virtualLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    virtualLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    virtualLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    virtualLink.EntityData.Children = types.NewOrderedMap()
    virtualLink.EntityData.Children.Append("authentication", types.YChild{"Authentication", &virtualLink.Authentication})
    virtualLink.EntityData.Children.Append("encryption", types.YChild{"Encryption", &virtualLink.Encryption})
    virtualLink.EntityData.Leafs = types.NewOrderedMap()
    virtualLink.EntityData.Leafs.Append("virtual-link-address", types.YLeaf{"VirtualLinkAddress", virtualLink.VirtualLinkAddress})
    virtualLink.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", virtualLink.Enable})
    virtualLink.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", virtualLink.HelloInterval})
    virtualLink.EntityData.Leafs.Append("dead-interval", types.YLeaf{"DeadInterval", virtualLink.DeadInterval})
    virtualLink.EntityData.Leafs.Append("retransmit-interval", types.YLeaf{"RetransmitInterval", virtualLink.RetransmitInterval})
    virtualLink.EntityData.Leafs.Append("transmit-delay", types.YLeaf{"TransmitDelay", virtualLink.TransmitDelay})

    virtualLink.EntityData.YListKeys = []string {"VirtualLinkAddress"}

    return &(virtualLink.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_VirtualLinks_VirtualLink_Authentication
// Authenticate OSPFv3 packets
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_VirtualLinks_VirtualLink_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec AH authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Use the MD5 or SHA1 algorithm. The type is Ospfv3Authentication.
    Algorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    Password interface{}
}

func (authentication *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_VirtualLinks_VirtualLink_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "virtual-link"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/virtual-links/virtual-link/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", authentication.Enable})
    authentication.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", authentication.Spi})
    authentication.EntityData.Leafs.Append("algorithm", types.YLeaf{"Algorithm", authentication.Algorithm})
    authentication.EntityData.Leafs.Append("password", types.YLeaf{"Password", authentication.Password})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_VirtualLinks_VirtualLink_Encryption
// Encrypt and authenticate OSPFv3 packets
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_VirtualLinks_VirtualLink_Encryption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec ESP authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Specify the encryption algorithm. The type is Ospfv3EncryptionAlgorithm.
    EncryptionAlgorithm interface{}

    // Encryption password. The type is string with pattern: b'(!.+)|([^!].+)'.
    EncryptionPassword interface{}

    // Use the NULL, MD5 or SHA1 algorithm. The type is Ospfv3AuthenticationType2.
    AuthenticationAlgorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    AuthenticationPassword interface{}
}

func (encryption *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAddress_VirtualLinks_VirtualLink_Encryption) GetEntityData() *types.CommonEntityData {
    encryption.EntityData.YFilter = encryption.YFilter
    encryption.EntityData.YangName = "encryption"
    encryption.EntityData.BundleName = "cisco_ios_xr"
    encryption.EntityData.ParentYangName = "virtual-link"
    encryption.EntityData.SegmentPath = "encryption"
    encryption.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-address/virtual-links/virtual-link/" + encryption.EntityData.SegmentPath
    encryption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryption.EntityData.Children = types.NewOrderedMap()
    encryption.EntityData.Leafs = types.NewOrderedMap()
    encryption.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", encryption.Enable})
    encryption.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", encryption.Spi})
    encryption.EntityData.Leafs.Append("encryption-algorithm", types.YLeaf{"EncryptionAlgorithm", encryption.EncryptionAlgorithm})
    encryption.EntityData.Leafs.Append("encryption-password", types.YLeaf{"EncryptionPassword", encryption.EncryptionPassword})
    encryption.EntityData.Leafs.Append("authentication-algorithm", types.YLeaf{"AuthenticationAlgorithm", encryption.AuthenticationAlgorithm})
    encryption.EntityData.Leafs.Append("authentication-password", types.YLeaf{"AuthenticationPassword", encryption.AuthenticationPassword})

    encryption.EntityData.YListKeys = []string {}

    return &(encryption.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId
// Configuration for a particular area
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Area ID if in integer format. The type is
    // interface{} with range: 0..4294967295.
    AreaId interface{}

    // Specify area as a stub area.  Allowed only in non-backbone areas. The type
    // is bool.
    Stub interface{}

    // Limit size of OSPFv3 packets. The type is interface{} with range:
    // 256..10000.
    PacketSize interface{}

    // Instance ID. The type is interface{} with range: 0..255.
    Instance interface{}

    // Enable/disable demand circuit operation. The type is bool.
    DemandCircuit interface{}

    // Specify router priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // Translate Type 7 to Type 5, even if not elected NSSA translator. The type
    // is bool.
    Type7TranslateAlways interface{}

    // Enable/disable prefix suppression on an interface. The type is bool.
    PrefixSuppression interface{}

    // Enable OSPFv3 area. The type is interface{}.
    Enable interface{}

    // Enable/disable ignoring of MTU in DBD packets. The type is bool.
    MtuIgnore interface{}

    // Enable/disable routing updates on an interface. The type is bool.
    Passive interface{}

    // Time between HELLO packets. The type is interface{} with range: 1..65535.
    // Units are second.
    HelloInterval interface{}

    // Interval after which a neighbor is declared dead (in seconds). The type is
    // interface{} with range: 1..65535. Units are second.
    DeadInterval interface{}

    // Set the summary default-cost of a NSSA/stub area. The type is interface{}
    // with range: 0..16777215.
    DefaultCost interface{}

    // Enable/disable flood reduction. The type is bool.
    FloodReduction interface{}

    // Specify the transmit interval in seconds. The type is interface{} with
    // range: 1..65535. Units are second.
    RetransmitInterval interface{}

    // Enable/Disable MPLS LDP sync. The type is bool.
    LdpSync interface{}

    // Specify network type. The type is Ospfv3Network.
    Network interface{}

    // Specify the transmit delay in seconds. The type is interface{} with range:
    // 1..65535. Units are second.
    TransmitDelay interface{}

    // Interface cost. The type is interface{} with range: 1..65535.
    Cost interface{}

    // Authenticate OSPFv3 packets.
    Authentication Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Authentication

    // Configure BFD parameters.
    Bfd Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Bfd

    // Range configuration.
    Ranges Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Ranges

    // Encrypt and authenticate OSPFv3 packets.
    Encryption Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Encryption

    // Specify area as a NSSA area.  Allowed only in non-backbone areas.
    Nssa Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Nssa

    // Database filter.
    DatabaseFilter Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_DatabaseFilter

    // Filter prefixes to/from RIB.
    DistributeList Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_DistributeList

    // OSPFv3 interfaces.
    Interfaces Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces

    // Area Scope Configuration.
    AreaScope Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope

    // Sham Link sub-mode.
    ShamLinks Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_ShamLinks

    // Virtual link sub-mode.
    VirtualLinks Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_VirtualLinks
}

func (areaAreaId *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId) GetEntityData() *types.CommonEntityData {
    areaAreaId.EntityData.YFilter = areaAreaId.YFilter
    areaAreaId.EntityData.YangName = "area-area-id"
    areaAreaId.EntityData.BundleName = "cisco_ios_xr"
    areaAreaId.EntityData.ParentYangName = "area-addresses"
    areaAreaId.EntityData.SegmentPath = "area-area-id" + types.AddKeyToken(areaAreaId.AreaId, "area-id")
    areaAreaId.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/" + areaAreaId.EntityData.SegmentPath
    areaAreaId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    areaAreaId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    areaAreaId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    areaAreaId.EntityData.Children = types.NewOrderedMap()
    areaAreaId.EntityData.Children.Append("authentication", types.YChild{"Authentication", &areaAreaId.Authentication})
    areaAreaId.EntityData.Children.Append("bfd", types.YChild{"Bfd", &areaAreaId.Bfd})
    areaAreaId.EntityData.Children.Append("ranges", types.YChild{"Ranges", &areaAreaId.Ranges})
    areaAreaId.EntityData.Children.Append("encryption", types.YChild{"Encryption", &areaAreaId.Encryption})
    areaAreaId.EntityData.Children.Append("nssa", types.YChild{"Nssa", &areaAreaId.Nssa})
    areaAreaId.EntityData.Children.Append("database-filter", types.YChild{"DatabaseFilter", &areaAreaId.DatabaseFilter})
    areaAreaId.EntityData.Children.Append("distribute-list", types.YChild{"DistributeList", &areaAreaId.DistributeList})
    areaAreaId.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &areaAreaId.Interfaces})
    areaAreaId.EntityData.Children.Append("area-scope", types.YChild{"AreaScope", &areaAreaId.AreaScope})
    areaAreaId.EntityData.Children.Append("sham-links", types.YChild{"ShamLinks", &areaAreaId.ShamLinks})
    areaAreaId.EntityData.Children.Append("virtual-links", types.YChild{"VirtualLinks", &areaAreaId.VirtualLinks})
    areaAreaId.EntityData.Leafs = types.NewOrderedMap()
    areaAreaId.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", areaAreaId.AreaId})
    areaAreaId.EntityData.Leafs.Append("stub", types.YLeaf{"Stub", areaAreaId.Stub})
    areaAreaId.EntityData.Leafs.Append("packet-size", types.YLeaf{"PacketSize", areaAreaId.PacketSize})
    areaAreaId.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", areaAreaId.Instance})
    areaAreaId.EntityData.Leafs.Append("demand-circuit", types.YLeaf{"DemandCircuit", areaAreaId.DemandCircuit})
    areaAreaId.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", areaAreaId.Priority})
    areaAreaId.EntityData.Leafs.Append("type7-translate-always", types.YLeaf{"Type7TranslateAlways", areaAreaId.Type7TranslateAlways})
    areaAreaId.EntityData.Leafs.Append("prefix-suppression", types.YLeaf{"PrefixSuppression", areaAreaId.PrefixSuppression})
    areaAreaId.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", areaAreaId.Enable})
    areaAreaId.EntityData.Leafs.Append("mtu-ignore", types.YLeaf{"MtuIgnore", areaAreaId.MtuIgnore})
    areaAreaId.EntityData.Leafs.Append("passive", types.YLeaf{"Passive", areaAreaId.Passive})
    areaAreaId.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", areaAreaId.HelloInterval})
    areaAreaId.EntityData.Leafs.Append("dead-interval", types.YLeaf{"DeadInterval", areaAreaId.DeadInterval})
    areaAreaId.EntityData.Leafs.Append("default-cost", types.YLeaf{"DefaultCost", areaAreaId.DefaultCost})
    areaAreaId.EntityData.Leafs.Append("flood-reduction", types.YLeaf{"FloodReduction", areaAreaId.FloodReduction})
    areaAreaId.EntityData.Leafs.Append("retransmit-interval", types.YLeaf{"RetransmitInterval", areaAreaId.RetransmitInterval})
    areaAreaId.EntityData.Leafs.Append("ldp-sync", types.YLeaf{"LdpSync", areaAreaId.LdpSync})
    areaAreaId.EntityData.Leafs.Append("network", types.YLeaf{"Network", areaAreaId.Network})
    areaAreaId.EntityData.Leafs.Append("transmit-delay", types.YLeaf{"TransmitDelay", areaAreaId.TransmitDelay})
    areaAreaId.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", areaAreaId.Cost})

    areaAreaId.EntityData.YListKeys = []string {"AreaId"}

    return &(areaAreaId.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Authentication
// Authenticate OSPFv3 packets
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec AH authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Use the MD5 or SHA1 algorithm. The type is Ospfv3Authentication.
    Algorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    Password interface{}
}

func (authentication *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "area-area-id"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", authentication.Enable})
    authentication.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", authentication.Spi})
    authentication.EntityData.Leafs.Append("algorithm", types.YLeaf{"Algorithm", authentication.Algorithm})
    authentication.EntityData.Leafs.Append("password", types.YLeaf{"Password", authentication.Password})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Bfd
// Configure BFD parameters
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detect multiplier. The type is interface{} with range: 2..50.
    DetectionMultiplier interface{}

    // Hello interval in milli-seconds. The type is interface{} with range:
    // 3..30000. Units are millisecond.
    Interval interface{}

    // Enable or disable BFD fast detection. The type is Ospfv3bfdEnableMode.
    FastDetectMode interface{}
}

func (bfd *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "area-area-id"
    bfd.EntityData.SegmentPath = "bfd"
    bfd.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/" + bfd.EntityData.SegmentPath
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = types.NewOrderedMap()
    bfd.EntityData.Leafs = types.NewOrderedMap()
    bfd.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", bfd.DetectionMultiplier})
    bfd.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", bfd.Interval})
    bfd.EntityData.Leafs.Append("fast-detect-mode", types.YLeaf{"FastDetectMode", bfd.FastDetectMode})

    bfd.EntityData.YListKeys = []string {}

    return &(bfd.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Ranges
// Range configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Ranges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summarize inter-area routes matching prefix/length. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Ranges_Range.
    Range []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Ranges_Range
}

func (ranges *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Ranges) GetEntityData() *types.CommonEntityData {
    ranges.EntityData.YFilter = ranges.YFilter
    ranges.EntityData.YangName = "ranges"
    ranges.EntityData.BundleName = "cisco_ios_xr"
    ranges.EntityData.ParentYangName = "area-area-id"
    ranges.EntityData.SegmentPath = "ranges"
    ranges.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/" + ranges.EntityData.SegmentPath
    ranges.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ranges.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ranges.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ranges.EntityData.Children = types.NewOrderedMap()
    ranges.EntityData.Children.Append("range", types.YChild{"Range", nil})
    for i := range ranges.Range {
        ranges.EntityData.Children.Append(types.GetSegmentPath(ranges.Range[i]), types.YChild{"Range", ranges.Range[i]})
    }
    ranges.EntityData.Leafs = types.NewOrderedMap()

    ranges.EntityData.YListKeys = []string {}

    return &(ranges.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Ranges_Range
// Summarize inter-area routes matching
// prefix/length
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Ranges_Range struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IPv6 prefix format. The type is string with
    // pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Prefix interface{}

    // This attribute is a key. IPV6 prefix length. The type is interface{} with
    // range: 0..128.
    PrefixLength interface{}

    // Do not advertise address range. The type is bool. The default value is
    // false.
    NotAdvertise interface{}

    // Specified metric for this range. The type is interface{} with range:
    // 1..16777214.
    Cost interface{}
}

func (self *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Ranges_Range) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "range"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "ranges"
    self.EntityData.SegmentPath = "range" + types.AddKeyToken(self.Prefix, "prefix") + types.AddKeyToken(self.PrefixLength, "prefix-length")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/ranges/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", self.Prefix})
    self.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", self.PrefixLength})
    self.EntityData.Leafs.Append("not-advertise", types.YLeaf{"NotAdvertise", self.NotAdvertise})
    self.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", self.Cost})

    self.EntityData.YListKeys = []string {"Prefix", "PrefixLength"}

    return &(self.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Encryption
// Encrypt and authenticate OSPFv3 packets
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Encryption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec ESP authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Specify the encryption algorithm. The type is Ospfv3EncryptionAlgorithm.
    EncryptionAlgorithm interface{}

    // Encryption password. The type is string with pattern: b'(!.+)|([^!].+)'.
    EncryptionPassword interface{}

    // Use the NULL, MD5 or SHA1 algorithm. The type is Ospfv3AuthenticationType2.
    AuthenticationAlgorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    AuthenticationPassword interface{}
}

func (encryption *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Encryption) GetEntityData() *types.CommonEntityData {
    encryption.EntityData.YFilter = encryption.YFilter
    encryption.EntityData.YangName = "encryption"
    encryption.EntityData.BundleName = "cisco_ios_xr"
    encryption.EntityData.ParentYangName = "area-area-id"
    encryption.EntityData.SegmentPath = "encryption"
    encryption.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/" + encryption.EntityData.SegmentPath
    encryption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryption.EntityData.Children = types.NewOrderedMap()
    encryption.EntityData.Leafs = types.NewOrderedMap()
    encryption.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", encryption.Enable})
    encryption.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", encryption.Spi})
    encryption.EntityData.Leafs.Append("encryption-algorithm", types.YLeaf{"EncryptionAlgorithm", encryption.EncryptionAlgorithm})
    encryption.EntityData.Leafs.Append("encryption-password", types.YLeaf{"EncryptionPassword", encryption.EncryptionPassword})
    encryption.EntityData.Leafs.Append("authentication-algorithm", types.YLeaf{"AuthenticationAlgorithm", encryption.AuthenticationAlgorithm})
    encryption.EntityData.Leafs.Append("authentication-password", types.YLeaf{"AuthenticationPassword", encryption.AuthenticationPassword})

    encryption.EntityData.YListKeys = []string {}

    return &(encryption.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Nssa
// Specify area as a NSSA area.  Allowed only in
// non-backbone areas
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Nssa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // No redistribution into this NSSA area. The type is bool. The default value
    // is false.
    NoRedistribution interface{}

    // Originate Type 7 default into NSSA area. The type is bool. The default
    // value is false.
    DefaultInfoOriginate interface{}

    // Only valid with DefaultInfoOriginate. The type is interface{} with range:
    // 0..16777214.
    Metric interface{}

    // Only valid with DefaultInfoOriginate. The type is Ospfv3Metric.
    MetricType interface{}

    // Do not send summary LSA into NSSA. The type is interface{}.
    NoSummary interface{}
}

func (nssa *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Nssa) GetEntityData() *types.CommonEntityData {
    nssa.EntityData.YFilter = nssa.YFilter
    nssa.EntityData.YangName = "nssa"
    nssa.EntityData.BundleName = "cisco_ios_xr"
    nssa.EntityData.ParentYangName = "area-area-id"
    nssa.EntityData.SegmentPath = "nssa"
    nssa.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/" + nssa.EntityData.SegmentPath
    nssa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nssa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nssa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nssa.EntityData.Children = types.NewOrderedMap()
    nssa.EntityData.Leafs = types.NewOrderedMap()
    nssa.EntityData.Leafs.Append("no-redistribution", types.YLeaf{"NoRedistribution", nssa.NoRedistribution})
    nssa.EntityData.Leafs.Append("default-info-originate", types.YLeaf{"DefaultInfoOriginate", nssa.DefaultInfoOriginate})
    nssa.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", nssa.Metric})
    nssa.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", nssa.MetricType})
    nssa.EntityData.Leafs.Append("no-summary", types.YLeaf{"NoSummary", nssa.NoSummary})

    nssa.EntityData.YListKeys = []string {}

    return &(nssa.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_DatabaseFilter
// Database filter
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_DatabaseFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All.
    All Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_DatabaseFilter_All
}

func (databaseFilter *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_DatabaseFilter) GetEntityData() *types.CommonEntityData {
    databaseFilter.EntityData.YFilter = databaseFilter.YFilter
    databaseFilter.EntityData.YangName = "database-filter"
    databaseFilter.EntityData.BundleName = "cisco_ios_xr"
    databaseFilter.EntityData.ParentYangName = "area-area-id"
    databaseFilter.EntityData.SegmentPath = "database-filter"
    databaseFilter.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/" + databaseFilter.EntityData.SegmentPath
    databaseFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    databaseFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    databaseFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    databaseFilter.EntityData.Children = types.NewOrderedMap()
    databaseFilter.EntityData.Children.Append("all", types.YChild{"All", &databaseFilter.All})
    databaseFilter.EntityData.Leafs = types.NewOrderedMap()

    databaseFilter.EntityData.YListKeys = []string {}

    return &(databaseFilter.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_DatabaseFilter_All
// All
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_DatabaseFilter_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable or disable database-filter. The type is bool.
    Out interface{}
}

func (all *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_DatabaseFilter_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "database-filter"
    all.EntityData.SegmentPath = "all"
    all.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/database-filter/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Leafs = types.NewOrderedMap()
    all.EntityData.Leafs.Append("out", types.YLeaf{"Out", all.Out})

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_DistributeList
// Filter prefixes to/from RIB
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_DistributeList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter prefixes installed to RIB.
    In Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_DistributeList_In
}

func (distributeList *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_DistributeList) GetEntityData() *types.CommonEntityData {
    distributeList.EntityData.YFilter = distributeList.YFilter
    distributeList.EntityData.YangName = "distribute-list"
    distributeList.EntityData.BundleName = "cisco_ios_xr"
    distributeList.EntityData.ParentYangName = "area-area-id"
    distributeList.EntityData.SegmentPath = "distribute-list"
    distributeList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/" + distributeList.EntityData.SegmentPath
    distributeList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    distributeList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    distributeList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    distributeList.EntityData.Children = types.NewOrderedMap()
    distributeList.EntityData.Children.Append("in", types.YChild{"In", &distributeList.In})
    distributeList.EntityData.Leafs = types.NewOrderedMap()

    distributeList.EntityData.YListKeys = []string {}

    return &(distributeList.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_DistributeList_In
// Filter prefixes installed to RIB
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_DistributeList_In struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter prefixes based on an IPv6 prefix-list. The type is string.
    PrefixList interface{}
}

func (in *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_DistributeList_In) GetEntityData() *types.CommonEntityData {
    in.EntityData.YFilter = in.YFilter
    in.EntityData.YangName = "in"
    in.EntityData.BundleName = "cisco_ios_xr"
    in.EntityData.ParentYangName = "distribute-list"
    in.EntityData.SegmentPath = "in"
    in.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/distribute-list/" + in.EntityData.SegmentPath
    in.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    in.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    in.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    in.EntityData.Children = types.NewOrderedMap()
    in.EntityData.Leafs = types.NewOrderedMap()
    in.EntityData.Leafs.Append("prefix-list", types.YLeaf{"PrefixList", in.PrefixList})

    in.EntityData.YListKeys = []string {}

    return &(in.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces
// OSPFv3 interfaces
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPFv3 interface. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface.
    Interface []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface
}

func (interfaces *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "area-area-id"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/" + interfaces.EntityData.SegmentPath
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface
// OSPFv3 interface
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface to configure. The type is string with
    // pattern: b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Enable OSPFv3 interface. The type is interface{}.
    Enable interface{}

    // Interval after which a neighbor is declared dead (in seconds). The type is
    // interface{} with range: 1..65535. Units are second.
    DeadInterval interface{}

    // Enable/disable flood reduction. The type is bool.
    FloodReduction interface{}

    // Interface cost. The type is interface{} with range: 1..65535.
    Cost interface{}

    // Specify the transmit delay in seconds. The type is interface{} with range:
    // 1..65535. Units are second.
    TransmitDelay interface{}

    // Instance ID. The type is interface{} with range: 0..255.
    Instance interface{}

    // Enable/Disable MPLS LDP sync. The type is bool.
    LdpSync interface{}

    // Enable/disable ignoring of MTU in DBD packets. The type is bool.
    MtuIgnore interface{}

    // Specify the transmit interval in seconds. The type is interface{} with
    // range: 1..65535. Units are second.
    RetransmitInterval interface{}

    // Time between HELLO packets. The type is interface{} with range: 1..65535.
    // Units are second.
    HelloInterval interface{}

    // Enable/disable routing updates on an interface. The type is bool.
    Passive interface{}

    // Limit size of OSPFv3 packets. The type is interface{} with range:
    // 256..10000.
    PacketSize interface{}

    // Enable/disable prefix suppression on an interface. The type is bool.
    PrefixSuppression interface{}

    // Specify router priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // Specify network type. The type is Ospfv3Network.
    Network interface{}

    // Enable/disable demand circuit operation. The type is bool.
    DemandCircuit interface{}

    // Authenticate OSPFv3 packets.
    Authentication Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Authentication

    // Specify a neighbor router.
    Neighbors Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Neighbors

    // Encrypt and authenticate OSPFv3 packets.
    Encryption Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Encryption

    // Configure BFD parameters.
    Bfd Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Bfd

    // Database filter.
    DatabaseFilter Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DatabaseFilter

    // Filter prefixes to/from RIB.
    DistributeList Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DistributeList

    // Fast-reroute configuration.
    FastReroute Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute
}

func (self *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("authentication", types.YChild{"Authentication", &self.Authentication})
    self.EntityData.Children.Append("neighbors", types.YChild{"Neighbors", &self.Neighbors})
    self.EntityData.Children.Append("encryption", types.YChild{"Encryption", &self.Encryption})
    self.EntityData.Children.Append("bfd", types.YChild{"Bfd", &self.Bfd})
    self.EntityData.Children.Append("database-filter", types.YChild{"DatabaseFilter", &self.DatabaseFilter})
    self.EntityData.Children.Append("distribute-list", types.YChild{"DistributeList", &self.DistributeList})
    self.EntityData.Children.Append("fast-reroute", types.YChild{"FastReroute", &self.FastReroute})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", self.Enable})
    self.EntityData.Leafs.Append("dead-interval", types.YLeaf{"DeadInterval", self.DeadInterval})
    self.EntityData.Leafs.Append("flood-reduction", types.YLeaf{"FloodReduction", self.FloodReduction})
    self.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", self.Cost})
    self.EntityData.Leafs.Append("transmit-delay", types.YLeaf{"TransmitDelay", self.TransmitDelay})
    self.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", self.Instance})
    self.EntityData.Leafs.Append("ldp-sync", types.YLeaf{"LdpSync", self.LdpSync})
    self.EntityData.Leafs.Append("mtu-ignore", types.YLeaf{"MtuIgnore", self.MtuIgnore})
    self.EntityData.Leafs.Append("retransmit-interval", types.YLeaf{"RetransmitInterval", self.RetransmitInterval})
    self.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", self.HelloInterval})
    self.EntityData.Leafs.Append("passive", types.YLeaf{"Passive", self.Passive})
    self.EntityData.Leafs.Append("packet-size", types.YLeaf{"PacketSize", self.PacketSize})
    self.EntityData.Leafs.Append("prefix-suppression", types.YLeaf{"PrefixSuppression", self.PrefixSuppression})
    self.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", self.Priority})
    self.EntityData.Leafs.Append("network", types.YLeaf{"Network", self.Network})
    self.EntityData.Leafs.Append("demand-circuit", types.YLeaf{"DemandCircuit", self.DemandCircuit})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Authentication
// Authenticate OSPFv3 packets
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec AH authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Use the MD5 or SHA1 algorithm. The type is Ospfv3Authentication.
    Algorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    Password interface{}
}

func (authentication *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "interface"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/interfaces/interface/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", authentication.Enable})
    authentication.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", authentication.Spi})
    authentication.EntityData.Leafs.Append("algorithm", types.YLeaf{"Algorithm", authentication.Algorithm})
    authentication.EntityData.Leafs.Append("password", types.YLeaf{"Password", authentication.Password})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Neighbors
// Specify a neighbor router
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 address. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Neighbors_Neighbor.
    Neighbor []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Neighbors_Neighbor
}

func (neighbors *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "cisco_ios_xr"
    neighbors.EntityData.ParentYangName = "interface"
    neighbors.EntityData.SegmentPath = "neighbors"
    neighbors.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/interfaces/interface/" + neighbors.EntityData.SegmentPath
    neighbors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbors.EntityData.Children = types.NewOrderedMap()
    neighbors.EntityData.Children.Append("neighbor", types.YChild{"Neighbor", nil})
    for i := range neighbors.Neighbor {
        neighbors.EntityData.Children.Append(types.GetSegmentPath(neighbors.Neighbor[i]), types.YChild{"Neighbor", neighbors.Neighbor[i]})
    }
    neighbors.EntityData.Leafs = types.NewOrderedMap()

    neighbors.EntityData.YListKeys = []string {}

    return &(neighbors.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Neighbors_Neighbor
// IPv6 address
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IPV6 address. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    NeighborAddress interface{}

    // OSPFv3 priority of non-broadcast neighbor. The type is interface{} with
    // range: 0..255.
    Priority interface{}

    // OSPFv3 dead-router polling interval (in seconds). The type is interface{}
    // with range: 0..65535. Units are second.
    PollInterval interface{}

    // OSPFv3 cost for point-to-multipoint neighbor. The type is interface{} with
    // range: 1..65535.
    Cost interface{}

    // Filter OSPFv3 LSA during synchronization and flooding for
    // point-to-multipoint neighbor. The type is bool.
    DatabaseFilter interface{}

    // Zone. The type is string.
    Zone interface{}
}

func (neighbor *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Neighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "neighbors"
    neighbor.EntityData.SegmentPath = "neighbor" + types.AddKeyToken(neighbor.NeighborAddress, "neighbor-address")
    neighbor.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/interfaces/interface/neighbors/" + neighbor.EntityData.SegmentPath
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", neighbor.NeighborAddress})
    neighbor.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", neighbor.Priority})
    neighbor.EntityData.Leafs.Append("poll-interval", types.YLeaf{"PollInterval", neighbor.PollInterval})
    neighbor.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", neighbor.Cost})
    neighbor.EntityData.Leafs.Append("database-filter", types.YLeaf{"DatabaseFilter", neighbor.DatabaseFilter})
    neighbor.EntityData.Leafs.Append("zone", types.YLeaf{"Zone", neighbor.Zone})

    neighbor.EntityData.YListKeys = []string {"NeighborAddress"}

    return &(neighbor.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Encryption
// Encrypt and authenticate OSPFv3 packets
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Encryption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec ESP authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Specify the encryption algorithm. The type is Ospfv3EncryptionAlgorithm.
    EncryptionAlgorithm interface{}

    // Encryption password. The type is string with pattern: b'(!.+)|([^!].+)'.
    EncryptionPassword interface{}

    // Use the NULL, MD5 or SHA1 algorithm. The type is Ospfv3AuthenticationType2.
    AuthenticationAlgorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    AuthenticationPassword interface{}
}

func (encryption *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Encryption) GetEntityData() *types.CommonEntityData {
    encryption.EntityData.YFilter = encryption.YFilter
    encryption.EntityData.YangName = "encryption"
    encryption.EntityData.BundleName = "cisco_ios_xr"
    encryption.EntityData.ParentYangName = "interface"
    encryption.EntityData.SegmentPath = "encryption"
    encryption.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/interfaces/interface/" + encryption.EntityData.SegmentPath
    encryption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryption.EntityData.Children = types.NewOrderedMap()
    encryption.EntityData.Leafs = types.NewOrderedMap()
    encryption.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", encryption.Enable})
    encryption.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", encryption.Spi})
    encryption.EntityData.Leafs.Append("encryption-algorithm", types.YLeaf{"EncryptionAlgorithm", encryption.EncryptionAlgorithm})
    encryption.EntityData.Leafs.Append("encryption-password", types.YLeaf{"EncryptionPassword", encryption.EncryptionPassword})
    encryption.EntityData.Leafs.Append("authentication-algorithm", types.YLeaf{"AuthenticationAlgorithm", encryption.AuthenticationAlgorithm})
    encryption.EntityData.Leafs.Append("authentication-password", types.YLeaf{"AuthenticationPassword", encryption.AuthenticationPassword})

    encryption.EntityData.YListKeys = []string {}

    return &(encryption.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Bfd
// Configure BFD parameters
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Hello interval in milli-seconds. The type is interface{} with range:
    // 3..30000. Units are millisecond.
    Interval interface{}

    // Enable or disable BFD fast detection. The type is Ospfv3bfdEnableMode.
    FastDetectMode interface{}

    // Detect multiplier. The type is interface{} with range: 2..50.
    DetectionMultiplier interface{}
}

func (bfd *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "interface"
    bfd.EntityData.SegmentPath = "bfd"
    bfd.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/interfaces/interface/" + bfd.EntityData.SegmentPath
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = types.NewOrderedMap()
    bfd.EntityData.Leafs = types.NewOrderedMap()
    bfd.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", bfd.Interval})
    bfd.EntityData.Leafs.Append("fast-detect-mode", types.YLeaf{"FastDetectMode", bfd.FastDetectMode})
    bfd.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", bfd.DetectionMultiplier})

    bfd.EntityData.YListKeys = []string {}

    return &(bfd.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DatabaseFilter
// Database filter
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DatabaseFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All.
    All Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DatabaseFilter_All
}

func (databaseFilter *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DatabaseFilter) GetEntityData() *types.CommonEntityData {
    databaseFilter.EntityData.YFilter = databaseFilter.YFilter
    databaseFilter.EntityData.YangName = "database-filter"
    databaseFilter.EntityData.BundleName = "cisco_ios_xr"
    databaseFilter.EntityData.ParentYangName = "interface"
    databaseFilter.EntityData.SegmentPath = "database-filter"
    databaseFilter.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/interfaces/interface/" + databaseFilter.EntityData.SegmentPath
    databaseFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    databaseFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    databaseFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    databaseFilter.EntityData.Children = types.NewOrderedMap()
    databaseFilter.EntityData.Children.Append("all", types.YChild{"All", &databaseFilter.All})
    databaseFilter.EntityData.Leafs = types.NewOrderedMap()

    databaseFilter.EntityData.YListKeys = []string {}

    return &(databaseFilter.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DatabaseFilter_All
// All
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DatabaseFilter_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable or disable database-filter. The type is bool.
    Out interface{}
}

func (all *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DatabaseFilter_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "database-filter"
    all.EntityData.SegmentPath = "all"
    all.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/interfaces/interface/database-filter/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Leafs = types.NewOrderedMap()
    all.EntityData.Leafs.Append("out", types.YLeaf{"Out", all.Out})

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DistributeList
// Filter prefixes to/from RIB
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DistributeList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter prefixes installed to RIB.
    In Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DistributeList_In
}

func (distributeList *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DistributeList) GetEntityData() *types.CommonEntityData {
    distributeList.EntityData.YFilter = distributeList.YFilter
    distributeList.EntityData.YangName = "distribute-list"
    distributeList.EntityData.BundleName = "cisco_ios_xr"
    distributeList.EntityData.ParentYangName = "interface"
    distributeList.EntityData.SegmentPath = "distribute-list"
    distributeList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/interfaces/interface/" + distributeList.EntityData.SegmentPath
    distributeList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    distributeList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    distributeList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    distributeList.EntityData.Children = types.NewOrderedMap()
    distributeList.EntityData.Children.Append("in", types.YChild{"In", &distributeList.In})
    distributeList.EntityData.Leafs = types.NewOrderedMap()

    distributeList.EntityData.YListKeys = []string {}

    return &(distributeList.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DistributeList_In
// Filter prefixes installed to RIB
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DistributeList_In struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter prefixes based on an IPv6 prefix-list. The type is string.
    PrefixList interface{}
}

func (in *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DistributeList_In) GetEntityData() *types.CommonEntityData {
    in.EntityData.YFilter = in.YFilter
    in.EntityData.YangName = "in"
    in.EntityData.BundleName = "cisco_ios_xr"
    in.EntityData.ParentYangName = "distribute-list"
    in.EntityData.SegmentPath = "in"
    in.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/interfaces/interface/distribute-list/" + in.EntityData.SegmentPath
    in.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    in.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    in.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    in.EntityData.Children = types.NewOrderedMap()
    in.EntityData.Leafs = types.NewOrderedMap()
    in.EntityData.Leafs.Append("prefix-list", types.YLeaf{"PrefixList", in.PrefixList})

    in.EntityData.YListKeys = []string {}

    return &(in.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute
// Fast-reroute configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable Fast-reroute per-link or per-prefix. The type is
    // Ospfv3FastReroute.
    FastRerouteEnable interface{}

    // Fast-reroute per-link configuration.
    PerLink Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink

    // Fast-reroute per-link configuration.
    PerPrefix Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix
}

func (fastReroute *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute) GetEntityData() *types.CommonEntityData {
    fastReroute.EntityData.YFilter = fastReroute.YFilter
    fastReroute.EntityData.YangName = "fast-reroute"
    fastReroute.EntityData.BundleName = "cisco_ios_xr"
    fastReroute.EntityData.ParentYangName = "interface"
    fastReroute.EntityData.SegmentPath = "fast-reroute"
    fastReroute.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/interfaces/interface/" + fastReroute.EntityData.SegmentPath
    fastReroute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fastReroute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fastReroute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fastReroute.EntityData.Children = types.NewOrderedMap()
    fastReroute.EntityData.Children.Append("per-link", types.YChild{"PerLink", &fastReroute.PerLink})
    fastReroute.EntityData.Children.Append("per-prefix", types.YChild{"PerPrefix", &fastReroute.PerPrefix})
    fastReroute.EntityData.Leafs = types.NewOrderedMap()
    fastReroute.EntityData.Leafs.Append("fast-reroute-enable", types.YLeaf{"FastRerouteEnable", fastReroute.FastRerouteEnable})

    fastReroute.EntityData.YListKeys = []string {}

    return &(fastReroute.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink
// Fast-reroute per-link configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Use only interfaces on the candidate list as a backup path. The type is
    // bool. The default value is false.
    FastRerouteUseCandidateOnly interface{}

    // Fast-reroute per-link/per-prefix candidate interface configuration.
    CandidateInterfaces Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces

    // Fast-reroute per-link/per-prefix exclude interface configuration.
    ExcludeInterfaces Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces
}

func (perLink *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink) GetEntityData() *types.CommonEntityData {
    perLink.EntityData.YFilter = perLink.YFilter
    perLink.EntityData.YangName = "per-link"
    perLink.EntityData.BundleName = "cisco_ios_xr"
    perLink.EntityData.ParentYangName = "fast-reroute"
    perLink.EntityData.SegmentPath = "per-link"
    perLink.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/interfaces/interface/fast-reroute/" + perLink.EntityData.SegmentPath
    perLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    perLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    perLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    perLink.EntityData.Children = types.NewOrderedMap()
    perLink.EntityData.Children.Append("candidate-interfaces", types.YChild{"CandidateInterfaces", &perLink.CandidateInterfaces})
    perLink.EntityData.Children.Append("exclude-interfaces", types.YChild{"ExcludeInterfaces", &perLink.ExcludeInterfaces})
    perLink.EntityData.Leafs = types.NewOrderedMap()
    perLink.EntityData.Leafs.Append("fast-reroute-use-candidate-only", types.YLeaf{"FastRerouteUseCandidateOnly", perLink.FastRerouteUseCandidateOnly})

    perLink.EntityData.YListKeys = []string {}

    return &(perLink.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces
// Fast-reroute per-link/per-prefix candidate
// interface configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Candidate backup interface. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces_CandidateInterface.
    CandidateInterface []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces_CandidateInterface
}

func (candidateInterfaces *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces) GetEntityData() *types.CommonEntityData {
    candidateInterfaces.EntityData.YFilter = candidateInterfaces.YFilter
    candidateInterfaces.EntityData.YangName = "candidate-interfaces"
    candidateInterfaces.EntityData.BundleName = "cisco_ios_xr"
    candidateInterfaces.EntityData.ParentYangName = "per-link"
    candidateInterfaces.EntityData.SegmentPath = "candidate-interfaces"
    candidateInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/interfaces/interface/fast-reroute/per-link/" + candidateInterfaces.EntityData.SegmentPath
    candidateInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterfaces.EntityData.Children = types.NewOrderedMap()
    candidateInterfaces.EntityData.Children.Append("candidate-interface", types.YChild{"CandidateInterface", nil})
    for i := range candidateInterfaces.CandidateInterface {
        candidateInterfaces.EntityData.Children.Append(types.GetSegmentPath(candidateInterfaces.CandidateInterface[i]), types.YChild{"CandidateInterface", candidateInterfaces.CandidateInterface[i]})
    }
    candidateInterfaces.EntityData.Leafs = types.NewOrderedMap()

    candidateInterfaces.EntityData.YListKeys = []string {}

    return &(candidateInterfaces.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces_CandidateInterface
// Candidate backup interface
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces_CandidateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (candidateInterface *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces_CandidateInterface) GetEntityData() *types.CommonEntityData {
    candidateInterface.EntityData.YFilter = candidateInterface.YFilter
    candidateInterface.EntityData.YangName = "candidate-interface"
    candidateInterface.EntityData.BundleName = "cisco_ios_xr"
    candidateInterface.EntityData.ParentYangName = "candidate-interfaces"
    candidateInterface.EntityData.SegmentPath = "candidate-interface" + types.AddKeyToken(candidateInterface.InterfaceName, "interface-name")
    candidateInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/interfaces/interface/fast-reroute/per-link/candidate-interfaces/" + candidateInterface.EntityData.SegmentPath
    candidateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterface.EntityData.Children = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", candidateInterface.InterfaceName})

    candidateInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(candidateInterface.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces
// Fast-reroute per-link/per-prefix exclude
// interface configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exclude an interface from becoming a backup. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface.
    ExcludeInterface []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface
}

func (excludeInterfaces *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces) GetEntityData() *types.CommonEntityData {
    excludeInterfaces.EntityData.YFilter = excludeInterfaces.YFilter
    excludeInterfaces.EntityData.YangName = "exclude-interfaces"
    excludeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    excludeInterfaces.EntityData.ParentYangName = "per-link"
    excludeInterfaces.EntityData.SegmentPath = "exclude-interfaces"
    excludeInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/interfaces/interface/fast-reroute/per-link/" + excludeInterfaces.EntityData.SegmentPath
    excludeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterfaces.EntityData.Children = types.NewOrderedMap()
    excludeInterfaces.EntityData.Children.Append("exclude-interface", types.YChild{"ExcludeInterface", nil})
    for i := range excludeInterfaces.ExcludeInterface {
        excludeInterfaces.EntityData.Children.Append(types.GetSegmentPath(excludeInterfaces.ExcludeInterface[i]), types.YChild{"ExcludeInterface", excludeInterfaces.ExcludeInterface[i]})
    }
    excludeInterfaces.EntityData.Leafs = types.NewOrderedMap()

    excludeInterfaces.EntityData.YListKeys = []string {}

    return &(excludeInterfaces.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface
// Exclude an interface from becoming a backup
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (excludeInterface *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface) GetEntityData() *types.CommonEntityData {
    excludeInterface.EntityData.YFilter = excludeInterface.YFilter
    excludeInterface.EntityData.YangName = "exclude-interface"
    excludeInterface.EntityData.BundleName = "cisco_ios_xr"
    excludeInterface.EntityData.ParentYangName = "exclude-interfaces"
    excludeInterface.EntityData.SegmentPath = "exclude-interface" + types.AddKeyToken(excludeInterface.InterfaceName, "interface-name")
    excludeInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/interfaces/interface/fast-reroute/per-link/exclude-interfaces/" + excludeInterface.EntityData.SegmentPath
    excludeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterface.EntityData.Children = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", excludeInterface.InterfaceName})

    excludeInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(excludeInterface.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix
// Fast-reroute per-link configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Use only interfaces on the candidate list as a backup path. The type is
    // bool. The default value is false.
    FastRerouteUseCandidateOnly interface{}

    // Fast-reroute per-link/per-prefix candidate interface configuration.
    CandidateInterfaces Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces

    // Fast-reroute per-link/per-prefix exclude interface configuration.
    ExcludeInterfaces Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces
}

func (perPrefix *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix) GetEntityData() *types.CommonEntityData {
    perPrefix.EntityData.YFilter = perPrefix.YFilter
    perPrefix.EntityData.YangName = "per-prefix"
    perPrefix.EntityData.BundleName = "cisco_ios_xr"
    perPrefix.EntityData.ParentYangName = "fast-reroute"
    perPrefix.EntityData.SegmentPath = "per-prefix"
    perPrefix.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/interfaces/interface/fast-reroute/" + perPrefix.EntityData.SegmentPath
    perPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    perPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    perPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    perPrefix.EntityData.Children = types.NewOrderedMap()
    perPrefix.EntityData.Children.Append("candidate-interfaces", types.YChild{"CandidateInterfaces", &perPrefix.CandidateInterfaces})
    perPrefix.EntityData.Children.Append("exclude-interfaces", types.YChild{"ExcludeInterfaces", &perPrefix.ExcludeInterfaces})
    perPrefix.EntityData.Leafs = types.NewOrderedMap()
    perPrefix.EntityData.Leafs.Append("fast-reroute-use-candidate-only", types.YLeaf{"FastRerouteUseCandidateOnly", perPrefix.FastRerouteUseCandidateOnly})

    perPrefix.EntityData.YListKeys = []string {}

    return &(perPrefix.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces
// Fast-reroute per-link/per-prefix candidate
// interface configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Candidate backup interface. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface.
    CandidateInterface []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface
}

func (candidateInterfaces *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces) GetEntityData() *types.CommonEntityData {
    candidateInterfaces.EntityData.YFilter = candidateInterfaces.YFilter
    candidateInterfaces.EntityData.YangName = "candidate-interfaces"
    candidateInterfaces.EntityData.BundleName = "cisco_ios_xr"
    candidateInterfaces.EntityData.ParentYangName = "per-prefix"
    candidateInterfaces.EntityData.SegmentPath = "candidate-interfaces"
    candidateInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/interfaces/interface/fast-reroute/per-prefix/" + candidateInterfaces.EntityData.SegmentPath
    candidateInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterfaces.EntityData.Children = types.NewOrderedMap()
    candidateInterfaces.EntityData.Children.Append("candidate-interface", types.YChild{"CandidateInterface", nil})
    for i := range candidateInterfaces.CandidateInterface {
        candidateInterfaces.EntityData.Children.Append(types.GetSegmentPath(candidateInterfaces.CandidateInterface[i]), types.YChild{"CandidateInterface", candidateInterfaces.CandidateInterface[i]})
    }
    candidateInterfaces.EntityData.Leafs = types.NewOrderedMap()

    candidateInterfaces.EntityData.YListKeys = []string {}

    return &(candidateInterfaces.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface
// Candidate backup interface
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (candidateInterface *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface) GetEntityData() *types.CommonEntityData {
    candidateInterface.EntityData.YFilter = candidateInterface.YFilter
    candidateInterface.EntityData.YangName = "candidate-interface"
    candidateInterface.EntityData.BundleName = "cisco_ios_xr"
    candidateInterface.EntityData.ParentYangName = "candidate-interfaces"
    candidateInterface.EntityData.SegmentPath = "candidate-interface" + types.AddKeyToken(candidateInterface.InterfaceName, "interface-name")
    candidateInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/interfaces/interface/fast-reroute/per-prefix/candidate-interfaces/" + candidateInterface.EntityData.SegmentPath
    candidateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterface.EntityData.Children = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", candidateInterface.InterfaceName})

    candidateInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(candidateInterface.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces
// Fast-reroute per-link/per-prefix exclude
// interface configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exclude an interface from becoming a backup. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface.
    ExcludeInterface []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface
}

func (excludeInterfaces *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces) GetEntityData() *types.CommonEntityData {
    excludeInterfaces.EntityData.YFilter = excludeInterfaces.YFilter
    excludeInterfaces.EntityData.YangName = "exclude-interfaces"
    excludeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    excludeInterfaces.EntityData.ParentYangName = "per-prefix"
    excludeInterfaces.EntityData.SegmentPath = "exclude-interfaces"
    excludeInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/interfaces/interface/fast-reroute/per-prefix/" + excludeInterfaces.EntityData.SegmentPath
    excludeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterfaces.EntityData.Children = types.NewOrderedMap()
    excludeInterfaces.EntityData.Children.Append("exclude-interface", types.YChild{"ExcludeInterface", nil})
    for i := range excludeInterfaces.ExcludeInterface {
        excludeInterfaces.EntityData.Children.Append(types.GetSegmentPath(excludeInterfaces.ExcludeInterface[i]), types.YChild{"ExcludeInterface", excludeInterfaces.ExcludeInterface[i]})
    }
    excludeInterfaces.EntityData.Leafs = types.NewOrderedMap()

    excludeInterfaces.EntityData.YListKeys = []string {}

    return &(excludeInterfaces.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface
// Exclude an interface from becoming a backup
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (excludeInterface *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface) GetEntityData() *types.CommonEntityData {
    excludeInterface.EntityData.YFilter = excludeInterface.YFilter
    excludeInterface.EntityData.YangName = "exclude-interface"
    excludeInterface.EntityData.BundleName = "cisco_ios_xr"
    excludeInterface.EntityData.ParentYangName = "exclude-interfaces"
    excludeInterface.EntityData.SegmentPath = "exclude-interface" + types.AddKeyToken(excludeInterface.InterfaceName, "interface-name")
    excludeInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/interfaces/interface/fast-reroute/per-prefix/exclude-interfaces/" + excludeInterface.EntityData.SegmentPath
    excludeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterface.EntityData.Children = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", excludeInterface.InterfaceName})

    excludeInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(excludeInterface.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope
// Area Scope Configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fast-reroute configuration.
    FastReroute Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute
}

func (areaScope *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope) GetEntityData() *types.CommonEntityData {
    areaScope.EntityData.YFilter = areaScope.YFilter
    areaScope.EntityData.YangName = "area-scope"
    areaScope.EntityData.BundleName = "cisco_ios_xr"
    areaScope.EntityData.ParentYangName = "area-area-id"
    areaScope.EntityData.SegmentPath = "area-scope"
    areaScope.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/" + areaScope.EntityData.SegmentPath
    areaScope.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    areaScope.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    areaScope.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    areaScope.EntityData.Children = types.NewOrderedMap()
    areaScope.EntityData.Children.Append("fast-reroute", types.YChild{"FastReroute", &areaScope.FastReroute})
    areaScope.EntityData.Leafs = types.NewOrderedMap()

    areaScope.EntityData.YListKeys = []string {}

    return &(areaScope.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute
// Fast-reroute configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable Fast-reroute per-link or per-prefix. The type is
    // Ospfv3FastReroute.
    FastRerouteEnable interface{}

    // Fast-reroute per-link configuration.
    PerLink Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink

    // Fast-reroute per-link configuration.
    PerPrefix Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix
}

func (fastReroute *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute) GetEntityData() *types.CommonEntityData {
    fastReroute.EntityData.YFilter = fastReroute.YFilter
    fastReroute.EntityData.YangName = "fast-reroute"
    fastReroute.EntityData.BundleName = "cisco_ios_xr"
    fastReroute.EntityData.ParentYangName = "area-scope"
    fastReroute.EntityData.SegmentPath = "fast-reroute"
    fastReroute.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/area-scope/" + fastReroute.EntityData.SegmentPath
    fastReroute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fastReroute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fastReroute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fastReroute.EntityData.Children = types.NewOrderedMap()
    fastReroute.EntityData.Children.Append("per-link", types.YChild{"PerLink", &fastReroute.PerLink})
    fastReroute.EntityData.Children.Append("per-prefix", types.YChild{"PerPrefix", &fastReroute.PerPrefix})
    fastReroute.EntityData.Leafs = types.NewOrderedMap()
    fastReroute.EntityData.Leafs.Append("fast-reroute-enable", types.YLeaf{"FastRerouteEnable", fastReroute.FastRerouteEnable})

    fastReroute.EntityData.YListKeys = []string {}

    return &(fastReroute.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink
// Fast-reroute per-link configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Use only interfaces on the candidate list as a backup path. The type is
    // bool. The default value is false.
    FastRerouteUseCandidateOnly interface{}

    // Fast-reroute per-link/per-prefix candidate interface configuration.
    CandidateInterfaces Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_CandidateInterfaces

    // Fast-reroute per-link/per-prefix exclude interface configuration.
    ExcludeInterfaces Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_ExcludeInterfaces
}

func (perLink *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink) GetEntityData() *types.CommonEntityData {
    perLink.EntityData.YFilter = perLink.YFilter
    perLink.EntityData.YangName = "per-link"
    perLink.EntityData.BundleName = "cisco_ios_xr"
    perLink.EntityData.ParentYangName = "fast-reroute"
    perLink.EntityData.SegmentPath = "per-link"
    perLink.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/area-scope/fast-reroute/" + perLink.EntityData.SegmentPath
    perLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    perLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    perLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    perLink.EntityData.Children = types.NewOrderedMap()
    perLink.EntityData.Children.Append("candidate-interfaces", types.YChild{"CandidateInterfaces", &perLink.CandidateInterfaces})
    perLink.EntityData.Children.Append("exclude-interfaces", types.YChild{"ExcludeInterfaces", &perLink.ExcludeInterfaces})
    perLink.EntityData.Leafs = types.NewOrderedMap()
    perLink.EntityData.Leafs.Append("fast-reroute-use-candidate-only", types.YLeaf{"FastRerouteUseCandidateOnly", perLink.FastRerouteUseCandidateOnly})

    perLink.EntityData.YListKeys = []string {}

    return &(perLink.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_CandidateInterfaces
// Fast-reroute per-link/per-prefix candidate
// interface configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_CandidateInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Candidate backup interface. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface.
    CandidateInterface []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface
}

func (candidateInterfaces *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_CandidateInterfaces) GetEntityData() *types.CommonEntityData {
    candidateInterfaces.EntityData.YFilter = candidateInterfaces.YFilter
    candidateInterfaces.EntityData.YangName = "candidate-interfaces"
    candidateInterfaces.EntityData.BundleName = "cisco_ios_xr"
    candidateInterfaces.EntityData.ParentYangName = "per-link"
    candidateInterfaces.EntityData.SegmentPath = "candidate-interfaces"
    candidateInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/area-scope/fast-reroute/per-link/" + candidateInterfaces.EntityData.SegmentPath
    candidateInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterfaces.EntityData.Children = types.NewOrderedMap()
    candidateInterfaces.EntityData.Children.Append("candidate-interface", types.YChild{"CandidateInterface", nil})
    for i := range candidateInterfaces.CandidateInterface {
        candidateInterfaces.EntityData.Children.Append(types.GetSegmentPath(candidateInterfaces.CandidateInterface[i]), types.YChild{"CandidateInterface", candidateInterfaces.CandidateInterface[i]})
    }
    candidateInterfaces.EntityData.Leafs = types.NewOrderedMap()

    candidateInterfaces.EntityData.YListKeys = []string {}

    return &(candidateInterfaces.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface
// Candidate backup interface
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (candidateInterface *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface) GetEntityData() *types.CommonEntityData {
    candidateInterface.EntityData.YFilter = candidateInterface.YFilter
    candidateInterface.EntityData.YangName = "candidate-interface"
    candidateInterface.EntityData.BundleName = "cisco_ios_xr"
    candidateInterface.EntityData.ParentYangName = "candidate-interfaces"
    candidateInterface.EntityData.SegmentPath = "candidate-interface" + types.AddKeyToken(candidateInterface.InterfaceName, "interface-name")
    candidateInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/area-scope/fast-reroute/per-link/candidate-interfaces/" + candidateInterface.EntityData.SegmentPath
    candidateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterface.EntityData.Children = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", candidateInterface.InterfaceName})

    candidateInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(candidateInterface.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_ExcludeInterfaces
// Fast-reroute per-link/per-prefix exclude
// interface configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_ExcludeInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exclude an interface from becoming a backup. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface.
    ExcludeInterface []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface
}

func (excludeInterfaces *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_ExcludeInterfaces) GetEntityData() *types.CommonEntityData {
    excludeInterfaces.EntityData.YFilter = excludeInterfaces.YFilter
    excludeInterfaces.EntityData.YangName = "exclude-interfaces"
    excludeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    excludeInterfaces.EntityData.ParentYangName = "per-link"
    excludeInterfaces.EntityData.SegmentPath = "exclude-interfaces"
    excludeInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/area-scope/fast-reroute/per-link/" + excludeInterfaces.EntityData.SegmentPath
    excludeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterfaces.EntityData.Children = types.NewOrderedMap()
    excludeInterfaces.EntityData.Children.Append("exclude-interface", types.YChild{"ExcludeInterface", nil})
    for i := range excludeInterfaces.ExcludeInterface {
        excludeInterfaces.EntityData.Children.Append(types.GetSegmentPath(excludeInterfaces.ExcludeInterface[i]), types.YChild{"ExcludeInterface", excludeInterfaces.ExcludeInterface[i]})
    }
    excludeInterfaces.EntityData.Leafs = types.NewOrderedMap()

    excludeInterfaces.EntityData.YListKeys = []string {}

    return &(excludeInterfaces.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface
// Exclude an interface from becoming a backup
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (excludeInterface *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface) GetEntityData() *types.CommonEntityData {
    excludeInterface.EntityData.YFilter = excludeInterface.YFilter
    excludeInterface.EntityData.YangName = "exclude-interface"
    excludeInterface.EntityData.BundleName = "cisco_ios_xr"
    excludeInterface.EntityData.ParentYangName = "exclude-interfaces"
    excludeInterface.EntityData.SegmentPath = "exclude-interface" + types.AddKeyToken(excludeInterface.InterfaceName, "interface-name")
    excludeInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/area-scope/fast-reroute/per-link/exclude-interfaces/" + excludeInterface.EntityData.SegmentPath
    excludeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterface.EntityData.Children = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", excludeInterface.InterfaceName})

    excludeInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(excludeInterface.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix
// Fast-reroute per-link configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Use only interfaces on the candidate list as a backup path. The type is
    // bool. The default value is false.
    FastRerouteUseCandidateOnly interface{}

    // Fast-reroute per-link/per-prefix candidate interface configuration.
    CandidateInterfaces Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_CandidateInterfaces

    // Fast-reroute per-link/per-prefix exclude interface configuration.
    ExcludeInterfaces Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces
}

func (perPrefix *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix) GetEntityData() *types.CommonEntityData {
    perPrefix.EntityData.YFilter = perPrefix.YFilter
    perPrefix.EntityData.YangName = "per-prefix"
    perPrefix.EntityData.BundleName = "cisco_ios_xr"
    perPrefix.EntityData.ParentYangName = "fast-reroute"
    perPrefix.EntityData.SegmentPath = "per-prefix"
    perPrefix.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/area-scope/fast-reroute/" + perPrefix.EntityData.SegmentPath
    perPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    perPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    perPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    perPrefix.EntityData.Children = types.NewOrderedMap()
    perPrefix.EntityData.Children.Append("candidate-interfaces", types.YChild{"CandidateInterfaces", &perPrefix.CandidateInterfaces})
    perPrefix.EntityData.Children.Append("exclude-interfaces", types.YChild{"ExcludeInterfaces", &perPrefix.ExcludeInterfaces})
    perPrefix.EntityData.Leafs = types.NewOrderedMap()
    perPrefix.EntityData.Leafs.Append("fast-reroute-use-candidate-only", types.YLeaf{"FastRerouteUseCandidateOnly", perPrefix.FastRerouteUseCandidateOnly})

    perPrefix.EntityData.YListKeys = []string {}

    return &(perPrefix.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_CandidateInterfaces
// Fast-reroute per-link/per-prefix candidate
// interface configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_CandidateInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Candidate backup interface. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface.
    CandidateInterface []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface
}

func (candidateInterfaces *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_CandidateInterfaces) GetEntityData() *types.CommonEntityData {
    candidateInterfaces.EntityData.YFilter = candidateInterfaces.YFilter
    candidateInterfaces.EntityData.YangName = "candidate-interfaces"
    candidateInterfaces.EntityData.BundleName = "cisco_ios_xr"
    candidateInterfaces.EntityData.ParentYangName = "per-prefix"
    candidateInterfaces.EntityData.SegmentPath = "candidate-interfaces"
    candidateInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/area-scope/fast-reroute/per-prefix/" + candidateInterfaces.EntityData.SegmentPath
    candidateInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterfaces.EntityData.Children = types.NewOrderedMap()
    candidateInterfaces.EntityData.Children.Append("candidate-interface", types.YChild{"CandidateInterface", nil})
    for i := range candidateInterfaces.CandidateInterface {
        candidateInterfaces.EntityData.Children.Append(types.GetSegmentPath(candidateInterfaces.CandidateInterface[i]), types.YChild{"CandidateInterface", candidateInterfaces.CandidateInterface[i]})
    }
    candidateInterfaces.EntityData.Leafs = types.NewOrderedMap()

    candidateInterfaces.EntityData.YListKeys = []string {}

    return &(candidateInterfaces.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface
// Candidate backup interface
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (candidateInterface *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface) GetEntityData() *types.CommonEntityData {
    candidateInterface.EntityData.YFilter = candidateInterface.YFilter
    candidateInterface.EntityData.YangName = "candidate-interface"
    candidateInterface.EntityData.BundleName = "cisco_ios_xr"
    candidateInterface.EntityData.ParentYangName = "candidate-interfaces"
    candidateInterface.EntityData.SegmentPath = "candidate-interface" + types.AddKeyToken(candidateInterface.InterfaceName, "interface-name")
    candidateInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/area-scope/fast-reroute/per-prefix/candidate-interfaces/" + candidateInterface.EntityData.SegmentPath
    candidateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterface.EntityData.Children = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", candidateInterface.InterfaceName})

    candidateInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(candidateInterface.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces
// Fast-reroute per-link/per-prefix exclude
// interface configuration
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exclude an interface from becoming a backup. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface.
    ExcludeInterface []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface
}

func (excludeInterfaces *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces) GetEntityData() *types.CommonEntityData {
    excludeInterfaces.EntityData.YFilter = excludeInterfaces.YFilter
    excludeInterfaces.EntityData.YangName = "exclude-interfaces"
    excludeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    excludeInterfaces.EntityData.ParentYangName = "per-prefix"
    excludeInterfaces.EntityData.SegmentPath = "exclude-interfaces"
    excludeInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/area-scope/fast-reroute/per-prefix/" + excludeInterfaces.EntityData.SegmentPath
    excludeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterfaces.EntityData.Children = types.NewOrderedMap()
    excludeInterfaces.EntityData.Children.Append("exclude-interface", types.YChild{"ExcludeInterface", nil})
    for i := range excludeInterfaces.ExcludeInterface {
        excludeInterfaces.EntityData.Children.Append(types.GetSegmentPath(excludeInterfaces.ExcludeInterface[i]), types.YChild{"ExcludeInterface", excludeInterfaces.ExcludeInterface[i]})
    }
    excludeInterfaces.EntityData.Leafs = types.NewOrderedMap()

    excludeInterfaces.EntityData.YListKeys = []string {}

    return &(excludeInterfaces.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface
// Exclude an interface from becoming a backup
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (excludeInterface *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface) GetEntityData() *types.CommonEntityData {
    excludeInterface.EntityData.YFilter = excludeInterface.YFilter
    excludeInterface.EntityData.YangName = "exclude-interface"
    excludeInterface.EntityData.BundleName = "cisco_ios_xr"
    excludeInterface.EntityData.ParentYangName = "exclude-interfaces"
    excludeInterface.EntityData.SegmentPath = "exclude-interface" + types.AddKeyToken(excludeInterface.InterfaceName, "interface-name")
    excludeInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/area-scope/fast-reroute/per-prefix/exclude-interfaces/" + excludeInterface.EntityData.SegmentPath
    excludeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterface.EntityData.Children = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", excludeInterface.InterfaceName})

    excludeInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(excludeInterface.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_ShamLinks
// Sham Link sub-mode
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_ShamLinks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ShamLink local and remote endpoints. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_ShamLinks_ShamLink.
    ShamLink []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_ShamLinks_ShamLink
}

func (shamLinks *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_ShamLinks) GetEntityData() *types.CommonEntityData {
    shamLinks.EntityData.YFilter = shamLinks.YFilter
    shamLinks.EntityData.YangName = "sham-links"
    shamLinks.EntityData.BundleName = "cisco_ios_xr"
    shamLinks.EntityData.ParentYangName = "area-area-id"
    shamLinks.EntityData.SegmentPath = "sham-links"
    shamLinks.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/" + shamLinks.EntityData.SegmentPath
    shamLinks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shamLinks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shamLinks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shamLinks.EntityData.Children = types.NewOrderedMap()
    shamLinks.EntityData.Children.Append("sham-link", types.YChild{"ShamLink", nil})
    for i := range shamLinks.ShamLink {
        shamLinks.EntityData.Children.Append(types.GetSegmentPath(shamLinks.ShamLink[i]), types.YChild{"ShamLink", shamLinks.ShamLink[i]})
    }
    shamLinks.EntityData.Leafs = types.NewOrderedMap()

    shamLinks.EntityData.YListKeys = []string {}

    return &(shamLinks.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_ShamLinks_ShamLink
// ShamLink local and remote endpoints
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_ShamLinks_ShamLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Local sham-link endpoint. The type is string with
    // pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // This attribute is a key. Remote sham-link endpoint. The type is string with
    // pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    DestinationAddress interface{}

    // Enable sham link. The type is interface{}.
    Enable interface{}

    // Time between HELLO packets. The type is interface{} with range: 1..65535.
    // Units are second.
    HelloInterval interface{}

    // Interval after which a neighbor is declared dead (in seconds). The type is
    // interface{} with range: 1..65535. Units are second.
    DeadInterval interface{}

    // Specify the transmit interval in seconds. The type is interface{} with
    // range: 1..65535. Units are second.
    RetransmitInterval interface{}

    // Specify the transmit delay in seconds. The type is interface{} with range:
    // 1..65535. Units are second.
    TransmitDelay interface{}

    // Authenticate OSPFv3 packets.
    Authentication Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_ShamLinks_ShamLink_Authentication

    // Encrypt and authenticate OSPFv3 packets.
    Encryption Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_ShamLinks_ShamLink_Encryption
}

func (shamLink *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_ShamLinks_ShamLink) GetEntityData() *types.CommonEntityData {
    shamLink.EntityData.YFilter = shamLink.YFilter
    shamLink.EntityData.YangName = "sham-link"
    shamLink.EntityData.BundleName = "cisco_ios_xr"
    shamLink.EntityData.ParentYangName = "sham-links"
    shamLink.EntityData.SegmentPath = "sham-link" + types.AddKeyToken(shamLink.SourceAddress, "source-address") + types.AddKeyToken(shamLink.DestinationAddress, "destination-address")
    shamLink.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/sham-links/" + shamLink.EntityData.SegmentPath
    shamLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shamLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shamLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shamLink.EntityData.Children = types.NewOrderedMap()
    shamLink.EntityData.Children.Append("authentication", types.YChild{"Authentication", &shamLink.Authentication})
    shamLink.EntityData.Children.Append("encryption", types.YChild{"Encryption", &shamLink.Encryption})
    shamLink.EntityData.Leafs = types.NewOrderedMap()
    shamLink.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", shamLink.SourceAddress})
    shamLink.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", shamLink.DestinationAddress})
    shamLink.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", shamLink.Enable})
    shamLink.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", shamLink.HelloInterval})
    shamLink.EntityData.Leafs.Append("dead-interval", types.YLeaf{"DeadInterval", shamLink.DeadInterval})
    shamLink.EntityData.Leafs.Append("retransmit-interval", types.YLeaf{"RetransmitInterval", shamLink.RetransmitInterval})
    shamLink.EntityData.Leafs.Append("transmit-delay", types.YLeaf{"TransmitDelay", shamLink.TransmitDelay})

    shamLink.EntityData.YListKeys = []string {"SourceAddress", "DestinationAddress"}

    return &(shamLink.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_ShamLinks_ShamLink_Authentication
// Authenticate OSPFv3 packets
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_ShamLinks_ShamLink_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec AH authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Use the MD5 or SHA1 algorithm. The type is Ospfv3Authentication.
    Algorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    Password interface{}
}

func (authentication *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_ShamLinks_ShamLink_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "sham-link"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/sham-links/sham-link/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", authentication.Enable})
    authentication.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", authentication.Spi})
    authentication.EntityData.Leafs.Append("algorithm", types.YLeaf{"Algorithm", authentication.Algorithm})
    authentication.EntityData.Leafs.Append("password", types.YLeaf{"Password", authentication.Password})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_ShamLinks_ShamLink_Encryption
// Encrypt and authenticate OSPFv3 packets
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_ShamLinks_ShamLink_Encryption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec ESP authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Specify the encryption algorithm. The type is Ospfv3EncryptionAlgorithm.
    EncryptionAlgorithm interface{}

    // Encryption password. The type is string with pattern: b'(!.+)|([^!].+)'.
    EncryptionPassword interface{}

    // Use the NULL, MD5 or SHA1 algorithm. The type is Ospfv3AuthenticationType2.
    AuthenticationAlgorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    AuthenticationPassword interface{}
}

func (encryption *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_ShamLinks_ShamLink_Encryption) GetEntityData() *types.CommonEntityData {
    encryption.EntityData.YFilter = encryption.YFilter
    encryption.EntityData.YangName = "encryption"
    encryption.EntityData.BundleName = "cisco_ios_xr"
    encryption.EntityData.ParentYangName = "sham-link"
    encryption.EntityData.SegmentPath = "encryption"
    encryption.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/sham-links/sham-link/" + encryption.EntityData.SegmentPath
    encryption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryption.EntityData.Children = types.NewOrderedMap()
    encryption.EntityData.Leafs = types.NewOrderedMap()
    encryption.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", encryption.Enable})
    encryption.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", encryption.Spi})
    encryption.EntityData.Leafs.Append("encryption-algorithm", types.YLeaf{"EncryptionAlgorithm", encryption.EncryptionAlgorithm})
    encryption.EntityData.Leafs.Append("encryption-password", types.YLeaf{"EncryptionPassword", encryption.EncryptionPassword})
    encryption.EntityData.Leafs.Append("authentication-algorithm", types.YLeaf{"AuthenticationAlgorithm", encryption.AuthenticationAlgorithm})
    encryption.EntityData.Leafs.Append("authentication-password", types.YLeaf{"AuthenticationPassword", encryption.AuthenticationPassword})

    encryption.EntityData.YListKeys = []string {}

    return &(encryption.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_VirtualLinks
// Virtual link sub-mode
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_VirtualLinks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router ID of virtual link neighbor. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_VirtualLinks_VirtualLink.
    VirtualLink []*Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_VirtualLinks_VirtualLink
}

func (virtualLinks *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_VirtualLinks) GetEntityData() *types.CommonEntityData {
    virtualLinks.EntityData.YFilter = virtualLinks.YFilter
    virtualLinks.EntityData.YangName = "virtual-links"
    virtualLinks.EntityData.BundleName = "cisco_ios_xr"
    virtualLinks.EntityData.ParentYangName = "area-area-id"
    virtualLinks.EntityData.SegmentPath = "virtual-links"
    virtualLinks.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/" + virtualLinks.EntityData.SegmentPath
    virtualLinks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    virtualLinks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    virtualLinks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    virtualLinks.EntityData.Children = types.NewOrderedMap()
    virtualLinks.EntityData.Children.Append("virtual-link", types.YChild{"VirtualLink", nil})
    for i := range virtualLinks.VirtualLink {
        virtualLinks.EntityData.Children.Append(types.GetSegmentPath(virtualLinks.VirtualLink[i]), types.YChild{"VirtualLink", virtualLinks.VirtualLink[i]})
    }
    virtualLinks.EntityData.Leafs = types.NewOrderedMap()

    virtualLinks.EntityData.YListKeys = []string {}

    return &(virtualLinks.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_VirtualLinks_VirtualLink
// Router ID of virtual link neighbor
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_VirtualLinks_VirtualLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Router ID of virtual link neighbor. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    VirtualLinkAddress interface{}

    // Enabled virtual link. The type is interface{}.
    Enable interface{}

    // Time between HELLO packets. The type is interface{} with range: 1..65535.
    // Units are second.
    HelloInterval interface{}

    // Interval after which a neighbor is declared dead (in seconds). The type is
    // interface{} with range: 1..65535. Units are second.
    DeadInterval interface{}

    // Specify the transmit interval in seconds. The type is interface{} with
    // range: 1..65535. Units are second.
    RetransmitInterval interface{}

    // Specify the transmit delay in seconds. The type is interface{} with range:
    // 1..65535. Units are second.
    TransmitDelay interface{}

    // Authenticate OSPFv3 packets.
    Authentication Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_VirtualLinks_VirtualLink_Authentication

    // Encrypt and authenticate OSPFv3 packets.
    Encryption Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_VirtualLinks_VirtualLink_Encryption
}

func (virtualLink *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_VirtualLinks_VirtualLink) GetEntityData() *types.CommonEntityData {
    virtualLink.EntityData.YFilter = virtualLink.YFilter
    virtualLink.EntityData.YangName = "virtual-link"
    virtualLink.EntityData.BundleName = "cisco_ios_xr"
    virtualLink.EntityData.ParentYangName = "virtual-links"
    virtualLink.EntityData.SegmentPath = "virtual-link" + types.AddKeyToken(virtualLink.VirtualLinkAddress, "virtual-link-address")
    virtualLink.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/virtual-links/" + virtualLink.EntityData.SegmentPath
    virtualLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    virtualLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    virtualLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    virtualLink.EntityData.Children = types.NewOrderedMap()
    virtualLink.EntityData.Children.Append("authentication", types.YChild{"Authentication", &virtualLink.Authentication})
    virtualLink.EntityData.Children.Append("encryption", types.YChild{"Encryption", &virtualLink.Encryption})
    virtualLink.EntityData.Leafs = types.NewOrderedMap()
    virtualLink.EntityData.Leafs.Append("virtual-link-address", types.YLeaf{"VirtualLinkAddress", virtualLink.VirtualLinkAddress})
    virtualLink.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", virtualLink.Enable})
    virtualLink.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", virtualLink.HelloInterval})
    virtualLink.EntityData.Leafs.Append("dead-interval", types.YLeaf{"DeadInterval", virtualLink.DeadInterval})
    virtualLink.EntityData.Leafs.Append("retransmit-interval", types.YLeaf{"RetransmitInterval", virtualLink.RetransmitInterval})
    virtualLink.EntityData.Leafs.Append("transmit-delay", types.YLeaf{"TransmitDelay", virtualLink.TransmitDelay})

    virtualLink.EntityData.YListKeys = []string {"VirtualLinkAddress"}

    return &(virtualLink.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_VirtualLinks_VirtualLink_Authentication
// Authenticate OSPFv3 packets
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_VirtualLinks_VirtualLink_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec AH authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Use the MD5 or SHA1 algorithm. The type is Ospfv3Authentication.
    Algorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    Password interface{}
}

func (authentication *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_VirtualLinks_VirtualLink_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "virtual-link"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/virtual-links/virtual-link/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", authentication.Enable})
    authentication.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", authentication.Spi})
    authentication.EntityData.Leafs.Append("algorithm", types.YLeaf{"Algorithm", authentication.Algorithm})
    authentication.EntityData.Leafs.Append("password", types.YLeaf{"Password", authentication.Password})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_VirtualLinks_VirtualLink_Encryption
// Encrypt and authenticate OSPFv3 packets
type Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_VirtualLinks_VirtualLink_Encryption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec ESP authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Specify the encryption algorithm. The type is Ospfv3EncryptionAlgorithm.
    EncryptionAlgorithm interface{}

    // Encryption password. The type is string with pattern: b'(!.+)|([^!].+)'.
    EncryptionPassword interface{}

    // Use the NULL, MD5 or SHA1 algorithm. The type is Ospfv3AuthenticationType2.
    AuthenticationAlgorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    AuthenticationPassword interface{}
}

func (encryption *Ospfv3_Processes_Process_DefaultVrf_AreaAddresses_AreaAreaId_VirtualLinks_VirtualLink_Encryption) GetEntityData() *types.CommonEntityData {
    encryption.EntityData.YFilter = encryption.YFilter
    encryption.EntityData.YangName = "encryption"
    encryption.EntityData.BundleName = "cisco_ios_xr"
    encryption.EntityData.ParentYangName = "virtual-link"
    encryption.EntityData.SegmentPath = "encryption"
    encryption.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/area-addresses/area-area-id/virtual-links/virtual-link/" + encryption.EntityData.SegmentPath
    encryption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryption.EntityData.Children = types.NewOrderedMap()
    encryption.EntityData.Leafs = types.NewOrderedMap()
    encryption.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", encryption.Enable})
    encryption.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", encryption.Spi})
    encryption.EntityData.Leafs.Append("encryption-algorithm", types.YLeaf{"EncryptionAlgorithm", encryption.EncryptionAlgorithm})
    encryption.EntityData.Leafs.Append("encryption-password", types.YLeaf{"EncryptionPassword", encryption.EncryptionPassword})
    encryption.EntityData.Leafs.Append("authentication-algorithm", types.YLeaf{"AuthenticationAlgorithm", encryption.AuthenticationAlgorithm})
    encryption.EntityData.Leafs.Append("authentication-password", types.YLeaf{"AuthenticationPassword", encryption.AuthenticationPassword})

    encryption.EntityData.YListKeys = []string {}

    return &(encryption.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_Timers
// Adjust routing timers
type Ospfv3_Processes_Process_DefaultVrf_Timers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pacing timers.
    Pacing Ospfv3_Processes_Process_DefaultVrf_Timers_Pacing

    // LSA timers.
    LsaTimers Ospfv3_Processes_Process_DefaultVrf_Timers_LsaTimers

    // Throttle timers.
    Throttle Ospfv3_Processes_Process_DefaultVrf_Timers_Throttle
}

func (timers *Ospfv3_Processes_Process_DefaultVrf_Timers) GetEntityData() *types.CommonEntityData {
    timers.EntityData.YFilter = timers.YFilter
    timers.EntityData.YangName = "timers"
    timers.EntityData.BundleName = "cisco_ios_xr"
    timers.EntityData.ParentYangName = "default-vrf"
    timers.EntityData.SegmentPath = "timers"
    timers.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/" + timers.EntityData.SegmentPath
    timers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timers.EntityData.Children = types.NewOrderedMap()
    timers.EntityData.Children.Append("pacing", types.YChild{"Pacing", &timers.Pacing})
    timers.EntityData.Children.Append("lsa-timers", types.YChild{"LsaTimers", &timers.LsaTimers})
    timers.EntityData.Children.Append("throttle", types.YChild{"Throttle", &timers.Throttle})
    timers.EntityData.Leafs = types.NewOrderedMap()

    timers.EntityData.YListKeys = []string {}

    return &(timers.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_Timers_Pacing
// Pacing timers
type Ospfv3_Processes_Process_DefaultVrf_Timers_Pacing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The minimum interval in milliseconds to pace limit flooding on interface.
    // The type is interface{} with range: 5..100. Units are millisecond.
    Flood interface{}

    // The minimum interval in msec between neighbor retransmissions. The type is
    // interface{} with range: 5..100.
    Retransmission interface{}

    // Interval in seconds at which LSAs are grouped and refreshed, checksummed,
    // or aged. The type is interface{} with range: 10..1800. Units are second.
    LsaGroup interface{}
}

func (pacing *Ospfv3_Processes_Process_DefaultVrf_Timers_Pacing) GetEntityData() *types.CommonEntityData {
    pacing.EntityData.YFilter = pacing.YFilter
    pacing.EntityData.YangName = "pacing"
    pacing.EntityData.BundleName = "cisco_ios_xr"
    pacing.EntityData.ParentYangName = "timers"
    pacing.EntityData.SegmentPath = "pacing"
    pacing.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/timers/" + pacing.EntityData.SegmentPath
    pacing.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pacing.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pacing.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pacing.EntityData.Children = types.NewOrderedMap()
    pacing.EntityData.Leafs = types.NewOrderedMap()
    pacing.EntityData.Leafs.Append("flood", types.YLeaf{"Flood", pacing.Flood})
    pacing.EntityData.Leafs.Append("retransmission", types.YLeaf{"Retransmission", pacing.Retransmission})
    pacing.EntityData.Leafs.Append("lsa-group", types.YLeaf{"LsaGroup", pacing.LsaGroup})

    pacing.EntityData.YListKeys = []string {}

    return &(pacing.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_Timers_LsaTimers
// LSA timers
type Ospfv3_Processes_Process_DefaultVrf_Timers_LsaTimers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The minimum interval in milliseconds between accepting the same LSA. The
    // type is interface{} with range: 0..60000. Units are millisecond.
    Arrival interface{}
}

func (lsaTimers *Ospfv3_Processes_Process_DefaultVrf_Timers_LsaTimers) GetEntityData() *types.CommonEntityData {
    lsaTimers.EntityData.YFilter = lsaTimers.YFilter
    lsaTimers.EntityData.YangName = "lsa-timers"
    lsaTimers.EntityData.BundleName = "cisco_ios_xr"
    lsaTimers.EntityData.ParentYangName = "timers"
    lsaTimers.EntityData.SegmentPath = "lsa-timers"
    lsaTimers.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/timers/" + lsaTimers.EntityData.SegmentPath
    lsaTimers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsaTimers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsaTimers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsaTimers.EntityData.Children = types.NewOrderedMap()
    lsaTimers.EntityData.Leafs = types.NewOrderedMap()
    lsaTimers.EntityData.Leafs.Append("arrival", types.YLeaf{"Arrival", lsaTimers.Arrival})

    lsaTimers.EntityData.YListKeys = []string {}

    return &(lsaTimers.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_Timers_Throttle
// Throttle timers
type Ospfv3_Processes_Process_DefaultVrf_Timers_Throttle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA throttle timers for all types of OSPF LSAs.
    Lsa Ospfv3_Processes_Process_DefaultVrf_Timers_Throttle_Lsa

    // SPF throttle timers.
    Spf Ospfv3_Processes_Process_DefaultVrf_Timers_Throttle_Spf
}

func (throttle *Ospfv3_Processes_Process_DefaultVrf_Timers_Throttle) GetEntityData() *types.CommonEntityData {
    throttle.EntityData.YFilter = throttle.YFilter
    throttle.EntityData.YangName = "throttle"
    throttle.EntityData.BundleName = "cisco_ios_xr"
    throttle.EntityData.ParentYangName = "timers"
    throttle.EntityData.SegmentPath = "throttle"
    throttle.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/timers/" + throttle.EntityData.SegmentPath
    throttle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    throttle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    throttle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    throttle.EntityData.Children = types.NewOrderedMap()
    throttle.EntityData.Children.Append("lsa", types.YChild{"Lsa", &throttle.Lsa})
    throttle.EntityData.Children.Append("spf", types.YChild{"Spf", &throttle.Spf})
    throttle.EntityData.Leafs = types.NewOrderedMap()

    throttle.EntityData.YListKeys = []string {}

    return &(throttle.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_Timers_Throttle_Lsa
// LSA throttle timers for all types of OSPF LSAs
type Ospfv3_Processes_Process_DefaultVrf_Timers_Throttle_Lsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Delay to generate first occurrence of LSA in milliseconds. The type is
    // interface{} with range: 0..600000. Units are millisecond.
    FirstDelay interface{}

    // Minimum delay between originating the same LSA in milliseconds. The type is
    // interface{} with range: 1..600000. Units are millisecond.
    MinimumDelay interface{}

    // Maximum delay between originating the same LSA in milliseconds. The type is
    // interface{} with range: 1..600000. Units are millisecond.
    MaximumDelay interface{}
}

func (lsa *Ospfv3_Processes_Process_DefaultVrf_Timers_Throttle_Lsa) GetEntityData() *types.CommonEntityData {
    lsa.EntityData.YFilter = lsa.YFilter
    lsa.EntityData.YangName = "lsa"
    lsa.EntityData.BundleName = "cisco_ios_xr"
    lsa.EntityData.ParentYangName = "throttle"
    lsa.EntityData.SegmentPath = "lsa"
    lsa.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/timers/throttle/" + lsa.EntityData.SegmentPath
    lsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsa.EntityData.Children = types.NewOrderedMap()
    lsa.EntityData.Leafs = types.NewOrderedMap()
    lsa.EntityData.Leafs.Append("first-delay", types.YLeaf{"FirstDelay", lsa.FirstDelay})
    lsa.EntityData.Leafs.Append("minimum-delay", types.YLeaf{"MinimumDelay", lsa.MinimumDelay})
    lsa.EntityData.Leafs.Append("maximum-delay", types.YLeaf{"MaximumDelay", lsa.MaximumDelay})

    lsa.EntityData.YListKeys = []string {}

    return &(lsa.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_Timers_Throttle_Spf
// SPF throttle timers
type Ospfv3_Processes_Process_DefaultVrf_Timers_Throttle_Spf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Initial delay between receiving a change and starting SPF in ms. The type
    // is interface{} with range: 1..600000.
    FirstDelay interface{}

    // Minimum hold time between consecutive SPF calculations in ms. The type is
    // interface{} with range: 1..600000.
    MinimumDelay interface{}

    // Maximum wait time between consecutive SPF calculations in ms. The type is
    // interface{} with range: 1..600000.
    MaximumDelay interface{}
}

func (spf *Ospfv3_Processes_Process_DefaultVrf_Timers_Throttle_Spf) GetEntityData() *types.CommonEntityData {
    spf.EntityData.YFilter = spf.YFilter
    spf.EntityData.YangName = "spf"
    spf.EntityData.BundleName = "cisco_ios_xr"
    spf.EntityData.ParentYangName = "throttle"
    spf.EntityData.SegmentPath = "spf"
    spf.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/timers/throttle/" + spf.EntityData.SegmentPath
    spf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spf.EntityData.Children = types.NewOrderedMap()
    spf.EntityData.Leafs = types.NewOrderedMap()
    spf.EntityData.Leafs.Append("first-delay", types.YLeaf{"FirstDelay", spf.FirstDelay})
    spf.EntityData.Leafs.Append("minimum-delay", types.YLeaf{"MinimumDelay", spf.MinimumDelay})
    spf.EntityData.Leafs.Append("maximum-delay", types.YLeaf{"MaximumDelay", spf.MaximumDelay})

    spf.EntityData.YListKeys = []string {}

    return &(spf.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_SummaryPrefixes
// Summarize redistributed routes matching
// prefix/length
type Ospfv3_Processes_Process_DefaultVrf_SummaryPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 address. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_SummaryPrefixes_SummaryPrefix.
    SummaryPrefix []*Ospfv3_Processes_Process_DefaultVrf_SummaryPrefixes_SummaryPrefix
}

func (summaryPrefixes *Ospfv3_Processes_Process_DefaultVrf_SummaryPrefixes) GetEntityData() *types.CommonEntityData {
    summaryPrefixes.EntityData.YFilter = summaryPrefixes.YFilter
    summaryPrefixes.EntityData.YangName = "summary-prefixes"
    summaryPrefixes.EntityData.BundleName = "cisco_ios_xr"
    summaryPrefixes.EntityData.ParentYangName = "default-vrf"
    summaryPrefixes.EntityData.SegmentPath = "summary-prefixes"
    summaryPrefixes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/" + summaryPrefixes.EntityData.SegmentPath
    summaryPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaryPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaryPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaryPrefixes.EntityData.Children = types.NewOrderedMap()
    summaryPrefixes.EntityData.Children.Append("summary-prefix", types.YChild{"SummaryPrefix", nil})
    for i := range summaryPrefixes.SummaryPrefix {
        summaryPrefixes.EntityData.Children.Append(types.GetSegmentPath(summaryPrefixes.SummaryPrefix[i]), types.YChild{"SummaryPrefix", summaryPrefixes.SummaryPrefix[i]})
    }
    summaryPrefixes.EntityData.Leafs = types.NewOrderedMap()

    summaryPrefixes.EntityData.YListKeys = []string {}

    return &(summaryPrefixes.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_SummaryPrefixes_SummaryPrefix
// IPv6 address
type Ospfv3_Processes_Process_DefaultVrf_SummaryPrefixes_SummaryPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IPv6 prefix string format. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Prefix interface{}

    // This attribute is a key. IPV6 prefix length. The type is interface{} with
    // range: 0..128.
    PrefixLength interface{}

    // Suppress routes matching prefix/length. The type is bool.
    NotAdvertise interface{}

    // Tag. The type is interface{} with range: 1..4294967295.
    Tag interface{}
}

func (summaryPrefix *Ospfv3_Processes_Process_DefaultVrf_SummaryPrefixes_SummaryPrefix) GetEntityData() *types.CommonEntityData {
    summaryPrefix.EntityData.YFilter = summaryPrefix.YFilter
    summaryPrefix.EntityData.YangName = "summary-prefix"
    summaryPrefix.EntityData.BundleName = "cisco_ios_xr"
    summaryPrefix.EntityData.ParentYangName = "summary-prefixes"
    summaryPrefix.EntityData.SegmentPath = "summary-prefix" + types.AddKeyToken(summaryPrefix.Prefix, "prefix") + types.AddKeyToken(summaryPrefix.PrefixLength, "prefix-length")
    summaryPrefix.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/summary-prefixes/" + summaryPrefix.EntityData.SegmentPath
    summaryPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaryPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaryPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaryPrefix.EntityData.Children = types.NewOrderedMap()
    summaryPrefix.EntityData.Leafs = types.NewOrderedMap()
    summaryPrefix.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", summaryPrefix.Prefix})
    summaryPrefix.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", summaryPrefix.PrefixLength})
    summaryPrefix.EntityData.Leafs.Append("not-advertise", types.YLeaf{"NotAdvertise", summaryPrefix.NotAdvertise})
    summaryPrefix.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", summaryPrefix.Tag})

    summaryPrefix.EntityData.YListKeys = []string {"Prefix", "PrefixLength"}

    return &(summaryPrefix.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_Snmp
// SNMP configuration
type Ospfv3_Processes_Process_DefaultVrf_Snmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP context configuration. The type is string.
    Context interface{}

    // SNMP trap rate configuration.
    TrapRateLimit Ospfv3_Processes_Process_DefaultVrf_Snmp_TrapRateLimit
}

func (snmp *Ospfv3_Processes_Process_DefaultVrf_Snmp) GetEntityData() *types.CommonEntityData {
    snmp.EntityData.YFilter = snmp.YFilter
    snmp.EntityData.YangName = "snmp"
    snmp.EntityData.BundleName = "cisco_ios_xr"
    snmp.EntityData.ParentYangName = "default-vrf"
    snmp.EntityData.SegmentPath = "snmp"
    snmp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/" + snmp.EntityData.SegmentPath
    snmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmp.EntityData.Children = types.NewOrderedMap()
    snmp.EntityData.Children.Append("trap-rate-limit", types.YChild{"TrapRateLimit", &snmp.TrapRateLimit})
    snmp.EntityData.Leafs = types.NewOrderedMap()
    snmp.EntityData.Leafs.Append("context", types.YLeaf{"Context", snmp.Context})

    snmp.EntityData.YListKeys = []string {}

    return &(snmp.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_Snmp_TrapRateLimit
// SNMP trap rate configuration
type Ospfv3_Processes_Process_DefaultVrf_Snmp_TrapRateLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Trap rate limit sliding window size. The type is interface{} with range:
    // 2..60.
    WindowSize interface{}

    // Max number of traps sent in window time. The type is interface{} with
    // range: 0..300.
    MaxWindowTraps interface{}
}

func (trapRateLimit *Ospfv3_Processes_Process_DefaultVrf_Snmp_TrapRateLimit) GetEntityData() *types.CommonEntityData {
    trapRateLimit.EntityData.YFilter = trapRateLimit.YFilter
    trapRateLimit.EntityData.YangName = "trap-rate-limit"
    trapRateLimit.EntityData.BundleName = "cisco_ios_xr"
    trapRateLimit.EntityData.ParentYangName = "snmp"
    trapRateLimit.EntityData.SegmentPath = "trap-rate-limit"
    trapRateLimit.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/snmp/" + trapRateLimit.EntityData.SegmentPath
    trapRateLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapRateLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapRateLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapRateLimit.EntityData.Children = types.NewOrderedMap()
    trapRateLimit.EntityData.Leafs = types.NewOrderedMap()
    trapRateLimit.EntityData.Leafs.Append("window-size", types.YLeaf{"WindowSize", trapRateLimit.WindowSize})
    trapRateLimit.EntityData.Leafs.Append("max-window-traps", types.YLeaf{"MaxWindowTraps", trapRateLimit.MaxWindowTraps})

    trapRateLimit.EntityData.YListKeys = []string {}

    return &(trapRateLimit.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_FastReroute
// Fast-reroute instance scoped parameters
type Ospfv3_Processes_Process_DefaultVrf_FastReroute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fast-reroute per-link global configuration.
    PerLink Ospfv3_Processes_Process_DefaultVrf_FastReroute_PerLink

    // Fast-reroute per-prefix global configuration.
    PerPrefix Ospfv3_Processes_Process_DefaultVrf_FastReroute_PerPrefix
}

func (fastReroute *Ospfv3_Processes_Process_DefaultVrf_FastReroute) GetEntityData() *types.CommonEntityData {
    fastReroute.EntityData.YFilter = fastReroute.YFilter
    fastReroute.EntityData.YangName = "fast-reroute"
    fastReroute.EntityData.BundleName = "cisco_ios_xr"
    fastReroute.EntityData.ParentYangName = "default-vrf"
    fastReroute.EntityData.SegmentPath = "fast-reroute"
    fastReroute.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/" + fastReroute.EntityData.SegmentPath
    fastReroute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fastReroute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fastReroute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fastReroute.EntityData.Children = types.NewOrderedMap()
    fastReroute.EntityData.Children.Append("per-link", types.YChild{"PerLink", &fastReroute.PerLink})
    fastReroute.EntityData.Children.Append("per-prefix", types.YChild{"PerPrefix", &fastReroute.PerPrefix})
    fastReroute.EntityData.Leafs = types.NewOrderedMap()

    fastReroute.EntityData.YListKeys = []string {}

    return &(fastReroute.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_FastReroute_PerLink
// Fast-reroute per-link global configuration
type Ospfv3_Processes_Process_DefaultVrf_FastReroute_PerLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fast-reroute per-link/per-prefix priority-limit command. The type is
    // Ospfv3FastReroutePriority.
    Priority interface{}
}

func (perLink *Ospfv3_Processes_Process_DefaultVrf_FastReroute_PerLink) GetEntityData() *types.CommonEntityData {
    perLink.EntityData.YFilter = perLink.YFilter
    perLink.EntityData.YangName = "per-link"
    perLink.EntityData.BundleName = "cisco_ios_xr"
    perLink.EntityData.ParentYangName = "fast-reroute"
    perLink.EntityData.SegmentPath = "per-link"
    perLink.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/fast-reroute/" + perLink.EntityData.SegmentPath
    perLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    perLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    perLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    perLink.EntityData.Children = types.NewOrderedMap()
    perLink.EntityData.Leafs = types.NewOrderedMap()
    perLink.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", perLink.Priority})

    perLink.EntityData.YListKeys = []string {}

    return &(perLink.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_FastReroute_PerPrefix
// Fast-reroute per-prefix global configuration
type Ospfv3_Processes_Process_DefaultVrf_FastReroute_PerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable load sharing between multiple backups. The type is interface{}.
    LoadSharingDisable interface{}

    // Fast-reroute per-link/per-prefix priority-limit command. The type is
    // Ospfv3FastReroutePriority.
    Priority interface{}

    // Fast-reroute tiebreakers configurations.
    Tiebreakers Ospfv3_Processes_Process_DefaultVrf_FastReroute_PerPrefix_Tiebreakers
}

func (perPrefix *Ospfv3_Processes_Process_DefaultVrf_FastReroute_PerPrefix) GetEntityData() *types.CommonEntityData {
    perPrefix.EntityData.YFilter = perPrefix.YFilter
    perPrefix.EntityData.YangName = "per-prefix"
    perPrefix.EntityData.BundleName = "cisco_ios_xr"
    perPrefix.EntityData.ParentYangName = "fast-reroute"
    perPrefix.EntityData.SegmentPath = "per-prefix"
    perPrefix.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/fast-reroute/" + perPrefix.EntityData.SegmentPath
    perPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    perPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    perPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    perPrefix.EntityData.Children = types.NewOrderedMap()
    perPrefix.EntityData.Children.Append("tiebreakers", types.YChild{"Tiebreakers", &perPrefix.Tiebreakers})
    perPrefix.EntityData.Leafs = types.NewOrderedMap()
    perPrefix.EntityData.Leafs.Append("load-sharing-disable", types.YLeaf{"LoadSharingDisable", perPrefix.LoadSharingDisable})
    perPrefix.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", perPrefix.Priority})

    perPrefix.EntityData.YListKeys = []string {}

    return &(perPrefix.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_FastReroute_PerPrefix_Tiebreakers
// Fast-reroute tiebreakers configurations
type Ospfv3_Processes_Process_DefaultVrf_FastReroute_PerPrefix_Tiebreakers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fast-reroute tiebreakers configuration. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_FastReroute_PerPrefix_Tiebreakers_Tiebreaker.
    Tiebreaker []*Ospfv3_Processes_Process_DefaultVrf_FastReroute_PerPrefix_Tiebreakers_Tiebreaker
}

func (tiebreakers *Ospfv3_Processes_Process_DefaultVrf_FastReroute_PerPrefix_Tiebreakers) GetEntityData() *types.CommonEntityData {
    tiebreakers.EntityData.YFilter = tiebreakers.YFilter
    tiebreakers.EntityData.YangName = "tiebreakers"
    tiebreakers.EntityData.BundleName = "cisco_ios_xr"
    tiebreakers.EntityData.ParentYangName = "per-prefix"
    tiebreakers.EntityData.SegmentPath = "tiebreakers"
    tiebreakers.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/fast-reroute/per-prefix/" + tiebreakers.EntityData.SegmentPath
    tiebreakers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tiebreakers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tiebreakers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tiebreakers.EntityData.Children = types.NewOrderedMap()
    tiebreakers.EntityData.Children.Append("tiebreaker", types.YChild{"Tiebreaker", nil})
    for i := range tiebreakers.Tiebreaker {
        tiebreakers.EntityData.Children.Append(types.GetSegmentPath(tiebreakers.Tiebreaker[i]), types.YChild{"Tiebreaker", tiebreakers.Tiebreaker[i]})
    }
    tiebreakers.EntityData.Leafs = types.NewOrderedMap()

    tiebreakers.EntityData.YListKeys = []string {}

    return &(tiebreakers.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_FastReroute_PerPrefix_Tiebreakers_Tiebreaker
// Fast-reroute tiebreakers configuration
type Ospfv3_Processes_Process_DefaultVrf_FastReroute_PerPrefix_Tiebreakers_Tiebreaker struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Tiebreaker type. The type is
    // Ospfv3FastRerouteTiebreakers.
    TiebreakerType interface{}

    // Index value for a tiebreaker. The type is interface{} with range: 1..255.
    // This attribute is mandatory.
    TiebreakerIndex interface{}
}

func (tiebreaker *Ospfv3_Processes_Process_DefaultVrf_FastReroute_PerPrefix_Tiebreakers_Tiebreaker) GetEntityData() *types.CommonEntityData {
    tiebreaker.EntityData.YFilter = tiebreaker.YFilter
    tiebreaker.EntityData.YangName = "tiebreaker"
    tiebreaker.EntityData.BundleName = "cisco_ios_xr"
    tiebreaker.EntityData.ParentYangName = "tiebreakers"
    tiebreaker.EntityData.SegmentPath = "tiebreaker" + types.AddKeyToken(tiebreaker.TiebreakerType, "tiebreaker-type")
    tiebreaker.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/fast-reroute/per-prefix/tiebreakers/" + tiebreaker.EntityData.SegmentPath
    tiebreaker.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tiebreaker.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tiebreaker.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tiebreaker.EntityData.Children = types.NewOrderedMap()
    tiebreaker.EntityData.Leafs = types.NewOrderedMap()
    tiebreaker.EntityData.Leafs.Append("tiebreaker-type", types.YLeaf{"TiebreakerType", tiebreaker.TiebreakerType})
    tiebreaker.EntityData.Leafs.Append("tiebreaker-index", types.YLeaf{"TiebreakerIndex", tiebreaker.TiebreakerIndex})

    tiebreaker.EntityData.YListKeys = []string {"TiebreakerType"}

    return &(tiebreaker.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_Distance
// Define an administrative distance
type Ospfv3_Processes_Process_DefaultVrf_Distance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Define an administrative distance. The type is interface{} with range:
    // 1..255.
    Administrative interface{}

    // OSPFv3 administrative distance.
    Ospfv3 Ospfv3_Processes_Process_DefaultVrf_Distance_Ospfv3
}

func (distance *Ospfv3_Processes_Process_DefaultVrf_Distance) GetEntityData() *types.CommonEntityData {
    distance.EntityData.YFilter = distance.YFilter
    distance.EntityData.YangName = "distance"
    distance.EntityData.BundleName = "cisco_ios_xr"
    distance.EntityData.ParentYangName = "default-vrf"
    distance.EntityData.SegmentPath = "distance"
    distance.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/" + distance.EntityData.SegmentPath
    distance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    distance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    distance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    distance.EntityData.Children = types.NewOrderedMap()
    distance.EntityData.Children.Append("ospfv3", types.YChild{"Ospfv3", &distance.Ospfv3})
    distance.EntityData.Leafs = types.NewOrderedMap()
    distance.EntityData.Leafs.Append("administrative", types.YLeaf{"Administrative", distance.Administrative})

    distance.EntityData.YListKeys = []string {}

    return &(distance.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_Distance_Ospfv3
// OSPFv3 administrative distance
type Ospfv3_Processes_Process_DefaultVrf_Distance_Ospfv3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Distance for intra-area routes. The type is interface{} with range: 1..255.
    IntraArea interface{}

    // Distance for inter-area routes. The type is interface{} with range: 1..255.
    InterArea interface{}

    // Distance for external type 5 and type 7 routes. The type is interface{}
    // with range: 1..255.
    External interface{}
}

func (ospfv3 *Ospfv3_Processes_Process_DefaultVrf_Distance_Ospfv3) GetEntityData() *types.CommonEntityData {
    ospfv3.EntityData.YFilter = ospfv3.YFilter
    ospfv3.EntityData.YangName = "ospfv3"
    ospfv3.EntityData.BundleName = "cisco_ios_xr"
    ospfv3.EntityData.ParentYangName = "distance"
    ospfv3.EntityData.SegmentPath = "ospfv3"
    ospfv3.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/distance/" + ospfv3.EntityData.SegmentPath
    ospfv3.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv3.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv3.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv3.EntityData.Children = types.NewOrderedMap()
    ospfv3.EntityData.Leafs = types.NewOrderedMap()
    ospfv3.EntityData.Leafs.Append("intra-area", types.YLeaf{"IntraArea", ospfv3.IntraArea})
    ospfv3.EntityData.Leafs.Append("inter-area", types.YLeaf{"InterArea", ospfv3.InterArea})
    ospfv3.EntityData.Leafs.Append("external", types.YLeaf{"External", ospfv3.External})

    ospfv3.EntityData.YListKeys = []string {}

    return &(ospfv3.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_Maximum
// Set OSPFv3 limits
type Ospfv3_Processes_Process_DefaultVrf_Maximum struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify maximum number of interfaces. The type is interface{} with range:
    // 1..4294967295.
    Interfaces interface{}

    // Specify maximum number of paths per route. The type is interface{} with
    // range: 1..64.
    Paths interface{}

    // Limit number of redistributed prefixes.
    RedistributedPrefixes Ospfv3_Processes_Process_DefaultVrf_Maximum_RedistributedPrefixes
}

func (maximum *Ospfv3_Processes_Process_DefaultVrf_Maximum) GetEntityData() *types.CommonEntityData {
    maximum.EntityData.YFilter = maximum.YFilter
    maximum.EntityData.YangName = "maximum"
    maximum.EntityData.BundleName = "cisco_ios_xr"
    maximum.EntityData.ParentYangName = "default-vrf"
    maximum.EntityData.SegmentPath = "maximum"
    maximum.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/" + maximum.EntityData.SegmentPath
    maximum.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maximum.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maximum.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maximum.EntityData.Children = types.NewOrderedMap()
    maximum.EntityData.Children.Append("redistributed-prefixes", types.YChild{"RedistributedPrefixes", &maximum.RedistributedPrefixes})
    maximum.EntityData.Leafs = types.NewOrderedMap()
    maximum.EntityData.Leafs.Append("interfaces", types.YLeaf{"Interfaces", maximum.Interfaces})
    maximum.EntityData.Leafs.Append("paths", types.YLeaf{"Paths", maximum.Paths})

    maximum.EntityData.YListKeys = []string {}

    return &(maximum.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_Maximum_RedistributedPrefixes
// Limit number of redistributed prefixes
type Ospfv3_Processes_Process_DefaultVrf_Maximum_RedistributedPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes redistributed to protocol. The type is
    // interface{} with range: 1..4294967295.
    Prefixes interface{}

    // Threshold value (%) at which to generate a warning message. The type is
    // interface{} with range: 1..100.
    Threshold interface{}

    // Only give warning message when limit is exceeded. The type is interface{}.
    WarningOnly interface{}
}

func (redistributedPrefixes *Ospfv3_Processes_Process_DefaultVrf_Maximum_RedistributedPrefixes) GetEntityData() *types.CommonEntityData {
    redistributedPrefixes.EntityData.YFilter = redistributedPrefixes.YFilter
    redistributedPrefixes.EntityData.YangName = "redistributed-prefixes"
    redistributedPrefixes.EntityData.BundleName = "cisco_ios_xr"
    redistributedPrefixes.EntityData.ParentYangName = "maximum"
    redistributedPrefixes.EntityData.SegmentPath = "redistributed-prefixes"
    redistributedPrefixes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/maximum/" + redistributedPrefixes.EntityData.SegmentPath
    redistributedPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistributedPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistributedPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistributedPrefixes.EntityData.Children = types.NewOrderedMap()
    redistributedPrefixes.EntityData.Leafs = types.NewOrderedMap()
    redistributedPrefixes.EntityData.Leafs.Append("prefixes", types.YLeaf{"Prefixes", redistributedPrefixes.Prefixes})
    redistributedPrefixes.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", redistributedPrefixes.Threshold})
    redistributedPrefixes.EntityData.Leafs.Append("warning-only", types.YLeaf{"WarningOnly", redistributedPrefixes.WarningOnly})

    redistributedPrefixes.EntityData.YListKeys = []string {}

    return &(redistributedPrefixes.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_Redistributes
// Redistribute information from another routing
// protocol
type Ospfv3_Processes_Process_DefaultVrf_Redistributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Redistribute information from another routing protocol. The type is slice
    // of Ospfv3_Processes_Process_DefaultVrf_Redistributes_Redistribute.
    Redistribute []*Ospfv3_Processes_Process_DefaultVrf_Redistributes_Redistribute
}

func (redistributes *Ospfv3_Processes_Process_DefaultVrf_Redistributes) GetEntityData() *types.CommonEntityData {
    redistributes.EntityData.YFilter = redistributes.YFilter
    redistributes.EntityData.YangName = "redistributes"
    redistributes.EntityData.BundleName = "cisco_ios_xr"
    redistributes.EntityData.ParentYangName = "default-vrf"
    redistributes.EntityData.SegmentPath = "redistributes"
    redistributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/" + redistributes.EntityData.SegmentPath
    redistributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistributes.EntityData.Children = types.NewOrderedMap()
    redistributes.EntityData.Children.Append("redistribute", types.YChild{"Redistribute", nil})
    for i := range redistributes.Redistribute {
        redistributes.EntityData.Children.Append(types.GetSegmentPath(redistributes.Redistribute[i]), types.YChild{"Redistribute", redistributes.Redistribute[i]})
    }
    redistributes.EntityData.Leafs = types.NewOrderedMap()

    redistributes.EntityData.YListKeys = []string {}

    return &(redistributes.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_Redistributes_Redistribute
// Redistribute information from another routing
// protocol
type Ospfv3_Processes_Process_DefaultVrf_Redistributes_Redistribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Protocol. The type is Ospfv3ProtocolType2.
    ProtocolName interface{}

    // connected or static or subscriber or mobile.
    ConnectedOrStaticOrSubscriberOrMobile Ospfv3_Processes_Process_DefaultVrf_Redistributes_Redistribute_ConnectedOrStaticOrSubscriberOrMobile

    // bgp. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_Redistributes_Redistribute_Bgp.
    Bgp []*Ospfv3_Processes_Process_DefaultVrf_Redistributes_Redistribute_Bgp

    // ospfv3 or isis or application. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_Redistributes_Redistribute_Ospfv3OrIsisOrApplication.
    Ospfv3OrIsisOrApplication []*Ospfv3_Processes_Process_DefaultVrf_Redistributes_Redistribute_Ospfv3OrIsisOrApplication

    // eigrp. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_Redistributes_Redistribute_Eigrp.
    Eigrp []*Ospfv3_Processes_Process_DefaultVrf_Redistributes_Redistribute_Eigrp
}

func (redistribute *Ospfv3_Processes_Process_DefaultVrf_Redistributes_Redistribute) GetEntityData() *types.CommonEntityData {
    redistribute.EntityData.YFilter = redistribute.YFilter
    redistribute.EntityData.YangName = "redistribute"
    redistribute.EntityData.BundleName = "cisco_ios_xr"
    redistribute.EntityData.ParentYangName = "redistributes"
    redistribute.EntityData.SegmentPath = "redistribute" + types.AddKeyToken(redistribute.ProtocolName, "protocol-name")
    redistribute.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/redistributes/" + redistribute.EntityData.SegmentPath
    redistribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistribute.EntityData.Children = types.NewOrderedMap()
    redistribute.EntityData.Children.Append("connected-or-static-or-subscriber-or-mobile", types.YChild{"ConnectedOrStaticOrSubscriberOrMobile", &redistribute.ConnectedOrStaticOrSubscriberOrMobile})
    redistribute.EntityData.Children.Append("bgp", types.YChild{"Bgp", nil})
    for i := range redistribute.Bgp {
        redistribute.EntityData.Children.Append(types.GetSegmentPath(redistribute.Bgp[i]), types.YChild{"Bgp", redistribute.Bgp[i]})
    }
    redistribute.EntityData.Children.Append("ospfv3-or-isis-or-application", types.YChild{"Ospfv3OrIsisOrApplication", nil})
    for i := range redistribute.Ospfv3OrIsisOrApplication {
        redistribute.EntityData.Children.Append(types.GetSegmentPath(redistribute.Ospfv3OrIsisOrApplication[i]), types.YChild{"Ospfv3OrIsisOrApplication", redistribute.Ospfv3OrIsisOrApplication[i]})
    }
    redistribute.EntityData.Children.Append("eigrp", types.YChild{"Eigrp", nil})
    for i := range redistribute.Eigrp {
        redistribute.EntityData.Children.Append(types.GetSegmentPath(redistribute.Eigrp[i]), types.YChild{"Eigrp", redistribute.Eigrp[i]})
    }
    redistribute.EntityData.Leafs = types.NewOrderedMap()
    redistribute.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", redistribute.ProtocolName})

    redistribute.EntityData.YListKeys = []string {"ProtocolName"}

    return &(redistribute.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_Redistributes_Redistribute_ConnectedOrStaticOrSubscriberOrMobile
// connected or static or subscriber or mobile
// This type is a presence type.
type Ospfv3_Processes_Process_DefaultVrf_Redistributes_Redistribute_ConnectedOrStaticOrSubscriberOrMobile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Redistribute OSPFv3 routes. The type is Ospfv3InternalRoute.
    InternalRouteType interface{}

    // OSPFv3 default metric. The type is interface{} with range: 0..16777214.
    DefaultMetric interface{}

    // OSPFv3 exterior metric type for redistributed routes. The type is
    // Ospfv3Metric.
    MetricType interface{}

    // Tag for routes redistributed into OSPFv3. The type is interface{} with
    // range: 0..4294967295.
    Tag interface{}

    // Route policy to redistribution. The type is string.
    RoutePolicyName interface{}

    // Redistribute OSPFv3 external routes. The type is Ospfv3ExternalRoute.
    ExternalRouteType interface{}

    // Redistribute OSPFv3 NSSA external routes. The type is
    // Ospfv3nssaExternalRoute.
    NssaExternalRouteType interface{}

    // Redistribution of OSPFv3 routes. The type is bool.
    RedistributeRoute interface{}

    // ISIS route type. The type is Ospfv3isisRoute.
    IsisRouteType interface{}

    // EIGRP route type. The type is Ospfv3EigrpRoute.
    EigrpRouteType interface{}

    // Preserve (Multi-Exit Discriminator) of BGP routes. The type is bool.
    PreserveMed interface{}

    // Preserve Metric and Metric Type ofBGP Default Route. The type is bool.
    BgpPreserveDefaultInfo interface{}

    // Use metric from RIB for redistributed routes. The type is bool.
    UseRibMetric interface{}
}

func (connectedOrStaticOrSubscriberOrMobile *Ospfv3_Processes_Process_DefaultVrf_Redistributes_Redistribute_ConnectedOrStaticOrSubscriberOrMobile) GetEntityData() *types.CommonEntityData {
    connectedOrStaticOrSubscriberOrMobile.EntityData.YFilter = connectedOrStaticOrSubscriberOrMobile.YFilter
    connectedOrStaticOrSubscriberOrMobile.EntityData.YangName = "connected-or-static-or-subscriber-or-mobile"
    connectedOrStaticOrSubscriberOrMobile.EntityData.BundleName = "cisco_ios_xr"
    connectedOrStaticOrSubscriberOrMobile.EntityData.ParentYangName = "redistribute"
    connectedOrStaticOrSubscriberOrMobile.EntityData.SegmentPath = "connected-or-static-or-subscriber-or-mobile"
    connectedOrStaticOrSubscriberOrMobile.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/redistributes/redistribute/" + connectedOrStaticOrSubscriberOrMobile.EntityData.SegmentPath
    connectedOrStaticOrSubscriberOrMobile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    connectedOrStaticOrSubscriberOrMobile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    connectedOrStaticOrSubscriberOrMobile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    connectedOrStaticOrSubscriberOrMobile.EntityData.Children = types.NewOrderedMap()
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs = types.NewOrderedMap()
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs.Append("internal-route-type", types.YLeaf{"InternalRouteType", connectedOrStaticOrSubscriberOrMobile.InternalRouteType})
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs.Append("default-metric", types.YLeaf{"DefaultMetric", connectedOrStaticOrSubscriberOrMobile.DefaultMetric})
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", connectedOrStaticOrSubscriberOrMobile.MetricType})
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", connectedOrStaticOrSubscriberOrMobile.Tag})
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", connectedOrStaticOrSubscriberOrMobile.RoutePolicyName})
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs.Append("external-route-type", types.YLeaf{"ExternalRouteType", connectedOrStaticOrSubscriberOrMobile.ExternalRouteType})
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs.Append("nssa-external-route-type", types.YLeaf{"NssaExternalRouteType", connectedOrStaticOrSubscriberOrMobile.NssaExternalRouteType})
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs.Append("redistribute-route", types.YLeaf{"RedistributeRoute", connectedOrStaticOrSubscriberOrMobile.RedistributeRoute})
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs.Append("isis-route-type", types.YLeaf{"IsisRouteType", connectedOrStaticOrSubscriberOrMobile.IsisRouteType})
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs.Append("eigrp-route-type", types.YLeaf{"EigrpRouteType", connectedOrStaticOrSubscriberOrMobile.EigrpRouteType})
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs.Append("preserve-med", types.YLeaf{"PreserveMed", connectedOrStaticOrSubscriberOrMobile.PreserveMed})
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs.Append("bgp-preserve-default-info", types.YLeaf{"BgpPreserveDefaultInfo", connectedOrStaticOrSubscriberOrMobile.BgpPreserveDefaultInfo})
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs.Append("use-rib-metric", types.YLeaf{"UseRibMetric", connectedOrStaticOrSubscriberOrMobile.UseRibMetric})

    connectedOrStaticOrSubscriberOrMobile.EntityData.YListKeys = []string {}

    return &(connectedOrStaticOrSubscriberOrMobile.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_Redistributes_Redistribute_Bgp
// bgp
type Ospfv3_Processes_Process_DefaultVrf_Redistributes_Redistribute_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. 4-byte AS number in asdot (X.Y) format - first
    // half (X). The type is interface{} with range: 0..65535.
    AsXx interface{}

    // This attribute is a key. 4-byte AS number in asdot (X.Y) format - second
    // half (Y), or 2-byte AS number, or 4-byte AS number in asplain format. The
    // type is interface{} with range: 0..4294967295.
    AsYy interface{}

    // Redistribute OSPFv3 routes. The type is Ospfv3InternalRoute.
    InternalRouteType interface{}

    // OSPFv3 default metric. The type is interface{} with range: 0..16777214.
    DefaultMetric interface{}

    // OSPFv3 exterior metric type for redistributed routes. The type is
    // Ospfv3Metric.
    MetricType interface{}

    // Tag for routes redistributed into OSPFv3. The type is interface{} with
    // range: 0..4294967295.
    Tag interface{}

    // Route policy to redistribution. The type is string.
    RoutePolicyName interface{}

    // Redistribute OSPFv3 external routes. The type is Ospfv3ExternalRoute.
    ExternalRouteType interface{}

    // Redistribute OSPFv3 NSSA external routes. The type is
    // Ospfv3nssaExternalRoute.
    NssaExternalRouteType interface{}

    // Redistribution of OSPFv3 routes. The type is bool.
    RedistributeRoute interface{}

    // ISIS route type. The type is Ospfv3isisRoute.
    IsisRouteType interface{}

    // EIGRP route type. The type is Ospfv3EigrpRoute.
    EigrpRouteType interface{}

    // Preserve (Multi-Exit Discriminator) of BGP routes. The type is bool.
    PreserveMed interface{}

    // Preserve Metric and Metric Type ofBGP Default Route. The type is bool.
    BgpPreserveDefaultInfo interface{}

    // Use metric from RIB for redistributed routes. The type is bool.
    UseRibMetric interface{}
}

func (bgp *Ospfv3_Processes_Process_DefaultVrf_Redistributes_Redistribute_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "redistribute"
    bgp.EntityData.SegmentPath = "bgp" + types.AddKeyToken(bgp.AsXx, "as-xx") + types.AddKeyToken(bgp.AsYy, "as-yy")
    bgp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/redistributes/redistribute/" + bgp.EntityData.SegmentPath
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = types.NewOrderedMap()
    bgp.EntityData.Leafs = types.NewOrderedMap()
    bgp.EntityData.Leafs.Append("as-xx", types.YLeaf{"AsXx", bgp.AsXx})
    bgp.EntityData.Leafs.Append("as-yy", types.YLeaf{"AsYy", bgp.AsYy})
    bgp.EntityData.Leafs.Append("internal-route-type", types.YLeaf{"InternalRouteType", bgp.InternalRouteType})
    bgp.EntityData.Leafs.Append("default-metric", types.YLeaf{"DefaultMetric", bgp.DefaultMetric})
    bgp.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", bgp.MetricType})
    bgp.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", bgp.Tag})
    bgp.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", bgp.RoutePolicyName})
    bgp.EntityData.Leafs.Append("external-route-type", types.YLeaf{"ExternalRouteType", bgp.ExternalRouteType})
    bgp.EntityData.Leafs.Append("nssa-external-route-type", types.YLeaf{"NssaExternalRouteType", bgp.NssaExternalRouteType})
    bgp.EntityData.Leafs.Append("redistribute-route", types.YLeaf{"RedistributeRoute", bgp.RedistributeRoute})
    bgp.EntityData.Leafs.Append("isis-route-type", types.YLeaf{"IsisRouteType", bgp.IsisRouteType})
    bgp.EntityData.Leafs.Append("eigrp-route-type", types.YLeaf{"EigrpRouteType", bgp.EigrpRouteType})
    bgp.EntityData.Leafs.Append("preserve-med", types.YLeaf{"PreserveMed", bgp.PreserveMed})
    bgp.EntityData.Leafs.Append("bgp-preserve-default-info", types.YLeaf{"BgpPreserveDefaultInfo", bgp.BgpPreserveDefaultInfo})
    bgp.EntityData.Leafs.Append("use-rib-metric", types.YLeaf{"UseRibMetric", bgp.UseRibMetric})

    bgp.EntityData.YListKeys = []string {"AsXx", "AsYy"}

    return &(bgp.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_Redistributes_Redistribute_Ospfv3OrIsisOrApplication
// ospfv3 or isis or application
type Ospfv3_Processes_Process_DefaultVrf_Redistributes_Redistribute_Ospfv3OrIsisOrApplication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. ISIS process name if protocol is ISIS, or OSPFv3
    // process name if protocol is OSPFv3. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ProcessName interface{}

    // Redistribute OSPFv3 routes. The type is Ospfv3InternalRoute.
    InternalRouteType interface{}

    // OSPFv3 default metric. The type is interface{} with range: 0..16777214.
    DefaultMetric interface{}

    // OSPFv3 exterior metric type for redistributed routes. The type is
    // Ospfv3Metric.
    MetricType interface{}

    // Tag for routes redistributed into OSPFv3. The type is interface{} with
    // range: 0..4294967295.
    Tag interface{}

    // Route policy to redistribution. The type is string.
    RoutePolicyName interface{}

    // Redistribute OSPFv3 external routes. The type is Ospfv3ExternalRoute.
    ExternalRouteType interface{}

    // Redistribute OSPFv3 NSSA external routes. The type is
    // Ospfv3nssaExternalRoute.
    NssaExternalRouteType interface{}

    // Redistribution of OSPFv3 routes. The type is bool.
    RedistributeRoute interface{}

    // ISIS route type. The type is Ospfv3isisRoute.
    IsisRouteType interface{}

    // EIGRP route type. The type is Ospfv3EigrpRoute.
    EigrpRouteType interface{}

    // Preserve (Multi-Exit Discriminator) of BGP routes. The type is bool.
    PreserveMed interface{}

    // Preserve Metric and Metric Type ofBGP Default Route. The type is bool.
    BgpPreserveDefaultInfo interface{}

    // Use metric from RIB for redistributed routes. The type is bool.
    UseRibMetric interface{}
}

func (ospfv3OrIsisOrApplication *Ospfv3_Processes_Process_DefaultVrf_Redistributes_Redistribute_Ospfv3OrIsisOrApplication) GetEntityData() *types.CommonEntityData {
    ospfv3OrIsisOrApplication.EntityData.YFilter = ospfv3OrIsisOrApplication.YFilter
    ospfv3OrIsisOrApplication.EntityData.YangName = "ospfv3-or-isis-or-application"
    ospfv3OrIsisOrApplication.EntityData.BundleName = "cisco_ios_xr"
    ospfv3OrIsisOrApplication.EntityData.ParentYangName = "redistribute"
    ospfv3OrIsisOrApplication.EntityData.SegmentPath = "ospfv3-or-isis-or-application" + types.AddKeyToken(ospfv3OrIsisOrApplication.ProcessName, "process-name")
    ospfv3OrIsisOrApplication.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/redistributes/redistribute/" + ospfv3OrIsisOrApplication.EntityData.SegmentPath
    ospfv3OrIsisOrApplication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv3OrIsisOrApplication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv3OrIsisOrApplication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv3OrIsisOrApplication.EntityData.Children = types.NewOrderedMap()
    ospfv3OrIsisOrApplication.EntityData.Leafs = types.NewOrderedMap()
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", ospfv3OrIsisOrApplication.ProcessName})
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("internal-route-type", types.YLeaf{"InternalRouteType", ospfv3OrIsisOrApplication.InternalRouteType})
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("default-metric", types.YLeaf{"DefaultMetric", ospfv3OrIsisOrApplication.DefaultMetric})
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", ospfv3OrIsisOrApplication.MetricType})
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", ospfv3OrIsisOrApplication.Tag})
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", ospfv3OrIsisOrApplication.RoutePolicyName})
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("external-route-type", types.YLeaf{"ExternalRouteType", ospfv3OrIsisOrApplication.ExternalRouteType})
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("nssa-external-route-type", types.YLeaf{"NssaExternalRouteType", ospfv3OrIsisOrApplication.NssaExternalRouteType})
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("redistribute-route", types.YLeaf{"RedistributeRoute", ospfv3OrIsisOrApplication.RedistributeRoute})
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("isis-route-type", types.YLeaf{"IsisRouteType", ospfv3OrIsisOrApplication.IsisRouteType})
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("eigrp-route-type", types.YLeaf{"EigrpRouteType", ospfv3OrIsisOrApplication.EigrpRouteType})
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("preserve-med", types.YLeaf{"PreserveMed", ospfv3OrIsisOrApplication.PreserveMed})
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("bgp-preserve-default-info", types.YLeaf{"BgpPreserveDefaultInfo", ospfv3OrIsisOrApplication.BgpPreserveDefaultInfo})
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("use-rib-metric", types.YLeaf{"UseRibMetric", ospfv3OrIsisOrApplication.UseRibMetric})

    ospfv3OrIsisOrApplication.EntityData.YListKeys = []string {"ProcessName"}

    return &(ospfv3OrIsisOrApplication.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_Redistributes_Redistribute_Eigrp
// eigrp
type Ospfv3_Processes_Process_DefaultVrf_Redistributes_Redistribute_Eigrp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. 4-byte AS number in asdot (X.Y) format - first
    // half (X). The type is interface{} with range: 0..65535.
    AsXx interface{}

    // Redistribute OSPFv3 routes. The type is Ospfv3InternalRoute.
    InternalRouteType interface{}

    // OSPFv3 default metric. The type is interface{} with range: 0..16777214.
    DefaultMetric interface{}

    // OSPFv3 exterior metric type for redistributed routes. The type is
    // Ospfv3Metric.
    MetricType interface{}

    // Tag for routes redistributed into OSPFv3. The type is interface{} with
    // range: 0..4294967295.
    Tag interface{}

    // Route policy to redistribution. The type is string.
    RoutePolicyName interface{}

    // Redistribute OSPFv3 external routes. The type is Ospfv3ExternalRoute.
    ExternalRouteType interface{}

    // Redistribute OSPFv3 NSSA external routes. The type is
    // Ospfv3nssaExternalRoute.
    NssaExternalRouteType interface{}

    // Redistribution of OSPFv3 routes. The type is bool.
    RedistributeRoute interface{}

    // ISIS route type. The type is Ospfv3isisRoute.
    IsisRouteType interface{}

    // EIGRP route type. The type is Ospfv3EigrpRoute.
    EigrpRouteType interface{}

    // Preserve (Multi-Exit Discriminator) of BGP routes. The type is bool.
    PreserveMed interface{}

    // Preserve Metric and Metric Type ofBGP Default Route. The type is bool.
    BgpPreserveDefaultInfo interface{}

    // Use metric from RIB for redistributed routes. The type is bool.
    UseRibMetric interface{}
}

func (eigrp *Ospfv3_Processes_Process_DefaultVrf_Redistributes_Redistribute_Eigrp) GetEntityData() *types.CommonEntityData {
    eigrp.EntityData.YFilter = eigrp.YFilter
    eigrp.EntityData.YangName = "eigrp"
    eigrp.EntityData.BundleName = "cisco_ios_xr"
    eigrp.EntityData.ParentYangName = "redistribute"
    eigrp.EntityData.SegmentPath = "eigrp" + types.AddKeyToken(eigrp.AsXx, "as-xx")
    eigrp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/redistributes/redistribute/" + eigrp.EntityData.SegmentPath
    eigrp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eigrp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eigrp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eigrp.EntityData.Children = types.NewOrderedMap()
    eigrp.EntityData.Leafs = types.NewOrderedMap()
    eigrp.EntityData.Leafs.Append("as-xx", types.YLeaf{"AsXx", eigrp.AsXx})
    eigrp.EntityData.Leafs.Append("internal-route-type", types.YLeaf{"InternalRouteType", eigrp.InternalRouteType})
    eigrp.EntityData.Leafs.Append("default-metric", types.YLeaf{"DefaultMetric", eigrp.DefaultMetric})
    eigrp.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", eigrp.MetricType})
    eigrp.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", eigrp.Tag})
    eigrp.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", eigrp.RoutePolicyName})
    eigrp.EntityData.Leafs.Append("external-route-type", types.YLeaf{"ExternalRouteType", eigrp.ExternalRouteType})
    eigrp.EntityData.Leafs.Append("nssa-external-route-type", types.YLeaf{"NssaExternalRouteType", eigrp.NssaExternalRouteType})
    eigrp.EntityData.Leafs.Append("redistribute-route", types.YLeaf{"RedistributeRoute", eigrp.RedistributeRoute})
    eigrp.EntityData.Leafs.Append("isis-route-type", types.YLeaf{"IsisRouteType", eigrp.IsisRouteType})
    eigrp.EntityData.Leafs.Append("eigrp-route-type", types.YLeaf{"EigrpRouteType", eigrp.EigrpRouteType})
    eigrp.EntityData.Leafs.Append("preserve-med", types.YLeaf{"PreserveMed", eigrp.PreserveMed})
    eigrp.EntityData.Leafs.Append("bgp-preserve-default-info", types.YLeaf{"BgpPreserveDefaultInfo", eigrp.BgpPreserveDefaultInfo})
    eigrp.EntityData.Leafs.Append("use-rib-metric", types.YLeaf{"UseRibMetric", eigrp.UseRibMetric})

    eigrp.EntityData.YListKeys = []string {"AsXx"}

    return &(eigrp.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_Ignore
// Do not complain about a specified event
type Ospfv3_Processes_Process_DefaultVrf_Ignore struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Do not complain upon receiving LSA of the specified type.
    Lsa Ospfv3_Processes_Process_DefaultVrf_Ignore_Lsa
}

func (ignore *Ospfv3_Processes_Process_DefaultVrf_Ignore) GetEntityData() *types.CommonEntityData {
    ignore.EntityData.YFilter = ignore.YFilter
    ignore.EntityData.YangName = "ignore"
    ignore.EntityData.BundleName = "cisco_ios_xr"
    ignore.EntityData.ParentYangName = "default-vrf"
    ignore.EntityData.SegmentPath = "ignore"
    ignore.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/" + ignore.EntityData.SegmentPath
    ignore.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ignore.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ignore.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ignore.EntityData.Children = types.NewOrderedMap()
    ignore.EntityData.Children.Append("lsa", types.YChild{"Lsa", &ignore.Lsa})
    ignore.EntityData.Leafs = types.NewOrderedMap()

    ignore.EntityData.YListKeys = []string {}

    return &(ignore.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_Ignore_Lsa
// Do not complain upon receiving LSA of the
// specified type
type Ospfv3_Processes_Process_DefaultVrf_Ignore_Lsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable ignore of MOSPF type 6 LSA. The type is interface{}.
    Mospf interface{}
}

func (lsa *Ospfv3_Processes_Process_DefaultVrf_Ignore_Lsa) GetEntityData() *types.CommonEntityData {
    lsa.EntityData.YFilter = lsa.YFilter
    lsa.EntityData.YangName = "lsa"
    lsa.EntityData.BundleName = "cisco_ios_xr"
    lsa.EntityData.ParentYangName = "ignore"
    lsa.EntityData.SegmentPath = "lsa"
    lsa.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/ignore/" + lsa.EntityData.SegmentPath
    lsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsa.EntityData.Children = types.NewOrderedMap()
    lsa.EntityData.Leafs = types.NewOrderedMap()
    lsa.EntityData.Leafs.Append("mospf", types.YLeaf{"Mospf", lsa.Mospf})

    lsa.EntityData.YListKeys = []string {}

    return &(lsa.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_DistributeListOut
// Filter prefixes from RIB 
type Ospfv3_Processes_Process_DefaultVrf_DistributeListOut struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter generated type-5 LSAs.
    DistributeOuts Ospfv3_Processes_Process_DefaultVrf_DistributeListOut_DistributeOuts
}

func (distributeListOut *Ospfv3_Processes_Process_DefaultVrf_DistributeListOut) GetEntityData() *types.CommonEntityData {
    distributeListOut.EntityData.YFilter = distributeListOut.YFilter
    distributeListOut.EntityData.YangName = "distribute-list-out"
    distributeListOut.EntityData.BundleName = "cisco_ios_xr"
    distributeListOut.EntityData.ParentYangName = "default-vrf"
    distributeListOut.EntityData.SegmentPath = "distribute-list-out"
    distributeListOut.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/" + distributeListOut.EntityData.SegmentPath
    distributeListOut.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    distributeListOut.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    distributeListOut.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    distributeListOut.EntityData.Children = types.NewOrderedMap()
    distributeListOut.EntityData.Children.Append("distribute-outs", types.YChild{"DistributeOuts", &distributeListOut.DistributeOuts})
    distributeListOut.EntityData.Leafs = types.NewOrderedMap()

    distributeListOut.EntityData.YListKeys = []string {}

    return &(distributeListOut.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_DistributeListOut_DistributeOuts
// Filter generated type-5 LSAs
type Ospfv3_Processes_Process_DefaultVrf_DistributeListOut_DistributeOuts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter generated type-5 LSAs. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_DistributeListOut_DistributeOuts_DistributeOut.
    DistributeOut []*Ospfv3_Processes_Process_DefaultVrf_DistributeListOut_DistributeOuts_DistributeOut
}

func (distributeOuts *Ospfv3_Processes_Process_DefaultVrf_DistributeListOut_DistributeOuts) GetEntityData() *types.CommonEntityData {
    distributeOuts.EntityData.YFilter = distributeOuts.YFilter
    distributeOuts.EntityData.YangName = "distribute-outs"
    distributeOuts.EntityData.BundleName = "cisco_ios_xr"
    distributeOuts.EntityData.ParentYangName = "distribute-list-out"
    distributeOuts.EntityData.SegmentPath = "distribute-outs"
    distributeOuts.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/distribute-list-out/" + distributeOuts.EntityData.SegmentPath
    distributeOuts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    distributeOuts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    distributeOuts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    distributeOuts.EntityData.Children = types.NewOrderedMap()
    distributeOuts.EntityData.Children.Append("distribute-out", types.YChild{"DistributeOut", nil})
    for i := range distributeOuts.DistributeOut {
        distributeOuts.EntityData.Children.Append(types.GetSegmentPath(distributeOuts.DistributeOut[i]), types.YChild{"DistributeOut", distributeOuts.DistributeOut[i]})
    }
    distributeOuts.EntityData.Leafs = types.NewOrderedMap()

    distributeOuts.EntityData.YListKeys = []string {}

    return &(distributeOuts.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_DistributeListOut_DistributeOuts_DistributeOut
// Filter generated type-5 LSAs
type Ospfv3_Processes_Process_DefaultVrf_DistributeListOut_DistributeOuts_DistributeOut struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. none. The type is Ospfv3Protocol.
    ProtocolName interface{}

    // Prefix-list name. The type is string.
    AllOrConnectedOrStaticPrefixList interface{}

    // bgp. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_DistributeListOut_DistributeOuts_DistributeOut_Bgp.
    Bgp []*Ospfv3_Processes_Process_DefaultVrf_DistributeListOut_DistributeOuts_DistributeOut_Bgp

    // ospfv3 or isis. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_DistributeListOut_DistributeOuts_DistributeOut_Ospfv3OrIsis.
    Ospfv3OrIsis []*Ospfv3_Processes_Process_DefaultVrf_DistributeListOut_DistributeOuts_DistributeOut_Ospfv3OrIsis

    // eigrp. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_DistributeListOut_DistributeOuts_DistributeOut_Eigrp.
    Eigrp []*Ospfv3_Processes_Process_DefaultVrf_DistributeListOut_DistributeOuts_DistributeOut_Eigrp
}

func (distributeOut *Ospfv3_Processes_Process_DefaultVrf_DistributeListOut_DistributeOuts_DistributeOut) GetEntityData() *types.CommonEntityData {
    distributeOut.EntityData.YFilter = distributeOut.YFilter
    distributeOut.EntityData.YangName = "distribute-out"
    distributeOut.EntityData.BundleName = "cisco_ios_xr"
    distributeOut.EntityData.ParentYangName = "distribute-outs"
    distributeOut.EntityData.SegmentPath = "distribute-out" + types.AddKeyToken(distributeOut.ProtocolName, "protocol-name")
    distributeOut.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/distribute-list-out/distribute-outs/" + distributeOut.EntityData.SegmentPath
    distributeOut.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    distributeOut.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    distributeOut.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    distributeOut.EntityData.Children = types.NewOrderedMap()
    distributeOut.EntityData.Children.Append("bgp", types.YChild{"Bgp", nil})
    for i := range distributeOut.Bgp {
        distributeOut.EntityData.Children.Append(types.GetSegmentPath(distributeOut.Bgp[i]), types.YChild{"Bgp", distributeOut.Bgp[i]})
    }
    distributeOut.EntityData.Children.Append("ospfv3-or-isis", types.YChild{"Ospfv3OrIsis", nil})
    for i := range distributeOut.Ospfv3OrIsis {
        distributeOut.EntityData.Children.Append(types.GetSegmentPath(distributeOut.Ospfv3OrIsis[i]), types.YChild{"Ospfv3OrIsis", distributeOut.Ospfv3OrIsis[i]})
    }
    distributeOut.EntityData.Children.Append("eigrp", types.YChild{"Eigrp", nil})
    for i := range distributeOut.Eigrp {
        distributeOut.EntityData.Children.Append(types.GetSegmentPath(distributeOut.Eigrp[i]), types.YChild{"Eigrp", distributeOut.Eigrp[i]})
    }
    distributeOut.EntityData.Leafs = types.NewOrderedMap()
    distributeOut.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", distributeOut.ProtocolName})
    distributeOut.EntityData.Leafs.Append("all-or-connected-or-static-prefix-list", types.YLeaf{"AllOrConnectedOrStaticPrefixList", distributeOut.AllOrConnectedOrStaticPrefixList})

    distributeOut.EntityData.YListKeys = []string {"ProtocolName"}

    return &(distributeOut.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_DistributeListOut_DistributeOuts_DistributeOut_Bgp
// bgp
type Ospfv3_Processes_Process_DefaultVrf_DistributeListOut_DistributeOuts_DistributeOut_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. 4-byte AS number in asdot (X.Y) format - first
    // half (X). The type is interface{} with range: 1..65535.
    AsXx interface{}

    // This attribute is a key. 4-byte AS number in asdot (X.Y) format - second
    // half (Y), or 2-byte AS number, or 4-byte AS number in asplain format. The
    // type is interface{} with range: 0..4294967295.
    AsYy interface{}

    // Prefix-list name. The type is string.
    PrefixList interface{}
}

func (bgp *Ospfv3_Processes_Process_DefaultVrf_DistributeListOut_DistributeOuts_DistributeOut_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "distribute-out"
    bgp.EntityData.SegmentPath = "bgp" + types.AddKeyToken(bgp.AsXx, "as-xx") + types.AddKeyToken(bgp.AsYy, "as-yy")
    bgp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/distribute-list-out/distribute-outs/distribute-out/" + bgp.EntityData.SegmentPath
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = types.NewOrderedMap()
    bgp.EntityData.Leafs = types.NewOrderedMap()
    bgp.EntityData.Leafs.Append("as-xx", types.YLeaf{"AsXx", bgp.AsXx})
    bgp.EntityData.Leafs.Append("as-yy", types.YLeaf{"AsYy", bgp.AsYy})
    bgp.EntityData.Leafs.Append("prefix-list", types.YLeaf{"PrefixList", bgp.PrefixList})

    bgp.EntityData.YListKeys = []string {"AsXx", "AsYy"}

    return &(bgp.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_DistributeListOut_DistributeOuts_DistributeOut_Ospfv3OrIsis
// ospfv3 or isis
type Ospfv3_Processes_Process_DefaultVrf_DistributeListOut_DistributeOuts_DistributeOut_Ospfv3OrIsis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. If ISIS or OSPFv3, specify the instance name. The
    // type is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ProcessName interface{}

    // Prefix-list name. The type is string.
    PrefixList interface{}
}

func (ospfv3OrIsis *Ospfv3_Processes_Process_DefaultVrf_DistributeListOut_DistributeOuts_DistributeOut_Ospfv3OrIsis) GetEntityData() *types.CommonEntityData {
    ospfv3OrIsis.EntityData.YFilter = ospfv3OrIsis.YFilter
    ospfv3OrIsis.EntityData.YangName = "ospfv3-or-isis"
    ospfv3OrIsis.EntityData.BundleName = "cisco_ios_xr"
    ospfv3OrIsis.EntityData.ParentYangName = "distribute-out"
    ospfv3OrIsis.EntityData.SegmentPath = "ospfv3-or-isis" + types.AddKeyToken(ospfv3OrIsis.ProcessName, "process-name")
    ospfv3OrIsis.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/distribute-list-out/distribute-outs/distribute-out/" + ospfv3OrIsis.EntityData.SegmentPath
    ospfv3OrIsis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv3OrIsis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv3OrIsis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv3OrIsis.EntityData.Children = types.NewOrderedMap()
    ospfv3OrIsis.EntityData.Leafs = types.NewOrderedMap()
    ospfv3OrIsis.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", ospfv3OrIsis.ProcessName})
    ospfv3OrIsis.EntityData.Leafs.Append("prefix-list", types.YLeaf{"PrefixList", ospfv3OrIsis.PrefixList})

    ospfv3OrIsis.EntityData.YListKeys = []string {"ProcessName"}

    return &(ospfv3OrIsis.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_DistributeListOut_DistributeOuts_DistributeOut_Eigrp
// eigrp
type Ospfv3_Processes_Process_DefaultVrf_DistributeListOut_DistributeOuts_DistributeOut_Eigrp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. 4-byte AS number in asdot (X.Y) format - first
    // half (X). The type is interface{} with range: 1..65535.
    AsXx interface{}

    // Prefix-list name. The type is string.
    PrefixList interface{}
}

func (eigrp *Ospfv3_Processes_Process_DefaultVrf_DistributeListOut_DistributeOuts_DistributeOut_Eigrp) GetEntityData() *types.CommonEntityData {
    eigrp.EntityData.YFilter = eigrp.YFilter
    eigrp.EntityData.YangName = "eigrp"
    eigrp.EntityData.BundleName = "cisco_ios_xr"
    eigrp.EntityData.ParentYangName = "distribute-out"
    eigrp.EntityData.SegmentPath = "eigrp" + types.AddKeyToken(eigrp.AsXx, "as-xx")
    eigrp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/distribute-list-out/distribute-outs/distribute-out/" + eigrp.EntityData.SegmentPath
    eigrp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eigrp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eigrp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eigrp.EntityData.Children = types.NewOrderedMap()
    eigrp.EntityData.Leafs = types.NewOrderedMap()
    eigrp.EntityData.Leafs.Append("as-xx", types.YLeaf{"AsXx", eigrp.AsXx})
    eigrp.EntityData.Leafs.Append("prefix-list", types.YLeaf{"PrefixList", eigrp.PrefixList})

    eigrp.EntityData.YListKeys = []string {"AsXx"}

    return &(eigrp.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_DistributeList
// Filter prefixes to/from RIB
type Ospfv3_Processes_Process_DefaultVrf_DistributeList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter prefixes installed to RIB.
    In Ospfv3_Processes_Process_DefaultVrf_DistributeList_In
}

func (distributeList *Ospfv3_Processes_Process_DefaultVrf_DistributeList) GetEntityData() *types.CommonEntityData {
    distributeList.EntityData.YFilter = distributeList.YFilter
    distributeList.EntityData.YangName = "distribute-list"
    distributeList.EntityData.BundleName = "cisco_ios_xr"
    distributeList.EntityData.ParentYangName = "default-vrf"
    distributeList.EntityData.SegmentPath = "distribute-list"
    distributeList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/" + distributeList.EntityData.SegmentPath
    distributeList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    distributeList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    distributeList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    distributeList.EntityData.Children = types.NewOrderedMap()
    distributeList.EntityData.Children.Append("in", types.YChild{"In", &distributeList.In})
    distributeList.EntityData.Leafs = types.NewOrderedMap()

    distributeList.EntityData.YListKeys = []string {}

    return &(distributeList.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_DistributeList_In
// Filter prefixes installed to RIB
type Ospfv3_Processes_Process_DefaultVrf_DistributeList_In struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter prefixes based on an IPv6 prefix-list. The type is string.
    PrefixList interface{}
}

func (in *Ospfv3_Processes_Process_DefaultVrf_DistributeList_In) GetEntityData() *types.CommonEntityData {
    in.EntityData.YFilter = in.YFilter
    in.EntityData.YangName = "in"
    in.EntityData.BundleName = "cisco_ios_xr"
    in.EntityData.ParentYangName = "distribute-list"
    in.EntityData.SegmentPath = "in"
    in.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/distribute-list/" + in.EntityData.SegmentPath
    in.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    in.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    in.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    in.EntityData.Children = types.NewOrderedMap()
    in.EntityData.Leafs = types.NewOrderedMap()
    in.EntityData.Leafs.Append("prefix-list", types.YLeaf{"PrefixList", in.PrefixList})

    in.EntityData.YListKeys = []string {}

    return &(in.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_StubRouter
// Stub router configuration
type Ospfv3_Processes_Process_DefaultVrf_StubRouter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Stub router R-bit configuration.
    Rbit Ospfv3_Processes_Process_DefaultVrf_StubRouter_Rbit

    // Stub router V6-bit configuration.
    V6bit Ospfv3_Processes_Process_DefaultVrf_StubRouter_V6bit

    // Stub router max-metric configuration.
    MaxMetric Ospfv3_Processes_Process_DefaultVrf_StubRouter_MaxMetric
}

func (stubRouter *Ospfv3_Processes_Process_DefaultVrf_StubRouter) GetEntityData() *types.CommonEntityData {
    stubRouter.EntityData.YFilter = stubRouter.YFilter
    stubRouter.EntityData.YangName = "stub-router"
    stubRouter.EntityData.BundleName = "cisco_ios_xr"
    stubRouter.EntityData.ParentYangName = "default-vrf"
    stubRouter.EntityData.SegmentPath = "stub-router"
    stubRouter.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/" + stubRouter.EntityData.SegmentPath
    stubRouter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stubRouter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stubRouter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stubRouter.EntityData.Children = types.NewOrderedMap()
    stubRouter.EntityData.Children.Append("rbit", types.YChild{"Rbit", &stubRouter.Rbit})
    stubRouter.EntityData.Children.Append("v6bit", types.YChild{"V6bit", &stubRouter.V6bit})
    stubRouter.EntityData.Children.Append("max-metric", types.YChild{"MaxMetric", &stubRouter.MaxMetric})
    stubRouter.EntityData.Leafs = types.NewOrderedMap()

    stubRouter.EntityData.YListKeys = []string {}

    return &(stubRouter.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_StubRouter_Rbit
// Stub router R-bit configuration
// This type is a presence type.
type Ospfv3_Processes_Process_DefaultVrf_StubRouter_Rbit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Time (in seconds) to stay in stub router operational state. The type is
    // interface{} with range: 5..86400. Units are second.
    OnSwitchover interface{}

    // Unconditionally enter stub router operational state. The type is
    // interface{}.
    Always interface{}

    // Advertise stub links with maximum metric in stub router mode. The type is
    // interface{}.
    IncludeStub interface{}

    // Time (in seconds) to stay in stub router operational state. The type is
    // interface{} with range: 5..86400. Units are second.
    OnProcMigration interface{}

    // Time (in seconds) to stay in stub router operational state. The type is
    // interface{} with range: 5..86400. Units are second.
    OnProcRestart interface{}

    // Enter stub router operational state on startup.
    OnStartup Ospfv3_Processes_Process_DefaultVrf_StubRouter_Rbit_OnStartup
}

func (rbit *Ospfv3_Processes_Process_DefaultVrf_StubRouter_Rbit) GetEntityData() *types.CommonEntityData {
    rbit.EntityData.YFilter = rbit.YFilter
    rbit.EntityData.YangName = "rbit"
    rbit.EntityData.BundleName = "cisco_ios_xr"
    rbit.EntityData.ParentYangName = "stub-router"
    rbit.EntityData.SegmentPath = "rbit"
    rbit.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/stub-router/" + rbit.EntityData.SegmentPath
    rbit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rbit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rbit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rbit.EntityData.Children = types.NewOrderedMap()
    rbit.EntityData.Children.Append("on-startup", types.YChild{"OnStartup", &rbit.OnStartup})
    rbit.EntityData.Leafs = types.NewOrderedMap()
    rbit.EntityData.Leafs.Append("on-switchover", types.YLeaf{"OnSwitchover", rbit.OnSwitchover})
    rbit.EntityData.Leafs.Append("always", types.YLeaf{"Always", rbit.Always})
    rbit.EntityData.Leafs.Append("include-stub", types.YLeaf{"IncludeStub", rbit.IncludeStub})
    rbit.EntityData.Leafs.Append("on-proc-migration", types.YLeaf{"OnProcMigration", rbit.OnProcMigration})
    rbit.EntityData.Leafs.Append("on-proc-restart", types.YLeaf{"OnProcRestart", rbit.OnProcRestart})

    rbit.EntityData.YListKeys = []string {}

    return &(rbit.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_StubRouter_Rbit_OnStartup
// Enter stub router operational state on startup
type Ospfv3_Processes_Process_DefaultVrf_StubRouter_Rbit_OnStartup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Wait until BGP converges (only applicable to default VRF). The type is
    // bool. The default value is false.
    WaitForBgp interface{}

    // Time (in seconds) to stay in stub router operational state. The type is
    // interface{} with range: 5..86400. Units are second.
    WaitTime interface{}
}

func (onStartup *Ospfv3_Processes_Process_DefaultVrf_StubRouter_Rbit_OnStartup) GetEntityData() *types.CommonEntityData {
    onStartup.EntityData.YFilter = onStartup.YFilter
    onStartup.EntityData.YangName = "on-startup"
    onStartup.EntityData.BundleName = "cisco_ios_xr"
    onStartup.EntityData.ParentYangName = "rbit"
    onStartup.EntityData.SegmentPath = "on-startup"
    onStartup.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/stub-router/rbit/" + onStartup.EntityData.SegmentPath
    onStartup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    onStartup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    onStartup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    onStartup.EntityData.Children = types.NewOrderedMap()
    onStartup.EntityData.Leafs = types.NewOrderedMap()
    onStartup.EntityData.Leafs.Append("wait-for-bgp", types.YLeaf{"WaitForBgp", onStartup.WaitForBgp})
    onStartup.EntityData.Leafs.Append("wait-time", types.YLeaf{"WaitTime", onStartup.WaitTime})

    onStartup.EntityData.YListKeys = []string {}

    return &(onStartup.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_StubRouter_V6bit
// Stub router V6-bit configuration
// This type is a presence type.
type Ospfv3_Processes_Process_DefaultVrf_StubRouter_V6bit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Time (in seconds) to stay in stub router operational state. The type is
    // interface{} with range: 5..86400. Units are second.
    OnSwitchover interface{}

    // Unconditionally enter stub router operational state. The type is
    // interface{}.
    Always interface{}

    // Time (in seconds) to stay in stub router operational state. The type is
    // interface{} with range: 5..86400. Units are second.
    OnProcMigration interface{}

    // Time (in seconds) to stay in stub router operational state. The type is
    // interface{} with range: 5..86400. Units are second.
    OnProcRestart interface{}

    // Enter stub router operational state on startup.
    OnStartup Ospfv3_Processes_Process_DefaultVrf_StubRouter_V6bit_OnStartup
}

func (v6bit *Ospfv3_Processes_Process_DefaultVrf_StubRouter_V6bit) GetEntityData() *types.CommonEntityData {
    v6bit.EntityData.YFilter = v6bit.YFilter
    v6bit.EntityData.YangName = "v6bit"
    v6bit.EntityData.BundleName = "cisco_ios_xr"
    v6bit.EntityData.ParentYangName = "stub-router"
    v6bit.EntityData.SegmentPath = "v6bit"
    v6bit.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/stub-router/" + v6bit.EntityData.SegmentPath
    v6bit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v6bit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v6bit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v6bit.EntityData.Children = types.NewOrderedMap()
    v6bit.EntityData.Children.Append("on-startup", types.YChild{"OnStartup", &v6bit.OnStartup})
    v6bit.EntityData.Leafs = types.NewOrderedMap()
    v6bit.EntityData.Leafs.Append("on-switchover", types.YLeaf{"OnSwitchover", v6bit.OnSwitchover})
    v6bit.EntityData.Leafs.Append("always", types.YLeaf{"Always", v6bit.Always})
    v6bit.EntityData.Leafs.Append("on-proc-migration", types.YLeaf{"OnProcMigration", v6bit.OnProcMigration})
    v6bit.EntityData.Leafs.Append("on-proc-restart", types.YLeaf{"OnProcRestart", v6bit.OnProcRestart})

    v6bit.EntityData.YListKeys = []string {}

    return &(v6bit.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_StubRouter_V6bit_OnStartup
// Enter stub router operational state on startup
type Ospfv3_Processes_Process_DefaultVrf_StubRouter_V6bit_OnStartup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Wait until BGP converges (only applicable to default VRF). The type is
    // bool. The default value is false.
    WaitForBgp interface{}

    // Time (in seconds) to stay in stub router operational state. The type is
    // interface{} with range: 5..86400. Units are second.
    WaitTime interface{}
}

func (onStartup *Ospfv3_Processes_Process_DefaultVrf_StubRouter_V6bit_OnStartup) GetEntityData() *types.CommonEntityData {
    onStartup.EntityData.YFilter = onStartup.YFilter
    onStartup.EntityData.YangName = "on-startup"
    onStartup.EntityData.BundleName = "cisco_ios_xr"
    onStartup.EntityData.ParentYangName = "v6bit"
    onStartup.EntityData.SegmentPath = "on-startup"
    onStartup.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/stub-router/v6bit/" + onStartup.EntityData.SegmentPath
    onStartup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    onStartup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    onStartup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    onStartup.EntityData.Children = types.NewOrderedMap()
    onStartup.EntityData.Leafs = types.NewOrderedMap()
    onStartup.EntityData.Leafs.Append("wait-for-bgp", types.YLeaf{"WaitForBgp", onStartup.WaitForBgp})
    onStartup.EntityData.Leafs.Append("wait-time", types.YLeaf{"WaitTime", onStartup.WaitTime})

    onStartup.EntityData.YListKeys = []string {}

    return &(onStartup.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_StubRouter_MaxMetric
// Stub router max-metric configuration
// This type is a presence type.
type Ospfv3_Processes_Process_DefaultVrf_StubRouter_MaxMetric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Advertise external LSAs with modified metric in stub router mode. The type
    // is interface{} with range: 1..16777214. The default value is 16711680.
    ExternalLsa interface{}

    // Advertise summary LSAs with modified metric in stub router mode. The type
    // is interface{} with range: 1..16777214. The default value is 16711680.
    SummaryLsa interface{}

    // Time (in seconds) to stay in stub router operational state. The type is
    // interface{} with range: 5..86400. Units are second.
    OnSwitchover interface{}

    // Unconditionally enter stub router operational state. The type is
    // interface{}.
    Always interface{}

    // Advertise stub links with maximum metric in stub router mode. The type is
    // interface{}.
    IncludeStub interface{}

    // Time (in seconds) to stay in stub router operational state. The type is
    // interface{} with range: 5..86400. Units are second.
    OnProcMigration interface{}

    // Time (in seconds) to stay in stub router operational state. The type is
    // interface{} with range: 5..86400. Units are second.
    OnProcRestart interface{}

    // Enter stub router operational state on startup.
    OnStartup Ospfv3_Processes_Process_DefaultVrf_StubRouter_MaxMetric_OnStartup
}

func (maxMetric *Ospfv3_Processes_Process_DefaultVrf_StubRouter_MaxMetric) GetEntityData() *types.CommonEntityData {
    maxMetric.EntityData.YFilter = maxMetric.YFilter
    maxMetric.EntityData.YangName = "max-metric"
    maxMetric.EntityData.BundleName = "cisco_ios_xr"
    maxMetric.EntityData.ParentYangName = "stub-router"
    maxMetric.EntityData.SegmentPath = "max-metric"
    maxMetric.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/stub-router/" + maxMetric.EntityData.SegmentPath
    maxMetric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maxMetric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maxMetric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maxMetric.EntityData.Children = types.NewOrderedMap()
    maxMetric.EntityData.Children.Append("on-startup", types.YChild{"OnStartup", &maxMetric.OnStartup})
    maxMetric.EntityData.Leafs = types.NewOrderedMap()
    maxMetric.EntityData.Leafs.Append("external-lsa", types.YLeaf{"ExternalLsa", maxMetric.ExternalLsa})
    maxMetric.EntityData.Leafs.Append("summary-lsa", types.YLeaf{"SummaryLsa", maxMetric.SummaryLsa})
    maxMetric.EntityData.Leafs.Append("on-switchover", types.YLeaf{"OnSwitchover", maxMetric.OnSwitchover})
    maxMetric.EntityData.Leafs.Append("always", types.YLeaf{"Always", maxMetric.Always})
    maxMetric.EntityData.Leafs.Append("include-stub", types.YLeaf{"IncludeStub", maxMetric.IncludeStub})
    maxMetric.EntityData.Leafs.Append("on-proc-migration", types.YLeaf{"OnProcMigration", maxMetric.OnProcMigration})
    maxMetric.EntityData.Leafs.Append("on-proc-restart", types.YLeaf{"OnProcRestart", maxMetric.OnProcRestart})

    maxMetric.EntityData.YListKeys = []string {}

    return &(maxMetric.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_StubRouter_MaxMetric_OnStartup
// Enter stub router operational state on startup
type Ospfv3_Processes_Process_DefaultVrf_StubRouter_MaxMetric_OnStartup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Wait until BGP converges (only applicable to default VRF). The type is
    // bool. The default value is false.
    WaitForBgp interface{}

    // Time (in seconds) to stay in stub router operational state. The type is
    // interface{} with range: 5..86400. Units are second.
    WaitTime interface{}
}

func (onStartup *Ospfv3_Processes_Process_DefaultVrf_StubRouter_MaxMetric_OnStartup) GetEntityData() *types.CommonEntityData {
    onStartup.EntityData.YFilter = onStartup.YFilter
    onStartup.EntityData.YangName = "on-startup"
    onStartup.EntityData.BundleName = "cisco_ios_xr"
    onStartup.EntityData.ParentYangName = "max-metric"
    onStartup.EntityData.SegmentPath = "on-startup"
    onStartup.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/stub-router/max-metric/" + onStartup.EntityData.SegmentPath
    onStartup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    onStartup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    onStartup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    onStartup.EntityData.Children = types.NewOrderedMap()
    onStartup.EntityData.Leafs = types.NewOrderedMap()
    onStartup.EntityData.Leafs.Append("wait-for-bgp", types.YLeaf{"WaitForBgp", onStartup.WaitForBgp})
    onStartup.EntityData.Leafs.Append("wait-time", types.YLeaf{"WaitTime", onStartup.WaitTime})

    onStartup.EntityData.YListKeys = []string {}

    return &(onStartup.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_Bfd
// Configure BFD parameters
type Ospfv3_Processes_Process_DefaultVrf_Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Hello interval in milli-seconds. The type is interface{} with range:
    // 3..30000. Units are millisecond.
    Interval interface{}

    // Detect multiplier. The type is interface{} with range: 2..50.
    DetectionMultiplier interface{}

    // Enable or disable BFD fast detection. The type is Ospfv3bfdEnableMode.
    FastDetectMode interface{}
}

func (bfd *Ospfv3_Processes_Process_DefaultVrf_Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "default-vrf"
    bfd.EntityData.SegmentPath = "bfd"
    bfd.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/" + bfd.EntityData.SegmentPath
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = types.NewOrderedMap()
    bfd.EntityData.Leafs = types.NewOrderedMap()
    bfd.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", bfd.Interval})
    bfd.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", bfd.DetectionMultiplier})
    bfd.EntityData.Leafs.Append("fast-detect-mode", types.YLeaf{"FastDetectMode", bfd.FastDetectMode})

    bfd.EntityData.YListKeys = []string {}

    return &(bfd.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_DatabaseFilter
// Database filter
type Ospfv3_Processes_Process_DefaultVrf_DatabaseFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All.
    All Ospfv3_Processes_Process_DefaultVrf_DatabaseFilter_All
}

func (databaseFilter *Ospfv3_Processes_Process_DefaultVrf_DatabaseFilter) GetEntityData() *types.CommonEntityData {
    databaseFilter.EntityData.YFilter = databaseFilter.YFilter
    databaseFilter.EntityData.YangName = "database-filter"
    databaseFilter.EntityData.BundleName = "cisco_ios_xr"
    databaseFilter.EntityData.ParentYangName = "default-vrf"
    databaseFilter.EntityData.SegmentPath = "database-filter"
    databaseFilter.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/" + databaseFilter.EntityData.SegmentPath
    databaseFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    databaseFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    databaseFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    databaseFilter.EntityData.Children = types.NewOrderedMap()
    databaseFilter.EntityData.Children.Append("all", types.YChild{"All", &databaseFilter.All})
    databaseFilter.EntityData.Leafs = types.NewOrderedMap()

    databaseFilter.EntityData.YListKeys = []string {}

    return &(databaseFilter.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_DatabaseFilter_All
// All
type Ospfv3_Processes_Process_DefaultVrf_DatabaseFilter_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable out. The type is interface{}.
    Out interface{}
}

func (all *Ospfv3_Processes_Process_DefaultVrf_DatabaseFilter_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "database-filter"
    all.EntityData.SegmentPath = "all"
    all.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/database-filter/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Leafs = types.NewOrderedMap()
    all.EntityData.Leafs.Append("out", types.YLeaf{"Out", all.Out})

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_Capability
// OSPFv3 Capability
type Ospfv3_Processes_Process_DefaultVrf_Capability struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NSSA capability to prefer Type 7 over Type 5. The type is bool.
    Type7Prefer interface{}

    // Enable VRF Lite. The type is bool.
    VrfLite interface{}

    // Enable capability to translate LSAs with fwd addr. The type is bool.
    Type7TranslateZeroForwardingAddr interface{}
}

func (capability *Ospfv3_Processes_Process_DefaultVrf_Capability) GetEntityData() *types.CommonEntityData {
    capability.EntityData.YFilter = capability.YFilter
    capability.EntityData.YangName = "capability"
    capability.EntityData.BundleName = "cisco_ios_xr"
    capability.EntityData.ParentYangName = "default-vrf"
    capability.EntityData.SegmentPath = "capability"
    capability.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/" + capability.EntityData.SegmentPath
    capability.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    capability.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    capability.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    capability.EntityData.Children = types.NewOrderedMap()
    capability.EntityData.Leafs = types.NewOrderedMap()
    capability.EntityData.Leafs.Append("type7-prefer", types.YLeaf{"Type7Prefer", capability.Type7Prefer})
    capability.EntityData.Leafs.Append("vrf-lite", types.YLeaf{"VrfLite", capability.VrfLite})
    capability.EntityData.Leafs.Append("type7-translate-zero-forwarding-addr", types.YLeaf{"Type7TranslateZeroForwardingAddr", capability.Type7TranslateZeroForwardingAddr})

    capability.EntityData.YListKeys = []string {}

    return &(capability.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_Authentication
// Authenticate OSPFv3 packets
type Ospfv3_Processes_Process_DefaultVrf_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec AH authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Use the MD5 or SHA1 algorithm. The type is Ospfv3Authentication.
    Algorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    Password interface{}
}

func (authentication *Ospfv3_Processes_Process_DefaultVrf_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "default-vrf"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", authentication.Enable})
    authentication.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", authentication.Spi})
    authentication.EntityData.Leafs.Append("algorithm", types.YLeaf{"Algorithm", authentication.Algorithm})
    authentication.EntityData.Leafs.Append("password", types.YLeaf{"Password", authentication.Password})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_GracefulRestart
// Graceful restart configuration
type Ospfv3_Processes_Process_DefaultVrf_GracefulRestart struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum interval between graceful restarts (seconds). The type is
    // interface{} with range: 90..3600. Units are second.
    Interval interface{}

    // Terminate graceful restart helper mode if LSA changed. The type is
    // interface{}.
    StrictLsaChecking interface{}

    // Disable router's helper support. The type is interface{}.
    Helper interface{}

    // Enable graceful restart. The type is interface{}.
    Enable interface{}

    // Maximum route lifetime following restart (seconds). The type is interface{}
    // with range: 90..1800. Units are second.
    Lifetime interface{}
}

func (gracefulRestart *Ospfv3_Processes_Process_DefaultVrf_GracefulRestart) GetEntityData() *types.CommonEntityData {
    gracefulRestart.EntityData.YFilter = gracefulRestart.YFilter
    gracefulRestart.EntityData.YangName = "graceful-restart"
    gracefulRestart.EntityData.BundleName = "cisco_ios_xr"
    gracefulRestart.EntityData.ParentYangName = "default-vrf"
    gracefulRestart.EntityData.SegmentPath = "graceful-restart"
    gracefulRestart.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/" + gracefulRestart.EntityData.SegmentPath
    gracefulRestart.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gracefulRestart.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gracefulRestart.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gracefulRestart.EntityData.Children = types.NewOrderedMap()
    gracefulRestart.EntityData.Leafs = types.NewOrderedMap()
    gracefulRestart.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", gracefulRestart.Interval})
    gracefulRestart.EntityData.Leafs.Append("strict-lsa-checking", types.YLeaf{"StrictLsaChecking", gracefulRestart.StrictLsaChecking})
    gracefulRestart.EntityData.Leafs.Append("helper", types.YLeaf{"Helper", gracefulRestart.Helper})
    gracefulRestart.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", gracefulRestart.Enable})
    gracefulRestart.EntityData.Leafs.Append("lifetime", types.YLeaf{"Lifetime", gracefulRestart.Lifetime})

    gracefulRestart.EntityData.YListKeys = []string {}

    return &(gracefulRestart.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_DefaultInformation
// Control distribution of default information
type Ospfv3_Processes_Process_DefaultVrf_DefaultInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Distribute a default route.
    Originate Ospfv3_Processes_Process_DefaultVrf_DefaultInformation_Originate
}

func (defaultInformation *Ospfv3_Processes_Process_DefaultVrf_DefaultInformation) GetEntityData() *types.CommonEntityData {
    defaultInformation.EntityData.YFilter = defaultInformation.YFilter
    defaultInformation.EntityData.YangName = "default-information"
    defaultInformation.EntityData.BundleName = "cisco_ios_xr"
    defaultInformation.EntityData.ParentYangName = "default-vrf"
    defaultInformation.EntityData.SegmentPath = "default-information"
    defaultInformation.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/" + defaultInformation.EntityData.SegmentPath
    defaultInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultInformation.EntityData.Children = types.NewOrderedMap()
    defaultInformation.EntityData.Children.Append("originate", types.YChild{"Originate", &defaultInformation.Originate})
    defaultInformation.EntityData.Leafs = types.NewOrderedMap()

    defaultInformation.EntityData.YListKeys = []string {}

    return &(defaultInformation.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_DefaultInformation_Originate
// Distribute a default route
// This type is a presence type.
type Ospfv3_Processes_Process_DefaultVrf_DefaultInformation_Originate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Always advertise default route. The type is bool. This attribute is
    // mandatory.
    Always interface{}

    // OSPFv3 default metric. The type is interface{} with range: 0..16777214.
    Metric interface{}

    // OSPFv3 metric type for default routes. The type is interface{} with range:
    // 1..2.
    MetricType interface{}

    // Tag for default route. The type is interface{} with range: 0..4294967295.
    Tag interface{}

    // Route policy to default-information origination. The type is string.
    RoutePolicyName interface{}
}

func (originate *Ospfv3_Processes_Process_DefaultVrf_DefaultInformation_Originate) GetEntityData() *types.CommonEntityData {
    originate.EntityData.YFilter = originate.YFilter
    originate.EntityData.YangName = "originate"
    originate.EntityData.BundleName = "cisco_ios_xr"
    originate.EntityData.ParentYangName = "default-information"
    originate.EntityData.SegmentPath = "originate"
    originate.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/default-information/" + originate.EntityData.SegmentPath
    originate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    originate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    originate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    originate.EntityData.Children = types.NewOrderedMap()
    originate.EntityData.Leafs = types.NewOrderedMap()
    originate.EntityData.Leafs.Append("always", types.YLeaf{"Always", originate.Always})
    originate.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", originate.Metric})
    originate.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", originate.MetricType})
    originate.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", originate.Tag})
    originate.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", originate.RoutePolicyName})

    originate.EntityData.YListKeys = []string {}

    return &(originate.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_ProcessScope
// Process scope configuration
type Ospfv3_Processes_Process_DefaultVrf_ProcessScope struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fast-reroute configuration.
    FastReroute Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute
}

func (processScope *Ospfv3_Processes_Process_DefaultVrf_ProcessScope) GetEntityData() *types.CommonEntityData {
    processScope.EntityData.YFilter = processScope.YFilter
    processScope.EntityData.YangName = "process-scope"
    processScope.EntityData.BundleName = "cisco_ios_xr"
    processScope.EntityData.ParentYangName = "default-vrf"
    processScope.EntityData.SegmentPath = "process-scope"
    processScope.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/" + processScope.EntityData.SegmentPath
    processScope.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processScope.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processScope.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processScope.EntityData.Children = types.NewOrderedMap()
    processScope.EntityData.Children.Append("fast-reroute", types.YChild{"FastReroute", &processScope.FastReroute})
    processScope.EntityData.Leafs = types.NewOrderedMap()

    processScope.EntityData.YListKeys = []string {}

    return &(processScope.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute
// Fast-reroute configuration
type Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable Fast-reroute per-link or per-prefix. The type is
    // Ospfv3FastReroute.
    FastRerouteEnable interface{}

    // Fast-reroute per-link configuration.
    PerLink Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerLink

    // Fast-reroute per-link configuration.
    PerPrefix Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerPrefix
}

func (fastReroute *Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute) GetEntityData() *types.CommonEntityData {
    fastReroute.EntityData.YFilter = fastReroute.YFilter
    fastReroute.EntityData.YangName = "fast-reroute"
    fastReroute.EntityData.BundleName = "cisco_ios_xr"
    fastReroute.EntityData.ParentYangName = "process-scope"
    fastReroute.EntityData.SegmentPath = "fast-reroute"
    fastReroute.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/process-scope/" + fastReroute.EntityData.SegmentPath
    fastReroute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fastReroute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fastReroute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fastReroute.EntityData.Children = types.NewOrderedMap()
    fastReroute.EntityData.Children.Append("per-link", types.YChild{"PerLink", &fastReroute.PerLink})
    fastReroute.EntityData.Children.Append("per-prefix", types.YChild{"PerPrefix", &fastReroute.PerPrefix})
    fastReroute.EntityData.Leafs = types.NewOrderedMap()
    fastReroute.EntityData.Leafs.Append("fast-reroute-enable", types.YLeaf{"FastRerouteEnable", fastReroute.FastRerouteEnable})

    fastReroute.EntityData.YListKeys = []string {}

    return &(fastReroute.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerLink
// Fast-reroute per-link configuration
type Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Use only interfaces on the candidate list as a backup path. The type is
    // bool. The default value is false.
    FastRerouteUseCandidateOnly interface{}

    // Fast-reroute per-link/per-prefix candidate interface configuration.
    CandidateInterfaces Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerLink_CandidateInterfaces

    // Fast-reroute per-link/per-prefix exclude interface configuration.
    ExcludeInterfaces Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerLink_ExcludeInterfaces
}

func (perLink *Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerLink) GetEntityData() *types.CommonEntityData {
    perLink.EntityData.YFilter = perLink.YFilter
    perLink.EntityData.YangName = "per-link"
    perLink.EntityData.BundleName = "cisco_ios_xr"
    perLink.EntityData.ParentYangName = "fast-reroute"
    perLink.EntityData.SegmentPath = "per-link"
    perLink.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/process-scope/fast-reroute/" + perLink.EntityData.SegmentPath
    perLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    perLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    perLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    perLink.EntityData.Children = types.NewOrderedMap()
    perLink.EntityData.Children.Append("candidate-interfaces", types.YChild{"CandidateInterfaces", &perLink.CandidateInterfaces})
    perLink.EntityData.Children.Append("exclude-interfaces", types.YChild{"ExcludeInterfaces", &perLink.ExcludeInterfaces})
    perLink.EntityData.Leafs = types.NewOrderedMap()
    perLink.EntityData.Leafs.Append("fast-reroute-use-candidate-only", types.YLeaf{"FastRerouteUseCandidateOnly", perLink.FastRerouteUseCandidateOnly})

    perLink.EntityData.YListKeys = []string {}

    return &(perLink.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerLink_CandidateInterfaces
// Fast-reroute per-link/per-prefix candidate
// interface configuration
type Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerLink_CandidateInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Candidate backup interface. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface.
    CandidateInterface []*Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface
}

func (candidateInterfaces *Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerLink_CandidateInterfaces) GetEntityData() *types.CommonEntityData {
    candidateInterfaces.EntityData.YFilter = candidateInterfaces.YFilter
    candidateInterfaces.EntityData.YangName = "candidate-interfaces"
    candidateInterfaces.EntityData.BundleName = "cisco_ios_xr"
    candidateInterfaces.EntityData.ParentYangName = "per-link"
    candidateInterfaces.EntityData.SegmentPath = "candidate-interfaces"
    candidateInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/process-scope/fast-reroute/per-link/" + candidateInterfaces.EntityData.SegmentPath
    candidateInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterfaces.EntityData.Children = types.NewOrderedMap()
    candidateInterfaces.EntityData.Children.Append("candidate-interface", types.YChild{"CandidateInterface", nil})
    for i := range candidateInterfaces.CandidateInterface {
        candidateInterfaces.EntityData.Children.Append(types.GetSegmentPath(candidateInterfaces.CandidateInterface[i]), types.YChild{"CandidateInterface", candidateInterfaces.CandidateInterface[i]})
    }
    candidateInterfaces.EntityData.Leafs = types.NewOrderedMap()

    candidateInterfaces.EntityData.YListKeys = []string {}

    return &(candidateInterfaces.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface
// Candidate backup interface
type Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (candidateInterface *Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface) GetEntityData() *types.CommonEntityData {
    candidateInterface.EntityData.YFilter = candidateInterface.YFilter
    candidateInterface.EntityData.YangName = "candidate-interface"
    candidateInterface.EntityData.BundleName = "cisco_ios_xr"
    candidateInterface.EntityData.ParentYangName = "candidate-interfaces"
    candidateInterface.EntityData.SegmentPath = "candidate-interface" + types.AddKeyToken(candidateInterface.InterfaceName, "interface-name")
    candidateInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/process-scope/fast-reroute/per-link/candidate-interfaces/" + candidateInterface.EntityData.SegmentPath
    candidateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterface.EntityData.Children = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", candidateInterface.InterfaceName})

    candidateInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(candidateInterface.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerLink_ExcludeInterfaces
// Fast-reroute per-link/per-prefix exclude
// interface configuration
type Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerLink_ExcludeInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exclude an interface from becoming a backup. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface.
    ExcludeInterface []*Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface
}

func (excludeInterfaces *Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerLink_ExcludeInterfaces) GetEntityData() *types.CommonEntityData {
    excludeInterfaces.EntityData.YFilter = excludeInterfaces.YFilter
    excludeInterfaces.EntityData.YangName = "exclude-interfaces"
    excludeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    excludeInterfaces.EntityData.ParentYangName = "per-link"
    excludeInterfaces.EntityData.SegmentPath = "exclude-interfaces"
    excludeInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/process-scope/fast-reroute/per-link/" + excludeInterfaces.EntityData.SegmentPath
    excludeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterfaces.EntityData.Children = types.NewOrderedMap()
    excludeInterfaces.EntityData.Children.Append("exclude-interface", types.YChild{"ExcludeInterface", nil})
    for i := range excludeInterfaces.ExcludeInterface {
        excludeInterfaces.EntityData.Children.Append(types.GetSegmentPath(excludeInterfaces.ExcludeInterface[i]), types.YChild{"ExcludeInterface", excludeInterfaces.ExcludeInterface[i]})
    }
    excludeInterfaces.EntityData.Leafs = types.NewOrderedMap()

    excludeInterfaces.EntityData.YListKeys = []string {}

    return &(excludeInterfaces.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface
// Exclude an interface from becoming a backup
type Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (excludeInterface *Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface) GetEntityData() *types.CommonEntityData {
    excludeInterface.EntityData.YFilter = excludeInterface.YFilter
    excludeInterface.EntityData.YangName = "exclude-interface"
    excludeInterface.EntityData.BundleName = "cisco_ios_xr"
    excludeInterface.EntityData.ParentYangName = "exclude-interfaces"
    excludeInterface.EntityData.SegmentPath = "exclude-interface" + types.AddKeyToken(excludeInterface.InterfaceName, "interface-name")
    excludeInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/process-scope/fast-reroute/per-link/exclude-interfaces/" + excludeInterface.EntityData.SegmentPath
    excludeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterface.EntityData.Children = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", excludeInterface.InterfaceName})

    excludeInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(excludeInterface.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerPrefix
// Fast-reroute per-link configuration
type Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Use only interfaces on the candidate list as a backup path. The type is
    // bool. The default value is false.
    FastRerouteUseCandidateOnly interface{}

    // Fast-reroute per-link/per-prefix candidate interface configuration.
    CandidateInterfaces Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerPrefix_CandidateInterfaces

    // Fast-reroute per-link/per-prefix exclude interface configuration.
    ExcludeInterfaces Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerPrefix_ExcludeInterfaces
}

func (perPrefix *Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerPrefix) GetEntityData() *types.CommonEntityData {
    perPrefix.EntityData.YFilter = perPrefix.YFilter
    perPrefix.EntityData.YangName = "per-prefix"
    perPrefix.EntityData.BundleName = "cisco_ios_xr"
    perPrefix.EntityData.ParentYangName = "fast-reroute"
    perPrefix.EntityData.SegmentPath = "per-prefix"
    perPrefix.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/process-scope/fast-reroute/" + perPrefix.EntityData.SegmentPath
    perPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    perPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    perPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    perPrefix.EntityData.Children = types.NewOrderedMap()
    perPrefix.EntityData.Children.Append("candidate-interfaces", types.YChild{"CandidateInterfaces", &perPrefix.CandidateInterfaces})
    perPrefix.EntityData.Children.Append("exclude-interfaces", types.YChild{"ExcludeInterfaces", &perPrefix.ExcludeInterfaces})
    perPrefix.EntityData.Leafs = types.NewOrderedMap()
    perPrefix.EntityData.Leafs.Append("fast-reroute-use-candidate-only", types.YLeaf{"FastRerouteUseCandidateOnly", perPrefix.FastRerouteUseCandidateOnly})

    perPrefix.EntityData.YListKeys = []string {}

    return &(perPrefix.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerPrefix_CandidateInterfaces
// Fast-reroute per-link/per-prefix candidate
// interface configuration
type Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerPrefix_CandidateInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Candidate backup interface. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface.
    CandidateInterface []*Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface
}

func (candidateInterfaces *Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerPrefix_CandidateInterfaces) GetEntityData() *types.CommonEntityData {
    candidateInterfaces.EntityData.YFilter = candidateInterfaces.YFilter
    candidateInterfaces.EntityData.YangName = "candidate-interfaces"
    candidateInterfaces.EntityData.BundleName = "cisco_ios_xr"
    candidateInterfaces.EntityData.ParentYangName = "per-prefix"
    candidateInterfaces.EntityData.SegmentPath = "candidate-interfaces"
    candidateInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/process-scope/fast-reroute/per-prefix/" + candidateInterfaces.EntityData.SegmentPath
    candidateInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterfaces.EntityData.Children = types.NewOrderedMap()
    candidateInterfaces.EntityData.Children.Append("candidate-interface", types.YChild{"CandidateInterface", nil})
    for i := range candidateInterfaces.CandidateInterface {
        candidateInterfaces.EntityData.Children.Append(types.GetSegmentPath(candidateInterfaces.CandidateInterface[i]), types.YChild{"CandidateInterface", candidateInterfaces.CandidateInterface[i]})
    }
    candidateInterfaces.EntityData.Leafs = types.NewOrderedMap()

    candidateInterfaces.EntityData.YListKeys = []string {}

    return &(candidateInterfaces.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface
// Candidate backup interface
type Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (candidateInterface *Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface) GetEntityData() *types.CommonEntityData {
    candidateInterface.EntityData.YFilter = candidateInterface.YFilter
    candidateInterface.EntityData.YangName = "candidate-interface"
    candidateInterface.EntityData.BundleName = "cisco_ios_xr"
    candidateInterface.EntityData.ParentYangName = "candidate-interfaces"
    candidateInterface.EntityData.SegmentPath = "candidate-interface" + types.AddKeyToken(candidateInterface.InterfaceName, "interface-name")
    candidateInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/process-scope/fast-reroute/per-prefix/candidate-interfaces/" + candidateInterface.EntityData.SegmentPath
    candidateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterface.EntityData.Children = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", candidateInterface.InterfaceName})

    candidateInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(candidateInterface.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerPrefix_ExcludeInterfaces
// Fast-reroute per-link/per-prefix exclude
// interface configuration
type Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerPrefix_ExcludeInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exclude an interface from becoming a backup. The type is slice of
    // Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface.
    ExcludeInterface []*Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface
}

func (excludeInterfaces *Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerPrefix_ExcludeInterfaces) GetEntityData() *types.CommonEntityData {
    excludeInterfaces.EntityData.YFilter = excludeInterfaces.YFilter
    excludeInterfaces.EntityData.YangName = "exclude-interfaces"
    excludeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    excludeInterfaces.EntityData.ParentYangName = "per-prefix"
    excludeInterfaces.EntityData.SegmentPath = "exclude-interfaces"
    excludeInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/process-scope/fast-reroute/per-prefix/" + excludeInterfaces.EntityData.SegmentPath
    excludeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterfaces.EntityData.Children = types.NewOrderedMap()
    excludeInterfaces.EntityData.Children.Append("exclude-interface", types.YChild{"ExcludeInterface", nil})
    for i := range excludeInterfaces.ExcludeInterface {
        excludeInterfaces.EntityData.Children.Append(types.GetSegmentPath(excludeInterfaces.ExcludeInterface[i]), types.YChild{"ExcludeInterface", excludeInterfaces.ExcludeInterface[i]})
    }
    excludeInterfaces.EntityData.Leafs = types.NewOrderedMap()

    excludeInterfaces.EntityData.YListKeys = []string {}

    return &(excludeInterfaces.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface
// Exclude an interface from becoming a backup
type Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (excludeInterface *Ospfv3_Processes_Process_DefaultVrf_ProcessScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface) GetEntityData() *types.CommonEntityData {
    excludeInterface.EntityData.YFilter = excludeInterface.YFilter
    excludeInterface.EntityData.YangName = "exclude-interface"
    excludeInterface.EntityData.BundleName = "cisco_ios_xr"
    excludeInterface.EntityData.ParentYangName = "exclude-interfaces"
    excludeInterface.EntityData.SegmentPath = "exclude-interface" + types.AddKeyToken(excludeInterface.InterfaceName, "interface-name")
    excludeInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/process-scope/fast-reroute/per-prefix/exclude-interfaces/" + excludeInterface.EntityData.SegmentPath
    excludeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterface.EntityData.Children = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", excludeInterface.InterfaceName})

    excludeInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(excludeInterface.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_Encryption
// Encrypt and authenticate OSPFv3 packets
type Ospfv3_Processes_Process_DefaultVrf_Encryption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec ESP authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Specify the encryption algorithm. The type is Ospfv3EncryptionAlgorithm.
    EncryptionAlgorithm interface{}

    // Encryption password. The type is string with pattern: b'(!.+)|([^!].+)'.
    EncryptionPassword interface{}

    // Use the NULL, MD5 or SHA1 algorithm. The type is Ospfv3AuthenticationType2.
    AuthenticationAlgorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    AuthenticationPassword interface{}
}

func (encryption *Ospfv3_Processes_Process_DefaultVrf_Encryption) GetEntityData() *types.CommonEntityData {
    encryption.EntityData.YFilter = encryption.YFilter
    encryption.EntityData.YangName = "encryption"
    encryption.EntityData.BundleName = "cisco_ios_xr"
    encryption.EntityData.ParentYangName = "default-vrf"
    encryption.EntityData.SegmentPath = "encryption"
    encryption.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/" + encryption.EntityData.SegmentPath
    encryption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryption.EntityData.Children = types.NewOrderedMap()
    encryption.EntityData.Leafs = types.NewOrderedMap()
    encryption.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", encryption.Enable})
    encryption.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", encryption.Spi})
    encryption.EntityData.Leafs.Append("encryption-algorithm", types.YLeaf{"EncryptionAlgorithm", encryption.EncryptionAlgorithm})
    encryption.EntityData.Leafs.Append("encryption-password", types.YLeaf{"EncryptionPassword", encryption.EncryptionPassword})
    encryption.EntityData.Leafs.Append("authentication-algorithm", types.YLeaf{"AuthenticationAlgorithm", encryption.AuthenticationAlgorithm})
    encryption.EntityData.Leafs.Append("authentication-password", types.YLeaf{"AuthenticationPassword", encryption.AuthenticationPassword})

    encryption.EntityData.YListKeys = []string {}

    return &(encryption.EntityData)
}

// Ospfv3_Processes_Process_DefaultVrf_AutoCost
// Calculate interface cost according to bandwidth
// This type is a presence type.
type Ospfv3_Processes_Process_DefaultVrf_AutoCost struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Specify 'true' to assign cost based on interface type. The type is
    // interface{}.
    Disable interface{}

    // Specify reference bandwidth for cost computations in terms of Mbits per
    // second. The type is interface{} with range: 1..4294967. Units are Mbit/s.
    ReferenceBandwidth interface{}
}

func (autoCost *Ospfv3_Processes_Process_DefaultVrf_AutoCost) GetEntityData() *types.CommonEntityData {
    autoCost.EntityData.YFilter = autoCost.YFilter
    autoCost.EntityData.YangName = "auto-cost"
    autoCost.EntityData.BundleName = "cisco_ios_xr"
    autoCost.EntityData.ParentYangName = "default-vrf"
    autoCost.EntityData.SegmentPath = "auto-cost"
    autoCost.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/default-vrf/" + autoCost.EntityData.SegmentPath
    autoCost.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoCost.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoCost.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoCost.EntityData.Children = types.NewOrderedMap()
    autoCost.EntityData.Leafs = types.NewOrderedMap()
    autoCost.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", autoCost.Disable})
    autoCost.EntityData.Leafs.Append("reference-bandwidth", types.YLeaf{"ReferenceBandwidth", autoCost.ReferenceBandwidth})

    autoCost.EntityData.YListKeys = []string {}

    return &(autoCost.EntityData)
}

// Ospfv3_Processes_Process_Vrfs
// VRF related configuration
type Ospfv3_Processes_Process_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for a particular OSPF VRF. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf.
    Vrf []*Ospfv3_Processes_Process_Vrfs_Vrf
}

func (vrfs *Ospfv3_Processes_Process_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "process"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/" + vrfs.EntityData.SegmentPath
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

// Ospfv3_Processes_Process_Vrfs_Vrf
// Configuration for a particular OSPF VRF
type Ospfv3_Processes_Process_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name for this VRF. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // Enable SNMP trap configuration in a VRF. The type is interface{}.
    SnmpvrfTrap interface{}

    // Enable/disable prefix suppression on an interface. The type is bool.
    PrefixSuppression interface{}

    // Specify the transmit interval in seconds. The type is interface{} with
    // range: 1..65535. Units are second.
    RetransmitInterval interface{}

    // Enable/disable routing updates on an interface. The type is bool.
    Passive interface{}

    // Set metric of redistributed routes. The type is interface{} with range:
    // 1..16777214.
    DefaultMetric interface{}

    // Enable/disable flood reduction. The type is bool.
    FloodReduction interface{}

    // Time between HELLO packets. The type is interface{} with range: 1..65535.
    // Units are second.
    HelloInterval interface{}

    // Specify router priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // Interface cost. The type is interface{} with range: 1..65535.
    Cost interface{}

    // Interval after which a neighbor is declared dead (in seconds). The type is
    // interface{} with range: 1..65535. Units are second.
    DeadInterval interface{}

    // Limit size of OSPFv3 packets. The type is interface{} with range:
    // 256..10000.
    PacketSize interface{}

    // Instance ID. The type is interface{} with range: 0..255.
    Instance interface{}

    // Route policy for SPF prefix prioritization. The type is string.
    SpfPrefixPriorityPolicy interface{}

    // Specify the router ID for this OSPFv3 process in IPv4 address format. The
    // type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    RouterId interface{}

    // Specify network type. The type is Ospfv3Network.
    Network interface{}

    // Enable/disable ignoring of MTU in DBD packets. The type is bool.
    MtuIgnore interface{}

    // Log changes in adjacency state. The type is Ospfv3LogAdj.
    LogAdjacencyChanges interface{}

    // Enable/disable demand circuit operation. The type is bool.
    DemandCircuit interface{}

    // Specify the transmit delay in seconds. The type is interface{} with range:
    // 1..65535. Units are second.
    TransmitDelay interface{}

    // OSPFv3 Domain ID.
    DomainId Ospfv3_Processes_Process_Vrfs_Vrf_DomainId

    // Area configuration.
    AreaAddresses Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses

    // Adjust routing timers.
    Timers Ospfv3_Processes_Process_Vrfs_Vrf_Timers

    // Summarize redistributed routes matching prefix/length.
    SummaryPrefixes Ospfv3_Processes_Process_Vrfs_Vrf_SummaryPrefixes

    // SNMP configuration.
    Snmp Ospfv3_Processes_Process_Vrfs_Vrf_Snmp

    // Fast-reroute instance scoped parameters.
    FastReroute Ospfv3_Processes_Process_Vrfs_Vrf_FastReroute

    // Define an administrative distance.
    Distance Ospfv3_Processes_Process_Vrfs_Vrf_Distance

    // Set OSPFv3 limits.
    Maximum Ospfv3_Processes_Process_Vrfs_Vrf_Maximum

    // Redistribute information from another routing protocol.
    Redistributes Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes

    // Do not complain about a specified event.
    Ignore Ospfv3_Processes_Process_Vrfs_Vrf_Ignore

    // Filter prefixes from RIB .
    DistributeListOut Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut

    // Filter prefixes to/from RIB.
    DistributeList Ospfv3_Processes_Process_Vrfs_Vrf_DistributeList

    // Stub router configuration.
    StubRouter Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter

    // Configure BFD parameters.
    Bfd Ospfv3_Processes_Process_Vrfs_Vrf_Bfd

    // Database filter.
    DatabaseFilter Ospfv3_Processes_Process_Vrfs_Vrf_DatabaseFilter

    // OSPFv3 Capability.
    Capability Ospfv3_Processes_Process_Vrfs_Vrf_Capability

    // Authenticate OSPFv3 packets.
    Authentication Ospfv3_Processes_Process_Vrfs_Vrf_Authentication

    // Graceful restart configuration.
    GracefulRestart Ospfv3_Processes_Process_Vrfs_Vrf_GracefulRestart

    // Control distribution of default information.
    DefaultInformation Ospfv3_Processes_Process_Vrfs_Vrf_DefaultInformation

    // Process scope configuration.
    ProcessScope Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope

    // Encrypt and authenticate OSPFv3 packets.
    Encryption Ospfv3_Processes_Process_Vrfs_Vrf_Encryption

    // Calculate interface cost according to bandwidth.
    AutoCost Ospfv3_Processes_Process_Vrfs_Vrf_AutoCost
}

func (vrf *Ospfv3_Processes_Process_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/" + vrf.EntityData.SegmentPath
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Children.Append("domain-id", types.YChild{"DomainId", &vrf.DomainId})
    vrf.EntityData.Children.Append("area-addresses", types.YChild{"AreaAddresses", &vrf.AreaAddresses})
    vrf.EntityData.Children.Append("timers", types.YChild{"Timers", &vrf.Timers})
    vrf.EntityData.Children.Append("summary-prefixes", types.YChild{"SummaryPrefixes", &vrf.SummaryPrefixes})
    vrf.EntityData.Children.Append("snmp", types.YChild{"Snmp", &vrf.Snmp})
    vrf.EntityData.Children.Append("fast-reroute", types.YChild{"FastReroute", &vrf.FastReroute})
    vrf.EntityData.Children.Append("distance", types.YChild{"Distance", &vrf.Distance})
    vrf.EntityData.Children.Append("maximum", types.YChild{"Maximum", &vrf.Maximum})
    vrf.EntityData.Children.Append("redistributes", types.YChild{"Redistributes", &vrf.Redistributes})
    vrf.EntityData.Children.Append("ignore", types.YChild{"Ignore", &vrf.Ignore})
    vrf.EntityData.Children.Append("distribute-list-out", types.YChild{"DistributeListOut", &vrf.DistributeListOut})
    vrf.EntityData.Children.Append("distribute-list", types.YChild{"DistributeList", &vrf.DistributeList})
    vrf.EntityData.Children.Append("stub-router", types.YChild{"StubRouter", &vrf.StubRouter})
    vrf.EntityData.Children.Append("bfd", types.YChild{"Bfd", &vrf.Bfd})
    vrf.EntityData.Children.Append("database-filter", types.YChild{"DatabaseFilter", &vrf.DatabaseFilter})
    vrf.EntityData.Children.Append("capability", types.YChild{"Capability", &vrf.Capability})
    vrf.EntityData.Children.Append("authentication", types.YChild{"Authentication", &vrf.Authentication})
    vrf.EntityData.Children.Append("graceful-restart", types.YChild{"GracefulRestart", &vrf.GracefulRestart})
    vrf.EntityData.Children.Append("default-information", types.YChild{"DefaultInformation", &vrf.DefaultInformation})
    vrf.EntityData.Children.Append("process-scope", types.YChild{"ProcessScope", &vrf.ProcessScope})
    vrf.EntityData.Children.Append("encryption", types.YChild{"Encryption", &vrf.Encryption})
    vrf.EntityData.Children.Append("auto-cost", types.YChild{"AutoCost", &vrf.AutoCost})
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})
    vrf.EntityData.Leafs.Append("snmpvrf-trap", types.YLeaf{"SnmpvrfTrap", vrf.SnmpvrfTrap})
    vrf.EntityData.Leafs.Append("prefix-suppression", types.YLeaf{"PrefixSuppression", vrf.PrefixSuppression})
    vrf.EntityData.Leafs.Append("retransmit-interval", types.YLeaf{"RetransmitInterval", vrf.RetransmitInterval})
    vrf.EntityData.Leafs.Append("passive", types.YLeaf{"Passive", vrf.Passive})
    vrf.EntityData.Leafs.Append("default-metric", types.YLeaf{"DefaultMetric", vrf.DefaultMetric})
    vrf.EntityData.Leafs.Append("flood-reduction", types.YLeaf{"FloodReduction", vrf.FloodReduction})
    vrf.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", vrf.HelloInterval})
    vrf.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", vrf.Priority})
    vrf.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", vrf.Cost})
    vrf.EntityData.Leafs.Append("dead-interval", types.YLeaf{"DeadInterval", vrf.DeadInterval})
    vrf.EntityData.Leafs.Append("packet-size", types.YLeaf{"PacketSize", vrf.PacketSize})
    vrf.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", vrf.Instance})
    vrf.EntityData.Leafs.Append("spf-prefix-priority-policy", types.YLeaf{"SpfPrefixPriorityPolicy", vrf.SpfPrefixPriorityPolicy})
    vrf.EntityData.Leafs.Append("router-id", types.YLeaf{"RouterId", vrf.RouterId})
    vrf.EntityData.Leafs.Append("network", types.YLeaf{"Network", vrf.Network})
    vrf.EntityData.Leafs.Append("mtu-ignore", types.YLeaf{"MtuIgnore", vrf.MtuIgnore})
    vrf.EntityData.Leafs.Append("log-adjacency-changes", types.YLeaf{"LogAdjacencyChanges", vrf.LogAdjacencyChanges})
    vrf.EntityData.Leafs.Append("demand-circuit", types.YLeaf{"DemandCircuit", vrf.DemandCircuit})
    vrf.EntityData.Leafs.Append("transmit-delay", types.YLeaf{"TransmitDelay", vrf.TransmitDelay})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_DomainId
// OSPFv3 Domain ID
type Ospfv3_Processes_Process_Vrfs_Vrf_DomainId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Secondary domain ID Table.
    SecondaryDomainIds Ospfv3_Processes_Process_Vrfs_Vrf_DomainId_SecondaryDomainIds

    // OSPF Primary domain ID.
    PrimaryDomainId Ospfv3_Processes_Process_Vrfs_Vrf_DomainId_PrimaryDomainId
}

func (domainId *Ospfv3_Processes_Process_Vrfs_Vrf_DomainId) GetEntityData() *types.CommonEntityData {
    domainId.EntityData.YFilter = domainId.YFilter
    domainId.EntityData.YangName = "domain-id"
    domainId.EntityData.BundleName = "cisco_ios_xr"
    domainId.EntityData.ParentYangName = "vrf"
    domainId.EntityData.SegmentPath = "domain-id"
    domainId.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/" + domainId.EntityData.SegmentPath
    domainId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    domainId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    domainId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    domainId.EntityData.Children = types.NewOrderedMap()
    domainId.EntityData.Children.Append("secondary-domain-ids", types.YChild{"SecondaryDomainIds", &domainId.SecondaryDomainIds})
    domainId.EntityData.Children.Append("primary-domain-id", types.YChild{"PrimaryDomainId", &domainId.PrimaryDomainId})
    domainId.EntityData.Leafs = types.NewOrderedMap()

    domainId.EntityData.YListKeys = []string {}

    return &(domainId.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_DomainId_SecondaryDomainIds
// Secondary domain ID Table
type Ospfv3_Processes_Process_Vrfs_Vrf_DomainId_SecondaryDomainIds struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPF Secondary domain ID. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_DomainId_SecondaryDomainIds_SecondaryDomainId.
    SecondaryDomainId []*Ospfv3_Processes_Process_Vrfs_Vrf_DomainId_SecondaryDomainIds_SecondaryDomainId
}

func (secondaryDomainIds *Ospfv3_Processes_Process_Vrfs_Vrf_DomainId_SecondaryDomainIds) GetEntityData() *types.CommonEntityData {
    secondaryDomainIds.EntityData.YFilter = secondaryDomainIds.YFilter
    secondaryDomainIds.EntityData.YangName = "secondary-domain-ids"
    secondaryDomainIds.EntityData.BundleName = "cisco_ios_xr"
    secondaryDomainIds.EntityData.ParentYangName = "domain-id"
    secondaryDomainIds.EntityData.SegmentPath = "secondary-domain-ids"
    secondaryDomainIds.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/domain-id/" + secondaryDomainIds.EntityData.SegmentPath
    secondaryDomainIds.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secondaryDomainIds.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secondaryDomainIds.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secondaryDomainIds.EntityData.Children = types.NewOrderedMap()
    secondaryDomainIds.EntityData.Children.Append("secondary-domain-id", types.YChild{"SecondaryDomainId", nil})
    for i := range secondaryDomainIds.SecondaryDomainId {
        secondaryDomainIds.EntityData.Children.Append(types.GetSegmentPath(secondaryDomainIds.SecondaryDomainId[i]), types.YChild{"SecondaryDomainId", secondaryDomainIds.SecondaryDomainId[i]})
    }
    secondaryDomainIds.EntityData.Leafs = types.NewOrderedMap()

    secondaryDomainIds.EntityData.YListKeys = []string {}

    return &(secondaryDomainIds.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_DomainId_SecondaryDomainIds_SecondaryDomainId
// OSPF Secondary domain ID
type Ospfv3_Processes_Process_Vrfs_Vrf_DomainId_SecondaryDomainIds_SecondaryDomainId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Secondary domain ID type. The type is
    // Ospfv3DomainId.
    DomainIdType interface{}

    // This attribute is a key. Secondary domain ID value. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    DomainIdName interface{}
}

func (secondaryDomainId *Ospfv3_Processes_Process_Vrfs_Vrf_DomainId_SecondaryDomainIds_SecondaryDomainId) GetEntityData() *types.CommonEntityData {
    secondaryDomainId.EntityData.YFilter = secondaryDomainId.YFilter
    secondaryDomainId.EntityData.YangName = "secondary-domain-id"
    secondaryDomainId.EntityData.BundleName = "cisco_ios_xr"
    secondaryDomainId.EntityData.ParentYangName = "secondary-domain-ids"
    secondaryDomainId.EntityData.SegmentPath = "secondary-domain-id" + types.AddKeyToken(secondaryDomainId.DomainIdType, "domain-id-type") + types.AddKeyToken(secondaryDomainId.DomainIdName, "domain-id-name")
    secondaryDomainId.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/domain-id/secondary-domain-ids/" + secondaryDomainId.EntityData.SegmentPath
    secondaryDomainId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secondaryDomainId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secondaryDomainId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secondaryDomainId.EntityData.Children = types.NewOrderedMap()
    secondaryDomainId.EntityData.Leafs = types.NewOrderedMap()
    secondaryDomainId.EntityData.Leafs.Append("domain-id-type", types.YLeaf{"DomainIdType", secondaryDomainId.DomainIdType})
    secondaryDomainId.EntityData.Leafs.Append("domain-id-name", types.YLeaf{"DomainIdName", secondaryDomainId.DomainIdName})

    secondaryDomainId.EntityData.YListKeys = []string {"DomainIdType", "DomainIdName"}

    return &(secondaryDomainId.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_DomainId_PrimaryDomainId
// OSPF Primary domain ID
type Ospfv3_Processes_Process_Vrfs_Vrf_DomainId_PrimaryDomainId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Primary domain ID type. The type is Ospfv3DomainId.
    DomainIdType interface{}

    // Primary domain ID value. The type is string.
    DomainIdName interface{}
}

func (primaryDomainId *Ospfv3_Processes_Process_Vrfs_Vrf_DomainId_PrimaryDomainId) GetEntityData() *types.CommonEntityData {
    primaryDomainId.EntityData.YFilter = primaryDomainId.YFilter
    primaryDomainId.EntityData.YangName = "primary-domain-id"
    primaryDomainId.EntityData.BundleName = "cisco_ios_xr"
    primaryDomainId.EntityData.ParentYangName = "domain-id"
    primaryDomainId.EntityData.SegmentPath = "primary-domain-id"
    primaryDomainId.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/domain-id/" + primaryDomainId.EntityData.SegmentPath
    primaryDomainId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    primaryDomainId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    primaryDomainId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    primaryDomainId.EntityData.Children = types.NewOrderedMap()
    primaryDomainId.EntityData.Leafs = types.NewOrderedMap()
    primaryDomainId.EntityData.Leafs.Append("domain-id-type", types.YLeaf{"DomainIdType", primaryDomainId.DomainIdType})
    primaryDomainId.EntityData.Leafs.Append("domain-id-name", types.YLeaf{"DomainIdName", primaryDomainId.DomainIdName})

    primaryDomainId.EntityData.YListKeys = []string {}

    return &(primaryDomainId.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses
// Area configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for a particular area. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress.
    AreaAddress []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress

    // Configuration for a particular area. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId.
    AreaAreaId []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId
}

func (areaAddresses *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses) GetEntityData() *types.CommonEntityData {
    areaAddresses.EntityData.YFilter = areaAddresses.YFilter
    areaAddresses.EntityData.YangName = "area-addresses"
    areaAddresses.EntityData.BundleName = "cisco_ios_xr"
    areaAddresses.EntityData.ParentYangName = "vrf"
    areaAddresses.EntityData.SegmentPath = "area-addresses"
    areaAddresses.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/" + areaAddresses.EntityData.SegmentPath
    areaAddresses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    areaAddresses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    areaAddresses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    areaAddresses.EntityData.Children = types.NewOrderedMap()
    areaAddresses.EntityData.Children.Append("area-address", types.YChild{"AreaAddress", nil})
    for i := range areaAddresses.AreaAddress {
        areaAddresses.EntityData.Children.Append(types.GetSegmentPath(areaAddresses.AreaAddress[i]), types.YChild{"AreaAddress", areaAddresses.AreaAddress[i]})
    }
    areaAddresses.EntityData.Children.Append("area-area-id", types.YChild{"AreaAreaId", nil})
    for i := range areaAddresses.AreaAreaId {
        areaAddresses.EntityData.Children.Append(types.GetSegmentPath(areaAddresses.AreaAreaId[i]), types.YChild{"AreaAreaId", areaAddresses.AreaAreaId[i]})
    }
    areaAddresses.EntityData.Leafs = types.NewOrderedMap()

    areaAddresses.EntityData.YListKeys = []string {}

    return &(areaAddresses.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress
// Configuration for a particular area
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Area ID if in IP address format. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // Specify area as a stub area.  Allowed only in non-backbone areas. The type
    // is bool.
    Stub interface{}

    // Limit size of OSPFv3 packets. The type is interface{} with range:
    // 256..10000.
    PacketSize interface{}

    // Instance ID. The type is interface{} with range: 0..255.
    Instance interface{}

    // Enable/disable demand circuit operation. The type is bool.
    DemandCircuit interface{}

    // Specify router priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // Translate Type 7 to Type 5, even if not elected NSSA translator. The type
    // is bool.
    Type7TranslateAlways interface{}

    // Enable/disable prefix suppression on an interface. The type is bool.
    PrefixSuppression interface{}

    // Enable OSPFv3 area. The type is interface{}.
    Enable interface{}

    // Enable/disable ignoring of MTU in DBD packets. The type is bool.
    MtuIgnore interface{}

    // Enable/disable routing updates on an interface. The type is bool.
    Passive interface{}

    // Time between HELLO packets. The type is interface{} with range: 1..65535.
    // Units are second.
    HelloInterval interface{}

    // Interval after which a neighbor is declared dead (in seconds). The type is
    // interface{} with range: 1..65535. Units are second.
    DeadInterval interface{}

    // Set the summary default-cost of a NSSA/stub area. The type is interface{}
    // with range: 0..16777215.
    DefaultCost interface{}

    // Enable/disable flood reduction. The type is bool.
    FloodReduction interface{}

    // Specify the transmit interval in seconds. The type is interface{} with
    // range: 1..65535. Units are second.
    RetransmitInterval interface{}

    // Enable/Disable MPLS LDP sync. The type is bool.
    LdpSync interface{}

    // Specify network type. The type is Ospfv3Network.
    Network interface{}

    // Specify the transmit delay in seconds. The type is interface{} with range:
    // 1..65535. Units are second.
    TransmitDelay interface{}

    // Interface cost. The type is interface{} with range: 1..65535.
    Cost interface{}

    // Authenticate OSPFv3 packets.
    Authentication Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Authentication

    // Configure BFD parameters.
    Bfd Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Bfd

    // Range configuration.
    Ranges Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Ranges

    // Encrypt and authenticate OSPFv3 packets.
    Encryption Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Encryption

    // Specify area as a NSSA area.  Allowed only in non-backbone areas.
    Nssa Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Nssa

    // Database filter.
    DatabaseFilter Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_DatabaseFilter

    // Filter prefixes to/from RIB.
    DistributeList Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_DistributeList

    // OSPFv3 interfaces.
    Interfaces Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces

    // Area Scope Configuration.
    AreaScope Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope

    // Sham Link sub-mode.
    ShamLinks Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_ShamLinks

    // Virtual link sub-mode.
    VirtualLinks Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_VirtualLinks
}

func (areaAddress *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress) GetEntityData() *types.CommonEntityData {
    areaAddress.EntityData.YFilter = areaAddress.YFilter
    areaAddress.EntityData.YangName = "area-address"
    areaAddress.EntityData.BundleName = "cisco_ios_xr"
    areaAddress.EntityData.ParentYangName = "area-addresses"
    areaAddress.EntityData.SegmentPath = "area-address" + types.AddKeyToken(areaAddress.Address, "address")
    areaAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/" + areaAddress.EntityData.SegmentPath
    areaAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    areaAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    areaAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    areaAddress.EntityData.Children = types.NewOrderedMap()
    areaAddress.EntityData.Children.Append("authentication", types.YChild{"Authentication", &areaAddress.Authentication})
    areaAddress.EntityData.Children.Append("bfd", types.YChild{"Bfd", &areaAddress.Bfd})
    areaAddress.EntityData.Children.Append("ranges", types.YChild{"Ranges", &areaAddress.Ranges})
    areaAddress.EntityData.Children.Append("encryption", types.YChild{"Encryption", &areaAddress.Encryption})
    areaAddress.EntityData.Children.Append("nssa", types.YChild{"Nssa", &areaAddress.Nssa})
    areaAddress.EntityData.Children.Append("database-filter", types.YChild{"DatabaseFilter", &areaAddress.DatabaseFilter})
    areaAddress.EntityData.Children.Append("distribute-list", types.YChild{"DistributeList", &areaAddress.DistributeList})
    areaAddress.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &areaAddress.Interfaces})
    areaAddress.EntityData.Children.Append("area-scope", types.YChild{"AreaScope", &areaAddress.AreaScope})
    areaAddress.EntityData.Children.Append("sham-links", types.YChild{"ShamLinks", &areaAddress.ShamLinks})
    areaAddress.EntityData.Children.Append("virtual-links", types.YChild{"VirtualLinks", &areaAddress.VirtualLinks})
    areaAddress.EntityData.Leafs = types.NewOrderedMap()
    areaAddress.EntityData.Leafs.Append("address", types.YLeaf{"Address", areaAddress.Address})
    areaAddress.EntityData.Leafs.Append("stub", types.YLeaf{"Stub", areaAddress.Stub})
    areaAddress.EntityData.Leafs.Append("packet-size", types.YLeaf{"PacketSize", areaAddress.PacketSize})
    areaAddress.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", areaAddress.Instance})
    areaAddress.EntityData.Leafs.Append("demand-circuit", types.YLeaf{"DemandCircuit", areaAddress.DemandCircuit})
    areaAddress.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", areaAddress.Priority})
    areaAddress.EntityData.Leafs.Append("type7-translate-always", types.YLeaf{"Type7TranslateAlways", areaAddress.Type7TranslateAlways})
    areaAddress.EntityData.Leafs.Append("prefix-suppression", types.YLeaf{"PrefixSuppression", areaAddress.PrefixSuppression})
    areaAddress.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", areaAddress.Enable})
    areaAddress.EntityData.Leafs.Append("mtu-ignore", types.YLeaf{"MtuIgnore", areaAddress.MtuIgnore})
    areaAddress.EntityData.Leafs.Append("passive", types.YLeaf{"Passive", areaAddress.Passive})
    areaAddress.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", areaAddress.HelloInterval})
    areaAddress.EntityData.Leafs.Append("dead-interval", types.YLeaf{"DeadInterval", areaAddress.DeadInterval})
    areaAddress.EntityData.Leafs.Append("default-cost", types.YLeaf{"DefaultCost", areaAddress.DefaultCost})
    areaAddress.EntityData.Leafs.Append("flood-reduction", types.YLeaf{"FloodReduction", areaAddress.FloodReduction})
    areaAddress.EntityData.Leafs.Append("retransmit-interval", types.YLeaf{"RetransmitInterval", areaAddress.RetransmitInterval})
    areaAddress.EntityData.Leafs.Append("ldp-sync", types.YLeaf{"LdpSync", areaAddress.LdpSync})
    areaAddress.EntityData.Leafs.Append("network", types.YLeaf{"Network", areaAddress.Network})
    areaAddress.EntityData.Leafs.Append("transmit-delay", types.YLeaf{"TransmitDelay", areaAddress.TransmitDelay})
    areaAddress.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", areaAddress.Cost})

    areaAddress.EntityData.YListKeys = []string {"Address"}

    return &(areaAddress.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Authentication
// Authenticate OSPFv3 packets
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec AH authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Use the MD5 or SHA1 algorithm. The type is Ospfv3Authentication.
    Algorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    Password interface{}
}

func (authentication *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "area-address"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", authentication.Enable})
    authentication.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", authentication.Spi})
    authentication.EntityData.Leafs.Append("algorithm", types.YLeaf{"Algorithm", authentication.Algorithm})
    authentication.EntityData.Leafs.Append("password", types.YLeaf{"Password", authentication.Password})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Bfd
// Configure BFD parameters
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detect multiplier. The type is interface{} with range: 2..50.
    DetectionMultiplier interface{}

    // Hello interval in milli-seconds. The type is interface{} with range:
    // 3..30000. Units are millisecond.
    Interval interface{}

    // Enable or disable BFD fast detection. The type is Ospfv3bfdEnableMode.
    FastDetectMode interface{}
}

func (bfd *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "area-address"
    bfd.EntityData.SegmentPath = "bfd"
    bfd.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/" + bfd.EntityData.SegmentPath
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = types.NewOrderedMap()
    bfd.EntityData.Leafs = types.NewOrderedMap()
    bfd.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", bfd.DetectionMultiplier})
    bfd.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", bfd.Interval})
    bfd.EntityData.Leafs.Append("fast-detect-mode", types.YLeaf{"FastDetectMode", bfd.FastDetectMode})

    bfd.EntityData.YListKeys = []string {}

    return &(bfd.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Ranges
// Range configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Ranges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summarize inter-area routes matching prefix/length. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Ranges_Range.
    Range []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Ranges_Range
}

func (ranges *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Ranges) GetEntityData() *types.CommonEntityData {
    ranges.EntityData.YFilter = ranges.YFilter
    ranges.EntityData.YangName = "ranges"
    ranges.EntityData.BundleName = "cisco_ios_xr"
    ranges.EntityData.ParentYangName = "area-address"
    ranges.EntityData.SegmentPath = "ranges"
    ranges.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/" + ranges.EntityData.SegmentPath
    ranges.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ranges.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ranges.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ranges.EntityData.Children = types.NewOrderedMap()
    ranges.EntityData.Children.Append("range", types.YChild{"Range", nil})
    for i := range ranges.Range {
        ranges.EntityData.Children.Append(types.GetSegmentPath(ranges.Range[i]), types.YChild{"Range", ranges.Range[i]})
    }
    ranges.EntityData.Leafs = types.NewOrderedMap()

    ranges.EntityData.YListKeys = []string {}

    return &(ranges.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Ranges_Range
// Summarize inter-area routes matching
// prefix/length
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Ranges_Range struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IPv6 prefix format. The type is string with
    // pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Prefix interface{}

    // This attribute is a key. IPV6 prefix length. The type is interface{} with
    // range: 0..128.
    PrefixLength interface{}

    // Do not advertise address range. The type is bool. The default value is
    // false.
    NotAdvertise interface{}

    // Specified metric for this range. The type is interface{} with range:
    // 1..16777214.
    Cost interface{}
}

func (self *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Ranges_Range) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "range"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "ranges"
    self.EntityData.SegmentPath = "range" + types.AddKeyToken(self.Prefix, "prefix") + types.AddKeyToken(self.PrefixLength, "prefix-length")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/ranges/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", self.Prefix})
    self.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", self.PrefixLength})
    self.EntityData.Leafs.Append("not-advertise", types.YLeaf{"NotAdvertise", self.NotAdvertise})
    self.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", self.Cost})

    self.EntityData.YListKeys = []string {"Prefix", "PrefixLength"}

    return &(self.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Encryption
// Encrypt and authenticate OSPFv3 packets
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Encryption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec ESP authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Specify the encryption algorithm. The type is Ospfv3EncryptionAlgorithm.
    EncryptionAlgorithm interface{}

    // Encryption password. The type is string with pattern: b'(!.+)|([^!].+)'.
    EncryptionPassword interface{}

    // Use the NULL, MD5 or SHA1 algorithm. The type is Ospfv3AuthenticationType2.
    AuthenticationAlgorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    AuthenticationPassword interface{}
}

func (encryption *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Encryption) GetEntityData() *types.CommonEntityData {
    encryption.EntityData.YFilter = encryption.YFilter
    encryption.EntityData.YangName = "encryption"
    encryption.EntityData.BundleName = "cisco_ios_xr"
    encryption.EntityData.ParentYangName = "area-address"
    encryption.EntityData.SegmentPath = "encryption"
    encryption.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/" + encryption.EntityData.SegmentPath
    encryption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryption.EntityData.Children = types.NewOrderedMap()
    encryption.EntityData.Leafs = types.NewOrderedMap()
    encryption.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", encryption.Enable})
    encryption.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", encryption.Spi})
    encryption.EntityData.Leafs.Append("encryption-algorithm", types.YLeaf{"EncryptionAlgorithm", encryption.EncryptionAlgorithm})
    encryption.EntityData.Leafs.Append("encryption-password", types.YLeaf{"EncryptionPassword", encryption.EncryptionPassword})
    encryption.EntityData.Leafs.Append("authentication-algorithm", types.YLeaf{"AuthenticationAlgorithm", encryption.AuthenticationAlgorithm})
    encryption.EntityData.Leafs.Append("authentication-password", types.YLeaf{"AuthenticationPassword", encryption.AuthenticationPassword})

    encryption.EntityData.YListKeys = []string {}

    return &(encryption.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Nssa
// Specify area as a NSSA area.  Allowed only in
// non-backbone areas
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Nssa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // No redistribution into this NSSA area. The type is bool. The default value
    // is false.
    NoRedistribution interface{}

    // Originate Type 7 default into NSSA area. The type is bool. The default
    // value is false.
    DefaultInfoOriginate interface{}

    // Only valid with DefaultInfoOriginate. The type is interface{} with range:
    // 0..16777214.
    Metric interface{}

    // Only valid with DefaultInfoOriginate. The type is Ospfv3Metric.
    MetricType interface{}

    // Do not send summary LSA into NSSA. The type is interface{}.
    NoSummary interface{}
}

func (nssa *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Nssa) GetEntityData() *types.CommonEntityData {
    nssa.EntityData.YFilter = nssa.YFilter
    nssa.EntityData.YangName = "nssa"
    nssa.EntityData.BundleName = "cisco_ios_xr"
    nssa.EntityData.ParentYangName = "area-address"
    nssa.EntityData.SegmentPath = "nssa"
    nssa.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/" + nssa.EntityData.SegmentPath
    nssa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nssa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nssa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nssa.EntityData.Children = types.NewOrderedMap()
    nssa.EntityData.Leafs = types.NewOrderedMap()
    nssa.EntityData.Leafs.Append("no-redistribution", types.YLeaf{"NoRedistribution", nssa.NoRedistribution})
    nssa.EntityData.Leafs.Append("default-info-originate", types.YLeaf{"DefaultInfoOriginate", nssa.DefaultInfoOriginate})
    nssa.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", nssa.Metric})
    nssa.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", nssa.MetricType})
    nssa.EntityData.Leafs.Append("no-summary", types.YLeaf{"NoSummary", nssa.NoSummary})

    nssa.EntityData.YListKeys = []string {}

    return &(nssa.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_DatabaseFilter
// Database filter
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_DatabaseFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All.
    All Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_DatabaseFilter_All
}

func (databaseFilter *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_DatabaseFilter) GetEntityData() *types.CommonEntityData {
    databaseFilter.EntityData.YFilter = databaseFilter.YFilter
    databaseFilter.EntityData.YangName = "database-filter"
    databaseFilter.EntityData.BundleName = "cisco_ios_xr"
    databaseFilter.EntityData.ParentYangName = "area-address"
    databaseFilter.EntityData.SegmentPath = "database-filter"
    databaseFilter.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/" + databaseFilter.EntityData.SegmentPath
    databaseFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    databaseFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    databaseFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    databaseFilter.EntityData.Children = types.NewOrderedMap()
    databaseFilter.EntityData.Children.Append("all", types.YChild{"All", &databaseFilter.All})
    databaseFilter.EntityData.Leafs = types.NewOrderedMap()

    databaseFilter.EntityData.YListKeys = []string {}

    return &(databaseFilter.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_DatabaseFilter_All
// All
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_DatabaseFilter_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable or disable database-filter. The type is bool.
    Out interface{}
}

func (all *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_DatabaseFilter_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "database-filter"
    all.EntityData.SegmentPath = "all"
    all.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/database-filter/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Leafs = types.NewOrderedMap()
    all.EntityData.Leafs.Append("out", types.YLeaf{"Out", all.Out})

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_DistributeList
// Filter prefixes to/from RIB
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_DistributeList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter prefixes installed to RIB.
    In Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_DistributeList_In
}

func (distributeList *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_DistributeList) GetEntityData() *types.CommonEntityData {
    distributeList.EntityData.YFilter = distributeList.YFilter
    distributeList.EntityData.YangName = "distribute-list"
    distributeList.EntityData.BundleName = "cisco_ios_xr"
    distributeList.EntityData.ParentYangName = "area-address"
    distributeList.EntityData.SegmentPath = "distribute-list"
    distributeList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/" + distributeList.EntityData.SegmentPath
    distributeList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    distributeList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    distributeList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    distributeList.EntityData.Children = types.NewOrderedMap()
    distributeList.EntityData.Children.Append("in", types.YChild{"In", &distributeList.In})
    distributeList.EntityData.Leafs = types.NewOrderedMap()

    distributeList.EntityData.YListKeys = []string {}

    return &(distributeList.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_DistributeList_In
// Filter prefixes installed to RIB
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_DistributeList_In struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter prefixes based on an IPv6 prefix-list. The type is string.
    PrefixList interface{}
}

func (in *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_DistributeList_In) GetEntityData() *types.CommonEntityData {
    in.EntityData.YFilter = in.YFilter
    in.EntityData.YangName = "in"
    in.EntityData.BundleName = "cisco_ios_xr"
    in.EntityData.ParentYangName = "distribute-list"
    in.EntityData.SegmentPath = "in"
    in.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/distribute-list/" + in.EntityData.SegmentPath
    in.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    in.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    in.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    in.EntityData.Children = types.NewOrderedMap()
    in.EntityData.Leafs = types.NewOrderedMap()
    in.EntityData.Leafs.Append("prefix-list", types.YLeaf{"PrefixList", in.PrefixList})

    in.EntityData.YListKeys = []string {}

    return &(in.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces
// OSPFv3 interfaces
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPFv3 interface. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface.
    Interface []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface
}

func (interfaces *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "area-address"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/" + interfaces.EntityData.SegmentPath
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface
// OSPFv3 interface
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface to configure. The type is string with
    // pattern: b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Enable OSPFv3 interface. The type is interface{}.
    Enable interface{}

    // Interval after which a neighbor is declared dead (in seconds). The type is
    // interface{} with range: 1..65535. Units are second.
    DeadInterval interface{}

    // Enable/disable flood reduction. The type is bool.
    FloodReduction interface{}

    // Interface cost. The type is interface{} with range: 1..65535.
    Cost interface{}

    // Specify the transmit delay in seconds. The type is interface{} with range:
    // 1..65535. Units are second.
    TransmitDelay interface{}

    // Instance ID. The type is interface{} with range: 0..255.
    Instance interface{}

    // Enable/Disable MPLS LDP sync. The type is bool.
    LdpSync interface{}

    // Enable/disable ignoring of MTU in DBD packets. The type is bool.
    MtuIgnore interface{}

    // Specify the transmit interval in seconds. The type is interface{} with
    // range: 1..65535. Units are second.
    RetransmitInterval interface{}

    // Time between HELLO packets. The type is interface{} with range: 1..65535.
    // Units are second.
    HelloInterval interface{}

    // Enable/disable routing updates on an interface. The type is bool.
    Passive interface{}

    // Limit size of OSPFv3 packets. The type is interface{} with range:
    // 256..10000.
    PacketSize interface{}

    // Enable/disable prefix suppression on an interface. The type is bool.
    PrefixSuppression interface{}

    // Specify router priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // Specify network type. The type is Ospfv3Network.
    Network interface{}

    // Enable/disable demand circuit operation. The type is bool.
    DemandCircuit interface{}

    // Authenticate OSPFv3 packets.
    Authentication Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_Authentication

    // Specify a neighbor router.
    Neighbors Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_Neighbors

    // Encrypt and authenticate OSPFv3 packets.
    Encryption Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_Encryption

    // Configure BFD parameters.
    Bfd Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_Bfd

    // Database filter.
    DatabaseFilter Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_DatabaseFilter

    // Filter prefixes to/from RIB.
    DistributeList Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_DistributeList

    // Fast-reroute configuration.
    FastReroute Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute
}

func (self *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("authentication", types.YChild{"Authentication", &self.Authentication})
    self.EntityData.Children.Append("neighbors", types.YChild{"Neighbors", &self.Neighbors})
    self.EntityData.Children.Append("encryption", types.YChild{"Encryption", &self.Encryption})
    self.EntityData.Children.Append("bfd", types.YChild{"Bfd", &self.Bfd})
    self.EntityData.Children.Append("database-filter", types.YChild{"DatabaseFilter", &self.DatabaseFilter})
    self.EntityData.Children.Append("distribute-list", types.YChild{"DistributeList", &self.DistributeList})
    self.EntityData.Children.Append("fast-reroute", types.YChild{"FastReroute", &self.FastReroute})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", self.Enable})
    self.EntityData.Leafs.Append("dead-interval", types.YLeaf{"DeadInterval", self.DeadInterval})
    self.EntityData.Leafs.Append("flood-reduction", types.YLeaf{"FloodReduction", self.FloodReduction})
    self.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", self.Cost})
    self.EntityData.Leafs.Append("transmit-delay", types.YLeaf{"TransmitDelay", self.TransmitDelay})
    self.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", self.Instance})
    self.EntityData.Leafs.Append("ldp-sync", types.YLeaf{"LdpSync", self.LdpSync})
    self.EntityData.Leafs.Append("mtu-ignore", types.YLeaf{"MtuIgnore", self.MtuIgnore})
    self.EntityData.Leafs.Append("retransmit-interval", types.YLeaf{"RetransmitInterval", self.RetransmitInterval})
    self.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", self.HelloInterval})
    self.EntityData.Leafs.Append("passive", types.YLeaf{"Passive", self.Passive})
    self.EntityData.Leafs.Append("packet-size", types.YLeaf{"PacketSize", self.PacketSize})
    self.EntityData.Leafs.Append("prefix-suppression", types.YLeaf{"PrefixSuppression", self.PrefixSuppression})
    self.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", self.Priority})
    self.EntityData.Leafs.Append("network", types.YLeaf{"Network", self.Network})
    self.EntityData.Leafs.Append("demand-circuit", types.YLeaf{"DemandCircuit", self.DemandCircuit})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_Authentication
// Authenticate OSPFv3 packets
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec AH authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Use the MD5 or SHA1 algorithm. The type is Ospfv3Authentication.
    Algorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    Password interface{}
}

func (authentication *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "interface"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/interfaces/interface/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", authentication.Enable})
    authentication.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", authentication.Spi})
    authentication.EntityData.Leafs.Append("algorithm", types.YLeaf{"Algorithm", authentication.Algorithm})
    authentication.EntityData.Leafs.Append("password", types.YLeaf{"Password", authentication.Password})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_Neighbors
// Specify a neighbor router
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 address. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_Neighbors_Neighbor.
    Neighbor []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_Neighbors_Neighbor
}

func (neighbors *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "cisco_ios_xr"
    neighbors.EntityData.ParentYangName = "interface"
    neighbors.EntityData.SegmentPath = "neighbors"
    neighbors.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/interfaces/interface/" + neighbors.EntityData.SegmentPath
    neighbors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbors.EntityData.Children = types.NewOrderedMap()
    neighbors.EntityData.Children.Append("neighbor", types.YChild{"Neighbor", nil})
    for i := range neighbors.Neighbor {
        neighbors.EntityData.Children.Append(types.GetSegmentPath(neighbors.Neighbor[i]), types.YChild{"Neighbor", neighbors.Neighbor[i]})
    }
    neighbors.EntityData.Leafs = types.NewOrderedMap()

    neighbors.EntityData.YListKeys = []string {}

    return &(neighbors.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_Neighbors_Neighbor
// IPv6 address
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IPV6 address. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    NeighborAddress interface{}

    // OSPFv3 priority of non-broadcast neighbor. The type is interface{} with
    // range: 0..255.
    Priority interface{}

    // OSPFv3 dead-router polling interval (in seconds). The type is interface{}
    // with range: 0..65535. Units are second.
    PollInterval interface{}

    // OSPFv3 cost for point-to-multipoint neighbor. The type is interface{} with
    // range: 1..65535.
    Cost interface{}

    // Filter OSPFv3 LSA during synchronization and flooding for
    // point-to-multipoint neighbor. The type is bool.
    DatabaseFilter interface{}

    // Zone. The type is string.
    Zone interface{}
}

func (neighbor *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_Neighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "neighbors"
    neighbor.EntityData.SegmentPath = "neighbor" + types.AddKeyToken(neighbor.NeighborAddress, "neighbor-address")
    neighbor.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/interfaces/interface/neighbors/" + neighbor.EntityData.SegmentPath
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", neighbor.NeighborAddress})
    neighbor.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", neighbor.Priority})
    neighbor.EntityData.Leafs.Append("poll-interval", types.YLeaf{"PollInterval", neighbor.PollInterval})
    neighbor.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", neighbor.Cost})
    neighbor.EntityData.Leafs.Append("database-filter", types.YLeaf{"DatabaseFilter", neighbor.DatabaseFilter})
    neighbor.EntityData.Leafs.Append("zone", types.YLeaf{"Zone", neighbor.Zone})

    neighbor.EntityData.YListKeys = []string {"NeighborAddress"}

    return &(neighbor.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_Encryption
// Encrypt and authenticate OSPFv3 packets
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_Encryption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec ESP authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Specify the encryption algorithm. The type is Ospfv3EncryptionAlgorithm.
    EncryptionAlgorithm interface{}

    // Encryption password. The type is string with pattern: b'(!.+)|([^!].+)'.
    EncryptionPassword interface{}

    // Use the NULL, MD5 or SHA1 algorithm. The type is Ospfv3AuthenticationType2.
    AuthenticationAlgorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    AuthenticationPassword interface{}
}

func (encryption *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_Encryption) GetEntityData() *types.CommonEntityData {
    encryption.EntityData.YFilter = encryption.YFilter
    encryption.EntityData.YangName = "encryption"
    encryption.EntityData.BundleName = "cisco_ios_xr"
    encryption.EntityData.ParentYangName = "interface"
    encryption.EntityData.SegmentPath = "encryption"
    encryption.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/interfaces/interface/" + encryption.EntityData.SegmentPath
    encryption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryption.EntityData.Children = types.NewOrderedMap()
    encryption.EntityData.Leafs = types.NewOrderedMap()
    encryption.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", encryption.Enable})
    encryption.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", encryption.Spi})
    encryption.EntityData.Leafs.Append("encryption-algorithm", types.YLeaf{"EncryptionAlgorithm", encryption.EncryptionAlgorithm})
    encryption.EntityData.Leafs.Append("encryption-password", types.YLeaf{"EncryptionPassword", encryption.EncryptionPassword})
    encryption.EntityData.Leafs.Append("authentication-algorithm", types.YLeaf{"AuthenticationAlgorithm", encryption.AuthenticationAlgorithm})
    encryption.EntityData.Leafs.Append("authentication-password", types.YLeaf{"AuthenticationPassword", encryption.AuthenticationPassword})

    encryption.EntityData.YListKeys = []string {}

    return &(encryption.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_Bfd
// Configure BFD parameters
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Hello interval in milli-seconds. The type is interface{} with range:
    // 3..30000. Units are millisecond.
    Interval interface{}

    // Enable or disable BFD fast detection. The type is Ospfv3bfdEnableMode.
    FastDetectMode interface{}

    // Detect multiplier. The type is interface{} with range: 2..50.
    DetectionMultiplier interface{}
}

func (bfd *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "interface"
    bfd.EntityData.SegmentPath = "bfd"
    bfd.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/interfaces/interface/" + bfd.EntityData.SegmentPath
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = types.NewOrderedMap()
    bfd.EntityData.Leafs = types.NewOrderedMap()
    bfd.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", bfd.Interval})
    bfd.EntityData.Leafs.Append("fast-detect-mode", types.YLeaf{"FastDetectMode", bfd.FastDetectMode})
    bfd.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", bfd.DetectionMultiplier})

    bfd.EntityData.YListKeys = []string {}

    return &(bfd.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_DatabaseFilter
// Database filter
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_DatabaseFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All.
    All Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_DatabaseFilter_All
}

func (databaseFilter *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_DatabaseFilter) GetEntityData() *types.CommonEntityData {
    databaseFilter.EntityData.YFilter = databaseFilter.YFilter
    databaseFilter.EntityData.YangName = "database-filter"
    databaseFilter.EntityData.BundleName = "cisco_ios_xr"
    databaseFilter.EntityData.ParentYangName = "interface"
    databaseFilter.EntityData.SegmentPath = "database-filter"
    databaseFilter.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/interfaces/interface/" + databaseFilter.EntityData.SegmentPath
    databaseFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    databaseFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    databaseFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    databaseFilter.EntityData.Children = types.NewOrderedMap()
    databaseFilter.EntityData.Children.Append("all", types.YChild{"All", &databaseFilter.All})
    databaseFilter.EntityData.Leafs = types.NewOrderedMap()

    databaseFilter.EntityData.YListKeys = []string {}

    return &(databaseFilter.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_DatabaseFilter_All
// All
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_DatabaseFilter_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable or disable database-filter. The type is bool.
    Out interface{}
}

func (all *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_DatabaseFilter_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "database-filter"
    all.EntityData.SegmentPath = "all"
    all.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/interfaces/interface/database-filter/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Leafs = types.NewOrderedMap()
    all.EntityData.Leafs.Append("out", types.YLeaf{"Out", all.Out})

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_DistributeList
// Filter prefixes to/from RIB
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_DistributeList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter prefixes installed to RIB.
    In Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_DistributeList_In
}

func (distributeList *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_DistributeList) GetEntityData() *types.CommonEntityData {
    distributeList.EntityData.YFilter = distributeList.YFilter
    distributeList.EntityData.YangName = "distribute-list"
    distributeList.EntityData.BundleName = "cisco_ios_xr"
    distributeList.EntityData.ParentYangName = "interface"
    distributeList.EntityData.SegmentPath = "distribute-list"
    distributeList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/interfaces/interface/" + distributeList.EntityData.SegmentPath
    distributeList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    distributeList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    distributeList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    distributeList.EntityData.Children = types.NewOrderedMap()
    distributeList.EntityData.Children.Append("in", types.YChild{"In", &distributeList.In})
    distributeList.EntityData.Leafs = types.NewOrderedMap()

    distributeList.EntityData.YListKeys = []string {}

    return &(distributeList.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_DistributeList_In
// Filter prefixes installed to RIB
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_DistributeList_In struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter prefixes based on an IPv6 prefix-list. The type is string.
    PrefixList interface{}
}

func (in *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_DistributeList_In) GetEntityData() *types.CommonEntityData {
    in.EntityData.YFilter = in.YFilter
    in.EntityData.YangName = "in"
    in.EntityData.BundleName = "cisco_ios_xr"
    in.EntityData.ParentYangName = "distribute-list"
    in.EntityData.SegmentPath = "in"
    in.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/interfaces/interface/distribute-list/" + in.EntityData.SegmentPath
    in.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    in.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    in.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    in.EntityData.Children = types.NewOrderedMap()
    in.EntityData.Leafs = types.NewOrderedMap()
    in.EntityData.Leafs.Append("prefix-list", types.YLeaf{"PrefixList", in.PrefixList})

    in.EntityData.YListKeys = []string {}

    return &(in.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute
// Fast-reroute configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable Fast-reroute per-link or per-prefix. The type is
    // Ospfv3FastReroute.
    FastRerouteEnable interface{}

    // Fast-reroute per-link configuration.
    PerLink Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink

    // Fast-reroute per-link configuration.
    PerPrefix Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix
}

func (fastReroute *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute) GetEntityData() *types.CommonEntityData {
    fastReroute.EntityData.YFilter = fastReroute.YFilter
    fastReroute.EntityData.YangName = "fast-reroute"
    fastReroute.EntityData.BundleName = "cisco_ios_xr"
    fastReroute.EntityData.ParentYangName = "interface"
    fastReroute.EntityData.SegmentPath = "fast-reroute"
    fastReroute.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/interfaces/interface/" + fastReroute.EntityData.SegmentPath
    fastReroute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fastReroute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fastReroute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fastReroute.EntityData.Children = types.NewOrderedMap()
    fastReroute.EntityData.Children.Append("per-link", types.YChild{"PerLink", &fastReroute.PerLink})
    fastReroute.EntityData.Children.Append("per-prefix", types.YChild{"PerPrefix", &fastReroute.PerPrefix})
    fastReroute.EntityData.Leafs = types.NewOrderedMap()
    fastReroute.EntityData.Leafs.Append("fast-reroute-enable", types.YLeaf{"FastRerouteEnable", fastReroute.FastRerouteEnable})

    fastReroute.EntityData.YListKeys = []string {}

    return &(fastReroute.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink
// Fast-reroute per-link configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Use only interfaces on the candidate list as a backup path. The type is
    // bool. The default value is false.
    FastRerouteUseCandidateOnly interface{}

    // Fast-reroute per-link/per-prefix candidate interface configuration.
    CandidateInterfaces Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces

    // Fast-reroute per-link/per-prefix exclude interface configuration.
    ExcludeInterfaces Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces
}

func (perLink *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink) GetEntityData() *types.CommonEntityData {
    perLink.EntityData.YFilter = perLink.YFilter
    perLink.EntityData.YangName = "per-link"
    perLink.EntityData.BundleName = "cisco_ios_xr"
    perLink.EntityData.ParentYangName = "fast-reroute"
    perLink.EntityData.SegmentPath = "per-link"
    perLink.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/interfaces/interface/fast-reroute/" + perLink.EntityData.SegmentPath
    perLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    perLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    perLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    perLink.EntityData.Children = types.NewOrderedMap()
    perLink.EntityData.Children.Append("candidate-interfaces", types.YChild{"CandidateInterfaces", &perLink.CandidateInterfaces})
    perLink.EntityData.Children.Append("exclude-interfaces", types.YChild{"ExcludeInterfaces", &perLink.ExcludeInterfaces})
    perLink.EntityData.Leafs = types.NewOrderedMap()
    perLink.EntityData.Leafs.Append("fast-reroute-use-candidate-only", types.YLeaf{"FastRerouteUseCandidateOnly", perLink.FastRerouteUseCandidateOnly})

    perLink.EntityData.YListKeys = []string {}

    return &(perLink.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces
// Fast-reroute per-link/per-prefix candidate
// interface configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Candidate backup interface. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces_CandidateInterface.
    CandidateInterface []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces_CandidateInterface
}

func (candidateInterfaces *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces) GetEntityData() *types.CommonEntityData {
    candidateInterfaces.EntityData.YFilter = candidateInterfaces.YFilter
    candidateInterfaces.EntityData.YangName = "candidate-interfaces"
    candidateInterfaces.EntityData.BundleName = "cisco_ios_xr"
    candidateInterfaces.EntityData.ParentYangName = "per-link"
    candidateInterfaces.EntityData.SegmentPath = "candidate-interfaces"
    candidateInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/interfaces/interface/fast-reroute/per-link/" + candidateInterfaces.EntityData.SegmentPath
    candidateInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterfaces.EntityData.Children = types.NewOrderedMap()
    candidateInterfaces.EntityData.Children.Append("candidate-interface", types.YChild{"CandidateInterface", nil})
    for i := range candidateInterfaces.CandidateInterface {
        candidateInterfaces.EntityData.Children.Append(types.GetSegmentPath(candidateInterfaces.CandidateInterface[i]), types.YChild{"CandidateInterface", candidateInterfaces.CandidateInterface[i]})
    }
    candidateInterfaces.EntityData.Leafs = types.NewOrderedMap()

    candidateInterfaces.EntityData.YListKeys = []string {}

    return &(candidateInterfaces.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces_CandidateInterface
// Candidate backup interface
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces_CandidateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (candidateInterface *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces_CandidateInterface) GetEntityData() *types.CommonEntityData {
    candidateInterface.EntityData.YFilter = candidateInterface.YFilter
    candidateInterface.EntityData.YangName = "candidate-interface"
    candidateInterface.EntityData.BundleName = "cisco_ios_xr"
    candidateInterface.EntityData.ParentYangName = "candidate-interfaces"
    candidateInterface.EntityData.SegmentPath = "candidate-interface" + types.AddKeyToken(candidateInterface.InterfaceName, "interface-name")
    candidateInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/interfaces/interface/fast-reroute/per-link/candidate-interfaces/" + candidateInterface.EntityData.SegmentPath
    candidateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterface.EntityData.Children = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", candidateInterface.InterfaceName})

    candidateInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(candidateInterface.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces
// Fast-reroute per-link/per-prefix exclude
// interface configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exclude an interface from becoming a backup. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface.
    ExcludeInterface []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface
}

func (excludeInterfaces *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces) GetEntityData() *types.CommonEntityData {
    excludeInterfaces.EntityData.YFilter = excludeInterfaces.YFilter
    excludeInterfaces.EntityData.YangName = "exclude-interfaces"
    excludeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    excludeInterfaces.EntityData.ParentYangName = "per-link"
    excludeInterfaces.EntityData.SegmentPath = "exclude-interfaces"
    excludeInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/interfaces/interface/fast-reroute/per-link/" + excludeInterfaces.EntityData.SegmentPath
    excludeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterfaces.EntityData.Children = types.NewOrderedMap()
    excludeInterfaces.EntityData.Children.Append("exclude-interface", types.YChild{"ExcludeInterface", nil})
    for i := range excludeInterfaces.ExcludeInterface {
        excludeInterfaces.EntityData.Children.Append(types.GetSegmentPath(excludeInterfaces.ExcludeInterface[i]), types.YChild{"ExcludeInterface", excludeInterfaces.ExcludeInterface[i]})
    }
    excludeInterfaces.EntityData.Leafs = types.NewOrderedMap()

    excludeInterfaces.EntityData.YListKeys = []string {}

    return &(excludeInterfaces.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface
// Exclude an interface from becoming a backup
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (excludeInterface *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface) GetEntityData() *types.CommonEntityData {
    excludeInterface.EntityData.YFilter = excludeInterface.YFilter
    excludeInterface.EntityData.YangName = "exclude-interface"
    excludeInterface.EntityData.BundleName = "cisco_ios_xr"
    excludeInterface.EntityData.ParentYangName = "exclude-interfaces"
    excludeInterface.EntityData.SegmentPath = "exclude-interface" + types.AddKeyToken(excludeInterface.InterfaceName, "interface-name")
    excludeInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/interfaces/interface/fast-reroute/per-link/exclude-interfaces/" + excludeInterface.EntityData.SegmentPath
    excludeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterface.EntityData.Children = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", excludeInterface.InterfaceName})

    excludeInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(excludeInterface.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix
// Fast-reroute per-link configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Use only interfaces on the candidate list as a backup path. The type is
    // bool. The default value is false.
    FastRerouteUseCandidateOnly interface{}

    // Fast-reroute per-link/per-prefix candidate interface configuration.
    CandidateInterfaces Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces

    // Fast-reroute per-link/per-prefix exclude interface configuration.
    ExcludeInterfaces Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces
}

func (perPrefix *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix) GetEntityData() *types.CommonEntityData {
    perPrefix.EntityData.YFilter = perPrefix.YFilter
    perPrefix.EntityData.YangName = "per-prefix"
    perPrefix.EntityData.BundleName = "cisco_ios_xr"
    perPrefix.EntityData.ParentYangName = "fast-reroute"
    perPrefix.EntityData.SegmentPath = "per-prefix"
    perPrefix.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/interfaces/interface/fast-reroute/" + perPrefix.EntityData.SegmentPath
    perPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    perPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    perPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    perPrefix.EntityData.Children = types.NewOrderedMap()
    perPrefix.EntityData.Children.Append("candidate-interfaces", types.YChild{"CandidateInterfaces", &perPrefix.CandidateInterfaces})
    perPrefix.EntityData.Children.Append("exclude-interfaces", types.YChild{"ExcludeInterfaces", &perPrefix.ExcludeInterfaces})
    perPrefix.EntityData.Leafs = types.NewOrderedMap()
    perPrefix.EntityData.Leafs.Append("fast-reroute-use-candidate-only", types.YLeaf{"FastRerouteUseCandidateOnly", perPrefix.FastRerouteUseCandidateOnly})

    perPrefix.EntityData.YListKeys = []string {}

    return &(perPrefix.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces
// Fast-reroute per-link/per-prefix candidate
// interface configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Candidate backup interface. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface.
    CandidateInterface []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface
}

func (candidateInterfaces *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces) GetEntityData() *types.CommonEntityData {
    candidateInterfaces.EntityData.YFilter = candidateInterfaces.YFilter
    candidateInterfaces.EntityData.YangName = "candidate-interfaces"
    candidateInterfaces.EntityData.BundleName = "cisco_ios_xr"
    candidateInterfaces.EntityData.ParentYangName = "per-prefix"
    candidateInterfaces.EntityData.SegmentPath = "candidate-interfaces"
    candidateInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/interfaces/interface/fast-reroute/per-prefix/" + candidateInterfaces.EntityData.SegmentPath
    candidateInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterfaces.EntityData.Children = types.NewOrderedMap()
    candidateInterfaces.EntityData.Children.Append("candidate-interface", types.YChild{"CandidateInterface", nil})
    for i := range candidateInterfaces.CandidateInterface {
        candidateInterfaces.EntityData.Children.Append(types.GetSegmentPath(candidateInterfaces.CandidateInterface[i]), types.YChild{"CandidateInterface", candidateInterfaces.CandidateInterface[i]})
    }
    candidateInterfaces.EntityData.Leafs = types.NewOrderedMap()

    candidateInterfaces.EntityData.YListKeys = []string {}

    return &(candidateInterfaces.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface
// Candidate backup interface
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (candidateInterface *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface) GetEntityData() *types.CommonEntityData {
    candidateInterface.EntityData.YFilter = candidateInterface.YFilter
    candidateInterface.EntityData.YangName = "candidate-interface"
    candidateInterface.EntityData.BundleName = "cisco_ios_xr"
    candidateInterface.EntityData.ParentYangName = "candidate-interfaces"
    candidateInterface.EntityData.SegmentPath = "candidate-interface" + types.AddKeyToken(candidateInterface.InterfaceName, "interface-name")
    candidateInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/interfaces/interface/fast-reroute/per-prefix/candidate-interfaces/" + candidateInterface.EntityData.SegmentPath
    candidateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterface.EntityData.Children = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", candidateInterface.InterfaceName})

    candidateInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(candidateInterface.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces
// Fast-reroute per-link/per-prefix exclude
// interface configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exclude an interface from becoming a backup. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface.
    ExcludeInterface []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface
}

func (excludeInterfaces *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces) GetEntityData() *types.CommonEntityData {
    excludeInterfaces.EntityData.YFilter = excludeInterfaces.YFilter
    excludeInterfaces.EntityData.YangName = "exclude-interfaces"
    excludeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    excludeInterfaces.EntityData.ParentYangName = "per-prefix"
    excludeInterfaces.EntityData.SegmentPath = "exclude-interfaces"
    excludeInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/interfaces/interface/fast-reroute/per-prefix/" + excludeInterfaces.EntityData.SegmentPath
    excludeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterfaces.EntityData.Children = types.NewOrderedMap()
    excludeInterfaces.EntityData.Children.Append("exclude-interface", types.YChild{"ExcludeInterface", nil})
    for i := range excludeInterfaces.ExcludeInterface {
        excludeInterfaces.EntityData.Children.Append(types.GetSegmentPath(excludeInterfaces.ExcludeInterface[i]), types.YChild{"ExcludeInterface", excludeInterfaces.ExcludeInterface[i]})
    }
    excludeInterfaces.EntityData.Leafs = types.NewOrderedMap()

    excludeInterfaces.EntityData.YListKeys = []string {}

    return &(excludeInterfaces.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface
// Exclude an interface from becoming a backup
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (excludeInterface *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface) GetEntityData() *types.CommonEntityData {
    excludeInterface.EntityData.YFilter = excludeInterface.YFilter
    excludeInterface.EntityData.YangName = "exclude-interface"
    excludeInterface.EntityData.BundleName = "cisco_ios_xr"
    excludeInterface.EntityData.ParentYangName = "exclude-interfaces"
    excludeInterface.EntityData.SegmentPath = "exclude-interface" + types.AddKeyToken(excludeInterface.InterfaceName, "interface-name")
    excludeInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/interfaces/interface/fast-reroute/per-prefix/exclude-interfaces/" + excludeInterface.EntityData.SegmentPath
    excludeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterface.EntityData.Children = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", excludeInterface.InterfaceName})

    excludeInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(excludeInterface.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope
// Area Scope Configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fast-reroute configuration.
    FastReroute Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute
}

func (areaScope *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope) GetEntityData() *types.CommonEntityData {
    areaScope.EntityData.YFilter = areaScope.YFilter
    areaScope.EntityData.YangName = "area-scope"
    areaScope.EntityData.BundleName = "cisco_ios_xr"
    areaScope.EntityData.ParentYangName = "area-address"
    areaScope.EntityData.SegmentPath = "area-scope"
    areaScope.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/" + areaScope.EntityData.SegmentPath
    areaScope.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    areaScope.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    areaScope.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    areaScope.EntityData.Children = types.NewOrderedMap()
    areaScope.EntityData.Children.Append("fast-reroute", types.YChild{"FastReroute", &areaScope.FastReroute})
    areaScope.EntityData.Leafs = types.NewOrderedMap()

    areaScope.EntityData.YListKeys = []string {}

    return &(areaScope.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute
// Fast-reroute configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable Fast-reroute per-link or per-prefix. The type is
    // Ospfv3FastReroute.
    FastRerouteEnable interface{}

    // Fast-reroute per-link configuration.
    PerLink Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink

    // Fast-reroute per-link configuration.
    PerPrefix Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix
}

func (fastReroute *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute) GetEntityData() *types.CommonEntityData {
    fastReroute.EntityData.YFilter = fastReroute.YFilter
    fastReroute.EntityData.YangName = "fast-reroute"
    fastReroute.EntityData.BundleName = "cisco_ios_xr"
    fastReroute.EntityData.ParentYangName = "area-scope"
    fastReroute.EntityData.SegmentPath = "fast-reroute"
    fastReroute.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/area-scope/" + fastReroute.EntityData.SegmentPath
    fastReroute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fastReroute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fastReroute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fastReroute.EntityData.Children = types.NewOrderedMap()
    fastReroute.EntityData.Children.Append("per-link", types.YChild{"PerLink", &fastReroute.PerLink})
    fastReroute.EntityData.Children.Append("per-prefix", types.YChild{"PerPrefix", &fastReroute.PerPrefix})
    fastReroute.EntityData.Leafs = types.NewOrderedMap()
    fastReroute.EntityData.Leafs.Append("fast-reroute-enable", types.YLeaf{"FastRerouteEnable", fastReroute.FastRerouteEnable})

    fastReroute.EntityData.YListKeys = []string {}

    return &(fastReroute.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink
// Fast-reroute per-link configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Use only interfaces on the candidate list as a backup path. The type is
    // bool. The default value is false.
    FastRerouteUseCandidateOnly interface{}

    // Fast-reroute per-link/per-prefix candidate interface configuration.
    CandidateInterfaces Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_CandidateInterfaces

    // Fast-reroute per-link/per-prefix exclude interface configuration.
    ExcludeInterfaces Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_ExcludeInterfaces
}

func (perLink *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink) GetEntityData() *types.CommonEntityData {
    perLink.EntityData.YFilter = perLink.YFilter
    perLink.EntityData.YangName = "per-link"
    perLink.EntityData.BundleName = "cisco_ios_xr"
    perLink.EntityData.ParentYangName = "fast-reroute"
    perLink.EntityData.SegmentPath = "per-link"
    perLink.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/area-scope/fast-reroute/" + perLink.EntityData.SegmentPath
    perLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    perLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    perLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    perLink.EntityData.Children = types.NewOrderedMap()
    perLink.EntityData.Children.Append("candidate-interfaces", types.YChild{"CandidateInterfaces", &perLink.CandidateInterfaces})
    perLink.EntityData.Children.Append("exclude-interfaces", types.YChild{"ExcludeInterfaces", &perLink.ExcludeInterfaces})
    perLink.EntityData.Leafs = types.NewOrderedMap()
    perLink.EntityData.Leafs.Append("fast-reroute-use-candidate-only", types.YLeaf{"FastRerouteUseCandidateOnly", perLink.FastRerouteUseCandidateOnly})

    perLink.EntityData.YListKeys = []string {}

    return &(perLink.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_CandidateInterfaces
// Fast-reroute per-link/per-prefix candidate
// interface configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_CandidateInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Candidate backup interface. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface.
    CandidateInterface []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface
}

func (candidateInterfaces *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_CandidateInterfaces) GetEntityData() *types.CommonEntityData {
    candidateInterfaces.EntityData.YFilter = candidateInterfaces.YFilter
    candidateInterfaces.EntityData.YangName = "candidate-interfaces"
    candidateInterfaces.EntityData.BundleName = "cisco_ios_xr"
    candidateInterfaces.EntityData.ParentYangName = "per-link"
    candidateInterfaces.EntityData.SegmentPath = "candidate-interfaces"
    candidateInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/area-scope/fast-reroute/per-link/" + candidateInterfaces.EntityData.SegmentPath
    candidateInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterfaces.EntityData.Children = types.NewOrderedMap()
    candidateInterfaces.EntityData.Children.Append("candidate-interface", types.YChild{"CandidateInterface", nil})
    for i := range candidateInterfaces.CandidateInterface {
        candidateInterfaces.EntityData.Children.Append(types.GetSegmentPath(candidateInterfaces.CandidateInterface[i]), types.YChild{"CandidateInterface", candidateInterfaces.CandidateInterface[i]})
    }
    candidateInterfaces.EntityData.Leafs = types.NewOrderedMap()

    candidateInterfaces.EntityData.YListKeys = []string {}

    return &(candidateInterfaces.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface
// Candidate backup interface
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (candidateInterface *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface) GetEntityData() *types.CommonEntityData {
    candidateInterface.EntityData.YFilter = candidateInterface.YFilter
    candidateInterface.EntityData.YangName = "candidate-interface"
    candidateInterface.EntityData.BundleName = "cisco_ios_xr"
    candidateInterface.EntityData.ParentYangName = "candidate-interfaces"
    candidateInterface.EntityData.SegmentPath = "candidate-interface" + types.AddKeyToken(candidateInterface.InterfaceName, "interface-name")
    candidateInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/area-scope/fast-reroute/per-link/candidate-interfaces/" + candidateInterface.EntityData.SegmentPath
    candidateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterface.EntityData.Children = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", candidateInterface.InterfaceName})

    candidateInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(candidateInterface.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_ExcludeInterfaces
// Fast-reroute per-link/per-prefix exclude
// interface configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_ExcludeInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exclude an interface from becoming a backup. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface.
    ExcludeInterface []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface
}

func (excludeInterfaces *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_ExcludeInterfaces) GetEntityData() *types.CommonEntityData {
    excludeInterfaces.EntityData.YFilter = excludeInterfaces.YFilter
    excludeInterfaces.EntityData.YangName = "exclude-interfaces"
    excludeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    excludeInterfaces.EntityData.ParentYangName = "per-link"
    excludeInterfaces.EntityData.SegmentPath = "exclude-interfaces"
    excludeInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/area-scope/fast-reroute/per-link/" + excludeInterfaces.EntityData.SegmentPath
    excludeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterfaces.EntityData.Children = types.NewOrderedMap()
    excludeInterfaces.EntityData.Children.Append("exclude-interface", types.YChild{"ExcludeInterface", nil})
    for i := range excludeInterfaces.ExcludeInterface {
        excludeInterfaces.EntityData.Children.Append(types.GetSegmentPath(excludeInterfaces.ExcludeInterface[i]), types.YChild{"ExcludeInterface", excludeInterfaces.ExcludeInterface[i]})
    }
    excludeInterfaces.EntityData.Leafs = types.NewOrderedMap()

    excludeInterfaces.EntityData.YListKeys = []string {}

    return &(excludeInterfaces.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface
// Exclude an interface from becoming a backup
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (excludeInterface *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface) GetEntityData() *types.CommonEntityData {
    excludeInterface.EntityData.YFilter = excludeInterface.YFilter
    excludeInterface.EntityData.YangName = "exclude-interface"
    excludeInterface.EntityData.BundleName = "cisco_ios_xr"
    excludeInterface.EntityData.ParentYangName = "exclude-interfaces"
    excludeInterface.EntityData.SegmentPath = "exclude-interface" + types.AddKeyToken(excludeInterface.InterfaceName, "interface-name")
    excludeInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/area-scope/fast-reroute/per-link/exclude-interfaces/" + excludeInterface.EntityData.SegmentPath
    excludeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterface.EntityData.Children = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", excludeInterface.InterfaceName})

    excludeInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(excludeInterface.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix
// Fast-reroute per-link configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Use only interfaces on the candidate list as a backup path. The type is
    // bool. The default value is false.
    FastRerouteUseCandidateOnly interface{}

    // Fast-reroute per-link/per-prefix candidate interface configuration.
    CandidateInterfaces Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_CandidateInterfaces

    // Fast-reroute per-link/per-prefix exclude interface configuration.
    ExcludeInterfaces Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces
}

func (perPrefix *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix) GetEntityData() *types.CommonEntityData {
    perPrefix.EntityData.YFilter = perPrefix.YFilter
    perPrefix.EntityData.YangName = "per-prefix"
    perPrefix.EntityData.BundleName = "cisco_ios_xr"
    perPrefix.EntityData.ParentYangName = "fast-reroute"
    perPrefix.EntityData.SegmentPath = "per-prefix"
    perPrefix.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/area-scope/fast-reroute/" + perPrefix.EntityData.SegmentPath
    perPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    perPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    perPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    perPrefix.EntityData.Children = types.NewOrderedMap()
    perPrefix.EntityData.Children.Append("candidate-interfaces", types.YChild{"CandidateInterfaces", &perPrefix.CandidateInterfaces})
    perPrefix.EntityData.Children.Append("exclude-interfaces", types.YChild{"ExcludeInterfaces", &perPrefix.ExcludeInterfaces})
    perPrefix.EntityData.Leafs = types.NewOrderedMap()
    perPrefix.EntityData.Leafs.Append("fast-reroute-use-candidate-only", types.YLeaf{"FastRerouteUseCandidateOnly", perPrefix.FastRerouteUseCandidateOnly})

    perPrefix.EntityData.YListKeys = []string {}

    return &(perPrefix.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_CandidateInterfaces
// Fast-reroute per-link/per-prefix candidate
// interface configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_CandidateInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Candidate backup interface. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface.
    CandidateInterface []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface
}

func (candidateInterfaces *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_CandidateInterfaces) GetEntityData() *types.CommonEntityData {
    candidateInterfaces.EntityData.YFilter = candidateInterfaces.YFilter
    candidateInterfaces.EntityData.YangName = "candidate-interfaces"
    candidateInterfaces.EntityData.BundleName = "cisco_ios_xr"
    candidateInterfaces.EntityData.ParentYangName = "per-prefix"
    candidateInterfaces.EntityData.SegmentPath = "candidate-interfaces"
    candidateInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/area-scope/fast-reroute/per-prefix/" + candidateInterfaces.EntityData.SegmentPath
    candidateInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterfaces.EntityData.Children = types.NewOrderedMap()
    candidateInterfaces.EntityData.Children.Append("candidate-interface", types.YChild{"CandidateInterface", nil})
    for i := range candidateInterfaces.CandidateInterface {
        candidateInterfaces.EntityData.Children.Append(types.GetSegmentPath(candidateInterfaces.CandidateInterface[i]), types.YChild{"CandidateInterface", candidateInterfaces.CandidateInterface[i]})
    }
    candidateInterfaces.EntityData.Leafs = types.NewOrderedMap()

    candidateInterfaces.EntityData.YListKeys = []string {}

    return &(candidateInterfaces.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface
// Candidate backup interface
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (candidateInterface *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface) GetEntityData() *types.CommonEntityData {
    candidateInterface.EntityData.YFilter = candidateInterface.YFilter
    candidateInterface.EntityData.YangName = "candidate-interface"
    candidateInterface.EntityData.BundleName = "cisco_ios_xr"
    candidateInterface.EntityData.ParentYangName = "candidate-interfaces"
    candidateInterface.EntityData.SegmentPath = "candidate-interface" + types.AddKeyToken(candidateInterface.InterfaceName, "interface-name")
    candidateInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/area-scope/fast-reroute/per-prefix/candidate-interfaces/" + candidateInterface.EntityData.SegmentPath
    candidateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterface.EntityData.Children = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", candidateInterface.InterfaceName})

    candidateInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(candidateInterface.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces
// Fast-reroute per-link/per-prefix exclude
// interface configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exclude an interface from becoming a backup. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface.
    ExcludeInterface []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface
}

func (excludeInterfaces *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces) GetEntityData() *types.CommonEntityData {
    excludeInterfaces.EntityData.YFilter = excludeInterfaces.YFilter
    excludeInterfaces.EntityData.YangName = "exclude-interfaces"
    excludeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    excludeInterfaces.EntityData.ParentYangName = "per-prefix"
    excludeInterfaces.EntityData.SegmentPath = "exclude-interfaces"
    excludeInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/area-scope/fast-reroute/per-prefix/" + excludeInterfaces.EntityData.SegmentPath
    excludeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterfaces.EntityData.Children = types.NewOrderedMap()
    excludeInterfaces.EntityData.Children.Append("exclude-interface", types.YChild{"ExcludeInterface", nil})
    for i := range excludeInterfaces.ExcludeInterface {
        excludeInterfaces.EntityData.Children.Append(types.GetSegmentPath(excludeInterfaces.ExcludeInterface[i]), types.YChild{"ExcludeInterface", excludeInterfaces.ExcludeInterface[i]})
    }
    excludeInterfaces.EntityData.Leafs = types.NewOrderedMap()

    excludeInterfaces.EntityData.YListKeys = []string {}

    return &(excludeInterfaces.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface
// Exclude an interface from becoming a backup
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (excludeInterface *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface) GetEntityData() *types.CommonEntityData {
    excludeInterface.EntityData.YFilter = excludeInterface.YFilter
    excludeInterface.EntityData.YangName = "exclude-interface"
    excludeInterface.EntityData.BundleName = "cisco_ios_xr"
    excludeInterface.EntityData.ParentYangName = "exclude-interfaces"
    excludeInterface.EntityData.SegmentPath = "exclude-interface" + types.AddKeyToken(excludeInterface.InterfaceName, "interface-name")
    excludeInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/area-scope/fast-reroute/per-prefix/exclude-interfaces/" + excludeInterface.EntityData.SegmentPath
    excludeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterface.EntityData.Children = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", excludeInterface.InterfaceName})

    excludeInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(excludeInterface.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_ShamLinks
// Sham Link sub-mode
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_ShamLinks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ShamLink local and remote endpoints. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_ShamLinks_ShamLink.
    ShamLink []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_ShamLinks_ShamLink
}

func (shamLinks *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_ShamLinks) GetEntityData() *types.CommonEntityData {
    shamLinks.EntityData.YFilter = shamLinks.YFilter
    shamLinks.EntityData.YangName = "sham-links"
    shamLinks.EntityData.BundleName = "cisco_ios_xr"
    shamLinks.EntityData.ParentYangName = "area-address"
    shamLinks.EntityData.SegmentPath = "sham-links"
    shamLinks.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/" + shamLinks.EntityData.SegmentPath
    shamLinks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shamLinks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shamLinks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shamLinks.EntityData.Children = types.NewOrderedMap()
    shamLinks.EntityData.Children.Append("sham-link", types.YChild{"ShamLink", nil})
    for i := range shamLinks.ShamLink {
        shamLinks.EntityData.Children.Append(types.GetSegmentPath(shamLinks.ShamLink[i]), types.YChild{"ShamLink", shamLinks.ShamLink[i]})
    }
    shamLinks.EntityData.Leafs = types.NewOrderedMap()

    shamLinks.EntityData.YListKeys = []string {}

    return &(shamLinks.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_ShamLinks_ShamLink
// ShamLink local and remote endpoints
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_ShamLinks_ShamLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Local sham-link endpoint. The type is string with
    // pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // This attribute is a key. Remote sham-link endpoint. The type is string with
    // pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    DestinationAddress interface{}

    // Enable sham link. The type is interface{}.
    Enable interface{}

    // Time between HELLO packets. The type is interface{} with range: 1..65535.
    // Units are second.
    HelloInterval interface{}

    // Interval after which a neighbor is declared dead (in seconds). The type is
    // interface{} with range: 1..65535. Units are second.
    DeadInterval interface{}

    // Specify the transmit interval in seconds. The type is interface{} with
    // range: 1..65535. Units are second.
    RetransmitInterval interface{}

    // Specify the transmit delay in seconds. The type is interface{} with range:
    // 1..65535. Units are second.
    TransmitDelay interface{}

    // Authenticate OSPFv3 packets.
    Authentication Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_ShamLinks_ShamLink_Authentication

    // Encrypt and authenticate OSPFv3 packets.
    Encryption Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_ShamLinks_ShamLink_Encryption
}

func (shamLink *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_ShamLinks_ShamLink) GetEntityData() *types.CommonEntityData {
    shamLink.EntityData.YFilter = shamLink.YFilter
    shamLink.EntityData.YangName = "sham-link"
    shamLink.EntityData.BundleName = "cisco_ios_xr"
    shamLink.EntityData.ParentYangName = "sham-links"
    shamLink.EntityData.SegmentPath = "sham-link" + types.AddKeyToken(shamLink.SourceAddress, "source-address") + types.AddKeyToken(shamLink.DestinationAddress, "destination-address")
    shamLink.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/sham-links/" + shamLink.EntityData.SegmentPath
    shamLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shamLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shamLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shamLink.EntityData.Children = types.NewOrderedMap()
    shamLink.EntityData.Children.Append("authentication", types.YChild{"Authentication", &shamLink.Authentication})
    shamLink.EntityData.Children.Append("encryption", types.YChild{"Encryption", &shamLink.Encryption})
    shamLink.EntityData.Leafs = types.NewOrderedMap()
    shamLink.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", shamLink.SourceAddress})
    shamLink.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", shamLink.DestinationAddress})
    shamLink.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", shamLink.Enable})
    shamLink.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", shamLink.HelloInterval})
    shamLink.EntityData.Leafs.Append("dead-interval", types.YLeaf{"DeadInterval", shamLink.DeadInterval})
    shamLink.EntityData.Leafs.Append("retransmit-interval", types.YLeaf{"RetransmitInterval", shamLink.RetransmitInterval})
    shamLink.EntityData.Leafs.Append("transmit-delay", types.YLeaf{"TransmitDelay", shamLink.TransmitDelay})

    shamLink.EntityData.YListKeys = []string {"SourceAddress", "DestinationAddress"}

    return &(shamLink.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_ShamLinks_ShamLink_Authentication
// Authenticate OSPFv3 packets
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_ShamLinks_ShamLink_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec AH authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Use the MD5 or SHA1 algorithm. The type is Ospfv3Authentication.
    Algorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    Password interface{}
}

func (authentication *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_ShamLinks_ShamLink_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "sham-link"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/sham-links/sham-link/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", authentication.Enable})
    authentication.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", authentication.Spi})
    authentication.EntityData.Leafs.Append("algorithm", types.YLeaf{"Algorithm", authentication.Algorithm})
    authentication.EntityData.Leafs.Append("password", types.YLeaf{"Password", authentication.Password})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_ShamLinks_ShamLink_Encryption
// Encrypt and authenticate OSPFv3 packets
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_ShamLinks_ShamLink_Encryption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec ESP authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Specify the encryption algorithm. The type is Ospfv3EncryptionAlgorithm.
    EncryptionAlgorithm interface{}

    // Encryption password. The type is string with pattern: b'(!.+)|([^!].+)'.
    EncryptionPassword interface{}

    // Use the NULL, MD5 or SHA1 algorithm. The type is Ospfv3AuthenticationType2.
    AuthenticationAlgorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    AuthenticationPassword interface{}
}

func (encryption *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_ShamLinks_ShamLink_Encryption) GetEntityData() *types.CommonEntityData {
    encryption.EntityData.YFilter = encryption.YFilter
    encryption.EntityData.YangName = "encryption"
    encryption.EntityData.BundleName = "cisco_ios_xr"
    encryption.EntityData.ParentYangName = "sham-link"
    encryption.EntityData.SegmentPath = "encryption"
    encryption.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/sham-links/sham-link/" + encryption.EntityData.SegmentPath
    encryption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryption.EntityData.Children = types.NewOrderedMap()
    encryption.EntityData.Leafs = types.NewOrderedMap()
    encryption.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", encryption.Enable})
    encryption.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", encryption.Spi})
    encryption.EntityData.Leafs.Append("encryption-algorithm", types.YLeaf{"EncryptionAlgorithm", encryption.EncryptionAlgorithm})
    encryption.EntityData.Leafs.Append("encryption-password", types.YLeaf{"EncryptionPassword", encryption.EncryptionPassword})
    encryption.EntityData.Leafs.Append("authentication-algorithm", types.YLeaf{"AuthenticationAlgorithm", encryption.AuthenticationAlgorithm})
    encryption.EntityData.Leafs.Append("authentication-password", types.YLeaf{"AuthenticationPassword", encryption.AuthenticationPassword})

    encryption.EntityData.YListKeys = []string {}

    return &(encryption.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_VirtualLinks
// Virtual link sub-mode
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_VirtualLinks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router ID of virtual link neighbor. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_VirtualLinks_VirtualLink.
    VirtualLink []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_VirtualLinks_VirtualLink
}

func (virtualLinks *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_VirtualLinks) GetEntityData() *types.CommonEntityData {
    virtualLinks.EntityData.YFilter = virtualLinks.YFilter
    virtualLinks.EntityData.YangName = "virtual-links"
    virtualLinks.EntityData.BundleName = "cisco_ios_xr"
    virtualLinks.EntityData.ParentYangName = "area-address"
    virtualLinks.EntityData.SegmentPath = "virtual-links"
    virtualLinks.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/" + virtualLinks.EntityData.SegmentPath
    virtualLinks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    virtualLinks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    virtualLinks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    virtualLinks.EntityData.Children = types.NewOrderedMap()
    virtualLinks.EntityData.Children.Append("virtual-link", types.YChild{"VirtualLink", nil})
    for i := range virtualLinks.VirtualLink {
        virtualLinks.EntityData.Children.Append(types.GetSegmentPath(virtualLinks.VirtualLink[i]), types.YChild{"VirtualLink", virtualLinks.VirtualLink[i]})
    }
    virtualLinks.EntityData.Leafs = types.NewOrderedMap()

    virtualLinks.EntityData.YListKeys = []string {}

    return &(virtualLinks.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_VirtualLinks_VirtualLink
// Router ID of virtual link neighbor
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_VirtualLinks_VirtualLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Router ID of virtual link neighbor. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    VirtualLinkAddress interface{}

    // Enabled virtual link. The type is interface{}.
    Enable interface{}

    // Time between HELLO packets. The type is interface{} with range: 1..65535.
    // Units are second.
    HelloInterval interface{}

    // Interval after which a neighbor is declared dead (in seconds). The type is
    // interface{} with range: 1..65535. Units are second.
    DeadInterval interface{}

    // Specify the transmit interval in seconds. The type is interface{} with
    // range: 1..65535. Units are second.
    RetransmitInterval interface{}

    // Specify the transmit delay in seconds. The type is interface{} with range:
    // 1..65535. Units are second.
    TransmitDelay interface{}

    // Authenticate OSPFv3 packets.
    Authentication Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_VirtualLinks_VirtualLink_Authentication

    // Encrypt and authenticate OSPFv3 packets.
    Encryption Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_VirtualLinks_VirtualLink_Encryption
}

func (virtualLink *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_VirtualLinks_VirtualLink) GetEntityData() *types.CommonEntityData {
    virtualLink.EntityData.YFilter = virtualLink.YFilter
    virtualLink.EntityData.YangName = "virtual-link"
    virtualLink.EntityData.BundleName = "cisco_ios_xr"
    virtualLink.EntityData.ParentYangName = "virtual-links"
    virtualLink.EntityData.SegmentPath = "virtual-link" + types.AddKeyToken(virtualLink.VirtualLinkAddress, "virtual-link-address")
    virtualLink.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/virtual-links/" + virtualLink.EntityData.SegmentPath
    virtualLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    virtualLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    virtualLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    virtualLink.EntityData.Children = types.NewOrderedMap()
    virtualLink.EntityData.Children.Append("authentication", types.YChild{"Authentication", &virtualLink.Authentication})
    virtualLink.EntityData.Children.Append("encryption", types.YChild{"Encryption", &virtualLink.Encryption})
    virtualLink.EntityData.Leafs = types.NewOrderedMap()
    virtualLink.EntityData.Leafs.Append("virtual-link-address", types.YLeaf{"VirtualLinkAddress", virtualLink.VirtualLinkAddress})
    virtualLink.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", virtualLink.Enable})
    virtualLink.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", virtualLink.HelloInterval})
    virtualLink.EntityData.Leafs.Append("dead-interval", types.YLeaf{"DeadInterval", virtualLink.DeadInterval})
    virtualLink.EntityData.Leafs.Append("retransmit-interval", types.YLeaf{"RetransmitInterval", virtualLink.RetransmitInterval})
    virtualLink.EntityData.Leafs.Append("transmit-delay", types.YLeaf{"TransmitDelay", virtualLink.TransmitDelay})

    virtualLink.EntityData.YListKeys = []string {"VirtualLinkAddress"}

    return &(virtualLink.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_VirtualLinks_VirtualLink_Authentication
// Authenticate OSPFv3 packets
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_VirtualLinks_VirtualLink_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec AH authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Use the MD5 or SHA1 algorithm. The type is Ospfv3Authentication.
    Algorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    Password interface{}
}

func (authentication *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_VirtualLinks_VirtualLink_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "virtual-link"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/virtual-links/virtual-link/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", authentication.Enable})
    authentication.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", authentication.Spi})
    authentication.EntityData.Leafs.Append("algorithm", types.YLeaf{"Algorithm", authentication.Algorithm})
    authentication.EntityData.Leafs.Append("password", types.YLeaf{"Password", authentication.Password})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_VirtualLinks_VirtualLink_Encryption
// Encrypt and authenticate OSPFv3 packets
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_VirtualLinks_VirtualLink_Encryption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec ESP authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Specify the encryption algorithm. The type is Ospfv3EncryptionAlgorithm.
    EncryptionAlgorithm interface{}

    // Encryption password. The type is string with pattern: b'(!.+)|([^!].+)'.
    EncryptionPassword interface{}

    // Use the NULL, MD5 or SHA1 algorithm. The type is Ospfv3AuthenticationType2.
    AuthenticationAlgorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    AuthenticationPassword interface{}
}

func (encryption *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAddress_VirtualLinks_VirtualLink_Encryption) GetEntityData() *types.CommonEntityData {
    encryption.EntityData.YFilter = encryption.YFilter
    encryption.EntityData.YangName = "encryption"
    encryption.EntityData.BundleName = "cisco_ios_xr"
    encryption.EntityData.ParentYangName = "virtual-link"
    encryption.EntityData.SegmentPath = "encryption"
    encryption.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-address/virtual-links/virtual-link/" + encryption.EntityData.SegmentPath
    encryption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryption.EntityData.Children = types.NewOrderedMap()
    encryption.EntityData.Leafs = types.NewOrderedMap()
    encryption.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", encryption.Enable})
    encryption.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", encryption.Spi})
    encryption.EntityData.Leafs.Append("encryption-algorithm", types.YLeaf{"EncryptionAlgorithm", encryption.EncryptionAlgorithm})
    encryption.EntityData.Leafs.Append("encryption-password", types.YLeaf{"EncryptionPassword", encryption.EncryptionPassword})
    encryption.EntityData.Leafs.Append("authentication-algorithm", types.YLeaf{"AuthenticationAlgorithm", encryption.AuthenticationAlgorithm})
    encryption.EntityData.Leafs.Append("authentication-password", types.YLeaf{"AuthenticationPassword", encryption.AuthenticationPassword})

    encryption.EntityData.YListKeys = []string {}

    return &(encryption.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId
// Configuration for a particular area
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Area ID if in integer format. The type is
    // interface{} with range: 0..4294967295.
    AreaId interface{}

    // Specify area as a stub area.  Allowed only in non-backbone areas. The type
    // is bool.
    Stub interface{}

    // Limit size of OSPFv3 packets. The type is interface{} with range:
    // 256..10000.
    PacketSize interface{}

    // Instance ID. The type is interface{} with range: 0..255.
    Instance interface{}

    // Enable/disable demand circuit operation. The type is bool.
    DemandCircuit interface{}

    // Specify router priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // Translate Type 7 to Type 5, even if not elected NSSA translator. The type
    // is bool.
    Type7TranslateAlways interface{}

    // Enable/disable prefix suppression on an interface. The type is bool.
    PrefixSuppression interface{}

    // Enable OSPFv3 area. The type is interface{}.
    Enable interface{}

    // Enable/disable ignoring of MTU in DBD packets. The type is bool.
    MtuIgnore interface{}

    // Enable/disable routing updates on an interface. The type is bool.
    Passive interface{}

    // Time between HELLO packets. The type is interface{} with range: 1..65535.
    // Units are second.
    HelloInterval interface{}

    // Interval after which a neighbor is declared dead (in seconds). The type is
    // interface{} with range: 1..65535. Units are second.
    DeadInterval interface{}

    // Set the summary default-cost of a NSSA/stub area. The type is interface{}
    // with range: 0..16777215.
    DefaultCost interface{}

    // Enable/disable flood reduction. The type is bool.
    FloodReduction interface{}

    // Specify the transmit interval in seconds. The type is interface{} with
    // range: 1..65535. Units are second.
    RetransmitInterval interface{}

    // Enable/Disable MPLS LDP sync. The type is bool.
    LdpSync interface{}

    // Specify network type. The type is Ospfv3Network.
    Network interface{}

    // Specify the transmit delay in seconds. The type is interface{} with range:
    // 1..65535. Units are second.
    TransmitDelay interface{}

    // Interface cost. The type is interface{} with range: 1..65535.
    Cost interface{}

    // Authenticate OSPFv3 packets.
    Authentication Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Authentication

    // Configure BFD parameters.
    Bfd Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Bfd

    // Range configuration.
    Ranges Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Ranges

    // Encrypt and authenticate OSPFv3 packets.
    Encryption Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Encryption

    // Specify area as a NSSA area.  Allowed only in non-backbone areas.
    Nssa Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Nssa

    // Database filter.
    DatabaseFilter Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_DatabaseFilter

    // Filter prefixes to/from RIB.
    DistributeList Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_DistributeList

    // OSPFv3 interfaces.
    Interfaces Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces

    // Area Scope Configuration.
    AreaScope Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope

    // Sham Link sub-mode.
    ShamLinks Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_ShamLinks

    // Virtual link sub-mode.
    VirtualLinks Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_VirtualLinks
}

func (areaAreaId *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId) GetEntityData() *types.CommonEntityData {
    areaAreaId.EntityData.YFilter = areaAreaId.YFilter
    areaAreaId.EntityData.YangName = "area-area-id"
    areaAreaId.EntityData.BundleName = "cisco_ios_xr"
    areaAreaId.EntityData.ParentYangName = "area-addresses"
    areaAreaId.EntityData.SegmentPath = "area-area-id" + types.AddKeyToken(areaAreaId.AreaId, "area-id")
    areaAreaId.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/" + areaAreaId.EntityData.SegmentPath
    areaAreaId.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    areaAreaId.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    areaAreaId.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    areaAreaId.EntityData.Children = types.NewOrderedMap()
    areaAreaId.EntityData.Children.Append("authentication", types.YChild{"Authentication", &areaAreaId.Authentication})
    areaAreaId.EntityData.Children.Append("bfd", types.YChild{"Bfd", &areaAreaId.Bfd})
    areaAreaId.EntityData.Children.Append("ranges", types.YChild{"Ranges", &areaAreaId.Ranges})
    areaAreaId.EntityData.Children.Append("encryption", types.YChild{"Encryption", &areaAreaId.Encryption})
    areaAreaId.EntityData.Children.Append("nssa", types.YChild{"Nssa", &areaAreaId.Nssa})
    areaAreaId.EntityData.Children.Append("database-filter", types.YChild{"DatabaseFilter", &areaAreaId.DatabaseFilter})
    areaAreaId.EntityData.Children.Append("distribute-list", types.YChild{"DistributeList", &areaAreaId.DistributeList})
    areaAreaId.EntityData.Children.Append("interfaces", types.YChild{"Interfaces", &areaAreaId.Interfaces})
    areaAreaId.EntityData.Children.Append("area-scope", types.YChild{"AreaScope", &areaAreaId.AreaScope})
    areaAreaId.EntityData.Children.Append("sham-links", types.YChild{"ShamLinks", &areaAreaId.ShamLinks})
    areaAreaId.EntityData.Children.Append("virtual-links", types.YChild{"VirtualLinks", &areaAreaId.VirtualLinks})
    areaAreaId.EntityData.Leafs = types.NewOrderedMap()
    areaAreaId.EntityData.Leafs.Append("area-id", types.YLeaf{"AreaId", areaAreaId.AreaId})
    areaAreaId.EntityData.Leafs.Append("stub", types.YLeaf{"Stub", areaAreaId.Stub})
    areaAreaId.EntityData.Leafs.Append("packet-size", types.YLeaf{"PacketSize", areaAreaId.PacketSize})
    areaAreaId.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", areaAreaId.Instance})
    areaAreaId.EntityData.Leafs.Append("demand-circuit", types.YLeaf{"DemandCircuit", areaAreaId.DemandCircuit})
    areaAreaId.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", areaAreaId.Priority})
    areaAreaId.EntityData.Leafs.Append("type7-translate-always", types.YLeaf{"Type7TranslateAlways", areaAreaId.Type7TranslateAlways})
    areaAreaId.EntityData.Leafs.Append("prefix-suppression", types.YLeaf{"PrefixSuppression", areaAreaId.PrefixSuppression})
    areaAreaId.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", areaAreaId.Enable})
    areaAreaId.EntityData.Leafs.Append("mtu-ignore", types.YLeaf{"MtuIgnore", areaAreaId.MtuIgnore})
    areaAreaId.EntityData.Leafs.Append("passive", types.YLeaf{"Passive", areaAreaId.Passive})
    areaAreaId.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", areaAreaId.HelloInterval})
    areaAreaId.EntityData.Leafs.Append("dead-interval", types.YLeaf{"DeadInterval", areaAreaId.DeadInterval})
    areaAreaId.EntityData.Leafs.Append("default-cost", types.YLeaf{"DefaultCost", areaAreaId.DefaultCost})
    areaAreaId.EntityData.Leafs.Append("flood-reduction", types.YLeaf{"FloodReduction", areaAreaId.FloodReduction})
    areaAreaId.EntityData.Leafs.Append("retransmit-interval", types.YLeaf{"RetransmitInterval", areaAreaId.RetransmitInterval})
    areaAreaId.EntityData.Leafs.Append("ldp-sync", types.YLeaf{"LdpSync", areaAreaId.LdpSync})
    areaAreaId.EntityData.Leafs.Append("network", types.YLeaf{"Network", areaAreaId.Network})
    areaAreaId.EntityData.Leafs.Append("transmit-delay", types.YLeaf{"TransmitDelay", areaAreaId.TransmitDelay})
    areaAreaId.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", areaAreaId.Cost})

    areaAreaId.EntityData.YListKeys = []string {"AreaId"}

    return &(areaAreaId.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Authentication
// Authenticate OSPFv3 packets
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec AH authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Use the MD5 or SHA1 algorithm. The type is Ospfv3Authentication.
    Algorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    Password interface{}
}

func (authentication *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "area-area-id"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", authentication.Enable})
    authentication.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", authentication.Spi})
    authentication.EntityData.Leafs.Append("algorithm", types.YLeaf{"Algorithm", authentication.Algorithm})
    authentication.EntityData.Leafs.Append("password", types.YLeaf{"Password", authentication.Password})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Bfd
// Configure BFD parameters
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Detect multiplier. The type is interface{} with range: 2..50.
    DetectionMultiplier interface{}

    // Hello interval in milli-seconds. The type is interface{} with range:
    // 3..30000. Units are millisecond.
    Interval interface{}

    // Enable or disable BFD fast detection. The type is Ospfv3bfdEnableMode.
    FastDetectMode interface{}
}

func (bfd *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "area-area-id"
    bfd.EntityData.SegmentPath = "bfd"
    bfd.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/" + bfd.EntityData.SegmentPath
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = types.NewOrderedMap()
    bfd.EntityData.Leafs = types.NewOrderedMap()
    bfd.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", bfd.DetectionMultiplier})
    bfd.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", bfd.Interval})
    bfd.EntityData.Leafs.Append("fast-detect-mode", types.YLeaf{"FastDetectMode", bfd.FastDetectMode})

    bfd.EntityData.YListKeys = []string {}

    return &(bfd.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Ranges
// Range configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Ranges struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Summarize inter-area routes matching prefix/length. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Ranges_Range.
    Range []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Ranges_Range
}

func (ranges *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Ranges) GetEntityData() *types.CommonEntityData {
    ranges.EntityData.YFilter = ranges.YFilter
    ranges.EntityData.YangName = "ranges"
    ranges.EntityData.BundleName = "cisco_ios_xr"
    ranges.EntityData.ParentYangName = "area-area-id"
    ranges.EntityData.SegmentPath = "ranges"
    ranges.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/" + ranges.EntityData.SegmentPath
    ranges.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ranges.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ranges.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ranges.EntityData.Children = types.NewOrderedMap()
    ranges.EntityData.Children.Append("range", types.YChild{"Range", nil})
    for i := range ranges.Range {
        ranges.EntityData.Children.Append(types.GetSegmentPath(ranges.Range[i]), types.YChild{"Range", ranges.Range[i]})
    }
    ranges.EntityData.Leafs = types.NewOrderedMap()

    ranges.EntityData.YListKeys = []string {}

    return &(ranges.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Ranges_Range
// Summarize inter-area routes matching
// prefix/length
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Ranges_Range struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IPv6 prefix format. The type is string with
    // pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Prefix interface{}

    // This attribute is a key. IPV6 prefix length. The type is interface{} with
    // range: 0..128.
    PrefixLength interface{}

    // Do not advertise address range. The type is bool. The default value is
    // false.
    NotAdvertise interface{}

    // Specified metric for this range. The type is interface{} with range:
    // 1..16777214.
    Cost interface{}
}

func (self *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Ranges_Range) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "range"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "ranges"
    self.EntityData.SegmentPath = "range" + types.AddKeyToken(self.Prefix, "prefix") + types.AddKeyToken(self.PrefixLength, "prefix-length")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/ranges/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", self.Prefix})
    self.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", self.PrefixLength})
    self.EntityData.Leafs.Append("not-advertise", types.YLeaf{"NotAdvertise", self.NotAdvertise})
    self.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", self.Cost})

    self.EntityData.YListKeys = []string {"Prefix", "PrefixLength"}

    return &(self.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Encryption
// Encrypt and authenticate OSPFv3 packets
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Encryption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec ESP authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Specify the encryption algorithm. The type is Ospfv3EncryptionAlgorithm.
    EncryptionAlgorithm interface{}

    // Encryption password. The type is string with pattern: b'(!.+)|([^!].+)'.
    EncryptionPassword interface{}

    // Use the NULL, MD5 or SHA1 algorithm. The type is Ospfv3AuthenticationType2.
    AuthenticationAlgorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    AuthenticationPassword interface{}
}

func (encryption *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Encryption) GetEntityData() *types.CommonEntityData {
    encryption.EntityData.YFilter = encryption.YFilter
    encryption.EntityData.YangName = "encryption"
    encryption.EntityData.BundleName = "cisco_ios_xr"
    encryption.EntityData.ParentYangName = "area-area-id"
    encryption.EntityData.SegmentPath = "encryption"
    encryption.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/" + encryption.EntityData.SegmentPath
    encryption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryption.EntityData.Children = types.NewOrderedMap()
    encryption.EntityData.Leafs = types.NewOrderedMap()
    encryption.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", encryption.Enable})
    encryption.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", encryption.Spi})
    encryption.EntityData.Leafs.Append("encryption-algorithm", types.YLeaf{"EncryptionAlgorithm", encryption.EncryptionAlgorithm})
    encryption.EntityData.Leafs.Append("encryption-password", types.YLeaf{"EncryptionPassword", encryption.EncryptionPassword})
    encryption.EntityData.Leafs.Append("authentication-algorithm", types.YLeaf{"AuthenticationAlgorithm", encryption.AuthenticationAlgorithm})
    encryption.EntityData.Leafs.Append("authentication-password", types.YLeaf{"AuthenticationPassword", encryption.AuthenticationPassword})

    encryption.EntityData.YListKeys = []string {}

    return &(encryption.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Nssa
// Specify area as a NSSA area.  Allowed only in
// non-backbone areas
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Nssa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // No redistribution into this NSSA area. The type is bool. The default value
    // is false.
    NoRedistribution interface{}

    // Originate Type 7 default into NSSA area. The type is bool. The default
    // value is false.
    DefaultInfoOriginate interface{}

    // Only valid with DefaultInfoOriginate. The type is interface{} with range:
    // 0..16777214.
    Metric interface{}

    // Only valid with DefaultInfoOriginate. The type is Ospfv3Metric.
    MetricType interface{}

    // Do not send summary LSA into NSSA. The type is interface{}.
    NoSummary interface{}
}

func (nssa *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Nssa) GetEntityData() *types.CommonEntityData {
    nssa.EntityData.YFilter = nssa.YFilter
    nssa.EntityData.YangName = "nssa"
    nssa.EntityData.BundleName = "cisco_ios_xr"
    nssa.EntityData.ParentYangName = "area-area-id"
    nssa.EntityData.SegmentPath = "nssa"
    nssa.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/" + nssa.EntityData.SegmentPath
    nssa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    nssa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    nssa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    nssa.EntityData.Children = types.NewOrderedMap()
    nssa.EntityData.Leafs = types.NewOrderedMap()
    nssa.EntityData.Leafs.Append("no-redistribution", types.YLeaf{"NoRedistribution", nssa.NoRedistribution})
    nssa.EntityData.Leafs.Append("default-info-originate", types.YLeaf{"DefaultInfoOriginate", nssa.DefaultInfoOriginate})
    nssa.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", nssa.Metric})
    nssa.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", nssa.MetricType})
    nssa.EntityData.Leafs.Append("no-summary", types.YLeaf{"NoSummary", nssa.NoSummary})

    nssa.EntityData.YListKeys = []string {}

    return &(nssa.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_DatabaseFilter
// Database filter
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_DatabaseFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All.
    All Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_DatabaseFilter_All
}

func (databaseFilter *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_DatabaseFilter) GetEntityData() *types.CommonEntityData {
    databaseFilter.EntityData.YFilter = databaseFilter.YFilter
    databaseFilter.EntityData.YangName = "database-filter"
    databaseFilter.EntityData.BundleName = "cisco_ios_xr"
    databaseFilter.EntityData.ParentYangName = "area-area-id"
    databaseFilter.EntityData.SegmentPath = "database-filter"
    databaseFilter.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/" + databaseFilter.EntityData.SegmentPath
    databaseFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    databaseFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    databaseFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    databaseFilter.EntityData.Children = types.NewOrderedMap()
    databaseFilter.EntityData.Children.Append("all", types.YChild{"All", &databaseFilter.All})
    databaseFilter.EntityData.Leafs = types.NewOrderedMap()

    databaseFilter.EntityData.YListKeys = []string {}

    return &(databaseFilter.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_DatabaseFilter_All
// All
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_DatabaseFilter_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable or disable database-filter. The type is bool.
    Out interface{}
}

func (all *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_DatabaseFilter_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "database-filter"
    all.EntityData.SegmentPath = "all"
    all.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/database-filter/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Leafs = types.NewOrderedMap()
    all.EntityData.Leafs.Append("out", types.YLeaf{"Out", all.Out})

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_DistributeList
// Filter prefixes to/from RIB
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_DistributeList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter prefixes installed to RIB.
    In Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_DistributeList_In
}

func (distributeList *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_DistributeList) GetEntityData() *types.CommonEntityData {
    distributeList.EntityData.YFilter = distributeList.YFilter
    distributeList.EntityData.YangName = "distribute-list"
    distributeList.EntityData.BundleName = "cisco_ios_xr"
    distributeList.EntityData.ParentYangName = "area-area-id"
    distributeList.EntityData.SegmentPath = "distribute-list"
    distributeList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/" + distributeList.EntityData.SegmentPath
    distributeList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    distributeList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    distributeList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    distributeList.EntityData.Children = types.NewOrderedMap()
    distributeList.EntityData.Children.Append("in", types.YChild{"In", &distributeList.In})
    distributeList.EntityData.Leafs = types.NewOrderedMap()

    distributeList.EntityData.YListKeys = []string {}

    return &(distributeList.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_DistributeList_In
// Filter prefixes installed to RIB
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_DistributeList_In struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter prefixes based on an IPv6 prefix-list. The type is string.
    PrefixList interface{}
}

func (in *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_DistributeList_In) GetEntityData() *types.CommonEntityData {
    in.EntityData.YFilter = in.YFilter
    in.EntityData.YangName = "in"
    in.EntityData.BundleName = "cisco_ios_xr"
    in.EntityData.ParentYangName = "distribute-list"
    in.EntityData.SegmentPath = "in"
    in.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/distribute-list/" + in.EntityData.SegmentPath
    in.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    in.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    in.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    in.EntityData.Children = types.NewOrderedMap()
    in.EntityData.Leafs = types.NewOrderedMap()
    in.EntityData.Leafs.Append("prefix-list", types.YLeaf{"PrefixList", in.PrefixList})

    in.EntityData.YListKeys = []string {}

    return &(in.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces
// OSPFv3 interfaces
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // OSPFv3 interface. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface.
    Interface []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface
}

func (interfaces *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces) GetEntityData() *types.CommonEntityData {
    interfaces.EntityData.YFilter = interfaces.YFilter
    interfaces.EntityData.YangName = "interfaces"
    interfaces.EntityData.BundleName = "cisco_ios_xr"
    interfaces.EntityData.ParentYangName = "area-area-id"
    interfaces.EntityData.SegmentPath = "interfaces"
    interfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/" + interfaces.EntityData.SegmentPath
    interfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interfaces.EntityData.Children = types.NewOrderedMap()
    interfaces.EntityData.Children.Append("interface", types.YChild{"Interface", nil})
    for i := range interfaces.Interface {
        interfaces.EntityData.Children.Append(types.GetSegmentPath(interfaces.Interface[i]), types.YChild{"Interface", interfaces.Interface[i]})
    }
    interfaces.EntityData.Leafs = types.NewOrderedMap()

    interfaces.EntityData.YListKeys = []string {}

    return &(interfaces.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface
// OSPFv3 interface
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface to configure. The type is string with
    // pattern: b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}

    // Enable OSPFv3 interface. The type is interface{}.
    Enable interface{}

    // Interval after which a neighbor is declared dead (in seconds). The type is
    // interface{} with range: 1..65535. Units are second.
    DeadInterval interface{}

    // Enable/disable flood reduction. The type is bool.
    FloodReduction interface{}

    // Interface cost. The type is interface{} with range: 1..65535.
    Cost interface{}

    // Specify the transmit delay in seconds. The type is interface{} with range:
    // 1..65535. Units are second.
    TransmitDelay interface{}

    // Instance ID. The type is interface{} with range: 0..255.
    Instance interface{}

    // Enable/Disable MPLS LDP sync. The type is bool.
    LdpSync interface{}

    // Enable/disable ignoring of MTU in DBD packets. The type is bool.
    MtuIgnore interface{}

    // Specify the transmit interval in seconds. The type is interface{} with
    // range: 1..65535. Units are second.
    RetransmitInterval interface{}

    // Time between HELLO packets. The type is interface{} with range: 1..65535.
    // Units are second.
    HelloInterval interface{}

    // Enable/disable routing updates on an interface. The type is bool.
    Passive interface{}

    // Limit size of OSPFv3 packets. The type is interface{} with range:
    // 256..10000.
    PacketSize interface{}

    // Enable/disable prefix suppression on an interface. The type is bool.
    PrefixSuppression interface{}

    // Specify router priority. The type is interface{} with range: 0..255.
    Priority interface{}

    // Specify network type. The type is Ospfv3Network.
    Network interface{}

    // Enable/disable demand circuit operation. The type is bool.
    DemandCircuit interface{}

    // Authenticate OSPFv3 packets.
    Authentication Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Authentication

    // Specify a neighbor router.
    Neighbors Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Neighbors

    // Encrypt and authenticate OSPFv3 packets.
    Encryption Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Encryption

    // Configure BFD parameters.
    Bfd Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Bfd

    // Database filter.
    DatabaseFilter Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DatabaseFilter

    // Filter prefixes to/from RIB.
    DistributeList Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DistributeList

    // Fast-reroute configuration.
    FastReroute Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute
}

func (self *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface) GetEntityData() *types.CommonEntityData {
    self.EntityData.YFilter = self.YFilter
    self.EntityData.YangName = "interface"
    self.EntityData.BundleName = "cisco_ios_xr"
    self.EntityData.ParentYangName = "interfaces"
    self.EntityData.SegmentPath = "interface" + types.AddKeyToken(self.InterfaceName, "interface-name")
    self.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/interfaces/" + self.EntityData.SegmentPath
    self.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    self.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    self.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    self.EntityData.Children = types.NewOrderedMap()
    self.EntityData.Children.Append("authentication", types.YChild{"Authentication", &self.Authentication})
    self.EntityData.Children.Append("neighbors", types.YChild{"Neighbors", &self.Neighbors})
    self.EntityData.Children.Append("encryption", types.YChild{"Encryption", &self.Encryption})
    self.EntityData.Children.Append("bfd", types.YChild{"Bfd", &self.Bfd})
    self.EntityData.Children.Append("database-filter", types.YChild{"DatabaseFilter", &self.DatabaseFilter})
    self.EntityData.Children.Append("distribute-list", types.YChild{"DistributeList", &self.DistributeList})
    self.EntityData.Children.Append("fast-reroute", types.YChild{"FastReroute", &self.FastReroute})
    self.EntityData.Leafs = types.NewOrderedMap()
    self.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", self.InterfaceName})
    self.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", self.Enable})
    self.EntityData.Leafs.Append("dead-interval", types.YLeaf{"DeadInterval", self.DeadInterval})
    self.EntityData.Leafs.Append("flood-reduction", types.YLeaf{"FloodReduction", self.FloodReduction})
    self.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", self.Cost})
    self.EntityData.Leafs.Append("transmit-delay", types.YLeaf{"TransmitDelay", self.TransmitDelay})
    self.EntityData.Leafs.Append("instance", types.YLeaf{"Instance", self.Instance})
    self.EntityData.Leafs.Append("ldp-sync", types.YLeaf{"LdpSync", self.LdpSync})
    self.EntityData.Leafs.Append("mtu-ignore", types.YLeaf{"MtuIgnore", self.MtuIgnore})
    self.EntityData.Leafs.Append("retransmit-interval", types.YLeaf{"RetransmitInterval", self.RetransmitInterval})
    self.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", self.HelloInterval})
    self.EntityData.Leafs.Append("passive", types.YLeaf{"Passive", self.Passive})
    self.EntityData.Leafs.Append("packet-size", types.YLeaf{"PacketSize", self.PacketSize})
    self.EntityData.Leafs.Append("prefix-suppression", types.YLeaf{"PrefixSuppression", self.PrefixSuppression})
    self.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", self.Priority})
    self.EntityData.Leafs.Append("network", types.YLeaf{"Network", self.Network})
    self.EntityData.Leafs.Append("demand-circuit", types.YLeaf{"DemandCircuit", self.DemandCircuit})

    self.EntityData.YListKeys = []string {"InterfaceName"}

    return &(self.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Authentication
// Authenticate OSPFv3 packets
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec AH authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Use the MD5 or SHA1 algorithm. The type is Ospfv3Authentication.
    Algorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    Password interface{}
}

func (authentication *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "interface"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/interfaces/interface/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", authentication.Enable})
    authentication.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", authentication.Spi})
    authentication.EntityData.Leafs.Append("algorithm", types.YLeaf{"Algorithm", authentication.Algorithm})
    authentication.EntityData.Leafs.Append("password", types.YLeaf{"Password", authentication.Password})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Neighbors
// Specify a neighbor router
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Neighbors struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 address. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Neighbors_Neighbor.
    Neighbor []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Neighbors_Neighbor
}

func (neighbors *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Neighbors) GetEntityData() *types.CommonEntityData {
    neighbors.EntityData.YFilter = neighbors.YFilter
    neighbors.EntityData.YangName = "neighbors"
    neighbors.EntityData.BundleName = "cisco_ios_xr"
    neighbors.EntityData.ParentYangName = "interface"
    neighbors.EntityData.SegmentPath = "neighbors"
    neighbors.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/interfaces/interface/" + neighbors.EntityData.SegmentPath
    neighbors.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbors.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbors.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbors.EntityData.Children = types.NewOrderedMap()
    neighbors.EntityData.Children.Append("neighbor", types.YChild{"Neighbor", nil})
    for i := range neighbors.Neighbor {
        neighbors.EntityData.Children.Append(types.GetSegmentPath(neighbors.Neighbor[i]), types.YChild{"Neighbor", neighbors.Neighbor[i]})
    }
    neighbors.EntityData.Leafs = types.NewOrderedMap()

    neighbors.EntityData.YListKeys = []string {}

    return &(neighbors.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Neighbors_Neighbor
// IPv6 address
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Neighbors_Neighbor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IPV6 address. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    NeighborAddress interface{}

    // OSPFv3 priority of non-broadcast neighbor. The type is interface{} with
    // range: 0..255.
    Priority interface{}

    // OSPFv3 dead-router polling interval (in seconds). The type is interface{}
    // with range: 0..65535. Units are second.
    PollInterval interface{}

    // OSPFv3 cost for point-to-multipoint neighbor. The type is interface{} with
    // range: 1..65535.
    Cost interface{}

    // Filter OSPFv3 LSA during synchronization and flooding for
    // point-to-multipoint neighbor. The type is bool.
    DatabaseFilter interface{}

    // Zone. The type is string.
    Zone interface{}
}

func (neighbor *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Neighbors_Neighbor) GetEntityData() *types.CommonEntityData {
    neighbor.EntityData.YFilter = neighbor.YFilter
    neighbor.EntityData.YangName = "neighbor"
    neighbor.EntityData.BundleName = "cisco_ios_xr"
    neighbor.EntityData.ParentYangName = "neighbors"
    neighbor.EntityData.SegmentPath = "neighbor" + types.AddKeyToken(neighbor.NeighborAddress, "neighbor-address")
    neighbor.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/interfaces/interface/neighbors/" + neighbor.EntityData.SegmentPath
    neighbor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    neighbor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    neighbor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    neighbor.EntityData.Children = types.NewOrderedMap()
    neighbor.EntityData.Leafs = types.NewOrderedMap()
    neighbor.EntityData.Leafs.Append("neighbor-address", types.YLeaf{"NeighborAddress", neighbor.NeighborAddress})
    neighbor.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", neighbor.Priority})
    neighbor.EntityData.Leafs.Append("poll-interval", types.YLeaf{"PollInterval", neighbor.PollInterval})
    neighbor.EntityData.Leafs.Append("cost", types.YLeaf{"Cost", neighbor.Cost})
    neighbor.EntityData.Leafs.Append("database-filter", types.YLeaf{"DatabaseFilter", neighbor.DatabaseFilter})
    neighbor.EntityData.Leafs.Append("zone", types.YLeaf{"Zone", neighbor.Zone})

    neighbor.EntityData.YListKeys = []string {"NeighborAddress"}

    return &(neighbor.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Encryption
// Encrypt and authenticate OSPFv3 packets
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Encryption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec ESP authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Specify the encryption algorithm. The type is Ospfv3EncryptionAlgorithm.
    EncryptionAlgorithm interface{}

    // Encryption password. The type is string with pattern: b'(!.+)|([^!].+)'.
    EncryptionPassword interface{}

    // Use the NULL, MD5 or SHA1 algorithm. The type is Ospfv3AuthenticationType2.
    AuthenticationAlgorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    AuthenticationPassword interface{}
}

func (encryption *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Encryption) GetEntityData() *types.CommonEntityData {
    encryption.EntityData.YFilter = encryption.YFilter
    encryption.EntityData.YangName = "encryption"
    encryption.EntityData.BundleName = "cisco_ios_xr"
    encryption.EntityData.ParentYangName = "interface"
    encryption.EntityData.SegmentPath = "encryption"
    encryption.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/interfaces/interface/" + encryption.EntityData.SegmentPath
    encryption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryption.EntityData.Children = types.NewOrderedMap()
    encryption.EntityData.Leafs = types.NewOrderedMap()
    encryption.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", encryption.Enable})
    encryption.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", encryption.Spi})
    encryption.EntityData.Leafs.Append("encryption-algorithm", types.YLeaf{"EncryptionAlgorithm", encryption.EncryptionAlgorithm})
    encryption.EntityData.Leafs.Append("encryption-password", types.YLeaf{"EncryptionPassword", encryption.EncryptionPassword})
    encryption.EntityData.Leafs.Append("authentication-algorithm", types.YLeaf{"AuthenticationAlgorithm", encryption.AuthenticationAlgorithm})
    encryption.EntityData.Leafs.Append("authentication-password", types.YLeaf{"AuthenticationPassword", encryption.AuthenticationPassword})

    encryption.EntityData.YListKeys = []string {}

    return &(encryption.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Bfd
// Configure BFD parameters
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Hello interval in milli-seconds. The type is interface{} with range:
    // 3..30000. Units are millisecond.
    Interval interface{}

    // Enable or disable BFD fast detection. The type is Ospfv3bfdEnableMode.
    FastDetectMode interface{}

    // Detect multiplier. The type is interface{} with range: 2..50.
    DetectionMultiplier interface{}
}

func (bfd *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "interface"
    bfd.EntityData.SegmentPath = "bfd"
    bfd.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/interfaces/interface/" + bfd.EntityData.SegmentPath
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = types.NewOrderedMap()
    bfd.EntityData.Leafs = types.NewOrderedMap()
    bfd.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", bfd.Interval})
    bfd.EntityData.Leafs.Append("fast-detect-mode", types.YLeaf{"FastDetectMode", bfd.FastDetectMode})
    bfd.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", bfd.DetectionMultiplier})

    bfd.EntityData.YListKeys = []string {}

    return &(bfd.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DatabaseFilter
// Database filter
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DatabaseFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All.
    All Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DatabaseFilter_All
}

func (databaseFilter *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DatabaseFilter) GetEntityData() *types.CommonEntityData {
    databaseFilter.EntityData.YFilter = databaseFilter.YFilter
    databaseFilter.EntityData.YangName = "database-filter"
    databaseFilter.EntityData.BundleName = "cisco_ios_xr"
    databaseFilter.EntityData.ParentYangName = "interface"
    databaseFilter.EntityData.SegmentPath = "database-filter"
    databaseFilter.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/interfaces/interface/" + databaseFilter.EntityData.SegmentPath
    databaseFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    databaseFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    databaseFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    databaseFilter.EntityData.Children = types.NewOrderedMap()
    databaseFilter.EntityData.Children.Append("all", types.YChild{"All", &databaseFilter.All})
    databaseFilter.EntityData.Leafs = types.NewOrderedMap()

    databaseFilter.EntityData.YListKeys = []string {}

    return &(databaseFilter.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DatabaseFilter_All
// All
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DatabaseFilter_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable or disable database-filter. The type is bool.
    Out interface{}
}

func (all *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DatabaseFilter_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "database-filter"
    all.EntityData.SegmentPath = "all"
    all.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/interfaces/interface/database-filter/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Leafs = types.NewOrderedMap()
    all.EntityData.Leafs.Append("out", types.YLeaf{"Out", all.Out})

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DistributeList
// Filter prefixes to/from RIB
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DistributeList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter prefixes installed to RIB.
    In Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DistributeList_In
}

func (distributeList *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DistributeList) GetEntityData() *types.CommonEntityData {
    distributeList.EntityData.YFilter = distributeList.YFilter
    distributeList.EntityData.YangName = "distribute-list"
    distributeList.EntityData.BundleName = "cisco_ios_xr"
    distributeList.EntityData.ParentYangName = "interface"
    distributeList.EntityData.SegmentPath = "distribute-list"
    distributeList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/interfaces/interface/" + distributeList.EntityData.SegmentPath
    distributeList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    distributeList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    distributeList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    distributeList.EntityData.Children = types.NewOrderedMap()
    distributeList.EntityData.Children.Append("in", types.YChild{"In", &distributeList.In})
    distributeList.EntityData.Leafs = types.NewOrderedMap()

    distributeList.EntityData.YListKeys = []string {}

    return &(distributeList.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DistributeList_In
// Filter prefixes installed to RIB
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DistributeList_In struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter prefixes based on an IPv6 prefix-list. The type is string.
    PrefixList interface{}
}

func (in *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_DistributeList_In) GetEntityData() *types.CommonEntityData {
    in.EntityData.YFilter = in.YFilter
    in.EntityData.YangName = "in"
    in.EntityData.BundleName = "cisco_ios_xr"
    in.EntityData.ParentYangName = "distribute-list"
    in.EntityData.SegmentPath = "in"
    in.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/interfaces/interface/distribute-list/" + in.EntityData.SegmentPath
    in.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    in.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    in.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    in.EntityData.Children = types.NewOrderedMap()
    in.EntityData.Leafs = types.NewOrderedMap()
    in.EntityData.Leafs.Append("prefix-list", types.YLeaf{"PrefixList", in.PrefixList})

    in.EntityData.YListKeys = []string {}

    return &(in.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute
// Fast-reroute configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable Fast-reroute per-link or per-prefix. The type is
    // Ospfv3FastReroute.
    FastRerouteEnable interface{}

    // Fast-reroute per-link configuration.
    PerLink Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink

    // Fast-reroute per-link configuration.
    PerPrefix Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix
}

func (fastReroute *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute) GetEntityData() *types.CommonEntityData {
    fastReroute.EntityData.YFilter = fastReroute.YFilter
    fastReroute.EntityData.YangName = "fast-reroute"
    fastReroute.EntityData.BundleName = "cisco_ios_xr"
    fastReroute.EntityData.ParentYangName = "interface"
    fastReroute.EntityData.SegmentPath = "fast-reroute"
    fastReroute.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/interfaces/interface/" + fastReroute.EntityData.SegmentPath
    fastReroute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fastReroute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fastReroute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fastReroute.EntityData.Children = types.NewOrderedMap()
    fastReroute.EntityData.Children.Append("per-link", types.YChild{"PerLink", &fastReroute.PerLink})
    fastReroute.EntityData.Children.Append("per-prefix", types.YChild{"PerPrefix", &fastReroute.PerPrefix})
    fastReroute.EntityData.Leafs = types.NewOrderedMap()
    fastReroute.EntityData.Leafs.Append("fast-reroute-enable", types.YLeaf{"FastRerouteEnable", fastReroute.FastRerouteEnable})

    fastReroute.EntityData.YListKeys = []string {}

    return &(fastReroute.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink
// Fast-reroute per-link configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Use only interfaces on the candidate list as a backup path. The type is
    // bool. The default value is false.
    FastRerouteUseCandidateOnly interface{}

    // Fast-reroute per-link/per-prefix candidate interface configuration.
    CandidateInterfaces Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces

    // Fast-reroute per-link/per-prefix exclude interface configuration.
    ExcludeInterfaces Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces
}

func (perLink *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink) GetEntityData() *types.CommonEntityData {
    perLink.EntityData.YFilter = perLink.YFilter
    perLink.EntityData.YangName = "per-link"
    perLink.EntityData.BundleName = "cisco_ios_xr"
    perLink.EntityData.ParentYangName = "fast-reroute"
    perLink.EntityData.SegmentPath = "per-link"
    perLink.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/interfaces/interface/fast-reroute/" + perLink.EntityData.SegmentPath
    perLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    perLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    perLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    perLink.EntityData.Children = types.NewOrderedMap()
    perLink.EntityData.Children.Append("candidate-interfaces", types.YChild{"CandidateInterfaces", &perLink.CandidateInterfaces})
    perLink.EntityData.Children.Append("exclude-interfaces", types.YChild{"ExcludeInterfaces", &perLink.ExcludeInterfaces})
    perLink.EntityData.Leafs = types.NewOrderedMap()
    perLink.EntityData.Leafs.Append("fast-reroute-use-candidate-only", types.YLeaf{"FastRerouteUseCandidateOnly", perLink.FastRerouteUseCandidateOnly})

    perLink.EntityData.YListKeys = []string {}

    return &(perLink.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces
// Fast-reroute per-link/per-prefix candidate
// interface configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Candidate backup interface. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces_CandidateInterface.
    CandidateInterface []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces_CandidateInterface
}

func (candidateInterfaces *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces) GetEntityData() *types.CommonEntityData {
    candidateInterfaces.EntityData.YFilter = candidateInterfaces.YFilter
    candidateInterfaces.EntityData.YangName = "candidate-interfaces"
    candidateInterfaces.EntityData.BundleName = "cisco_ios_xr"
    candidateInterfaces.EntityData.ParentYangName = "per-link"
    candidateInterfaces.EntityData.SegmentPath = "candidate-interfaces"
    candidateInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/interfaces/interface/fast-reroute/per-link/" + candidateInterfaces.EntityData.SegmentPath
    candidateInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterfaces.EntityData.Children = types.NewOrderedMap()
    candidateInterfaces.EntityData.Children.Append("candidate-interface", types.YChild{"CandidateInterface", nil})
    for i := range candidateInterfaces.CandidateInterface {
        candidateInterfaces.EntityData.Children.Append(types.GetSegmentPath(candidateInterfaces.CandidateInterface[i]), types.YChild{"CandidateInterface", candidateInterfaces.CandidateInterface[i]})
    }
    candidateInterfaces.EntityData.Leafs = types.NewOrderedMap()

    candidateInterfaces.EntityData.YListKeys = []string {}

    return &(candidateInterfaces.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces_CandidateInterface
// Candidate backup interface
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces_CandidateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (candidateInterface *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_CandidateInterfaces_CandidateInterface) GetEntityData() *types.CommonEntityData {
    candidateInterface.EntityData.YFilter = candidateInterface.YFilter
    candidateInterface.EntityData.YangName = "candidate-interface"
    candidateInterface.EntityData.BundleName = "cisco_ios_xr"
    candidateInterface.EntityData.ParentYangName = "candidate-interfaces"
    candidateInterface.EntityData.SegmentPath = "candidate-interface" + types.AddKeyToken(candidateInterface.InterfaceName, "interface-name")
    candidateInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/interfaces/interface/fast-reroute/per-link/candidate-interfaces/" + candidateInterface.EntityData.SegmentPath
    candidateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterface.EntityData.Children = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", candidateInterface.InterfaceName})

    candidateInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(candidateInterface.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces
// Fast-reroute per-link/per-prefix exclude
// interface configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exclude an interface from becoming a backup. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface.
    ExcludeInterface []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface
}

func (excludeInterfaces *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces) GetEntityData() *types.CommonEntityData {
    excludeInterfaces.EntityData.YFilter = excludeInterfaces.YFilter
    excludeInterfaces.EntityData.YangName = "exclude-interfaces"
    excludeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    excludeInterfaces.EntityData.ParentYangName = "per-link"
    excludeInterfaces.EntityData.SegmentPath = "exclude-interfaces"
    excludeInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/interfaces/interface/fast-reroute/per-link/" + excludeInterfaces.EntityData.SegmentPath
    excludeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterfaces.EntityData.Children = types.NewOrderedMap()
    excludeInterfaces.EntityData.Children.Append("exclude-interface", types.YChild{"ExcludeInterface", nil})
    for i := range excludeInterfaces.ExcludeInterface {
        excludeInterfaces.EntityData.Children.Append(types.GetSegmentPath(excludeInterfaces.ExcludeInterface[i]), types.YChild{"ExcludeInterface", excludeInterfaces.ExcludeInterface[i]})
    }
    excludeInterfaces.EntityData.Leafs = types.NewOrderedMap()

    excludeInterfaces.EntityData.YListKeys = []string {}

    return &(excludeInterfaces.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface
// Exclude an interface from becoming a backup
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (excludeInterface *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface) GetEntityData() *types.CommonEntityData {
    excludeInterface.EntityData.YFilter = excludeInterface.YFilter
    excludeInterface.EntityData.YangName = "exclude-interface"
    excludeInterface.EntityData.BundleName = "cisco_ios_xr"
    excludeInterface.EntityData.ParentYangName = "exclude-interfaces"
    excludeInterface.EntityData.SegmentPath = "exclude-interface" + types.AddKeyToken(excludeInterface.InterfaceName, "interface-name")
    excludeInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/interfaces/interface/fast-reroute/per-link/exclude-interfaces/" + excludeInterface.EntityData.SegmentPath
    excludeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterface.EntityData.Children = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", excludeInterface.InterfaceName})

    excludeInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(excludeInterface.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix
// Fast-reroute per-link configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Use only interfaces on the candidate list as a backup path. The type is
    // bool. The default value is false.
    FastRerouteUseCandidateOnly interface{}

    // Fast-reroute per-link/per-prefix candidate interface configuration.
    CandidateInterfaces Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces

    // Fast-reroute per-link/per-prefix exclude interface configuration.
    ExcludeInterfaces Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces
}

func (perPrefix *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix) GetEntityData() *types.CommonEntityData {
    perPrefix.EntityData.YFilter = perPrefix.YFilter
    perPrefix.EntityData.YangName = "per-prefix"
    perPrefix.EntityData.BundleName = "cisco_ios_xr"
    perPrefix.EntityData.ParentYangName = "fast-reroute"
    perPrefix.EntityData.SegmentPath = "per-prefix"
    perPrefix.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/interfaces/interface/fast-reroute/" + perPrefix.EntityData.SegmentPath
    perPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    perPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    perPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    perPrefix.EntityData.Children = types.NewOrderedMap()
    perPrefix.EntityData.Children.Append("candidate-interfaces", types.YChild{"CandidateInterfaces", &perPrefix.CandidateInterfaces})
    perPrefix.EntityData.Children.Append("exclude-interfaces", types.YChild{"ExcludeInterfaces", &perPrefix.ExcludeInterfaces})
    perPrefix.EntityData.Leafs = types.NewOrderedMap()
    perPrefix.EntityData.Leafs.Append("fast-reroute-use-candidate-only", types.YLeaf{"FastRerouteUseCandidateOnly", perPrefix.FastRerouteUseCandidateOnly})

    perPrefix.EntityData.YListKeys = []string {}

    return &(perPrefix.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces
// Fast-reroute per-link/per-prefix candidate
// interface configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Candidate backup interface. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface.
    CandidateInterface []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface
}

func (candidateInterfaces *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces) GetEntityData() *types.CommonEntityData {
    candidateInterfaces.EntityData.YFilter = candidateInterfaces.YFilter
    candidateInterfaces.EntityData.YangName = "candidate-interfaces"
    candidateInterfaces.EntityData.BundleName = "cisco_ios_xr"
    candidateInterfaces.EntityData.ParentYangName = "per-prefix"
    candidateInterfaces.EntityData.SegmentPath = "candidate-interfaces"
    candidateInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/interfaces/interface/fast-reroute/per-prefix/" + candidateInterfaces.EntityData.SegmentPath
    candidateInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterfaces.EntityData.Children = types.NewOrderedMap()
    candidateInterfaces.EntityData.Children.Append("candidate-interface", types.YChild{"CandidateInterface", nil})
    for i := range candidateInterfaces.CandidateInterface {
        candidateInterfaces.EntityData.Children.Append(types.GetSegmentPath(candidateInterfaces.CandidateInterface[i]), types.YChild{"CandidateInterface", candidateInterfaces.CandidateInterface[i]})
    }
    candidateInterfaces.EntityData.Leafs = types.NewOrderedMap()

    candidateInterfaces.EntityData.YListKeys = []string {}

    return &(candidateInterfaces.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface
// Candidate backup interface
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (candidateInterface *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface) GetEntityData() *types.CommonEntityData {
    candidateInterface.EntityData.YFilter = candidateInterface.YFilter
    candidateInterface.EntityData.YangName = "candidate-interface"
    candidateInterface.EntityData.BundleName = "cisco_ios_xr"
    candidateInterface.EntityData.ParentYangName = "candidate-interfaces"
    candidateInterface.EntityData.SegmentPath = "candidate-interface" + types.AddKeyToken(candidateInterface.InterfaceName, "interface-name")
    candidateInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/interfaces/interface/fast-reroute/per-prefix/candidate-interfaces/" + candidateInterface.EntityData.SegmentPath
    candidateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterface.EntityData.Children = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", candidateInterface.InterfaceName})

    candidateInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(candidateInterface.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces
// Fast-reroute per-link/per-prefix exclude
// interface configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exclude an interface from becoming a backup. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface.
    ExcludeInterface []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface
}

func (excludeInterfaces *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces) GetEntityData() *types.CommonEntityData {
    excludeInterfaces.EntityData.YFilter = excludeInterfaces.YFilter
    excludeInterfaces.EntityData.YangName = "exclude-interfaces"
    excludeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    excludeInterfaces.EntityData.ParentYangName = "per-prefix"
    excludeInterfaces.EntityData.SegmentPath = "exclude-interfaces"
    excludeInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/interfaces/interface/fast-reroute/per-prefix/" + excludeInterfaces.EntityData.SegmentPath
    excludeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterfaces.EntityData.Children = types.NewOrderedMap()
    excludeInterfaces.EntityData.Children.Append("exclude-interface", types.YChild{"ExcludeInterface", nil})
    for i := range excludeInterfaces.ExcludeInterface {
        excludeInterfaces.EntityData.Children.Append(types.GetSegmentPath(excludeInterfaces.ExcludeInterface[i]), types.YChild{"ExcludeInterface", excludeInterfaces.ExcludeInterface[i]})
    }
    excludeInterfaces.EntityData.Leafs = types.NewOrderedMap()

    excludeInterfaces.EntityData.YListKeys = []string {}

    return &(excludeInterfaces.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface
// Exclude an interface from becoming a backup
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (excludeInterface *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_Interfaces_Interface_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface) GetEntityData() *types.CommonEntityData {
    excludeInterface.EntityData.YFilter = excludeInterface.YFilter
    excludeInterface.EntityData.YangName = "exclude-interface"
    excludeInterface.EntityData.BundleName = "cisco_ios_xr"
    excludeInterface.EntityData.ParentYangName = "exclude-interfaces"
    excludeInterface.EntityData.SegmentPath = "exclude-interface" + types.AddKeyToken(excludeInterface.InterfaceName, "interface-name")
    excludeInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/interfaces/interface/fast-reroute/per-prefix/exclude-interfaces/" + excludeInterface.EntityData.SegmentPath
    excludeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterface.EntityData.Children = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", excludeInterface.InterfaceName})

    excludeInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(excludeInterface.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope
// Area Scope Configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fast-reroute configuration.
    FastReroute Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute
}

func (areaScope *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope) GetEntityData() *types.CommonEntityData {
    areaScope.EntityData.YFilter = areaScope.YFilter
    areaScope.EntityData.YangName = "area-scope"
    areaScope.EntityData.BundleName = "cisco_ios_xr"
    areaScope.EntityData.ParentYangName = "area-area-id"
    areaScope.EntityData.SegmentPath = "area-scope"
    areaScope.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/" + areaScope.EntityData.SegmentPath
    areaScope.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    areaScope.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    areaScope.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    areaScope.EntityData.Children = types.NewOrderedMap()
    areaScope.EntityData.Children.Append("fast-reroute", types.YChild{"FastReroute", &areaScope.FastReroute})
    areaScope.EntityData.Leafs = types.NewOrderedMap()

    areaScope.EntityData.YListKeys = []string {}

    return &(areaScope.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute
// Fast-reroute configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable Fast-reroute per-link or per-prefix. The type is
    // Ospfv3FastReroute.
    FastRerouteEnable interface{}

    // Fast-reroute per-link configuration.
    PerLink Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink

    // Fast-reroute per-link configuration.
    PerPrefix Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix
}

func (fastReroute *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute) GetEntityData() *types.CommonEntityData {
    fastReroute.EntityData.YFilter = fastReroute.YFilter
    fastReroute.EntityData.YangName = "fast-reroute"
    fastReroute.EntityData.BundleName = "cisco_ios_xr"
    fastReroute.EntityData.ParentYangName = "area-scope"
    fastReroute.EntityData.SegmentPath = "fast-reroute"
    fastReroute.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/area-scope/" + fastReroute.EntityData.SegmentPath
    fastReroute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fastReroute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fastReroute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fastReroute.EntityData.Children = types.NewOrderedMap()
    fastReroute.EntityData.Children.Append("per-link", types.YChild{"PerLink", &fastReroute.PerLink})
    fastReroute.EntityData.Children.Append("per-prefix", types.YChild{"PerPrefix", &fastReroute.PerPrefix})
    fastReroute.EntityData.Leafs = types.NewOrderedMap()
    fastReroute.EntityData.Leafs.Append("fast-reroute-enable", types.YLeaf{"FastRerouteEnable", fastReroute.FastRerouteEnable})

    fastReroute.EntityData.YListKeys = []string {}

    return &(fastReroute.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink
// Fast-reroute per-link configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Use only interfaces on the candidate list as a backup path. The type is
    // bool. The default value is false.
    FastRerouteUseCandidateOnly interface{}

    // Fast-reroute per-link/per-prefix candidate interface configuration.
    CandidateInterfaces Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_CandidateInterfaces

    // Fast-reroute per-link/per-prefix exclude interface configuration.
    ExcludeInterfaces Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_ExcludeInterfaces
}

func (perLink *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink) GetEntityData() *types.CommonEntityData {
    perLink.EntityData.YFilter = perLink.YFilter
    perLink.EntityData.YangName = "per-link"
    perLink.EntityData.BundleName = "cisco_ios_xr"
    perLink.EntityData.ParentYangName = "fast-reroute"
    perLink.EntityData.SegmentPath = "per-link"
    perLink.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/area-scope/fast-reroute/" + perLink.EntityData.SegmentPath
    perLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    perLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    perLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    perLink.EntityData.Children = types.NewOrderedMap()
    perLink.EntityData.Children.Append("candidate-interfaces", types.YChild{"CandidateInterfaces", &perLink.CandidateInterfaces})
    perLink.EntityData.Children.Append("exclude-interfaces", types.YChild{"ExcludeInterfaces", &perLink.ExcludeInterfaces})
    perLink.EntityData.Leafs = types.NewOrderedMap()
    perLink.EntityData.Leafs.Append("fast-reroute-use-candidate-only", types.YLeaf{"FastRerouteUseCandidateOnly", perLink.FastRerouteUseCandidateOnly})

    perLink.EntityData.YListKeys = []string {}

    return &(perLink.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_CandidateInterfaces
// Fast-reroute per-link/per-prefix candidate
// interface configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_CandidateInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Candidate backup interface. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface.
    CandidateInterface []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface
}

func (candidateInterfaces *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_CandidateInterfaces) GetEntityData() *types.CommonEntityData {
    candidateInterfaces.EntityData.YFilter = candidateInterfaces.YFilter
    candidateInterfaces.EntityData.YangName = "candidate-interfaces"
    candidateInterfaces.EntityData.BundleName = "cisco_ios_xr"
    candidateInterfaces.EntityData.ParentYangName = "per-link"
    candidateInterfaces.EntityData.SegmentPath = "candidate-interfaces"
    candidateInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/area-scope/fast-reroute/per-link/" + candidateInterfaces.EntityData.SegmentPath
    candidateInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterfaces.EntityData.Children = types.NewOrderedMap()
    candidateInterfaces.EntityData.Children.Append("candidate-interface", types.YChild{"CandidateInterface", nil})
    for i := range candidateInterfaces.CandidateInterface {
        candidateInterfaces.EntityData.Children.Append(types.GetSegmentPath(candidateInterfaces.CandidateInterface[i]), types.YChild{"CandidateInterface", candidateInterfaces.CandidateInterface[i]})
    }
    candidateInterfaces.EntityData.Leafs = types.NewOrderedMap()

    candidateInterfaces.EntityData.YListKeys = []string {}

    return &(candidateInterfaces.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface
// Candidate backup interface
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (candidateInterface *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface) GetEntityData() *types.CommonEntityData {
    candidateInterface.EntityData.YFilter = candidateInterface.YFilter
    candidateInterface.EntityData.YangName = "candidate-interface"
    candidateInterface.EntityData.BundleName = "cisco_ios_xr"
    candidateInterface.EntityData.ParentYangName = "candidate-interfaces"
    candidateInterface.EntityData.SegmentPath = "candidate-interface" + types.AddKeyToken(candidateInterface.InterfaceName, "interface-name")
    candidateInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/area-scope/fast-reroute/per-link/candidate-interfaces/" + candidateInterface.EntityData.SegmentPath
    candidateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterface.EntityData.Children = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", candidateInterface.InterfaceName})

    candidateInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(candidateInterface.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_ExcludeInterfaces
// Fast-reroute per-link/per-prefix exclude
// interface configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_ExcludeInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exclude an interface from becoming a backup. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface.
    ExcludeInterface []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface
}

func (excludeInterfaces *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_ExcludeInterfaces) GetEntityData() *types.CommonEntityData {
    excludeInterfaces.EntityData.YFilter = excludeInterfaces.YFilter
    excludeInterfaces.EntityData.YangName = "exclude-interfaces"
    excludeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    excludeInterfaces.EntityData.ParentYangName = "per-link"
    excludeInterfaces.EntityData.SegmentPath = "exclude-interfaces"
    excludeInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/area-scope/fast-reroute/per-link/" + excludeInterfaces.EntityData.SegmentPath
    excludeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterfaces.EntityData.Children = types.NewOrderedMap()
    excludeInterfaces.EntityData.Children.Append("exclude-interface", types.YChild{"ExcludeInterface", nil})
    for i := range excludeInterfaces.ExcludeInterface {
        excludeInterfaces.EntityData.Children.Append(types.GetSegmentPath(excludeInterfaces.ExcludeInterface[i]), types.YChild{"ExcludeInterface", excludeInterfaces.ExcludeInterface[i]})
    }
    excludeInterfaces.EntityData.Leafs = types.NewOrderedMap()

    excludeInterfaces.EntityData.YListKeys = []string {}

    return &(excludeInterfaces.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface
// Exclude an interface from becoming a backup
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (excludeInterface *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface) GetEntityData() *types.CommonEntityData {
    excludeInterface.EntityData.YFilter = excludeInterface.YFilter
    excludeInterface.EntityData.YangName = "exclude-interface"
    excludeInterface.EntityData.BundleName = "cisco_ios_xr"
    excludeInterface.EntityData.ParentYangName = "exclude-interfaces"
    excludeInterface.EntityData.SegmentPath = "exclude-interface" + types.AddKeyToken(excludeInterface.InterfaceName, "interface-name")
    excludeInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/area-scope/fast-reroute/per-link/exclude-interfaces/" + excludeInterface.EntityData.SegmentPath
    excludeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterface.EntityData.Children = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", excludeInterface.InterfaceName})

    excludeInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(excludeInterface.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix
// Fast-reroute per-link configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Use only interfaces on the candidate list as a backup path. The type is
    // bool. The default value is false.
    FastRerouteUseCandidateOnly interface{}

    // Fast-reroute per-link/per-prefix candidate interface configuration.
    CandidateInterfaces Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_CandidateInterfaces

    // Fast-reroute per-link/per-prefix exclude interface configuration.
    ExcludeInterfaces Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces
}

func (perPrefix *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix) GetEntityData() *types.CommonEntityData {
    perPrefix.EntityData.YFilter = perPrefix.YFilter
    perPrefix.EntityData.YangName = "per-prefix"
    perPrefix.EntityData.BundleName = "cisco_ios_xr"
    perPrefix.EntityData.ParentYangName = "fast-reroute"
    perPrefix.EntityData.SegmentPath = "per-prefix"
    perPrefix.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/area-scope/fast-reroute/" + perPrefix.EntityData.SegmentPath
    perPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    perPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    perPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    perPrefix.EntityData.Children = types.NewOrderedMap()
    perPrefix.EntityData.Children.Append("candidate-interfaces", types.YChild{"CandidateInterfaces", &perPrefix.CandidateInterfaces})
    perPrefix.EntityData.Children.Append("exclude-interfaces", types.YChild{"ExcludeInterfaces", &perPrefix.ExcludeInterfaces})
    perPrefix.EntityData.Leafs = types.NewOrderedMap()
    perPrefix.EntityData.Leafs.Append("fast-reroute-use-candidate-only", types.YLeaf{"FastRerouteUseCandidateOnly", perPrefix.FastRerouteUseCandidateOnly})

    perPrefix.EntityData.YListKeys = []string {}

    return &(perPrefix.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_CandidateInterfaces
// Fast-reroute per-link/per-prefix candidate
// interface configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_CandidateInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Candidate backup interface. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface.
    CandidateInterface []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface
}

func (candidateInterfaces *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_CandidateInterfaces) GetEntityData() *types.CommonEntityData {
    candidateInterfaces.EntityData.YFilter = candidateInterfaces.YFilter
    candidateInterfaces.EntityData.YangName = "candidate-interfaces"
    candidateInterfaces.EntityData.BundleName = "cisco_ios_xr"
    candidateInterfaces.EntityData.ParentYangName = "per-prefix"
    candidateInterfaces.EntityData.SegmentPath = "candidate-interfaces"
    candidateInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/area-scope/fast-reroute/per-prefix/" + candidateInterfaces.EntityData.SegmentPath
    candidateInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterfaces.EntityData.Children = types.NewOrderedMap()
    candidateInterfaces.EntityData.Children.Append("candidate-interface", types.YChild{"CandidateInterface", nil})
    for i := range candidateInterfaces.CandidateInterface {
        candidateInterfaces.EntityData.Children.Append(types.GetSegmentPath(candidateInterfaces.CandidateInterface[i]), types.YChild{"CandidateInterface", candidateInterfaces.CandidateInterface[i]})
    }
    candidateInterfaces.EntityData.Leafs = types.NewOrderedMap()

    candidateInterfaces.EntityData.YListKeys = []string {}

    return &(candidateInterfaces.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface
// Candidate backup interface
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (candidateInterface *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface) GetEntityData() *types.CommonEntityData {
    candidateInterface.EntityData.YFilter = candidateInterface.YFilter
    candidateInterface.EntityData.YangName = "candidate-interface"
    candidateInterface.EntityData.BundleName = "cisco_ios_xr"
    candidateInterface.EntityData.ParentYangName = "candidate-interfaces"
    candidateInterface.EntityData.SegmentPath = "candidate-interface" + types.AddKeyToken(candidateInterface.InterfaceName, "interface-name")
    candidateInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/area-scope/fast-reroute/per-prefix/candidate-interfaces/" + candidateInterface.EntityData.SegmentPath
    candidateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterface.EntityData.Children = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", candidateInterface.InterfaceName})

    candidateInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(candidateInterface.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces
// Fast-reroute per-link/per-prefix exclude
// interface configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exclude an interface from becoming a backup. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface.
    ExcludeInterface []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface
}

func (excludeInterfaces *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces) GetEntityData() *types.CommonEntityData {
    excludeInterfaces.EntityData.YFilter = excludeInterfaces.YFilter
    excludeInterfaces.EntityData.YangName = "exclude-interfaces"
    excludeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    excludeInterfaces.EntityData.ParentYangName = "per-prefix"
    excludeInterfaces.EntityData.SegmentPath = "exclude-interfaces"
    excludeInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/area-scope/fast-reroute/per-prefix/" + excludeInterfaces.EntityData.SegmentPath
    excludeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterfaces.EntityData.Children = types.NewOrderedMap()
    excludeInterfaces.EntityData.Children.Append("exclude-interface", types.YChild{"ExcludeInterface", nil})
    for i := range excludeInterfaces.ExcludeInterface {
        excludeInterfaces.EntityData.Children.Append(types.GetSegmentPath(excludeInterfaces.ExcludeInterface[i]), types.YChild{"ExcludeInterface", excludeInterfaces.ExcludeInterface[i]})
    }
    excludeInterfaces.EntityData.Leafs = types.NewOrderedMap()

    excludeInterfaces.EntityData.YListKeys = []string {}

    return &(excludeInterfaces.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface
// Exclude an interface from becoming a backup
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (excludeInterface *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_AreaScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface) GetEntityData() *types.CommonEntityData {
    excludeInterface.EntityData.YFilter = excludeInterface.YFilter
    excludeInterface.EntityData.YangName = "exclude-interface"
    excludeInterface.EntityData.BundleName = "cisco_ios_xr"
    excludeInterface.EntityData.ParentYangName = "exclude-interfaces"
    excludeInterface.EntityData.SegmentPath = "exclude-interface" + types.AddKeyToken(excludeInterface.InterfaceName, "interface-name")
    excludeInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/area-scope/fast-reroute/per-prefix/exclude-interfaces/" + excludeInterface.EntityData.SegmentPath
    excludeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterface.EntityData.Children = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", excludeInterface.InterfaceName})

    excludeInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(excludeInterface.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_ShamLinks
// Sham Link sub-mode
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_ShamLinks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ShamLink local and remote endpoints. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_ShamLinks_ShamLink.
    ShamLink []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_ShamLinks_ShamLink
}

func (shamLinks *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_ShamLinks) GetEntityData() *types.CommonEntityData {
    shamLinks.EntityData.YFilter = shamLinks.YFilter
    shamLinks.EntityData.YangName = "sham-links"
    shamLinks.EntityData.BundleName = "cisco_ios_xr"
    shamLinks.EntityData.ParentYangName = "area-area-id"
    shamLinks.EntityData.SegmentPath = "sham-links"
    shamLinks.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/" + shamLinks.EntityData.SegmentPath
    shamLinks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shamLinks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shamLinks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shamLinks.EntityData.Children = types.NewOrderedMap()
    shamLinks.EntityData.Children.Append("sham-link", types.YChild{"ShamLink", nil})
    for i := range shamLinks.ShamLink {
        shamLinks.EntityData.Children.Append(types.GetSegmentPath(shamLinks.ShamLink[i]), types.YChild{"ShamLink", shamLinks.ShamLink[i]})
    }
    shamLinks.EntityData.Leafs = types.NewOrderedMap()

    shamLinks.EntityData.YListKeys = []string {}

    return &(shamLinks.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_ShamLinks_ShamLink
// ShamLink local and remote endpoints
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_ShamLinks_ShamLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Local sham-link endpoint. The type is string with
    // pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    SourceAddress interface{}

    // This attribute is a key. Remote sham-link endpoint. The type is string with
    // pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    DestinationAddress interface{}

    // Enable sham link. The type is interface{}.
    Enable interface{}

    // Time between HELLO packets. The type is interface{} with range: 1..65535.
    // Units are second.
    HelloInterval interface{}

    // Interval after which a neighbor is declared dead (in seconds). The type is
    // interface{} with range: 1..65535. Units are second.
    DeadInterval interface{}

    // Specify the transmit interval in seconds. The type is interface{} with
    // range: 1..65535. Units are second.
    RetransmitInterval interface{}

    // Specify the transmit delay in seconds. The type is interface{} with range:
    // 1..65535. Units are second.
    TransmitDelay interface{}

    // Authenticate OSPFv3 packets.
    Authentication Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_ShamLinks_ShamLink_Authentication

    // Encrypt and authenticate OSPFv3 packets.
    Encryption Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_ShamLinks_ShamLink_Encryption
}

func (shamLink *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_ShamLinks_ShamLink) GetEntityData() *types.CommonEntityData {
    shamLink.EntityData.YFilter = shamLink.YFilter
    shamLink.EntityData.YangName = "sham-link"
    shamLink.EntityData.BundleName = "cisco_ios_xr"
    shamLink.EntityData.ParentYangName = "sham-links"
    shamLink.EntityData.SegmentPath = "sham-link" + types.AddKeyToken(shamLink.SourceAddress, "source-address") + types.AddKeyToken(shamLink.DestinationAddress, "destination-address")
    shamLink.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/sham-links/" + shamLink.EntityData.SegmentPath
    shamLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    shamLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    shamLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    shamLink.EntityData.Children = types.NewOrderedMap()
    shamLink.EntityData.Children.Append("authentication", types.YChild{"Authentication", &shamLink.Authentication})
    shamLink.EntityData.Children.Append("encryption", types.YChild{"Encryption", &shamLink.Encryption})
    shamLink.EntityData.Leafs = types.NewOrderedMap()
    shamLink.EntityData.Leafs.Append("source-address", types.YLeaf{"SourceAddress", shamLink.SourceAddress})
    shamLink.EntityData.Leafs.Append("destination-address", types.YLeaf{"DestinationAddress", shamLink.DestinationAddress})
    shamLink.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", shamLink.Enable})
    shamLink.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", shamLink.HelloInterval})
    shamLink.EntityData.Leafs.Append("dead-interval", types.YLeaf{"DeadInterval", shamLink.DeadInterval})
    shamLink.EntityData.Leafs.Append("retransmit-interval", types.YLeaf{"RetransmitInterval", shamLink.RetransmitInterval})
    shamLink.EntityData.Leafs.Append("transmit-delay", types.YLeaf{"TransmitDelay", shamLink.TransmitDelay})

    shamLink.EntityData.YListKeys = []string {"SourceAddress", "DestinationAddress"}

    return &(shamLink.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_ShamLinks_ShamLink_Authentication
// Authenticate OSPFv3 packets
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_ShamLinks_ShamLink_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec AH authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Use the MD5 or SHA1 algorithm. The type is Ospfv3Authentication.
    Algorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    Password interface{}
}

func (authentication *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_ShamLinks_ShamLink_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "sham-link"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/sham-links/sham-link/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", authentication.Enable})
    authentication.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", authentication.Spi})
    authentication.EntityData.Leafs.Append("algorithm", types.YLeaf{"Algorithm", authentication.Algorithm})
    authentication.EntityData.Leafs.Append("password", types.YLeaf{"Password", authentication.Password})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_ShamLinks_ShamLink_Encryption
// Encrypt and authenticate OSPFv3 packets
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_ShamLinks_ShamLink_Encryption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec ESP authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Specify the encryption algorithm. The type is Ospfv3EncryptionAlgorithm.
    EncryptionAlgorithm interface{}

    // Encryption password. The type is string with pattern: b'(!.+)|([^!].+)'.
    EncryptionPassword interface{}

    // Use the NULL, MD5 or SHA1 algorithm. The type is Ospfv3AuthenticationType2.
    AuthenticationAlgorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    AuthenticationPassword interface{}
}

func (encryption *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_ShamLinks_ShamLink_Encryption) GetEntityData() *types.CommonEntityData {
    encryption.EntityData.YFilter = encryption.YFilter
    encryption.EntityData.YangName = "encryption"
    encryption.EntityData.BundleName = "cisco_ios_xr"
    encryption.EntityData.ParentYangName = "sham-link"
    encryption.EntityData.SegmentPath = "encryption"
    encryption.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/sham-links/sham-link/" + encryption.EntityData.SegmentPath
    encryption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryption.EntityData.Children = types.NewOrderedMap()
    encryption.EntityData.Leafs = types.NewOrderedMap()
    encryption.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", encryption.Enable})
    encryption.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", encryption.Spi})
    encryption.EntityData.Leafs.Append("encryption-algorithm", types.YLeaf{"EncryptionAlgorithm", encryption.EncryptionAlgorithm})
    encryption.EntityData.Leafs.Append("encryption-password", types.YLeaf{"EncryptionPassword", encryption.EncryptionPassword})
    encryption.EntityData.Leafs.Append("authentication-algorithm", types.YLeaf{"AuthenticationAlgorithm", encryption.AuthenticationAlgorithm})
    encryption.EntityData.Leafs.Append("authentication-password", types.YLeaf{"AuthenticationPassword", encryption.AuthenticationPassword})

    encryption.EntityData.YListKeys = []string {}

    return &(encryption.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_VirtualLinks
// Virtual link sub-mode
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_VirtualLinks struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Router ID of virtual link neighbor. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_VirtualLinks_VirtualLink.
    VirtualLink []*Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_VirtualLinks_VirtualLink
}

func (virtualLinks *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_VirtualLinks) GetEntityData() *types.CommonEntityData {
    virtualLinks.EntityData.YFilter = virtualLinks.YFilter
    virtualLinks.EntityData.YangName = "virtual-links"
    virtualLinks.EntityData.BundleName = "cisco_ios_xr"
    virtualLinks.EntityData.ParentYangName = "area-area-id"
    virtualLinks.EntityData.SegmentPath = "virtual-links"
    virtualLinks.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/" + virtualLinks.EntityData.SegmentPath
    virtualLinks.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    virtualLinks.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    virtualLinks.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    virtualLinks.EntityData.Children = types.NewOrderedMap()
    virtualLinks.EntityData.Children.Append("virtual-link", types.YChild{"VirtualLink", nil})
    for i := range virtualLinks.VirtualLink {
        virtualLinks.EntityData.Children.Append(types.GetSegmentPath(virtualLinks.VirtualLink[i]), types.YChild{"VirtualLink", virtualLinks.VirtualLink[i]})
    }
    virtualLinks.EntityData.Leafs = types.NewOrderedMap()

    virtualLinks.EntityData.YListKeys = []string {}

    return &(virtualLinks.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_VirtualLinks_VirtualLink
// Router ID of virtual link neighbor
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_VirtualLinks_VirtualLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Router ID of virtual link neighbor. The type is
    // string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    VirtualLinkAddress interface{}

    // Enabled virtual link. The type is interface{}.
    Enable interface{}

    // Time between HELLO packets. The type is interface{} with range: 1..65535.
    // Units are second.
    HelloInterval interface{}

    // Interval after which a neighbor is declared dead (in seconds). The type is
    // interface{} with range: 1..65535. Units are second.
    DeadInterval interface{}

    // Specify the transmit interval in seconds. The type is interface{} with
    // range: 1..65535. Units are second.
    RetransmitInterval interface{}

    // Specify the transmit delay in seconds. The type is interface{} with range:
    // 1..65535. Units are second.
    TransmitDelay interface{}

    // Authenticate OSPFv3 packets.
    Authentication Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_VirtualLinks_VirtualLink_Authentication

    // Encrypt and authenticate OSPFv3 packets.
    Encryption Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_VirtualLinks_VirtualLink_Encryption
}

func (virtualLink *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_VirtualLinks_VirtualLink) GetEntityData() *types.CommonEntityData {
    virtualLink.EntityData.YFilter = virtualLink.YFilter
    virtualLink.EntityData.YangName = "virtual-link"
    virtualLink.EntityData.BundleName = "cisco_ios_xr"
    virtualLink.EntityData.ParentYangName = "virtual-links"
    virtualLink.EntityData.SegmentPath = "virtual-link" + types.AddKeyToken(virtualLink.VirtualLinkAddress, "virtual-link-address")
    virtualLink.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/virtual-links/" + virtualLink.EntityData.SegmentPath
    virtualLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    virtualLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    virtualLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    virtualLink.EntityData.Children = types.NewOrderedMap()
    virtualLink.EntityData.Children.Append("authentication", types.YChild{"Authentication", &virtualLink.Authentication})
    virtualLink.EntityData.Children.Append("encryption", types.YChild{"Encryption", &virtualLink.Encryption})
    virtualLink.EntityData.Leafs = types.NewOrderedMap()
    virtualLink.EntityData.Leafs.Append("virtual-link-address", types.YLeaf{"VirtualLinkAddress", virtualLink.VirtualLinkAddress})
    virtualLink.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", virtualLink.Enable})
    virtualLink.EntityData.Leafs.Append("hello-interval", types.YLeaf{"HelloInterval", virtualLink.HelloInterval})
    virtualLink.EntityData.Leafs.Append("dead-interval", types.YLeaf{"DeadInterval", virtualLink.DeadInterval})
    virtualLink.EntityData.Leafs.Append("retransmit-interval", types.YLeaf{"RetransmitInterval", virtualLink.RetransmitInterval})
    virtualLink.EntityData.Leafs.Append("transmit-delay", types.YLeaf{"TransmitDelay", virtualLink.TransmitDelay})

    virtualLink.EntityData.YListKeys = []string {"VirtualLinkAddress"}

    return &(virtualLink.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_VirtualLinks_VirtualLink_Authentication
// Authenticate OSPFv3 packets
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_VirtualLinks_VirtualLink_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec AH authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Use the MD5 or SHA1 algorithm. The type is Ospfv3Authentication.
    Algorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    Password interface{}
}

func (authentication *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_VirtualLinks_VirtualLink_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "virtual-link"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/virtual-links/virtual-link/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", authentication.Enable})
    authentication.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", authentication.Spi})
    authentication.EntityData.Leafs.Append("algorithm", types.YLeaf{"Algorithm", authentication.Algorithm})
    authentication.EntityData.Leafs.Append("password", types.YLeaf{"Password", authentication.Password})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_VirtualLinks_VirtualLink_Encryption
// Encrypt and authenticate OSPFv3 packets
type Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_VirtualLinks_VirtualLink_Encryption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec ESP authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Specify the encryption algorithm. The type is Ospfv3EncryptionAlgorithm.
    EncryptionAlgorithm interface{}

    // Encryption password. The type is string with pattern: b'(!.+)|([^!].+)'.
    EncryptionPassword interface{}

    // Use the NULL, MD5 or SHA1 algorithm. The type is Ospfv3AuthenticationType2.
    AuthenticationAlgorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    AuthenticationPassword interface{}
}

func (encryption *Ospfv3_Processes_Process_Vrfs_Vrf_AreaAddresses_AreaAreaId_VirtualLinks_VirtualLink_Encryption) GetEntityData() *types.CommonEntityData {
    encryption.EntityData.YFilter = encryption.YFilter
    encryption.EntityData.YangName = "encryption"
    encryption.EntityData.BundleName = "cisco_ios_xr"
    encryption.EntityData.ParentYangName = "virtual-link"
    encryption.EntityData.SegmentPath = "encryption"
    encryption.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/area-addresses/area-area-id/virtual-links/virtual-link/" + encryption.EntityData.SegmentPath
    encryption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryption.EntityData.Children = types.NewOrderedMap()
    encryption.EntityData.Leafs = types.NewOrderedMap()
    encryption.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", encryption.Enable})
    encryption.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", encryption.Spi})
    encryption.EntityData.Leafs.Append("encryption-algorithm", types.YLeaf{"EncryptionAlgorithm", encryption.EncryptionAlgorithm})
    encryption.EntityData.Leafs.Append("encryption-password", types.YLeaf{"EncryptionPassword", encryption.EncryptionPassword})
    encryption.EntityData.Leafs.Append("authentication-algorithm", types.YLeaf{"AuthenticationAlgorithm", encryption.AuthenticationAlgorithm})
    encryption.EntityData.Leafs.Append("authentication-password", types.YLeaf{"AuthenticationPassword", encryption.AuthenticationPassword})

    encryption.EntityData.YListKeys = []string {}

    return &(encryption.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_Timers
// Adjust routing timers
type Ospfv3_Processes_Process_Vrfs_Vrf_Timers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Pacing timers.
    Pacing Ospfv3_Processes_Process_Vrfs_Vrf_Timers_Pacing

    // LSA timers.
    LsaTimers Ospfv3_Processes_Process_Vrfs_Vrf_Timers_LsaTimers

    // Throttle timers.
    Throttle Ospfv3_Processes_Process_Vrfs_Vrf_Timers_Throttle
}

func (timers *Ospfv3_Processes_Process_Vrfs_Vrf_Timers) GetEntityData() *types.CommonEntityData {
    timers.EntityData.YFilter = timers.YFilter
    timers.EntityData.YangName = "timers"
    timers.EntityData.BundleName = "cisco_ios_xr"
    timers.EntityData.ParentYangName = "vrf"
    timers.EntityData.SegmentPath = "timers"
    timers.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/" + timers.EntityData.SegmentPath
    timers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timers.EntityData.Children = types.NewOrderedMap()
    timers.EntityData.Children.Append("pacing", types.YChild{"Pacing", &timers.Pacing})
    timers.EntityData.Children.Append("lsa-timers", types.YChild{"LsaTimers", &timers.LsaTimers})
    timers.EntityData.Children.Append("throttle", types.YChild{"Throttle", &timers.Throttle})
    timers.EntityData.Leafs = types.NewOrderedMap()

    timers.EntityData.YListKeys = []string {}

    return &(timers.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_Timers_Pacing
// Pacing timers
type Ospfv3_Processes_Process_Vrfs_Vrf_Timers_Pacing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The minimum interval in milliseconds to pace limit flooding on interface.
    // The type is interface{} with range: 5..100. Units are millisecond.
    Flood interface{}

    // The minimum interval in msec between neighbor retransmissions. The type is
    // interface{} with range: 5..100.
    Retransmission interface{}

    // Interval in seconds at which LSAs are grouped and refreshed, checksummed,
    // or aged. The type is interface{} with range: 10..1800. Units are second.
    LsaGroup interface{}
}

func (pacing *Ospfv3_Processes_Process_Vrfs_Vrf_Timers_Pacing) GetEntityData() *types.CommonEntityData {
    pacing.EntityData.YFilter = pacing.YFilter
    pacing.EntityData.YangName = "pacing"
    pacing.EntityData.BundleName = "cisco_ios_xr"
    pacing.EntityData.ParentYangName = "timers"
    pacing.EntityData.SegmentPath = "pacing"
    pacing.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/timers/" + pacing.EntityData.SegmentPath
    pacing.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    pacing.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    pacing.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    pacing.EntityData.Children = types.NewOrderedMap()
    pacing.EntityData.Leafs = types.NewOrderedMap()
    pacing.EntityData.Leafs.Append("flood", types.YLeaf{"Flood", pacing.Flood})
    pacing.EntityData.Leafs.Append("retransmission", types.YLeaf{"Retransmission", pacing.Retransmission})
    pacing.EntityData.Leafs.Append("lsa-group", types.YLeaf{"LsaGroup", pacing.LsaGroup})

    pacing.EntityData.YListKeys = []string {}

    return &(pacing.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_Timers_LsaTimers
// LSA timers
type Ospfv3_Processes_Process_Vrfs_Vrf_Timers_LsaTimers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The minimum interval in milliseconds between accepting the same LSA. The
    // type is interface{} with range: 0..60000. Units are millisecond.
    Arrival interface{}
}

func (lsaTimers *Ospfv3_Processes_Process_Vrfs_Vrf_Timers_LsaTimers) GetEntityData() *types.CommonEntityData {
    lsaTimers.EntityData.YFilter = lsaTimers.YFilter
    lsaTimers.EntityData.YangName = "lsa-timers"
    lsaTimers.EntityData.BundleName = "cisco_ios_xr"
    lsaTimers.EntityData.ParentYangName = "timers"
    lsaTimers.EntityData.SegmentPath = "lsa-timers"
    lsaTimers.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/timers/" + lsaTimers.EntityData.SegmentPath
    lsaTimers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsaTimers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsaTimers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsaTimers.EntityData.Children = types.NewOrderedMap()
    lsaTimers.EntityData.Leafs = types.NewOrderedMap()
    lsaTimers.EntityData.Leafs.Append("arrival", types.YLeaf{"Arrival", lsaTimers.Arrival})

    lsaTimers.EntityData.YListKeys = []string {}

    return &(lsaTimers.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_Timers_Throttle
// Throttle timers
type Ospfv3_Processes_Process_Vrfs_Vrf_Timers_Throttle struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // LSA throttle timers for all types of OSPF LSAs.
    Lsa Ospfv3_Processes_Process_Vrfs_Vrf_Timers_Throttle_Lsa

    // SPF throttle timers.
    Spf Ospfv3_Processes_Process_Vrfs_Vrf_Timers_Throttle_Spf
}

func (throttle *Ospfv3_Processes_Process_Vrfs_Vrf_Timers_Throttle) GetEntityData() *types.CommonEntityData {
    throttle.EntityData.YFilter = throttle.YFilter
    throttle.EntityData.YangName = "throttle"
    throttle.EntityData.BundleName = "cisco_ios_xr"
    throttle.EntityData.ParentYangName = "timers"
    throttle.EntityData.SegmentPath = "throttle"
    throttle.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/timers/" + throttle.EntityData.SegmentPath
    throttle.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    throttle.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    throttle.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    throttle.EntityData.Children = types.NewOrderedMap()
    throttle.EntityData.Children.Append("lsa", types.YChild{"Lsa", &throttle.Lsa})
    throttle.EntityData.Children.Append("spf", types.YChild{"Spf", &throttle.Spf})
    throttle.EntityData.Leafs = types.NewOrderedMap()

    throttle.EntityData.YListKeys = []string {}

    return &(throttle.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_Timers_Throttle_Lsa
// LSA throttle timers for all types of OSPF LSAs
type Ospfv3_Processes_Process_Vrfs_Vrf_Timers_Throttle_Lsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Delay to generate first occurrence of LSA in milliseconds. The type is
    // interface{} with range: 0..600000. Units are millisecond.
    FirstDelay interface{}

    // Minimum delay between originating the same LSA in milliseconds. The type is
    // interface{} with range: 1..600000. Units are millisecond.
    MinimumDelay interface{}

    // Maximum delay between originating the same LSA in milliseconds. The type is
    // interface{} with range: 1..600000. Units are millisecond.
    MaximumDelay interface{}
}

func (lsa *Ospfv3_Processes_Process_Vrfs_Vrf_Timers_Throttle_Lsa) GetEntityData() *types.CommonEntityData {
    lsa.EntityData.YFilter = lsa.YFilter
    lsa.EntityData.YangName = "lsa"
    lsa.EntityData.BundleName = "cisco_ios_xr"
    lsa.EntityData.ParentYangName = "throttle"
    lsa.EntityData.SegmentPath = "lsa"
    lsa.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/timers/throttle/" + lsa.EntityData.SegmentPath
    lsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsa.EntityData.Children = types.NewOrderedMap()
    lsa.EntityData.Leafs = types.NewOrderedMap()
    lsa.EntityData.Leafs.Append("first-delay", types.YLeaf{"FirstDelay", lsa.FirstDelay})
    lsa.EntityData.Leafs.Append("minimum-delay", types.YLeaf{"MinimumDelay", lsa.MinimumDelay})
    lsa.EntityData.Leafs.Append("maximum-delay", types.YLeaf{"MaximumDelay", lsa.MaximumDelay})

    lsa.EntityData.YListKeys = []string {}

    return &(lsa.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_Timers_Throttle_Spf
// SPF throttle timers
type Ospfv3_Processes_Process_Vrfs_Vrf_Timers_Throttle_Spf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Initial delay between receiving a change and starting SPF in ms. The type
    // is interface{} with range: 1..600000.
    FirstDelay interface{}

    // Minimum hold time between consecutive SPF calculations in ms. The type is
    // interface{} with range: 1..600000.
    MinimumDelay interface{}

    // Maximum wait time between consecutive SPF calculations in ms. The type is
    // interface{} with range: 1..600000.
    MaximumDelay interface{}
}

func (spf *Ospfv3_Processes_Process_Vrfs_Vrf_Timers_Throttle_Spf) GetEntityData() *types.CommonEntityData {
    spf.EntityData.YFilter = spf.YFilter
    spf.EntityData.YangName = "spf"
    spf.EntityData.BundleName = "cisco_ios_xr"
    spf.EntityData.ParentYangName = "throttle"
    spf.EntityData.SegmentPath = "spf"
    spf.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/timers/throttle/" + spf.EntityData.SegmentPath
    spf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    spf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    spf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    spf.EntityData.Children = types.NewOrderedMap()
    spf.EntityData.Leafs = types.NewOrderedMap()
    spf.EntityData.Leafs.Append("first-delay", types.YLeaf{"FirstDelay", spf.FirstDelay})
    spf.EntityData.Leafs.Append("minimum-delay", types.YLeaf{"MinimumDelay", spf.MinimumDelay})
    spf.EntityData.Leafs.Append("maximum-delay", types.YLeaf{"MaximumDelay", spf.MaximumDelay})

    spf.EntityData.YListKeys = []string {}

    return &(spf.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_SummaryPrefixes
// Summarize redistributed routes matching
// prefix/length
type Ospfv3_Processes_Process_Vrfs_Vrf_SummaryPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IPv6 address. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_SummaryPrefixes_SummaryPrefix.
    SummaryPrefix []*Ospfv3_Processes_Process_Vrfs_Vrf_SummaryPrefixes_SummaryPrefix
}

func (summaryPrefixes *Ospfv3_Processes_Process_Vrfs_Vrf_SummaryPrefixes) GetEntityData() *types.CommonEntityData {
    summaryPrefixes.EntityData.YFilter = summaryPrefixes.YFilter
    summaryPrefixes.EntityData.YangName = "summary-prefixes"
    summaryPrefixes.EntityData.BundleName = "cisco_ios_xr"
    summaryPrefixes.EntityData.ParentYangName = "vrf"
    summaryPrefixes.EntityData.SegmentPath = "summary-prefixes"
    summaryPrefixes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/" + summaryPrefixes.EntityData.SegmentPath
    summaryPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaryPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaryPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaryPrefixes.EntityData.Children = types.NewOrderedMap()
    summaryPrefixes.EntityData.Children.Append("summary-prefix", types.YChild{"SummaryPrefix", nil})
    for i := range summaryPrefixes.SummaryPrefix {
        summaryPrefixes.EntityData.Children.Append(types.GetSegmentPath(summaryPrefixes.SummaryPrefix[i]), types.YChild{"SummaryPrefix", summaryPrefixes.SummaryPrefix[i]})
    }
    summaryPrefixes.EntityData.Leafs = types.NewOrderedMap()

    summaryPrefixes.EntityData.YListKeys = []string {}

    return &(summaryPrefixes.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_SummaryPrefixes_SummaryPrefix
// IPv6 address
type Ospfv3_Processes_Process_Vrfs_Vrf_SummaryPrefixes_SummaryPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. IPv6 prefix string format. The type is string with
    // pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    Prefix interface{}

    // This attribute is a key. IPV6 prefix length. The type is interface{} with
    // range: 0..128.
    PrefixLength interface{}

    // Suppress routes matching prefix/length. The type is bool.
    NotAdvertise interface{}

    // Tag. The type is interface{} with range: 1..4294967295.
    Tag interface{}
}

func (summaryPrefix *Ospfv3_Processes_Process_Vrfs_Vrf_SummaryPrefixes_SummaryPrefix) GetEntityData() *types.CommonEntityData {
    summaryPrefix.EntityData.YFilter = summaryPrefix.YFilter
    summaryPrefix.EntityData.YangName = "summary-prefix"
    summaryPrefix.EntityData.BundleName = "cisco_ios_xr"
    summaryPrefix.EntityData.ParentYangName = "summary-prefixes"
    summaryPrefix.EntityData.SegmentPath = "summary-prefix" + types.AddKeyToken(summaryPrefix.Prefix, "prefix") + types.AddKeyToken(summaryPrefix.PrefixLength, "prefix-length")
    summaryPrefix.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/summary-prefixes/" + summaryPrefix.EntityData.SegmentPath
    summaryPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    summaryPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    summaryPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    summaryPrefix.EntityData.Children = types.NewOrderedMap()
    summaryPrefix.EntityData.Leafs = types.NewOrderedMap()
    summaryPrefix.EntityData.Leafs.Append("prefix", types.YLeaf{"Prefix", summaryPrefix.Prefix})
    summaryPrefix.EntityData.Leafs.Append("prefix-length", types.YLeaf{"PrefixLength", summaryPrefix.PrefixLength})
    summaryPrefix.EntityData.Leafs.Append("not-advertise", types.YLeaf{"NotAdvertise", summaryPrefix.NotAdvertise})
    summaryPrefix.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", summaryPrefix.Tag})

    summaryPrefix.EntityData.YListKeys = []string {"Prefix", "PrefixLength"}

    return &(summaryPrefix.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_Snmp
// SNMP configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_Snmp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SNMP context configuration. The type is string.
    Context interface{}

    // SNMP trap rate configuration.
    TrapRateLimit Ospfv3_Processes_Process_Vrfs_Vrf_Snmp_TrapRateLimit
}

func (snmp *Ospfv3_Processes_Process_Vrfs_Vrf_Snmp) GetEntityData() *types.CommonEntityData {
    snmp.EntityData.YFilter = snmp.YFilter
    snmp.EntityData.YangName = "snmp"
    snmp.EntityData.BundleName = "cisco_ios_xr"
    snmp.EntityData.ParentYangName = "vrf"
    snmp.EntityData.SegmentPath = "snmp"
    snmp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/" + snmp.EntityData.SegmentPath
    snmp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    snmp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    snmp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    snmp.EntityData.Children = types.NewOrderedMap()
    snmp.EntityData.Children.Append("trap-rate-limit", types.YChild{"TrapRateLimit", &snmp.TrapRateLimit})
    snmp.EntityData.Leafs = types.NewOrderedMap()
    snmp.EntityData.Leafs.Append("context", types.YLeaf{"Context", snmp.Context})

    snmp.EntityData.YListKeys = []string {}

    return &(snmp.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_Snmp_TrapRateLimit
// SNMP trap rate configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_Snmp_TrapRateLimit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Trap rate limit sliding window size. The type is interface{} with range:
    // 2..60.
    WindowSize interface{}

    // Max number of traps sent in window time. The type is interface{} with
    // range: 0..300.
    MaxWindowTraps interface{}
}

func (trapRateLimit *Ospfv3_Processes_Process_Vrfs_Vrf_Snmp_TrapRateLimit) GetEntityData() *types.CommonEntityData {
    trapRateLimit.EntityData.YFilter = trapRateLimit.YFilter
    trapRateLimit.EntityData.YangName = "trap-rate-limit"
    trapRateLimit.EntityData.BundleName = "cisco_ios_xr"
    trapRateLimit.EntityData.ParentYangName = "snmp"
    trapRateLimit.EntityData.SegmentPath = "trap-rate-limit"
    trapRateLimit.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/snmp/" + trapRateLimit.EntityData.SegmentPath
    trapRateLimit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    trapRateLimit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    trapRateLimit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    trapRateLimit.EntityData.Children = types.NewOrderedMap()
    trapRateLimit.EntityData.Leafs = types.NewOrderedMap()
    trapRateLimit.EntityData.Leafs.Append("window-size", types.YLeaf{"WindowSize", trapRateLimit.WindowSize})
    trapRateLimit.EntityData.Leafs.Append("max-window-traps", types.YLeaf{"MaxWindowTraps", trapRateLimit.MaxWindowTraps})

    trapRateLimit.EntityData.YListKeys = []string {}

    return &(trapRateLimit.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_FastReroute
// Fast-reroute instance scoped parameters
type Ospfv3_Processes_Process_Vrfs_Vrf_FastReroute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fast-reroute per-link global configuration.
    PerLink Ospfv3_Processes_Process_Vrfs_Vrf_FastReroute_PerLink

    // Fast-reroute per-prefix global configuration.
    PerPrefix Ospfv3_Processes_Process_Vrfs_Vrf_FastReroute_PerPrefix
}

func (fastReroute *Ospfv3_Processes_Process_Vrfs_Vrf_FastReroute) GetEntityData() *types.CommonEntityData {
    fastReroute.EntityData.YFilter = fastReroute.YFilter
    fastReroute.EntityData.YangName = "fast-reroute"
    fastReroute.EntityData.BundleName = "cisco_ios_xr"
    fastReroute.EntityData.ParentYangName = "vrf"
    fastReroute.EntityData.SegmentPath = "fast-reroute"
    fastReroute.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/" + fastReroute.EntityData.SegmentPath
    fastReroute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fastReroute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fastReroute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fastReroute.EntityData.Children = types.NewOrderedMap()
    fastReroute.EntityData.Children.Append("per-link", types.YChild{"PerLink", &fastReroute.PerLink})
    fastReroute.EntityData.Children.Append("per-prefix", types.YChild{"PerPrefix", &fastReroute.PerPrefix})
    fastReroute.EntityData.Leafs = types.NewOrderedMap()

    fastReroute.EntityData.YListKeys = []string {}

    return &(fastReroute.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_FastReroute_PerLink
// Fast-reroute per-link global configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_FastReroute_PerLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fast-reroute per-link/per-prefix priority-limit command. The type is
    // Ospfv3FastReroutePriority.
    Priority interface{}
}

func (perLink *Ospfv3_Processes_Process_Vrfs_Vrf_FastReroute_PerLink) GetEntityData() *types.CommonEntityData {
    perLink.EntityData.YFilter = perLink.YFilter
    perLink.EntityData.YangName = "per-link"
    perLink.EntityData.BundleName = "cisco_ios_xr"
    perLink.EntityData.ParentYangName = "fast-reroute"
    perLink.EntityData.SegmentPath = "per-link"
    perLink.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/fast-reroute/" + perLink.EntityData.SegmentPath
    perLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    perLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    perLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    perLink.EntityData.Children = types.NewOrderedMap()
    perLink.EntityData.Leafs = types.NewOrderedMap()
    perLink.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", perLink.Priority})

    perLink.EntityData.YListKeys = []string {}

    return &(perLink.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_FastReroute_PerPrefix
// Fast-reroute per-prefix global configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_FastReroute_PerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Disable load sharing between multiple backups. The type is interface{}.
    LoadSharingDisable interface{}

    // Fast-reroute per-link/per-prefix priority-limit command. The type is
    // Ospfv3FastReroutePriority.
    Priority interface{}

    // Fast-reroute tiebreakers configurations.
    Tiebreakers Ospfv3_Processes_Process_Vrfs_Vrf_FastReroute_PerPrefix_Tiebreakers
}

func (perPrefix *Ospfv3_Processes_Process_Vrfs_Vrf_FastReroute_PerPrefix) GetEntityData() *types.CommonEntityData {
    perPrefix.EntityData.YFilter = perPrefix.YFilter
    perPrefix.EntityData.YangName = "per-prefix"
    perPrefix.EntityData.BundleName = "cisco_ios_xr"
    perPrefix.EntityData.ParentYangName = "fast-reroute"
    perPrefix.EntityData.SegmentPath = "per-prefix"
    perPrefix.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/fast-reroute/" + perPrefix.EntityData.SegmentPath
    perPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    perPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    perPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    perPrefix.EntityData.Children = types.NewOrderedMap()
    perPrefix.EntityData.Children.Append("tiebreakers", types.YChild{"Tiebreakers", &perPrefix.Tiebreakers})
    perPrefix.EntityData.Leafs = types.NewOrderedMap()
    perPrefix.EntityData.Leafs.Append("load-sharing-disable", types.YLeaf{"LoadSharingDisable", perPrefix.LoadSharingDisable})
    perPrefix.EntityData.Leafs.Append("priority", types.YLeaf{"Priority", perPrefix.Priority})

    perPrefix.EntityData.YListKeys = []string {}

    return &(perPrefix.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_FastReroute_PerPrefix_Tiebreakers
// Fast-reroute tiebreakers configurations
type Ospfv3_Processes_Process_Vrfs_Vrf_FastReroute_PerPrefix_Tiebreakers struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fast-reroute tiebreakers configuration. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_FastReroute_PerPrefix_Tiebreakers_Tiebreaker.
    Tiebreaker []*Ospfv3_Processes_Process_Vrfs_Vrf_FastReroute_PerPrefix_Tiebreakers_Tiebreaker
}

func (tiebreakers *Ospfv3_Processes_Process_Vrfs_Vrf_FastReroute_PerPrefix_Tiebreakers) GetEntityData() *types.CommonEntityData {
    tiebreakers.EntityData.YFilter = tiebreakers.YFilter
    tiebreakers.EntityData.YangName = "tiebreakers"
    tiebreakers.EntityData.BundleName = "cisco_ios_xr"
    tiebreakers.EntityData.ParentYangName = "per-prefix"
    tiebreakers.EntityData.SegmentPath = "tiebreakers"
    tiebreakers.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/fast-reroute/per-prefix/" + tiebreakers.EntityData.SegmentPath
    tiebreakers.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tiebreakers.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tiebreakers.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tiebreakers.EntityData.Children = types.NewOrderedMap()
    tiebreakers.EntityData.Children.Append("tiebreaker", types.YChild{"Tiebreaker", nil})
    for i := range tiebreakers.Tiebreaker {
        tiebreakers.EntityData.Children.Append(types.GetSegmentPath(tiebreakers.Tiebreaker[i]), types.YChild{"Tiebreaker", tiebreakers.Tiebreaker[i]})
    }
    tiebreakers.EntityData.Leafs = types.NewOrderedMap()

    tiebreakers.EntityData.YListKeys = []string {}

    return &(tiebreakers.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_FastReroute_PerPrefix_Tiebreakers_Tiebreaker
// Fast-reroute tiebreakers configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_FastReroute_PerPrefix_Tiebreakers_Tiebreaker struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Tiebreaker type. The type is
    // Ospfv3FastRerouteTiebreakers.
    TiebreakerType interface{}

    // Index value for a tiebreaker. The type is interface{} with range: 1..255.
    // This attribute is mandatory.
    TiebreakerIndex interface{}
}

func (tiebreaker *Ospfv3_Processes_Process_Vrfs_Vrf_FastReroute_PerPrefix_Tiebreakers_Tiebreaker) GetEntityData() *types.CommonEntityData {
    tiebreaker.EntityData.YFilter = tiebreaker.YFilter
    tiebreaker.EntityData.YangName = "tiebreaker"
    tiebreaker.EntityData.BundleName = "cisco_ios_xr"
    tiebreaker.EntityData.ParentYangName = "tiebreakers"
    tiebreaker.EntityData.SegmentPath = "tiebreaker" + types.AddKeyToken(tiebreaker.TiebreakerType, "tiebreaker-type")
    tiebreaker.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/fast-reroute/per-prefix/tiebreakers/" + tiebreaker.EntityData.SegmentPath
    tiebreaker.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tiebreaker.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tiebreaker.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tiebreaker.EntityData.Children = types.NewOrderedMap()
    tiebreaker.EntityData.Leafs = types.NewOrderedMap()
    tiebreaker.EntityData.Leafs.Append("tiebreaker-type", types.YLeaf{"TiebreakerType", tiebreaker.TiebreakerType})
    tiebreaker.EntityData.Leafs.Append("tiebreaker-index", types.YLeaf{"TiebreakerIndex", tiebreaker.TiebreakerIndex})

    tiebreaker.EntityData.YListKeys = []string {"TiebreakerType"}

    return &(tiebreaker.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_Distance
// Define an administrative distance
type Ospfv3_Processes_Process_Vrfs_Vrf_Distance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Define an administrative distance. The type is interface{} with range:
    // 1..255.
    Administrative interface{}

    // OSPFv3 administrative distance.
    Ospfv3 Ospfv3_Processes_Process_Vrfs_Vrf_Distance_Ospfv3
}

func (distance *Ospfv3_Processes_Process_Vrfs_Vrf_Distance) GetEntityData() *types.CommonEntityData {
    distance.EntityData.YFilter = distance.YFilter
    distance.EntityData.YangName = "distance"
    distance.EntityData.BundleName = "cisco_ios_xr"
    distance.EntityData.ParentYangName = "vrf"
    distance.EntityData.SegmentPath = "distance"
    distance.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/" + distance.EntityData.SegmentPath
    distance.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    distance.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    distance.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    distance.EntityData.Children = types.NewOrderedMap()
    distance.EntityData.Children.Append("ospfv3", types.YChild{"Ospfv3", &distance.Ospfv3})
    distance.EntityData.Leafs = types.NewOrderedMap()
    distance.EntityData.Leafs.Append("administrative", types.YLeaf{"Administrative", distance.Administrative})

    distance.EntityData.YListKeys = []string {}

    return &(distance.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_Distance_Ospfv3
// OSPFv3 administrative distance
type Ospfv3_Processes_Process_Vrfs_Vrf_Distance_Ospfv3 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Distance for intra-area routes. The type is interface{} with range: 1..255.
    IntraArea interface{}

    // Distance for inter-area routes. The type is interface{} with range: 1..255.
    InterArea interface{}

    // Distance for external type 5 and type 7 routes. The type is interface{}
    // with range: 1..255.
    External interface{}
}

func (ospfv3 *Ospfv3_Processes_Process_Vrfs_Vrf_Distance_Ospfv3) GetEntityData() *types.CommonEntityData {
    ospfv3.EntityData.YFilter = ospfv3.YFilter
    ospfv3.EntityData.YangName = "ospfv3"
    ospfv3.EntityData.BundleName = "cisco_ios_xr"
    ospfv3.EntityData.ParentYangName = "distance"
    ospfv3.EntityData.SegmentPath = "ospfv3"
    ospfv3.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/distance/" + ospfv3.EntityData.SegmentPath
    ospfv3.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv3.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv3.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv3.EntityData.Children = types.NewOrderedMap()
    ospfv3.EntityData.Leafs = types.NewOrderedMap()
    ospfv3.EntityData.Leafs.Append("intra-area", types.YLeaf{"IntraArea", ospfv3.IntraArea})
    ospfv3.EntityData.Leafs.Append("inter-area", types.YLeaf{"InterArea", ospfv3.InterArea})
    ospfv3.EntityData.Leafs.Append("external", types.YLeaf{"External", ospfv3.External})

    ospfv3.EntityData.YListKeys = []string {}

    return &(ospfv3.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_Maximum
// Set OSPFv3 limits
type Ospfv3_Processes_Process_Vrfs_Vrf_Maximum struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify maximum number of interfaces. The type is interface{} with range:
    // 1..4294967295.
    Interfaces interface{}

    // Specify maximum number of paths per route. The type is interface{} with
    // range: 1..64.
    Paths interface{}

    // Limit number of redistributed prefixes.
    RedistributedPrefixes Ospfv3_Processes_Process_Vrfs_Vrf_Maximum_RedistributedPrefixes
}

func (maximum *Ospfv3_Processes_Process_Vrfs_Vrf_Maximum) GetEntityData() *types.CommonEntityData {
    maximum.EntityData.YFilter = maximum.YFilter
    maximum.EntityData.YangName = "maximum"
    maximum.EntityData.BundleName = "cisco_ios_xr"
    maximum.EntityData.ParentYangName = "vrf"
    maximum.EntityData.SegmentPath = "maximum"
    maximum.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/" + maximum.EntityData.SegmentPath
    maximum.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maximum.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maximum.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maximum.EntityData.Children = types.NewOrderedMap()
    maximum.EntityData.Children.Append("redistributed-prefixes", types.YChild{"RedistributedPrefixes", &maximum.RedistributedPrefixes})
    maximum.EntityData.Leafs = types.NewOrderedMap()
    maximum.EntityData.Leafs.Append("interfaces", types.YLeaf{"Interfaces", maximum.Interfaces})
    maximum.EntityData.Leafs.Append("paths", types.YLeaf{"Paths", maximum.Paths})

    maximum.EntityData.YListKeys = []string {}

    return &(maximum.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_Maximum_RedistributedPrefixes
// Limit number of redistributed prefixes
type Ospfv3_Processes_Process_Vrfs_Vrf_Maximum_RedistributedPrefixes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Maximum number of prefixes redistributed to protocol. The type is
    // interface{} with range: 1..4294967295.
    Prefixes interface{}

    // Threshold value (%) at which to generate a warning message. The type is
    // interface{} with range: 1..100.
    Threshold interface{}

    // Only give warning message when limit is exceeded. The type is interface{}.
    WarningOnly interface{}
}

func (redistributedPrefixes *Ospfv3_Processes_Process_Vrfs_Vrf_Maximum_RedistributedPrefixes) GetEntityData() *types.CommonEntityData {
    redistributedPrefixes.EntityData.YFilter = redistributedPrefixes.YFilter
    redistributedPrefixes.EntityData.YangName = "redistributed-prefixes"
    redistributedPrefixes.EntityData.BundleName = "cisco_ios_xr"
    redistributedPrefixes.EntityData.ParentYangName = "maximum"
    redistributedPrefixes.EntityData.SegmentPath = "redistributed-prefixes"
    redistributedPrefixes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/maximum/" + redistributedPrefixes.EntityData.SegmentPath
    redistributedPrefixes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistributedPrefixes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistributedPrefixes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistributedPrefixes.EntityData.Children = types.NewOrderedMap()
    redistributedPrefixes.EntityData.Leafs = types.NewOrderedMap()
    redistributedPrefixes.EntityData.Leafs.Append("prefixes", types.YLeaf{"Prefixes", redistributedPrefixes.Prefixes})
    redistributedPrefixes.EntityData.Leafs.Append("threshold", types.YLeaf{"Threshold", redistributedPrefixes.Threshold})
    redistributedPrefixes.EntityData.Leafs.Append("warning-only", types.YLeaf{"WarningOnly", redistributedPrefixes.WarningOnly})

    redistributedPrefixes.EntityData.YListKeys = []string {}

    return &(redistributedPrefixes.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes
// Redistribute information from another routing
// protocol
type Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Redistribute information from another routing protocol. The type is slice
    // of Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes_Redistribute.
    Redistribute []*Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes_Redistribute
}

func (redistributes *Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes) GetEntityData() *types.CommonEntityData {
    redistributes.EntityData.YFilter = redistributes.YFilter
    redistributes.EntityData.YangName = "redistributes"
    redistributes.EntityData.BundleName = "cisco_ios_xr"
    redistributes.EntityData.ParentYangName = "vrf"
    redistributes.EntityData.SegmentPath = "redistributes"
    redistributes.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/" + redistributes.EntityData.SegmentPath
    redistributes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistributes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistributes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistributes.EntityData.Children = types.NewOrderedMap()
    redistributes.EntityData.Children.Append("redistribute", types.YChild{"Redistribute", nil})
    for i := range redistributes.Redistribute {
        redistributes.EntityData.Children.Append(types.GetSegmentPath(redistributes.Redistribute[i]), types.YChild{"Redistribute", redistributes.Redistribute[i]})
    }
    redistributes.EntityData.Leafs = types.NewOrderedMap()

    redistributes.EntityData.YListKeys = []string {}

    return &(redistributes.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes_Redistribute
// Redistribute information from another routing
// protocol
type Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes_Redistribute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Protocol. The type is Ospfv3ProtocolType2.
    ProtocolName interface{}

    // connected or static or subscriber or mobile.
    ConnectedOrStaticOrSubscriberOrMobile Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes_Redistribute_ConnectedOrStaticOrSubscriberOrMobile

    // bgp. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes_Redistribute_Bgp.
    Bgp []*Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes_Redistribute_Bgp

    // ospfv3 or isis or application. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes_Redistribute_Ospfv3OrIsisOrApplication.
    Ospfv3OrIsisOrApplication []*Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes_Redistribute_Ospfv3OrIsisOrApplication

    // eigrp. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes_Redistribute_Eigrp.
    Eigrp []*Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes_Redistribute_Eigrp
}

func (redistribute *Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes_Redistribute) GetEntityData() *types.CommonEntityData {
    redistribute.EntityData.YFilter = redistribute.YFilter
    redistribute.EntityData.YangName = "redistribute"
    redistribute.EntityData.BundleName = "cisco_ios_xr"
    redistribute.EntityData.ParentYangName = "redistributes"
    redistribute.EntityData.SegmentPath = "redistribute" + types.AddKeyToken(redistribute.ProtocolName, "protocol-name")
    redistribute.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/redistributes/" + redistribute.EntityData.SegmentPath
    redistribute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    redistribute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    redistribute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    redistribute.EntityData.Children = types.NewOrderedMap()
    redistribute.EntityData.Children.Append("connected-or-static-or-subscriber-or-mobile", types.YChild{"ConnectedOrStaticOrSubscriberOrMobile", &redistribute.ConnectedOrStaticOrSubscriberOrMobile})
    redistribute.EntityData.Children.Append("bgp", types.YChild{"Bgp", nil})
    for i := range redistribute.Bgp {
        redistribute.EntityData.Children.Append(types.GetSegmentPath(redistribute.Bgp[i]), types.YChild{"Bgp", redistribute.Bgp[i]})
    }
    redistribute.EntityData.Children.Append("ospfv3-or-isis-or-application", types.YChild{"Ospfv3OrIsisOrApplication", nil})
    for i := range redistribute.Ospfv3OrIsisOrApplication {
        redistribute.EntityData.Children.Append(types.GetSegmentPath(redistribute.Ospfv3OrIsisOrApplication[i]), types.YChild{"Ospfv3OrIsisOrApplication", redistribute.Ospfv3OrIsisOrApplication[i]})
    }
    redistribute.EntityData.Children.Append("eigrp", types.YChild{"Eigrp", nil})
    for i := range redistribute.Eigrp {
        redistribute.EntityData.Children.Append(types.GetSegmentPath(redistribute.Eigrp[i]), types.YChild{"Eigrp", redistribute.Eigrp[i]})
    }
    redistribute.EntityData.Leafs = types.NewOrderedMap()
    redistribute.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", redistribute.ProtocolName})

    redistribute.EntityData.YListKeys = []string {"ProtocolName"}

    return &(redistribute.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes_Redistribute_ConnectedOrStaticOrSubscriberOrMobile
// connected or static or subscriber or mobile
// This type is a presence type.
type Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes_Redistribute_ConnectedOrStaticOrSubscriberOrMobile struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Redistribute OSPFv3 routes. The type is Ospfv3InternalRoute.
    InternalRouteType interface{}

    // OSPFv3 default metric. The type is interface{} with range: 0..16777214.
    DefaultMetric interface{}

    // OSPFv3 exterior metric type for redistributed routes. The type is
    // Ospfv3Metric.
    MetricType interface{}

    // Tag for routes redistributed into OSPFv3. The type is interface{} with
    // range: 0..4294967295.
    Tag interface{}

    // Route policy to redistribution. The type is string.
    RoutePolicyName interface{}

    // Redistribute OSPFv3 external routes. The type is Ospfv3ExternalRoute.
    ExternalRouteType interface{}

    // Redistribute OSPFv3 NSSA external routes. The type is
    // Ospfv3nssaExternalRoute.
    NssaExternalRouteType interface{}

    // Redistribution of OSPFv3 routes. The type is bool.
    RedistributeRoute interface{}

    // ISIS route type. The type is Ospfv3isisRoute.
    IsisRouteType interface{}

    // EIGRP route type. The type is Ospfv3EigrpRoute.
    EigrpRouteType interface{}

    // Preserve (Multi-Exit Discriminator) of BGP routes. The type is bool.
    PreserveMed interface{}

    // Preserve Metric and Metric Type ofBGP Default Route. The type is bool.
    BgpPreserveDefaultInfo interface{}

    // Use metric from RIB for redistributed routes. The type is bool.
    UseRibMetric interface{}
}

func (connectedOrStaticOrSubscriberOrMobile *Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes_Redistribute_ConnectedOrStaticOrSubscriberOrMobile) GetEntityData() *types.CommonEntityData {
    connectedOrStaticOrSubscriberOrMobile.EntityData.YFilter = connectedOrStaticOrSubscriberOrMobile.YFilter
    connectedOrStaticOrSubscriberOrMobile.EntityData.YangName = "connected-or-static-or-subscriber-or-mobile"
    connectedOrStaticOrSubscriberOrMobile.EntityData.BundleName = "cisco_ios_xr"
    connectedOrStaticOrSubscriberOrMobile.EntityData.ParentYangName = "redistribute"
    connectedOrStaticOrSubscriberOrMobile.EntityData.SegmentPath = "connected-or-static-or-subscriber-or-mobile"
    connectedOrStaticOrSubscriberOrMobile.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/redistributes/redistribute/" + connectedOrStaticOrSubscriberOrMobile.EntityData.SegmentPath
    connectedOrStaticOrSubscriberOrMobile.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    connectedOrStaticOrSubscriberOrMobile.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    connectedOrStaticOrSubscriberOrMobile.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    connectedOrStaticOrSubscriberOrMobile.EntityData.Children = types.NewOrderedMap()
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs = types.NewOrderedMap()
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs.Append("internal-route-type", types.YLeaf{"InternalRouteType", connectedOrStaticOrSubscriberOrMobile.InternalRouteType})
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs.Append("default-metric", types.YLeaf{"DefaultMetric", connectedOrStaticOrSubscriberOrMobile.DefaultMetric})
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", connectedOrStaticOrSubscriberOrMobile.MetricType})
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", connectedOrStaticOrSubscriberOrMobile.Tag})
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", connectedOrStaticOrSubscriberOrMobile.RoutePolicyName})
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs.Append("external-route-type", types.YLeaf{"ExternalRouteType", connectedOrStaticOrSubscriberOrMobile.ExternalRouteType})
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs.Append("nssa-external-route-type", types.YLeaf{"NssaExternalRouteType", connectedOrStaticOrSubscriberOrMobile.NssaExternalRouteType})
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs.Append("redistribute-route", types.YLeaf{"RedistributeRoute", connectedOrStaticOrSubscriberOrMobile.RedistributeRoute})
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs.Append("isis-route-type", types.YLeaf{"IsisRouteType", connectedOrStaticOrSubscriberOrMobile.IsisRouteType})
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs.Append("eigrp-route-type", types.YLeaf{"EigrpRouteType", connectedOrStaticOrSubscriberOrMobile.EigrpRouteType})
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs.Append("preserve-med", types.YLeaf{"PreserveMed", connectedOrStaticOrSubscriberOrMobile.PreserveMed})
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs.Append("bgp-preserve-default-info", types.YLeaf{"BgpPreserveDefaultInfo", connectedOrStaticOrSubscriberOrMobile.BgpPreserveDefaultInfo})
    connectedOrStaticOrSubscriberOrMobile.EntityData.Leafs.Append("use-rib-metric", types.YLeaf{"UseRibMetric", connectedOrStaticOrSubscriberOrMobile.UseRibMetric})

    connectedOrStaticOrSubscriberOrMobile.EntityData.YListKeys = []string {}

    return &(connectedOrStaticOrSubscriberOrMobile.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes_Redistribute_Bgp
// bgp
type Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes_Redistribute_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. 4-byte AS number in asdot (X.Y) format - first
    // half (X). The type is interface{} with range: 0..65535.
    AsXx interface{}

    // This attribute is a key. 4-byte AS number in asdot (X.Y) format - second
    // half (Y), or 2-byte AS number, or 4-byte AS number in asplain format. The
    // type is interface{} with range: 0..4294967295.
    AsYy interface{}

    // Redistribute OSPFv3 routes. The type is Ospfv3InternalRoute.
    InternalRouteType interface{}

    // OSPFv3 default metric. The type is interface{} with range: 0..16777214.
    DefaultMetric interface{}

    // OSPFv3 exterior metric type for redistributed routes. The type is
    // Ospfv3Metric.
    MetricType interface{}

    // Tag for routes redistributed into OSPFv3. The type is interface{} with
    // range: 0..4294967295.
    Tag interface{}

    // Route policy to redistribution. The type is string.
    RoutePolicyName interface{}

    // Redistribute OSPFv3 external routes. The type is Ospfv3ExternalRoute.
    ExternalRouteType interface{}

    // Redistribute OSPFv3 NSSA external routes. The type is
    // Ospfv3nssaExternalRoute.
    NssaExternalRouteType interface{}

    // Redistribution of OSPFv3 routes. The type is bool.
    RedistributeRoute interface{}

    // ISIS route type. The type is Ospfv3isisRoute.
    IsisRouteType interface{}

    // EIGRP route type. The type is Ospfv3EigrpRoute.
    EigrpRouteType interface{}

    // Preserve (Multi-Exit Discriminator) of BGP routes. The type is bool.
    PreserveMed interface{}

    // Preserve Metric and Metric Type ofBGP Default Route. The type is bool.
    BgpPreserveDefaultInfo interface{}

    // Use metric from RIB for redistributed routes. The type is bool.
    UseRibMetric interface{}
}

func (bgp *Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes_Redistribute_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "redistribute"
    bgp.EntityData.SegmentPath = "bgp" + types.AddKeyToken(bgp.AsXx, "as-xx") + types.AddKeyToken(bgp.AsYy, "as-yy")
    bgp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/redistributes/redistribute/" + bgp.EntityData.SegmentPath
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = types.NewOrderedMap()
    bgp.EntityData.Leafs = types.NewOrderedMap()
    bgp.EntityData.Leafs.Append("as-xx", types.YLeaf{"AsXx", bgp.AsXx})
    bgp.EntityData.Leafs.Append("as-yy", types.YLeaf{"AsYy", bgp.AsYy})
    bgp.EntityData.Leafs.Append("internal-route-type", types.YLeaf{"InternalRouteType", bgp.InternalRouteType})
    bgp.EntityData.Leafs.Append("default-metric", types.YLeaf{"DefaultMetric", bgp.DefaultMetric})
    bgp.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", bgp.MetricType})
    bgp.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", bgp.Tag})
    bgp.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", bgp.RoutePolicyName})
    bgp.EntityData.Leafs.Append("external-route-type", types.YLeaf{"ExternalRouteType", bgp.ExternalRouteType})
    bgp.EntityData.Leafs.Append("nssa-external-route-type", types.YLeaf{"NssaExternalRouteType", bgp.NssaExternalRouteType})
    bgp.EntityData.Leafs.Append("redistribute-route", types.YLeaf{"RedistributeRoute", bgp.RedistributeRoute})
    bgp.EntityData.Leafs.Append("isis-route-type", types.YLeaf{"IsisRouteType", bgp.IsisRouteType})
    bgp.EntityData.Leafs.Append("eigrp-route-type", types.YLeaf{"EigrpRouteType", bgp.EigrpRouteType})
    bgp.EntityData.Leafs.Append("preserve-med", types.YLeaf{"PreserveMed", bgp.PreserveMed})
    bgp.EntityData.Leafs.Append("bgp-preserve-default-info", types.YLeaf{"BgpPreserveDefaultInfo", bgp.BgpPreserveDefaultInfo})
    bgp.EntityData.Leafs.Append("use-rib-metric", types.YLeaf{"UseRibMetric", bgp.UseRibMetric})

    bgp.EntityData.YListKeys = []string {"AsXx", "AsYy"}

    return &(bgp.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes_Redistribute_Ospfv3OrIsisOrApplication
// ospfv3 or isis or application
type Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes_Redistribute_Ospfv3OrIsisOrApplication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. ISIS process name if protocol is ISIS, or OSPFv3
    // process name if protocol is OSPFv3. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ProcessName interface{}

    // Redistribute OSPFv3 routes. The type is Ospfv3InternalRoute.
    InternalRouteType interface{}

    // OSPFv3 default metric. The type is interface{} with range: 0..16777214.
    DefaultMetric interface{}

    // OSPFv3 exterior metric type for redistributed routes. The type is
    // Ospfv3Metric.
    MetricType interface{}

    // Tag for routes redistributed into OSPFv3. The type is interface{} with
    // range: 0..4294967295.
    Tag interface{}

    // Route policy to redistribution. The type is string.
    RoutePolicyName interface{}

    // Redistribute OSPFv3 external routes. The type is Ospfv3ExternalRoute.
    ExternalRouteType interface{}

    // Redistribute OSPFv3 NSSA external routes. The type is
    // Ospfv3nssaExternalRoute.
    NssaExternalRouteType interface{}

    // Redistribution of OSPFv3 routes. The type is bool.
    RedistributeRoute interface{}

    // ISIS route type. The type is Ospfv3isisRoute.
    IsisRouteType interface{}

    // EIGRP route type. The type is Ospfv3EigrpRoute.
    EigrpRouteType interface{}

    // Preserve (Multi-Exit Discriminator) of BGP routes. The type is bool.
    PreserveMed interface{}

    // Preserve Metric and Metric Type ofBGP Default Route. The type is bool.
    BgpPreserveDefaultInfo interface{}

    // Use metric from RIB for redistributed routes. The type is bool.
    UseRibMetric interface{}
}

func (ospfv3OrIsisOrApplication *Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes_Redistribute_Ospfv3OrIsisOrApplication) GetEntityData() *types.CommonEntityData {
    ospfv3OrIsisOrApplication.EntityData.YFilter = ospfv3OrIsisOrApplication.YFilter
    ospfv3OrIsisOrApplication.EntityData.YangName = "ospfv3-or-isis-or-application"
    ospfv3OrIsisOrApplication.EntityData.BundleName = "cisco_ios_xr"
    ospfv3OrIsisOrApplication.EntityData.ParentYangName = "redistribute"
    ospfv3OrIsisOrApplication.EntityData.SegmentPath = "ospfv3-or-isis-or-application" + types.AddKeyToken(ospfv3OrIsisOrApplication.ProcessName, "process-name")
    ospfv3OrIsisOrApplication.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/redistributes/redistribute/" + ospfv3OrIsisOrApplication.EntityData.SegmentPath
    ospfv3OrIsisOrApplication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv3OrIsisOrApplication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv3OrIsisOrApplication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv3OrIsisOrApplication.EntityData.Children = types.NewOrderedMap()
    ospfv3OrIsisOrApplication.EntityData.Leafs = types.NewOrderedMap()
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", ospfv3OrIsisOrApplication.ProcessName})
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("internal-route-type", types.YLeaf{"InternalRouteType", ospfv3OrIsisOrApplication.InternalRouteType})
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("default-metric", types.YLeaf{"DefaultMetric", ospfv3OrIsisOrApplication.DefaultMetric})
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", ospfv3OrIsisOrApplication.MetricType})
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", ospfv3OrIsisOrApplication.Tag})
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", ospfv3OrIsisOrApplication.RoutePolicyName})
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("external-route-type", types.YLeaf{"ExternalRouteType", ospfv3OrIsisOrApplication.ExternalRouteType})
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("nssa-external-route-type", types.YLeaf{"NssaExternalRouteType", ospfv3OrIsisOrApplication.NssaExternalRouteType})
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("redistribute-route", types.YLeaf{"RedistributeRoute", ospfv3OrIsisOrApplication.RedistributeRoute})
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("isis-route-type", types.YLeaf{"IsisRouteType", ospfv3OrIsisOrApplication.IsisRouteType})
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("eigrp-route-type", types.YLeaf{"EigrpRouteType", ospfv3OrIsisOrApplication.EigrpRouteType})
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("preserve-med", types.YLeaf{"PreserveMed", ospfv3OrIsisOrApplication.PreserveMed})
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("bgp-preserve-default-info", types.YLeaf{"BgpPreserveDefaultInfo", ospfv3OrIsisOrApplication.BgpPreserveDefaultInfo})
    ospfv3OrIsisOrApplication.EntityData.Leafs.Append("use-rib-metric", types.YLeaf{"UseRibMetric", ospfv3OrIsisOrApplication.UseRibMetric})

    ospfv3OrIsisOrApplication.EntityData.YListKeys = []string {"ProcessName"}

    return &(ospfv3OrIsisOrApplication.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes_Redistribute_Eigrp
// eigrp
type Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes_Redistribute_Eigrp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. 4-byte AS number in asdot (X.Y) format - first
    // half (X). The type is interface{} with range: 0..65535.
    AsXx interface{}

    // Redistribute OSPFv3 routes. The type is Ospfv3InternalRoute.
    InternalRouteType interface{}

    // OSPFv3 default metric. The type is interface{} with range: 0..16777214.
    DefaultMetric interface{}

    // OSPFv3 exterior metric type for redistributed routes. The type is
    // Ospfv3Metric.
    MetricType interface{}

    // Tag for routes redistributed into OSPFv3. The type is interface{} with
    // range: 0..4294967295.
    Tag interface{}

    // Route policy to redistribution. The type is string.
    RoutePolicyName interface{}

    // Redistribute OSPFv3 external routes. The type is Ospfv3ExternalRoute.
    ExternalRouteType interface{}

    // Redistribute OSPFv3 NSSA external routes. The type is
    // Ospfv3nssaExternalRoute.
    NssaExternalRouteType interface{}

    // Redistribution of OSPFv3 routes. The type is bool.
    RedistributeRoute interface{}

    // ISIS route type. The type is Ospfv3isisRoute.
    IsisRouteType interface{}

    // EIGRP route type. The type is Ospfv3EigrpRoute.
    EigrpRouteType interface{}

    // Preserve (Multi-Exit Discriminator) of BGP routes. The type is bool.
    PreserveMed interface{}

    // Preserve Metric and Metric Type ofBGP Default Route. The type is bool.
    BgpPreserveDefaultInfo interface{}

    // Use metric from RIB for redistributed routes. The type is bool.
    UseRibMetric interface{}
}

func (eigrp *Ospfv3_Processes_Process_Vrfs_Vrf_Redistributes_Redistribute_Eigrp) GetEntityData() *types.CommonEntityData {
    eigrp.EntityData.YFilter = eigrp.YFilter
    eigrp.EntityData.YangName = "eigrp"
    eigrp.EntityData.BundleName = "cisco_ios_xr"
    eigrp.EntityData.ParentYangName = "redistribute"
    eigrp.EntityData.SegmentPath = "eigrp" + types.AddKeyToken(eigrp.AsXx, "as-xx")
    eigrp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/redistributes/redistribute/" + eigrp.EntityData.SegmentPath
    eigrp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eigrp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eigrp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eigrp.EntityData.Children = types.NewOrderedMap()
    eigrp.EntityData.Leafs = types.NewOrderedMap()
    eigrp.EntityData.Leafs.Append("as-xx", types.YLeaf{"AsXx", eigrp.AsXx})
    eigrp.EntityData.Leafs.Append("internal-route-type", types.YLeaf{"InternalRouteType", eigrp.InternalRouteType})
    eigrp.EntityData.Leafs.Append("default-metric", types.YLeaf{"DefaultMetric", eigrp.DefaultMetric})
    eigrp.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", eigrp.MetricType})
    eigrp.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", eigrp.Tag})
    eigrp.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", eigrp.RoutePolicyName})
    eigrp.EntityData.Leafs.Append("external-route-type", types.YLeaf{"ExternalRouteType", eigrp.ExternalRouteType})
    eigrp.EntityData.Leafs.Append("nssa-external-route-type", types.YLeaf{"NssaExternalRouteType", eigrp.NssaExternalRouteType})
    eigrp.EntityData.Leafs.Append("redistribute-route", types.YLeaf{"RedistributeRoute", eigrp.RedistributeRoute})
    eigrp.EntityData.Leafs.Append("isis-route-type", types.YLeaf{"IsisRouteType", eigrp.IsisRouteType})
    eigrp.EntityData.Leafs.Append("eigrp-route-type", types.YLeaf{"EigrpRouteType", eigrp.EigrpRouteType})
    eigrp.EntityData.Leafs.Append("preserve-med", types.YLeaf{"PreserveMed", eigrp.PreserveMed})
    eigrp.EntityData.Leafs.Append("bgp-preserve-default-info", types.YLeaf{"BgpPreserveDefaultInfo", eigrp.BgpPreserveDefaultInfo})
    eigrp.EntityData.Leafs.Append("use-rib-metric", types.YLeaf{"UseRibMetric", eigrp.UseRibMetric})

    eigrp.EntityData.YListKeys = []string {"AsXx"}

    return &(eigrp.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_Ignore
// Do not complain about a specified event
type Ospfv3_Processes_Process_Vrfs_Vrf_Ignore struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Do not complain upon receiving LSA of the specified type.
    Lsa Ospfv3_Processes_Process_Vrfs_Vrf_Ignore_Lsa
}

func (ignore *Ospfv3_Processes_Process_Vrfs_Vrf_Ignore) GetEntityData() *types.CommonEntityData {
    ignore.EntityData.YFilter = ignore.YFilter
    ignore.EntityData.YangName = "ignore"
    ignore.EntityData.BundleName = "cisco_ios_xr"
    ignore.EntityData.ParentYangName = "vrf"
    ignore.EntityData.SegmentPath = "ignore"
    ignore.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/" + ignore.EntityData.SegmentPath
    ignore.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ignore.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ignore.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ignore.EntityData.Children = types.NewOrderedMap()
    ignore.EntityData.Children.Append("lsa", types.YChild{"Lsa", &ignore.Lsa})
    ignore.EntityData.Leafs = types.NewOrderedMap()

    ignore.EntityData.YListKeys = []string {}

    return &(ignore.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_Ignore_Lsa
// Do not complain upon receiving LSA of the
// specified type
type Ospfv3_Processes_Process_Vrfs_Vrf_Ignore_Lsa struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable ignore of MOSPF type 6 LSA. The type is interface{}.
    Mospf interface{}
}

func (lsa *Ospfv3_Processes_Process_Vrfs_Vrf_Ignore_Lsa) GetEntityData() *types.CommonEntityData {
    lsa.EntityData.YFilter = lsa.YFilter
    lsa.EntityData.YangName = "lsa"
    lsa.EntityData.BundleName = "cisco_ios_xr"
    lsa.EntityData.ParentYangName = "ignore"
    lsa.EntityData.SegmentPath = "lsa"
    lsa.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/ignore/" + lsa.EntityData.SegmentPath
    lsa.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    lsa.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    lsa.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    lsa.EntityData.Children = types.NewOrderedMap()
    lsa.EntityData.Leafs = types.NewOrderedMap()
    lsa.EntityData.Leafs.Append("mospf", types.YLeaf{"Mospf", lsa.Mospf})

    lsa.EntityData.YListKeys = []string {}

    return &(lsa.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut
// Filter prefixes from RIB 
type Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter generated type-5 LSAs.
    DistributeOuts Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut_DistributeOuts
}

func (distributeListOut *Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut) GetEntityData() *types.CommonEntityData {
    distributeListOut.EntityData.YFilter = distributeListOut.YFilter
    distributeListOut.EntityData.YangName = "distribute-list-out"
    distributeListOut.EntityData.BundleName = "cisco_ios_xr"
    distributeListOut.EntityData.ParentYangName = "vrf"
    distributeListOut.EntityData.SegmentPath = "distribute-list-out"
    distributeListOut.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/" + distributeListOut.EntityData.SegmentPath
    distributeListOut.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    distributeListOut.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    distributeListOut.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    distributeListOut.EntityData.Children = types.NewOrderedMap()
    distributeListOut.EntityData.Children.Append("distribute-outs", types.YChild{"DistributeOuts", &distributeListOut.DistributeOuts})
    distributeListOut.EntityData.Leafs = types.NewOrderedMap()

    distributeListOut.EntityData.YListKeys = []string {}

    return &(distributeListOut.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut_DistributeOuts
// Filter generated type-5 LSAs
type Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut_DistributeOuts struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter generated type-5 LSAs. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut_DistributeOuts_DistributeOut.
    DistributeOut []*Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut_DistributeOuts_DistributeOut
}

func (distributeOuts *Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut_DistributeOuts) GetEntityData() *types.CommonEntityData {
    distributeOuts.EntityData.YFilter = distributeOuts.YFilter
    distributeOuts.EntityData.YangName = "distribute-outs"
    distributeOuts.EntityData.BundleName = "cisco_ios_xr"
    distributeOuts.EntityData.ParentYangName = "distribute-list-out"
    distributeOuts.EntityData.SegmentPath = "distribute-outs"
    distributeOuts.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/distribute-list-out/" + distributeOuts.EntityData.SegmentPath
    distributeOuts.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    distributeOuts.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    distributeOuts.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    distributeOuts.EntityData.Children = types.NewOrderedMap()
    distributeOuts.EntityData.Children.Append("distribute-out", types.YChild{"DistributeOut", nil})
    for i := range distributeOuts.DistributeOut {
        distributeOuts.EntityData.Children.Append(types.GetSegmentPath(distributeOuts.DistributeOut[i]), types.YChild{"DistributeOut", distributeOuts.DistributeOut[i]})
    }
    distributeOuts.EntityData.Leafs = types.NewOrderedMap()

    distributeOuts.EntityData.YListKeys = []string {}

    return &(distributeOuts.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut_DistributeOuts_DistributeOut
// Filter generated type-5 LSAs
type Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut_DistributeOuts_DistributeOut struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. none. The type is Ospfv3Protocol.
    ProtocolName interface{}

    // Prefix-list name. The type is string.
    AllOrConnectedOrStaticPrefixList interface{}

    // bgp. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut_DistributeOuts_DistributeOut_Bgp.
    Bgp []*Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut_DistributeOuts_DistributeOut_Bgp

    // ospfv3 or isis. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut_DistributeOuts_DistributeOut_Ospfv3OrIsis.
    Ospfv3OrIsis []*Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut_DistributeOuts_DistributeOut_Ospfv3OrIsis

    // eigrp. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut_DistributeOuts_DistributeOut_Eigrp.
    Eigrp []*Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut_DistributeOuts_DistributeOut_Eigrp
}

func (distributeOut *Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut_DistributeOuts_DistributeOut) GetEntityData() *types.CommonEntityData {
    distributeOut.EntityData.YFilter = distributeOut.YFilter
    distributeOut.EntityData.YangName = "distribute-out"
    distributeOut.EntityData.BundleName = "cisco_ios_xr"
    distributeOut.EntityData.ParentYangName = "distribute-outs"
    distributeOut.EntityData.SegmentPath = "distribute-out" + types.AddKeyToken(distributeOut.ProtocolName, "protocol-name")
    distributeOut.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/distribute-list-out/distribute-outs/" + distributeOut.EntityData.SegmentPath
    distributeOut.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    distributeOut.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    distributeOut.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    distributeOut.EntityData.Children = types.NewOrderedMap()
    distributeOut.EntityData.Children.Append("bgp", types.YChild{"Bgp", nil})
    for i := range distributeOut.Bgp {
        distributeOut.EntityData.Children.Append(types.GetSegmentPath(distributeOut.Bgp[i]), types.YChild{"Bgp", distributeOut.Bgp[i]})
    }
    distributeOut.EntityData.Children.Append("ospfv3-or-isis", types.YChild{"Ospfv3OrIsis", nil})
    for i := range distributeOut.Ospfv3OrIsis {
        distributeOut.EntityData.Children.Append(types.GetSegmentPath(distributeOut.Ospfv3OrIsis[i]), types.YChild{"Ospfv3OrIsis", distributeOut.Ospfv3OrIsis[i]})
    }
    distributeOut.EntityData.Children.Append("eigrp", types.YChild{"Eigrp", nil})
    for i := range distributeOut.Eigrp {
        distributeOut.EntityData.Children.Append(types.GetSegmentPath(distributeOut.Eigrp[i]), types.YChild{"Eigrp", distributeOut.Eigrp[i]})
    }
    distributeOut.EntityData.Leafs = types.NewOrderedMap()
    distributeOut.EntityData.Leafs.Append("protocol-name", types.YLeaf{"ProtocolName", distributeOut.ProtocolName})
    distributeOut.EntityData.Leafs.Append("all-or-connected-or-static-prefix-list", types.YLeaf{"AllOrConnectedOrStaticPrefixList", distributeOut.AllOrConnectedOrStaticPrefixList})

    distributeOut.EntityData.YListKeys = []string {"ProtocolName"}

    return &(distributeOut.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut_DistributeOuts_DistributeOut_Bgp
// bgp
type Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut_DistributeOuts_DistributeOut_Bgp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. 4-byte AS number in asdot (X.Y) format - first
    // half (X). The type is interface{} with range: 1..65535.
    AsXx interface{}

    // This attribute is a key. 4-byte AS number in asdot (X.Y) format - second
    // half (Y), or 2-byte AS number, or 4-byte AS number in asplain format. The
    // type is interface{} with range: 0..4294967295.
    AsYy interface{}

    // Prefix-list name. The type is string.
    PrefixList interface{}
}

func (bgp *Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut_DistributeOuts_DistributeOut_Bgp) GetEntityData() *types.CommonEntityData {
    bgp.EntityData.YFilter = bgp.YFilter
    bgp.EntityData.YangName = "bgp"
    bgp.EntityData.BundleName = "cisco_ios_xr"
    bgp.EntityData.ParentYangName = "distribute-out"
    bgp.EntityData.SegmentPath = "bgp" + types.AddKeyToken(bgp.AsXx, "as-xx") + types.AddKeyToken(bgp.AsYy, "as-yy")
    bgp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/distribute-list-out/distribute-outs/distribute-out/" + bgp.EntityData.SegmentPath
    bgp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bgp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bgp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bgp.EntityData.Children = types.NewOrderedMap()
    bgp.EntityData.Leafs = types.NewOrderedMap()
    bgp.EntityData.Leafs.Append("as-xx", types.YLeaf{"AsXx", bgp.AsXx})
    bgp.EntityData.Leafs.Append("as-yy", types.YLeaf{"AsYy", bgp.AsYy})
    bgp.EntityData.Leafs.Append("prefix-list", types.YLeaf{"PrefixList", bgp.PrefixList})

    bgp.EntityData.YListKeys = []string {"AsXx", "AsYy"}

    return &(bgp.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut_DistributeOuts_DistributeOut_Ospfv3OrIsis
// ospfv3 or isis
type Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut_DistributeOuts_DistributeOut_Ospfv3OrIsis struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. If ISIS or OSPFv3, specify the instance name. The
    // type is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    ProcessName interface{}

    // Prefix-list name. The type is string.
    PrefixList interface{}
}

func (ospfv3OrIsis *Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut_DistributeOuts_DistributeOut_Ospfv3OrIsis) GetEntityData() *types.CommonEntityData {
    ospfv3OrIsis.EntityData.YFilter = ospfv3OrIsis.YFilter
    ospfv3OrIsis.EntityData.YangName = "ospfv3-or-isis"
    ospfv3OrIsis.EntityData.BundleName = "cisco_ios_xr"
    ospfv3OrIsis.EntityData.ParentYangName = "distribute-out"
    ospfv3OrIsis.EntityData.SegmentPath = "ospfv3-or-isis" + types.AddKeyToken(ospfv3OrIsis.ProcessName, "process-name")
    ospfv3OrIsis.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/distribute-list-out/distribute-outs/distribute-out/" + ospfv3OrIsis.EntityData.SegmentPath
    ospfv3OrIsis.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ospfv3OrIsis.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ospfv3OrIsis.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ospfv3OrIsis.EntityData.Children = types.NewOrderedMap()
    ospfv3OrIsis.EntityData.Leafs = types.NewOrderedMap()
    ospfv3OrIsis.EntityData.Leafs.Append("process-name", types.YLeaf{"ProcessName", ospfv3OrIsis.ProcessName})
    ospfv3OrIsis.EntityData.Leafs.Append("prefix-list", types.YLeaf{"PrefixList", ospfv3OrIsis.PrefixList})

    ospfv3OrIsis.EntityData.YListKeys = []string {"ProcessName"}

    return &(ospfv3OrIsis.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut_DistributeOuts_DistributeOut_Eigrp
// eigrp
type Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut_DistributeOuts_DistributeOut_Eigrp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. 4-byte AS number in asdot (X.Y) format - first
    // half (X). The type is interface{} with range: 1..65535.
    AsXx interface{}

    // Prefix-list name. The type is string.
    PrefixList interface{}
}

func (eigrp *Ospfv3_Processes_Process_Vrfs_Vrf_DistributeListOut_DistributeOuts_DistributeOut_Eigrp) GetEntityData() *types.CommonEntityData {
    eigrp.EntityData.YFilter = eigrp.YFilter
    eigrp.EntityData.YangName = "eigrp"
    eigrp.EntityData.BundleName = "cisco_ios_xr"
    eigrp.EntityData.ParentYangName = "distribute-out"
    eigrp.EntityData.SegmentPath = "eigrp" + types.AddKeyToken(eigrp.AsXx, "as-xx")
    eigrp.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/distribute-list-out/distribute-outs/distribute-out/" + eigrp.EntityData.SegmentPath
    eigrp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    eigrp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    eigrp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    eigrp.EntityData.Children = types.NewOrderedMap()
    eigrp.EntityData.Leafs = types.NewOrderedMap()
    eigrp.EntityData.Leafs.Append("as-xx", types.YLeaf{"AsXx", eigrp.AsXx})
    eigrp.EntityData.Leafs.Append("prefix-list", types.YLeaf{"PrefixList", eigrp.PrefixList})

    eigrp.EntityData.YListKeys = []string {"AsXx"}

    return &(eigrp.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_DistributeList
// Filter prefixes to/from RIB
type Ospfv3_Processes_Process_Vrfs_Vrf_DistributeList struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter prefixes installed to RIB.
    In Ospfv3_Processes_Process_Vrfs_Vrf_DistributeList_In
}

func (distributeList *Ospfv3_Processes_Process_Vrfs_Vrf_DistributeList) GetEntityData() *types.CommonEntityData {
    distributeList.EntityData.YFilter = distributeList.YFilter
    distributeList.EntityData.YangName = "distribute-list"
    distributeList.EntityData.BundleName = "cisco_ios_xr"
    distributeList.EntityData.ParentYangName = "vrf"
    distributeList.EntityData.SegmentPath = "distribute-list"
    distributeList.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/" + distributeList.EntityData.SegmentPath
    distributeList.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    distributeList.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    distributeList.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    distributeList.EntityData.Children = types.NewOrderedMap()
    distributeList.EntityData.Children.Append("in", types.YChild{"In", &distributeList.In})
    distributeList.EntityData.Leafs = types.NewOrderedMap()

    distributeList.EntityData.YListKeys = []string {}

    return &(distributeList.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_DistributeList_In
// Filter prefixes installed to RIB
type Ospfv3_Processes_Process_Vrfs_Vrf_DistributeList_In struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Filter prefixes based on an IPv6 prefix-list. The type is string.
    PrefixList interface{}
}

func (in *Ospfv3_Processes_Process_Vrfs_Vrf_DistributeList_In) GetEntityData() *types.CommonEntityData {
    in.EntityData.YFilter = in.YFilter
    in.EntityData.YangName = "in"
    in.EntityData.BundleName = "cisco_ios_xr"
    in.EntityData.ParentYangName = "distribute-list"
    in.EntityData.SegmentPath = "in"
    in.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/distribute-list/" + in.EntityData.SegmentPath
    in.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    in.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    in.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    in.EntityData.Children = types.NewOrderedMap()
    in.EntityData.Leafs = types.NewOrderedMap()
    in.EntityData.Leafs.Append("prefix-list", types.YLeaf{"PrefixList", in.PrefixList})

    in.EntityData.YListKeys = []string {}

    return &(in.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter
// Stub router configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Stub router R-bit configuration.
    Rbit Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter_Rbit

    // Stub router V6-bit configuration.
    V6bit Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter_V6bit

    // Stub router max-metric configuration.
    MaxMetric Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter_MaxMetric
}

func (stubRouter *Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter) GetEntityData() *types.CommonEntityData {
    stubRouter.EntityData.YFilter = stubRouter.YFilter
    stubRouter.EntityData.YangName = "stub-router"
    stubRouter.EntityData.BundleName = "cisco_ios_xr"
    stubRouter.EntityData.ParentYangName = "vrf"
    stubRouter.EntityData.SegmentPath = "stub-router"
    stubRouter.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/" + stubRouter.EntityData.SegmentPath
    stubRouter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    stubRouter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    stubRouter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    stubRouter.EntityData.Children = types.NewOrderedMap()
    stubRouter.EntityData.Children.Append("rbit", types.YChild{"Rbit", &stubRouter.Rbit})
    stubRouter.EntityData.Children.Append("v6bit", types.YChild{"V6bit", &stubRouter.V6bit})
    stubRouter.EntityData.Children.Append("max-metric", types.YChild{"MaxMetric", &stubRouter.MaxMetric})
    stubRouter.EntityData.Leafs = types.NewOrderedMap()

    stubRouter.EntityData.YListKeys = []string {}

    return &(stubRouter.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter_Rbit
// Stub router R-bit configuration
// This type is a presence type.
type Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter_Rbit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Time (in seconds) to stay in stub router operational state. The type is
    // interface{} with range: 5..86400. Units are second.
    OnSwitchover interface{}

    // Unconditionally enter stub router operational state. The type is
    // interface{}.
    Always interface{}

    // Advertise stub links with maximum metric in stub router mode. The type is
    // interface{}.
    IncludeStub interface{}

    // Time (in seconds) to stay in stub router operational state. The type is
    // interface{} with range: 5..86400. Units are second.
    OnProcMigration interface{}

    // Time (in seconds) to stay in stub router operational state. The type is
    // interface{} with range: 5..86400. Units are second.
    OnProcRestart interface{}

    // Enter stub router operational state on startup.
    OnStartup Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter_Rbit_OnStartup
}

func (rbit *Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter_Rbit) GetEntityData() *types.CommonEntityData {
    rbit.EntityData.YFilter = rbit.YFilter
    rbit.EntityData.YangName = "rbit"
    rbit.EntityData.BundleName = "cisco_ios_xr"
    rbit.EntityData.ParentYangName = "stub-router"
    rbit.EntityData.SegmentPath = "rbit"
    rbit.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/stub-router/" + rbit.EntityData.SegmentPath
    rbit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rbit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rbit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rbit.EntityData.Children = types.NewOrderedMap()
    rbit.EntityData.Children.Append("on-startup", types.YChild{"OnStartup", &rbit.OnStartup})
    rbit.EntityData.Leafs = types.NewOrderedMap()
    rbit.EntityData.Leafs.Append("on-switchover", types.YLeaf{"OnSwitchover", rbit.OnSwitchover})
    rbit.EntityData.Leafs.Append("always", types.YLeaf{"Always", rbit.Always})
    rbit.EntityData.Leafs.Append("include-stub", types.YLeaf{"IncludeStub", rbit.IncludeStub})
    rbit.EntityData.Leafs.Append("on-proc-migration", types.YLeaf{"OnProcMigration", rbit.OnProcMigration})
    rbit.EntityData.Leafs.Append("on-proc-restart", types.YLeaf{"OnProcRestart", rbit.OnProcRestart})

    rbit.EntityData.YListKeys = []string {}

    return &(rbit.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter_Rbit_OnStartup
// Enter stub router operational state on startup
type Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter_Rbit_OnStartup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Wait until BGP converges (only applicable to default VRF). The type is
    // bool. The default value is false.
    WaitForBgp interface{}

    // Time (in seconds) to stay in stub router operational state. The type is
    // interface{} with range: 5..86400. Units are second.
    WaitTime interface{}
}

func (onStartup *Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter_Rbit_OnStartup) GetEntityData() *types.CommonEntityData {
    onStartup.EntityData.YFilter = onStartup.YFilter
    onStartup.EntityData.YangName = "on-startup"
    onStartup.EntityData.BundleName = "cisco_ios_xr"
    onStartup.EntityData.ParentYangName = "rbit"
    onStartup.EntityData.SegmentPath = "on-startup"
    onStartup.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/stub-router/rbit/" + onStartup.EntityData.SegmentPath
    onStartup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    onStartup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    onStartup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    onStartup.EntityData.Children = types.NewOrderedMap()
    onStartup.EntityData.Leafs = types.NewOrderedMap()
    onStartup.EntityData.Leafs.Append("wait-for-bgp", types.YLeaf{"WaitForBgp", onStartup.WaitForBgp})
    onStartup.EntityData.Leafs.Append("wait-time", types.YLeaf{"WaitTime", onStartup.WaitTime})

    onStartup.EntityData.YListKeys = []string {}

    return &(onStartup.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter_V6bit
// Stub router V6-bit configuration
// This type is a presence type.
type Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter_V6bit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Time (in seconds) to stay in stub router operational state. The type is
    // interface{} with range: 5..86400. Units are second.
    OnSwitchover interface{}

    // Unconditionally enter stub router operational state. The type is
    // interface{}.
    Always interface{}

    // Time (in seconds) to stay in stub router operational state. The type is
    // interface{} with range: 5..86400. Units are second.
    OnProcMigration interface{}

    // Time (in seconds) to stay in stub router operational state. The type is
    // interface{} with range: 5..86400. Units are second.
    OnProcRestart interface{}

    // Enter stub router operational state on startup.
    OnStartup Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter_V6bit_OnStartup
}

func (v6bit *Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter_V6bit) GetEntityData() *types.CommonEntityData {
    v6bit.EntityData.YFilter = v6bit.YFilter
    v6bit.EntityData.YangName = "v6bit"
    v6bit.EntityData.BundleName = "cisco_ios_xr"
    v6bit.EntityData.ParentYangName = "stub-router"
    v6bit.EntityData.SegmentPath = "v6bit"
    v6bit.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/stub-router/" + v6bit.EntityData.SegmentPath
    v6bit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    v6bit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    v6bit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    v6bit.EntityData.Children = types.NewOrderedMap()
    v6bit.EntityData.Children.Append("on-startup", types.YChild{"OnStartup", &v6bit.OnStartup})
    v6bit.EntityData.Leafs = types.NewOrderedMap()
    v6bit.EntityData.Leafs.Append("on-switchover", types.YLeaf{"OnSwitchover", v6bit.OnSwitchover})
    v6bit.EntityData.Leafs.Append("always", types.YLeaf{"Always", v6bit.Always})
    v6bit.EntityData.Leafs.Append("on-proc-migration", types.YLeaf{"OnProcMigration", v6bit.OnProcMigration})
    v6bit.EntityData.Leafs.Append("on-proc-restart", types.YLeaf{"OnProcRestart", v6bit.OnProcRestart})

    v6bit.EntityData.YListKeys = []string {}

    return &(v6bit.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter_V6bit_OnStartup
// Enter stub router operational state on startup
type Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter_V6bit_OnStartup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Wait until BGP converges (only applicable to default VRF). The type is
    // bool. The default value is false.
    WaitForBgp interface{}

    // Time (in seconds) to stay in stub router operational state. The type is
    // interface{} with range: 5..86400. Units are second.
    WaitTime interface{}
}

func (onStartup *Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter_V6bit_OnStartup) GetEntityData() *types.CommonEntityData {
    onStartup.EntityData.YFilter = onStartup.YFilter
    onStartup.EntityData.YangName = "on-startup"
    onStartup.EntityData.BundleName = "cisco_ios_xr"
    onStartup.EntityData.ParentYangName = "v6bit"
    onStartup.EntityData.SegmentPath = "on-startup"
    onStartup.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/stub-router/v6bit/" + onStartup.EntityData.SegmentPath
    onStartup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    onStartup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    onStartup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    onStartup.EntityData.Children = types.NewOrderedMap()
    onStartup.EntityData.Leafs = types.NewOrderedMap()
    onStartup.EntityData.Leafs.Append("wait-for-bgp", types.YLeaf{"WaitForBgp", onStartup.WaitForBgp})
    onStartup.EntityData.Leafs.Append("wait-time", types.YLeaf{"WaitTime", onStartup.WaitTime})

    onStartup.EntityData.YListKeys = []string {}

    return &(onStartup.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter_MaxMetric
// Stub router max-metric configuration
// This type is a presence type.
type Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter_MaxMetric struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Advertise external LSAs with modified metric in stub router mode. The type
    // is interface{} with range: 1..16777214. The default value is 16711680.
    ExternalLsa interface{}

    // Advertise summary LSAs with modified metric in stub router mode. The type
    // is interface{} with range: 1..16777214. The default value is 16711680.
    SummaryLsa interface{}

    // Time (in seconds) to stay in stub router operational state. The type is
    // interface{} with range: 5..86400. Units are second.
    OnSwitchover interface{}

    // Unconditionally enter stub router operational state. The type is
    // interface{}.
    Always interface{}

    // Advertise stub links with maximum metric in stub router mode. The type is
    // interface{}.
    IncludeStub interface{}

    // Time (in seconds) to stay in stub router operational state. The type is
    // interface{} with range: 5..86400. Units are second.
    OnProcMigration interface{}

    // Time (in seconds) to stay in stub router operational state. The type is
    // interface{} with range: 5..86400. Units are second.
    OnProcRestart interface{}

    // Enter stub router operational state on startup.
    OnStartup Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter_MaxMetric_OnStartup
}

func (maxMetric *Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter_MaxMetric) GetEntityData() *types.CommonEntityData {
    maxMetric.EntityData.YFilter = maxMetric.YFilter
    maxMetric.EntityData.YangName = "max-metric"
    maxMetric.EntityData.BundleName = "cisco_ios_xr"
    maxMetric.EntityData.ParentYangName = "stub-router"
    maxMetric.EntityData.SegmentPath = "max-metric"
    maxMetric.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/stub-router/" + maxMetric.EntityData.SegmentPath
    maxMetric.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    maxMetric.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    maxMetric.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    maxMetric.EntityData.Children = types.NewOrderedMap()
    maxMetric.EntityData.Children.Append("on-startup", types.YChild{"OnStartup", &maxMetric.OnStartup})
    maxMetric.EntityData.Leafs = types.NewOrderedMap()
    maxMetric.EntityData.Leafs.Append("external-lsa", types.YLeaf{"ExternalLsa", maxMetric.ExternalLsa})
    maxMetric.EntityData.Leafs.Append("summary-lsa", types.YLeaf{"SummaryLsa", maxMetric.SummaryLsa})
    maxMetric.EntityData.Leafs.Append("on-switchover", types.YLeaf{"OnSwitchover", maxMetric.OnSwitchover})
    maxMetric.EntityData.Leafs.Append("always", types.YLeaf{"Always", maxMetric.Always})
    maxMetric.EntityData.Leafs.Append("include-stub", types.YLeaf{"IncludeStub", maxMetric.IncludeStub})
    maxMetric.EntityData.Leafs.Append("on-proc-migration", types.YLeaf{"OnProcMigration", maxMetric.OnProcMigration})
    maxMetric.EntityData.Leafs.Append("on-proc-restart", types.YLeaf{"OnProcRestart", maxMetric.OnProcRestart})

    maxMetric.EntityData.YListKeys = []string {}

    return &(maxMetric.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter_MaxMetric_OnStartup
// Enter stub router operational state on startup
type Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter_MaxMetric_OnStartup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Wait until BGP converges (only applicable to default VRF). The type is
    // bool. The default value is false.
    WaitForBgp interface{}

    // Time (in seconds) to stay in stub router operational state. The type is
    // interface{} with range: 5..86400. Units are second.
    WaitTime interface{}
}

func (onStartup *Ospfv3_Processes_Process_Vrfs_Vrf_StubRouter_MaxMetric_OnStartup) GetEntityData() *types.CommonEntityData {
    onStartup.EntityData.YFilter = onStartup.YFilter
    onStartup.EntityData.YangName = "on-startup"
    onStartup.EntityData.BundleName = "cisco_ios_xr"
    onStartup.EntityData.ParentYangName = "max-metric"
    onStartup.EntityData.SegmentPath = "on-startup"
    onStartup.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/stub-router/max-metric/" + onStartup.EntityData.SegmentPath
    onStartup.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    onStartup.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    onStartup.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    onStartup.EntityData.Children = types.NewOrderedMap()
    onStartup.EntityData.Leafs = types.NewOrderedMap()
    onStartup.EntityData.Leafs.Append("wait-for-bgp", types.YLeaf{"WaitForBgp", onStartup.WaitForBgp})
    onStartup.EntityData.Leafs.Append("wait-time", types.YLeaf{"WaitTime", onStartup.WaitTime})

    onStartup.EntityData.YListKeys = []string {}

    return &(onStartup.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_Bfd
// Configure BFD parameters
type Ospfv3_Processes_Process_Vrfs_Vrf_Bfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Hello interval in milli-seconds. The type is interface{} with range:
    // 3..30000. Units are millisecond.
    Interval interface{}

    // Detect multiplier. The type is interface{} with range: 2..50.
    DetectionMultiplier interface{}

    // Enable or disable BFD fast detection. The type is Ospfv3bfdEnableMode.
    FastDetectMode interface{}
}

func (bfd *Ospfv3_Processes_Process_Vrfs_Vrf_Bfd) GetEntityData() *types.CommonEntityData {
    bfd.EntityData.YFilter = bfd.YFilter
    bfd.EntityData.YangName = "bfd"
    bfd.EntityData.BundleName = "cisco_ios_xr"
    bfd.EntityData.ParentYangName = "vrf"
    bfd.EntityData.SegmentPath = "bfd"
    bfd.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/" + bfd.EntityData.SegmentPath
    bfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bfd.EntityData.Children = types.NewOrderedMap()
    bfd.EntityData.Leafs = types.NewOrderedMap()
    bfd.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", bfd.Interval})
    bfd.EntityData.Leafs.Append("detection-multiplier", types.YLeaf{"DetectionMultiplier", bfd.DetectionMultiplier})
    bfd.EntityData.Leafs.Append("fast-detect-mode", types.YLeaf{"FastDetectMode", bfd.FastDetectMode})

    bfd.EntityData.YListKeys = []string {}

    return &(bfd.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_DatabaseFilter
// Database filter
type Ospfv3_Processes_Process_Vrfs_Vrf_DatabaseFilter struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // All.
    All Ospfv3_Processes_Process_Vrfs_Vrf_DatabaseFilter_All
}

func (databaseFilter *Ospfv3_Processes_Process_Vrfs_Vrf_DatabaseFilter) GetEntityData() *types.CommonEntityData {
    databaseFilter.EntityData.YFilter = databaseFilter.YFilter
    databaseFilter.EntityData.YangName = "database-filter"
    databaseFilter.EntityData.BundleName = "cisco_ios_xr"
    databaseFilter.EntityData.ParentYangName = "vrf"
    databaseFilter.EntityData.SegmentPath = "database-filter"
    databaseFilter.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/" + databaseFilter.EntityData.SegmentPath
    databaseFilter.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    databaseFilter.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    databaseFilter.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    databaseFilter.EntityData.Children = types.NewOrderedMap()
    databaseFilter.EntityData.Children.Append("all", types.YChild{"All", &databaseFilter.All})
    databaseFilter.EntityData.Leafs = types.NewOrderedMap()

    databaseFilter.EntityData.YListKeys = []string {}

    return &(databaseFilter.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_DatabaseFilter_All
// All
type Ospfv3_Processes_Process_Vrfs_Vrf_DatabaseFilter_All struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable out. The type is interface{}.
    Out interface{}
}

func (all *Ospfv3_Processes_Process_Vrfs_Vrf_DatabaseFilter_All) GetEntityData() *types.CommonEntityData {
    all.EntityData.YFilter = all.YFilter
    all.EntityData.YangName = "all"
    all.EntityData.BundleName = "cisco_ios_xr"
    all.EntityData.ParentYangName = "database-filter"
    all.EntityData.SegmentPath = "all"
    all.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/database-filter/" + all.EntityData.SegmentPath
    all.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    all.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    all.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    all.EntityData.Children = types.NewOrderedMap()
    all.EntityData.Leafs = types.NewOrderedMap()
    all.EntityData.Leafs.Append("out", types.YLeaf{"Out", all.Out})

    all.EntityData.YListKeys = []string {}

    return &(all.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_Capability
// OSPFv3 Capability
type Ospfv3_Processes_Process_Vrfs_Vrf_Capability struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // NSSA capability to prefer Type 7 over Type 5. The type is bool.
    Type7Prefer interface{}

    // Enable VRF Lite. The type is bool.
    VrfLite interface{}

    // Enable capability to translate LSAs with fwd addr. The type is bool.
    Type7TranslateZeroForwardingAddr interface{}
}

func (capability *Ospfv3_Processes_Process_Vrfs_Vrf_Capability) GetEntityData() *types.CommonEntityData {
    capability.EntityData.YFilter = capability.YFilter
    capability.EntityData.YangName = "capability"
    capability.EntityData.BundleName = "cisco_ios_xr"
    capability.EntityData.ParentYangName = "vrf"
    capability.EntityData.SegmentPath = "capability"
    capability.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/" + capability.EntityData.SegmentPath
    capability.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    capability.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    capability.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    capability.EntityData.Children = types.NewOrderedMap()
    capability.EntityData.Leafs = types.NewOrderedMap()
    capability.EntityData.Leafs.Append("type7-prefer", types.YLeaf{"Type7Prefer", capability.Type7Prefer})
    capability.EntityData.Leafs.Append("vrf-lite", types.YLeaf{"VrfLite", capability.VrfLite})
    capability.EntityData.Leafs.Append("type7-translate-zero-forwarding-addr", types.YLeaf{"Type7TranslateZeroForwardingAddr", capability.Type7TranslateZeroForwardingAddr})

    capability.EntityData.YListKeys = []string {}

    return &(capability.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_Authentication
// Authenticate OSPFv3 packets
type Ospfv3_Processes_Process_Vrfs_Vrf_Authentication struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec AH authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Use the MD5 or SHA1 algorithm. The type is Ospfv3Authentication.
    Algorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    Password interface{}
}

func (authentication *Ospfv3_Processes_Process_Vrfs_Vrf_Authentication) GetEntityData() *types.CommonEntityData {
    authentication.EntityData.YFilter = authentication.YFilter
    authentication.EntityData.YangName = "authentication"
    authentication.EntityData.BundleName = "cisco_ios_xr"
    authentication.EntityData.ParentYangName = "vrf"
    authentication.EntityData.SegmentPath = "authentication"
    authentication.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/" + authentication.EntityData.SegmentPath
    authentication.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    authentication.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    authentication.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    authentication.EntityData.Children = types.NewOrderedMap()
    authentication.EntityData.Leafs = types.NewOrderedMap()
    authentication.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", authentication.Enable})
    authentication.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", authentication.Spi})
    authentication.EntityData.Leafs.Append("algorithm", types.YLeaf{"Algorithm", authentication.Algorithm})
    authentication.EntityData.Leafs.Append("password", types.YLeaf{"Password", authentication.Password})

    authentication.EntityData.YListKeys = []string {}

    return &(authentication.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_GracefulRestart
// Graceful restart configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_GracefulRestart struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Minimum interval between graceful restarts (seconds). The type is
    // interface{} with range: 90..3600. Units are second.
    Interval interface{}

    // Terminate graceful restart helper mode if LSA changed. The type is
    // interface{}.
    StrictLsaChecking interface{}

    // Disable router's helper support. The type is interface{}.
    Helper interface{}

    // Enable graceful restart. The type is interface{}.
    Enable interface{}

    // Maximum route lifetime following restart (seconds). The type is interface{}
    // with range: 90..1800. Units are second.
    Lifetime interface{}
}

func (gracefulRestart *Ospfv3_Processes_Process_Vrfs_Vrf_GracefulRestart) GetEntityData() *types.CommonEntityData {
    gracefulRestart.EntityData.YFilter = gracefulRestart.YFilter
    gracefulRestart.EntityData.YangName = "graceful-restart"
    gracefulRestart.EntityData.BundleName = "cisco_ios_xr"
    gracefulRestart.EntityData.ParentYangName = "vrf"
    gracefulRestart.EntityData.SegmentPath = "graceful-restart"
    gracefulRestart.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/" + gracefulRestart.EntityData.SegmentPath
    gracefulRestart.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    gracefulRestart.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    gracefulRestart.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    gracefulRestart.EntityData.Children = types.NewOrderedMap()
    gracefulRestart.EntityData.Leafs = types.NewOrderedMap()
    gracefulRestart.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", gracefulRestart.Interval})
    gracefulRestart.EntityData.Leafs.Append("strict-lsa-checking", types.YLeaf{"StrictLsaChecking", gracefulRestart.StrictLsaChecking})
    gracefulRestart.EntityData.Leafs.Append("helper", types.YLeaf{"Helper", gracefulRestart.Helper})
    gracefulRestart.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", gracefulRestart.Enable})
    gracefulRestart.EntityData.Leafs.Append("lifetime", types.YLeaf{"Lifetime", gracefulRestart.Lifetime})

    gracefulRestart.EntityData.YListKeys = []string {}

    return &(gracefulRestart.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_DefaultInformation
// Control distribution of default information
type Ospfv3_Processes_Process_Vrfs_Vrf_DefaultInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Distribute a default route.
    Originate Ospfv3_Processes_Process_Vrfs_Vrf_DefaultInformation_Originate
}

func (defaultInformation *Ospfv3_Processes_Process_Vrfs_Vrf_DefaultInformation) GetEntityData() *types.CommonEntityData {
    defaultInformation.EntityData.YFilter = defaultInformation.YFilter
    defaultInformation.EntityData.YangName = "default-information"
    defaultInformation.EntityData.BundleName = "cisco_ios_xr"
    defaultInformation.EntityData.ParentYangName = "vrf"
    defaultInformation.EntityData.SegmentPath = "default-information"
    defaultInformation.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/" + defaultInformation.EntityData.SegmentPath
    defaultInformation.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    defaultInformation.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    defaultInformation.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    defaultInformation.EntityData.Children = types.NewOrderedMap()
    defaultInformation.EntityData.Children.Append("originate", types.YChild{"Originate", &defaultInformation.Originate})
    defaultInformation.EntityData.Leafs = types.NewOrderedMap()

    defaultInformation.EntityData.YListKeys = []string {}

    return &(defaultInformation.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_DefaultInformation_Originate
// Distribute a default route
// This type is a presence type.
type Ospfv3_Processes_Process_Vrfs_Vrf_DefaultInformation_Originate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Always advertise default route. The type is bool. This attribute is
    // mandatory.
    Always interface{}

    // OSPFv3 default metric. The type is interface{} with range: 0..16777214.
    Metric interface{}

    // OSPFv3 metric type for default routes. The type is interface{} with range:
    // 1..2.
    MetricType interface{}

    // Tag for default route. The type is interface{} with range: 0..4294967295.
    Tag interface{}

    // Route policy to default-information origination. The type is string.
    RoutePolicyName interface{}
}

func (originate *Ospfv3_Processes_Process_Vrfs_Vrf_DefaultInformation_Originate) GetEntityData() *types.CommonEntityData {
    originate.EntityData.YFilter = originate.YFilter
    originate.EntityData.YangName = "originate"
    originate.EntityData.BundleName = "cisco_ios_xr"
    originate.EntityData.ParentYangName = "default-information"
    originate.EntityData.SegmentPath = "originate"
    originate.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/default-information/" + originate.EntityData.SegmentPath
    originate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    originate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    originate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    originate.EntityData.Children = types.NewOrderedMap()
    originate.EntityData.Leafs = types.NewOrderedMap()
    originate.EntityData.Leafs.Append("always", types.YLeaf{"Always", originate.Always})
    originate.EntityData.Leafs.Append("metric", types.YLeaf{"Metric", originate.Metric})
    originate.EntityData.Leafs.Append("metric-type", types.YLeaf{"MetricType", originate.MetricType})
    originate.EntityData.Leafs.Append("tag", types.YLeaf{"Tag", originate.Tag})
    originate.EntityData.Leafs.Append("route-policy-name", types.YLeaf{"RoutePolicyName", originate.RoutePolicyName})

    originate.EntityData.YListKeys = []string {}

    return &(originate.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope
// Process scope configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Fast-reroute configuration.
    FastReroute Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute
}

func (processScope *Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope) GetEntityData() *types.CommonEntityData {
    processScope.EntityData.YFilter = processScope.YFilter
    processScope.EntityData.YangName = "process-scope"
    processScope.EntityData.BundleName = "cisco_ios_xr"
    processScope.EntityData.ParentYangName = "vrf"
    processScope.EntityData.SegmentPath = "process-scope"
    processScope.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/" + processScope.EntityData.SegmentPath
    processScope.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    processScope.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    processScope.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    processScope.EntityData.Children = types.NewOrderedMap()
    processScope.EntityData.Children.Append("fast-reroute", types.YChild{"FastReroute", &processScope.FastReroute})
    processScope.EntityData.Leafs = types.NewOrderedMap()

    processScope.EntityData.YListKeys = []string {}

    return &(processScope.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute
// Fast-reroute configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable/Disable Fast-reroute per-link or per-prefix. The type is
    // Ospfv3FastReroute.
    FastRerouteEnable interface{}

    // Fast-reroute per-link configuration.
    PerLink Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerLink

    // Fast-reroute per-link configuration.
    PerPrefix Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerPrefix
}

func (fastReroute *Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute) GetEntityData() *types.CommonEntityData {
    fastReroute.EntityData.YFilter = fastReroute.YFilter
    fastReroute.EntityData.YangName = "fast-reroute"
    fastReroute.EntityData.BundleName = "cisco_ios_xr"
    fastReroute.EntityData.ParentYangName = "process-scope"
    fastReroute.EntityData.SegmentPath = "fast-reroute"
    fastReroute.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/process-scope/" + fastReroute.EntityData.SegmentPath
    fastReroute.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    fastReroute.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    fastReroute.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    fastReroute.EntityData.Children = types.NewOrderedMap()
    fastReroute.EntityData.Children.Append("per-link", types.YChild{"PerLink", &fastReroute.PerLink})
    fastReroute.EntityData.Children.Append("per-prefix", types.YChild{"PerPrefix", &fastReroute.PerPrefix})
    fastReroute.EntityData.Leafs = types.NewOrderedMap()
    fastReroute.EntityData.Leafs.Append("fast-reroute-enable", types.YLeaf{"FastRerouteEnable", fastReroute.FastRerouteEnable})

    fastReroute.EntityData.YListKeys = []string {}

    return &(fastReroute.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerLink
// Fast-reroute per-link configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerLink struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Use only interfaces on the candidate list as a backup path. The type is
    // bool. The default value is false.
    FastRerouteUseCandidateOnly interface{}

    // Fast-reroute per-link/per-prefix candidate interface configuration.
    CandidateInterfaces Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerLink_CandidateInterfaces

    // Fast-reroute per-link/per-prefix exclude interface configuration.
    ExcludeInterfaces Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerLink_ExcludeInterfaces
}

func (perLink *Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerLink) GetEntityData() *types.CommonEntityData {
    perLink.EntityData.YFilter = perLink.YFilter
    perLink.EntityData.YangName = "per-link"
    perLink.EntityData.BundleName = "cisco_ios_xr"
    perLink.EntityData.ParentYangName = "fast-reroute"
    perLink.EntityData.SegmentPath = "per-link"
    perLink.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/process-scope/fast-reroute/" + perLink.EntityData.SegmentPath
    perLink.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    perLink.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    perLink.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    perLink.EntityData.Children = types.NewOrderedMap()
    perLink.EntityData.Children.Append("candidate-interfaces", types.YChild{"CandidateInterfaces", &perLink.CandidateInterfaces})
    perLink.EntityData.Children.Append("exclude-interfaces", types.YChild{"ExcludeInterfaces", &perLink.ExcludeInterfaces})
    perLink.EntityData.Leafs = types.NewOrderedMap()
    perLink.EntityData.Leafs.Append("fast-reroute-use-candidate-only", types.YLeaf{"FastRerouteUseCandidateOnly", perLink.FastRerouteUseCandidateOnly})

    perLink.EntityData.YListKeys = []string {}

    return &(perLink.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerLink_CandidateInterfaces
// Fast-reroute per-link/per-prefix candidate
// interface configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerLink_CandidateInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Candidate backup interface. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface.
    CandidateInterface []*Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface
}

func (candidateInterfaces *Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerLink_CandidateInterfaces) GetEntityData() *types.CommonEntityData {
    candidateInterfaces.EntityData.YFilter = candidateInterfaces.YFilter
    candidateInterfaces.EntityData.YangName = "candidate-interfaces"
    candidateInterfaces.EntityData.BundleName = "cisco_ios_xr"
    candidateInterfaces.EntityData.ParentYangName = "per-link"
    candidateInterfaces.EntityData.SegmentPath = "candidate-interfaces"
    candidateInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/process-scope/fast-reroute/per-link/" + candidateInterfaces.EntityData.SegmentPath
    candidateInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterfaces.EntityData.Children = types.NewOrderedMap()
    candidateInterfaces.EntityData.Children.Append("candidate-interface", types.YChild{"CandidateInterface", nil})
    for i := range candidateInterfaces.CandidateInterface {
        candidateInterfaces.EntityData.Children.Append(types.GetSegmentPath(candidateInterfaces.CandidateInterface[i]), types.YChild{"CandidateInterface", candidateInterfaces.CandidateInterface[i]})
    }
    candidateInterfaces.EntityData.Leafs = types.NewOrderedMap()

    candidateInterfaces.EntityData.YListKeys = []string {}

    return &(candidateInterfaces.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface
// Candidate backup interface
type Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (candidateInterface *Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerLink_CandidateInterfaces_CandidateInterface) GetEntityData() *types.CommonEntityData {
    candidateInterface.EntityData.YFilter = candidateInterface.YFilter
    candidateInterface.EntityData.YangName = "candidate-interface"
    candidateInterface.EntityData.BundleName = "cisco_ios_xr"
    candidateInterface.EntityData.ParentYangName = "candidate-interfaces"
    candidateInterface.EntityData.SegmentPath = "candidate-interface" + types.AddKeyToken(candidateInterface.InterfaceName, "interface-name")
    candidateInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/process-scope/fast-reroute/per-link/candidate-interfaces/" + candidateInterface.EntityData.SegmentPath
    candidateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterface.EntityData.Children = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", candidateInterface.InterfaceName})

    candidateInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(candidateInterface.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerLink_ExcludeInterfaces
// Fast-reroute per-link/per-prefix exclude
// interface configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerLink_ExcludeInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exclude an interface from becoming a backup. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface.
    ExcludeInterface []*Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface
}

func (excludeInterfaces *Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerLink_ExcludeInterfaces) GetEntityData() *types.CommonEntityData {
    excludeInterfaces.EntityData.YFilter = excludeInterfaces.YFilter
    excludeInterfaces.EntityData.YangName = "exclude-interfaces"
    excludeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    excludeInterfaces.EntityData.ParentYangName = "per-link"
    excludeInterfaces.EntityData.SegmentPath = "exclude-interfaces"
    excludeInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/process-scope/fast-reroute/per-link/" + excludeInterfaces.EntityData.SegmentPath
    excludeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterfaces.EntityData.Children = types.NewOrderedMap()
    excludeInterfaces.EntityData.Children.Append("exclude-interface", types.YChild{"ExcludeInterface", nil})
    for i := range excludeInterfaces.ExcludeInterface {
        excludeInterfaces.EntityData.Children.Append(types.GetSegmentPath(excludeInterfaces.ExcludeInterface[i]), types.YChild{"ExcludeInterface", excludeInterfaces.ExcludeInterface[i]})
    }
    excludeInterfaces.EntityData.Leafs = types.NewOrderedMap()

    excludeInterfaces.EntityData.YListKeys = []string {}

    return &(excludeInterfaces.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface
// Exclude an interface from becoming a backup
type Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (excludeInterface *Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerLink_ExcludeInterfaces_ExcludeInterface) GetEntityData() *types.CommonEntityData {
    excludeInterface.EntityData.YFilter = excludeInterface.YFilter
    excludeInterface.EntityData.YangName = "exclude-interface"
    excludeInterface.EntityData.BundleName = "cisco_ios_xr"
    excludeInterface.EntityData.ParentYangName = "exclude-interfaces"
    excludeInterface.EntityData.SegmentPath = "exclude-interface" + types.AddKeyToken(excludeInterface.InterfaceName, "interface-name")
    excludeInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/process-scope/fast-reroute/per-link/exclude-interfaces/" + excludeInterface.EntityData.SegmentPath
    excludeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterface.EntityData.Children = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", excludeInterface.InterfaceName})

    excludeInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(excludeInterface.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerPrefix
// Fast-reroute per-link configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerPrefix struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Use only interfaces on the candidate list as a backup path. The type is
    // bool. The default value is false.
    FastRerouteUseCandidateOnly interface{}

    // Fast-reroute per-link/per-prefix candidate interface configuration.
    CandidateInterfaces Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerPrefix_CandidateInterfaces

    // Fast-reroute per-link/per-prefix exclude interface configuration.
    ExcludeInterfaces Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerPrefix_ExcludeInterfaces
}

func (perPrefix *Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerPrefix) GetEntityData() *types.CommonEntityData {
    perPrefix.EntityData.YFilter = perPrefix.YFilter
    perPrefix.EntityData.YangName = "per-prefix"
    perPrefix.EntityData.BundleName = "cisco_ios_xr"
    perPrefix.EntityData.ParentYangName = "fast-reroute"
    perPrefix.EntityData.SegmentPath = "per-prefix"
    perPrefix.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/process-scope/fast-reroute/" + perPrefix.EntityData.SegmentPath
    perPrefix.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    perPrefix.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    perPrefix.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    perPrefix.EntityData.Children = types.NewOrderedMap()
    perPrefix.EntityData.Children.Append("candidate-interfaces", types.YChild{"CandidateInterfaces", &perPrefix.CandidateInterfaces})
    perPrefix.EntityData.Children.Append("exclude-interfaces", types.YChild{"ExcludeInterfaces", &perPrefix.ExcludeInterfaces})
    perPrefix.EntityData.Leafs = types.NewOrderedMap()
    perPrefix.EntityData.Leafs.Append("fast-reroute-use-candidate-only", types.YLeaf{"FastRerouteUseCandidateOnly", perPrefix.FastRerouteUseCandidateOnly})

    perPrefix.EntityData.YListKeys = []string {}

    return &(perPrefix.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerPrefix_CandidateInterfaces
// Fast-reroute per-link/per-prefix candidate
// interface configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerPrefix_CandidateInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Candidate backup interface. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface.
    CandidateInterface []*Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface
}

func (candidateInterfaces *Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerPrefix_CandidateInterfaces) GetEntityData() *types.CommonEntityData {
    candidateInterfaces.EntityData.YFilter = candidateInterfaces.YFilter
    candidateInterfaces.EntityData.YangName = "candidate-interfaces"
    candidateInterfaces.EntityData.BundleName = "cisco_ios_xr"
    candidateInterfaces.EntityData.ParentYangName = "per-prefix"
    candidateInterfaces.EntityData.SegmentPath = "candidate-interfaces"
    candidateInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/process-scope/fast-reroute/per-prefix/" + candidateInterfaces.EntityData.SegmentPath
    candidateInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterfaces.EntityData.Children = types.NewOrderedMap()
    candidateInterfaces.EntityData.Children.Append("candidate-interface", types.YChild{"CandidateInterface", nil})
    for i := range candidateInterfaces.CandidateInterface {
        candidateInterfaces.EntityData.Children.Append(types.GetSegmentPath(candidateInterfaces.CandidateInterface[i]), types.YChild{"CandidateInterface", candidateInterfaces.CandidateInterface[i]})
    }
    candidateInterfaces.EntityData.Leafs = types.NewOrderedMap()

    candidateInterfaces.EntityData.YListKeys = []string {}

    return &(candidateInterfaces.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface
// Candidate backup interface
type Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (candidateInterface *Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerPrefix_CandidateInterfaces_CandidateInterface) GetEntityData() *types.CommonEntityData {
    candidateInterface.EntityData.YFilter = candidateInterface.YFilter
    candidateInterface.EntityData.YangName = "candidate-interface"
    candidateInterface.EntityData.BundleName = "cisco_ios_xr"
    candidateInterface.EntityData.ParentYangName = "candidate-interfaces"
    candidateInterface.EntityData.SegmentPath = "candidate-interface" + types.AddKeyToken(candidateInterface.InterfaceName, "interface-name")
    candidateInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/process-scope/fast-reroute/per-prefix/candidate-interfaces/" + candidateInterface.EntityData.SegmentPath
    candidateInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    candidateInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    candidateInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    candidateInterface.EntityData.Children = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs = types.NewOrderedMap()
    candidateInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", candidateInterface.InterfaceName})

    candidateInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(candidateInterface.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerPrefix_ExcludeInterfaces
// Fast-reroute per-link/per-prefix exclude
// interface configuration
type Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerPrefix_ExcludeInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Exclude an interface from becoming a backup. The type is slice of
    // Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface.
    ExcludeInterface []*Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface
}

func (excludeInterfaces *Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerPrefix_ExcludeInterfaces) GetEntityData() *types.CommonEntityData {
    excludeInterfaces.EntityData.YFilter = excludeInterfaces.YFilter
    excludeInterfaces.EntityData.YangName = "exclude-interfaces"
    excludeInterfaces.EntityData.BundleName = "cisco_ios_xr"
    excludeInterfaces.EntityData.ParentYangName = "per-prefix"
    excludeInterfaces.EntityData.SegmentPath = "exclude-interfaces"
    excludeInterfaces.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/process-scope/fast-reroute/per-prefix/" + excludeInterfaces.EntityData.SegmentPath
    excludeInterfaces.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterfaces.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterfaces.EntityData.Children = types.NewOrderedMap()
    excludeInterfaces.EntityData.Children.Append("exclude-interface", types.YChild{"ExcludeInterface", nil})
    for i := range excludeInterfaces.ExcludeInterface {
        excludeInterfaces.EntityData.Children.Append(types.GetSegmentPath(excludeInterfaces.ExcludeInterface[i]), types.YChild{"ExcludeInterface", excludeInterfaces.ExcludeInterface[i]})
    }
    excludeInterfaces.EntityData.Leafs = types.NewOrderedMap()

    excludeInterfaces.EntityData.YListKeys = []string {}

    return &(excludeInterfaces.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface
// Exclude an interface from becoming a backup
type Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Interface. The type is string with pattern:
    // b'[a-zA-Z0-9._/-]+'.
    InterfaceName interface{}
}

func (excludeInterface *Ospfv3_Processes_Process_Vrfs_Vrf_ProcessScope_FastReroute_PerPrefix_ExcludeInterfaces_ExcludeInterface) GetEntityData() *types.CommonEntityData {
    excludeInterface.EntityData.YFilter = excludeInterface.YFilter
    excludeInterface.EntityData.YangName = "exclude-interface"
    excludeInterface.EntityData.BundleName = "cisco_ios_xr"
    excludeInterface.EntityData.ParentYangName = "exclude-interfaces"
    excludeInterface.EntityData.SegmentPath = "exclude-interface" + types.AddKeyToken(excludeInterface.InterfaceName, "interface-name")
    excludeInterface.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/process-scope/fast-reroute/per-prefix/exclude-interfaces/" + excludeInterface.EntityData.SegmentPath
    excludeInterface.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    excludeInterface.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    excludeInterface.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    excludeInterface.EntityData.Children = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs = types.NewOrderedMap()
    excludeInterface.EntityData.Leafs.Append("interface-name", types.YLeaf{"InterfaceName", excludeInterface.InterfaceName})

    excludeInterface.EntityData.YListKeys = []string {"InterfaceName"}

    return &(excludeInterface.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_Encryption
// Encrypt and authenticate OSPFv3 packets
type Ospfv3_Processes_Process_Vrfs_Vrf_Encryption struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Authenticate packets. The type is bool.
    Enable interface{}

    // Use IPSec ESP authentication. Specify the Security Parameter Index (SPI)
    // value. The type is interface{} with range: 256..4294967295.
    Spi interface{}

    // Specify the encryption algorithm. The type is Ospfv3EncryptionAlgorithm.
    EncryptionAlgorithm interface{}

    // Encryption password. The type is string with pattern: b'(!.+)|([^!].+)'.
    EncryptionPassword interface{}

    // Use the NULL, MD5 or SHA1 algorithm. The type is Ospfv3AuthenticationType2.
    AuthenticationAlgorithm interface{}

    // Specify MD5 or SHA1 password. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    AuthenticationPassword interface{}
}

func (encryption *Ospfv3_Processes_Process_Vrfs_Vrf_Encryption) GetEntityData() *types.CommonEntityData {
    encryption.EntityData.YFilter = encryption.YFilter
    encryption.EntityData.YangName = "encryption"
    encryption.EntityData.BundleName = "cisco_ios_xr"
    encryption.EntityData.ParentYangName = "vrf"
    encryption.EntityData.SegmentPath = "encryption"
    encryption.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/" + encryption.EntityData.SegmentPath
    encryption.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryption.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryption.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryption.EntityData.Children = types.NewOrderedMap()
    encryption.EntityData.Leafs = types.NewOrderedMap()
    encryption.EntityData.Leafs.Append("enable", types.YLeaf{"Enable", encryption.Enable})
    encryption.EntityData.Leafs.Append("spi", types.YLeaf{"Spi", encryption.Spi})
    encryption.EntityData.Leafs.Append("encryption-algorithm", types.YLeaf{"EncryptionAlgorithm", encryption.EncryptionAlgorithm})
    encryption.EntityData.Leafs.Append("encryption-password", types.YLeaf{"EncryptionPassword", encryption.EncryptionPassword})
    encryption.EntityData.Leafs.Append("authentication-algorithm", types.YLeaf{"AuthenticationAlgorithm", encryption.AuthenticationAlgorithm})
    encryption.EntityData.Leafs.Append("authentication-password", types.YLeaf{"AuthenticationPassword", encryption.AuthenticationPassword})

    encryption.EntityData.YListKeys = []string {}

    return &(encryption.EntityData)
}

// Ospfv3_Processes_Process_Vrfs_Vrf_AutoCost
// Calculate interface cost according to bandwidth
// This type is a presence type.
type Ospfv3_Processes_Process_Vrfs_Vrf_AutoCost struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Specify 'true' to assign cost based on interface type. The type is
    // interface{}.
    Disable interface{}

    // Specify reference bandwidth for cost computations in terms of Mbits per
    // second. The type is interface{} with range: 1..4294967. Units are Mbit/s.
    ReferenceBandwidth interface{}
}

func (autoCost *Ospfv3_Processes_Process_Vrfs_Vrf_AutoCost) GetEntityData() *types.CommonEntityData {
    autoCost.EntityData.YFilter = autoCost.YFilter
    autoCost.EntityData.YangName = "auto-cost"
    autoCost.EntityData.BundleName = "cisco_ios_xr"
    autoCost.EntityData.ParentYangName = "vrf"
    autoCost.EntityData.SegmentPath = "auto-cost"
    autoCost.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/vrfs/vrf/" + autoCost.EntityData.SegmentPath
    autoCost.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    autoCost.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    autoCost.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    autoCost.EntityData.Children = types.NewOrderedMap()
    autoCost.EntityData.Leafs = types.NewOrderedMap()
    autoCost.EntityData.Leafs.Append("disable", types.YLeaf{"Disable", autoCost.Disable})
    autoCost.EntityData.Leafs.Append("reference-bandwidth", types.YLeaf{"ReferenceBandwidth", autoCost.ReferenceBandwidth})

    autoCost.EntityData.YListKeys = []string {}

    return &(autoCost.EntityData)
}

// Ospfv3_Processes_Process_Af
// Address Family (AF)
// This type is a presence type.
type Ospfv3_Processes_Process_Af struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Address Family (AF) identifier. The type is Ospfv3AddressFamily. This
    // attribute is mandatory.
    AfName interface{}

    // Subsequent Address Family (SAF) identifier. The type is
    // Ospfv3SubsequentAddressFamily.
    SafName interface{}
}

func (af *Ospfv3_Processes_Process_Af) GetEntityData() *types.CommonEntityData {
    af.EntityData.YFilter = af.YFilter
    af.EntityData.YangName = "af"
    af.EntityData.BundleName = "cisco_ios_xr"
    af.EntityData.ParentYangName = "process"
    af.EntityData.SegmentPath = "af"
    af.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/" + af.EntityData.SegmentPath
    af.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    af.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    af.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    af.EntityData.Children = types.NewOrderedMap()
    af.EntityData.Leafs = types.NewOrderedMap()
    af.EntityData.Leafs.Append("af-name", types.YLeaf{"AfName", af.AfName})
    af.EntityData.Leafs.Append("saf-name", types.YLeaf{"SafName", af.SafName})

    af.EntityData.YListKeys = []string {}

    return &(af.EntityData)
}

// Ospfv3_Processes_Process_TraceBufs
// Configuration to change size of trace buffer
type Ospfv3_Processes_Process_TraceBufs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Changes the size of the specified trace buffer. The type is slice of
    // Ospfv3_Processes_Process_TraceBufs_TraceBuf.
    TraceBuf []*Ospfv3_Processes_Process_TraceBufs_TraceBuf
}

func (traceBufs *Ospfv3_Processes_Process_TraceBufs) GetEntityData() *types.CommonEntityData {
    traceBufs.EntityData.YFilter = traceBufs.YFilter
    traceBufs.EntityData.YangName = "trace-bufs"
    traceBufs.EntityData.BundleName = "cisco_ios_xr"
    traceBufs.EntityData.ParentYangName = "process"
    traceBufs.EntityData.SegmentPath = "trace-bufs"
    traceBufs.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/" + traceBufs.EntityData.SegmentPath
    traceBufs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traceBufs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traceBufs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traceBufs.EntityData.Children = types.NewOrderedMap()
    traceBufs.EntityData.Children.Append("trace-buf", types.YChild{"TraceBuf", nil})
    for i := range traceBufs.TraceBuf {
        traceBufs.EntityData.Children.Append(types.GetSegmentPath(traceBufs.TraceBuf[i]), types.YChild{"TraceBuf", traceBufs.TraceBuf[i]})
    }
    traceBufs.EntityData.Leafs = types.NewOrderedMap()

    traceBufs.EntityData.YListKeys = []string {}

    return &(traceBufs.EntityData)
}

// Ospfv3_Processes_Process_TraceBufs_TraceBuf
// Changes the size of the specified trace
// buffer
type Ospfv3_Processes_Process_TraceBufs_TraceBuf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name for this trace buffer. The type is string
    // with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    TraceBufName interface{}

    // Buffer size. The type is Ospfv3TraceBufSize. This attribute is mandatory.
    Bufsize interface{}
}

func (traceBuf *Ospfv3_Processes_Process_TraceBufs_TraceBuf) GetEntityData() *types.CommonEntityData {
    traceBuf.EntityData.YFilter = traceBuf.YFilter
    traceBuf.EntityData.YangName = "trace-buf"
    traceBuf.EntityData.BundleName = "cisco_ios_xr"
    traceBuf.EntityData.ParentYangName = "trace-bufs"
    traceBuf.EntityData.SegmentPath = "trace-buf" + types.AddKeyToken(traceBuf.TraceBufName, "trace-buf-name")
    traceBuf.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv6-ospfv3-cfg:ospfv3/processes/process/trace-bufs/" + traceBuf.EntityData.SegmentPath
    traceBuf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    traceBuf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    traceBuf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    traceBuf.EntityData.Children = types.NewOrderedMap()
    traceBuf.EntityData.Leafs = types.NewOrderedMap()
    traceBuf.EntityData.Leafs.Append("trace-buf-name", types.YLeaf{"TraceBufName", traceBuf.TraceBufName})
    traceBuf.EntityData.Leafs.Append("bufsize", types.YLeaf{"Bufsize", traceBuf.Bufsize})

    traceBuf.EntityData.YListKeys = []string {"TraceBufName"}

    return &(traceBuf.EntityData)
}

