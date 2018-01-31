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
    parent types.Entity
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

func (cISCOIPSECFLOWMONITORMIB *CISCOIPSECFLOWMONITORMIB) GetFilter() yfilter.YFilter { return cISCOIPSECFLOWMONITORMIB.YFilter }

func (cISCOIPSECFLOWMONITORMIB *CISCOIPSECFLOWMONITORMIB) SetFilter(yf yfilter.YFilter) { cISCOIPSECFLOWMONITORMIB.YFilter = yf }

func (cISCOIPSECFLOWMONITORMIB *CISCOIPSECFLOWMONITORMIB) GetGoName(yname string) string {
    if yname == "cipSecLevels" { return "Cipseclevels" }
    if yname == "cikeGlobalStats" { return "Cikeglobalstats" }
    if yname == "cipSecGlobalStats" { return "Cipsecglobalstats" }
    if yname == "cipSecHistGlobalCntl" { return "Cipsechistglobalcntl" }
    if yname == "cipSecFailGlobalCntl" { return "Cipsecfailglobalcntl" }
    if yname == "cipSecTrapCntl" { return "Cipsectrapcntl" }
    if yname == "cikePeerTable" { return "Cikepeertable" }
    if yname == "cikeTunnelTable" { return "Ciketunneltable" }
    if yname == "cikePeerCorrTable" { return "Cikepeercorrtable" }
    if yname == "cikePhase1GWStatsTable" { return "Cikephase1Gwstatstable" }
    if yname == "cipSecTunnelTable" { return "Cipsectunneltable" }
    if yname == "cipSecEndPtTable" { return "Cipsecendpttable" }
    if yname == "cipSecSpiTable" { return "Cipsecspitable" }
    if yname == "cipSecPhase2GWStatsTable" { return "Cipsecphase2Gwstatstable" }
    if yname == "cikeTunnelHistTable" { return "Ciketunnelhisttable" }
    if yname == "cipSecTunnelHistTable" { return "Cipsectunnelhisttable" }
    if yname == "cipSecEndPtHistTable" { return "Cipsecendpthisttable" }
    if yname == "cikeFailTable" { return "Cikefailtable" }
    if yname == "cipSecFailTable" { return "Cipsecfailtable" }
    return ""
}

func (cISCOIPSECFLOWMONITORMIB *CISCOIPSECFLOWMONITORMIB) GetSegmentPath() string {
    return "CISCO-IPSEC-FLOW-MONITOR-MIB:CISCO-IPSEC-FLOW-MONITOR-MIB"
}

func (cISCOIPSECFLOWMONITORMIB *CISCOIPSECFLOWMONITORMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipSecLevels" {
        return &cISCOIPSECFLOWMONITORMIB.Cipseclevels
    }
    if childYangName == "cikeGlobalStats" {
        return &cISCOIPSECFLOWMONITORMIB.Cikeglobalstats
    }
    if childYangName == "cipSecGlobalStats" {
        return &cISCOIPSECFLOWMONITORMIB.Cipsecglobalstats
    }
    if childYangName == "cipSecHistGlobalCntl" {
        return &cISCOIPSECFLOWMONITORMIB.Cipsechistglobalcntl
    }
    if childYangName == "cipSecFailGlobalCntl" {
        return &cISCOIPSECFLOWMONITORMIB.Cipsecfailglobalcntl
    }
    if childYangName == "cipSecTrapCntl" {
        return &cISCOIPSECFLOWMONITORMIB.Cipsectrapcntl
    }
    if childYangName == "cikePeerTable" {
        return &cISCOIPSECFLOWMONITORMIB.Cikepeertable
    }
    if childYangName == "cikeTunnelTable" {
        return &cISCOIPSECFLOWMONITORMIB.Ciketunneltable
    }
    if childYangName == "cikePeerCorrTable" {
        return &cISCOIPSECFLOWMONITORMIB.Cikepeercorrtable
    }
    if childYangName == "cikePhase1GWStatsTable" {
        return &cISCOIPSECFLOWMONITORMIB.Cikephase1Gwstatstable
    }
    if childYangName == "cipSecTunnelTable" {
        return &cISCOIPSECFLOWMONITORMIB.Cipsectunneltable
    }
    if childYangName == "cipSecEndPtTable" {
        return &cISCOIPSECFLOWMONITORMIB.Cipsecendpttable
    }
    if childYangName == "cipSecSpiTable" {
        return &cISCOIPSECFLOWMONITORMIB.Cipsecspitable
    }
    if childYangName == "cipSecPhase2GWStatsTable" {
        return &cISCOIPSECFLOWMONITORMIB.Cipsecphase2Gwstatstable
    }
    if childYangName == "cikeTunnelHistTable" {
        return &cISCOIPSECFLOWMONITORMIB.Ciketunnelhisttable
    }
    if childYangName == "cipSecTunnelHistTable" {
        return &cISCOIPSECFLOWMONITORMIB.Cipsectunnelhisttable
    }
    if childYangName == "cipSecEndPtHistTable" {
        return &cISCOIPSECFLOWMONITORMIB.Cipsecendpthisttable
    }
    if childYangName == "cikeFailTable" {
        return &cISCOIPSECFLOWMONITORMIB.Cikefailtable
    }
    if childYangName == "cipSecFailTable" {
        return &cISCOIPSECFLOWMONITORMIB.Cipsecfailtable
    }
    return nil
}

func (cISCOIPSECFLOWMONITORMIB *CISCOIPSECFLOWMONITORMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cipSecLevels"] = &cISCOIPSECFLOWMONITORMIB.Cipseclevels
    children["cikeGlobalStats"] = &cISCOIPSECFLOWMONITORMIB.Cikeglobalstats
    children["cipSecGlobalStats"] = &cISCOIPSECFLOWMONITORMIB.Cipsecglobalstats
    children["cipSecHistGlobalCntl"] = &cISCOIPSECFLOWMONITORMIB.Cipsechistglobalcntl
    children["cipSecFailGlobalCntl"] = &cISCOIPSECFLOWMONITORMIB.Cipsecfailglobalcntl
    children["cipSecTrapCntl"] = &cISCOIPSECFLOWMONITORMIB.Cipsectrapcntl
    children["cikePeerTable"] = &cISCOIPSECFLOWMONITORMIB.Cikepeertable
    children["cikeTunnelTable"] = &cISCOIPSECFLOWMONITORMIB.Ciketunneltable
    children["cikePeerCorrTable"] = &cISCOIPSECFLOWMONITORMIB.Cikepeercorrtable
    children["cikePhase1GWStatsTable"] = &cISCOIPSECFLOWMONITORMIB.Cikephase1Gwstatstable
    children["cipSecTunnelTable"] = &cISCOIPSECFLOWMONITORMIB.Cipsectunneltable
    children["cipSecEndPtTable"] = &cISCOIPSECFLOWMONITORMIB.Cipsecendpttable
    children["cipSecSpiTable"] = &cISCOIPSECFLOWMONITORMIB.Cipsecspitable
    children["cipSecPhase2GWStatsTable"] = &cISCOIPSECFLOWMONITORMIB.Cipsecphase2Gwstatstable
    children["cikeTunnelHistTable"] = &cISCOIPSECFLOWMONITORMIB.Ciketunnelhisttable
    children["cipSecTunnelHistTable"] = &cISCOIPSECFLOWMONITORMIB.Cipsectunnelhisttable
    children["cipSecEndPtHistTable"] = &cISCOIPSECFLOWMONITORMIB.Cipsecendpthisttable
    children["cikeFailTable"] = &cISCOIPSECFLOWMONITORMIB.Cikefailtable
    children["cipSecFailTable"] = &cISCOIPSECFLOWMONITORMIB.Cipsecfailtable
    return children
}

func (cISCOIPSECFLOWMONITORMIB *CISCOIPSECFLOWMONITORMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOIPSECFLOWMONITORMIB *CISCOIPSECFLOWMONITORMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOIPSECFLOWMONITORMIB *CISCOIPSECFLOWMONITORMIB) GetYangName() string { return "CISCO-IPSEC-FLOW-MONITOR-MIB" }

func (cISCOIPSECFLOWMONITORMIB *CISCOIPSECFLOWMONITORMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOIPSECFLOWMONITORMIB *CISCOIPSECFLOWMONITORMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOIPSECFLOWMONITORMIB *CISCOIPSECFLOWMONITORMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOIPSECFLOWMONITORMIB *CISCOIPSECFLOWMONITORMIB) SetParent(parent types.Entity) { cISCOIPSECFLOWMONITORMIB.parent = parent }

func (cISCOIPSECFLOWMONITORMIB *CISCOIPSECFLOWMONITORMIB) GetParent() types.Entity { return cISCOIPSECFLOWMONITORMIB.parent }

func (cISCOIPSECFLOWMONITORMIB *CISCOIPSECFLOWMONITORMIB) GetParentYangName() string { return "CISCO-IPSEC-FLOW-MONITOR-MIB" }

// CISCOIPSECFLOWMONITORMIB_Cipseclevels
type CISCOIPSECFLOWMONITORMIB_Cipseclevels struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The level of the IPsec MIB. The type is interface{} with range: 1..4096.
    Cipsecmiblevel interface{}
}

func (cipseclevels *CISCOIPSECFLOWMONITORMIB_Cipseclevels) GetFilter() yfilter.YFilter { return cipseclevels.YFilter }

func (cipseclevels *CISCOIPSECFLOWMONITORMIB_Cipseclevels) SetFilter(yf yfilter.YFilter) { cipseclevels.YFilter = yf }

func (cipseclevels *CISCOIPSECFLOWMONITORMIB_Cipseclevels) GetGoName(yname string) string {
    if yname == "cipSecMibLevel" { return "Cipsecmiblevel" }
    return ""
}

func (cipseclevels *CISCOIPSECFLOWMONITORMIB_Cipseclevels) GetSegmentPath() string {
    return "cipSecLevels"
}

func (cipseclevels *CISCOIPSECFLOWMONITORMIB_Cipseclevels) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipseclevels *CISCOIPSECFLOWMONITORMIB_Cipseclevels) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipseclevels *CISCOIPSECFLOWMONITORMIB_Cipseclevels) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipSecMibLevel"] = cipseclevels.Cipsecmiblevel
    return leafs
}

func (cipseclevels *CISCOIPSECFLOWMONITORMIB_Cipseclevels) GetBundleName() string { return "cisco_ios_xe" }

func (cipseclevels *CISCOIPSECFLOWMONITORMIB_Cipseclevels) GetYangName() string { return "cipSecLevels" }

func (cipseclevels *CISCOIPSECFLOWMONITORMIB_Cipseclevels) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipseclevels *CISCOIPSECFLOWMONITORMIB_Cipseclevels) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipseclevels *CISCOIPSECFLOWMONITORMIB_Cipseclevels) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipseclevels *CISCOIPSECFLOWMONITORMIB_Cipseclevels) SetParent(parent types.Entity) { cipseclevels.parent = parent }

func (cipseclevels *CISCOIPSECFLOWMONITORMIB_Cipseclevels) GetParent() types.Entity { return cipseclevels.parent }

func (cipseclevels *CISCOIPSECFLOWMONITORMIB_Cipseclevels) GetParentYangName() string { return "CISCO-IPSEC-FLOW-MONITOR-MIB" }

// CISCOIPSECFLOWMONITORMIB_Cikeglobalstats
type CISCOIPSECFLOWMONITORMIB_Cikeglobalstats struct {
    parent types.Entity
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

func (cikeglobalstats *CISCOIPSECFLOWMONITORMIB_Cikeglobalstats) GetFilter() yfilter.YFilter { return cikeglobalstats.YFilter }

func (cikeglobalstats *CISCOIPSECFLOWMONITORMIB_Cikeglobalstats) SetFilter(yf yfilter.YFilter) { cikeglobalstats.YFilter = yf }

func (cikeglobalstats *CISCOIPSECFLOWMONITORMIB_Cikeglobalstats) GetGoName(yname string) string {
    if yname == "cikeGlobalActiveTunnels" { return "Cikeglobalactivetunnels" }
    if yname == "cikeGlobalPreviousTunnels" { return "Cikeglobalprevioustunnels" }
    if yname == "cikeGlobalInOctets" { return "Cikeglobalinoctets" }
    if yname == "cikeGlobalInPkts" { return "Cikeglobalinpkts" }
    if yname == "cikeGlobalInDropPkts" { return "Cikeglobalindroppkts" }
    if yname == "cikeGlobalInNotifys" { return "Cikeglobalinnotifys" }
    if yname == "cikeGlobalInP2Exchgs" { return "Cikeglobalinp2Exchgs" }
    if yname == "cikeGlobalInP2ExchgInvalids" { return "Cikeglobalinp2Exchginvalids" }
    if yname == "cikeGlobalInP2ExchgRejects" { return "Cikeglobalinp2Exchgrejects" }
    if yname == "cikeGlobalInP2SaDelRequests" { return "Cikeglobalinp2Sadelrequests" }
    if yname == "cikeGlobalOutOctets" { return "Cikeglobaloutoctets" }
    if yname == "cikeGlobalOutPkts" { return "Cikeglobaloutpkts" }
    if yname == "cikeGlobalOutDropPkts" { return "Cikeglobaloutdroppkts" }
    if yname == "cikeGlobalOutNotifys" { return "Cikeglobaloutnotifys" }
    if yname == "cikeGlobalOutP2Exchgs" { return "Cikeglobaloutp2Exchgs" }
    if yname == "cikeGlobalOutP2ExchgInvalids" { return "Cikeglobaloutp2Exchginvalids" }
    if yname == "cikeGlobalOutP2ExchgRejects" { return "Cikeglobaloutp2Exchgrejects" }
    if yname == "cikeGlobalOutP2SaDelRequests" { return "Cikeglobaloutp2Sadelrequests" }
    if yname == "cikeGlobalInitTunnels" { return "Cikeglobalinittunnels" }
    if yname == "cikeGlobalInitTunnelFails" { return "Cikeglobalinittunnelfails" }
    if yname == "cikeGlobalRespTunnelFails" { return "Cikeglobalresptunnelfails" }
    if yname == "cikeGlobalSysCapFails" { return "Cikeglobalsyscapfails" }
    if yname == "cikeGlobalAuthFails" { return "Cikeglobalauthfails" }
    if yname == "cikeGlobalDecryptFails" { return "Cikeglobaldecryptfails" }
    if yname == "cikeGlobalHashValidFails" { return "Cikeglobalhashvalidfails" }
    if yname == "cikeGlobalNoSaFails" { return "Cikeglobalnosafails" }
    return ""
}

func (cikeglobalstats *CISCOIPSECFLOWMONITORMIB_Cikeglobalstats) GetSegmentPath() string {
    return "cikeGlobalStats"
}

func (cikeglobalstats *CISCOIPSECFLOWMONITORMIB_Cikeglobalstats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cikeglobalstats *CISCOIPSECFLOWMONITORMIB_Cikeglobalstats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cikeglobalstats *CISCOIPSECFLOWMONITORMIB_Cikeglobalstats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cikeGlobalActiveTunnels"] = cikeglobalstats.Cikeglobalactivetunnels
    leafs["cikeGlobalPreviousTunnels"] = cikeglobalstats.Cikeglobalprevioustunnels
    leafs["cikeGlobalInOctets"] = cikeglobalstats.Cikeglobalinoctets
    leafs["cikeGlobalInPkts"] = cikeglobalstats.Cikeglobalinpkts
    leafs["cikeGlobalInDropPkts"] = cikeglobalstats.Cikeglobalindroppkts
    leafs["cikeGlobalInNotifys"] = cikeglobalstats.Cikeglobalinnotifys
    leafs["cikeGlobalInP2Exchgs"] = cikeglobalstats.Cikeglobalinp2Exchgs
    leafs["cikeGlobalInP2ExchgInvalids"] = cikeglobalstats.Cikeglobalinp2Exchginvalids
    leafs["cikeGlobalInP2ExchgRejects"] = cikeglobalstats.Cikeglobalinp2Exchgrejects
    leafs["cikeGlobalInP2SaDelRequests"] = cikeglobalstats.Cikeglobalinp2Sadelrequests
    leafs["cikeGlobalOutOctets"] = cikeglobalstats.Cikeglobaloutoctets
    leafs["cikeGlobalOutPkts"] = cikeglobalstats.Cikeglobaloutpkts
    leafs["cikeGlobalOutDropPkts"] = cikeglobalstats.Cikeglobaloutdroppkts
    leafs["cikeGlobalOutNotifys"] = cikeglobalstats.Cikeglobaloutnotifys
    leafs["cikeGlobalOutP2Exchgs"] = cikeglobalstats.Cikeglobaloutp2Exchgs
    leafs["cikeGlobalOutP2ExchgInvalids"] = cikeglobalstats.Cikeglobaloutp2Exchginvalids
    leafs["cikeGlobalOutP2ExchgRejects"] = cikeglobalstats.Cikeglobaloutp2Exchgrejects
    leafs["cikeGlobalOutP2SaDelRequests"] = cikeglobalstats.Cikeglobaloutp2Sadelrequests
    leafs["cikeGlobalInitTunnels"] = cikeglobalstats.Cikeglobalinittunnels
    leafs["cikeGlobalInitTunnelFails"] = cikeglobalstats.Cikeglobalinittunnelfails
    leafs["cikeGlobalRespTunnelFails"] = cikeglobalstats.Cikeglobalresptunnelfails
    leafs["cikeGlobalSysCapFails"] = cikeglobalstats.Cikeglobalsyscapfails
    leafs["cikeGlobalAuthFails"] = cikeglobalstats.Cikeglobalauthfails
    leafs["cikeGlobalDecryptFails"] = cikeglobalstats.Cikeglobaldecryptfails
    leafs["cikeGlobalHashValidFails"] = cikeglobalstats.Cikeglobalhashvalidfails
    leafs["cikeGlobalNoSaFails"] = cikeglobalstats.Cikeglobalnosafails
    return leafs
}

func (cikeglobalstats *CISCOIPSECFLOWMONITORMIB_Cikeglobalstats) GetBundleName() string { return "cisco_ios_xe" }

func (cikeglobalstats *CISCOIPSECFLOWMONITORMIB_Cikeglobalstats) GetYangName() string { return "cikeGlobalStats" }

func (cikeglobalstats *CISCOIPSECFLOWMONITORMIB_Cikeglobalstats) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cikeglobalstats *CISCOIPSECFLOWMONITORMIB_Cikeglobalstats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cikeglobalstats *CISCOIPSECFLOWMONITORMIB_Cikeglobalstats) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cikeglobalstats *CISCOIPSECFLOWMONITORMIB_Cikeglobalstats) SetParent(parent types.Entity) { cikeglobalstats.parent = parent }

func (cikeglobalstats *CISCOIPSECFLOWMONITORMIB_Cikeglobalstats) GetParent() types.Entity { return cikeglobalstats.parent }

func (cikeglobalstats *CISCOIPSECFLOWMONITORMIB_Cikeglobalstats) GetParentYangName() string { return "CISCO-IPSEC-FLOW-MONITOR-MIB" }

// CISCOIPSECFLOWMONITORMIB_Cipsecglobalstats
type CISCOIPSECFLOWMONITORMIB_Cipsecglobalstats struct {
    parent types.Entity
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

func (cipsecglobalstats *CISCOIPSECFLOWMONITORMIB_Cipsecglobalstats) GetFilter() yfilter.YFilter { return cipsecglobalstats.YFilter }

func (cipsecglobalstats *CISCOIPSECFLOWMONITORMIB_Cipsecglobalstats) SetFilter(yf yfilter.YFilter) { cipsecglobalstats.YFilter = yf }

func (cipsecglobalstats *CISCOIPSECFLOWMONITORMIB_Cipsecglobalstats) GetGoName(yname string) string {
    if yname == "cipSecGlobalActiveTunnels" { return "Cipsecglobalactivetunnels" }
    if yname == "cipSecGlobalPreviousTunnels" { return "Cipsecglobalprevioustunnels" }
    if yname == "cipSecGlobalInOctets" { return "Cipsecglobalinoctets" }
    if yname == "cipSecGlobalHcInOctets" { return "Cipsecglobalhcinoctets" }
    if yname == "cipSecGlobalInOctWraps" { return "Cipsecglobalinoctwraps" }
    if yname == "cipSecGlobalInDecompOctets" { return "Cipsecglobalindecompoctets" }
    if yname == "cipSecGlobalHcInDecompOctets" { return "Cipsecglobalhcindecompoctets" }
    if yname == "cipSecGlobalInDecompOctWraps" { return "Cipsecglobalindecompoctwraps" }
    if yname == "cipSecGlobalInPkts" { return "Cipsecglobalinpkts" }
    if yname == "cipSecGlobalInDrops" { return "Cipsecglobalindrops" }
    if yname == "cipSecGlobalInReplayDrops" { return "Cipsecglobalinreplaydrops" }
    if yname == "cipSecGlobalInAuths" { return "Cipsecglobalinauths" }
    if yname == "cipSecGlobalInAuthFails" { return "Cipsecglobalinauthfails" }
    if yname == "cipSecGlobalInDecrypts" { return "Cipsecglobalindecrypts" }
    if yname == "cipSecGlobalInDecryptFails" { return "Cipsecglobalindecryptfails" }
    if yname == "cipSecGlobalOutOctets" { return "Cipsecglobaloutoctets" }
    if yname == "cipSecGlobalHcOutOctets" { return "Cipsecglobalhcoutoctets" }
    if yname == "cipSecGlobalOutOctWraps" { return "Cipsecglobaloutoctwraps" }
    if yname == "cipSecGlobalOutUncompOctets" { return "Cipsecglobaloutuncompoctets" }
    if yname == "cipSecGlobalHcOutUncompOctets" { return "Cipsecglobalhcoutuncompoctets" }
    if yname == "cipSecGlobalOutUncompOctWraps" { return "Cipsecglobaloutuncompoctwraps" }
    if yname == "cipSecGlobalOutPkts" { return "Cipsecglobaloutpkts" }
    if yname == "cipSecGlobalOutDrops" { return "Cipsecglobaloutdrops" }
    if yname == "cipSecGlobalOutAuths" { return "Cipsecglobaloutauths" }
    if yname == "cipSecGlobalOutAuthFails" { return "Cipsecglobaloutauthfails" }
    if yname == "cipSecGlobalOutEncrypts" { return "Cipsecglobaloutencrypts" }
    if yname == "cipSecGlobalOutEncryptFails" { return "Cipsecglobaloutencryptfails" }
    if yname == "cipSecGlobalProtocolUseFails" { return "Cipsecglobalprotocolusefails" }
    if yname == "cipSecGlobalNoSaFails" { return "Cipsecglobalnosafails" }
    if yname == "cipSecGlobalSysCapFails" { return "Cipsecglobalsyscapfails" }
    return ""
}

func (cipsecglobalstats *CISCOIPSECFLOWMONITORMIB_Cipsecglobalstats) GetSegmentPath() string {
    return "cipSecGlobalStats"
}

func (cipsecglobalstats *CISCOIPSECFLOWMONITORMIB_Cipsecglobalstats) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipsecglobalstats *CISCOIPSECFLOWMONITORMIB_Cipsecglobalstats) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipsecglobalstats *CISCOIPSECFLOWMONITORMIB_Cipsecglobalstats) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipSecGlobalActiveTunnels"] = cipsecglobalstats.Cipsecglobalactivetunnels
    leafs["cipSecGlobalPreviousTunnels"] = cipsecglobalstats.Cipsecglobalprevioustunnels
    leafs["cipSecGlobalInOctets"] = cipsecglobalstats.Cipsecglobalinoctets
    leafs["cipSecGlobalHcInOctets"] = cipsecglobalstats.Cipsecglobalhcinoctets
    leafs["cipSecGlobalInOctWraps"] = cipsecglobalstats.Cipsecglobalinoctwraps
    leafs["cipSecGlobalInDecompOctets"] = cipsecglobalstats.Cipsecglobalindecompoctets
    leafs["cipSecGlobalHcInDecompOctets"] = cipsecglobalstats.Cipsecglobalhcindecompoctets
    leafs["cipSecGlobalInDecompOctWraps"] = cipsecglobalstats.Cipsecglobalindecompoctwraps
    leafs["cipSecGlobalInPkts"] = cipsecglobalstats.Cipsecglobalinpkts
    leafs["cipSecGlobalInDrops"] = cipsecglobalstats.Cipsecglobalindrops
    leafs["cipSecGlobalInReplayDrops"] = cipsecglobalstats.Cipsecglobalinreplaydrops
    leafs["cipSecGlobalInAuths"] = cipsecglobalstats.Cipsecglobalinauths
    leafs["cipSecGlobalInAuthFails"] = cipsecglobalstats.Cipsecglobalinauthfails
    leafs["cipSecGlobalInDecrypts"] = cipsecglobalstats.Cipsecglobalindecrypts
    leafs["cipSecGlobalInDecryptFails"] = cipsecglobalstats.Cipsecglobalindecryptfails
    leafs["cipSecGlobalOutOctets"] = cipsecglobalstats.Cipsecglobaloutoctets
    leafs["cipSecGlobalHcOutOctets"] = cipsecglobalstats.Cipsecglobalhcoutoctets
    leafs["cipSecGlobalOutOctWraps"] = cipsecglobalstats.Cipsecglobaloutoctwraps
    leafs["cipSecGlobalOutUncompOctets"] = cipsecglobalstats.Cipsecglobaloutuncompoctets
    leafs["cipSecGlobalHcOutUncompOctets"] = cipsecglobalstats.Cipsecglobalhcoutuncompoctets
    leafs["cipSecGlobalOutUncompOctWraps"] = cipsecglobalstats.Cipsecglobaloutuncompoctwraps
    leafs["cipSecGlobalOutPkts"] = cipsecglobalstats.Cipsecglobaloutpkts
    leafs["cipSecGlobalOutDrops"] = cipsecglobalstats.Cipsecglobaloutdrops
    leafs["cipSecGlobalOutAuths"] = cipsecglobalstats.Cipsecglobaloutauths
    leafs["cipSecGlobalOutAuthFails"] = cipsecglobalstats.Cipsecglobaloutauthfails
    leafs["cipSecGlobalOutEncrypts"] = cipsecglobalstats.Cipsecglobaloutencrypts
    leafs["cipSecGlobalOutEncryptFails"] = cipsecglobalstats.Cipsecglobaloutencryptfails
    leafs["cipSecGlobalProtocolUseFails"] = cipsecglobalstats.Cipsecglobalprotocolusefails
    leafs["cipSecGlobalNoSaFails"] = cipsecglobalstats.Cipsecglobalnosafails
    leafs["cipSecGlobalSysCapFails"] = cipsecglobalstats.Cipsecglobalsyscapfails
    return leafs
}

func (cipsecglobalstats *CISCOIPSECFLOWMONITORMIB_Cipsecglobalstats) GetBundleName() string { return "cisco_ios_xe" }

func (cipsecglobalstats *CISCOIPSECFLOWMONITORMIB_Cipsecglobalstats) GetYangName() string { return "cipSecGlobalStats" }

func (cipsecglobalstats *CISCOIPSECFLOWMONITORMIB_Cipsecglobalstats) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsecglobalstats *CISCOIPSECFLOWMONITORMIB_Cipsecglobalstats) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsecglobalstats *CISCOIPSECFLOWMONITORMIB_Cipsecglobalstats) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsecglobalstats *CISCOIPSECFLOWMONITORMIB_Cipsecglobalstats) SetParent(parent types.Entity) { cipsecglobalstats.parent = parent }

func (cipsecglobalstats *CISCOIPSECFLOWMONITORMIB_Cipsecglobalstats) GetParent() types.Entity { return cipsecglobalstats.parent }

func (cipsecglobalstats *CISCOIPSECFLOWMONITORMIB_Cipsecglobalstats) GetParentYangName() string { return "CISCO-IPSEC-FLOW-MONITOR-MIB" }

// CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl
type CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl struct {
    parent types.Entity
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

func (cipsechistglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl) GetFilter() yfilter.YFilter { return cipsechistglobalcntl.YFilter }

func (cipsechistglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl) SetFilter(yf yfilter.YFilter) { cipsechistglobalcntl.YFilter = yf }

func (cipsechistglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl) GetGoName(yname string) string {
    if yname == "cipSecHistTableSize" { return "Cipsechisttablesize" }
    if yname == "cipSecHistCheckPoint" { return "Cipsechistcheckpoint" }
    return ""
}

func (cipsechistglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl) GetSegmentPath() string {
    return "cipSecHistGlobalCntl"
}

func (cipsechistglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipsechistglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipsechistglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipSecHistTableSize"] = cipsechistglobalcntl.Cipsechisttablesize
    leafs["cipSecHistCheckPoint"] = cipsechistglobalcntl.Cipsechistcheckpoint
    return leafs
}

func (cipsechistglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl) GetBundleName() string { return "cisco_ios_xe" }

func (cipsechistglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl) GetYangName() string { return "cipSecHistGlobalCntl" }

func (cipsechistglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsechistglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsechistglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsechistglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl) SetParent(parent types.Entity) { cipsechistglobalcntl.parent = parent }

func (cipsechistglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl) GetParent() types.Entity { return cipsechistglobalcntl.parent }

func (cipsechistglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl) GetParentYangName() string { return "CISCO-IPSEC-FLOW-MONITOR-MIB" }

// CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl_Cipsechistcheckpoint represents    for each active IPsec Phase-2 Tunnel.
type CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl_Cipsechistcheckpoint string

const (
    CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl_Cipsechistcheckpoint_ready CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl_Cipsechistcheckpoint = "ready"

    CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl_Cipsechistcheckpoint_checkPoint CISCOIPSECFLOWMONITORMIB_Cipsechistglobalcntl_Cipsechistcheckpoint = "checkPoint"
)

// CISCOIPSECFLOWMONITORMIB_Cipsecfailglobalcntl
type CISCOIPSECFLOWMONITORMIB_Cipsecfailglobalcntl struct {
    parent types.Entity
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

func (cipsecfailglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsecfailglobalcntl) GetFilter() yfilter.YFilter { return cipsecfailglobalcntl.YFilter }

func (cipsecfailglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsecfailglobalcntl) SetFilter(yf yfilter.YFilter) { cipsecfailglobalcntl.YFilter = yf }

func (cipsecfailglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsecfailglobalcntl) GetGoName(yname string) string {
    if yname == "cipSecFailTableSize" { return "Cipsecfailtablesize" }
    return ""
}

func (cipsecfailglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsecfailglobalcntl) GetSegmentPath() string {
    return "cipSecFailGlobalCntl"
}

func (cipsecfailglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsecfailglobalcntl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipsecfailglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsecfailglobalcntl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipsecfailglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsecfailglobalcntl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipSecFailTableSize"] = cipsecfailglobalcntl.Cipsecfailtablesize
    return leafs
}

func (cipsecfailglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsecfailglobalcntl) GetBundleName() string { return "cisco_ios_xe" }

func (cipsecfailglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsecfailglobalcntl) GetYangName() string { return "cipSecFailGlobalCntl" }

func (cipsecfailglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsecfailglobalcntl) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsecfailglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsecfailglobalcntl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsecfailglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsecfailglobalcntl) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsecfailglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsecfailglobalcntl) SetParent(parent types.Entity) { cipsecfailglobalcntl.parent = parent }

func (cipsecfailglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsecfailglobalcntl) GetParent() types.Entity { return cipsecfailglobalcntl.parent }

func (cipsecfailglobalcntl *CISCOIPSECFLOWMONITORMIB_Cipsecfailglobalcntl) GetParentYangName() string { return "CISCO-IPSEC-FLOW-MONITOR-MIB" }

// CISCOIPSECFLOWMONITORMIB_Cipsectrapcntl
type CISCOIPSECFLOWMONITORMIB_Cipsectrapcntl struct {
    parent types.Entity
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

func (cipsectrapcntl *CISCOIPSECFLOWMONITORMIB_Cipsectrapcntl) GetFilter() yfilter.YFilter { return cipsectrapcntl.YFilter }

func (cipsectrapcntl *CISCOIPSECFLOWMONITORMIB_Cipsectrapcntl) SetFilter(yf yfilter.YFilter) { cipsectrapcntl.YFilter = yf }

func (cipsectrapcntl *CISCOIPSECFLOWMONITORMIB_Cipsectrapcntl) GetGoName(yname string) string {
    if yname == "cipSecTrapCntlIkeTunnelStart" { return "Cipsectrapcntliketunnelstart" }
    if yname == "cipSecTrapCntlIkeTunnelStop" { return "Cipsectrapcntliketunnelstop" }
    if yname == "cipSecTrapCntlIkeSysFailure" { return "Cipsectrapcntlikesysfailure" }
    if yname == "cipSecTrapCntlIkeCertCrlFailure" { return "Cipsectrapcntlikecertcrlfailure" }
    if yname == "cipSecTrapCntlIkeProtocolFail" { return "Cipsectrapcntlikeprotocolfail" }
    if yname == "cipSecTrapCntlIkeNoSa" { return "Cipsectrapcntlikenosa" }
    if yname == "cipSecTrapCntlIpSecTunnelStart" { return "Cipsectrapcntlipsectunnelstart" }
    if yname == "cipSecTrapCntlIpSecTunnelStop" { return "Cipsectrapcntlipsectunnelstop" }
    if yname == "cipSecTrapCntlIpSecSysFailure" { return "Cipsectrapcntlipsecsysfailure" }
    if yname == "cipSecTrapCntlIpSecSetUpFailure" { return "Cipsectrapcntlipsecsetupfailure" }
    if yname == "cipSecTrapCntlIpSecEarlyTunTerm" { return "Cipsectrapcntlipsecearlytunterm" }
    if yname == "cipSecTrapCntlIpSecProtocolFail" { return "Cipsectrapcntlipsecprotocolfail" }
    if yname == "cipSecTrapCntlIpSecNoSa" { return "Cipsectrapcntlipsecnosa" }
    return ""
}

func (cipsectrapcntl *CISCOIPSECFLOWMONITORMIB_Cipsectrapcntl) GetSegmentPath() string {
    return "cipSecTrapCntl"
}

func (cipsectrapcntl *CISCOIPSECFLOWMONITORMIB_Cipsectrapcntl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipsectrapcntl *CISCOIPSECFLOWMONITORMIB_Cipsectrapcntl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipsectrapcntl *CISCOIPSECFLOWMONITORMIB_Cipsectrapcntl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipSecTrapCntlIkeTunnelStart"] = cipsectrapcntl.Cipsectrapcntliketunnelstart
    leafs["cipSecTrapCntlIkeTunnelStop"] = cipsectrapcntl.Cipsectrapcntliketunnelstop
    leafs["cipSecTrapCntlIkeSysFailure"] = cipsectrapcntl.Cipsectrapcntlikesysfailure
    leafs["cipSecTrapCntlIkeCertCrlFailure"] = cipsectrapcntl.Cipsectrapcntlikecertcrlfailure
    leafs["cipSecTrapCntlIkeProtocolFail"] = cipsectrapcntl.Cipsectrapcntlikeprotocolfail
    leafs["cipSecTrapCntlIkeNoSa"] = cipsectrapcntl.Cipsectrapcntlikenosa
    leafs["cipSecTrapCntlIpSecTunnelStart"] = cipsectrapcntl.Cipsectrapcntlipsectunnelstart
    leafs["cipSecTrapCntlIpSecTunnelStop"] = cipsectrapcntl.Cipsectrapcntlipsectunnelstop
    leafs["cipSecTrapCntlIpSecSysFailure"] = cipsectrapcntl.Cipsectrapcntlipsecsysfailure
    leafs["cipSecTrapCntlIpSecSetUpFailure"] = cipsectrapcntl.Cipsectrapcntlipsecsetupfailure
    leafs["cipSecTrapCntlIpSecEarlyTunTerm"] = cipsectrapcntl.Cipsectrapcntlipsecearlytunterm
    leafs["cipSecTrapCntlIpSecProtocolFail"] = cipsectrapcntl.Cipsectrapcntlipsecprotocolfail
    leafs["cipSecTrapCntlIpSecNoSa"] = cipsectrapcntl.Cipsectrapcntlipsecnosa
    return leafs
}

func (cipsectrapcntl *CISCOIPSECFLOWMONITORMIB_Cipsectrapcntl) GetBundleName() string { return "cisco_ios_xe" }

func (cipsectrapcntl *CISCOIPSECFLOWMONITORMIB_Cipsectrapcntl) GetYangName() string { return "cipSecTrapCntl" }

func (cipsectrapcntl *CISCOIPSECFLOWMONITORMIB_Cipsectrapcntl) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsectrapcntl *CISCOIPSECFLOWMONITORMIB_Cipsectrapcntl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsectrapcntl *CISCOIPSECFLOWMONITORMIB_Cipsectrapcntl) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsectrapcntl *CISCOIPSECFLOWMONITORMIB_Cipsectrapcntl) SetParent(parent types.Entity) { cipsectrapcntl.parent = parent }

func (cipsectrapcntl *CISCOIPSECFLOWMONITORMIB_Cipsectrapcntl) GetParent() types.Entity { return cipsectrapcntl.parent }

func (cipsectrapcntl *CISCOIPSECFLOWMONITORMIB_Cipsectrapcntl) GetParentYangName() string { return "CISCO-IPSEC-FLOW-MONITOR-MIB" }

// CISCOIPSECFLOWMONITORMIB_Cikepeertable
// The IPsec Phase-1 Internet Key Exchange Peer Table.
// There is one entry in this table for each IPsec
// Phase-1 IKE peer association which is currently
// associated with an active IPsec Phase-1 Tunnel.
// The IPsec Phase-1 IKE Tunnel associated with this
// IPsec Phase-1 IKE peer association may or may not
// be currently active.
type CISCOIPSECFLOWMONITORMIB_Cikepeertable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with an IPsec Phase-1 IKE
    // peer association. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_Cikepeertable_Cikepeerentry.
    Cikepeerentry []CISCOIPSECFLOWMONITORMIB_Cikepeertable_Cikepeerentry
}

func (cikepeertable *CISCOIPSECFLOWMONITORMIB_Cikepeertable) GetFilter() yfilter.YFilter { return cikepeertable.YFilter }

func (cikepeertable *CISCOIPSECFLOWMONITORMIB_Cikepeertable) SetFilter(yf yfilter.YFilter) { cikepeertable.YFilter = yf }

func (cikepeertable *CISCOIPSECFLOWMONITORMIB_Cikepeertable) GetGoName(yname string) string {
    if yname == "cikePeerEntry" { return "Cikepeerentry" }
    return ""
}

func (cikepeertable *CISCOIPSECFLOWMONITORMIB_Cikepeertable) GetSegmentPath() string {
    return "cikePeerTable"
}

func (cikepeertable *CISCOIPSECFLOWMONITORMIB_Cikepeertable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cikePeerEntry" {
        for _, c := range cikepeertable.Cikepeerentry {
            if cikepeertable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSECFLOWMONITORMIB_Cikepeertable_Cikepeerentry{}
        cikepeertable.Cikepeerentry = append(cikepeertable.Cikepeerentry, child)
        return &cikepeertable.Cikepeerentry[len(cikepeertable.Cikepeerentry)-1]
    }
    return nil
}

func (cikepeertable *CISCOIPSECFLOWMONITORMIB_Cikepeertable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cikepeertable.Cikepeerentry {
        children[cikepeertable.Cikepeerentry[i].GetSegmentPath()] = &cikepeertable.Cikepeerentry[i]
    }
    return children
}

func (cikepeertable *CISCOIPSECFLOWMONITORMIB_Cikepeertable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cikepeertable *CISCOIPSECFLOWMONITORMIB_Cikepeertable) GetBundleName() string { return "cisco_ios_xe" }

func (cikepeertable *CISCOIPSECFLOWMONITORMIB_Cikepeertable) GetYangName() string { return "cikePeerTable" }

func (cikepeertable *CISCOIPSECFLOWMONITORMIB_Cikepeertable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cikepeertable *CISCOIPSECFLOWMONITORMIB_Cikepeertable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cikepeertable *CISCOIPSECFLOWMONITORMIB_Cikepeertable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cikepeertable *CISCOIPSECFLOWMONITORMIB_Cikepeertable) SetParent(parent types.Entity) { cikepeertable.parent = parent }

func (cikepeertable *CISCOIPSECFLOWMONITORMIB_Cikepeertable) GetParent() types.Entity { return cikepeertable.parent }

func (cikepeertable *CISCOIPSECFLOWMONITORMIB_Cikepeertable) GetParentYangName() string { return "CISCO-IPSEC-FLOW-MONITOR-MIB" }

// CISCOIPSECFLOWMONITORMIB_Cikepeertable_Cikepeerentry
// Each entry contains the attributes associated
// with an IPsec Phase-1 IKE peer association.
type CISCOIPSECFLOWMONITORMIB_Cikepeertable_Cikepeerentry struct {
    parent types.Entity
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

func (cikepeerentry *CISCOIPSECFLOWMONITORMIB_Cikepeertable_Cikepeerentry) GetFilter() yfilter.YFilter { return cikepeerentry.YFilter }

func (cikepeerentry *CISCOIPSECFLOWMONITORMIB_Cikepeertable_Cikepeerentry) SetFilter(yf yfilter.YFilter) { cikepeerentry.YFilter = yf }

func (cikepeerentry *CISCOIPSECFLOWMONITORMIB_Cikepeertable_Cikepeerentry) GetGoName(yname string) string {
    if yname == "cikePeerLocalType" { return "Cikepeerlocaltype" }
    if yname == "cikePeerLocalValue" { return "Cikepeerlocalvalue" }
    if yname == "cikePeerRemoteType" { return "Cikepeerremotetype" }
    if yname == "cikePeerRemoteValue" { return "Cikepeerremotevalue" }
    if yname == "cikePeerIntIndex" { return "Cikepeerintindex" }
    if yname == "cikePeerLocalAddr" { return "Cikepeerlocaladdr" }
    if yname == "cikePeerRemoteAddr" { return "Cikepeerremoteaddr" }
    if yname == "cikePeerActiveTime" { return "Cikepeeractivetime" }
    if yname == "cikePeerActiveTunnelIndex" { return "Cikepeeractivetunnelindex" }
    return ""
}

func (cikepeerentry *CISCOIPSECFLOWMONITORMIB_Cikepeertable_Cikepeerentry) GetSegmentPath() string {
    return "cikePeerEntry" + "[cikePeerLocalType='" + fmt.Sprintf("%v", cikepeerentry.Cikepeerlocaltype) + "']" + "[cikePeerLocalValue='" + fmt.Sprintf("%v", cikepeerentry.Cikepeerlocalvalue) + "']" + "[cikePeerRemoteType='" + fmt.Sprintf("%v", cikepeerentry.Cikepeerremotetype) + "']" + "[cikePeerRemoteValue='" + fmt.Sprintf("%v", cikepeerentry.Cikepeerremotevalue) + "']" + "[cikePeerIntIndex='" + fmt.Sprintf("%v", cikepeerentry.Cikepeerintindex) + "']"
}

func (cikepeerentry *CISCOIPSECFLOWMONITORMIB_Cikepeertable_Cikepeerentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cikepeerentry *CISCOIPSECFLOWMONITORMIB_Cikepeertable_Cikepeerentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cikepeerentry *CISCOIPSECFLOWMONITORMIB_Cikepeertable_Cikepeerentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cikePeerLocalType"] = cikepeerentry.Cikepeerlocaltype
    leafs["cikePeerLocalValue"] = cikepeerentry.Cikepeerlocalvalue
    leafs["cikePeerRemoteType"] = cikepeerentry.Cikepeerremotetype
    leafs["cikePeerRemoteValue"] = cikepeerentry.Cikepeerremotevalue
    leafs["cikePeerIntIndex"] = cikepeerentry.Cikepeerintindex
    leafs["cikePeerLocalAddr"] = cikepeerentry.Cikepeerlocaladdr
    leafs["cikePeerRemoteAddr"] = cikepeerentry.Cikepeerremoteaddr
    leafs["cikePeerActiveTime"] = cikepeerentry.Cikepeeractivetime
    leafs["cikePeerActiveTunnelIndex"] = cikepeerentry.Cikepeeractivetunnelindex
    return leafs
}

func (cikepeerentry *CISCOIPSECFLOWMONITORMIB_Cikepeertable_Cikepeerentry) GetBundleName() string { return "cisco_ios_xe" }

func (cikepeerentry *CISCOIPSECFLOWMONITORMIB_Cikepeertable_Cikepeerentry) GetYangName() string { return "cikePeerEntry" }

func (cikepeerentry *CISCOIPSECFLOWMONITORMIB_Cikepeertable_Cikepeerentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cikepeerentry *CISCOIPSECFLOWMONITORMIB_Cikepeertable_Cikepeerentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cikepeerentry *CISCOIPSECFLOWMONITORMIB_Cikepeertable_Cikepeerentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cikepeerentry *CISCOIPSECFLOWMONITORMIB_Cikepeertable_Cikepeerentry) SetParent(parent types.Entity) { cikepeerentry.parent = parent }

func (cikepeerentry *CISCOIPSECFLOWMONITORMIB_Cikepeertable_Cikepeerentry) GetParent() types.Entity { return cikepeerentry.parent }

func (cikepeerentry *CISCOIPSECFLOWMONITORMIB_Cikepeertable_Cikepeerentry) GetParentYangName() string { return "cikePeerTable" }

// CISCOIPSECFLOWMONITORMIB_Ciketunneltable
// The IPsec Phase-1 Internet Key Exchange Tunnel Table.
// There is one entry in this table for each active IPsec
// Phase-1 IKE Tunnel.
type CISCOIPSECFLOWMONITORMIB_Ciketunneltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with an active IPsec Phase-1
    // IKE Tunnel. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_Ciketunneltable_Ciketunnelentry.
    Ciketunnelentry []CISCOIPSECFLOWMONITORMIB_Ciketunneltable_Ciketunnelentry
}

func (ciketunneltable *CISCOIPSECFLOWMONITORMIB_Ciketunneltable) GetFilter() yfilter.YFilter { return ciketunneltable.YFilter }

func (ciketunneltable *CISCOIPSECFLOWMONITORMIB_Ciketunneltable) SetFilter(yf yfilter.YFilter) { ciketunneltable.YFilter = yf }

func (ciketunneltable *CISCOIPSECFLOWMONITORMIB_Ciketunneltable) GetGoName(yname string) string {
    if yname == "cikeTunnelEntry" { return "Ciketunnelentry" }
    return ""
}

func (ciketunneltable *CISCOIPSECFLOWMONITORMIB_Ciketunneltable) GetSegmentPath() string {
    return "cikeTunnelTable"
}

func (ciketunneltable *CISCOIPSECFLOWMONITORMIB_Ciketunneltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cikeTunnelEntry" {
        for _, c := range ciketunneltable.Ciketunnelentry {
            if ciketunneltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSECFLOWMONITORMIB_Ciketunneltable_Ciketunnelentry{}
        ciketunneltable.Ciketunnelentry = append(ciketunneltable.Ciketunnelentry, child)
        return &ciketunneltable.Ciketunnelentry[len(ciketunneltable.Ciketunnelentry)-1]
    }
    return nil
}

func (ciketunneltable *CISCOIPSECFLOWMONITORMIB_Ciketunneltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ciketunneltable.Ciketunnelentry {
        children[ciketunneltable.Ciketunnelentry[i].GetSegmentPath()] = &ciketunneltable.Ciketunnelentry[i]
    }
    return children
}

func (ciketunneltable *CISCOIPSECFLOWMONITORMIB_Ciketunneltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciketunneltable *CISCOIPSECFLOWMONITORMIB_Ciketunneltable) GetBundleName() string { return "cisco_ios_xe" }

func (ciketunneltable *CISCOIPSECFLOWMONITORMIB_Ciketunneltable) GetYangName() string { return "cikeTunnelTable" }

func (ciketunneltable *CISCOIPSECFLOWMONITORMIB_Ciketunneltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciketunneltable *CISCOIPSECFLOWMONITORMIB_Ciketunneltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciketunneltable *CISCOIPSECFLOWMONITORMIB_Ciketunneltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciketunneltable *CISCOIPSECFLOWMONITORMIB_Ciketunneltable) SetParent(parent types.Entity) { ciketunneltable.parent = parent }

func (ciketunneltable *CISCOIPSECFLOWMONITORMIB_Ciketunneltable) GetParent() types.Entity { return ciketunneltable.parent }

func (ciketunneltable *CISCOIPSECFLOWMONITORMIB_Ciketunneltable) GetParentYangName() string { return "CISCO-IPSEC-FLOW-MONITOR-MIB" }

// CISCOIPSECFLOWMONITORMIB_Ciketunneltable_Ciketunnelentry
// Each entry contains the attributes associated with
// an active IPsec Phase-1 IKE Tunnel.
type CISCOIPSECFLOWMONITORMIB_Ciketunneltable_Ciketunnelentry struct {
    parent types.Entity
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

func (ciketunnelentry *CISCOIPSECFLOWMONITORMIB_Ciketunneltable_Ciketunnelentry) GetFilter() yfilter.YFilter { return ciketunnelentry.YFilter }

func (ciketunnelentry *CISCOIPSECFLOWMONITORMIB_Ciketunneltable_Ciketunnelentry) SetFilter(yf yfilter.YFilter) { ciketunnelentry.YFilter = yf }

func (ciketunnelentry *CISCOIPSECFLOWMONITORMIB_Ciketunneltable_Ciketunnelentry) GetGoName(yname string) string {
    if yname == "cikeTunIndex" { return "Ciketunindex" }
    if yname == "cikeTunLocalType" { return "Ciketunlocaltype" }
    if yname == "cikeTunLocalValue" { return "Ciketunlocalvalue" }
    if yname == "cikeTunLocalAddr" { return "Ciketunlocaladdr" }
    if yname == "cikeTunLocalName" { return "Ciketunlocalname" }
    if yname == "cikeTunRemoteType" { return "Ciketunremotetype" }
    if yname == "cikeTunRemoteValue" { return "Ciketunremotevalue" }
    if yname == "cikeTunRemoteAddr" { return "Ciketunremoteaddr" }
    if yname == "cikeTunRemoteName" { return "Ciketunremotename" }
    if yname == "cikeTunNegoMode" { return "Ciketunnegomode" }
    if yname == "cikeTunDiffHellmanGrp" { return "Ciketundiffhellmangrp" }
    if yname == "cikeTunEncryptAlgo" { return "Ciketunencryptalgo" }
    if yname == "cikeTunHashAlgo" { return "Ciketunhashalgo" }
    if yname == "cikeTunAuthMethod" { return "Ciketunauthmethod" }
    if yname == "cikeTunLifeTime" { return "Ciketunlifetime" }
    if yname == "cikeTunActiveTime" { return "Ciketunactivetime" }
    if yname == "cikeTunSaRefreshThreshold" { return "Ciketunsarefreshthreshold" }
    if yname == "cikeTunTotalRefreshes" { return "Ciketuntotalrefreshes" }
    if yname == "cikeTunInOctets" { return "Ciketuninoctets" }
    if yname == "cikeTunInPkts" { return "Ciketuninpkts" }
    if yname == "cikeTunInDropPkts" { return "Ciketunindroppkts" }
    if yname == "cikeTunInNotifys" { return "Ciketuninnotifys" }
    if yname == "cikeTunInP2Exchgs" { return "Ciketuninp2Exchgs" }
    if yname == "cikeTunInP2ExchgInvalids" { return "Ciketuninp2Exchginvalids" }
    if yname == "cikeTunInP2ExchgRejects" { return "Ciketuninp2Exchgrejects" }
    if yname == "cikeTunInP2SaDelRequests" { return "Ciketuninp2Sadelrequests" }
    if yname == "cikeTunOutOctets" { return "Ciketunoutoctets" }
    if yname == "cikeTunOutPkts" { return "Ciketunoutpkts" }
    if yname == "cikeTunOutDropPkts" { return "Ciketunoutdroppkts" }
    if yname == "cikeTunOutNotifys" { return "Ciketunoutnotifys" }
    if yname == "cikeTunOutP2Exchgs" { return "Ciketunoutp2Exchgs" }
    if yname == "cikeTunOutP2ExchgInvalids" { return "Ciketunoutp2Exchginvalids" }
    if yname == "cikeTunOutP2ExchgRejects" { return "Ciketunoutp2Exchgrejects" }
    if yname == "cikeTunOutP2SaDelRequests" { return "Ciketunoutp2Sadelrequests" }
    if yname == "cikeTunStatus" { return "Ciketunstatus" }
    return ""
}

func (ciketunnelentry *CISCOIPSECFLOWMONITORMIB_Ciketunneltable_Ciketunnelentry) GetSegmentPath() string {
    return "cikeTunnelEntry" + "[cikeTunIndex='" + fmt.Sprintf("%v", ciketunnelentry.Ciketunindex) + "']"
}

func (ciketunnelentry *CISCOIPSECFLOWMONITORMIB_Ciketunneltable_Ciketunnelentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciketunnelentry *CISCOIPSECFLOWMONITORMIB_Ciketunneltable_Ciketunnelentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciketunnelentry *CISCOIPSECFLOWMONITORMIB_Ciketunneltable_Ciketunnelentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cikeTunIndex"] = ciketunnelentry.Ciketunindex
    leafs["cikeTunLocalType"] = ciketunnelentry.Ciketunlocaltype
    leafs["cikeTunLocalValue"] = ciketunnelentry.Ciketunlocalvalue
    leafs["cikeTunLocalAddr"] = ciketunnelentry.Ciketunlocaladdr
    leafs["cikeTunLocalName"] = ciketunnelentry.Ciketunlocalname
    leafs["cikeTunRemoteType"] = ciketunnelentry.Ciketunremotetype
    leafs["cikeTunRemoteValue"] = ciketunnelentry.Ciketunremotevalue
    leafs["cikeTunRemoteAddr"] = ciketunnelentry.Ciketunremoteaddr
    leafs["cikeTunRemoteName"] = ciketunnelentry.Ciketunremotename
    leafs["cikeTunNegoMode"] = ciketunnelentry.Ciketunnegomode
    leafs["cikeTunDiffHellmanGrp"] = ciketunnelentry.Ciketundiffhellmangrp
    leafs["cikeTunEncryptAlgo"] = ciketunnelentry.Ciketunencryptalgo
    leafs["cikeTunHashAlgo"] = ciketunnelentry.Ciketunhashalgo
    leafs["cikeTunAuthMethod"] = ciketunnelentry.Ciketunauthmethod
    leafs["cikeTunLifeTime"] = ciketunnelentry.Ciketunlifetime
    leafs["cikeTunActiveTime"] = ciketunnelentry.Ciketunactivetime
    leafs["cikeTunSaRefreshThreshold"] = ciketunnelentry.Ciketunsarefreshthreshold
    leafs["cikeTunTotalRefreshes"] = ciketunnelentry.Ciketuntotalrefreshes
    leafs["cikeTunInOctets"] = ciketunnelentry.Ciketuninoctets
    leafs["cikeTunInPkts"] = ciketunnelentry.Ciketuninpkts
    leafs["cikeTunInDropPkts"] = ciketunnelentry.Ciketunindroppkts
    leafs["cikeTunInNotifys"] = ciketunnelentry.Ciketuninnotifys
    leafs["cikeTunInP2Exchgs"] = ciketunnelentry.Ciketuninp2Exchgs
    leafs["cikeTunInP2ExchgInvalids"] = ciketunnelentry.Ciketuninp2Exchginvalids
    leafs["cikeTunInP2ExchgRejects"] = ciketunnelentry.Ciketuninp2Exchgrejects
    leafs["cikeTunInP2SaDelRequests"] = ciketunnelentry.Ciketuninp2Sadelrequests
    leafs["cikeTunOutOctets"] = ciketunnelentry.Ciketunoutoctets
    leafs["cikeTunOutPkts"] = ciketunnelentry.Ciketunoutpkts
    leafs["cikeTunOutDropPkts"] = ciketunnelentry.Ciketunoutdroppkts
    leafs["cikeTunOutNotifys"] = ciketunnelentry.Ciketunoutnotifys
    leafs["cikeTunOutP2Exchgs"] = ciketunnelentry.Ciketunoutp2Exchgs
    leafs["cikeTunOutP2ExchgInvalids"] = ciketunnelentry.Ciketunoutp2Exchginvalids
    leafs["cikeTunOutP2ExchgRejects"] = ciketunnelentry.Ciketunoutp2Exchgrejects
    leafs["cikeTunOutP2SaDelRequests"] = ciketunnelentry.Ciketunoutp2Sadelrequests
    leafs["cikeTunStatus"] = ciketunnelentry.Ciketunstatus
    return leafs
}

func (ciketunnelentry *CISCOIPSECFLOWMONITORMIB_Ciketunneltable_Ciketunnelentry) GetBundleName() string { return "cisco_ios_xe" }

func (ciketunnelentry *CISCOIPSECFLOWMONITORMIB_Ciketunneltable_Ciketunnelentry) GetYangName() string { return "cikeTunnelEntry" }

func (ciketunnelentry *CISCOIPSECFLOWMONITORMIB_Ciketunneltable_Ciketunnelentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciketunnelentry *CISCOIPSECFLOWMONITORMIB_Ciketunneltable_Ciketunnelentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciketunnelentry *CISCOIPSECFLOWMONITORMIB_Ciketunneltable_Ciketunnelentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciketunnelentry *CISCOIPSECFLOWMONITORMIB_Ciketunneltable_Ciketunnelentry) SetParent(parent types.Entity) { ciketunnelentry.parent = parent }

func (ciketunnelentry *CISCOIPSECFLOWMONITORMIB_Ciketunneltable_Ciketunnelentry) GetParent() types.Entity { return ciketunnelentry.parent }

func (ciketunnelentry *CISCOIPSECFLOWMONITORMIB_Ciketunneltable_Ciketunnelentry) GetParentYangName() string { return "cikeTunnelTable" }

// CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable
// The IPsec Phase-1 Internet Key Exchange Peer
// Association to IPsec Phase-2 Tunnel
// Correlation Table. There is one entry in
// this table for each active IPsec Phase-2
// Tunnel.
type CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains the attributes of an IPsec Phase-1 IKE Peer Association
    // to IPsec Phase-2 Tunnel Correlation. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable_Cikepeercorrentry.
    Cikepeercorrentry []CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable_Cikepeercorrentry
}

func (cikepeercorrtable *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable) GetFilter() yfilter.YFilter { return cikepeercorrtable.YFilter }

func (cikepeercorrtable *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable) SetFilter(yf yfilter.YFilter) { cikepeercorrtable.YFilter = yf }

func (cikepeercorrtable *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable) GetGoName(yname string) string {
    if yname == "cikePeerCorrEntry" { return "Cikepeercorrentry" }
    return ""
}

func (cikepeercorrtable *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable) GetSegmentPath() string {
    return "cikePeerCorrTable"
}

func (cikepeercorrtable *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cikePeerCorrEntry" {
        for _, c := range cikepeercorrtable.Cikepeercorrentry {
            if cikepeercorrtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable_Cikepeercorrentry{}
        cikepeercorrtable.Cikepeercorrentry = append(cikepeercorrtable.Cikepeercorrentry, child)
        return &cikepeercorrtable.Cikepeercorrentry[len(cikepeercorrtable.Cikepeercorrentry)-1]
    }
    return nil
}

func (cikepeercorrtable *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cikepeercorrtable.Cikepeercorrentry {
        children[cikepeercorrtable.Cikepeercorrentry[i].GetSegmentPath()] = &cikepeercorrtable.Cikepeercorrentry[i]
    }
    return children
}

func (cikepeercorrtable *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cikepeercorrtable *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable) GetBundleName() string { return "cisco_ios_xe" }

func (cikepeercorrtable *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable) GetYangName() string { return "cikePeerCorrTable" }

func (cikepeercorrtable *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cikepeercorrtable *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cikepeercorrtable *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cikepeercorrtable *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable) SetParent(parent types.Entity) { cikepeercorrtable.parent = parent }

func (cikepeercorrtable *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable) GetParent() types.Entity { return cikepeercorrtable.parent }

func (cikepeercorrtable *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable) GetParentYangName() string { return "CISCO-IPSEC-FLOW-MONITOR-MIB" }

// CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable_Cikepeercorrentry
// Each entry contains the attributes of an
// IPsec Phase-1 IKE Peer Association to IPsec
// Phase-2 Tunnel Correlation.
type CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable_Cikepeercorrentry struct {
    parent types.Entity
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

func (cikepeercorrentry *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable_Cikepeercorrentry) GetFilter() yfilter.YFilter { return cikepeercorrentry.YFilter }

func (cikepeercorrentry *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable_Cikepeercorrentry) SetFilter(yf yfilter.YFilter) { cikepeercorrentry.YFilter = yf }

func (cikepeercorrentry *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable_Cikepeercorrentry) GetGoName(yname string) string {
    if yname == "cikePeerCorrLocalType" { return "Cikepeercorrlocaltype" }
    if yname == "cikePeerCorrLocalValue" { return "Cikepeercorrlocalvalue" }
    if yname == "cikePeerCorrRemoteType" { return "Cikepeercorrremotetype" }
    if yname == "cikePeerCorrRemoteValue" { return "Cikepeercorrremotevalue" }
    if yname == "cikePeerCorrIntIndex" { return "Cikepeercorrintindex" }
    if yname == "cikePeerCorrSeqNum" { return "Cikepeercorrseqnum" }
    if yname == "cikePeerCorrIpSecTunIndex" { return "Cikepeercorripsectunindex" }
    return ""
}

func (cikepeercorrentry *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable_Cikepeercorrentry) GetSegmentPath() string {
    return "cikePeerCorrEntry" + "[cikePeerCorrLocalType='" + fmt.Sprintf("%v", cikepeercorrentry.Cikepeercorrlocaltype) + "']" + "[cikePeerCorrLocalValue='" + fmt.Sprintf("%v", cikepeercorrentry.Cikepeercorrlocalvalue) + "']" + "[cikePeerCorrRemoteType='" + fmt.Sprintf("%v", cikepeercorrentry.Cikepeercorrremotetype) + "']" + "[cikePeerCorrRemoteValue='" + fmt.Sprintf("%v", cikepeercorrentry.Cikepeercorrremotevalue) + "']" + "[cikePeerCorrIntIndex='" + fmt.Sprintf("%v", cikepeercorrentry.Cikepeercorrintindex) + "']" + "[cikePeerCorrSeqNum='" + fmt.Sprintf("%v", cikepeercorrentry.Cikepeercorrseqnum) + "']"
}

func (cikepeercorrentry *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable_Cikepeercorrentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cikepeercorrentry *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable_Cikepeercorrentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cikepeercorrentry *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable_Cikepeercorrentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cikePeerCorrLocalType"] = cikepeercorrentry.Cikepeercorrlocaltype
    leafs["cikePeerCorrLocalValue"] = cikepeercorrentry.Cikepeercorrlocalvalue
    leafs["cikePeerCorrRemoteType"] = cikepeercorrentry.Cikepeercorrremotetype
    leafs["cikePeerCorrRemoteValue"] = cikepeercorrentry.Cikepeercorrremotevalue
    leafs["cikePeerCorrIntIndex"] = cikepeercorrentry.Cikepeercorrintindex
    leafs["cikePeerCorrSeqNum"] = cikepeercorrentry.Cikepeercorrseqnum
    leafs["cikePeerCorrIpSecTunIndex"] = cikepeercorrentry.Cikepeercorripsectunindex
    return leafs
}

func (cikepeercorrentry *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable_Cikepeercorrentry) GetBundleName() string { return "cisco_ios_xe" }

func (cikepeercorrentry *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable_Cikepeercorrentry) GetYangName() string { return "cikePeerCorrEntry" }

func (cikepeercorrentry *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable_Cikepeercorrentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cikepeercorrentry *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable_Cikepeercorrentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cikepeercorrentry *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable_Cikepeercorrentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cikepeercorrentry *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable_Cikepeercorrentry) SetParent(parent types.Entity) { cikepeercorrentry.parent = parent }

func (cikepeercorrentry *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable_Cikepeercorrentry) GetParent() types.Entity { return cikepeercorrentry.parent }

func (cikepeercorrentry *CISCOIPSECFLOWMONITORMIB_Cikepeercorrtable_Cikepeercorrentry) GetParentYangName() string { return "cikePeerCorrTable" }

// CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable
// Phase-1 IKE stats information is included in this table.
// Each entry is related to a specific gateway which is 
// identified by 'cmgwIndex'.
type CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains the attributes of an Phase-1 IKE stats information for
    // the related gateway.  There is only one entry for each gateway. The entry 
    // is created when a gateway up and cannot be deleted. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable_Cikephase1Gwstatsentry.
    Cikephase1Gwstatsentry []CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable_Cikephase1Gwstatsentry
}

func (cikephase1Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable) GetFilter() yfilter.YFilter { return cikephase1Gwstatstable.YFilter }

func (cikephase1Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable) SetFilter(yf yfilter.YFilter) { cikephase1Gwstatstable.YFilter = yf }

func (cikephase1Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable) GetGoName(yname string) string {
    if yname == "cikePhase1GWStatsEntry" { return "Cikephase1Gwstatsentry" }
    return ""
}

func (cikephase1Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable) GetSegmentPath() string {
    return "cikePhase1GWStatsTable"
}

func (cikephase1Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cikePhase1GWStatsEntry" {
        for _, c := range cikephase1Gwstatstable.Cikephase1Gwstatsentry {
            if cikephase1Gwstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable_Cikephase1Gwstatsentry{}
        cikephase1Gwstatstable.Cikephase1Gwstatsentry = append(cikephase1Gwstatstable.Cikephase1Gwstatsentry, child)
        return &cikephase1Gwstatstable.Cikephase1Gwstatsentry[len(cikephase1Gwstatstable.Cikephase1Gwstatsentry)-1]
    }
    return nil
}

func (cikephase1Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cikephase1Gwstatstable.Cikephase1Gwstatsentry {
        children[cikephase1Gwstatstable.Cikephase1Gwstatsentry[i].GetSegmentPath()] = &cikephase1Gwstatstable.Cikephase1Gwstatsentry[i]
    }
    return children
}

func (cikephase1Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cikephase1Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (cikephase1Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable) GetYangName() string { return "cikePhase1GWStatsTable" }

func (cikephase1Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cikephase1Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cikephase1Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cikephase1Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable) SetParent(parent types.Entity) { cikephase1Gwstatstable.parent = parent }

func (cikephase1Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable) GetParent() types.Entity { return cikephase1Gwstatstable.parent }

func (cikephase1Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable) GetParentYangName() string { return "CISCO-IPSEC-FLOW-MONITOR-MIB" }

// CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable_Cikephase1Gwstatsentry
// Each entry contains the attributes of an Phase-1 IKE stats
// information for the related gateway.
// 
// There is only one entry for each gateway. The entry 
// is created when a gateway up and cannot be deleted.
type CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable_Cikephase1Gwstatsentry struct {
    parent types.Entity
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

func (cikephase1Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable_Cikephase1Gwstatsentry) GetFilter() yfilter.YFilter { return cikephase1Gwstatsentry.YFilter }

func (cikephase1Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable_Cikephase1Gwstatsentry) SetFilter(yf yfilter.YFilter) { cikephase1Gwstatsentry.YFilter = yf }

func (cikephase1Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable_Cikephase1Gwstatsentry) GetGoName(yname string) string {
    if yname == "cmgwIndex" { return "Cmgwindex" }
    if yname == "cikePhase1GWActiveTunnels" { return "Cikephase1Gwactivetunnels" }
    if yname == "cikePhase1GWPreviousTunnels" { return "Cikephase1Gwprevioustunnels" }
    if yname == "cikePhase1GWInOctets" { return "Cikephase1Gwinoctets" }
    if yname == "cikePhase1GWInPkts" { return "Cikephase1Gwinpkts" }
    if yname == "cikePhase1GWInDropPkts" { return "Cikephase1Gwindroppkts" }
    if yname == "cikePhase1GWInNotifys" { return "Cikephase1Gwinnotifys" }
    if yname == "cikePhase1GWInP2Exchgs" { return "Cikephase1Gwinp2Exchgs" }
    if yname == "cikePhase1GWInP2ExchgInvalids" { return "Cikephase1Gwinp2Exchginvalids" }
    if yname == "cikePhase1GWInP2ExchgRejects" { return "Cikephase1Gwinp2Exchgrejects" }
    if yname == "cikePhase1GWInP2SaDelRequests" { return "Cikephase1Gwinp2Sadelrequests" }
    if yname == "cikePhase1GWOutOctets" { return "Cikephase1Gwoutoctets" }
    if yname == "cikePhase1GWOutPkts" { return "Cikephase1Gwoutpkts" }
    if yname == "cikePhase1GWOutDropPkts" { return "Cikephase1Gwoutdroppkts" }
    if yname == "cikePhase1GWOutNotifys" { return "Cikephase1Gwoutnotifys" }
    if yname == "cikePhase1GWOutP2Exchgs" { return "Cikephase1Gwoutp2Exchgs" }
    if yname == "cikePhase1GWOutP2ExchgInvalids" { return "Cikephase1Gwoutp2Exchginvalids" }
    if yname == "cikePhase1GWOutP2ExchgRejects" { return "Cikephase1Gwoutp2Exchgrejects" }
    if yname == "cikePhase1GWOutP2SaDelRequests" { return "Cikephase1Gwoutp2Sadelrequests" }
    if yname == "cikePhase1GWInitTunnels" { return "Cikephase1Gwinittunnels" }
    if yname == "cikePhase1GWInitTunnelFails" { return "Cikephase1Gwinittunnelfails" }
    if yname == "cikePhase1GWRespTunnelFails" { return "Cikephase1Gwresptunnelfails" }
    if yname == "cikePhase1GWSysCapFails" { return "Cikephase1Gwsyscapfails" }
    if yname == "cikePhase1GWAuthFails" { return "Cikephase1Gwauthfails" }
    if yname == "cikePhase1GWDecryptFails" { return "Cikephase1Gwdecryptfails" }
    if yname == "cikePhase1GWHashValidFails" { return "Cikephase1Gwhashvalidfails" }
    if yname == "cikePhase1GWNoSaFails" { return "Cikephase1Gwnosafails" }
    return ""
}

func (cikephase1Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable_Cikephase1Gwstatsentry) GetSegmentPath() string {
    return "cikePhase1GWStatsEntry" + "[cmgwIndex='" + fmt.Sprintf("%v", cikephase1Gwstatsentry.Cmgwindex) + "']"
}

func (cikephase1Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable_Cikephase1Gwstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cikephase1Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable_Cikephase1Gwstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cikephase1Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable_Cikephase1Gwstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cmgwIndex"] = cikephase1Gwstatsentry.Cmgwindex
    leafs["cikePhase1GWActiveTunnels"] = cikephase1Gwstatsentry.Cikephase1Gwactivetunnels
    leafs["cikePhase1GWPreviousTunnels"] = cikephase1Gwstatsentry.Cikephase1Gwprevioustunnels
    leafs["cikePhase1GWInOctets"] = cikephase1Gwstatsentry.Cikephase1Gwinoctets
    leafs["cikePhase1GWInPkts"] = cikephase1Gwstatsentry.Cikephase1Gwinpkts
    leafs["cikePhase1GWInDropPkts"] = cikephase1Gwstatsentry.Cikephase1Gwindroppkts
    leafs["cikePhase1GWInNotifys"] = cikephase1Gwstatsentry.Cikephase1Gwinnotifys
    leafs["cikePhase1GWInP2Exchgs"] = cikephase1Gwstatsentry.Cikephase1Gwinp2Exchgs
    leafs["cikePhase1GWInP2ExchgInvalids"] = cikephase1Gwstatsentry.Cikephase1Gwinp2Exchginvalids
    leafs["cikePhase1GWInP2ExchgRejects"] = cikephase1Gwstatsentry.Cikephase1Gwinp2Exchgrejects
    leafs["cikePhase1GWInP2SaDelRequests"] = cikephase1Gwstatsentry.Cikephase1Gwinp2Sadelrequests
    leafs["cikePhase1GWOutOctets"] = cikephase1Gwstatsentry.Cikephase1Gwoutoctets
    leafs["cikePhase1GWOutPkts"] = cikephase1Gwstatsentry.Cikephase1Gwoutpkts
    leafs["cikePhase1GWOutDropPkts"] = cikephase1Gwstatsentry.Cikephase1Gwoutdroppkts
    leafs["cikePhase1GWOutNotifys"] = cikephase1Gwstatsentry.Cikephase1Gwoutnotifys
    leafs["cikePhase1GWOutP2Exchgs"] = cikephase1Gwstatsentry.Cikephase1Gwoutp2Exchgs
    leafs["cikePhase1GWOutP2ExchgInvalids"] = cikephase1Gwstatsentry.Cikephase1Gwoutp2Exchginvalids
    leafs["cikePhase1GWOutP2ExchgRejects"] = cikephase1Gwstatsentry.Cikephase1Gwoutp2Exchgrejects
    leafs["cikePhase1GWOutP2SaDelRequests"] = cikephase1Gwstatsentry.Cikephase1Gwoutp2Sadelrequests
    leafs["cikePhase1GWInitTunnels"] = cikephase1Gwstatsentry.Cikephase1Gwinittunnels
    leafs["cikePhase1GWInitTunnelFails"] = cikephase1Gwstatsentry.Cikephase1Gwinittunnelfails
    leafs["cikePhase1GWRespTunnelFails"] = cikephase1Gwstatsentry.Cikephase1Gwresptunnelfails
    leafs["cikePhase1GWSysCapFails"] = cikephase1Gwstatsentry.Cikephase1Gwsyscapfails
    leafs["cikePhase1GWAuthFails"] = cikephase1Gwstatsentry.Cikephase1Gwauthfails
    leafs["cikePhase1GWDecryptFails"] = cikephase1Gwstatsentry.Cikephase1Gwdecryptfails
    leafs["cikePhase1GWHashValidFails"] = cikephase1Gwstatsentry.Cikephase1Gwhashvalidfails
    leafs["cikePhase1GWNoSaFails"] = cikephase1Gwstatsentry.Cikephase1Gwnosafails
    return leafs
}

func (cikephase1Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable_Cikephase1Gwstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (cikephase1Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable_Cikephase1Gwstatsentry) GetYangName() string { return "cikePhase1GWStatsEntry" }

func (cikephase1Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable_Cikephase1Gwstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cikephase1Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable_Cikephase1Gwstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cikephase1Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable_Cikephase1Gwstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cikephase1Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable_Cikephase1Gwstatsentry) SetParent(parent types.Entity) { cikephase1Gwstatsentry.parent = parent }

func (cikephase1Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable_Cikephase1Gwstatsentry) GetParent() types.Entity { return cikephase1Gwstatsentry.parent }

func (cikephase1Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cikephase1Gwstatstable_Cikephase1Gwstatsentry) GetParentYangName() string { return "cikePhase1GWStatsTable" }

// CISCOIPSECFLOWMONITORMIB_Cipsectunneltable
// The IPsec Phase-2 Tunnel Table.
// There is one entry in this table for 
// each active IPsec Phase-2 Tunnel.
type CISCOIPSECFLOWMONITORMIB_Cipsectunneltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with an active IPsec Phase-2
    // Tunnel. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_Cipsectunneltable_Cipsectunnelentry.
    Cipsectunnelentry []CISCOIPSECFLOWMONITORMIB_Cipsectunneltable_Cipsectunnelentry
}

func (cipsectunneltable *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable) GetFilter() yfilter.YFilter { return cipsectunneltable.YFilter }

func (cipsectunneltable *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable) SetFilter(yf yfilter.YFilter) { cipsectunneltable.YFilter = yf }

func (cipsectunneltable *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable) GetGoName(yname string) string {
    if yname == "cipSecTunnelEntry" { return "Cipsectunnelentry" }
    return ""
}

func (cipsectunneltable *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable) GetSegmentPath() string {
    return "cipSecTunnelTable"
}

func (cipsectunneltable *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipSecTunnelEntry" {
        for _, c := range cipsectunneltable.Cipsectunnelentry {
            if cipsectunneltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSECFLOWMONITORMIB_Cipsectunneltable_Cipsectunnelentry{}
        cipsectunneltable.Cipsectunnelentry = append(cipsectunneltable.Cipsectunnelentry, child)
        return &cipsectunneltable.Cipsectunnelentry[len(cipsectunneltable.Cipsectunnelentry)-1]
    }
    return nil
}

func (cipsectunneltable *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cipsectunneltable.Cipsectunnelentry {
        children[cipsectunneltable.Cipsectunnelentry[i].GetSegmentPath()] = &cipsectunneltable.Cipsectunnelentry[i]
    }
    return children
}

func (cipsectunneltable *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cipsectunneltable *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable) GetBundleName() string { return "cisco_ios_xe" }

func (cipsectunneltable *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable) GetYangName() string { return "cipSecTunnelTable" }

func (cipsectunneltable *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsectunneltable *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsectunneltable *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsectunneltable *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable) SetParent(parent types.Entity) { cipsectunneltable.parent = parent }

func (cipsectunneltable *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable) GetParent() types.Entity { return cipsectunneltable.parent }

func (cipsectunneltable *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable) GetParentYangName() string { return "CISCO-IPSEC-FLOW-MONITOR-MIB" }

// CISCOIPSECFLOWMONITORMIB_Cipsectunneltable_Cipsectunnelentry
// Each entry contains the attributes
// associated with an active IPsec Phase-2 Tunnel.
type CISCOIPSECFLOWMONITORMIB_Cipsectunneltable_Cipsectunnelentry struct {
    parent types.Entity
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

func (cipsectunnelentry *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable_Cipsectunnelentry) GetFilter() yfilter.YFilter { return cipsectunnelentry.YFilter }

func (cipsectunnelentry *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable_Cipsectunnelentry) SetFilter(yf yfilter.YFilter) { cipsectunnelentry.YFilter = yf }

func (cipsectunnelentry *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable_Cipsectunnelentry) GetGoName(yname string) string {
    if yname == "cipSecTunIndex" { return "Cipsectunindex" }
    if yname == "cipSecTunIkeTunnelIndex" { return "Cipsectuniketunnelindex" }
    if yname == "cipSecTunIkeTunnelAlive" { return "Cipsectuniketunnelalive" }
    if yname == "cipSecTunLocalAddr" { return "Cipsectunlocaladdr" }
    if yname == "cipSecTunRemoteAddr" { return "Cipsectunremoteaddr" }
    if yname == "cipSecTunKeyType" { return "Cipsectunkeytype" }
    if yname == "cipSecTunEncapMode" { return "Cipsectunencapmode" }
    if yname == "cipSecTunLifeSize" { return "Cipsectunlifesize" }
    if yname == "cipSecTunLifeTime" { return "Cipsectunlifetime" }
    if yname == "cipSecTunActiveTime" { return "Cipsectunactivetime" }
    if yname == "cipSecTunSaLifeSizeThreshold" { return "Cipsectunsalifesizethreshold" }
    if yname == "cipSecTunSaLifeTimeThreshold" { return "Cipsectunsalifetimethreshold" }
    if yname == "cipSecTunTotalRefreshes" { return "Cipsectuntotalrefreshes" }
    if yname == "cipSecTunExpiredSaInstances" { return "Cipsectunexpiredsainstances" }
    if yname == "cipSecTunCurrentSaInstances" { return "Cipsectuncurrentsainstances" }
    if yname == "cipSecTunInSaDiffHellmanGrp" { return "Cipsectuninsadiffhellmangrp" }
    if yname == "cipSecTunInSaEncryptAlgo" { return "Cipsectuninsaencryptalgo" }
    if yname == "cipSecTunInSaAhAuthAlgo" { return "Cipsectuninsaahauthalgo" }
    if yname == "cipSecTunInSaEspAuthAlgo" { return "Cipsectuninsaespauthalgo" }
    if yname == "cipSecTunInSaDecompAlgo" { return "Cipsectuninsadecompalgo" }
    if yname == "cipSecTunOutSaDiffHellmanGrp" { return "Cipsectunoutsadiffhellmangrp" }
    if yname == "cipSecTunOutSaEncryptAlgo" { return "Cipsectunoutsaencryptalgo" }
    if yname == "cipSecTunOutSaAhAuthAlgo" { return "Cipsectunoutsaahauthalgo" }
    if yname == "cipSecTunOutSaEspAuthAlgo" { return "Cipsectunoutsaespauthalgo" }
    if yname == "cipSecTunOutSaCompAlgo" { return "Cipsectunoutsacompalgo" }
    if yname == "cipSecTunInOctets" { return "Cipsectuninoctets" }
    if yname == "cipSecTunHcInOctets" { return "Cipsectunhcinoctets" }
    if yname == "cipSecTunInOctWraps" { return "Cipsectuninoctwraps" }
    if yname == "cipSecTunInDecompOctets" { return "Cipsectunindecompoctets" }
    if yname == "cipSecTunHcInDecompOctets" { return "Cipsectunhcindecompoctets" }
    if yname == "cipSecTunInDecompOctWraps" { return "Cipsectunindecompoctwraps" }
    if yname == "cipSecTunInPkts" { return "Cipsectuninpkts" }
    if yname == "cipSecTunInDropPkts" { return "Cipsectunindroppkts" }
    if yname == "cipSecTunInReplayDropPkts" { return "Cipsectuninreplaydroppkts" }
    if yname == "cipSecTunInAuths" { return "Cipsectuninauths" }
    if yname == "cipSecTunInAuthFails" { return "Cipsectuninauthfails" }
    if yname == "cipSecTunInDecrypts" { return "Cipsectunindecrypts" }
    if yname == "cipSecTunInDecryptFails" { return "Cipsectunindecryptfails" }
    if yname == "cipSecTunOutOctets" { return "Cipsectunoutoctets" }
    if yname == "cipSecTunHcOutOctets" { return "Cipsectunhcoutoctets" }
    if yname == "cipSecTunOutOctWraps" { return "Cipsectunoutoctwraps" }
    if yname == "cipSecTunOutUncompOctets" { return "Cipsectunoutuncompoctets" }
    if yname == "cipSecTunHcOutUncompOctets" { return "Cipsectunhcoutuncompoctets" }
    if yname == "cipSecTunOutUncompOctWraps" { return "Cipsectunoutuncompoctwraps" }
    if yname == "cipSecTunOutPkts" { return "Cipsectunoutpkts" }
    if yname == "cipSecTunOutDropPkts" { return "Cipsectunoutdroppkts" }
    if yname == "cipSecTunOutAuths" { return "Cipsectunoutauths" }
    if yname == "cipSecTunOutAuthFails" { return "Cipsectunoutauthfails" }
    if yname == "cipSecTunOutEncrypts" { return "Cipsectunoutencrypts" }
    if yname == "cipSecTunOutEncryptFails" { return "Cipsectunoutencryptfails" }
    if yname == "cipSecTunStatus" { return "Cipsectunstatus" }
    return ""
}

func (cipsectunnelentry *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable_Cipsectunnelentry) GetSegmentPath() string {
    return "cipSecTunnelEntry" + "[cipSecTunIndex='" + fmt.Sprintf("%v", cipsectunnelentry.Cipsectunindex) + "']"
}

func (cipsectunnelentry *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable_Cipsectunnelentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipsectunnelentry *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable_Cipsectunnelentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipsectunnelentry *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable_Cipsectunnelentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipSecTunIndex"] = cipsectunnelentry.Cipsectunindex
    leafs["cipSecTunIkeTunnelIndex"] = cipsectunnelentry.Cipsectuniketunnelindex
    leafs["cipSecTunIkeTunnelAlive"] = cipsectunnelentry.Cipsectuniketunnelalive
    leafs["cipSecTunLocalAddr"] = cipsectunnelentry.Cipsectunlocaladdr
    leafs["cipSecTunRemoteAddr"] = cipsectunnelentry.Cipsectunremoteaddr
    leafs["cipSecTunKeyType"] = cipsectunnelentry.Cipsectunkeytype
    leafs["cipSecTunEncapMode"] = cipsectunnelentry.Cipsectunencapmode
    leafs["cipSecTunLifeSize"] = cipsectunnelentry.Cipsectunlifesize
    leafs["cipSecTunLifeTime"] = cipsectunnelentry.Cipsectunlifetime
    leafs["cipSecTunActiveTime"] = cipsectunnelentry.Cipsectunactivetime
    leafs["cipSecTunSaLifeSizeThreshold"] = cipsectunnelentry.Cipsectunsalifesizethreshold
    leafs["cipSecTunSaLifeTimeThreshold"] = cipsectunnelentry.Cipsectunsalifetimethreshold
    leafs["cipSecTunTotalRefreshes"] = cipsectunnelentry.Cipsectuntotalrefreshes
    leafs["cipSecTunExpiredSaInstances"] = cipsectunnelentry.Cipsectunexpiredsainstances
    leafs["cipSecTunCurrentSaInstances"] = cipsectunnelentry.Cipsectuncurrentsainstances
    leafs["cipSecTunInSaDiffHellmanGrp"] = cipsectunnelentry.Cipsectuninsadiffhellmangrp
    leafs["cipSecTunInSaEncryptAlgo"] = cipsectunnelentry.Cipsectuninsaencryptalgo
    leafs["cipSecTunInSaAhAuthAlgo"] = cipsectunnelentry.Cipsectuninsaahauthalgo
    leafs["cipSecTunInSaEspAuthAlgo"] = cipsectunnelentry.Cipsectuninsaespauthalgo
    leafs["cipSecTunInSaDecompAlgo"] = cipsectunnelentry.Cipsectuninsadecompalgo
    leafs["cipSecTunOutSaDiffHellmanGrp"] = cipsectunnelentry.Cipsectunoutsadiffhellmangrp
    leafs["cipSecTunOutSaEncryptAlgo"] = cipsectunnelentry.Cipsectunoutsaencryptalgo
    leafs["cipSecTunOutSaAhAuthAlgo"] = cipsectunnelentry.Cipsectunoutsaahauthalgo
    leafs["cipSecTunOutSaEspAuthAlgo"] = cipsectunnelentry.Cipsectunoutsaespauthalgo
    leafs["cipSecTunOutSaCompAlgo"] = cipsectunnelentry.Cipsectunoutsacompalgo
    leafs["cipSecTunInOctets"] = cipsectunnelentry.Cipsectuninoctets
    leafs["cipSecTunHcInOctets"] = cipsectunnelentry.Cipsectunhcinoctets
    leafs["cipSecTunInOctWraps"] = cipsectunnelentry.Cipsectuninoctwraps
    leafs["cipSecTunInDecompOctets"] = cipsectunnelentry.Cipsectunindecompoctets
    leafs["cipSecTunHcInDecompOctets"] = cipsectunnelentry.Cipsectunhcindecompoctets
    leafs["cipSecTunInDecompOctWraps"] = cipsectunnelentry.Cipsectunindecompoctwraps
    leafs["cipSecTunInPkts"] = cipsectunnelentry.Cipsectuninpkts
    leafs["cipSecTunInDropPkts"] = cipsectunnelentry.Cipsectunindroppkts
    leafs["cipSecTunInReplayDropPkts"] = cipsectunnelentry.Cipsectuninreplaydroppkts
    leafs["cipSecTunInAuths"] = cipsectunnelentry.Cipsectuninauths
    leafs["cipSecTunInAuthFails"] = cipsectunnelentry.Cipsectuninauthfails
    leafs["cipSecTunInDecrypts"] = cipsectunnelentry.Cipsectunindecrypts
    leafs["cipSecTunInDecryptFails"] = cipsectunnelentry.Cipsectunindecryptfails
    leafs["cipSecTunOutOctets"] = cipsectunnelentry.Cipsectunoutoctets
    leafs["cipSecTunHcOutOctets"] = cipsectunnelentry.Cipsectunhcoutoctets
    leafs["cipSecTunOutOctWraps"] = cipsectunnelentry.Cipsectunoutoctwraps
    leafs["cipSecTunOutUncompOctets"] = cipsectunnelentry.Cipsectunoutuncompoctets
    leafs["cipSecTunHcOutUncompOctets"] = cipsectunnelentry.Cipsectunhcoutuncompoctets
    leafs["cipSecTunOutUncompOctWraps"] = cipsectunnelentry.Cipsectunoutuncompoctwraps
    leafs["cipSecTunOutPkts"] = cipsectunnelentry.Cipsectunoutpkts
    leafs["cipSecTunOutDropPkts"] = cipsectunnelentry.Cipsectunoutdroppkts
    leafs["cipSecTunOutAuths"] = cipsectunnelentry.Cipsectunoutauths
    leafs["cipSecTunOutAuthFails"] = cipsectunnelentry.Cipsectunoutauthfails
    leafs["cipSecTunOutEncrypts"] = cipsectunnelentry.Cipsectunoutencrypts
    leafs["cipSecTunOutEncryptFails"] = cipsectunnelentry.Cipsectunoutencryptfails
    leafs["cipSecTunStatus"] = cipsectunnelentry.Cipsectunstatus
    return leafs
}

func (cipsectunnelentry *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable_Cipsectunnelentry) GetBundleName() string { return "cisco_ios_xe" }

func (cipsectunnelentry *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable_Cipsectunnelentry) GetYangName() string { return "cipSecTunnelEntry" }

func (cipsectunnelentry *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable_Cipsectunnelentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsectunnelentry *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable_Cipsectunnelentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsectunnelentry *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable_Cipsectunnelentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsectunnelentry *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable_Cipsectunnelentry) SetParent(parent types.Entity) { cipsectunnelentry.parent = parent }

func (cipsectunnelentry *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable_Cipsectunnelentry) GetParent() types.Entity { return cipsectunnelentry.parent }

func (cipsectunnelentry *CISCOIPSECFLOWMONITORMIB_Cipsectunneltable_Cipsectunnelentry) GetParentYangName() string { return "cipSecTunnelTable" }

// CISCOIPSECFLOWMONITORMIB_Cipsecendpttable
// The IPsec Phase-2 Tunnel Endpoint Table.
// This table contains an entry for each 
// active endpoint associated with an IPsec
//  Phase-2 Tunnel.
type CISCOIPSECFLOWMONITORMIB_Cipsecendpttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An IPsec Phase-2 Tunnel Endpoint entry. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_Cipsecendpttable_Cipsecendptentry.
    Cipsecendptentry []CISCOIPSECFLOWMONITORMIB_Cipsecendpttable_Cipsecendptentry
}

func (cipsecendpttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable) GetFilter() yfilter.YFilter { return cipsecendpttable.YFilter }

func (cipsecendpttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable) SetFilter(yf yfilter.YFilter) { cipsecendpttable.YFilter = yf }

func (cipsecendpttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable) GetGoName(yname string) string {
    if yname == "cipSecEndPtEntry" { return "Cipsecendptentry" }
    return ""
}

func (cipsecendpttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable) GetSegmentPath() string {
    return "cipSecEndPtTable"
}

func (cipsecendpttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipSecEndPtEntry" {
        for _, c := range cipsecendpttable.Cipsecendptentry {
            if cipsecendpttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSECFLOWMONITORMIB_Cipsecendpttable_Cipsecendptentry{}
        cipsecendpttable.Cipsecendptentry = append(cipsecendpttable.Cipsecendptentry, child)
        return &cipsecendpttable.Cipsecendptentry[len(cipsecendpttable.Cipsecendptentry)-1]
    }
    return nil
}

func (cipsecendpttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cipsecendpttable.Cipsecendptentry {
        children[cipsecendpttable.Cipsecendptentry[i].GetSegmentPath()] = &cipsecendpttable.Cipsecendptentry[i]
    }
    return children
}

func (cipsecendpttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cipsecendpttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable) GetBundleName() string { return "cisco_ios_xe" }

func (cipsecendpttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable) GetYangName() string { return "cipSecEndPtTable" }

func (cipsecendpttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsecendpttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsecendpttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsecendpttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable) SetParent(parent types.Entity) { cipsecendpttable.parent = parent }

func (cipsecendpttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable) GetParent() types.Entity { return cipsecendpttable.parent }

func (cipsecendpttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable) GetParentYangName() string { return "CISCO-IPSEC-FLOW-MONITOR-MIB" }

// CISCOIPSECFLOWMONITORMIB_Cipsecendpttable_Cipsecendptentry
// An IPsec Phase-2 Tunnel Endpoint entry.
type CISCOIPSECFLOWMONITORMIB_Cipsecendpttable_Cipsecendptentry struct {
    parent types.Entity
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

func (cipsecendptentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable_Cipsecendptentry) GetFilter() yfilter.YFilter { return cipsecendptentry.YFilter }

func (cipsecendptentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable_Cipsecendptentry) SetFilter(yf yfilter.YFilter) { cipsecendptentry.YFilter = yf }

func (cipsecendptentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable_Cipsecendptentry) GetGoName(yname string) string {
    if yname == "cipSecTunIndex" { return "Cipsectunindex" }
    if yname == "cipSecEndPtIndex" { return "Cipsecendptindex" }
    if yname == "cipSecEndPtLocalName" { return "Cipsecendptlocalname" }
    if yname == "cipSecEndPtLocalType" { return "Cipsecendptlocaltype" }
    if yname == "cipSecEndPtLocalAddr1" { return "Cipsecendptlocaladdr1" }
    if yname == "cipSecEndPtLocalAddr2" { return "Cipsecendptlocaladdr2" }
    if yname == "cipSecEndPtLocalProtocol" { return "Cipsecendptlocalprotocol" }
    if yname == "cipSecEndPtLocalPort" { return "Cipsecendptlocalport" }
    if yname == "cipSecEndPtRemoteName" { return "Cipsecendptremotename" }
    if yname == "cipSecEndPtRemoteType" { return "Cipsecendptremotetype" }
    if yname == "cipSecEndPtRemoteAddr1" { return "Cipsecendptremoteaddr1" }
    if yname == "cipSecEndPtRemoteAddr2" { return "Cipsecendptremoteaddr2" }
    if yname == "cipSecEndPtRemoteProtocol" { return "Cipsecendptremoteprotocol" }
    if yname == "cipSecEndPtRemotePort" { return "Cipsecendptremoteport" }
    return ""
}

func (cipsecendptentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable_Cipsecendptentry) GetSegmentPath() string {
    return "cipSecEndPtEntry" + "[cipSecTunIndex='" + fmt.Sprintf("%v", cipsecendptentry.Cipsectunindex) + "']" + "[cipSecEndPtIndex='" + fmt.Sprintf("%v", cipsecendptentry.Cipsecendptindex) + "']"
}

func (cipsecendptentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable_Cipsecendptentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipsecendptentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable_Cipsecendptentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipsecendptentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable_Cipsecendptentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipSecTunIndex"] = cipsecendptentry.Cipsectunindex
    leafs["cipSecEndPtIndex"] = cipsecendptentry.Cipsecendptindex
    leafs["cipSecEndPtLocalName"] = cipsecendptentry.Cipsecendptlocalname
    leafs["cipSecEndPtLocalType"] = cipsecendptentry.Cipsecendptlocaltype
    leafs["cipSecEndPtLocalAddr1"] = cipsecendptentry.Cipsecendptlocaladdr1
    leafs["cipSecEndPtLocalAddr2"] = cipsecendptentry.Cipsecendptlocaladdr2
    leafs["cipSecEndPtLocalProtocol"] = cipsecendptentry.Cipsecendptlocalprotocol
    leafs["cipSecEndPtLocalPort"] = cipsecendptentry.Cipsecendptlocalport
    leafs["cipSecEndPtRemoteName"] = cipsecendptentry.Cipsecendptremotename
    leafs["cipSecEndPtRemoteType"] = cipsecendptentry.Cipsecendptremotetype
    leafs["cipSecEndPtRemoteAddr1"] = cipsecendptentry.Cipsecendptremoteaddr1
    leafs["cipSecEndPtRemoteAddr2"] = cipsecendptentry.Cipsecendptremoteaddr2
    leafs["cipSecEndPtRemoteProtocol"] = cipsecendptentry.Cipsecendptremoteprotocol
    leafs["cipSecEndPtRemotePort"] = cipsecendptentry.Cipsecendptremoteport
    return leafs
}

func (cipsecendptentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable_Cipsecendptentry) GetBundleName() string { return "cisco_ios_xe" }

func (cipsecendptentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable_Cipsecendptentry) GetYangName() string { return "cipSecEndPtEntry" }

func (cipsecendptentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable_Cipsecendptentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsecendptentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable_Cipsecendptentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsecendptentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable_Cipsecendptentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsecendptentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable_Cipsecendptentry) SetParent(parent types.Entity) { cipsecendptentry.parent = parent }

func (cipsecendptentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable_Cipsecendptentry) GetParent() types.Entity { return cipsecendptentry.parent }

func (cipsecendptentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpttable_Cipsecendptentry) GetParentYangName() string { return "cipSecEndPtTable" }

// CISCOIPSECFLOWMONITORMIB_Cipsecspitable
// The IPsec Phase-2 Security Protection Index Table.
// This table contains an entry for each active 
// and expiring security
//  association.
type CISCOIPSECFLOWMONITORMIB_Cipsecspitable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with active and expiring
    // IPsec Phase-2  security associations. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry.
    Cipsecspientry []CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry
}

func (cipsecspitable *CISCOIPSECFLOWMONITORMIB_Cipsecspitable) GetFilter() yfilter.YFilter { return cipsecspitable.YFilter }

func (cipsecspitable *CISCOIPSECFLOWMONITORMIB_Cipsecspitable) SetFilter(yf yfilter.YFilter) { cipsecspitable.YFilter = yf }

func (cipsecspitable *CISCOIPSECFLOWMONITORMIB_Cipsecspitable) GetGoName(yname string) string {
    if yname == "cipSecSpiEntry" { return "Cipsecspientry" }
    return ""
}

func (cipsecspitable *CISCOIPSECFLOWMONITORMIB_Cipsecspitable) GetSegmentPath() string {
    return "cipSecSpiTable"
}

func (cipsecspitable *CISCOIPSECFLOWMONITORMIB_Cipsecspitable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipSecSpiEntry" {
        for _, c := range cipsecspitable.Cipsecspientry {
            if cipsecspitable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry{}
        cipsecspitable.Cipsecspientry = append(cipsecspitable.Cipsecspientry, child)
        return &cipsecspitable.Cipsecspientry[len(cipsecspitable.Cipsecspientry)-1]
    }
    return nil
}

func (cipsecspitable *CISCOIPSECFLOWMONITORMIB_Cipsecspitable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cipsecspitable.Cipsecspientry {
        children[cipsecspitable.Cipsecspientry[i].GetSegmentPath()] = &cipsecspitable.Cipsecspientry[i]
    }
    return children
}

func (cipsecspitable *CISCOIPSECFLOWMONITORMIB_Cipsecspitable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cipsecspitable *CISCOIPSECFLOWMONITORMIB_Cipsecspitable) GetBundleName() string { return "cisco_ios_xe" }

func (cipsecspitable *CISCOIPSECFLOWMONITORMIB_Cipsecspitable) GetYangName() string { return "cipSecSpiTable" }

func (cipsecspitable *CISCOIPSECFLOWMONITORMIB_Cipsecspitable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsecspitable *CISCOIPSECFLOWMONITORMIB_Cipsecspitable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsecspitable *CISCOIPSECFLOWMONITORMIB_Cipsecspitable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsecspitable *CISCOIPSECFLOWMONITORMIB_Cipsecspitable) SetParent(parent types.Entity) { cipsecspitable.parent = parent }

func (cipsecspitable *CISCOIPSECFLOWMONITORMIB_Cipsecspitable) GetParent() types.Entity { return cipsecspitable.parent }

func (cipsecspitable *CISCOIPSECFLOWMONITORMIB_Cipsecspitable) GetParentYangName() string { return "CISCO-IPSEC-FLOW-MONITOR-MIB" }

// CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry
// Each entry contains the attributes associated with
// active and expiring IPsec Phase-2 
// security associations.
type CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry struct {
    parent types.Entity
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

func (cipsecspientry *CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry) GetFilter() yfilter.YFilter { return cipsecspientry.YFilter }

func (cipsecspientry *CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry) SetFilter(yf yfilter.YFilter) { cipsecspientry.YFilter = yf }

func (cipsecspientry *CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry) GetGoName(yname string) string {
    if yname == "cipSecTunIndex" { return "Cipsectunindex" }
    if yname == "cipSecSpiIndex" { return "Cipsecspiindex" }
    if yname == "cipSecSpiDirection" { return "Cipsecspidirection" }
    if yname == "cipSecSpiValue" { return "Cipsecspivalue" }
    if yname == "cipSecSpiProtocol" { return "Cipsecspiprotocol" }
    if yname == "cipSecSpiStatus" { return "Cipsecspistatus" }
    return ""
}

func (cipsecspientry *CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry) GetSegmentPath() string {
    return "cipSecSpiEntry" + "[cipSecTunIndex='" + fmt.Sprintf("%v", cipsecspientry.Cipsectunindex) + "']" + "[cipSecSpiIndex='" + fmt.Sprintf("%v", cipsecspientry.Cipsecspiindex) + "']"
}

func (cipsecspientry *CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipsecspientry *CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipsecspientry *CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipSecTunIndex"] = cipsecspientry.Cipsectunindex
    leafs["cipSecSpiIndex"] = cipsecspientry.Cipsecspiindex
    leafs["cipSecSpiDirection"] = cipsecspientry.Cipsecspidirection
    leafs["cipSecSpiValue"] = cipsecspientry.Cipsecspivalue
    leafs["cipSecSpiProtocol"] = cipsecspientry.Cipsecspiprotocol
    leafs["cipSecSpiStatus"] = cipsecspientry.Cipsecspistatus
    return leafs
}

func (cipsecspientry *CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry) GetBundleName() string { return "cisco_ios_xe" }

func (cipsecspientry *CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry) GetYangName() string { return "cipSecSpiEntry" }

func (cipsecspientry *CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsecspientry *CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsecspientry *CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsecspientry *CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry) SetParent(parent types.Entity) { cipsecspientry.parent = parent }

func (cipsecspientry *CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry) GetParent() types.Entity { return cipsecspientry.parent }

func (cipsecspientry *CISCOIPSECFLOWMONITORMIB_Cipsecspitable_Cipsecspientry) GetParentYangName() string { return "cipSecSpiTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains the attributes of an Phase-2 IPsec stats information
    // for the related gateway.  There is only one entry for each gateway. The
    // entry  is created when a gateway up and cannot be deleted. The type is
    // slice of
    // CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable_Cipsecphase2Gwstatsentry.
    Cipsecphase2Gwstatsentry []CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable_Cipsecphase2Gwstatsentry
}

func (cipsecphase2Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable) GetFilter() yfilter.YFilter { return cipsecphase2Gwstatstable.YFilter }

func (cipsecphase2Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable) SetFilter(yf yfilter.YFilter) { cipsecphase2Gwstatstable.YFilter = yf }

func (cipsecphase2Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable) GetGoName(yname string) string {
    if yname == "cipSecPhase2GWStatsEntry" { return "Cipsecphase2Gwstatsentry" }
    return ""
}

func (cipsecphase2Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable) GetSegmentPath() string {
    return "cipSecPhase2GWStatsTable"
}

func (cipsecphase2Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipSecPhase2GWStatsEntry" {
        for _, c := range cipsecphase2Gwstatstable.Cipsecphase2Gwstatsentry {
            if cipsecphase2Gwstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable_Cipsecphase2Gwstatsentry{}
        cipsecphase2Gwstatstable.Cipsecphase2Gwstatsentry = append(cipsecphase2Gwstatstable.Cipsecphase2Gwstatsentry, child)
        return &cipsecphase2Gwstatstable.Cipsecphase2Gwstatsentry[len(cipsecphase2Gwstatstable.Cipsecphase2Gwstatsentry)-1]
    }
    return nil
}

func (cipsecphase2Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cipsecphase2Gwstatstable.Cipsecphase2Gwstatsentry {
        children[cipsecphase2Gwstatstable.Cipsecphase2Gwstatsentry[i].GetSegmentPath()] = &cipsecphase2Gwstatstable.Cipsecphase2Gwstatsentry[i]
    }
    return children
}

func (cipsecphase2Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cipsecphase2Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (cipsecphase2Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable) GetYangName() string { return "cipSecPhase2GWStatsTable" }

func (cipsecphase2Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsecphase2Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsecphase2Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsecphase2Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable) SetParent(parent types.Entity) { cipsecphase2Gwstatstable.parent = parent }

func (cipsecphase2Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable) GetParent() types.Entity { return cipsecphase2Gwstatstable.parent }

func (cipsecphase2Gwstatstable *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable) GetParentYangName() string { return "CISCO-IPSEC-FLOW-MONITOR-MIB" }

// CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable_Cipsecphase2Gwstatsentry
// Each entry contains the attributes of an Phase-2 IPsec stats
// information for the related gateway.
// 
// There is only one entry for each gateway. The entry 
// is created when a gateway up and cannot be deleted.
type CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable_Cipsecphase2Gwstatsentry struct {
    parent types.Entity
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

func (cipsecphase2Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable_Cipsecphase2Gwstatsentry) GetFilter() yfilter.YFilter { return cipsecphase2Gwstatsentry.YFilter }

func (cipsecphase2Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable_Cipsecphase2Gwstatsentry) SetFilter(yf yfilter.YFilter) { cipsecphase2Gwstatsentry.YFilter = yf }

func (cipsecphase2Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable_Cipsecphase2Gwstatsentry) GetGoName(yname string) string {
    if yname == "cmgwIndex" { return "Cmgwindex" }
    if yname == "cipSecPhase2GWActiveTunnels" { return "Cipsecphase2Gwactivetunnels" }
    if yname == "cipSecPhase2GWPreviousTunnels" { return "Cipsecphase2Gwprevioustunnels" }
    if yname == "cipSecPhase2GWInOctets" { return "Cipsecphase2Gwinoctets" }
    if yname == "cipSecPhase2GWInOctWraps" { return "Cipsecphase2Gwinoctwraps" }
    if yname == "cipSecPhase2GWInDecompOctets" { return "Cipsecphase2Gwindecompoctets" }
    if yname == "cipSecPhase2GWInDecompOctWraps" { return "Cipsecphase2Gwindecompoctwraps" }
    if yname == "cipSecPhase2GWInPkts" { return "Cipsecphase2Gwinpkts" }
    if yname == "cipSecPhase2GWInDrops" { return "Cipsecphase2Gwindrops" }
    if yname == "cipSecPhase2GWInReplayDrops" { return "Cipsecphase2Gwinreplaydrops" }
    if yname == "cipSecPhase2GWInAuths" { return "Cipsecphase2Gwinauths" }
    if yname == "cipSecPhase2GWInAuthFails" { return "Cipsecphase2Gwinauthfails" }
    if yname == "cipSecPhase2GWInDecrypts" { return "Cipsecphase2Gwindecrypts" }
    if yname == "cipSecPhase2GWInDecryptFails" { return "Cipsecphase2Gwindecryptfails" }
    if yname == "cipSecPhase2GWOutOctets" { return "Cipsecphase2Gwoutoctets" }
    if yname == "cipSecPhase2GWOutOctWraps" { return "Cipsecphase2Gwoutoctwraps" }
    if yname == "cipSecPhase2GWOutUncompOctets" { return "Cipsecphase2Gwoutuncompoctets" }
    if yname == "cipSecPhase2GWOutUncompOctWraps" { return "Cipsecphase2Gwoutuncompoctwraps" }
    if yname == "cipSecPhase2GWOutPkts" { return "Cipsecphase2Gwoutpkts" }
    if yname == "cipSecPhase2GWOutDrops" { return "Cipsecphase2Gwoutdrops" }
    if yname == "cipSecPhase2GWOutAuths" { return "Cipsecphase2Gwoutauths" }
    if yname == "cipSecPhase2GWOutAuthFails" { return "Cipsecphase2Gwoutauthfails" }
    if yname == "cipSecPhase2GWOutEncrypts" { return "Cipsecphase2Gwoutencrypts" }
    if yname == "cipSecPhase2GWOutEncryptFails" { return "Cipsecphase2Gwoutencryptfails" }
    if yname == "cipSecPhase2GWProtocolUseFails" { return "Cipsecphase2Gwprotocolusefails" }
    if yname == "cipSecPhase2GWNoSaFails" { return "Cipsecphase2Gwnosafails" }
    if yname == "cipSecPhase2GWSysCapFails" { return "Cipsecphase2Gwsyscapfails" }
    return ""
}

func (cipsecphase2Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable_Cipsecphase2Gwstatsentry) GetSegmentPath() string {
    return "cipSecPhase2GWStatsEntry" + "[cmgwIndex='" + fmt.Sprintf("%v", cipsecphase2Gwstatsentry.Cmgwindex) + "']"
}

func (cipsecphase2Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable_Cipsecphase2Gwstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipsecphase2Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable_Cipsecphase2Gwstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipsecphase2Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable_Cipsecphase2Gwstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cmgwIndex"] = cipsecphase2Gwstatsentry.Cmgwindex
    leafs["cipSecPhase2GWActiveTunnels"] = cipsecphase2Gwstatsentry.Cipsecphase2Gwactivetunnels
    leafs["cipSecPhase2GWPreviousTunnels"] = cipsecphase2Gwstatsentry.Cipsecphase2Gwprevioustunnels
    leafs["cipSecPhase2GWInOctets"] = cipsecphase2Gwstatsentry.Cipsecphase2Gwinoctets
    leafs["cipSecPhase2GWInOctWraps"] = cipsecphase2Gwstatsentry.Cipsecphase2Gwinoctwraps
    leafs["cipSecPhase2GWInDecompOctets"] = cipsecphase2Gwstatsentry.Cipsecphase2Gwindecompoctets
    leafs["cipSecPhase2GWInDecompOctWraps"] = cipsecphase2Gwstatsentry.Cipsecphase2Gwindecompoctwraps
    leafs["cipSecPhase2GWInPkts"] = cipsecphase2Gwstatsentry.Cipsecphase2Gwinpkts
    leafs["cipSecPhase2GWInDrops"] = cipsecphase2Gwstatsentry.Cipsecphase2Gwindrops
    leafs["cipSecPhase2GWInReplayDrops"] = cipsecphase2Gwstatsentry.Cipsecphase2Gwinreplaydrops
    leafs["cipSecPhase2GWInAuths"] = cipsecphase2Gwstatsentry.Cipsecphase2Gwinauths
    leafs["cipSecPhase2GWInAuthFails"] = cipsecphase2Gwstatsentry.Cipsecphase2Gwinauthfails
    leafs["cipSecPhase2GWInDecrypts"] = cipsecphase2Gwstatsentry.Cipsecphase2Gwindecrypts
    leafs["cipSecPhase2GWInDecryptFails"] = cipsecphase2Gwstatsentry.Cipsecphase2Gwindecryptfails
    leafs["cipSecPhase2GWOutOctets"] = cipsecphase2Gwstatsentry.Cipsecphase2Gwoutoctets
    leafs["cipSecPhase2GWOutOctWraps"] = cipsecphase2Gwstatsentry.Cipsecphase2Gwoutoctwraps
    leafs["cipSecPhase2GWOutUncompOctets"] = cipsecphase2Gwstatsentry.Cipsecphase2Gwoutuncompoctets
    leafs["cipSecPhase2GWOutUncompOctWraps"] = cipsecphase2Gwstatsentry.Cipsecphase2Gwoutuncompoctwraps
    leafs["cipSecPhase2GWOutPkts"] = cipsecphase2Gwstatsentry.Cipsecphase2Gwoutpkts
    leafs["cipSecPhase2GWOutDrops"] = cipsecphase2Gwstatsentry.Cipsecphase2Gwoutdrops
    leafs["cipSecPhase2GWOutAuths"] = cipsecphase2Gwstatsentry.Cipsecphase2Gwoutauths
    leafs["cipSecPhase2GWOutAuthFails"] = cipsecphase2Gwstatsentry.Cipsecphase2Gwoutauthfails
    leafs["cipSecPhase2GWOutEncrypts"] = cipsecphase2Gwstatsentry.Cipsecphase2Gwoutencrypts
    leafs["cipSecPhase2GWOutEncryptFails"] = cipsecphase2Gwstatsentry.Cipsecphase2Gwoutencryptfails
    leafs["cipSecPhase2GWProtocolUseFails"] = cipsecphase2Gwstatsentry.Cipsecphase2Gwprotocolusefails
    leafs["cipSecPhase2GWNoSaFails"] = cipsecphase2Gwstatsentry.Cipsecphase2Gwnosafails
    leafs["cipSecPhase2GWSysCapFails"] = cipsecphase2Gwstatsentry.Cipsecphase2Gwsyscapfails
    return leafs
}

func (cipsecphase2Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable_Cipsecphase2Gwstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (cipsecphase2Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable_Cipsecphase2Gwstatsentry) GetYangName() string { return "cipSecPhase2GWStatsEntry" }

func (cipsecphase2Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable_Cipsecphase2Gwstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsecphase2Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable_Cipsecphase2Gwstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsecphase2Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable_Cipsecphase2Gwstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsecphase2Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable_Cipsecphase2Gwstatsentry) SetParent(parent types.Entity) { cipsecphase2Gwstatsentry.parent = parent }

func (cipsecphase2Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable_Cipsecphase2Gwstatsentry) GetParent() types.Entity { return cipsecphase2Gwstatsentry.parent }

func (cipsecphase2Gwstatsentry *CISCOIPSECFLOWMONITORMIB_Cipsecphase2Gwstatstable_Cipsecphase2Gwstatsentry) GetParentYangName() string { return "cipSecPhase2GWStatsTable" }

// CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable
// The IPsec Phase-1 Internet Key Exchange Tunnel
// History Table.  This table is implemented as a 
// sliding window in which only the last n entries 
// are maintained.  The maximum number of entries
//  is specified by the cipSecHistTableSize object.
type CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with a previously active
    // IPsec  Phase-1 IKE Tunnel. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry.
    Ciketunnelhistentry []CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry
}

func (ciketunnelhisttable *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable) GetFilter() yfilter.YFilter { return ciketunnelhisttable.YFilter }

func (ciketunnelhisttable *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable) SetFilter(yf yfilter.YFilter) { ciketunnelhisttable.YFilter = yf }

func (ciketunnelhisttable *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable) GetGoName(yname string) string {
    if yname == "cikeTunnelHistEntry" { return "Ciketunnelhistentry" }
    return ""
}

func (ciketunnelhisttable *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable) GetSegmentPath() string {
    return "cikeTunnelHistTable"
}

func (ciketunnelhisttable *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cikeTunnelHistEntry" {
        for _, c := range ciketunnelhisttable.Ciketunnelhistentry {
            if ciketunnelhisttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry{}
        ciketunnelhisttable.Ciketunnelhistentry = append(ciketunnelhisttable.Ciketunnelhistentry, child)
        return &ciketunnelhisttable.Ciketunnelhistentry[len(ciketunnelhisttable.Ciketunnelhistentry)-1]
    }
    return nil
}

func (ciketunnelhisttable *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ciketunnelhisttable.Ciketunnelhistentry {
        children[ciketunnelhisttable.Ciketunnelhistentry[i].GetSegmentPath()] = &ciketunnelhisttable.Ciketunnelhistentry[i]
    }
    return children
}

func (ciketunnelhisttable *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciketunnelhisttable *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable) GetBundleName() string { return "cisco_ios_xe" }

func (ciketunnelhisttable *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable) GetYangName() string { return "cikeTunnelHistTable" }

func (ciketunnelhisttable *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciketunnelhisttable *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciketunnelhisttable *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciketunnelhisttable *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable) SetParent(parent types.Entity) { ciketunnelhisttable.parent = parent }

func (ciketunnelhisttable *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable) GetParent() types.Entity { return ciketunnelhisttable.parent }

func (ciketunnelhisttable *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable) GetParentYangName() string { return "CISCO-IPSEC-FLOW-MONITOR-MIB" }

// CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry
// Each entry contains the attributes
// associated with a previously active IPsec 
// Phase-1 IKE Tunnel.
type CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry struct {
    parent types.Entity
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

func (ciketunnelhistentry *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry) GetFilter() yfilter.YFilter { return ciketunnelhistentry.YFilter }

func (ciketunnelhistentry *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry) SetFilter(yf yfilter.YFilter) { ciketunnelhistentry.YFilter = yf }

func (ciketunnelhistentry *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry) GetGoName(yname string) string {
    if yname == "cikeTunHistIndex" { return "Ciketunhistindex" }
    if yname == "cikeTunHistTermReason" { return "Ciketunhisttermreason" }
    if yname == "cikeTunHistActiveIndex" { return "Ciketunhistactiveindex" }
    if yname == "cikeTunHistPeerLocalType" { return "Ciketunhistpeerlocaltype" }
    if yname == "cikeTunHistPeerLocalValue" { return "Ciketunhistpeerlocalvalue" }
    if yname == "cikeTunHistPeerIntIndex" { return "Ciketunhistpeerintindex" }
    if yname == "cikeTunHistPeerRemoteType" { return "Ciketunhistpeerremotetype" }
    if yname == "cikeTunHistPeerRemoteValue" { return "Ciketunhistpeerremotevalue" }
    if yname == "cikeTunHistLocalAddr" { return "Ciketunhistlocaladdr" }
    if yname == "cikeTunHistLocalName" { return "Ciketunhistlocalname" }
    if yname == "cikeTunHistRemoteAddr" { return "Ciketunhistremoteaddr" }
    if yname == "cikeTunHistRemoteName" { return "Ciketunhistremotename" }
    if yname == "cikeTunHistNegoMode" { return "Ciketunhistnegomode" }
    if yname == "cikeTunHistDiffHellmanGrp" { return "Ciketunhistdiffhellmangrp" }
    if yname == "cikeTunHistEncryptAlgo" { return "Ciketunhistencryptalgo" }
    if yname == "cikeTunHistHashAlgo" { return "Ciketunhisthashalgo" }
    if yname == "cikeTunHistAuthMethod" { return "Ciketunhistauthmethod" }
    if yname == "cikeTunHistLifeTime" { return "Ciketunhistlifetime" }
    if yname == "cikeTunHistStartTime" { return "Ciketunhiststarttime" }
    if yname == "cikeTunHistActiveTime" { return "Ciketunhistactivetime" }
    if yname == "cikeTunHistTotalRefreshes" { return "Ciketunhisttotalrefreshes" }
    if yname == "cikeTunHistTotalSas" { return "Ciketunhisttotalsas" }
    if yname == "cikeTunHistInOctets" { return "Ciketunhistinoctets" }
    if yname == "cikeTunHistInPkts" { return "Ciketunhistinpkts" }
    if yname == "cikeTunHistInDropPkts" { return "Ciketunhistindroppkts" }
    if yname == "cikeTunHistInNotifys" { return "Ciketunhistinnotifys" }
    if yname == "cikeTunHistInP2Exchgs" { return "Ciketunhistinp2Exchgs" }
    if yname == "cikeTunHistInP2ExchgInvalids" { return "Ciketunhistinp2Exchginvalids" }
    if yname == "cikeTunHistInP2ExchgRejects" { return "Ciketunhistinp2Exchgrejects" }
    if yname == "cikeTunHistInP2SaDelRequests" { return "Ciketunhistinp2Sadelrequests" }
    if yname == "cikeTunHistOutOctets" { return "Ciketunhistoutoctets" }
    if yname == "cikeTunHistOutPkts" { return "Ciketunhistoutpkts" }
    if yname == "cikeTunHistOutDropPkts" { return "Ciketunhistoutdroppkts" }
    if yname == "cikeTunHistOutNotifys" { return "Ciketunhistoutnotifys" }
    if yname == "cikeTunHistOutP2Exchgs" { return "Ciketunhistoutp2Exchgs" }
    if yname == "cikeTunHistOutP2ExchgInvalids" { return "Ciketunhistoutp2Exchginvalids" }
    if yname == "cikeTunHistOutP2ExchgRejects" { return "Ciketunhistoutp2Exchgrejects" }
    if yname == "cikeTunHistOutP2SaDelRequests" { return "Ciketunhistoutp2Sadelrequests" }
    return ""
}

func (ciketunnelhistentry *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry) GetSegmentPath() string {
    return "cikeTunnelHistEntry" + "[cikeTunHistIndex='" + fmt.Sprintf("%v", ciketunnelhistentry.Ciketunhistindex) + "']"
}

func (ciketunnelhistentry *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciketunnelhistentry *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciketunnelhistentry *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cikeTunHistIndex"] = ciketunnelhistentry.Ciketunhistindex
    leafs["cikeTunHistTermReason"] = ciketunnelhistentry.Ciketunhisttermreason
    leafs["cikeTunHistActiveIndex"] = ciketunnelhistentry.Ciketunhistactiveindex
    leafs["cikeTunHistPeerLocalType"] = ciketunnelhistentry.Ciketunhistpeerlocaltype
    leafs["cikeTunHistPeerLocalValue"] = ciketunnelhistentry.Ciketunhistpeerlocalvalue
    leafs["cikeTunHistPeerIntIndex"] = ciketunnelhistentry.Ciketunhistpeerintindex
    leafs["cikeTunHistPeerRemoteType"] = ciketunnelhistentry.Ciketunhistpeerremotetype
    leafs["cikeTunHistPeerRemoteValue"] = ciketunnelhistentry.Ciketunhistpeerremotevalue
    leafs["cikeTunHistLocalAddr"] = ciketunnelhistentry.Ciketunhistlocaladdr
    leafs["cikeTunHistLocalName"] = ciketunnelhistentry.Ciketunhistlocalname
    leafs["cikeTunHistRemoteAddr"] = ciketunnelhistentry.Ciketunhistremoteaddr
    leafs["cikeTunHistRemoteName"] = ciketunnelhistentry.Ciketunhistremotename
    leafs["cikeTunHistNegoMode"] = ciketunnelhistentry.Ciketunhistnegomode
    leafs["cikeTunHistDiffHellmanGrp"] = ciketunnelhistentry.Ciketunhistdiffhellmangrp
    leafs["cikeTunHistEncryptAlgo"] = ciketunnelhistentry.Ciketunhistencryptalgo
    leafs["cikeTunHistHashAlgo"] = ciketunnelhistentry.Ciketunhisthashalgo
    leafs["cikeTunHistAuthMethod"] = ciketunnelhistentry.Ciketunhistauthmethod
    leafs["cikeTunHistLifeTime"] = ciketunnelhistentry.Ciketunhistlifetime
    leafs["cikeTunHistStartTime"] = ciketunnelhistentry.Ciketunhiststarttime
    leafs["cikeTunHistActiveTime"] = ciketunnelhistentry.Ciketunhistactivetime
    leafs["cikeTunHistTotalRefreshes"] = ciketunnelhistentry.Ciketunhisttotalrefreshes
    leafs["cikeTunHistTotalSas"] = ciketunnelhistentry.Ciketunhisttotalsas
    leafs["cikeTunHistInOctets"] = ciketunnelhistentry.Ciketunhistinoctets
    leafs["cikeTunHistInPkts"] = ciketunnelhistentry.Ciketunhistinpkts
    leafs["cikeTunHistInDropPkts"] = ciketunnelhistentry.Ciketunhistindroppkts
    leafs["cikeTunHistInNotifys"] = ciketunnelhistentry.Ciketunhistinnotifys
    leafs["cikeTunHistInP2Exchgs"] = ciketunnelhistentry.Ciketunhistinp2Exchgs
    leafs["cikeTunHistInP2ExchgInvalids"] = ciketunnelhistentry.Ciketunhistinp2Exchginvalids
    leafs["cikeTunHistInP2ExchgRejects"] = ciketunnelhistentry.Ciketunhistinp2Exchgrejects
    leafs["cikeTunHistInP2SaDelRequests"] = ciketunnelhistentry.Ciketunhistinp2Sadelrequests
    leafs["cikeTunHistOutOctets"] = ciketunnelhistentry.Ciketunhistoutoctets
    leafs["cikeTunHistOutPkts"] = ciketunnelhistentry.Ciketunhistoutpkts
    leafs["cikeTunHistOutDropPkts"] = ciketunnelhistentry.Ciketunhistoutdroppkts
    leafs["cikeTunHistOutNotifys"] = ciketunnelhistentry.Ciketunhistoutnotifys
    leafs["cikeTunHistOutP2Exchgs"] = ciketunnelhistentry.Ciketunhistoutp2Exchgs
    leafs["cikeTunHistOutP2ExchgInvalids"] = ciketunnelhistentry.Ciketunhistoutp2Exchginvalids
    leafs["cikeTunHistOutP2ExchgRejects"] = ciketunnelhistentry.Ciketunhistoutp2Exchgrejects
    leafs["cikeTunHistOutP2SaDelRequests"] = ciketunnelhistentry.Ciketunhistoutp2Sadelrequests
    return leafs
}

func (ciketunnelhistentry *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry) GetBundleName() string { return "cisco_ios_xe" }

func (ciketunnelhistentry *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry) GetYangName() string { return "cikeTunnelHistEntry" }

func (ciketunnelhistentry *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciketunnelhistentry *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciketunnelhistentry *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciketunnelhistentry *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry) SetParent(parent types.Entity) { ciketunnelhistentry.parent = parent }

func (ciketunnelhistentry *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry) GetParent() types.Entity { return ciketunnelhistentry.parent }

func (ciketunnelhistentry *CISCOIPSECFLOWMONITORMIB_Ciketunnelhisttable_Ciketunnelhistentry) GetParentYangName() string { return "cikeTunnelHistTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with a previously active
    // IPsec Phase-2 Tunnel. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry.
    Cipsectunnelhistentry []CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry
}

func (cipsectunnelhisttable *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable) GetFilter() yfilter.YFilter { return cipsectunnelhisttable.YFilter }

func (cipsectunnelhisttable *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable) SetFilter(yf yfilter.YFilter) { cipsectunnelhisttable.YFilter = yf }

func (cipsectunnelhisttable *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable) GetGoName(yname string) string {
    if yname == "cipSecTunnelHistEntry" { return "Cipsectunnelhistentry" }
    return ""
}

func (cipsectunnelhisttable *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable) GetSegmentPath() string {
    return "cipSecTunnelHistTable"
}

func (cipsectunnelhisttable *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipSecTunnelHistEntry" {
        for _, c := range cipsectunnelhisttable.Cipsectunnelhistentry {
            if cipsectunnelhisttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry{}
        cipsectunnelhisttable.Cipsectunnelhistentry = append(cipsectunnelhisttable.Cipsectunnelhistentry, child)
        return &cipsectunnelhisttable.Cipsectunnelhistentry[len(cipsectunnelhisttable.Cipsectunnelhistentry)-1]
    }
    return nil
}

func (cipsectunnelhisttable *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cipsectunnelhisttable.Cipsectunnelhistentry {
        children[cipsectunnelhisttable.Cipsectunnelhistentry[i].GetSegmentPath()] = &cipsectunnelhisttable.Cipsectunnelhistentry[i]
    }
    return children
}

func (cipsectunnelhisttable *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cipsectunnelhisttable *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable) GetBundleName() string { return "cisco_ios_xe" }

func (cipsectunnelhisttable *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable) GetYangName() string { return "cipSecTunnelHistTable" }

func (cipsectunnelhisttable *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsectunnelhisttable *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsectunnelhisttable *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsectunnelhisttable *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable) SetParent(parent types.Entity) { cipsectunnelhisttable.parent = parent }

func (cipsectunnelhisttable *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable) GetParent() types.Entity { return cipsectunnelhisttable.parent }

func (cipsectunnelhisttable *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable) GetParentYangName() string { return "CISCO-IPSEC-FLOW-MONITOR-MIB" }

// CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry
// Each entry contains the attributes associated with
// a previously active IPsec Phase-2 Tunnel.
type CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry struct {
    parent types.Entity
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

func (cipsectunnelhistentry *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry) GetFilter() yfilter.YFilter { return cipsectunnelhistentry.YFilter }

func (cipsectunnelhistentry *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry) SetFilter(yf yfilter.YFilter) { cipsectunnelhistentry.YFilter = yf }

func (cipsectunnelhistentry *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry) GetGoName(yname string) string {
    if yname == "cipSecTunHistIndex" { return "Cipsectunhistindex" }
    if yname == "cipSecTunHistTermReason" { return "Cipsectunhisttermreason" }
    if yname == "cipSecTunHistActiveIndex" { return "Cipsectunhistactiveindex" }
    if yname == "cipSecTunHistIkeTunnelIndex" { return "Cipsectunhistiketunnelindex" }
    if yname == "cipSecTunHistLocalAddr" { return "Cipsectunhistlocaladdr" }
    if yname == "cipSecTunHistRemoteAddr" { return "Cipsectunhistremoteaddr" }
    if yname == "cipSecTunHistKeyType" { return "Cipsectunhistkeytype" }
    if yname == "cipSecTunHistEncapMode" { return "Cipsectunhistencapmode" }
    if yname == "cipSecTunHistLifeSize" { return "Cipsectunhistlifesize" }
    if yname == "cipSecTunHistLifeTime" { return "Cipsectunhistlifetime" }
    if yname == "cipSecTunHistStartTime" { return "Cipsectunhiststarttime" }
    if yname == "cipSecTunHistActiveTime" { return "Cipsectunhistactivetime" }
    if yname == "cipSecTunHistTotalRefreshes" { return "Cipsectunhisttotalrefreshes" }
    if yname == "cipSecTunHistTotalSas" { return "Cipsectunhisttotalsas" }
    if yname == "cipSecTunHistInSaDiffHellmanGrp" { return "Cipsectunhistinsadiffhellmangrp" }
    if yname == "cipSecTunHistInSaEncryptAlgo" { return "Cipsectunhistinsaencryptalgo" }
    if yname == "cipSecTunHistInSaAhAuthAlgo" { return "Cipsectunhistinsaahauthalgo" }
    if yname == "cipSecTunHistInSaEspAuthAlgo" { return "Cipsectunhistinsaespauthalgo" }
    if yname == "cipSecTunHistInSaDecompAlgo" { return "Cipsectunhistinsadecompalgo" }
    if yname == "cipSecTunHistOutSaDiffHellmanGrp" { return "Cipsectunhistoutsadiffhellmangrp" }
    if yname == "cipSecTunHistOutSaEncryptAlgo" { return "Cipsectunhistoutsaencryptalgo" }
    if yname == "cipSecTunHistOutSaAhAuthAlgo" { return "Cipsectunhistoutsaahauthalgo" }
    if yname == "cipSecTunHistOutSaEspAuthAlgo" { return "Cipsectunhistoutsaespauthalgo" }
    if yname == "cipSecTunHistOutSaCompAlgo" { return "Cipsectunhistoutsacompalgo" }
    if yname == "cipSecTunHistInOctets" { return "Cipsectunhistinoctets" }
    if yname == "cipSecTunHistHcInOctets" { return "Cipsectunhisthcinoctets" }
    if yname == "cipSecTunHistInOctWraps" { return "Cipsectunhistinoctwraps" }
    if yname == "cipSecTunHistInDecompOctets" { return "Cipsectunhistindecompoctets" }
    if yname == "cipSecTunHistHcInDecompOctets" { return "Cipsectunhisthcindecompoctets" }
    if yname == "cipSecTunHistInDecompOctWraps" { return "Cipsectunhistindecompoctwraps" }
    if yname == "cipSecTunHistInPkts" { return "Cipsectunhistinpkts" }
    if yname == "cipSecTunHistInDropPkts" { return "Cipsectunhistindroppkts" }
    if yname == "cipSecTunHistInReplayDropPkts" { return "Cipsectunhistinreplaydroppkts" }
    if yname == "cipSecTunHistInAuths" { return "Cipsectunhistinauths" }
    if yname == "cipSecTunHistInAuthFails" { return "Cipsectunhistinauthfails" }
    if yname == "cipSecTunHistInDecrypts" { return "Cipsectunhistindecrypts" }
    if yname == "cipSecTunHistInDecryptFails" { return "Cipsectunhistindecryptfails" }
    if yname == "cipSecTunHistOutOctets" { return "Cipsectunhistoutoctets" }
    if yname == "cipSecTunHistHcOutOctets" { return "Cipsectunhisthcoutoctets" }
    if yname == "cipSecTunHistOutOctWraps" { return "Cipsectunhistoutoctwraps" }
    if yname == "cipSecTunHistOutUncompOctets" { return "Cipsectunhistoutuncompoctets" }
    if yname == "cipSecTunHistHcOutUncompOctets" { return "Cipsectunhisthcoutuncompoctets" }
    if yname == "cipSecTunHistOutUncompOctWraps" { return "Cipsectunhistoutuncompoctwraps" }
    if yname == "cipSecTunHistOutPkts" { return "Cipsectunhistoutpkts" }
    if yname == "cipSecTunHistOutDropPkts" { return "Cipsectunhistoutdroppkts" }
    if yname == "cipSecTunHistOutAuths" { return "Cipsectunhistoutauths" }
    if yname == "cipSecTunHistOutAuthFails" { return "Cipsectunhistoutauthfails" }
    if yname == "cipSecTunHistOutEncrypts" { return "Cipsectunhistoutencrypts" }
    if yname == "cipSecTunHistOutEncryptFails" { return "Cipsectunhistoutencryptfails" }
    return ""
}

func (cipsectunnelhistentry *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry) GetSegmentPath() string {
    return "cipSecTunnelHistEntry" + "[cipSecTunHistIndex='" + fmt.Sprintf("%v", cipsectunnelhistentry.Cipsectunhistindex) + "']"
}

func (cipsectunnelhistentry *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipsectunnelhistentry *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipsectunnelhistentry *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipSecTunHistIndex"] = cipsectunnelhistentry.Cipsectunhistindex
    leafs["cipSecTunHistTermReason"] = cipsectunnelhistentry.Cipsectunhisttermreason
    leafs["cipSecTunHistActiveIndex"] = cipsectunnelhistentry.Cipsectunhistactiveindex
    leafs["cipSecTunHistIkeTunnelIndex"] = cipsectunnelhistentry.Cipsectunhistiketunnelindex
    leafs["cipSecTunHistLocalAddr"] = cipsectunnelhistentry.Cipsectunhistlocaladdr
    leafs["cipSecTunHistRemoteAddr"] = cipsectunnelhistentry.Cipsectunhistremoteaddr
    leafs["cipSecTunHistKeyType"] = cipsectunnelhistentry.Cipsectunhistkeytype
    leafs["cipSecTunHistEncapMode"] = cipsectunnelhistentry.Cipsectunhistencapmode
    leafs["cipSecTunHistLifeSize"] = cipsectunnelhistentry.Cipsectunhistlifesize
    leafs["cipSecTunHistLifeTime"] = cipsectunnelhistentry.Cipsectunhistlifetime
    leafs["cipSecTunHistStartTime"] = cipsectunnelhistentry.Cipsectunhiststarttime
    leafs["cipSecTunHistActiveTime"] = cipsectunnelhistentry.Cipsectunhistactivetime
    leafs["cipSecTunHistTotalRefreshes"] = cipsectunnelhistentry.Cipsectunhisttotalrefreshes
    leafs["cipSecTunHistTotalSas"] = cipsectunnelhistentry.Cipsectunhisttotalsas
    leafs["cipSecTunHistInSaDiffHellmanGrp"] = cipsectunnelhistentry.Cipsectunhistinsadiffhellmangrp
    leafs["cipSecTunHistInSaEncryptAlgo"] = cipsectunnelhistentry.Cipsectunhistinsaencryptalgo
    leafs["cipSecTunHistInSaAhAuthAlgo"] = cipsectunnelhistentry.Cipsectunhistinsaahauthalgo
    leafs["cipSecTunHistInSaEspAuthAlgo"] = cipsectunnelhistentry.Cipsectunhistinsaespauthalgo
    leafs["cipSecTunHistInSaDecompAlgo"] = cipsectunnelhistentry.Cipsectunhistinsadecompalgo
    leafs["cipSecTunHistOutSaDiffHellmanGrp"] = cipsectunnelhistentry.Cipsectunhistoutsadiffhellmangrp
    leafs["cipSecTunHistOutSaEncryptAlgo"] = cipsectunnelhistentry.Cipsectunhistoutsaencryptalgo
    leafs["cipSecTunHistOutSaAhAuthAlgo"] = cipsectunnelhistentry.Cipsectunhistoutsaahauthalgo
    leafs["cipSecTunHistOutSaEspAuthAlgo"] = cipsectunnelhistentry.Cipsectunhistoutsaespauthalgo
    leafs["cipSecTunHistOutSaCompAlgo"] = cipsectunnelhistentry.Cipsectunhistoutsacompalgo
    leafs["cipSecTunHistInOctets"] = cipsectunnelhistentry.Cipsectunhistinoctets
    leafs["cipSecTunHistHcInOctets"] = cipsectunnelhistentry.Cipsectunhisthcinoctets
    leafs["cipSecTunHistInOctWraps"] = cipsectunnelhistentry.Cipsectunhistinoctwraps
    leafs["cipSecTunHistInDecompOctets"] = cipsectunnelhistentry.Cipsectunhistindecompoctets
    leafs["cipSecTunHistHcInDecompOctets"] = cipsectunnelhistentry.Cipsectunhisthcindecompoctets
    leafs["cipSecTunHistInDecompOctWraps"] = cipsectunnelhistentry.Cipsectunhistindecompoctwraps
    leafs["cipSecTunHistInPkts"] = cipsectunnelhistentry.Cipsectunhistinpkts
    leafs["cipSecTunHistInDropPkts"] = cipsectunnelhistentry.Cipsectunhistindroppkts
    leafs["cipSecTunHistInReplayDropPkts"] = cipsectunnelhistentry.Cipsectunhistinreplaydroppkts
    leafs["cipSecTunHistInAuths"] = cipsectunnelhistentry.Cipsectunhistinauths
    leafs["cipSecTunHistInAuthFails"] = cipsectunnelhistentry.Cipsectunhistinauthfails
    leafs["cipSecTunHistInDecrypts"] = cipsectunnelhistentry.Cipsectunhistindecrypts
    leafs["cipSecTunHistInDecryptFails"] = cipsectunnelhistentry.Cipsectunhistindecryptfails
    leafs["cipSecTunHistOutOctets"] = cipsectunnelhistentry.Cipsectunhistoutoctets
    leafs["cipSecTunHistHcOutOctets"] = cipsectunnelhistentry.Cipsectunhisthcoutoctets
    leafs["cipSecTunHistOutOctWraps"] = cipsectunnelhistentry.Cipsectunhistoutoctwraps
    leafs["cipSecTunHistOutUncompOctets"] = cipsectunnelhistentry.Cipsectunhistoutuncompoctets
    leafs["cipSecTunHistHcOutUncompOctets"] = cipsectunnelhistentry.Cipsectunhisthcoutuncompoctets
    leafs["cipSecTunHistOutUncompOctWraps"] = cipsectunnelhistentry.Cipsectunhistoutuncompoctwraps
    leafs["cipSecTunHistOutPkts"] = cipsectunnelhistentry.Cipsectunhistoutpkts
    leafs["cipSecTunHistOutDropPkts"] = cipsectunnelhistentry.Cipsectunhistoutdroppkts
    leafs["cipSecTunHistOutAuths"] = cipsectunnelhistentry.Cipsectunhistoutauths
    leafs["cipSecTunHistOutAuthFails"] = cipsectunnelhistentry.Cipsectunhistoutauthfails
    leafs["cipSecTunHistOutEncrypts"] = cipsectunnelhistentry.Cipsectunhistoutencrypts
    leafs["cipSecTunHistOutEncryptFails"] = cipsectunnelhistentry.Cipsectunhistoutencryptfails
    return leafs
}

func (cipsectunnelhistentry *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry) GetBundleName() string { return "cisco_ios_xe" }

func (cipsectunnelhistentry *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry) GetYangName() string { return "cipSecTunnelHistEntry" }

func (cipsectunnelhistentry *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsectunnelhistentry *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsectunnelhistentry *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsectunnelhistentry *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry) SetParent(parent types.Entity) { cipsectunnelhistentry.parent = parent }

func (cipsectunnelhistentry *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry) GetParent() types.Entity { return cipsectunnelhistentry.parent }

func (cipsectunnelhistentry *CISCOIPSECFLOWMONITORMIB_Cipsectunnelhisttable_Cipsectunnelhistentry) GetParentYangName() string { return "cipSecTunnelHistTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with a previously active
    // IPsec Phase-2 Tunnel Endpoint. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable_Cipsecendpthistentry.
    Cipsecendpthistentry []CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable_Cipsecendpthistentry
}

func (cipsecendpthisttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable) GetFilter() yfilter.YFilter { return cipsecendpthisttable.YFilter }

func (cipsecendpthisttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable) SetFilter(yf yfilter.YFilter) { cipsecendpthisttable.YFilter = yf }

func (cipsecendpthisttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable) GetGoName(yname string) string {
    if yname == "cipSecEndPtHistEntry" { return "Cipsecendpthistentry" }
    return ""
}

func (cipsecendpthisttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable) GetSegmentPath() string {
    return "cipSecEndPtHistTable"
}

func (cipsecendpthisttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipSecEndPtHistEntry" {
        for _, c := range cipsecendpthisttable.Cipsecendpthistentry {
            if cipsecendpthisttable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable_Cipsecendpthistentry{}
        cipsecendpthisttable.Cipsecendpthistentry = append(cipsecendpthisttable.Cipsecendpthistentry, child)
        return &cipsecendpthisttable.Cipsecendpthistentry[len(cipsecendpthisttable.Cipsecendpthistentry)-1]
    }
    return nil
}

func (cipsecendpthisttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cipsecendpthisttable.Cipsecendpthistentry {
        children[cipsecendpthisttable.Cipsecendpthistentry[i].GetSegmentPath()] = &cipsecendpthisttable.Cipsecendpthistentry[i]
    }
    return children
}

func (cipsecendpthisttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cipsecendpthisttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable) GetBundleName() string { return "cisco_ios_xe" }

func (cipsecendpthisttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable) GetYangName() string { return "cipSecEndPtHistTable" }

func (cipsecendpthisttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsecendpthisttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsecendpthisttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsecendpthisttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable) SetParent(parent types.Entity) { cipsecendpthisttable.parent = parent }

func (cipsecendpthisttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable) GetParent() types.Entity { return cipsecendpthisttable.parent }

func (cipsecendpthisttable *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable) GetParentYangName() string { return "CISCO-IPSEC-FLOW-MONITOR-MIB" }

// CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable_Cipsecendpthistentry
// Each entry contains the attributes associated with
// a previously active IPsec Phase-2 Tunnel Endpoint.
type CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable_Cipsecendpthistentry struct {
    parent types.Entity
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

func (cipsecendpthistentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable_Cipsecendpthistentry) GetFilter() yfilter.YFilter { return cipsecendpthistentry.YFilter }

func (cipsecendpthistentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable_Cipsecendpthistentry) SetFilter(yf yfilter.YFilter) { cipsecendpthistentry.YFilter = yf }

func (cipsecendpthistentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable_Cipsecendpthistentry) GetGoName(yname string) string {
    if yname == "cipSecEndPtHistIndex" { return "Cipsecendpthistindex" }
    if yname == "cipSecEndPtHistTunIndex" { return "Cipsecendpthisttunindex" }
    if yname == "cipSecEndPtHistActiveIndex" { return "Cipsecendpthistactiveindex" }
    if yname == "cipSecEndPtHistLocalName" { return "Cipsecendpthistlocalname" }
    if yname == "cipSecEndPtHistLocalType" { return "Cipsecendpthistlocaltype" }
    if yname == "cipSecEndPtHistLocalAddr1" { return "Cipsecendpthistlocaladdr1" }
    if yname == "cipSecEndPtHistLocalAddr2" { return "Cipsecendpthistlocaladdr2" }
    if yname == "cipSecEndPtHistLocalProtocol" { return "Cipsecendpthistlocalprotocol" }
    if yname == "cipSecEndPtHistLocalPort" { return "Cipsecendpthistlocalport" }
    if yname == "cipSecEndPtHistRemoteName" { return "Cipsecendpthistremotename" }
    if yname == "cipSecEndPtHistRemoteType" { return "Cipsecendpthistremotetype" }
    if yname == "cipSecEndPtHistRemoteAddr1" { return "Cipsecendpthistremoteaddr1" }
    if yname == "cipSecEndPtHistRemoteAddr2" { return "Cipsecendpthistremoteaddr2" }
    if yname == "cipSecEndPtHistRemoteProtocol" { return "Cipsecendpthistremoteprotocol" }
    if yname == "cipSecEndPtHistRemotePort" { return "Cipsecendpthistremoteport" }
    return ""
}

func (cipsecendpthistentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable_Cipsecendpthistentry) GetSegmentPath() string {
    return "cipSecEndPtHistEntry" + "[cipSecEndPtHistIndex='" + fmt.Sprintf("%v", cipsecendpthistentry.Cipsecendpthistindex) + "']"
}

func (cipsecendpthistentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable_Cipsecendpthistentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipsecendpthistentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable_Cipsecendpthistentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipsecendpthistentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable_Cipsecendpthistentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipSecEndPtHistIndex"] = cipsecendpthistentry.Cipsecendpthistindex
    leafs["cipSecEndPtHistTunIndex"] = cipsecendpthistentry.Cipsecendpthisttunindex
    leafs["cipSecEndPtHistActiveIndex"] = cipsecendpthistentry.Cipsecendpthistactiveindex
    leafs["cipSecEndPtHistLocalName"] = cipsecendpthistentry.Cipsecendpthistlocalname
    leafs["cipSecEndPtHistLocalType"] = cipsecendpthistentry.Cipsecendpthistlocaltype
    leafs["cipSecEndPtHistLocalAddr1"] = cipsecendpthistentry.Cipsecendpthistlocaladdr1
    leafs["cipSecEndPtHistLocalAddr2"] = cipsecendpthistentry.Cipsecendpthistlocaladdr2
    leafs["cipSecEndPtHistLocalProtocol"] = cipsecendpthistentry.Cipsecendpthistlocalprotocol
    leafs["cipSecEndPtHistLocalPort"] = cipsecendpthistentry.Cipsecendpthistlocalport
    leafs["cipSecEndPtHistRemoteName"] = cipsecendpthistentry.Cipsecendpthistremotename
    leafs["cipSecEndPtHistRemoteType"] = cipsecendpthistentry.Cipsecendpthistremotetype
    leafs["cipSecEndPtHistRemoteAddr1"] = cipsecendpthistentry.Cipsecendpthistremoteaddr1
    leafs["cipSecEndPtHistRemoteAddr2"] = cipsecendpthistentry.Cipsecendpthistremoteaddr2
    leafs["cipSecEndPtHistRemoteProtocol"] = cipsecendpthistentry.Cipsecendpthistremoteprotocol
    leafs["cipSecEndPtHistRemotePort"] = cipsecendpthistentry.Cipsecendpthistremoteport
    return leafs
}

func (cipsecendpthistentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable_Cipsecendpthistentry) GetBundleName() string { return "cisco_ios_xe" }

func (cipsecendpthistentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable_Cipsecendpthistentry) GetYangName() string { return "cipSecEndPtHistEntry" }

func (cipsecendpthistentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable_Cipsecendpthistentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsecendpthistentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable_Cipsecendpthistentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsecendpthistentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable_Cipsecendpthistentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsecendpthistentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable_Cipsecendpthistentry) SetParent(parent types.Entity) { cipsecendpthistentry.parent = parent }

func (cipsecendpthistentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable_Cipsecendpthistentry) GetParent() types.Entity { return cipsecendpthistentry.parent }

func (cipsecendpthistentry *CISCOIPSECFLOWMONITORMIB_Cipsecendpthisttable_Cipsecendpthistentry) GetParentYangName() string { return "cipSecEndPtHistTable" }

// CISCOIPSECFLOWMONITORMIB_Cikefailtable
// The IPsec Phase-1 Failure Table.
// This table is implemented as a sliding 
// window in which only the last n entries are 
// maintained.  The maximum number of entries
// is specified by the cipSecFailTableSize object.
type CISCOIPSECFLOWMONITORMIB_Cikefailtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with  an IPsec Phase-1
    // failure. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry.
    Cikefailentry []CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry
}

func (cikefailtable *CISCOIPSECFLOWMONITORMIB_Cikefailtable) GetFilter() yfilter.YFilter { return cikefailtable.YFilter }

func (cikefailtable *CISCOIPSECFLOWMONITORMIB_Cikefailtable) SetFilter(yf yfilter.YFilter) { cikefailtable.YFilter = yf }

func (cikefailtable *CISCOIPSECFLOWMONITORMIB_Cikefailtable) GetGoName(yname string) string {
    if yname == "cikeFailEntry" { return "Cikefailentry" }
    return ""
}

func (cikefailtable *CISCOIPSECFLOWMONITORMIB_Cikefailtable) GetSegmentPath() string {
    return "cikeFailTable"
}

func (cikefailtable *CISCOIPSECFLOWMONITORMIB_Cikefailtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cikeFailEntry" {
        for _, c := range cikefailtable.Cikefailentry {
            if cikefailtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry{}
        cikefailtable.Cikefailentry = append(cikefailtable.Cikefailentry, child)
        return &cikefailtable.Cikefailentry[len(cikefailtable.Cikefailentry)-1]
    }
    return nil
}

func (cikefailtable *CISCOIPSECFLOWMONITORMIB_Cikefailtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cikefailtable.Cikefailentry {
        children[cikefailtable.Cikefailentry[i].GetSegmentPath()] = &cikefailtable.Cikefailentry[i]
    }
    return children
}

func (cikefailtable *CISCOIPSECFLOWMONITORMIB_Cikefailtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cikefailtable *CISCOIPSECFLOWMONITORMIB_Cikefailtable) GetBundleName() string { return "cisco_ios_xe" }

func (cikefailtable *CISCOIPSECFLOWMONITORMIB_Cikefailtable) GetYangName() string { return "cikeFailTable" }

func (cikefailtable *CISCOIPSECFLOWMONITORMIB_Cikefailtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cikefailtable *CISCOIPSECFLOWMONITORMIB_Cikefailtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cikefailtable *CISCOIPSECFLOWMONITORMIB_Cikefailtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cikefailtable *CISCOIPSECFLOWMONITORMIB_Cikefailtable) SetParent(parent types.Entity) { cikefailtable.parent = parent }

func (cikefailtable *CISCOIPSECFLOWMONITORMIB_Cikefailtable) GetParent() types.Entity { return cikefailtable.parent }

func (cikefailtable *CISCOIPSECFLOWMONITORMIB_Cikefailtable) GetParentYangName() string { return "CISCO-IPSEC-FLOW-MONITOR-MIB" }

// CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry
// Each entry contains the attributes associated
// with
//  an IPsec Phase-1 failure.
type CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry struct {
    parent types.Entity
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

func (cikefailentry *CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry) GetFilter() yfilter.YFilter { return cikefailentry.YFilter }

func (cikefailentry *CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry) SetFilter(yf yfilter.YFilter) { cikefailentry.YFilter = yf }

func (cikefailentry *CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry) GetGoName(yname string) string {
    if yname == "cikeFailIndex" { return "Cikefailindex" }
    if yname == "cikeFailReason" { return "Cikefailreason" }
    if yname == "cikeFailTime" { return "Cikefailtime" }
    if yname == "cikeFailLocalType" { return "Cikefaillocaltype" }
    if yname == "cikeFailLocalValue" { return "Cikefaillocalvalue" }
    if yname == "cikeFailRemoteType" { return "Cikefailremotetype" }
    if yname == "cikeFailRemoteValue" { return "Cikefailremotevalue" }
    if yname == "cikeFailLocalAddr" { return "Cikefaillocaladdr" }
    if yname == "cikeFailRemoteAddr" { return "Cikefailremoteaddr" }
    return ""
}

func (cikefailentry *CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry) GetSegmentPath() string {
    return "cikeFailEntry" + "[cikeFailIndex='" + fmt.Sprintf("%v", cikefailentry.Cikefailindex) + "']"
}

func (cikefailentry *CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cikefailentry *CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cikefailentry *CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cikeFailIndex"] = cikefailentry.Cikefailindex
    leafs["cikeFailReason"] = cikefailentry.Cikefailreason
    leafs["cikeFailTime"] = cikefailentry.Cikefailtime
    leafs["cikeFailLocalType"] = cikefailentry.Cikefaillocaltype
    leafs["cikeFailLocalValue"] = cikefailentry.Cikefaillocalvalue
    leafs["cikeFailRemoteType"] = cikefailentry.Cikefailremotetype
    leafs["cikeFailRemoteValue"] = cikefailentry.Cikefailremotevalue
    leafs["cikeFailLocalAddr"] = cikefailentry.Cikefaillocaladdr
    leafs["cikeFailRemoteAddr"] = cikefailentry.Cikefailremoteaddr
    return leafs
}

func (cikefailentry *CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry) GetBundleName() string { return "cisco_ios_xe" }

func (cikefailentry *CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry) GetYangName() string { return "cikeFailEntry" }

func (cikefailentry *CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cikefailentry *CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cikefailentry *CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cikefailentry *CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry) SetParent(parent types.Entity) { cikefailentry.parent = parent }

func (cikefailentry *CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry) GetParent() types.Entity { return cikefailentry.parent }

func (cikefailentry *CISCOIPSECFLOWMONITORMIB_Cikefailtable_Cikefailentry) GetParentYangName() string { return "cikeFailTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with an IPsec Phase-1
    // failure. The type is slice of
    // CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry.
    Cipsecfailentry []CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry
}

func (cipsecfailtable *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable) GetFilter() yfilter.YFilter { return cipsecfailtable.YFilter }

func (cipsecfailtable *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable) SetFilter(yf yfilter.YFilter) { cipsecfailtable.YFilter = yf }

func (cipsecfailtable *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable) GetGoName(yname string) string {
    if yname == "cipSecFailEntry" { return "Cipsecfailentry" }
    return ""
}

func (cipsecfailtable *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable) GetSegmentPath() string {
    return "cipSecFailTable"
}

func (cipsecfailtable *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipSecFailEntry" {
        for _, c := range cipsecfailtable.Cipsecfailentry {
            if cipsecfailtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry{}
        cipsecfailtable.Cipsecfailentry = append(cipsecfailtable.Cipsecfailentry, child)
        return &cipsecfailtable.Cipsecfailentry[len(cipsecfailtable.Cipsecfailentry)-1]
    }
    return nil
}

func (cipsecfailtable *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cipsecfailtable.Cipsecfailentry {
        children[cipsecfailtable.Cipsecfailentry[i].GetSegmentPath()] = &cipsecfailtable.Cipsecfailentry[i]
    }
    return children
}

func (cipsecfailtable *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cipsecfailtable *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable) GetBundleName() string { return "cisco_ios_xe" }

func (cipsecfailtable *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable) GetYangName() string { return "cipSecFailTable" }

func (cipsecfailtable *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsecfailtable *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsecfailtable *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsecfailtable *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable) SetParent(parent types.Entity) { cipsecfailtable.parent = parent }

func (cipsecfailtable *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable) GetParent() types.Entity { return cipsecfailtable.parent }

func (cipsecfailtable *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable) GetParentYangName() string { return "CISCO-IPSEC-FLOW-MONITOR-MIB" }

// CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry
// Each entry contains the attributes associated with
// an IPsec Phase-1 failure.
type CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry struct {
    parent types.Entity
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

func (cipsecfailentry *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry) GetFilter() yfilter.YFilter { return cipsecfailentry.YFilter }

func (cipsecfailentry *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry) SetFilter(yf yfilter.YFilter) { cipsecfailentry.YFilter = yf }

func (cipsecfailentry *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry) GetGoName(yname string) string {
    if yname == "cipSecFailIndex" { return "Cipsecfailindex" }
    if yname == "cipSecFailReason" { return "Cipsecfailreason" }
    if yname == "cipSecFailTime" { return "Cipsecfailtime" }
    if yname == "cipSecFailTunnelIndex" { return "Cipsecfailtunnelindex" }
    if yname == "cipSecFailSaSpi" { return "Cipsecfailsaspi" }
    if yname == "cipSecFailPktSrcAddr" { return "Cipsecfailpktsrcaddr" }
    if yname == "cipSecFailPktDstAddr" { return "Cipsecfailpktdstaddr" }
    return ""
}

func (cipsecfailentry *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry) GetSegmentPath() string {
    return "cipSecFailEntry" + "[cipSecFailIndex='" + fmt.Sprintf("%v", cipsecfailentry.Cipsecfailindex) + "']"
}

func (cipsecfailentry *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipsecfailentry *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipsecfailentry *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipSecFailIndex"] = cipsecfailentry.Cipsecfailindex
    leafs["cipSecFailReason"] = cipsecfailentry.Cipsecfailreason
    leafs["cipSecFailTime"] = cipsecfailentry.Cipsecfailtime
    leafs["cipSecFailTunnelIndex"] = cipsecfailentry.Cipsecfailtunnelindex
    leafs["cipSecFailSaSpi"] = cipsecfailentry.Cipsecfailsaspi
    leafs["cipSecFailPktSrcAddr"] = cipsecfailentry.Cipsecfailpktsrcaddr
    leafs["cipSecFailPktDstAddr"] = cipsecfailentry.Cipsecfailpktdstaddr
    return leafs
}

func (cipsecfailentry *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry) GetBundleName() string { return "cisco_ios_xe" }

func (cipsecfailentry *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry) GetYangName() string { return "cipSecFailEntry" }

func (cipsecfailentry *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsecfailentry *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsecfailentry *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsecfailentry *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry) SetParent(parent types.Entity) { cipsecfailentry.parent = parent }

func (cipsecfailentry *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry) GetParent() types.Entity { return cipsecfailentry.parent }

func (cipsecfailentry *CISCOIPSECFLOWMONITORMIB_Cipsecfailtable_Cipsecfailentry) GetParentYangName() string { return "cipSecFailTable" }

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

