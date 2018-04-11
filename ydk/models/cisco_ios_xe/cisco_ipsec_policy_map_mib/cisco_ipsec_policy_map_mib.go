// The MIB module maps the IPSec
// entities created dynamically to the policy entities
// that caused them. This is an appendix to the
// IPSEC-MONITOR-MIB that has been proposed to
// IETF for monitoring IPSec based Virtual Private 
// Networks.
// 
// Overview of Cisco IPsec Policy Map MIB
// 
// MIB description
// 
// There are two components to this MIB:  
//   #1 a table that maps an IPSec Phase-1 
//      tunnel to the Internet Security Association 
//      and Key Exchange (ISAKMP) Policy 
//      
//   and 
// 
//   #2 a table that maps an IPSec Phase-2 
//      tunnel to the corresponding IPSec Policy
//      element - called 'cryptomaps' - in IOS 
//      (Internet Operating System)
// 
// The first mappin (also called Internet Key Exchange
//  or IKE mapping) yields, given the index of
// the IKE tunnel in the ikeTunnelTable
// (IPSEC-MONITOR-MIB), the ISAKMP policy definition
// defined using the CLI on the managed entity.
// 
// The IPSec mapping yields, given the index
// of the IPSec tunnel in the ipSecTunnelTable
// (IPSEC-MONITOR-MIB), the IPSec transform and
// the cryptomap definition that gave rise to
// this tunnel.
// 
// In implementation and usage, this MIB cannot
// exist independent of the IPSEC-MONITOR-MIB. 
package cisco_ipsec_policy_map_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ipsec_policy_map_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IPSEC-POLICY-MAP-MIB CISCO-IPSEC-POLICY-MAP-MIB}", reflect.TypeOf(CISCOIPSECPOLICYMAPMIB{}))
    ydk.RegisterEntity("CISCO-IPSEC-POLICY-MAP-MIB:CISCO-IPSEC-POLICY-MAP-MIB", reflect.TypeOf(CISCOIPSECPOLICYMAPMIB{}))
}

// CISCOIPSECPOLICYMAPMIB
type CISCOIPSECPOLICYMAPMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The IPSec Phase-1 Internet Key Exchange Tunnel to Policy Mapping Table.
    // There is one entry in this table for each active IPSec Phase-1 Tunnel.
    Ikepolmaptable CISCOIPSECPOLICYMAPMIB_Ikepolmaptable

    // The IPSec Phase-2 Tunnel to Policy Mapping Table. There is one entry in
    // this table for each active IPSec Phase-2 Tunnel.
    Ipsecpolmaptable CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable
}

func (cISCOIPSECPOLICYMAPMIB *CISCOIPSECPOLICYMAPMIB) GetEntityData() *types.CommonEntityData {
    cISCOIPSECPOLICYMAPMIB.EntityData.YFilter = cISCOIPSECPOLICYMAPMIB.YFilter
    cISCOIPSECPOLICYMAPMIB.EntityData.YangName = "CISCO-IPSEC-POLICY-MAP-MIB"
    cISCOIPSECPOLICYMAPMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIPSECPOLICYMAPMIB.EntityData.ParentYangName = "CISCO-IPSEC-POLICY-MAP-MIB"
    cISCOIPSECPOLICYMAPMIB.EntityData.SegmentPath = "CISCO-IPSEC-POLICY-MAP-MIB:CISCO-IPSEC-POLICY-MAP-MIB"
    cISCOIPSECPOLICYMAPMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIPSECPOLICYMAPMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIPSECPOLICYMAPMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIPSECPOLICYMAPMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOIPSECPOLICYMAPMIB.EntityData.Children["ikePolMapTable"] = types.YChild{"Ikepolmaptable", &cISCOIPSECPOLICYMAPMIB.Ikepolmaptable}
    cISCOIPSECPOLICYMAPMIB.EntityData.Children["ipSecPolMapTable"] = types.YChild{"Ipsecpolmaptable", &cISCOIPSECPOLICYMAPMIB.Ipsecpolmaptable}
    cISCOIPSECPOLICYMAPMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOIPSECPOLICYMAPMIB.EntityData)
}

// CISCOIPSECPOLICYMAPMIB_Ikepolmaptable
// The IPSec Phase-1 Internet Key Exchange Tunnel
// to Policy Mapping Table. There is one entry in
// this table for each active IPSec Phase-1
// Tunnel.
type CISCOIPSECPOLICYMAPMIB_Ikepolmaptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with mapping an active IPSec
    // Phase-1 IKE Tunnel to it's configured Policy definition. The type is slice
    // of CISCOIPSECPOLICYMAPMIB_Ikepolmaptable_Ikepolmapentry.
    Ikepolmapentry []CISCOIPSECPOLICYMAPMIB_Ikepolmaptable_Ikepolmapentry
}

func (ikepolmaptable *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable) GetEntityData() *types.CommonEntityData {
    ikepolmaptable.EntityData.YFilter = ikepolmaptable.YFilter
    ikepolmaptable.EntityData.YangName = "ikePolMapTable"
    ikepolmaptable.EntityData.BundleName = "cisco_ios_xe"
    ikepolmaptable.EntityData.ParentYangName = "CISCO-IPSEC-POLICY-MAP-MIB"
    ikepolmaptable.EntityData.SegmentPath = "ikePolMapTable"
    ikepolmaptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ikepolmaptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ikepolmaptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ikepolmaptable.EntityData.Children = make(map[string]types.YChild)
    ikepolmaptable.EntityData.Children["ikePolMapEntry"] = types.YChild{"Ikepolmapentry", nil}
    for i := range ikepolmaptable.Ikepolmapentry {
        ikepolmaptable.EntityData.Children[types.GetSegmentPath(&ikepolmaptable.Ikepolmapentry[i])] = types.YChild{"Ikepolmapentry", &ikepolmaptable.Ikepolmapentry[i]}
    }
    ikepolmaptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ikepolmaptable.EntityData)
}

// CISCOIPSECPOLICYMAPMIB_Ikepolmaptable_Ikepolmapentry
// Each entry contains the attributes associated
// with mapping an active IPSec Phase-1 IKE Tunnel
// to it's configured Policy definition.
type CISCOIPSECPOLICYMAPMIB_Ikepolmaptable_Ikepolmapentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index of the IPSec Phase-1 Tunnel to Policy
    // Map Table.  The value of the index is the number used to represent this
    // IPSec Phase-1 Tunnel in the IPSec MIB (ikeTunIndex in the ikeTunnelTable).
    // The type is interface{} with range: 1..2147483647.
    Ikepolmaptunindex interface{}

    // The number of the locally defined ISAKMP policy used to establish the IPSec
    // IKE Phase-1 Tunnel. This is the number which was used on the crypto
    // command. For example, if the configuration command was:   ==>  crypto
    // isakmp policy 15  then the value of this object would be 15. If ISAKMP was
    // not used to establish this tunnel, then the value of this object will be
    // zero. The type is interface{} with range: 1..2147483647.
    Ikepolmappolicynum interface{}
}

func (ikepolmapentry *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable_Ikepolmapentry) GetEntityData() *types.CommonEntityData {
    ikepolmapentry.EntityData.YFilter = ikepolmapentry.YFilter
    ikepolmapentry.EntityData.YangName = "ikePolMapEntry"
    ikepolmapentry.EntityData.BundleName = "cisco_ios_xe"
    ikepolmapentry.EntityData.ParentYangName = "ikePolMapTable"
    ikepolmapentry.EntityData.SegmentPath = "ikePolMapEntry" + "[ikePolMapTunIndex='" + fmt.Sprintf("%v", ikepolmapentry.Ikepolmaptunindex) + "']"
    ikepolmapentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ikepolmapentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ikepolmapentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ikepolmapentry.EntityData.Children = make(map[string]types.YChild)
    ikepolmapentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ikepolmapentry.EntityData.Leafs["ikePolMapTunIndex"] = types.YLeaf{"Ikepolmaptunindex", ikepolmapentry.Ikepolmaptunindex}
    ikepolmapentry.EntityData.Leafs["ikePolMapPolicyNum"] = types.YLeaf{"Ikepolmappolicynum", ikepolmapentry.Ikepolmappolicynum}
    return &(ikepolmapentry.EntityData)
}

// CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable
// The IPSec Phase-2 Tunnel to Policy Mapping Table.
// There is one entry in this table for each active
// IPSec Phase-2 Tunnel.
type CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with mapping an active IPSec
    // Phase-2 Tunnel to its configured Policy definition. The type is slice of
    // CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable_Ipsecpolmapentry.
    Ipsecpolmapentry []CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable_Ipsecpolmapentry
}

func (ipsecpolmaptable *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable) GetEntityData() *types.CommonEntityData {
    ipsecpolmaptable.EntityData.YFilter = ipsecpolmaptable.YFilter
    ipsecpolmaptable.EntityData.YangName = "ipSecPolMapTable"
    ipsecpolmaptable.EntityData.BundleName = "cisco_ios_xe"
    ipsecpolmaptable.EntityData.ParentYangName = "CISCO-IPSEC-POLICY-MAP-MIB"
    ipsecpolmaptable.EntityData.SegmentPath = "ipSecPolMapTable"
    ipsecpolmaptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipsecpolmaptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipsecpolmaptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipsecpolmaptable.EntityData.Children = make(map[string]types.YChild)
    ipsecpolmaptable.EntityData.Children["ipSecPolMapEntry"] = types.YChild{"Ipsecpolmapentry", nil}
    for i := range ipsecpolmaptable.Ipsecpolmapentry {
        ipsecpolmaptable.EntityData.Children[types.GetSegmentPath(&ipsecpolmaptable.Ipsecpolmapentry[i])] = types.YChild{"Ipsecpolmapentry", &ipsecpolmaptable.Ipsecpolmapentry[i]}
    }
    ipsecpolmaptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ipsecpolmaptable.EntityData)
}

// CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable_Ipsecpolmapentry
// Each entry contains the attributes associated
// with mapping an active IPSec Phase-2 Tunnel
// to its configured Policy definition.
type CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable_Ipsecpolmapentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index of the IPSec Phase-2 Tunnel to Policy
    // Map Table. The value of the index is the number used to represent this
    // IPSec Phase-2 Tunnel in the IPSec MIB (ipSecTunIndex in the
    // ipSecTunnelTable). The type is interface{} with range: 1..2147483647.
    Ipsecpolmaptunindex interface{}

    // The value of this object should be the name of  the IPSec Policy
    // (cryptomap) as assigned by the  operator while configuring the policy of 
    // the IPSec traffic.  For instance, on an IOS router, the if the command
    // entered to configure the IPSec policy was   ==>  crypto map ftpPolicy 10
    // ipsec-isakmp  then the value of this object would be 'ftpPolicy'. The type
    // is string.
    Ipsecpolmapcryptomapname interface{}

    // The value of this object should be the priority of the IPSec Policy
    // (cryptomap) assigned by the  operator while configuring the policy of  this
    // IPSec tunnel.  For instance, on an IOS router, the if the command entered
    // to configure the IPSec policy was   ==>  crypto map ftpPolicy 10
    // ipsec-isakmp  then the value of this object would be 10. The type is
    // interface{} with range: 1..2147483647.
    Ipsecpolmapcryptomapnum interface{}

    // The value of this object is the number or the name of the access control
    // string (ACL)  that caused this IPSec tunnel to be established.   The ACL
    // that causes an IPSec tunnel  to be established is referenced by the  
    // cryptomap of the tunnel.   The ACL identifies the traffic that requires 
    // protection as defined by the policy.   For instance, the ACL that requires
    // FTP  traffic between local subnet 172.16.14.0 and a  remote subnet
    // 172.16.16.0 to be protected  is defined as   ==>access-list 101 permit tcp
    // 172.16.14.0 0.0.0.255                   172.16.16.0 0.0.0.255 eq ftp   When
    // this command causes an IPSec tunnel to be   established, the object
    // 'ipSecPolMapAclString'   assumes the string value '101'.   If the ACL is a
    // named list such as   ==> ip access-list standard myAcl        permit
    // 172.16.16.8 0.0.0.0   then the value of this MIB element corresponding to  
    // IPSec tunnel that was created by this ACL would   be 'myAcl'. The type is
    // string.
    Ipsecpolmapaclstring interface{}

    // The value of this object is the access control  entry (ACE) within the ACL
    // that caused this IPSec  tunnel to be established.   For instance, if an ACL
    // defines access for two traffic streams (FTP and SNMP) as follows: 
    // access-list 101 permit tcp 172.16.14.0 0.0.0.255                 
    // 172.16.16.0 0.0.0.255 eq ftp access-list 101 permit udp 172.16.14.0
    // 0.0.0.255                  host 172.16.16.1 eq 161   When associated with
    // an IPSec policy, the second element of the ACL gives rise to an IPSec
    // tunnel in the wake of SNMP traffic. The value of the object
    // 'ipSecPolMapAceString' for the IPSec tunnel would be then the string
    // 'access-list 101 permit udp 172.16.14.0 0.0.0.255                  host
    // 172.16.16.1 eq 161'. The type is string.
    Ipsecpolmapacestring interface{}
}

func (ipsecpolmapentry *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable_Ipsecpolmapentry) GetEntityData() *types.CommonEntityData {
    ipsecpolmapentry.EntityData.YFilter = ipsecpolmapentry.YFilter
    ipsecpolmapentry.EntityData.YangName = "ipSecPolMapEntry"
    ipsecpolmapentry.EntityData.BundleName = "cisco_ios_xe"
    ipsecpolmapentry.EntityData.ParentYangName = "ipSecPolMapTable"
    ipsecpolmapentry.EntityData.SegmentPath = "ipSecPolMapEntry" + "[ipSecPolMapTunIndex='" + fmt.Sprintf("%v", ipsecpolmapentry.Ipsecpolmaptunindex) + "']"
    ipsecpolmapentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipsecpolmapentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipsecpolmapentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipsecpolmapentry.EntityData.Children = make(map[string]types.YChild)
    ipsecpolmapentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ipsecpolmapentry.EntityData.Leafs["ipSecPolMapTunIndex"] = types.YLeaf{"Ipsecpolmaptunindex", ipsecpolmapentry.Ipsecpolmaptunindex}
    ipsecpolmapentry.EntityData.Leafs["ipSecPolMapCryptoMapName"] = types.YLeaf{"Ipsecpolmapcryptomapname", ipsecpolmapentry.Ipsecpolmapcryptomapname}
    ipsecpolmapentry.EntityData.Leafs["ipSecPolMapCryptoMapNum"] = types.YLeaf{"Ipsecpolmapcryptomapnum", ipsecpolmapentry.Ipsecpolmapcryptomapnum}
    ipsecpolmapentry.EntityData.Leafs["ipSecPolMapAclString"] = types.YLeaf{"Ipsecpolmapaclstring", ipsecpolmapentry.Ipsecpolmapaclstring}
    ipsecpolmapentry.EntityData.Leafs["ipSecPolMapAceString"] = types.YLeaf{"Ipsecpolmapacestring", ipsecpolmapentry.Ipsecpolmapacestring}
    return &(ipsecpolmapentry.EntityData)
}

