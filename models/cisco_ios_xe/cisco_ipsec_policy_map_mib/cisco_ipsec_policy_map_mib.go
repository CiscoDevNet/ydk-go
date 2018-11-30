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
    IkePolMapTable CISCOIPSECPOLICYMAPMIB_IkePolMapTable

    // The IPSec Phase-2 Tunnel to Policy Mapping Table. There is one entry in
    // this table for each active IPSec Phase-2 Tunnel.
    IpSecPolMapTable CISCOIPSECPOLICYMAPMIB_IpSecPolMapTable
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

    cISCOIPSECPOLICYMAPMIB.EntityData.Children = types.NewOrderedMap()
    cISCOIPSECPOLICYMAPMIB.EntityData.Children.Append("ikePolMapTable", types.YChild{"IkePolMapTable", &cISCOIPSECPOLICYMAPMIB.IkePolMapTable})
    cISCOIPSECPOLICYMAPMIB.EntityData.Children.Append("ipSecPolMapTable", types.YChild{"IpSecPolMapTable", &cISCOIPSECPOLICYMAPMIB.IpSecPolMapTable})
    cISCOIPSECPOLICYMAPMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOIPSECPOLICYMAPMIB.EntityData.YListKeys = []string {}

    return &(cISCOIPSECPOLICYMAPMIB.EntityData)
}

// CISCOIPSECPOLICYMAPMIB_IkePolMapTable
// The IPSec Phase-1 Internet Key Exchange Tunnel
// to Policy Mapping Table. There is one entry in
// this table for each active IPSec Phase-1
// Tunnel.
type CISCOIPSECPOLICYMAPMIB_IkePolMapTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with mapping an active IPSec
    // Phase-1 IKE Tunnel to it's configured Policy definition. The type is slice
    // of CISCOIPSECPOLICYMAPMIB_IkePolMapTable_IkePolMapEntry.
    IkePolMapEntry []*CISCOIPSECPOLICYMAPMIB_IkePolMapTable_IkePolMapEntry
}

func (ikePolMapTable *CISCOIPSECPOLICYMAPMIB_IkePolMapTable) GetEntityData() *types.CommonEntityData {
    ikePolMapTable.EntityData.YFilter = ikePolMapTable.YFilter
    ikePolMapTable.EntityData.YangName = "ikePolMapTable"
    ikePolMapTable.EntityData.BundleName = "cisco_ios_xe"
    ikePolMapTable.EntityData.ParentYangName = "CISCO-IPSEC-POLICY-MAP-MIB"
    ikePolMapTable.EntityData.SegmentPath = "ikePolMapTable"
    ikePolMapTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ikePolMapTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ikePolMapTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ikePolMapTable.EntityData.Children = types.NewOrderedMap()
    ikePolMapTable.EntityData.Children.Append("ikePolMapEntry", types.YChild{"IkePolMapEntry", nil})
    for i := range ikePolMapTable.IkePolMapEntry {
        ikePolMapTable.EntityData.Children.Append(types.GetSegmentPath(ikePolMapTable.IkePolMapEntry[i]), types.YChild{"IkePolMapEntry", ikePolMapTable.IkePolMapEntry[i]})
    }
    ikePolMapTable.EntityData.Leafs = types.NewOrderedMap()

    ikePolMapTable.EntityData.YListKeys = []string {}

    return &(ikePolMapTable.EntityData)
}

// CISCOIPSECPOLICYMAPMIB_IkePolMapTable_IkePolMapEntry
// Each entry contains the attributes associated
// with mapping an active IPSec Phase-1 IKE Tunnel
// to it's configured Policy definition.
type CISCOIPSECPOLICYMAPMIB_IkePolMapTable_IkePolMapEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index of the IPSec Phase-1 Tunnel to Policy
    // Map Table.  The value of the index is the number used to represent this
    // IPSec Phase-1 Tunnel in the IPSec MIB (ikeTunIndex in the ikeTunnelTable).
    // The type is interface{} with range: 1..2147483647.
    IkePolMapTunIndex interface{}

    // The number of the locally defined ISAKMP policy used to establish the IPSec
    // IKE Phase-1 Tunnel. This is the number which was used on the crypto
    // command. For example, if the configuration command was:   ==>  crypto
    // isakmp policy 15  then the value of this object would be 15. If ISAKMP was
    // not used to establish this tunnel, then the value of this object will be
    // zero. The type is interface{} with range: 1..2147483647.
    IkePolMapPolicyNum interface{}
}

func (ikePolMapEntry *CISCOIPSECPOLICYMAPMIB_IkePolMapTable_IkePolMapEntry) GetEntityData() *types.CommonEntityData {
    ikePolMapEntry.EntityData.YFilter = ikePolMapEntry.YFilter
    ikePolMapEntry.EntityData.YangName = "ikePolMapEntry"
    ikePolMapEntry.EntityData.BundleName = "cisco_ios_xe"
    ikePolMapEntry.EntityData.ParentYangName = "ikePolMapTable"
    ikePolMapEntry.EntityData.SegmentPath = "ikePolMapEntry" + types.AddKeyToken(ikePolMapEntry.IkePolMapTunIndex, "ikePolMapTunIndex")
    ikePolMapEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ikePolMapEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ikePolMapEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ikePolMapEntry.EntityData.Children = types.NewOrderedMap()
    ikePolMapEntry.EntityData.Leafs = types.NewOrderedMap()
    ikePolMapEntry.EntityData.Leafs.Append("ikePolMapTunIndex", types.YLeaf{"IkePolMapTunIndex", ikePolMapEntry.IkePolMapTunIndex})
    ikePolMapEntry.EntityData.Leafs.Append("ikePolMapPolicyNum", types.YLeaf{"IkePolMapPolicyNum", ikePolMapEntry.IkePolMapPolicyNum})

    ikePolMapEntry.EntityData.YListKeys = []string {"IkePolMapTunIndex"}

    return &(ikePolMapEntry.EntityData)
}

// CISCOIPSECPOLICYMAPMIB_IpSecPolMapTable
// The IPSec Phase-2 Tunnel to Policy Mapping Table.
// There is one entry in this table for each active
// IPSec Phase-2 Tunnel.
type CISCOIPSECPOLICYMAPMIB_IpSecPolMapTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with mapping an active IPSec
    // Phase-2 Tunnel to its configured Policy definition. The type is slice of
    // CISCOIPSECPOLICYMAPMIB_IpSecPolMapTable_IpSecPolMapEntry.
    IpSecPolMapEntry []*CISCOIPSECPOLICYMAPMIB_IpSecPolMapTable_IpSecPolMapEntry
}

func (ipSecPolMapTable *CISCOIPSECPOLICYMAPMIB_IpSecPolMapTable) GetEntityData() *types.CommonEntityData {
    ipSecPolMapTable.EntityData.YFilter = ipSecPolMapTable.YFilter
    ipSecPolMapTable.EntityData.YangName = "ipSecPolMapTable"
    ipSecPolMapTable.EntityData.BundleName = "cisco_ios_xe"
    ipSecPolMapTable.EntityData.ParentYangName = "CISCO-IPSEC-POLICY-MAP-MIB"
    ipSecPolMapTable.EntityData.SegmentPath = "ipSecPolMapTable"
    ipSecPolMapTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipSecPolMapTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipSecPolMapTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipSecPolMapTable.EntityData.Children = types.NewOrderedMap()
    ipSecPolMapTable.EntityData.Children.Append("ipSecPolMapEntry", types.YChild{"IpSecPolMapEntry", nil})
    for i := range ipSecPolMapTable.IpSecPolMapEntry {
        ipSecPolMapTable.EntityData.Children.Append(types.GetSegmentPath(ipSecPolMapTable.IpSecPolMapEntry[i]), types.YChild{"IpSecPolMapEntry", ipSecPolMapTable.IpSecPolMapEntry[i]})
    }
    ipSecPolMapTable.EntityData.Leafs = types.NewOrderedMap()

    ipSecPolMapTable.EntityData.YListKeys = []string {}

    return &(ipSecPolMapTable.EntityData)
}

// CISCOIPSECPOLICYMAPMIB_IpSecPolMapTable_IpSecPolMapEntry
// Each entry contains the attributes associated
// with mapping an active IPSec Phase-2 Tunnel
// to its configured Policy definition.
type CISCOIPSECPOLICYMAPMIB_IpSecPolMapTable_IpSecPolMapEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The index of the IPSec Phase-2 Tunnel to Policy
    // Map Table. The value of the index is the number used to represent this
    // IPSec Phase-2 Tunnel in the IPSec MIB (ipSecTunIndex in the
    // ipSecTunnelTable). The type is interface{} with range: 1..2147483647.
    IpSecPolMapTunIndex interface{}

    // The value of this object should be the name of  the IPSec Policy
    // (cryptomap) as assigned by the  operator while configuring the policy of 
    // the IPSec traffic.  For instance, on an IOS router, the if the command
    // entered to configure the IPSec policy was   ==>  crypto map ftpPolicy 10
    // ipsec-isakmp  then the value of this object would be 'ftpPolicy'. The type
    // is string.
    IpSecPolMapCryptoMapName interface{}

    // The value of this object should be the priority of the IPSec Policy
    // (cryptomap) assigned by the  operator while configuring the policy of  this
    // IPSec tunnel.  For instance, on an IOS router, the if the command entered
    // to configure the IPSec policy was   ==>  crypto map ftpPolicy 10
    // ipsec-isakmp  then the value of this object would be 10. The type is
    // interface{} with range: 1..2147483647.
    IpSecPolMapCryptoMapNum interface{}

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
    IpSecPolMapAclString interface{}

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
    IpSecPolMapAceString interface{}
}

func (ipSecPolMapEntry *CISCOIPSECPOLICYMAPMIB_IpSecPolMapTable_IpSecPolMapEntry) GetEntityData() *types.CommonEntityData {
    ipSecPolMapEntry.EntityData.YFilter = ipSecPolMapEntry.YFilter
    ipSecPolMapEntry.EntityData.YangName = "ipSecPolMapEntry"
    ipSecPolMapEntry.EntityData.BundleName = "cisco_ios_xe"
    ipSecPolMapEntry.EntityData.ParentYangName = "ipSecPolMapTable"
    ipSecPolMapEntry.EntityData.SegmentPath = "ipSecPolMapEntry" + types.AddKeyToken(ipSecPolMapEntry.IpSecPolMapTunIndex, "ipSecPolMapTunIndex")
    ipSecPolMapEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ipSecPolMapEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ipSecPolMapEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ipSecPolMapEntry.EntityData.Children = types.NewOrderedMap()
    ipSecPolMapEntry.EntityData.Leafs = types.NewOrderedMap()
    ipSecPolMapEntry.EntityData.Leafs.Append("ipSecPolMapTunIndex", types.YLeaf{"IpSecPolMapTunIndex", ipSecPolMapEntry.IpSecPolMapTunIndex})
    ipSecPolMapEntry.EntityData.Leafs.Append("ipSecPolMapCryptoMapName", types.YLeaf{"IpSecPolMapCryptoMapName", ipSecPolMapEntry.IpSecPolMapCryptoMapName})
    ipSecPolMapEntry.EntityData.Leafs.Append("ipSecPolMapCryptoMapNum", types.YLeaf{"IpSecPolMapCryptoMapNum", ipSecPolMapEntry.IpSecPolMapCryptoMapNum})
    ipSecPolMapEntry.EntityData.Leafs.Append("ipSecPolMapAclString", types.YLeaf{"IpSecPolMapAclString", ipSecPolMapEntry.IpSecPolMapAclString})
    ipSecPolMapEntry.EntityData.Leafs.Append("ipSecPolMapAceString", types.YLeaf{"IpSecPolMapAceString", ipSecPolMapEntry.IpSecPolMapAceString})

    ipSecPolMapEntry.EntityData.YListKeys = []string {"IpSecPolMapTunIndex"}

    return &(ipSecPolMapEntry.EntityData)
}

