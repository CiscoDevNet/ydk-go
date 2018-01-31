// This MIB defines the configuration and monitoring capabilities
// relating to local IP pools.
// 
// Local IP pools have the following characteristics:
// 
// - An IP local pool consists of one or more IP address ranges.
// 
// - An IP pool group consists of one or more IP local pools.
// 
// - An IP local pool can only belong to one IP pool group.
// 
// - IP local pools that belong to different groups can have
//   overlapping addresses.
// 
// - IP local pool names are unique even when they belong to
//   different groups.
// 
// - Addresses within an IP pool group can not overlap.
// 
// - IP local pools without an explicit group name are considered
//   members of the base system group.  In this MIB, the base
//   system group is represented by a null IP pool group name.
// 
// This MIB defines objects that expose the relationship between
// IP pool groups and IP local pools.  There exist other objects
// that maintain statistics about the address usage of IP local
// pools.
package cisco_ip_local_pool_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ip_local_pool_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IP-LOCAL-POOL-MIB CISCO-IP-LOCAL-POOL-MIB}", reflect.TypeOf(CISCOIPLOCALPOOLMIB{}))
    ydk.RegisterEntity("CISCO-IP-LOCAL-POOL-MIB:CISCO-IP-LOCAL-POOL-MIB", reflect.TypeOf(CISCOIPLOCALPOOLMIB{}))
}

// CISCOIPLOCALPOOLMIB
type CISCOIPLOCALPOOLMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Ciplocalpoolconfig CISCOIPLOCALPOOLMIB_Ciplocalpoolconfig

    // This table manages the creation, modification, and deletion of IP local
    // pools using the RowStatus textual convention.  An entry in this table
    // defines an IP address range that is associated with an IP local pool.  A
    // conceptual row in this table can not be modified while
    // cIpLocalPoolRowStatus is set to 'active'.  Since IP local pool names are
    // unique even when they belong to different groups, and addresses within a
    // group can not overlap, a row in this table is uniquely indexed by the pool
    // name, and by the low address of the IP local pool together with its address
    // type.
    Ciplocalpoolconfigtable CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable

    // A table which exposes the container/'containee' relationships between local
    // IP pools and IP pool groups.  Entries in this table are created or deleted
    // as a by-product of creating or deleting entries in the
    // cIpLocalPoolConfigTable.  When an entry is created and activated in the
    // cIpLocalPoolConfigTable table, an entry in this table will come into
    // existence if it does not already exist.  When an entry is deleted in the
    // cIpLocalPoolConfigTable table, if there is no other entry existing in that
    // table with the same cIpLocalPoolGroupContainedIn and cIpLocalPoolName, the
    // entry in this table with the respective cIpLocalPoolGroupName and
    // cIpLocalPoolName indices will be removed.
    Ciplocalpoolgroupcontainstable CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable

    // This table provides statistics for configured IP pool groups.  Entries in
    // this table are created as the result of adding a new IP pool group to the
    // cIpLocalPoolConfigTable.  Entries in this table are deleted as the result
    // of removing all IP local pools that are contained in an IP pool group in
    // the cIpLocalPoolConfigTable.  An entry in this table is uniquely indexed by
    // IP pool group name.
    Ciplocalpoolgrouptable CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable

    // A table providing statistics for each IP local pool.  Entries in this table
    // are created as the result of adding a new IP local pool to the
    // cIpLocalPoolConfigTable.  Entries in this table are deleted as the result
    // of removing all the address ranges that are contained in an IP local pool
    // in the cIpLocalPoolConfigTable.  Entries in this table are uniquely indexed
    // by the name of the IP local pool.
    Ciplocalpoolstatstable CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable

    // This table lists all addresses that have been allocated out of an IP local
    // pool.  Entries in this table are created when a remote peer allocates an
    // address from one of the IP local pools in the cIpLocalPoolConfigTable. 
    // Entries in this table are deleted when a remote peer deallocates an address
    // from one of the IP local pool in the cIpLocalPoolConfigTable.  Entries in
    // this table are uniquely indexed by the name of the IP local pool, and the
    // allocated address, together with its address type.
    Ciplocalpoolalloctable CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable
}

func (cISCOIPLOCALPOOLMIB *CISCOIPLOCALPOOLMIB) GetFilter() yfilter.YFilter { return cISCOIPLOCALPOOLMIB.YFilter }

func (cISCOIPLOCALPOOLMIB *CISCOIPLOCALPOOLMIB) SetFilter(yf yfilter.YFilter) { cISCOIPLOCALPOOLMIB.YFilter = yf }

func (cISCOIPLOCALPOOLMIB *CISCOIPLOCALPOOLMIB) GetGoName(yname string) string {
    if yname == "cIpLocalPoolConfig" { return "Ciplocalpoolconfig" }
    if yname == "cIpLocalPoolConfigTable" { return "Ciplocalpoolconfigtable" }
    if yname == "cIpLocalPoolGroupContainsTable" { return "Ciplocalpoolgroupcontainstable" }
    if yname == "cIpLocalPoolGroupTable" { return "Ciplocalpoolgrouptable" }
    if yname == "cIpLocalPoolStatsTable" { return "Ciplocalpoolstatstable" }
    if yname == "cIpLocalPoolAllocTable" { return "Ciplocalpoolalloctable" }
    return ""
}

func (cISCOIPLOCALPOOLMIB *CISCOIPLOCALPOOLMIB) GetSegmentPath() string {
    return "CISCO-IP-LOCAL-POOL-MIB:CISCO-IP-LOCAL-POOL-MIB"
}

func (cISCOIPLOCALPOOLMIB *CISCOIPLOCALPOOLMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cIpLocalPoolConfig" {
        return &cISCOIPLOCALPOOLMIB.Ciplocalpoolconfig
    }
    if childYangName == "cIpLocalPoolConfigTable" {
        return &cISCOIPLOCALPOOLMIB.Ciplocalpoolconfigtable
    }
    if childYangName == "cIpLocalPoolGroupContainsTable" {
        return &cISCOIPLOCALPOOLMIB.Ciplocalpoolgroupcontainstable
    }
    if childYangName == "cIpLocalPoolGroupTable" {
        return &cISCOIPLOCALPOOLMIB.Ciplocalpoolgrouptable
    }
    if childYangName == "cIpLocalPoolStatsTable" {
        return &cISCOIPLOCALPOOLMIB.Ciplocalpoolstatstable
    }
    if childYangName == "cIpLocalPoolAllocTable" {
        return &cISCOIPLOCALPOOLMIB.Ciplocalpoolalloctable
    }
    return nil
}

func (cISCOIPLOCALPOOLMIB *CISCOIPLOCALPOOLMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cIpLocalPoolConfig"] = &cISCOIPLOCALPOOLMIB.Ciplocalpoolconfig
    children["cIpLocalPoolConfigTable"] = &cISCOIPLOCALPOOLMIB.Ciplocalpoolconfigtable
    children["cIpLocalPoolGroupContainsTable"] = &cISCOIPLOCALPOOLMIB.Ciplocalpoolgroupcontainstable
    children["cIpLocalPoolGroupTable"] = &cISCOIPLOCALPOOLMIB.Ciplocalpoolgrouptable
    children["cIpLocalPoolStatsTable"] = &cISCOIPLOCALPOOLMIB.Ciplocalpoolstatstable
    children["cIpLocalPoolAllocTable"] = &cISCOIPLOCALPOOLMIB.Ciplocalpoolalloctable
    return children
}

func (cISCOIPLOCALPOOLMIB *CISCOIPLOCALPOOLMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOIPLOCALPOOLMIB *CISCOIPLOCALPOOLMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOIPLOCALPOOLMIB *CISCOIPLOCALPOOLMIB) GetYangName() string { return "CISCO-IP-LOCAL-POOL-MIB" }

func (cISCOIPLOCALPOOLMIB *CISCOIPLOCALPOOLMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOIPLOCALPOOLMIB *CISCOIPLOCALPOOLMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOIPLOCALPOOLMIB *CISCOIPLOCALPOOLMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOIPLOCALPOOLMIB *CISCOIPLOCALPOOLMIB) SetParent(parent types.Entity) { cISCOIPLOCALPOOLMIB.parent = parent }

func (cISCOIPLOCALPOOLMIB *CISCOIPLOCALPOOLMIB) GetParent() types.Entity { return cISCOIPLOCALPOOLMIB.parent }

func (cISCOIPLOCALPOOLMIB *CISCOIPLOCALPOOLMIB) GetParentYangName() string { return "CISCO-IP-LOCAL-POOL-MIB" }

// CISCOIPLOCALPOOLMIB_Ciplocalpoolconfig
type CISCOIPLOCALPOOLMIB_Ciplocalpoolconfig struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An indication of whether the notifications defined by the
    // ciscoIpLocalPoolNotifGroup are enabled. The type is bool.
    Ciplocalpoolnotificationsenable interface{}
}

func (ciplocalpoolconfig *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfig) GetFilter() yfilter.YFilter { return ciplocalpoolconfig.YFilter }

func (ciplocalpoolconfig *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfig) SetFilter(yf yfilter.YFilter) { ciplocalpoolconfig.YFilter = yf }

func (ciplocalpoolconfig *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfig) GetGoName(yname string) string {
    if yname == "cIpLocalPoolNotificationsEnable" { return "Ciplocalpoolnotificationsenable" }
    return ""
}

func (ciplocalpoolconfig *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfig) GetSegmentPath() string {
    return "cIpLocalPoolConfig"
}

func (ciplocalpoolconfig *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciplocalpoolconfig *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciplocalpoolconfig *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cIpLocalPoolNotificationsEnable"] = ciplocalpoolconfig.Ciplocalpoolnotificationsenable
    return leafs
}

func (ciplocalpoolconfig *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfig) GetBundleName() string { return "cisco_ios_xe" }

func (ciplocalpoolconfig *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfig) GetYangName() string { return "cIpLocalPoolConfig" }

func (ciplocalpoolconfig *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfig) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciplocalpoolconfig *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfig) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciplocalpoolconfig *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfig) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciplocalpoolconfig *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfig) SetParent(parent types.Entity) { ciplocalpoolconfig.parent = parent }

func (ciplocalpoolconfig *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfig) GetParent() types.Entity { return ciplocalpoolconfig.parent }

func (ciplocalpoolconfig *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfig) GetParentYangName() string { return "CISCO-IP-LOCAL-POOL-MIB" }

// CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable
// This table manages the creation, modification, and deletion
// of IP local pools using the RowStatus textual convention.  An
// entry in this table defines an IP address range that is
// associated with an IP local pool.
// 
// A conceptual row in this table can not be modified while
// cIpLocalPoolRowStatus is set to 'active'.
// 
// Since IP local pool names are unique even when they belong to
// different groups, and addresses within a group can not overlap,
// a row in this table is uniquely indexed by the pool name, and by
// the low address of the IP local pool together with its address
// type.
type CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry provides information about a particular IP local pool, including
    // the number of free and used addresses and its priority. The type is slice
    // of CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable_Ciplocalpoolconfigentry.
    Ciplocalpoolconfigentry []CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable_Ciplocalpoolconfigentry
}

func (ciplocalpoolconfigtable *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable) GetFilter() yfilter.YFilter { return ciplocalpoolconfigtable.YFilter }

func (ciplocalpoolconfigtable *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable) SetFilter(yf yfilter.YFilter) { ciplocalpoolconfigtable.YFilter = yf }

func (ciplocalpoolconfigtable *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable) GetGoName(yname string) string {
    if yname == "cIpLocalPoolConfigEntry" { return "Ciplocalpoolconfigentry" }
    return ""
}

func (ciplocalpoolconfigtable *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable) GetSegmentPath() string {
    return "cIpLocalPoolConfigTable"
}

func (ciplocalpoolconfigtable *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cIpLocalPoolConfigEntry" {
        for _, c := range ciplocalpoolconfigtable.Ciplocalpoolconfigentry {
            if ciplocalpoolconfigtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable_Ciplocalpoolconfigentry{}
        ciplocalpoolconfigtable.Ciplocalpoolconfigentry = append(ciplocalpoolconfigtable.Ciplocalpoolconfigentry, child)
        return &ciplocalpoolconfigtable.Ciplocalpoolconfigentry[len(ciplocalpoolconfigtable.Ciplocalpoolconfigentry)-1]
    }
    return nil
}

func (ciplocalpoolconfigtable *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ciplocalpoolconfigtable.Ciplocalpoolconfigentry {
        children[ciplocalpoolconfigtable.Ciplocalpoolconfigentry[i].GetSegmentPath()] = &ciplocalpoolconfigtable.Ciplocalpoolconfigentry[i]
    }
    return children
}

func (ciplocalpoolconfigtable *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciplocalpoolconfigtable *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable) GetBundleName() string { return "cisco_ios_xe" }

func (ciplocalpoolconfigtable *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable) GetYangName() string { return "cIpLocalPoolConfigTable" }

func (ciplocalpoolconfigtable *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciplocalpoolconfigtable *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciplocalpoolconfigtable *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciplocalpoolconfigtable *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable) SetParent(parent types.Entity) { ciplocalpoolconfigtable.parent = parent }

func (ciplocalpoolconfigtable *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable) GetParent() types.Entity { return ciplocalpoolconfigtable.parent }

func (ciplocalpoolconfigtable *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable) GetParentYangName() string { return "CISCO-IP-LOCAL-POOL-MIB" }

// CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable_Ciplocalpoolconfigentry
// Each entry provides information about a particular IP local
// pool, including the number of free and used addresses and its priority.
type CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable_Ciplocalpoolconfigentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. An arbitrary non-empty string that uniquely
    // identifies the IP local pool.  This name must be unique among all the local
    // IP pools even when they belong to different pool groups. The type is string
    // with length: 1..48.
    Ciplocalpoolname interface{}

    // This attribute is a key. This object specifies the address type of
    // cIpLocalPoolAddressLo and cIpLocalPoolAddressHi. The type is
    // InetAddressType.
    Ciplocalpooladdrtype interface{}

    // This attribute is a key. This object specifies the first IP address of the
    // range of IP addresses contained by this pool entry.  The address type of
    // this object is described by cIpLocalPoolAddrType.  This address must be
    // less than or equal to the address in cIpLocalPoolAddressHi. The type is
    // string with length: 0..255.
    Ciplocalpooladdresslo interface{}

    // This object specifies the last IP address of the range of IP addresses
    // mapped by this pool entry.  The address type of this object is described by
    // cIpLocalPoolAddrType.  If only a single address is being mapped, the value
    // of this object is equal to the value of cIpLocalPoolAddressLo. The type is
    // string with length: 0..255.
    Ciplocalpooladdresshi interface{}

    // The number of IP addresses available for use in the range of IP addresses.
    // The type is interface{} with range: 0..4294967295.
    Ciplocalpoolfreeaddrs interface{}

    // The number of IP addresses being used in the range of IP addresses. The
    // type is interface{} with range: 0..4294967295.
    Ciplocalpoolinuseaddrs interface{}

    // This object relates an IP local pool to its IP pool group.  A null string
    // indicates this IP local pool is not contained in a named IP pool group, but
    // that it is contained in the base IP pool group.  An IP local pool can only
    // belong to one IP pool group. The type is string with length: 0..48.
    Ciplocalpoolgroupcontainedin interface{}

    // This object facilitates the creation, or deletion of a conceptual row in
    // this table. The type is RowStatus.
    Ciplocalpoolrowstatus interface{}

    // This object specifies priority of the IP local pool, where smaller value
    // indicates the lower priority. The priority value is used in assigning IP
    // Address  from local pools. The type is interface{} with range: 1..255.
    Ciplocalpoolpriority interface{}
}

func (ciplocalpoolconfigentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable_Ciplocalpoolconfigentry) GetFilter() yfilter.YFilter { return ciplocalpoolconfigentry.YFilter }

func (ciplocalpoolconfigentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable_Ciplocalpoolconfigentry) SetFilter(yf yfilter.YFilter) { ciplocalpoolconfigentry.YFilter = yf }

func (ciplocalpoolconfigentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable_Ciplocalpoolconfigentry) GetGoName(yname string) string {
    if yname == "cIpLocalPoolName" { return "Ciplocalpoolname" }
    if yname == "cIpLocalPoolAddrType" { return "Ciplocalpooladdrtype" }
    if yname == "cIpLocalPoolAddressLo" { return "Ciplocalpooladdresslo" }
    if yname == "cIpLocalPoolAddressHi" { return "Ciplocalpooladdresshi" }
    if yname == "cIpLocalPoolFreeAddrs" { return "Ciplocalpoolfreeaddrs" }
    if yname == "cIpLocalPoolInUseAddrs" { return "Ciplocalpoolinuseaddrs" }
    if yname == "cIpLocalPoolGroupContainedIn" { return "Ciplocalpoolgroupcontainedin" }
    if yname == "cIpLocalPoolRowStatus" { return "Ciplocalpoolrowstatus" }
    if yname == "cIpLocalPoolPriority" { return "Ciplocalpoolpriority" }
    return ""
}

func (ciplocalpoolconfigentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable_Ciplocalpoolconfigentry) GetSegmentPath() string {
    return "cIpLocalPoolConfigEntry" + "[cIpLocalPoolName='" + fmt.Sprintf("%v", ciplocalpoolconfigentry.Ciplocalpoolname) + "']" + "[cIpLocalPoolAddrType='" + fmt.Sprintf("%v", ciplocalpoolconfigentry.Ciplocalpooladdrtype) + "']" + "[cIpLocalPoolAddressLo='" + fmt.Sprintf("%v", ciplocalpoolconfigentry.Ciplocalpooladdresslo) + "']"
}

func (ciplocalpoolconfigentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable_Ciplocalpoolconfigentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciplocalpoolconfigentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable_Ciplocalpoolconfigentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciplocalpoolconfigentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable_Ciplocalpoolconfigentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cIpLocalPoolName"] = ciplocalpoolconfigentry.Ciplocalpoolname
    leafs["cIpLocalPoolAddrType"] = ciplocalpoolconfigentry.Ciplocalpooladdrtype
    leafs["cIpLocalPoolAddressLo"] = ciplocalpoolconfigentry.Ciplocalpooladdresslo
    leafs["cIpLocalPoolAddressHi"] = ciplocalpoolconfigentry.Ciplocalpooladdresshi
    leafs["cIpLocalPoolFreeAddrs"] = ciplocalpoolconfigentry.Ciplocalpoolfreeaddrs
    leafs["cIpLocalPoolInUseAddrs"] = ciplocalpoolconfigentry.Ciplocalpoolinuseaddrs
    leafs["cIpLocalPoolGroupContainedIn"] = ciplocalpoolconfigentry.Ciplocalpoolgroupcontainedin
    leafs["cIpLocalPoolRowStatus"] = ciplocalpoolconfigentry.Ciplocalpoolrowstatus
    leafs["cIpLocalPoolPriority"] = ciplocalpoolconfigentry.Ciplocalpoolpriority
    return leafs
}

func (ciplocalpoolconfigentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable_Ciplocalpoolconfigentry) GetBundleName() string { return "cisco_ios_xe" }

func (ciplocalpoolconfigentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable_Ciplocalpoolconfigentry) GetYangName() string { return "cIpLocalPoolConfigEntry" }

func (ciplocalpoolconfigentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable_Ciplocalpoolconfigentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciplocalpoolconfigentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable_Ciplocalpoolconfigentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciplocalpoolconfigentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable_Ciplocalpoolconfigentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciplocalpoolconfigentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable_Ciplocalpoolconfigentry) SetParent(parent types.Entity) { ciplocalpoolconfigentry.parent = parent }

func (ciplocalpoolconfigentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable_Ciplocalpoolconfigentry) GetParent() types.Entity { return ciplocalpoolconfigentry.parent }

func (ciplocalpoolconfigentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable_Ciplocalpoolconfigentry) GetParentYangName() string { return "cIpLocalPoolConfigTable" }

// CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable
// A table which exposes the container/'containee' relationships
// between local IP pools and IP pool groups.
// 
// Entries in this table are created or deleted as a by-product of
// creating or deleting entries in the cIpLocalPoolConfigTable.
// 
// When an entry is created and activated in the
// cIpLocalPoolConfigTable table, an entry in this table will come
// into existence if it does not already exist.
// 
// When an entry is deleted in the cIpLocalPoolConfigTable table,
// if there is no other entry existing in that table with the same
// cIpLocalPoolGroupContainedIn and cIpLocalPoolName, the entry in
// this table with the respective cIpLocalPoolGroupName and
// cIpLocalPoolName indices will be removed.
type CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry describes single container/'containee' relationship.  Pool names
    // can only be associated with one group.  Pools carry implicit group
    // identifiers because pool names can only be associated with one group.  An
    // entry in this table describes such an association. The type is slice of
    // CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable_Ciplocalpoolgroupcontainsentry.
    Ciplocalpoolgroupcontainsentry []CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable_Ciplocalpoolgroupcontainsentry
}

func (ciplocalpoolgroupcontainstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable) GetFilter() yfilter.YFilter { return ciplocalpoolgroupcontainstable.YFilter }

func (ciplocalpoolgroupcontainstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable) SetFilter(yf yfilter.YFilter) { ciplocalpoolgroupcontainstable.YFilter = yf }

func (ciplocalpoolgroupcontainstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable) GetGoName(yname string) string {
    if yname == "cIpLocalPoolGroupContainsEntry" { return "Ciplocalpoolgroupcontainsentry" }
    return ""
}

func (ciplocalpoolgroupcontainstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable) GetSegmentPath() string {
    return "cIpLocalPoolGroupContainsTable"
}

func (ciplocalpoolgroupcontainstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cIpLocalPoolGroupContainsEntry" {
        for _, c := range ciplocalpoolgroupcontainstable.Ciplocalpoolgroupcontainsentry {
            if ciplocalpoolgroupcontainstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable_Ciplocalpoolgroupcontainsentry{}
        ciplocalpoolgroupcontainstable.Ciplocalpoolgroupcontainsentry = append(ciplocalpoolgroupcontainstable.Ciplocalpoolgroupcontainsentry, child)
        return &ciplocalpoolgroupcontainstable.Ciplocalpoolgroupcontainsentry[len(ciplocalpoolgroupcontainstable.Ciplocalpoolgroupcontainsentry)-1]
    }
    return nil
}

func (ciplocalpoolgroupcontainstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ciplocalpoolgroupcontainstable.Ciplocalpoolgroupcontainsentry {
        children[ciplocalpoolgroupcontainstable.Ciplocalpoolgroupcontainsentry[i].GetSegmentPath()] = &ciplocalpoolgroupcontainstable.Ciplocalpoolgroupcontainsentry[i]
    }
    return children
}

func (ciplocalpoolgroupcontainstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciplocalpoolgroupcontainstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable) GetBundleName() string { return "cisco_ios_xe" }

func (ciplocalpoolgroupcontainstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable) GetYangName() string { return "cIpLocalPoolGroupContainsTable" }

func (ciplocalpoolgroupcontainstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciplocalpoolgroupcontainstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciplocalpoolgroupcontainstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciplocalpoolgroupcontainstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable) SetParent(parent types.Entity) { ciplocalpoolgroupcontainstable.parent = parent }

func (ciplocalpoolgroupcontainstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable) GetParent() types.Entity { return ciplocalpoolgroupcontainstable.parent }

func (ciplocalpoolgroupcontainstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable) GetParentYangName() string { return "CISCO-IP-LOCAL-POOL-MIB" }

// CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable_Ciplocalpoolgroupcontainsentry
// Each entry describes single container/'containee'
// relationship.
// 
// Pool names can only be associated with one group.  Pools carry
// implicit group identifiers because pool names can only be
// associated with one group.  An entry in this table describes
// such an association.
type CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable_Ciplocalpoolgroupcontainsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. A unique group name that identifies the IP pool
    // group.  The null string represents the base IP pool group. The type is
    // string with length: 0..48.
    Ciplocalpoolgroupname interface{}

    // This attribute is a key. The value of cIpLocalPoolName for the contained IP
    // local pool. The type is string with length: 1..48.
    Ciplocalpoolchildindex interface{}
}

func (ciplocalpoolgroupcontainsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable_Ciplocalpoolgroupcontainsentry) GetFilter() yfilter.YFilter { return ciplocalpoolgroupcontainsentry.YFilter }

func (ciplocalpoolgroupcontainsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable_Ciplocalpoolgroupcontainsentry) SetFilter(yf yfilter.YFilter) { ciplocalpoolgroupcontainsentry.YFilter = yf }

func (ciplocalpoolgroupcontainsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable_Ciplocalpoolgroupcontainsentry) GetGoName(yname string) string {
    if yname == "cIpLocalPoolGroupName" { return "Ciplocalpoolgroupname" }
    if yname == "cIpLocalPoolChildIndex" { return "Ciplocalpoolchildindex" }
    return ""
}

func (ciplocalpoolgroupcontainsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable_Ciplocalpoolgroupcontainsentry) GetSegmentPath() string {
    return "cIpLocalPoolGroupContainsEntry" + "[cIpLocalPoolGroupName='" + fmt.Sprintf("%v", ciplocalpoolgroupcontainsentry.Ciplocalpoolgroupname) + "']" + "[cIpLocalPoolChildIndex='" + fmt.Sprintf("%v", ciplocalpoolgroupcontainsentry.Ciplocalpoolchildindex) + "']"
}

func (ciplocalpoolgroupcontainsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable_Ciplocalpoolgroupcontainsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciplocalpoolgroupcontainsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable_Ciplocalpoolgroupcontainsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciplocalpoolgroupcontainsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable_Ciplocalpoolgroupcontainsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cIpLocalPoolGroupName"] = ciplocalpoolgroupcontainsentry.Ciplocalpoolgroupname
    leafs["cIpLocalPoolChildIndex"] = ciplocalpoolgroupcontainsentry.Ciplocalpoolchildindex
    return leafs
}

func (ciplocalpoolgroupcontainsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable_Ciplocalpoolgroupcontainsentry) GetBundleName() string { return "cisco_ios_xe" }

func (ciplocalpoolgroupcontainsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable_Ciplocalpoolgroupcontainsentry) GetYangName() string { return "cIpLocalPoolGroupContainsEntry" }

func (ciplocalpoolgroupcontainsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable_Ciplocalpoolgroupcontainsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciplocalpoolgroupcontainsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable_Ciplocalpoolgroupcontainsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciplocalpoolgroupcontainsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable_Ciplocalpoolgroupcontainsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciplocalpoolgroupcontainsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable_Ciplocalpoolgroupcontainsentry) SetParent(parent types.Entity) { ciplocalpoolgroupcontainsentry.parent = parent }

func (ciplocalpoolgroupcontainsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable_Ciplocalpoolgroupcontainsentry) GetParent() types.Entity { return ciplocalpoolgroupcontainsentry.parent }

func (ciplocalpoolgroupcontainsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable_Ciplocalpoolgroupcontainsentry) GetParentYangName() string { return "cIpLocalPoolGroupContainsTable" }

// CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable
// This table provides statistics for configured IP pool groups.
// 
// Entries in this table are created as the result of adding a new
// IP pool group to the cIpLocalPoolConfigTable.
// 
// Entries in this table are deleted as the result of removing all
// IP local pools that are contained in an IP pool group in the
// cIpLocalPoolConfigTable.
// 
// An entry in this table is uniquely indexed by IP pool group
// name.
type CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry provides information about a particular IP pool group and the
    // number of free and used addresses in an IP pool group. The type is slice of
    // CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable_Ciplocalpoolgroupentry.
    Ciplocalpoolgroupentry []CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable_Ciplocalpoolgroupentry
}

func (ciplocalpoolgrouptable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable) GetFilter() yfilter.YFilter { return ciplocalpoolgrouptable.YFilter }

func (ciplocalpoolgrouptable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable) SetFilter(yf yfilter.YFilter) { ciplocalpoolgrouptable.YFilter = yf }

func (ciplocalpoolgrouptable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable) GetGoName(yname string) string {
    if yname == "cIpLocalPoolGroupEntry" { return "Ciplocalpoolgroupentry" }
    return ""
}

func (ciplocalpoolgrouptable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable) GetSegmentPath() string {
    return "cIpLocalPoolGroupTable"
}

func (ciplocalpoolgrouptable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cIpLocalPoolGroupEntry" {
        for _, c := range ciplocalpoolgrouptable.Ciplocalpoolgroupentry {
            if ciplocalpoolgrouptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable_Ciplocalpoolgroupentry{}
        ciplocalpoolgrouptable.Ciplocalpoolgroupentry = append(ciplocalpoolgrouptable.Ciplocalpoolgroupentry, child)
        return &ciplocalpoolgrouptable.Ciplocalpoolgroupentry[len(ciplocalpoolgrouptable.Ciplocalpoolgroupentry)-1]
    }
    return nil
}

func (ciplocalpoolgrouptable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ciplocalpoolgrouptable.Ciplocalpoolgroupentry {
        children[ciplocalpoolgrouptable.Ciplocalpoolgroupentry[i].GetSegmentPath()] = &ciplocalpoolgrouptable.Ciplocalpoolgroupentry[i]
    }
    return children
}

func (ciplocalpoolgrouptable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciplocalpoolgrouptable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable) GetBundleName() string { return "cisco_ios_xe" }

func (ciplocalpoolgrouptable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable) GetYangName() string { return "cIpLocalPoolGroupTable" }

func (ciplocalpoolgrouptable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciplocalpoolgrouptable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciplocalpoolgrouptable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciplocalpoolgrouptable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable) SetParent(parent types.Entity) { ciplocalpoolgrouptable.parent = parent }

func (ciplocalpoolgrouptable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable) GetParent() types.Entity { return ciplocalpoolgrouptable.parent }

func (ciplocalpoolgrouptable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable) GetParentYangName() string { return "CISCO-IP-LOCAL-POOL-MIB" }

// CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable_Ciplocalpoolgroupentry
// Each entry provides information about a particular IP pool
// group and the number of free and used addresses in an IP pool
// group.
type CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable_Ciplocalpoolgroupentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..48. Refers to
    // cisco_ip_local_pool_mib.CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable_Ciplocalpoolgroupcontainsentry_Ciplocalpoolgroupname
    Ciplocalpoolgroupname interface{}

    // The number of IP addresses available for use in the IP pool group. The type
    // is interface{} with range: 0..4294967295.
    Ciplocalpoolgroupfreeaddrs interface{}

    // The number of IP addresses that have been allocated from the IP pool group.
    // The type is interface{} with range: 0..4294967295.
    Ciplocalpoolgroupinuseaddrs interface{}
}

func (ciplocalpoolgroupentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable_Ciplocalpoolgroupentry) GetFilter() yfilter.YFilter { return ciplocalpoolgroupentry.YFilter }

func (ciplocalpoolgroupentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable_Ciplocalpoolgroupentry) SetFilter(yf yfilter.YFilter) { ciplocalpoolgroupentry.YFilter = yf }

func (ciplocalpoolgroupentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable_Ciplocalpoolgroupentry) GetGoName(yname string) string {
    if yname == "cIpLocalPoolGroupName" { return "Ciplocalpoolgroupname" }
    if yname == "cIpLocalPoolGroupFreeAddrs" { return "Ciplocalpoolgroupfreeaddrs" }
    if yname == "cIpLocalPoolGroupInUseAddrs" { return "Ciplocalpoolgroupinuseaddrs" }
    return ""
}

func (ciplocalpoolgroupentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable_Ciplocalpoolgroupentry) GetSegmentPath() string {
    return "cIpLocalPoolGroupEntry" + "[cIpLocalPoolGroupName='" + fmt.Sprintf("%v", ciplocalpoolgroupentry.Ciplocalpoolgroupname) + "']"
}

func (ciplocalpoolgroupentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable_Ciplocalpoolgroupentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciplocalpoolgroupentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable_Ciplocalpoolgroupentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciplocalpoolgroupentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable_Ciplocalpoolgroupentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cIpLocalPoolGroupName"] = ciplocalpoolgroupentry.Ciplocalpoolgroupname
    leafs["cIpLocalPoolGroupFreeAddrs"] = ciplocalpoolgroupentry.Ciplocalpoolgroupfreeaddrs
    leafs["cIpLocalPoolGroupInUseAddrs"] = ciplocalpoolgroupentry.Ciplocalpoolgroupinuseaddrs
    return leafs
}

func (ciplocalpoolgroupentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable_Ciplocalpoolgroupentry) GetBundleName() string { return "cisco_ios_xe" }

func (ciplocalpoolgroupentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable_Ciplocalpoolgroupentry) GetYangName() string { return "cIpLocalPoolGroupEntry" }

func (ciplocalpoolgroupentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable_Ciplocalpoolgroupentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciplocalpoolgroupentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable_Ciplocalpoolgroupentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciplocalpoolgroupentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable_Ciplocalpoolgroupentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciplocalpoolgroupentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable_Ciplocalpoolgroupentry) SetParent(parent types.Entity) { ciplocalpoolgroupentry.parent = parent }

func (ciplocalpoolgroupentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable_Ciplocalpoolgroupentry) GetParent() types.Entity { return ciplocalpoolgroupentry.parent }

func (ciplocalpoolgroupentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable_Ciplocalpoolgroupentry) GetParentYangName() string { return "cIpLocalPoolGroupTable" }

// CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable
// A table providing statistics for each IP local pool.
// 
// Entries in this table are created as the result of adding a new
// IP local pool to the cIpLocalPoolConfigTable.
// 
// Entries in this table are deleted as the result of removing all
// the address ranges that are contained in an IP local pool in the
// cIpLocalPoolConfigTable.
// 
// Entries in this table are uniquely indexed by the name of the IP
// local pool.
type CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry provides statistical information about a particular IP local
    // pool, and the total number of free and used addresses of all the ranges in
    // an IP local pool. The type is slice of
    // CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable_Ciplocalpoolstatsentry.
    Ciplocalpoolstatsentry []CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable_Ciplocalpoolstatsentry
}

func (ciplocalpoolstatstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable) GetFilter() yfilter.YFilter { return ciplocalpoolstatstable.YFilter }

func (ciplocalpoolstatstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable) SetFilter(yf yfilter.YFilter) { ciplocalpoolstatstable.YFilter = yf }

func (ciplocalpoolstatstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable) GetGoName(yname string) string {
    if yname == "cIpLocalPoolStatsEntry" { return "Ciplocalpoolstatsentry" }
    return ""
}

func (ciplocalpoolstatstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable) GetSegmentPath() string {
    return "cIpLocalPoolStatsTable"
}

func (ciplocalpoolstatstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cIpLocalPoolStatsEntry" {
        for _, c := range ciplocalpoolstatstable.Ciplocalpoolstatsentry {
            if ciplocalpoolstatstable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable_Ciplocalpoolstatsentry{}
        ciplocalpoolstatstable.Ciplocalpoolstatsentry = append(ciplocalpoolstatstable.Ciplocalpoolstatsentry, child)
        return &ciplocalpoolstatstable.Ciplocalpoolstatsentry[len(ciplocalpoolstatstable.Ciplocalpoolstatsentry)-1]
    }
    return nil
}

func (ciplocalpoolstatstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ciplocalpoolstatstable.Ciplocalpoolstatsentry {
        children[ciplocalpoolstatstable.Ciplocalpoolstatsentry[i].GetSegmentPath()] = &ciplocalpoolstatstable.Ciplocalpoolstatsentry[i]
    }
    return children
}

func (ciplocalpoolstatstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciplocalpoolstatstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable) GetBundleName() string { return "cisco_ios_xe" }

func (ciplocalpoolstatstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable) GetYangName() string { return "cIpLocalPoolStatsTable" }

func (ciplocalpoolstatstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciplocalpoolstatstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciplocalpoolstatstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciplocalpoolstatstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable) SetParent(parent types.Entity) { ciplocalpoolstatstable.parent = parent }

func (ciplocalpoolstatstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable) GetParent() types.Entity { return ciplocalpoolstatstable.parent }

func (ciplocalpoolstatstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable) GetParentYangName() string { return "CISCO-IP-LOCAL-POOL-MIB" }

// CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable_Ciplocalpoolstatsentry
// Each entry provides statistical information about a particular
// IP local pool, and the total number of free and used addresses
// of all the ranges in an IP local pool.
type CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable_Ciplocalpoolstatsentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 1..48. Refers to
    // cisco_ip_local_pool_mib.CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable_Ciplocalpoolconfigentry_Ciplocalpoolname
    Ciplocalpoolname interface{}

    // The number of IP addresses available for use in this IP local pool. The
    // type is interface{} with range: 0..4294967295.
    Ciplocalpoolstatfreeaddrs interface{}

    // The number of IP addresses being used in this IP local pool. The type is
    // interface{} with range: 0..4294967295.
    Ciplocalpoolstatinuseaddrs interface{}

    // This object contains the high water mark of used addresses in an IP local
    // pool since pool creation, since the system was restarted, or since this
    // object was reset, whichever occurred last.  This object can only be set to
    // zero, and by doing so, it is reset to the value of
    // cIpLocalPoolStatInUseAddrs.  Since the number of addresses in a pool can be
    // reduced (e.g. by deleting one of its ranges), the value of this object may
    // be greater than the sum of cIpLocalPoolStatFreeAddrs and
    // cIpLocalPoolStatInUseAddrs. The type is interface{} with range:
    // 0..4294967295.
    Ciplocalpoolstathiwaterusedaddrs interface{}

    // When the number of used addresses in an IP local pool falls below this
    // threshold value, the ciscoIpLocalPoolInUseAddrNoti notification will be
    // rearmed.  The value of this object should never be greater than the value
    // of cIpLocalPoolStatInUseAddrThldHi. The type is interface{} with range:
    // 0..4294967295.
    Ciplocalpoolstatinuseaddrthldlo interface{}

    // When the number of used addresses in an IP local pool is equal or exceeds
    // this threshold value, a ciscoIpLocalPoolInUseAddrNoti notification will be
    // generated. Once this notification is generated, it will be disarmed and it
    // will not be generated again until the number of used address falls below
    // the value indicated by cIpLocalPoolStatInUseAddrThldLo.  The value of this
    // object should never be smaller than the value of
    // cIpLocalPoolStatInUseAddrThldLo. The type is interface{} with range:
    // 0..4294967295.
    Ciplocalpoolstatinuseaddrthldhi interface{}

    // When the percentage of used addresses in an IP local pool falls below this
    // threshold value, a cilpPercentAddrUsedLoNotif notification will be
    // generated.  Once the notification is generated, it will be disarmed and it
    // will not be generated again until the number of used addresses equals or
    // exceeds the value indicated by cIpLocalPoolPercentAddrThldHi.  The value of
    // this object should never be greater than the value of
    // cIpLocalPoolPercentAddrThldHi. The type is interface{} with range: 0..100.
    Ciplocalpoolpercentaddrthldlo interface{}

    // When the percentage of used addresses in an IP local pool is equal or
    // exceeds this threshold value, a cilpPercentAddrUsedHiNotif notification
    // will be generated. Once the notification is generated, it will be disarmed
    // and it will not be generated again until the number of used addresses falls
    // below the value indicated by cIpLocalPoolPercentAddrThldLo.  The value of
    // this object should never be smaller than the value of
    // cIpLocalPoolPercentAddrThldLo. The type is interface{} with range: 0..100.
    Ciplocalpoolpercentaddrthldhi interface{}
}

func (ciplocalpoolstatsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable_Ciplocalpoolstatsentry) GetFilter() yfilter.YFilter { return ciplocalpoolstatsentry.YFilter }

func (ciplocalpoolstatsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable_Ciplocalpoolstatsentry) SetFilter(yf yfilter.YFilter) { ciplocalpoolstatsentry.YFilter = yf }

func (ciplocalpoolstatsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable_Ciplocalpoolstatsentry) GetGoName(yname string) string {
    if yname == "cIpLocalPoolName" { return "Ciplocalpoolname" }
    if yname == "cIpLocalPoolStatFreeAddrs" { return "Ciplocalpoolstatfreeaddrs" }
    if yname == "cIpLocalPoolStatInUseAddrs" { return "Ciplocalpoolstatinuseaddrs" }
    if yname == "cIpLocalPoolStatHiWaterUsedAddrs" { return "Ciplocalpoolstathiwaterusedaddrs" }
    if yname == "cIpLocalPoolStatInUseAddrThldLo" { return "Ciplocalpoolstatinuseaddrthldlo" }
    if yname == "cIpLocalPoolStatInUseAddrThldHi" { return "Ciplocalpoolstatinuseaddrthldhi" }
    if yname == "cIpLocalPoolPercentAddrThldLo" { return "Ciplocalpoolpercentaddrthldlo" }
    if yname == "cIpLocalPoolPercentAddrThldHi" { return "Ciplocalpoolpercentaddrthldhi" }
    return ""
}

func (ciplocalpoolstatsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable_Ciplocalpoolstatsentry) GetSegmentPath() string {
    return "cIpLocalPoolStatsEntry" + "[cIpLocalPoolName='" + fmt.Sprintf("%v", ciplocalpoolstatsentry.Ciplocalpoolname) + "']"
}

func (ciplocalpoolstatsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable_Ciplocalpoolstatsentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciplocalpoolstatsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable_Ciplocalpoolstatsentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciplocalpoolstatsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable_Ciplocalpoolstatsentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cIpLocalPoolName"] = ciplocalpoolstatsentry.Ciplocalpoolname
    leafs["cIpLocalPoolStatFreeAddrs"] = ciplocalpoolstatsentry.Ciplocalpoolstatfreeaddrs
    leafs["cIpLocalPoolStatInUseAddrs"] = ciplocalpoolstatsentry.Ciplocalpoolstatinuseaddrs
    leafs["cIpLocalPoolStatHiWaterUsedAddrs"] = ciplocalpoolstatsentry.Ciplocalpoolstathiwaterusedaddrs
    leafs["cIpLocalPoolStatInUseAddrThldLo"] = ciplocalpoolstatsentry.Ciplocalpoolstatinuseaddrthldlo
    leafs["cIpLocalPoolStatInUseAddrThldHi"] = ciplocalpoolstatsentry.Ciplocalpoolstatinuseaddrthldhi
    leafs["cIpLocalPoolPercentAddrThldLo"] = ciplocalpoolstatsentry.Ciplocalpoolpercentaddrthldlo
    leafs["cIpLocalPoolPercentAddrThldHi"] = ciplocalpoolstatsentry.Ciplocalpoolpercentaddrthldhi
    return leafs
}

func (ciplocalpoolstatsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable_Ciplocalpoolstatsentry) GetBundleName() string { return "cisco_ios_xe" }

func (ciplocalpoolstatsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable_Ciplocalpoolstatsentry) GetYangName() string { return "cIpLocalPoolStatsEntry" }

func (ciplocalpoolstatsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable_Ciplocalpoolstatsentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciplocalpoolstatsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable_Ciplocalpoolstatsentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciplocalpoolstatsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable_Ciplocalpoolstatsentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciplocalpoolstatsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable_Ciplocalpoolstatsentry) SetParent(parent types.Entity) { ciplocalpoolstatsentry.parent = parent }

func (ciplocalpoolstatsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable_Ciplocalpoolstatsentry) GetParent() types.Entity { return ciplocalpoolstatsentry.parent }

func (ciplocalpoolstatsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable_Ciplocalpoolstatsentry) GetParentYangName() string { return "cIpLocalPoolStatsTable" }

// CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable
// This table lists all addresses that have been allocated out of
// an IP local pool.
// 
// Entries in this table are created when a remote peer allocates
// an address from one of the IP local pools in the
// cIpLocalPoolConfigTable.
// 
// Entries in this table are deleted when a remote peer deallocates
// an address from one of the IP local pool in the
// cIpLocalPoolConfigTable.
// 
// Entries in this table are uniquely indexed by the name of the IP
// local pool, and the allocated address, together with its address
// type.
type CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry refers to conceptual row that associates an IP addresses with
    // the interface where the request was received, and the user that requested
    // the address. The type is slice of
    // CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable_Ciplocalpoolallocentry.
    Ciplocalpoolallocentry []CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable_Ciplocalpoolallocentry
}

func (ciplocalpoolalloctable *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable) GetFilter() yfilter.YFilter { return ciplocalpoolalloctable.YFilter }

func (ciplocalpoolalloctable *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable) SetFilter(yf yfilter.YFilter) { ciplocalpoolalloctable.YFilter = yf }

func (ciplocalpoolalloctable *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable) GetGoName(yname string) string {
    if yname == "cIpLocalPoolAllocEntry" { return "Ciplocalpoolallocentry" }
    return ""
}

func (ciplocalpoolalloctable *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable) GetSegmentPath() string {
    return "cIpLocalPoolAllocTable"
}

func (ciplocalpoolalloctable *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cIpLocalPoolAllocEntry" {
        for _, c := range ciplocalpoolalloctable.Ciplocalpoolallocentry {
            if ciplocalpoolalloctable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable_Ciplocalpoolallocentry{}
        ciplocalpoolalloctable.Ciplocalpoolallocentry = append(ciplocalpoolalloctable.Ciplocalpoolallocentry, child)
        return &ciplocalpoolalloctable.Ciplocalpoolallocentry[len(ciplocalpoolalloctable.Ciplocalpoolallocentry)-1]
    }
    return nil
}

func (ciplocalpoolalloctable *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ciplocalpoolalloctable.Ciplocalpoolallocentry {
        children[ciplocalpoolalloctable.Ciplocalpoolallocentry[i].GetSegmentPath()] = &ciplocalpoolalloctable.Ciplocalpoolallocentry[i]
    }
    return children
}

func (ciplocalpoolalloctable *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciplocalpoolalloctable *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable) GetBundleName() string { return "cisco_ios_xe" }

func (ciplocalpoolalloctable *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable) GetYangName() string { return "cIpLocalPoolAllocTable" }

func (ciplocalpoolalloctable *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciplocalpoolalloctable *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciplocalpoolalloctable *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciplocalpoolalloctable *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable) SetParent(parent types.Entity) { ciplocalpoolalloctable.parent = parent }

func (ciplocalpoolalloctable *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable) GetParent() types.Entity { return ciplocalpoolalloctable.parent }

func (ciplocalpoolalloctable *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable) GetParentYangName() string { return "CISCO-IP-LOCAL-POOL-MIB" }

// CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable_Ciplocalpoolallocentry
// Each entry refers to conceptual row that associates an IP
// addresses with the interface where the request was received, and
// the user that requested the address.
type CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable_Ciplocalpoolallocentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 1..48. Refers to
    // cisco_ip_local_pool_mib.CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable_Ciplocalpoolconfigentry_Ciplocalpoolname
    Ciplocalpoolname interface{}

    // This attribute is a key. This object specifies the address type of
    // cIpLocalPoolAllocAddr. The type is InetAddressType.
    Ciplocalpoolallocaddrtype interface{}

    // This attribute is a key. This object specifies the allocated IP address. 
    // The address type of this object is described by cIpLocalPoolAllocAddrType.
    // The type is string with length: 0..255.
    Ciplocalpoolallocaddr interface{}

    // This object indicates the interface from which the allocation message was
    // sent.  In the case that the interface can not be determined, the value of
    // this object will be zero. The type is interface{} with range:
    // 0..2147483647.
    Ciplocalpoolallocifindex interface{}

    // This object indicates the user name of the person from whom the allocation
    // message was sent.  In the case that the user name can not be determined,
    // the value of this object will the null string. The type is string.
    Ciplocalpoolallocuser interface{}
}

func (ciplocalpoolallocentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable_Ciplocalpoolallocentry) GetFilter() yfilter.YFilter { return ciplocalpoolallocentry.YFilter }

func (ciplocalpoolallocentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable_Ciplocalpoolallocentry) SetFilter(yf yfilter.YFilter) { ciplocalpoolallocentry.YFilter = yf }

func (ciplocalpoolallocentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable_Ciplocalpoolallocentry) GetGoName(yname string) string {
    if yname == "cIpLocalPoolName" { return "Ciplocalpoolname" }
    if yname == "cIpLocalPoolAllocAddrType" { return "Ciplocalpoolallocaddrtype" }
    if yname == "cIpLocalPoolAllocAddr" { return "Ciplocalpoolallocaddr" }
    if yname == "cIpLocalPoolAllocIfIndex" { return "Ciplocalpoolallocifindex" }
    if yname == "cIpLocalPoolAllocUser" { return "Ciplocalpoolallocuser" }
    return ""
}

func (ciplocalpoolallocentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable_Ciplocalpoolallocentry) GetSegmentPath() string {
    return "cIpLocalPoolAllocEntry" + "[cIpLocalPoolName='" + fmt.Sprintf("%v", ciplocalpoolallocentry.Ciplocalpoolname) + "']" + "[cIpLocalPoolAllocAddrType='" + fmt.Sprintf("%v", ciplocalpoolallocentry.Ciplocalpoolallocaddrtype) + "']" + "[cIpLocalPoolAllocAddr='" + fmt.Sprintf("%v", ciplocalpoolallocentry.Ciplocalpoolallocaddr) + "']"
}

func (ciplocalpoolallocentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable_Ciplocalpoolallocentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciplocalpoolallocentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable_Ciplocalpoolallocentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciplocalpoolallocentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable_Ciplocalpoolallocentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cIpLocalPoolName"] = ciplocalpoolallocentry.Ciplocalpoolname
    leafs["cIpLocalPoolAllocAddrType"] = ciplocalpoolallocentry.Ciplocalpoolallocaddrtype
    leafs["cIpLocalPoolAllocAddr"] = ciplocalpoolallocentry.Ciplocalpoolallocaddr
    leafs["cIpLocalPoolAllocIfIndex"] = ciplocalpoolallocentry.Ciplocalpoolallocifindex
    leafs["cIpLocalPoolAllocUser"] = ciplocalpoolallocentry.Ciplocalpoolallocuser
    return leafs
}

func (ciplocalpoolallocentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable_Ciplocalpoolallocentry) GetBundleName() string { return "cisco_ios_xe" }

func (ciplocalpoolallocentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable_Ciplocalpoolallocentry) GetYangName() string { return "cIpLocalPoolAllocEntry" }

func (ciplocalpoolallocentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable_Ciplocalpoolallocentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciplocalpoolallocentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable_Ciplocalpoolallocentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciplocalpoolallocentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable_Ciplocalpoolallocentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciplocalpoolallocentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable_Ciplocalpoolallocentry) SetParent(parent types.Entity) { ciplocalpoolallocentry.parent = parent }

func (ciplocalpoolallocentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable_Ciplocalpoolallocentry) GetParent() types.Entity { return ciplocalpoolallocentry.parent }

func (ciplocalpoolallocentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable_Ciplocalpoolallocentry) GetParentYangName() string { return "cIpLocalPoolAllocTable" }

