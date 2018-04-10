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

func (cISCOIPLOCALPOOLMIB *CISCOIPLOCALPOOLMIB) GetEntityData() *types.CommonEntityData {
    cISCOIPLOCALPOOLMIB.EntityData.YFilter = cISCOIPLOCALPOOLMIB.YFilter
    cISCOIPLOCALPOOLMIB.EntityData.YangName = "CISCO-IP-LOCAL-POOL-MIB"
    cISCOIPLOCALPOOLMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIPLOCALPOOLMIB.EntityData.ParentYangName = "CISCO-IP-LOCAL-POOL-MIB"
    cISCOIPLOCALPOOLMIB.EntityData.SegmentPath = "CISCO-IP-LOCAL-POOL-MIB:CISCO-IP-LOCAL-POOL-MIB"
    cISCOIPLOCALPOOLMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIPLOCALPOOLMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIPLOCALPOOLMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIPLOCALPOOLMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOIPLOCALPOOLMIB.EntityData.Children["cIpLocalPoolConfig"] = types.YChild{"Ciplocalpoolconfig", &cISCOIPLOCALPOOLMIB.Ciplocalpoolconfig}
    cISCOIPLOCALPOOLMIB.EntityData.Children["cIpLocalPoolConfigTable"] = types.YChild{"Ciplocalpoolconfigtable", &cISCOIPLOCALPOOLMIB.Ciplocalpoolconfigtable}
    cISCOIPLOCALPOOLMIB.EntityData.Children["cIpLocalPoolGroupContainsTable"] = types.YChild{"Ciplocalpoolgroupcontainstable", &cISCOIPLOCALPOOLMIB.Ciplocalpoolgroupcontainstable}
    cISCOIPLOCALPOOLMIB.EntityData.Children["cIpLocalPoolGroupTable"] = types.YChild{"Ciplocalpoolgrouptable", &cISCOIPLOCALPOOLMIB.Ciplocalpoolgrouptable}
    cISCOIPLOCALPOOLMIB.EntityData.Children["cIpLocalPoolStatsTable"] = types.YChild{"Ciplocalpoolstatstable", &cISCOIPLOCALPOOLMIB.Ciplocalpoolstatstable}
    cISCOIPLOCALPOOLMIB.EntityData.Children["cIpLocalPoolAllocTable"] = types.YChild{"Ciplocalpoolalloctable", &cISCOIPLOCALPOOLMIB.Ciplocalpoolalloctable}
    cISCOIPLOCALPOOLMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOIPLOCALPOOLMIB.EntityData)
}

// CISCOIPLOCALPOOLMIB_Ciplocalpoolconfig
type CISCOIPLOCALPOOLMIB_Ciplocalpoolconfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An indication of whether the notifications defined by the
    // ciscoIpLocalPoolNotifGroup are enabled. The type is bool.
    Ciplocalpoolnotificationsenable interface{}
}

func (ciplocalpoolconfig *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfig) GetEntityData() *types.CommonEntityData {
    ciplocalpoolconfig.EntityData.YFilter = ciplocalpoolconfig.YFilter
    ciplocalpoolconfig.EntityData.YangName = "cIpLocalPoolConfig"
    ciplocalpoolconfig.EntityData.BundleName = "cisco_ios_xe"
    ciplocalpoolconfig.EntityData.ParentYangName = "CISCO-IP-LOCAL-POOL-MIB"
    ciplocalpoolconfig.EntityData.SegmentPath = "cIpLocalPoolConfig"
    ciplocalpoolconfig.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciplocalpoolconfig.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciplocalpoolconfig.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciplocalpoolconfig.EntityData.Children = make(map[string]types.YChild)
    ciplocalpoolconfig.EntityData.Leafs = make(map[string]types.YLeaf)
    ciplocalpoolconfig.EntityData.Leafs["cIpLocalPoolNotificationsEnable"] = types.YLeaf{"Ciplocalpoolnotificationsenable", ciplocalpoolconfig.Ciplocalpoolnotificationsenable}
    return &(ciplocalpoolconfig.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry provides information about a particular IP local pool, including
    // the number of free and used addresses and its priority. The type is slice
    // of CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable_Ciplocalpoolconfigentry.
    Ciplocalpoolconfigentry []CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable_Ciplocalpoolconfigentry
}

func (ciplocalpoolconfigtable *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable) GetEntityData() *types.CommonEntityData {
    ciplocalpoolconfigtable.EntityData.YFilter = ciplocalpoolconfigtable.YFilter
    ciplocalpoolconfigtable.EntityData.YangName = "cIpLocalPoolConfigTable"
    ciplocalpoolconfigtable.EntityData.BundleName = "cisco_ios_xe"
    ciplocalpoolconfigtable.EntityData.ParentYangName = "CISCO-IP-LOCAL-POOL-MIB"
    ciplocalpoolconfigtable.EntityData.SegmentPath = "cIpLocalPoolConfigTable"
    ciplocalpoolconfigtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciplocalpoolconfigtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciplocalpoolconfigtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciplocalpoolconfigtable.EntityData.Children = make(map[string]types.YChild)
    ciplocalpoolconfigtable.EntityData.Children["cIpLocalPoolConfigEntry"] = types.YChild{"Ciplocalpoolconfigentry", nil}
    for i := range ciplocalpoolconfigtable.Ciplocalpoolconfigentry {
        ciplocalpoolconfigtable.EntityData.Children[types.GetSegmentPath(&ciplocalpoolconfigtable.Ciplocalpoolconfigentry[i])] = types.YChild{"Ciplocalpoolconfigentry", &ciplocalpoolconfigtable.Ciplocalpoolconfigentry[i]}
    }
    ciplocalpoolconfigtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ciplocalpoolconfigtable.EntityData)
}

// CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable_Ciplocalpoolconfigentry
// Each entry provides information about a particular IP local
// pool, including the number of free and used addresses and its priority.
type CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable_Ciplocalpoolconfigentry struct {
    EntityData types.CommonEntityData
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

func (ciplocalpoolconfigentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolconfigtable_Ciplocalpoolconfigentry) GetEntityData() *types.CommonEntityData {
    ciplocalpoolconfigentry.EntityData.YFilter = ciplocalpoolconfigentry.YFilter
    ciplocalpoolconfigentry.EntityData.YangName = "cIpLocalPoolConfigEntry"
    ciplocalpoolconfigentry.EntityData.BundleName = "cisco_ios_xe"
    ciplocalpoolconfigentry.EntityData.ParentYangName = "cIpLocalPoolConfigTable"
    ciplocalpoolconfigentry.EntityData.SegmentPath = "cIpLocalPoolConfigEntry" + "[cIpLocalPoolName='" + fmt.Sprintf("%v", ciplocalpoolconfigentry.Ciplocalpoolname) + "']" + "[cIpLocalPoolAddrType='" + fmt.Sprintf("%v", ciplocalpoolconfigentry.Ciplocalpooladdrtype) + "']" + "[cIpLocalPoolAddressLo='" + fmt.Sprintf("%v", ciplocalpoolconfigentry.Ciplocalpooladdresslo) + "']"
    ciplocalpoolconfigentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciplocalpoolconfigentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciplocalpoolconfigentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciplocalpoolconfigentry.EntityData.Children = make(map[string]types.YChild)
    ciplocalpoolconfigentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ciplocalpoolconfigentry.EntityData.Leafs["cIpLocalPoolName"] = types.YLeaf{"Ciplocalpoolname", ciplocalpoolconfigentry.Ciplocalpoolname}
    ciplocalpoolconfigentry.EntityData.Leafs["cIpLocalPoolAddrType"] = types.YLeaf{"Ciplocalpooladdrtype", ciplocalpoolconfigentry.Ciplocalpooladdrtype}
    ciplocalpoolconfigentry.EntityData.Leafs["cIpLocalPoolAddressLo"] = types.YLeaf{"Ciplocalpooladdresslo", ciplocalpoolconfigentry.Ciplocalpooladdresslo}
    ciplocalpoolconfigentry.EntityData.Leafs["cIpLocalPoolAddressHi"] = types.YLeaf{"Ciplocalpooladdresshi", ciplocalpoolconfigentry.Ciplocalpooladdresshi}
    ciplocalpoolconfigentry.EntityData.Leafs["cIpLocalPoolFreeAddrs"] = types.YLeaf{"Ciplocalpoolfreeaddrs", ciplocalpoolconfigentry.Ciplocalpoolfreeaddrs}
    ciplocalpoolconfigentry.EntityData.Leafs["cIpLocalPoolInUseAddrs"] = types.YLeaf{"Ciplocalpoolinuseaddrs", ciplocalpoolconfigentry.Ciplocalpoolinuseaddrs}
    ciplocalpoolconfigentry.EntityData.Leafs["cIpLocalPoolGroupContainedIn"] = types.YLeaf{"Ciplocalpoolgroupcontainedin", ciplocalpoolconfigentry.Ciplocalpoolgroupcontainedin}
    ciplocalpoolconfigentry.EntityData.Leafs["cIpLocalPoolRowStatus"] = types.YLeaf{"Ciplocalpoolrowstatus", ciplocalpoolconfigentry.Ciplocalpoolrowstatus}
    ciplocalpoolconfigentry.EntityData.Leafs["cIpLocalPoolPriority"] = types.YLeaf{"Ciplocalpoolpriority", ciplocalpoolconfigentry.Ciplocalpoolpriority}
    return &(ciplocalpoolconfigentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry describes single container/'containee' relationship.  Pool names
    // can only be associated with one group.  Pools carry implicit group
    // identifiers because pool names can only be associated with one group.  An
    // entry in this table describes such an association. The type is slice of
    // CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable_Ciplocalpoolgroupcontainsentry.
    Ciplocalpoolgroupcontainsentry []CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable_Ciplocalpoolgroupcontainsentry
}

func (ciplocalpoolgroupcontainstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable) GetEntityData() *types.CommonEntityData {
    ciplocalpoolgroupcontainstable.EntityData.YFilter = ciplocalpoolgroupcontainstable.YFilter
    ciplocalpoolgroupcontainstable.EntityData.YangName = "cIpLocalPoolGroupContainsTable"
    ciplocalpoolgroupcontainstable.EntityData.BundleName = "cisco_ios_xe"
    ciplocalpoolgroupcontainstable.EntityData.ParentYangName = "CISCO-IP-LOCAL-POOL-MIB"
    ciplocalpoolgroupcontainstable.EntityData.SegmentPath = "cIpLocalPoolGroupContainsTable"
    ciplocalpoolgroupcontainstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciplocalpoolgroupcontainstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciplocalpoolgroupcontainstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciplocalpoolgroupcontainstable.EntityData.Children = make(map[string]types.YChild)
    ciplocalpoolgroupcontainstable.EntityData.Children["cIpLocalPoolGroupContainsEntry"] = types.YChild{"Ciplocalpoolgroupcontainsentry", nil}
    for i := range ciplocalpoolgroupcontainstable.Ciplocalpoolgroupcontainsentry {
        ciplocalpoolgroupcontainstable.EntityData.Children[types.GetSegmentPath(&ciplocalpoolgroupcontainstable.Ciplocalpoolgroupcontainsentry[i])] = types.YChild{"Ciplocalpoolgroupcontainsentry", &ciplocalpoolgroupcontainstable.Ciplocalpoolgroupcontainsentry[i]}
    }
    ciplocalpoolgroupcontainstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ciplocalpoolgroupcontainstable.EntityData)
}

// CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable_Ciplocalpoolgroupcontainsentry
// Each entry describes single container/'containee'
// relationship.
// 
// Pool names can only be associated with one group.  Pools carry
// implicit group identifiers because pool names can only be
// associated with one group.  An entry in this table describes
// such an association.
type CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable_Ciplocalpoolgroupcontainsentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. A unique group name that identifies the IP pool
    // group.  The null string represents the base IP pool group. The type is
    // string with length: 0..48.
    Ciplocalpoolgroupname interface{}

    // This attribute is a key. The value of cIpLocalPoolName for the contained IP
    // local pool. The type is string with length: 1..48.
    Ciplocalpoolchildindex interface{}
}

func (ciplocalpoolgroupcontainsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgroupcontainstable_Ciplocalpoolgroupcontainsentry) GetEntityData() *types.CommonEntityData {
    ciplocalpoolgroupcontainsentry.EntityData.YFilter = ciplocalpoolgroupcontainsentry.YFilter
    ciplocalpoolgroupcontainsentry.EntityData.YangName = "cIpLocalPoolGroupContainsEntry"
    ciplocalpoolgroupcontainsentry.EntityData.BundleName = "cisco_ios_xe"
    ciplocalpoolgroupcontainsentry.EntityData.ParentYangName = "cIpLocalPoolGroupContainsTable"
    ciplocalpoolgroupcontainsentry.EntityData.SegmentPath = "cIpLocalPoolGroupContainsEntry" + "[cIpLocalPoolGroupName='" + fmt.Sprintf("%v", ciplocalpoolgroupcontainsentry.Ciplocalpoolgroupname) + "']" + "[cIpLocalPoolChildIndex='" + fmt.Sprintf("%v", ciplocalpoolgroupcontainsentry.Ciplocalpoolchildindex) + "']"
    ciplocalpoolgroupcontainsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciplocalpoolgroupcontainsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciplocalpoolgroupcontainsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciplocalpoolgroupcontainsentry.EntityData.Children = make(map[string]types.YChild)
    ciplocalpoolgroupcontainsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ciplocalpoolgroupcontainsentry.EntityData.Leafs["cIpLocalPoolGroupName"] = types.YLeaf{"Ciplocalpoolgroupname", ciplocalpoolgroupcontainsentry.Ciplocalpoolgroupname}
    ciplocalpoolgroupcontainsentry.EntityData.Leafs["cIpLocalPoolChildIndex"] = types.YLeaf{"Ciplocalpoolchildindex", ciplocalpoolgroupcontainsentry.Ciplocalpoolchildindex}
    return &(ciplocalpoolgroupcontainsentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry provides information about a particular IP pool group and the
    // number of free and used addresses in an IP pool group. The type is slice of
    // CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable_Ciplocalpoolgroupentry.
    Ciplocalpoolgroupentry []CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable_Ciplocalpoolgroupentry
}

func (ciplocalpoolgrouptable *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable) GetEntityData() *types.CommonEntityData {
    ciplocalpoolgrouptable.EntityData.YFilter = ciplocalpoolgrouptable.YFilter
    ciplocalpoolgrouptable.EntityData.YangName = "cIpLocalPoolGroupTable"
    ciplocalpoolgrouptable.EntityData.BundleName = "cisco_ios_xe"
    ciplocalpoolgrouptable.EntityData.ParentYangName = "CISCO-IP-LOCAL-POOL-MIB"
    ciplocalpoolgrouptable.EntityData.SegmentPath = "cIpLocalPoolGroupTable"
    ciplocalpoolgrouptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciplocalpoolgrouptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciplocalpoolgrouptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciplocalpoolgrouptable.EntityData.Children = make(map[string]types.YChild)
    ciplocalpoolgrouptable.EntityData.Children["cIpLocalPoolGroupEntry"] = types.YChild{"Ciplocalpoolgroupentry", nil}
    for i := range ciplocalpoolgrouptable.Ciplocalpoolgroupentry {
        ciplocalpoolgrouptable.EntityData.Children[types.GetSegmentPath(&ciplocalpoolgrouptable.Ciplocalpoolgroupentry[i])] = types.YChild{"Ciplocalpoolgroupentry", &ciplocalpoolgrouptable.Ciplocalpoolgroupentry[i]}
    }
    ciplocalpoolgrouptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ciplocalpoolgrouptable.EntityData)
}

// CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable_Ciplocalpoolgroupentry
// Each entry provides information about a particular IP pool
// group and the number of free and used addresses in an IP pool
// group.
type CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable_Ciplocalpoolgroupentry struct {
    EntityData types.CommonEntityData
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

func (ciplocalpoolgroupentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolgrouptable_Ciplocalpoolgroupentry) GetEntityData() *types.CommonEntityData {
    ciplocalpoolgroupentry.EntityData.YFilter = ciplocalpoolgroupentry.YFilter
    ciplocalpoolgroupentry.EntityData.YangName = "cIpLocalPoolGroupEntry"
    ciplocalpoolgroupentry.EntityData.BundleName = "cisco_ios_xe"
    ciplocalpoolgroupentry.EntityData.ParentYangName = "cIpLocalPoolGroupTable"
    ciplocalpoolgroupentry.EntityData.SegmentPath = "cIpLocalPoolGroupEntry" + "[cIpLocalPoolGroupName='" + fmt.Sprintf("%v", ciplocalpoolgroupentry.Ciplocalpoolgroupname) + "']"
    ciplocalpoolgroupentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciplocalpoolgroupentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciplocalpoolgroupentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciplocalpoolgroupentry.EntityData.Children = make(map[string]types.YChild)
    ciplocalpoolgroupentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ciplocalpoolgroupentry.EntityData.Leafs["cIpLocalPoolGroupName"] = types.YLeaf{"Ciplocalpoolgroupname", ciplocalpoolgroupentry.Ciplocalpoolgroupname}
    ciplocalpoolgroupentry.EntityData.Leafs["cIpLocalPoolGroupFreeAddrs"] = types.YLeaf{"Ciplocalpoolgroupfreeaddrs", ciplocalpoolgroupentry.Ciplocalpoolgroupfreeaddrs}
    ciplocalpoolgroupentry.EntityData.Leafs["cIpLocalPoolGroupInUseAddrs"] = types.YLeaf{"Ciplocalpoolgroupinuseaddrs", ciplocalpoolgroupentry.Ciplocalpoolgroupinuseaddrs}
    return &(ciplocalpoolgroupentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry provides statistical information about a particular IP local
    // pool, and the total number of free and used addresses of all the ranges in
    // an IP local pool. The type is slice of
    // CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable_Ciplocalpoolstatsentry.
    Ciplocalpoolstatsentry []CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable_Ciplocalpoolstatsentry
}

func (ciplocalpoolstatstable *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable) GetEntityData() *types.CommonEntityData {
    ciplocalpoolstatstable.EntityData.YFilter = ciplocalpoolstatstable.YFilter
    ciplocalpoolstatstable.EntityData.YangName = "cIpLocalPoolStatsTable"
    ciplocalpoolstatstable.EntityData.BundleName = "cisco_ios_xe"
    ciplocalpoolstatstable.EntityData.ParentYangName = "CISCO-IP-LOCAL-POOL-MIB"
    ciplocalpoolstatstable.EntityData.SegmentPath = "cIpLocalPoolStatsTable"
    ciplocalpoolstatstable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciplocalpoolstatstable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciplocalpoolstatstable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciplocalpoolstatstable.EntityData.Children = make(map[string]types.YChild)
    ciplocalpoolstatstable.EntityData.Children["cIpLocalPoolStatsEntry"] = types.YChild{"Ciplocalpoolstatsentry", nil}
    for i := range ciplocalpoolstatstable.Ciplocalpoolstatsentry {
        ciplocalpoolstatstable.EntityData.Children[types.GetSegmentPath(&ciplocalpoolstatstable.Ciplocalpoolstatsentry[i])] = types.YChild{"Ciplocalpoolstatsentry", &ciplocalpoolstatstable.Ciplocalpoolstatsentry[i]}
    }
    ciplocalpoolstatstable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ciplocalpoolstatstable.EntityData)
}

// CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable_Ciplocalpoolstatsentry
// Each entry provides statistical information about a particular
// IP local pool, and the total number of free and used addresses
// of all the ranges in an IP local pool.
type CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable_Ciplocalpoolstatsentry struct {
    EntityData types.CommonEntityData
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

func (ciplocalpoolstatsentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolstatstable_Ciplocalpoolstatsentry) GetEntityData() *types.CommonEntityData {
    ciplocalpoolstatsentry.EntityData.YFilter = ciplocalpoolstatsentry.YFilter
    ciplocalpoolstatsentry.EntityData.YangName = "cIpLocalPoolStatsEntry"
    ciplocalpoolstatsentry.EntityData.BundleName = "cisco_ios_xe"
    ciplocalpoolstatsentry.EntityData.ParentYangName = "cIpLocalPoolStatsTable"
    ciplocalpoolstatsentry.EntityData.SegmentPath = "cIpLocalPoolStatsEntry" + "[cIpLocalPoolName='" + fmt.Sprintf("%v", ciplocalpoolstatsentry.Ciplocalpoolname) + "']"
    ciplocalpoolstatsentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciplocalpoolstatsentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciplocalpoolstatsentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciplocalpoolstatsentry.EntityData.Children = make(map[string]types.YChild)
    ciplocalpoolstatsentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ciplocalpoolstatsentry.EntityData.Leafs["cIpLocalPoolName"] = types.YLeaf{"Ciplocalpoolname", ciplocalpoolstatsentry.Ciplocalpoolname}
    ciplocalpoolstatsentry.EntityData.Leafs["cIpLocalPoolStatFreeAddrs"] = types.YLeaf{"Ciplocalpoolstatfreeaddrs", ciplocalpoolstatsentry.Ciplocalpoolstatfreeaddrs}
    ciplocalpoolstatsentry.EntityData.Leafs["cIpLocalPoolStatInUseAddrs"] = types.YLeaf{"Ciplocalpoolstatinuseaddrs", ciplocalpoolstatsentry.Ciplocalpoolstatinuseaddrs}
    ciplocalpoolstatsentry.EntityData.Leafs["cIpLocalPoolStatHiWaterUsedAddrs"] = types.YLeaf{"Ciplocalpoolstathiwaterusedaddrs", ciplocalpoolstatsentry.Ciplocalpoolstathiwaterusedaddrs}
    ciplocalpoolstatsentry.EntityData.Leafs["cIpLocalPoolStatInUseAddrThldLo"] = types.YLeaf{"Ciplocalpoolstatinuseaddrthldlo", ciplocalpoolstatsentry.Ciplocalpoolstatinuseaddrthldlo}
    ciplocalpoolstatsentry.EntityData.Leafs["cIpLocalPoolStatInUseAddrThldHi"] = types.YLeaf{"Ciplocalpoolstatinuseaddrthldhi", ciplocalpoolstatsentry.Ciplocalpoolstatinuseaddrthldhi}
    ciplocalpoolstatsentry.EntityData.Leafs["cIpLocalPoolPercentAddrThldLo"] = types.YLeaf{"Ciplocalpoolpercentaddrthldlo", ciplocalpoolstatsentry.Ciplocalpoolpercentaddrthldlo}
    ciplocalpoolstatsentry.EntityData.Leafs["cIpLocalPoolPercentAddrThldHi"] = types.YLeaf{"Ciplocalpoolpercentaddrthldhi", ciplocalpoolstatsentry.Ciplocalpoolpercentaddrthldhi}
    return &(ciplocalpoolstatsentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry refers to conceptual row that associates an IP addresses with
    // the interface where the request was received, and the user that requested
    // the address. The type is slice of
    // CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable_Ciplocalpoolallocentry.
    Ciplocalpoolallocentry []CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable_Ciplocalpoolallocentry
}

func (ciplocalpoolalloctable *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable) GetEntityData() *types.CommonEntityData {
    ciplocalpoolalloctable.EntityData.YFilter = ciplocalpoolalloctable.YFilter
    ciplocalpoolalloctable.EntityData.YangName = "cIpLocalPoolAllocTable"
    ciplocalpoolalloctable.EntityData.BundleName = "cisco_ios_xe"
    ciplocalpoolalloctable.EntityData.ParentYangName = "CISCO-IP-LOCAL-POOL-MIB"
    ciplocalpoolalloctable.EntityData.SegmentPath = "cIpLocalPoolAllocTable"
    ciplocalpoolalloctable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciplocalpoolalloctable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciplocalpoolalloctable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciplocalpoolalloctable.EntityData.Children = make(map[string]types.YChild)
    ciplocalpoolalloctable.EntityData.Children["cIpLocalPoolAllocEntry"] = types.YChild{"Ciplocalpoolallocentry", nil}
    for i := range ciplocalpoolalloctable.Ciplocalpoolallocentry {
        ciplocalpoolalloctable.EntityData.Children[types.GetSegmentPath(&ciplocalpoolalloctable.Ciplocalpoolallocentry[i])] = types.YChild{"Ciplocalpoolallocentry", &ciplocalpoolalloctable.Ciplocalpoolallocentry[i]}
    }
    ciplocalpoolalloctable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ciplocalpoolalloctable.EntityData)
}

// CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable_Ciplocalpoolallocentry
// Each entry refers to conceptual row that associates an IP
// addresses with the interface where the request was received, and
// the user that requested the address.
type CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable_Ciplocalpoolallocentry struct {
    EntityData types.CommonEntityData
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

func (ciplocalpoolallocentry *CISCOIPLOCALPOOLMIB_Ciplocalpoolalloctable_Ciplocalpoolallocentry) GetEntityData() *types.CommonEntityData {
    ciplocalpoolallocentry.EntityData.YFilter = ciplocalpoolallocentry.YFilter
    ciplocalpoolallocentry.EntityData.YangName = "cIpLocalPoolAllocEntry"
    ciplocalpoolallocentry.EntityData.BundleName = "cisco_ios_xe"
    ciplocalpoolallocentry.EntityData.ParentYangName = "cIpLocalPoolAllocTable"
    ciplocalpoolallocentry.EntityData.SegmentPath = "cIpLocalPoolAllocEntry" + "[cIpLocalPoolName='" + fmt.Sprintf("%v", ciplocalpoolallocentry.Ciplocalpoolname) + "']" + "[cIpLocalPoolAllocAddrType='" + fmt.Sprintf("%v", ciplocalpoolallocentry.Ciplocalpoolallocaddrtype) + "']" + "[cIpLocalPoolAllocAddr='" + fmt.Sprintf("%v", ciplocalpoolallocentry.Ciplocalpoolallocaddr) + "']"
    ciplocalpoolallocentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciplocalpoolallocentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciplocalpoolallocentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciplocalpoolallocentry.EntityData.Children = make(map[string]types.YChild)
    ciplocalpoolallocentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ciplocalpoolallocentry.EntityData.Leafs["cIpLocalPoolName"] = types.YLeaf{"Ciplocalpoolname", ciplocalpoolallocentry.Ciplocalpoolname}
    ciplocalpoolallocentry.EntityData.Leafs["cIpLocalPoolAllocAddrType"] = types.YLeaf{"Ciplocalpoolallocaddrtype", ciplocalpoolallocentry.Ciplocalpoolallocaddrtype}
    ciplocalpoolallocentry.EntityData.Leafs["cIpLocalPoolAllocAddr"] = types.YLeaf{"Ciplocalpoolallocaddr", ciplocalpoolallocentry.Ciplocalpoolallocaddr}
    ciplocalpoolallocentry.EntityData.Leafs["cIpLocalPoolAllocIfIndex"] = types.YLeaf{"Ciplocalpoolallocifindex", ciplocalpoolallocentry.Ciplocalpoolallocifindex}
    ciplocalpoolallocentry.EntityData.Leafs["cIpLocalPoolAllocUser"] = types.YLeaf{"Ciplocalpoolallocuser", ciplocalpoolallocentry.Ciplocalpoolallocuser}
    return &(ciplocalpoolallocentry.EntityData)
}

