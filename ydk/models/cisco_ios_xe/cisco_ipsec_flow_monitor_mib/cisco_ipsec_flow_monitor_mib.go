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

// DiffHellmanGrp represents The Diffie Hellman Group used in negotiations.
type DiffHellmanGrp string

const (
    DiffHellmanGrp_none DiffHellmanGrp = "none"

    DiffHellmanGrp_dhGroup1 DiffHellmanGrp = "dhGroup1"

    DiffHellmanGrp_dhGroup2 DiffHellmanGrp = "dhGroup2"
)

// KeyType represents The type of key used by an IPsec Phase-2 Tunnel.
type KeyType string

const (
    KeyType_ike KeyType = "ike"

    KeyType_manual KeyType = "manual"
)

// EncapMode represents Tunnel.
type EncapMode string

const (
    EncapMode_tunnel EncapMode = "tunnel"

    EncapMode_transport EncapMode = "transport"
)

// EncryptAlgo represents The encryption algorithm used in negotiations.
type EncryptAlgo string

const (
    EncryptAlgo_none EncryptAlgo = "none"

    EncryptAlgo_des EncryptAlgo = "des"

    EncryptAlgo_des3 EncryptAlgo = "des3"
)

// AuthAlgo represents security association of an IPsec Phase-2 Tunnel.
type AuthAlgo string

const (
    AuthAlgo_none AuthAlgo = "none"

    AuthAlgo_hmacMd5 AuthAlgo = "hmacMd5"

    AuthAlgo_hmacSha AuthAlgo = "hmacSha"
)

// CompAlgo represents security association of an IPsec Phase-2 Tunnel.
type CompAlgo string

const (
    CompAlgo_none CompAlgo = "none"

    CompAlgo_ldf CompAlgo = "ldf"
)

// EndPtType represents The type of identity use to specify an IPsec End Point.
type EndPtType string

const (
    EndPtType_singleIpAddr EndPtType = "singleIpAddr"

    EndPtType_ipAddrRange EndPtType = "ipAddrRange"

    EndPtType_ipSubnet EndPtType = "ipSubnet"
)

// TunnelStatus represents type cannot be used to create a Tunnel.
type TunnelStatus string

const (
    TunnelStatus_active TunnelStatus = "active"

    TunnelStatus_destroy TunnelStatus = "destroy"
)

// TrapStatus represents The administrative status for sending a TRAP.
type TrapStatus string

const (
    TrapStatus_enabled TrapStatus = "enabled"

    TrapStatus_disabled TrapStatus = "disabled"
)

// CISCOIPSECFLOWMONITORMIB
type CISCOIPSECFLOWMONITORMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Cipseclevels CISCOIPSECFLOWMONITORMIB_Cipseclevels

    
    Cikeglobalstats CISCOIPSECFLOWMONITORMIB_Cikeglobalstats

    
    Cipsecglobalstats CISCOIPSECFLOWMONITORMIB_Cipsecglobalstats

    
    Cipsechistglobalcntl CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl

    
    Cipsecfailglobalcntl CISCOIPSECFLOWMONITORMIB_Cipsecfailglobalcntl

    
    Cipsectrapcntl CISCOIPSECFLOWMONITORMIB_Cipsectrapcntl

    // The IPsec Phase-1 Internet Key Exchange Peer Table. There is one entry in
    // this table for each IPsec Phase-1 IKE peer association which is currently
    // associated with an active IPsec Phase-1 Tunnel. The IPsec Phase-1 IKE
    // Tunnel associated with this IPsec Phase-1 IKE peer association may or may
    // not be currently active.
    Cikepeertable CISCOIPSECFLOWMONITORMIB_Cikepeertable

    // The IPsec Phase-1 Internet Key Exchange Tunnel Table. There is one entry in
    // this table for each active IPsec Phase-1 IKE Tunnel.
    Ciketunneltable CISCOIPSECFLOWMONITORMIB_Ciketunneltable

    // The IPsec Phase-1 Internet Key Exchange Peer Association to IPsec Phase-2
    // Tunnel Correlation Table. There is one entry in this table for each active
    // IPsec Phase-2 Tunnel.
    Cikepeercorrtable CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable

    // Phase-1 IKE stats information is included in this table. Each entry is
    // related to a specific gateway which is  identified by 'cmgwIndex'.
    Cikephase1Gwstatstable CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable

    // The IPsec Phase-2 Tunnel Table. There is one entry in this table for  each
    // active IPsec Phase-2 Tunnel.
    Cipsectunneltable CISCOIPSECFLOWMONITORMIB_Cipsectunneltable

    // The IPsec Phase-2 Tunnel Endpoint Table. This table contains an entry for
    // each  active endpoint associated with an IPsec  Phase-2 Tunnel.
    Cipsecendpttable CISCOIPSECFLOWMONITORMIB_Cipsecendpttable

    // The IPsec Phase-2 Security Protection Index Table. This table contains an
    // entry for each active  and expiring security  association.
    Cipsecspitable CISCOIPSECFLOWMONITORMIB_Cipsecspitable

    // Phase-2 IPsec stats information is included in this table. Each entry is
    // related to a specific gateway which is  identified by 'cmgwIndex'.
    Cipsecphase2Gwstatstable CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable

    // The IPsec Phase-1 Internet Key Exchange Tunnel History Table.  This table
    // is implemented as a  sliding window in which only the last n entries  are
    // maintained.  The maximum number of entries  is specified by the
    // cipSecHistTableSize object.
    Ciketunnelhisttable CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable

    // The IPsec Phase-2 Tunnel History Table. This table is implemented as a
    // sliding  window in which only the last n entries are maintained.  The
    // maximum number  of entries is specified by the cipSecHistTableSize object.
    Cipsectunnelhisttable CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable

    // The IPsec Phase-2 Tunnel Endpoint History Table. This table is implemented
    // as a  sliding window in which only the last n entries are maintained.   The
    // maximum number of entries is specified by the cipSecHistTableSize object.
    Cipsecendpthisttable CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable

    // The IPsec Phase-1 Failure Table. This table is implemented as a sliding 
    // window in which only the last n entries are  maintained.  The maximum
    // number of entries is specified by the cipSecFailTableSize object.
    Cikefailtable CISCOIPSECFLOWMONITORMIB_Cikefailtable

    // The IPsec Phase-2 Failure Table. This table is implemented as a sliding
    // window  in which only the last n entries are maintained.   The maximum
    // number of entries is specified by the cipSecFailTableSize object.
    Cipsecfailtable CISCOIPSECFLOWMONITORMIB_Cipsecfailtable
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

    cISCOIPSECFLOWMONITORMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOIPSECFLOWMONITORMIB.EntityData.Children["cipSecLevels"] = types.YChild{"Cipseclevels", &cISCOIPSECFLOWMONITORMIB.Cipseclevels}
    cISCOIPSECFLOWMONITORMIB.EntityData.Children["cikeGlobalStats"] = types.YChild{"Cikeglobalstats", &cISCOIPSECFLOWMONITORMIB.Cikeglobalstats}
    cISCOIPSECFLOWMONITORMIB.EntityData.Children["cipSecGlobalStats"] = types.YChild{"Cipsecglobalstats", &cISCOIPSECFLOWMONITORMIB.Cipsecglobalstats}
    cISCOIPSECFLOWMONITORMIB.EntityData.Children["cipSecHistGlobalCntl"] = types.YChild{"Cipsechistglobalcntl", &cISCOIPSECFLOWMONITORMIB.Cipsechistglobalcntl}
    cISCOIPSECFLOWMONITORMIB.EntityData.Children["cipSecFailGlobalCntl"] = types.YChild{"Cipsecfailglobalcntl", &cISCOIPSECFLOWMONITORMIB.Cipsecfailglobalcntl}
    cISCOIPSECFLOWMONITORMIB.EntityData.Children["cipSecTrapCntl"] = types.YChild{"Cipsectrapcntl", &cISCOIPSECFLOWMONITORMIB.Cipsectrapcntl}
    cISCOIPSECFLOWMONITORMIB.EntityData.Children["cikePeerTable"] = types.YChild{"Cikepeertable", &cISCOIPSECFLOWMONITORMIB.Cikepeertable}
    cISCOIPSECFLOWMONITORMIB.EntityData.Children["cikeTunnelTable"] = types.YChild{"Ciketunneltable", &cISCOIPSECFLOWMONITORMIB.Ciketunneltable}
    cISCOIPSECFLOWMONITORMIB.EntityData.Children["cikePeerCorrTable"] = types.YChild{"Cikepeercorrtable", &cISCOIPSECFLOWMONITORMIB.Cikepeercorrtable}
    cISCOIPSECFLOWMONITORMIB.EntityData.Children["cikePhase1GWStatsTable"] = types.YChild{"Cikephase1Gwstatstable", &cISCOIPSECFLOWMONITORMIB.Cikephase1Gwstatstable}
    cISCOIPSECFLOWMONITORMIB.EntityData.Children["cipSecTunnelTable"] = types.YChild{"Cipsectunneltable", &cISCOIPSECFLOWMONITORMIB.Cipsectunneltable}
    cISCOIPSECFLOWMONITORMIB.EntityData.Children["cipSecEndPtTable"] = types.YChild{"Cipsecendpttable", &cISCOIPSECFLOWMONITORMIB.Cipsecendpttable}
    cISCOIPSECFLOWMONITORMIB.EntityData.Children["cipSecSpiTable"] = types.YChild{"Cipsecspitable", &cISCOIPSECFLOWMONITORMIB.Cipsecspitable}
    cISCOIPSECFLOWMONITORMIB.EntityData.Children["cipSecPhase2GWStatsTable"] = types.YChild{"Cipsecphase2Gwstatstable", &cISCOIPSECFLOWMONITORMIB.Cipsecphase2Gwstatstable}
    cISCOIPSECFLOWMONITORMIB.EntityData.Children["cikeTunnelHistTable"] = types.YChild{"Ciketunnelhisttable", &cISCOIPSECFLOWMONITORMIB.Ciketunnelhisttable}
    cISCOIPSECFLOWMONITORMIB.EntityData.Children["cipSecTunnelHistTable"] = types.YChild{"Cipsectunnelhisttable", &cISCOIPSECFLOWMONITORMIB.Cipsectunnelhisttable}
    cISCOIPSECFLOWMONITORMIB.EntityData.Children["cipSecEndPtHistTable"] = types.YChild{"Cipsecendpthisttable", &cISCOIPSECFLOWMONITORMIB.Cipsecendpthisttable}
    cISCOIPSECFLOWMONITORMIB.EntityData.Children["cikeFailTable"] = types.YChild{"Cikefailtable", &cISCOIPSECFLOWMONITORMIB.Cikefailtable}
    cISCOIPSECFLOWMONITORMIB.EntityData.Children["cipSecFailTable"] = types.YChild{"Cipsecfailtable", &cISCOIPSECFLOWMONITORMIB.Cipsecfailtable}
    cISCOIPSECFLOWMONITORMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOIPSECFLOWMONITORMIB.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cipseclevels
type CISCOIPSECFLOWMONITORMIB_Cipseclevels struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The level of the IPsec MIB. The type is interface{} with range: 1..4096.
    Cipsecmiblevel interface{}
}

func (cipseclevels *CISCOIPSECFLOWMONITORMIB_Cipseclevels) GetEntityData() *types.CommonEntityData {
    cipseclevels.EntityData.YFilter = cipseclevels.YFilter
    cipseclevels.EntityData.YangName = "cipSecLevels"
    cipseclevels.EntityData.BundleName = "cisco_ios_xe"
    cipseclevels.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cipseclevels.EntityData.SegmentPath = "cipSecLevels"
    cipseclevels.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipseclevels.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipseclevels.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipseclevels.EntityData.Children = make(map[string]types.YChild)
    cipseclevels.EntityData.Leafs = make(map[string]types.YLeaf)
    cipseclevels.EntityData.Leafs["cipSecMibLevel"] = types.YLeaf{"Cipsecmiblevel", cipseclevels.Cipsecmiblevel}
    return &(cipseclevels.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cikeglobalstats
type CISCOIPSECFLOWMONITORMIB_Cikeglobalstats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of currently active IPsec Phase-1 IKE Tunnels. The type is
    // interface{} with range: 0..4294967295.
    Cikeglobalactivetunnels interface{}

    // The total number of previously active IPsec Phase-1 IKE Tunnels. The type
    // is interface{} with range: 0..4294967295. Units are SAs.
    Cikeglobalprevioustunnels interface{}

    // The total number of octets received by all currently and previously active
    // IPsec Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Octets.
    Cikeglobalinoctets interface{}

    // The total number of packets received by all currently and previously active
    // IPsec  Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    Cikeglobalinpkts interface{}

    // The total number of packets which were dropped during receive processing by
    // all  currently and previously  active IPsec Phase-1 IKE Tunnels. The type
    // is interface{} with range: 0..4294967295. Units are Packets.
    Cikeglobalindroppkts interface{}

    // The total number of notifys received by all currently and previously active
    // IPsec  Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Notification Payloads.
    Cikeglobalinnotifys interface{}

    // The total number of IPsec Phase-2 exchanges received by all currently and
    // previously  active IPsec Phase-1 IKE Tunnels. The type is interface{} with
    // range: 0..4294967295. Units are SA Payloads.
    Cikeglobalinp2Exchgs interface{}

    // The total number of IPsec Phase-2 exchanges which were received and found
    // to be invalid  by all currently and previously active IPsec  Phase-1 IKE
    // Tunnels. The type is interface{} with range: 0..4294967295. Units are SA
    // Payloads.
    Cikeglobalinp2Exchginvalids interface{}

    // The total number of IPsec Phase-2 exchanges which were received and
    // rejected by all  currently and previously active IPsec Phase-1  IKE
    // Tunnels. The type is interface{} with range: 0..4294967295. Units are SA
    // Payloads.
    Cikeglobalinp2Exchgrejects interface{}

    // The total number of IPsec Phase-2 security association delete requests
    // received by all  currently and previously  active and IPsec Phase-1 IKE
    // Tunnels. The type is interface{} with range: 0..4294967295. Units are
    // Notification Payloads.
    Cikeglobalinp2Sadelrequests interface{}

    // The total number of octets sent by all currently and previously active and
    // IPsec Phase-1  IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Octets.
    Cikeglobaloutoctets interface{}

    // The total number of packets sent by all currently and previously active and
    // IPsec Phase-1  Tunnels. The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    Cikeglobaloutpkts interface{}

    // The total number of packets which were dropped during send processing by
    // all currently  and previously  active IPsec Phase-1 IKE Tunnels. The type
    // is interface{} with range: 0..4294967295. Units are Packets.
    Cikeglobaloutdroppkts interface{}

    // The total number of notifys sent by all currently and previously active
    // IPsec Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Notification Payloads.
    Cikeglobaloutnotifys interface{}

    // The total number of IPsec Phase-2 exchanges which were sent by all
    // currently and previously  active IPsec Phase-1 IKE Tunnels. The type is
    // interface{} with range: 0..4294967295. Units are SA Payloads.
    Cikeglobaloutp2Exchgs interface{}

    // The total number of IPsec Phase-2 exchanges which were sent and found to be
    // invalid by  all currently and previously active IPsec Phase-1  Tunnels. The
    // type is interface{} with range: 0..4294967295. Units are SA Payloads.
    Cikeglobaloutp2Exchginvalids interface{}

    // The total number of IPsec Phase-2 exchanges which were sent and rejected by
    // all currently and  previously active IPsec Phase-1 IKE Tunnels. The type is
    // interface{} with range: 0..4294967295. Units are SA Payloads.
    Cikeglobaloutp2Exchgrejects interface{}

    // The total number of IPsec Phase-2 SA delete requests sent by all currently
    // and  previously active IPsec Phase-1 IKE Tunnels. The type is interface{}
    // with range: 0..4294967295. Units are Notification Payloads.
    Cikeglobaloutp2Sadelrequests interface{}

    // The total number of IPsec Phase-1 IKE Tunnels which were locally initiated.
    // The type is interface{} with range: 0..4294967295. Units are SAs.
    Cikeglobalinittunnels interface{}

    // The total number of IPsec Phase-1 IKE Tunnels which were locally initiated
    // and failed to activate. The type is interface{} with range: 0..4294967295.
    // Units are SAs.
    Cikeglobalinittunnelfails interface{}

    // The total number of IPsec Phase-1 IKE Tunnels which were remotely initiated
    // and failed to activate. The type is interface{} with range: 0..4294967295.
    // Units are SAs.
    Cikeglobalresptunnelfails interface{}

    // The total number of system capacity failures which occurred during
    // processing of all current  and previously active IPsec Phase-1 IKE Tunnels.
    // The type is interface{} with range: 0..4294967295. Units are Failures.
    Cikeglobalsyscapfails interface{}

    // The total number of authentications which ended in failure by all current
    // and previous IPsec Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Failures.
    Cikeglobalauthfails interface{}

    // The total number of decryptions which ended in failure by all current and
    // previous IPsec Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Failures.
    Cikeglobaldecryptfails interface{}

    // The total number of hash validations which ended in failure by all current
    // and previous IPsec Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Failures.
    Cikeglobalhashvalidfails interface{}

    // The total number of non-existent Security Association in failures which
    // occurred during processing of  all current and previous IPsec Phase-1 IKE
    // Tunnels. The type is interface{} with range: 0..4294967295. Units are
    // Failures.
    Cikeglobalnosafails interface{}
}

func (cikeglobalstats *CISCOIPSECFLOWMONITORMIB_Cikeglobalstats) GetEntityData() *types.CommonEntityData {
    cikeglobalstats.EntityData.YFilter = cikeglobalstats.YFilter
    cikeglobalstats.EntityData.YangName = "cikeGlobalStats"
    cikeglobalstats.EntityData.BundleName = "cisco_ios_xe"
    cikeglobalstats.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cikeglobalstats.EntityData.SegmentPath = "cikeGlobalStats"
    cikeglobalstats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cikeglobalstats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cikeglobalstats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cikeglobalstats.EntityData.Children = make(map[string]types.YChild)
    cikeglobalstats.EntityData.Leafs = make(map[string]types.YLeaf)
    cikeglobalstats.EntityData.Leafs["cikeGlobalActiveTunnels"] = types.YLeaf{"Cikeglobalactivetunnels", cikeglobalstats.Cikeglobalactivetunnels}
    cikeglobalstats.EntityData.Leafs["cikeGlobalPreviousTunnels"] = types.YLeaf{"Cikeglobalprevioustunnels", cikeglobalstats.Cikeglobalprevioustunnels}
    cikeglobalstats.EntityData.Leafs["cikeGlobalInOctets"] = types.YLeaf{"Cikeglobalinoctets", cikeglobalstats.Cikeglobalinoctets}
    cikeglobalstats.EntityData.Leafs["cikeGlobalInPkts"] = types.YLeaf{"Cikeglobalinpkts", cikeglobalstats.Cikeglobalinpkts}
    cikeglobalstats.EntityData.Leafs["cikeGlobalInDropPkts"] = types.YLeaf{"Cikeglobalindroppkts", cikeglobalstats.Cikeglobalindroppkts}
    cikeglobalstats.EntityData.Leafs["cikeGlobalInNotifys"] = types.YLeaf{"Cikeglobalinnotifys", cikeglobalstats.Cikeglobalinnotifys}
    cikeglobalstats.EntityData.Leafs["cikeGlobalInP2Exchgs"] = types.YLeaf{"Cikeglobalinp2Exchgs", cikeglobalstats.Cikeglobalinp2Exchgs}
    cikeglobalstats.EntityData.Leafs["cikeGlobalInP2ExchgInvalids"] = types.YLeaf{"Cikeglobalinp2Exchginvalids", cikeglobalstats.Cikeglobalinp2Exchginvalids}
    cikeglobalstats.EntityData.Leafs["cikeGlobalInP2ExchgRejects"] = types.YLeaf{"Cikeglobalinp2Exchgrejects", cikeglobalstats.Cikeglobalinp2Exchgrejects}
    cikeglobalstats.EntityData.Leafs["cikeGlobalInP2SaDelRequests"] = types.YLeaf{"Cikeglobalinp2Sadelrequests", cikeglobalstats.Cikeglobalinp2Sadelrequests}
    cikeglobalstats.EntityData.Leafs["cikeGlobalOutOctets"] = types.YLeaf{"Cikeglobaloutoctets", cikeglobalstats.Cikeglobaloutoctets}
    cikeglobalstats.EntityData.Leafs["cikeGlobalOutPkts"] = types.YLeaf{"Cikeglobaloutpkts", cikeglobalstats.Cikeglobaloutpkts}
    cikeglobalstats.EntityData.Leafs["cikeGlobalOutDropPkts"] = types.YLeaf{"Cikeglobaloutdroppkts", cikeglobalstats.Cikeglobaloutdroppkts}
    cikeglobalstats.EntityData.Leafs["cikeGlobalOutNotifys"] = types.YLeaf{"Cikeglobaloutnotifys", cikeglobalstats.Cikeglobaloutnotifys}
    cikeglobalstats.EntityData.Leafs["cikeGlobalOutP2Exchgs"] = types.YLeaf{"Cikeglobaloutp2Exchgs", cikeglobalstats.Cikeglobaloutp2Exchgs}
    cikeglobalstats.EntityData.Leafs["cikeGlobalOutP2ExchgInvalids"] = types.YLeaf{"Cikeglobaloutp2Exchginvalids", cikeglobalstats.Cikeglobaloutp2Exchginvalids}
    cikeglobalstats.EntityData.Leafs["cikeGlobalOutP2ExchgRejects"] = types.YLeaf{"Cikeglobaloutp2Exchgrejects", cikeglobalstats.Cikeglobaloutp2Exchgrejects}
    cikeglobalstats.EntityData.Leafs["cikeGlobalOutP2SaDelRequests"] = types.YLeaf{"Cikeglobaloutp2Sadelrequests", cikeglobalstats.Cikeglobaloutp2Sadelrequests}
    cikeglobalstats.EntityData.Leafs["cikeGlobalInitTunnels"] = types.YLeaf{"Cikeglobalinittunnels", cikeglobalstats.Cikeglobalinittunnels}
    cikeglobalstats.EntityData.Leafs["cikeGlobalInitTunnelFails"] = types.YLeaf{"Cikeglobalinittunnelfails", cikeglobalstats.Cikeglobalinittunnelfails}
    cikeglobalstats.EntityData.Leafs["cikeGlobalRespTunnelFails"] = types.YLeaf{"Cikeglobalresptunnelfails", cikeglobalstats.Cikeglobalresptunnelfails}
    cikeglobalstats.EntityData.Leafs["cikeGlobalSysCapFails"] = types.YLeaf{"Cikeglobalsyscapfails", cikeglobalstats.Cikeglobalsyscapfails}
    cikeglobalstats.EntityData.Leafs["cikeGlobalAuthFails"] = types.YLeaf{"Cikeglobalauthfails", cikeglobalstats.Cikeglobalauthfails}
    cikeglobalstats.EntityData.Leafs["cikeGlobalDecryptFails"] = types.YLeaf{"Cikeglobaldecryptfails", cikeglobalstats.Cikeglobaldecryptfails}
    cikeglobalstats.EntityData.Leafs["cikeGlobalHashValidFails"] = types.YLeaf{"Cikeglobalhashvalidfails", cikeglobalstats.Cikeglobalhashvalidfails}
    cikeglobalstats.EntityData.Leafs["cikeGlobalNoSaFails"] = types.YLeaf{"Cikeglobalnosafails", cikeglobalstats.Cikeglobalnosafails}
    return &(cikeglobalstats.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cipsecglobalstats
type CISCOIPSECFLOWMONITORMIB_Cipsecglobalstats struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The total number of currently active IPsec Phase-2 Tunnels. The type is
    // interface{} with range: 0..4294967295.
    Cipsecglobalactivetunnels interface{}

    // The total number of previously active IPsec Phase-2 Tunnels. The type is
    // interface{} with range: 0..4294967295. Units are Phase-2 Tunnels.
    Cipsecglobalprevioustunnels interface{}

    // The total number of octets received by all current and previous IPsec
    // Phase-2 Tunnels.  This value is accumulated BEFORE determining whether or
    // not the packet should be decompressed. See also cipSecGlobalInOctWraps for
    // the number of times this counter has wrapped. The type is interface{} with
    // range: 0..4294967295. Units are Octets.
    Cipsecglobalinoctets interface{}

    // A high capacity count of the total number of octets received by all current
    // and previous IPsec Phase-2 Tunnels. This value is accumulated BEFORE
    // determining whether or not the packet should be decompressed. The type is
    // interface{} with range: 0..18446744073709551615.
    Cipsecglobalhcinoctets interface{}

    // The number of times the global octets received counter
    // (cipSecGlobalInOctets) has wrapped. The type is interface{} with range:
    // 0..4294967295. Units are Integral units.
    Cipsecglobalinoctwraps interface{}

    // The total number of decompressed octets received by all current and
    // previous IPsec Phase-2 Tunnels.   This value is accumulated AFTER the
    // packet is  decompressed. If compression is not being used,  this value will
    // match the value of cipSecGlobalInOctets.  See also
    // cipSecGlobalInDecompOctWraps  for the number of times this counter has
    // wrapped. The type is interface{} with range: 0..4294967295. Units are
    // Octets.
    Cipsecglobalindecompoctets interface{}

    // A high capacity count of the total number of decompressed octets received
    // by all current  and previous IPsec Phase-2 Tunnels.  This value  is
    // accumulated AFTER the packet is decompressed.  If compression is not being
    // used, this value   will match the value of cipSecGlobalHcInOctets. The type
    // is interface{} with range: 0..18446744073709551615.
    Cipsecglobalhcindecompoctets interface{}

    // The number of times the global decompressed octets received counter 
    // (cipSecGlobalInDecompOctets) has wrapped. The type is interface{} with
    // range: 0..4294967295. Units are Integral units.
    Cipsecglobalindecompoctwraps interface{}

    // The total number of packets received by all current and previous  IPsec
    // Phase-2 Tunnels. The type is interface{} with range: 0..4294967295. Units
    // are Packets.
    Cipsecglobalinpkts interface{}

    // The total number of packets dropped during receive processing by all
    // current and previous  IPsec Phase-2 Tunnels. This count does NOT include
    // packets dropped due to  Anti-Replay processing. The type is interface{}
    // with range: 0..4294967295. Units are Packets.
    Cipsecglobalindrops interface{}

    // The total number of packets dropped during receive processing due to
    // Anti-Replay  processing by all current and previous IPsec  Phase-2 Tunnels.
    // The type is interface{} with range: 0..4294967295. Units are Packets.
    Cipsecglobalinreplaydrops interface{}

    // The total number of inbound authentication's performed by all current and
    // previous IPsec  Phase-2 Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Events.
    Cipsecglobalinauths interface{}

    // The total number of inbound authentication's which ended in failure by all
    // current and previous  IPsec Phase-2 Tunnels. The type is interface{} with
    // range: 0..4294967295. Units are Failures.
    Cipsecglobalinauthfails interface{}

    // The total number of inbound decryption's performed by all current and
    // previous IPsec  Phase-2 Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    Cipsecglobalindecrypts interface{}

    // The total number of inbound decryption's which ended in failure by all
    // current and  previous IPsec Phase-2 Tunnels. The type is interface{} with
    // range: 0..4294967295. Units are Packets.
    Cipsecglobalindecryptfails interface{}

    // The total number of octets sent by all current and previous IPsec Phase-2
    // Tunnels.   This value is accumulated AFTER determining  whether or not the
    // packet should be compressed.   See also cipSecGlobalOutOctWraps for the 
    // number of times this counter has wrapped. The type is interface{} with
    // range: 0..4294967295. Units are Octets.
    Cipsecglobaloutoctets interface{}

    // A high capacity count of the total number of octets sent by all current and
    // previous  IPsec Phase-2 Tunnels.  This value is accumulated  AFTER
    // determining whether or not the packet should  be compressed. The type is
    // interface{} with range: 0..18446744073709551615.
    Cipsecglobalhcoutoctets interface{}

    // The number of times the global octets sent counter (cipSecGlobalOutOctets)
    // has wrapped. The type is interface{} with range: 0..4294967295. Units are
    // Integral units.
    Cipsecglobaloutoctwraps interface{}

    // The total number of uncompressed octets sent by all current and previous
    // IPsec Phase-2 Tunnels.   This value is accumulated BEFORE the packet is 
    // compressed. If compression is not being used, this  value will match the
    // value of cipSecGlobalOutOctets.  See also cipSecGlobalOutDecompOctWraps for
    // the number  of times this counter has wrapped. The type is interface{} with
    // range: 0..4294967295. Units are Octets.
    Cipsecglobaloutuncompoctets interface{}

    // A high capacity count of the total number of uncompressed octets sent by
    // all current and previous  IPsec Phase-2 Tunnels.  This value is accumulated
    // BEFORE the packet is compressed.  If compression is  not being used, this
    // value will match the       value of cipSecGlobalHcOutOctets. The type is
    // interface{} with range: 0..18446744073709551615. Units are Octets.
    Cipsecglobalhcoutuncompoctets interface{}

    // The number of times the global uncompressed octets sent counter
    // (cipSecGlobalOutUncompOctets)  has wrapped. The type is interface{} with
    // range: 0..4294967295. Units are Integral units.
    Cipsecglobaloutuncompoctwraps interface{}

    // The total number of packets sent by all current and previous  IPsec Phase-2
    // Tunnels. The type is interface{} with range: 0..4294967295. Units are
    // Packets.
    Cipsecglobaloutpkts interface{}

    // The total number of packets dropped during send processing by all current
    // and previous IPsec  Phase-2 Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    Cipsecglobaloutdrops interface{}

    // The total number of outbound authentication's performed by all current and
    // previous IPsec  Phase-2 Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Events.
    Cipsecglobaloutauths interface{}

    // The total number of outbound authentication's which ended in failure  by
    // all current and previous IPsec Phase-2 Tunnels. The type is interface{}
    // with range: 0..4294967295. Units are Failures.
    Cipsecglobaloutauthfails interface{}

    // The total number of outbound encryption's performed by all current and
    // previous IPsec Phase-2 Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    Cipsecglobaloutencrypts interface{}

    // The total number of outbound encryption's which ended in failure by all
    // current and  previous IPsec Phase-2 Tunnels. The type is interface{} with
    // range: 0..4294967295. Units are Failures.
    Cipsecglobaloutencryptfails interface{}

    // The total number of protocol use failures which occurred during processing
    // of all current  and previously active IPsec Phase-2 Tunnels. The type is
    // interface{} with range: 0..4294967295. Units are Failures.
    Cipsecglobalprotocolusefails interface{}

    // The total number of non-existent Security Association in failures which
    // occurred  during processing of all current  and previous IPsec Phase-2
    // Tunnels. The type is interface{} with range: 0..4294967295. Units are
    // Failures.
    Cipsecglobalnosafails interface{}

    // The total number of system capacity failures which occurred during
    // processing of all current  and previously active IPsec Phase-2 Tunnels. The
    // type is interface{} with range: 0..4294967295. Units are Failures.
    Cipsecglobalsyscapfails interface{}
}

func (cipsecglobalstats *CISCOIPSECFLOWMONITORMIB_Cipsecglobalstats) GetEntityData() *types.CommonEntityData {
    cipsecglobalstats.EntityData.YFilter = cipsecglobalstats.YFilter
    cipsecglobalstats.EntityData.YangName = "cipSecGlobalStats"
    cipsecglobalstats.EntityData.BundleName = "cisco_ios_xe"
    cipsecglobalstats.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cipsecglobalstats.EntityData.SegmentPath = "cipSecGlobalStats"
    cipsecglobalstats.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsecglobalstats.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsecglobalstats.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsecglobalstats.EntityData.Children = make(map[string]types.YChild)
    cipsecglobalstats.EntityData.Leafs = make(map[string]types.YLeaf)
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalActiveTunnels"] = types.YLeaf{"Cipsecglobalactivetunnels", cipsecglobalstats.Cipsecglobalactivetunnels}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalPreviousTunnels"] = types.YLeaf{"Cipsecglobalprevioustunnels", cipsecglobalstats.Cipsecglobalprevioustunnels}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalInOctets"] = types.YLeaf{"Cipsecglobalinoctets", cipsecglobalstats.Cipsecglobalinoctets}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalHcInOctets"] = types.YLeaf{"Cipsecglobalhcinoctets", cipsecglobalstats.Cipsecglobalhcinoctets}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalInOctWraps"] = types.YLeaf{"Cipsecglobalinoctwraps", cipsecglobalstats.Cipsecglobalinoctwraps}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalInDecompOctets"] = types.YLeaf{"Cipsecglobalindecompoctets", cipsecglobalstats.Cipsecglobalindecompoctets}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalHcInDecompOctets"] = types.YLeaf{"Cipsecglobalhcindecompoctets", cipsecglobalstats.Cipsecglobalhcindecompoctets}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalInDecompOctWraps"] = types.YLeaf{"Cipsecglobalindecompoctwraps", cipsecglobalstats.Cipsecglobalindecompoctwraps}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalInPkts"] = types.YLeaf{"Cipsecglobalinpkts", cipsecglobalstats.Cipsecglobalinpkts}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalInDrops"] = types.YLeaf{"Cipsecglobalindrops", cipsecglobalstats.Cipsecglobalindrops}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalInReplayDrops"] = types.YLeaf{"Cipsecglobalinreplaydrops", cipsecglobalstats.Cipsecglobalinreplaydrops}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalInAuths"] = types.YLeaf{"Cipsecglobalinauths", cipsecglobalstats.Cipsecglobalinauths}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalInAuthFails"] = types.YLeaf{"Cipsecglobalinauthfails", cipsecglobalstats.Cipsecglobalinauthfails}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalInDecrypts"] = types.YLeaf{"Cipsecglobalindecrypts", cipsecglobalstats.Cipsecglobalindecrypts}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalInDecryptFails"] = types.YLeaf{"Cipsecglobalindecryptfails", cipsecglobalstats.Cipsecglobalindecryptfails}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalOutOctets"] = types.YLeaf{"Cipsecglobaloutoctets", cipsecglobalstats.Cipsecglobaloutoctets}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalHcOutOctets"] = types.YLeaf{"Cipsecglobalhcoutoctets", cipsecglobalstats.Cipsecglobalhcoutoctets}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalOutOctWraps"] = types.YLeaf{"Cipsecglobaloutoctwraps", cipsecglobalstats.Cipsecglobaloutoctwraps}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalOutUncompOctets"] = types.YLeaf{"Cipsecglobaloutuncompoctets", cipsecglobalstats.Cipsecglobaloutuncompoctets}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalHcOutUncompOctets"] = types.YLeaf{"Cipsecglobalhcoutuncompoctets", cipsecglobalstats.Cipsecglobalhcoutuncompoctets}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalOutUncompOctWraps"] = types.YLeaf{"Cipsecglobaloutuncompoctwraps", cipsecglobalstats.Cipsecglobaloutuncompoctwraps}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalOutPkts"] = types.YLeaf{"Cipsecglobaloutpkts", cipsecglobalstats.Cipsecglobaloutpkts}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalOutDrops"] = types.YLeaf{"Cipsecglobaloutdrops", cipsecglobalstats.Cipsecglobaloutdrops}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalOutAuths"] = types.YLeaf{"Cipsecglobaloutauths", cipsecglobalstats.Cipsecglobaloutauths}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalOutAuthFails"] = types.YLeaf{"Cipsecglobaloutauthfails", cipsecglobalstats.Cipsecglobaloutauthfails}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalOutEncrypts"] = types.YLeaf{"Cipsecglobaloutencrypts", cipsecglobalstats.Cipsecglobaloutencrypts}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalOutEncryptFails"] = types.YLeaf{"Cipsecglobaloutencryptfails", cipsecglobalstats.Cipsecglobaloutencryptfails}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalProtocolUseFails"] = types.YLeaf{"Cipsecglobalprotocolusefails", cipsecglobalstats.Cipsecglobalprotocolusefails}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalNoSaFails"] = types.YLeaf{"Cipsecglobalnosafails", cipsecglobalstats.Cipsecglobalnosafails}
    cipsecglobalstats.EntityData.Leafs["cipSecGlobalSysCapFails"] = types.YLeaf{"Cipsecglobalsyscapfails", cipsecglobalstats.Cipsecglobalsyscapfails}
    return &(cipsecglobalstats.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl
type CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl struct {
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
    Cipsechisttablesize interface{}

    // The current state of check point processing.  This object will return ready
    // when the agent is  ready to create on-demand history entries for  active
    // IPsec Tunnels or checkPoint when the  agent is currently creating on-demand
    // history  entries for active IPsec Tunnels.  By setting this value to
    // checkPoint, the agent  will create: a) an entry in the IPsec Phase-1 Tunnel
    // History     for each active IPsec Phase-1 Tunnel and b) an entry in the
    // IPsec Phase-2 Tunnel History     Table and an entry in the IPsec Phase-2   
    // Tunnel EndPoint History Table    for each active IPsec Phase-2 Tunnel. The
    // type is Cipsechistcheckpoint.
    Cipsechistcheckpoint interface{}
}

func (cipsechistglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl) GetEntityData() *types.CommonEntityData {
    cipsechistglobalcntl.EntityData.YFilter = cipsechistglobalcntl.YFilter
    cipsechistglobalcntl.EntityData.YangName = "cipSecHistGlobalCntl"
    cipsechistglobalcntl.EntityData.BundleName = "cisco_ios_xe"
    cipsechistglobalcntl.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cipsechistglobalcntl.EntityData.SegmentPath = "cipSecHistGlobalCntl"
    cipsechistglobalcntl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsechistglobalcntl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsechistglobalcntl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsechistglobalcntl.EntityData.Children = make(map[string]types.YChild)
    cipsechistglobalcntl.EntityData.Leafs = make(map[string]types.YLeaf)
    cipsechistglobalcntl.EntityData.Leafs["cipSecHistTableSize"] = types.YLeaf{"Cipsechisttablesize", cipsechistglobalcntl.Cipsechisttablesize}
    cipsechistglobalcntl.EntityData.Leafs["cipSecHistCheckPoint"] = types.YLeaf{"Cipsechistcheckpoint", cipsechistglobalcntl.Cipsechistcheckpoint}
    return &(cipsechistglobalcntl.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl_Cipsechistcheckpoint represents    for each active IPsec Phase-2 Tunnel.
type CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl_Cipsechistcheckpoint string

const (
    CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl_Cipsechistcheckpoint_ready CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl_Cipsechistcheckpoint = "ready"

    CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl_Cipsechistcheckpoint_checkPoint CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl_Cipsechistcheckpoint = "checkPoint"
)

// CISCOIPSECFLOWMONITORMIB_Cipsecfailglobalcntl
type CISCOIPSECFLOWMONITORMIB_Cipsecfailglobalcntl struct {
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
    Cipsecfailtablesize interface{}
}

func (cipsecfailglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsecfailglobalcntl) GetEntityData() *types.CommonEntityData {
    cipsecfailglobalcntl.EntityData.YFilter = cipsecfailglobalcntl.YFilter
    cipsecfailglobalcntl.EntityData.YangName = "cipSecFailGlobalCntl"
    cipsecfailglobalcntl.EntityData.BundleName = "cisco_ios_xe"
    cipsecfailglobalcntl.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cipsecfailglobalcntl.EntityData.SegmentPath = "cipSecFailGlobalCntl"
    cipsecfailglobalcntl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsecfailglobalcntl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsecfailglobalcntl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsecfailglobalcntl.EntityData.Children = make(map[string]types.YChild)
    cipsecfailglobalcntl.EntityData.Leafs = make(map[string]types.YLeaf)
    cipsecfailglobalcntl.EntityData.Leafs["cipSecFailTableSize"] = types.YLeaf{"Cipsecfailtablesize", cipsecfailglobalcntl.Cipsecfailtablesize}
    return &(cipsecfailglobalcntl.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cipsectrapcntl
type CISCOIPSECFLOWMONITORMIB_Cipsectrapcntl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object defines the administrative state of sending the IPsec IKE
    // Phase-1 Tunnel Start TRAP. The type is TrapStatus.
    Cipsectrapcntliketunnelstart interface{}

    // This object defines the administrative state of sending the  IPsec IKE
    // Phase-1 Tunnel Stop TRAP. The type is TrapStatus.
    Cipsectrapcntliketunnelstop interface{}

    // This object defines the administrative state of sending the  IPsec IKE
    // Phase-1 System Failure TRAP. The type is TrapStatus.
    Cipsectrapcntlikesysfailure interface{}

    // This object defines the administrative state of sending the  IPsec IKE
    // Phase-1 Certificate/CRL Failure TRAP. The type is TrapStatus.
    Cipsectrapcntlikecertcrlfailure interface{}

    // This object defines the administrative state of sending the  IPsec IKE
    // Phase-1 Protocol Failure TRAP. The type is TrapStatus.
    Cipsectrapcntlikeprotocolfail interface{}

    // This object defines the administrative state of sending the  IPsec IKE
    // Phase-1 No Security Association TRAP. The type is TrapStatus.
    Cipsectrapcntlikenosa interface{}

    // This object defines the administrative state of sending the IPsec  Phase-2
    // Tunnel Start TRAP. The type is TrapStatus.
    Cipsectrapcntlipsectunnelstart interface{}

    // This object defines the administrative state of sending the IPsec  Phase-2
    // Tunnel Stop TRAP. The type is TrapStatus.
    Cipsectrapcntlipsectunnelstop interface{}

    // This object defines the administrative state of sending the IPsec  Phase-2
    // System Failure TRAP. The type is TrapStatus.
    Cipsectrapcntlipsecsysfailure interface{}

    // This object defines the administrative state of sending the IPsec  Phase-2
    // Set Up Failure TRAP. The type is TrapStatus.
    Cipsectrapcntlipsecsetupfailure interface{}

    // This object defines the administrative state of sending the IPsec  Phase-2
    // Early Tunnel Termination TRAP. The type is TrapStatus.
    Cipsectrapcntlipsecearlytunterm interface{}

    // This object defines the administrative state of sending the IPsec  Phase-2
    // Protocol Failure TRAP. The type is TrapStatus.
    Cipsectrapcntlipsecprotocolfail interface{}

    // This object defines the administrative state of sending the IPsec  Phase-2 
    // No Security Association TRAP. The type is TrapStatus.
    Cipsectrapcntlipsecnosa interface{}
}

func (cipsectrapcntl *CISCOIPSECFLOWMONITORMIB_Cipsectrapcntl) GetEntityData() *types.CommonEntityData {
    cipsectrapcntl.EntityData.YFilter = cipsectrapcntl.YFilter
    cipsectrapcntl.EntityData.YangName = "cipSecTrapCntl"
    cipsectrapcntl.EntityData.BundleName = "cisco_ios_xe"
    cipsectrapcntl.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cipsectrapcntl.EntityData.SegmentPath = "cipSecTrapCntl"
    cipsectrapcntl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsectrapcntl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsectrapcntl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsectrapcntl.EntityData.Children = make(map[string]types.YChild)
    cipsectrapcntl.EntityData.Leafs = make(map[string]types.YLeaf)
    cipsectrapcntl.EntityData.Leafs["cipSecTrapCntlIkeTunnelStart"] = types.YLeaf{"Cipsectrapcntliketunnelstart", cipsectrapcntl.Cipsectrapcntliketunnelstart}
    cipsectrapcntl.EntityData.Leafs["cipSecTrapCntlIkeTunnelStop"] = types.YLeaf{"Cipsectrapcntliketunnelstop", cipsectrapcntl.Cipsectrapcntliketunnelstop}
    cipsectrapcntl.EntityData.Leafs["cipSecTrapCntlIkeSysFailure"] = types.YLeaf{"Cipsectrapcntlikesysfailure", cipsectrapcntl.Cipsectrapcntlikesysfailure}
    cipsectrapcntl.EntityData.Leafs["cipSecTrapCntlIkeCertCrlFailure"] = types.YLeaf{"Cipsectrapcntlikecertcrlfailure", cipsectrapcntl.Cipsectrapcntlikecertcrlfailure}
    cipsectrapcntl.EntityData.Leafs["cipSecTrapCntlIkeProtocolFail"] = types.YLeaf{"Cipsectrapcntlikeprotocolfail", cipsectrapcntl.Cipsectrapcntlikeprotocolfail}
    cipsectrapcntl.EntityData.Leafs["cipSecTrapCntlIkeNoSa"] = types.YLeaf{"Cipsectrapcntlikenosa", cipsectrapcntl.Cipsectrapcntlikenosa}
    cipsectrapcntl.EntityData.Leafs["cipSecTrapCntlIpSecTunnelStart"] = types.YLeaf{"Cipsectrapcntlipsectunnelstart", cipsectrapcntl.Cipsectrapcntlipsectunnelstart}
    cipsectrapcntl.EntityData.Leafs["cipSecTrapCntlIpSecTunnelStop"] = types.YLeaf{"Cipsectrapcntlipsectunnelstop", cipsectrapcntl.Cipsectrapcntlipsectunnelstop}
    cipsectrapcntl.EntityData.Leafs["cipSecTrapCntlIpSecSysFailure"] = types.YLeaf{"Cipsectrapcntlipsecsysfailure", cipsectrapcntl.Cipsectrapcntlipsecsysfailure}
    cipsectrapcntl.EntityData.Leafs["cipSecTrapCntlIpSecSetUpFailure"] = types.YLeaf{"Cipsectrapcntlipsecsetupfailure", cipsectrapcntl.Cipsectrapcntlipsecsetupfailure}
    cipsectrapcntl.EntityData.Leafs["cipSecTrapCntlIpSecEarlyTunTerm"] = types.YLeaf{"Cipsectrapcntlipsecearlytunterm", cipsectrapcntl.Cipsectrapcntlipsecearlytunterm}
    cipsectrapcntl.EntityData.Leafs["cipSecTrapCntlIpSecProtocolFail"] = types.YLeaf{"Cipsectrapcntlipsecprotocolfail", cipsectrapcntl.Cipsectrapcntlipsecprotocolfail}
    cipsectrapcntl.EntityData.Leafs["cipSecTrapCntlIpSecNoSa"] = types.YLeaf{"Cipsectrapcntlipsecnosa", cipsectrapcntl.Cipsectrapcntlipsecnosa}
    return &(cipsectrapcntl.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cikepeertable
// The IPsec Phase-1 Internet Key Exchange Peer Table.
// There is one entry in this table for each IPsec
// Phase-1 IKE peer association which is currently
// associated with an active IPsec Phase-1 Tunnel.
// The IPsec Phase-1 IKE Tunnel associated with this
// IPsec Phase-1 IKE peer association may or may not
// be currently active.
type CISCOIPSECFLOWMONITORMIB_Cikepeertable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with an IPsec Phase-1 IKE
    // peer association. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_Cikepeertable_Cikepeerentry.
    Cikepeerentry []CISCOIPSECFLOWMONITORMIB_Cikepeertable_Cikepeerentry
}

func (cikepeertable *CISCOIPSECFLOWMONITORMIB_Cikepeertable) GetEntityData() *types.CommonEntityData {
    cikepeertable.EntityData.YFilter = cikepeertable.YFilter
    cikepeertable.EntityData.YangName = "cikePeerTable"
    cikepeertable.EntityData.BundleName = "cisco_ios_xe"
    cikepeertable.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cikepeertable.EntityData.SegmentPath = "cikePeerTable"
    cikepeertable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cikepeertable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cikepeertable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cikepeertable.EntityData.Children = make(map[string]types.YChild)
    cikepeertable.EntityData.Children["cikePeerEntry"] = types.YChild{"Cikepeerentry", nil}
    for i := range cikepeertable.Cikepeerentry {
        cikepeertable.EntityData.Children[types.GetSegmentPath(&cikepeertable.Cikepeerentry[i])] = types.YChild{"Cikepeerentry", &cikepeertable.Cikepeerentry[i]}
    }
    cikepeertable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cikepeertable.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cikepeertable_Cikepeerentry
// Each entry contains the attributes associated
// with an IPsec Phase-1 IKE peer association.
type CISCOIPSECFLOWMONITORMIB_Cikepeertable_Cikepeerentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of local peer identity.  The local peer
    // may be identified by: 1. an IP address, or 2. a host name. The type is
    // IkePeerType.
    Cikepeerlocaltype interface{}

    // This attribute is a key. The value of the local peer identity.  If the
    // local peer type is an IP Address, then this is the IP Address used to
    // identify the local peer.  If the local peer type is a host name, then this
    // is the host name used to identify the local peer. The type is string.
    Cikepeerlocalvalue interface{}

    // This attribute is a key. The type of remote peer identity.  The remote peer
    // may be identified by: 1. an IP address, or 2. a host name. The type is
    // IkePeerType.
    Cikepeerremotetype interface{}

    // This attribute is a key. The value of the remote peer identity.  If the
    // remote peer type is an IP Address, then this is the IP Address used to
    // identify the remote peer.  If the remote peer type is a host name, then
    // this is the host name used to identify the remote peer. The type is string.
    Cikepeerremotevalue interface{}

    // This attribute is a key. The internal index of the local-remote peer
    // association.  This internal index is used  to uniquely identify multiple
    // associations between  the local and remote peer. The type is interface{}
    // with range: 1..2147483647.
    Cikepeerintindex interface{}

    // The IP address of the local peer. The type is string with length: 4 | 16.
    Cikepeerlocaladdr interface{}

    // The IP address of the remote peer. The type is string with length: 4 | 16.
    Cikepeerremoteaddr interface{}

    // The length of time that the peer association has existed in hundredths of a
    // second. The type is interface{} with range: 0..2147483647.
    Cikepeeractivetime interface{}

    // The index of the active IPsec Phase-1 IKE Tunnel (cikeTunIndex in the
    // cikeTunnelTable) for this peer association.  If an IPsec Phase-1 IKE Tunnel
    // is not currently active, then the value of this object will be zero. The
    // type is interface{} with range: 1..2147483647.
    Cikepeeractivetunnelindex interface{}
}

func (cikepeerentry *CISCOIPSECFLOWMONITORMIB_Cikepeertable_Cikepeerentry) GetEntityData() *types.CommonEntityData {
    cikepeerentry.EntityData.YFilter = cikepeerentry.YFilter
    cikepeerentry.EntityData.YangName = "cikePeerEntry"
    cikepeerentry.EntityData.BundleName = "cisco_ios_xe"
    cikepeerentry.EntityData.ParentYangName = "cikePeerTable"
    cikepeerentry.EntityData.SegmentPath = "cikePeerEntry" + "[cikePeerLocalType='" + fmt.Sprintf("%v", cikepeerentry.Cikepeerlocaltype) + "']" + "[cikePeerLocalValue='" + fmt.Sprintf("%v", cikepeerentry.Cikepeerlocalvalue) + "']" + "[cikePeerRemoteType='" + fmt.Sprintf("%v", cikepeerentry.Cikepeerremotetype) + "']" + "[cikePeerRemoteValue='" + fmt.Sprintf("%v", cikepeerentry.Cikepeerremotevalue) + "']" + "[cikePeerIntIndex='" + fmt.Sprintf("%v", cikepeerentry.Cikepeerintindex) + "']"
    cikepeerentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cikepeerentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cikepeerentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cikepeerentry.EntityData.Children = make(map[string]types.YChild)
    cikepeerentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cikepeerentry.EntityData.Leafs["cikePeerLocalType"] = types.YLeaf{"Cikepeerlocaltype", cikepeerentry.Cikepeerlocaltype}
    cikepeerentry.EntityData.Leafs["cikePeerLocalValue"] = types.YLeaf{"Cikepeerlocalvalue", cikepeerentry.Cikepeerlocalvalue}
    cikepeerentry.EntityData.Leafs["cikePeerRemoteType"] = types.YLeaf{"Cikepeerremotetype", cikepeerentry.Cikepeerremotetype}
    cikepeerentry.EntityData.Leafs["cikePeerRemoteValue"] = types.YLeaf{"Cikepeerremotevalue", cikepeerentry.Cikepeerremotevalue}
    cikepeerentry.EntityData.Leafs["cikePeerIntIndex"] = types.YLeaf{"Cikepeerintindex", cikepeerentry.Cikepeerintindex}
    cikepeerentry.EntityData.Leafs["cikePeerLocalAddr"] = types.YLeaf{"Cikepeerlocaladdr", cikepeerentry.Cikepeerlocaladdr}
    cikepeerentry.EntityData.Leafs["cikePeerRemoteAddr"] = types.YLeaf{"Cikepeerremoteaddr", cikepeerentry.Cikepeerremoteaddr}
    cikepeerentry.EntityData.Leafs["cikePeerActiveTime"] = types.YLeaf{"Cikepeeractivetime", cikepeerentry.Cikepeeractivetime}
    cikepeerentry.EntityData.Leafs["cikePeerActiveTunnelIndex"] = types.YLeaf{"Cikepeeractivetunnelindex", cikepeerentry.Cikepeeractivetunnelindex}
    return &(cikepeerentry.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Ciketunneltable
// The IPsec Phase-1 Internet Key Exchange Tunnel Table.
// There is one entry in this table for each active IPsec
// Phase-1 IKE Tunnel.
type CISCOIPSECFLOWMONITORMIB_Ciketunneltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with an active IPsec Phase-1
    // IKE Tunnel. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_Ciketunneltable_Ciketunnelentry.
    Ciketunnelentry []CISCOIPSECFLOWMONITORMIB_Ciketunneltable_Ciketunnelentry
}

func (ciketunneltable *CISCOIPSECFLOWMONITORMIB_Ciketunneltable) GetEntityData() *types.CommonEntityData {
    ciketunneltable.EntityData.YFilter = ciketunneltable.YFilter
    ciketunneltable.EntityData.YangName = "cikeTunnelTable"
    ciketunneltable.EntityData.BundleName = "cisco_ios_xe"
    ciketunneltable.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    ciketunneltable.EntityData.SegmentPath = "cikeTunnelTable"
    ciketunneltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciketunneltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciketunneltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciketunneltable.EntityData.Children = make(map[string]types.YChild)
    ciketunneltable.EntityData.Children["cikeTunnelEntry"] = types.YChild{"Ciketunnelentry", nil}
    for i := range ciketunneltable.Ciketunnelentry {
        ciketunneltable.EntityData.Children[types.GetSegmentPath(&ciketunneltable.Ciketunnelentry[i])] = types.YChild{"Ciketunnelentry", &ciketunneltable.Ciketunnelentry[i]}
    }
    ciketunneltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ciketunneltable.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Ciketunneltable_Ciketunnelentry
// Each entry contains the attributes associated with
// an active IPsec Phase-1 IKE Tunnel.
type CISCOIPSECFLOWMONITORMIB_Ciketunneltable_Ciketunnelentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index of the IPsec Phase-1 IKE Tunnel Table.
    // The value of the index is a number which begins  at one and is incremented
    // with each tunnel that  is created. The value of this object will  wrap at
    // 2,147,483,647. The type is interface{} with range: 1..2147483647.
    Ciketunindex interface{}

    // The type of local peer identity.  The local peer may be identified by:  1.
    // an IP address, or  2. a host name. The type is IkePeerType.
    Ciketunlocaltype interface{}

    // The value of the local peer identity.  If the local peer type is an IP
    // Address, then this is the IP Address used to identify the local peer.  If
    // the local peer type is a host name, then this is the host name used to
    // identify the local peer. The type is string.
    Ciketunlocalvalue interface{}

    // The IP address of the local endpoint for the IPsec Phase-1 IKE Tunnel. The
    // type is string with length: 4 | 16.
    Ciketunlocaladdr interface{}

    // The DNS name of the local IP address for the IPsec Phase-1 IKE Tunnel. If
    // the DNS  name associated with the local tunnel endpoint  is not known, then
    // the value of this  object will be a NULL string. The type is string.
    Ciketunlocalname interface{}

    // The type of remote peer identity. The remote peer may be identified by:  1.
    // an IP address, or  2. a host name. The type is IkePeerType.
    Ciketunremotetype interface{}

    // The value of the remote peer identity.  If the remote peer type is an IP
    // Address, then this is the IP Address used to identify the remote peer.  If
    // the remote peer type is a host name, then  this is the host name used to
    // identify the  remote peer. The type is string.
    Ciketunremotevalue interface{}

    // The IP address of the remote endpoint for the IPsec Phase-1 IKE Tunnel. The
    // type is string with length: 4 | 16.
    Ciketunremoteaddr interface{}

    // The DNS name of the remote IP address of IPsec Phase-1 IKE Tunnel. If the
    // DNS name associated with the remote tunnel endpoint is not known, then the
    // value of this object will be a NULL string. The type is string.
    Ciketunremotename interface{}

    // The negotiation mode of the IPsec Phase-1 IKE Tunnel. The type is
    // IkeNegoMode.
    Ciketunnegomode interface{}

    // The Diffie Hellman Group used in IPsec Phase-1 IKE negotiations. The type
    // is DiffHellmanGrp.
    Ciketundiffhellmangrp interface{}

    // The encryption algorithm used in IPsec Phase-1 IKE negotiations. The type
    // is EncryptAlgo.
    Ciketunencryptalgo interface{}

    // The hash algorithm used in IPsec Phase-1 IKE negotiations. The type is
    // IkeHashAlgo.
    Ciketunhashalgo interface{}

    // The authentication method used in IPsec Phase-1 IKE negotiations. The type
    // is IkeAuthMethod.
    Ciketunauthmethod interface{}

    // The negotiated LifeTime of the IPsec Phase-1 IKE Tunnel in seconds. The
    // type is interface{} with range: 1..2147483647. Units are seconds.
    Ciketunlifetime interface{}

    // The length of time the IPsec Phase-1 IKE tunnel has been active in
    // hundredths of seconds. The type is interface{} with range: 0..2147483647.
    Ciketunactivetime interface{}

    // The security association refresh threshold in seconds. The type is
    // interface{} with range: 1..2147483647. Units are seconds.
    Ciketunsarefreshthreshold interface{}

    // The total number of security associations refreshes performed. The type is
    // interface{} with range: 0..4294967295. Units are QM Exchanges.
    Ciketuntotalrefreshes interface{}

    // The total number of octets received by this IPsec Phase-1 IKE Tunnel. The
    // type is interface{} with range: 0..4294967295. Units are Octets.
    Ciketuninoctets interface{}

    // The total number of packets received by this IPsec Phase-1 IKE Tunnel. The
    // type is interface{} with range: 0..4294967295. Units are Packets.
    Ciketuninpkts interface{}

    // The total number of packets dropped by this IPsec Phase-1 IKE Tunnel during
    // receive processing. The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    Ciketunindroppkts interface{}

    // The total number of notifys received by this IPsec Phase-1 IKE Tunnel. The
    // type is interface{} with range: 0..4294967295. Units are Notification
    // Payloads.
    Ciketuninnotifys interface{}

    // The total number of IPsec Phase-2 exchanges received by  this IPsec Phase-1
    // IKE Tunnel. The type is interface{} with range: 0..4294967295. Units are SA
    // Payloads.
    Ciketuninp2Exchgs interface{}

    // The total number of IPsec Phase-2 exchanges received and found to be
    // invalid  by this IPsec Phase-1 IKE Tunnel. The type is interface{} with
    // range: 0..4294967295. Units are SA Payloads.
    Ciketuninp2Exchginvalids interface{}

    // The total number of IPsec Phase-2 exchanges received and rejected by this
    // IPsec Phase-1  Tunnel. The type is interface{} with range: 0..4294967295.
    // Units are SA Payloads.
    Ciketuninp2Exchgrejects interface{}

    // The total number of IPsec Phase-2 security association delete requests
    // received  by this IPsec Phase-1 IKE Tunnel. The type is interface{} with
    // range: 0..4294967295. Units are Notification Payloads.
    Ciketuninp2Sadelrequests interface{}

    // The total number of octets sent by this IPsec Phase-1 IKE Tunnel. The type
    // is interface{} with range: 0..4294967295. Units are Octets.
    Ciketunoutoctets interface{}

    // The total number of packets sent by this IPsec Phase-1 IKE Tunnel. The type
    // is interface{} with range: 0..4294967295. Units are Packets.
    Ciketunoutpkts interface{}

    // The total number of packets dropped by this IPsec Phase-1 IKE Tunnel during
    // send processing. The type is interface{} with range: 0..4294967295. Units
    // are Packets.
    Ciketunoutdroppkts interface{}

    // The total number of notifys sent by this IPsec Phase-1 Tunnel. The type is
    // interface{} with range: 0..4294967295. Units are Notification Payloads.
    Ciketunoutnotifys interface{}

    // The total number of IPsec Phase-2 exchanges sent by this IPsec Phase-1 IKE
    // Tunnel. The type is interface{} with range: 0..4294967295. Units are SA
    // Payloads.
    Ciketunoutp2Exchgs interface{}

    // The total number of IPsec Phase-2 exchanges sent and found to be invalid by
    // this IPsec Phase-1 IKE Tunnel. The type is interface{} with range:
    // 0..4294967295. Units are SA Payloads.
    Ciketunoutp2Exchginvalids interface{}

    // The total number of IPsec Phase-2 exchanges sent and rejected by this IPsec
    // Phase-1 IKE Tunnel. The type is interface{} with range: 0..4294967295.
    // Units are SA Payloads.
    Ciketunoutp2Exchgrejects interface{}

    // The total number of IPsec Phase-2 security association delete requests sent
    // by this IPsec Phase-1 IKE Tunnel. The type is interface{} with range:
    // 0..4294967295. Units are Notification Payloads.
    Ciketunoutp2Sadelrequests interface{}

    // The status of the MIB table row.  This object can be used to bring the
    // tunnel down  by setting value of this object to destroy(2).  This object
    // cannot be used to create  a MIB table row. The type is TunnelStatus.
    Ciketunstatus interface{}
}

func (ciketunnelentry *CISCOIPSECFLOWMONITORMIB_Ciketunneltable_Ciketunnelentry) GetEntityData() *types.CommonEntityData {
    ciketunnelentry.EntityData.YFilter = ciketunnelentry.YFilter
    ciketunnelentry.EntityData.YangName = "cikeTunnelEntry"
    ciketunnelentry.EntityData.BundleName = "cisco_ios_xe"
    ciketunnelentry.EntityData.ParentYangName = "cikeTunnelTable"
    ciketunnelentry.EntityData.SegmentPath = "cikeTunnelEntry" + "[cikeTunIndex='" + fmt.Sprintf("%v", ciketunnelentry.Ciketunindex) + "']"
    ciketunnelentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciketunnelentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciketunnelentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciketunnelentry.EntityData.Children = make(map[string]types.YChild)
    ciketunnelentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ciketunnelentry.EntityData.Leafs["cikeTunIndex"] = types.YLeaf{"Ciketunindex", ciketunnelentry.Ciketunindex}
    ciketunnelentry.EntityData.Leafs["cikeTunLocalType"] = types.YLeaf{"Ciketunlocaltype", ciketunnelentry.Ciketunlocaltype}
    ciketunnelentry.EntityData.Leafs["cikeTunLocalValue"] = types.YLeaf{"Ciketunlocalvalue", ciketunnelentry.Ciketunlocalvalue}
    ciketunnelentry.EntityData.Leafs["cikeTunLocalAddr"] = types.YLeaf{"Ciketunlocaladdr", ciketunnelentry.Ciketunlocaladdr}
    ciketunnelentry.EntityData.Leafs["cikeTunLocalName"] = types.YLeaf{"Ciketunlocalname", ciketunnelentry.Ciketunlocalname}
    ciketunnelentry.EntityData.Leafs["cikeTunRemoteType"] = types.YLeaf{"Ciketunremotetype", ciketunnelentry.Ciketunremotetype}
    ciketunnelentry.EntityData.Leafs["cikeTunRemoteValue"] = types.YLeaf{"Ciketunremotevalue", ciketunnelentry.Ciketunremotevalue}
    ciketunnelentry.EntityData.Leafs["cikeTunRemoteAddr"] = types.YLeaf{"Ciketunremoteaddr", ciketunnelentry.Ciketunremoteaddr}
    ciketunnelentry.EntityData.Leafs["cikeTunRemoteName"] = types.YLeaf{"Ciketunremotename", ciketunnelentry.Ciketunremotename}
    ciketunnelentry.EntityData.Leafs["cikeTunNegoMode"] = types.YLeaf{"Ciketunnegomode", ciketunnelentry.Ciketunnegomode}
    ciketunnelentry.EntityData.Leafs["cikeTunDiffHellmanGrp"] = types.YLeaf{"Ciketundiffhellmangrp", ciketunnelentry.Ciketundiffhellmangrp}
    ciketunnelentry.EntityData.Leafs["cikeTunEncryptAlgo"] = types.YLeaf{"Ciketunencryptalgo", ciketunnelentry.Ciketunencryptalgo}
    ciketunnelentry.EntityData.Leafs["cikeTunHashAlgo"] = types.YLeaf{"Ciketunhashalgo", ciketunnelentry.Ciketunhashalgo}
    ciketunnelentry.EntityData.Leafs["cikeTunAuthMethod"] = types.YLeaf{"Ciketunauthmethod", ciketunnelentry.Ciketunauthmethod}
    ciketunnelentry.EntityData.Leafs["cikeTunLifeTime"] = types.YLeaf{"Ciketunlifetime", ciketunnelentry.Ciketunlifetime}
    ciketunnelentry.EntityData.Leafs["cikeTunActiveTime"] = types.YLeaf{"Ciketunactivetime", ciketunnelentry.Ciketunactivetime}
    ciketunnelentry.EntityData.Leafs["cikeTunSaRefreshThreshold"] = types.YLeaf{"Ciketunsarefreshthreshold", ciketunnelentry.Ciketunsarefreshthreshold}
    ciketunnelentry.EntityData.Leafs["cikeTunTotalRefreshes"] = types.YLeaf{"Ciketuntotalrefreshes", ciketunnelentry.Ciketuntotalrefreshes}
    ciketunnelentry.EntityData.Leafs["cikeTunInOctets"] = types.YLeaf{"Ciketuninoctets", ciketunnelentry.Ciketuninoctets}
    ciketunnelentry.EntityData.Leafs["cikeTunInPkts"] = types.YLeaf{"Ciketuninpkts", ciketunnelentry.Ciketuninpkts}
    ciketunnelentry.EntityData.Leafs["cikeTunInDropPkts"] = types.YLeaf{"Ciketunindroppkts", ciketunnelentry.Ciketunindroppkts}
    ciketunnelentry.EntityData.Leafs["cikeTunInNotifys"] = types.YLeaf{"Ciketuninnotifys", ciketunnelentry.Ciketuninnotifys}
    ciketunnelentry.EntityData.Leafs["cikeTunInP2Exchgs"] = types.YLeaf{"Ciketuninp2Exchgs", ciketunnelentry.Ciketuninp2Exchgs}
    ciketunnelentry.EntityData.Leafs["cikeTunInP2ExchgInvalids"] = types.YLeaf{"Ciketuninp2Exchginvalids", ciketunnelentry.Ciketuninp2Exchginvalids}
    ciketunnelentry.EntityData.Leafs["cikeTunInP2ExchgRejects"] = types.YLeaf{"Ciketuninp2Exchgrejects", ciketunnelentry.Ciketuninp2Exchgrejects}
    ciketunnelentry.EntityData.Leafs["cikeTunInP2SaDelRequests"] = types.YLeaf{"Ciketuninp2Sadelrequests", ciketunnelentry.Ciketuninp2Sadelrequests}
    ciketunnelentry.EntityData.Leafs["cikeTunOutOctets"] = types.YLeaf{"Ciketunoutoctets", ciketunnelentry.Ciketunoutoctets}
    ciketunnelentry.EntityData.Leafs["cikeTunOutPkts"] = types.YLeaf{"Ciketunoutpkts", ciketunnelentry.Ciketunoutpkts}
    ciketunnelentry.EntityData.Leafs["cikeTunOutDropPkts"] = types.YLeaf{"Ciketunoutdroppkts", ciketunnelentry.Ciketunoutdroppkts}
    ciketunnelentry.EntityData.Leafs["cikeTunOutNotifys"] = types.YLeaf{"Ciketunoutnotifys", ciketunnelentry.Ciketunoutnotifys}
    ciketunnelentry.EntityData.Leafs["cikeTunOutP2Exchgs"] = types.YLeaf{"Ciketunoutp2Exchgs", ciketunnelentry.Ciketunoutp2Exchgs}
    ciketunnelentry.EntityData.Leafs["cikeTunOutP2ExchgInvalids"] = types.YLeaf{"Ciketunoutp2Exchginvalids", ciketunnelentry.Ciketunoutp2Exchginvalids}
    ciketunnelentry.EntityData.Leafs["cikeTunOutP2ExchgRejects"] = types.YLeaf{"Ciketunoutp2Exchgrejects", ciketunnelentry.Ciketunoutp2Exchgrejects}
    ciketunnelentry.EntityData.Leafs["cikeTunOutP2SaDelRequests"] = types.YLeaf{"Ciketunoutp2Sadelrequests", ciketunnelentry.Ciketunoutp2Sadelrequests}
    ciketunnelentry.EntityData.Leafs["cikeTunStatus"] = types.YLeaf{"Ciketunstatus", ciketunnelentry.Ciketunstatus}
    return &(ciketunnelentry.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable
// The IPsec Phase-1 Internet Key Exchange Peer
// Association to IPsec Phase-2 Tunnel
// Correlation Table. There is one entry in
// this table for each active IPsec Phase-2
// Tunnel.
type CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes of an IPsec Phase-1 IKE Peer Association
    // to IPsec Phase-2 Tunnel Correlation. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable_Cikepeercorrentry.
    Cikepeercorrentry []CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable_Cikepeercorrentry
}

func (cikepeercorrtable *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable) GetEntityData() *types.CommonEntityData {
    cikepeercorrtable.EntityData.YFilter = cikepeercorrtable.YFilter
    cikepeercorrtable.EntityData.YangName = "cikePeerCorrTable"
    cikepeercorrtable.EntityData.BundleName = "cisco_ios_xe"
    cikepeercorrtable.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cikepeercorrtable.EntityData.SegmentPath = "cikePeerCorrTable"
    cikepeercorrtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cikepeercorrtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cikepeercorrtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cikepeercorrtable.EntityData.Children = make(map[string]types.YChild)
    cikepeercorrtable.EntityData.Children["cikePeerCorrEntry"] = types.YChild{"Cikepeercorrentry", nil}
    for i := range cikepeercorrtable.Cikepeercorrentry {
        cikepeercorrtable.EntityData.Children[types.GetSegmentPath(&cikepeercorrtable.Cikepeercorrentry[i])] = types.YChild{"Cikepeercorrentry", &cikepeercorrtable.Cikepeercorrentry[i]}
    }
    cikepeercorrtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cikepeercorrtable.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable_Cikepeercorrentry
// Each entry contains the attributes of an
// IPsec Phase-1 IKE Peer Association to IPsec
// Phase-2 Tunnel Correlation.
type CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable_Cikepeercorrentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type of local peer identity. The local peer
    // may be identified by: 1. an IP address, or 2. a host name. The type is
    // IkePeerType.
    Cikepeercorrlocaltype interface{}

    // This attribute is a key. The value of the local peer identity.  If the
    // local peer type is an IP Address, then this is the IP Address used to
    // identify the local peer.  If the local peer type is a host name, then this
    // is the host name used to identify the local peer. The type is string.
    Cikepeercorrlocalvalue interface{}

    // This attribute is a key. The type of remote peer identity. The remote peer
    // may be identified by: 1. an IP address, or 2. a host name. The type is
    // IkePeerType.
    Cikepeercorrremotetype interface{}

    // This attribute is a key. The value of the remote peer identity.  If the
    // remote peer type is an IP Address, then this is the IP Address used to
    // identify the remote peer.  If the remote peer type is a host name, then
    // this is the host name used to identify the remote peer. The type is string.
    Cikepeercorrremotevalue interface{}

    // This attribute is a key. The internal index of the local-remote peer
    // association.  This internal index is  used to uniquely identify multiple
    // associations  between the local and remote peer. The type is interface{}
    // with range: 1..2147483647.
    Cikepeercorrintindex interface{}

    // This attribute is a key. The sequence number of the local-remote peer
    // association.  This sequence number is  used to uniquely identify multiple
    // instances  of an unique association between  the local and remote peer. The
    // type is interface{} with range: 1..2147483647.
    Cikepeercorrseqnum interface{}

    // The index of the active IPsec Phase-2 Tunnel (cipSecTunIndex in the
    // cipSecTunnelTable) for this IPsec Phase-1 IKE Peer Association. The type is
    // interface{} with range: 1..2147483647.
    Cikepeercorripsectunindex interface{}
}

func (cikepeercorrentry *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable_Cikepeercorrentry) GetEntityData() *types.CommonEntityData {
    cikepeercorrentry.EntityData.YFilter = cikepeercorrentry.YFilter
    cikepeercorrentry.EntityData.YangName = "cikePeerCorrEntry"
    cikepeercorrentry.EntityData.BundleName = "cisco_ios_xe"
    cikepeercorrentry.EntityData.ParentYangName = "cikePeerCorrTable"
    cikepeercorrentry.EntityData.SegmentPath = "cikePeerCorrEntry" + "[cikePeerCorrLocalType='" + fmt.Sprintf("%v", cikepeercorrentry.Cikepeercorrlocaltype) + "']" + "[cikePeerCorrLocalValue='" + fmt.Sprintf("%v", cikepeercorrentry.Cikepeercorrlocalvalue) + "']" + "[cikePeerCorrRemoteType='" + fmt.Sprintf("%v", cikepeercorrentry.Cikepeercorrremotetype) + "']" + "[cikePeerCorrRemoteValue='" + fmt.Sprintf("%v", cikepeercorrentry.Cikepeercorrremotevalue) + "']" + "[cikePeerCorrIntIndex='" + fmt.Sprintf("%v", cikepeercorrentry.Cikepeercorrintindex) + "']" + "[cikePeerCorrSeqNum='" + fmt.Sprintf("%v", cikepeercorrentry.Cikepeercorrseqnum) + "']"
    cikepeercorrentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cikepeercorrentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cikepeercorrentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cikepeercorrentry.EntityData.Children = make(map[string]types.YChild)
    cikepeercorrentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cikepeercorrentry.EntityData.Leafs["cikePeerCorrLocalType"] = types.YLeaf{"Cikepeercorrlocaltype", cikepeercorrentry.Cikepeercorrlocaltype}
    cikepeercorrentry.EntityData.Leafs["cikePeerCorrLocalValue"] = types.YLeaf{"Cikepeercorrlocalvalue", cikepeercorrentry.Cikepeercorrlocalvalue}
    cikepeercorrentry.EntityData.Leafs["cikePeerCorrRemoteType"] = types.YLeaf{"Cikepeercorrremotetype", cikepeercorrentry.Cikepeercorrremotetype}
    cikepeercorrentry.EntityData.Leafs["cikePeerCorrRemoteValue"] = types.YLeaf{"Cikepeercorrremotevalue", cikepeercorrentry.Cikepeercorrremotevalue}
    cikepeercorrentry.EntityData.Leafs["cikePeerCorrIntIndex"] = types.YLeaf{"Cikepeercorrintindex", cikepeercorrentry.Cikepeercorrintindex}
    cikepeercorrentry.EntityData.Leafs["cikePeerCorrSeqNum"] = types.YLeaf{"Cikepeercorrseqnum", cikepeercorrentry.Cikepeercorrseqnum}
    cikepeercorrentry.EntityData.Leafs["cikePeerCorrIpSecTunIndex"] = types.YLeaf{"Cikepeercorripsectunindex", cikepeercorrentry.Cikepeercorripsectunindex}
    return &(cikepeercorrentry.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable
// Phase-1 IKE stats information is included in this table.
// Each entry is related to a specific gateway which is 
// identified by 'cmgwIndex'.
type CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes of an Phase-1 IKE stats information for
    // the related gateway.  There is only one entry for each gateway. The entry 
    // is created when a gateway up and cannot be deleted. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable_Cikephase1Gwstatsentry.
    Cikephase1Gwstatsentry []CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable_Cikephase1Gwstatsentry
}

func (cikephase1Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable) GetEntityData() *types.CommonEntityData {
    cikephase1Gwstatstable.EntityData.YFilter = cikephase1Gwstatstable.YFilter
    cikephase1Gwstatstable.EntityData.YangName = "cikePhase1GWStatsTable"
    cikephase1Gwstatstable.EntityData.BundleName = "cisco_ios_xe"
    cikephase1Gwstatstable.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cikephase1Gwstatstable.EntityData.SegmentPath = "cikePhase1GWStatsTable"
    cikephase1Gwstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cikephase1Gwstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cikephase1Gwstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cikephase1Gwstatstable.EntityData.Children = make(map[string]types.YChild)
    cikephase1Gwstatstable.EntityData.Children["cikePhase1GWStatsEntry"] = types.YChild{"Cikephase1Gwstatsentry", nil}
    for i := range cikephase1Gwstatstable.Cikephase1Gwstatsentry {
        cikephase1Gwstatstable.EntityData.Children[types.GetSegmentPath(&cikephase1Gwstatstable.Cikephase1Gwstatsentry[i])] = types.YChild{"Cikephase1Gwstatsentry", &cikephase1Gwstatstable.Cikephase1Gwstatsentry[i]}
    }
    cikephase1Gwstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cikephase1Gwstatstable.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable_Cikephase1Gwstatsentry
// Each entry contains the attributes of an Phase-1 IKE stats
// information for the related gateway.
// 
// There is only one entry for each gateway. The entry 
// is created when a gateway up and cannot be deleted.
type CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable_Cikephase1Gwstatsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_media_gateway_mib.CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry_Cmgwindex
    Cmgwindex interface{}

    // The number of currently active IPsec Phase-1 IKE Tunnels. The type is
    // interface{} with range: 0..4294967295.
    Cikephase1Gwactivetunnels interface{}

    // The total number of previously active IPsec Phase-1 IKE Tunnels. The type
    // is interface{} with range: 0..4294967295. Units are SAs.
    Cikephase1Gwprevioustunnels interface{}

    // The total number of octets received by all currently and previously active
    // IPsec Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Octets.
    Cikephase1Gwinoctets interface{}

    // The total number of packets received by all currently and previously active
    // IPsec  Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    Cikephase1Gwinpkts interface{}

    // The total number of packets which were dropped during receive processing by
    // all  currently and previously active IPsec Phase-1 IKE Tunnels. The type is
    // interface{} with range: 0..4294967295. Units are Packets.
    Cikephase1Gwindroppkts interface{}

    // The total number of notifys received by all currently and previously active
    // IPsec  Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Notification Payloads.
    Cikephase1Gwinnotifys interface{}

    // The total number of IPsec Phase-2 exchanges received by all currently and
    // previously  active IPsec Phase-1 IKE Tunnels. The type is interface{} with
    // range: 0..4294967295. Units are SA Payloads.
    Cikephase1Gwinp2Exchgs interface{}

    // The total number of IPsec Phase-2 exchanges which were received and found
    // to be invalid  by all currently and previously active IPsec  Phase-1 IKE
    // Tunnels. The type is interface{} with range: 0..4294967295. Units are SA
    // Payloads.
    Cikephase1Gwinp2Exchginvalids interface{}

    // The total number of IPsec Phase-2 exchanges which were received and
    // rejected by all  currently and previously active IPsec Phase-1  IKE
    // Tunnels. The type is interface{} with range: 0..4294967295. Units are SA
    // Payloads.
    Cikephase1Gwinp2Exchgrejects interface{}

    // The total number of IPsec Phase-2 'Security Association' delete requests
    // received by all  currently and previously active and IPsec  Phase-1 IKE
    // Tunnels. The type is interface{} with range: 0..4294967295. Units are
    // Notification Payloads.
    Cikephase1Gwinp2Sadelrequests interface{}

    // The total number of octets sent by all currently and previously active and
    // IPsec Phase-1  IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Octets.
    Cikephase1Gwoutoctets interface{}

    // The total number of packets sent by all currently and previously active and
    // IPsec Phase-1  Tunnels. The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    Cikephase1Gwoutpkts interface{}

    // The total number of packets which were dropped during send processing by
    // all currently  and previously active IPsec Phase-1 IKE Tunnels. The type is
    // interface{} with range: 0..4294967295. Units are Packets.
    Cikephase1Gwoutdroppkts interface{}

    // The total number of notifys sent by all currently and previously active
    // IPsec Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Notification Payloads.
    Cikephase1Gwoutnotifys interface{}

    // The total number of IPsec Phase-2 exchanges which were sent by all
    // currently and previously  active IPsec Phase-1 IKE Tunnels. The type is
    // interface{} with range: 0..4294967295. Units are SA Payloads.
    Cikephase1Gwoutp2Exchgs interface{}

    // The total number of IPsec Phase-2 exchanges which were sent and found to be
    // invalid by  all currently and previously active IPsec Phase-1  Tunnels. The
    // type is interface{} with range: 0..4294967295. Units are SA Payloads.
    Cikephase1Gwoutp2Exchginvalids interface{}

    // The total number of IPsec Phase-2 exchanges which were sent and rejected by
    // all currently and previously active IPsec Phase-1 IKE Tunnels. The type is
    // interface{} with range: 0..4294967295. Units are SA Payloads.
    Cikephase1Gwoutp2Exchgrejects interface{}

    // The total number of IPsec Phase-2 SA delete requests sent by all currently
    // and  previously active IPsec Phase-1 IKE Tunnels. The type is interface{}
    // with range: 0..4294967295. Units are Notification Payloads.
    Cikephase1Gwoutp2Sadelrequests interface{}

    // The total number of IPsec Phase-1 IKE Tunnels which were locally initiated.
    // The type is interface{} with range: 0..4294967295. Units are SAs.
    Cikephase1Gwinittunnels interface{}

    // The total number of IPsec Phase-1 IKE Tunnels which were locally initiated
    // and failed to activate. The type is interface{} with range: 0..4294967295.
    // Units are SAs.
    Cikephase1Gwinittunnelfails interface{}

    // The total number of IPsec Phase-1 IKE Tunnels which were remotely initiated
    // and failed to activate. The type is interface{} with range: 0..4294967295.
    // Units are SAs.
    Cikephase1Gwresptunnelfails interface{}

    // The total number of system capacity failures which occurred during
    // processing of all current  and previously active IPsec Phase-1 IKE Tunnels.
    // The type is interface{} with range: 0..4294967295. Units are Failures.
    Cikephase1Gwsyscapfails interface{}

    // The total number of authentications which ended in failure by all current
    // and previous IPsec Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Failures.
    Cikephase1Gwauthfails interface{}

    // The total number of decryptions which ended in failure by all current and
    // previous IPsec Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Failures.
    Cikephase1Gwdecryptfails interface{}

    // The total number of hash validations which ended in failure by all current
    // and previous IPsec Phase-1 IKE Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Failures.
    Cikephase1Gwhashvalidfails interface{}

    // The total number of non-existent 'Security Association' failures occurred
    // during processing of current and  previous IPsec Phase-1 IKE Tunnels. The
    // type is interface{} with range: 0..4294967295. Units are Failures.
    Cikephase1Gwnosafails interface{}
}

func (cikephase1Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable_Cikephase1Gwstatsentry) GetEntityData() *types.CommonEntityData {
    cikephase1Gwstatsentry.EntityData.YFilter = cikephase1Gwstatsentry.YFilter
    cikephase1Gwstatsentry.EntityData.YangName = "cikePhase1GWStatsEntry"
    cikephase1Gwstatsentry.EntityData.BundleName = "cisco_ios_xe"
    cikephase1Gwstatsentry.EntityData.ParentYangName = "cikePhase1GWStatsTable"
    cikephase1Gwstatsentry.EntityData.SegmentPath = "cikePhase1GWStatsEntry" + "[cmgwIndex='" + fmt.Sprintf("%v", cikephase1Gwstatsentry.Cmgwindex) + "']"
    cikephase1Gwstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cikephase1Gwstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cikephase1Gwstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cikephase1Gwstatsentry.EntityData.Children = make(map[string]types.YChild)
    cikephase1Gwstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cikephase1Gwstatsentry.EntityData.Leafs["cmgwIndex"] = types.YLeaf{"Cmgwindex", cikephase1Gwstatsentry.Cmgwindex}
    cikephase1Gwstatsentry.EntityData.Leafs["cikePhase1GWActiveTunnels"] = types.YLeaf{"Cikephase1Gwactivetunnels", cikephase1Gwstatsentry.Cikephase1Gwactivetunnels}
    cikephase1Gwstatsentry.EntityData.Leafs["cikePhase1GWPreviousTunnels"] = types.YLeaf{"Cikephase1Gwprevioustunnels", cikephase1Gwstatsentry.Cikephase1Gwprevioustunnels}
    cikephase1Gwstatsentry.EntityData.Leafs["cikePhase1GWInOctets"] = types.YLeaf{"Cikephase1Gwinoctets", cikephase1Gwstatsentry.Cikephase1Gwinoctets}
    cikephase1Gwstatsentry.EntityData.Leafs["cikePhase1GWInPkts"] = types.YLeaf{"Cikephase1Gwinpkts", cikephase1Gwstatsentry.Cikephase1Gwinpkts}
    cikephase1Gwstatsentry.EntityData.Leafs["cikePhase1GWInDropPkts"] = types.YLeaf{"Cikephase1Gwindroppkts", cikephase1Gwstatsentry.Cikephase1Gwindroppkts}
    cikephase1Gwstatsentry.EntityData.Leafs["cikePhase1GWInNotifys"] = types.YLeaf{"Cikephase1Gwinnotifys", cikephase1Gwstatsentry.Cikephase1Gwinnotifys}
    cikephase1Gwstatsentry.EntityData.Leafs["cikePhase1GWInP2Exchgs"] = types.YLeaf{"Cikephase1Gwinp2Exchgs", cikephase1Gwstatsentry.Cikephase1Gwinp2Exchgs}
    cikephase1Gwstatsentry.EntityData.Leafs["cikePhase1GWInP2ExchgInvalids"] = types.YLeaf{"Cikephase1Gwinp2Exchginvalids", cikephase1Gwstatsentry.Cikephase1Gwinp2Exchginvalids}
    cikephase1Gwstatsentry.EntityData.Leafs["cikePhase1GWInP2ExchgRejects"] = types.YLeaf{"Cikephase1Gwinp2Exchgrejects", cikephase1Gwstatsentry.Cikephase1Gwinp2Exchgrejects}
    cikephase1Gwstatsentry.EntityData.Leafs["cikePhase1GWInP2SaDelRequests"] = types.YLeaf{"Cikephase1Gwinp2Sadelrequests", cikephase1Gwstatsentry.Cikephase1Gwinp2Sadelrequests}
    cikephase1Gwstatsentry.EntityData.Leafs["cikePhase1GWOutOctets"] = types.YLeaf{"Cikephase1Gwoutoctets", cikephase1Gwstatsentry.Cikephase1Gwoutoctets}
    cikephase1Gwstatsentry.EntityData.Leafs["cikePhase1GWOutPkts"] = types.YLeaf{"Cikephase1Gwoutpkts", cikephase1Gwstatsentry.Cikephase1Gwoutpkts}
    cikephase1Gwstatsentry.EntityData.Leafs["cikePhase1GWOutDropPkts"] = types.YLeaf{"Cikephase1Gwoutdroppkts", cikephase1Gwstatsentry.Cikephase1Gwoutdroppkts}
    cikephase1Gwstatsentry.EntityData.Leafs["cikePhase1GWOutNotifys"] = types.YLeaf{"Cikephase1Gwoutnotifys", cikephase1Gwstatsentry.Cikephase1Gwoutnotifys}
    cikephase1Gwstatsentry.EntityData.Leafs["cikePhase1GWOutP2Exchgs"] = types.YLeaf{"Cikephase1Gwoutp2Exchgs", cikephase1Gwstatsentry.Cikephase1Gwoutp2Exchgs}
    cikephase1Gwstatsentry.EntityData.Leafs["cikePhase1GWOutP2ExchgInvalids"] = types.YLeaf{"Cikephase1Gwoutp2Exchginvalids", cikephase1Gwstatsentry.Cikephase1Gwoutp2Exchginvalids}
    cikephase1Gwstatsentry.EntityData.Leafs["cikePhase1GWOutP2ExchgRejects"] = types.YLeaf{"Cikephase1Gwoutp2Exchgrejects", cikephase1Gwstatsentry.Cikephase1Gwoutp2Exchgrejects}
    cikephase1Gwstatsentry.EntityData.Leafs["cikePhase1GWOutP2SaDelRequests"] = types.YLeaf{"Cikephase1Gwoutp2Sadelrequests", cikephase1Gwstatsentry.Cikephase1Gwoutp2Sadelrequests}
    cikephase1Gwstatsentry.EntityData.Leafs["cikePhase1GWInitTunnels"] = types.YLeaf{"Cikephase1Gwinittunnels", cikephase1Gwstatsentry.Cikephase1Gwinittunnels}
    cikephase1Gwstatsentry.EntityData.Leafs["cikePhase1GWInitTunnelFails"] = types.YLeaf{"Cikephase1Gwinittunnelfails", cikephase1Gwstatsentry.Cikephase1Gwinittunnelfails}
    cikephase1Gwstatsentry.EntityData.Leafs["cikePhase1GWRespTunnelFails"] = types.YLeaf{"Cikephase1Gwresptunnelfails", cikephase1Gwstatsentry.Cikephase1Gwresptunnelfails}
    cikephase1Gwstatsentry.EntityData.Leafs["cikePhase1GWSysCapFails"] = types.YLeaf{"Cikephase1Gwsyscapfails", cikephase1Gwstatsentry.Cikephase1Gwsyscapfails}
    cikephase1Gwstatsentry.EntityData.Leafs["cikePhase1GWAuthFails"] = types.YLeaf{"Cikephase1Gwauthfails", cikephase1Gwstatsentry.Cikephase1Gwauthfails}
    cikephase1Gwstatsentry.EntityData.Leafs["cikePhase1GWDecryptFails"] = types.YLeaf{"Cikephase1Gwdecryptfails", cikephase1Gwstatsentry.Cikephase1Gwdecryptfails}
    cikephase1Gwstatsentry.EntityData.Leafs["cikePhase1GWHashValidFails"] = types.YLeaf{"Cikephase1Gwhashvalidfails", cikephase1Gwstatsentry.Cikephase1Gwhashvalidfails}
    cikephase1Gwstatsentry.EntityData.Leafs["cikePhase1GWNoSaFails"] = types.YLeaf{"Cikephase1Gwnosafails", cikephase1Gwstatsentry.Cikephase1Gwnosafails}
    return &(cikephase1Gwstatsentry.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cipsectunneltable
// The IPsec Phase-2 Tunnel Table.
// There is one entry in this table for 
// each active IPsec Phase-2 Tunnel.
type CISCOIPSECFLOWMONITORMIB_Cipsectunneltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with an active IPsec Phase-2
    // Tunnel. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_Cipsectunneltable_Cipsectunnelentry.
    Cipsectunnelentry []CISCOIPSECFLOWMONITORMIB_Cipsectunneltable_Cipsectunnelentry
}

func (cipsectunneltable *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable) GetEntityData() *types.CommonEntityData {
    cipsectunneltable.EntityData.YFilter = cipsectunneltable.YFilter
    cipsectunneltable.EntityData.YangName = "cipSecTunnelTable"
    cipsectunneltable.EntityData.BundleName = "cisco_ios_xe"
    cipsectunneltable.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cipsectunneltable.EntityData.SegmentPath = "cipSecTunnelTable"
    cipsectunneltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsectunneltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsectunneltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsectunneltable.EntityData.Children = make(map[string]types.YChild)
    cipsectunneltable.EntityData.Children["cipSecTunnelEntry"] = types.YChild{"Cipsectunnelentry", nil}
    for i := range cipsectunneltable.Cipsectunnelentry {
        cipsectunneltable.EntityData.Children[types.GetSegmentPath(&cipsectunneltable.Cipsectunnelentry[i])] = types.YChild{"Cipsectunnelentry", &cipsectunneltable.Cipsectunnelentry[i]}
    }
    cipsectunneltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cipsectunneltable.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cipsectunneltable_Cipsectunnelentry
// Each entry contains the attributes
// associated with an active IPsec Phase-2 Tunnel.
type CISCOIPSECFLOWMONITORMIB_Cipsectunneltable_Cipsectunnelentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index of the IPsec Phase-2 Tunnel Table. The
    // value of the index is a number which begins  at one and is incremented with
    // each tunnel that  is created. The value of this object will wrap  at
    // 2,147,483,647. The type is interface{} with range: 1..2147483647.
    Cipsectunindex interface{}

    // The index of the associated IPsec Phase-1 IKE Tunnel.  (cikeTunIndex in the
    // cikeTunnelTable). The type is interface{} with range: 1..2147483647.
    Cipsectuniketunnelindex interface{}

    // An indicator which specifies whether or not the IPsec Phase-1 IKE Tunnel
    // currently exists. The type is bool.
    Cipsectuniketunnelalive interface{}

    // The IP address of the local endpoint for the IPsec Phase-2 Tunnel. The type
    // is string with length: 4 | 16.
    Cipsectunlocaladdr interface{}

    // The IP address of the remote endpoint for the IPsec Phase-2 Tunnel. The
    // type is string with length: 4 | 16.
    Cipsectunremoteaddr interface{}

    // The type of key used by the IPsec Phase-2 Tunnel. The type is KeyType.
    Cipsectunkeytype interface{}

    // The encapsulation mode used by the IPsec Phase-2 Tunnel. The type is
    // EncapMode.
    Cipsectunencapmode interface{}

    // The negotiated LifeSize of the IPsec Phase-2 Tunnel in kilobytes. The type
    // is interface{} with range: 1..2147483647. Units are KBytes.
    Cipsectunlifesize interface{}

    // The negotiated LifeTime of the IPsec Phase-2 Tunnel in seconds. The type is
    // interface{} with range: 1..2147483647. Units are Seconds.
    Cipsectunlifetime interface{}

    // The length of time the IPsec Phase-2 Tunnel has been  active in hundredths
    // of seconds. The type is interface{} with range: 0..2147483647.
    Cipsectunactivetime interface{}

    // The security association LifeSize refresh threshold in kilobytes. The type
    // is interface{} with range: 1..2147483647. Units are KBytes.
    Cipsectunsalifesizethreshold interface{}

    // The security association LifeTime refresh threshold in seconds. The type is
    // interface{} with range: 1..2147483647. Units are Seconds.
    Cipsectunsalifetimethreshold interface{}

    // The total number of security association refreshes performed. The type is
    // interface{} with range: 0..4294967295. Units are QM Exchanges.
    Cipsectuntotalrefreshes interface{}

    // The total number of security associations which have expired. The type is
    // interface{} with range: 0..4294967295. Units are SAs.
    Cipsectunexpiredsainstances interface{}

    // The number of security associations which are currently active or expiring.
    // The type is interface{} with range: 0..4294967295.
    Cipsectuncurrentsainstances interface{}

    // The Diffie Hellman Group used by the inbound security association of the 
    // IPsec Phase-2 Tunnel. The type is DiffHellmanGrp.
    Cipsectuninsadiffhellmangrp interface{}

    // The encryption algorithm used by the inbound security association of the
    // IPsec Phase-2 Tunnel. The type is EncryptAlgo.
    Cipsectuninsaencryptalgo interface{}

    // The authentication algorithm used by the inbound authentication header (AH)
    // security association of the IPsec Phase-2 Tunnel. The type is AuthAlgo.
    Cipsectuninsaahauthalgo interface{}

    // The authentication algorithm used by the inbound encapsulation security
    // protocol (ESP) security  association of the IPsec Phase-2 Tunnel. The type
    // is AuthAlgo.
    Cipsectuninsaespauthalgo interface{}

    // The decompression algorithm used by the inbound security association of the
    // IPsec Phase-2 Tunnel. The type is CompAlgo.
    Cipsectuninsadecompalgo interface{}

    // The Diffie Hellman Group used by the outbound security association of the
    // IPsec Phase-2 Tunnel. The type is DiffHellmanGrp.
    Cipsectunoutsadiffhellmangrp interface{}

    // The encryption algorithm used by the outbound security association of the
    // IPsec Phase-2 Tunnel. The type is EncryptAlgo.
    Cipsectunoutsaencryptalgo interface{}

    // The authentication algorithm used by the outbound authentication header
    // (AH) security association of the IPsec Phase-2 Tunnel. The type is
    // AuthAlgo.
    Cipsectunoutsaahauthalgo interface{}

    // The authentication algorithm used by the inbound encapsulation security
    // protocol (ESP)  security association of the IPsec Phase-2 Tunnel. The type
    // is AuthAlgo.
    Cipsectunoutsaespauthalgo interface{}

    // The compression algorithm used by the inbound security association of the
    // IPsec Phase-2 Tunnel. The type is CompAlgo.
    Cipsectunoutsacompalgo interface{}

    // The total number of octets received by this IPsec Phase-2 Tunnel.  This
    // value is accumulated BEFORE determining whether or not the packet should be
    // decompressed.  See also cipSecTunInOctWraps for the number of times this
    // counter has wrapped. The type is interface{} with range: 0..4294967295.
    // Units are Octets.
    Cipsectuninoctets interface{}

    // A high capacity count of the total number of octets received by this IPsec
    // Phase-2 Tunnel.  This value is accumulated BEFORE determining whether or
    // not the packet should be decompressed. The type is interface{} with range:
    // 0..18446744073709551615. Units are Octets.
    Cipsectunhcinoctets interface{}

    // The number of times the octets received counter (cipSecTunInOctets) has
    // wrapped. The type is interface{} with range: 0..4294967295. Units are
    // Integral units.
    Cipsectuninoctwraps interface{}

    // The total number of decompressed octets received by this IPsec Phase-2
    // Tunnel. This value is  accumulated AFTER the packet is decompressed.  If
    // compression is not being  used, this value will match the value of  
    // cipSecTunInOctets.  See also cipSecTunInDecompOctWraps   for the number of
    // times  this counter has wrapped. The type is interface{} with range:
    // 0..4294967295. Units are Octets.
    Cipsectunindecompoctets interface{}

    // A high capacity count of the total number of decompressed octets received
    // by this IPsec Phase-2 Tunnel.  This value is accumulated AFTER the packet
    // is decompressed. If compression is not being used, this value will match
    // the value of cipSecTunHcInOctets. The type is interface{} with range:
    // 0..18446744073709551615.
    Cipsectunhcindecompoctets interface{}

    // The number of times the decompressed octets received counter 
    // (cipSecTunInDecompOctets) has wrapped. The type is interface{} with range:
    // 0..4294967295. Units are Integral units.
    Cipsectunindecompoctwraps interface{}

    // The total number of packets received by this IPsec Phase-2 Tunnel. The type
    // is interface{} with range: 0..4294967295. Units are Packets.
    Cipsectuninpkts interface{}

    // The total number of packets dropped during receive processing by this IPsec
    // Phase-2  Tunnel. This count does NOT include  packets dropped due to
    // Anti-Replay processing. The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    Cipsectunindroppkts interface{}

    // The total number of packets dropped during receive processing due to
    // Anti-Replay processing  by this IPsec Phase-2 Tunnel. The type is
    // interface{} with range: 0..4294967295. Units are Packets.
    Cipsectuninreplaydroppkts interface{}

    // The total number of inbound authentication's performed by this  IPsec
    // Phase-2 Tunnel. The type is interface{} with range: 0..4294967295. Units
    // are Events.
    Cipsectuninauths interface{}

    // The total number of inbound authentication's which ended in  failure by
    // this IPsec Phase-2 Tunnel . The type is interface{} with range:
    // 0..4294967295. Units are Failures.
    Cipsectuninauthfails interface{}

    // The total number of inbound decryption's performed by this IPsec Phase-2
    // Tunnel. The type is interface{} with range: 0..4294967295. Units are
    // Packets.
    Cipsectunindecrypts interface{}

    // The total number of inbound decryption's which ended in failure  by this
    // IPsec Phase-2 Tunnel. The type is interface{} with range: 0..4294967295.
    // Units are Failures.
    Cipsectunindecryptfails interface{}

    // The total number of octets sent by this IPsec Phase-2 Tunnel.  This value
    // is accumulated AFTER determining whether or not the packet should  be
    // compressed.  See also cipSecTunOutOctWraps for the number of times this
    // counter has wrapped. The type is interface{} with range: 0..4294967295.
    // Units are Octets.
    Cipsectunoutoctets interface{}

    // A high capacity count of the total number of octets sent by this IPsec
    // Phase-2 Tunnel.  This value is accumulated AFTER determining whether or not
    // the  packet should be compressed. The type is interface{} with range:
    // 0..18446744073709551615.
    Cipsectunhcoutoctets interface{}

    // The number of times the out octets counter (cipSecTunOutOctets) has
    // wrapped. The type is interface{} with range: 0..4294967295. Units are
    // Integral units.
    Cipsectunoutoctwraps interface{}

    // The total number of uncompressed octets sent by this IPsec Phase-2 Tunnel. 
    // This value  is accumulated BEFORE the packet is compressed.  If compression
    // is not being used, this value  will match the value of cipSecTunOutOctets. 
    // See also cipSecTunOutDecompOctWraps for the   number of times this counter
    // has wrapped. The type is interface{} with range: 0..4294967295. Units are
    // Octets.
    Cipsectunoutuncompoctets interface{}

    // A high capacity count of the total number of uncompressed octets sent by
    // this IPsec  Phase-2 Tunnel.  This value is accumulated BEFORE  the packet
    // is compressed. If compression  is not being used, this value will match the
    // value  of cipSecTunHcOutOctets. The type is interface{} with range:
    // 0..18446744073709551615.
    Cipsectunhcoutuncompoctets interface{}

    // The number of times the uncompressed octets sent counter
    // (cipSecTunOutUncompOctets) has wrapped. The type is interface{} with range:
    // 0..4294967295. Units are Integral units.
    Cipsectunoutuncompoctwraps interface{}

    // The total number of packets sent by this IPsec Phase-2 Tunnel. The type is
    // interface{} with range: 0..4294967295. Units are Packets.
    Cipsectunoutpkts interface{}

    // The total number of packets dropped during send processing by this IPsec
    // Phase-2 Tunnel. The type is interface{} with range: 0..4294967295. Units
    // are Packets.
    Cipsectunoutdroppkts interface{}

    // The total number of outbound authentication's performed by this IPsec
    // Phase-2 Tunnel. The type is interface{} with range: 0..4294967295. Units
    // are Events.
    Cipsectunoutauths interface{}

    // The total number of outbound authentication's which ended in failure  by
    // this IPsec Phase-2 Tunnel. The type is interface{} with range:
    // 0..4294967295. Units are Failures.
    Cipsectunoutauthfails interface{}

    // The total number of outbound encryption's performed by this IPsec Phase-2
    // Tunnel. The type is interface{} with range: 0..4294967295. Units are
    // Packets.
    Cipsectunoutencrypts interface{}

    // The total number of outbound encryption's which ended in failure by this
    // IPsec Phase-2 Tunnel. The type is interface{} with range: 0..4294967295.
    // Units are Failures.
    Cipsectunoutencryptfails interface{}

    // The status of the MIB table row.  This object can be used to bring the
    // tunnel down by setting value of this object to destroy(2). When the value
    // is set to destroy(2), the SA bundle is destroyed and this row is deleted
    // from this table.  When this MIB value is queried, the value of active(1) is
    // always returned, if the instance  exists.  This object cannot be used to
    // create a MIB  table row. The type is TunnelStatus.
    Cipsectunstatus interface{}
}

func (cipsectunnelentry *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable_Cipsectunnelentry) GetEntityData() *types.CommonEntityData {
    cipsectunnelentry.EntityData.YFilter = cipsectunnelentry.YFilter
    cipsectunnelentry.EntityData.YangName = "cipSecTunnelEntry"
    cipsectunnelentry.EntityData.BundleName = "cisco_ios_xe"
    cipsectunnelentry.EntityData.ParentYangName = "cipSecTunnelTable"
    cipsectunnelentry.EntityData.SegmentPath = "cipSecTunnelEntry" + "[cipSecTunIndex='" + fmt.Sprintf("%v", cipsectunnelentry.Cipsectunindex) + "']"
    cipsectunnelentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsectunnelentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsectunnelentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsectunnelentry.EntityData.Children = make(map[string]types.YChild)
    cipsectunnelentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cipsectunnelentry.EntityData.Leafs["cipSecTunIndex"] = types.YLeaf{"Cipsectunindex", cipsectunnelentry.Cipsectunindex}
    cipsectunnelentry.EntityData.Leafs["cipSecTunIkeTunnelIndex"] = types.YLeaf{"Cipsectuniketunnelindex", cipsectunnelentry.Cipsectuniketunnelindex}
    cipsectunnelentry.EntityData.Leafs["cipSecTunIkeTunnelAlive"] = types.YLeaf{"Cipsectuniketunnelalive", cipsectunnelentry.Cipsectuniketunnelalive}
    cipsectunnelentry.EntityData.Leafs["cipSecTunLocalAddr"] = types.YLeaf{"Cipsectunlocaladdr", cipsectunnelentry.Cipsectunlocaladdr}
    cipsectunnelentry.EntityData.Leafs["cipSecTunRemoteAddr"] = types.YLeaf{"Cipsectunremoteaddr", cipsectunnelentry.Cipsectunremoteaddr}
    cipsectunnelentry.EntityData.Leafs["cipSecTunKeyType"] = types.YLeaf{"Cipsectunkeytype", cipsectunnelentry.Cipsectunkeytype}
    cipsectunnelentry.EntityData.Leafs["cipSecTunEncapMode"] = types.YLeaf{"Cipsectunencapmode", cipsectunnelentry.Cipsectunencapmode}
    cipsectunnelentry.EntityData.Leafs["cipSecTunLifeSize"] = types.YLeaf{"Cipsectunlifesize", cipsectunnelentry.Cipsectunlifesize}
    cipsectunnelentry.EntityData.Leafs["cipSecTunLifeTime"] = types.YLeaf{"Cipsectunlifetime", cipsectunnelentry.Cipsectunlifetime}
    cipsectunnelentry.EntityData.Leafs["cipSecTunActiveTime"] = types.YLeaf{"Cipsectunactivetime", cipsectunnelentry.Cipsectunactivetime}
    cipsectunnelentry.EntityData.Leafs["cipSecTunSaLifeSizeThreshold"] = types.YLeaf{"Cipsectunsalifesizethreshold", cipsectunnelentry.Cipsectunsalifesizethreshold}
    cipsectunnelentry.EntityData.Leafs["cipSecTunSaLifeTimeThreshold"] = types.YLeaf{"Cipsectunsalifetimethreshold", cipsectunnelentry.Cipsectunsalifetimethreshold}
    cipsectunnelentry.EntityData.Leafs["cipSecTunTotalRefreshes"] = types.YLeaf{"Cipsectuntotalrefreshes", cipsectunnelentry.Cipsectuntotalrefreshes}
    cipsectunnelentry.EntityData.Leafs["cipSecTunExpiredSaInstances"] = types.YLeaf{"Cipsectunexpiredsainstances", cipsectunnelentry.Cipsectunexpiredsainstances}
    cipsectunnelentry.EntityData.Leafs["cipSecTunCurrentSaInstances"] = types.YLeaf{"Cipsectuncurrentsainstances", cipsectunnelentry.Cipsectuncurrentsainstances}
    cipsectunnelentry.EntityData.Leafs["cipSecTunInSaDiffHellmanGrp"] = types.YLeaf{"Cipsectuninsadiffhellmangrp", cipsectunnelentry.Cipsectuninsadiffhellmangrp}
    cipsectunnelentry.EntityData.Leafs["cipSecTunInSaEncryptAlgo"] = types.YLeaf{"Cipsectuninsaencryptalgo", cipsectunnelentry.Cipsectuninsaencryptalgo}
    cipsectunnelentry.EntityData.Leafs["cipSecTunInSaAhAuthAlgo"] = types.YLeaf{"Cipsectuninsaahauthalgo", cipsectunnelentry.Cipsectuninsaahauthalgo}
    cipsectunnelentry.EntityData.Leafs["cipSecTunInSaEspAuthAlgo"] = types.YLeaf{"Cipsectuninsaespauthalgo", cipsectunnelentry.Cipsectuninsaespauthalgo}
    cipsectunnelentry.EntityData.Leafs["cipSecTunInSaDecompAlgo"] = types.YLeaf{"Cipsectuninsadecompalgo", cipsectunnelentry.Cipsectuninsadecompalgo}
    cipsectunnelentry.EntityData.Leafs["cipSecTunOutSaDiffHellmanGrp"] = types.YLeaf{"Cipsectunoutsadiffhellmangrp", cipsectunnelentry.Cipsectunoutsadiffhellmangrp}
    cipsectunnelentry.EntityData.Leafs["cipSecTunOutSaEncryptAlgo"] = types.YLeaf{"Cipsectunoutsaencryptalgo", cipsectunnelentry.Cipsectunoutsaencryptalgo}
    cipsectunnelentry.EntityData.Leafs["cipSecTunOutSaAhAuthAlgo"] = types.YLeaf{"Cipsectunoutsaahauthalgo", cipsectunnelentry.Cipsectunoutsaahauthalgo}
    cipsectunnelentry.EntityData.Leafs["cipSecTunOutSaEspAuthAlgo"] = types.YLeaf{"Cipsectunoutsaespauthalgo", cipsectunnelentry.Cipsectunoutsaespauthalgo}
    cipsectunnelentry.EntityData.Leafs["cipSecTunOutSaCompAlgo"] = types.YLeaf{"Cipsectunoutsacompalgo", cipsectunnelentry.Cipsectunoutsacompalgo}
    cipsectunnelentry.EntityData.Leafs["cipSecTunInOctets"] = types.YLeaf{"Cipsectuninoctets", cipsectunnelentry.Cipsectuninoctets}
    cipsectunnelentry.EntityData.Leafs["cipSecTunHcInOctets"] = types.YLeaf{"Cipsectunhcinoctets", cipsectunnelentry.Cipsectunhcinoctets}
    cipsectunnelentry.EntityData.Leafs["cipSecTunInOctWraps"] = types.YLeaf{"Cipsectuninoctwraps", cipsectunnelentry.Cipsectuninoctwraps}
    cipsectunnelentry.EntityData.Leafs["cipSecTunInDecompOctets"] = types.YLeaf{"Cipsectunindecompoctets", cipsectunnelentry.Cipsectunindecompoctets}
    cipsectunnelentry.EntityData.Leafs["cipSecTunHcInDecompOctets"] = types.YLeaf{"Cipsectunhcindecompoctets", cipsectunnelentry.Cipsectunhcindecompoctets}
    cipsectunnelentry.EntityData.Leafs["cipSecTunInDecompOctWraps"] = types.YLeaf{"Cipsectunindecompoctwraps", cipsectunnelentry.Cipsectunindecompoctwraps}
    cipsectunnelentry.EntityData.Leafs["cipSecTunInPkts"] = types.YLeaf{"Cipsectuninpkts", cipsectunnelentry.Cipsectuninpkts}
    cipsectunnelentry.EntityData.Leafs["cipSecTunInDropPkts"] = types.YLeaf{"Cipsectunindroppkts", cipsectunnelentry.Cipsectunindroppkts}
    cipsectunnelentry.EntityData.Leafs["cipSecTunInReplayDropPkts"] = types.YLeaf{"Cipsectuninreplaydroppkts", cipsectunnelentry.Cipsectuninreplaydroppkts}
    cipsectunnelentry.EntityData.Leafs["cipSecTunInAuths"] = types.YLeaf{"Cipsectuninauths", cipsectunnelentry.Cipsectuninauths}
    cipsectunnelentry.EntityData.Leafs["cipSecTunInAuthFails"] = types.YLeaf{"Cipsectuninauthfails", cipsectunnelentry.Cipsectuninauthfails}
    cipsectunnelentry.EntityData.Leafs["cipSecTunInDecrypts"] = types.YLeaf{"Cipsectunindecrypts", cipsectunnelentry.Cipsectunindecrypts}
    cipsectunnelentry.EntityData.Leafs["cipSecTunInDecryptFails"] = types.YLeaf{"Cipsectunindecryptfails", cipsectunnelentry.Cipsectunindecryptfails}
    cipsectunnelentry.EntityData.Leafs["cipSecTunOutOctets"] = types.YLeaf{"Cipsectunoutoctets", cipsectunnelentry.Cipsectunoutoctets}
    cipsectunnelentry.EntityData.Leafs["cipSecTunHcOutOctets"] = types.YLeaf{"Cipsectunhcoutoctets", cipsectunnelentry.Cipsectunhcoutoctets}
    cipsectunnelentry.EntityData.Leafs["cipSecTunOutOctWraps"] = types.YLeaf{"Cipsectunoutoctwraps", cipsectunnelentry.Cipsectunoutoctwraps}
    cipsectunnelentry.EntityData.Leafs["cipSecTunOutUncompOctets"] = types.YLeaf{"Cipsectunoutuncompoctets", cipsectunnelentry.Cipsectunoutuncompoctets}
    cipsectunnelentry.EntityData.Leafs["cipSecTunHcOutUncompOctets"] = types.YLeaf{"Cipsectunhcoutuncompoctets", cipsectunnelentry.Cipsectunhcoutuncompoctets}
    cipsectunnelentry.EntityData.Leafs["cipSecTunOutUncompOctWraps"] = types.YLeaf{"Cipsectunoutuncompoctwraps", cipsectunnelentry.Cipsectunoutuncompoctwraps}
    cipsectunnelentry.EntityData.Leafs["cipSecTunOutPkts"] = types.YLeaf{"Cipsectunoutpkts", cipsectunnelentry.Cipsectunoutpkts}
    cipsectunnelentry.EntityData.Leafs["cipSecTunOutDropPkts"] = types.YLeaf{"Cipsectunoutdroppkts", cipsectunnelentry.Cipsectunoutdroppkts}
    cipsectunnelentry.EntityData.Leafs["cipSecTunOutAuths"] = types.YLeaf{"Cipsectunoutauths", cipsectunnelentry.Cipsectunoutauths}
    cipsectunnelentry.EntityData.Leafs["cipSecTunOutAuthFails"] = types.YLeaf{"Cipsectunoutauthfails", cipsectunnelentry.Cipsectunoutauthfails}
    cipsectunnelentry.EntityData.Leafs["cipSecTunOutEncrypts"] = types.YLeaf{"Cipsectunoutencrypts", cipsectunnelentry.Cipsectunoutencrypts}
    cipsectunnelentry.EntityData.Leafs["cipSecTunOutEncryptFails"] = types.YLeaf{"Cipsectunoutencryptfails", cipsectunnelentry.Cipsectunoutencryptfails}
    cipsectunnelentry.EntityData.Leafs["cipSecTunStatus"] = types.YLeaf{"Cipsectunstatus", cipsectunnelentry.Cipsectunstatus}
    return &(cipsectunnelentry.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cipsecendpttable
// The IPsec Phase-2 Tunnel Endpoint Table.
// This table contains an entry for each 
// active endpoint associated with an IPsec
//  Phase-2 Tunnel.
type CISCOIPSECFLOWMONITORMIB_Cipsecendpttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An IPsec Phase-2 Tunnel Endpoint entry. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_Cipsecendpttable_Cipsecendptentry.
    Cipsecendptentry []CISCOIPSECFLOWMONITORMIB_Cipsecendpttable_Cipsecendptentry
}

func (cipsecendpttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable) GetEntityData() *types.CommonEntityData {
    cipsecendpttable.EntityData.YFilter = cipsecendpttable.YFilter
    cipsecendpttable.EntityData.YangName = "cipSecEndPtTable"
    cipsecendpttable.EntityData.BundleName = "cisco_ios_xe"
    cipsecendpttable.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cipsecendpttable.EntityData.SegmentPath = "cipSecEndPtTable"
    cipsecendpttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsecendpttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsecendpttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsecendpttable.EntityData.Children = make(map[string]types.YChild)
    cipsecendpttable.EntityData.Children["cipSecEndPtEntry"] = types.YChild{"Cipsecendptentry", nil}
    for i := range cipsecendpttable.Cipsecendptentry {
        cipsecendpttable.EntityData.Children[types.GetSegmentPath(&cipsecendpttable.Cipsecendptentry[i])] = types.YChild{"Cipsecendptentry", &cipsecendpttable.Cipsecendptentry[i]}
    }
    cipsecendpttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cipsecendpttable.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cipsecendpttable_Cipsecendptentry
// An IPsec Phase-2 Tunnel Endpoint entry.
type CISCOIPSECFLOWMONITORMIB_Cipsecendpttable_Cipsecendptentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_ipsec_flow_monitor_mib.CISCOIPSECFLOWMONITORMIB_Cipsectunneltable_Cipsectunnelentry_Cipsectunindex
    Cipsectunindex interface{}

    // This attribute is a key. The number of the Endpoint associated with the
    // IPsec Phase-2 Tunnel Table.  The value of this index is a number which
    // begins at one and  is incremented with each Endpoint associated  with an
    // IPsec Phase-2 Tunnel. The value of this object will wrap at 2,147,483,647.
    // The type is interface{} with range: 1..2147483647.
    Cipsecendptindex interface{}

    // The DNS name of the local Endpoint. The type is string.
    Cipsecendptlocalname interface{}

    // The type of identity for the local Endpoint. Possible values are: 1) a
    // single IP address, or 2) an IP address range, or 3) an IP subnet. The type
    // is EndPtType.
    Cipsecendptlocaltype interface{}

    // The local Endpoint's first IP address specification.  If the local Endpoint
    // type is single IP address,  then this is the value of the IP address.  If
    // the local Endpoint type is IP subnet, then this is the value of the subnet.
    // If the local Endpoint type is IP address range,  then this is the value of
    // beginning IP address  of the range. The type is string with length: 4 | 16.
    Cipsecendptlocaladdr1 interface{}

    // The local Endpoint's second IP address specification.  If the local
    // Endpoint type is single IP address,  then this is the value of the IP
    // address.  If the local Endpoint type is IP subnet, then this is the value
    // of the subnet mask.  If the local Endpoint type is IP address range,  then
    // this is the value of ending IP address  of the range. The type is string
    // with length: 4 | 16.
    Cipsecendptlocaladdr2 interface{}

    // The protocol number of the local Endpoint's traffic. The type is
    // interface{} with range: 0..255.
    Cipsecendptlocalprotocol interface{}

    // The port number of the local Endpoint's traffic. The type is interface{}
    // with range: 0..65535.
    Cipsecendptlocalport interface{}

    // The DNS name of the remote Endpoint. The type is string.
    Cipsecendptremotename interface{}

    // The type of identity for the remote Endpoint. Possible values are: 1) a
    // single IP address, or 2) an IP address range, or 3) an IP subnet. The type
    // is EndPtType.
    Cipsecendptremotetype interface{}

    // The remote Endpoint's first IP address specification.  If the remote
    // Endpoint type is single IP address,  then this is the value of the IP
    // address.  If the remote Endpoint type is IP subnet, then this is the value
    // of the subnet.  If the remote Endpoint type is IP address range,  then this
    // is the value of beginning IP address  of the range. The type is string with
    // length: 4 | 16.
    Cipsecendptremoteaddr1 interface{}

    // The remote Endpoint's second IP address specification.  If the remote
    // Endpoint type is single IP address,  then this is the value of the IP
    // address.  If the remote Endpoint type is IP subnet, then this is the value
    // of the subnet mask.  If the remote Endpoint type is IP address range,  then
    // this is the value of ending IP address of  the range. The type is string
    // with length: 4 | 16.
    Cipsecendptremoteaddr2 interface{}

    // The protocol number of the remote Endpoint's traffic. The type is
    // interface{} with range: 0..255.
    Cipsecendptremoteprotocol interface{}

    // The port number of the remote Endpoint's traffic. The type is interface{}
    // with range: 0..65535.
    Cipsecendptremoteport interface{}
}

func (cipsecendptentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable_Cipsecendptentry) GetEntityData() *types.CommonEntityData {
    cipsecendptentry.EntityData.YFilter = cipsecendptentry.YFilter
    cipsecendptentry.EntityData.YangName = "cipSecEndPtEntry"
    cipsecendptentry.EntityData.BundleName = "cisco_ios_xe"
    cipsecendptentry.EntityData.ParentYangName = "cipSecEndPtTable"
    cipsecendptentry.EntityData.SegmentPath = "cipSecEndPtEntry" + "[cipSecTunIndex='" + fmt.Sprintf("%v", cipsecendptentry.Cipsectunindex) + "']" + "[cipSecEndPtIndex='" + fmt.Sprintf("%v", cipsecendptentry.Cipsecendptindex) + "']"
    cipsecendptentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsecendptentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsecendptentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsecendptentry.EntityData.Children = make(map[string]types.YChild)
    cipsecendptentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cipsecendptentry.EntityData.Leafs["cipSecTunIndex"] = types.YLeaf{"Cipsectunindex", cipsecendptentry.Cipsectunindex}
    cipsecendptentry.EntityData.Leafs["cipSecEndPtIndex"] = types.YLeaf{"Cipsecendptindex", cipsecendptentry.Cipsecendptindex}
    cipsecendptentry.EntityData.Leafs["cipSecEndPtLocalName"] = types.YLeaf{"Cipsecendptlocalname", cipsecendptentry.Cipsecendptlocalname}
    cipsecendptentry.EntityData.Leafs["cipSecEndPtLocalType"] = types.YLeaf{"Cipsecendptlocaltype", cipsecendptentry.Cipsecendptlocaltype}
    cipsecendptentry.EntityData.Leafs["cipSecEndPtLocalAddr1"] = types.YLeaf{"Cipsecendptlocaladdr1", cipsecendptentry.Cipsecendptlocaladdr1}
    cipsecendptentry.EntityData.Leafs["cipSecEndPtLocalAddr2"] = types.YLeaf{"Cipsecendptlocaladdr2", cipsecendptentry.Cipsecendptlocaladdr2}
    cipsecendptentry.EntityData.Leafs["cipSecEndPtLocalProtocol"] = types.YLeaf{"Cipsecendptlocalprotocol", cipsecendptentry.Cipsecendptlocalprotocol}
    cipsecendptentry.EntityData.Leafs["cipSecEndPtLocalPort"] = types.YLeaf{"Cipsecendptlocalport", cipsecendptentry.Cipsecendptlocalport}
    cipsecendptentry.EntityData.Leafs["cipSecEndPtRemoteName"] = types.YLeaf{"Cipsecendptremotename", cipsecendptentry.Cipsecendptremotename}
    cipsecendptentry.EntityData.Leafs["cipSecEndPtRemoteType"] = types.YLeaf{"Cipsecendptremotetype", cipsecendptentry.Cipsecendptremotetype}
    cipsecendptentry.EntityData.Leafs["cipSecEndPtRemoteAddr1"] = types.YLeaf{"Cipsecendptremoteaddr1", cipsecendptentry.Cipsecendptremoteaddr1}
    cipsecendptentry.EntityData.Leafs["cipSecEndPtRemoteAddr2"] = types.YLeaf{"Cipsecendptremoteaddr2", cipsecendptentry.Cipsecendptremoteaddr2}
    cipsecendptentry.EntityData.Leafs["cipSecEndPtRemoteProtocol"] = types.YLeaf{"Cipsecendptremoteprotocol", cipsecendptentry.Cipsecendptremoteprotocol}
    cipsecendptentry.EntityData.Leafs["cipSecEndPtRemotePort"] = types.YLeaf{"Cipsecendptremoteport", cipsecendptentry.Cipsecendptremoteport}
    return &(cipsecendptentry.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cipsecspitable
// The IPsec Phase-2 Security Protection Index Table.
// This table contains an entry for each active 
// and expiring security
//  association.
type CISCOIPSECFLOWMONITORMIB_Cipsecspitable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with active and expiring
    // IPsec Phase-2  security associations. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry.
    Cipsecspientry []CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry
}

func (cipsecspitable *CISCOIPSECFLOWMONITORMIB_Cipsecspitable) GetEntityData() *types.CommonEntityData {
    cipsecspitable.EntityData.YFilter = cipsecspitable.YFilter
    cipsecspitable.EntityData.YangName = "cipSecSpiTable"
    cipsecspitable.EntityData.BundleName = "cisco_ios_xe"
    cipsecspitable.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cipsecspitable.EntityData.SegmentPath = "cipSecSpiTable"
    cipsecspitable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsecspitable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsecspitable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsecspitable.EntityData.Children = make(map[string]types.YChild)
    cipsecspitable.EntityData.Children["cipSecSpiEntry"] = types.YChild{"Cipsecspientry", nil}
    for i := range cipsecspitable.Cipsecspientry {
        cipsecspitable.EntityData.Children[types.GetSegmentPath(&cipsecspitable.Cipsecspientry[i])] = types.YChild{"Cipsecspientry", &cipsecspitable.Cipsecspientry[i]}
    }
    cipsecspitable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cipsecspitable.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry
// Each entry contains the attributes associated with
// active and expiring IPsec Phase-2 
// security associations.
type CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_ipsec_flow_monitor_mib.CISCOIPSECFLOWMONITORMIB_Cipsectunneltable_Cipsectunnelentry_Cipsectunindex
    Cipsectunindex interface{}

    // This attribute is a key. The number of the SPI associated with the Phase-2
    // Tunnel Table.  The value of this  index is a number which begins at one and
    // is  incremented with each SPI associated with an  IPsec Phase-2 Tunnel. 
    // The value of this  object will wrap at 2,147,483,647. The type is
    // interface{} with range: 1..2147483647.
    Cipsecspiindex interface{}

    // The direction of the SPI. The type is Cipsecspidirection.
    Cipsecspidirection interface{}

    // The value of the SPI. The type is interface{} with range: 1..4294967295.
    Cipsecspivalue interface{}

    // The protocol of the SPI. The type is Cipsecspiprotocol.
    Cipsecspiprotocol interface{}

    // The status of the SPI. The type is Cipsecspistatus.
    Cipsecspistatus interface{}
}

func (cipsecspientry *CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry) GetEntityData() *types.CommonEntityData {
    cipsecspientry.EntityData.YFilter = cipsecspientry.YFilter
    cipsecspientry.EntityData.YangName = "cipSecSpiEntry"
    cipsecspientry.EntityData.BundleName = "cisco_ios_xe"
    cipsecspientry.EntityData.ParentYangName = "cipSecSpiTable"
    cipsecspientry.EntityData.SegmentPath = "cipSecSpiEntry" + "[cipSecTunIndex='" + fmt.Sprintf("%v", cipsecspientry.Cipsectunindex) + "']" + "[cipSecSpiIndex='" + fmt.Sprintf("%v", cipsecspientry.Cipsecspiindex) + "']"
    cipsecspientry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsecspientry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsecspientry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsecspientry.EntityData.Children = make(map[string]types.YChild)
    cipsecspientry.EntityData.Leafs = make(map[string]types.YLeaf)
    cipsecspientry.EntityData.Leafs["cipSecTunIndex"] = types.YLeaf{"Cipsectunindex", cipsecspientry.Cipsectunindex}
    cipsecspientry.EntityData.Leafs["cipSecSpiIndex"] = types.YLeaf{"Cipsecspiindex", cipsecspientry.Cipsecspiindex}
    cipsecspientry.EntityData.Leafs["cipSecSpiDirection"] = types.YLeaf{"Cipsecspidirection", cipsecspientry.Cipsecspidirection}
    cipsecspientry.EntityData.Leafs["cipSecSpiValue"] = types.YLeaf{"Cipsecspivalue", cipsecspientry.Cipsecspivalue}
    cipsecspientry.EntityData.Leafs["cipSecSpiProtocol"] = types.YLeaf{"Cipsecspiprotocol", cipsecspientry.Cipsecspiprotocol}
    cipsecspientry.EntityData.Leafs["cipSecSpiStatus"] = types.YLeaf{"Cipsecspistatus", cipsecspientry.Cipsecspistatus}
    return &(cipsecspientry.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry_Cipsecspidirection represents The direction of the SPI.
type CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry_Cipsecspidirection string

const (
    CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry_Cipsecspidirection_in CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry_Cipsecspidirection = "in"

    CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry_Cipsecspidirection_out CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry_Cipsecspidirection = "out"
)

// CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry_Cipsecspiprotocol represents The protocol of the SPI.
type CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry_Cipsecspiprotocol string

const (
    CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry_Cipsecspiprotocol_ah CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry_Cipsecspiprotocol = "ah"

    CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry_Cipsecspiprotocol_esp CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry_Cipsecspiprotocol = "esp"

    CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry_Cipsecspiprotocol_ipcomp CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry_Cipsecspiprotocol = "ipcomp"
)

// CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry_Cipsecspistatus represents The status of the SPI.
type CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry_Cipsecspistatus string

const (
    CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry_Cipsecspistatus_active CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry_Cipsecspistatus = "active"

    CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry_Cipsecspistatus_expiring CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry_Cipsecspistatus = "expiring"
)

// CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable
// Phase-2 IPsec stats information is included in this table.
// Each entry is related to a specific gateway which is 
// identified by 'cmgwIndex'
type CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes of an Phase-2 IPsec stats information
    // for the related gateway.  There is only one entry for each gateway. The
    // entry  is created when a gateway up and cannot be deleted. The type is
    // slice of
    // CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable_Cipsecphase2Gwstatsentry.
    Cipsecphase2Gwstatsentry []CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable_Cipsecphase2Gwstatsentry
}

func (cipsecphase2Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable) GetEntityData() *types.CommonEntityData {
    cipsecphase2Gwstatstable.EntityData.YFilter = cipsecphase2Gwstatstable.YFilter
    cipsecphase2Gwstatstable.EntityData.YangName = "cipSecPhase2GWStatsTable"
    cipsecphase2Gwstatstable.EntityData.BundleName = "cisco_ios_xe"
    cipsecphase2Gwstatstable.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cipsecphase2Gwstatstable.EntityData.SegmentPath = "cipSecPhase2GWStatsTable"
    cipsecphase2Gwstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsecphase2Gwstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsecphase2Gwstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsecphase2Gwstatstable.EntityData.Children = make(map[string]types.YChild)
    cipsecphase2Gwstatstable.EntityData.Children["cipSecPhase2GWStatsEntry"] = types.YChild{"Cipsecphase2Gwstatsentry", nil}
    for i := range cipsecphase2Gwstatstable.Cipsecphase2Gwstatsentry {
        cipsecphase2Gwstatstable.EntityData.Children[types.GetSegmentPath(&cipsecphase2Gwstatstable.Cipsecphase2Gwstatsentry[i])] = types.YChild{"Cipsecphase2Gwstatsentry", &cipsecphase2Gwstatstable.Cipsecphase2Gwstatsentry[i]}
    }
    cipsecphase2Gwstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cipsecphase2Gwstatstable.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable_Cipsecphase2Gwstatsentry
// Each entry contains the attributes of an Phase-2 IPsec stats
// information for the related gateway.
// 
// There is only one entry for each gateway. The entry 
// is created when a gateway up and cannot be deleted.
type CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable_Cipsecphase2Gwstatsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // cisco_media_gateway_mib.CISCOMEDIAGATEWAYMIB_Cmediagwtable_Cmediagwentry_Cmgwindex
    Cmgwindex interface{}

    // The total number of currently active IPsec Phase-2 Tunnels. The type is
    // interface{} with range: 0..4294967295.
    Cipsecphase2Gwactivetunnels interface{}

    // The total number of previously active IPsec Phase-2 Tunnels. The type is
    // interface{} with range: 0..4294967295. Units are Phase-2 Tunnels.
    Cipsecphase2Gwprevioustunnels interface{}

    // The total number of octets received by all current and previous IPsec
    // Phase-2 Tunnels.  This value is accumulated BEFORE determining  whether or
    // not the packet should be decompressed.  See also cipSecGlobalInOctWraps for
    // the number of times this counter has wrapped. The type is interface{} with
    // range: 0..4294967295. Units are Octets.
    Cipsecphase2Gwinoctets interface{}

    // The number of times the global octets received counter
    // (cipSecGlobalInOctets) has wrapped. The type is interface{} with range:
    // 0..4294967295. Units are Integral units.
    Cipsecphase2Gwinoctwraps interface{}

    // The total number of decompressed octets received by all current and
    // previous IPsec Phase-2 Tunnels.   This value is accumulated AFTER the
    // packet is  decompressed. If compression is not being used,  this value will
    // match the value of cipSecGlobalInOctets.  See also
    // cipSecGlobalInDecompOctWraps for the number of times this counter has
    // wrapped. The type is interface{} with range: 0..4294967295. Units are
    // Octets.
    Cipsecphase2Gwindecompoctets interface{}

    // The number of times the global decompressed octets received counter
    // (cipSecGlobalInDecompOctets)  has wrapped. The type is interface{} with
    // range: 0..4294967295. Units are Integral units.
    Cipsecphase2Gwindecompoctwraps interface{}

    // The total number of packets received by all current and previous IPsec
    // Phase-2 Tunnels. The type is interface{} with range: 0..4294967295. Units
    // are Packets.
    Cipsecphase2Gwinpkts interface{}

    // The total number of packets dropped during receive processing by all
    // current and previous  IPsec Phase-2 Tunnels. This count does NOT include 
    // packets dropped due to Anti-Replay processing. The type is interface{} with
    // range: 0..4294967295. Units are Packets.
    Cipsecphase2Gwindrops interface{}

    // The total number of packets dropped during receive processing due to
    // Anti-Replay  processing by all current and previous IPsec Phase-2 Tunnels.
    // The type is interface{} with range: 0..4294967295. Units are Packets.
    Cipsecphase2Gwinreplaydrops interface{}

    // The total number of inbound authentication's performed by all current and
    // previous IPsec  Phase-2 Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Events.
    Cipsecphase2Gwinauths interface{}

    // The total number of inbound authentication's which ended in failure by all
    // current and previous  IPsec Phase-2 Tunnels. The type is interface{} with
    // range: 0..4294967295. Units are Failures.
    Cipsecphase2Gwinauthfails interface{}

    // The total number of inbound decryption's performed by all current and
    // previous IPsec  Phase-2 Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    Cipsecphase2Gwindecrypts interface{}

    // The total number of inbound decryption's which ended in failure by all
    // current and  previous IPsec Phase-2 Tunnels. The type is interface{} with
    // range: 0..4294967295. Units are Packets.
    Cipsecphase2Gwindecryptfails interface{}

    // The total number of octets sent by all current and previous IPsec Phase-2
    // Tunnels.   This value is accumulated AFTER determining  whether or not the
    // packet should be compressed.   See also cipSecGlobalOutOctWraps for the
    // number of times this counter has wrapped. The type is interface{} with
    // range: 0..4294967295. Units are Octets.
    Cipsecphase2Gwoutoctets interface{}

    // The number of times the global octets sent counter (cipSecGlobalOutOctets)
    // has wrapped. The type is interface{} with range: 0..4294967295. Units are
    // Integral units.
    Cipsecphase2Gwoutoctwraps interface{}

    // The total number of uncompressed octets sent by all current and previous
    // IPsec Phase-2 Tunnels.   This value is accumulated BEFORE the packet is 
    // compressed. If compression is not being used, this  value will match the
    // value of cipSecGlobalOutOctets.  See also cipSecGlobalOutDecompOctWraps for
    // the number  of times this counter has wrapped. The type is interface{} with
    // range: 0..4294967295. Units are Octets.
    Cipsecphase2Gwoutuncompoctets interface{}

    // The number of times the global uncompressed octets sent counter
    // (cipSecGlobalOutUncompOctets)  has wrapped. The type is interface{} with
    // range: 0..4294967295. Units are Integral units.
    Cipsecphase2Gwoutuncompoctwraps interface{}

    // The total number of packets sent by all current and previous IPsec Phase-2 
    // Tunnels. The type is interface{} with range: 0..4294967295. Units are
    // Packets.
    Cipsecphase2Gwoutpkts interface{}

    // The total number of packets dropped during send processing by all current
    // and previous IPsec  Phase-2 Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    Cipsecphase2Gwoutdrops interface{}

    // The total number of outbound authentication's performed by all current and
    // previous IPsec  Phase-2 Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Events.
    Cipsecphase2Gwoutauths interface{}

    // The total number of outbound authentication's which ended in failure by all
    // current and previous IPsec Phase-2 Tunnels. The type is interface{} with
    // range: 0..4294967295. Units are Failures.
    Cipsecphase2Gwoutauthfails interface{}

    // The total number of outbound encryption's performed by all current and
    // previous IPsec Phase-2 Tunnels. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    Cipsecphase2Gwoutencrypts interface{}

    // The total number of outbound encryption's which ended in failure by all
    // current and  previous IPsec Phase-2 Tunnels. The type is interface{} with
    // range: 0..4294967295. Units are Failures.
    Cipsecphase2Gwoutencryptfails interface{}

    // The total number of protocol use failures which occurred during processing
    // of all current  and previously active IPsec Phase-2 Tunnels. The type is
    // interface{} with range: 0..4294967295. Units are Failures.
    Cipsecphase2Gwprotocolusefails interface{}

    // The total number of non-existent Security Association in failures which
    // occurred  during processing of all current and previous IPsec Phase-2
    // Tunnels. The type is interface{} with range: 0..4294967295. Units are
    // Failures.
    Cipsecphase2Gwnosafails interface{}

    // The total number of system capacity failures which occurred during
    // processing of all current  and previously active IPsec Phase-2 Tunnels. The
    // type is interface{} with range: 0..4294967295. Units are Failures.
    Cipsecphase2Gwsyscapfails interface{}
}

func (cipsecphase2Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable_Cipsecphase2Gwstatsentry) GetEntityData() *types.CommonEntityData {
    cipsecphase2Gwstatsentry.EntityData.YFilter = cipsecphase2Gwstatsentry.YFilter
    cipsecphase2Gwstatsentry.EntityData.YangName = "cipSecPhase2GWStatsEntry"
    cipsecphase2Gwstatsentry.EntityData.BundleName = "cisco_ios_xe"
    cipsecphase2Gwstatsentry.EntityData.ParentYangName = "cipSecPhase2GWStatsTable"
    cipsecphase2Gwstatsentry.EntityData.SegmentPath = "cipSecPhase2GWStatsEntry" + "[cmgwIndex='" + fmt.Sprintf("%v", cipsecphase2Gwstatsentry.Cmgwindex) + "']"
    cipsecphase2Gwstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsecphase2Gwstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsecphase2Gwstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsecphase2Gwstatsentry.EntityData.Children = make(map[string]types.YChild)
    cipsecphase2Gwstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cipsecphase2Gwstatsentry.EntityData.Leafs["cmgwIndex"] = types.YLeaf{"Cmgwindex", cipsecphase2Gwstatsentry.Cmgwindex}
    cipsecphase2Gwstatsentry.EntityData.Leafs["cipSecPhase2GWActiveTunnels"] = types.YLeaf{"Cipsecphase2Gwactivetunnels", cipsecphase2Gwstatsentry.Cipsecphase2Gwactivetunnels}
    cipsecphase2Gwstatsentry.EntityData.Leafs["cipSecPhase2GWPreviousTunnels"] = types.YLeaf{"Cipsecphase2Gwprevioustunnels", cipsecphase2Gwstatsentry.Cipsecphase2Gwprevioustunnels}
    cipsecphase2Gwstatsentry.EntityData.Leafs["cipSecPhase2GWInOctets"] = types.YLeaf{"Cipsecphase2Gwinoctets", cipsecphase2Gwstatsentry.Cipsecphase2Gwinoctets}
    cipsecphase2Gwstatsentry.EntityData.Leafs["cipSecPhase2GWInOctWraps"] = types.YLeaf{"Cipsecphase2Gwinoctwraps", cipsecphase2Gwstatsentry.Cipsecphase2Gwinoctwraps}
    cipsecphase2Gwstatsentry.EntityData.Leafs["cipSecPhase2GWInDecompOctets"] = types.YLeaf{"Cipsecphase2Gwindecompoctets", cipsecphase2Gwstatsentry.Cipsecphase2Gwindecompoctets}
    cipsecphase2Gwstatsentry.EntityData.Leafs["cipSecPhase2GWInDecompOctWraps"] = types.YLeaf{"Cipsecphase2Gwindecompoctwraps", cipsecphase2Gwstatsentry.Cipsecphase2Gwindecompoctwraps}
    cipsecphase2Gwstatsentry.EntityData.Leafs["cipSecPhase2GWInPkts"] = types.YLeaf{"Cipsecphase2Gwinpkts", cipsecphase2Gwstatsentry.Cipsecphase2Gwinpkts}
    cipsecphase2Gwstatsentry.EntityData.Leafs["cipSecPhase2GWInDrops"] = types.YLeaf{"Cipsecphase2Gwindrops", cipsecphase2Gwstatsentry.Cipsecphase2Gwindrops}
    cipsecphase2Gwstatsentry.EntityData.Leafs["cipSecPhase2GWInReplayDrops"] = types.YLeaf{"Cipsecphase2Gwinreplaydrops", cipsecphase2Gwstatsentry.Cipsecphase2Gwinreplaydrops}
    cipsecphase2Gwstatsentry.EntityData.Leafs["cipSecPhase2GWInAuths"] = types.YLeaf{"Cipsecphase2Gwinauths", cipsecphase2Gwstatsentry.Cipsecphase2Gwinauths}
    cipsecphase2Gwstatsentry.EntityData.Leafs["cipSecPhase2GWInAuthFails"] = types.YLeaf{"Cipsecphase2Gwinauthfails", cipsecphase2Gwstatsentry.Cipsecphase2Gwinauthfails}
    cipsecphase2Gwstatsentry.EntityData.Leafs["cipSecPhase2GWInDecrypts"] = types.YLeaf{"Cipsecphase2Gwindecrypts", cipsecphase2Gwstatsentry.Cipsecphase2Gwindecrypts}
    cipsecphase2Gwstatsentry.EntityData.Leafs["cipSecPhase2GWInDecryptFails"] = types.YLeaf{"Cipsecphase2Gwindecryptfails", cipsecphase2Gwstatsentry.Cipsecphase2Gwindecryptfails}
    cipsecphase2Gwstatsentry.EntityData.Leafs["cipSecPhase2GWOutOctets"] = types.YLeaf{"Cipsecphase2Gwoutoctets", cipsecphase2Gwstatsentry.Cipsecphase2Gwoutoctets}
    cipsecphase2Gwstatsentry.EntityData.Leafs["cipSecPhase2GWOutOctWraps"] = types.YLeaf{"Cipsecphase2Gwoutoctwraps", cipsecphase2Gwstatsentry.Cipsecphase2Gwoutoctwraps}
    cipsecphase2Gwstatsentry.EntityData.Leafs["cipSecPhase2GWOutUncompOctets"] = types.YLeaf{"Cipsecphase2Gwoutuncompoctets", cipsecphase2Gwstatsentry.Cipsecphase2Gwoutuncompoctets}
    cipsecphase2Gwstatsentry.EntityData.Leafs["cipSecPhase2GWOutUncompOctWraps"] = types.YLeaf{"Cipsecphase2Gwoutuncompoctwraps", cipsecphase2Gwstatsentry.Cipsecphase2Gwoutuncompoctwraps}
    cipsecphase2Gwstatsentry.EntityData.Leafs["cipSecPhase2GWOutPkts"] = types.YLeaf{"Cipsecphase2Gwoutpkts", cipsecphase2Gwstatsentry.Cipsecphase2Gwoutpkts}
    cipsecphase2Gwstatsentry.EntityData.Leafs["cipSecPhase2GWOutDrops"] = types.YLeaf{"Cipsecphase2Gwoutdrops", cipsecphase2Gwstatsentry.Cipsecphase2Gwoutdrops}
    cipsecphase2Gwstatsentry.EntityData.Leafs["cipSecPhase2GWOutAuths"] = types.YLeaf{"Cipsecphase2Gwoutauths", cipsecphase2Gwstatsentry.Cipsecphase2Gwoutauths}
    cipsecphase2Gwstatsentry.EntityData.Leafs["cipSecPhase2GWOutAuthFails"] = types.YLeaf{"Cipsecphase2Gwoutauthfails", cipsecphase2Gwstatsentry.Cipsecphase2Gwoutauthfails}
    cipsecphase2Gwstatsentry.EntityData.Leafs["cipSecPhase2GWOutEncrypts"] = types.YLeaf{"Cipsecphase2Gwoutencrypts", cipsecphase2Gwstatsentry.Cipsecphase2Gwoutencrypts}
    cipsecphase2Gwstatsentry.EntityData.Leafs["cipSecPhase2GWOutEncryptFails"] = types.YLeaf{"Cipsecphase2Gwoutencryptfails", cipsecphase2Gwstatsentry.Cipsecphase2Gwoutencryptfails}
    cipsecphase2Gwstatsentry.EntityData.Leafs["cipSecPhase2GWProtocolUseFails"] = types.YLeaf{"Cipsecphase2Gwprotocolusefails", cipsecphase2Gwstatsentry.Cipsecphase2Gwprotocolusefails}
    cipsecphase2Gwstatsentry.EntityData.Leafs["cipSecPhase2GWNoSaFails"] = types.YLeaf{"Cipsecphase2Gwnosafails", cipsecphase2Gwstatsentry.Cipsecphase2Gwnosafails}
    cipsecphase2Gwstatsentry.EntityData.Leafs["cipSecPhase2GWSysCapFails"] = types.YLeaf{"Cipsecphase2Gwsyscapfails", cipsecphase2Gwstatsentry.Cipsecphase2Gwsyscapfails}
    return &(cipsecphase2Gwstatsentry.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable
// The IPsec Phase-1 Internet Key Exchange Tunnel
// History Table.  This table is implemented as a 
// sliding window in which only the last n entries 
// are maintained.  The maximum number of entries
//  is specified by the cipSecHistTableSize object.
type CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with a previously active
    // IPsec  Phase-1 IKE Tunnel. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry.
    Ciketunnelhistentry []CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry
}

func (ciketunnelhisttable *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable) GetEntityData() *types.CommonEntityData {
    ciketunnelhisttable.EntityData.YFilter = ciketunnelhisttable.YFilter
    ciketunnelhisttable.EntityData.YangName = "cikeTunnelHistTable"
    ciketunnelhisttable.EntityData.BundleName = "cisco_ios_xe"
    ciketunnelhisttable.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    ciketunnelhisttable.EntityData.SegmentPath = "cikeTunnelHistTable"
    ciketunnelhisttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciketunnelhisttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciketunnelhisttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciketunnelhisttable.EntityData.Children = make(map[string]types.YChild)
    ciketunnelhisttable.EntityData.Children["cikeTunnelHistEntry"] = types.YChild{"Ciketunnelhistentry", nil}
    for i := range ciketunnelhisttable.Ciketunnelhistentry {
        ciketunnelhisttable.EntityData.Children[types.GetSegmentPath(&ciketunnelhisttable.Ciketunnelhistentry[i])] = types.YChild{"Ciketunnelhistentry", &ciketunnelhisttable.Ciketunnelhistentry[i]}
    }
    ciketunnelhisttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ciketunnelhisttable.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry
// Each entry contains the attributes
// associated with a previously active IPsec 
// Phase-1 IKE Tunnel.
type CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index of the IPsec Phase-1 IKE Tunnel History
    // Table.  The value of the index is a number which  begins at one and is
    // incremented with each  tunnel that ends. The value of this object  will
    // wrap at 2,147,483,647. The type is interface{} with range: 1..2147483647.
    Ciketunhistindex interface{}

    // The reason the IPsec Phase-1 IKE Tunnel was terminated. Possible reasons
    // include: 1 = other 2 = normal termination 3 = operator request 4 = peer
    // delete request was received 5 = contact with peer was lost 6 = local
    // failure occurred. 7 = operator initiated check point request. The type is
    // Ciketunhisttermreason.
    Ciketunhisttermreason interface{}

    // The index of the previously active IPsec Phase-1 IKE Tunnel. The type is
    // interface{} with range: 1..2147483647.
    Ciketunhistactiveindex interface{}

    // The type of local peer identity.  The local peer may be identified by:  1.
    // an IP address, or  2. a host name. The type is IkePeerType.
    Ciketunhistpeerlocaltype interface{}

    // The value of the local peer identity.  If the local peer type is an IP
    // Address, then this is the IP Address used to identify the local peer.  If
    // the local peer type is a host name, then this is the host name used to
    // identify the local peer. The type is string.
    Ciketunhistpeerlocalvalue interface{}

    // The internal index of the local-remote peer association.  This internal
    // index is used to  uniquely identify multiple associations between  the
    // local and remote peer. The type is interface{} with range: 1..2147483647.
    Ciketunhistpeerintindex interface{}

    // The type of remote peer identity.  The remote peer may be identified by: 
    // 1. an IP address, or  2. a host name. The type is IkePeerType.
    Ciketunhistpeerremotetype interface{}

    // The value of the remote peer identity.  If the remote peer type is an IP
    // Address, then this is the IP Address used to identify the remote peer.  If
    // the remote peer type is a host name, then this is the host name used to
    // identify the remote peer. The type is string.
    Ciketunhistpeerremotevalue interface{}

    // The IP address of the local endpoint for the IPsec Phase-1 IKE Tunnel. The
    // type is string with length: 4 | 16.
    Ciketunhistlocaladdr interface{}

    // The DNS name of the local IP address for the IPsec Phase-1 IKE Tunnel. If
    // the DNS  name associated with the local tunnel endpoint  is not known, then
    // the value of this  object will be a NULL string. The type is string.
    Ciketunhistlocalname interface{}

    // The IP address of the remote endpoint for the IPsec Phase-1 IKE Tunnel. The
    // type is string with length: 4 | 16.
    Ciketunhistremoteaddr interface{}

    // The DNS name of the remote IP address of IPsec Phase-1 IKE Tunnel. If the
    // DNS name associated with the remote tunnel endpoint is not known, then the
    // value of this object will be a NULL string. The type is string.
    Ciketunhistremotename interface{}

    // The negotiation mode of the IPsec Phase-1 IKE Tunnel. The type is
    // IkeNegoMode.
    Ciketunhistnegomode interface{}

    // The Diffie Hellman Group used in IPsec Phase-1 IKE negotiations. The type
    // is DiffHellmanGrp.
    Ciketunhistdiffhellmangrp interface{}

    // The encryption algorithm used in IPsec Phase-1 IKE negotiations. The type
    // is EncryptAlgo.
    Ciketunhistencryptalgo interface{}

    // The hash algorithm used in IPsec Phase-1 IKE negotiations. The type is
    // IkeHashAlgo.
    Ciketunhisthashalgo interface{}

    // The authentication method used in IPsec Phase-1 IKE negotiations. The type
    // is IkeAuthMethod.
    Ciketunhistauthmethod interface{}

    // The negotiated LifeTime of the IPsec Phase-1 IKE Tunnel in seconds. The
    // type is interface{} with range: 1..2147483647.
    Ciketunhistlifetime interface{}

    // The value of sysUpTime in hundredths of seconds when the IPsec Phase-1 IKE
    // tunnel was started. The type is interface{} with range: 0..4294967295.
    Ciketunhiststarttime interface{}

    // The length of time the IPsec Phase-1 IKE tunnel was been active in
    // hundredths of seconds. The type is interface{} with range: 0..2147483647.
    Ciketunhistactivetime interface{}

    // The total number of security associations refreshes performed. The type is
    // interface{} with range: 0..4294967295. Units are QM Exchanges.
    Ciketunhisttotalrefreshes interface{}

    // The total number of security associations used during the  life of the
    // IPsec Phase-1 IKE Tunnel. The type is interface{} with range:
    // 0..4294967295. Units are SAs.
    Ciketunhisttotalsas interface{}

    // The total number of octets received by this IPsec Phase-1  IKE Tunnel. The
    // type is interface{} with range: 0..4294967295. Units are Octets.
    Ciketunhistinoctets interface{}

    // The total number of packets received by this IPsec Phase-1  IKE Tunnel. The
    // type is interface{} with range: 0..4294967295. Units are Packets.
    Ciketunhistinpkts interface{}

    // The total number of packets dropped by this IPsec Phase-1  IKE Tunnel
    // during receive processing. The type is interface{} with range:
    // 0..4294967295. Units are Packets.
    Ciketunhistindroppkts interface{}

    // The total number of notifys received by this IPsec Phase-1  IKE Tunnel. The
    // type is interface{} with range: 0..4294967295. Units are Notification
    // Payloads.
    Ciketunhistinnotifys interface{}

    // The total number of IPsec Phase-2 exchanges received by  this IPsec Phase-1
    // IKE Tunnel. The type is interface{} with range: 0..4294967295. Units are SA
    // Payloads.
    Ciketunhistinp2Exchgs interface{}

    // The total number of IPsec Phase-2 exchanges received and  found to be
    // invalid by this IPsec Phase-1 IKE Tunnel. The type is interface{} with
    // range: 0..4294967295. Units are SA Payloads.
    Ciketunhistinp2Exchginvalids interface{}

    // The total number of IPsec Phase-2 exchanges received and  rejected by this
    // IPsec Phase-1 IKE Tunnel. The type is interface{} with range:
    // 0..4294967295. Units are SA Payloads.
    Ciketunhistinp2Exchgrejects interface{}

    // The total number of IPsec Phase-2 security association delete requests
    // received by this IPsec  Phase-1 IKE Tunnel. The type is interface{} with
    // range: 0..4294967295. Units are Notification Payloads.
    Ciketunhistinp2Sadelrequests interface{}

    // The total number of octets sent by this IPsec Phase-1 IKE Tunnel. The type
    // is interface{} with range: 0..4294967295. Units are Octets.
    Ciketunhistoutoctets interface{}

    // The total number of packets sent by this IPsec Phase-1 IKE Tunnel. The type
    // is interface{} with range: 0..4294967295. Units are Packets.
    Ciketunhistoutpkts interface{}

    // The total number of packets dropped by this IPsec Phase-1  IKE Tunnel
    // during send processing. The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    Ciketunhistoutdroppkts interface{}

    // The total number of notifys sent by this IPsec Phase-1 IKE Tunnel. The type
    // is interface{} with range: 0..4294967295. Units are Notification Payloads.
    Ciketunhistoutnotifys interface{}

    // The total number of IPsec Phase-2 exchanges sent by this IPsec Phase-1 IKE
    // Tunnel. The type is interface{} with range: 0..4294967295. Units are SA
    // Payloads.
    Ciketunhistoutp2Exchgs interface{}

    // The total number of IPsec Phase-2 exchanges sent and found to be invalid by
    // this IPsec Phase-1 IKE Tunnel. The type is interface{} with range:
    // 0..4294967295. Units are SA Payloads.
    Ciketunhistoutp2Exchginvalids interface{}

    // The total number of IPsec Phase-2 exchanges sent and rejected by this IPsec
    // Phase-1 IKE Tunnel. The type is interface{} with range: 0..4294967295.
    // Units are SA Payloads.
    Ciketunhistoutp2Exchgrejects interface{}

    // The total number of IPsec Phase-2 security association delete requests sent
    // by this IPsec Phase-1 IKE Tunnel. The type is interface{} with range:
    // 0..4294967295. Units are Notification Payloads.
    Ciketunhistoutp2Sadelrequests interface{}
}

func (ciketunnelhistentry *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry) GetEntityData() *types.CommonEntityData {
    ciketunnelhistentry.EntityData.YFilter = ciketunnelhistentry.YFilter
    ciketunnelhistentry.EntityData.YangName = "cikeTunnelHistEntry"
    ciketunnelhistentry.EntityData.BundleName = "cisco_ios_xe"
    ciketunnelhistentry.EntityData.ParentYangName = "cikeTunnelHistTable"
    ciketunnelhistentry.EntityData.SegmentPath = "cikeTunnelHistEntry" + "[cikeTunHistIndex='" + fmt.Sprintf("%v", ciketunnelhistentry.Ciketunhistindex) + "']"
    ciketunnelhistentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciketunnelhistentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciketunnelhistentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciketunnelhistentry.EntityData.Children = make(map[string]types.YChild)
    ciketunnelhistentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistIndex"] = types.YLeaf{"Ciketunhistindex", ciketunnelhistentry.Ciketunhistindex}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistTermReason"] = types.YLeaf{"Ciketunhisttermreason", ciketunnelhistentry.Ciketunhisttermreason}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistActiveIndex"] = types.YLeaf{"Ciketunhistactiveindex", ciketunnelhistentry.Ciketunhistactiveindex}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistPeerLocalType"] = types.YLeaf{"Ciketunhistpeerlocaltype", ciketunnelhistentry.Ciketunhistpeerlocaltype}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistPeerLocalValue"] = types.YLeaf{"Ciketunhistpeerlocalvalue", ciketunnelhistentry.Ciketunhistpeerlocalvalue}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistPeerIntIndex"] = types.YLeaf{"Ciketunhistpeerintindex", ciketunnelhistentry.Ciketunhistpeerintindex}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistPeerRemoteType"] = types.YLeaf{"Ciketunhistpeerremotetype", ciketunnelhistentry.Ciketunhistpeerremotetype}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistPeerRemoteValue"] = types.YLeaf{"Ciketunhistpeerremotevalue", ciketunnelhistentry.Ciketunhistpeerremotevalue}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistLocalAddr"] = types.YLeaf{"Ciketunhistlocaladdr", ciketunnelhistentry.Ciketunhistlocaladdr}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistLocalName"] = types.YLeaf{"Ciketunhistlocalname", ciketunnelhistentry.Ciketunhistlocalname}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistRemoteAddr"] = types.YLeaf{"Ciketunhistremoteaddr", ciketunnelhistentry.Ciketunhistremoteaddr}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistRemoteName"] = types.YLeaf{"Ciketunhistremotename", ciketunnelhistentry.Ciketunhistremotename}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistNegoMode"] = types.YLeaf{"Ciketunhistnegomode", ciketunnelhistentry.Ciketunhistnegomode}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistDiffHellmanGrp"] = types.YLeaf{"Ciketunhistdiffhellmangrp", ciketunnelhistentry.Ciketunhistdiffhellmangrp}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistEncryptAlgo"] = types.YLeaf{"Ciketunhistencryptalgo", ciketunnelhistentry.Ciketunhistencryptalgo}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistHashAlgo"] = types.YLeaf{"Ciketunhisthashalgo", ciketunnelhistentry.Ciketunhisthashalgo}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistAuthMethod"] = types.YLeaf{"Ciketunhistauthmethod", ciketunnelhistentry.Ciketunhistauthmethod}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistLifeTime"] = types.YLeaf{"Ciketunhistlifetime", ciketunnelhistentry.Ciketunhistlifetime}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistStartTime"] = types.YLeaf{"Ciketunhiststarttime", ciketunnelhistentry.Ciketunhiststarttime}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistActiveTime"] = types.YLeaf{"Ciketunhistactivetime", ciketunnelhistentry.Ciketunhistactivetime}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistTotalRefreshes"] = types.YLeaf{"Ciketunhisttotalrefreshes", ciketunnelhistentry.Ciketunhisttotalrefreshes}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistTotalSas"] = types.YLeaf{"Ciketunhisttotalsas", ciketunnelhistentry.Ciketunhisttotalsas}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistInOctets"] = types.YLeaf{"Ciketunhistinoctets", ciketunnelhistentry.Ciketunhistinoctets}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistInPkts"] = types.YLeaf{"Ciketunhistinpkts", ciketunnelhistentry.Ciketunhistinpkts}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistInDropPkts"] = types.YLeaf{"Ciketunhistindroppkts", ciketunnelhistentry.Ciketunhistindroppkts}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistInNotifys"] = types.YLeaf{"Ciketunhistinnotifys", ciketunnelhistentry.Ciketunhistinnotifys}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistInP2Exchgs"] = types.YLeaf{"Ciketunhistinp2Exchgs", ciketunnelhistentry.Ciketunhistinp2Exchgs}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistInP2ExchgInvalids"] = types.YLeaf{"Ciketunhistinp2Exchginvalids", ciketunnelhistentry.Ciketunhistinp2Exchginvalids}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistInP2ExchgRejects"] = types.YLeaf{"Ciketunhistinp2Exchgrejects", ciketunnelhistentry.Ciketunhistinp2Exchgrejects}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistInP2SaDelRequests"] = types.YLeaf{"Ciketunhistinp2Sadelrequests", ciketunnelhistentry.Ciketunhistinp2Sadelrequests}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistOutOctets"] = types.YLeaf{"Ciketunhistoutoctets", ciketunnelhistentry.Ciketunhistoutoctets}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistOutPkts"] = types.YLeaf{"Ciketunhistoutpkts", ciketunnelhistentry.Ciketunhistoutpkts}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistOutDropPkts"] = types.YLeaf{"Ciketunhistoutdroppkts", ciketunnelhistentry.Ciketunhistoutdroppkts}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistOutNotifys"] = types.YLeaf{"Ciketunhistoutnotifys", ciketunnelhistentry.Ciketunhistoutnotifys}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistOutP2Exchgs"] = types.YLeaf{"Ciketunhistoutp2Exchgs", ciketunnelhistentry.Ciketunhistoutp2Exchgs}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistOutP2ExchgInvalids"] = types.YLeaf{"Ciketunhistoutp2Exchginvalids", ciketunnelhistentry.Ciketunhistoutp2Exchginvalids}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistOutP2ExchgRejects"] = types.YLeaf{"Ciketunhistoutp2Exchgrejects", ciketunnelhistentry.Ciketunhistoutp2Exchgrejects}
    ciketunnelhistentry.EntityData.Leafs["cikeTunHistOutP2SaDelRequests"] = types.YLeaf{"Ciketunhistoutp2Sadelrequests", ciketunnelhistentry.Ciketunhistoutp2Sadelrequests}
    return &(ciketunnelhistentry.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry_Ciketunhisttermreason represents 7 = operator initiated check point request
type CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry_Ciketunhisttermreason string

const (
    CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry_Ciketunhisttermreason_other CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry_Ciketunhisttermreason = "other"

    CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry_Ciketunhisttermreason_normal CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry_Ciketunhisttermreason = "normal"

    CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry_Ciketunhisttermreason_operRequest CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry_Ciketunhisttermreason = "operRequest"

    CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry_Ciketunhisttermreason_peerDelRequest CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry_Ciketunhisttermreason = "peerDelRequest"

    CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry_Ciketunhisttermreason_peerLost CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry_Ciketunhisttermreason = "peerLost"

    CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry_Ciketunhisttermreason_localFailure CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry_Ciketunhisttermreason = "localFailure"

    CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry_Ciketunhisttermreason_checkPointReg CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry_Ciketunhisttermreason = "checkPointReg"
)

// CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable
// The IPsec Phase-2 Tunnel History Table.
// This table is implemented as a sliding 
// window in which only the
// last n entries are maintained.  The maximum number 
// of entries
// is specified by the cipSecHistTableSize object.
type CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with a previously active
    // IPsec Phase-2 Tunnel. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry.
    Cipsectunnelhistentry []CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry
}

func (cipsectunnelhisttable *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable) GetEntityData() *types.CommonEntityData {
    cipsectunnelhisttable.EntityData.YFilter = cipsectunnelhisttable.YFilter
    cipsectunnelhisttable.EntityData.YangName = "cipSecTunnelHistTable"
    cipsectunnelhisttable.EntityData.BundleName = "cisco_ios_xe"
    cipsectunnelhisttable.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cipsectunnelhisttable.EntityData.SegmentPath = "cipSecTunnelHistTable"
    cipsectunnelhisttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsectunnelhisttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsectunnelhisttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsectunnelhisttable.EntityData.Children = make(map[string]types.YChild)
    cipsectunnelhisttable.EntityData.Children["cipSecTunnelHistEntry"] = types.YChild{"Cipsectunnelhistentry", nil}
    for i := range cipsectunnelhisttable.Cipsectunnelhistentry {
        cipsectunnelhisttable.EntityData.Children[types.GetSegmentPath(&cipsectunnelhisttable.Cipsectunnelhistentry[i])] = types.YChild{"Cipsectunnelhistentry", &cipsectunnelhisttable.Cipsectunnelhistentry[i]}
    }
    cipsectunnelhisttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cipsectunnelhisttable.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry
// Each entry contains the attributes associated with
// a previously active IPsec Phase-2 Tunnel.
type CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index of the IPsec Phase-2 Tunnel History
    // Table. The value of the index is a number which  begins at one and is
    // incremented with each tunnel  that ends. The value of this object will wrap
    // at 2,147,483,647. The type is interface{} with range: 1..2147483647.
    Cipsectunhistindex interface{}

    // The reason the IPsec Phase-2 Tunnel was terminated. Possible reasons
    // include: 1 = other 2 = normal termination 3 = operator request 4 = peer
    // delete request was received 5 = contact with peer was lost 6 = local
    // failure occurred 7 = operator initiated check point request. The type is
    // Cipsectunhisttermreason.
    Cipsectunhisttermreason interface{}

    // The index of the previously active IPsec Phase-2 Tunnel. The type is
    // interface{} with range: 1..2147483647.
    Cipsectunhistactiveindex interface{}

    // The index of the associated IPsec Phase-1 Tunnel (cikeTunIndex in the
    // cikeTunnelTable). The type is interface{} with range: 1..2147483647.
    Cipsectunhistiketunnelindex interface{}

    // The IP address of the local endpoint for the IPsec Phase-2 Tunnel. The type
    // is string with length: 4 | 16.
    Cipsectunhistlocaladdr interface{}

    // The IP address of the remote endpoint for the IPsec Phase-2 Tunnel. The
    // type is string with length: 4 | 16.
    Cipsectunhistremoteaddr interface{}

    // The type of key used by the IPsec Phase-2 Tunnel. The type is KeyType.
    Cipsectunhistkeytype interface{}

    // The encapsulation mode used by the IPsec Phase-2 Tunnel. The type is
    // EncapMode.
    Cipsectunhistencapmode interface{}

    // The negotiated LifeSize of the IPsec Phase-2 Tunnel in kilobytes. The type
    // is interface{} with range: 1..2147483647. Units are KBytes.
    Cipsectunhistlifesize interface{}

    // The negotiated LifeTime of the IPsec Phase-2 Tunnel in seconds. The type is
    // interface{} with range: 1..2147483647. Units are Seconds.
    Cipsectunhistlifetime interface{}

    // The value of sysUpTime in hundredths of seconds when the IPsec Phase-2
    // Tunnel was started. The type is interface{} with range: 0..4294967295.
    Cipsectunhiststarttime interface{}

    // The length of time the IPsec Phase-2 Tunnel has been active in hundredths
    // of seconds. The type is interface{} with range: 0..2147483647.
    Cipsectunhistactivetime interface{}

    // The total number of security association refreshes performed. The type is
    // interface{} with range: 0..4294967295. Units are QM Exchanges.
    Cipsectunhisttotalrefreshes interface{}

    // The total number of security associations used during the  life of the
    // IPsec Phase-2 Tunnel. The type is interface{} with range: 0..4294967295.
    // Units are SAs.
    Cipsectunhisttotalsas interface{}

    // The Diffie Hellman Group used by the inbound security association of the
    // IPsec Phase-2 Tunnel. The type is DiffHellmanGrp.
    Cipsectunhistinsadiffhellmangrp interface{}

    // The encryption algorithm used by the inbound security association of the
    // IPsec Phase-2 Tunnel. The type is EncryptAlgo.
    Cipsectunhistinsaencryptalgo interface{}

    // The authentication algorithm used by the inbound authentication header (AH)
    // security association of the IPsec Phase-2 Tunnel. The type is AuthAlgo.
    Cipsectunhistinsaahauthalgo interface{}

    // The authentication algorithm used by the inbound encapsulation security
    // protocol (ESP)  security association of the IPsec Phase-2 Tunnel. The type
    // is AuthAlgo.
    Cipsectunhistinsaespauthalgo interface{}

    // The decompression algorithm used by the inbound security association of the
    // IPsec Phase-2 Tunnel. The type is CompAlgo.
    Cipsectunhistinsadecompalgo interface{}

    // The Diffie Hellman Group used by the outbound security association of the
    // IPsec Phase-2 Tunnel. The type is DiffHellmanGrp.
    Cipsectunhistoutsadiffhellmangrp interface{}

    // The encryption algorithm used by the outbound security association of the
    // IPsec Phase-2 Tunnel. The type is EncryptAlgo.
    Cipsectunhistoutsaencryptalgo interface{}

    // The authentication algorithm used by the outbound authentication header
    // (AH) security association of the IPsec Phase-2 Tunnel. The type is
    // AuthAlgo.
    Cipsectunhistoutsaahauthalgo interface{}

    // The authentication algorithm used by the inbound encapsulation security
    // protocol (ESP)  security association of the IPsec Phase-2 Tunnel. The type
    // is AuthAlgo.
    Cipsectunhistoutsaespauthalgo interface{}

    // The compression algorithm used by the inbound security association of the
    // IPsec Phase-2 Tunnel. The type is CompAlgo.
    Cipsectunhistoutsacompalgo interface{}

    // The total number of octets received by this IPsec Phase-2 Tunnel.  This
    // value is accumulated BEFORE determining whether or not the packet should 
    // be decompressed.  See also cipSecTunInOctWraps for  the number of times
    // this counter has wrapped. The type is interface{} with range:
    // 0..4294967295. Units are Octets.
    Cipsectunhistinoctets interface{}

    // A high capacity count of the total number of octets received by this IPsec
    // Phase-2 Tunnel.  This value is accumulated BEFORE determining whether or
    // not  the packet should be decompressed. The type is interface{} with range:
    // 0..18446744073709551615.
    Cipsectunhisthcinoctets interface{}

    // The number of times the octets received counter (cipSecTunInOctets) has
    // wrapped. The type is interface{} with range: 0..4294967295. Units are
    // Integral units.
    Cipsectunhistinoctwraps interface{}

    // The total number of decompressed octets received by this IPsec Phase-2
    // Tunnel.  This value is accumulated AFTER the packet is decompressed. If
    // compression is not being used, this value will match the value of
    // cipSecTunHistInOctets. See also cipSecTunInDecompOctWraps for the number of
    // times this counter has wrapped. The type is interface{} with range:
    // 0..4294967295. Units are Octets.
    Cipsectunhistindecompoctets interface{}

    // A high capacity count of the total number of decompressed octets received
    // by this IPsec Phase-2 Tunnel.  This value is accumulated AFTER the packet
    // is decompressed. If compression is not being used, this value will match
    // the value of cipSecTunHistHcInOctets. The type is interface{} with range:
    // 0..18446744073709551615.
    Cipsectunhisthcindecompoctets interface{}

    // The number of times the decompressed octets received counter
    // (cipSecTunInDecompOctets) has wrapped. The type is interface{} with range:
    // 0..4294967295. Units are Integral units.
    Cipsectunhistindecompoctwraps interface{}

    // The total number of packets received by this IPsec Phase-2 Tunnel. The type
    // is interface{} with range: 0..4294967295. Units are Packets.
    Cipsectunhistinpkts interface{}

    // The total number of packets dropped during receive processing by this IPsec
    // Phase-2 Tunnel.  This count does NOT include packets  dropped due to
    // Anti-Replay processing. The type is interface{} with range: 0..4294967295.
    // Units are Packets.
    Cipsectunhistindroppkts interface{}

    // The total number of packets dropped during receive processing due to
    // Anti-Replay processing  by this IPsec Phase-2 Tunnel. The type is
    // interface{} with range: 0..4294967295. Units are Packets.
    Cipsectunhistinreplaydroppkts interface{}

    // The total number of inbound authentication's performed  by this IPsec
    // Phase-2 Tunnel. The type is interface{} with range: 0..4294967295. Units
    // are Events.
    Cipsectunhistinauths interface{}

    // The total number of inbound authentication's which ended in  failure by
    // this IPsec Phase-2 Tunnel . The type is interface{} with range:
    // 0..4294967295. Units are Failures.
    Cipsectunhistinauthfails interface{}

    // The total number of inbound decryption's performed by this IPsec Phase-2
    // Tunnel. The type is interface{} with range: 0..4294967295. Units are
    // Packets.
    Cipsectunhistindecrypts interface{}

    // The total number of inbound decryption's which ended in failure  by this
    // IPsec Phase-2 Tunnel. The type is interface{} with range: 0..4294967295.
    // Units are Failures.
    Cipsectunhistindecryptfails interface{}

    // The total number of octets sent by this IPsec Phase-2 Tunnel.  This value
    // is accumulated AFTER determining whether or not the  packet should be
    // compressed.  See also cipSecTunOutOctWraps for the number of times this
    // counter has wrapped. The type is interface{} with range: 0..4294967295.
    // Units are Octets.
    Cipsectunhistoutoctets interface{}

    // A high capacity count of the total number of octets sent by this IPsec
    // Phase-2 Tunnel.  This value  is accumulated AFTER determining whether or
    // not  the packet should be compressed. The type is interface{} with range:
    // 0..18446744073709551615.
    Cipsectunhisthcoutoctets interface{}

    // The number of times the octets sent counter (cipSecTunOutOctets) has
    // wrapped. The type is interface{} with range: 0..4294967295. Units are
    // Integral units.
    Cipsectunhistoutoctwraps interface{}

    // The total number of uncompressed octets sent by this IPsec Phase-2 Tunnel. 
    // This value is accumulated BEFORE the packet is compressed. If compression
    // is not being used, this value will match the value of 
    // cipSecTunHistOutOctets.  See also  cipSecTunOutDecompOctWraps for the
    // number of times this counter has wrapped. The type is interface{} with
    // range: 0..4294967295. Units are Octets.
    Cipsectunhistoutuncompoctets interface{}

    // A high capacity count of the total number of uncompressed octets sent by
    // this  IPsec Phase-2 Tunnel.  This value is accumulated  BEFORE the packet
    // is compressed. If compression is not being used, this value will match the
    // value of cipSecTunHistHcOutOctets. The type is interface{} with range:
    // 0..18446744073709551615. Units are Octets.
    Cipsectunhisthcoutuncompoctets interface{}

    // The number of times the uncompressed octets sent counter
    // (cipSecTunOutUncompOctets) has wrapped. The type is interface{} with range:
    // 0..4294967295. Units are Integral units.
    Cipsectunhistoutuncompoctwraps interface{}

    // The total number of packets sent by this IPsec Phase-2 Tunnel. The type is
    // interface{} with range: 0..4294967295. Units are Packets.
    Cipsectunhistoutpkts interface{}

    // The total number of packets dropped during send processing  by this IPsec
    // Phase-2 Tunnel. The type is interface{} with range: 0..4294967295. Units
    // are Packets.
    Cipsectunhistoutdroppkts interface{}

    // The total number of outbound authentication's performed by this IPsec
    // Phase-2 Tunnel. The type is interface{} with range: 0..4294967295. Units
    // are Events.
    Cipsectunhistoutauths interface{}

    // The total number of outbound authentication's which ended in  failure by
    // this IPsec Phase-2 Tunnel. The type is interface{} with range:
    // 0..4294967295. Units are Failures.
    Cipsectunhistoutauthfails interface{}

    // The total number of outbound encryption's performed by this IPsec Phase-2
    // Tunnel. The type is interface{} with range: 0..4294967295. Units are
    // Packets.
    Cipsectunhistoutencrypts interface{}

    // The total number of outbound encryption's which ended in failure  by this
    // IPsec Phase-2 Tunnel. The type is interface{} with range: 0..4294967295.
    // Units are Failures.
    Cipsectunhistoutencryptfails interface{}
}

func (cipsectunnelhistentry *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry) GetEntityData() *types.CommonEntityData {
    cipsectunnelhistentry.EntityData.YFilter = cipsectunnelhistentry.YFilter
    cipsectunnelhistentry.EntityData.YangName = "cipSecTunnelHistEntry"
    cipsectunnelhistentry.EntityData.BundleName = "cisco_ios_xe"
    cipsectunnelhistentry.EntityData.ParentYangName = "cipSecTunnelHistTable"
    cipsectunnelhistentry.EntityData.SegmentPath = "cipSecTunnelHistEntry" + "[cipSecTunHistIndex='" + fmt.Sprintf("%v", cipsectunnelhistentry.Cipsectunhistindex) + "']"
    cipsectunnelhistentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsectunnelhistentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsectunnelhistentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsectunnelhistentry.EntityData.Children = make(map[string]types.YChild)
    cipsectunnelhistentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistIndex"] = types.YLeaf{"Cipsectunhistindex", cipsectunnelhistentry.Cipsectunhistindex}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistTermReason"] = types.YLeaf{"Cipsectunhisttermreason", cipsectunnelhistentry.Cipsectunhisttermreason}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistActiveIndex"] = types.YLeaf{"Cipsectunhistactiveindex", cipsectunnelhistentry.Cipsectunhistactiveindex}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistIkeTunnelIndex"] = types.YLeaf{"Cipsectunhistiketunnelindex", cipsectunnelhistentry.Cipsectunhistiketunnelindex}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistLocalAddr"] = types.YLeaf{"Cipsectunhistlocaladdr", cipsectunnelhistentry.Cipsectunhistlocaladdr}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistRemoteAddr"] = types.YLeaf{"Cipsectunhistremoteaddr", cipsectunnelhistentry.Cipsectunhistremoteaddr}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistKeyType"] = types.YLeaf{"Cipsectunhistkeytype", cipsectunnelhistentry.Cipsectunhistkeytype}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistEncapMode"] = types.YLeaf{"Cipsectunhistencapmode", cipsectunnelhistentry.Cipsectunhistencapmode}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistLifeSize"] = types.YLeaf{"Cipsectunhistlifesize", cipsectunnelhistentry.Cipsectunhistlifesize}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistLifeTime"] = types.YLeaf{"Cipsectunhistlifetime", cipsectunnelhistentry.Cipsectunhistlifetime}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistStartTime"] = types.YLeaf{"Cipsectunhiststarttime", cipsectunnelhistentry.Cipsectunhiststarttime}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistActiveTime"] = types.YLeaf{"Cipsectunhistactivetime", cipsectunnelhistentry.Cipsectunhistactivetime}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistTotalRefreshes"] = types.YLeaf{"Cipsectunhisttotalrefreshes", cipsectunnelhistentry.Cipsectunhisttotalrefreshes}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistTotalSas"] = types.YLeaf{"Cipsectunhisttotalsas", cipsectunnelhistentry.Cipsectunhisttotalsas}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistInSaDiffHellmanGrp"] = types.YLeaf{"Cipsectunhistinsadiffhellmangrp", cipsectunnelhistentry.Cipsectunhistinsadiffhellmangrp}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistInSaEncryptAlgo"] = types.YLeaf{"Cipsectunhistinsaencryptalgo", cipsectunnelhistentry.Cipsectunhistinsaencryptalgo}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistInSaAhAuthAlgo"] = types.YLeaf{"Cipsectunhistinsaahauthalgo", cipsectunnelhistentry.Cipsectunhistinsaahauthalgo}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistInSaEspAuthAlgo"] = types.YLeaf{"Cipsectunhistinsaespauthalgo", cipsectunnelhistentry.Cipsectunhistinsaespauthalgo}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistInSaDecompAlgo"] = types.YLeaf{"Cipsectunhistinsadecompalgo", cipsectunnelhistentry.Cipsectunhistinsadecompalgo}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistOutSaDiffHellmanGrp"] = types.YLeaf{"Cipsectunhistoutsadiffhellmangrp", cipsectunnelhistentry.Cipsectunhistoutsadiffhellmangrp}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistOutSaEncryptAlgo"] = types.YLeaf{"Cipsectunhistoutsaencryptalgo", cipsectunnelhistentry.Cipsectunhistoutsaencryptalgo}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistOutSaAhAuthAlgo"] = types.YLeaf{"Cipsectunhistoutsaahauthalgo", cipsectunnelhistentry.Cipsectunhistoutsaahauthalgo}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistOutSaEspAuthAlgo"] = types.YLeaf{"Cipsectunhistoutsaespauthalgo", cipsectunnelhistentry.Cipsectunhistoutsaespauthalgo}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistOutSaCompAlgo"] = types.YLeaf{"Cipsectunhistoutsacompalgo", cipsectunnelhistentry.Cipsectunhistoutsacompalgo}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistInOctets"] = types.YLeaf{"Cipsectunhistinoctets", cipsectunnelhistentry.Cipsectunhistinoctets}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistHcInOctets"] = types.YLeaf{"Cipsectunhisthcinoctets", cipsectunnelhistentry.Cipsectunhisthcinoctets}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistInOctWraps"] = types.YLeaf{"Cipsectunhistinoctwraps", cipsectunnelhistentry.Cipsectunhistinoctwraps}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistInDecompOctets"] = types.YLeaf{"Cipsectunhistindecompoctets", cipsectunnelhistentry.Cipsectunhistindecompoctets}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistHcInDecompOctets"] = types.YLeaf{"Cipsectunhisthcindecompoctets", cipsectunnelhistentry.Cipsectunhisthcindecompoctets}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistInDecompOctWraps"] = types.YLeaf{"Cipsectunhistindecompoctwraps", cipsectunnelhistentry.Cipsectunhistindecompoctwraps}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistInPkts"] = types.YLeaf{"Cipsectunhistinpkts", cipsectunnelhistentry.Cipsectunhistinpkts}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistInDropPkts"] = types.YLeaf{"Cipsectunhistindroppkts", cipsectunnelhistentry.Cipsectunhistindroppkts}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistInReplayDropPkts"] = types.YLeaf{"Cipsectunhistinreplaydroppkts", cipsectunnelhistentry.Cipsectunhistinreplaydroppkts}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistInAuths"] = types.YLeaf{"Cipsectunhistinauths", cipsectunnelhistentry.Cipsectunhistinauths}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistInAuthFails"] = types.YLeaf{"Cipsectunhistinauthfails", cipsectunnelhistentry.Cipsectunhistinauthfails}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistInDecrypts"] = types.YLeaf{"Cipsectunhistindecrypts", cipsectunnelhistentry.Cipsectunhistindecrypts}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistInDecryptFails"] = types.YLeaf{"Cipsectunhistindecryptfails", cipsectunnelhistentry.Cipsectunhistindecryptfails}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistOutOctets"] = types.YLeaf{"Cipsectunhistoutoctets", cipsectunnelhistentry.Cipsectunhistoutoctets}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistHcOutOctets"] = types.YLeaf{"Cipsectunhisthcoutoctets", cipsectunnelhistentry.Cipsectunhisthcoutoctets}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistOutOctWraps"] = types.YLeaf{"Cipsectunhistoutoctwraps", cipsectunnelhistentry.Cipsectunhistoutoctwraps}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistOutUncompOctets"] = types.YLeaf{"Cipsectunhistoutuncompoctets", cipsectunnelhistentry.Cipsectunhistoutuncompoctets}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistHcOutUncompOctets"] = types.YLeaf{"Cipsectunhisthcoutuncompoctets", cipsectunnelhistentry.Cipsectunhisthcoutuncompoctets}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistOutUncompOctWraps"] = types.YLeaf{"Cipsectunhistoutuncompoctwraps", cipsectunnelhistentry.Cipsectunhistoutuncompoctwraps}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistOutPkts"] = types.YLeaf{"Cipsectunhistoutpkts", cipsectunnelhistentry.Cipsectunhistoutpkts}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistOutDropPkts"] = types.YLeaf{"Cipsectunhistoutdroppkts", cipsectunnelhistentry.Cipsectunhistoutdroppkts}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistOutAuths"] = types.YLeaf{"Cipsectunhistoutauths", cipsectunnelhistentry.Cipsectunhistoutauths}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistOutAuthFails"] = types.YLeaf{"Cipsectunhistoutauthfails", cipsectunnelhistentry.Cipsectunhistoutauthfails}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistOutEncrypts"] = types.YLeaf{"Cipsectunhistoutencrypts", cipsectunnelhistentry.Cipsectunhistoutencrypts}
    cipsectunnelhistentry.EntityData.Leafs["cipSecTunHistOutEncryptFails"] = types.YLeaf{"Cipsectunhistoutencryptfails", cipsectunnelhistentry.Cipsectunhistoutencryptfails}
    return &(cipsectunnelhistentry.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry_Cipsectunhisttermreason represents 7 = operator initiated check point request
type CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry_Cipsectunhisttermreason string

const (
    CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry_Cipsectunhisttermreason_other CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry_Cipsectunhisttermreason = "other"

    CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry_Cipsectunhisttermreason_normal CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry_Cipsectunhisttermreason = "normal"

    CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry_Cipsectunhisttermreason_operRequest CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry_Cipsectunhisttermreason = "operRequest"

    CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry_Cipsectunhisttermreason_peerDelRequest CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry_Cipsectunhisttermreason = "peerDelRequest"

    CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry_Cipsectunhisttermreason_peerLost CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry_Cipsectunhisttermreason = "peerLost"

    CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry_Cipsectunhisttermreason_seqNumRollOver CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry_Cipsectunhisttermreason = "seqNumRollOver"

    CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry_Cipsectunhisttermreason_checkPointReq CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry_Cipsectunhisttermreason = "checkPointReq"
)

// CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable
// The IPsec Phase-2 Tunnel Endpoint History Table.
// This table is implemented as a 
// sliding window in which only the
// last n entries are maintained.  
// The maximum number of entries
// is specified by the cipSecHistTableSize object.
type CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with a previously active
    // IPsec Phase-2 Tunnel Endpoint. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable_Cipsecendpthistentry.
    Cipsecendpthistentry []CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable_Cipsecendpthistentry
}

func (cipsecendpthisttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable) GetEntityData() *types.CommonEntityData {
    cipsecendpthisttable.EntityData.YFilter = cipsecendpthisttable.YFilter
    cipsecendpthisttable.EntityData.YangName = "cipSecEndPtHistTable"
    cipsecendpthisttable.EntityData.BundleName = "cisco_ios_xe"
    cipsecendpthisttable.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cipsecendpthisttable.EntityData.SegmentPath = "cipSecEndPtHistTable"
    cipsecendpthisttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsecendpthisttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsecendpthisttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsecendpthisttable.EntityData.Children = make(map[string]types.YChild)
    cipsecendpthisttable.EntityData.Children["cipSecEndPtHistEntry"] = types.YChild{"Cipsecendpthistentry", nil}
    for i := range cipsecendpthisttable.Cipsecendpthistentry {
        cipsecendpthisttable.EntityData.Children[types.GetSegmentPath(&cipsecendpthisttable.Cipsecendpthistentry[i])] = types.YChild{"Cipsecendpthistentry", &cipsecendpthisttable.Cipsecendpthistentry[i]}
    }
    cipsecendpthisttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cipsecendpthisttable.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable_Cipsecendpthistentry
// Each entry contains the attributes associated with
// a previously active IPsec Phase-2 Tunnel Endpoint.
type CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable_Cipsecendpthistentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The number of the previously active Endpoint
    // associated  with a IPsec Phase-2 Tunnel Table.  The value   of this index
    // is a number which begins at   one and is incremented with each Endpoint  
    // associated with an IPsec Phase-2 Tunnel.  The value of this object will
    // wrap at 2,147,483,647. The type is interface{} with range: 1..2147483647.
    Cipsecendpthistindex interface{}

    // The index  of the previously active IPsec Phase-2 Tunnel Table. The type is
    // interface{} with range: 1..2147483647.
    Cipsecendpthisttunindex interface{}

    // The index  of the previously active Endpoint. The type is interface{} with
    // range: 1..2147483647.
    Cipsecendpthistactiveindex interface{}

    // The DNS name of the local Endpoint. The type is string.
    Cipsecendpthistlocalname interface{}

    // The type of identity for the local Endpoint. Possible values are: 1) a
    // single IP address, or 2) an IP address range, or 3) an IP subnet. The type
    // is EndPtType.
    Cipsecendpthistlocaltype interface{}

    // The local Endpoint's first IP address specification.  If the local Endpoint
    // type is single IP address,  then this is the value of the IP address.  If
    // the local Endpoint type is IP subnet, then this is the value of the subnet.
    // If the local Endpoint type is IP address range,  then this is the value of
    // beginning IP address of  the range. The type is string with length: 4 | 16.
    Cipsecendpthistlocaladdr1 interface{}

    // The local Endpoint's second IP address specification.  If the local
    // Endpoint type is single IP address,  then this is the value of the IP
    // address.  If the local Endpoint type is IP subnet, then this is the value
    // of the subnet mask.  If the local Endpoint type is IP address range,  then
    // this is the value of ending IP address of the range. The type is string
    // with length: 4 | 16.
    Cipsecendpthistlocaladdr2 interface{}

    // The protocol number of the local Endpoint's traffic. The type is
    // interface{} with range: 0..255.
    Cipsecendpthistlocalprotocol interface{}

    // The port number of the local Endpoint's traffic. The type is interface{}
    // with range: 0..65535.
    Cipsecendpthistlocalport interface{}

    // The DNS name of the remote Endpoint. The type is string.
    Cipsecendpthistremotename interface{}

    // The type of identity for the remote Endpoint. Possible values are: 1) a
    // single IP address, or 2) an IP address range, or 3) an IP subnet. The type
    // is EndPtType.
    Cipsecendpthistremotetype interface{}

    // The remote Endpoint's first IP address specification.  If the remote
    // Endpoint type is single IP address,  then this is the value of the IP
    // address.  If the remote Endpoint type is IP subnet, then this is the value
    // of the subnet.  If the remote Endpoint type is IP address range,  then this
    // is the value of beginning IP address of the range. The type is string with
    // length: 4 | 16.
    Cipsecendpthistremoteaddr1 interface{}

    // The remote Endpoint's second IP address specification.  If the remote
    // Endpoint type is single IP address,  then this is the value of the IP
    // address.  If the remote Endpoint type is IP subnet, then this is the value
    // of the subnet mask.  If the remote Endpoint type is IP address range,  then
    // this is the value of ending IP address of the range. The type is string
    // with length: 4 | 16.
    Cipsecendpthistremoteaddr2 interface{}

    // The protocol number of the remote Endpoint's traffic. The type is
    // interface{} with range: 0..255.
    Cipsecendpthistremoteprotocol interface{}

    // The port number of the remote Endpoint's traffic. The type is interface{}
    // with range: 0..65535.
    Cipsecendpthistremoteport interface{}
}

func (cipsecendpthistentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable_Cipsecendpthistentry) GetEntityData() *types.CommonEntityData {
    cipsecendpthistentry.EntityData.YFilter = cipsecendpthistentry.YFilter
    cipsecendpthistentry.EntityData.YangName = "cipSecEndPtHistEntry"
    cipsecendpthistentry.EntityData.BundleName = "cisco_ios_xe"
    cipsecendpthistentry.EntityData.ParentYangName = "cipSecEndPtHistTable"
    cipsecendpthistentry.EntityData.SegmentPath = "cipSecEndPtHistEntry" + "[cipSecEndPtHistIndex='" + fmt.Sprintf("%v", cipsecendpthistentry.Cipsecendpthistindex) + "']"
    cipsecendpthistentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsecendpthistentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsecendpthistentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsecendpthistentry.EntityData.Children = make(map[string]types.YChild)
    cipsecendpthistentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cipsecendpthistentry.EntityData.Leafs["cipSecEndPtHistIndex"] = types.YLeaf{"Cipsecendpthistindex", cipsecendpthistentry.Cipsecendpthistindex}
    cipsecendpthistentry.EntityData.Leafs["cipSecEndPtHistTunIndex"] = types.YLeaf{"Cipsecendpthisttunindex", cipsecendpthistentry.Cipsecendpthisttunindex}
    cipsecendpthistentry.EntityData.Leafs["cipSecEndPtHistActiveIndex"] = types.YLeaf{"Cipsecendpthistactiveindex", cipsecendpthistentry.Cipsecendpthistactiveindex}
    cipsecendpthistentry.EntityData.Leafs["cipSecEndPtHistLocalName"] = types.YLeaf{"Cipsecendpthistlocalname", cipsecendpthistentry.Cipsecendpthistlocalname}
    cipsecendpthistentry.EntityData.Leafs["cipSecEndPtHistLocalType"] = types.YLeaf{"Cipsecendpthistlocaltype", cipsecendpthistentry.Cipsecendpthistlocaltype}
    cipsecendpthistentry.EntityData.Leafs["cipSecEndPtHistLocalAddr1"] = types.YLeaf{"Cipsecendpthistlocaladdr1", cipsecendpthistentry.Cipsecendpthistlocaladdr1}
    cipsecendpthistentry.EntityData.Leafs["cipSecEndPtHistLocalAddr2"] = types.YLeaf{"Cipsecendpthistlocaladdr2", cipsecendpthistentry.Cipsecendpthistlocaladdr2}
    cipsecendpthistentry.EntityData.Leafs["cipSecEndPtHistLocalProtocol"] = types.YLeaf{"Cipsecendpthistlocalprotocol", cipsecendpthistentry.Cipsecendpthistlocalprotocol}
    cipsecendpthistentry.EntityData.Leafs["cipSecEndPtHistLocalPort"] = types.YLeaf{"Cipsecendpthistlocalport", cipsecendpthistentry.Cipsecendpthistlocalport}
    cipsecendpthistentry.EntityData.Leafs["cipSecEndPtHistRemoteName"] = types.YLeaf{"Cipsecendpthistremotename", cipsecendpthistentry.Cipsecendpthistremotename}
    cipsecendpthistentry.EntityData.Leafs["cipSecEndPtHistRemoteType"] = types.YLeaf{"Cipsecendpthistremotetype", cipsecendpthistentry.Cipsecendpthistremotetype}
    cipsecendpthistentry.EntityData.Leafs["cipSecEndPtHistRemoteAddr1"] = types.YLeaf{"Cipsecendpthistremoteaddr1", cipsecendpthistentry.Cipsecendpthistremoteaddr1}
    cipsecendpthistentry.EntityData.Leafs["cipSecEndPtHistRemoteAddr2"] = types.YLeaf{"Cipsecendpthistremoteaddr2", cipsecendpthistentry.Cipsecendpthistremoteaddr2}
    cipsecendpthistentry.EntityData.Leafs["cipSecEndPtHistRemoteProtocol"] = types.YLeaf{"Cipsecendpthistremoteprotocol", cipsecendpthistentry.Cipsecendpthistremoteprotocol}
    cipsecendpthistentry.EntityData.Leafs["cipSecEndPtHistRemotePort"] = types.YLeaf{"Cipsecendpthistremoteport", cipsecendpthistentry.Cipsecendpthistremoteport}
    return &(cipsecendpthistentry.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cikefailtable
// The IPsec Phase-1 Failure Table.
// This table is implemented as a sliding 
// window in which only the last n entries are 
// maintained.  The maximum number of entries
// is specified by the cipSecFailTableSize object.
type CISCOIPSECFLOWMONITORMIB_Cikefailtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with  an IPsec Phase-1
    // failure. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry.
    Cikefailentry []CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry
}

func (cikefailtable *CISCOIPSECFLOWMONITORMIB_Cikefailtable) GetEntityData() *types.CommonEntityData {
    cikefailtable.EntityData.YFilter = cikefailtable.YFilter
    cikefailtable.EntityData.YangName = "cikeFailTable"
    cikefailtable.EntityData.BundleName = "cisco_ios_xe"
    cikefailtable.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cikefailtable.EntityData.SegmentPath = "cikeFailTable"
    cikefailtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cikefailtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cikefailtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cikefailtable.EntityData.Children = make(map[string]types.YChild)
    cikefailtable.EntityData.Children["cikeFailEntry"] = types.YChild{"Cikefailentry", nil}
    for i := range cikefailtable.Cikefailentry {
        cikefailtable.EntityData.Children[types.GetSegmentPath(&cikefailtable.Cikefailentry[i])] = types.YChild{"Cikefailentry", &cikefailtable.Cikefailentry[i]}
    }
    cikefailtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cikefailtable.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry
// Each entry contains the attributes associated
// with
//  an IPsec Phase-1 failure.
type CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IPsec Phase-1 Failure Table index. The value
    // of the index is a number which  begins at one and is incremented with each 
    // IPsec Phase-1 failure. The value of this object will wrap at 2,147,483,647.
    // The type is interface{} with range: 1..2147483647.
    Cikefailindex interface{}

    // The reason for the failure.  Possible reasons include: 1 = other 2 = peer
    // delete request was received 3 = contact with peer was lost 4 = local
    // failure occurred 5 = authentication failure 6 = hash validation failure 7 =
    // encryption failure 8 = internal error occurred 9 = system capacity failure
    // 10 = proposal failure 11 = peer's certificate is unavailable 12 = peer's
    // certificate was found invalid 13 = local certificate expired 14 =
    // certificate revoke list (crl) failure 15 = peer encoding error 16 =
    // non-existent security association 17 = operator requested termination. The
    // type is Cikefailreason.
    Cikefailreason interface{}

    // The value of sysUpTime in hundredths of seconds at the time of the failure.
    // The type is interface{} with range: 0..4294967295.
    Cikefailtime interface{}

    // The type of local peer identity.  The local peer may be identified by:  1.
    // an IP address, or  2. a host name. The type is IkePeerType.
    Cikefaillocaltype interface{}

    // The value of the local peer identity.  If the local peer type is an IP
    // Address, then this is the IP Address used to identify the local peer.  If
    // the local peer type is a host name, then this is the host name used to
    // identify the local peer. The type is string.
    Cikefaillocalvalue interface{}

    // The type of remote peer identity.  The remote peer may be identified by: 
    // 1. an IP address, or  2. a host name. The type is IkePeerType.
    Cikefailremotetype interface{}

    // The value of the remote peer identity.  If the remote peer type is an IP
    // Address, then this is the IP Address used to identify the remote peer.  If
    // the remote peer type is a host name, then this is the host name used to
    // identify the remote peer. The type is string.
    Cikefailremotevalue interface{}

    // The IP address of the local peer. The type is string with length: 4 | 16.
    Cikefaillocaladdr interface{}

    // The IP address of the remote peer. The type is string with length: 4 | 16.
    Cikefailremoteaddr interface{}
}

func (cikefailentry *CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry) GetEntityData() *types.CommonEntityData {
    cikefailentry.EntityData.YFilter = cikefailentry.YFilter
    cikefailentry.EntityData.YangName = "cikeFailEntry"
    cikefailentry.EntityData.BundleName = "cisco_ios_xe"
    cikefailentry.EntityData.ParentYangName = "cikeFailTable"
    cikefailentry.EntityData.SegmentPath = "cikeFailEntry" + "[cikeFailIndex='" + fmt.Sprintf("%v", cikefailentry.Cikefailindex) + "']"
    cikefailentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cikefailentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cikefailentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cikefailentry.EntityData.Children = make(map[string]types.YChild)
    cikefailentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cikefailentry.EntityData.Leafs["cikeFailIndex"] = types.YLeaf{"Cikefailindex", cikefailentry.Cikefailindex}
    cikefailentry.EntityData.Leafs["cikeFailReason"] = types.YLeaf{"Cikefailreason", cikefailentry.Cikefailreason}
    cikefailentry.EntityData.Leafs["cikeFailTime"] = types.YLeaf{"Cikefailtime", cikefailentry.Cikefailtime}
    cikefailentry.EntityData.Leafs["cikeFailLocalType"] = types.YLeaf{"Cikefaillocaltype", cikefailentry.Cikefaillocaltype}
    cikefailentry.EntityData.Leafs["cikeFailLocalValue"] = types.YLeaf{"Cikefaillocalvalue", cikefailentry.Cikefaillocalvalue}
    cikefailentry.EntityData.Leafs["cikeFailRemoteType"] = types.YLeaf{"Cikefailremotetype", cikefailentry.Cikefailremotetype}
    cikefailentry.EntityData.Leafs["cikeFailRemoteValue"] = types.YLeaf{"Cikefailremotevalue", cikefailentry.Cikefailremotevalue}
    cikefailentry.EntityData.Leafs["cikeFailLocalAddr"] = types.YLeaf{"Cikefaillocaladdr", cikefailentry.Cikefaillocaladdr}
    cikefailentry.EntityData.Leafs["cikeFailRemoteAddr"] = types.YLeaf{"Cikefailremoteaddr", cikefailentry.Cikefailremoteaddr}
    return &(cikefailentry.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason represents 17 = operator requested termination.
type CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason string

const (
    CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason_other CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason = "other"

    CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason_peerDelRequest CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason = "peerDelRequest"

    CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason_peerLost CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason = "peerLost"

    CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason_localFailure CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason = "localFailure"

    CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason_authFailure CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason = "authFailure"

    CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason_hashValidation CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason = "hashValidation"

    CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason_encryptFailure CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason = "encryptFailure"

    CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason_internalError CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason = "internalError"

    CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason_sysCapExceeded CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason = "sysCapExceeded"

    CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason_proposalFailure CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason = "proposalFailure"

    CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason_peerCertUnavailable CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason = "peerCertUnavailable"

    CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason_peerCertNotValid CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason = "peerCertNotValid"

    CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason_localCertExpired CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason = "localCertExpired"

    CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason_crlFailure CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason = "crlFailure"

    CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason_peerEncodingError CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason = "peerEncodingError"

    CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason_nonExistentSa CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason = "nonExistentSa"

    CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason_operRequest CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry_Cikefailreason = "operRequest"
)

// CISCOIPSECFLOWMONITORMIB_Cipsecfailtable
// The IPsec Phase-2 Failure Table.
// This table is implemented as a sliding window 
// in which only the last n entries are maintained.  
// The maximum number of entries
// is specified by the cipSecFailTableSize object.
type CISCOIPSECFLOWMONITORMIB_Cipsecfailtable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with an IPsec Phase-1
    // failure. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry.
    Cipsecfailentry []CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry
}

func (cipsecfailtable *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable) GetEntityData() *types.CommonEntityData {
    cipsecfailtable.EntityData.YFilter = cipsecfailtable.YFilter
    cipsecfailtable.EntityData.YangName = "cipSecFailTable"
    cipsecfailtable.EntityData.BundleName = "cisco_ios_xe"
    cipsecfailtable.EntityData.ParentYangName = "CISCO-IPSEC-FLOW-MONITOR-MIB"
    cipsecfailtable.EntityData.SegmentPath = "cipSecFailTable"
    cipsecfailtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsecfailtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsecfailtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsecfailtable.EntityData.Children = make(map[string]types.YChild)
    cipsecfailtable.EntityData.Children["cipSecFailEntry"] = types.YChild{"Cipsecfailentry", nil}
    for i := range cipsecfailtable.Cipsecfailentry {
        cipsecfailtable.EntityData.Children[types.GetSegmentPath(&cipsecfailtable.Cipsecfailentry[i])] = types.YChild{"Cipsecfailentry", &cipsecfailtable.Cipsecfailentry[i]}
    }
    cipsecfailtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cipsecfailtable.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry
// Each entry contains the attributes associated with
// an IPsec Phase-1 failure.
type CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The IPsec Phase-2 Failure Table index. The value
    // of the index is a number which  begins at one and is incremented with each 
    // IPsec Phase-1 failure. The value of this object will wrap at 2,147,483,647.
    // The type is interface{} with range: 1..2147483647.
    Cipsecfailindex interface{}

    // The reason for the failure.  Possible reasons include:   1 = other   2 =
    // internal error occurred   3 = peer encoding error   4 = proposal failure  
    // 5 = protocol use failure   6 = non-existent security association   7 =
    // decryption failure   8 = encryption failure   9 = inbound authentication
    // failure  10 = outbound authentication failure  11 = compression failure  12
    // = system capacity failure  13 = peer delete request was received  14 =
    // contact with peer was lost  15 = sequence number rolled over  16 = operator
    // requested termination. The type is Cipsecfailreason.
    Cipsecfailreason interface{}

    // The value of sysUpTime in hundredths of seconds at the time of the failure.
    // The type is interface{} with range: 0..4294967295.
    Cipsecfailtime interface{}

    // The Phase-2 Tunnel index (cipSecTunIndex). The type is interface{} with
    // range: 1..2147483647.
    Cipsecfailtunnelindex interface{}

    // The security association SPI value. The type is interface{} with range:
    // 0..2147483647.
    Cipsecfailsaspi interface{}

    // The packet's source IP address. The type is string with length: 4 | 16.
    Cipsecfailpktsrcaddr interface{}

    // The packet's destination IP address. The type is string with length: 4 |
    // 16.
    Cipsecfailpktdstaddr interface{}
}

func (cipsecfailentry *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry) GetEntityData() *types.CommonEntityData {
    cipsecfailentry.EntityData.YFilter = cipsecfailentry.YFilter
    cipsecfailentry.EntityData.YangName = "cipSecFailEntry"
    cipsecfailentry.EntityData.BundleName = "cisco_ios_xe"
    cipsecfailentry.EntityData.ParentYangName = "cipSecFailTable"
    cipsecfailentry.EntityData.SegmentPath = "cipSecFailEntry" + "[cipSecFailIndex='" + fmt.Sprintf("%v", cipsecfailentry.Cipsecfailindex) + "']"
    cipsecfailentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsecfailentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsecfailentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsecfailentry.EntityData.Children = make(map[string]types.YChild)
    cipsecfailentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cipsecfailentry.EntityData.Leafs["cipSecFailIndex"] = types.YLeaf{"Cipsecfailindex", cipsecfailentry.Cipsecfailindex}
    cipsecfailentry.EntityData.Leafs["cipSecFailReason"] = types.YLeaf{"Cipsecfailreason", cipsecfailentry.Cipsecfailreason}
    cipsecfailentry.EntityData.Leafs["cipSecFailTime"] = types.YLeaf{"Cipsecfailtime", cipsecfailentry.Cipsecfailtime}
    cipsecfailentry.EntityData.Leafs["cipSecFailTunnelIndex"] = types.YLeaf{"Cipsecfailtunnelindex", cipsecfailentry.Cipsecfailtunnelindex}
    cipsecfailentry.EntityData.Leafs["cipSecFailSaSpi"] = types.YLeaf{"Cipsecfailsaspi", cipsecfailentry.Cipsecfailsaspi}
    cipsecfailentry.EntityData.Leafs["cipSecFailPktSrcAddr"] = types.YLeaf{"Cipsecfailpktsrcaddr", cipsecfailentry.Cipsecfailpktsrcaddr}
    cipsecfailentry.EntityData.Leafs["cipSecFailPktDstAddr"] = types.YLeaf{"Cipsecfailpktdstaddr", cipsecfailentry.Cipsecfailpktdstaddr}
    return &(cipsecfailentry.EntityData)
}

// CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason represents  16 = operator requested termination.
type CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason string

const (
    CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason_other CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason = "other"

    CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason_internalError CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason = "internalError"

    CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason_peerEncodingError CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason = "peerEncodingError"

    CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason_proposalFailure CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason = "proposalFailure"

    CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason_protocolUseFail CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason = "protocolUseFail"

    CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason_nonExistentSa CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason = "nonExistentSa"

    CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason_decryptFailure CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason = "decryptFailure"

    CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason_encryptFailure CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason = "encryptFailure"

    CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason_inAuthFailure CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason = "inAuthFailure"

    CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason_outAuthFailure CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason = "outAuthFailure"

    CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason_compression CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason = "compression"

    CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason_sysCapExceeded CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason = "sysCapExceeded"

    CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason_peerDelRequest CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason = "peerDelRequest"

    CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason_peerLost CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason = "peerLost"

    CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason_seqNumRollOver CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason = "seqNumRollOver"

    CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason_operRequest CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry_Cipsecfailreason = "operRequest"
)

