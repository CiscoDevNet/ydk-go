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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CipsIsakmpGroup CISCOIPSECMIB_CipsIsakmpGroup

    
    CipsIPsecGlobals CISCOIPSECMIB_CipsIPsecGlobals

    
    CipsIPsecStatistics CISCOIPSECMIB_CipsIPsecStatistics

    
    CipsSysCapacityGroup CISCOIPSECMIB_CipsSysCapacityGroup

    
    CipsTrapCntlGroup CISCOIPSECMIB_CipsTrapCntlGroup

    // The table containing the list of all ISAKMP policy entries configured by
    // the operator.
    CipsIsakmpPolicyTable CISCOIPSECMIB_CipsIsakmpPolicyTable

    // The table containing the list of all cryptomap sets that are fully
    // specified and are not wild-carded.  The operator may include different
    // types of cryptomaps in such a set - manual, CET, ISAKMP or dynamic.
    CipsStaticCryptomapSetTable CISCOIPSECMIB_CipsStaticCryptomapSetTable

    // The table containing the list of all dynamic cryptomaps that use IKE,
    // defined on   the managed entity.
    CipsDynamicCryptomapSetTable CISCOIPSECMIB_CipsDynamicCryptomapSetTable

    // The table ilisting the member cryptomaps of the cryptomap sets that are
    // configured on the managed entity.
    CipsStaticCryptomapTable CISCOIPSECMIB_CipsStaticCryptomapTable

    // The table lists the binding of cryptomap sets to the interfaces of the
    // managed entity.
    CipsCryptomapSetIfTable CISCOIPSECMIB_CipsCryptomapSetIfTable
}

func (cISCOIPSECMIB *CISCOIPSECMIB) GetEntityData() *types.CommonEntityData {
    cISCOIPSECMIB.EntityData.YFilter = cISCOIPSECMIB.YFilter
    cISCOIPSECMIB.EntityData.YangName = "CISCO-IPSEC-MIB"
    cISCOIPSECMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIPSECMIB.EntityData.ParentYangName = "CISCO-IPSEC-MIB"
    cISCOIPSECMIB.EntityData.SegmentPath = "CISCO-IPSEC-MIB:CISCO-IPSEC-MIB"
    cISCOIPSECMIB.EntityData.AbsolutePath = cISCOIPSECMIB.EntityData.SegmentPath
    cISCOIPSECMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIPSECMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIPSECMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIPSECMIB.EntityData.Children = types.NewOrderedMap()
    cISCOIPSECMIB.EntityData.Children.Append("cipsIsakmpGroup", types.YChild{"CipsIsakmpGroup", &cISCOIPSECMIB.CipsIsakmpGroup})
    cISCOIPSECMIB.EntityData.Children.Append("cipsIPsecGlobals", types.YChild{"CipsIPsecGlobals", &cISCOIPSECMIB.CipsIPsecGlobals})
    cISCOIPSECMIB.EntityData.Children.Append("cipsIPsecStatistics", types.YChild{"CipsIPsecStatistics", &cISCOIPSECMIB.CipsIPsecStatistics})
    cISCOIPSECMIB.EntityData.Children.Append("cipsSysCapacityGroup", types.YChild{"CipsSysCapacityGroup", &cISCOIPSECMIB.CipsSysCapacityGroup})
    cISCOIPSECMIB.EntityData.Children.Append("cipsTrapCntlGroup", types.YChild{"CipsTrapCntlGroup", &cISCOIPSECMIB.CipsTrapCntlGroup})
    cISCOIPSECMIB.EntityData.Children.Append("cipsIsakmpPolicyTable", types.YChild{"CipsIsakmpPolicyTable", &cISCOIPSECMIB.CipsIsakmpPolicyTable})
    cISCOIPSECMIB.EntityData.Children.Append("cipsStaticCryptomapSetTable", types.YChild{"CipsStaticCryptomapSetTable", &cISCOIPSECMIB.CipsStaticCryptomapSetTable})
    cISCOIPSECMIB.EntityData.Children.Append("cipsDynamicCryptomapSetTable", types.YChild{"CipsDynamicCryptomapSetTable", &cISCOIPSECMIB.CipsDynamicCryptomapSetTable})
    cISCOIPSECMIB.EntityData.Children.Append("cipsStaticCryptomapTable", types.YChild{"CipsStaticCryptomapTable", &cISCOIPSECMIB.CipsStaticCryptomapTable})
    cISCOIPSECMIB.EntityData.Children.Append("cipsCryptomapSetIfTable", types.YChild{"CipsCryptomapSetIfTable", &cISCOIPSECMIB.CipsCryptomapSetIfTable})
    cISCOIPSECMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOIPSECMIB.EntityData.YListKeys = []string {}

    return &(cISCOIPSECMIB.EntityData)
}

// CISCOIPSECMIB_CipsIsakmpGroup
type CISCOIPSECMIB_CipsIsakmpGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The value of this object is TRUE if ISAKMP has been enabled on the managed
    // entity. Otherise the value of this object is FALSE. The type is bool.
    CipsIsakmpEnabled interface{}

    // The value of this object is shows the type of identity used by the managed
    // entity in ISAKMP negotiations with another peer. The type is
    // IkeIdentityType.
    CipsIsakmpIdentity interface{}

    // The value of this object is time interval in seconds between successive
    // ISAKMP keepalive heartbeats issued to the peers to which IKE tunnels have
    // been setup. The type is interface{} with range: 10..3600. Units are
    // seconds.
    CipsIsakmpKeepaliveInterval interface{}

    // The value of this object is the number of ISAKMP policies that have been
    // configured on the  managed entity. The type is interface{} with range:
    // 0..2147483647.
    CipsNumIsakmpPolicies interface{}
}

func (cipsIsakmpGroup *CISCOIPSECMIB_CipsIsakmpGroup) GetEntityData() *types.CommonEntityData {
    cipsIsakmpGroup.EntityData.YFilter = cipsIsakmpGroup.YFilter
    cipsIsakmpGroup.EntityData.YangName = "cipsIsakmpGroup"
    cipsIsakmpGroup.EntityData.BundleName = "cisco_ios_xe"
    cipsIsakmpGroup.EntityData.ParentYangName = "CISCO-IPSEC-MIB"
    cipsIsakmpGroup.EntityData.SegmentPath = "cipsIsakmpGroup"
    cipsIsakmpGroup.EntityData.AbsolutePath = "CISCO-IPSEC-MIB:CISCO-IPSEC-MIB/" + cipsIsakmpGroup.EntityData.SegmentPath
    cipsIsakmpGroup.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsIsakmpGroup.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsIsakmpGroup.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsIsakmpGroup.EntityData.Children = types.NewOrderedMap()
    cipsIsakmpGroup.EntityData.Leafs = types.NewOrderedMap()
    cipsIsakmpGroup.EntityData.Leafs.Append("cipsIsakmpEnabled", types.YLeaf{"CipsIsakmpEnabled", cipsIsakmpGroup.CipsIsakmpEnabled})
    cipsIsakmpGroup.EntityData.Leafs.Append("cipsIsakmpIdentity", types.YLeaf{"CipsIsakmpIdentity", cipsIsakmpGroup.CipsIsakmpIdentity})
    cipsIsakmpGroup.EntityData.Leafs.Append("cipsIsakmpKeepaliveInterval", types.YLeaf{"CipsIsakmpKeepaliveInterval", cipsIsakmpGroup.CipsIsakmpKeepaliveInterval})
    cipsIsakmpGroup.EntityData.Leafs.Append("cipsNumIsakmpPolicies", types.YLeaf{"CipsNumIsakmpPolicies", cipsIsakmpGroup.CipsNumIsakmpPolicies})

    cipsIsakmpGroup.EntityData.YListKeys = []string {}

    return &(cipsIsakmpGroup.EntityData)
}

// CISCOIPSECMIB_CipsIPsecGlobals
type CISCOIPSECMIB_CipsIPsecGlobals struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The default lifetime (in seconds) assigned  to an SA as a global policy
    // (maybe overridden  in specific cryptomap definitions). The type is
    // interface{} with range: 120..86400. Units are Seconds.
    CipsSALifetime interface{}

    // The default lifesize in KBytes assigned to an SA  as a global policy
    // (unless overridden in cryptomap  definition). The type is interface{} with
    // range: 2560..536870912. Units are KBytes.
    CipsSALifesize interface{}

    // The number of Cryptomap Sets that are are fully configured. Statically
    // defined cryptomap sets  are ones where the operator has fully specified all
    // the parameters required set up IPSec  Virtual Private Networks (VPNs). The
    // type is interface{} with range: 0..2147483647. Units are Integral Units.
    CipsNumStaticCryptomapSets interface{}

    // The number of static Cryptomap Sets that have  at least one CET cryptomap
    // element as a member of the set. The type is interface{} with range:
    // 0..2147483647. Units are Integral Units.
    CipsNumCETCryptomapSets interface{}

    // The number of dynamic IPSec Policy templates (called 'dynamic cryptomap
    // templates') configured on the managed entity. The type is interface{} with
    // range: 0..2147483647. Units are Integral Units.
    CipsNumDynamicCryptomapSets interface{}

    // The number of static Cryptomap Sets that have  at least one dynamic
    // cryptomap template  bound to them which has the Tunnel Endpoint Discovery
    // (TED) enabled. The type is interface{} with range: 0..2147483647. Units are
    // Integral Units.
    CipsNumTEDCryptomapSets interface{}
}

func (cipsIPsecGlobals *CISCOIPSECMIB_CipsIPsecGlobals) GetEntityData() *types.CommonEntityData {
    cipsIPsecGlobals.EntityData.YFilter = cipsIPsecGlobals.YFilter
    cipsIPsecGlobals.EntityData.YangName = "cipsIPsecGlobals"
    cipsIPsecGlobals.EntityData.BundleName = "cisco_ios_xe"
    cipsIPsecGlobals.EntityData.ParentYangName = "CISCO-IPSEC-MIB"
    cipsIPsecGlobals.EntityData.SegmentPath = "cipsIPsecGlobals"
    cipsIPsecGlobals.EntityData.AbsolutePath = "CISCO-IPSEC-MIB:CISCO-IPSEC-MIB/" + cipsIPsecGlobals.EntityData.SegmentPath
    cipsIPsecGlobals.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsIPsecGlobals.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsIPsecGlobals.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsIPsecGlobals.EntityData.Children = types.NewOrderedMap()
    cipsIPsecGlobals.EntityData.Leafs = types.NewOrderedMap()
    cipsIPsecGlobals.EntityData.Leafs.Append("cipsSALifetime", types.YLeaf{"CipsSALifetime", cipsIPsecGlobals.CipsSALifetime})
    cipsIPsecGlobals.EntityData.Leafs.Append("cipsSALifesize", types.YLeaf{"CipsSALifesize", cipsIPsecGlobals.CipsSALifesize})
    cipsIPsecGlobals.EntityData.Leafs.Append("cipsNumStaticCryptomapSets", types.YLeaf{"CipsNumStaticCryptomapSets", cipsIPsecGlobals.CipsNumStaticCryptomapSets})
    cipsIPsecGlobals.EntityData.Leafs.Append("cipsNumCETCryptomapSets", types.YLeaf{"CipsNumCETCryptomapSets", cipsIPsecGlobals.CipsNumCETCryptomapSets})
    cipsIPsecGlobals.EntityData.Leafs.Append("cipsNumDynamicCryptomapSets", types.YLeaf{"CipsNumDynamicCryptomapSets", cipsIPsecGlobals.CipsNumDynamicCryptomapSets})
    cipsIPsecGlobals.EntityData.Leafs.Append("cipsNumTEDCryptomapSets", types.YLeaf{"CipsNumTEDCryptomapSets", cipsIPsecGlobals.CipsNumTEDCryptomapSets})

    cipsIPsecGlobals.EntityData.YListKeys = []string {}

    return &(cipsIPsecGlobals.EntityData)
}

// CISCOIPSECMIB_CipsIPsecStatistics
type CISCOIPSECMIB_CipsIPsecStatistics struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The number of TED probes that were received by this  managed entity since
    // bootup. Not affected by any  CLI operation. The type is interface{} with
    // range: 0..4294967295. Units are Integral Units.
    CipsNumTEDProbesReceived interface{}

    // The number of TED probes that were dispatched by all the dynamic cryptomaps
    // in this managed entity since  bootup. Not affected by any CLI operation.
    // The type is interface{} with range: 0..4294967295. Units are Integral
    // Units.
    CipsNumTEDProbesSent interface{}

    // The number of TED probes that were dispatched by  the local entity and that
    // failed to locate crypto  endpoint.  Not affected by any CLI operation. The
    // type is interface{} with range: 0..4294967295. Units are Integral Units.
    CipsNumTEDFailures interface{}
}

func (cipsIPsecStatistics *CISCOIPSECMIB_CipsIPsecStatistics) GetEntityData() *types.CommonEntityData {
    cipsIPsecStatistics.EntityData.YFilter = cipsIPsecStatistics.YFilter
    cipsIPsecStatistics.EntityData.YangName = "cipsIPsecStatistics"
    cipsIPsecStatistics.EntityData.BundleName = "cisco_ios_xe"
    cipsIPsecStatistics.EntityData.ParentYangName = "CISCO-IPSEC-MIB"
    cipsIPsecStatistics.EntityData.SegmentPath = "cipsIPsecStatistics"
    cipsIPsecStatistics.EntityData.AbsolutePath = "CISCO-IPSEC-MIB:CISCO-IPSEC-MIB/" + cipsIPsecStatistics.EntityData.SegmentPath
    cipsIPsecStatistics.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsIPsecStatistics.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsIPsecStatistics.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsIPsecStatistics.EntityData.Children = types.NewOrderedMap()
    cipsIPsecStatistics.EntityData.Leafs = types.NewOrderedMap()
    cipsIPsecStatistics.EntityData.Leafs.Append("cipsNumTEDProbesReceived", types.YLeaf{"CipsNumTEDProbesReceived", cipsIPsecStatistics.CipsNumTEDProbesReceived})
    cipsIPsecStatistics.EntityData.Leafs.Append("cipsNumTEDProbesSent", types.YLeaf{"CipsNumTEDProbesSent", cipsIPsecStatistics.CipsNumTEDProbesSent})
    cipsIPsecStatistics.EntityData.Leafs.Append("cipsNumTEDFailures", types.YLeaf{"CipsNumTEDFailures", cipsIPsecStatistics.CipsNumTEDFailures})

    cipsIPsecStatistics.EntityData.YListKeys = []string {}

    return &(cipsIPsecStatistics.EntityData)
}

// CISCOIPSECMIB_CipsSysCapacityGroup
type CISCOIPSECMIB_CipsSysCapacityGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The maximum number of IPsec Security Associations that can be established
    // on this managed entity. If no theoretical limit exists, this returns value
    // 0.  Not affected by any CLI operation. The type is interface{} with range:
    // 0..65535. Units are Integral Units.
    CipsMaxSAs interface{}

    // The value of this object is TRUE if the  managed entity has the hardware
    // nad software  features to support 3DES encryption algorithm.  Not affected
    // by any CLI operation. The type is bool.
    Cips3DesCapable interface{}
}

func (cipsSysCapacityGroup *CISCOIPSECMIB_CipsSysCapacityGroup) GetEntityData() *types.CommonEntityData {
    cipsSysCapacityGroup.EntityData.YFilter = cipsSysCapacityGroup.YFilter
    cipsSysCapacityGroup.EntityData.YangName = "cipsSysCapacityGroup"
    cipsSysCapacityGroup.EntityData.BundleName = "cisco_ios_xe"
    cipsSysCapacityGroup.EntityData.ParentYangName = "CISCO-IPSEC-MIB"
    cipsSysCapacityGroup.EntityData.SegmentPath = "cipsSysCapacityGroup"
    cipsSysCapacityGroup.EntityData.AbsolutePath = "CISCO-IPSEC-MIB:CISCO-IPSEC-MIB/" + cipsSysCapacityGroup.EntityData.SegmentPath
    cipsSysCapacityGroup.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsSysCapacityGroup.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsSysCapacityGroup.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsSysCapacityGroup.EntityData.Children = types.NewOrderedMap()
    cipsSysCapacityGroup.EntityData.Leafs = types.NewOrderedMap()
    cipsSysCapacityGroup.EntityData.Leafs.Append("cipsMaxSAs", types.YLeaf{"CipsMaxSAs", cipsSysCapacityGroup.CipsMaxSAs})
    cipsSysCapacityGroup.EntityData.Leafs.Append("cips3DesCapable", types.YLeaf{"Cips3DesCapable", cipsSysCapacityGroup.Cips3DesCapable})

    cipsSysCapacityGroup.EntityData.YListKeys = []string {}

    return &(cipsSysCapacityGroup.EntityData)
}

// CISCOIPSECMIB_CipsTrapCntlGroup
type CISCOIPSECMIB_CipsTrapCntlGroup struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object defines the administrative state of  sending the IOS IPsec
    // ISAKMP Policy Add trap. The type is TrapStatus.
    CipsCntlIsakmpPolicyAdded interface{}

    // This object defines the administrative state of  sending the IOS IPsec
    // ISAKMP Policy Delete trap. The type is TrapStatus.
    CipsCntlIsakmpPolicyDeleted interface{}

    // This object defines the administrative state of  sending the IOS IPsec
    // Cryptomap Add trap. The type is TrapStatus.
    CipsCntlCryptomapAdded interface{}

    // This object defines the administrative state of  sending the IOS IPsec
    // Cryptomap Delete trap. The type is TrapStatus.
    CipsCntlCryptomapDeleted interface{}

    // This object defines the administrative state of  sending the IOS IPsec trap
    // that is issued when a cryptomap set is attached to an interface. The type
    // is TrapStatus.
    CipsCntlCryptomapSetAttached interface{}

    // This object defines the administrative state of  sending the IOS IPsec trap
    // that is issued when a cryptomap set is detached from an interface. to which
    // it was earlier bound. The type is TrapStatus.
    CipsCntlCryptomapSetDetached interface{}

    // This object defines the administrative state of  sending the IOS IPsec trap
    // that is issued when the number of SAs crosses the maximum number of SAs
    // that may be supported on the managed entity. The type is TrapStatus.
    CipsCntlTooManySAs interface{}
}

func (cipsTrapCntlGroup *CISCOIPSECMIB_CipsTrapCntlGroup) GetEntityData() *types.CommonEntityData {
    cipsTrapCntlGroup.EntityData.YFilter = cipsTrapCntlGroup.YFilter
    cipsTrapCntlGroup.EntityData.YangName = "cipsTrapCntlGroup"
    cipsTrapCntlGroup.EntityData.BundleName = "cisco_ios_xe"
    cipsTrapCntlGroup.EntityData.ParentYangName = "CISCO-IPSEC-MIB"
    cipsTrapCntlGroup.EntityData.SegmentPath = "cipsTrapCntlGroup"
    cipsTrapCntlGroup.EntityData.AbsolutePath = "CISCO-IPSEC-MIB:CISCO-IPSEC-MIB/" + cipsTrapCntlGroup.EntityData.SegmentPath
    cipsTrapCntlGroup.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsTrapCntlGroup.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsTrapCntlGroup.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsTrapCntlGroup.EntityData.Children = types.NewOrderedMap()
    cipsTrapCntlGroup.EntityData.Leafs = types.NewOrderedMap()
    cipsTrapCntlGroup.EntityData.Leafs.Append("cipsCntlIsakmpPolicyAdded", types.YLeaf{"CipsCntlIsakmpPolicyAdded", cipsTrapCntlGroup.CipsCntlIsakmpPolicyAdded})
    cipsTrapCntlGroup.EntityData.Leafs.Append("cipsCntlIsakmpPolicyDeleted", types.YLeaf{"CipsCntlIsakmpPolicyDeleted", cipsTrapCntlGroup.CipsCntlIsakmpPolicyDeleted})
    cipsTrapCntlGroup.EntityData.Leafs.Append("cipsCntlCryptomapAdded", types.YLeaf{"CipsCntlCryptomapAdded", cipsTrapCntlGroup.CipsCntlCryptomapAdded})
    cipsTrapCntlGroup.EntityData.Leafs.Append("cipsCntlCryptomapDeleted", types.YLeaf{"CipsCntlCryptomapDeleted", cipsTrapCntlGroup.CipsCntlCryptomapDeleted})
    cipsTrapCntlGroup.EntityData.Leafs.Append("cipsCntlCryptomapSetAttached", types.YLeaf{"CipsCntlCryptomapSetAttached", cipsTrapCntlGroup.CipsCntlCryptomapSetAttached})
    cipsTrapCntlGroup.EntityData.Leafs.Append("cipsCntlCryptomapSetDetached", types.YLeaf{"CipsCntlCryptomapSetDetached", cipsTrapCntlGroup.CipsCntlCryptomapSetDetached})
    cipsTrapCntlGroup.EntityData.Leafs.Append("cipsCntlTooManySAs", types.YLeaf{"CipsCntlTooManySAs", cipsTrapCntlGroup.CipsCntlTooManySAs})

    cipsTrapCntlGroup.EntityData.YListKeys = []string {}

    return &(cipsTrapCntlGroup.EntityData)
}

// CISCOIPSECMIB_CipsIsakmpPolicyTable
// The table containing the list of all
// ISAKMP policy entries configured by the operator.
type CISCOIPSECMIB_CipsIsakmpPolicyTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes  associated with a single ISAKMP Policy
    // entry. The type is slice of
    // CISCOIPSECMIB_CipsIsakmpPolicyTable_CipsIsakmpPolicyEntry.
    CipsIsakmpPolicyEntry []*CISCOIPSECMIB_CipsIsakmpPolicyTable_CipsIsakmpPolicyEntry
}

func (cipsIsakmpPolicyTable *CISCOIPSECMIB_CipsIsakmpPolicyTable) GetEntityData() *types.CommonEntityData {
    cipsIsakmpPolicyTable.EntityData.YFilter = cipsIsakmpPolicyTable.YFilter
    cipsIsakmpPolicyTable.EntityData.YangName = "cipsIsakmpPolicyTable"
    cipsIsakmpPolicyTable.EntityData.BundleName = "cisco_ios_xe"
    cipsIsakmpPolicyTable.EntityData.ParentYangName = "CISCO-IPSEC-MIB"
    cipsIsakmpPolicyTable.EntityData.SegmentPath = "cipsIsakmpPolicyTable"
    cipsIsakmpPolicyTable.EntityData.AbsolutePath = "CISCO-IPSEC-MIB:CISCO-IPSEC-MIB/" + cipsIsakmpPolicyTable.EntityData.SegmentPath
    cipsIsakmpPolicyTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsIsakmpPolicyTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsIsakmpPolicyTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsIsakmpPolicyTable.EntityData.Children = types.NewOrderedMap()
    cipsIsakmpPolicyTable.EntityData.Children.Append("cipsIsakmpPolicyEntry", types.YChild{"CipsIsakmpPolicyEntry", nil})
    for i := range cipsIsakmpPolicyTable.CipsIsakmpPolicyEntry {
        cipsIsakmpPolicyTable.EntityData.Children.Append(types.GetSegmentPath(cipsIsakmpPolicyTable.CipsIsakmpPolicyEntry[i]), types.YChild{"CipsIsakmpPolicyEntry", cipsIsakmpPolicyTable.CipsIsakmpPolicyEntry[i]})
    }
    cipsIsakmpPolicyTable.EntityData.Leafs = types.NewOrderedMap()

    cipsIsakmpPolicyTable.EntityData.YListKeys = []string {}

    return &(cipsIsakmpPolicyTable.EntityData)
}

// CISCOIPSECMIB_CipsIsakmpPolicyTable_CipsIsakmpPolicyEntry
// Each entry contains the attributes 
// associated with a single ISAKMP
// Policy entry.
type CISCOIPSECMIB_CipsIsakmpPolicyTable_CipsIsakmpPolicyEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The priotity of this ISAKMP Policy entry. This is
    // also the index of this table. The type is interface{} with range: 0..65535.
    CipsIsakmpPolPriority interface{}

    // The encryption transform specified by this  ISAKMP policy specification.
    // The Internet Key Exchange (IKE) tunnels setup using this policy item would
    // use the specified encryption transform to protect the ISAKMP PDUs. The type
    // is EncryptAlgo.
    CipsIsakmpPolEncr interface{}

    // The hash transform specified by this  ISAKMP policy specification. The IKE
    // tunnels setup using this policy item would use the  specified hash
    // transform to protect the ISAKMP PDUs. The type is IkeHashAlgo.
    CipsIsakmpPolHash interface{}

    // The peer authentication mthod specified by this ISAKMP policy
    // specification. If this policy entity is selected for negotiation with a
    // peer, the local entity would authenticate the peer using  the method
    // specified by this object. The type is IkeAuthMethod.
    CipsIsakmpPolAuth interface{}

    // This object specifies the Oakley group used  for Diffie Hellman exchange in
    // the Main Mode.  If this policy item is selected to negotiate Main Mode with
    // an IKE peer, the local entity  chooses the group specified by this object
    // to perform Diffie Hellman exchange with the peer. The type is
    // DiffHellmanGrp.
    CipsIsakmpPolGroup interface{}

    // This object specifies the lifetime in seconds of the IKE tunnels generated
    // using this  policy specification. The type is interface{} with range:
    // 60..86400. Units are seconds.
    CipsIsakmpPolLifetime interface{}
}

func (cipsIsakmpPolicyEntry *CISCOIPSECMIB_CipsIsakmpPolicyTable_CipsIsakmpPolicyEntry) GetEntityData() *types.CommonEntityData {
    cipsIsakmpPolicyEntry.EntityData.YFilter = cipsIsakmpPolicyEntry.YFilter
    cipsIsakmpPolicyEntry.EntityData.YangName = "cipsIsakmpPolicyEntry"
    cipsIsakmpPolicyEntry.EntityData.BundleName = "cisco_ios_xe"
    cipsIsakmpPolicyEntry.EntityData.ParentYangName = "cipsIsakmpPolicyTable"
    cipsIsakmpPolicyEntry.EntityData.SegmentPath = "cipsIsakmpPolicyEntry" + types.AddKeyToken(cipsIsakmpPolicyEntry.CipsIsakmpPolPriority, "cipsIsakmpPolPriority")
    cipsIsakmpPolicyEntry.EntityData.AbsolutePath = "CISCO-IPSEC-MIB:CISCO-IPSEC-MIB/cipsIsakmpPolicyTable/" + cipsIsakmpPolicyEntry.EntityData.SegmentPath
    cipsIsakmpPolicyEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsIsakmpPolicyEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsIsakmpPolicyEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsIsakmpPolicyEntry.EntityData.Children = types.NewOrderedMap()
    cipsIsakmpPolicyEntry.EntityData.Leafs = types.NewOrderedMap()
    cipsIsakmpPolicyEntry.EntityData.Leafs.Append("cipsIsakmpPolPriority", types.YLeaf{"CipsIsakmpPolPriority", cipsIsakmpPolicyEntry.CipsIsakmpPolPriority})
    cipsIsakmpPolicyEntry.EntityData.Leafs.Append("cipsIsakmpPolEncr", types.YLeaf{"CipsIsakmpPolEncr", cipsIsakmpPolicyEntry.CipsIsakmpPolEncr})
    cipsIsakmpPolicyEntry.EntityData.Leafs.Append("cipsIsakmpPolHash", types.YLeaf{"CipsIsakmpPolHash", cipsIsakmpPolicyEntry.CipsIsakmpPolHash})
    cipsIsakmpPolicyEntry.EntityData.Leafs.Append("cipsIsakmpPolAuth", types.YLeaf{"CipsIsakmpPolAuth", cipsIsakmpPolicyEntry.CipsIsakmpPolAuth})
    cipsIsakmpPolicyEntry.EntityData.Leafs.Append("cipsIsakmpPolGroup", types.YLeaf{"CipsIsakmpPolGroup", cipsIsakmpPolicyEntry.CipsIsakmpPolGroup})
    cipsIsakmpPolicyEntry.EntityData.Leafs.Append("cipsIsakmpPolLifetime", types.YLeaf{"CipsIsakmpPolLifetime", cipsIsakmpPolicyEntry.CipsIsakmpPolLifetime})

    cipsIsakmpPolicyEntry.EntityData.YListKeys = []string {"CipsIsakmpPolPriority"}

    return &(cipsIsakmpPolicyEntry.EntityData)
}

// CISCOIPSECMIB_CipsStaticCryptomapSetTable
// The table containing the list of all
// cryptomap sets that are fully specified
// and are not wild-carded.
// 
// The operator may include different types of
// cryptomaps in such a set - manual, CET,
// ISAKMP or dynamic.
type CISCOIPSECMIB_CipsStaticCryptomapSetTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes  associated with a single static 
    // cryptomap set. The type is slice of
    // CISCOIPSECMIB_CipsStaticCryptomapSetTable_CipsStaticCryptomapSetEntry.
    CipsStaticCryptomapSetEntry []*CISCOIPSECMIB_CipsStaticCryptomapSetTable_CipsStaticCryptomapSetEntry
}

func (cipsStaticCryptomapSetTable *CISCOIPSECMIB_CipsStaticCryptomapSetTable) GetEntityData() *types.CommonEntityData {
    cipsStaticCryptomapSetTable.EntityData.YFilter = cipsStaticCryptomapSetTable.YFilter
    cipsStaticCryptomapSetTable.EntityData.YangName = "cipsStaticCryptomapSetTable"
    cipsStaticCryptomapSetTable.EntityData.BundleName = "cisco_ios_xe"
    cipsStaticCryptomapSetTable.EntityData.ParentYangName = "CISCO-IPSEC-MIB"
    cipsStaticCryptomapSetTable.EntityData.SegmentPath = "cipsStaticCryptomapSetTable"
    cipsStaticCryptomapSetTable.EntityData.AbsolutePath = "CISCO-IPSEC-MIB:CISCO-IPSEC-MIB/" + cipsStaticCryptomapSetTable.EntityData.SegmentPath
    cipsStaticCryptomapSetTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsStaticCryptomapSetTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsStaticCryptomapSetTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsStaticCryptomapSetTable.EntityData.Children = types.NewOrderedMap()
    cipsStaticCryptomapSetTable.EntityData.Children.Append("cipsStaticCryptomapSetEntry", types.YChild{"CipsStaticCryptomapSetEntry", nil})
    for i := range cipsStaticCryptomapSetTable.CipsStaticCryptomapSetEntry {
        cipsStaticCryptomapSetTable.EntityData.Children.Append(types.GetSegmentPath(cipsStaticCryptomapSetTable.CipsStaticCryptomapSetEntry[i]), types.YChild{"CipsStaticCryptomapSetEntry", cipsStaticCryptomapSetTable.CipsStaticCryptomapSetEntry[i]})
    }
    cipsStaticCryptomapSetTable.EntityData.Leafs = types.NewOrderedMap()

    cipsStaticCryptomapSetTable.EntityData.YListKeys = []string {}

    return &(cipsStaticCryptomapSetTable.EntityData)
}

// CISCOIPSECMIB_CipsStaticCryptomapSetTable_CipsStaticCryptomapSetEntry
// Each entry contains the attributes 
// associated with a single static 
// cryptomap set.
type CISCOIPSECMIB_CipsStaticCryptomapSetTable_CipsStaticCryptomapSetEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index of the static cryptomap table. The value
    // of the string is the name string assigned by the  operator in defining the
    // cryptomap set. The type is string.
    CipsStaticCryptomapSetName interface{}

    // The total number of cryptomap entries contained in this cryptomap set. .
    // The type is interface{} with range: 0..4294967295.
    CipsStaticCryptomapSetSize interface{}

    // The number of cryptomaps associated with this  cryptomap set that use
    // ISAKMP protocol to do key exchange. The type is interface{} with range:
    // 0..4294967295.
    CipsStaticCryptomapSetNumIsakmp interface{}

    // The number of cryptomaps associated with this  cryptomap set that require
    // the operator to manually setup the keys and SPIs. The type is interface{}
    // with range: 0..4294967295.
    CipsStaticCryptomapSetNumManual interface{}

    // The number of cryptomaps of type 'ipsec-cisco'  associated with this
    // cryptomap set. Such cryptomap elements implement Cisco Encryption
    // Technology based Virtual Private Networks. The type is interface{} with
    // range: 0..4294967295.
    CipsStaticCryptomapSetNumCET interface{}

    // The number of dynamic cryptomap templates linked to this cryptomap set. The
    // type is interface{} with range: 0..4294967295.
    CipsStaticCryptomapSetNumDynamic interface{}

    // The number of dynamic cryptomap templates linked to this cryptomap set that
    // have Tunnel Endpoint Discovery (TED) enabled. The type is interface{} with
    // range: 0..4294967295.
    CipsStaticCryptomapSetNumDisc interface{}

    // The number of and IPsec Security Associations that are active and were
    // setup using this cryptomap.  . The type is interface{} with range:
    // 0..4294967295.
    CipsStaticCryptomapSetNumSAs interface{}
}

func (cipsStaticCryptomapSetEntry *CISCOIPSECMIB_CipsStaticCryptomapSetTable_CipsStaticCryptomapSetEntry) GetEntityData() *types.CommonEntityData {
    cipsStaticCryptomapSetEntry.EntityData.YFilter = cipsStaticCryptomapSetEntry.YFilter
    cipsStaticCryptomapSetEntry.EntityData.YangName = "cipsStaticCryptomapSetEntry"
    cipsStaticCryptomapSetEntry.EntityData.BundleName = "cisco_ios_xe"
    cipsStaticCryptomapSetEntry.EntityData.ParentYangName = "cipsStaticCryptomapSetTable"
    cipsStaticCryptomapSetEntry.EntityData.SegmentPath = "cipsStaticCryptomapSetEntry" + types.AddKeyToken(cipsStaticCryptomapSetEntry.CipsStaticCryptomapSetName, "cipsStaticCryptomapSetName")
    cipsStaticCryptomapSetEntry.EntityData.AbsolutePath = "CISCO-IPSEC-MIB:CISCO-IPSEC-MIB/cipsStaticCryptomapSetTable/" + cipsStaticCryptomapSetEntry.EntityData.SegmentPath
    cipsStaticCryptomapSetEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsStaticCryptomapSetEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsStaticCryptomapSetEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsStaticCryptomapSetEntry.EntityData.Children = types.NewOrderedMap()
    cipsStaticCryptomapSetEntry.EntityData.Leafs = types.NewOrderedMap()
    cipsStaticCryptomapSetEntry.EntityData.Leafs.Append("cipsStaticCryptomapSetName", types.YLeaf{"CipsStaticCryptomapSetName", cipsStaticCryptomapSetEntry.CipsStaticCryptomapSetName})
    cipsStaticCryptomapSetEntry.EntityData.Leafs.Append("cipsStaticCryptomapSetSize", types.YLeaf{"CipsStaticCryptomapSetSize", cipsStaticCryptomapSetEntry.CipsStaticCryptomapSetSize})
    cipsStaticCryptomapSetEntry.EntityData.Leafs.Append("cipsStaticCryptomapSetNumIsakmp", types.YLeaf{"CipsStaticCryptomapSetNumIsakmp", cipsStaticCryptomapSetEntry.CipsStaticCryptomapSetNumIsakmp})
    cipsStaticCryptomapSetEntry.EntityData.Leafs.Append("cipsStaticCryptomapSetNumManual", types.YLeaf{"CipsStaticCryptomapSetNumManual", cipsStaticCryptomapSetEntry.CipsStaticCryptomapSetNumManual})
    cipsStaticCryptomapSetEntry.EntityData.Leafs.Append("cipsStaticCryptomapSetNumCET", types.YLeaf{"CipsStaticCryptomapSetNumCET", cipsStaticCryptomapSetEntry.CipsStaticCryptomapSetNumCET})
    cipsStaticCryptomapSetEntry.EntityData.Leafs.Append("cipsStaticCryptomapSetNumDynamic", types.YLeaf{"CipsStaticCryptomapSetNumDynamic", cipsStaticCryptomapSetEntry.CipsStaticCryptomapSetNumDynamic})
    cipsStaticCryptomapSetEntry.EntityData.Leafs.Append("cipsStaticCryptomapSetNumDisc", types.YLeaf{"CipsStaticCryptomapSetNumDisc", cipsStaticCryptomapSetEntry.CipsStaticCryptomapSetNumDisc})
    cipsStaticCryptomapSetEntry.EntityData.Leafs.Append("cipsStaticCryptomapSetNumSAs", types.YLeaf{"CipsStaticCryptomapSetNumSAs", cipsStaticCryptomapSetEntry.CipsStaticCryptomapSetNumSAs})

    cipsStaticCryptomapSetEntry.EntityData.YListKeys = []string {"CipsStaticCryptomapSetName"}

    return &(cipsStaticCryptomapSetEntry.EntityData)
}

// CISCOIPSECMIB_CipsDynamicCryptomapSetTable
// The table containing the list of all dynamic
// cryptomaps that use IKE, defined on 
//  the managed entity.
type CISCOIPSECMIB_CipsDynamicCryptomapSetTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with a single dynamic
    // cryptomap template. The type is slice of
    // CISCOIPSECMIB_CipsDynamicCryptomapSetTable_CipsDynamicCryptomapSetEntry.
    CipsDynamicCryptomapSetEntry []*CISCOIPSECMIB_CipsDynamicCryptomapSetTable_CipsDynamicCryptomapSetEntry
}

func (cipsDynamicCryptomapSetTable *CISCOIPSECMIB_CipsDynamicCryptomapSetTable) GetEntityData() *types.CommonEntityData {
    cipsDynamicCryptomapSetTable.EntityData.YFilter = cipsDynamicCryptomapSetTable.YFilter
    cipsDynamicCryptomapSetTable.EntityData.YangName = "cipsDynamicCryptomapSetTable"
    cipsDynamicCryptomapSetTable.EntityData.BundleName = "cisco_ios_xe"
    cipsDynamicCryptomapSetTable.EntityData.ParentYangName = "CISCO-IPSEC-MIB"
    cipsDynamicCryptomapSetTable.EntityData.SegmentPath = "cipsDynamicCryptomapSetTable"
    cipsDynamicCryptomapSetTable.EntityData.AbsolutePath = "CISCO-IPSEC-MIB:CISCO-IPSEC-MIB/" + cipsDynamicCryptomapSetTable.EntityData.SegmentPath
    cipsDynamicCryptomapSetTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsDynamicCryptomapSetTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsDynamicCryptomapSetTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsDynamicCryptomapSetTable.EntityData.Children = types.NewOrderedMap()
    cipsDynamicCryptomapSetTable.EntityData.Children.Append("cipsDynamicCryptomapSetEntry", types.YChild{"CipsDynamicCryptomapSetEntry", nil})
    for i := range cipsDynamicCryptomapSetTable.CipsDynamicCryptomapSetEntry {
        cipsDynamicCryptomapSetTable.EntityData.Children.Append(types.GetSegmentPath(cipsDynamicCryptomapSetTable.CipsDynamicCryptomapSetEntry[i]), types.YChild{"CipsDynamicCryptomapSetEntry", cipsDynamicCryptomapSetTable.CipsDynamicCryptomapSetEntry[i]})
    }
    cipsDynamicCryptomapSetTable.EntityData.Leafs = types.NewOrderedMap()

    cipsDynamicCryptomapSetTable.EntityData.YListKeys = []string {}

    return &(cipsDynamicCryptomapSetTable.EntityData)
}

// CISCOIPSECMIB_CipsDynamicCryptomapSetTable_CipsDynamicCryptomapSetEntry
// Each entry contains the attributes associated
// with a single dynamic cryptomap template.
type CISCOIPSECMIB_CipsDynamicCryptomapSetTable_CipsDynamicCryptomapSetEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The index of the dynamic cryptomap table.  The
    // value of the string is the one assigned  by the operator in defining the
    // cryptomap set. The type is string.
    CipsDynamicCryptomapSetName interface{}

    // The number of cryptomap entries in this cryptomap. The type is interface{}
    // with range: 0..4294967295.
    CipsDynamicCryptomapSetSize interface{}

    // The number of static cryptomap sets with which this dynamic cryptomap is
    // associated.  . The type is interface{} with range: 0..4294967295.
    CipsDynamicCryptomapSetNumAssoc interface{}
}

func (cipsDynamicCryptomapSetEntry *CISCOIPSECMIB_CipsDynamicCryptomapSetTable_CipsDynamicCryptomapSetEntry) GetEntityData() *types.CommonEntityData {
    cipsDynamicCryptomapSetEntry.EntityData.YFilter = cipsDynamicCryptomapSetEntry.YFilter
    cipsDynamicCryptomapSetEntry.EntityData.YangName = "cipsDynamicCryptomapSetEntry"
    cipsDynamicCryptomapSetEntry.EntityData.BundleName = "cisco_ios_xe"
    cipsDynamicCryptomapSetEntry.EntityData.ParentYangName = "cipsDynamicCryptomapSetTable"
    cipsDynamicCryptomapSetEntry.EntityData.SegmentPath = "cipsDynamicCryptomapSetEntry" + types.AddKeyToken(cipsDynamicCryptomapSetEntry.CipsDynamicCryptomapSetName, "cipsDynamicCryptomapSetName")
    cipsDynamicCryptomapSetEntry.EntityData.AbsolutePath = "CISCO-IPSEC-MIB:CISCO-IPSEC-MIB/cipsDynamicCryptomapSetTable/" + cipsDynamicCryptomapSetEntry.EntityData.SegmentPath
    cipsDynamicCryptomapSetEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsDynamicCryptomapSetEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsDynamicCryptomapSetEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsDynamicCryptomapSetEntry.EntityData.Children = types.NewOrderedMap()
    cipsDynamicCryptomapSetEntry.EntityData.Leafs = types.NewOrderedMap()
    cipsDynamicCryptomapSetEntry.EntityData.Leafs.Append("cipsDynamicCryptomapSetName", types.YLeaf{"CipsDynamicCryptomapSetName", cipsDynamicCryptomapSetEntry.CipsDynamicCryptomapSetName})
    cipsDynamicCryptomapSetEntry.EntityData.Leafs.Append("cipsDynamicCryptomapSetSize", types.YLeaf{"CipsDynamicCryptomapSetSize", cipsDynamicCryptomapSetEntry.CipsDynamicCryptomapSetSize})
    cipsDynamicCryptomapSetEntry.EntityData.Leafs.Append("cipsDynamicCryptomapSetNumAssoc", types.YLeaf{"CipsDynamicCryptomapSetNumAssoc", cipsDynamicCryptomapSetEntry.CipsDynamicCryptomapSetNumAssoc})

    cipsDynamicCryptomapSetEntry.EntityData.YListKeys = []string {"CipsDynamicCryptomapSetName"}

    return &(cipsDynamicCryptomapSetEntry.EntityData)
}

// CISCOIPSECMIB_CipsStaticCryptomapTable
// The table ilisting the member cryptomaps
// of the cryptomap sets that are configured
// on the managed entity.
type CISCOIPSECMIB_CipsStaticCryptomapTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes  associated with a single static  (fully
    // specified) cryptomap entry. This table does not include the members  of
    // dynamic cryptomap sets that may be linked with the parent static cryptomap
    // set. The type is slice of
    // CISCOIPSECMIB_CipsStaticCryptomapTable_CipsStaticCryptomapEntry.
    CipsStaticCryptomapEntry []*CISCOIPSECMIB_CipsStaticCryptomapTable_CipsStaticCryptomapEntry
}

func (cipsStaticCryptomapTable *CISCOIPSECMIB_CipsStaticCryptomapTable) GetEntityData() *types.CommonEntityData {
    cipsStaticCryptomapTable.EntityData.YFilter = cipsStaticCryptomapTable.YFilter
    cipsStaticCryptomapTable.EntityData.YangName = "cipsStaticCryptomapTable"
    cipsStaticCryptomapTable.EntityData.BundleName = "cisco_ios_xe"
    cipsStaticCryptomapTable.EntityData.ParentYangName = "CISCO-IPSEC-MIB"
    cipsStaticCryptomapTable.EntityData.SegmentPath = "cipsStaticCryptomapTable"
    cipsStaticCryptomapTable.EntityData.AbsolutePath = "CISCO-IPSEC-MIB:CISCO-IPSEC-MIB/" + cipsStaticCryptomapTable.EntityData.SegmentPath
    cipsStaticCryptomapTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsStaticCryptomapTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsStaticCryptomapTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsStaticCryptomapTable.EntityData.Children = types.NewOrderedMap()
    cipsStaticCryptomapTable.EntityData.Children.Append("cipsStaticCryptomapEntry", types.YChild{"CipsStaticCryptomapEntry", nil})
    for i := range cipsStaticCryptomapTable.CipsStaticCryptomapEntry {
        cipsStaticCryptomapTable.EntityData.Children.Append(types.GetSegmentPath(cipsStaticCryptomapTable.CipsStaticCryptomapEntry[i]), types.YChild{"CipsStaticCryptomapEntry", cipsStaticCryptomapTable.CipsStaticCryptomapEntry[i]})
    }
    cipsStaticCryptomapTable.EntityData.Leafs = types.NewOrderedMap()

    cipsStaticCryptomapTable.EntityData.YListKeys = []string {}

    return &(cipsStaticCryptomapTable.EntityData)
}

// CISCOIPSECMIB_CipsStaticCryptomapTable_CipsStaticCryptomapEntry
// Each entry contains the attributes 
// associated with a single static 
// (fully specified) cryptomap entry.
// This table does not include the members 
// of dynamic cryptomap sets that may be
// linked with the parent static cryptomap set.
type CISCOIPSECMIB_CipsStaticCryptomapTable_CipsStaticCryptomapEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string. Refers to
    // cisco_ipsec_mib.CISCOIPSECMIB_CipsStaticCryptomapSetTable_CipsStaticCryptomapSetEntry_CipsStaticCryptomapSetName
    CipsStaticCryptomapSetName interface{}

    // This attribute is a key. The priority of the cryptomap entry in the 
    // cryptomap set. This is the second index component of this table. The type
    // is interface{} with range: 0..65535.
    CipsStaticCryptomapPriority interface{}

    // The type of the cryptomap entry. This can be an ISAKMP cryptomap, CET or
    // manual. Dynamic cryptomaps are not counted in this table. The type is
    // CryptomapType.
    CipsStaticCryptomapType interface{}

    // The description string entered by the operatoir while creating this
    // cryptomap. The string generally identifies a description and the purpose of
    // this policy. The type is string.
    CipsStaticCryptomapDescr interface{}

    // The IP address of the current peer associated with  this IPSec policy item.
    // Traffic that is protected by this cryptomap is protected by a tunnel that
    // terminates at the device whose IP address is specified by this object. The
    // type is string with length: 4 | 16.
    CipsStaticCryptomapPeer interface{}

    // The number of peers associated with this cryptomap  entry. The peers other
    // than the one identified by  'cipsStaticCryptomapPeer' are backup peers.  
    // Manual cryptomaps may have only one peer. The type is interface{} with
    // range: 0..40.
    CipsStaticCryptomapNumPeers interface{}

    // This object identifies if the tunnels instantiated due to this policy item
    // should use Perfect Forward Secrecy  (PFS) and if so, what group of Oakley
    // they should use. The type is DiffHellmanGrp.
    CipsStaticCryptomapPfs interface{}

    // This object identifies the lifetime of the IPSec Security Associations (SA)
    // created using this IPSec policy entry. If this value is zero, the lifetime
    // assumes the  value specified by the global lifetime parameter. The type is
    // interface{} with range: 0..None | 120..86400.
    CipsStaticCryptomapLifetime interface{}

    // This object identifies the lifesize (maximum traffic in bytes that may be
    // carried) of the IPSec SAs created using this IPSec policy entry.  If this
    // value is zero, the lifetime assumes the  value specified by the global
    // lifesize parameter. The type is interface{} with range: 0..None |
    // 2560..536870912.
    CipsStaticCryptomapLifesize interface{}

    // This object identifies the granularity of the IPSec SAs created using this
    // IPSec policy entry.  If this value is TRUE, distinct SA bundles are created
    // for distinct hosts at the end of the application traffic. The type is bool.
    CipsStaticCryptomapLevelHost interface{}
}

func (cipsStaticCryptomapEntry *CISCOIPSECMIB_CipsStaticCryptomapTable_CipsStaticCryptomapEntry) GetEntityData() *types.CommonEntityData {
    cipsStaticCryptomapEntry.EntityData.YFilter = cipsStaticCryptomapEntry.YFilter
    cipsStaticCryptomapEntry.EntityData.YangName = "cipsStaticCryptomapEntry"
    cipsStaticCryptomapEntry.EntityData.BundleName = "cisco_ios_xe"
    cipsStaticCryptomapEntry.EntityData.ParentYangName = "cipsStaticCryptomapTable"
    cipsStaticCryptomapEntry.EntityData.SegmentPath = "cipsStaticCryptomapEntry" + types.AddKeyToken(cipsStaticCryptomapEntry.CipsStaticCryptomapSetName, "cipsStaticCryptomapSetName") + types.AddKeyToken(cipsStaticCryptomapEntry.CipsStaticCryptomapPriority, "cipsStaticCryptomapPriority")
    cipsStaticCryptomapEntry.EntityData.AbsolutePath = "CISCO-IPSEC-MIB:CISCO-IPSEC-MIB/cipsStaticCryptomapTable/" + cipsStaticCryptomapEntry.EntityData.SegmentPath
    cipsStaticCryptomapEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsStaticCryptomapEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsStaticCryptomapEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsStaticCryptomapEntry.EntityData.Children = types.NewOrderedMap()
    cipsStaticCryptomapEntry.EntityData.Leafs = types.NewOrderedMap()
    cipsStaticCryptomapEntry.EntityData.Leafs.Append("cipsStaticCryptomapSetName", types.YLeaf{"CipsStaticCryptomapSetName", cipsStaticCryptomapEntry.CipsStaticCryptomapSetName})
    cipsStaticCryptomapEntry.EntityData.Leafs.Append("cipsStaticCryptomapPriority", types.YLeaf{"CipsStaticCryptomapPriority", cipsStaticCryptomapEntry.CipsStaticCryptomapPriority})
    cipsStaticCryptomapEntry.EntityData.Leafs.Append("cipsStaticCryptomapType", types.YLeaf{"CipsStaticCryptomapType", cipsStaticCryptomapEntry.CipsStaticCryptomapType})
    cipsStaticCryptomapEntry.EntityData.Leafs.Append("cipsStaticCryptomapDescr", types.YLeaf{"CipsStaticCryptomapDescr", cipsStaticCryptomapEntry.CipsStaticCryptomapDescr})
    cipsStaticCryptomapEntry.EntityData.Leafs.Append("cipsStaticCryptomapPeer", types.YLeaf{"CipsStaticCryptomapPeer", cipsStaticCryptomapEntry.CipsStaticCryptomapPeer})
    cipsStaticCryptomapEntry.EntityData.Leafs.Append("cipsStaticCryptomapNumPeers", types.YLeaf{"CipsStaticCryptomapNumPeers", cipsStaticCryptomapEntry.CipsStaticCryptomapNumPeers})
    cipsStaticCryptomapEntry.EntityData.Leafs.Append("cipsStaticCryptomapPfs", types.YLeaf{"CipsStaticCryptomapPfs", cipsStaticCryptomapEntry.CipsStaticCryptomapPfs})
    cipsStaticCryptomapEntry.EntityData.Leafs.Append("cipsStaticCryptomapLifetime", types.YLeaf{"CipsStaticCryptomapLifetime", cipsStaticCryptomapEntry.CipsStaticCryptomapLifetime})
    cipsStaticCryptomapEntry.EntityData.Leafs.Append("cipsStaticCryptomapLifesize", types.YLeaf{"CipsStaticCryptomapLifesize", cipsStaticCryptomapEntry.CipsStaticCryptomapLifesize})
    cipsStaticCryptomapEntry.EntityData.Leafs.Append("cipsStaticCryptomapLevelHost", types.YLeaf{"CipsStaticCryptomapLevelHost", cipsStaticCryptomapEntry.CipsStaticCryptomapLevelHost})

    cipsStaticCryptomapEntry.EntityData.YListKeys = []string {"CipsStaticCryptomapSetName", "CipsStaticCryptomapPriority"}

    return &(cipsStaticCryptomapEntry.EntityData)
}

// CISCOIPSECMIB_CipsCryptomapSetIfTable
// The table lists the binding of cryptomap sets
// to the interfaces of the managed entity.
type CISCOIPSECMIB_CipsCryptomapSetIfTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the record of the association between an interface and
    // a cryptomap set (static) that is defined on the managed entity.  Note that
    // the cryptomap set identified in  this binding must static. Dynamic
    // cryptomaps cannot be bound to interfaces. The type is slice of
    // CISCOIPSECMIB_CipsCryptomapSetIfTable_CipsCryptomapSetIfEntry.
    CipsCryptomapSetIfEntry []*CISCOIPSECMIB_CipsCryptomapSetIfTable_CipsCryptomapSetIfEntry
}

func (cipsCryptomapSetIfTable *CISCOIPSECMIB_CipsCryptomapSetIfTable) GetEntityData() *types.CommonEntityData {
    cipsCryptomapSetIfTable.EntityData.YFilter = cipsCryptomapSetIfTable.YFilter
    cipsCryptomapSetIfTable.EntityData.YangName = "cipsCryptomapSetIfTable"
    cipsCryptomapSetIfTable.EntityData.BundleName = "cisco_ios_xe"
    cipsCryptomapSetIfTable.EntityData.ParentYangName = "CISCO-IPSEC-MIB"
    cipsCryptomapSetIfTable.EntityData.SegmentPath = "cipsCryptomapSetIfTable"
    cipsCryptomapSetIfTable.EntityData.AbsolutePath = "CISCO-IPSEC-MIB:CISCO-IPSEC-MIB/" + cipsCryptomapSetIfTable.EntityData.SegmentPath
    cipsCryptomapSetIfTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsCryptomapSetIfTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsCryptomapSetIfTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsCryptomapSetIfTable.EntityData.Children = types.NewOrderedMap()
    cipsCryptomapSetIfTable.EntityData.Children.Append("cipsCryptomapSetIfEntry", types.YChild{"CipsCryptomapSetIfEntry", nil})
    for i := range cipsCryptomapSetIfTable.CipsCryptomapSetIfEntry {
        cipsCryptomapSetIfTable.EntityData.Children.Append(types.GetSegmentPath(cipsCryptomapSetIfTable.CipsCryptomapSetIfEntry[i]), types.YChild{"CipsCryptomapSetIfEntry", cipsCryptomapSetIfTable.CipsCryptomapSetIfEntry[i]})
    }
    cipsCryptomapSetIfTable.EntityData.Leafs = types.NewOrderedMap()

    cipsCryptomapSetIfTable.EntityData.YListKeys = []string {}

    return &(cipsCryptomapSetIfTable.EntityData)
}

// CISCOIPSECMIB_CipsCryptomapSetIfTable_CipsCryptomapSetIfEntry
// Each entry contains the record of
// the association between an interface
// and a cryptomap set (static) that is defined
// on the managed entity.
// 
// Note that the cryptomap set identified in 
// this binding must static. Dynamic cryptomaps cannot
// be bound to interfaces.
type CISCOIPSECMIB_CipsCryptomapSetIfTable_CipsCryptomapSetIfEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to if_mib.IFMIB_IfTable_IfEntry_IfIndex
    IfIndex interface{}

    // This attribute is a key. The type is string. Refers to
    // cisco_ipsec_mib.CISCOIPSECMIB_CipsStaticCryptomapSetTable_CipsStaticCryptomapSetEntry_CipsStaticCryptomapSetName
    CipsStaticCryptomapSetName interface{}

    // The value of this object identifies if the interface to which the cryptomap
    // set is attached is a tunnel (such as a GRE or PPTP tunnel). The type is
    // bool.
    CipsCryptomapSetIfVirtual interface{}

    // This object identifies the status of the binding  of the specified
    // cryptomap set with the specified interface. The value when queried is
    // always 'attached'.  When set to 'detached', the cryptomap set if detached 
    // from the specified interface. The effect of this is same  as the CLI
    // command  	config-if# no crypto map cryptomapSetName  Setting the value to
    // 'attached' will result in  SNMP General Error. The type is
    // CryptomapSetBindStatus.
    CipsCryptomapSetIfStatus interface{}
}

func (cipsCryptomapSetIfEntry *CISCOIPSECMIB_CipsCryptomapSetIfTable_CipsCryptomapSetIfEntry) GetEntityData() *types.CommonEntityData {
    cipsCryptomapSetIfEntry.EntityData.YFilter = cipsCryptomapSetIfEntry.YFilter
    cipsCryptomapSetIfEntry.EntityData.YangName = "cipsCryptomapSetIfEntry"
    cipsCryptomapSetIfEntry.EntityData.BundleName = "cisco_ios_xe"
    cipsCryptomapSetIfEntry.EntityData.ParentYangName = "cipsCryptomapSetIfTable"
    cipsCryptomapSetIfEntry.EntityData.SegmentPath = "cipsCryptomapSetIfEntry" + types.AddKeyToken(cipsCryptomapSetIfEntry.IfIndex, "ifIndex") + types.AddKeyToken(cipsCryptomapSetIfEntry.CipsStaticCryptomapSetName, "cipsStaticCryptomapSetName")
    cipsCryptomapSetIfEntry.EntityData.AbsolutePath = "CISCO-IPSEC-MIB:CISCO-IPSEC-MIB/cipsCryptomapSetIfTable/" + cipsCryptomapSetIfEntry.EntityData.SegmentPath
    cipsCryptomapSetIfEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cipsCryptomapSetIfEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cipsCryptomapSetIfEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cipsCryptomapSetIfEntry.EntityData.Children = types.NewOrderedMap()
    cipsCryptomapSetIfEntry.EntityData.Leafs = types.NewOrderedMap()
    cipsCryptomapSetIfEntry.EntityData.Leafs.Append("ifIndex", types.YLeaf{"IfIndex", cipsCryptomapSetIfEntry.IfIndex})
    cipsCryptomapSetIfEntry.EntityData.Leafs.Append("cipsStaticCryptomapSetName", types.YLeaf{"CipsStaticCryptomapSetName", cipsCryptomapSetIfEntry.CipsStaticCryptomapSetName})
    cipsCryptomapSetIfEntry.EntityData.Leafs.Append("cipsCryptomapSetIfVirtual", types.YLeaf{"CipsCryptomapSetIfVirtual", cipsCryptomapSetIfEntry.CipsCryptomapSetIfVirtual})
    cipsCryptomapSetIfEntry.EntityData.Leafs.Append("cipsCryptomapSetIfStatus", types.YLeaf{"CipsCryptomapSetIfStatus", cipsCryptomapSetIfEntry.CipsCryptomapSetIfStatus})

    cipsCryptomapSetIfEntry.EntityData.YListKeys = []string {"IfIndex", "CipsStaticCryptomapSetName"}

    return &(cipsCryptomapSetIfEntry.EntityData)
}

