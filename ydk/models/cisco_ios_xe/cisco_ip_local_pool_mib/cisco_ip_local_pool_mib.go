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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CIpLocalPoolConfig CISCOIPLOCALPOOLMIB_CIpLocalPoolConfig

    // This table manages the creation, modification, and deletion of IP local
    // pools using the RowStatus textual convention.  An entry in this table
    // defines an IP address range that is associated with an IP local pool.  A
    // conceptual row in this table can not be modified while
    // cIpLocalPoolRowStatus is set to 'active'.  Since IP local pool names are
    // unique even when they belong to different groups, and addresses within a
    // group can not overlap, a row in this table is uniquely indexed by the pool
    // name, and by the low address of the IP local pool together with its address
    // type.
    CIpLocalPoolConfigTable CISCOIPLOCALPOOLMIB_CIpLocalPoolConfigTable

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
    CIpLocalPoolGroupContainsTable CISCOIPLOCALPOOLMIB_CIpLocalPoolGroupContainsTable

    // This table provides statistics for configured IP pool groups.  Entries in
    // this table are created as the result of adding a new IP pool group to the
    // cIpLocalPoolConfigTable.  Entries in this table are deleted as the result
    // of removing all IP local pools that are contained in an IP pool group in
    // the cIpLocalPoolConfigTable.  An entry in this table is uniquely indexed by
    // IP pool group name.
    CIpLocalPoolGroupTable CISCOIPLOCALPOOLMIB_CIpLocalPoolGroupTable

    // A table providing statistics for each IP local pool.  Entries in this table
    // are created as the result of adding a new IP local pool to the
    // cIpLocalPoolConfigTable.  Entries in this table are deleted as the result
    // of removing all the address ranges that are contained in an IP local pool
    // in the cIpLocalPoolConfigTable.  Entries in this table are uniquely indexed
    // by the name of the IP local pool.
    CIpLocalPoolStatsTable CISCOIPLOCALPOOLMIB_CIpLocalPoolStatsTable

    // This table lists all addresses that have been allocated out of an IP local
    // pool.  Entries in this table are created when a remote peer allocates an
    // address from one of the IP local pools in the cIpLocalPoolConfigTable. 
    // Entries in this table are deleted when a remote peer deallocates an address
    // from one of the IP local pool in the cIpLocalPoolConfigTable.  Entries in
    // this table are uniquely indexed by the name of the IP local pool, and the
    // allocated address, together with its address type.
    CIpLocalPoolAllocTable CISCOIPLOCALPOOLMIB_CIpLocalPoolAllocTable
}

func (cISCOIPLOCALPOOLMIB *CISCOIPLOCALPOOLMIB) GetEntityData() *types.CommonEntityData {
    cISCOIPLOCALPOOLMIB.EntityData.YFilter = cISCOIPLOCALPOOLMIB.YFilter
    cISCOIPLOCALPOOLMIB.EntityData.YangName = "CISCO-IP-LOCAL-POOL-MIB"
    cISCOIPLOCALPOOLMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIPLOCALPOOLMIB.EntityData.ParentYangName = "CISCO-IP-LOCAL-POOL-MIB"
    cISCOIPLOCALPOOLMIB.EntityData.SegmentPath = "CISCO-IP-LOCAL-POOL-MIB:CISCO-IP-LOCAL-POOL-MIB"
    cISCOIPLOCALPOOLMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIPLOCALPOOLMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIPLOCALPOOLMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIPLOCALPOOLMIB.EntityData.Children = types.NewOrderedMap()
    cISCOIPLOCALPOOLMIB.EntityData.Children.Append("cIpLocalPoolConfig", types.YChild{"CIpLocalPoolConfig", &cISCOIPLOCALPOOLMIB.CIpLocalPoolConfig})
    cISCOIPLOCALPOOLMIB.EntityData.Children.Append("cIpLocalPoolConfigTable", types.YChild{"CIpLocalPoolConfigTable", &cISCOIPLOCALPOOLMIB.CIpLocalPoolConfigTable})
    cISCOIPLOCALPOOLMIB.EntityData.Children.Append("cIpLocalPoolGroupContainsTable", types.YChild{"CIpLocalPoolGroupContainsTable", &cISCOIPLOCALPOOLMIB.CIpLocalPoolGroupContainsTable})
    cISCOIPLOCALPOOLMIB.EntityData.Children.Append("cIpLocalPoolGroupTable", types.YChild{"CIpLocalPoolGroupTable", &cISCOIPLOCALPOOLMIB.CIpLocalPoolGroupTable})
    cISCOIPLOCALPOOLMIB.EntityData.Children.Append("cIpLocalPoolStatsTable", types.YChild{"CIpLocalPoolStatsTable", &cISCOIPLOCALPOOLMIB.CIpLocalPoolStatsTable})
    cISCOIPLOCALPOOLMIB.EntityData.Children.Append("cIpLocalPoolAllocTable", types.YChild{"CIpLocalPoolAllocTable", &cISCOIPLOCALPOOLMIB.CIpLocalPoolAllocTable})
    cISCOIPLOCALPOOLMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOIPLOCALPOOLMIB.EntityData.YListKeys = []string {}

    return &(cISCOIPLOCALPOOLMIB.EntityData)
}

// CISCOIPLOCALPOOLMIB_CIpLocalPoolConfig
type CISCOIPLOCALPOOLMIB_CIpLocalPoolConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An indication of whether the notifications defined by the
    // ciscoIpLocalPoolNotifGroup are enabled. The type is bool.
    CIpLocalPoolNotificationsEnable interface{}
}

func (cIpLocalPoolConfig *CISCOIPLOCALPOOLMIB_CIpLocalPoolConfig) GetEntityData() *types.CommonEntityData {
    cIpLocalPoolConfig.EntityData.YFilter = cIpLocalPoolConfig.YFilter
    cIpLocalPoolConfig.EntityData.YangName = "cIpLocalPoolConfig"
    cIpLocalPoolConfig.EntityData.BundleName = "cisco_ios_xe"
    cIpLocalPoolConfig.EntityData.ParentYangName = "CISCO-IP-LOCAL-POOL-MIB"
    cIpLocalPoolConfig.EntityData.SegmentPath = "cIpLocalPoolConfig"
    cIpLocalPoolConfig.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cIpLocalPoolConfig.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cIpLocalPoolConfig.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cIpLocalPoolConfig.EntityData.Children = types.NewOrderedMap()
    cIpLocalPoolConfig.EntityData.Leafs = types.NewOrderedMap()
    cIpLocalPoolConfig.EntityData.Leafs.Append("cIpLocalPoolNotificationsEnable", types.YLeaf{"CIpLocalPoolNotificationsEnable", cIpLocalPoolConfig.CIpLocalPoolNotificationsEnable})

    cIpLocalPoolConfig.EntityData.YListKeys = []string {}

    return &(cIpLocalPoolConfig.EntityData)
}

// CISCOIPLOCALPOOLMIB_CIpLocalPoolConfigTable
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
type CISCOIPLOCALPOOLMIB_CIpLocalPoolConfigTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry provides information about a particular IP local pool, including
    // the number of free and used addresses and its priority. The type is slice
    // of CISCOIPLOCALPOOLMIB_CIpLocalPoolConfigTable_CIpLocalPoolConfigEntry.
    CIpLocalPoolConfigEntry []*CISCOIPLOCALPOOLMIB_CIpLocalPoolConfigTable_CIpLocalPoolConfigEntry
}

func (cIpLocalPoolConfigTable *CISCOIPLOCALPOOLMIB_CIpLocalPoolConfigTable) GetEntityData() *types.CommonEntityData {
    cIpLocalPoolConfigTable.EntityData.YFilter = cIpLocalPoolConfigTable.YFilter
    cIpLocalPoolConfigTable.EntityData.YangName = "cIpLocalPoolConfigTable"
    cIpLocalPoolConfigTable.EntityData.BundleName = "cisco_ios_xe"
    cIpLocalPoolConfigTable.EntityData.ParentYangName = "CISCO-IP-LOCAL-POOL-MIB"
    cIpLocalPoolConfigTable.EntityData.SegmentPath = "cIpLocalPoolConfigTable"
    cIpLocalPoolConfigTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cIpLocalPoolConfigTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cIpLocalPoolConfigTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cIpLocalPoolConfigTable.EntityData.Children = types.NewOrderedMap()
    cIpLocalPoolConfigTable.EntityData.Children.Append("cIpLocalPoolConfigEntry", types.YChild{"CIpLocalPoolConfigEntry", nil})
    for i := range cIpLocalPoolConfigTable.CIpLocalPoolConfigEntry {
        cIpLocalPoolConfigTable.EntityData.Children.Append(types.GetSegmentPath(cIpLocalPoolConfigTable.CIpLocalPoolConfigEntry[i]), types.YChild{"CIpLocalPoolConfigEntry", cIpLocalPoolConfigTable.CIpLocalPoolConfigEntry[i]})
    }
    cIpLocalPoolConfigTable.EntityData.Leafs = types.NewOrderedMap()

    cIpLocalPoolConfigTable.EntityData.YListKeys = []string {}

    return &(cIpLocalPoolConfigTable.EntityData)
}

// CISCOIPLOCALPOOLMIB_CIpLocalPoolConfigTable_CIpLocalPoolConfigEntry
// Each entry provides information about a particular IP local
// pool, including the number of free and used addresses and its priority.
type CISCOIPLOCALPOOLMIB_CIpLocalPoolConfigTable_CIpLocalPoolConfigEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. An arbitrary non-empty string that uniquely
    // identifies the IP local pool.  This name must be unique among all the local
    // IP pools even when they belong to different pool groups. The type is string
    // with length: 1..48.
    CIpLocalPoolName interface{}

    // This attribute is a key. This object specifies the address type of
    // cIpLocalPoolAddressLo and cIpLocalPoolAddressHi. The type is
    // InetAddressType.
    CIpLocalPoolAddrType interface{}

    // This attribute is a key. This object specifies the first IP address of the
    // range of IP addresses contained by this pool entry.  The address type of
    // this object is described by cIpLocalPoolAddrType.  This address must be
    // less than or equal to the address in cIpLocalPoolAddressHi. The type is
    // string with length: 0..255.
    CIpLocalPoolAddressLo interface{}

    // This object specifies the last IP address of the range of IP addresses
    // mapped by this pool entry.  The address type of this object is described by
    // cIpLocalPoolAddrType.  If only a single address is being mapped, the value
    // of this object is equal to the value of cIpLocalPoolAddressLo. The type is
    // string with length: 0..255.
    CIpLocalPoolAddressHi interface{}

    // The number of IP addresses available for use in the range of IP addresses.
    // The type is interface{} with range: 0..4294967295.
    CIpLocalPoolFreeAddrs interface{}

    // The number of IP addresses being used in the range of IP addresses. The
    // type is interface{} with range: 0..4294967295.
    CIpLocalPoolInUseAddrs interface{}

    // This object relates an IP local pool to its IP pool group.  A null string
    // indicates this IP local pool is not contained in a named IP pool group, but
    // that it is contained in the base IP pool group.  An IP local pool can only
    // belong to one IP pool group. The type is string with length: 0..48.
    CIpLocalPoolGroupContainedIn interface{}

    // This object facilitates the creation, or deletion of a conceptual row in
    // this table. The type is RowStatus.
    CIpLocalPoolRowStatus interface{}

    // This object specifies priority of the IP local pool, where smaller value
    // indicates the lower priority. The priority value is used in assigning IP
    // Address  from local pools. The type is interface{} with range: 1..255.
    CIpLocalPoolPriority interface{}
}

func (cIpLocalPoolConfigEntry *CISCOIPLOCALPOOLMIB_CIpLocalPoolConfigTable_CIpLocalPoolConfigEntry) GetEntityData() *types.CommonEntityData {
    cIpLocalPoolConfigEntry.EntityData.YFilter = cIpLocalPoolConfigEntry.YFilter
    cIpLocalPoolConfigEntry.EntityData.YangName = "cIpLocalPoolConfigEntry"
    cIpLocalPoolConfigEntry.EntityData.BundleName = "cisco_ios_xe"
    cIpLocalPoolConfigEntry.EntityData.ParentYangName = "cIpLocalPoolConfigTable"
    cIpLocalPoolConfigEntry.EntityData.SegmentPath = "cIpLocalPoolConfigEntry" + types.AddKeyToken(cIpLocalPoolConfigEntry.CIpLocalPoolName, "cIpLocalPoolName") + types.AddKeyToken(cIpLocalPoolConfigEntry.CIpLocalPoolAddrType, "cIpLocalPoolAddrType") + types.AddKeyToken(cIpLocalPoolConfigEntry.CIpLocalPoolAddressLo, "cIpLocalPoolAddressLo")
    cIpLocalPoolConfigEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cIpLocalPoolConfigEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cIpLocalPoolConfigEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cIpLocalPoolConfigEntry.EntityData.Children = types.NewOrderedMap()
    cIpLocalPoolConfigEntry.EntityData.Leafs = types.NewOrderedMap()
    cIpLocalPoolConfigEntry.EntityData.Leafs.Append("cIpLocalPoolName", types.YLeaf{"CIpLocalPoolName", cIpLocalPoolConfigEntry.CIpLocalPoolName})
    cIpLocalPoolConfigEntry.EntityData.Leafs.Append("cIpLocalPoolAddrType", types.YLeaf{"CIpLocalPoolAddrType", cIpLocalPoolConfigEntry.CIpLocalPoolAddrType})
    cIpLocalPoolConfigEntry.EntityData.Leafs.Append("cIpLocalPoolAddressLo", types.YLeaf{"CIpLocalPoolAddressLo", cIpLocalPoolConfigEntry.CIpLocalPoolAddressLo})
    cIpLocalPoolConfigEntry.EntityData.Leafs.Append("cIpLocalPoolAddressHi", types.YLeaf{"CIpLocalPoolAddressHi", cIpLocalPoolConfigEntry.CIpLocalPoolAddressHi})
    cIpLocalPoolConfigEntry.EntityData.Leafs.Append("cIpLocalPoolFreeAddrs", types.YLeaf{"CIpLocalPoolFreeAddrs", cIpLocalPoolConfigEntry.CIpLocalPoolFreeAddrs})
    cIpLocalPoolConfigEntry.EntityData.Leafs.Append("cIpLocalPoolInUseAddrs", types.YLeaf{"CIpLocalPoolInUseAddrs", cIpLocalPoolConfigEntry.CIpLocalPoolInUseAddrs})
    cIpLocalPoolConfigEntry.EntityData.Leafs.Append("cIpLocalPoolGroupContainedIn", types.YLeaf{"CIpLocalPoolGroupContainedIn", cIpLocalPoolConfigEntry.CIpLocalPoolGroupContainedIn})
    cIpLocalPoolConfigEntry.EntityData.Leafs.Append("cIpLocalPoolRowStatus", types.YLeaf{"CIpLocalPoolRowStatus", cIpLocalPoolConfigEntry.CIpLocalPoolRowStatus})
    cIpLocalPoolConfigEntry.EntityData.Leafs.Append("cIpLocalPoolPriority", types.YLeaf{"CIpLocalPoolPriority", cIpLocalPoolConfigEntry.CIpLocalPoolPriority})

    cIpLocalPoolConfigEntry.EntityData.YListKeys = []string {"CIpLocalPoolName", "CIpLocalPoolAddrType", "CIpLocalPoolAddressLo"}

    return &(cIpLocalPoolConfigEntry.EntityData)
}

// CISCOIPLOCALPOOLMIB_CIpLocalPoolGroupContainsTable
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
type CISCOIPLOCALPOOLMIB_CIpLocalPoolGroupContainsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry describes single container/'containee' relationship.  Pool names
    // can only be associated with one group.  Pools carry implicit group
    // identifiers because pool names can only be associated with one group.  An
    // entry in this table describes such an association. The type is slice of
    // CISCOIPLOCALPOOLMIB_CIpLocalPoolGroupContainsTable_CIpLocalPoolGroupContainsEntry.
    CIpLocalPoolGroupContainsEntry []*CISCOIPLOCALPOOLMIB_CIpLocalPoolGroupContainsTable_CIpLocalPoolGroupContainsEntry
}

func (cIpLocalPoolGroupContainsTable *CISCOIPLOCALPOOLMIB_CIpLocalPoolGroupContainsTable) GetEntityData() *types.CommonEntityData {
    cIpLocalPoolGroupContainsTable.EntityData.YFilter = cIpLocalPoolGroupContainsTable.YFilter
    cIpLocalPoolGroupContainsTable.EntityData.YangName = "cIpLocalPoolGroupContainsTable"
    cIpLocalPoolGroupContainsTable.EntityData.BundleName = "cisco_ios_xe"
    cIpLocalPoolGroupContainsTable.EntityData.ParentYangName = "CISCO-IP-LOCAL-POOL-MIB"
    cIpLocalPoolGroupContainsTable.EntityData.SegmentPath = "cIpLocalPoolGroupContainsTable"
    cIpLocalPoolGroupContainsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cIpLocalPoolGroupContainsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cIpLocalPoolGroupContainsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cIpLocalPoolGroupContainsTable.EntityData.Children = types.NewOrderedMap()
    cIpLocalPoolGroupContainsTable.EntityData.Children.Append("cIpLocalPoolGroupContainsEntry", types.YChild{"CIpLocalPoolGroupContainsEntry", nil})
    for i := range cIpLocalPoolGroupContainsTable.CIpLocalPoolGroupContainsEntry {
        cIpLocalPoolGroupContainsTable.EntityData.Children.Append(types.GetSegmentPath(cIpLocalPoolGroupContainsTable.CIpLocalPoolGroupContainsEntry[i]), types.YChild{"CIpLocalPoolGroupContainsEntry", cIpLocalPoolGroupContainsTable.CIpLocalPoolGroupContainsEntry[i]})
    }
    cIpLocalPoolGroupContainsTable.EntityData.Leafs = types.NewOrderedMap()

    cIpLocalPoolGroupContainsTable.EntityData.YListKeys = []string {}

    return &(cIpLocalPoolGroupContainsTable.EntityData)
}

// CISCOIPLOCALPOOLMIB_CIpLocalPoolGroupContainsTable_CIpLocalPoolGroupContainsEntry
// Each entry describes single container/'containee'
// relationship.
// 
// Pool names can only be associated with one group.  Pools carry
// implicit group identifiers because pool names can only be
// associated with one group.  An entry in this table describes
// such an association.
type CISCOIPLOCALPOOLMIB_CIpLocalPoolGroupContainsTable_CIpLocalPoolGroupContainsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A unique group name that identifies the IP pool
    // group.  The null string represents the base IP pool group. The type is
    // string with length: 0..48.
    CIpLocalPoolGroupName interface{}

    // This attribute is a key. The value of cIpLocalPoolName for the contained IP
    // local pool. The type is string with length: 1..48.
    CIpLocalPoolChildIndex interface{}
}

func (cIpLocalPoolGroupContainsEntry *CISCOIPLOCALPOOLMIB_CIpLocalPoolGroupContainsTable_CIpLocalPoolGroupContainsEntry) GetEntityData() *types.CommonEntityData {
    cIpLocalPoolGroupContainsEntry.EntityData.YFilter = cIpLocalPoolGroupContainsEntry.YFilter
    cIpLocalPoolGroupContainsEntry.EntityData.YangName = "cIpLocalPoolGroupContainsEntry"
    cIpLocalPoolGroupContainsEntry.EntityData.BundleName = "cisco_ios_xe"
    cIpLocalPoolGroupContainsEntry.EntityData.ParentYangName = "cIpLocalPoolGroupContainsTable"
    cIpLocalPoolGroupContainsEntry.EntityData.SegmentPath = "cIpLocalPoolGroupContainsEntry" + types.AddKeyToken(cIpLocalPoolGroupContainsEntry.CIpLocalPoolGroupName, "cIpLocalPoolGroupName") + types.AddKeyToken(cIpLocalPoolGroupContainsEntry.CIpLocalPoolChildIndex, "cIpLocalPoolChildIndex")
    cIpLocalPoolGroupContainsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cIpLocalPoolGroupContainsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cIpLocalPoolGroupContainsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cIpLocalPoolGroupContainsEntry.EntityData.Children = types.NewOrderedMap()
    cIpLocalPoolGroupContainsEntry.EntityData.Leafs = types.NewOrderedMap()
    cIpLocalPoolGroupContainsEntry.EntityData.Leafs.Append("cIpLocalPoolGroupName", types.YLeaf{"CIpLocalPoolGroupName", cIpLocalPoolGroupContainsEntry.CIpLocalPoolGroupName})
    cIpLocalPoolGroupContainsEntry.EntityData.Leafs.Append("cIpLocalPoolChildIndex", types.YLeaf{"CIpLocalPoolChildIndex", cIpLocalPoolGroupContainsEntry.CIpLocalPoolChildIndex})

    cIpLocalPoolGroupContainsEntry.EntityData.YListKeys = []string {"CIpLocalPoolGroupName", "CIpLocalPoolChildIndex"}

    return &(cIpLocalPoolGroupContainsEntry.EntityData)
}

// CISCOIPLOCALPOOLMIB_CIpLocalPoolGroupTable
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
type CISCOIPLOCALPOOLMIB_CIpLocalPoolGroupTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry provides information about a particular IP pool group and the
    // number of free and used addresses in an IP pool group. The type is slice of
    // CISCOIPLOCALPOOLMIB_CIpLocalPoolGroupTable_CIpLocalPoolGroupEntry.
    CIpLocalPoolGroupEntry []*CISCOIPLOCALPOOLMIB_CIpLocalPoolGroupTable_CIpLocalPoolGroupEntry
}

func (cIpLocalPoolGroupTable *CISCOIPLOCALPOOLMIB_CIpLocalPoolGroupTable) GetEntityData() *types.CommonEntityData {
    cIpLocalPoolGroupTable.EntityData.YFilter = cIpLocalPoolGroupTable.YFilter
    cIpLocalPoolGroupTable.EntityData.YangName = "cIpLocalPoolGroupTable"
    cIpLocalPoolGroupTable.EntityData.BundleName = "cisco_ios_xe"
    cIpLocalPoolGroupTable.EntityData.ParentYangName = "CISCO-IP-LOCAL-POOL-MIB"
    cIpLocalPoolGroupTable.EntityData.SegmentPath = "cIpLocalPoolGroupTable"
    cIpLocalPoolGroupTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cIpLocalPoolGroupTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cIpLocalPoolGroupTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cIpLocalPoolGroupTable.EntityData.Children = types.NewOrderedMap()
    cIpLocalPoolGroupTable.EntityData.Children.Append("cIpLocalPoolGroupEntry", types.YChild{"CIpLocalPoolGroupEntry", nil})
    for i := range cIpLocalPoolGroupTable.CIpLocalPoolGroupEntry {
        cIpLocalPoolGroupTable.EntityData.Children.Append(types.GetSegmentPath(cIpLocalPoolGroupTable.CIpLocalPoolGroupEntry[i]), types.YChild{"CIpLocalPoolGroupEntry", cIpLocalPoolGroupTable.CIpLocalPoolGroupEntry[i]})
    }
    cIpLocalPoolGroupTable.EntityData.Leafs = types.NewOrderedMap()

    cIpLocalPoolGroupTable.EntityData.YListKeys = []string {}

    return &(cIpLocalPoolGroupTable.EntityData)
}

// CISCOIPLOCALPOOLMIB_CIpLocalPoolGroupTable_CIpLocalPoolGroupEntry
// Each entry provides information about a particular IP pool
// group and the number of free and used addresses in an IP pool
// group.
type CISCOIPLOCALPOOLMIB_CIpLocalPoolGroupTable_CIpLocalPoolGroupEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 0..48. Refers to
    // cisco_ip_local_pool_mib.CISCOIPLOCALPOOLMIB_CIpLocalPoolGroupContainsTable_CIpLocalPoolGroupContainsEntry_CIpLocalPoolGroupName
    CIpLocalPoolGroupName interface{}

    // The number of IP addresses available for use in the IP pool group. The type
    // is interface{} with range: 0..4294967295.
    CIpLocalPoolGroupFreeAddrs interface{}

    // The number of IP addresses that have been allocated from the IP pool group.
    // The type is interface{} with range: 0..4294967295.
    CIpLocalPoolGroupInUseAddrs interface{}
}

func (cIpLocalPoolGroupEntry *CISCOIPLOCALPOOLMIB_CIpLocalPoolGroupTable_CIpLocalPoolGroupEntry) GetEntityData() *types.CommonEntityData {
    cIpLocalPoolGroupEntry.EntityData.YFilter = cIpLocalPoolGroupEntry.YFilter
    cIpLocalPoolGroupEntry.EntityData.YangName = "cIpLocalPoolGroupEntry"
    cIpLocalPoolGroupEntry.EntityData.BundleName = "cisco_ios_xe"
    cIpLocalPoolGroupEntry.EntityData.ParentYangName = "cIpLocalPoolGroupTable"
    cIpLocalPoolGroupEntry.EntityData.SegmentPath = "cIpLocalPoolGroupEntry" + types.AddKeyToken(cIpLocalPoolGroupEntry.CIpLocalPoolGroupName, "cIpLocalPoolGroupName")
    cIpLocalPoolGroupEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cIpLocalPoolGroupEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cIpLocalPoolGroupEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cIpLocalPoolGroupEntry.EntityData.Children = types.NewOrderedMap()
    cIpLocalPoolGroupEntry.EntityData.Leafs = types.NewOrderedMap()
    cIpLocalPoolGroupEntry.EntityData.Leafs.Append("cIpLocalPoolGroupName", types.YLeaf{"CIpLocalPoolGroupName", cIpLocalPoolGroupEntry.CIpLocalPoolGroupName})
    cIpLocalPoolGroupEntry.EntityData.Leafs.Append("cIpLocalPoolGroupFreeAddrs", types.YLeaf{"CIpLocalPoolGroupFreeAddrs", cIpLocalPoolGroupEntry.CIpLocalPoolGroupFreeAddrs})
    cIpLocalPoolGroupEntry.EntityData.Leafs.Append("cIpLocalPoolGroupInUseAddrs", types.YLeaf{"CIpLocalPoolGroupInUseAddrs", cIpLocalPoolGroupEntry.CIpLocalPoolGroupInUseAddrs})

    cIpLocalPoolGroupEntry.EntityData.YListKeys = []string {"CIpLocalPoolGroupName"}

    return &(cIpLocalPoolGroupEntry.EntityData)
}

// CISCOIPLOCALPOOLMIB_CIpLocalPoolStatsTable
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
type CISCOIPLOCALPOOLMIB_CIpLocalPoolStatsTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry provides statistical information about a particular IP local
    // pool, and the total number of free and used addresses of all the ranges in
    // an IP local pool. The type is slice of
    // CISCOIPLOCALPOOLMIB_CIpLocalPoolStatsTable_CIpLocalPoolStatsEntry.
    CIpLocalPoolStatsEntry []*CISCOIPLOCALPOOLMIB_CIpLocalPoolStatsTable_CIpLocalPoolStatsEntry
}

func (cIpLocalPoolStatsTable *CISCOIPLOCALPOOLMIB_CIpLocalPoolStatsTable) GetEntityData() *types.CommonEntityData {
    cIpLocalPoolStatsTable.EntityData.YFilter = cIpLocalPoolStatsTable.YFilter
    cIpLocalPoolStatsTable.EntityData.YangName = "cIpLocalPoolStatsTable"
    cIpLocalPoolStatsTable.EntityData.BundleName = "cisco_ios_xe"
    cIpLocalPoolStatsTable.EntityData.ParentYangName = "CISCO-IP-LOCAL-POOL-MIB"
    cIpLocalPoolStatsTable.EntityData.SegmentPath = "cIpLocalPoolStatsTable"
    cIpLocalPoolStatsTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cIpLocalPoolStatsTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cIpLocalPoolStatsTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cIpLocalPoolStatsTable.EntityData.Children = types.NewOrderedMap()
    cIpLocalPoolStatsTable.EntityData.Children.Append("cIpLocalPoolStatsEntry", types.YChild{"CIpLocalPoolStatsEntry", nil})
    for i := range cIpLocalPoolStatsTable.CIpLocalPoolStatsEntry {
        cIpLocalPoolStatsTable.EntityData.Children.Append(types.GetSegmentPath(cIpLocalPoolStatsTable.CIpLocalPoolStatsEntry[i]), types.YChild{"CIpLocalPoolStatsEntry", cIpLocalPoolStatsTable.CIpLocalPoolStatsEntry[i]})
    }
    cIpLocalPoolStatsTable.EntityData.Leafs = types.NewOrderedMap()

    cIpLocalPoolStatsTable.EntityData.YListKeys = []string {}

    return &(cIpLocalPoolStatsTable.EntityData)
}

// CISCOIPLOCALPOOLMIB_CIpLocalPoolStatsTable_CIpLocalPoolStatsEntry
// Each entry provides statistical information about a particular
// IP local pool, and the total number of free and used addresses
// of all the ranges in an IP local pool.
type CISCOIPLOCALPOOLMIB_CIpLocalPoolStatsTable_CIpLocalPoolStatsEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 1..48. Refers to
    // cisco_ip_local_pool_mib.CISCOIPLOCALPOOLMIB_CIpLocalPoolConfigTable_CIpLocalPoolConfigEntry_CIpLocalPoolName
    CIpLocalPoolName interface{}

    // The number of IP addresses available for use in this IP local pool. The
    // type is interface{} with range: 0..4294967295.
    CIpLocalPoolStatFreeAddrs interface{}

    // The number of IP addresses being used in this IP local pool. The type is
    // interface{} with range: 0..4294967295.
    CIpLocalPoolStatInUseAddrs interface{}

    // This object contains the high water mark of used addresses in an IP local
    // pool since pool creation, since the system was restarted, or since this
    // object was reset, whichever occurred last.  This object can only be set to
    // zero, and by doing so, it is reset to the value of
    // cIpLocalPoolStatInUseAddrs.  Since the number of addresses in a pool can be
    // reduced (e.g. by deleting one of its ranges), the value of this object may
    // be greater than the sum of cIpLocalPoolStatFreeAddrs and
    // cIpLocalPoolStatInUseAddrs. The type is interface{} with range:
    // 0..4294967295.
    CIpLocalPoolStatHiWaterUsedAddrs interface{}

    // When the number of used addresses in an IP local pool falls below this
    // threshold value, the ciscoIpLocalPoolInUseAddrNoti notification will be
    // rearmed.  The value of this object should never be greater than the value
    // of cIpLocalPoolStatInUseAddrThldHi. The type is interface{} with range:
    // 0..4294967295.
    CIpLocalPoolStatInUseAddrThldLo interface{}

    // When the number of used addresses in an IP local pool is equal or exceeds
    // this threshold value, a ciscoIpLocalPoolInUseAddrNoti notification will be
    // generated. Once this notification is generated, it will be disarmed and it
    // will not be generated again until the number of used address falls below
    // the value indicated by cIpLocalPoolStatInUseAddrThldLo.  The value of this
    // object should never be smaller than the value of
    // cIpLocalPoolStatInUseAddrThldLo. The type is interface{} with range:
    // 0..4294967295.
    CIpLocalPoolStatInUseAddrThldHi interface{}

    // When the percentage of used addresses in an IP local pool falls below this
    // threshold value, a cilpPercentAddrUsedLoNotif notification will be
    // generated.  Once the notification is generated, it will be disarmed and it
    // will not be generated again until the number of used addresses equals or
    // exceeds the value indicated by cIpLocalPoolPercentAddrThldHi.  The value of
    // this object should never be greater than the value of
    // cIpLocalPoolPercentAddrThldHi. The type is interface{} with range: 0..100.
    CIpLocalPoolPercentAddrThldLo interface{}

    // When the percentage of used addresses in an IP local pool is equal or
    // exceeds this threshold value, a cilpPercentAddrUsedHiNotif notification
    // will be generated. Once the notification is generated, it will be disarmed
    // and it will not be generated again until the number of used addresses falls
    // below the value indicated by cIpLocalPoolPercentAddrThldLo.  The value of
    // this object should never be smaller than the value of
    // cIpLocalPoolPercentAddrThldLo. The type is interface{} with range: 0..100.
    CIpLocalPoolPercentAddrThldHi interface{}
}

func (cIpLocalPoolStatsEntry *CISCOIPLOCALPOOLMIB_CIpLocalPoolStatsTable_CIpLocalPoolStatsEntry) GetEntityData() *types.CommonEntityData {
    cIpLocalPoolStatsEntry.EntityData.YFilter = cIpLocalPoolStatsEntry.YFilter
    cIpLocalPoolStatsEntry.EntityData.YangName = "cIpLocalPoolStatsEntry"
    cIpLocalPoolStatsEntry.EntityData.BundleName = "cisco_ios_xe"
    cIpLocalPoolStatsEntry.EntityData.ParentYangName = "cIpLocalPoolStatsTable"
    cIpLocalPoolStatsEntry.EntityData.SegmentPath = "cIpLocalPoolStatsEntry" + types.AddKeyToken(cIpLocalPoolStatsEntry.CIpLocalPoolName, "cIpLocalPoolName")
    cIpLocalPoolStatsEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cIpLocalPoolStatsEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cIpLocalPoolStatsEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cIpLocalPoolStatsEntry.EntityData.Children = types.NewOrderedMap()
    cIpLocalPoolStatsEntry.EntityData.Leafs = types.NewOrderedMap()
    cIpLocalPoolStatsEntry.EntityData.Leafs.Append("cIpLocalPoolName", types.YLeaf{"CIpLocalPoolName", cIpLocalPoolStatsEntry.CIpLocalPoolName})
    cIpLocalPoolStatsEntry.EntityData.Leafs.Append("cIpLocalPoolStatFreeAddrs", types.YLeaf{"CIpLocalPoolStatFreeAddrs", cIpLocalPoolStatsEntry.CIpLocalPoolStatFreeAddrs})
    cIpLocalPoolStatsEntry.EntityData.Leafs.Append("cIpLocalPoolStatInUseAddrs", types.YLeaf{"CIpLocalPoolStatInUseAddrs", cIpLocalPoolStatsEntry.CIpLocalPoolStatInUseAddrs})
    cIpLocalPoolStatsEntry.EntityData.Leafs.Append("cIpLocalPoolStatHiWaterUsedAddrs", types.YLeaf{"CIpLocalPoolStatHiWaterUsedAddrs", cIpLocalPoolStatsEntry.CIpLocalPoolStatHiWaterUsedAddrs})
    cIpLocalPoolStatsEntry.EntityData.Leafs.Append("cIpLocalPoolStatInUseAddrThldLo", types.YLeaf{"CIpLocalPoolStatInUseAddrThldLo", cIpLocalPoolStatsEntry.CIpLocalPoolStatInUseAddrThldLo})
    cIpLocalPoolStatsEntry.EntityData.Leafs.Append("cIpLocalPoolStatInUseAddrThldHi", types.YLeaf{"CIpLocalPoolStatInUseAddrThldHi", cIpLocalPoolStatsEntry.CIpLocalPoolStatInUseAddrThldHi})
    cIpLocalPoolStatsEntry.EntityData.Leafs.Append("cIpLocalPoolPercentAddrThldLo", types.YLeaf{"CIpLocalPoolPercentAddrThldLo", cIpLocalPoolStatsEntry.CIpLocalPoolPercentAddrThldLo})
    cIpLocalPoolStatsEntry.EntityData.Leafs.Append("cIpLocalPoolPercentAddrThldHi", types.YLeaf{"CIpLocalPoolPercentAddrThldHi", cIpLocalPoolStatsEntry.CIpLocalPoolPercentAddrThldHi})

    cIpLocalPoolStatsEntry.EntityData.YListKeys = []string {"CIpLocalPoolName"}

    return &(cIpLocalPoolStatsEntry.EntityData)
}

// CISCOIPLOCALPOOLMIB_CIpLocalPoolAllocTable
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
type CISCOIPLOCALPOOLMIB_CIpLocalPoolAllocTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry refers to conceptual row that associates an IP addresses with
    // the interface where the request was received, and the user that requested
    // the address. The type is slice of
    // CISCOIPLOCALPOOLMIB_CIpLocalPoolAllocTable_CIpLocalPoolAllocEntry.
    CIpLocalPoolAllocEntry []*CISCOIPLOCALPOOLMIB_CIpLocalPoolAllocTable_CIpLocalPoolAllocEntry
}

func (cIpLocalPoolAllocTable *CISCOIPLOCALPOOLMIB_CIpLocalPoolAllocTable) GetEntityData() *types.CommonEntityData {
    cIpLocalPoolAllocTable.EntityData.YFilter = cIpLocalPoolAllocTable.YFilter
    cIpLocalPoolAllocTable.EntityData.YangName = "cIpLocalPoolAllocTable"
    cIpLocalPoolAllocTable.EntityData.BundleName = "cisco_ios_xe"
    cIpLocalPoolAllocTable.EntityData.ParentYangName = "CISCO-IP-LOCAL-POOL-MIB"
    cIpLocalPoolAllocTable.EntityData.SegmentPath = "cIpLocalPoolAllocTable"
    cIpLocalPoolAllocTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cIpLocalPoolAllocTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cIpLocalPoolAllocTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cIpLocalPoolAllocTable.EntityData.Children = types.NewOrderedMap()
    cIpLocalPoolAllocTable.EntityData.Children.Append("cIpLocalPoolAllocEntry", types.YChild{"CIpLocalPoolAllocEntry", nil})
    for i := range cIpLocalPoolAllocTable.CIpLocalPoolAllocEntry {
        cIpLocalPoolAllocTable.EntityData.Children.Append(types.GetSegmentPath(cIpLocalPoolAllocTable.CIpLocalPoolAllocEntry[i]), types.YChild{"CIpLocalPoolAllocEntry", cIpLocalPoolAllocTable.CIpLocalPoolAllocEntry[i]})
    }
    cIpLocalPoolAllocTable.EntityData.Leafs = types.NewOrderedMap()

    cIpLocalPoolAllocTable.EntityData.YListKeys = []string {}

    return &(cIpLocalPoolAllocTable.EntityData)
}

// CISCOIPLOCALPOOLMIB_CIpLocalPoolAllocTable_CIpLocalPoolAllocEntry
// Each entry refers to conceptual row that associates an IP
// addresses with the interface where the request was received, and
// the user that requested the address.
type CISCOIPLOCALPOOLMIB_CIpLocalPoolAllocTable_CIpLocalPoolAllocEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with length: 1..48. Refers to
    // cisco_ip_local_pool_mib.CISCOIPLOCALPOOLMIB_CIpLocalPoolConfigTable_CIpLocalPoolConfigEntry_CIpLocalPoolName
    CIpLocalPoolName interface{}

    // This attribute is a key. This object specifies the address type of
    // cIpLocalPoolAllocAddr. The type is InetAddressType.
    CIpLocalPoolAllocAddrType interface{}

    // This attribute is a key. This object specifies the allocated IP address. 
    // The address type of this object is described by cIpLocalPoolAllocAddrType.
    // The type is string with length: 0..255.
    CIpLocalPoolAllocAddr interface{}

    // This object indicates the interface from which the allocation message was
    // sent.  In the case that the interface can not be determined, the value of
    // this object will be zero. The type is interface{} with range:
    // 0..2147483647.
    CIpLocalPoolAllocIfIndex interface{}

    // This object indicates the user name of the person from whom the allocation
    // message was sent.  In the case that the user name can not be determined,
    // the value of this object will the null string. The type is string.
    CIpLocalPoolAllocUser interface{}
}

func (cIpLocalPoolAllocEntry *CISCOIPLOCALPOOLMIB_CIpLocalPoolAllocTable_CIpLocalPoolAllocEntry) GetEntityData() *types.CommonEntityData {
    cIpLocalPoolAllocEntry.EntityData.YFilter = cIpLocalPoolAllocEntry.YFilter
    cIpLocalPoolAllocEntry.EntityData.YangName = "cIpLocalPoolAllocEntry"
    cIpLocalPoolAllocEntry.EntityData.BundleName = "cisco_ios_xe"
    cIpLocalPoolAllocEntry.EntityData.ParentYangName = "cIpLocalPoolAllocTable"
    cIpLocalPoolAllocEntry.EntityData.SegmentPath = "cIpLocalPoolAllocEntry" + types.AddKeyToken(cIpLocalPoolAllocEntry.CIpLocalPoolName, "cIpLocalPoolName") + types.AddKeyToken(cIpLocalPoolAllocEntry.CIpLocalPoolAllocAddrType, "cIpLocalPoolAllocAddrType") + types.AddKeyToken(cIpLocalPoolAllocEntry.CIpLocalPoolAllocAddr, "cIpLocalPoolAllocAddr")
    cIpLocalPoolAllocEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cIpLocalPoolAllocEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cIpLocalPoolAllocEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cIpLocalPoolAllocEntry.EntityData.Children = types.NewOrderedMap()
    cIpLocalPoolAllocEntry.EntityData.Leafs = types.NewOrderedMap()
    cIpLocalPoolAllocEntry.EntityData.Leafs.Append("cIpLocalPoolName", types.YLeaf{"CIpLocalPoolName", cIpLocalPoolAllocEntry.CIpLocalPoolName})
    cIpLocalPoolAllocEntry.EntityData.Leafs.Append("cIpLocalPoolAllocAddrType", types.YLeaf{"CIpLocalPoolAllocAddrType", cIpLocalPoolAllocEntry.CIpLocalPoolAllocAddrType})
    cIpLocalPoolAllocEntry.EntityData.Leafs.Append("cIpLocalPoolAllocAddr", types.YLeaf{"CIpLocalPoolAllocAddr", cIpLocalPoolAllocEntry.CIpLocalPoolAllocAddr})
    cIpLocalPoolAllocEntry.EntityData.Leafs.Append("cIpLocalPoolAllocIfIndex", types.YLeaf{"CIpLocalPoolAllocIfIndex", cIpLocalPoolAllocEntry.CIpLocalPoolAllocIfIndex})
    cIpLocalPoolAllocEntry.EntityData.Leafs.Append("cIpLocalPoolAllocUser", types.YLeaf{"CIpLocalPoolAllocUser", cIpLocalPoolAllocEntry.CIpLocalPoolAllocUser})

    cIpLocalPoolAllocEntry.EntityData.YListKeys = []string {"CIpLocalPoolName", "CIpLocalPoolAllocAddrType", "CIpLocalPoolAllocAddr"}

    return &(cIpLocalPoolAllocEntry.EntityData)
}

