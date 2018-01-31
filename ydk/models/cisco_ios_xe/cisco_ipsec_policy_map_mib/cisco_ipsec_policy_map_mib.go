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
    parent types.Entity
    YFilter yfilter.YFilter

    // The IPSec Phase-1 Internet Key Exchange Tunnel to Policy Mapping Table.
    // There is one entry in this table for each active IPSec Phase-1 Tunnel.
    Ikepolmaptable CISCOIPSECPOLICYMAPMIB_Ikepolmaptable

    // The IPSec Phase-2 Tunnel to Policy Mapping Table. There is one entry in
    // this table for each active IPSec Phase-2 Tunnel.
    Ipsecpolmaptable CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable
}

func (cISCOIPSECPOLICYMAPMIB *CISCOIPSECPOLICYMAPMIB) GetFilter() yfilter.YFilter { return cISCOIPSECPOLICYMAPMIB.YFilter }

func (cISCOIPSECPOLICYMAPMIB *CISCOIPSECPOLICYMAPMIB) SetFilter(yf yfilter.YFilter) { cISCOIPSECPOLICYMAPMIB.YFilter = yf }

func (cISCOIPSECPOLICYMAPMIB *CISCOIPSECPOLICYMAPMIB) GetGoName(yname string) string {
    if yname == "ikePolMapTable" { return "Ikepolmaptable" }
    if yname == "ipSecPolMapTable" { return "Ipsecpolmaptable" }
    return ""
}

func (cISCOIPSECPOLICYMAPMIB *CISCOIPSECPOLICYMAPMIB) GetSegmentPath() string {
    return "CISCO-IPSEC-POLICY-MAP-MIB:CISCO-IPSEC-POLICY-MAP-MIB"
}

func (cISCOIPSECPOLICYMAPMIB *CISCOIPSECPOLICYMAPMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ikePolMapTable" {
        return &cISCOIPSECPOLICYMAPMIB.Ikepolmaptable
    }
    if childYangName == "ipSecPolMapTable" {
        return &cISCOIPSECPOLICYMAPMIB.Ipsecpolmaptable
    }
    return nil
}

func (cISCOIPSECPOLICYMAPMIB *CISCOIPSECPOLICYMAPMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ikePolMapTable"] = &cISCOIPSECPOLICYMAPMIB.Ikepolmaptable
    children["ipSecPolMapTable"] = &cISCOIPSECPOLICYMAPMIB.Ipsecpolmaptable
    return children
}

func (cISCOIPSECPOLICYMAPMIB *CISCOIPSECPOLICYMAPMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOIPSECPOLICYMAPMIB *CISCOIPSECPOLICYMAPMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOIPSECPOLICYMAPMIB *CISCOIPSECPOLICYMAPMIB) GetYangName() string { return "CISCO-IPSEC-POLICY-MAP-MIB" }

func (cISCOIPSECPOLICYMAPMIB *CISCOIPSECPOLICYMAPMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOIPSECPOLICYMAPMIB *CISCOIPSECPOLICYMAPMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOIPSECPOLICYMAPMIB *CISCOIPSECPOLICYMAPMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOIPSECPOLICYMAPMIB *CISCOIPSECPOLICYMAPMIB) SetParent(parent types.Entity) { cISCOIPSECPOLICYMAPMIB.parent = parent }

func (cISCOIPSECPOLICYMAPMIB *CISCOIPSECPOLICYMAPMIB) GetParent() types.Entity { return cISCOIPSECPOLICYMAPMIB.parent }

func (cISCOIPSECPOLICYMAPMIB *CISCOIPSECPOLICYMAPMIB) GetParentYangName() string { return "CISCO-IPSEC-POLICY-MAP-MIB" }

// CISCOIPSECPOLICYMAPMIB_Ikepolmaptable
// The IPSec Phase-1 Internet Key Exchange Tunnel
// to Policy Mapping Table. There is one entry in
// this table for each active IPSec Phase-1
// Tunnel.
type CISCOIPSECPOLICYMAPMIB_Ikepolmaptable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with mapping an active IPSec
    // Phase-1 IKE Tunnel to it's configured Policy definition. The type is slice
    // of CISCOIPSECPOLICYMAPMIB_Ikepolmaptable_Ikepolmapentry.
    Ikepolmapentry []CISCOIPSECPOLICYMAPMIB_Ikepolmaptable_Ikepolmapentry
}

func (ikepolmaptable *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable) GetFilter() yfilter.YFilter { return ikepolmaptable.YFilter }

func (ikepolmaptable *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable) SetFilter(yf yfilter.YFilter) { ikepolmaptable.YFilter = yf }

func (ikepolmaptable *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable) GetGoName(yname string) string {
    if yname == "ikePolMapEntry" { return "Ikepolmapentry" }
    return ""
}

func (ikepolmaptable *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable) GetSegmentPath() string {
    return "ikePolMapTable"
}

func (ikepolmaptable *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ikePolMapEntry" {
        for _, c := range ikepolmaptable.Ikepolmapentry {
            if ikepolmaptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSECPOLICYMAPMIB_Ikepolmaptable_Ikepolmapentry{}
        ikepolmaptable.Ikepolmapentry = append(ikepolmaptable.Ikepolmapentry, child)
        return &ikepolmaptable.Ikepolmapentry[len(ikepolmaptable.Ikepolmapentry)-1]
    }
    return nil
}

func (ikepolmaptable *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ikepolmaptable.Ikepolmapentry {
        children[ikepolmaptable.Ikepolmapentry[i].GetSegmentPath()] = &ikepolmaptable.Ikepolmapentry[i]
    }
    return children
}

func (ikepolmaptable *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ikepolmaptable *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable) GetBundleName() string { return "cisco_ios_xe" }

func (ikepolmaptable *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable) GetYangName() string { return "ikePolMapTable" }

func (ikepolmaptable *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ikepolmaptable *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ikepolmaptable *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ikepolmaptable *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable) SetParent(parent types.Entity) { ikepolmaptable.parent = parent }

func (ikepolmaptable *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable) GetParent() types.Entity { return ikepolmaptable.parent }

func (ikepolmaptable *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable) GetParentYangName() string { return "CISCO-IPSEC-POLICY-MAP-MIB" }

// CISCOIPSECPOLICYMAPMIB_Ikepolmaptable_Ikepolmapentry
// Each entry contains the attributes associated
// with mapping an active IPSec Phase-1 IKE Tunnel
// to it's configured Policy definition.
type CISCOIPSECPOLICYMAPMIB_Ikepolmaptable_Ikepolmapentry struct {
    parent types.Entity
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

func (ikepolmapentry *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable_Ikepolmapentry) GetFilter() yfilter.YFilter { return ikepolmapentry.YFilter }

func (ikepolmapentry *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable_Ikepolmapentry) SetFilter(yf yfilter.YFilter) { ikepolmapentry.YFilter = yf }

func (ikepolmapentry *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable_Ikepolmapentry) GetGoName(yname string) string {
    if yname == "ikePolMapTunIndex" { return "Ikepolmaptunindex" }
    if yname == "ikePolMapPolicyNum" { return "Ikepolmappolicynum" }
    return ""
}

func (ikepolmapentry *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable_Ikepolmapentry) GetSegmentPath() string {
    return "ikePolMapEntry" + "[ikePolMapTunIndex='" + fmt.Sprintf("%v", ikepolmapentry.Ikepolmaptunindex) + "']"
}

func (ikepolmapentry *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable_Ikepolmapentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ikepolmapentry *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable_Ikepolmapentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ikepolmapentry *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable_Ikepolmapentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ikePolMapTunIndex"] = ikepolmapentry.Ikepolmaptunindex
    leafs["ikePolMapPolicyNum"] = ikepolmapentry.Ikepolmappolicynum
    return leafs
}

func (ikepolmapentry *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable_Ikepolmapentry) GetBundleName() string { return "cisco_ios_xe" }

func (ikepolmapentry *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable_Ikepolmapentry) GetYangName() string { return "ikePolMapEntry" }

func (ikepolmapentry *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable_Ikepolmapentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ikepolmapentry *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable_Ikepolmapentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ikepolmapentry *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable_Ikepolmapentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ikepolmapentry *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable_Ikepolmapentry) SetParent(parent types.Entity) { ikepolmapentry.parent = parent }

func (ikepolmapentry *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable_Ikepolmapentry) GetParent() types.Entity { return ikepolmapentry.parent }

func (ikepolmapentry *CISCOIPSECPOLICYMAPMIB_Ikepolmaptable_Ikepolmapentry) GetParentYangName() string { return "ikePolMapTable" }

// CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable
// The IPSec Phase-2 Tunnel to Policy Mapping Table.
// There is one entry in this table for each active
// IPSec Phase-2 Tunnel.
type CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry contains the attributes associated with mapping an active IPSec
    // Phase-2 Tunnel to its configured Policy definition. The type is slice of
    // CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable_Ipsecpolmapentry.
    Ipsecpolmapentry []CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable_Ipsecpolmapentry
}

func (ipsecpolmaptable *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable) GetFilter() yfilter.YFilter { return ipsecpolmaptable.YFilter }

func (ipsecpolmaptable *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable) SetFilter(yf yfilter.YFilter) { ipsecpolmaptable.YFilter = yf }

func (ipsecpolmaptable *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable) GetGoName(yname string) string {
    if yname == "ipSecPolMapEntry" { return "Ipsecpolmapentry" }
    return ""
}

func (ipsecpolmaptable *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable) GetSegmentPath() string {
    return "ipSecPolMapTable"
}

func (ipsecpolmaptable *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ipSecPolMapEntry" {
        for _, c := range ipsecpolmaptable.Ipsecpolmapentry {
            if ipsecpolmaptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable_Ipsecpolmapentry{}
        ipsecpolmaptable.Ipsecpolmapentry = append(ipsecpolmaptable.Ipsecpolmapentry, child)
        return &ipsecpolmaptable.Ipsecpolmapentry[len(ipsecpolmaptable.Ipsecpolmapentry)-1]
    }
    return nil
}

func (ipsecpolmaptable *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ipsecpolmaptable.Ipsecpolmapentry {
        children[ipsecpolmaptable.Ipsecpolmapentry[i].GetSegmentPath()] = &ipsecpolmaptable.Ipsecpolmapentry[i]
    }
    return children
}

func (ipsecpolmaptable *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ipsecpolmaptable *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable) GetBundleName() string { return "cisco_ios_xe" }

func (ipsecpolmaptable *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable) GetYangName() string { return "ipSecPolMapTable" }

func (ipsecpolmaptable *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipsecpolmaptable *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipsecpolmaptable *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipsecpolmaptable *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable) SetParent(parent types.Entity) { ipsecpolmaptable.parent = parent }

func (ipsecpolmaptable *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable) GetParent() types.Entity { return ipsecpolmaptable.parent }

func (ipsecpolmaptable *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable) GetParentYangName() string { return "CISCO-IPSEC-POLICY-MAP-MIB" }

// CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable_Ipsecpolmapentry
// Each entry contains the attributes associated
// with mapping an active IPSec Phase-2 Tunnel
// to its configured Policy definition.
type CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable_Ipsecpolmapentry struct {
    parent types.Entity
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

func (ipsecpolmapentry *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable_Ipsecpolmapentry) GetFilter() yfilter.YFilter { return ipsecpolmapentry.YFilter }

func (ipsecpolmapentry *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable_Ipsecpolmapentry) SetFilter(yf yfilter.YFilter) { ipsecpolmapentry.YFilter = yf }

func (ipsecpolmapentry *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable_Ipsecpolmapentry) GetGoName(yname string) string {
    if yname == "ipSecPolMapTunIndex" { return "Ipsecpolmaptunindex" }
    if yname == "ipSecPolMapCryptoMapName" { return "Ipsecpolmapcryptomapname" }
    if yname == "ipSecPolMapCryptoMapNum" { return "Ipsecpolmapcryptomapnum" }
    if yname == "ipSecPolMapAclString" { return "Ipsecpolmapaclstring" }
    if yname == "ipSecPolMapAceString" { return "Ipsecpolmapacestring" }
    return ""
}

func (ipsecpolmapentry *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable_Ipsecpolmapentry) GetSegmentPath() string {
    return "ipSecPolMapEntry" + "[ipSecPolMapTunIndex='" + fmt.Sprintf("%v", ipsecpolmapentry.Ipsecpolmaptunindex) + "']"
}

func (ipsecpolmapentry *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable_Ipsecpolmapentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ipsecpolmapentry *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable_Ipsecpolmapentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ipsecpolmapentry *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable_Ipsecpolmapentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ipSecPolMapTunIndex"] = ipsecpolmapentry.Ipsecpolmaptunindex
    leafs["ipSecPolMapCryptoMapName"] = ipsecpolmapentry.Ipsecpolmapcryptomapname
    leafs["ipSecPolMapCryptoMapNum"] = ipsecpolmapentry.Ipsecpolmapcryptomapnum
    leafs["ipSecPolMapAclString"] = ipsecpolmapentry.Ipsecpolmapaclstring
    leafs["ipSecPolMapAceString"] = ipsecpolmapentry.Ipsecpolmapacestring
    return leafs
}

func (ipsecpolmapentry *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable_Ipsecpolmapentry) GetBundleName() string { return "cisco_ios_xe" }

func (ipsecpolmapentry *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable_Ipsecpolmapentry) GetYangName() string { return "ipSecPolMapEntry" }

func (ipsecpolmapentry *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable_Ipsecpolmapentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ipsecpolmapentry *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable_Ipsecpolmapentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ipsecpolmapentry *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable_Ipsecpolmapentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ipsecpolmapentry *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable_Ipsecpolmapentry) SetParent(parent types.Entity) { ipsecpolmapentry.parent = parent }

func (ipsecpolmapentry *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable_Ipsecpolmapentry) GetParent() types.Entity { return ipsecpolmapentry.parent }

func (ipsecpolmapentry *CISCOIPSECPOLICYMAPMIB_Ipsecpolmaptable_Ipsecpolmapentry) GetParentYangName() string { return "ipSecPolMapTable" }

