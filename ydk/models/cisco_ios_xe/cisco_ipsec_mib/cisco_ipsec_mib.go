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

// CryptomapSetBindStatus represents SNMP General Error.
type CryptomapSetBindStatus string

const (
    CryptomapSetBindStatus_unknown CryptomapSetBindStatus = "unknown"

    CryptomapSetBindStatus_attached CryptomapSetBindStatus = "attached"

    CryptomapSetBindStatus_detached CryptomapSetBindStatus = "detached"
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

// IkeIdentityType represents 	Main Mode of IPSec tunnel setup.
type IkeIdentityType string

const (
    IkeIdentityType_isakmpIdTypeUNKNOWN IkeIdentityType = "isakmpIdTypeUNKNOWN"

    IkeIdentityType_isakmpIdTypeADDRESS IkeIdentityType = "isakmpIdTypeADDRESS"

    IkeIdentityType_isakmpIdTypeHOSTNAME IkeIdentityType = "isakmpIdTypeHOSTNAME"
)

// DiffHellmanGrp represents The Diffie Hellman Group used in negotiations.
type DiffHellmanGrp string

const (
    DiffHellmanGrp_none DiffHellmanGrp = "none"

    DiffHellmanGrp_dhGroup1 DiffHellmanGrp = "dhGroup1"

    DiffHellmanGrp_dhGroup2 DiffHellmanGrp = "dhGroup2"
)

// EncryptAlgo represents The encryption algorithm used in negotiations.
type EncryptAlgo string

const (
    EncryptAlgo_none EncryptAlgo = "none"

    EncryptAlgo_des EncryptAlgo = "des"

    EncryptAlgo_des3 EncryptAlgo = "des3"
)

// TrapStatus represents The administrative status for sending a TRAP.
type TrapStatus string

const (
    TrapStatus_enabled TrapStatus = "enabled"

    TrapStatus_disabled TrapStatus = "disabled"
)

// CISCOIPSECMIB
type CISCOIPSECMIB struct {
    EntityData types.CommonEntityData
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

func (cISCOIPSECMIB *CISCOIPSECMIB) GetEntityData() *types.CommonEntityData {
    cISCOIPSECMIB.EntityData.YFilter = cISCOIPSECMIB.YFilter
    cISCOIPSECMIB.EntityData.YangName = "CISCO-IPSEC-MIB"
    cISCOIPSECMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIPSECMIB.EntityData.ParentYangName = "CISCO-IPSEC-MIB"
    cISCOIPSECMIB.EntityData.SegmentPath = "CISCO-IPSEC-MIB:CISCO-IPSEC-MIB"
    cISCOIPSECMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIPSECMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIPSECMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIPSECMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOIPSECMIB.EntityData.Children["cipsIsakmpGroup"] = types.YChild{"Cipsisakmpgroup", &cISCOIPSECMIB.Cipsisakmpgroup}
    cISCOIPSECMIB.EntityData.Children["cipsIPsecGlobals"] = types.YChild{"Cipsipsecglobals", &cISCOIPSECMIB.Cipsipsecglobals}
    cISCOIPSECMIB.EntityData.Children["cipsIPsecStatistics"] = types.YChild{"Cipsipsecstatistics", &cISCOIPSECMIB.Cipsipsecstatistics}
    cISCOIPSECMIB.EntityData.Children["cipsSysCapacityGroup"] = types.YChild{"Cipssyscapacitygroup", &cISCOIPSECMIB.Cipssyscapacitygroup}
    cISCOIPSECMIB.EntityData.Children["cipsTrapCntlGroup"] = types.YChild{"Cipstrapcntlgroup", &cISCOIPSECMIB.Cipstrapcntlgroup}
    cISCOIPSECMIB.EntityData.Children["cipsIsakmpPolicyTable"] = types.YChild{"Cipsisakmppolicytable", &cISCOIPSECMIB.Cipsisakmppolicytable}
    cISCOIPSECMIB.EntityData.Children["cipsStaticCryptomapSetTable"] = types.YChild{"Cipsstaticcryptomapsettable", &cISCOIPSECMIB.Cipsstaticcryptomapsettable}
    cISCOIPSECMIB.EntityData.Children["cipsDynamicCryptomapSetTable"] = types.YChild{"Cipsdynamiccryptomapsettable", &cISCOIPSECMIB.Cipsdynamiccryptomapsettable}
    cISCOIPSECMIB.EntityData.Children["cipsStaticCryptomapTable"] = types.YChild{"Cipsstaticcryptomaptable", &cISCOIPSECMIB.Cipsstaticcryptomaptable}
    cISCOIPSECMIB.EntityData.Children["cipsCryptomapSetIfTable"] = types.YChild{"Cipscryptomapsetiftable", &cISCOIPSECMIB.Cipscryptomapsetiftable}
    cISCOIPSECMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOIPSECMIB.EntityData)
}

// CISCOIPSECMIB_Cipsisakmpgroup
type CISCOIPSECMIB_Cipsisakmpgroup struct {
    EntityData types.CommonEntityData
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

func (cipsisakmpgroup *CISCOIPSECMIB_Cipsisakmpgroup) GetEntityData() *types.CommonEntityData {
    cipsisakmpgroup.EntityData.YFilter = cipsisakmpgroup.YFilter
    cipsisakmpgroup.EntityData.YangName = "cipsIsakmpGroup"
    cipsisakmpgroup.EntityData.BundleName = "cisco_ios_xe"
    cipsisakmpgroup.EntityData.ParentYangName = "CISCO-IPSEC-MIB"
    cipsisakmpgroup.EntityData.SegmentPath = "cipsIsakmpGroup"
    cipsisakmpgroup.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsisakmpgroup.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsisakmpgroup.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsisakmpgroup.EntityData.Children = make(map[string]types.YChild)
    cipsisakmpgroup.EntityData.Leafs = make(map[string]types.YLeaf)
    cipsisakmpgroup.EntityData.Leafs["cipsIsakmpEnabled"] = types.YLeaf{"Cipsisakmpenabled", cipsisakmpgroup.Cipsisakmpenabled}
    cipsisakmpgroup.EntityData.Leafs["cipsIsakmpIdentity"] = types.YLeaf{"Cipsisakmpidentity", cipsisakmpgroup.Cipsisakmpidentity}
    cipsisakmpgroup.EntityData.Leafs["cipsIsakmpKeepaliveInterval"] = types.YLeaf{"Cipsisakmpkeepaliveinterval", cipsisakmpgroup.Cipsisakmpkeepaliveinterval}
    cipsisakmpgroup.EntityData.Leafs["cipsNumIsakmpPolicies"] = types.YLeaf{"Cipsnumisakmppolicies", cipsisakmpgroup.Cipsnumisakmppolicies}
    return &(cipsisakmpgroup.EntityData)
}

// CISCOIPSECMIB_Cipsipsecglobals
type CISCOIPSECMIB_Cipsipsecglobals struct {
    EntityData types.CommonEntityData
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

func (cipsipsecglobals *CISCOIPSECMIB_Cipsipsecglobals) GetEntityData() *types.CommonEntityData {
    cipsipsecglobals.EntityData.YFilter = cipsipsecglobals.YFilter
    cipsipsecglobals.EntityData.YangName = "cipsIPsecGlobals"
    cipsipsecglobals.EntityData.BundleName = "cisco_ios_xe"
    cipsipsecglobals.EntityData.ParentYangName = "CISCO-IPSEC-MIB"
    cipsipsecglobals.EntityData.SegmentPath = "cipsIPsecGlobals"
    cipsipsecglobals.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsipsecglobals.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsipsecglobals.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsipsecglobals.EntityData.Children = make(map[string]types.YChild)
    cipsipsecglobals.EntityData.Leafs = make(map[string]types.YLeaf)
    cipsipsecglobals.EntityData.Leafs["cipsSALifetime"] = types.YLeaf{"Cipssalifetime", cipsipsecglobals.Cipssalifetime}
    cipsipsecglobals.EntityData.Leafs["cipsSALifesize"] = types.YLeaf{"Cipssalifesize", cipsipsecglobals.Cipssalifesize}
    cipsipsecglobals.EntityData.Leafs["cipsNumStaticCryptomapSets"] = types.YLeaf{"Cipsnumstaticcryptomapsets", cipsipsecglobals.Cipsnumstaticcryptomapsets}
    cipsipsecglobals.EntityData.Leafs["cipsNumCETCryptomapSets"] = types.YLeaf{"Cipsnumcetcryptomapsets", cipsipsecglobals.Cipsnumcetcryptomapsets}
    cipsipsecglobals.EntityData.Leafs["cipsNumDynamicCryptomapSets"] = types.YLeaf{"Cipsnumdynamiccryptomapsets", cipsipsecglobals.Cipsnumdynamiccryptomapsets}
    cipsipsecglobals.EntityData.Leafs["cipsNumTEDCryptomapSets"] = types.YLeaf{"Cipsnumtedcryptomapsets", cipsipsecglobals.Cipsnumtedcryptomapsets}
    return &(cipsipsecglobals.EntityData)
}

// CISCOIPSECMIB_Cipsipsecstatistics
type CISCOIPSECMIB_Cipsipsecstatistics struct {
    EntityData types.CommonEntityData
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

func (cipsipsecstatistics *CISCOIPSECMIB_Cipsipsecstatistics) GetEntityData() *types.CommonEntityData {
    cipsipsecstatistics.EntityData.YFilter = cipsipsecstatistics.YFilter
    cipsipsecstatistics.EntityData.YangName = "cipsIPsecStatistics"
    cipsipsecstatistics.EntityData.BundleName = "cisco_ios_xe"
    cipsipsecstatistics.EntityData.ParentYangName = "CISCO-IPSEC-MIB"
    cipsipsecstatistics.EntityData.SegmentPath = "cipsIPsecStatistics"
    cipsipsecstatistics.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsipsecstatistics.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsipsecstatistics.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsipsecstatistics.EntityData.Children = make(map[string]types.YChild)
    cipsipsecstatistics.EntityData.Leafs = make(map[string]types.YLeaf)
    cipsipsecstatistics.EntityData.Leafs["cipsNumTEDProbesReceived"] = types.YLeaf{"Cipsnumtedprobesreceived", cipsipsecstatistics.Cipsnumtedprobesreceived}
    cipsipsecstatistics.EntityData.Leafs["cipsNumTEDProbesSent"] = types.YLeaf{"Cipsnumtedprobessent", cipsipsecstatistics.Cipsnumtedprobessent}
    cipsipsecstatistics.EntityData.Leafs["cipsNumTEDFailures"] = types.YLeaf{"Cipsnumtedfailures", cipsipsecstatistics.Cipsnumtedfailures}
    return &(cipsipsecstatistics.EntityData)
}

// CISCOIPSECMIB_Cipssyscapacitygroup
type CISCOIPSECMIB_Cipssyscapacitygroup struct {
    EntityData types.CommonEntityData
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

func (cipssyscapacitygroup *CISCOIPSECMIB_Cipssyscapacitygroup) GetEntityData() *types.CommonEntityData {
    cipssyscapacitygroup.EntityData.YFilter = cipssyscapacitygroup.YFilter
    cipssyscapacitygroup.EntityData.YangName = "cipsSysCapacityGroup"
    cipssyscapacitygroup.EntityData.BundleName = "cisco_ios_xe"
    cipssyscapacitygroup.EntityData.ParentYangName = "CISCO-IPSEC-MIB"
    cipssyscapacitygroup.EntityData.SegmentPath = "cipsSysCapacityGroup"
    cipssyscapacitygroup.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipssyscapacitygroup.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipssyscapacitygroup.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipssyscapacitygroup.EntityData.Children = make(map[string]types.YChild)
    cipssyscapacitygroup.EntityData.Leafs = make(map[string]types.YLeaf)
    cipssyscapacitygroup.EntityData.Leafs["cipsMaxSAs"] = types.YLeaf{"Cipsmaxsas", cipssyscapacitygroup.Cipsmaxsas}
    cipssyscapacitygroup.EntityData.Leafs["cips3DesCapable"] = types.YLeaf{"Cips3Descapable", cipssyscapacitygroup.Cips3Descapable}
    return &(cipssyscapacitygroup.EntityData)
}

// CISCOIPSECMIB_Cipstrapcntlgroup
type CISCOIPSECMIB_Cipstrapcntlgroup struct {
    EntityData types.CommonEntityData
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

func (cipstrapcntlgroup *CISCOIPSECMIB_Cipstrapcntlgroup) GetEntityData() *types.CommonEntityData {
    cipstrapcntlgroup.EntityData.YFilter = cipstrapcntlgroup.YFilter
    cipstrapcntlgroup.EntityData.YangName = "cipsTrapCntlGroup"
    cipstrapcntlgroup.EntityData.BundleName = "cisco_ios_xe"
    cipstrapcntlgroup.EntityData.ParentYangName = "CISCO-IPSEC-MIB"
    cipstrapcntlgroup.EntityData.SegmentPath = "cipsTrapCntlGroup"
    cipstrapcntlgroup.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipstrapcntlgroup.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipstrapcntlgroup.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipstrapcntlgroup.EntityData.Children = make(map[string]types.YChild)
    cipstrapcntlgroup.EntityData.Leafs = make(map[string]types.YLeaf)
    cipstrapcntlgroup.EntityData.Leafs["cipsCntlIsakmpPolicyAdded"] = types.YLeaf{"Cipscntlisakmppolicyadded", cipstrapcntlgroup.Cipscntlisakmppolicyadded}
    cipstrapcntlgroup.EntityData.Leafs["cipsCntlIsakmpPolicyDeleted"] = types.YLeaf{"Cipscntlisakmppolicydeleted", cipstrapcntlgroup.Cipscntlisakmppolicydeleted}
    cipstrapcntlgroup.EntityData.Leafs["cipsCntlCryptomapAdded"] = types.YLeaf{"Cipscntlcryptomapadded", cipstrapcntlgroup.Cipscntlcryptomapadded}
    cipstrapcntlgroup.EntityData.Leafs["cipsCntlCryptomapDeleted"] = types.YLeaf{"Cipscntlcryptomapdeleted", cipstrapcntlgroup.Cipscntlcryptomapdeleted}
    cipstrapcntlgroup.EntityData.Leafs["cipsCntlCryptomapSetAttached"] = types.YLeaf{"Cipscntlcryptomapsetattached", cipstrapcntlgroup.Cipscntlcryptomapsetattached}
    cipstrapcntlgroup.EntityData.Leafs["cipsCntlCryptomapSetDetached"] = types.YLeaf{"Cipscntlcryptomapsetdetached", cipstrapcntlgroup.Cipscntlcryptomapsetdetached}
    cipstrapcntlgroup.EntityData.Leafs["cipsCntlTooManySAs"] = types.YLeaf{"Cipscntltoomanysas", cipstrapcntlgroup.Cipscntltoomanysas}
    return &(cipstrapcntlgroup.EntityData)
}

// CISCOIPSECMIB_Cipsisakmppolicytable
// The table containing the list of all
// ISAKMP policy entries configured by the operator.
type CISCOIPSECMIB_Cipsisakmppolicytable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes  associated with a single ISAKMP Policy
    // entry. The type is slice of
    // CISCOIPSECMIB_Cipsisakmppolicytable_Cipsisakmppolicyentry.
    Cipsisakmppolicyentry []CISCOIPSECMIB_Cipsisakmppolicytable_Cipsisakmppolicyentry
}

func (cipsisakmppolicytable *CISCOIPSECMIB_Cipsisakmppolicytable) GetEntityData() *types.CommonEntityData {
    cipsisakmppolicytable.EntityData.YFilter = cipsisakmppolicytable.YFilter
    cipsisakmppolicytable.EntityData.YangName = "cipsIsakmpPolicyTable"
    cipsisakmppolicytable.EntityData.BundleName = "cisco_ios_xe"
    cipsisakmppolicytable.EntityData.ParentYangName = "CISCO-IPSEC-MIB"
    cipsisakmppolicytable.EntityData.SegmentPath = "cipsIsakmpPolicyTable"
    cipsisakmppolicytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsisakmppolicytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsisakmppolicytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsisakmppolicytable.EntityData.Children = make(map[string]types.YChild)
    cipsisakmppolicytable.EntityData.Children["cipsIsakmpPolicyEntry"] = types.YChild{"Cipsisakmppolicyentry", nil}
    for i := range cipsisakmppolicytable.Cipsisakmppolicyentry {
        cipsisakmppolicytable.EntityData.Children[types.GetSegmentPath(&cipsisakmppolicytable.Cipsisakmppolicyentry[i])] = types.YChild{"Cipsisakmppolicyentry", &cipsisakmppolicytable.Cipsisakmppolicyentry[i]}
    }
    cipsisakmppolicytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cipsisakmppolicytable.EntityData)
}

// CISCOIPSECMIB_Cipsisakmppolicytable_Cipsisakmppolicyentry
// Each entry contains the attributes 
// associated with a single ISAKMP
// Policy entry.
type CISCOIPSECMIB_Cipsisakmppolicytable_Cipsisakmppolicyentry struct {
    EntityData types.CommonEntityData
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

func (cipsisakmppolicyentry *CISCOIPSECMIB_Cipsisakmppolicytable_Cipsisakmppolicyentry) GetEntityData() *types.CommonEntityData {
    cipsisakmppolicyentry.EntityData.YFilter = cipsisakmppolicyentry.YFilter
    cipsisakmppolicyentry.EntityData.YangName = "cipsIsakmpPolicyEntry"
    cipsisakmppolicyentry.EntityData.BundleName = "cisco_ios_xe"
    cipsisakmppolicyentry.EntityData.ParentYangName = "cipsIsakmpPolicyTable"
    cipsisakmppolicyentry.EntityData.SegmentPath = "cipsIsakmpPolicyEntry" + "[cipsIsakmpPolPriority='" + fmt.Sprintf("%v", cipsisakmppolicyentry.Cipsisakmppolpriority) + "']"
    cipsisakmppolicyentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsisakmppolicyentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsisakmppolicyentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsisakmppolicyentry.EntityData.Children = make(map[string]types.YChild)
    cipsisakmppolicyentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cipsisakmppolicyentry.EntityData.Leafs["cipsIsakmpPolPriority"] = types.YLeaf{"Cipsisakmppolpriority", cipsisakmppolicyentry.Cipsisakmppolpriority}
    cipsisakmppolicyentry.EntityData.Leafs["cipsIsakmpPolEncr"] = types.YLeaf{"Cipsisakmppolencr", cipsisakmppolicyentry.Cipsisakmppolencr}
    cipsisakmppolicyentry.EntityData.Leafs["cipsIsakmpPolHash"] = types.YLeaf{"Cipsisakmppolhash", cipsisakmppolicyentry.Cipsisakmppolhash}
    cipsisakmppolicyentry.EntityData.Leafs["cipsIsakmpPolAuth"] = types.YLeaf{"Cipsisakmppolauth", cipsisakmppolicyentry.Cipsisakmppolauth}
    cipsisakmppolicyentry.EntityData.Leafs["cipsIsakmpPolGroup"] = types.YLeaf{"Cipsisakmppolgroup", cipsisakmppolicyentry.Cipsisakmppolgroup}
    cipsisakmppolicyentry.EntityData.Leafs["cipsIsakmpPolLifetime"] = types.YLeaf{"Cipsisakmppollifetime", cipsisakmppolicyentry.Cipsisakmppollifetime}
    return &(cipsisakmppolicyentry.EntityData)
}

// CISCOIPSECMIB_Cipsstaticcryptomapsettable
// The table containing the list of all
// cryptomap sets that are fully specified
// and are not wild-carded.
// 
// The operator may include different types of
// cryptomaps in such a set - manual, CET,
// ISAKMP or dynamic.
type CISCOIPSECMIB_Cipsstaticcryptomapsettable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes  associated with a single static 
    // cryptomap set. The type is slice of
    // CISCOIPSECMIB_Cipsstaticcryptomapsettable_Cipsstaticcryptomapsetentry.
    Cipsstaticcryptomapsetentry []CISCOIPSECMIB_Cipsstaticcryptomapsettable_Cipsstaticcryptomapsetentry
}

func (cipsstaticcryptomapsettable *CISCOIPSECMIB_Cipsstaticcryptomapsettable) GetEntityData() *types.CommonEntityData {
    cipsstaticcryptomapsettable.EntityData.YFilter = cipsstaticcryptomapsettable.YFilter
    cipsstaticcryptomapsettable.EntityData.YangName = "cipsStaticCryptomapSetTable"
    cipsstaticcryptomapsettable.EntityData.BundleName = "cisco_ios_xe"
    cipsstaticcryptomapsettable.EntityData.ParentYangName = "CISCO-IPSEC-MIB"
    cipsstaticcryptomapsettable.EntityData.SegmentPath = "cipsStaticCryptomapSetTable"
    cipsstaticcryptomapsettable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsstaticcryptomapsettable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsstaticcryptomapsettable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsstaticcryptomapsettable.EntityData.Children = make(map[string]types.YChild)
    cipsstaticcryptomapsettable.EntityData.Children["cipsStaticCryptomapSetEntry"] = types.YChild{"Cipsstaticcryptomapsetentry", nil}
    for i := range cipsstaticcryptomapsettable.Cipsstaticcryptomapsetentry {
        cipsstaticcryptomapsettable.EntityData.Children[types.GetSegmentPath(&cipsstaticcryptomapsettable.Cipsstaticcryptomapsetentry[i])] = types.YChild{"Cipsstaticcryptomapsetentry", &cipsstaticcryptomapsettable.Cipsstaticcryptomapsetentry[i]}
    }
    cipsstaticcryptomapsettable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cipsstaticcryptomapsettable.EntityData)
}

// CISCOIPSECMIB_Cipsstaticcryptomapsettable_Cipsstaticcryptomapsetentry
// Each entry contains the attributes 
// associated with a single static 
// cryptomap set.
type CISCOIPSECMIB_Cipsstaticcryptomapsettable_Cipsstaticcryptomapsetentry struct {
    EntityData types.CommonEntityData
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

func (cipsstaticcryptomapsetentry *CISCOIPSECMIB_Cipsstaticcryptomapsettable_Cipsstaticcryptomapsetentry) GetEntityData() *types.CommonEntityData {
    cipsstaticcryptomapsetentry.EntityData.YFilter = cipsstaticcryptomapsetentry.YFilter
    cipsstaticcryptomapsetentry.EntityData.YangName = "cipsStaticCryptomapSetEntry"
    cipsstaticcryptomapsetentry.EntityData.BundleName = "cisco_ios_xe"
    cipsstaticcryptomapsetentry.EntityData.ParentYangName = "cipsStaticCryptomapSetTable"
    cipsstaticcryptomapsetentry.EntityData.SegmentPath = "cipsStaticCryptomapSetEntry" + "[cipsStaticCryptomapSetName='" + fmt.Sprintf("%v", cipsstaticcryptomapsetentry.Cipsstaticcryptomapsetname) + "']"
    cipsstaticcryptomapsetentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsstaticcryptomapsetentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsstaticcryptomapsetentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsstaticcryptomapsetentry.EntityData.Children = make(map[string]types.YChild)
    cipsstaticcryptomapsetentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cipsstaticcryptomapsetentry.EntityData.Leafs["cipsStaticCryptomapSetName"] = types.YLeaf{"Cipsstaticcryptomapsetname", cipsstaticcryptomapsetentry.Cipsstaticcryptomapsetname}
    cipsstaticcryptomapsetentry.EntityData.Leafs["cipsStaticCryptomapSetSize"] = types.YLeaf{"Cipsstaticcryptomapsetsize", cipsstaticcryptomapsetentry.Cipsstaticcryptomapsetsize}
    cipsstaticcryptomapsetentry.EntityData.Leafs["cipsStaticCryptomapSetNumIsakmp"] = types.YLeaf{"Cipsstaticcryptomapsetnumisakmp", cipsstaticcryptomapsetentry.Cipsstaticcryptomapsetnumisakmp}
    cipsstaticcryptomapsetentry.EntityData.Leafs["cipsStaticCryptomapSetNumManual"] = types.YLeaf{"Cipsstaticcryptomapsetnummanual", cipsstaticcryptomapsetentry.Cipsstaticcryptomapsetnummanual}
    cipsstaticcryptomapsetentry.EntityData.Leafs["cipsStaticCryptomapSetNumCET"] = types.YLeaf{"Cipsstaticcryptomapsetnumcet", cipsstaticcryptomapsetentry.Cipsstaticcryptomapsetnumcet}
    cipsstaticcryptomapsetentry.EntityData.Leafs["cipsStaticCryptomapSetNumDynamic"] = types.YLeaf{"Cipsstaticcryptomapsetnumdynamic", cipsstaticcryptomapsetentry.Cipsstaticcryptomapsetnumdynamic}
    cipsstaticcryptomapsetentry.EntityData.Leafs["cipsStaticCryptomapSetNumDisc"] = types.YLeaf{"Cipsstaticcryptomapsetnumdisc", cipsstaticcryptomapsetentry.Cipsstaticcryptomapsetnumdisc}
    cipsstaticcryptomapsetentry.EntityData.Leafs["cipsStaticCryptomapSetNumSAs"] = types.YLeaf{"Cipsstaticcryptomapsetnumsas", cipsstaticcryptomapsetentry.Cipsstaticcryptomapsetnumsas}
    return &(cipsstaticcryptomapsetentry.EntityData)
}

// CISCOIPSECMIB_Cipsdynamiccryptomapsettable
// The table containing the list of all dynamic
// cryptomaps that use IKE, defined on 
//  the managed entity.
type CISCOIPSECMIB_Cipsdynamiccryptomapsettable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with a single dynamic
    // cryptomap template. The type is slice of
    // CISCOIPSECMIB_Cipsdynamiccryptomapsettable_Cipsdynamiccryptomapsetentry.
    Cipsdynamiccryptomapsetentry []CISCOIPSECMIB_Cipsdynamiccryptomapsettable_Cipsdynamiccryptomapsetentry
}

func (cipsdynamiccryptomapsettable *CISCOIPSECMIB_Cipsdynamiccryptomapsettable) GetEntityData() *types.CommonEntityData {
    cipsdynamiccryptomapsettable.EntityData.YFilter = cipsdynamiccryptomapsettable.YFilter
    cipsdynamiccryptomapsettable.EntityData.YangName = "cipsDynamicCryptomapSetTable"
    cipsdynamiccryptomapsettable.EntityData.BundleName = "cisco_ios_xe"
    cipsdynamiccryptomapsettable.EntityData.ParentYangName = "CISCO-IPSEC-MIB"
    cipsdynamiccryptomapsettable.EntityData.SegmentPath = "cipsDynamicCryptomapSetTable"
    cipsdynamiccryptomapsettable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsdynamiccryptomapsettable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsdynamiccryptomapsettable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsdynamiccryptomapsettable.EntityData.Children = make(map[string]types.YChild)
    cipsdynamiccryptomapsettable.EntityData.Children["cipsDynamicCryptomapSetEntry"] = types.YChild{"Cipsdynamiccryptomapsetentry", nil}
    for i := range cipsdynamiccryptomapsettable.Cipsdynamiccryptomapsetentry {
        cipsdynamiccryptomapsettable.EntityData.Children[types.GetSegmentPath(&cipsdynamiccryptomapsettable.Cipsdynamiccryptomapsetentry[i])] = types.YChild{"Cipsdynamiccryptomapsetentry", &cipsdynamiccryptomapsettable.Cipsdynamiccryptomapsetentry[i]}
    }
    cipsdynamiccryptomapsettable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cipsdynamiccryptomapsettable.EntityData)
}

// CISCOIPSECMIB_Cipsdynamiccryptomapsettable_Cipsdynamiccryptomapsetentry
// Each entry contains the attributes associated
// with a single dynamic cryptomap template.
type CISCOIPSECMIB_Cipsdynamiccryptomapsettable_Cipsdynamiccryptomapsetentry struct {
    EntityData types.CommonEntityData
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

func (cipsdynamiccryptomapsetentry *CISCOIPSECMIB_Cipsdynamiccryptomapsettable_Cipsdynamiccryptomapsetentry) GetEntityData() *types.CommonEntityData {
    cipsdynamiccryptomapsetentry.EntityData.YFilter = cipsdynamiccryptomapsetentry.YFilter
    cipsdynamiccryptomapsetentry.EntityData.YangName = "cipsDynamicCryptomapSetEntry"
    cipsdynamiccryptomapsetentry.EntityData.BundleName = "cisco_ios_xe"
    cipsdynamiccryptomapsetentry.EntityData.ParentYangName = "cipsDynamicCryptomapSetTable"
    cipsdynamiccryptomapsetentry.EntityData.SegmentPath = "cipsDynamicCryptomapSetEntry" + "[cipsDynamicCryptomapSetName='" + fmt.Sprintf("%v", cipsdynamiccryptomapsetentry.Cipsdynamiccryptomapsetname) + "']"
    cipsdynamiccryptomapsetentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsdynamiccryptomapsetentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsdynamiccryptomapsetentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsdynamiccryptomapsetentry.EntityData.Children = make(map[string]types.YChild)
    cipsdynamiccryptomapsetentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cipsdynamiccryptomapsetentry.EntityData.Leafs["cipsDynamicCryptomapSetName"] = types.YLeaf{"Cipsdynamiccryptomapsetname", cipsdynamiccryptomapsetentry.Cipsdynamiccryptomapsetname}
    cipsdynamiccryptomapsetentry.EntityData.Leafs["cipsDynamicCryptomapSetSize"] = types.YLeaf{"Cipsdynamiccryptomapsetsize", cipsdynamiccryptomapsetentry.Cipsdynamiccryptomapsetsize}
    cipsdynamiccryptomapsetentry.EntityData.Leafs["cipsDynamicCryptomapSetNumAssoc"] = types.YLeaf{"Cipsdynamiccryptomapsetnumassoc", cipsdynamiccryptomapsetentry.Cipsdynamiccryptomapsetnumassoc}
    return &(cipsdynamiccryptomapsetentry.EntityData)
}

// CISCOIPSECMIB_Cipsstaticcryptomaptable
// The table ilisting the member cryptomaps
// of the cryptomap sets that are configured
// on the managed entity.
type CISCOIPSECMIB_Cipsstaticcryptomaptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes  associated with a single static  (fully
    // specified) cryptomap entry. This table does not include the members  of
    // dynamic cryptomap sets that may be linked with the parent static cryptomap
    // set. The type is slice of
    // CISCOIPSECMIB_Cipsstaticcryptomaptable_Cipsstaticcryptomapentry.
    Cipsstaticcryptomapentry []CISCOIPSECMIB_Cipsstaticcryptomaptable_Cipsstaticcryptomapentry
}

func (cipsstaticcryptomaptable *CISCOIPSECMIB_Cipsstaticcryptomaptable) GetEntityData() *types.CommonEntityData {
    cipsstaticcryptomaptable.EntityData.YFilter = cipsstaticcryptomaptable.YFilter
    cipsstaticcryptomaptable.EntityData.YangName = "cipsStaticCryptomapTable"
    cipsstaticcryptomaptable.EntityData.BundleName = "cisco_ios_xe"
    cipsstaticcryptomaptable.EntityData.ParentYangName = "CISCO-IPSEC-MIB"
    cipsstaticcryptomaptable.EntityData.SegmentPath = "cipsStaticCryptomapTable"
    cipsstaticcryptomaptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsstaticcryptomaptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsstaticcryptomaptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsstaticcryptomaptable.EntityData.Children = make(map[string]types.YChild)
    cipsstaticcryptomaptable.EntityData.Children["cipsStaticCryptomapEntry"] = types.YChild{"Cipsstaticcryptomapentry", nil}
    for i := range cipsstaticcryptomaptable.Cipsstaticcryptomapentry {
        cipsstaticcryptomaptable.EntityData.Children[types.GetSegmentPath(&cipsstaticcryptomaptable.Cipsstaticcryptomapentry[i])] = types.YChild{"Cipsstaticcryptomapentry", &cipsstaticcryptomaptable.Cipsstaticcryptomapentry[i]}
    }
    cipsstaticcryptomaptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cipsstaticcryptomaptable.EntityData)
}

// CISCOIPSECMIB_Cipsstaticcryptomaptable_Cipsstaticcryptomapentry
// Each entry contains the attributes 
// associated with a single static 
// (fully specified) cryptomap entry.
// This table does not include the members 
// of dynamic cryptomap sets that may be
// linked with the parent static cryptomap set.
type CISCOIPSECMIB_Cipsstaticcryptomaptable_Cipsstaticcryptomapentry struct {
    EntityData types.CommonEntityData
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

func (cipsstaticcryptomapentry *CISCOIPSECMIB_Cipsstaticcryptomaptable_Cipsstaticcryptomapentry) GetEntityData() *types.CommonEntityData {
    cipsstaticcryptomapentry.EntityData.YFilter = cipsstaticcryptomapentry.YFilter
    cipsstaticcryptomapentry.EntityData.YangName = "cipsStaticCryptomapEntry"
    cipsstaticcryptomapentry.EntityData.BundleName = "cisco_ios_xe"
    cipsstaticcryptomapentry.EntityData.ParentYangName = "cipsStaticCryptomapTable"
    cipsstaticcryptomapentry.EntityData.SegmentPath = "cipsStaticCryptomapEntry" + "[cipsStaticCryptomapSetName='" + fmt.Sprintf("%v", cipsstaticcryptomapentry.Cipsstaticcryptomapsetname) + "']" + "[cipsStaticCryptomapPriority='" + fmt.Sprintf("%v", cipsstaticcryptomapentry.Cipsstaticcryptomappriority) + "']"
    cipsstaticcryptomapentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsstaticcryptomapentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsstaticcryptomapentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsstaticcryptomapentry.EntityData.Children = make(map[string]types.YChild)
    cipsstaticcryptomapentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cipsstaticcryptomapentry.EntityData.Leafs["cipsStaticCryptomapSetName"] = types.YLeaf{"Cipsstaticcryptomapsetname", cipsstaticcryptomapentry.Cipsstaticcryptomapsetname}
    cipsstaticcryptomapentry.EntityData.Leafs["cipsStaticCryptomapPriority"] = types.YLeaf{"Cipsstaticcryptomappriority", cipsstaticcryptomapentry.Cipsstaticcryptomappriority}
    cipsstaticcryptomapentry.EntityData.Leafs["cipsStaticCryptomapType"] = types.YLeaf{"Cipsstaticcryptomaptype", cipsstaticcryptomapentry.Cipsstaticcryptomaptype}
    cipsstaticcryptomapentry.EntityData.Leafs["cipsStaticCryptomapDescr"] = types.YLeaf{"Cipsstaticcryptomapdescr", cipsstaticcryptomapentry.Cipsstaticcryptomapdescr}
    cipsstaticcryptomapentry.EntityData.Leafs["cipsStaticCryptomapPeer"] = types.YLeaf{"Cipsstaticcryptomappeer", cipsstaticcryptomapentry.Cipsstaticcryptomappeer}
    cipsstaticcryptomapentry.EntityData.Leafs["cipsStaticCryptomapNumPeers"] = types.YLeaf{"Cipsstaticcryptomapnumpeers", cipsstaticcryptomapentry.Cipsstaticcryptomapnumpeers}
    cipsstaticcryptomapentry.EntityData.Leafs["cipsStaticCryptomapPfs"] = types.YLeaf{"Cipsstaticcryptomappfs", cipsstaticcryptomapentry.Cipsstaticcryptomappfs}
    cipsstaticcryptomapentry.EntityData.Leafs["cipsStaticCryptomapLifetime"] = types.YLeaf{"Cipsstaticcryptomaplifetime", cipsstaticcryptomapentry.Cipsstaticcryptomaplifetime}
    cipsstaticcryptomapentry.EntityData.Leafs["cipsStaticCryptomapLifesize"] = types.YLeaf{"Cipsstaticcryptomaplifesize", cipsstaticcryptomapentry.Cipsstaticcryptomaplifesize}
    cipsstaticcryptomapentry.EntityData.Leafs["cipsStaticCryptomapLevelHost"] = types.YLeaf{"Cipsstaticcryptomaplevelhost", cipsstaticcryptomapentry.Cipsstaticcryptomaplevelhost}
    return &(cipsstaticcryptomapentry.EntityData)
}

// CISCOIPSECMIB_Cipscryptomapsetiftable
// The table lists the binding of cryptomap sets
// to the interfaces of the managed entity.
type CISCOIPSECMIB_Cipscryptomapsetiftable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the record of the association between an interface and
    // a cryptomap set (static) that is defined on the managed entity.  Note that
    // the cryptomap set identified in  this binding must static. Dynamic
    // cryptomaps cannot be bound to interfaces. The type is slice of
    // CISCOIPSECMIB_Cipscryptomapsetiftable_Cipscryptomapsetifentry.
    Cipscryptomapsetifentry []CISCOIPSECMIB_Cipscryptomapsetiftable_Cipscryptomapsetifentry
}

func (cipscryptomapsetiftable *CISCOIPSECMIB_Cipscryptomapsetiftable) GetEntityData() *types.CommonEntityData {
    cipscryptomapsetiftable.EntityData.YFilter = cipscryptomapsetiftable.YFilter
    cipscryptomapsetiftable.EntityData.YangName = "cipsCryptomapSetIfTable"
    cipscryptomapsetiftable.EntityData.BundleName = "cisco_ios_xe"
    cipscryptomapsetiftable.EntityData.ParentYangName = "CISCO-IPSEC-MIB"
    cipscryptomapsetiftable.EntityData.SegmentPath = "cipsCryptomapSetIfTable"
    cipscryptomapsetiftable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipscryptomapsetiftable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipscryptomapsetiftable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipscryptomapsetiftable.EntityData.Children = make(map[string]types.YChild)
    cipscryptomapsetiftable.EntityData.Children["cipsCryptomapSetIfEntry"] = types.YChild{"Cipscryptomapsetifentry", nil}
    for i := range cipscryptomapsetiftable.Cipscryptomapsetifentry {
        cipscryptomapsetiftable.EntityData.Children[types.GetSegmentPath(&cipscryptomapsetiftable.Cipscryptomapsetifentry[i])] = types.YChild{"Cipscryptomapsetifentry", &cipscryptomapsetiftable.Cipscryptomapsetifentry[i]}
    }
    cipscryptomapsetiftable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cipscryptomapsetiftable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (cipscryptomapsetifentry *CISCOIPSECMIB_Cipscryptomapsetiftable_Cipscryptomapsetifentry) GetEntityData() *types.CommonEntityData {
    cipscryptomapsetifentry.EntityData.YFilter = cipscryptomapsetifentry.YFilter
    cipscryptomapsetifentry.EntityData.YangName = "cipsCryptomapSetIfEntry"
    cipscryptomapsetifentry.EntityData.BundleName = "cisco_ios_xe"
    cipscryptomapsetifentry.EntityData.ParentYangName = "cipsCryptomapSetIfTable"
    cipscryptomapsetifentry.EntityData.SegmentPath = "cipsCryptomapSetIfEntry" + "[ifIndex='" + fmt.Sprintf("%v", cipscryptomapsetifentry.Ifindex) + "']" + "[cipsStaticCryptomapSetName='" + fmt.Sprintf("%v", cipscryptomapsetifentry.Cipsstaticcryptomapsetname) + "']"
    cipscryptomapsetifentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipscryptomapsetifentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipscryptomapsetifentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipscryptomapsetifentry.EntityData.Children = make(map[string]types.YChild)
    cipscryptomapsetifentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cipscryptomapsetifentry.EntityData.Leafs["ifIndex"] = types.YLeaf{"Ifindex", cipscryptomapsetifentry.Ifindex}
    cipscryptomapsetifentry.EntityData.Leafs["cipsStaticCryptomapSetName"] = types.YLeaf{"Cipsstaticcryptomapsetname", cipscryptomapsetifentry.Cipsstaticcryptomapsetname}
    cipscryptomapsetifentry.EntityData.Leafs["cipsCryptomapSetIfVirtual"] = types.YLeaf{"Cipscryptomapsetifvirtual", cipscryptomapsetifentry.Cipscryptomapsetifvirtual}
    cipscryptomapsetifentry.EntityData.Leafs["cipsCryptomapSetIfStatus"] = types.YLeaf{"Cipscryptomapsetifstatus", cipscryptomapsetifentry.Cipscryptomapsetifstatus}
    return &(cipscryptomapsetifentry.EntityData)
}

