// The MIB module for modeling Cisco-specific 
// IPsec attributes
// 
// Overview of Cisco IPsec MIB
// 
// MIB description
// 
// This MIB models the Cisco implementation-specific 
// attributes of a Cisco entity that implements IPsec. 
// This MIB is complementary to the standard IPsec MIB 
// proposed jointly by Tivoli and Cisco.
// 
// The ciscoIPsec MIB provides the operational information 
// on Cisco's IPsec tunnelling implementation.  
// The following entities are managed:
// 1) ISAKMP Group:
// a) ISAKMP global parameters
// b) ISAKMP Policy Table
// 
// 2) IPSec Group:
// a) IPSec Global Parameters
// b) IPSec Global Traffic Parameters
// c) Cryptomap Group
// - Cryptomap Set Table
// - Cryptomap Table
// - CryptomapSet Binding Table
// 
// 3) System Capacity & Capability Group:
// a) Capacity Parameters
// b) Capability Parameters
// 
// 4) Trap Control Group
// 5) Notifications Group
package cisco_ipsec_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ipsec_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IPSEC-MIB CISCO-IPSEC-MIB}", reflect.TypeOf(CISCOIPSECMIB{}))
    ydk.RegisterEntity("CISCO-IPSEC-MIB:CISCO-IPSEC-MIB", reflect.TypeOf(CISCOIPSECMIB{}))
}

// IkeHashAlgo represents IKE negotiations.
type IkeHashAlgo string

const (
    IkeHashAlgo_none IkeHashAlgo = "none"

    IkeHashAlgo_md5 IkeHashAlgo = "md5"

    IkeHashAlgo_sha IkeHashAlgo = "sha"
)

// CryptomapType represents is a unit of IOS IPSec policy specification.
type CryptomapType string

const (
    CryptomapType_cryptomapTypeNONE CryptomapType = "cryptomapTypeNONE"

    CryptomapType_cryptomapTypeMANUAL CryptomapType = "cryptomapTypeMANUAL"

    CryptomapType_cryptomapTypeISAKMP CryptomapType = "cryptomapTypeISAKMP"

    CryptomapType_cryptomapTypeCET CryptomapType = "cryptomapTypeCET"

    CryptomapType_cryptomapTypeDYNAMIC CryptomapType = "cryptomapTypeDYNAMIC"

    CryptomapType_cryptomapTypeDYNAMICDISCOVERY CryptomapType = "cryptomapTypeDYNAMICDISCOVERY"
)

// IkeIdentityType represents 	Main Mode of IPSec tunnel setup.
type IkeIdentityType string

const (
    IkeIdentityType_isakmpIdTypeUNKNOWN IkeIdentityType = "isakmpIdTypeUNKNOWN"

    IkeIdentityType_isakmpIdTypeADDRESS IkeIdentityType = "isakmpIdTypeADDRESS"

    IkeIdentityType_isakmpIdTypeHOSTNAME IkeIdentityType = "isakmpIdTypeHOSTNAME"
)

// TrapStatus represents The administrative status for sending a TRAP.
type TrapStatus string

const (
    TrapStatus_enabled TrapStatus = "enabled"

    TrapStatus_disabled TrapStatus = "disabled"
)

// EncryptAlgo represents The encryption algorithm used in negotiations.
type EncryptAlgo string

const (
    EncryptAlgo_none EncryptAlgo = "none"

    EncryptAlgo_des EncryptAlgo = "des"

    EncryptAlgo_des3 EncryptAlgo = "des3"
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

// CryptomapSetBindStatus represents SNMP General Error.
type CryptomapSetBindStatus string

const (
    CryptomapSetBindStatus_unknown CryptomapSetBindStatus = "unknown"

    CryptomapSetBindStatus_attached CryptomapSetBindStatus = "attached"

    CryptomapSetBindStatus_detached CryptomapSetBindStatus = "detached"
)

// DiffHellmanGrp represents The Diffie Hellman Group used in negotiations.
type DiffHellmanGrp string

const (
    DiffHellmanGrp_none DiffHellmanGrp = "none"

    DiffHellmanGrp_dhGroup1 DiffHellmanGrp = "dhGroup1"

    DiffHellmanGrp_dhGroup2 DiffHellmanGrp = "dhGroup2"
)

// CISCOIPSECMIB
type CISCOIPSECMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Cipsisakmpgroup CISCOIPSECMIB_Cipsisakmpgroup

    
    Cipsipsecglobals CISCOIPSECMIB_Cipsipsecglobals

    
    Cipsipsecstatistics CISCOIPSECMIB_Cipsipsecstatistics

    
    Cipssyscapacitygroup CISCOIPSECMIB_Cipssyscapacitygroup

    
    Cipstrapcntlgroup CISCOIPSECMIB_Cipstrapcntlgroup

    // The table containing the list of all ISAKMP policy entries configured by
    // the operator.
    Cipsisakmppolicytable CISCOIPSECMIB_Cipsisakmppolicytable

    // The table containing the list of all cryptomap sets that are fully
    // specified and are not wild-carded.  The operator may include different
    // types of cryptomaps in such a set - manual, CET, ISAKMP or dynamic.
    Cipsstaticcryptomapsettable CISCOIPSECMIB_Cipsstaticcryptomapsettable

    // The table containing the list of all dynamic cryptomaps that use IKE,
    // defined on   the managed entity.
    Cipsdynamiccryptomapsettable CISCOIPSECMIB_Cipsdynamiccryptomapsettable

    // The table ilisting the member cryptomaps of the cryptomap sets that are
    // configured on the managed entity.
    Cipsstaticcryptomaptable CISCOIPSECMIB_Cipsstaticcryptomaptable

    // The table lists the binding of cryptomap sets to the interfaces of the
    // managed entity.
    Cipscryptomapsetiftable CISCOIPSECMIB_Cipscryptomapsetiftable
}

func (cISCOIPSECMIB *CISCOIPSECMIB) GetFilter() yfilter.YFilter { return cISCOIPSECMIB.YFilter }

func (cISCOIPSECMIB *CISCOIPSECMIB) SetFilter(yf yfilter.YFilter) { cISCOIPSECMIB.YFilter = yf }

func (cISCOIPSECMIB *CISCOIPSECMIB) GetGoName(yname string) string {
    if yname == "cipsIsakmpGroup" { return "Cipsisakmpgroup" }
    if yname == "cipsIPsecGlobals" { return "Cipsipsecglobals" }
    if yname == "cipsIPsecStatistics" { return "Cipsipsecstatistics" }
    if yname == "cipsSysCapacityGroup" { return "Cipssyscapacitygroup" }
    if yname == "cipsTrapCntlGroup" { return "Cipstrapcntlgroup" }
    if yname == "cipsIsakmpPolicyTable" { return "Cipsisakmppolicytable" }
    if yname == "cipsStaticCryptomapSetTable" { return "Cipsstaticcryptomapsettable" }
    if yname == "cipsDynamicCryptomapSetTable" { return "Cipsdynamiccryptomapsettable" }
    if yname == "cipsStaticCryptomapTable" { return "Cipsstaticcryptomaptable" }
    if yname == "cipsCryptomapSetIfTable" { return "Cipscryptomapsetiftable" }
    return ""
}

func (cISCOIPSECMIB *CISCOIPSECMIB) GetSegmentPath() string {
    return "CISCO-IPSEC-MIB:CISCO-IPSEC-MIB"
}

func (cISCOIPSECMIB *CISCOIPSECMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipsIsakmpGroup" {
        return &cISCOIPSECMIB.Cipsisakmpgroup
    }
    if childYangName == "cipsIPsecGlobals" {
        return &cISCOIPSECMIB.Cipsipsecglobals
    }
    if childYangName == "cipsIPsecStatistics" {
        return &cISCOIPSECMIB.Cipsipsecstatistics
    }
    if childYangName == "cipsSysCapacityGroup" {
        return &cISCOIPSECMIB.Cipssyscapacitygroup
    }
    if childYangName == "cipsTrapCntlGroup" {
        return &cISCOIPSECMIB.Cipstrapcntlgroup
    }
    if childYangName == "cipsIsakmpPolicyTable" {
        return &cISCOIPSECMIB.Cipsisakmppolicytable
    }
    if childYangName == "cipsStaticCryptomapSetTable" {
        return &cISCOIPSECMIB.Cipsstaticcryptomapsettable
    }
    if childYangName == "cipsDynamicCryptomapSetTable" {
        return &cISCOIPSECMIB.Cipsdynamiccryptomapsettable
    }
    if childYangName == "cipsStaticCryptomapTable" {
        return &cISCOIPSECMIB.Cipsstaticcryptomaptable
    }
    if childYangName == "cipsCryptomapSetIfTable" {
        return &cISCOIPSECMIB.Cipscryptomapsetiftable
    }
    return nil
}

func (cISCOIPSECMIB *CISCOIPSECMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cipsIsakmpGroup"] = &cISCOIPSECMIB.Cipsisakmpgroup
    children["cipsIPsecGlobals"] = &cISCOIPSECMIB.Cipsipsecglobals
    children["cipsIPsecStatistics"] = &cISCOIPSECMIB.Cipsipsecstatistics
    children["cipsSysCapacityGroup"] = &cISCOIPSECMIB.Cipssyscapacitygroup
    children["cipsTrapCntlGroup"] = &cISCOIPSECMIB.Cipstrapcntlgroup
    children["cipsIsakmpPolicyTable"] = &cISCOIPSECMIB.Cipsisakmppolicytable
    children["cipsStaticCryptomapSetTable"] = &cISCOIPSECMIB.Cipsstaticcryptomapsettable
    children["cipsDynamicCryptomapSetTable"] = &cISCOIPSECMIB.Cipsdynamiccryptomapsettable
    children["cipsStaticCryptomapTable"] = &cISCOIPSECMIB.Cipsstaticcryptomaptable
    children["cipsCryptomapSetIfTable"] = &cISCOIPSECMIB.Cipscryptomapsetiftable
    return children
}

func (cISCOIPSECMIB *CISCOIPSECMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOIPSECMIB *CISCOIPSECMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOIPSECMIB *CISCOIPSECMIB) GetYangName() string { return "CISCO-IPSEC-MIB" }

func (cISCOIPSECMIB *CISCOIPSECMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOIPSECMIB *CISCOIPSECMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOIPSECMIB *CISCOIPSECMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOIPSECMIB *CISCOIPSECMIB) SetParent(parent types.Entity) { cISCOIPSECMIB.parent = parent }

func (cISCOIPSECMIB *CISCOIPSECMIB) GetParent() types.Entity { return cISCOIPSECMIB.parent }

func (cISCOIPSECMIB *CISCOIPSECMIB) GetParentYangName() string { return "CISCO-IPSEC-MIB" }

// CISCOIPSECMIB_Cipsisakmpgroup
type CISCOIPSECMIB_Cipsisakmpgroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The value of this object is TRUE if ISAKMP has been enabled on the managed
    // entity. Otherise the value of this object is FALSE. The type is bool.
    Cipsisakmpenabled interface{}

    // The value of this object is shows the type of identity used by the managed
    // entity in ISAKMP negotiations with another peer. The type is
    // IkeIdentityType.
    Cipsisakmpidentity interface{}

    // The value of this object is time interval in seconds between successive
    // ISAKMP keepalive heartbeats issued to the peers to which IKE tunnels have
    // been setup. The type is interface{} with range: 10..3600. Units are
    // seconds.
    Cipsisakmpkeepaliveinterval interface{}

    // The value of this object is the number of ISAKMP policies that have been
    // configured on the  managed entity. The type is interface{} with range:
    // 0..2147483647.
    Cipsnumisakmppolicies interface{}
}

func (cipsisakmpgroup *CISCOIPSECMIB_Cipsisakmpgroup) GetFilter() yfilter.YFilter { return cipsisakmpgroup.YFilter }

func (cipsisakmpgroup *CISCOIPSECMIB_Cipsisakmpgroup) SetFilter(yf yfilter.YFilter) { cipsisakmpgroup.YFilter = yf }

func (cipsisakmpgroup *CISCOIPSECMIB_Cipsisakmpgroup) GetGoName(yname string) string {
    if yname == "cipsIsakmpEnabled" { return "Cipsisakmpenabled" }
    if yname == "cipsIsakmpIdentity" { return "Cipsisakmpidentity" }
    if yname == "cipsIsakmpKeepaliveInterval" { return "Cipsisakmpkeepaliveinterval" }
    if yname == "cipsNumIsakmpPolicies" { return "Cipsnumisakmppolicies" }
    return ""
}

func (cipsisakmpgroup *CISCOIPSECMIB_Cipsisakmpgroup) GetSegmentPath() string {
    return "cipsIsakmpGroup"
}

func (cipsisakmpgroup *CISCOIPSECMIB_Cipsisakmpgroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipsisakmpgroup *CISCOIPSECMIB_Cipsisakmpgroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipsisakmpgroup *CISCOIPSECMIB_Cipsisakmpgroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipsIsakmpEnabled"] = cipsisakmpgroup.Cipsisakmpenabled
    leafs["cipsIsakmpIdentity"] = cipsisakmpgroup.Cipsisakmpidentity
    leafs["cipsIsakmpKeepaliveInterval"] = cipsisakmpgroup.Cipsisakmpkeepaliveinterval
    leafs["cipsNumIsakmpPolicies"] = cipsisakmpgroup.Cipsnumisakmppolicies
    return leafs
}

func (cipsisakmpgroup *CISCOIPSECMIB_Cipsisakmpgroup) GetBundleName() string { return "cisco_ios_xe" }

func (cipsisakmpgroup *CISCOIPSECMIB_Cipsisakmpgroup) GetYangName() string { return "cipsIsakmpGroup" }

func (cipsisakmpgroup *CISCOIPSECMIB_Cipsisakmpgroup) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsisakmpgroup *CISCOIPSECMIB_Cipsisakmpgroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsisakmpgroup *CISCOIPSECMIB_Cipsisakmpgroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsisakmpgroup *CISCOIPSECMIB_Cipsisakmpgroup) SetParent(parent types.Entity) { cipsisakmpgroup.parent = parent }

func (cipsisakmpgroup *CISCOIPSECMIB_Cipsisakmpgroup) GetParent() types.Entity { return cipsisakmpgroup.parent }

func (cipsisakmpgroup *CISCOIPSECMIB_Cipsisakmpgroup) GetParentYangName() string { return "CISCO-IPSEC-MIB" }

// CISCOIPSECMIB_Cipsipsecglobals
type CISCOIPSECMIB_Cipsipsecglobals struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The default lifetime (in seconds) assigned  to an SA as a global policy
    // (maybe overridden  in specific cryptomap definitions). The type is
    // interface{} with range: 120..86400. Units are Seconds.
    Cipssalifetime interface{}

    // The default lifesize in KBytes assigned to an SA  as a global policy
    // (unless overridden in cryptomap  definition). The type is interface{} with
    // range: 2560..536870912. Units are KBytes.
    Cipssalifesize interface{}

    // The number of Cryptomap Sets that are are fully configured. Statically
    // defined cryptomap sets  are ones where the operator has fully specified all
    // the parameters required set up IPSec  Virtual Private Networks (VPNs). The
    // type is interface{} with range: 0..2147483647. Units are Integral Units.
    Cipsnumstaticcryptomapsets interface{}

    // The number of static Cryptomap Sets that have  at least one CET cryptomap
    // element as a member of the set. The type is interface{} with range:
    // 0..2147483647. Units are Integral Units.
    Cipsnumcetcryptomapsets interface{}

    // The number of dynamic IPSec Policy templates (called 'dynamic cryptomap
    // templates') configured on the managed entity. The type is interface{} with
    // range: 0..2147483647. Units are Integral Units.
    Cipsnumdynamiccryptomapsets interface{}

    // The number of static Cryptomap Sets that have  at least one dynamic
    // cryptomap template  bound to them which has the Tunnel Endpoint Discovery
    // (TED) enabled. The type is interface{} with range: 0..2147483647. Units are
    // Integral Units.
    Cipsnumtedcryptomapsets interface{}
}

func (cipsipsecglobals *CISCOIPSECMIB_Cipsipsecglobals) GetFilter() yfilter.YFilter { return cipsipsecglobals.YFilter }

func (cipsipsecglobals *CISCOIPSECMIB_Cipsipsecglobals) SetFilter(yf yfilter.YFilter) { cipsipsecglobals.YFilter = yf }

func (cipsipsecglobals *CISCOIPSECMIB_Cipsipsecglobals) GetGoName(yname string) string {
    if yname == "cipsSALifetime" { return "Cipssalifetime" }
    if yname == "cipsSALifesize" { return "Cipssalifesize" }
    if yname == "cipsNumStaticCryptomapSets" { return "Cipsnumstaticcryptomapsets" }
    if yname == "cipsNumCETCryptomapSets" { return "Cipsnumcetcryptomapsets" }
    if yname == "cipsNumDynamicCryptomapSets" { return "Cipsnumdynamiccryptomapsets" }
    if yname == "cipsNumTEDCryptomapSets" { return "Cipsnumtedcryptomapsets" }
    return ""
}

func (cipsipsecglobals *CISCOIPSECMIB_Cipsipsecglobals) GetSegmentPath() string {
    return "cipsIPsecGlobals"
}

func (cipsipsecglobals *CISCOIPSECMIB_Cipsipsecglobals) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipsipsecglobals *CISCOIPSECMIB_Cipsipsecglobals) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipsipsecglobals *CISCOIPSECMIB_Cipsipsecglobals) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipsSALifetime"] = cipsipsecglobals.Cipssalifetime
    leafs["cipsSALifesize"] = cipsipsecglobals.Cipssalifesize
    leafs["cipsNumStaticCryptomapSets"] = cipsipsecglobals.Cipsnumstaticcryptomapsets
    leafs["cipsNumCETCryptomapSets"] = cipsipsecglobals.Cipsnumcetcryptomapsets
    leafs["cipsNumDynamicCryptomapSets"] = cipsipsecglobals.Cipsnumdynamiccryptomapsets
    leafs["cipsNumTEDCryptomapSets"] = cipsipsecglobals.Cipsnumtedcryptomapsets
    return leafs
}

func (cipsipsecglobals *CISCOIPSECMIB_Cipsipsecglobals) GetBundleName() string { return "cisco_ios_xe" }

func (cipsipsecglobals *CISCOIPSECMIB_Cipsipsecglobals) GetYangName() string { return "cipsIPsecGlobals" }

func (cipsipsecglobals *CISCOIPSECMIB_Cipsipsecglobals) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsipsecglobals *CISCOIPSECMIB_Cipsipsecglobals) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsipsecglobals *CISCOIPSECMIB_Cipsipsecglobals) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsipsecglobals *CISCOIPSECMIB_Cipsipsecglobals) SetParent(parent types.Entity) { cipsipsecglobals.parent = parent }

func (cipsipsecglobals *CISCOIPSECMIB_Cipsipsecglobals) GetParent() types.Entity { return cipsipsecglobals.parent }

func (cipsipsecglobals *CISCOIPSECMIB_Cipsipsecglobals) GetParentYangName() string { return "CISCO-IPSEC-MIB" }

// CISCOIPSECMIB_Cipsipsecstatistics
type CISCOIPSECMIB_Cipsipsecstatistics struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The number of TED probes that were received by this  managed entity since
    // bootup. Not affected by any  CLI operation. The type is interface{} with
    // range: 0..4294967295. Units are Integral Units.
    Cipsnumtedprobesreceived interface{}

    // The number of TED probes that were dispatched by all the dynamic cryptomaps
    // in this managed entity since  bootup. Not affected by any CLI operation.
    // The type is interface{} with range: 0..4294967295. Units are Integral
    // Units.
    Cipsnumtedprobessent interface{}

    // The number of TED probes that were dispatched by  the local entity and that
    // failed to locate crypto  endpoint.  Not affected by any CLI operation. The
    // type is interface{} with range: 0..4294967295. Units are Integral Units.
    Cipsnumtedfailures interface{}
}

func (cipsipsecstatistics *CISCOIPSECMIB_Cipsipsecstatistics) GetFilter() yfilter.YFilter { return cipsipsecstatistics.YFilter }

func (cipsipsecstatistics *CISCOIPSECMIB_Cipsipsecstatistics) SetFilter(yf yfilter.YFilter) { cipsipsecstatistics.YFilter = yf }

func (cipsipsecstatistics *CISCOIPSECMIB_Cipsipsecstatistics) GetGoName(yname string) string {
    if yname == "cipsNumTEDProbesReceived" { return "Cipsnumtedprobesreceived" }
    if yname == "cipsNumTEDProbesSent" { return "Cipsnumtedprobessent" }
    if yname == "cipsNumTEDFailures" { return "Cipsnumtedfailures" }
    return ""
}

func (cipsipsecstatistics *CISCOIPSECMIB_Cipsipsecstatistics) GetSegmentPath() string {
    return "cipsIPsecStatistics"
}

func (cipsipsecstatistics *CISCOIPSECMIB_Cipsipsecstatistics) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipsipsecstatistics *CISCOIPSECMIB_Cipsipsecstatistics) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipsipsecstatistics *CISCOIPSECMIB_Cipsipsecstatistics) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipsNumTEDProbesReceived"] = cipsipsecstatistics.Cipsnumtedprobesreceived
    leafs["cipsNumTEDProbesSent"] = cipsipsecstatistics.Cipsnumtedprobessent
    leafs["cipsNumTEDFailures"] = cipsipsecstatistics.Cipsnumtedfailures
    return leafs
}

func (cipsipsecstatistics *CISCOIPSECMIB_Cipsipsecstatistics) GetBundleName() string { return "cisco_ios_xe" }

func (cipsipsecstatistics *CISCOIPSECMIB_Cipsipsecstatistics) GetYangName() string { return "cipsIPsecStatistics" }

func (cipsipsecstatistics *CISCOIPSECMIB_Cipsipsecstatistics) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsipsecstatistics *CISCOIPSECMIB_Cipsipsecstatistics) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsipsecstatistics *CISCOIPSECMIB_Cipsipsecstatistics) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsipsecstatistics *CISCOIPSECMIB_Cipsipsecstatistics) SetParent(parent types.Entity) { cipsipsecstatistics.parent = parent }

func (cipsipsecstatistics *CISCOIPSECMIB_Cipsipsecstatistics) GetParent() types.Entity { return cipsipsecstatistics.parent }

func (cipsipsecstatistics *CISCOIPSECMIB_Cipsipsecstatistics) GetParentYangName() string { return "CISCO-IPSEC-MIB" }

// CISCOIPSECMIB_Cipssyscapacitygroup
type CISCOIPSECMIB_Cipssyscapacitygroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The maximum number of IPsec Security Associations that can be established
    // on this managed entity. If no theoretical limit exists, this returns value
    // 0.  Not affected by any CLI operation. The type is interface{} with range:
    // 0..65535. Units are Integral Units.
    Cipsmaxsas interface{}

    // The value of this object is TRUE if the  managed entity has the hardware
    // nad software  features to support 3DES encryption algorithm.  Not affected
    // by any CLI operation. The type is bool.
    Cips3Descapable interface{}
}

func (cipssyscapacitygroup *CISCOIPSECMIB_Cipssyscapacitygroup) GetFilter() yfilter.YFilter { return cipssyscapacitygroup.YFilter }

func (cipssyscapacitygroup *CISCOIPSECMIB_Cipssyscapacitygroup) SetFilter(yf yfilter.YFilter) { cipssyscapacitygroup.YFilter = yf }

func (cipssyscapacitygroup *CISCOIPSECMIB_Cipssyscapacitygroup) GetGoName(yname string) string {
    if yname == "cipsMaxSAs" { return "Cipsmaxsas" }
    if yname == "cips3DesCapable" { return "Cips3Descapable" }
    return ""
}

func (cipssyscapacitygroup *CISCOIPSECMIB_Cipssyscapacitygroup) GetSegmentPath() string {
    return "cipsSysCapacityGroup"
}

func (cipssyscapacitygroup *CISCOIPSECMIB_Cipssyscapacitygroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipssyscapacitygroup *CISCOIPSECMIB_Cipssyscapacitygroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipssyscapacitygroup *CISCOIPSECMIB_Cipssyscapacitygroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipsMaxSAs"] = cipssyscapacitygroup.Cipsmaxsas
    leafs["cips3DesCapable"] = cipssyscapacitygroup.Cips3Descapable
    return leafs
}

func (cipssyscapacitygroup *CISCOIPSECMIB_Cipssyscapacitygroup) GetBundleName() string { return "cisco_ios_xe" }

func (cipssyscapacitygroup *CISCOIPSECMIB_Cipssyscapacitygroup) GetYangName() string { return "cipsSysCapacityGroup" }

func (cipssyscapacitygroup *CISCOIPSECMIB_Cipssyscapacitygroup) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipssyscapacitygroup *CISCOIPSECMIB_Cipssyscapacitygroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipssyscapacitygroup *CISCOIPSECMIB_Cipssyscapacitygroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipssyscapacitygroup *CISCOIPSECMIB_Cipssyscapacitygroup) SetParent(parent types.Entity) { cipssyscapacitygroup.parent = parent }

func (cipssyscapacitygroup *CISCOIPSECMIB_Cipssyscapacitygroup) GetParent() types.Entity { return cipssyscapacitygroup.parent }

func (cipssyscapacitygroup *CISCOIPSECMIB_Cipssyscapacitygroup) GetParentYangName() string { return "CISCO-IPSEC-MIB" }

// CISCOIPSECMIB_Cipstrapcntlgroup
type CISCOIPSECMIB_Cipstrapcntlgroup struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This object defines the administrative state of  sending the IOS IPsec
    // ISAKMP Policy Add trap. The type is TrapStatus.
    Cipscntlisakmppolicyadded interface{}

    // This object defines the administrative state of  sending the IOS IPsec
    // ISAKMP Policy Delete trap. The type is TrapStatus.
    Cipscntlisakmppolicydeleted interface{}

    // This object defines the administrative state of  sending the IOS IPsec
    // Cryptomap Add trap. The type is TrapStatus.
    Cipscntlcryptomapadded interface{}

    // This object defines the administrative state of  sending the IOS IPsec
    // Cryptomap Delete trap. The type is TrapStatus.
    Cipscntlcryptomapdeleted interface{}

    // This object defines the administrative state of  sending the IOS IPsec trap
    // that is issued when a cryptomap set is attached to an interface. The type
    // is TrapStatus.
    Cipscntlcryptomapsetattached interface{}

    // This object defines the administrative state of  sending the IOS IPsec trap
    // that is issued when a cryptomap set is detached from an interface. to which
    // it was earlier bound. The type is TrapStatus.
    Cipscntlcryptomapsetdetached interface{}

    // This object defines the administrative state of  sending the IOS IPsec trap
    // that is issued when the number of SAs crosses the maximum number of SAs
    // that may be supported on the managed entity. The type is TrapStatus.
    Cipscntltoomanysas interface{}
}

func (cipstrapcntlgroup *CISCOIPSECMIB_Cipstrapcntlgroup) GetFilter() yfilter.YFilter { return cipstrapcntlgroup.YFilter }

func (cipstrapcntlgroup *CISCOIPSECMIB_Cipstrapcntlgroup) SetFilter(yf yfilter.YFilter) { cipstrapcntlgroup.YFilter = yf }

func (cipstrapcntlgroup *CISCOIPSECMIB_Cipstrapcntlgroup) GetGoName(yname string) string {
    if yname == "cipsCntlIsakmpPolicyAdded" { return "Cipscntlisakmppolicyadded" }
    if yname == "cipsCntlIsakmpPolicyDeleted" { return "Cipscntlisakmppolicydeleted" }
    if yname == "cipsCntlCryptomapAdded" { return "Cipscntlcryptomapadded" }
    if yname == "cipsCntlCryptomapDeleted" { return "Cipscntlcryptomapdeleted" }
    if yname == "cipsCntlCryptomapSetAttached" { return "Cipscntlcryptomapsetattached" }
    if yname == "cipsCntlCryptomapSetDetached" { return "Cipscntlcryptomapsetdetached" }
    if yname == "cipsCntlTooManySAs" { return "Cipscntltoomanysas" }
    return ""
}

func (cipstrapcntlgroup *CISCOIPSECMIB_Cipstrapcntlgroup) GetSegmentPath() string {
    return "cipsTrapCntlGroup"
}

func (cipstrapcntlgroup *CISCOIPSECMIB_Cipstrapcntlgroup) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipstrapcntlgroup *CISCOIPSECMIB_Cipstrapcntlgroup) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipstrapcntlgroup *CISCOIPSECMIB_Cipstrapcntlgroup) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipsCntlIsakmpPolicyAdded"] = cipstrapcntlgroup.Cipscntlisakmppolicyadded
    leafs["cipsCntlIsakmpPolicyDeleted"] = cipstrapcntlgroup.Cipscntlisakmppolicydeleted
    leafs["cipsCntlCryptomapAdded"] = cipstrapcntlgroup.Cipscntlcryptomapadded
    leafs["cipsCntlCryptomapDeleted"] = cipstrapcntlgroup.Cipscntlcryptomapdeleted
    leafs["cipsCntlCryptomapSetAttached"] = cipstrapcntlgroup.Cipscntlcryptomapsetattached
    leafs["cipsCntlCryptomapSetDetached"] = cipstrapcntlgroup.Cipscntlcryptomapsetdetached
    leafs["cipsCntlTooManySAs"] = cipstrapcntlgroup.Cipscntltoomanysas
    return leafs
}

func (cipstrapcntlgroup *CISCOIPSECMIB_Cipstrapcntlgroup) GetBundleName() string { return "cisco_ios_xe" }

func (cipstrapcntlgroup *CISCOIPSECMIB_Cipstrapcntlgroup) GetYangName() string { return "cipsTrapCntlGroup" }

func (cipstrapcntlgroup *CISCOIPSECMIB_Cipstrapcntlgroup) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipstrapcntlgroup *CISCOIPSECMIB_Cipstrapcntlgroup) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipstrapcntlgroup *CISCOIPSECMIB_Cipstrapcntlgroup) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipstrapcntlgroup *CISCOIPSECMIB_Cipstrapcntlgroup) SetParent(parent types.Entity) { cipstrapcntlgroup.parent = parent }

func (cipstrapcntlgroup *CISCOIPSECMIB_Cipstrapcntlgroup) GetParent() types.Entity { return cipstrapcntlgroup.parent }

func (cipstrapcntlgroup *CISCOIPSECMIB_Cipstrapcntlgroup) GetParentYangName() string { return "CISCO-IPSEC-MIB" }

// CISCOIPSECMIB_Cipsisakmppolicytable
// The table containing the list of all
// ISAKMP policy entries configured by the operator.
type CISCOIPSECMIB_Cipsisakmppolicytable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains the attributes  associated with a single ISAKMP Policy
    // entry. The type is slice of
    // CISCOIPSECMIB_Cipsisakmppolicytable_Cipsisakmppolicyentry.
    Cipsisakmppolicyentry []CISCOIPSECMIB_Cipsisakmppolicytable_Cipsisakmppolicyentry
}

func (cipsisakmppolicytable *CISCOIPSECMIB_Cipsisakmppolicytable) GetFilter() yfilter.YFilter { return cipsisakmppolicytable.YFilter }

func (cipsisakmppolicytable *CISCOIPSECMIB_Cipsisakmppolicytable) SetFilter(yf yfilter.YFilter) { cipsisakmppolicytable.YFilter = yf }

func (cipsisakmppolicytable *CISCOIPSECMIB_Cipsisakmppolicytable) GetGoName(yname string) string {
    if yname == "cipsIsakmpPolicyEntry" { return "Cipsisakmppolicyentry" }
    return ""
}

func (cipsisakmppolicytable *CISCOIPSECMIB_Cipsisakmppolicytable) GetSegmentPath() string {
    return "cipsIsakmpPolicyTable"
}

func (cipsisakmppolicytable *CISCOIPSECMIB_Cipsisakmppolicytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipsIsakmpPolicyEntry" {
        for _, c := range cipsisakmppolicytable.Cipsisakmppolicyentry {
            if cipsisakmppolicytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSECMIB_Cipsisakmppolicytable_Cipsisakmppolicyentry{}
        cipsisakmppolicytable.Cipsisakmppolicyentry = append(cipsisakmppolicytable.Cipsisakmppolicyentry, child)
        return &cipsisakmppolicytable.Cipsisakmppolicyentry[len(cipsisakmppolicytable.Cipsisakmppolicyentry)-1]
    }
    return nil
}

func (cipsisakmppolicytable *CISCOIPSECMIB_Cipsisakmppolicytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cipsisakmppolicytable.Cipsisakmppolicyentry {
        children[cipsisakmppolicytable.Cipsisakmppolicyentry[i].GetSegmentPath()] = &cipsisakmppolicytable.Cipsisakmppolicyentry[i]
    }
    return children
}

func (cipsisakmppolicytable *CISCOIPSECMIB_Cipsisakmppolicytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cipsisakmppolicytable *CISCOIPSECMIB_Cipsisakmppolicytable) GetBundleName() string { return "cisco_ios_xe" }

func (cipsisakmppolicytable *CISCOIPSECMIB_Cipsisakmppolicytable) GetYangName() string { return "cipsIsakmpPolicyTable" }

func (cipsisakmppolicytable *CISCOIPSECMIB_Cipsisakmppolicytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsisakmppolicytable *CISCOIPSECMIB_Cipsisakmppolicytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsisakmppolicytable *CISCOIPSECMIB_Cipsisakmppolicytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsisakmppolicytable *CISCOIPSECMIB_Cipsisakmppolicytable) SetParent(parent types.Entity) { cipsisakmppolicytable.parent = parent }

func (cipsisakmppolicytable *CISCOIPSECMIB_Cipsisakmppolicytable) GetParent() types.Entity { return cipsisakmppolicytable.parent }

func (cipsisakmppolicytable *CISCOIPSECMIB_Cipsisakmppolicytable) GetParentYangName() string { return "CISCO-IPSEC-MIB" }

// CISCOIPSECMIB_Cipsisakmppolicytable_Cipsisakmppolicyentry
// Each entry contains the attributes 
// associated with a single ISAKMP
// Policy entry.
type CISCOIPSECMIB_Cipsisakmppolicytable_Cipsisakmppolicyentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The priotity of this ISAKMP Policy entry. This is
    // also the index of this table. The type is interface{} with range: 0..65535.
    Cipsisakmppolpriority interface{}

    // The encryption transform specified by this  ISAKMP policy specification.
    // The Internet Key Exchange (IKE) tunnels setup using this policy item would
    // use the specified encryption transform to protect the ISAKMP PDUs. The type
    // is EncryptAlgo.
    Cipsisakmppolencr interface{}

    // The hash transform specified by this  ISAKMP policy specification. The IKE
    // tunnels setup using this policy item would use the  specified hash
    // transform to protect the ISAKMP PDUs. The type is IkeHashAlgo.
    Cipsisakmppolhash interface{}

    // The peer authentication mthod specified by this ISAKMP policy
    // specification. If this policy entity is selected for negotiation with a
    // peer, the local entity would authenticate the peer using  the method
    // specified by this object. The type is IkeAuthMethod.
    Cipsisakmppolauth interface{}

    // This object specifies the Oakley group used  for Diffie Hellman exchange in
    // the Main Mode.  If this policy item is selected to negotiate Main Mode with
    // an IKE peer, the local entity  chooses the group specified by this object
    // to perform Diffie Hellman exchange with the peer. The type is
    // DiffHellmanGrp.
    Cipsisakmppolgroup interface{}

    // This object specifies the lifetime in seconds of the IKE tunnels generated
    // using this  policy specification. The type is interface{} with range:
    // 60..86400. Units are seconds.
    Cipsisakmppollifetime interface{}
}

func (cipsisakmppolicyentry *CISCOIPSECMIB_Cipsisakmppolicytable_Cipsisakmppolicyentry) GetFilter() yfilter.YFilter { return cipsisakmppolicyentry.YFilter }

func (cipsisakmppolicyentry *CISCOIPSECMIB_Cipsisakmppolicytable_Cipsisakmppolicyentry) SetFilter(yf yfilter.YFilter) { cipsisakmppolicyentry.YFilter = yf }

func (cipsisakmppolicyentry *CISCOIPSECMIB_Cipsisakmppolicytable_Cipsisakmppolicyentry) GetGoName(yname string) string {
    if yname == "cipsIsakmpPolPriority" { return "Cipsisakmppolpriority" }
    if yname == "cipsIsakmpPolEncr" { return "Cipsisakmppolencr" }
    if yname == "cipsIsakmpPolHash" { return "Cipsisakmppolhash" }
    if yname == "cipsIsakmpPolAuth" { return "Cipsisakmppolauth" }
    if yname == "cipsIsakmpPolGroup" { return "Cipsisakmppolgroup" }
    if yname == "cipsIsakmpPolLifetime" { return "Cipsisakmppollifetime" }
    return ""
}

func (cipsisakmppolicyentry *CISCOIPSECMIB_Cipsisakmppolicytable_Cipsisakmppolicyentry) GetSegmentPath() string {
    return "cipsIsakmpPolicyEntry" + "[cipsIsakmpPolPriority='" + fmt.Sprintf("%v", cipsisakmppolicyentry.Cipsisakmppolpriority) + "']"
}

func (cipsisakmppolicyentry *CISCOIPSECMIB_Cipsisakmppolicytable_Cipsisakmppolicyentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipsisakmppolicyentry *CISCOIPSECMIB_Cipsisakmppolicytable_Cipsisakmppolicyentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipsisakmppolicyentry *CISCOIPSECMIB_Cipsisakmppolicytable_Cipsisakmppolicyentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipsIsakmpPolPriority"] = cipsisakmppolicyentry.Cipsisakmppolpriority
    leafs["cipsIsakmpPolEncr"] = cipsisakmppolicyentry.Cipsisakmppolencr
    leafs["cipsIsakmpPolHash"] = cipsisakmppolicyentry.Cipsisakmppolhash
    leafs["cipsIsakmpPolAuth"] = cipsisakmppolicyentry.Cipsisakmppolauth
    leafs["cipsIsakmpPolGroup"] = cipsisakmppolicyentry.Cipsisakmppolgroup
    leafs["cipsIsakmpPolLifetime"] = cipsisakmppolicyentry.Cipsisakmppollifetime
    return leafs
}

func (cipsisakmppolicyentry *CISCOIPSECMIB_Cipsisakmppolicytable_Cipsisakmppolicyentry) GetBundleName() string { return "cisco_ios_xe" }

func (cipsisakmppolicyentry *CISCOIPSECMIB_Cipsisakmppolicytable_Cipsisakmppolicyentry) GetYangName() string { return "cipsIsakmpPolicyEntry" }

func (cipsisakmppolicyentry *CISCOIPSECMIB_Cipsisakmppolicytable_Cipsisakmppolicyentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsisakmppolicyentry *CISCOIPSECMIB_Cipsisakmppolicytable_Cipsisakmppolicyentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsisakmppolicyentry *CISCOIPSECMIB_Cipsisakmppolicytable_Cipsisakmppolicyentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsisakmppolicyentry *CISCOIPSECMIB_Cipsisakmppolicytable_Cipsisakmppolicyentry) SetParent(parent types.Entity) { cipsisakmppolicyentry.parent = parent }

func (cipsisakmppolicyentry *CISCOIPSECMIB_Cipsisakmppolicytable_Cipsisakmppolicyentry) GetParent() types.Entity { return cipsisakmppolicyentry.parent }

func (cipsisakmppolicyentry *CISCOIPSECMIB_Cipsisakmppolicytable_Cipsisakmppolicyentry) GetParentYangName() string { return "cipsIsakmpPolicyTable" }

// CISCOIPSECMIB_Cipsstaticcryptomapsettable
// The table containing the list of all
// cryptomap sets that are fully specified
// and are not wild-carded.
// 
// The operator may include different types of
// cryptomaps in such a set - manual, CET,
// ISAKMP or dynamic.
type CISCOIPSECMIB_Cipsstaticcryptomapsettable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains the attributes  associated with a single static 
    // cryptomap set. The type is slice of
    // CISCOIPSECMIB_Cipsstaticcryptomapsettable_Cipsstaticcryptomapsetentry.
    Cipsstaticcryptomapsetentry []CISCOIPSECMIB_Cipsstaticcryptomapsettable_Cipsstaticcryptomapsetentry
}

func (cipsstaticcryptomapsettable *CISCOIPSECMIB_Cipsstaticcryptomapsettable) GetFilter() yfilter.YFilter { return cipsstaticcryptomapsettable.YFilter }

func (cipsstaticcryptomapsettable *CISCOIPSECMIB_Cipsstaticcryptomapsettable) SetFilter(yf yfilter.YFilter) { cipsstaticcryptomapsettable.YFilter = yf }

func (cipsstaticcryptomapsettable *CISCOIPSECMIB_Cipsstaticcryptomapsettable) GetGoName(yname string) string {
    if yname == "cipsStaticCryptomapSetEntry" { return "Cipsstaticcryptomapsetentry" }
    return ""
}

func (cipsstaticcryptomapsettable *CISCOIPSECMIB_Cipsstaticcryptomapsettable) GetSegmentPath() string {
    return "cipsStaticCryptomapSetTable"
}

func (cipsstaticcryptomapsettable *CISCOIPSECMIB_Cipsstaticcryptomapsettable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipsStaticCryptomapSetEntry" {
        for _, c := range cipsstaticcryptomapsettable.Cipsstaticcryptomapsetentry {
            if cipsstaticcryptomapsettable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSECMIB_Cipsstaticcryptomapsettable_Cipsstaticcryptomapsetentry{}
        cipsstaticcryptomapsettable.Cipsstaticcryptomapsetentry = append(cipsstaticcryptomapsettable.Cipsstaticcryptomapsetentry, child)
        return &cipsstaticcryptomapsettable.Cipsstaticcryptomapsetentry[len(cipsstaticcryptomapsettable.Cipsstaticcryptomapsetentry)-1]
    }
    return nil
}

func (cipsstaticcryptomapsettable *CISCOIPSECMIB_Cipsstaticcryptomapsettable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cipsstaticcryptomapsettable.Cipsstaticcryptomapsetentry {
        children[cipsstaticcryptomapsettable.Cipsstaticcryptomapsetentry[i].GetSegmentPath()] = &cipsstaticcryptomapsettable.Cipsstaticcryptomapsetentry[i]
    }
    return children
}

func (cipsstaticcryptomapsettable *CISCOIPSECMIB_Cipsstaticcryptomapsettable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cipsstaticcryptomapsettable *CISCOIPSECMIB_Cipsstaticcryptomapsettable) GetBundleName() string { return "cisco_ios_xe" }

func (cipsstaticcryptomapsettable *CISCOIPSECMIB_Cipsstaticcryptomapsettable) GetYangName() string { return "cipsStaticCryptomapSetTable" }

func (cipsstaticcryptomapsettable *CISCOIPSECMIB_Cipsstaticcryptomapsettable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsstaticcryptomapsettable *CISCOIPSECMIB_Cipsstaticcryptomapsettable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsstaticcryptomapsettable *CISCOIPSECMIB_Cipsstaticcryptomapsettable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsstaticcryptomapsettable *CISCOIPSECMIB_Cipsstaticcryptomapsettable) SetParent(parent types.Entity) { cipsstaticcryptomapsettable.parent = parent }

func (cipsstaticcryptomapsettable *CISCOIPSECMIB_Cipsstaticcryptomapsettable) GetParent() types.Entity { return cipsstaticcryptomapsettable.parent }

func (cipsstaticcryptomapsettable *CISCOIPSECMIB_Cipsstaticcryptomapsettable) GetParentYangName() string { return "CISCO-IPSEC-MIB" }

// CISCOIPSECMIB_Cipsstaticcryptomapsettable_Cipsstaticcryptomapsetentry
// Each entry contains the attributes 
// associated with a single static 
// cryptomap set.
type CISCOIPSECMIB_Cipsstaticcryptomapsettable_Cipsstaticcryptomapsetentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The index of the static cryptomap table. The value
    // of the string is the name string assigned by the  operator in defining the
    // cryptomap set. The type is string.
    Cipsstaticcryptomapsetname interface{}

    // The total number of cryptomap entries contained in this cryptomap set. .
    // The type is interface{} with range: 0..4294967295.
    Cipsstaticcryptomapsetsize interface{}

    // The number of cryptomaps associated with this  cryptomap set that use
    // ISAKMP protocol to do key exchange. The type is interface{} with range:
    // 0..4294967295.
    Cipsstaticcryptomapsetnumisakmp interface{}

    // The number of cryptomaps associated with this  cryptomap set that require
    // the operator to manually setup the keys and SPIs. The type is interface{}
    // with range: 0..4294967295.
    Cipsstaticcryptomapsetnummanual interface{}

    // The number of cryptomaps of type 'ipsec-cisco'  associated with this
    // cryptomap set. Such cryptomap elements implement Cisco Encryption
    // Technology based Virtual Private Networks. The type is interface{} with
    // range: 0..4294967295.
    Cipsstaticcryptomapsetnumcet interface{}

    // The number of dynamic cryptomap templates linked to this cryptomap set. The
    // type is interface{} with range: 0..4294967295.
    Cipsstaticcryptomapsetnumdynamic interface{}

    // The number of dynamic cryptomap templates linked to this cryptomap set that
    // have Tunnel Endpoint Discovery (TED) enabled. The type is interface{} with
    // range: 0..4294967295.
    Cipsstaticcryptomapsetnumdisc interface{}

    // The number of and IPsec Security Associations that are active and were
    // setup using this cryptomap.  . The type is interface{} with range:
    // 0..4294967295.
    Cipsstaticcryptomapsetnumsas interface{}
}

func (cipsstaticcryptomapsetentry *CISCOIPSECMIB_Cipsstaticcryptomapsettable_Cipsstaticcryptomapsetentry) GetFilter() yfilter.YFilter { return cipsstaticcryptomapsetentry.YFilter }

func (cipsstaticcryptomapsetentry *CISCOIPSECMIB_Cipsstaticcryptomapsettable_Cipsstaticcryptomapsetentry) SetFilter(yf yfilter.YFilter) { cipsstaticcryptomapsetentry.YFilter = yf }

func (cipsstaticcryptomapsetentry *CISCOIPSECMIB_Cipsstaticcryptomapsettable_Cipsstaticcryptomapsetentry) GetGoName(yname string) string {
    if yname == "cipsStaticCryptomapSetName" { return "Cipsstaticcryptomapsetname" }
    if yname == "cipsStaticCryptomapSetSize" { return "Cipsstaticcryptomapsetsize" }
    if yname == "cipsStaticCryptomapSetNumIsakmp" { return "Cipsstaticcryptomapsetnumisakmp" }
    if yname == "cipsStaticCryptomapSetNumManual" { return "Cipsstaticcryptomapsetnummanual" }
    if yname == "cipsStaticCryptomapSetNumCET" { return "Cipsstaticcryptomapsetnumcet" }
    if yname == "cipsStaticCryptomapSetNumDynamic" { return "Cipsstaticcryptomapsetnumdynamic" }
    if yname == "cipsStaticCryptomapSetNumDisc" { return "Cipsstaticcryptomapsetnumdisc" }
    if yname == "cipsStaticCryptomapSetNumSAs" { return "Cipsstaticcryptomapsetnumsas" }
    return ""
}

func (cipsstaticcryptomapsetentry *CISCOIPSECMIB_Cipsstaticcryptomapsettable_Cipsstaticcryptomapsetentry) GetSegmentPath() string {
    return "cipsStaticCryptomapSetEntry" + "[cipsStaticCryptomapSetName='" + fmt.Sprintf("%v", cipsstaticcryptomapsetentry.Cipsstaticcryptomapsetname) + "']"
}

func (cipsstaticcryptomapsetentry *CISCOIPSECMIB_Cipsstaticcryptomapsettable_Cipsstaticcryptomapsetentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipsstaticcryptomapsetentry *CISCOIPSECMIB_Cipsstaticcryptomapsettable_Cipsstaticcryptomapsetentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipsstaticcryptomapsetentry *CISCOIPSECMIB_Cipsstaticcryptomapsettable_Cipsstaticcryptomapsetentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipsStaticCryptomapSetName"] = cipsstaticcryptomapsetentry.Cipsstaticcryptomapsetname
    leafs["cipsStaticCryptomapSetSize"] = cipsstaticcryptomapsetentry.Cipsstaticcryptomapsetsize
    leafs["cipsStaticCryptomapSetNumIsakmp"] = cipsstaticcryptomapsetentry.Cipsstaticcryptomapsetnumisakmp
    leafs["cipsStaticCryptomapSetNumManual"] = cipsstaticcryptomapsetentry.Cipsstaticcryptomapsetnummanual
    leafs["cipsStaticCryptomapSetNumCET"] = cipsstaticcryptomapsetentry.Cipsstaticcryptomapsetnumcet
    leafs["cipsStaticCryptomapSetNumDynamic"] = cipsstaticcryptomapsetentry.Cipsstaticcryptomapsetnumdynamic
    leafs["cipsStaticCryptomapSetNumDisc"] = cipsstaticcryptomapsetentry.Cipsstaticcryptomapsetnumdisc
    leafs["cipsStaticCryptomapSetNumSAs"] = cipsstaticcryptomapsetentry.Cipsstaticcryptomapsetnumsas
    return leafs
}

func (cipsstaticcryptomapsetentry *CISCOIPSECMIB_Cipsstaticcryptomapsettable_Cipsstaticcryptomapsetentry) GetBundleName() string { return "cisco_ios_xe" }

func (cipsstaticcryptomapsetentry *CISCOIPSECMIB_Cipsstaticcryptomapsettable_Cipsstaticcryptomapsetentry) GetYangName() string { return "cipsStaticCryptomapSetEntry" }

func (cipsstaticcryptomapsetentry *CISCOIPSECMIB_Cipsstaticcryptomapsettable_Cipsstaticcryptomapsetentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsstaticcryptomapsetentry *CISCOIPSECMIB_Cipsstaticcryptomapsettable_Cipsstaticcryptomapsetentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsstaticcryptomapsetentry *CISCOIPSECMIB_Cipsstaticcryptomapsettable_Cipsstaticcryptomapsetentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsstaticcryptomapsetentry *CISCOIPSECMIB_Cipsstaticcryptomapsettable_Cipsstaticcryptomapsetentry) SetParent(parent types.Entity) { cipsstaticcryptomapsetentry.parent = parent }

func (cipsstaticcryptomapsetentry *CISCOIPSECMIB_Cipsstaticcryptomapsettable_Cipsstaticcryptomapsetentry) GetParent() types.Entity { return cipsstaticcryptomapsetentry.parent }

func (cipsstaticcryptomapsetentry *CISCOIPSECMIB_Cipsstaticcryptomapsettable_Cipsstaticcryptomapsetentry) GetParentYangName() string { return "cipsStaticCryptomapSetTable" }

// CISCOIPSECMIB_Cipsdynamiccryptomapsettable
// The table containing the list of all dynamic
// cryptomaps that use IKE, defined on 
//  the managed entity.
type CISCOIPSECMIB_Cipsdynamiccryptomapsettable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with a single dynamic
    // cryptomap template. The type is slice of
    // CISCOIPSECMIB_Cipsdynamiccryptomapsettable_Cipsdynamiccryptomapsetentry.
    Cipsdynamiccryptomapsetentry []CISCOIPSECMIB_Cipsdynamiccryptomapsettable_Cipsdynamiccryptomapsetentry
}

func (cipsdynamiccryptomapsettable *CISCOIPSECMIB_Cipsdynamiccryptomapsettable) GetFilter() yfilter.YFilter { return cipsdynamiccryptomapsettable.YFilter }

func (cipsdynamiccryptomapsettable *CISCOIPSECMIB_Cipsdynamiccryptomapsettable) SetFilter(yf yfilter.YFilter) { cipsdynamiccryptomapsettable.YFilter = yf }

func (cipsdynamiccryptomapsettable *CISCOIPSECMIB_Cipsdynamiccryptomapsettable) GetGoName(yname string) string {
    if yname == "cipsDynamicCryptomapSetEntry" { return "Cipsdynamiccryptomapsetentry" }
    return ""
}

func (cipsdynamiccryptomapsettable *CISCOIPSECMIB_Cipsdynamiccryptomapsettable) GetSegmentPath() string {
    return "cipsDynamicCryptomapSetTable"
}

func (cipsdynamiccryptomapsettable *CISCOIPSECMIB_Cipsdynamiccryptomapsettable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipsDynamicCryptomapSetEntry" {
        for _, c := range cipsdynamiccryptomapsettable.Cipsdynamiccryptomapsetentry {
            if cipsdynamiccryptomapsettable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSECMIB_Cipsdynamiccryptomapsettable_Cipsdynamiccryptomapsetentry{}
        cipsdynamiccryptomapsettable.Cipsdynamiccryptomapsetentry = append(cipsdynamiccryptomapsettable.Cipsdynamiccryptomapsetentry, child)
        return &cipsdynamiccryptomapsettable.Cipsdynamiccryptomapsetentry[len(cipsdynamiccryptomapsettable.Cipsdynamiccryptomapsetentry)-1]
    }
    return nil
}

func (cipsdynamiccryptomapsettable *CISCOIPSECMIB_Cipsdynamiccryptomapsettable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cipsdynamiccryptomapsettable.Cipsdynamiccryptomapsetentry {
        children[cipsdynamiccryptomapsettable.Cipsdynamiccryptomapsetentry[i].GetSegmentPath()] = &cipsdynamiccryptomapsettable.Cipsdynamiccryptomapsetentry[i]
    }
    return children
}

func (cipsdynamiccryptomapsettable *CISCOIPSECMIB_Cipsdynamiccryptomapsettable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cipsdynamiccryptomapsettable *CISCOIPSECMIB_Cipsdynamiccryptomapsettable) GetBundleName() string { return "cisco_ios_xe" }

func (cipsdynamiccryptomapsettable *CISCOIPSECMIB_Cipsdynamiccryptomapsettable) GetYangName() string { return "cipsDynamicCryptomapSetTable" }

func (cipsdynamiccryptomapsettable *CISCOIPSECMIB_Cipsdynamiccryptomapsettable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsdynamiccryptomapsettable *CISCOIPSECMIB_Cipsdynamiccryptomapsettable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsdynamiccryptomapsettable *CISCOIPSECMIB_Cipsdynamiccryptomapsettable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsdynamiccryptomapsettable *CISCOIPSECMIB_Cipsdynamiccryptomapsettable) SetParent(parent types.Entity) { cipsdynamiccryptomapsettable.parent = parent }

func (cipsdynamiccryptomapsettable *CISCOIPSECMIB_Cipsdynamiccryptomapsettable) GetParent() types.Entity { return cipsdynamiccryptomapsettable.parent }

func (cipsdynamiccryptomapsettable *CISCOIPSECMIB_Cipsdynamiccryptomapsettable) GetParentYangName() string { return "CISCO-IPSEC-MIB" }

// CISCOIPSECMIB_Cipsdynamiccryptomapsettable_Cipsdynamiccryptomapsetentry
// Each entry contains the attributes associated
// with a single dynamic cryptomap template.
type CISCOIPSECMIB_Cipsdynamiccryptomapsettable_Cipsdynamiccryptomapsetentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The index of the dynamic cryptomap table.  The
    // value of the string is the one assigned  by the operator in defining the
    // cryptomap set. The type is string.
    Cipsdynamiccryptomapsetname interface{}

    // The number of cryptomap entries in this cryptomap. The type is interface{}
    // with range: 0..4294967295.
    Cipsdynamiccryptomapsetsize interface{}

    // The number of static cryptomap sets with which this dynamic cryptomap is
    // associated.  . The type is interface{} with range: 0..4294967295.
    Cipsdynamiccryptomapsetnumassoc interface{}
}

func (cipsdynamiccryptomapsetentry *CISCOIPSECMIB_Cipsdynamiccryptomapsettable_Cipsdynamiccryptomapsetentry) GetFilter() yfilter.YFilter { return cipsdynamiccryptomapsetentry.YFilter }

func (cipsdynamiccryptomapsetentry *CISCOIPSECMIB_Cipsdynamiccryptomapsettable_Cipsdynamiccryptomapsetentry) SetFilter(yf yfilter.YFilter) { cipsdynamiccryptomapsetentry.YFilter = yf }

func (cipsdynamiccryptomapsetentry *CISCOIPSECMIB_Cipsdynamiccryptomapsettable_Cipsdynamiccryptomapsetentry) GetGoName(yname string) string {
    if yname == "cipsDynamicCryptomapSetName" { return "Cipsdynamiccryptomapsetname" }
    if yname == "cipsDynamicCryptomapSetSize" { return "Cipsdynamiccryptomapsetsize" }
    if yname == "cipsDynamicCryptomapSetNumAssoc" { return "Cipsdynamiccryptomapsetnumassoc" }
    return ""
}

func (cipsdynamiccryptomapsetentry *CISCOIPSECMIB_Cipsdynamiccryptomapsettable_Cipsdynamiccryptomapsetentry) GetSegmentPath() string {
    return "cipsDynamicCryptomapSetEntry" + "[cipsDynamicCryptomapSetName='" + fmt.Sprintf("%v", cipsdynamiccryptomapsetentry.Cipsdynamiccryptomapsetname) + "']"
}

func (cipsdynamiccryptomapsetentry *CISCOIPSECMIB_Cipsdynamiccryptomapsettable_Cipsdynamiccryptomapsetentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipsdynamiccryptomapsetentry *CISCOIPSECMIB_Cipsdynamiccryptomapsettable_Cipsdynamiccryptomapsetentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipsdynamiccryptomapsetentry *CISCOIPSECMIB_Cipsdynamiccryptomapsettable_Cipsdynamiccryptomapsetentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipsDynamicCryptomapSetName"] = cipsdynamiccryptomapsetentry.Cipsdynamiccryptomapsetname
    leafs["cipsDynamicCryptomapSetSize"] = cipsdynamiccryptomapsetentry.Cipsdynamiccryptomapsetsize
    leafs["cipsDynamicCryptomapSetNumAssoc"] = cipsdynamiccryptomapsetentry.Cipsdynamiccryptomapsetnumassoc
    return leafs
}

func (cipsdynamiccryptomapsetentry *CISCOIPSECMIB_Cipsdynamiccryptomapsettable_Cipsdynamiccryptomapsetentry) GetBundleName() string { return "cisco_ios_xe" }

func (cipsdynamiccryptomapsetentry *CISCOIPSECMIB_Cipsdynamiccryptomapsettable_Cipsdynamiccryptomapsetentry) GetYangName() string { return "cipsDynamicCryptomapSetEntry" }

func (cipsdynamiccryptomapsetentry *CISCOIPSECMIB_Cipsdynamiccryptomapsettable_Cipsdynamiccryptomapsetentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsdynamiccryptomapsetentry *CISCOIPSECMIB_Cipsdynamiccryptomapsettable_Cipsdynamiccryptomapsetentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsdynamiccryptomapsetentry *CISCOIPSECMIB_Cipsdynamiccryptomapsettable_Cipsdynamiccryptomapsetentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsdynamiccryptomapsetentry *CISCOIPSECMIB_Cipsdynamiccryptomapsettable_Cipsdynamiccryptomapsetentry) SetParent(parent types.Entity) { cipsdynamiccryptomapsetentry.parent = parent }

func (cipsdynamiccryptomapsetentry *CISCOIPSECMIB_Cipsdynamiccryptomapsettable_Cipsdynamiccryptomapsetentry) GetParent() types.Entity { return cipsdynamiccryptomapsetentry.parent }

func (cipsdynamiccryptomapsetentry *CISCOIPSECMIB_Cipsdynamiccryptomapsettable_Cipsdynamiccryptomapsetentry) GetParentYangName() string { return "cipsDynamicCryptomapSetTable" }

// CISCOIPSECMIB_Cipsstaticcryptomaptable
// The table ilisting the member cryptomaps
// of the cryptomap sets that are configured
// on the managed entity.
type CISCOIPSECMIB_Cipsstaticcryptomaptable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains the attributes  associated with a single static  (fully
    // specified) cryptomap entry. This table does not include the members  of
    // dynamic cryptomap sets that may be linked with the parent static cryptomap
    // set. The type is slice of
    // CISCOIPSECMIB_Cipsstaticcryptomaptable_Cipsstaticcryptomapentry.
    Cipsstaticcryptomapentry []CISCOIPSECMIB_Cipsstaticcryptomaptable_Cipsstaticcryptomapentry
}

func (cipsstaticcryptomaptable *CISCOIPSECMIB_Cipsstaticcryptomaptable) GetFilter() yfilter.YFilter { return cipsstaticcryptomaptable.YFilter }

func (cipsstaticcryptomaptable *CISCOIPSECMIB_Cipsstaticcryptomaptable) SetFilter(yf yfilter.YFilter) { cipsstaticcryptomaptable.YFilter = yf }

func (cipsstaticcryptomaptable *CISCOIPSECMIB_Cipsstaticcryptomaptable) GetGoName(yname string) string {
    if yname == "cipsStaticCryptomapEntry" { return "Cipsstaticcryptomapentry" }
    return ""
}

func (cipsstaticcryptomaptable *CISCOIPSECMIB_Cipsstaticcryptomaptable) GetSegmentPath() string {
    return "cipsStaticCryptomapTable"
}

func (cipsstaticcryptomaptable *CISCOIPSECMIB_Cipsstaticcryptomaptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipsStaticCryptomapEntry" {
        for _, c := range cipsstaticcryptomaptable.Cipsstaticcryptomapentry {
            if cipsstaticcryptomaptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSECMIB_Cipsstaticcryptomaptable_Cipsstaticcryptomapentry{}
        cipsstaticcryptomaptable.Cipsstaticcryptomapentry = append(cipsstaticcryptomaptable.Cipsstaticcryptomapentry, child)
        return &cipsstaticcryptomaptable.Cipsstaticcryptomapentry[len(cipsstaticcryptomaptable.Cipsstaticcryptomapentry)-1]
    }
    return nil
}

func (cipsstaticcryptomaptable *CISCOIPSECMIB_Cipsstaticcryptomaptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cipsstaticcryptomaptable.Cipsstaticcryptomapentry {
        children[cipsstaticcryptomaptable.Cipsstaticcryptomapentry[i].GetSegmentPath()] = &cipsstaticcryptomaptable.Cipsstaticcryptomapentry[i]
    }
    return children
}

func (cipsstaticcryptomaptable *CISCOIPSECMIB_Cipsstaticcryptomaptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cipsstaticcryptomaptable *CISCOIPSECMIB_Cipsstaticcryptomaptable) GetBundleName() string { return "cisco_ios_xe" }

func (cipsstaticcryptomaptable *CISCOIPSECMIB_Cipsstaticcryptomaptable) GetYangName() string { return "cipsStaticCryptomapTable" }

func (cipsstaticcryptomaptable *CISCOIPSECMIB_Cipsstaticcryptomaptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsstaticcryptomaptable *CISCOIPSECMIB_Cipsstaticcryptomaptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsstaticcryptomaptable *CISCOIPSECMIB_Cipsstaticcryptomaptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsstaticcryptomaptable *CISCOIPSECMIB_Cipsstaticcryptomaptable) SetParent(parent types.Entity) { cipsstaticcryptomaptable.parent = parent }

func (cipsstaticcryptomaptable *CISCOIPSECMIB_Cipsstaticcryptomaptable) GetParent() types.Entity { return cipsstaticcryptomaptable.parent }

func (cipsstaticcryptomaptable *CISCOIPSECMIB_Cipsstaticcryptomaptable) GetParentYangName() string { return "CISCO-IPSEC-MIB" }

// CISCOIPSECMIB_Cipsstaticcryptomaptable_Cipsstaticcryptomapentry
// Each entry contains the attributes 
// associated with a single static 
// (fully specified) cryptomap entry.
// This table does not include the members 
// of dynamic cryptomap sets that may be
// linked with the parent static cryptomap set.
type CISCOIPSECMIB_Cipsstaticcryptomaptable_Cipsstaticcryptomapentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string. Refers to
    // cisco_ipsec_mib.CISCOIPSECMIB_Cipsstaticcryptomapsettable_Cipsstaticcryptomapsetentry_Cipsstaticcryptomapsetname
    Cipsstaticcryptomapsetname interface{}

    // This attribute is a key. The priority of the cryptomap entry in the 
    // cryptomap set. This is the second index component of this table. The type
    // is interface{} with range: 0..65535.
    Cipsstaticcryptomappriority interface{}

    // The type of the cryptomap entry. This can be an ISAKMP cryptomap, CET or
    // manual. Dynamic cryptomaps are not counted in this table. The type is
    // CryptomapType.
    Cipsstaticcryptomaptype interface{}

    // The description string entered by the operatoir while creating this
    // cryptomap. The string generally identifies a description and the purpose of
    // this policy. The type is string.
    Cipsstaticcryptomapdescr interface{}

    // The IP address of the current peer associated with  this IPSec policy item.
    // Traffic that is protected by this cryptomap is protected by a tunnel that
    // terminates at the device whose IP address is specified by this object. The
    // type is string with length: 4 | 16.
    Cipsstaticcryptomappeer interface{}

    // The number of peers associated with this cryptomap  entry. The peers other
    // than the one identified by  'cipsStaticCryptomapPeer' are backup peers.  
    // Manual cryptomaps may have only one peer. The type is interface{} with
    // range: 0..40.
    Cipsstaticcryptomapnumpeers interface{}

    // This object identifies if the tunnels instantiated due to this policy item
    // should use Perfect Forward Secrecy  (PFS) and if so, what group of Oakley
    // they should use. The type is DiffHellmanGrp.
    Cipsstaticcryptomappfs interface{}

    // This object identifies the lifetime of the IPSec Security Associations (SA)
    // created using this IPSec policy entry. If this value is zero, the lifetime
    // assumes the  value specified by the global lifetime parameter. The type is
    // interface{} with range: 0..None | 120..86400.
    Cipsstaticcryptomaplifetime interface{}

    // This object identifies the lifesize (maximum traffic in bytes that may be
    // carried) of the IPSec SAs created using this IPSec policy entry.  If this
    // value is zero, the lifetime assumes the  value specified by the global
    // lifesize parameter. The type is interface{} with range: 0..None |
    // 2560..536870912.
    Cipsstaticcryptomaplifesize interface{}

    // This object identifies the granularity of the IPSec SAs created using this
    // IPSec policy entry.  If this value is TRUE, distinct SA bundles are created
    // for distinct hosts at the end of the application traffic. The type is bool.
    Cipsstaticcryptomaplevelhost interface{}
}

func (cipsstaticcryptomapentry *CISCOIPSECMIB_Cipsstaticcryptomaptable_Cipsstaticcryptomapentry) GetFilter() yfilter.YFilter { return cipsstaticcryptomapentry.YFilter }

func (cipsstaticcryptomapentry *CISCOIPSECMIB_Cipsstaticcryptomaptable_Cipsstaticcryptomapentry) SetFilter(yf yfilter.YFilter) { cipsstaticcryptomapentry.YFilter = yf }

func (cipsstaticcryptomapentry *CISCOIPSECMIB_Cipsstaticcryptomaptable_Cipsstaticcryptomapentry) GetGoName(yname string) string {
    if yname == "cipsStaticCryptomapSetName" { return "Cipsstaticcryptomapsetname" }
    if yname == "cipsStaticCryptomapPriority" { return "Cipsstaticcryptomappriority" }
    if yname == "cipsStaticCryptomapType" { return "Cipsstaticcryptomaptype" }
    if yname == "cipsStaticCryptomapDescr" { return "Cipsstaticcryptomapdescr" }
    if yname == "cipsStaticCryptomapPeer" { return "Cipsstaticcryptomappeer" }
    if yname == "cipsStaticCryptomapNumPeers" { return "Cipsstaticcryptomapnumpeers" }
    if yname == "cipsStaticCryptomapPfs" { return "Cipsstaticcryptomappfs" }
    if yname == "cipsStaticCryptomapLifetime" { return "Cipsstaticcryptomaplifetime" }
    if yname == "cipsStaticCryptomapLifesize" { return "Cipsstaticcryptomaplifesize" }
    if yname == "cipsStaticCryptomapLevelHost" { return "Cipsstaticcryptomaplevelhost" }
    return ""
}

func (cipsstaticcryptomapentry *CISCOIPSECMIB_Cipsstaticcryptomaptable_Cipsstaticcryptomapentry) GetSegmentPath() string {
    return "cipsStaticCryptomapEntry" + "[cipsStaticCryptomapSetName='" + fmt.Sprintf("%v", cipsstaticcryptomapentry.Cipsstaticcryptomapsetname) + "']" + "[cipsStaticCryptomapPriority='" + fmt.Sprintf("%v", cipsstaticcryptomapentry.Cipsstaticcryptomappriority) + "']"
}

func (cipsstaticcryptomapentry *CISCOIPSECMIB_Cipsstaticcryptomaptable_Cipsstaticcryptomapentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipsstaticcryptomapentry *CISCOIPSECMIB_Cipsstaticcryptomaptable_Cipsstaticcryptomapentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipsstaticcryptomapentry *CISCOIPSECMIB_Cipsstaticcryptomaptable_Cipsstaticcryptomapentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cipsStaticCryptomapSetName"] = cipsstaticcryptomapentry.Cipsstaticcryptomapsetname
    leafs["cipsStaticCryptomapPriority"] = cipsstaticcryptomapentry.Cipsstaticcryptomappriority
    leafs["cipsStaticCryptomapType"] = cipsstaticcryptomapentry.Cipsstaticcryptomaptype
    leafs["cipsStaticCryptomapDescr"] = cipsstaticcryptomapentry.Cipsstaticcryptomapdescr
    leafs["cipsStaticCryptomapPeer"] = cipsstaticcryptomapentry.Cipsstaticcryptomappeer
    leafs["cipsStaticCryptomapNumPeers"] = cipsstaticcryptomapentry.Cipsstaticcryptomapnumpeers
    leafs["cipsStaticCryptomapPfs"] = cipsstaticcryptomapentry.Cipsstaticcryptomappfs
    leafs["cipsStaticCryptomapLifetime"] = cipsstaticcryptomapentry.Cipsstaticcryptomaplifetime
    leafs["cipsStaticCryptomapLifesize"] = cipsstaticcryptomapentry.Cipsstaticcryptomaplifesize
    leafs["cipsStaticCryptomapLevelHost"] = cipsstaticcryptomapentry.Cipsstaticcryptomaplevelhost
    return leafs
}

func (cipsstaticcryptomapentry *CISCOIPSECMIB_Cipsstaticcryptomaptable_Cipsstaticcryptomapentry) GetBundleName() string { return "cisco_ios_xe" }

func (cipsstaticcryptomapentry *CISCOIPSECMIB_Cipsstaticcryptomaptable_Cipsstaticcryptomapentry) GetYangName() string { return "cipsStaticCryptomapEntry" }

func (cipsstaticcryptomapentry *CISCOIPSECMIB_Cipsstaticcryptomaptable_Cipsstaticcryptomapentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipsstaticcryptomapentry *CISCOIPSECMIB_Cipsstaticcryptomaptable_Cipsstaticcryptomapentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipsstaticcryptomapentry *CISCOIPSECMIB_Cipsstaticcryptomaptable_Cipsstaticcryptomapentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipsstaticcryptomapentry *CISCOIPSECMIB_Cipsstaticcryptomaptable_Cipsstaticcryptomapentry) SetParent(parent types.Entity) { cipsstaticcryptomapentry.parent = parent }

func (cipsstaticcryptomapentry *CISCOIPSECMIB_Cipsstaticcryptomaptable_Cipsstaticcryptomapentry) GetParent() types.Entity { return cipsstaticcryptomapentry.parent }

func (cipsstaticcryptomapentry *CISCOIPSECMIB_Cipsstaticcryptomaptable_Cipsstaticcryptomapentry) GetParentYangName() string { return "cipsStaticCryptomapTable" }

// CISCOIPSECMIB_Cipscryptomapsetiftable
// The table lists the binding of cryptomap sets
// to the interfaces of the managed entity.
type CISCOIPSECMIB_Cipscryptomapsetiftable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains the record of the association between an interface and
    // a cryptomap set (static) that is defined on the managed entity.  Note that
    // the cryptomap set identified in  this binding must static. Dynamic
    // cryptomaps cannot be bound to interfaces. The type is slice of
    // CISCOIPSECMIB_Cipscryptomapsetiftable_Cipscryptomapsetifentry.
    Cipscryptomapsetifentry []CISCOIPSECMIB_Cipscryptomapsetiftable_Cipscryptomapsetifentry
}

func (cipscryptomapsetiftable *CISCOIPSECMIB_Cipscryptomapsetiftable) GetFilter() yfilter.YFilter { return cipscryptomapsetiftable.YFilter }

func (cipscryptomapsetiftable *CISCOIPSECMIB_Cipscryptomapsetiftable) SetFilter(yf yfilter.YFilter) { cipscryptomapsetiftable.YFilter = yf }

func (cipscryptomapsetiftable *CISCOIPSECMIB_Cipscryptomapsetiftable) GetGoName(yname string) string {
    if yname == "cipsCryptomapSetIfEntry" { return "Cipscryptomapsetifentry" }
    return ""
}

func (cipscryptomapsetiftable *CISCOIPSECMIB_Cipscryptomapsetiftable) GetSegmentPath() string {
    return "cipsCryptomapSetIfTable"
}

func (cipscryptomapsetiftable *CISCOIPSECMIB_Cipscryptomapsetiftable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cipsCryptomapSetIfEntry" {
        for _, c := range cipscryptomapsetiftable.Cipscryptomapsetifentry {
            if cipscryptomapsetiftable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSECMIB_Cipscryptomapsetiftable_Cipscryptomapsetifentry{}
        cipscryptomapsetiftable.Cipscryptomapsetifentry = append(cipscryptomapsetiftable.Cipscryptomapsetifentry, child)
        return &cipscryptomapsetiftable.Cipscryptomapsetifentry[len(cipscryptomapsetiftable.Cipscryptomapsetifentry)-1]
    }
    return nil
}

func (cipscryptomapsetiftable *CISCOIPSECMIB_Cipscryptomapsetiftable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cipscryptomapsetiftable.Cipscryptomapsetifentry {
        children[cipscryptomapsetiftable.Cipscryptomapsetifentry[i].GetSegmentPath()] = &cipscryptomapsetiftable.Cipscryptomapsetifentry[i]
    }
    return children
}

func (cipscryptomapsetiftable *CISCOIPSECMIB_Cipscryptomapsetiftable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cipscryptomapsetiftable *CISCOIPSECMIB_Cipscryptomapsetiftable) GetBundleName() string { return "cisco_ios_xe" }

func (cipscryptomapsetiftable *CISCOIPSECMIB_Cipscryptomapsetiftable) GetYangName() string { return "cipsCryptomapSetIfTable" }

func (cipscryptomapsetiftable *CISCOIPSECMIB_Cipscryptomapsetiftable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipscryptomapsetiftable *CISCOIPSECMIB_Cipscryptomapsetiftable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipscryptomapsetiftable *CISCOIPSECMIB_Cipscryptomapsetiftable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipscryptomapsetiftable *CISCOIPSECMIB_Cipscryptomapsetiftable) SetParent(parent types.Entity) { cipscryptomapsetiftable.parent = parent }

func (cipscryptomapsetiftable *CISCOIPSECMIB_Cipscryptomapsetiftable) GetParent() types.Entity { return cipscryptomapsetiftable.parent }

func (cipscryptomapsetiftable *CISCOIPSECMIB_Cipscryptomapsetiftable) GetParentYangName() string { return "CISCO-IPSEC-MIB" }

// CISCOIPSECMIB_Cipscryptomapsetiftable_Cipscryptomapsetifentry
// Each entry contains the record of
// the association between an interface
// and a cryptomap set (static) that is defined
// on the managed entity.
// 
// Note that the cryptomap set identified in 
// this binding must static. Dynamic cryptomaps cannot
// be bound to interfaces.
type CISCOIPSECMIB_Cipscryptomapsetiftable_Cipscryptomapsetifentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_Iftable_Ifentry_Ifindex
    Ifindex interface{}

    // This attribute is a key. The type is string. Refers to
    // cisco_ipsec_mib.CISCOIPSECMIB_Cipsstaticcryptomapsettable_Cipsstaticcryptomapsetentry_Cipsstaticcryptomapsetname
    Cipsstaticcryptomapsetname interface{}

    // The value of this object identifies if the interface to which the cryptomap
    // set is attached is a tunnel (such as a GRE or PPTP tunnel). The type is
    // bool.
    Cipscryptomapsetifvirtual interface{}

    // This object identifies the status of the binding  of the specified
    // cryptomap set with the specified interface. The value when queried is
    // always 'attached'.  When set to 'detached', the cryptomap set if detached 
    // from the specified interface. The effect of this is same  as the CLI
    // command  	config-if# no crypto map cryptomapSetName  Setting the value to
    // 'attached' will result in  SNMP General Error. The type is
    // CryptomapSetBindStatus.
    Cipscryptomapsetifstatus interface{}
}

func (cipscryptomapsetifentry *CISCOIPSECMIB_Cipscryptomapsetiftable_Cipscryptomapsetifentry) GetFilter() yfilter.YFilter { return cipscryptomapsetifentry.YFilter }

func (cipscryptomapsetifentry *CISCOIPSECMIB_Cipscryptomapsetiftable_Cipscryptomapsetifentry) SetFilter(yf yfilter.YFilter) { cipscryptomapsetifentry.YFilter = yf }

func (cipscryptomapsetifentry *CISCOIPSECMIB_Cipscryptomapsetiftable_Cipscryptomapsetifentry) GetGoName(yname string) string {
    if yname == "ifIndex" { return "Ifindex" }
    if yname == "cipsStaticCryptomapSetName" { return "Cipsstaticcryptomapsetname" }
    if yname == "cipsCryptomapSetIfVirtual" { return "Cipscryptomapsetifvirtual" }
    if yname == "cipsCryptomapSetIfStatus" { return "Cipscryptomapsetifstatus" }
    return ""
}

func (cipscryptomapsetifentry *CISCOIPSECMIB_Cipscryptomapsetiftable_Cipscryptomapsetifentry) GetSegmentPath() string {
    return "cipsCryptomapSetIfEntry" + "[ifIndex='" + fmt.Sprintf("%v", cipscryptomapsetifentry.Ifindex) + "']" + "[cipsStaticCryptomapSetName='" + fmt.Sprintf("%v", cipscryptomapsetifentry.Cipsstaticcryptomapsetname) + "']"
}

func (cipscryptomapsetifentry *CISCOIPSECMIB_Cipscryptomapsetiftable_Cipscryptomapsetifentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cipscryptomapsetifentry *CISCOIPSECMIB_Cipscryptomapsetiftable_Cipscryptomapsetifentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cipscryptomapsetifentry *CISCOIPSECMIB_Cipscryptomapsetiftable_Cipscryptomapsetifentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ifIndex"] = cipscryptomapsetifentry.Ifindex
    leafs["cipsStaticCryptomapSetName"] = cipscryptomapsetifentry.Cipsstaticcryptomapsetname
    leafs["cipsCryptomapSetIfVirtual"] = cipscryptomapsetifentry.Cipscryptomapsetifvirtual
    leafs["cipsCryptomapSetIfStatus"] = cipscryptomapsetifentry.Cipscryptomapsetifstatus
    return leafs
}

func (cipscryptomapsetifentry *CISCOIPSECMIB_Cipscryptomapsetiftable_Cipscryptomapsetifentry) GetBundleName() string { return "cisco_ios_xe" }

func (cipscryptomapsetifentry *CISCOIPSECMIB_Cipscryptomapsetiftable_Cipscryptomapsetifentry) GetYangName() string { return "cipsCryptomapSetIfEntry" }

func (cipscryptomapsetifentry *CISCOIPSECMIB_Cipscryptomapsetiftable_Cipscryptomapsetifentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cipscryptomapsetifentry *CISCOIPSECMIB_Cipscryptomapsetiftable_Cipscryptomapsetifentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cipscryptomapsetifentry *CISCOIPSECMIB_Cipscryptomapsetiftable_Cipscryptomapsetifentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cipscryptomapsetifentry *CISCOIPSECMIB_Cipscryptomapsetiftable_Cipscryptomapsetifentry) SetParent(parent types.Entity) { cipscryptomapsetifentry.parent = parent }

func (cipscryptomapsetifentry *CISCOIPSECMIB_Cipscryptomapsetiftable_Cipscryptomapsetifentry) GetParent() types.Entity { return cipscryptomapsetifentry.parent }

func (cipscryptomapsetifentry *CISCOIPSECMIB_Cipscryptomapsetiftable_Cipscryptomapsetifentry) GetParentYangName() string { return "cipsCryptomapSetIfTable" }

