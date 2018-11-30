// This is a MIB Module for monitoring the
// structures in IPSec-based Virtual Private Networks.
// The MIB has been designed to be adopted as an IETF
// standard. Hence Cisco-specific features of IPSec
// protocol are excluded from this MIB. 
// 
// Acronyms
// The following acronyms are used in this document:
// 
//  IPSec:      Secure IP Protocol
// 
//  VPN:        Virtual Private Network
// 
//  ISAKMP:     Internet Security Association and Key Exchange
//              Protocol
// 
//  IKE:        Internet Key Exchange Protocol
// 
//  SA:         Security Association
// 
//  MM:         Main Mode - the process of setting up
//              a Phase 1 SA to secure the exchanges
//              required to setup Phase 2 SAs
// 
//  QM:         Quick Mode - the process of setting up
//              Phase 2 Security Associations using 
//              a Phase 1 SA.
// 
// 
//  Overview of IPsec MIB
// 
// The MIB contains six major groups of objects which are
// used to manage the IPSec Protocol. These groups include
// a Levels Group, a Phase-1 Group, a Phase-2 Group,
// a History Group, a Failure Group and a TRAP Control Group.
// The following table illustrates the structure of the
// IPSec MIB.
// 
// The Phase 1 group models objects pertaining to
// IKE negotiations and tunnels.
// 
// The Phase 2 group models objects pertaining to
// IPSec data tunnels.
// 
// The History group is to aid applications that do
// trending analysis.
// 
// The Failure group is to enable an operator to
// do troubleshooting and debugging of the VPN Router.
// Further, counters are supported to aid Intrusion 
// Detection.
// 
// In addition to the five major MIB Groups, there are
// a number of Notifications. The following table
// illustrates the name and description of the 
// IPSec TRAPs.
// 
// For a detailed discussion, please refer to the IETF
// draft draft-ietf-ipsec-flow-monitoring-mib-00.txt.
package cisco_ipsec_flow_monitor_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ipsec_flow_monitor_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IPSEC-FLOW-MONITOR-MIB CISCO-IPSEC-FLOW-MONITOR-MIB}", reflect.TypeOf(CISCOIPSECFLOWMONITORMIB{}))
    ydk.RegisterEntity("CISCO-IPSEC-FLOW-MONITOR-MIB:CISCO-IPSEC-FLOW-MONITOR-MIB", reflect.TypeOf(CISCOIPSECFLOWMONITORMIB{}))
}

// TunnelStatus represents type cannot be used to create a Tunnel.
type TunnelStatus string

const (
    TunnelStatus_active TunnelStatus = "active"

    TunnelStatus_destroy TunnelStatus = "destroy"
)

// IkeHashAlgo represents IKE negotiations.
type IkeHashAlgo string

const (
    IkeHashAlgo_none IkeHashAlgo = "none"

    IkeHashAlgo_md5 IkeHashAlgo = "md5"

    IkeHashAlgo_sha IkeHashAlgo = "sha"
)

// IkeAuthMethod represents negotiations.
type IkeAuthMethod string

const (
    IkeAuthMethod_none IkeAuthMethod = "none"

    IkeAuthMethod_preSharedKey IkeAuthMethod = "preSharedKey"

    IkeAuthMethod_rsaSig IkeAuthMethod = "rsaSig"

    IkeAuthMethod_rsaEncrypt IkeAuthMethod = "rsaEncrypt"

    IkeAuthMethod_revPublicKey IkeAuthMethod = "revPublicKey"
)

// KeyType represents The type of key used by an IPsec Phase-2 Tunnel.
type KeyType string

const (
    KeyType_ike KeyType = "ike"

    KeyType_manual KeyType = "manual"
)

// TrapStatus represents The administrative status for sending a TRAP.
type TrapStatus string

const (
    TrapStatus_enabled TrapStatus = "enabled"

    TrapStatus_disabled TrapStatus = "disabled"
)

// CompAlgo represents security association of an IPsec Phase-2 Tunnel.
type CompAlgo string

const (
    CompAlgo_none CompAlgo = "none"

    CompAlgo_ldf CompAlgo = "ldf"
)

// EncryptAlgo represents The encryption algorithm used in negotiations.
type EncryptAlgo string

const (
    EncryptAlgo_none EncryptAlgo = "none"

    EncryptAlgo_des EncryptAlgo = "des"

    EncryptAlgo_des3 EncryptAlgo = "des3"
)

// IkePeerType represents  2. a host name.
type IkePeerType string

const (
    IkePeerType_ipAddrPeer IkePeerType = "ipAddrPeer"

    IkePeerType_namePeer IkePeerType = "namePeer"
)

// IkeNegoMode represents The IPsec Phase-1 IKE negotiation mode.
type IkeNegoMode string

const (
    IkeNegoMode_main IkeNegoMode = "main"

    IkeNegoMode_aggressive IkeNegoMode = "aggressive"
)

// EncapMode represents Tunnel.
type EncapMode string

const (
    EncapMode_tunnel EncapMode = "tunnel"

    EncapMode_transport EncapMode = "transport"
)

// AuthAlgo represents security association of an IPsec Phase-2 Tunnel.
type AuthAlgo string

const (
    AuthAlgo_none AuthAlgo = "none"

    AuthAlgo_hmacMd5 AuthAlgo = "hmacMd5"

    AuthAlgo_hmacSha AuthAlgo = "hmacSha"
)

// EndPtType represents The type of identity use to specify an IPsec End Point.
type EndPtType string

const (
    EndPtType_singleIpAddr EndPtType = "singleIpAddr"

    EndPtType_ipAddrRange EndPtType = "ipAddrRange"

    EndPtType_ipSubnet EndPtType = "ipSubnet"
)

// DiffHellmanGrp represents The Diffie Hellman Group used in negotiations.
type DiffHellmanGrp string

const (
    DiffHellmanGrp_none DiffHellmanGrp = "none"

    DiffHellmanGrp_dhGroup1 DiffHellmanGrp = "dhGroup1"

    DiffHellmanGrp_dhGroup2 DiffHellmanGrp = "dhGroup2"
)

// CISCOIPSECFLOWMONITORMIB
type CISCOIPSECFLOWMONITORMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CipSecLevels CISCOIPSECFLOWMONITORMIB_CipSecLevels

    
    CikeGlobalStats CISCOIPSECFLOWMONITORMIB_CikeGlobalStats

    
    CipSecGlobalStats CISCOIPSECFLOWMONITORMIB_CipSecGlobalStats

    
    CipSecHistGlobalCntl CISCOIPSECFLOWMONITORMIB_CipSecHistGlobalCntl

    
    CipSecFailGlobalCntl CISCOIPSECFLOWMONITORMIB_CipSecFailGlobalCntl

    
    CipSecTrapCntl CISCOIPSECFLOWMONITORMIB_CipSecTrapCntl

    // The IPsec Phase-1 Internet Key Exchange Peer Table. There is one entry in
    // this table for each IPsec Phase-1 IKE peer association which is currently
    // associated with an active IPsec Phase-1 Tunnel. The IPsec Phase-1 IKE
    // Tunnel associated with this IPsec Phase-1 IKE peer association may or may
    // not be currently active.
    CikePeerTable CISCOIPSECFLOWMONITORMIB_CikePeerTable

    // The IPsec Phase-1 Internet Key Exchange Tunnel Table. There is one entry in
    // this table for each active IPsec Phase-1 IKE Tunnel.
    CikeTunnelTable CISCOIPSECFLOWMONITORMIB_CikeTunnelTable

    // The IPsec Phase-1 Internet Key Exchange Peer Association to IPsec Phase-2
    // Tunnel Correlation Table. There is one entry in this table for each active
    // IPsec Phase-2 Tunnel.
    CikePeerCorrTable CISCOIPSECFLOWMONITORMIB_CikePeerCorrTable

    // Phase-1 IKE stats information is included in this table. Each entry is
    // related to a specific gateway which is  identified by 'cmgwIndex'.
    CikePhase1GWStatsTable CISCOIPSECFLOWMONITORMIB_CikePhase1GWStatsTable

    // The IPsec Phase-2 Tunnel Table. There is one entry in this table for  each
    // active IPsec Phase-2 Tunnel.
    CipSecTunnelTable CISCOIPSECFLOWMONITORMIB_CipSecTunnelTable

    // The IPsec Phase-2 Tunnel Endpoint Table. This table contains an entry for
    // each  active endpoint associated with an IPsec  Phase-2 Tunnel.
    CipSecEndPtTable CISCOIPSECFLOWMONITORMIB_CipSecEndPtTable

    // The IPsec Phase-2 Security Protection Index Table. This table contains an
    // entry for each active  and expiring security  association.
    CipSecSpiTable CISCOIPSECFLOWMONITORMIB_CipSecSpiTable

    // Phase-2 IPsec stats information is included in this table. Each entry is
    // related to a specific gateway which is  identified by 'cmgwIndex'.
    CipSecPhase2GWStatsTable CISCOIPSECFLOWMONITORMIB_CipSecPhase2GWStatsTable

    // The IPsec Phase-1 Internet Key Exchange Tunnel History Table.  This table
    // is implemented as a  sliding window in which only the last n entries  are
    // maintained.  The maximum number of entries  is specified by the
    // cipSecHistTableSize object.
    CikeTunnelHistTable CISCOIPSECFLOWMONITORMIB_CikeTunnelHistTable

    // The IPsec Phase-2 Tunnel History Table. This table is implemented as a
    // sliding  window in which only the last n entries are maintained.  The
    // maximum number  of entries is specified by the cipSecHistTableSize object.
    CipSecTunnelHistTable CISCOIPSECFLOWMONITORMIB_CipSecTunnelHistTable

    // The IPsec Phase-2 Tunnel Endpoint History Table. This table is implemented
    // as a  sliding window in which only the last n entries are maintained.   The
    // maximum number of entries is specified by the cipSecHistTableSize object.
    CipSecEndPtHistTable CISCOIPSECFLOWMONITORMIB_CipSecEndPtHistTable

    // The IPsec Phase-1 Failure Table. This table is implemented as a sliding 
    // window in which only the last n entries are  maintained.  The maximum
    // number of entries is specified by the cipSecFailTableSize object.
    CikeFailTable CISCOIPSECFLOWMONITORMIB_CikeFailTable

    // The IPsec Phase-2 Failure Table. This table is implemented as a sliding
    // window  in which only the last n entries are maintained.   The maximum
    // number of entries is specified by the cipSecFailTableSize object.
    CipSecFailTable CISCOIPSECFLOWMONITORMIB_CipSecFailTable
}

func (cISCOIPSECFLOWMONITORMIB *CISCOIPSECFLOWMONITORMIB) GetEntityData() *types.CommonEntityData {
    cISCOIPSECFLOWMONITORMIB.EntityData.YFilter = cISCOIPSECFLOWMONITORMIB.YFilter
    cISCOIPSECFLOWMONITORMIB.EntityData.YangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cISCOIPSECFLOWMONITORMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIPSECFLOWMONITORMIB.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cISCOIPSECFLOWMONITORMIB.EntityData.SegmentPath = "CISCO-IPSEC-FLOW-MONITOR-MIB:CISCO-IPSEC-FLOW-MONITOR-MIB"
    cISCOIPSECFLOWMONITORMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIPSECFLOWMONITORMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIPSECFLOWMONITORMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIPSECFLOWMONITORMIB.EntityData.Children = types.NewOrderedMap()
    cISCOIPSECFLOWMONITORMIB.EntityData.Children.Append("cipSecLevels", types.YChild{"CipSecLevels", &cISCOIPSECFLOWMONITORMIB.CipSecLevels})
    cISCOIPSECFLOWMONITORMIB.EntityData.Children.Append("cikeGlobalStats", types.YChild{"CikeGlobalStats", &cISCOIPSECFLOWMONITORMIB.CikeGlobalStats})
    cISCOIPSECFLOWMONITORMIB.EntityData.Children.Append("cipSecGlobalStats", types.YChild{"CipSecGlobalStats", &cISCOIPSECFLOWMONITORMIB.CipSecGlobalStats})
    cISCOIPSECFLOWMONITORMIB.EntityData.Children.Append("cipSecHistGlobalCntl", types.YChild{"CipSecHistGlobalCntl", &cISCOIPSECFLOWMONITORMIB.CipSecHistGlobalCntl})
    cISCOIPSECFLOWMONITORMIB.EntityData.Children.Append("cipSecFailGlobalCntl", types.YChild{"CipSecFailGlobalCntl", &cISCOIPSECFLOWMONITORMIB.CipSecFailGlobalCntl})
    cISCOIPSECFLOWMONITORMIB.EntityData.Children.Append("cipSecTrapCntl", types.YChild{"CipSecTrapCntl", &cISCOIPSECFLOWMONITORMIB.CipSecTrapCntl})
    cISCOIPSECFLOWMONITORMIB.EntityData.Children.Append("cikePeerTable", types.YChild{"CikePeerTable", &cISCOIPSECFLOWMONITORMIB.CikePeerTable})
    cISCOIPSECFLOWMONITORMIB.EntityData.Children.Append("cikeTunnelTable", types.YChild{"CikeTunnelTable", &cISCOIPSECFLOWMONITORMIB.CikeTunnelTable})
    cISCOIPSECFLOWMONITORMIB.EntityData.Children.Append("cikePeerCorrTable", types.YChild{"CikePeerCorrTable", &cISCOIPSECFLOWMONITORMIB.CikePeerCorrTable})
    cISCOIPSECFLOWMONITORMIB.EntityData.Children.Append("cikePhase1GWStatsTable", types.YChild{"CikePhase1GWStatsTable", &cISCOIPSECFLOWMONITORMIB.CikePhase1GWStatsTable})
    cISCOIPSECFLOWMONITORMIB.EntityData.Children.Append("cipSecTunnelTable", types.YChild{"CipSecTunnelTable", &cISCOIPSECFLOWMONITORMIB.CipSecTunnelTable})
    cISCOIPSECFLOWMONITORMIB.EntityData.Children.Append("cipSecEndPtTable", types.YChild{"CipSecEndPtTable", &cISCOIPSECFLOWMONITORMIB.CipSecEndPtTable})
    cISCOIPSECFLOWMONITORMIB.EntityData.Children.Append("cipSecSpiTable", types.YChild{"CipSecSpiTable", &cISCOIPSECFLOWMONITORMIB.CipSecSpiTable})
    cISCOIPSECFLOWMONITORMIB.EntityData.Children.Append("cipSecPhase2GWStatsTable", types.YChild{"CipSecPhase2GWStatsTable", &cISCOIPSECFLOWMONITORMIB.CipSecPhase2GWStatsTable})
    cISCOIPSECFLOWMONITORMIB.EntityData.Children.Append("cikeTunnelHistTable", types.YChild{"CikeTunnelHistTable", &cISCOIPSECFLOWMONITORMIB.CikeTunnelHistTable})
    cISCOIPSECFLOWMONITORMIB.EntityData.Children.Append("cipSecTunnelHistTable", types.YChild{"CipSecTunnelHistTable", &cISCOIPSECFLOWMONITORMIB.CipSecTunnelHistTable})
    cISCOIPSECFLOWMONITORMIB.EntityData.Children.Append("cipSecEndPtHistTable", types.YChild{"CipSecEndPtHistTable", &cISCOIPSECFLOWMONITORMIB.CipSecEndPtHistTable})
    cISCOIPSECFLOWMONITORMIB.EntityData.Children.Append("cikeFailTable", types.YChild{"CikeFailTable", &cISCOIPSECFLOWMONITORMIB.CikeFailTable})
    cISCOIPSECFLOWMONITORMIB.EntityData.Children.Append("cipSecFailTable", types.YChild{"CipSecFailTable", &cISCOIPSECFLOWMONITORMIB.CipSecFailTable})
    cISCOIPSECFLOWMONITORMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOIPSECFLOWMONITORMIB.EntityData.YListKeys = []string {}

    return &(cISCOIPSECFLOWMONITORMIB.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CipSecLevels
type CISCOIPSECFLOWMONITORMIB_CipSecLevels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The level of the IPsec MIB. The type is interface{} with range: 1..4096.
    CipSecMibLevel interface{}
}

func (cipSecLevels *CISCOIPSECFLOWMONITORMIB_CipSecLevels) GetEntityData() *types.CommonEntityData {
    cipSecLevels.EntityData.YFilter = cipSecLevels.YFilter
    cipSecLevels.EntityData.YangName = "cipSecLevels"
    cipSecLevels.EntityData.BundleName = "cisco_ios_xe"
    cipSecLevels.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cipSecLevels.EntityData.SegmentPath = "cipSecLevels"
    cipSecLevels.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipSecLevels.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipSecLevels.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipSecLevels.EntityData.Children = types.NewOrderedMap()
    cipSecLevels.EntityData.Leafs = types.NewOrderedMap()
    cipSecLevels.EntityData.Leafs.Append("cipSecMibLevel", types.YLeaf{"CipSecMibLevel", cipSecLevels.CipSecMibLevel})

    cipSecLevels.EntityData.YListKeys = []string {}

    return &(cipSecLevels.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CikeGlobalStats
type CISCOIPSECFLOWMONITORMIB_CikeGlobalStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of currently active IPsec Phase-1 IKE Tunnels. The type is
    // interface{} with range: 0..4294967295.
    CikeGlobalActiveTunnels interface{}

    // The total number of previously active IPsec Phase-1 IKE Tunnels. The type
    // is interface{} with range: 0..4294967295. Units are SAs.
    CikeGlobalPreviousTunnels interface{}

    // The total number of octets received by all currently and previously active
    // IPsec Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Octets.
    CikeGlobalInOctets interface{}

    // The total number of packets received by all currently and previously active
    // IPsec  Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    CikeGlobalInPkts interface{}

    // The total number of packets which were dropped during receive processing by
    // all  currently and previously  active IPsec Phase-1 IKE Tunnels. The type
    // is interface{} with range: 0..4294967295. Units are Packets.
    CikeGlobalInDropPkts interface{}

    // The total number of notifys received by all currently and previously active
    // IPsec  Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Notification Payloads.
    CikeGlobalInNotifys interface{}

    // The total number of IPsec Phase-2 exchanges received by all currently and
    // previously  active IPsec Phase-1 IKE Tunnels. The type is interface{} with
    // range: 0..4294967295. Units are SA Payloads.
    CikeGlobalInP2Exchgs interface{}

    // The total number of IPsec Phase-2 exchanges which were received and found
    // to be invalid  by all currently and previously active IPsec  Phase-1 IKE
    // Tunnels. The type is interface{} with range: 0..4294967295. Units are SA
    // Payloads.
    CikeGlobalInP2ExchgInvalids interface{}

    // The total number of IPsec Phase-2 exchanges which were received and
    // rejected by all  currently and previously active IPsec Phase-1  IKE
    // Tunnels. The type is interface{} with range: 0..4294967295. Units are SA
    // Payloads.
    CikeGlobalInP2ExchgRejects interface{}

    // The total number of IPsec Phase-2 security association delete requests
    // received by all  currently and previously  active and IPsec Phase-1 IKE
    // Tunnels. The type is interface{} with range: 0..4294967295. Units are
    // Notification Payloads.
    CikeGlobalInP2SaDelRequests interface{}

    // The total number of octets sent by all currently and previously active and
    // IPsec Phase-1  IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Octets.
    CikeGlobalOutOctets interface{}

    // The total number of packets sent by all currently and previously active and
    // IPsec Phase-1  Tunnels. The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    CikeGlobalOutPkts interface{}

    // The total number of packets which were dropped during send processing by
    // all currently  and previously  active IPsec Phase-1 IKE Tunnels. The type
    // is interface{} with range: 0..4294967295. Units are Packets.
    CikeGlobalOutDropPkts interface{}

    // The total number of notifys sent by all currently and previously active
    // IPsec Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Notification Payloads.
    CikeGlobalOutNotifys interface{}

    // The total number of IPsec Phase-2 exchanges which were sent by all
    // currently and previously  active IPsec Phase-1 IKE Tunnels. The type is
    // interface{} with range: 0..4294967295. Units are SA Payloads.
    CikeGlobalOutP2Exchgs interface{}

    // The total number of IPsec Phase-2 exchanges which were sent and found to be
    // invalid by  all currently and previously active IPsec Phase-1  Tunnels. The
    // type is interface{} with range: 0..4294967295. Units are SA Payloads.
    CikeGlobalOutP2ExchgInvalids interface{}

    // The total number of IPsec Phase-2 exchanges which were sent and rejected by
    // all currently and  previously active IPsec Phase-1 IKE Tunnels. The type is
    // interface{} with range: 0..4294967295. Units are SA Payloads.
    CikeGlobalOutP2ExchgRejects interface{}

    // The total number of IPsec Phase-2 SA delete requests sent by all currently
    // and  previously active IPsec Phase-1 IKE Tunnels. The type is interface{}
    // with range: 0..4294967295. Units are Notification Payloads.
    CikeGlobalOutP2SaDelRequests interface{}

    // The total number of IPsec Phase-1 IKE Tunnels which were locally initiated.
    // The type is interface{} with range: 0..4294967295. Units are SAs.
    CikeGlobalInitTunnels interface{}

    // The total number of IPsec Phase-1 IKE Tunnels which were locally initiated
    // and failed to activate. The type is interface{} with range: 0..4294967295.
    // Units are SAs.
    CikeGlobalInitTunnelFails interface{}

    // The total number of IPsec Phase-1 IKE Tunnels which were remotely initiated
    // and failed to activate. The type is interface{} with range: 0..4294967295.
    // Units are SAs.
    CikeGlobalRespTunnelFails interface{}

    // The total number of system capacity failures which occurred during
    // processing of all current  and previously active IPsec Phase-1 IKE Tunnels.
    // The type is interface{} with range: 0..4294967295. Units are Failures.
    CikeGlobalSysCapFails interface{}

    // The total number of authentications which ended in failure by all current
    // and previous IPsec Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Failures.
    CikeGlobalAuthFails interface{}

    // The total number of decryptions which ended in failure by all current and
    // previous IPsec Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Failures.
    CikeGlobalDecryptFails interface{}

    // The total number of hash validations which ended in failure by all current
    // and previous IPsec Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Failures.
    CikeGlobalHashValidFails interface{}

    // The total number of non-existent Security Association in failures which
    // occurred during processing of  all current and previous IPsec Phase-1 IKE
    // Tunnels. The type is interface{} with range: 0..4294967295. Units are
    // Failures.
    CikeGlobalNoSaFails interface{}
}

func (cikeGlobalStats *CISCOIPSECFLOWMONITORMIB_CikeGlobalStats) GetEntityData() *types.CommonEntityData {
    cikeGlobalStats.EntityData.YFilter = cikeGlobalStats.YFilter
    cikeGlobalStats.EntityData.YangName = "cikeGlobalStats"
    cikeGlobalStats.EntityData.BundleName = "cisco_ios_xe"
    cikeGlobalStats.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cikeGlobalStats.EntityData.SegmentPath = "cikeGlobalStats"
    cikeGlobalStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cikeGlobalStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cikeGlobalStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cikeGlobalStats.EntityData.Children = types.NewOrderedMap()
    cikeGlobalStats.EntityData.Leafs = types.NewOrderedMap()
    cikeGlobalStats.EntityData.Leafs.Append("cikeGlobalActiveTunnels", types.YLeaf{"CikeGlobalActiveTunnels", cikeGlobalStats.CikeGlobalActiveTunnels})
    cikeGlobalStats.EntityData.Leafs.Append("cikeGlobalPreviousTunnels", types.YLeaf{"CikeGlobalPreviousTunnels", cikeGlobalStats.CikeGlobalPreviousTunnels})
    cikeGlobalStats.EntityData.Leafs.Append("cikeGlobalInOctets", types.YLeaf{"CikeGlobalInOctets", cikeGlobalStats.CikeGlobalInOctets})
    cikeGlobalStats.EntityData.Leafs.Append("cikeGlobalInPkts", types.YLeaf{"CikeGlobalInPkts", cikeGlobalStats.CikeGlobalInPkts})
    cikeGlobalStats.EntityData.Leafs.Append("cikeGlobalInDropPkts", types.YLeaf{"CikeGlobalInDropPkts", cikeGlobalStats.CikeGlobalInDropPkts})
    cikeGlobalStats.EntityData.Leafs.Append("cikeGlobalInNotifys", types.YLeaf{"CikeGlobalInNotifys", cikeGlobalStats.CikeGlobalInNotifys})
    cikeGlobalStats.EntityData.Leafs.Append("cikeGlobalInP2Exchgs", types.YLeaf{"CikeGlobalInP2Exchgs", cikeGlobalStats.CikeGlobalInP2Exchgs})
    cikeGlobalStats.EntityData.Leafs.Append("cikeGlobalInP2ExchgInvalids", types.YLeaf{"CikeGlobalInP2ExchgInvalids", cikeGlobalStats.CikeGlobalInP2ExchgInvalids})
    cikeGlobalStats.EntityData.Leafs.Append("cikeGlobalInP2ExchgRejects", types.YLeaf{"CikeGlobalInP2ExchgRejects", cikeGlobalStats.CikeGlobalInP2ExchgRejects})
    cikeGlobalStats.EntityData.Leafs.Append("cikeGlobalInP2SaDelRequests", types.YLeaf{"CikeGlobalInP2SaDelRequests", cikeGlobalStats.CikeGlobalInP2SaDelRequests})
    cikeGlobalStats.EntityData.Leafs.Append("cikeGlobalOutOctets", types.YLeaf{"CikeGlobalOutOctets", cikeGlobalStats.CikeGlobalOutOctets})
    cikeGlobalStats.EntityData.Leafs.Append("cikeGlobalOutPkts", types.YLeaf{"CikeGlobalOutPkts", cikeGlobalStats.CikeGlobalOutPkts})
    cikeGlobalStats.EntityData.Leafs.Append("cikeGlobalOutDropPkts", types.YLeaf{"CikeGlobalOutDropPkts", cikeGlobalStats.CikeGlobalOutDropPkts})
    cikeGlobalStats.EntityData.Leafs.Append("cikeGlobalOutNotifys", types.YLeaf{"CikeGlobalOutNotifys", cikeGlobalStats.CikeGlobalOutNotifys})
    cikeGlobalStats.EntityData.Leafs.Append("cikeGlobalOutP2Exchgs", types.YLeaf{"CikeGlobalOutP2Exchgs", cikeGlobalStats.CikeGlobalOutP2Exchgs})
    cikeGlobalStats.EntityData.Leafs.Append("cikeGlobalOutP2ExchgInvalids", types.YLeaf{"CikeGlobalOutP2ExchgInvalids", cikeGlobalStats.CikeGlobalOutP2ExchgInvalids})
    cikeGlobalStats.EntityData.Leafs.Append("cikeGlobalOutP2ExchgRejects", types.YLeaf{"CikeGlobalOutP2ExchgRejects", cikeGlobalStats.CikeGlobalOutP2ExchgRejects})
    cikeGlobalStats.EntityData.Leafs.Append("cikeGlobalOutP2SaDelRequests", types.YLeaf{"CikeGlobalOutP2SaDelRequests", cikeGlobalStats.CikeGlobalOutP2SaDelRequests})
    cikeGlobalStats.EntityData.Leafs.Append("cikeGlobalInitTunnels", types.YLeaf{"CikeGlobalInitTunnels", cikeGlobalStats.CikeGlobalInitTunnels})
    cikeGlobalStats.EntityData.Leafs.Append("cikeGlobalInitTunnelFails", types.YLeaf{"CikeGlobalInitTunnelFails", cikeGlobalStats.CikeGlobalInitTunnelFails})
    cikeGlobalStats.EntityData.Leafs.Append("cikeGlobalRespTunnelFails", types.YLeaf{"CikeGlobalRespTunnelFails", cikeGlobalStats.CikeGlobalRespTunnelFails})
    cikeGlobalStats.EntityData.Leafs.Append("cikeGlobalSysCapFails", types.YLeaf{"CikeGlobalSysCapFails", cikeGlobalStats.CikeGlobalSysCapFails})
    cikeGlobalStats.EntityData.Leafs.Append("cikeGlobalAuthFails", types.YLeaf{"CikeGlobalAuthFails", cikeGlobalStats.CikeGlobalAuthFails})
    cikeGlobalStats.EntityData.Leafs.Append("cikeGlobalDecryptFails", types.YLeaf{"CikeGlobalDecryptFails", cikeGlobalStats.CikeGlobalDecryptFails})
    cikeGlobalStats.EntityData.Leafs.Append("cikeGlobalHashValidFails", types.YLeaf{"CikeGlobalHashValidFails", cikeGlobalStats.CikeGlobalHashValidFails})
    cikeGlobalStats.EntityData.Leafs.Append("cikeGlobalNoSaFails", types.YLeaf{"CikeGlobalNoSaFails", cikeGlobalStats.CikeGlobalNoSaFails})

    cikeGlobalStats.EntityData.YListKeys = []string {}

    return &(cikeGlobalStats.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CipSecGlobalStats
type CISCOIPSECFLOWMONITORMIB_CipSecGlobalStats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The total number of currently active IPsec Phase-2 Tunnels. The type is
    // interface{} with range: 0..4294967295.
    CipSecGlobalActiveTunnels interface{}

    // The total number of previously active IPsec Phase-2 Tunnels. The type is
    // interface{} with range: 0..4294967295. Units are Phase-2 Tunnels.
    CipSecGlobalPreviousTunnels interface{}

    // The total number of octets received by all current and previous IPsec
    // Phase-2 Tunnels.  This value is accumulated BEFORE determining whether or
    // not the packet should be decompressed. See also cipSecGlobalInOctWraps for
    // the number of times this counter has wrapped. The type is interface{} with
    // range: 0..4294967295. Units are Octets.
    CipSecGlobalInOctets interface{}

    // A high capacity count of the total number of octets received by all current
    // and previous IPsec Phase-2 Tunnels. This value is accumulated BEFORE
    // determining whether or not the packet should be decompressed. The type is
    // interface{} with range: 0..18446744073709551615.
    CipSecGlobalHcInOctets interface{}

    // The number of times the global octets received counter
    // (cipSecGlobalInOctets) has wrapped. The type is interface{} with range:
    // 0..4294967295. Units are Integral units.
    CipSecGlobalInOctWraps interface{}

    // The total number of decompressed octets received by all current and
    // previous IPsec Phase-2 Tunnels.   This value is accumulated AFTER the
    // packet is  decompressed. If compression is not being used,  this value will
    // match the value of cipSecGlobalInOctets.  See also
    // cipSecGlobalInDecompOctWraps  for the number of times this counter has
    // wrapped. The type is interface{} with range: 0..4294967295. Units are
    // Octets.
    CipSecGlobalInDecompOctets interface{}

    // A high capacity count of the total number of decompressed octets received
    // by all current  and previous IPsec Phase-2 Tunnels.  This value  is
    // accumulated AFTER the packet is decompressed.  If compression is not being
    // used, this value   will match the value of cipSecGlobalHcInOctets. The type
    // is interface{} with range: 0..18446744073709551615.
    CipSecGlobalHcInDecompOctets interface{}

    // The number of times the global decompressed octets received counter 
    // (cipSecGlobalInDecompOctets) has wrapped. The type is interface{} with
    // range: 0..4294967295. Units are Integral units.
    CipSecGlobalInDecompOctWraps interface{}

    // The total number of packets received by all current and previous  IPsec
    // Phase-2 Tunnels. The type is interface{} with range: 0..4294967295. Units
    // are Packets.
    CipSecGlobalInPkts interface{}

    // The total number of packets dropped during receive processing by all
    // current and previous  IPsec Phase-2 Tunnels. This count does NOT include
    // packets dropped due to  Anti-Replay processing. The type is interface{}
    // with range: 0..4294967295. Units are Packets.
    CipSecGlobalInDrops interface{}

    // The total number of packets dropped during receive processing due to
    // Anti-Replay  processing by all current and previous IPsec  Phase-2 Tunnels.
    // The type is interface{} with range: 0..4294967295. Units are Packets.
    CipSecGlobalInReplayDrops interface{}

    // The total number of inbound authentication's performed by all current and
    // previous IPsec  Phase-2 Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Events.
    CipSecGlobalInAuths interface{}

    // The total number of inbound authentication's which ended in failure by all
    // current and previous  IPsec Phase-2 Tunnels. The type is interface{} with
    // range: 0..4294967295. Units are Failures.
    CipSecGlobalInAuthFails interface{}

    // The total number of inbound decryption's performed by all current and
    // previous IPsec  Phase-2 Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    CipSecGlobalInDecrypts interface{}

    // The total number of inbound decryption's which ended in failure by all
    // current and  previous IPsec Phase-2 Tunnels. The type is interface{} with
    // range: 0..4294967295. Units are Packets.
    CipSecGlobalInDecryptFails interface{}

    // The total number of octets sent by all current and previous IPsec Phase-2
    // Tunnels.   This value is accumulated AFTER determining  whether or not the
    // packet should be compressed.   See also cipSecGlobalOutOctWraps for the 
    // number of times this counter has wrapped. The type is interface{} with
    // range: 0..4294967295. Units are Octets.
    CipSecGlobalOutOctets interface{}

    // A high capacity count of the total number of octets sent by all current and
    // previous  IPsec Phase-2 Tunnels.  This value is accumulated  AFTER
    // determining whether or not the packet should  be compressed. The type is
    // interface{} with range: 0..18446744073709551615.
    CipSecGlobalHcOutOctets interface{}

    // The number of times the global octets sent counter (cipSecGlobalOutOctets)
    // has wrapped. The type is interface{} with range: 0..4294967295. Units are
    // Integral units.
    CipSecGlobalOutOctWraps interface{}

    // The total number of uncompressed octets sent by all current and previous
    // IPsec Phase-2 Tunnels.   This value is accumulated BEFORE the packet is 
    // compressed. If compression is not being used, this  value will match the
    // value of cipSecGlobalOutOctets.  See also cipSecGlobalOutDecompOctWraps for
    // the number  of times this counter has wrapped. The type is interface{} with
    // range: 0..4294967295. Units are Octets.
    CipSecGlobalOutUncompOctets interface{}

    // A high capacity count of the total number of uncompressed octets sent by
    // all current and previous  IPsec Phase-2 Tunnels.  This value is accumulated
    // BEFORE the packet is compressed.  If compression is  not being used, this
    // value will match the       value of cipSecGlobalHcOutOctets. The type is
    // interface{} with range: 0..18446744073709551615. Units are Octets.
    CipSecGlobalHcOutUncompOctets interface{}

    // The number of times the global uncompressed octets sent counter
    // (cipSecGlobalOutUncompOctets)  has wrapped. The type is interface{} with
    // range: 0..4294967295. Units are Integral units.
    CipSecGlobalOutUncompOctWraps interface{}

    // The total number of packets sent by all current and previous  IPsec Phase-2
    // Tunnels. The type is interface{} with range: 0..4294967295. Units are
    // Packets.
    CipSecGlobalOutPkts interface{}

    // The total number of packets dropped during send processing by all current
    // and previous IPsec  Phase-2 Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    CipSecGlobalOutDrops interface{}

    // The total number of outbound authentication's performed by all current and
    // previous IPsec  Phase-2 Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Events.
    CipSecGlobalOutAuths interface{}

    // The total number of outbound authentication's which ended in failure  by
    // all current and previous IPsec Phase-2 Tunnels. The type is interface{}
    // with range: 0..4294967295. Units are Failures.
    CipSecGlobalOutAuthFails interface{}

    // The total number of outbound encryption's performed by all current and
    // previous IPsec Phase-2 Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    CipSecGlobalOutEncrypts interface{}

    // The total number of outbound encryption's which ended in failure by all
    // current and  previous IPsec Phase-2 Tunnels. The type is interface{} with
    // range: 0..4294967295. Units are Failures.
    CipSecGlobalOutEncryptFails interface{}

    // The total number of protocol use failures which occurred during processing
    // of all current  and previously active IPsec Phase-2 Tunnels. The type is
    // interface{} with range: 0..4294967295. Units are Failures.
    CipSecGlobalProtocolUseFails interface{}

    // The total number of non-existent Security Association in failures which
    // occurred  during processing of all current  and previous IPsec Phase-2
    // Tunnels. The type is interface{} with range: 0..4294967295. Units are
    // Failures.
    CipSecGlobalNoSaFails interface{}

    // The total number of system capacity failures which occurred during
    // processing of all current  and previously active IPsec Phase-2 Tunnels. The
    // type is interface{} with range: 0..4294967295. Units are Failures.
    CipSecGlobalSysCapFails interface{}
}

func (cipSecGlobalStats *CISCOIPSECFLOWMONITORMIB_CipSecGlobalStats) GetEntityData() *types.CommonEntityData {
    cipSecGlobalStats.EntityData.YFilter = cipSecGlobalStats.YFilter
    cipSecGlobalStats.EntityData.YangName = "cipSecGlobalStats"
    cipSecGlobalStats.EntityData.BundleName = "cisco_ios_xe"
    cipSecGlobalStats.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cipSecGlobalStats.EntityData.SegmentPath = "cipSecGlobalStats"
    cipSecGlobalStats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipSecGlobalStats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipSecGlobalStats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipSecGlobalStats.EntityData.Children = types.NewOrderedMap()
    cipSecGlobalStats.EntityData.Leafs = types.NewOrderedMap()
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalActiveTunnels", types.YLeaf{"CipSecGlobalActiveTunnels", cipSecGlobalStats.CipSecGlobalActiveTunnels})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalPreviousTunnels", types.YLeaf{"CipSecGlobalPreviousTunnels", cipSecGlobalStats.CipSecGlobalPreviousTunnels})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalInOctets", types.YLeaf{"CipSecGlobalInOctets", cipSecGlobalStats.CipSecGlobalInOctets})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalHcInOctets", types.YLeaf{"CipSecGlobalHcInOctets", cipSecGlobalStats.CipSecGlobalHcInOctets})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalInOctWraps", types.YLeaf{"CipSecGlobalInOctWraps", cipSecGlobalStats.CipSecGlobalInOctWraps})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalInDecompOctets", types.YLeaf{"CipSecGlobalInDecompOctets", cipSecGlobalStats.CipSecGlobalInDecompOctets})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalHcInDecompOctets", types.YLeaf{"CipSecGlobalHcInDecompOctets", cipSecGlobalStats.CipSecGlobalHcInDecompOctets})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalInDecompOctWraps", types.YLeaf{"CipSecGlobalInDecompOctWraps", cipSecGlobalStats.CipSecGlobalInDecompOctWraps})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalInPkts", types.YLeaf{"CipSecGlobalInPkts", cipSecGlobalStats.CipSecGlobalInPkts})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalInDrops", types.YLeaf{"CipSecGlobalInDrops", cipSecGlobalStats.CipSecGlobalInDrops})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalInReplayDrops", types.YLeaf{"CipSecGlobalInReplayDrops", cipSecGlobalStats.CipSecGlobalInReplayDrops})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalInAuths", types.YLeaf{"CipSecGlobalInAuths", cipSecGlobalStats.CipSecGlobalInAuths})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalInAuthFails", types.YLeaf{"CipSecGlobalInAuthFails", cipSecGlobalStats.CipSecGlobalInAuthFails})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalInDecrypts", types.YLeaf{"CipSecGlobalInDecrypts", cipSecGlobalStats.CipSecGlobalInDecrypts})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalInDecryptFails", types.YLeaf{"CipSecGlobalInDecryptFails", cipSecGlobalStats.CipSecGlobalInDecryptFails})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalOutOctets", types.YLeaf{"CipSecGlobalOutOctets", cipSecGlobalStats.CipSecGlobalOutOctets})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalHcOutOctets", types.YLeaf{"CipSecGlobalHcOutOctets", cipSecGlobalStats.CipSecGlobalHcOutOctets})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalOutOctWraps", types.YLeaf{"CipSecGlobalOutOctWraps", cipSecGlobalStats.CipSecGlobalOutOctWraps})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalOutUncompOctets", types.YLeaf{"CipSecGlobalOutUncompOctets", cipSecGlobalStats.CipSecGlobalOutUncompOctets})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalHcOutUncompOctets", types.YLeaf{"CipSecGlobalHcOutUncompOctets", cipSecGlobalStats.CipSecGlobalHcOutUncompOctets})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalOutUncompOctWraps", types.YLeaf{"CipSecGlobalOutUncompOctWraps", cipSecGlobalStats.CipSecGlobalOutUncompOctWraps})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalOutPkts", types.YLeaf{"CipSecGlobalOutPkts", cipSecGlobalStats.CipSecGlobalOutPkts})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalOutDrops", types.YLeaf{"CipSecGlobalOutDrops", cipSecGlobalStats.CipSecGlobalOutDrops})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalOutAuths", types.YLeaf{"CipSecGlobalOutAuths", cipSecGlobalStats.CipSecGlobalOutAuths})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalOutAuthFails", types.YLeaf{"CipSecGlobalOutAuthFails", cipSecGlobalStats.CipSecGlobalOutAuthFails})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalOutEncrypts", types.YLeaf{"CipSecGlobalOutEncrypts", cipSecGlobalStats.CipSecGlobalOutEncrypts})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalOutEncryptFails", types.YLeaf{"CipSecGlobalOutEncryptFails", cipSecGlobalStats.CipSecGlobalOutEncryptFails})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalProtocolUseFails", types.YLeaf{"CipSecGlobalProtocolUseFails", cipSecGlobalStats.CipSecGlobalProtocolUseFails})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalNoSaFails", types.YLeaf{"CipSecGlobalNoSaFails", cipSecGlobalStats.CipSecGlobalNoSaFails})
    cipSecGlobalStats.EntityData.Leafs.Append("cipSecGlobalSysCapFails", types.YLeaf{"CipSecGlobalSysCapFails", cipSecGlobalStats.CipSecGlobalSysCapFails})

    cipSecGlobalStats.EntityData.YListKeys = []string {}

    return &(cipSecGlobalStats.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CipSecHistGlobalCntl
type CISCOIPSECFLOWMONITORMIB_CipSecHistGlobalCntl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The window size of the IPsec Phase-1 and Phase-2 History Tables.  The IPsec
    // Phase-1 and Phase-2 History Tables are implemented as a sliding window in
    // which only the last n entries are maintained.  This object is used specify
    // the number of entries which will be  maintained in the IPsec Phase-1 and 
    // Phase-2 History Tables.  An implementation may choose suitable minimum and 
    // maximum values for this element based on the local  policy and available
    // resources. If an SNMP SET request  specifies a value outside this window
    // for this element,  a BAD VALUE may be returned. The type is interface{}
    // with range: 1..2147483647.
    CipSecHistTableSize interface{}

    // The current state of check point processing.  This object will return ready
    // when the agent is  ready to create on-demand history entries for  active
    // IPsec Tunnels or checkPoint when the  agent is currently creating on-demand
    // history  entries for active IPsec Tunnels.  By setting this value to
    // checkPoint, the agent  will create: a) an entry in the IPsec Phase-1 Tunnel
    // History     for each active IPsec Phase-1 Tunnel and b) an entry in the
    // IPsec Phase-2 Tunnel History     Table and an entry in the IPsec Phase-2   
    // Tunnel EndPoint History Table    for each active IPsec Phase-2 Tunnel. The
    // type is CipSecHistCheckPoint.
    CipSecHistCheckPoint interface{}
}

func (cipSecHistGlobalCntl *CISCOIPSECFLOWMONITORMIB_CipSecHistGlobalCntl) GetEntityData() *types.CommonEntityData {
    cipSecHistGlobalCntl.EntityData.YFilter = cipSecHistGlobalCntl.YFilter
    cipSecHistGlobalCntl.EntityData.YangName = "cipSecHistGlobalCntl"
    cipSecHistGlobalCntl.EntityData.BundleName = "cisco_ios_xe"
    cipSecHistGlobalCntl.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cipSecHistGlobalCntl.EntityData.SegmentPath = "cipSecHistGlobalCntl"
    cipSecHistGlobalCntl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipSecHistGlobalCntl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipSecHistGlobalCntl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipSecHistGlobalCntl.EntityData.Children = types.NewOrderedMap()
    cipSecHistGlobalCntl.EntityData.Leafs = types.NewOrderedMap()
    cipSecHistGlobalCntl.EntityData.Leafs.Append("cipSecHistTableSize", types.YLeaf{"CipSecHistTableSize", cipSecHistGlobalCntl.CipSecHistTableSize})
    cipSecHistGlobalCntl.EntityData.Leafs.Append("cipSecHistCheckPoint", types.YLeaf{"CipSecHistCheckPoint", cipSecHistGlobalCntl.CipSecHistCheckPoint})

    cipSecHistGlobalCntl.EntityData.YListKeys = []string {}

    return &(cipSecHistGlobalCntl.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CipSecHistGlobalCntl_CipSecHistCheckPoint represents    for each active IPsec Phase-2 Tunnel.
type CISCOIPSECFLOWMONITORMIB_CipSecHistGlobalCntl_CipSecHistCheckPoint string

const (
    CISCOIPSECFLOWMONITORMIB_CipSecHistGlobalCntl_CipSecHistCheckPoint_ready CISCOIPSECFLOWMONITORMIB_CipSecHistGlobalCntl_CipSecHistCheckPoint = "ready"

    CISCOIPSECFLOWMONITORMIB_CipSecHistGlobalCntl_CipSecHistCheckPoint_checkPoint CISCOIPSECFLOWMONITORMIB_CipSecHistGlobalCntl_CipSecHistCheckPoint = "checkPoint"
)

// CISCOIPSECFLOWMONITORMIB_CipSecFailGlobalCntl
type CISCOIPSECFLOWMONITORMIB_CipSecFailGlobalCntl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The window size of the IPsec Phase-1 and Phase-2 Failure Tables.  The IPsec
    // Phase-1 and Phase-2 Failure Tables are implemented as a sliding window in
    // which only the last n entries are maintained.  This object is used specify
    // the number of entries which will be  maintained in the IPsec Phase-1 and
    // Phase-2 Failure  Tables.  An implementation may choose suitable minimum and
    // maximum values for this element based on the local  policy and available
    // resources. If an SNMP SET request  specifies a value outside this window
    // for this element,  a BAD VALUE may be returned. The type is interface{}
    // with range: 1..2147483647.
    CipSecFailTableSize interface{}
}

func (cipSecFailGlobalCntl *CISCOIPSECFLOWMONITORMIB_CipSecFailGlobalCntl) GetEntityData() *types.CommonEntityData {
    cipSecFailGlobalCntl.EntityData.YFilter = cipSecFailGlobalCntl.YFilter
    cipSecFailGlobalCntl.EntityData.YangName = "cipSecFailGlobalCntl"
    cipSecFailGlobalCntl.EntityData.BundleName = "cisco_ios_xe"
    cipSecFailGlobalCntl.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cipSecFailGlobalCntl.EntityData.SegmentPath = "cipSecFailGlobalCntl"
    cipSecFailGlobalCntl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipSecFailGlobalCntl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipSecFailGlobalCntl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipSecFailGlobalCntl.EntityData.Children = types.NewOrderedMap()
    cipSecFailGlobalCntl.EntityData.Leafs = types.NewOrderedMap()
    cipSecFailGlobalCntl.EntityData.Leafs.Append("cipSecFailTableSize", types.YLeaf{"CipSecFailTableSize", cipSecFailGlobalCntl.CipSecFailTableSize})

    cipSecFailGlobalCntl.EntityData.YListKeys = []string {}

    return &(cipSecFailGlobalCntl.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CipSecTrapCntl
type CISCOIPSECFLOWMONITORMIB_CipSecTrapCntl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object defines the administrative state of sending the IPsec IKE
    // Phase-1 Tunnel Start TRAP. The type is TrapStatus.
    CipSecTrapCntlIkeTunnelStart interface{}

    // This object defines the administrative state of sending the  IPsec IKE
    // Phase-1 Tunnel Stop TRAP. The type is TrapStatus.
    CipSecTrapCntlIkeTunnelStop interface{}

    // This object defines the administrative state of sending the  IPsec IKE
    // Phase-1 System Failure TRAP. The type is TrapStatus.
    CipSecTrapCntlIkeSysFailure interface{}

    // This object defines the administrative state of sending the  IPsec IKE
    // Phase-1 Certificate/CRL Failure TRAP. The type is TrapStatus.
    CipSecTrapCntlIkeCertCrlFailure interface{}

    // This object defines the administrative state of sending the  IPsec IKE
    // Phase-1 Protocol Failure TRAP. The type is TrapStatus.
    CipSecTrapCntlIkeProtocolFail interface{}

    // This object defines the administrative state of sending the  IPsec IKE
    // Phase-1 No Security Association TRAP. The type is TrapStatus.
    CipSecTrapCntlIkeNoSa interface{}

    // This object defines the administrative state of sending the IPsec  Phase-2
    // Tunnel Start TRAP. The type is TrapStatus.
    CipSecTrapCntlIpSecTunnelStart interface{}

    // This object defines the administrative state of sending the IPsec  Phase-2
    // Tunnel Stop TRAP. The type is TrapStatus.
    CipSecTrapCntlIpSecTunnelStop interface{}

    // This object defines the administrative state of sending the IPsec  Phase-2
    // System Failure TRAP. The type is TrapStatus.
    CipSecTrapCntlIpSecSysFailure interface{}

    // This object defines the administrative state of sending the IPsec  Phase-2
    // Set Up Failure TRAP. The type is TrapStatus.
    CipSecTrapCntlIpSecSetUpFailure interface{}

    // This object defines the administrative state of sending the IPsec  Phase-2
    // Early Tunnel Termination TRAP. The type is TrapStatus.
    CipSecTrapCntlIpSecEarlyTunTerm interface{}

    // This object defines the administrative state of sending the IPsec  Phase-2
    // Protocol Failure TRAP. The type is TrapStatus.
    CipSecTrapCntlIpSecProtocolFail interface{}

    // This object defines the administrative state of sending the IPsec  Phase-2 
    // No Security Association TRAP. The type is TrapStatus.
    CipSecTrapCntlIpSecNoSa interface{}
}

func (cipSecTrapCntl *CISCOIPSECFLOWMONITORMIB_CipSecTrapCntl) GetEntityData() *types.CommonEntityData {
    cipSecTrapCntl.EntityData.YFilter = cipSecTrapCntl.YFilter
    cipSecTrapCntl.EntityData.YangName = "cipSecTrapCntl"
    cipSecTrapCntl.EntityData.BundleName = "cisco_ios_xe"
    cipSecTrapCntl.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cipSecTrapCntl.EntityData.SegmentPath = "cipSecTrapCntl"
    cipSecTrapCntl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipSecTrapCntl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipSecTrapCntl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipSecTrapCntl.EntityData.Children = types.NewOrderedMap()
    cipSecTrapCntl.EntityData.Leafs = types.NewOrderedMap()
    cipSecTrapCntl.EntityData.Leafs.Append("cipSecTrapCntlIkeTunnelStart", types.YLeaf{"CipSecTrapCntlIkeTunnelStart", cipSecTrapCntl.CipSecTrapCntlIkeTunnelStart})
    cipSecTrapCntl.EntityData.Leafs.Append("cipSecTrapCntlIkeTunnelStop", types.YLeaf{"CipSecTrapCntlIkeTunnelStop", cipSecTrapCntl.CipSecTrapCntlIkeTunnelStop})
    cipSecTrapCntl.EntityData.Leafs.Append("cipSecTrapCntlIkeSysFailure", types.YLeaf{"CipSecTrapCntlIkeSysFailure", cipSecTrapCntl.CipSecTrapCntlIkeSysFailure})
    cipSecTrapCntl.EntityData.Leafs.Append("cipSecTrapCntlIkeCertCrlFailure", types.YLeaf{"CipSecTrapCntlIkeCertCrlFailure", cipSecTrapCntl.CipSecTrapCntlIkeCertCrlFailure})
    cipSecTrapCntl.EntityData.Leafs.Append("cipSecTrapCntlIkeProtocolFail", types.YLeaf{"CipSecTrapCntlIkeProtocolFail", cipSecTrapCntl.CipSecTrapCntlIkeProtocolFail})
    cipSecTrapCntl.EntityData.Leafs.Append("cipSecTrapCntlIkeNoSa", types.YLeaf{"CipSecTrapCntlIkeNoSa", cipSecTrapCntl.CipSecTrapCntlIkeNoSa})
    cipSecTrapCntl.EntityData.Leafs.Append("cipSecTrapCntlIpSecTunnelStart", types.YLeaf{"CipSecTrapCntlIpSecTunnelStart", cipSecTrapCntl.CipSecTrapCntlIpSecTunnelStart})
    cipSecTrapCntl.EntityData.Leafs.Append("cipSecTrapCntlIpSecTunnelStop", types.YLeaf{"CipSecTrapCntlIpSecTunnelStop", cipSecTrapCntl.CipSecTrapCntlIpSecTunnelStop})
    cipSecTrapCntl.EntityData.Leafs.Append("cipSecTrapCntlIpSecSysFailure", types.YLeaf{"CipSecTrapCntlIpSecSysFailure", cipSecTrapCntl.CipSecTrapCntlIpSecSysFailure})
    cipSecTrapCntl.EntityData.Leafs.Append("cipSecTrapCntlIpSecSetUpFailure", types.YLeaf{"CipSecTrapCntlIpSecSetUpFailure", cipSecTrapCntl.CipSecTrapCntlIpSecSetUpFailure})
    cipSecTrapCntl.EntityData.Leafs.Append("cipSecTrapCntlIpSecEarlyTunTerm", types.YLeaf{"CipSecTrapCntlIpSecEarlyTunTerm", cipSecTrapCntl.CipSecTrapCntlIpSecEarlyTunTerm})
    cipSecTrapCntl.EntityData.Leafs.Append("cipSecTrapCntlIpSecProtocolFail", types.YLeaf{"CipSecTrapCntlIpSecProtocolFail", cipSecTrapCntl.CipSecTrapCntlIpSecProtocolFail})
    cipSecTrapCntl.EntityData.Leafs.Append("cipSecTrapCntlIpSecNoSa", types.YLeaf{"CipSecTrapCntlIpSecNoSa", cipSecTrapCntl.CipSecTrapCntlIpSecNoSa})

    cipSecTrapCntl.EntityData.YListKeys = []string {}

    return &(cipSecTrapCntl.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CikePeerTable
// The IPsec Phase-1 Internet Key Exchange Peer Table.
// There is one entry in this table for each IPsec
// Phase-1 IKE peer association which is currently
// associated with an active IPsec Phase-1 Tunnel.
// The IPsec Phase-1 IKE Tunnel associated with this
// IPsec Phase-1 IKE peer association may or may not
// be currently active.
type CISCOIPSECFLOWMONITORMIB_CikePeerTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with an IPsec Phase-1 IKE
    // peer association. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_CikePeerTable_CikePeerEntry.
    CikePeerEntry []*CISCOIPSECFLOWMONITORMIB_CikePeerTable_CikePeerEntry
}

func (cikePeerTable *CISCOIPSECFLOWMONITORMIB_CikePeerTable) GetEntityData() *types.CommonEntityData {
    cikePeerTable.EntityData.YFilter = cikePeerTable.YFilter
    cikePeerTable.EntityData.YangName = "cikePeerTable"
    cikePeerTable.EntityData.BundleName = "cisco_ios_xe"
    cikePeerTable.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cikePeerTable.EntityData.SegmentPath = "cikePeerTable"
    cikePeerTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cikePeerTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cikePeerTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cikePeerTable.EntityData.Children = types.NewOrderedMap()
    cikePeerTable.EntityData.Children.Append("cikePeerEntry", types.YChild{"CikePeerEntry", nil})
    for i := range cikePeerTable.CikePeerEntry {
        cikePeerTable.EntityData.Children.Append(types.GetSegmentPath(cikePeerTable.CikePeerEntry[i]), types.YChild{"CikePeerEntry", cikePeerTable.CikePeerEntry[i]})
    }
    cikePeerTable.EntityData.Leafs = types.NewOrderedMap()

    cikePeerTable.EntityData.YListKeys = []string {}

    return &(cikePeerTable.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CikePeerTable_CikePeerEntry
// Each entry contains the attributes associated
// with an IPsec Phase-1 IKE peer association.
type CISCOIPSECFLOWMONITORMIB_CikePeerTable_CikePeerEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of local peer identity.  The local peer
    // may be identified by: 1. an IP address, or 2. a host name. The type is
    // IkePeerType.
    CikePeerLocalType interface{}

    // This attribute is a key. The value of the local peer identity.  If the
    // local peer type is an IP Address, then this is the IP Address used to
    // identify the local peer.  If the local peer type is a host name, then this
    // is the host name used to identify the local peer. The type is string.
    CikePeerLocalValue interface{}

    // This attribute is a key. The type of remote peer identity.  The remote peer
    // may be identified by: 1. an IP address, or 2. a host name. The type is
    // IkePeerType.
    CikePeerRemoteType interface{}

    // This attribute is a key. The value of the remote peer identity.  If the
    // remote peer type is an IP Address, then this is the IP Address used to
    // identify the remote peer.  If the remote peer type is a host name, then
    // this is the host name used to identify the remote peer. The type is string.
    CikePeerRemoteValue interface{}

    // This attribute is a key. The internal index of the local-remote peer
    // association.  This internal index is used  to uniquely identify multiple
    // associations between  the local and remote peer. The type is interface{}
    // with range: 1..2147483647.
    CikePeerIntIndex interface{}

    // The IP address of the local peer. The type is string with length: 4 | 16.
    CikePeerLocalAddr interface{}

    // The IP address of the remote peer. The type is string with length: 4 | 16.
    CikePeerRemoteAddr interface{}

    // The length of time that the peer association has existed in hundredths of a
    // second. The type is interface{} with range: 0..2147483647.
    CikePeerActiveTime interface{}

    // The index of the active IPsec Phase-1 IKE Tunnel (cikeTunIndex in the
    // cikeTunnelTable) for this peer association.  If an IPsec Phase-1 IKE Tunnel
    // is not currently active, then the value of this object will be zero. The
    // type is interface{} with range: 1..2147483647.
    CikePeerActiveTunnelIndex interface{}
}

func (cikePeerEntry *CISCOIPSECFLOWMONITORMIB_CikePeerTable_CikePeerEntry) GetEntityData() *types.CommonEntityData {
    cikePeerEntry.EntityData.YFilter = cikePeerEntry.YFilter
    cikePeerEntry.EntityData.YangName = "cikePeerEntry"
    cikePeerEntry.EntityData.BundleName = "cisco_ios_xe"
    cikePeerEntry.EntityData.ParentYangName = "cikePeerTable"
    cikePeerEntry.EntityData.SegmentPath = "cikePeerEntry" + types.AddKeyToken(cikePeerEntry.CikePeerLocalType, "cikePeerLocalType") + types.AddKeyToken(cikePeerEntry.CikePeerLocalValue, "cikePeerLocalValue") + types.AddKeyToken(cikePeerEntry.CikePeerRemoteType, "cikePeerRemoteType") + types.AddKeyToken(cikePeerEntry.CikePeerRemoteValue, "cikePeerRemoteValue") + types.AddKeyToken(cikePeerEntry.CikePeerIntIndex, "cikePeerIntIndex")
    cikePeerEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cikePeerEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cikePeerEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cikePeerEntry.EntityData.Children = types.NewOrderedMap()
    cikePeerEntry.EntityData.Leafs = types.NewOrderedMap()
    cikePeerEntry.EntityData.Leafs.Append("cikePeerLocalType", types.YLeaf{"CikePeerLocalType", cikePeerEntry.CikePeerLocalType})
    cikePeerEntry.EntityData.Leafs.Append("cikePeerLocalValue", types.YLeaf{"CikePeerLocalValue", cikePeerEntry.CikePeerLocalValue})
    cikePeerEntry.EntityData.Leafs.Append("cikePeerRemoteType", types.YLeaf{"CikePeerRemoteType", cikePeerEntry.CikePeerRemoteType})
    cikePeerEntry.EntityData.Leafs.Append("cikePeerRemoteValue", types.YLeaf{"CikePeerRemoteValue", cikePeerEntry.CikePeerRemoteValue})
    cikePeerEntry.EntityData.Leafs.Append("cikePeerIntIndex", types.YLeaf{"CikePeerIntIndex", cikePeerEntry.CikePeerIntIndex})
    cikePeerEntry.EntityData.Leafs.Append("cikePeerLocalAddr", types.YLeaf{"CikePeerLocalAddr", cikePeerEntry.CikePeerLocalAddr})
    cikePeerEntry.EntityData.Leafs.Append("cikePeerRemoteAddr", types.YLeaf{"CikePeerRemoteAddr", cikePeerEntry.CikePeerRemoteAddr})
    cikePeerEntry.EntityData.Leafs.Append("cikePeerActiveTime", types.YLeaf{"CikePeerActiveTime", cikePeerEntry.CikePeerActiveTime})
    cikePeerEntry.EntityData.Leafs.Append("cikePeerActiveTunnelIndex", types.YLeaf{"CikePeerActiveTunnelIndex", cikePeerEntry.CikePeerActiveTunnelIndex})

    cikePeerEntry.EntityData.YListKeys = []string {"CikePeerLocalType", "CikePeerLocalValue", "CikePeerRemoteType", "CikePeerRemoteValue", "CikePeerIntIndex"}

    return &(cikePeerEntry.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CikeTunnelTable
// The IPsec Phase-1 Internet Key Exchange Tunnel Table.
// There is one entry in this table for each active IPsec
// Phase-1 IKE Tunnel.
type CISCOIPSECFLOWMONITORMIB_CikeTunnelTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with an active IPsec Phase-1
    // IKE Tunnel. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_CikeTunnelTable_CikeTunnelEntry.
    CikeTunnelEntry []*CISCOIPSECFLOWMONITORMIB_CikeTunnelTable_CikeTunnelEntry
}

func (cikeTunnelTable *CISCOIPSECFLOWMONITORMIB_CikeTunnelTable) GetEntityData() *types.CommonEntityData {
    cikeTunnelTable.EntityData.YFilter = cikeTunnelTable.YFilter
    cikeTunnelTable.EntityData.YangName = "cikeTunnelTable"
    cikeTunnelTable.EntityData.BundleName = "cisco_ios_xe"
    cikeTunnelTable.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cikeTunnelTable.EntityData.SegmentPath = "cikeTunnelTable"
    cikeTunnelTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cikeTunnelTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cikeTunnelTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cikeTunnelTable.EntityData.Children = types.NewOrderedMap()
    cikeTunnelTable.EntityData.Children.Append("cikeTunnelEntry", types.YChild{"CikeTunnelEntry", nil})
    for i := range cikeTunnelTable.CikeTunnelEntry {
        cikeTunnelTable.EntityData.Children.Append(types.GetSegmentPath(cikeTunnelTable.CikeTunnelEntry[i]), types.YChild{"CikeTunnelEntry", cikeTunnelTable.CikeTunnelEntry[i]})
    }
    cikeTunnelTable.EntityData.Leafs = types.NewOrderedMap()

    cikeTunnelTable.EntityData.YListKeys = []string {}

    return &(cikeTunnelTable.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CikeTunnelTable_CikeTunnelEntry
// Each entry contains the attributes associated with
// an active IPsec Phase-1 IKE Tunnel.
type CISCOIPSECFLOWMONITORMIB_CikeTunnelTable_CikeTunnelEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index of the IPsec Phase-1 IKE Tunnel Table.
    // The value of the index is a number which begins  at one and is incremented
    // with each tunnel that  is created. The value of this object will  wrap at
    // 2,147,483,647. The type is interface{} with range: 1..2147483647.
    CikeTunIndex interface{}

    // The type of local peer identity.  The local peer may be identified by:  1.
    // an IP address, or  2. a host name. The type is IkePeerType.
    CikeTunLocalType interface{}

    // The value of the local peer identity.  If the local peer type is an IP
    // Address, then this is the IP Address used to identify the local peer.  If
    // the local peer type is a host name, then this is the host name used to
    // identify the local peer. The type is string.
    CikeTunLocalValue interface{}

    // The IP address of the local endpoint for the IPsec Phase-1 IKE Tunnel. The
    // type is string with length: 4 | 16.
    CikeTunLocalAddr interface{}

    // The DNS name of the local IP address for the IPsec Phase-1 IKE Tunnel. If
    // the DNS  name associated with the local tunnel endpoint  is not known, then
    // the value of this  object will be a NULL string. The type is string.
    CikeTunLocalName interface{}

    // The type of remote peer identity. The remote peer may be identified by:  1.
    // an IP address, or  2. a host name. The type is IkePeerType.
    CikeTunRemoteType interface{}

    // The value of the remote peer identity.  If the remote peer type is an IP
    // Address, then this is the IP Address used to identify the remote peer.  If
    // the remote peer type is a host name, then  this is the host name used to
    // identify the  remote peer. The type is string.
    CikeTunRemoteValue interface{}

    // The IP address of the remote endpoint for the IPsec Phase-1 IKE Tunnel. The
    // type is string with length: 4 | 16.
    CikeTunRemoteAddr interface{}

    // The DNS name of the remote IP address of IPsec Phase-1 IKE Tunnel. If the
    // DNS name associated with the remote tunnel endpoint is not known, then the
    // value of this object will be a NULL string. The type is string.
    CikeTunRemoteName interface{}

    // The negotiation mode of the IPsec Phase-1 IKE Tunnel. The type is
    // IkeNegoMode.
    CikeTunNegoMode interface{}

    // The Diffie Hellman Group used in IPsec Phase-1 IKE negotiations. The type
    // is DiffHellmanGrp.
    CikeTunDiffHellmanGrp interface{}

    // The encryption algorithm used in IPsec Phase-1 IKE negotiations. The type
    // is EncryptAlgo.
    CikeTunEncryptAlgo interface{}

    // The hash algorithm used in IPsec Phase-1 IKE negotiations. The type is
    // IkeHashAlgo.
    CikeTunHashAlgo interface{}

    // The authentication method used in IPsec Phase-1 IKE negotiations. The type
    // is IkeAuthMethod.
    CikeTunAuthMethod interface{}

    // The negotiated LifeTime of the IPsec Phase-1 IKE Tunnel in seconds. The
    // type is interface{} with range: 1..2147483647. Units are seconds.
    CikeTunLifeTime interface{}

    // The length of time the IPsec Phase-1 IKE tunnel has been active in
    // hundredths of seconds. The type is interface{} with range: 0..2147483647.
    CikeTunActiveTime interface{}

    // The security association refresh threshold in seconds. The type is
    // interface{} with range: 1..2147483647. Units are seconds.
    CikeTunSaRefreshThreshold interface{}

    // The total number of security associations refreshes performed. The type is
    // interface{} with range: 0..4294967295. Units are QM Exchanges.
    CikeTunTotalRefreshes interface{}

    // The total number of octets received by this IPsec Phase-1 IKE Tunnel. The
    // type is interface{} with range: 0..4294967295. Units are Octets.
    CikeTunInOctets interface{}

    // The total number of packets received by this IPsec Phase-1 IKE Tunnel. The
    // type is interface{} with range: 0..4294967295. Units are Packets.
    CikeTunInPkts interface{}

    // The total number of packets dropped by this IPsec Phase-1 IKE Tunnel during
    // receive processing. The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    CikeTunInDropPkts interface{}

    // The total number of notifys received by this IPsec Phase-1 IKE Tunnel. The
    // type is interface{} with range: 0..4294967295. Units are Notification
    // Payloads.
    CikeTunInNotifys interface{}

    // The total number of IPsec Phase-2 exchanges received by  this IPsec Phase-1
    // IKE Tunnel. The type is interface{} with range: 0..4294967295. Units are SA
    // Payloads.
    CikeTunInP2Exchgs interface{}

    // The total number of IPsec Phase-2 exchanges received and found to be
    // invalid  by this IPsec Phase-1 IKE Tunnel. The type is interface{} with
    // range: 0..4294967295. Units are SA Payloads.
    CikeTunInP2ExchgInvalids interface{}

    // The total number of IPsec Phase-2 exchanges received and rejected by this
    // IPsec Phase-1  Tunnel. The type is interface{} with range: 0..4294967295.
    // Units are SA Payloads.
    CikeTunInP2ExchgRejects interface{}

    // The total number of IPsec Phase-2 security association delete requests
    // received  by this IPsec Phase-1 IKE Tunnel. The type is interface{} with
    // range: 0..4294967295. Units are Notification Payloads.
    CikeTunInP2SaDelRequests interface{}

    // The total number of octets sent by this IPsec Phase-1 IKE Tunnel. The type
    // is interface{} with range: 0..4294967295. Units are Octets.
    CikeTunOutOctets interface{}

    // The total number of packets sent by this IPsec Phase-1 IKE Tunnel. The type
    // is interface{} with range: 0..4294967295. Units are Packets.
    CikeTunOutPkts interface{}

    // The total number of packets dropped by this IPsec Phase-1 IKE Tunnel during
    // send processing. The type is interface{} with range: 0..4294967295. Units
    // are Packets.
    CikeTunOutDropPkts interface{}

    // The total number of notifys sent by this IPsec Phase-1 Tunnel. The type is
    // interface{} with range: 0..4294967295. Units are Notification Payloads.
    CikeTunOutNotifys interface{}

    // The total number of IPsec Phase-2 exchanges sent by this IPsec Phase-1 IKE
    // Tunnel. The type is interface{} with range: 0..4294967295. Units are SA
    // Payloads.
    CikeTunOutP2Exchgs interface{}

    // The total number of IPsec Phase-2 exchanges sent and found to be invalid by
    // this IPsec Phase-1 IKE Tunnel. The type is interface{} with range:
    // 0..4294967295. Units are SA Payloads.
    CikeTunOutP2ExchgInvalids interface{}

    // The total number of IPsec Phase-2 exchanges sent and rejected by this IPsec
    // Phase-1 IKE Tunnel. The type is interface{} with range: 0..4294967295.
    // Units are SA Payloads.
    CikeTunOutP2ExchgRejects interface{}

    // The total number of IPsec Phase-2 security association delete requests sent
    // by this IPsec Phase-1 IKE Tunnel. The type is interface{} with range:
    // 0..4294967295. Units are Notification Payloads.
    CikeTunOutP2SaDelRequests interface{}

    // The status of the MIB table row.  This object can be used to bring the
    // tunnel down  by setting value of this object to destroy(2).  This object
    // cannot be used to create  a MIB table row. The type is TunnelStatus.
    CikeTunStatus interface{}
}

func (cikeTunnelEntry *CISCOIPSECFLOWMONITORMIB_CikeTunnelTable_CikeTunnelEntry) GetEntityData() *types.CommonEntityData {
    cikeTunnelEntry.EntityData.YFilter = cikeTunnelEntry.YFilter
    cikeTunnelEntry.EntityData.YangName = "cikeTunnelEntry"
    cikeTunnelEntry.EntityData.BundleName = "cisco_ios_xe"
    cikeTunnelEntry.EntityData.ParentYangName = "cikeTunnelTable"
    cikeTunnelEntry.EntityData.SegmentPath = "cikeTunnelEntry" + types.AddKeyToken(cikeTunnelEntry.CikeTunIndex, "cikeTunIndex")
    cikeTunnelEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cikeTunnelEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cikeTunnelEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cikeTunnelEntry.EntityData.Children = types.NewOrderedMap()
    cikeTunnelEntry.EntityData.Leafs = types.NewOrderedMap()
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunIndex", types.YLeaf{"CikeTunIndex", cikeTunnelEntry.CikeTunIndex})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunLocalType", types.YLeaf{"CikeTunLocalType", cikeTunnelEntry.CikeTunLocalType})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunLocalValue", types.YLeaf{"CikeTunLocalValue", cikeTunnelEntry.CikeTunLocalValue})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunLocalAddr", types.YLeaf{"CikeTunLocalAddr", cikeTunnelEntry.CikeTunLocalAddr})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunLocalName", types.YLeaf{"CikeTunLocalName", cikeTunnelEntry.CikeTunLocalName})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunRemoteType", types.YLeaf{"CikeTunRemoteType", cikeTunnelEntry.CikeTunRemoteType})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunRemoteValue", types.YLeaf{"CikeTunRemoteValue", cikeTunnelEntry.CikeTunRemoteValue})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunRemoteAddr", types.YLeaf{"CikeTunRemoteAddr", cikeTunnelEntry.CikeTunRemoteAddr})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunRemoteName", types.YLeaf{"CikeTunRemoteName", cikeTunnelEntry.CikeTunRemoteName})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunNegoMode", types.YLeaf{"CikeTunNegoMode", cikeTunnelEntry.CikeTunNegoMode})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunDiffHellmanGrp", types.YLeaf{"CikeTunDiffHellmanGrp", cikeTunnelEntry.CikeTunDiffHellmanGrp})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunEncryptAlgo", types.YLeaf{"CikeTunEncryptAlgo", cikeTunnelEntry.CikeTunEncryptAlgo})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunHashAlgo", types.YLeaf{"CikeTunHashAlgo", cikeTunnelEntry.CikeTunHashAlgo})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunAuthMethod", types.YLeaf{"CikeTunAuthMethod", cikeTunnelEntry.CikeTunAuthMethod})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunLifeTime", types.YLeaf{"CikeTunLifeTime", cikeTunnelEntry.CikeTunLifeTime})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunActiveTime", types.YLeaf{"CikeTunActiveTime", cikeTunnelEntry.CikeTunActiveTime})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunSaRefreshThreshold", types.YLeaf{"CikeTunSaRefreshThreshold", cikeTunnelEntry.CikeTunSaRefreshThreshold})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunTotalRefreshes", types.YLeaf{"CikeTunTotalRefreshes", cikeTunnelEntry.CikeTunTotalRefreshes})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunInOctets", types.YLeaf{"CikeTunInOctets", cikeTunnelEntry.CikeTunInOctets})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunInPkts", types.YLeaf{"CikeTunInPkts", cikeTunnelEntry.CikeTunInPkts})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunInDropPkts", types.YLeaf{"CikeTunInDropPkts", cikeTunnelEntry.CikeTunInDropPkts})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunInNotifys", types.YLeaf{"CikeTunInNotifys", cikeTunnelEntry.CikeTunInNotifys})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunInP2Exchgs", types.YLeaf{"CikeTunInP2Exchgs", cikeTunnelEntry.CikeTunInP2Exchgs})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunInP2ExchgInvalids", types.YLeaf{"CikeTunInP2ExchgInvalids", cikeTunnelEntry.CikeTunInP2ExchgInvalids})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunInP2ExchgRejects", types.YLeaf{"CikeTunInP2ExchgRejects", cikeTunnelEntry.CikeTunInP2ExchgRejects})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunInP2SaDelRequests", types.YLeaf{"CikeTunInP2SaDelRequests", cikeTunnelEntry.CikeTunInP2SaDelRequests})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunOutOctets", types.YLeaf{"CikeTunOutOctets", cikeTunnelEntry.CikeTunOutOctets})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunOutPkts", types.YLeaf{"CikeTunOutPkts", cikeTunnelEntry.CikeTunOutPkts})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunOutDropPkts", types.YLeaf{"CikeTunOutDropPkts", cikeTunnelEntry.CikeTunOutDropPkts})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunOutNotifys", types.YLeaf{"CikeTunOutNotifys", cikeTunnelEntry.CikeTunOutNotifys})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunOutP2Exchgs", types.YLeaf{"CikeTunOutP2Exchgs", cikeTunnelEntry.CikeTunOutP2Exchgs})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunOutP2ExchgInvalids", types.YLeaf{"CikeTunOutP2ExchgInvalids", cikeTunnelEntry.CikeTunOutP2ExchgInvalids})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunOutP2ExchgRejects", types.YLeaf{"CikeTunOutP2ExchgRejects", cikeTunnelEntry.CikeTunOutP2ExchgRejects})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunOutP2SaDelRequests", types.YLeaf{"CikeTunOutP2SaDelRequests", cikeTunnelEntry.CikeTunOutP2SaDelRequests})
    cikeTunnelEntry.EntityData.Leafs.Append("cikeTunStatus", types.YLeaf{"CikeTunStatus", cikeTunnelEntry.CikeTunStatus})

    cikeTunnelEntry.EntityData.YListKeys = []string {"CikeTunIndex"}

    return &(cikeTunnelEntry.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CikePeerCorrTable
// The IPsec Phase-1 Internet Key Exchange Peer
// Association to IPsec Phase-2 Tunnel
// Correlation Table. There is one entry in
// this table for each active IPsec Phase-2
// Tunnel.
type CISCOIPSECFLOWMONITORMIB_CikePeerCorrTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes of an IPsec Phase-1 IKE Peer Association
    // to IPsec Phase-2 Tunnel Correlation. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_CikePeerCorrTable_CikePeerCorrEntry.
    CikePeerCorrEntry []*CISCOIPSECFLOWMONITORMIB_CikePeerCorrTable_CikePeerCorrEntry
}

func (cikePeerCorrTable *CISCOIPSECFLOWMONITORMIB_CikePeerCorrTable) GetEntityData() *types.CommonEntityData {
    cikePeerCorrTable.EntityData.YFilter = cikePeerCorrTable.YFilter
    cikePeerCorrTable.EntityData.YangName = "cikePeerCorrTable"
    cikePeerCorrTable.EntityData.BundleName = "cisco_ios_xe"
    cikePeerCorrTable.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cikePeerCorrTable.EntityData.SegmentPath = "cikePeerCorrTable"
    cikePeerCorrTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cikePeerCorrTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cikePeerCorrTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cikePeerCorrTable.EntityData.Children = types.NewOrderedMap()
    cikePeerCorrTable.EntityData.Children.Append("cikePeerCorrEntry", types.YChild{"CikePeerCorrEntry", nil})
    for i := range cikePeerCorrTable.CikePeerCorrEntry {
        cikePeerCorrTable.EntityData.Children.Append(types.GetSegmentPath(cikePeerCorrTable.CikePeerCorrEntry[i]), types.YChild{"CikePeerCorrEntry", cikePeerCorrTable.CikePeerCorrEntry[i]})
    }
    cikePeerCorrTable.EntityData.Leafs = types.NewOrderedMap()

    cikePeerCorrTable.EntityData.YListKeys = []string {}

    return &(cikePeerCorrTable.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CikePeerCorrTable_CikePeerCorrEntry
// Each entry contains the attributes of an
// IPsec Phase-1 IKE Peer Association to IPsec
// Phase-2 Tunnel Correlation.
type CISCOIPSECFLOWMONITORMIB_CikePeerCorrTable_CikePeerCorrEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of local peer identity. The local peer
    // may be identified by: 1. an IP address, or 2. a host name. The type is
    // IkePeerType.
    CikePeerCorrLocalType interface{}

    // This attribute is a key. The value of the local peer identity.  If the
    // local peer type is an IP Address, then this is the IP Address used to
    // identify the local peer.  If the local peer type is a host name, then this
    // is the host name used to identify the local peer. The type is string.
    CikePeerCorrLocalValue interface{}

    // This attribute is a key. The type of remote peer identity. The remote peer
    // may be identified by: 1. an IP address, or 2. a host name. The type is
    // IkePeerType.
    CikePeerCorrRemoteType interface{}

    // This attribute is a key. The value of the remote peer identity.  If the
    // remote peer type is an IP Address, then this is the IP Address used to
    // identify the remote peer.  If the remote peer type is a host name, then
    // this is the host name used to identify the remote peer. The type is string.
    CikePeerCorrRemoteValue interface{}

    // This attribute is a key. The internal index of the local-remote peer
    // association.  This internal index is  used to uniquely identify multiple
    // associations  between the local and remote peer. The type is interface{}
    // with range: 1..2147483647.
    CikePeerCorrIntIndex interface{}

    // This attribute is a key. The sequence number of the local-remote peer
    // association.  This sequence number is  used to uniquely identify multiple
    // instances  of an unique association between  the local and remote peer. The
    // type is interface{} with range: 1..2147483647.
    CikePeerCorrSeqNum interface{}

    // The index of the active IPsec Phase-2 Tunnel (cipSecTunIndex in the
    // cipSecTunnelTable) for this IPsec Phase-1 IKE Peer Association. The type is
    // interface{} with range: 1..2147483647.
    CikePeerCorrIpSecTunIndex interface{}
}

func (cikePeerCorrEntry *CISCOIPSECFLOWMONITORMIB_CikePeerCorrTable_CikePeerCorrEntry) GetEntityData() *types.CommonEntityData {
    cikePeerCorrEntry.EntityData.YFilter = cikePeerCorrEntry.YFilter
    cikePeerCorrEntry.EntityData.YangName = "cikePeerCorrEntry"
    cikePeerCorrEntry.EntityData.BundleName = "cisco_ios_xe"
    cikePeerCorrEntry.EntityData.ParentYangName = "cikePeerCorrTable"
    cikePeerCorrEntry.EntityData.SegmentPath = "cikePeerCorrEntry" + types.AddKeyToken(cikePeerCorrEntry.CikePeerCorrLocalType, "cikePeerCorrLocalType") + types.AddKeyToken(cikePeerCorrEntry.CikePeerCorrLocalValue, "cikePeerCorrLocalValue") + types.AddKeyToken(cikePeerCorrEntry.CikePeerCorrRemoteType, "cikePeerCorrRemoteType") + types.AddKeyToken(cikePeerCorrEntry.CikePeerCorrRemoteValue, "cikePeerCorrRemoteValue") + types.AddKeyToken(cikePeerCorrEntry.CikePeerCorrIntIndex, "cikePeerCorrIntIndex") + types.AddKeyToken(cikePeerCorrEntry.CikePeerCorrSeqNum, "cikePeerCorrSeqNum")
    cikePeerCorrEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cikePeerCorrEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cikePeerCorrEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cikePeerCorrEntry.EntityData.Children = types.NewOrderedMap()
    cikePeerCorrEntry.EntityData.Leafs = types.NewOrderedMap()
    cikePeerCorrEntry.EntityData.Leafs.Append("cikePeerCorrLocalType", types.YLeaf{"CikePeerCorrLocalType", cikePeerCorrEntry.CikePeerCorrLocalType})
    cikePeerCorrEntry.EntityData.Leafs.Append("cikePeerCorrLocalValue", types.YLeaf{"CikePeerCorrLocalValue", cikePeerCorrEntry.CikePeerCorrLocalValue})
    cikePeerCorrEntry.EntityData.Leafs.Append("cikePeerCorrRemoteType", types.YLeaf{"CikePeerCorrRemoteType", cikePeerCorrEntry.CikePeerCorrRemoteType})
    cikePeerCorrEntry.EntityData.Leafs.Append("cikePeerCorrRemoteValue", types.YLeaf{"CikePeerCorrRemoteValue", cikePeerCorrEntry.CikePeerCorrRemoteValue})
    cikePeerCorrEntry.EntityData.Leafs.Append("cikePeerCorrIntIndex", types.YLeaf{"CikePeerCorrIntIndex", cikePeerCorrEntry.CikePeerCorrIntIndex})
    cikePeerCorrEntry.EntityData.Leafs.Append("cikePeerCorrSeqNum", types.YLeaf{"CikePeerCorrSeqNum", cikePeerCorrEntry.CikePeerCorrSeqNum})
    cikePeerCorrEntry.EntityData.Leafs.Append("cikePeerCorrIpSecTunIndex", types.YLeaf{"CikePeerCorrIpSecTunIndex", cikePeerCorrEntry.CikePeerCorrIpSecTunIndex})

    cikePeerCorrEntry.EntityData.YListKeys = []string {"CikePeerCorrLocalType", "CikePeerCorrLocalValue", "CikePeerCorrRemoteType", "CikePeerCorrRemoteValue", "CikePeerCorrIntIndex", "CikePeerCorrSeqNum"}

    return &(cikePeerCorrEntry.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CikePhase1GWStatsTable
// Phase-1 IKE stats information is included in this table.
// Each entry is related to a specific gateway which is 
// identified by 'cmgwIndex'.
type CISCOIPSECFLOWMONITORMIB_CikePhase1GWStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes of an Phase-1 IKE stats information for
    // the related gateway.  There is only one entry for each gateway. The entry 
    // is created when a gateway up and cannot be deleted. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_CikePhase1GWStatsTable_CikePhase1GWStatsEntry.
    CikePhase1GWStatsEntry []*CISCOIPSECFLOWMONITORMIB_CikePhase1GWStatsTable_CikePhase1GWStatsEntry
}

func (cikePhase1GWStatsTable *CISCOIPSECFLOWMONITORMIB_CikePhase1GWStatsTable) GetEntityData() *types.CommonEntityData {
    cikePhase1GWStatsTable.EntityData.YFilter = cikePhase1GWStatsTable.YFilter
    cikePhase1GWStatsTable.EntityData.YangName = "cikePhase1GWStatsTable"
    cikePhase1GWStatsTable.EntityData.BundleName = "cisco_ios_xe"
    cikePhase1GWStatsTable.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cikePhase1GWStatsTable.EntityData.SegmentPath = "cikePhase1GWStatsTable"
    cikePhase1GWStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cikePhase1GWStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cikePhase1GWStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cikePhase1GWStatsTable.EntityData.Children = types.NewOrderedMap()
    cikePhase1GWStatsTable.EntityData.Children.Append("cikePhase1GWStatsEntry", types.YChild{"CikePhase1GWStatsEntry", nil})
    for i := range cikePhase1GWStatsTable.CikePhase1GWStatsEntry {
        cikePhase1GWStatsTable.EntityData.Children.Append(types.GetSegmentPath(cikePhase1GWStatsTable.CikePhase1GWStatsEntry[i]), types.YChild{"CikePhase1GWStatsEntry", cikePhase1GWStatsTable.CikePhase1GWStatsEntry[i]})
    }
    cikePhase1GWStatsTable.EntityData.Leafs = types.NewOrderedMap()

    cikePhase1GWStatsTable.EntityData.YListKeys = []string {}

    return &(cikePhase1GWStatsTable.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CikePhase1GWStatsTable_CikePhase1GWStatsEntry
// Each entry contains the attributes of an Phase-1 IKE stats
// information for the related gateway.
// 
// There is only one entry for each gateway. The entry 
// is created when a gateway up and cannot be deleted.
type CISCOIPSECFLOWMONITORMIB_CikePhase1GWStatsTable_CikePhase1GWStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_media_gateway_mib.CISCOMEDIAGATEWAYMIB_CMediaGwTable_CMediaGwEntry_CmgwIndex
    CmgwIndex interface{}

    // The number of currently active IPsec Phase-1 IKE Tunnels. The type is
    // interface{} with range: 0..4294967295.
    CikePhase1GWActiveTunnels interface{}

    // The total number of previously active IPsec Phase-1 IKE Tunnels. The type
    // is interface{} with range: 0..4294967295. Units are SAs.
    CikePhase1GWPreviousTunnels interface{}

    // The total number of octets received by all currently and previously active
    // IPsec Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Octets.
    CikePhase1GWInOctets interface{}

    // The total number of packets received by all currently and previously active
    // IPsec  Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    CikePhase1GWInPkts interface{}

    // The total number of packets which were dropped during receive processing by
    // all  currently and previously active IPsec Phase-1 IKE Tunnels. The type is
    // interface{} with range: 0..4294967295. Units are Packets.
    CikePhase1GWInDropPkts interface{}

    // The total number of notifys received by all currently and previously active
    // IPsec  Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Notification Payloads.
    CikePhase1GWInNotifys interface{}

    // The total number of IPsec Phase-2 exchanges received by all currently and
    // previously  active IPsec Phase-1 IKE Tunnels. The type is interface{} with
    // range: 0..4294967295. Units are SA Payloads.
    CikePhase1GWInP2Exchgs interface{}

    // The total number of IPsec Phase-2 exchanges which were received and found
    // to be invalid  by all currently and previously active IPsec  Phase-1 IKE
    // Tunnels. The type is interface{} with range: 0..4294967295. Units are SA
    // Payloads.
    CikePhase1GWInP2ExchgInvalids interface{}

    // The total number of IPsec Phase-2 exchanges which were received and
    // rejected by all  currently and previously active IPsec Phase-1  IKE
    // Tunnels. The type is interface{} with range: 0..4294967295. Units are SA
    // Payloads.
    CikePhase1GWInP2ExchgRejects interface{}

    // The total number of IPsec Phase-2 'Security Association' delete requests
    // received by all  currently and previously active and IPsec  Phase-1 IKE
    // Tunnels. The type is interface{} with range: 0..4294967295. Units are
    // Notification Payloads.
    CikePhase1GWInP2SaDelRequests interface{}

    // The total number of octets sent by all currently and previously active and
    // IPsec Phase-1  IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Octets.
    CikePhase1GWOutOctets interface{}

    // The total number of packets sent by all currently and previously active and
    // IPsec Phase-1  Tunnels. The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    CikePhase1GWOutPkts interface{}

    // The total number of packets which were dropped during send processing by
    // all currently  and previously active IPsec Phase-1 IKE Tunnels. The type is
    // interface{} with range: 0..4294967295. Units are Packets.
    CikePhase1GWOutDropPkts interface{}

    // The total number of notifys sent by all currently and previously active
    // IPsec Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Notification Payloads.
    CikePhase1GWOutNotifys interface{}

    // The total number of IPsec Phase-2 exchanges which were sent by all
    // currently and previously  active IPsec Phase-1 IKE Tunnels. The type is
    // interface{} with range: 0..4294967295. Units are SA Payloads.
    CikePhase1GWOutP2Exchgs interface{}

    // The total number of IPsec Phase-2 exchanges which were sent and found to be
    // invalid by  all currently and previously active IPsec Phase-1  Tunnels. The
    // type is interface{} with range: 0..4294967295. Units are SA Payloads.
    CikePhase1GWOutP2ExchgInvalids interface{}

    // The total number of IPsec Phase-2 exchanges which were sent and rejected by
    // all currently and previously active IPsec Phase-1 IKE Tunnels. The type is
    // interface{} with range: 0..4294967295. Units are SA Payloads.
    CikePhase1GWOutP2ExchgRejects interface{}

    // The total number of IPsec Phase-2 SA delete requests sent by all currently
    // and  previously active IPsec Phase-1 IKE Tunnels. The type is interface{}
    // with range: 0..4294967295. Units are Notification Payloads.
    CikePhase1GWOutP2SaDelRequests interface{}

    // The total number of IPsec Phase-1 IKE Tunnels which were locally initiated.
    // The type is interface{} with range: 0..4294967295. Units are SAs.
    CikePhase1GWInitTunnels interface{}

    // The total number of IPsec Phase-1 IKE Tunnels which were locally initiated
    // and failed to activate. The type is interface{} with range: 0..4294967295.
    // Units are SAs.
    CikePhase1GWInitTunnelFails interface{}

    // The total number of IPsec Phase-1 IKE Tunnels which were remotely initiated
    // and failed to activate. The type is interface{} with range: 0..4294967295.
    // Units are SAs.
    CikePhase1GWRespTunnelFails interface{}

    // The total number of system capacity failures which occurred during
    // processing of all current  and previously active IPsec Phase-1 IKE Tunnels.
    // The type is interface{} with range: 0..4294967295. Units are Failures.
    CikePhase1GWSysCapFails interface{}

    // The total number of authentications which ended in failure by all current
    // and previous IPsec Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Failures.
    CikePhase1GWAuthFails interface{}

    // The total number of decryptions which ended in failure by all current and
    // previous IPsec Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Failures.
    CikePhase1GWDecryptFails interface{}

    // The total number of hash validations which ended in failure by all current
    // and previous IPsec Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Failures.
    CikePhase1GWHashValidFails interface{}

    // The total number of non-existent 'Security Association' failures occurred
    // during processing of current and  previous IPsec Phase-1 IKE Tunnels. The
    // type is interface{} with range: 0..4294967295. Units are Failures.
    CikePhase1GWNoSaFails interface{}
}

func (cikePhase1GWStatsEntry *CISCOIPSECFLOWMONITORMIB_CikePhase1GWStatsTable_CikePhase1GWStatsEntry) GetEntityData() *types.CommonEntityData {
    cikePhase1GWStatsEntry.EntityData.YFilter = cikePhase1GWStatsEntry.YFilter
    cikePhase1GWStatsEntry.EntityData.YangName = "cikePhase1GWStatsEntry"
    cikePhase1GWStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    cikePhase1GWStatsEntry.EntityData.ParentYangName = "cikePhase1GWStatsTable"
    cikePhase1GWStatsEntry.EntityData.SegmentPath = "cikePhase1GWStatsEntry" + types.AddKeyToken(cikePhase1GWStatsEntry.CmgwIndex, "cmgwIndex")
    cikePhase1GWStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cikePhase1GWStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cikePhase1GWStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cikePhase1GWStatsEntry.EntityData.Children = types.NewOrderedMap()
    cikePhase1GWStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    cikePhase1GWStatsEntry.EntityData.Leafs.Append("cmgwIndex", types.YLeaf{"CmgwIndex", cikePhase1GWStatsEntry.CmgwIndex})
    cikePhase1GWStatsEntry.EntityData.Leafs.Append("cikePhase1GWActiveTunnels", types.YLeaf{"CikePhase1GWActiveTunnels", cikePhase1GWStatsEntry.CikePhase1GWActiveTunnels})
    cikePhase1GWStatsEntry.EntityData.Leafs.Append("cikePhase1GWPreviousTunnels", types.YLeaf{"CikePhase1GWPreviousTunnels", cikePhase1GWStatsEntry.CikePhase1GWPreviousTunnels})
    cikePhase1GWStatsEntry.EntityData.Leafs.Append("cikePhase1GWInOctets", types.YLeaf{"CikePhase1GWInOctets", cikePhase1GWStatsEntry.CikePhase1GWInOctets})
    cikePhase1GWStatsEntry.EntityData.Leafs.Append("cikePhase1GWInPkts", types.YLeaf{"CikePhase1GWInPkts", cikePhase1GWStatsEntry.CikePhase1GWInPkts})
    cikePhase1GWStatsEntry.EntityData.Leafs.Append("cikePhase1GWInDropPkts", types.YLeaf{"CikePhase1GWInDropPkts", cikePhase1GWStatsEntry.CikePhase1GWInDropPkts})
    cikePhase1GWStatsEntry.EntityData.Leafs.Append("cikePhase1GWInNotifys", types.YLeaf{"CikePhase1GWInNotifys", cikePhase1GWStatsEntry.CikePhase1GWInNotifys})
    cikePhase1GWStatsEntry.EntityData.Leafs.Append("cikePhase1GWInP2Exchgs", types.YLeaf{"CikePhase1GWInP2Exchgs", cikePhase1GWStatsEntry.CikePhase1GWInP2Exchgs})
    cikePhase1GWStatsEntry.EntityData.Leafs.Append("cikePhase1GWInP2ExchgInvalids", types.YLeaf{"CikePhase1GWInP2ExchgInvalids", cikePhase1GWStatsEntry.CikePhase1GWInP2ExchgInvalids})
    cikePhase1GWStatsEntry.EntityData.Leafs.Append("cikePhase1GWInP2ExchgRejects", types.YLeaf{"CikePhase1GWInP2ExchgRejects", cikePhase1GWStatsEntry.CikePhase1GWInP2ExchgRejects})
    cikePhase1GWStatsEntry.EntityData.Leafs.Append("cikePhase1GWInP2SaDelRequests", types.YLeaf{"CikePhase1GWInP2SaDelRequests", cikePhase1GWStatsEntry.CikePhase1GWInP2SaDelRequests})
    cikePhase1GWStatsEntry.EntityData.Leafs.Append("cikePhase1GWOutOctets", types.YLeaf{"CikePhase1GWOutOctets", cikePhase1GWStatsEntry.CikePhase1GWOutOctets})
    cikePhase1GWStatsEntry.EntityData.Leafs.Append("cikePhase1GWOutPkts", types.YLeaf{"CikePhase1GWOutPkts", cikePhase1GWStatsEntry.CikePhase1GWOutPkts})
    cikePhase1GWStatsEntry.EntityData.Leafs.Append("cikePhase1GWOutDropPkts", types.YLeaf{"CikePhase1GWOutDropPkts", cikePhase1GWStatsEntry.CikePhase1GWOutDropPkts})
    cikePhase1GWStatsEntry.EntityData.Leafs.Append("cikePhase1GWOutNotifys", types.YLeaf{"CikePhase1GWOutNotifys", cikePhase1GWStatsEntry.CikePhase1GWOutNotifys})
    cikePhase1GWStatsEntry.EntityData.Leafs.Append("cikePhase1GWOutP2Exchgs", types.YLeaf{"CikePhase1GWOutP2Exchgs", cikePhase1GWStatsEntry.CikePhase1GWOutP2Exchgs})
    cikePhase1GWStatsEntry.EntityData.Leafs.Append("cikePhase1GWOutP2ExchgInvalids", types.YLeaf{"CikePhase1GWOutP2ExchgInvalids", cikePhase1GWStatsEntry.CikePhase1GWOutP2ExchgInvalids})
    cikePhase1GWStatsEntry.EntityData.Leafs.Append("cikePhase1GWOutP2ExchgRejects", types.YLeaf{"CikePhase1GWOutP2ExchgRejects", cikePhase1GWStatsEntry.CikePhase1GWOutP2ExchgRejects})
    cikePhase1GWStatsEntry.EntityData.Leafs.Append("cikePhase1GWOutP2SaDelRequests", types.YLeaf{"CikePhase1GWOutP2SaDelRequests", cikePhase1GWStatsEntry.CikePhase1GWOutP2SaDelRequests})
    cikePhase1GWStatsEntry.EntityData.Leafs.Append("cikePhase1GWInitTunnels", types.YLeaf{"CikePhase1GWInitTunnels", cikePhase1GWStatsEntry.CikePhase1GWInitTunnels})
    cikePhase1GWStatsEntry.EntityData.Leafs.Append("cikePhase1GWInitTunnelFails", types.YLeaf{"CikePhase1GWInitTunnelFails", cikePhase1GWStatsEntry.CikePhase1GWInitTunnelFails})
    cikePhase1GWStatsEntry.EntityData.Leafs.Append("cikePhase1GWRespTunnelFails", types.YLeaf{"CikePhase1GWRespTunnelFails", cikePhase1GWStatsEntry.CikePhase1GWRespTunnelFails})
    cikePhase1GWStatsEntry.EntityData.Leafs.Append("cikePhase1GWSysCapFails", types.YLeaf{"CikePhase1GWSysCapFails", cikePhase1GWStatsEntry.CikePhase1GWSysCapFails})
    cikePhase1GWStatsEntry.EntityData.Leafs.Append("cikePhase1GWAuthFails", types.YLeaf{"CikePhase1GWAuthFails", cikePhase1GWStatsEntry.CikePhase1GWAuthFails})
    cikePhase1GWStatsEntry.EntityData.Leafs.Append("cikePhase1GWDecryptFails", types.YLeaf{"CikePhase1GWDecryptFails", cikePhase1GWStatsEntry.CikePhase1GWDecryptFails})
    cikePhase1GWStatsEntry.EntityData.Leafs.Append("cikePhase1GWHashValidFails", types.YLeaf{"CikePhase1GWHashValidFails", cikePhase1GWStatsEntry.CikePhase1GWHashValidFails})
    cikePhase1GWStatsEntry.EntityData.Leafs.Append("cikePhase1GWNoSaFails", types.YLeaf{"CikePhase1GWNoSaFails", cikePhase1GWStatsEntry.CikePhase1GWNoSaFails})

    cikePhase1GWStatsEntry.EntityData.YListKeys = []string {"CmgwIndex"}

    return &(cikePhase1GWStatsEntry.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CipSecTunnelTable
// The IPsec Phase-2 Tunnel Table.
// There is one entry in this table for 
// each active IPsec Phase-2 Tunnel.
type CISCOIPSECFLOWMONITORMIB_CipSecTunnelTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with an active IPsec Phase-2
    // Tunnel. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_CipSecTunnelTable_CipSecTunnelEntry.
    CipSecTunnelEntry []*CISCOIPSECFLOWMONITORMIB_CipSecTunnelTable_CipSecTunnelEntry
}

func (cipSecTunnelTable *CISCOIPSECFLOWMONITORMIB_CipSecTunnelTable) GetEntityData() *types.CommonEntityData {
    cipSecTunnelTable.EntityData.YFilter = cipSecTunnelTable.YFilter
    cipSecTunnelTable.EntityData.YangName = "cipSecTunnelTable"
    cipSecTunnelTable.EntityData.BundleName = "cisco_ios_xe"
    cipSecTunnelTable.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cipSecTunnelTable.EntityData.SegmentPath = "cipSecTunnelTable"
    cipSecTunnelTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipSecTunnelTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipSecTunnelTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipSecTunnelTable.EntityData.Children = types.NewOrderedMap()
    cipSecTunnelTable.EntityData.Children.Append("cipSecTunnelEntry", types.YChild{"CipSecTunnelEntry", nil})
    for i := range cipSecTunnelTable.CipSecTunnelEntry {
        cipSecTunnelTable.EntityData.Children.Append(types.GetSegmentPath(cipSecTunnelTable.CipSecTunnelEntry[i]), types.YChild{"CipSecTunnelEntry", cipSecTunnelTable.CipSecTunnelEntry[i]})
    }
    cipSecTunnelTable.EntityData.Leafs = types.NewOrderedMap()

    cipSecTunnelTable.EntityData.YListKeys = []string {}

    return &(cipSecTunnelTable.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CipSecTunnelTable_CipSecTunnelEntry
// Each entry contains the attributes
// associated with an active IPsec Phase-2 Tunnel.
type CISCOIPSECFLOWMONITORMIB_CipSecTunnelTable_CipSecTunnelEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index of the IPsec Phase-2 Tunnel Table. The
    // value of the index is a number which begins  at one and is incremented with
    // each tunnel that  is created. The value of this object will wrap  at
    // 2,147,483,647. The type is interface{} with range: 1..2147483647.
    CipSecTunIndex interface{}

    // The index of the associated IPsec Phase-1 IKE Tunnel.  (cikeTunIndex in the
    // cikeTunnelTable). The type is interface{} with range: 1..2147483647.
    CipSecTunIkeTunnelIndex interface{}

    // An indicator which specifies whether or not the IPsec Phase-1 IKE Tunnel
    // currently exists. The type is bool.
    CipSecTunIkeTunnelAlive interface{}

    // The IP address of the local endpoint for the IPsec Phase-2 Tunnel. The type
    // is string with length: 4 | 16.
    CipSecTunLocalAddr interface{}

    // The IP address of the remote endpoint for the IPsec Phase-2 Tunnel. The
    // type is string with length: 4 | 16.
    CipSecTunRemoteAddr interface{}

    // The type of key used by the IPsec Phase-2 Tunnel. The type is KeyType.
    CipSecTunKeyType interface{}

    // The encapsulation mode used by the IPsec Phase-2 Tunnel. The type is
    // EncapMode.
    CipSecTunEncapMode interface{}

    // The negotiated LifeSize of the IPsec Phase-2 Tunnel in kilobytes. The type
    // is interface{} with range: 1..2147483647. Units are KBytes.
    CipSecTunLifeSize interface{}

    // The negotiated LifeTime of the IPsec Phase-2 Tunnel in seconds. The type is
    // interface{} with range: 1..2147483647. Units are Seconds.
    CipSecTunLifeTime interface{}

    // The length of time the IPsec Phase-2 Tunnel has been  active in hundredths
    // of seconds. The type is interface{} with range: 0..2147483647.
    CipSecTunActiveTime interface{}

    // The security association LifeSize refresh threshold in kilobytes. The type
    // is interface{} with range: 1..2147483647. Units are KBytes.
    CipSecTunSaLifeSizeThreshold interface{}

    // The security association LifeTime refresh threshold in seconds. The type is
    // interface{} with range: 1..2147483647. Units are Seconds.
    CipSecTunSaLifeTimeThreshold interface{}

    // The total number of security association refreshes performed. The type is
    // interface{} with range: 0..4294967295. Units are QM Exchanges.
    CipSecTunTotalRefreshes interface{}

    // The total number of security associations which have expired. The type is
    // interface{} with range: 0..4294967295. Units are SAs.
    CipSecTunExpiredSaInstances interface{}

    // The number of security associations which are currently active or expiring.
    // The type is interface{} with range: 0..4294967295.
    CipSecTunCurrentSaInstances interface{}

    // The Diffie Hellman Group used by the inbound security association of the 
    // IPsec Phase-2 Tunnel. The type is DiffHellmanGrp.
    CipSecTunInSaDiffHellmanGrp interface{}

    // The encryption algorithm used by the inbound security association of the
    // IPsec Phase-2 Tunnel. The type is EncryptAlgo.
    CipSecTunInSaEncryptAlgo interface{}

    // The authentication algorithm used by the inbound authentication header (AH)
    // security association of the IPsec Phase-2 Tunnel. The type is AuthAlgo.
    CipSecTunInSaAhAuthAlgo interface{}

    // The authentication algorithm used by the inbound encapsulation security
    // protocol (ESP) security  association of the IPsec Phase-2 Tunnel. The type
    // is AuthAlgo.
    CipSecTunInSaEspAuthAlgo interface{}

    // The decompression algorithm used by the inbound security association of the
    // IPsec Phase-2 Tunnel. The type is CompAlgo.
    CipSecTunInSaDecompAlgo interface{}

    // The Diffie Hellman Group used by the outbound security association of the
    // IPsec Phase-2 Tunnel. The type is DiffHellmanGrp.
    CipSecTunOutSaDiffHellmanGrp interface{}

    // The encryption algorithm used by the outbound security association of the
    // IPsec Phase-2 Tunnel. The type is EncryptAlgo.
    CipSecTunOutSaEncryptAlgo interface{}

    // The authentication algorithm used by the outbound authentication header
    // (AH) security association of the IPsec Phase-2 Tunnel. The type is
    // AuthAlgo.
    CipSecTunOutSaAhAuthAlgo interface{}

    // The authentication algorithm used by the inbound encapsulation security
    // protocol (ESP)  security association of the IPsec Phase-2 Tunnel. The type
    // is AuthAlgo.
    CipSecTunOutSaEspAuthAlgo interface{}

    // The compression algorithm used by the inbound security association of the
    // IPsec Phase-2 Tunnel. The type is CompAlgo.
    CipSecTunOutSaCompAlgo interface{}

    // The total number of octets received by this IPsec Phase-2 Tunnel.  This
    // value is accumulated BEFORE determining whether or not the packet should be
    // decompressed.  See also cipSecTunInOctWraps for the number of times this
    // counter has wrapped. The type is interface{} with range: 0..4294967295.
    // Units are Octets.
    CipSecTunInOctets interface{}

    // A high capacity count of the total number of octets received by this IPsec
    // Phase-2 Tunnel.  This value is accumulated BEFORE determining whether or
    // not the packet should be decompressed. The type is interface{} with range:
    // 0..18446744073709551615. Units are Octets.
    CipSecTunHcInOctets interface{}

    // The number of times the octets received counter (cipSecTunInOctets) has
    // wrapped. The type is interface{} with range: 0..4294967295. Units are
    // Integral units.
    CipSecTunInOctWraps interface{}

    // The total number of decompressed octets received by this IPsec Phase-2
    // Tunnel. This value is  accumulated AFTER the packet is decompressed.  If
    // compression is not being  used, this value will match the value of  
    // cipSecTunInOctets.  See also cipSecTunInDecompOctWraps   for the number of
    // times  this counter has wrapped. The type is interface{} with range:
    // 0..4294967295. Units are Octets.
    CipSecTunInDecompOctets interface{}

    // A high capacity count of the total number of decompressed octets received
    // by this IPsec Phase-2 Tunnel.  This value is accumulated AFTER the packet
    // is decompressed. If compression is not being used, this value will match
    // the value of cipSecTunHcInOctets. The type is interface{} with range:
    // 0..18446744073709551615.
    CipSecTunHcInDecompOctets interface{}

    // The number of times the decompressed octets received counter 
    // (cipSecTunInDecompOctets) has wrapped. The type is interface{} with range:
    // 0..4294967295. Units are Integral units.
    CipSecTunInDecompOctWraps interface{}

    // The total number of packets received by this IPsec Phase-2 Tunnel. The type
    // is interface{} with range: 0..4294967295. Units are Packets.
    CipSecTunInPkts interface{}

    // The total number of packets dropped during receive processing by this IPsec
    // Phase-2  Tunnel. This count does NOT include  packets dropped due to
    // Anti-Replay processing. The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    CipSecTunInDropPkts interface{}

    // The total number of packets dropped during receive processing due to
    // Anti-Replay processing  by this IPsec Phase-2 Tunnel. The type is
    // interface{} with range: 0..4294967295. Units are Packets.
    CipSecTunInReplayDropPkts interface{}

    // The total number of inbound authentication's performed by this  IPsec
    // Phase-2 Tunnel. The type is interface{} with range: 0..4294967295. Units
    // are Events.
    CipSecTunInAuths interface{}

    // The total number of inbound authentication's which ended in  failure by
    // this IPsec Phase-2 Tunnel . The type is interface{} with range:
    // 0..4294967295. Units are Failures.
    CipSecTunInAuthFails interface{}

    // The total number of inbound decryption's performed by this IPsec Phase-2
    // Tunnel. The type is interface{} with range: 0..4294967295. Units are
    // Packets.
    CipSecTunInDecrypts interface{}

    // The total number of inbound decryption's which ended in failure  by this
    // IPsec Phase-2 Tunnel. The type is interface{} with range: 0..4294967295.
    // Units are Failures.
    CipSecTunInDecryptFails interface{}

    // The total number of octets sent by this IPsec Phase-2 Tunnel.  This value
    // is accumulated AFTER determining whether or not the packet should  be
    // compressed.  See also cipSecTunOutOctWraps for the number of times this
    // counter has wrapped. The type is interface{} with range: 0..4294967295.
    // Units are Octets.
    CipSecTunOutOctets interface{}

    // A high capacity count of the total number of octets sent by this IPsec
    // Phase-2 Tunnel.  This value is accumulated AFTER determining whether or not
    // the  packet should be compressed. The type is interface{} with range:
    // 0..18446744073709551615.
    CipSecTunHcOutOctets interface{}

    // The number of times the out octets counter (cipSecTunOutOctets) has
    // wrapped. The type is interface{} with range: 0..4294967295. Units are
    // Integral units.
    CipSecTunOutOctWraps interface{}

    // The total number of uncompressed octets sent by this IPsec Phase-2 Tunnel. 
    // This value  is accumulated BEFORE the packet is compressed.  If compression
    // is not being used, this value  will match the value of cipSecTunOutOctets. 
    // See also cipSecTunOutDecompOctWraps for the   number of times this counter
    // has wrapped. The type is interface{} with range: 0..4294967295. Units are
    // Octets.
    CipSecTunOutUncompOctets interface{}

    // A high capacity count of the total number of uncompressed octets sent by
    // this IPsec  Phase-2 Tunnel.  This value is accumulated BEFORE  the packet
    // is compressed. If compression  is not being used, this value will match the
    // value  of cipSecTunHcOutOctets. The type is interface{} with range:
    // 0..18446744073709551615.
    CipSecTunHcOutUncompOctets interface{}

    // The number of times the uncompressed octets sent counter
    // (cipSecTunOutUncompOctets) has wrapped. The type is interface{} with range:
    // 0..4294967295. Units are Integral units.
    CipSecTunOutUncompOctWraps interface{}

    // The total number of packets sent by this IPsec Phase-2 Tunnel. The type is
    // interface{} with range: 0..4294967295. Units are Packets.
    CipSecTunOutPkts interface{}

    // The total number of packets dropped during send processing by this IPsec
    // Phase-2 Tunnel. The type is interface{} with range: 0..4294967295. Units
    // are Packets.
    CipSecTunOutDropPkts interface{}

    // The total number of outbound authentication's performed by this IPsec
    // Phase-2 Tunnel. The type is interface{} with range: 0..4294967295. Units
    // are Events.
    CipSecTunOutAuths interface{}

    // The total number of outbound authentication's which ended in failure  by
    // this IPsec Phase-2 Tunnel. The type is interface{} with range:
    // 0..4294967295. Units are Failures.
    CipSecTunOutAuthFails interface{}

    // The total number of outbound encryption's performed by this IPsec Phase-2
    // Tunnel. The type is interface{} with range: 0..4294967295. Units are
    // Packets.
    CipSecTunOutEncrypts interface{}

    // The total number of outbound encryption's which ended in failure by this
    // IPsec Phase-2 Tunnel. The type is interface{} with range: 0..4294967295.
    // Units are Failures.
    CipSecTunOutEncryptFails interface{}

    // The status of the MIB table row.  This object can be used to bring the
    // tunnel down by setting value of this object to destroy(2). When the value
    // is set to destroy(2), the SA bundle is destroyed and this row is deleted
    // from this table.  When this MIB value is queried, the value of active(1) is
    // always returned, if the instance  exists.  This object cannot be used to
    // create a MIB  table row. The type is TunnelStatus.
    CipSecTunStatus interface{}
}

func (cipSecTunnelEntry *CISCOIPSECFLOWMONITORMIB_CipSecTunnelTable_CipSecTunnelEntry) GetEntityData() *types.CommonEntityData {
    cipSecTunnelEntry.EntityData.YFilter = cipSecTunnelEntry.YFilter
    cipSecTunnelEntry.EntityData.YangName = "cipSecTunnelEntry"
    cipSecTunnelEntry.EntityData.BundleName = "cisco_ios_xe"
    cipSecTunnelEntry.EntityData.ParentYangName = "cipSecTunnelTable"
    cipSecTunnelEntry.EntityData.SegmentPath = "cipSecTunnelEntry" + types.AddKeyToken(cipSecTunnelEntry.CipSecTunIndex, "cipSecTunIndex")
    cipSecTunnelEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipSecTunnelEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipSecTunnelEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipSecTunnelEntry.EntityData.Children = types.NewOrderedMap()
    cipSecTunnelEntry.EntityData.Leafs = types.NewOrderedMap()
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunIndex", types.YLeaf{"CipSecTunIndex", cipSecTunnelEntry.CipSecTunIndex})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunIkeTunnelIndex", types.YLeaf{"CipSecTunIkeTunnelIndex", cipSecTunnelEntry.CipSecTunIkeTunnelIndex})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunIkeTunnelAlive", types.YLeaf{"CipSecTunIkeTunnelAlive", cipSecTunnelEntry.CipSecTunIkeTunnelAlive})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunLocalAddr", types.YLeaf{"CipSecTunLocalAddr", cipSecTunnelEntry.CipSecTunLocalAddr})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunRemoteAddr", types.YLeaf{"CipSecTunRemoteAddr", cipSecTunnelEntry.CipSecTunRemoteAddr})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunKeyType", types.YLeaf{"CipSecTunKeyType", cipSecTunnelEntry.CipSecTunKeyType})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunEncapMode", types.YLeaf{"CipSecTunEncapMode", cipSecTunnelEntry.CipSecTunEncapMode})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunLifeSize", types.YLeaf{"CipSecTunLifeSize", cipSecTunnelEntry.CipSecTunLifeSize})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunLifeTime", types.YLeaf{"CipSecTunLifeTime", cipSecTunnelEntry.CipSecTunLifeTime})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunActiveTime", types.YLeaf{"CipSecTunActiveTime", cipSecTunnelEntry.CipSecTunActiveTime})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunSaLifeSizeThreshold", types.YLeaf{"CipSecTunSaLifeSizeThreshold", cipSecTunnelEntry.CipSecTunSaLifeSizeThreshold})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunSaLifeTimeThreshold", types.YLeaf{"CipSecTunSaLifeTimeThreshold", cipSecTunnelEntry.CipSecTunSaLifeTimeThreshold})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunTotalRefreshes", types.YLeaf{"CipSecTunTotalRefreshes", cipSecTunnelEntry.CipSecTunTotalRefreshes})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunExpiredSaInstances", types.YLeaf{"CipSecTunExpiredSaInstances", cipSecTunnelEntry.CipSecTunExpiredSaInstances})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunCurrentSaInstances", types.YLeaf{"CipSecTunCurrentSaInstances", cipSecTunnelEntry.CipSecTunCurrentSaInstances})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunInSaDiffHellmanGrp", types.YLeaf{"CipSecTunInSaDiffHellmanGrp", cipSecTunnelEntry.CipSecTunInSaDiffHellmanGrp})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunInSaEncryptAlgo", types.YLeaf{"CipSecTunInSaEncryptAlgo", cipSecTunnelEntry.CipSecTunInSaEncryptAlgo})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunInSaAhAuthAlgo", types.YLeaf{"CipSecTunInSaAhAuthAlgo", cipSecTunnelEntry.CipSecTunInSaAhAuthAlgo})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunInSaEspAuthAlgo", types.YLeaf{"CipSecTunInSaEspAuthAlgo", cipSecTunnelEntry.CipSecTunInSaEspAuthAlgo})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunInSaDecompAlgo", types.YLeaf{"CipSecTunInSaDecompAlgo", cipSecTunnelEntry.CipSecTunInSaDecompAlgo})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunOutSaDiffHellmanGrp", types.YLeaf{"CipSecTunOutSaDiffHellmanGrp", cipSecTunnelEntry.CipSecTunOutSaDiffHellmanGrp})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunOutSaEncryptAlgo", types.YLeaf{"CipSecTunOutSaEncryptAlgo", cipSecTunnelEntry.CipSecTunOutSaEncryptAlgo})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunOutSaAhAuthAlgo", types.YLeaf{"CipSecTunOutSaAhAuthAlgo", cipSecTunnelEntry.CipSecTunOutSaAhAuthAlgo})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunOutSaEspAuthAlgo", types.YLeaf{"CipSecTunOutSaEspAuthAlgo", cipSecTunnelEntry.CipSecTunOutSaEspAuthAlgo})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunOutSaCompAlgo", types.YLeaf{"CipSecTunOutSaCompAlgo", cipSecTunnelEntry.CipSecTunOutSaCompAlgo})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunInOctets", types.YLeaf{"CipSecTunInOctets", cipSecTunnelEntry.CipSecTunInOctets})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunHcInOctets", types.YLeaf{"CipSecTunHcInOctets", cipSecTunnelEntry.CipSecTunHcInOctets})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunInOctWraps", types.YLeaf{"CipSecTunInOctWraps", cipSecTunnelEntry.CipSecTunInOctWraps})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunInDecompOctets", types.YLeaf{"CipSecTunInDecompOctets", cipSecTunnelEntry.CipSecTunInDecompOctets})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunHcInDecompOctets", types.YLeaf{"CipSecTunHcInDecompOctets", cipSecTunnelEntry.CipSecTunHcInDecompOctets})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunInDecompOctWraps", types.YLeaf{"CipSecTunInDecompOctWraps", cipSecTunnelEntry.CipSecTunInDecompOctWraps})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunInPkts", types.YLeaf{"CipSecTunInPkts", cipSecTunnelEntry.CipSecTunInPkts})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunInDropPkts", types.YLeaf{"CipSecTunInDropPkts", cipSecTunnelEntry.CipSecTunInDropPkts})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunInReplayDropPkts", types.YLeaf{"CipSecTunInReplayDropPkts", cipSecTunnelEntry.CipSecTunInReplayDropPkts})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunInAuths", types.YLeaf{"CipSecTunInAuths", cipSecTunnelEntry.CipSecTunInAuths})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunInAuthFails", types.YLeaf{"CipSecTunInAuthFails", cipSecTunnelEntry.CipSecTunInAuthFails})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunInDecrypts", types.YLeaf{"CipSecTunInDecrypts", cipSecTunnelEntry.CipSecTunInDecrypts})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunInDecryptFails", types.YLeaf{"CipSecTunInDecryptFails", cipSecTunnelEntry.CipSecTunInDecryptFails})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunOutOctets", types.YLeaf{"CipSecTunOutOctets", cipSecTunnelEntry.CipSecTunOutOctets})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunHcOutOctets", types.YLeaf{"CipSecTunHcOutOctets", cipSecTunnelEntry.CipSecTunHcOutOctets})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunOutOctWraps", types.YLeaf{"CipSecTunOutOctWraps", cipSecTunnelEntry.CipSecTunOutOctWraps})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunOutUncompOctets", types.YLeaf{"CipSecTunOutUncompOctets", cipSecTunnelEntry.CipSecTunOutUncompOctets})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunHcOutUncompOctets", types.YLeaf{"CipSecTunHcOutUncompOctets", cipSecTunnelEntry.CipSecTunHcOutUncompOctets})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunOutUncompOctWraps", types.YLeaf{"CipSecTunOutUncompOctWraps", cipSecTunnelEntry.CipSecTunOutUncompOctWraps})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunOutPkts", types.YLeaf{"CipSecTunOutPkts", cipSecTunnelEntry.CipSecTunOutPkts})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunOutDropPkts", types.YLeaf{"CipSecTunOutDropPkts", cipSecTunnelEntry.CipSecTunOutDropPkts})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunOutAuths", types.YLeaf{"CipSecTunOutAuths", cipSecTunnelEntry.CipSecTunOutAuths})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunOutAuthFails", types.YLeaf{"CipSecTunOutAuthFails", cipSecTunnelEntry.CipSecTunOutAuthFails})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunOutEncrypts", types.YLeaf{"CipSecTunOutEncrypts", cipSecTunnelEntry.CipSecTunOutEncrypts})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunOutEncryptFails", types.YLeaf{"CipSecTunOutEncryptFails", cipSecTunnelEntry.CipSecTunOutEncryptFails})
    cipSecTunnelEntry.EntityData.Leafs.Append("cipSecTunStatus", types.YLeaf{"CipSecTunStatus", cipSecTunnelEntry.CipSecTunStatus})

    cipSecTunnelEntry.EntityData.YListKeys = []string {"CipSecTunIndex"}

    return &(cipSecTunnelEntry.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CipSecEndPtTable
// The IPsec Phase-2 Tunnel Endpoint Table.
// This table contains an entry for each 
// active endpoint associated with an IPsec
//  Phase-2 Tunnel.
type CISCOIPSECFLOWMONITORMIB_CipSecEndPtTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An IPsec Phase-2 Tunnel Endpoint entry. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_CipSecEndPtTable_CipSecEndPtEntry.
    CipSecEndPtEntry []*CISCOIPSECFLOWMONITORMIB_CipSecEndPtTable_CipSecEndPtEntry
}

func (cipSecEndPtTable *CISCOIPSECFLOWMONITORMIB_CipSecEndPtTable) GetEntityData() *types.CommonEntityData {
    cipSecEndPtTable.EntityData.YFilter = cipSecEndPtTable.YFilter
    cipSecEndPtTable.EntityData.YangName = "cipSecEndPtTable"
    cipSecEndPtTable.EntityData.BundleName = "cisco_ios_xe"
    cipSecEndPtTable.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cipSecEndPtTable.EntityData.SegmentPath = "cipSecEndPtTable"
    cipSecEndPtTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipSecEndPtTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipSecEndPtTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipSecEndPtTable.EntityData.Children = types.NewOrderedMap()
    cipSecEndPtTable.EntityData.Children.Append("cipSecEndPtEntry", types.YChild{"CipSecEndPtEntry", nil})
    for i := range cipSecEndPtTable.CipSecEndPtEntry {
        cipSecEndPtTable.EntityData.Children.Append(types.GetSegmentPath(cipSecEndPtTable.CipSecEndPtEntry[i]), types.YChild{"CipSecEndPtEntry", cipSecEndPtTable.CipSecEndPtEntry[i]})
    }
    cipSecEndPtTable.EntityData.Leafs = types.NewOrderedMap()

    cipSecEndPtTable.EntityData.YListKeys = []string {}

    return &(cipSecEndPtTable.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CipSecEndPtTable_CipSecEndPtEntry
// An IPsec Phase-2 Tunnel Endpoint entry.
type CISCOIPSECFLOWMONITORMIB_CipSecEndPtTable_CipSecEndPtEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_ipsec_flow_monitor_mib.CISCOIPSECFLOWMONITORMIB_CipSecTunnelTable_CipSecTunnelEntry_CipSecTunIndex
    CipSecTunIndex interface{}

    // This attribute is a key. The number of the Endpoint associated with the
    // IPsec Phase-2 Tunnel Table.  The value of this index is a number which
    // begins at one and  is incremented with each Endpoint associated  with an
    // IPsec Phase-2 Tunnel. The value of this object will wrap at 2,147,483,647.
    // The type is interface{} with range: 1..2147483647.
    CipSecEndPtIndex interface{}

    // The DNS name of the local Endpoint. The type is string.
    CipSecEndPtLocalName interface{}

    // The type of identity for the local Endpoint. Possible values are: 1) a
    // single IP address, or 2) an IP address range, or 3) an IP subnet. The type
    // is EndPtType.
    CipSecEndPtLocalType interface{}

    // The local Endpoint's first IP address specification.  If the local Endpoint
    // type is single IP address,  then this is the value of the IP address.  If
    // the local Endpoint type is IP subnet, then this is the value of the subnet.
    // If the local Endpoint type is IP address range,  then this is the value of
    // beginning IP address  of the range. The type is string with length: 4 | 16.
    CipSecEndPtLocalAddr1 interface{}

    // The local Endpoint's second IP address specification.  If the local
    // Endpoint type is single IP address,  then this is the value of the IP
    // address.  If the local Endpoint type is IP subnet, then this is the value
    // of the subnet mask.  If the local Endpoint type is IP address range,  then
    // this is the value of ending IP address  of the range. The type is string
    // with length: 4 | 16.
    CipSecEndPtLocalAddr2 interface{}

    // The protocol number of the local Endpoint's traffic. The type is
    // interface{} with range: 0..255.
    CipSecEndPtLocalProtocol interface{}

    // The port number of the local Endpoint's traffic. The type is interface{}
    // with range: 0..65535.
    CipSecEndPtLocalPort interface{}

    // The DNS name of the remote Endpoint. The type is string.
    CipSecEndPtRemoteName interface{}

    // The type of identity for the remote Endpoint. Possible values are: 1) a
    // single IP address, or 2) an IP address range, or 3) an IP subnet. The type
    // is EndPtType.
    CipSecEndPtRemoteType interface{}

    // The remote Endpoint's first IP address specification.  If the remote
    // Endpoint type is single IP address,  then this is the value of the IP
    // address.  If the remote Endpoint type is IP subnet, then this is the value
    // of the subnet.  If the remote Endpoint type is IP address range,  then this
    // is the value of beginning IP address  of the range. The type is string with
    // length: 4 | 16.
    CipSecEndPtRemoteAddr1 interface{}

    // The remote Endpoint's second IP address specification.  If the remote
    // Endpoint type is single IP address,  then this is the value of the IP
    // address.  If the remote Endpoint type is IP subnet, then this is the value
    // of the subnet mask.  If the remote Endpoint type is IP address range,  then
    // this is the value of ending IP address of  the range. The type is string
    // with length: 4 | 16.
    CipSecEndPtRemoteAddr2 interface{}

    // The protocol number of the remote Endpoint's traffic. The type is
    // interface{} with range: 0..255.
    CipSecEndPtRemoteProtocol interface{}

    // The port number of the remote Endpoint's traffic. The type is interface{}
    // with range: 0..65535.
    CipSecEndPtRemotePort interface{}
}

func (cipSecEndPtEntry *CISCOIPSECFLOWMONITORMIB_CipSecEndPtTable_CipSecEndPtEntry) GetEntityData() *types.CommonEntityData {
    cipSecEndPtEntry.EntityData.YFilter = cipSecEndPtEntry.YFilter
    cipSecEndPtEntry.EntityData.YangName = "cipSecEndPtEntry"
    cipSecEndPtEntry.EntityData.BundleName = "cisco_ios_xe"
    cipSecEndPtEntry.EntityData.ParentYangName = "cipSecEndPtTable"
    cipSecEndPtEntry.EntityData.SegmentPath = "cipSecEndPtEntry" + types.AddKeyToken(cipSecEndPtEntry.CipSecTunIndex, "cipSecTunIndex") + types.AddKeyToken(cipSecEndPtEntry.CipSecEndPtIndex, "cipSecEndPtIndex")
    cipSecEndPtEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipSecEndPtEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipSecEndPtEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipSecEndPtEntry.EntityData.Children = types.NewOrderedMap()
    cipSecEndPtEntry.EntityData.Leafs = types.NewOrderedMap()
    cipSecEndPtEntry.EntityData.Leafs.Append("cipSecTunIndex", types.YLeaf{"CipSecTunIndex", cipSecEndPtEntry.CipSecTunIndex})
    cipSecEndPtEntry.EntityData.Leafs.Append("cipSecEndPtIndex", types.YLeaf{"CipSecEndPtIndex", cipSecEndPtEntry.CipSecEndPtIndex})
    cipSecEndPtEntry.EntityData.Leafs.Append("cipSecEndPtLocalName", types.YLeaf{"CipSecEndPtLocalName", cipSecEndPtEntry.CipSecEndPtLocalName})
    cipSecEndPtEntry.EntityData.Leafs.Append("cipSecEndPtLocalType", types.YLeaf{"CipSecEndPtLocalType", cipSecEndPtEntry.CipSecEndPtLocalType})
    cipSecEndPtEntry.EntityData.Leafs.Append("cipSecEndPtLocalAddr1", types.YLeaf{"CipSecEndPtLocalAddr1", cipSecEndPtEntry.CipSecEndPtLocalAddr1})
    cipSecEndPtEntry.EntityData.Leafs.Append("cipSecEndPtLocalAddr2", types.YLeaf{"CipSecEndPtLocalAddr2", cipSecEndPtEntry.CipSecEndPtLocalAddr2})
    cipSecEndPtEntry.EntityData.Leafs.Append("cipSecEndPtLocalProtocol", types.YLeaf{"CipSecEndPtLocalProtocol", cipSecEndPtEntry.CipSecEndPtLocalProtocol})
    cipSecEndPtEntry.EntityData.Leafs.Append("cipSecEndPtLocalPort", types.YLeaf{"CipSecEndPtLocalPort", cipSecEndPtEntry.CipSecEndPtLocalPort})
    cipSecEndPtEntry.EntityData.Leafs.Append("cipSecEndPtRemoteName", types.YLeaf{"CipSecEndPtRemoteName", cipSecEndPtEntry.CipSecEndPtRemoteName})
    cipSecEndPtEntry.EntityData.Leafs.Append("cipSecEndPtRemoteType", types.YLeaf{"CipSecEndPtRemoteType", cipSecEndPtEntry.CipSecEndPtRemoteType})
    cipSecEndPtEntry.EntityData.Leafs.Append("cipSecEndPtRemoteAddr1", types.YLeaf{"CipSecEndPtRemoteAddr1", cipSecEndPtEntry.CipSecEndPtRemoteAddr1})
    cipSecEndPtEntry.EntityData.Leafs.Append("cipSecEndPtRemoteAddr2", types.YLeaf{"CipSecEndPtRemoteAddr2", cipSecEndPtEntry.CipSecEndPtRemoteAddr2})
    cipSecEndPtEntry.EntityData.Leafs.Append("cipSecEndPtRemoteProtocol", types.YLeaf{"CipSecEndPtRemoteProtocol", cipSecEndPtEntry.CipSecEndPtRemoteProtocol})
    cipSecEndPtEntry.EntityData.Leafs.Append("cipSecEndPtRemotePort", types.YLeaf{"CipSecEndPtRemotePort", cipSecEndPtEntry.CipSecEndPtRemotePort})

    cipSecEndPtEntry.EntityData.YListKeys = []string {"CipSecTunIndex", "CipSecEndPtIndex"}

    return &(cipSecEndPtEntry.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CipSecSpiTable
// The IPsec Phase-2 Security Protection Index Table.
// This table contains an entry for each active 
// and expiring security
//  association.
type CISCOIPSECFLOWMONITORMIB_CipSecSpiTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with active and expiring
    // IPsec Phase-2  security associations. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_CipSecSpiTable_CipSecSpiEntry.
    CipSecSpiEntry []*CISCOIPSECFLOWMONITORMIB_CipSecSpiTable_CipSecSpiEntry
}

func (cipSecSpiTable *CISCOIPSECFLOWMONITORMIB_CipSecSpiTable) GetEntityData() *types.CommonEntityData {
    cipSecSpiTable.EntityData.YFilter = cipSecSpiTable.YFilter
    cipSecSpiTable.EntityData.YangName = "cipSecSpiTable"
    cipSecSpiTable.EntityData.BundleName = "cisco_ios_xe"
    cipSecSpiTable.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cipSecSpiTable.EntityData.SegmentPath = "cipSecSpiTable"
    cipSecSpiTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipSecSpiTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipSecSpiTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipSecSpiTable.EntityData.Children = types.NewOrderedMap()
    cipSecSpiTable.EntityData.Children.Append("cipSecSpiEntry", types.YChild{"CipSecSpiEntry", nil})
    for i := range cipSecSpiTable.CipSecSpiEntry {
        cipSecSpiTable.EntityData.Children.Append(types.GetSegmentPath(cipSecSpiTable.CipSecSpiEntry[i]), types.YChild{"CipSecSpiEntry", cipSecSpiTable.CipSecSpiEntry[i]})
    }
    cipSecSpiTable.EntityData.Leafs = types.NewOrderedMap()

    cipSecSpiTable.EntityData.YListKeys = []string {}

    return &(cipSecSpiTable.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CipSecSpiTable_CipSecSpiEntry
// Each entry contains the attributes associated with
// active and expiring IPsec Phase-2 
// security associations.
type CISCOIPSECFLOWMONITORMIB_CipSecSpiTable_CipSecSpiEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_ipsec_flow_monitor_mib.CISCOIPSECFLOWMONITORMIB_CipSecTunnelTable_CipSecTunnelEntry_CipSecTunIndex
    CipSecTunIndex interface{}

    // This attribute is a key. The number of the SPI associated with the Phase-2
    // Tunnel Table.  The value of this  index is a number which begins at one and
    // is  incremented with each SPI associated with an  IPsec Phase-2 Tunnel. 
    // The value of this  object will wrap at 2,147,483,647. The type is
    // interface{} with range: 1..2147483647.
    CipSecSpiIndex interface{}

    // The direction of the SPI. The type is CipSecSpiDirection.
    CipSecSpiDirection interface{}

    // The value of the SPI. The type is interface{} with range: 1..4294967295.
    CipSecSpiValue interface{}

    // The protocol of the SPI. The type is CipSecSpiProtocol.
    CipSecSpiProtocol interface{}

    // The status of the SPI. The type is CipSecSpiStatus.
    CipSecSpiStatus interface{}
}

func (cipSecSpiEntry *CISCOIPSECFLOWMONITORMIB_CipSecSpiTable_CipSecSpiEntry) GetEntityData() *types.CommonEntityData {
    cipSecSpiEntry.EntityData.YFilter = cipSecSpiEntry.YFilter
    cipSecSpiEntry.EntityData.YangName = "cipSecSpiEntry"
    cipSecSpiEntry.EntityData.BundleName = "cisco_ios_xe"
    cipSecSpiEntry.EntityData.ParentYangName = "cipSecSpiTable"
    cipSecSpiEntry.EntityData.SegmentPath = "cipSecSpiEntry" + types.AddKeyToken(cipSecSpiEntry.CipSecTunIndex, "cipSecTunIndex") + types.AddKeyToken(cipSecSpiEntry.CipSecSpiIndex, "cipSecSpiIndex")
    cipSecSpiEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipSecSpiEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipSecSpiEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipSecSpiEntry.EntityData.Children = types.NewOrderedMap()
    cipSecSpiEntry.EntityData.Leafs = types.NewOrderedMap()
    cipSecSpiEntry.EntityData.Leafs.Append("cipSecTunIndex", types.YLeaf{"CipSecTunIndex", cipSecSpiEntry.CipSecTunIndex})
    cipSecSpiEntry.EntityData.Leafs.Append("cipSecSpiIndex", types.YLeaf{"CipSecSpiIndex", cipSecSpiEntry.CipSecSpiIndex})
    cipSecSpiEntry.EntityData.Leafs.Append("cipSecSpiDirection", types.YLeaf{"CipSecSpiDirection", cipSecSpiEntry.CipSecSpiDirection})
    cipSecSpiEntry.EntityData.Leafs.Append("cipSecSpiValue", types.YLeaf{"CipSecSpiValue", cipSecSpiEntry.CipSecSpiValue})
    cipSecSpiEntry.EntityData.Leafs.Append("cipSecSpiProtocol", types.YLeaf{"CipSecSpiProtocol", cipSecSpiEntry.CipSecSpiProtocol})
    cipSecSpiEntry.EntityData.Leafs.Append("cipSecSpiStatus", types.YLeaf{"CipSecSpiStatus", cipSecSpiEntry.CipSecSpiStatus})

    cipSecSpiEntry.EntityData.YListKeys = []string {"CipSecTunIndex", "CipSecSpiIndex"}

    return &(cipSecSpiEntry.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CipSecSpiTable_CipSecSpiEntry_CipSecSpiDirection represents The direction of the SPI.
type CISCOIPSECFLOWMONITORMIB_CipSecSpiTable_CipSecSpiEntry_CipSecSpiDirection string

const (
    CISCOIPSECFLOWMONITORMIB_CipSecSpiTable_CipSecSpiEntry_CipSecSpiDirection_in CISCOIPSECFLOWMONITORMIB_CipSecSpiTable_CipSecSpiEntry_CipSecSpiDirection = "in"

    CISCOIPSECFLOWMONITORMIB_CipSecSpiTable_CipSecSpiEntry_CipSecSpiDirection_out CISCOIPSECFLOWMONITORMIB_CipSecSpiTable_CipSecSpiEntry_CipSecSpiDirection = "out"
)

// CISCOIPSECFLOWMONITORMIB_CipSecSpiTable_CipSecSpiEntry_CipSecSpiProtocol represents The protocol of the SPI.
type CISCOIPSECFLOWMONITORMIB_CipSecSpiTable_CipSecSpiEntry_CipSecSpiProtocol string

const (
    CISCOIPSECFLOWMONITORMIB_CipSecSpiTable_CipSecSpiEntry_CipSecSpiProtocol_ah CISCOIPSECFLOWMONITORMIB_CipSecSpiTable_CipSecSpiEntry_CipSecSpiProtocol = "ah"

    CISCOIPSECFLOWMONITORMIB_CipSecSpiTable_CipSecSpiEntry_CipSecSpiProtocol_esp CISCOIPSECFLOWMONITORMIB_CipSecSpiTable_CipSecSpiEntry_CipSecSpiProtocol = "esp"

    CISCOIPSECFLOWMONITORMIB_CipSecSpiTable_CipSecSpiEntry_CipSecSpiProtocol_ipcomp CISCOIPSECFLOWMONITORMIB_CipSecSpiTable_CipSecSpiEntry_CipSecSpiProtocol = "ipcomp"
)

// CISCOIPSECFLOWMONITORMIB_CipSecSpiTable_CipSecSpiEntry_CipSecSpiStatus represents The status of the SPI.
type CISCOIPSECFLOWMONITORMIB_CipSecSpiTable_CipSecSpiEntry_CipSecSpiStatus string

const (
    CISCOIPSECFLOWMONITORMIB_CipSecSpiTable_CipSecSpiEntry_CipSecSpiStatus_active CISCOIPSECFLOWMONITORMIB_CipSecSpiTable_CipSecSpiEntry_CipSecSpiStatus = "active"

    CISCOIPSECFLOWMONITORMIB_CipSecSpiTable_CipSecSpiEntry_CipSecSpiStatus_expiring CISCOIPSECFLOWMONITORMIB_CipSecSpiTable_CipSecSpiEntry_CipSecSpiStatus = "expiring"
)

// CISCOIPSECFLOWMONITORMIB_CipSecPhase2GWStatsTable
// Phase-2 IPsec stats information is included in this table.
// Each entry is related to a specific gateway which is 
// identified by 'cmgwIndex'
type CISCOIPSECFLOWMONITORMIB_CipSecPhase2GWStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes of an Phase-2 IPsec stats information
    // for the related gateway.  There is only one entry for each gateway. The
    // entry  is created when a gateway up and cannot be deleted. The type is
    // slice of
    // CISCOIPSECFLOWMONITORMIB_CipSecPhase2GWStatsTable_CipSecPhase2GWStatsEntry.
    CipSecPhase2GWStatsEntry []*CISCOIPSECFLOWMONITORMIB_CipSecPhase2GWStatsTable_CipSecPhase2GWStatsEntry
}

func (cipSecPhase2GWStatsTable *CISCOIPSECFLOWMONITORMIB_CipSecPhase2GWStatsTable) GetEntityData() *types.CommonEntityData {
    cipSecPhase2GWStatsTable.EntityData.YFilter = cipSecPhase2GWStatsTable.YFilter
    cipSecPhase2GWStatsTable.EntityData.YangName = "cipSecPhase2GWStatsTable"
    cipSecPhase2GWStatsTable.EntityData.BundleName = "cisco_ios_xe"
    cipSecPhase2GWStatsTable.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cipSecPhase2GWStatsTable.EntityData.SegmentPath = "cipSecPhase2GWStatsTable"
    cipSecPhase2GWStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipSecPhase2GWStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipSecPhase2GWStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipSecPhase2GWStatsTable.EntityData.Children = types.NewOrderedMap()
    cipSecPhase2GWStatsTable.EntityData.Children.Append("cipSecPhase2GWStatsEntry", types.YChild{"CipSecPhase2GWStatsEntry", nil})
    for i := range cipSecPhase2GWStatsTable.CipSecPhase2GWStatsEntry {
        cipSecPhase2GWStatsTable.EntityData.Children.Append(types.GetSegmentPath(cipSecPhase2GWStatsTable.CipSecPhase2GWStatsEntry[i]), types.YChild{"CipSecPhase2GWStatsEntry", cipSecPhase2GWStatsTable.CipSecPhase2GWStatsEntry[i]})
    }
    cipSecPhase2GWStatsTable.EntityData.Leafs = types.NewOrderedMap()

    cipSecPhase2GWStatsTable.EntityData.YListKeys = []string {}

    return &(cipSecPhase2GWStatsTable.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CipSecPhase2GWStatsTable_CipSecPhase2GWStatsEntry
// Each entry contains the attributes of an Phase-2 IPsec stats
// information for the related gateway.
// 
// There is only one entry for each gateway. The entry 
// is created when a gateway up and cannot be deleted.
type CISCOIPSECFLOWMONITORMIB_CipSecPhase2GWStatsTable_CipSecPhase2GWStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_media_gateway_mib.CISCOMEDIAGATEWAYMIB_CMediaGwTable_CMediaGwEntry_CmgwIndex
    CmgwIndex interface{}

    // The total number of currently active IPsec Phase-2 Tunnels. The type is
    // interface{} with range: 0..4294967295.
    CipSecPhase2GWActiveTunnels interface{}

    // The total number of previously active IPsec Phase-2 Tunnels. The type is
    // interface{} with range: 0..4294967295. Units are Phase-2 Tunnels.
    CipSecPhase2GWPreviousTunnels interface{}

    // The total number of octets received by all current and previous IPsec
    // Phase-2 Tunnels.  This value is accumulated BEFORE determining  whether or
    // not the packet should be decompressed.  See also cipSecGlobalInOctWraps for
    // the number of times this counter has wrapped. The type is interface{} with
    // range: 0..4294967295. Units are Octets.
    CipSecPhase2GWInOctets interface{}

    // The number of times the global octets received counter
    // (cipSecGlobalInOctets) has wrapped. The type is interface{} with range:
    // 0..4294967295. Units are Integral units.
    CipSecPhase2GWInOctWraps interface{}

    // The total number of decompressed octets received by all current and
    // previous IPsec Phase-2 Tunnels.   This value is accumulated AFTER the
    // packet is  decompressed. If compression is not being used,  this value will
    // match the value of cipSecGlobalInOctets.  See also
    // cipSecGlobalInDecompOctWraps for the number of times this counter has
    // wrapped. The type is interface{} with range: 0..4294967295. Units are
    // Octets.
    CipSecPhase2GWInDecompOctets interface{}

    // The number of times the global decompressed octets received counter
    // (cipSecGlobalInDecompOctets)  has wrapped. The type is interface{} with
    // range: 0..4294967295. Units are Integral units.
    CipSecPhase2GWInDecompOctWraps interface{}

    // The total number of packets received by all current and previous IPsec
    // Phase-2 Tunnels. The type is interface{} with range: 0..4294967295. Units
    // are Packets.
    CipSecPhase2GWInPkts interface{}

    // The total number of packets dropped during receive processing by all
    // current and previous  IPsec Phase-2 Tunnels. This count does NOT include 
    // packets dropped due to Anti-Replay processing. The type is interface{} with
    // range: 0..4294967295. Units are Packets.
    CipSecPhase2GWInDrops interface{}

    // The total number of packets dropped during receive processing due to
    // Anti-Replay  processing by all current and previous IPsec Phase-2 Tunnels.
    // The type is interface{} with range: 0..4294967295. Units are Packets.
    CipSecPhase2GWInReplayDrops interface{}

    // The total number of inbound authentication's performed by all current and
    // previous IPsec  Phase-2 Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Events.
    CipSecPhase2GWInAuths interface{}

    // The total number of inbound authentication's which ended in failure by all
    // current and previous  IPsec Phase-2 Tunnels. The type is interface{} with
    // range: 0..4294967295. Units are Failures.
    CipSecPhase2GWInAuthFails interface{}

    // The total number of inbound decryption's performed by all current and
    // previous IPsec  Phase-2 Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    CipSecPhase2GWInDecrypts interface{}

    // The total number of inbound decryption's which ended in failure by all
    // current and  previous IPsec Phase-2 Tunnels. The type is interface{} with
    // range: 0..4294967295. Units are Packets.
    CipSecPhase2GWInDecryptFails interface{}

    // The total number of octets sent by all current and previous IPsec Phase-2
    // Tunnels.   This value is accumulated AFTER determining  whether or not the
    // packet should be compressed.   See also cipSecGlobalOutOctWraps for the
    // number of times this counter has wrapped. The type is interface{} with
    // range: 0..4294967295. Units are Octets.
    CipSecPhase2GWOutOctets interface{}

    // The number of times the global octets sent counter (cipSecGlobalOutOctets)
    // has wrapped. The type is interface{} with range: 0..4294967295. Units are
    // Integral units.
    CipSecPhase2GWOutOctWraps interface{}

    // The total number of uncompressed octets sent by all current and previous
    // IPsec Phase-2 Tunnels.   This value is accumulated BEFORE the packet is 
    // compressed. If compression is not being used, this  value will match the
    // value of cipSecGlobalOutOctets.  See also cipSecGlobalOutDecompOctWraps for
    // the number  of times this counter has wrapped. The type is interface{} with
    // range: 0..4294967295. Units are Octets.
    CipSecPhase2GWOutUncompOctets interface{}

    // The number of times the global uncompressed octets sent counter
    // (cipSecGlobalOutUncompOctets)  has wrapped. The type is interface{} with
    // range: 0..4294967295. Units are Integral units.
    CipSecPhase2GWOutUncompOctWraps interface{}

    // The total number of packets sent by all current and previous IPsec Phase-2 
    // Tunnels. The type is interface{} with range: 0..4294967295. Units are
    // Packets.
    CipSecPhase2GWOutPkts interface{}

    // The total number of packets dropped during send processing by all current
    // and previous IPsec  Phase-2 Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    CipSecPhase2GWOutDrops interface{}

    // The total number of outbound authentication's performed by all current and
    // previous IPsec  Phase-2 Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Events.
    CipSecPhase2GWOutAuths interface{}

    // The total number of outbound authentication's which ended in failure by all
    // current and previous IPsec Phase-2 Tunnels. The type is interface{} with
    // range: 0..4294967295. Units are Failures.
    CipSecPhase2GWOutAuthFails interface{}

    // The total number of outbound encryption's performed by all current and
    // previous IPsec Phase-2 Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    CipSecPhase2GWOutEncrypts interface{}

    // The total number of outbound encryption's which ended in failure by all
    // current and  previous IPsec Phase-2 Tunnels. The type is interface{} with
    // range: 0..4294967295. Units are Failures.
    CipSecPhase2GWOutEncryptFails interface{}

    // The total number of protocol use failures which occurred during processing
    // of all current  and previously active IPsec Phase-2 Tunnels. The type is
    // interface{} with range: 0..4294967295. Units are Failures.
    CipSecPhase2GWProtocolUseFails interface{}

    // The total number of non-existent Security Association in failures which
    // occurred  during processing of all current and previous IPsec Phase-2
    // Tunnels. The type is interface{} with range: 0..4294967295. Units are
    // Failures.
    CipSecPhase2GWNoSaFails interface{}

    // The total number of system capacity failures which occurred during
    // processing of all current  and previously active IPsec Phase-2 Tunnels. The
    // type is interface{} with range: 0..4294967295. Units are Failures.
    CipSecPhase2GWSysCapFails interface{}
}

func (cipSecPhase2GWStatsEntry *CISCOIPSECFLOWMONITORMIB_CipSecPhase2GWStatsTable_CipSecPhase2GWStatsEntry) GetEntityData() *types.CommonEntityData {
    cipSecPhase2GWStatsEntry.EntityData.YFilter = cipSecPhase2GWStatsEntry.YFilter
    cipSecPhase2GWStatsEntry.EntityData.YangName = "cipSecPhase2GWStatsEntry"
    cipSecPhase2GWStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    cipSecPhase2GWStatsEntry.EntityData.ParentYangName = "cipSecPhase2GWStatsTable"
    cipSecPhase2GWStatsEntry.EntityData.SegmentPath = "cipSecPhase2GWStatsEntry" + types.AddKeyToken(cipSecPhase2GWStatsEntry.CmgwIndex, "cmgwIndex")
    cipSecPhase2GWStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipSecPhase2GWStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipSecPhase2GWStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipSecPhase2GWStatsEntry.EntityData.Children = types.NewOrderedMap()
    cipSecPhase2GWStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    cipSecPhase2GWStatsEntry.EntityData.Leafs.Append("cmgwIndex", types.YLeaf{"CmgwIndex", cipSecPhase2GWStatsEntry.CmgwIndex})
    cipSecPhase2GWStatsEntry.EntityData.Leafs.Append("cipSecPhase2GWActiveTunnels", types.YLeaf{"CipSecPhase2GWActiveTunnels", cipSecPhase2GWStatsEntry.CipSecPhase2GWActiveTunnels})
    cipSecPhase2GWStatsEntry.EntityData.Leafs.Append("cipSecPhase2GWPreviousTunnels", types.YLeaf{"CipSecPhase2GWPreviousTunnels", cipSecPhase2GWStatsEntry.CipSecPhase2GWPreviousTunnels})
    cipSecPhase2GWStatsEntry.EntityData.Leafs.Append("cipSecPhase2GWInOctets", types.YLeaf{"CipSecPhase2GWInOctets", cipSecPhase2GWStatsEntry.CipSecPhase2GWInOctets})
    cipSecPhase2GWStatsEntry.EntityData.Leafs.Append("cipSecPhase2GWInOctWraps", types.YLeaf{"CipSecPhase2GWInOctWraps", cipSecPhase2GWStatsEntry.CipSecPhase2GWInOctWraps})
    cipSecPhase2GWStatsEntry.EntityData.Leafs.Append("cipSecPhase2GWInDecompOctets", types.YLeaf{"CipSecPhase2GWInDecompOctets", cipSecPhase2GWStatsEntry.CipSecPhase2GWInDecompOctets})
    cipSecPhase2GWStatsEntry.EntityData.Leafs.Append("cipSecPhase2GWInDecompOctWraps", types.YLeaf{"CipSecPhase2GWInDecompOctWraps", cipSecPhase2GWStatsEntry.CipSecPhase2GWInDecompOctWraps})
    cipSecPhase2GWStatsEntry.EntityData.Leafs.Append("cipSecPhase2GWInPkts", types.YLeaf{"CipSecPhase2GWInPkts", cipSecPhase2GWStatsEntry.CipSecPhase2GWInPkts})
    cipSecPhase2GWStatsEntry.EntityData.Leafs.Append("cipSecPhase2GWInDrops", types.YLeaf{"CipSecPhase2GWInDrops", cipSecPhase2GWStatsEntry.CipSecPhase2GWInDrops})
    cipSecPhase2GWStatsEntry.EntityData.Leafs.Append("cipSecPhase2GWInReplayDrops", types.YLeaf{"CipSecPhase2GWInReplayDrops", cipSecPhase2GWStatsEntry.CipSecPhase2GWInReplayDrops})
    cipSecPhase2GWStatsEntry.EntityData.Leafs.Append("cipSecPhase2GWInAuths", types.YLeaf{"CipSecPhase2GWInAuths", cipSecPhase2GWStatsEntry.CipSecPhase2GWInAuths})
    cipSecPhase2GWStatsEntry.EntityData.Leafs.Append("cipSecPhase2GWInAuthFails", types.YLeaf{"CipSecPhase2GWInAuthFails", cipSecPhase2GWStatsEntry.CipSecPhase2GWInAuthFails})
    cipSecPhase2GWStatsEntry.EntityData.Leafs.Append("cipSecPhase2GWInDecrypts", types.YLeaf{"CipSecPhase2GWInDecrypts", cipSecPhase2GWStatsEntry.CipSecPhase2GWInDecrypts})
    cipSecPhase2GWStatsEntry.EntityData.Leafs.Append("cipSecPhase2GWInDecryptFails", types.YLeaf{"CipSecPhase2GWInDecryptFails", cipSecPhase2GWStatsEntry.CipSecPhase2GWInDecryptFails})
    cipSecPhase2GWStatsEntry.EntityData.Leafs.Append("cipSecPhase2GWOutOctets", types.YLeaf{"CipSecPhase2GWOutOctets", cipSecPhase2GWStatsEntry.CipSecPhase2GWOutOctets})
    cipSecPhase2GWStatsEntry.EntityData.Leafs.Append("cipSecPhase2GWOutOctWraps", types.YLeaf{"CipSecPhase2GWOutOctWraps", cipSecPhase2GWStatsEntry.CipSecPhase2GWOutOctWraps})
    cipSecPhase2GWStatsEntry.EntityData.Leafs.Append("cipSecPhase2GWOutUncompOctets", types.YLeaf{"CipSecPhase2GWOutUncompOctets", cipSecPhase2GWStatsEntry.CipSecPhase2GWOutUncompOctets})
    cipSecPhase2GWStatsEntry.EntityData.Leafs.Append("cipSecPhase2GWOutUncompOctWraps", types.YLeaf{"CipSecPhase2GWOutUncompOctWraps", cipSecPhase2GWStatsEntry.CipSecPhase2GWOutUncompOctWraps})
    cipSecPhase2GWStatsEntry.EntityData.Leafs.Append("cipSecPhase2GWOutPkts", types.YLeaf{"CipSecPhase2GWOutPkts", cipSecPhase2GWStatsEntry.CipSecPhase2GWOutPkts})
    cipSecPhase2GWStatsEntry.EntityData.Leafs.Append("cipSecPhase2GWOutDrops", types.YLeaf{"CipSecPhase2GWOutDrops", cipSecPhase2GWStatsEntry.CipSecPhase2GWOutDrops})
    cipSecPhase2GWStatsEntry.EntityData.Leafs.Append("cipSecPhase2GWOutAuths", types.YLeaf{"CipSecPhase2GWOutAuths", cipSecPhase2GWStatsEntry.CipSecPhase2GWOutAuths})
    cipSecPhase2GWStatsEntry.EntityData.Leafs.Append("cipSecPhase2GWOutAuthFails", types.YLeaf{"CipSecPhase2GWOutAuthFails", cipSecPhase2GWStatsEntry.CipSecPhase2GWOutAuthFails})
    cipSecPhase2GWStatsEntry.EntityData.Leafs.Append("cipSecPhase2GWOutEncrypts", types.YLeaf{"CipSecPhase2GWOutEncrypts", cipSecPhase2GWStatsEntry.CipSecPhase2GWOutEncrypts})
    cipSecPhase2GWStatsEntry.EntityData.Leafs.Append("cipSecPhase2GWOutEncryptFails", types.YLeaf{"CipSecPhase2GWOutEncryptFails", cipSecPhase2GWStatsEntry.CipSecPhase2GWOutEncryptFails})
    cipSecPhase2GWStatsEntry.EntityData.Leafs.Append("cipSecPhase2GWProtocolUseFails", types.YLeaf{"CipSecPhase2GWProtocolUseFails", cipSecPhase2GWStatsEntry.CipSecPhase2GWProtocolUseFails})
    cipSecPhase2GWStatsEntry.EntityData.Leafs.Append("cipSecPhase2GWNoSaFails", types.YLeaf{"CipSecPhase2GWNoSaFails", cipSecPhase2GWStatsEntry.CipSecPhase2GWNoSaFails})
    cipSecPhase2GWStatsEntry.EntityData.Leafs.Append("cipSecPhase2GWSysCapFails", types.YLeaf{"CipSecPhase2GWSysCapFails", cipSecPhase2GWStatsEntry.CipSecPhase2GWSysCapFails})

    cipSecPhase2GWStatsEntry.EntityData.YListKeys = []string {"CmgwIndex"}

    return &(cipSecPhase2GWStatsEntry.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CikeTunnelHistTable
// The IPsec Phase-1 Internet Key Exchange Tunnel
// History Table.  This table is implemented as a 
// sliding window in which only the last n entries 
// are maintained.  The maximum number of entries
//  is specified by the cipSecHistTableSize object.
type CISCOIPSECFLOWMONITORMIB_CikeTunnelHistTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with a previously active
    // IPsec  Phase-1 IKE Tunnel. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_CikeTunnelHistTable_CikeTunnelHistEntry.
    CikeTunnelHistEntry []*CISCOIPSECFLOWMONITORMIB_CikeTunnelHistTable_CikeTunnelHistEntry
}

func (cikeTunnelHistTable *CISCOIPSECFLOWMONITORMIB_CikeTunnelHistTable) GetEntityData() *types.CommonEntityData {
    cikeTunnelHistTable.EntityData.YFilter = cikeTunnelHistTable.YFilter
    cikeTunnelHistTable.EntityData.YangName = "cikeTunnelHistTable"
    cikeTunnelHistTable.EntityData.BundleName = "cisco_ios_xe"
    cikeTunnelHistTable.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cikeTunnelHistTable.EntityData.SegmentPath = "cikeTunnelHistTable"
    cikeTunnelHistTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cikeTunnelHistTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cikeTunnelHistTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cikeTunnelHistTable.EntityData.Children = types.NewOrderedMap()
    cikeTunnelHistTable.EntityData.Children.Append("cikeTunnelHistEntry", types.YChild{"CikeTunnelHistEntry", nil})
    for i := range cikeTunnelHistTable.CikeTunnelHistEntry {
        cikeTunnelHistTable.EntityData.Children.Append(types.GetSegmentPath(cikeTunnelHistTable.CikeTunnelHistEntry[i]), types.YChild{"CikeTunnelHistEntry", cikeTunnelHistTable.CikeTunnelHistEntry[i]})
    }
    cikeTunnelHistTable.EntityData.Leafs = types.NewOrderedMap()

    cikeTunnelHistTable.EntityData.YListKeys = []string {}

    return &(cikeTunnelHistTable.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CikeTunnelHistTable_CikeTunnelHistEntry
// Each entry contains the attributes
// associated with a previously active IPsec 
// Phase-1 IKE Tunnel.
type CISCOIPSECFLOWMONITORMIB_CikeTunnelHistTable_CikeTunnelHistEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index of the IPsec Phase-1 IKE Tunnel History
    // Table.  The value of the index is a number which  begins at one and is
    // incremented with each  tunnel that ends. The value of this object  will
    // wrap at 2,147,483,647. The type is interface{} with range: 1..2147483647.
    CikeTunHistIndex interface{}

    // The reason the IPsec Phase-1 IKE Tunnel was terminated. Possible reasons
    // include: 1 = other 2 = normal termination 3 = operator request 4 = peer
    // delete request was received 5 = contact with peer was lost 6 = local
    // failure occurred. 7 = operator initiated check point request. The type is
    // CikeTunHistTermReason.
    CikeTunHistTermReason interface{}

    // The index of the previously active IPsec Phase-1 IKE Tunnel. The type is
    // interface{} with range: 1..2147483647.
    CikeTunHistActiveIndex interface{}

    // The type of local peer identity.  The local peer may be identified by:  1.
    // an IP address, or  2. a host name. The type is IkePeerType.
    CikeTunHistPeerLocalType interface{}

    // The value of the local peer identity.  If the local peer type is an IP
    // Address, then this is the IP Address used to identify the local peer.  If
    // the local peer type is a host name, then this is the host name used to
    // identify the local peer. The type is string.
    CikeTunHistPeerLocalValue interface{}

    // The internal index of the local-remote peer association.  This internal
    // index is used to  uniquely identify multiple associations between  the
    // local and remote peer. The type is interface{} with range: 1..2147483647.
    CikeTunHistPeerIntIndex interface{}

    // The type of remote peer identity.  The remote peer may be identified by: 
    // 1. an IP address, or  2. a host name. The type is IkePeerType.
    CikeTunHistPeerRemoteType interface{}

    // The value of the remote peer identity.  If the remote peer type is an IP
    // Address, then this is the IP Address used to identify the remote peer.  If
    // the remote peer type is a host name, then this is the host name used to
    // identify the remote peer. The type is string.
    CikeTunHistPeerRemoteValue interface{}

    // The IP address of the local endpoint for the IPsec Phase-1 IKE Tunnel. The
    // type is string with length: 4 | 16.
    CikeTunHistLocalAddr interface{}

    // The DNS name of the local IP address for the IPsec Phase-1 IKE Tunnel. If
    // the DNS  name associated with the local tunnel endpoint  is not known, then
    // the value of this  object will be a NULL string. The type is string.
    CikeTunHistLocalName interface{}

    // The IP address of the remote endpoint for the IPsec Phase-1 IKE Tunnel. The
    // type is string with length: 4 | 16.
    CikeTunHistRemoteAddr interface{}

    // The DNS name of the remote IP address of IPsec Phase-1 IKE Tunnel. If the
    // DNS name associated with the remote tunnel endpoint is not known, then the
    // value of this object will be a NULL string. The type is string.
    CikeTunHistRemoteName interface{}

    // The negotiation mode of the IPsec Phase-1 IKE Tunnel. The type is
    // IkeNegoMode.
    CikeTunHistNegoMode interface{}

    // The Diffie Hellman Group used in IPsec Phase-1 IKE negotiations. The type
    // is DiffHellmanGrp.
    CikeTunHistDiffHellmanGrp interface{}

    // The encryption algorithm used in IPsec Phase-1 IKE negotiations. The type
    // is EncryptAlgo.
    CikeTunHistEncryptAlgo interface{}

    // The hash algorithm used in IPsec Phase-1 IKE negotiations. The type is
    // IkeHashAlgo.
    CikeTunHistHashAlgo interface{}

    // The authentication method used in IPsec Phase-1 IKE negotiations. The type
    // is IkeAuthMethod.
    CikeTunHistAuthMethod interface{}

    // The negotiated LifeTime of the IPsec Phase-1 IKE Tunnel in seconds. The
    // type is interface{} with range: 1..2147483647.
    CikeTunHistLifeTime interface{}

    // The value of sysUpTime in hundredths of seconds when the IPsec Phase-1 IKE
    // tunnel was started. The type is interface{} with range: 0..4294967295.
    CikeTunHistStartTime interface{}

    // The length of time the IPsec Phase-1 IKE tunnel was been active in
    // hundredths of seconds. The type is interface{} with range: 0..2147483647.
    CikeTunHistActiveTime interface{}

    // The total number of security associations refreshes performed. The type is
    // interface{} with range: 0..4294967295. Units are QM Exchanges.
    CikeTunHistTotalRefreshes interface{}

    // The total number of security associations used during the  life of the
    // IPsec Phase-1 IKE Tunnel. The type is interface{} with range:
    // 0..4294967295. Units are SAs.
    CikeTunHistTotalSas interface{}

    // The total number of octets received by this IPsec Phase-1  IKE Tunnel. The
    // type is interface{} with range: 0..4294967295. Units are Octets.
    CikeTunHistInOctets interface{}

    // The total number of packets received by this IPsec Phase-1  IKE Tunnel. The
    // type is interface{} with range: 0..4294967295. Units are Packets.
    CikeTunHistInPkts interface{}

    // The total number of packets dropped by this IPsec Phase-1  IKE Tunnel
    // during receive processing. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    CikeTunHistInDropPkts interface{}

    // The total number of notifys received by this IPsec Phase-1  IKE Tunnel. The
    // type is interface{} with range: 0..4294967295. Units are Notification
    // Payloads.
    CikeTunHistInNotifys interface{}

    // The total number of IPsec Phase-2 exchanges received by  this IPsec Phase-1
    // IKE Tunnel. The type is interface{} with range: 0..4294967295. Units are SA
    // Payloads.
    CikeTunHistInP2Exchgs interface{}

    // The total number of IPsec Phase-2 exchanges received and  found to be
    // invalid by this IPsec Phase-1 IKE Tunnel. The type is interface{} with
    // range: 0..4294967295. Units are SA Payloads.
    CikeTunHistInP2ExchgInvalids interface{}

    // The total number of IPsec Phase-2 exchanges received and  rejected by this
    // IPsec Phase-1 IKE Tunnel. The type is interface{} with range:
    // 0..4294967295. Units are SA Payloads.
    CikeTunHistInP2ExchgRejects interface{}

    // The total number of IPsec Phase-2 security association delete requests
    // received by this IPsec  Phase-1 IKE Tunnel. The type is interface{} with
    // range: 0..4294967295. Units are Notification Payloads.
    CikeTunHistInP2SaDelRequests interface{}

    // The total number of octets sent by this IPsec Phase-1 IKE Tunnel. The type
    // is interface{} with range: 0..4294967295. Units are Octets.
    CikeTunHistOutOctets interface{}

    // The total number of packets sent by this IPsec Phase-1 IKE Tunnel. The type
    // is interface{} with range: 0..4294967295. Units are Packets.
    CikeTunHistOutPkts interface{}

    // The total number of packets dropped by this IPsec Phase-1  IKE Tunnel
    // during send processing. The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    CikeTunHistOutDropPkts interface{}

    // The total number of notifys sent by this IPsec Phase-1 IKE Tunnel. The type
    // is interface{} with range: 0..4294967295. Units are Notification Payloads.
    CikeTunHistOutNotifys interface{}

    // The total number of IPsec Phase-2 exchanges sent by this IPsec Phase-1 IKE
    // Tunnel. The type is interface{} with range: 0..4294967295. Units are SA
    // Payloads.
    CikeTunHistOutP2Exchgs interface{}

    // The total number of IPsec Phase-2 exchanges sent and found to be invalid by
    // this IPsec Phase-1 IKE Tunnel. The type is interface{} with range:
    // 0..4294967295. Units are SA Payloads.
    CikeTunHistOutP2ExchgInvalids interface{}

    // The total number of IPsec Phase-2 exchanges sent and rejected by this IPsec
    // Phase-1 IKE Tunnel. The type is interface{} with range: 0..4294967295.
    // Units are SA Payloads.
    CikeTunHistOutP2ExchgRejects interface{}

    // The total number of IPsec Phase-2 security association delete requests sent
    // by this IPsec Phase-1 IKE Tunnel. The type is interface{} with range:
    // 0..4294967295. Units are Notification Payloads.
    CikeTunHistOutP2SaDelRequests interface{}
}

func (cikeTunnelHistEntry *CISCOIPSECFLOWMONITORMIB_CikeTunnelHistTable_CikeTunnelHistEntry) GetEntityData() *types.CommonEntityData {
    cikeTunnelHistEntry.EntityData.YFilter = cikeTunnelHistEntry.YFilter
    cikeTunnelHistEntry.EntityData.YangName = "cikeTunnelHistEntry"
    cikeTunnelHistEntry.EntityData.BundleName = "cisco_ios_xe"
    cikeTunnelHistEntry.EntityData.ParentYangName = "cikeTunnelHistTable"
    cikeTunnelHistEntry.EntityData.SegmentPath = "cikeTunnelHistEntry" + types.AddKeyToken(cikeTunnelHistEntry.CikeTunHistIndex, "cikeTunHistIndex")
    cikeTunnelHistEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cikeTunnelHistEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cikeTunnelHistEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cikeTunnelHistEntry.EntityData.Children = types.NewOrderedMap()
    cikeTunnelHistEntry.EntityData.Leafs = types.NewOrderedMap()
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistIndex", types.YLeaf{"CikeTunHistIndex", cikeTunnelHistEntry.CikeTunHistIndex})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistTermReason", types.YLeaf{"CikeTunHistTermReason", cikeTunnelHistEntry.CikeTunHistTermReason})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistActiveIndex", types.YLeaf{"CikeTunHistActiveIndex", cikeTunnelHistEntry.CikeTunHistActiveIndex})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistPeerLocalType", types.YLeaf{"CikeTunHistPeerLocalType", cikeTunnelHistEntry.CikeTunHistPeerLocalType})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistPeerLocalValue", types.YLeaf{"CikeTunHistPeerLocalValue", cikeTunnelHistEntry.CikeTunHistPeerLocalValue})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistPeerIntIndex", types.YLeaf{"CikeTunHistPeerIntIndex", cikeTunnelHistEntry.CikeTunHistPeerIntIndex})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistPeerRemoteType", types.YLeaf{"CikeTunHistPeerRemoteType", cikeTunnelHistEntry.CikeTunHistPeerRemoteType})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistPeerRemoteValue", types.YLeaf{"CikeTunHistPeerRemoteValue", cikeTunnelHistEntry.CikeTunHistPeerRemoteValue})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistLocalAddr", types.YLeaf{"CikeTunHistLocalAddr", cikeTunnelHistEntry.CikeTunHistLocalAddr})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistLocalName", types.YLeaf{"CikeTunHistLocalName", cikeTunnelHistEntry.CikeTunHistLocalName})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistRemoteAddr", types.YLeaf{"CikeTunHistRemoteAddr", cikeTunnelHistEntry.CikeTunHistRemoteAddr})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistRemoteName", types.YLeaf{"CikeTunHistRemoteName", cikeTunnelHistEntry.CikeTunHistRemoteName})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistNegoMode", types.YLeaf{"CikeTunHistNegoMode", cikeTunnelHistEntry.CikeTunHistNegoMode})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistDiffHellmanGrp", types.YLeaf{"CikeTunHistDiffHellmanGrp", cikeTunnelHistEntry.CikeTunHistDiffHellmanGrp})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistEncryptAlgo", types.YLeaf{"CikeTunHistEncryptAlgo", cikeTunnelHistEntry.CikeTunHistEncryptAlgo})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistHashAlgo", types.YLeaf{"CikeTunHistHashAlgo", cikeTunnelHistEntry.CikeTunHistHashAlgo})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistAuthMethod", types.YLeaf{"CikeTunHistAuthMethod", cikeTunnelHistEntry.CikeTunHistAuthMethod})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistLifeTime", types.YLeaf{"CikeTunHistLifeTime", cikeTunnelHistEntry.CikeTunHistLifeTime})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistStartTime", types.YLeaf{"CikeTunHistStartTime", cikeTunnelHistEntry.CikeTunHistStartTime})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistActiveTime", types.YLeaf{"CikeTunHistActiveTime", cikeTunnelHistEntry.CikeTunHistActiveTime})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistTotalRefreshes", types.YLeaf{"CikeTunHistTotalRefreshes", cikeTunnelHistEntry.CikeTunHistTotalRefreshes})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistTotalSas", types.YLeaf{"CikeTunHistTotalSas", cikeTunnelHistEntry.CikeTunHistTotalSas})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistInOctets", types.YLeaf{"CikeTunHistInOctets", cikeTunnelHistEntry.CikeTunHistInOctets})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistInPkts", types.YLeaf{"CikeTunHistInPkts", cikeTunnelHistEntry.CikeTunHistInPkts})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistInDropPkts", types.YLeaf{"CikeTunHistInDropPkts", cikeTunnelHistEntry.CikeTunHistInDropPkts})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistInNotifys", types.YLeaf{"CikeTunHistInNotifys", cikeTunnelHistEntry.CikeTunHistInNotifys})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistInP2Exchgs", types.YLeaf{"CikeTunHistInP2Exchgs", cikeTunnelHistEntry.CikeTunHistInP2Exchgs})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistInP2ExchgInvalids", types.YLeaf{"CikeTunHistInP2ExchgInvalids", cikeTunnelHistEntry.CikeTunHistInP2ExchgInvalids})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistInP2ExchgRejects", types.YLeaf{"CikeTunHistInP2ExchgRejects", cikeTunnelHistEntry.CikeTunHistInP2ExchgRejects})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistInP2SaDelRequests", types.YLeaf{"CikeTunHistInP2SaDelRequests", cikeTunnelHistEntry.CikeTunHistInP2SaDelRequests})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistOutOctets", types.YLeaf{"CikeTunHistOutOctets", cikeTunnelHistEntry.CikeTunHistOutOctets})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistOutPkts", types.YLeaf{"CikeTunHistOutPkts", cikeTunnelHistEntry.CikeTunHistOutPkts})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistOutDropPkts", types.YLeaf{"CikeTunHistOutDropPkts", cikeTunnelHistEntry.CikeTunHistOutDropPkts})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistOutNotifys", types.YLeaf{"CikeTunHistOutNotifys", cikeTunnelHistEntry.CikeTunHistOutNotifys})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistOutP2Exchgs", types.YLeaf{"CikeTunHistOutP2Exchgs", cikeTunnelHistEntry.CikeTunHistOutP2Exchgs})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistOutP2ExchgInvalids", types.YLeaf{"CikeTunHistOutP2ExchgInvalids", cikeTunnelHistEntry.CikeTunHistOutP2ExchgInvalids})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistOutP2ExchgRejects", types.YLeaf{"CikeTunHistOutP2ExchgRejects", cikeTunnelHistEntry.CikeTunHistOutP2ExchgRejects})
    cikeTunnelHistEntry.EntityData.Leafs.Append("cikeTunHistOutP2SaDelRequests", types.YLeaf{"CikeTunHistOutP2SaDelRequests", cikeTunnelHistEntry.CikeTunHistOutP2SaDelRequests})

    cikeTunnelHistEntry.EntityData.YListKeys = []string {"CikeTunHistIndex"}

    return &(cikeTunnelHistEntry.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CikeTunnelHistTable_CikeTunnelHistEntry_CikeTunHistTermReason represents 7 = operator initiated check point request
type CISCOIPSECFLOWMONITORMIB_CikeTunnelHistTable_CikeTunnelHistEntry_CikeTunHistTermReason string

const (
    CISCOIPSECFLOWMONITORMIB_CikeTunnelHistTable_CikeTunnelHistEntry_CikeTunHistTermReason_other CISCOIPSECFLOWMONITORMIB_CikeTunnelHistTable_CikeTunnelHistEntry_CikeTunHistTermReason = "other"

    CISCOIPSECFLOWMONITORMIB_CikeTunnelHistTable_CikeTunnelHistEntry_CikeTunHistTermReason_normal CISCOIPSECFLOWMONITORMIB_CikeTunnelHistTable_CikeTunnelHistEntry_CikeTunHistTermReason = "normal"

    CISCOIPSECFLOWMONITORMIB_CikeTunnelHistTable_CikeTunnelHistEntry_CikeTunHistTermReason_operRequest CISCOIPSECFLOWMONITORMIB_CikeTunnelHistTable_CikeTunnelHistEntry_CikeTunHistTermReason = "operRequest"

    CISCOIPSECFLOWMONITORMIB_CikeTunnelHistTable_CikeTunnelHistEntry_CikeTunHistTermReason_peerDelRequest CISCOIPSECFLOWMONITORMIB_CikeTunnelHistTable_CikeTunnelHistEntry_CikeTunHistTermReason = "peerDelRequest"

    CISCOIPSECFLOWMONITORMIB_CikeTunnelHistTable_CikeTunnelHistEntry_CikeTunHistTermReason_peerLost CISCOIPSECFLOWMONITORMIB_CikeTunnelHistTable_CikeTunnelHistEntry_CikeTunHistTermReason = "peerLost"

    CISCOIPSECFLOWMONITORMIB_CikeTunnelHistTable_CikeTunnelHistEntry_CikeTunHistTermReason_localFailure CISCOIPSECFLOWMONITORMIB_CikeTunnelHistTable_CikeTunnelHistEntry_CikeTunHistTermReason = "localFailure"

    CISCOIPSECFLOWMONITORMIB_CikeTunnelHistTable_CikeTunnelHistEntry_CikeTunHistTermReason_checkPointReg CISCOIPSECFLOWMONITORMIB_CikeTunnelHistTable_CikeTunnelHistEntry_CikeTunHistTermReason = "checkPointReg"
)

// CISCOIPSECFLOWMONITORMIB_CipSecTunnelHistTable
// The IPsec Phase-2 Tunnel History Table.
// This table is implemented as a sliding 
// window in which only the
// last n entries are maintained.  The maximum number 
// of entries
// is specified by the cipSecHistTableSize object.
type CISCOIPSECFLOWMONITORMIB_CipSecTunnelHistTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with a previously active
    // IPsec Phase-2 Tunnel. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_CipSecTunnelHistTable_CipSecTunnelHistEntry.
    CipSecTunnelHistEntry []*CISCOIPSECFLOWMONITORMIB_CipSecTunnelHistTable_CipSecTunnelHistEntry
}

func (cipSecTunnelHistTable *CISCOIPSECFLOWMONITORMIB_CipSecTunnelHistTable) GetEntityData() *types.CommonEntityData {
    cipSecTunnelHistTable.EntityData.YFilter = cipSecTunnelHistTable.YFilter
    cipSecTunnelHistTable.EntityData.YangName = "cipSecTunnelHistTable"
    cipSecTunnelHistTable.EntityData.BundleName = "cisco_ios_xe"
    cipSecTunnelHistTable.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cipSecTunnelHistTable.EntityData.SegmentPath = "cipSecTunnelHistTable"
    cipSecTunnelHistTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipSecTunnelHistTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipSecTunnelHistTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipSecTunnelHistTable.EntityData.Children = types.NewOrderedMap()
    cipSecTunnelHistTable.EntityData.Children.Append("cipSecTunnelHistEntry", types.YChild{"CipSecTunnelHistEntry", nil})
    for i := range cipSecTunnelHistTable.CipSecTunnelHistEntry {
        cipSecTunnelHistTable.EntityData.Children.Append(types.GetSegmentPath(cipSecTunnelHistTable.CipSecTunnelHistEntry[i]), types.YChild{"CipSecTunnelHistEntry", cipSecTunnelHistTable.CipSecTunnelHistEntry[i]})
    }
    cipSecTunnelHistTable.EntityData.Leafs = types.NewOrderedMap()

    cipSecTunnelHistTable.EntityData.YListKeys = []string {}

    return &(cipSecTunnelHistTable.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CipSecTunnelHistTable_CipSecTunnelHistEntry
// Each entry contains the attributes associated with
// a previously active IPsec Phase-2 Tunnel.
type CISCOIPSECFLOWMONITORMIB_CipSecTunnelHistTable_CipSecTunnelHistEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index of the IPsec Phase-2 Tunnel History
    // Table. The value of the index is a number which  begins at one and is
    // incremented with each tunnel  that ends. The value of this object will wrap
    // at 2,147,483,647. The type is interface{} with range: 1..2147483647.
    CipSecTunHistIndex interface{}

    // The reason the IPsec Phase-2 Tunnel was terminated. Possible reasons
    // include: 1 = other 2 = normal termination 3 = operator request 4 = peer
    // delete request was received 5 = contact with peer was lost 6 = local
    // failure occurred 7 = operator initiated check point request. The type is
    // CipSecTunHistTermReason.
    CipSecTunHistTermReason interface{}

    // The index of the previously active IPsec Phase-2 Tunnel. The type is
    // interface{} with range: 1..2147483647.
    CipSecTunHistActiveIndex interface{}

    // The index of the associated IPsec Phase-1 Tunnel (cikeTunIndex in the
    // cikeTunnelTable). The type is interface{} with range: 1..2147483647.
    CipSecTunHistIkeTunnelIndex interface{}

    // The IP address of the local endpoint for the IPsec Phase-2 Tunnel. The type
    // is string with length: 4 | 16.
    CipSecTunHistLocalAddr interface{}

    // The IP address of the remote endpoint for the IPsec Phase-2 Tunnel. The
    // type is string with length: 4 | 16.
    CipSecTunHistRemoteAddr interface{}

    // The type of key used by the IPsec Phase-2 Tunnel. The type is KeyType.
    CipSecTunHistKeyType interface{}

    // The encapsulation mode used by the IPsec Phase-2 Tunnel. The type is
    // EncapMode.
    CipSecTunHistEncapMode interface{}

    // The negotiated LifeSize of the IPsec Phase-2 Tunnel in kilobytes. The type
    // is interface{} with range: 1..2147483647. Units are KBytes.
    CipSecTunHistLifeSize interface{}

    // The negotiated LifeTime of the IPsec Phase-2 Tunnel in seconds. The type is
    // interface{} with range: 1..2147483647. Units are Seconds.
    CipSecTunHistLifeTime interface{}

    // The value of sysUpTime in hundredths of seconds when the IPsec Phase-2
    // Tunnel was started. The type is interface{} with range: 0..4294967295.
    CipSecTunHistStartTime interface{}

    // The length of time the IPsec Phase-2 Tunnel has been active in hundredths
    // of seconds. The type is interface{} with range: 0..2147483647.
    CipSecTunHistActiveTime interface{}

    // The total number of security association refreshes performed. The type is
    // interface{} with range: 0..4294967295. Units are QM Exchanges.
    CipSecTunHistTotalRefreshes interface{}

    // The total number of security associations used during the  life of the
    // IPsec Phase-2 Tunnel. The type is interface{} with range: 0..4294967295.
    // Units are SAs.
    CipSecTunHistTotalSas interface{}

    // The Diffie Hellman Group used by the inbound security association of the
    // IPsec Phase-2 Tunnel. The type is DiffHellmanGrp.
    CipSecTunHistInSaDiffHellmanGrp interface{}

    // The encryption algorithm used by the inbound security association of the
    // IPsec Phase-2 Tunnel. The type is EncryptAlgo.
    CipSecTunHistInSaEncryptAlgo interface{}

    // The authentication algorithm used by the inbound authentication header (AH)
    // security association of the IPsec Phase-2 Tunnel. The type is AuthAlgo.
    CipSecTunHistInSaAhAuthAlgo interface{}

    // The authentication algorithm used by the inbound encapsulation security
    // protocol (ESP)  security association of the IPsec Phase-2 Tunnel. The type
    // is AuthAlgo.
    CipSecTunHistInSaEspAuthAlgo interface{}

    // The decompression algorithm used by the inbound security association of the
    // IPsec Phase-2 Tunnel. The type is CompAlgo.
    CipSecTunHistInSaDecompAlgo interface{}

    // The Diffie Hellman Group used by the outbound security association of the
    // IPsec Phase-2 Tunnel. The type is DiffHellmanGrp.
    CipSecTunHistOutSaDiffHellmanGrp interface{}

    // The encryption algorithm used by the outbound security association of the
    // IPsec Phase-2 Tunnel. The type is EncryptAlgo.
    CipSecTunHistOutSaEncryptAlgo interface{}

    // The authentication algorithm used by the outbound authentication header
    // (AH) security association of the IPsec Phase-2 Tunnel. The type is
    // AuthAlgo.
    CipSecTunHistOutSaAhAuthAlgo interface{}

    // The authentication algorithm used by the inbound encapsulation security
    // protocol (ESP)  security association of the IPsec Phase-2 Tunnel. The type
    // is AuthAlgo.
    CipSecTunHistOutSaEspAuthAlgo interface{}

    // The compression algorithm used by the inbound security association of the
    // IPsec Phase-2 Tunnel. The type is CompAlgo.
    CipSecTunHistOutSaCompAlgo interface{}

    // The total number of octets received by this IPsec Phase-2 Tunnel.  This
    // value is accumulated BEFORE determining whether or not the packet should 
    // be decompressed.  See also cipSecTunInOctWraps for  the number of times
    // this counter has wrapped. The type is interface{} with range:
    // 0..4294967295. Units are Octets.
    CipSecTunHistInOctets interface{}

    // A high capacity count of the total number of octets received by this IPsec
    // Phase-2 Tunnel.  This value is accumulated BEFORE determining whether or
    // not  the packet should be decompressed. The type is interface{} with range:
    // 0..18446744073709551615.
    CipSecTunHistHcInOctets interface{}

    // The number of times the octets received counter (cipSecTunInOctets) has
    // wrapped. The type is interface{} with range: 0..4294967295. Units are
    // Integral units.
    CipSecTunHistInOctWraps interface{}

    // The total number of decompressed octets received by this IPsec Phase-2
    // Tunnel.  This value is accumulated AFTER the packet is decompressed. If
    // compression is not being used, this value will match the value of
    // cipSecTunHistInOctets. See also cipSecTunInDecompOctWraps for the number of
    // times this counter has wrapped. The type is interface{} with range:
    // 0..4294967295. Units are Octets.
    CipSecTunHistInDecompOctets interface{}

    // A high capacity count of the total number of decompressed octets received
    // by this IPsec Phase-2 Tunnel.  This value is accumulated AFTER the packet
    // is decompressed. If compression is not being used, this value will match
    // the value of cipSecTunHistHcInOctets. The type is interface{} with range:
    // 0..18446744073709551615.
    CipSecTunHistHcInDecompOctets interface{}

    // The number of times the decompressed octets received counter
    // (cipSecTunInDecompOctets) has wrapped. The type is interface{} with range:
    // 0..4294967295. Units are Integral units.
    CipSecTunHistInDecompOctWraps interface{}

    // The total number of packets received by this IPsec Phase-2 Tunnel. The type
    // is interface{} with range: 0..4294967295. Units are Packets.
    CipSecTunHistInPkts interface{}

    // The total number of packets dropped during receive processing by this IPsec
    // Phase-2 Tunnel.  This count does NOT include packets  dropped due to
    // Anti-Replay processing. The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    CipSecTunHistInDropPkts interface{}

    // The total number of packets dropped during receive processing due to
    // Anti-Replay processing  by this IPsec Phase-2 Tunnel. The type is
    // interface{} with range: 0..4294967295. Units are Packets.
    CipSecTunHistInReplayDropPkts interface{}

    // The total number of inbound authentication's performed  by this IPsec
    // Phase-2 Tunnel. The type is interface{} with range: 0..4294967295. Units
    // are Events.
    CipSecTunHistInAuths interface{}

    // The total number of inbound authentication's which ended in  failure by
    // this IPsec Phase-2 Tunnel . The type is interface{} with range:
    // 0..4294967295. Units are Failures.
    CipSecTunHistInAuthFails interface{}

    // The total number of inbound decryption's performed by this IPsec Phase-2
    // Tunnel. The type is interface{} with range: 0..4294967295. Units are
    // Packets.
    CipSecTunHistInDecrypts interface{}

    // The total number of inbound decryption's which ended in failure  by this
    // IPsec Phase-2 Tunnel. The type is interface{} with range: 0..4294967295.
    // Units are Failures.
    CipSecTunHistInDecryptFails interface{}

    // The total number of octets sent by this IPsec Phase-2 Tunnel.  This value
    // is accumulated AFTER determining whether or not the  packet should be
    // compressed.  See also cipSecTunOutOctWraps for the number of times this
    // counter has wrapped. The type is interface{} with range: 0..4294967295.
    // Units are Octets.
    CipSecTunHistOutOctets interface{}

    // A high capacity count of the total number of octets sent by this IPsec
    // Phase-2 Tunnel.  This value  is accumulated AFTER determining whether or
    // not  the packet should be compressed. The type is interface{} with range:
    // 0..18446744073709551615.
    CipSecTunHistHcOutOctets interface{}

    // The number of times the octets sent counter (cipSecTunOutOctets) has
    // wrapped. The type is interface{} with range: 0..4294967295. Units are
    // Integral units.
    CipSecTunHistOutOctWraps interface{}

    // The total number of uncompressed octets sent by this IPsec Phase-2 Tunnel. 
    // This value is accumulated BEFORE the packet is compressed. If compression
    // is not being used, this value will match the value of 
    // cipSecTunHistOutOctets.  See also  cipSecTunOutDecompOctWraps for the
    // number of times this counter has wrapped. The type is interface{} with
    // range: 0..4294967295. Units are Octets.
    CipSecTunHistOutUncompOctets interface{}

    // A high capacity count of the total number of uncompressed octets sent by
    // this  IPsec Phase-2 Tunnel.  This value is accumulated  BEFORE the packet
    // is compressed. If compression is not being used, this value will match the
    // value of cipSecTunHistHcOutOctets. The type is interface{} with range:
    // 0..18446744073709551615. Units are Octets.
    CipSecTunHistHcOutUncompOctets interface{}

    // The number of times the uncompressed octets sent counter
    // (cipSecTunOutUncompOctets) has wrapped. The type is interface{} with range:
    // 0..4294967295. Units are Integral units.
    CipSecTunHistOutUncompOctWraps interface{}

    // The total number of packets sent by this IPsec Phase-2 Tunnel. The type is
    // interface{} with range: 0..4294967295. Units are Packets.
    CipSecTunHistOutPkts interface{}

    // The total number of packets dropped during send processing  by this IPsec
    // Phase-2 Tunnel. The type is interface{} with range: 0..4294967295. Units
    // are Packets.
    CipSecTunHistOutDropPkts interface{}

    // The total number of outbound authentication's performed by this IPsec
    // Phase-2 Tunnel. The type is interface{} with range: 0..4294967295. Units
    // are Events.
    CipSecTunHistOutAuths interface{}

    // The total number of outbound authentication's which ended in  failure by
    // this IPsec Phase-2 Tunnel. The type is interface{} with range:
    // 0..4294967295. Units are Failures.
    CipSecTunHistOutAuthFails interface{}

    // The total number of outbound encryption's performed by this IPsec Phase-2
    // Tunnel. The type is interface{} with range: 0..4294967295. Units are
    // Packets.
    CipSecTunHistOutEncrypts interface{}

    // The total number of outbound encryption's which ended in failure  by this
    // IPsec Phase-2 Tunnel. The type is interface{} with range: 0..4294967295.
    // Units are Failures.
    CipSecTunHistOutEncryptFails interface{}
}

func (cipSecTunnelHistEntry *CISCOIPSECFLOWMONITORMIB_CipSecTunnelHistTable_CipSecTunnelHistEntry) GetEntityData() *types.CommonEntityData {
    cipSecTunnelHistEntry.EntityData.YFilter = cipSecTunnelHistEntry.YFilter
    cipSecTunnelHistEntry.EntityData.YangName = "cipSecTunnelHistEntry"
    cipSecTunnelHistEntry.EntityData.BundleName = "cisco_ios_xe"
    cipSecTunnelHistEntry.EntityData.ParentYangName = "cipSecTunnelHistTable"
    cipSecTunnelHistEntry.EntityData.SegmentPath = "cipSecTunnelHistEntry" + types.AddKeyToken(cipSecTunnelHistEntry.CipSecTunHistIndex, "cipSecTunHistIndex")
    cipSecTunnelHistEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipSecTunnelHistEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipSecTunnelHistEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipSecTunnelHistEntry.EntityData.Children = types.NewOrderedMap()
    cipSecTunnelHistEntry.EntityData.Leafs = types.NewOrderedMap()
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistIndex", types.YLeaf{"CipSecTunHistIndex", cipSecTunnelHistEntry.CipSecTunHistIndex})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistTermReason", types.YLeaf{"CipSecTunHistTermReason", cipSecTunnelHistEntry.CipSecTunHistTermReason})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistActiveIndex", types.YLeaf{"CipSecTunHistActiveIndex", cipSecTunnelHistEntry.CipSecTunHistActiveIndex})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistIkeTunnelIndex", types.YLeaf{"CipSecTunHistIkeTunnelIndex", cipSecTunnelHistEntry.CipSecTunHistIkeTunnelIndex})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistLocalAddr", types.YLeaf{"CipSecTunHistLocalAddr", cipSecTunnelHistEntry.CipSecTunHistLocalAddr})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistRemoteAddr", types.YLeaf{"CipSecTunHistRemoteAddr", cipSecTunnelHistEntry.CipSecTunHistRemoteAddr})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistKeyType", types.YLeaf{"CipSecTunHistKeyType", cipSecTunnelHistEntry.CipSecTunHistKeyType})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistEncapMode", types.YLeaf{"CipSecTunHistEncapMode", cipSecTunnelHistEntry.CipSecTunHistEncapMode})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistLifeSize", types.YLeaf{"CipSecTunHistLifeSize", cipSecTunnelHistEntry.CipSecTunHistLifeSize})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistLifeTime", types.YLeaf{"CipSecTunHistLifeTime", cipSecTunnelHistEntry.CipSecTunHistLifeTime})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistStartTime", types.YLeaf{"CipSecTunHistStartTime", cipSecTunnelHistEntry.CipSecTunHistStartTime})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistActiveTime", types.YLeaf{"CipSecTunHistActiveTime", cipSecTunnelHistEntry.CipSecTunHistActiveTime})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistTotalRefreshes", types.YLeaf{"CipSecTunHistTotalRefreshes", cipSecTunnelHistEntry.CipSecTunHistTotalRefreshes})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistTotalSas", types.YLeaf{"CipSecTunHistTotalSas", cipSecTunnelHistEntry.CipSecTunHistTotalSas})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistInSaDiffHellmanGrp", types.YLeaf{"CipSecTunHistInSaDiffHellmanGrp", cipSecTunnelHistEntry.CipSecTunHistInSaDiffHellmanGrp})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistInSaEncryptAlgo", types.YLeaf{"CipSecTunHistInSaEncryptAlgo", cipSecTunnelHistEntry.CipSecTunHistInSaEncryptAlgo})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistInSaAhAuthAlgo", types.YLeaf{"CipSecTunHistInSaAhAuthAlgo", cipSecTunnelHistEntry.CipSecTunHistInSaAhAuthAlgo})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistInSaEspAuthAlgo", types.YLeaf{"CipSecTunHistInSaEspAuthAlgo", cipSecTunnelHistEntry.CipSecTunHistInSaEspAuthAlgo})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistInSaDecompAlgo", types.YLeaf{"CipSecTunHistInSaDecompAlgo", cipSecTunnelHistEntry.CipSecTunHistInSaDecompAlgo})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistOutSaDiffHellmanGrp", types.YLeaf{"CipSecTunHistOutSaDiffHellmanGrp", cipSecTunnelHistEntry.CipSecTunHistOutSaDiffHellmanGrp})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistOutSaEncryptAlgo", types.YLeaf{"CipSecTunHistOutSaEncryptAlgo", cipSecTunnelHistEntry.CipSecTunHistOutSaEncryptAlgo})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistOutSaAhAuthAlgo", types.YLeaf{"CipSecTunHistOutSaAhAuthAlgo", cipSecTunnelHistEntry.CipSecTunHistOutSaAhAuthAlgo})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistOutSaEspAuthAlgo", types.YLeaf{"CipSecTunHistOutSaEspAuthAlgo", cipSecTunnelHistEntry.CipSecTunHistOutSaEspAuthAlgo})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistOutSaCompAlgo", types.YLeaf{"CipSecTunHistOutSaCompAlgo", cipSecTunnelHistEntry.CipSecTunHistOutSaCompAlgo})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistInOctets", types.YLeaf{"CipSecTunHistInOctets", cipSecTunnelHistEntry.CipSecTunHistInOctets})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistHcInOctets", types.YLeaf{"CipSecTunHistHcInOctets", cipSecTunnelHistEntry.CipSecTunHistHcInOctets})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistInOctWraps", types.YLeaf{"CipSecTunHistInOctWraps", cipSecTunnelHistEntry.CipSecTunHistInOctWraps})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistInDecompOctets", types.YLeaf{"CipSecTunHistInDecompOctets", cipSecTunnelHistEntry.CipSecTunHistInDecompOctets})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistHcInDecompOctets", types.YLeaf{"CipSecTunHistHcInDecompOctets", cipSecTunnelHistEntry.CipSecTunHistHcInDecompOctets})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistInDecompOctWraps", types.YLeaf{"CipSecTunHistInDecompOctWraps", cipSecTunnelHistEntry.CipSecTunHistInDecompOctWraps})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistInPkts", types.YLeaf{"CipSecTunHistInPkts", cipSecTunnelHistEntry.CipSecTunHistInPkts})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistInDropPkts", types.YLeaf{"CipSecTunHistInDropPkts", cipSecTunnelHistEntry.CipSecTunHistInDropPkts})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistInReplayDropPkts", types.YLeaf{"CipSecTunHistInReplayDropPkts", cipSecTunnelHistEntry.CipSecTunHistInReplayDropPkts})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistInAuths", types.YLeaf{"CipSecTunHistInAuths", cipSecTunnelHistEntry.CipSecTunHistInAuths})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistInAuthFails", types.YLeaf{"CipSecTunHistInAuthFails", cipSecTunnelHistEntry.CipSecTunHistInAuthFails})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistInDecrypts", types.YLeaf{"CipSecTunHistInDecrypts", cipSecTunnelHistEntry.CipSecTunHistInDecrypts})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistInDecryptFails", types.YLeaf{"CipSecTunHistInDecryptFails", cipSecTunnelHistEntry.CipSecTunHistInDecryptFails})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistOutOctets", types.YLeaf{"CipSecTunHistOutOctets", cipSecTunnelHistEntry.CipSecTunHistOutOctets})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistHcOutOctets", types.YLeaf{"CipSecTunHistHcOutOctets", cipSecTunnelHistEntry.CipSecTunHistHcOutOctets})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistOutOctWraps", types.YLeaf{"CipSecTunHistOutOctWraps", cipSecTunnelHistEntry.CipSecTunHistOutOctWraps})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistOutUncompOctets", types.YLeaf{"CipSecTunHistOutUncompOctets", cipSecTunnelHistEntry.CipSecTunHistOutUncompOctets})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistHcOutUncompOctets", types.YLeaf{"CipSecTunHistHcOutUncompOctets", cipSecTunnelHistEntry.CipSecTunHistHcOutUncompOctets})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistOutUncompOctWraps", types.YLeaf{"CipSecTunHistOutUncompOctWraps", cipSecTunnelHistEntry.CipSecTunHistOutUncompOctWraps})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistOutPkts", types.YLeaf{"CipSecTunHistOutPkts", cipSecTunnelHistEntry.CipSecTunHistOutPkts})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistOutDropPkts", types.YLeaf{"CipSecTunHistOutDropPkts", cipSecTunnelHistEntry.CipSecTunHistOutDropPkts})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistOutAuths", types.YLeaf{"CipSecTunHistOutAuths", cipSecTunnelHistEntry.CipSecTunHistOutAuths})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistOutAuthFails", types.YLeaf{"CipSecTunHistOutAuthFails", cipSecTunnelHistEntry.CipSecTunHistOutAuthFails})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistOutEncrypts", types.YLeaf{"CipSecTunHistOutEncrypts", cipSecTunnelHistEntry.CipSecTunHistOutEncrypts})
    cipSecTunnelHistEntry.EntityData.Leafs.Append("cipSecTunHistOutEncryptFails", types.YLeaf{"CipSecTunHistOutEncryptFails", cipSecTunnelHistEntry.CipSecTunHistOutEncryptFails})

    cipSecTunnelHistEntry.EntityData.YListKeys = []string {"CipSecTunHistIndex"}

    return &(cipSecTunnelHistEntry.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CipSecTunnelHistTable_CipSecTunnelHistEntry_CipSecTunHistTermReason represents 7 = operator initiated check point request
type CISCOIPSECFLOWMONITORMIB_CipSecTunnelHistTable_CipSecTunnelHistEntry_CipSecTunHistTermReason string

const (
    CISCOIPSECFLOWMONITORMIB_CipSecTunnelHistTable_CipSecTunnelHistEntry_CipSecTunHistTermReason_other CISCOIPSECFLOWMONITORMIB_CipSecTunnelHistTable_CipSecTunnelHistEntry_CipSecTunHistTermReason = "other"

    CISCOIPSECFLOWMONITORMIB_CipSecTunnelHistTable_CipSecTunnelHistEntry_CipSecTunHistTermReason_normal CISCOIPSECFLOWMONITORMIB_CipSecTunnelHistTable_CipSecTunnelHistEntry_CipSecTunHistTermReason = "normal"

    CISCOIPSECFLOWMONITORMIB_CipSecTunnelHistTable_CipSecTunnelHistEntry_CipSecTunHistTermReason_operRequest CISCOIPSECFLOWMONITORMIB_CipSecTunnelHistTable_CipSecTunnelHistEntry_CipSecTunHistTermReason = "operRequest"

    CISCOIPSECFLOWMONITORMIB_CipSecTunnelHistTable_CipSecTunnelHistEntry_CipSecTunHistTermReason_peerDelRequest CISCOIPSECFLOWMONITORMIB_CipSecTunnelHistTable_CipSecTunnelHistEntry_CipSecTunHistTermReason = "peerDelRequest"

    CISCOIPSECFLOWMONITORMIB_CipSecTunnelHistTable_CipSecTunnelHistEntry_CipSecTunHistTermReason_peerLost CISCOIPSECFLOWMONITORMIB_CipSecTunnelHistTable_CipSecTunnelHistEntry_CipSecTunHistTermReason = "peerLost"

    CISCOIPSECFLOWMONITORMIB_CipSecTunnelHistTable_CipSecTunnelHistEntry_CipSecTunHistTermReason_seqNumRollOver CISCOIPSECFLOWMONITORMIB_CipSecTunnelHistTable_CipSecTunnelHistEntry_CipSecTunHistTermReason = "seqNumRollOver"

    CISCOIPSECFLOWMONITORMIB_CipSecTunnelHistTable_CipSecTunnelHistEntry_CipSecTunHistTermReason_checkPointReq CISCOIPSECFLOWMONITORMIB_CipSecTunnelHistTable_CipSecTunnelHistEntry_CipSecTunHistTermReason = "checkPointReq"
)

// CISCOIPSECFLOWMONITORMIB_CipSecEndPtHistTable
// The IPsec Phase-2 Tunnel Endpoint History Table.
// This table is implemented as a 
// sliding window in which only the
// last n entries are maintained.  
// The maximum number of entries
// is specified by the cipSecHistTableSize object.
type CISCOIPSECFLOWMONITORMIB_CipSecEndPtHistTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with a previously active
    // IPsec Phase-2 Tunnel Endpoint. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_CipSecEndPtHistTable_CipSecEndPtHistEntry.
    CipSecEndPtHistEntry []*CISCOIPSECFLOWMONITORMIB_CipSecEndPtHistTable_CipSecEndPtHistEntry
}

func (cipSecEndPtHistTable *CISCOIPSECFLOWMONITORMIB_CipSecEndPtHistTable) GetEntityData() *types.CommonEntityData {
    cipSecEndPtHistTable.EntityData.YFilter = cipSecEndPtHistTable.YFilter
    cipSecEndPtHistTable.EntityData.YangName = "cipSecEndPtHistTable"
    cipSecEndPtHistTable.EntityData.BundleName = "cisco_ios_xe"
    cipSecEndPtHistTable.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cipSecEndPtHistTable.EntityData.SegmentPath = "cipSecEndPtHistTable"
    cipSecEndPtHistTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipSecEndPtHistTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipSecEndPtHistTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipSecEndPtHistTable.EntityData.Children = types.NewOrderedMap()
    cipSecEndPtHistTable.EntityData.Children.Append("cipSecEndPtHistEntry", types.YChild{"CipSecEndPtHistEntry", nil})
    for i := range cipSecEndPtHistTable.CipSecEndPtHistEntry {
        cipSecEndPtHistTable.EntityData.Children.Append(types.GetSegmentPath(cipSecEndPtHistTable.CipSecEndPtHistEntry[i]), types.YChild{"CipSecEndPtHistEntry", cipSecEndPtHistTable.CipSecEndPtHistEntry[i]})
    }
    cipSecEndPtHistTable.EntityData.Leafs = types.NewOrderedMap()

    cipSecEndPtHistTable.EntityData.YListKeys = []string {}

    return &(cipSecEndPtHistTable.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CipSecEndPtHistTable_CipSecEndPtHistEntry
// Each entry contains the attributes associated with
// a previously active IPsec Phase-2 Tunnel Endpoint.
type CISCOIPSECFLOWMONITORMIB_CipSecEndPtHistTable_CipSecEndPtHistEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The number of the previously active Endpoint
    // associated  with a IPsec Phase-2 Tunnel Table.  The value   of this index
    // is a number which begins at   one and is incremented with each Endpoint  
    // associated with an IPsec Phase-2 Tunnel.  The value of this object will
    // wrap at 2,147,483,647. The type is interface{} with range: 1..2147483647.
    CipSecEndPtHistIndex interface{}

    // The index  of the previously active IPsec Phase-2 Tunnel Table. The type is
    // interface{} with range: 1..2147483647.
    CipSecEndPtHistTunIndex interface{}

    // The index  of the previously active Endpoint. The type is interface{} with
    // range: 1..2147483647.
    CipSecEndPtHistActiveIndex interface{}

    // The DNS name of the local Endpoint. The type is string.
    CipSecEndPtHistLocalName interface{}

    // The type of identity for the local Endpoint. Possible values are: 1) a
    // single IP address, or 2) an IP address range, or 3) an IP subnet. The type
    // is EndPtType.
    CipSecEndPtHistLocalType interface{}

    // The local Endpoint's first IP address specification.  If the local Endpoint
    // type is single IP address,  then this is the value of the IP address.  If
    // the local Endpoint type is IP subnet, then this is the value of the subnet.
    // If the local Endpoint type is IP address range,  then this is the value of
    // beginning IP address of  the range. The type is string with length: 4 | 16.
    CipSecEndPtHistLocalAddr1 interface{}

    // The local Endpoint's second IP address specification.  If the local
    // Endpoint type is single IP address,  then this is the value of the IP
    // address.  If the local Endpoint type is IP subnet, then this is the value
    // of the subnet mask.  If the local Endpoint type is IP address range,  then
    // this is the value of ending IP address of the range. The type is string
    // with length: 4 | 16.
    CipSecEndPtHistLocalAddr2 interface{}

    // The protocol number of the local Endpoint's traffic. The type is
    // interface{} with range: 0..255.
    CipSecEndPtHistLocalProtocol interface{}

    // The port number of the local Endpoint's traffic. The type is interface{}
    // with range: 0..65535.
    CipSecEndPtHistLocalPort interface{}

    // The DNS name of the remote Endpoint. The type is string.
    CipSecEndPtHistRemoteName interface{}

    // The type of identity for the remote Endpoint. Possible values are: 1) a
    // single IP address, or 2) an IP address range, or 3) an IP subnet. The type
    // is EndPtType.
    CipSecEndPtHistRemoteType interface{}

    // The remote Endpoint's first IP address specification.  If the remote
    // Endpoint type is single IP address,  then this is the value of the IP
    // address.  If the remote Endpoint type is IP subnet, then this is the value
    // of the subnet.  If the remote Endpoint type is IP address range,  then this
    // is the value of beginning IP address of the range. The type is string with
    // length: 4 | 16.
    CipSecEndPtHistRemoteAddr1 interface{}

    // The remote Endpoint's second IP address specification.  If the remote
    // Endpoint type is single IP address,  then this is the value of the IP
    // address.  If the remote Endpoint type is IP subnet, then this is the value
    // of the subnet mask.  If the remote Endpoint type is IP address range,  then
    // this is the value of ending IP address of the range. The type is string
    // with length: 4 | 16.
    CipSecEndPtHistRemoteAddr2 interface{}

    // The protocol number of the remote Endpoint's traffic. The type is
    // interface{} with range: 0..255.
    CipSecEndPtHistRemoteProtocol interface{}

    // The port number of the remote Endpoint's traffic. The type is interface{}
    // with range: 0..65535.
    CipSecEndPtHistRemotePort interface{}
}

func (cipSecEndPtHistEntry *CISCOIPSECFLOWMONITORMIB_CipSecEndPtHistTable_CipSecEndPtHistEntry) GetEntityData() *types.CommonEntityData {
    cipSecEndPtHistEntry.EntityData.YFilter = cipSecEndPtHistEntry.YFilter
    cipSecEndPtHistEntry.EntityData.YangName = "cipSecEndPtHistEntry"
    cipSecEndPtHistEntry.EntityData.BundleName = "cisco_ios_xe"
    cipSecEndPtHistEntry.EntityData.ParentYangName = "cipSecEndPtHistTable"
    cipSecEndPtHistEntry.EntityData.SegmentPath = "cipSecEndPtHistEntry" + types.AddKeyToken(cipSecEndPtHistEntry.CipSecEndPtHistIndex, "cipSecEndPtHistIndex")
    cipSecEndPtHistEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipSecEndPtHistEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipSecEndPtHistEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipSecEndPtHistEntry.EntityData.Children = types.NewOrderedMap()
    cipSecEndPtHistEntry.EntityData.Leafs = types.NewOrderedMap()
    cipSecEndPtHistEntry.EntityData.Leafs.Append("cipSecEndPtHistIndex", types.YLeaf{"CipSecEndPtHistIndex", cipSecEndPtHistEntry.CipSecEndPtHistIndex})
    cipSecEndPtHistEntry.EntityData.Leafs.Append("cipSecEndPtHistTunIndex", types.YLeaf{"CipSecEndPtHistTunIndex", cipSecEndPtHistEntry.CipSecEndPtHistTunIndex})
    cipSecEndPtHistEntry.EntityData.Leafs.Append("cipSecEndPtHistActiveIndex", types.YLeaf{"CipSecEndPtHistActiveIndex", cipSecEndPtHistEntry.CipSecEndPtHistActiveIndex})
    cipSecEndPtHistEntry.EntityData.Leafs.Append("cipSecEndPtHistLocalName", types.YLeaf{"CipSecEndPtHistLocalName", cipSecEndPtHistEntry.CipSecEndPtHistLocalName})
    cipSecEndPtHistEntry.EntityData.Leafs.Append("cipSecEndPtHistLocalType", types.YLeaf{"CipSecEndPtHistLocalType", cipSecEndPtHistEntry.CipSecEndPtHistLocalType})
    cipSecEndPtHistEntry.EntityData.Leafs.Append("cipSecEndPtHistLocalAddr1", types.YLeaf{"CipSecEndPtHistLocalAddr1", cipSecEndPtHistEntry.CipSecEndPtHistLocalAddr1})
    cipSecEndPtHistEntry.EntityData.Leafs.Append("cipSecEndPtHistLocalAddr2", types.YLeaf{"CipSecEndPtHistLocalAddr2", cipSecEndPtHistEntry.CipSecEndPtHistLocalAddr2})
    cipSecEndPtHistEntry.EntityData.Leafs.Append("cipSecEndPtHistLocalProtocol", types.YLeaf{"CipSecEndPtHistLocalProtocol", cipSecEndPtHistEntry.CipSecEndPtHistLocalProtocol})
    cipSecEndPtHistEntry.EntityData.Leafs.Append("cipSecEndPtHistLocalPort", types.YLeaf{"CipSecEndPtHistLocalPort", cipSecEndPtHistEntry.CipSecEndPtHistLocalPort})
    cipSecEndPtHistEntry.EntityData.Leafs.Append("cipSecEndPtHistRemoteName", types.YLeaf{"CipSecEndPtHistRemoteName", cipSecEndPtHistEntry.CipSecEndPtHistRemoteName})
    cipSecEndPtHistEntry.EntityData.Leafs.Append("cipSecEndPtHistRemoteType", types.YLeaf{"CipSecEndPtHistRemoteType", cipSecEndPtHistEntry.CipSecEndPtHistRemoteType})
    cipSecEndPtHistEntry.EntityData.Leafs.Append("cipSecEndPtHistRemoteAddr1", types.YLeaf{"CipSecEndPtHistRemoteAddr1", cipSecEndPtHistEntry.CipSecEndPtHistRemoteAddr1})
    cipSecEndPtHistEntry.EntityData.Leafs.Append("cipSecEndPtHistRemoteAddr2", types.YLeaf{"CipSecEndPtHistRemoteAddr2", cipSecEndPtHistEntry.CipSecEndPtHistRemoteAddr2})
    cipSecEndPtHistEntry.EntityData.Leafs.Append("cipSecEndPtHistRemoteProtocol", types.YLeaf{"CipSecEndPtHistRemoteProtocol", cipSecEndPtHistEntry.CipSecEndPtHistRemoteProtocol})
    cipSecEndPtHistEntry.EntityData.Leafs.Append("cipSecEndPtHistRemotePort", types.YLeaf{"CipSecEndPtHistRemotePort", cipSecEndPtHistEntry.CipSecEndPtHistRemotePort})

    cipSecEndPtHistEntry.EntityData.YListKeys = []string {"CipSecEndPtHistIndex"}

    return &(cipSecEndPtHistEntry.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CikeFailTable
// The IPsec Phase-1 Failure Table.
// This table is implemented as a sliding 
// window in which only the last n entries are 
// maintained.  The maximum number of entries
// is specified by the cipSecFailTableSize object.
type CISCOIPSECFLOWMONITORMIB_CikeFailTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with  an IPsec Phase-1
    // failure. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry.
    CikeFailEntry []*CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry
}

func (cikeFailTable *CISCOIPSECFLOWMONITORMIB_CikeFailTable) GetEntityData() *types.CommonEntityData {
    cikeFailTable.EntityData.YFilter = cikeFailTable.YFilter
    cikeFailTable.EntityData.YangName = "cikeFailTable"
    cikeFailTable.EntityData.BundleName = "cisco_ios_xe"
    cikeFailTable.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cikeFailTable.EntityData.SegmentPath = "cikeFailTable"
    cikeFailTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cikeFailTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cikeFailTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cikeFailTable.EntityData.Children = types.NewOrderedMap()
    cikeFailTable.EntityData.Children.Append("cikeFailEntry", types.YChild{"CikeFailEntry", nil})
    for i := range cikeFailTable.CikeFailEntry {
        cikeFailTable.EntityData.Children.Append(types.GetSegmentPath(cikeFailTable.CikeFailEntry[i]), types.YChild{"CikeFailEntry", cikeFailTable.CikeFailEntry[i]})
    }
    cikeFailTable.EntityData.Leafs = types.NewOrderedMap()

    cikeFailTable.EntityData.YListKeys = []string {}

    return &(cikeFailTable.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry
// Each entry contains the attributes associated
// with
//  an IPsec Phase-1 failure.
type CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IPsec Phase-1 Failure Table index. The value
    // of the index is a number which  begins at one and is incremented with each 
    // IPsec Phase-1 failure. The value of this object will wrap at 2,147,483,647.
    // The type is interface{} with range: 1..2147483647.
    CikeFailIndex interface{}

    // The reason for the failure.  Possible reasons include: 1 = other 2 = peer
    // delete request was received 3 = contact with peer was lost 4 = local
    // failure occurred 5 = authentication failure 6 = hash validation failure 7 =
    // encryption failure 8 = internal error occurred 9 = system capacity failure
    // 10 = proposal failure 11 = peer's certificate is unavailable 12 = peer's
    // certificate was found invalid 13 = local certificate expired 14 =
    // certificate revoke list (crl) failure 15 = peer encoding error 16 =
    // non-existent security association 17 = operator requested termination. The
    // type is CikeFailReason.
    CikeFailReason interface{}

    // The value of sysUpTime in hundredths of seconds at the time of the failure.
    // The type is interface{} with range: 0..4294967295.
    CikeFailTime interface{}

    // The type of local peer identity.  The local peer may be identified by:  1.
    // an IP address, or  2. a host name. The type is IkePeerType.
    CikeFailLocalType interface{}

    // The value of the local peer identity.  If the local peer type is an IP
    // Address, then this is the IP Address used to identify the local peer.  If
    // the local peer type is a host name, then this is the host name used to
    // identify the local peer. The type is string.
    CikeFailLocalValue interface{}

    // The type of remote peer identity.  The remote peer may be identified by: 
    // 1. an IP address, or  2. a host name. The type is IkePeerType.
    CikeFailRemoteType interface{}

    // The value of the remote peer identity.  If the remote peer type is an IP
    // Address, then this is the IP Address used to identify the remote peer.  If
    // the remote peer type is a host name, then this is the host name used to
    // identify the remote peer. The type is string.
    CikeFailRemoteValue interface{}

    // The IP address of the local peer. The type is string with length: 4 | 16.
    CikeFailLocalAddr interface{}

    // The IP address of the remote peer. The type is string with length: 4 | 16.
    CikeFailRemoteAddr interface{}
}

func (cikeFailEntry *CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry) GetEntityData() *types.CommonEntityData {
    cikeFailEntry.EntityData.YFilter = cikeFailEntry.YFilter
    cikeFailEntry.EntityData.YangName = "cikeFailEntry"
    cikeFailEntry.EntityData.BundleName = "cisco_ios_xe"
    cikeFailEntry.EntityData.ParentYangName = "cikeFailTable"
    cikeFailEntry.EntityData.SegmentPath = "cikeFailEntry" + types.AddKeyToken(cikeFailEntry.CikeFailIndex, "cikeFailIndex")
    cikeFailEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cikeFailEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cikeFailEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cikeFailEntry.EntityData.Children = types.NewOrderedMap()
    cikeFailEntry.EntityData.Leafs = types.NewOrderedMap()
    cikeFailEntry.EntityData.Leafs.Append("cikeFailIndex", types.YLeaf{"CikeFailIndex", cikeFailEntry.CikeFailIndex})
    cikeFailEntry.EntityData.Leafs.Append("cikeFailReason", types.YLeaf{"CikeFailReason", cikeFailEntry.CikeFailReason})
    cikeFailEntry.EntityData.Leafs.Append("cikeFailTime", types.YLeaf{"CikeFailTime", cikeFailEntry.CikeFailTime})
    cikeFailEntry.EntityData.Leafs.Append("cikeFailLocalType", types.YLeaf{"CikeFailLocalType", cikeFailEntry.CikeFailLocalType})
    cikeFailEntry.EntityData.Leafs.Append("cikeFailLocalValue", types.YLeaf{"CikeFailLocalValue", cikeFailEntry.CikeFailLocalValue})
    cikeFailEntry.EntityData.Leafs.Append("cikeFailRemoteType", types.YLeaf{"CikeFailRemoteType", cikeFailEntry.CikeFailRemoteType})
    cikeFailEntry.EntityData.Leafs.Append("cikeFailRemoteValue", types.YLeaf{"CikeFailRemoteValue", cikeFailEntry.CikeFailRemoteValue})
    cikeFailEntry.EntityData.Leafs.Append("cikeFailLocalAddr", types.YLeaf{"CikeFailLocalAddr", cikeFailEntry.CikeFailLocalAddr})
    cikeFailEntry.EntityData.Leafs.Append("cikeFailRemoteAddr", types.YLeaf{"CikeFailRemoteAddr", cikeFailEntry.CikeFailRemoteAddr})

    cikeFailEntry.EntityData.YListKeys = []string {"CikeFailIndex"}

    return &(cikeFailEntry.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason represents 17 = operator requested termination.
type CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason string

const (
    CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason_other CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason = "other"

    CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason_peerDelRequest CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason = "peerDelRequest"

    CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason_peerLost CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason = "peerLost"

    CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason_localFailure CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason = "localFailure"

    CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason_authFailure CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason = "authFailure"

    CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason_hashValidation CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason = "hashValidation"

    CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason_encryptFailure CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason = "encryptFailure"

    CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason_internalError CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason = "internalError"

    CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason_sysCapExceeded CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason = "sysCapExceeded"

    CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason_proposalFailure CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason = "proposalFailure"

    CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason_peerCertUnavailable CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason = "peerCertUnavailable"

    CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason_peerCertNotValid CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason = "peerCertNotValid"

    CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason_localCertExpired CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason = "localCertExpired"

    CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason_crlFailure CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason = "crlFailure"

    CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason_peerEncodingError CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason = "peerEncodingError"

    CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason_nonExistentSa CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason = "nonExistentSa"

    CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason_operRequest CISCOIPSECFLOWMONITORMIB_CikeFailTable_CikeFailEntry_CikeFailReason = "operRequest"
)

// CISCOIPSECFLOWMONITORMIB_CipSecFailTable
// The IPsec Phase-2 Failure Table.
// This table is implemented as a sliding window 
// in which only the last n entries are maintained.  
// The maximum number of entries
// is specified by the cipSecFailTableSize object.
type CISCOIPSECFLOWMONITORMIB_CipSecFailTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with an IPsec Phase-1
    // failure. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry.
    CipSecFailEntry []*CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry
}

func (cipSecFailTable *CISCOIPSECFLOWMONITORMIB_CipSecFailTable) GetEntityData() *types.CommonEntityData {
    cipSecFailTable.EntityData.YFilter = cipSecFailTable.YFilter
    cipSecFailTable.EntityData.YangName = "cipSecFailTable"
    cipSecFailTable.EntityData.BundleName = "cisco_ios_xe"
    cipSecFailTable.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cipSecFailTable.EntityData.SegmentPath = "cipSecFailTable"
    cipSecFailTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipSecFailTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipSecFailTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipSecFailTable.EntityData.Children = types.NewOrderedMap()
    cipSecFailTable.EntityData.Children.Append("cipSecFailEntry", types.YChild{"CipSecFailEntry", nil})
    for i := range cipSecFailTable.CipSecFailEntry {
        cipSecFailTable.EntityData.Children.Append(types.GetSegmentPath(cipSecFailTable.CipSecFailEntry[i]), types.YChild{"CipSecFailEntry", cipSecFailTable.CipSecFailEntry[i]})
    }
    cipSecFailTable.EntityData.Leafs = types.NewOrderedMap()

    cipSecFailTable.EntityData.YListKeys = []string {}

    return &(cipSecFailTable.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry
// Each entry contains the attributes associated with
// an IPsec Phase-1 failure.
type CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IPsec Phase-2 Failure Table index. The value
    // of the index is a number which  begins at one and is incremented with each 
    // IPsec Phase-1 failure. The value of this object will wrap at 2,147,483,647.
    // The type is interface{} with range: 1..2147483647.
    CipSecFailIndex interface{}

    // The reason for the failure.  Possible reasons include:   1 = other   2 =
    // internal error occurred   3 = peer encoding error   4 = proposal failure  
    // 5 = protocol use failure   6 = non-existent security association   7 =
    // decryption failure   8 = encryption failure   9 = inbound authentication
    // failure  10 = outbound authentication failure  11 = compression failure  12
    // = system capacity failure  13 = peer delete request was received  14 =
    // contact with peer was lost  15 = sequence number rolled over  16 = operator
    // requested termination. The type is CipSecFailReason.
    CipSecFailReason interface{}

    // The value of sysUpTime in hundredths of seconds at the time of the failure.
    // The type is interface{} with range: 0..4294967295.
    CipSecFailTime interface{}

    // The Phase-2 Tunnel index (cipSecTunIndex). The type is interface{} with
    // range: 1..2147483647.
    CipSecFailTunnelIndex interface{}

    // The security association SPI value. The type is interface{} with range:
    // 0..2147483647.
    CipSecFailSaSpi interface{}

    // The packet's source IP address. The type is string with length: 4 | 16.
    CipSecFailPktSrcAddr interface{}

    // The packet's destination IP address. The type is string with length: 4 |
    // 16.
    CipSecFailPktDstAddr interface{}
}

func (cipSecFailEntry *CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry) GetEntityData() *types.CommonEntityData {
    cipSecFailEntry.EntityData.YFilter = cipSecFailEntry.YFilter
    cipSecFailEntry.EntityData.YangName = "cipSecFailEntry"
    cipSecFailEntry.EntityData.BundleName = "cisco_ios_xe"
    cipSecFailEntry.EntityData.ParentYangName = "cipSecFailTable"
    cipSecFailEntry.EntityData.SegmentPath = "cipSecFailEntry" + types.AddKeyToken(cipSecFailEntry.CipSecFailIndex, "cipSecFailIndex")
    cipSecFailEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipSecFailEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipSecFailEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipSecFailEntry.EntityData.Children = types.NewOrderedMap()
    cipSecFailEntry.EntityData.Leafs = types.NewOrderedMap()
    cipSecFailEntry.EntityData.Leafs.Append("cipSecFailIndex", types.YLeaf{"CipSecFailIndex", cipSecFailEntry.CipSecFailIndex})
    cipSecFailEntry.EntityData.Leafs.Append("cipSecFailReason", types.YLeaf{"CipSecFailReason", cipSecFailEntry.CipSecFailReason})
    cipSecFailEntry.EntityData.Leafs.Append("cipSecFailTime", types.YLeaf{"CipSecFailTime", cipSecFailEntry.CipSecFailTime})
    cipSecFailEntry.EntityData.Leafs.Append("cipSecFailTunnelIndex", types.YLeaf{"CipSecFailTunnelIndex", cipSecFailEntry.CipSecFailTunnelIndex})
    cipSecFailEntry.EntityData.Leafs.Append("cipSecFailSaSpi", types.YLeaf{"CipSecFailSaSpi", cipSecFailEntry.CipSecFailSaSpi})
    cipSecFailEntry.EntityData.Leafs.Append("cipSecFailPktSrcAddr", types.YLeaf{"CipSecFailPktSrcAddr", cipSecFailEntry.CipSecFailPktSrcAddr})
    cipSecFailEntry.EntityData.Leafs.Append("cipSecFailPktDstAddr", types.YLeaf{"CipSecFailPktDstAddr", cipSecFailEntry.CipSecFailPktDstAddr})

    cipSecFailEntry.EntityData.YListKeys = []string {"CipSecFailIndex"}

    return &(cipSecFailEntry.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason represents  16 = operator requested termination.
type CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason string

const (
    CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason_other CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason = "other"

    CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason_internalError CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason = "internalError"

    CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason_peerEncodingError CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason = "peerEncodingError"

    CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason_proposalFailure CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason = "proposalFailure"

    CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason_protocolUseFail CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason = "protocolUseFail"

    CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason_nonExistentSa CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason = "nonExistentSa"

    CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason_decryptFailure CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason = "decryptFailure"

    CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason_encryptFailure CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason = "encryptFailure"

    CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason_inAuthFailure CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason = "inAuthFailure"

    CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason_outAuthFailure CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason = "outAuthFailure"

    CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason_compression CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason = "compression"

    CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason_sysCapExceeded CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason = "sysCapExceeded"

    CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason_peerDelRequest CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason = "peerDelRequest"

    CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason_peerLost CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason = "peerLost"

    CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason_seqNumRollOver CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason = "seqNumRollOver"

    CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason_operRequest CISCOIPSECFLOWMONITORMIB_CipSecFailTable_CipSecFailEntry_CipSecFailReason = "operRequest"
)

