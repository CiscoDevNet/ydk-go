// New MIB module for monitoring the memory pools
// of all physical entities on a managed system.
package cisco_enhanced_mempool_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_enhanced_mempool_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-ENHANCED-MEMPOOL-MIB CISCO-ENHANCED-MEMPOOL-MIB}", reflect.TypeOf(CISCOENHANCEDMEMPOOLMIB{}))
    ydk.RegisterEntity("CISCO-ENHANCED-MEMPOOL-MIB:CISCO-ENHANCED-MEMPOOL-MIB", reflect.TypeOf(CISCOENHANCEDMEMPOOLMIB{}))
}

// CempMemPoolTypes represents           processes in ion.
type CempMemPoolTypes string

const (
    CempMemPoolTypes_other CempMemPoolTypes = "other"

    CempMemPoolTypes_processorMemory CempMemPoolTypes = "processorMemory"

    CempMemPoolTypes_ioMemory CempMemPoolTypes = "ioMemory"

    CempMemPoolTypes_pciMemory CempMemPoolTypes = "pciMemory"

    CempMemPoolTypes_fastMemory CempMemPoolTypes = "fastMemory"

    CempMemPoolTypes_multibusMemory CempMemPoolTypes = "multibusMemory"

    CempMemPoolTypes_interruptStackMemory CempMemPoolTypes = "interruptStackMemory"

    CempMemPoolTypes_processStackMemory CempMemPoolTypes = "processStackMemory"

    CempMemPoolTypes_localExceptionMemory CempMemPoolTypes = "localExceptionMemory"

    CempMemPoolTypes_virtualMemory CempMemPoolTypes = "virtualMemory"

    CempMemPoolTypes_reservedMemory CempMemPoolTypes = "reservedMemory"

    CempMemPoolTypes_imageMemory CempMemPoolTypes = "imageMemory"

    CempMemPoolTypes_asicMemory CempMemPoolTypes = "asicMemory"

    CempMemPoolTypes_posixMemory CempMemPoolTypes = "posixMemory"
)

// CISCOENHANCEDMEMPOOLMIB
type CISCOENHANCEDMEMPOOLMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CempNotificationConfig CISCOENHANCEDMEMPOOLMIB_CempNotificationConfig

    // A table of memory pool monitoring entries for all physical entities on a
    // managed system.
    CempMemPoolTable CISCOENHANCEDMEMPOOLMIB_CempMemPoolTable

    // Entries in this table define entities (buffer pools in this case) which are
    // contained in an entity  (memory pool) defined by an entry from
    // cempMemPoolTable. -- Basic Pool Architecture -- 1)Pools are classified as
    // being either Static or    Dynamic. Static pools make no attempt to increase
    // the number of buffers contained within them if the    number of free
    // buffers (cempMemBufferFree) are less   than the number of minimum buffers
    // (cempMemBufferMin).   With Dynamic pools, the pool attempts to meet the   
    // demands of its users. 2)Buffers in a pool are classified as being either   
    // Permanent or Temporary. Permanent buffers, as their   name suggests, are
    // always in the pool and are never   destroyed unless the number of permanent
    // buffers    (cempMemBufferPermanent) is changed. Temporary   buffers are
    // transient buffers that are created in   dynamic pools whenever the free
    // count    (cempMemBufferFree) of buffers in the pool drops    below the
    // minimum (cempMemBufferMin). 3)Buffers pools are classified as either Public
    // or    Private. Public pools are available for all users    to allocate
    // buffers from. Private pools are   primarily used by interface drivers.
    CempMemBufferPoolTable CISCOENHANCEDMEMPOOLMIB_CempMemBufferPoolTable

    // A table that lists the cache buffer pools configured on a managed system. 
    // 1)To provide a noticeable performance boost,    Cache Pool can be used. A
    // Cache Pool is effectively   a lookaside list of free buffers that can be   
    // accessed quickly. Cache Pool is tied to Buffer Pool.  2)Cache pools can
    // optionally have a threshold value   on the number of cache buffers used in
    // a pool. This   can provide flow control management by having a   
    // implementation specific approach such as invoking a   vector when pool
    // cache rises above the optional    threshold set for it on creation.
    CempMemBufferCachePoolTable CISCOENHANCEDMEMPOOLMIB_CempMemBufferCachePoolTable
}

func (cISCOENHANCEDMEMPOOLMIB *CISCOENHANCEDMEMPOOLMIB) GetEntityData() *types.CommonEntityData {
    cISCOENHANCEDMEMPOOLMIB.EntityData.YFilter = cISCOENHANCEDMEMPOOLMIB.YFilter
    cISCOENHANCEDMEMPOOLMIB.EntityData.YangName = "CISCO-ENHANCED-MEMPOOL-MIB"
    cISCOENHANCEDMEMPOOLMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOENHANCEDMEMPOOLMIB.EntityData.ParentYangName = "CISCO-ENHANCED-MEMPOOL-MIB"
    cISCOENHANCEDMEMPOOLMIB.EntityData.SegmentPath = "CISCO-ENHANCED-MEMPOOL-MIB:CISCO-ENHANCED-MEMPOOL-MIB"
    cISCOENHANCEDMEMPOOLMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOENHANCEDMEMPOOLMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOENHANCEDMEMPOOLMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOENHANCEDMEMPOOLMIB.EntityData.Children = types.NewOrderedMap()
    cISCOENHANCEDMEMPOOLMIB.EntityData.Children.Append("cempNotificationConfig", types.YChild{"CempNotificationConfig", &cISCOENHANCEDMEMPOOLMIB.CempNotificationConfig})
    cISCOENHANCEDMEMPOOLMIB.EntityData.Children.Append("cempMemPoolTable", types.YChild{"CempMemPoolTable", &cISCOENHANCEDMEMPOOLMIB.CempMemPoolTable})
    cISCOENHANCEDMEMPOOLMIB.EntityData.Children.Append("cempMemBufferPoolTable", types.YChild{"CempMemBufferPoolTable", &cISCOENHANCEDMEMPOOLMIB.CempMemBufferPoolTable})
    cISCOENHANCEDMEMPOOLMIB.EntityData.Children.Append("cempMemBufferCachePoolTable", types.YChild{"CempMemBufferCachePoolTable", &cISCOENHANCEDMEMPOOLMIB.CempMemBufferCachePoolTable})
    cISCOENHANCEDMEMPOOLMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOENHANCEDMEMPOOLMIB.EntityData.YListKeys = []string {}

    return &(cISCOENHANCEDMEMPOOLMIB.EntityData)
}

// CISCOENHANCEDMEMPOOLMIB_CempNotificationConfig
type CISCOENHANCEDMEMPOOLMIB_CempNotificationConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This variable controls generation of the cempMemBufferNotify.  When this
    // variable is 'true', generation of cempMemBufferNotify is enabled.  When
    // this variable is 'false', generation of cempMemBufferNotify is disabled.
    // The type is bool.
    CempMemBufferNotifyEnabled interface{}
}

func (cempNotificationConfig *CISCOENHANCEDMEMPOOLMIB_CempNotificationConfig) GetEntityData() *types.CommonEntityData {
    cempNotificationConfig.EntityData.YFilter = cempNotificationConfig.YFilter
    cempNotificationConfig.EntityData.YangName = "cempNotificationConfig"
    cempNotificationConfig.EntityData.BundleName = "cisco_ios_xe"
    cempNotificationConfig.EntityData.ParentYangName = "CISCO-ENHANCED-MEMPOOL-MIB"
    cempNotificationConfig.EntityData.SegmentPath = "cempNotificationConfig"
    cempNotificationConfig.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cempNotificationConfig.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cempNotificationConfig.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cempNotificationConfig.EntityData.Children = types.NewOrderedMap()
    cempNotificationConfig.EntityData.Leafs = types.NewOrderedMap()
    cempNotificationConfig.EntityData.Leafs.Append("cempMemBufferNotifyEnabled", types.YLeaf{"CempMemBufferNotifyEnabled", cempNotificationConfig.CempMemBufferNotifyEnabled})

    cempNotificationConfig.EntityData.YListKeys = []string {}

    return &(cempNotificationConfig.EntityData)
}

// CISCOENHANCEDMEMPOOLMIB_CempMemPoolTable
// A table of memory pool monitoring entries for all
// physical entities on a managed system.
type CISCOENHANCEDMEMPOOLMIB_CempMemPoolTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the memory pool monitoring table. The type is slice of
    // CISCOENHANCEDMEMPOOLMIB_CempMemPoolTable_CempMemPoolEntry.
    CempMemPoolEntry []*CISCOENHANCEDMEMPOOLMIB_CempMemPoolTable_CempMemPoolEntry
}

func (cempMemPoolTable *CISCOENHANCEDMEMPOOLMIB_CempMemPoolTable) GetEntityData() *types.CommonEntityData {
    cempMemPoolTable.EntityData.YFilter = cempMemPoolTable.YFilter
    cempMemPoolTable.EntityData.YangName = "cempMemPoolTable"
    cempMemPoolTable.EntityData.BundleName = "cisco_ios_xe"
    cempMemPoolTable.EntityData.ParentYangName = "CISCO-ENHANCED-MEMPOOL-MIB"
    cempMemPoolTable.EntityData.SegmentPath = "cempMemPoolTable"
    cempMemPoolTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cempMemPoolTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cempMemPoolTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cempMemPoolTable.EntityData.Children = types.NewOrderedMap()
    cempMemPoolTable.EntityData.Children.Append("cempMemPoolEntry", types.YChild{"CempMemPoolEntry", nil})
    for i := range cempMemPoolTable.CempMemPoolEntry {
        cempMemPoolTable.EntityData.Children.Append(types.GetSegmentPath(cempMemPoolTable.CempMemPoolEntry[i]), types.YChild{"CempMemPoolEntry", cempMemPoolTable.CempMemPoolEntry[i]})
    }
    cempMemPoolTable.EntityData.Leafs = types.NewOrderedMap()

    cempMemPoolTable.EntityData.YListKeys = []string {}

    return &(cempMemPoolTable.EntityData)
}

// CISCOENHANCEDMEMPOOLMIB_CempMemPoolTable_CempMemPoolEntry
// An entry in the memory pool monitoring table.
type CISCOENHANCEDMEMPOOLMIB_CempMemPoolTable_CempMemPoolEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. Within each physical entity, the unique value
    // greater than zero, used to represent each memory pool.   It is recommended
    // that values are assigned contiguously starting from 1. The type is
    // interface{} with range: 1..2147483647.
    CempMemPoolIndex interface{}

    // The type of memory pool for which this entry contains information. The type
    // is CempMemPoolTypes.
    CempMemPoolType interface{}

    // A textual name assigned to the memory pool. This object is suitable for
    // output to a human operator, and may also be used to distinguish among the
    // various pool types. The type is string.
    CempMemPoolName interface{}

    // An indication of the platform-specific memory pool type. The associated
    // instance of cempMemPoolType is used to indicate the general type of memory
    // pool.  If no platform specific memory hardware type identifier exists for
    // this physical entity, or the value is unknown by this agent, then the value
    // { 0 0 } is returned. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    CempMemPoolPlatformMemory interface{}

    // Indicates whether or not this memory pool has an alternate pool configured.
    // Alternate pools are used for fallback when the current pool runs out of
    // memory.  If an instance of this object has a value of zero, then this pool
    // does not have an alternate.  Otherwise the value of this object is the same
    // as the value of cempMemPoolType of the alternate pool. The type is
    // interface{} with range: 0..2147483647.
    CempMemPoolAlternate interface{}

    // Indicates whether or not cempMemPoolUsed, cempMemPoolFree,
    // cempMemPoolLargestFree and  cempMemPoolLowestFree in this entry contain
    // accurate  data. If an instance of this object has the value  false (which
    // in and of itself indicates an internal  error condition), the values of
    // these objects in the conceptual row may contain inaccurate  information
    // (specifically, the reported values may be  less than the actual values).
    // The type is bool.
    CempMemPoolValid interface{}

    // Indicates the number of bytes from the memory pool that are currently in
    // use by applications on the physical entity. The type is interface{} with
    // range: 0..4294967295. Units are bytes.
    CempMemPoolUsed interface{}

    // Indicates the number of bytes from the memory pool that are currently
    // unused on the physical entity.  Note that the sum of cempMemPoolUsed and
    // cempMemPoolFree  is the total amount of memory in the pool. The type is
    // interface{} with range: 0..4294967295. Units are bytes.
    CempMemPoolFree interface{}

    // Indicates the largest number of contiguous bytes from the memory pool that
    // are currently unused on the physical entity. The type is interface{} with
    // range: 0..4294967295. Units are bytes.
    CempMemPoolLargestFree interface{}

    // The lowest amount of available memory in the memory pool recorded at any
    // time during the operation of the system. The type is interface{} with
    // range: 0..4294967295. Units are bytes.
    CempMemPoolLowestFree interface{}

    // Indicates the lowest number of bytes from the memory pool that have been
    // used by applications on the physical entity since sysUpTime.Similarly,the
    // Used High Watermark indicates the largest number of bytes from the memory
    // pool that have been used by applications on the physical entity since
    // sysUpTime.This can be derived as follows: Used High Watermark =
    // cempMemPoolUsed + cempMemPoolFree  - cempMemPoolLowestFree. The type is
    // interface{} with range: 0..4294967295.
    CempMemPoolUsedLowWaterMark interface{}

    // Indicates the number of successful allocations from the memory pool. The
    // type is interface{} with range: 0..4294967295.
    CempMemPoolAllocHit interface{}

    // Indicates the number of unsuccessful allocations from the memory pool. The
    // type is interface{} with range: 0..4294967295.
    CempMemPoolAllocMiss interface{}

    // Indicates the number of successful frees/ deallocations from the memory
    // pool. The type is interface{} with range: 0..4294967295.
    CempMemPoolFreeHit interface{}

    // Indicates the number of unsuccessful attempts to free/deallocate memory
    // from the memory pool. For example, this could be due to ownership errors 
    // where the application that did not assign the  memory is trying to free it.
    // The type is interface{} with range: 0..4294967295.
    CempMemPoolFreeMiss interface{}

    // Indicates the number of bytes from the memory pool that are currently
    // shared on the physical entity. The type is interface{} with range:
    // 0..4294967295. Units are bytes.
    CempMemPoolShared interface{}

    // This object represents the upper 32-bits of cempMemPoolUsed. This object
    // needs to be supported only if the used bytes in the memory pool exceeds
    // 32-bits, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are bytes.
    CempMemPoolUsedOvrflw interface{}

    // Indicates the number of bytes from the memory pool that are currently in
    // use by applications on the physical entity. This object is a 64-bit version
    // of cempMemPoolUsed. The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    CempMemPoolHCUsed interface{}

    // This object represents the upper 32-bits of cempMemPoolFree. This object
    // needs to be supported only if the unused bytes in the memory pool exceeds
    // 32-bits, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are bytes.
    CempMemPoolFreeOvrflw interface{}

    // Indicates the number of bytes from the memory pool that are currently
    // unused on the physical entity. This object is a 64-bit version of
    // cempMemPoolFree. The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    CempMemPoolHCFree interface{}

    // This object represents the upper 32-bits of cempMemPoolLargestFree. This
    // object needs to  be supported only if the value of  cempMemPoolLargestFree
    // exceeds 32-bits, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are bytes.
    CempMemPoolLargestFreeOvrflw interface{}

    // Indicates the largest number of contiguous bytes from the memory pool that
    // are currently unused on the physical entity. This object is a 64-bit
    // version of cempMemPoolLargestFree. The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    CempMemPoolHCLargestFree interface{}

    // This object represents the upper 32-bits of cempMemPoolLowestFree. This
    // object needs to be supported only if the value of cempMemPoolLowestFree
    // exceeds 32-bits, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are bytes.
    CempMemPoolLowestFreeOvrflw interface{}

    // The lowest amount of available memory in the memory pool recorded at any
    // time during the operation of the system. This object is a 64-bit version of
    // cempMemPoolLowestFree. The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    CempMemPoolHCLowestFree interface{}

    // This object represents the upper 32-bits of cempMemPoolUsedLowWaterMark.
    // This object needs to be supported only if the value of
    // cempMemPoolUsedLowWaterMark exceeds 32-bits, otherwise this object value
    // would be set to 0. The type is interface{} with range: 0..4294967295. Units
    // are bytes.
    CempMemPoolUsedLowWaterMarkOvrflw interface{}

    // Indicates the lowest number of bytes from the memory pool that have been
    // used by applications on the physical entity since sysUpTime. This object is
    // a 64-bit version of cempMemPoolUsedLowWaterMark. The type is interface{}
    // with range: 0..18446744073709551615. Units are bytes.
    CempMemPoolHCUsedLowWaterMark interface{}

    // This object represents the upper 32-bits of cempMemPoolShared. This object
    // needs to be supported only if the value of cempMemPoolShared exceeds
    // 32-bits, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are bytes.
    CempMemPoolSharedOvrflw interface{}

    // Indicates the number of bytes from the memory pool that are currently
    // shared on the physical entity. This object is a 64-bit version of
    // cempMemPoolShared. The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    CempMemPoolHCShared interface{}
}

func (cempMemPoolEntry *CISCOENHANCEDMEMPOOLMIB_CempMemPoolTable_CempMemPoolEntry) GetEntityData() *types.CommonEntityData {
    cempMemPoolEntry.EntityData.YFilter = cempMemPoolEntry.YFilter
    cempMemPoolEntry.EntityData.YangName = "cempMemPoolEntry"
    cempMemPoolEntry.EntityData.BundleName = "cisco_ios_xe"
    cempMemPoolEntry.EntityData.ParentYangName = "cempMemPoolTable"
    cempMemPoolEntry.EntityData.SegmentPath = "cempMemPoolEntry" + types.AddKeyToken(cempMemPoolEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(cempMemPoolEntry.CempMemPoolIndex, "cempMemPoolIndex")
    cempMemPoolEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cempMemPoolEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cempMemPoolEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cempMemPoolEntry.EntityData.Children = types.NewOrderedMap()
    cempMemPoolEntry.EntityData.Leafs = types.NewOrderedMap()
    cempMemPoolEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cempMemPoolEntry.EntPhysicalIndex})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolIndex", types.YLeaf{"CempMemPoolIndex", cempMemPoolEntry.CempMemPoolIndex})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolType", types.YLeaf{"CempMemPoolType", cempMemPoolEntry.CempMemPoolType})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolName", types.YLeaf{"CempMemPoolName", cempMemPoolEntry.CempMemPoolName})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolPlatformMemory", types.YLeaf{"CempMemPoolPlatformMemory", cempMemPoolEntry.CempMemPoolPlatformMemory})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolAlternate", types.YLeaf{"CempMemPoolAlternate", cempMemPoolEntry.CempMemPoolAlternate})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolValid", types.YLeaf{"CempMemPoolValid", cempMemPoolEntry.CempMemPoolValid})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolUsed", types.YLeaf{"CempMemPoolUsed", cempMemPoolEntry.CempMemPoolUsed})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolFree", types.YLeaf{"CempMemPoolFree", cempMemPoolEntry.CempMemPoolFree})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolLargestFree", types.YLeaf{"CempMemPoolLargestFree", cempMemPoolEntry.CempMemPoolLargestFree})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolLowestFree", types.YLeaf{"CempMemPoolLowestFree", cempMemPoolEntry.CempMemPoolLowestFree})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolUsedLowWaterMark", types.YLeaf{"CempMemPoolUsedLowWaterMark", cempMemPoolEntry.CempMemPoolUsedLowWaterMark})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolAllocHit", types.YLeaf{"CempMemPoolAllocHit", cempMemPoolEntry.CempMemPoolAllocHit})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolAllocMiss", types.YLeaf{"CempMemPoolAllocMiss", cempMemPoolEntry.CempMemPoolAllocMiss})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolFreeHit", types.YLeaf{"CempMemPoolFreeHit", cempMemPoolEntry.CempMemPoolFreeHit})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolFreeMiss", types.YLeaf{"CempMemPoolFreeMiss", cempMemPoolEntry.CempMemPoolFreeMiss})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolShared", types.YLeaf{"CempMemPoolShared", cempMemPoolEntry.CempMemPoolShared})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolUsedOvrflw", types.YLeaf{"CempMemPoolUsedOvrflw", cempMemPoolEntry.CempMemPoolUsedOvrflw})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolHCUsed", types.YLeaf{"CempMemPoolHCUsed", cempMemPoolEntry.CempMemPoolHCUsed})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolFreeOvrflw", types.YLeaf{"CempMemPoolFreeOvrflw", cempMemPoolEntry.CempMemPoolFreeOvrflw})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolHCFree", types.YLeaf{"CempMemPoolHCFree", cempMemPoolEntry.CempMemPoolHCFree})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolLargestFreeOvrflw", types.YLeaf{"CempMemPoolLargestFreeOvrflw", cempMemPoolEntry.CempMemPoolLargestFreeOvrflw})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolHCLargestFree", types.YLeaf{"CempMemPoolHCLargestFree", cempMemPoolEntry.CempMemPoolHCLargestFree})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolLowestFreeOvrflw", types.YLeaf{"CempMemPoolLowestFreeOvrflw", cempMemPoolEntry.CempMemPoolLowestFreeOvrflw})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolHCLowestFree", types.YLeaf{"CempMemPoolHCLowestFree", cempMemPoolEntry.CempMemPoolHCLowestFree})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolUsedLowWaterMarkOvrflw", types.YLeaf{"CempMemPoolUsedLowWaterMarkOvrflw", cempMemPoolEntry.CempMemPoolUsedLowWaterMarkOvrflw})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolHCUsedLowWaterMark", types.YLeaf{"CempMemPoolHCUsedLowWaterMark", cempMemPoolEntry.CempMemPoolHCUsedLowWaterMark})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolSharedOvrflw", types.YLeaf{"CempMemPoolSharedOvrflw", cempMemPoolEntry.CempMemPoolSharedOvrflw})
    cempMemPoolEntry.EntityData.Leafs.Append("cempMemPoolHCShared", types.YLeaf{"CempMemPoolHCShared", cempMemPoolEntry.CempMemPoolHCShared})

    cempMemPoolEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "CempMemPoolIndex"}

    return &(cempMemPoolEntry.EntityData)
}

// CISCOENHANCEDMEMPOOLMIB_CempMemBufferPoolTable
// Entries in this table define entities (buffer pools
// in this case) which are contained in an entity 
// (memory pool) defined by an entry from
// cempMemPoolTable.
// -- Basic Pool Architecture --
// 1)Pools are classified as being either Static or 
//   Dynamic. Static pools make no attempt to increase 
//   the number of buffers contained within them if the 
//   number of free buffers (cempMemBufferFree) are less
//   than the number of minimum buffers (cempMemBufferMin).
//   With Dynamic pools, the pool attempts to meet the 
//   demands of its users.
// 2)Buffers in a pool are classified as being either 
//   Permanent or Temporary. Permanent buffers, as their
//   name suggests, are always in the pool and are never
//   destroyed unless the number of permanent buffers 
//   (cempMemBufferPermanent) is changed. Temporary
//   buffers are transient buffers that are created in
//   dynamic pools whenever the free count 
//   (cempMemBufferFree) of buffers in the pool drops 
//   below the minimum (cempMemBufferMin).
// 3)Buffers pools are classified as either Public or 
//   Private. Public pools are available for all users 
//   to allocate buffers from. Private pools are
//   primarily used by interface drivers.
type CISCOENHANCEDMEMPOOLMIB_CempMemBufferPoolTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This contains all the memory buffer pool configurations object values. The 
    // entPhysicalIndex identifies the entity on which memory buffer pools are
    // present. The type is slice of
    // CISCOENHANCEDMEMPOOLMIB_CempMemBufferPoolTable_CempMemBufferPoolEntry.
    CempMemBufferPoolEntry []*CISCOENHANCEDMEMPOOLMIB_CempMemBufferPoolTable_CempMemBufferPoolEntry
}

func (cempMemBufferPoolTable *CISCOENHANCEDMEMPOOLMIB_CempMemBufferPoolTable) GetEntityData() *types.CommonEntityData {
    cempMemBufferPoolTable.EntityData.YFilter = cempMemBufferPoolTable.YFilter
    cempMemBufferPoolTable.EntityData.YangName = "cempMemBufferPoolTable"
    cempMemBufferPoolTable.EntityData.BundleName = "cisco_ios_xe"
    cempMemBufferPoolTable.EntityData.ParentYangName = "CISCO-ENHANCED-MEMPOOL-MIB"
    cempMemBufferPoolTable.EntityData.SegmentPath = "cempMemBufferPoolTable"
    cempMemBufferPoolTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cempMemBufferPoolTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cempMemBufferPoolTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cempMemBufferPoolTable.EntityData.Children = types.NewOrderedMap()
    cempMemBufferPoolTable.EntityData.Children.Append("cempMemBufferPoolEntry", types.YChild{"CempMemBufferPoolEntry", nil})
    for i := range cempMemBufferPoolTable.CempMemBufferPoolEntry {
        cempMemBufferPoolTable.EntityData.Children.Append(types.GetSegmentPath(cempMemBufferPoolTable.CempMemBufferPoolEntry[i]), types.YChild{"CempMemBufferPoolEntry", cempMemBufferPoolTable.CempMemBufferPoolEntry[i]})
    }
    cempMemBufferPoolTable.EntityData.Leafs = types.NewOrderedMap()

    cempMemBufferPoolTable.EntityData.YListKeys = []string {}

    return &(cempMemBufferPoolTable.EntityData)
}

// CISCOENHANCEDMEMPOOLMIB_CempMemBufferPoolTable_CempMemBufferPoolEntry
// This contains all the memory buffer pool
// configurations object values. The 
// entPhysicalIndex identifies the entity on which
// memory buffer pools are present.
type CISCOENHANCEDMEMPOOLMIB_CempMemBufferPoolTable_CempMemBufferPoolEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. Within a physical entity, a unique value used to
    // represent each buffer pool. The type is interface{} with range:
    // 0..4294967295.
    CempMemBufferPoolIndex interface{}

    // This index corresponds to the memory pool (with cemMemPoolIndex as index in
    // cempMemPoolTable)  from which buffers are allocated. The type is
    // interface{} with range: 0..2147483647.
    CempMemBufferMemPoolIndex interface{}

    // A textual name assigned to the buffer pool. This object is suitable for
    // output to a human operator, and may also be used to distinguish among the
    // various buffer types. For example: 'Small', 'Big', 'Serial0/1' etc. The
    // type is string.
    CempMemBufferName interface{}

    // Boolean poolDynamic; if TRUE, the number of buffers in the pool is adjusted
    // (adding more packet buffers  or deleting excesses) dynamically by the
    // background  process. If FALSE, the number of buffers in the pool  is never
    // adjusted, even if it falls below the minimum, or to zero. The type is bool.
    CempMemBufferDynamic interface{}

    // Indicates the size of buffer element in number of bytes on the physical
    // entity. The type is interface{} with range: 0..4294967295. Units are bytes.
    CempMemBufferSize interface{}

    // Indicates the minimum number of free buffers allowed in the buffer pool or
    // low-water mark (lwm).  For example of its usage : If cempMemBufferFree <
    // cempMemBufferMin & pool is  dynamic, then signal for growth of particular
    // buffer pool. The type is interface{} with range: 0..4294967295.
    CempMemBufferMin interface{}

    // Indicates the maximum number of free buffers allowed in the buffer pool or
    // high-water mark (hwm). For example of its usage : If cempMemBufferFree >
    // cempMemBufferMax & pool is  dynamic, then signal for trim of particular
    // buffer pool. The type is interface{} with range: 0..4294967295.
    CempMemBufferMax interface{}

    // Indicates the total number of permanent buffers in the pool on the physical
    // entity. The type is interface{} with range: 0..4294967295.
    CempMemBufferPermanent interface{}

    // Indicates the initial number of temporary buffers in the pool on the
    // physical entity. This object  instructs the system to create this many
    // number of temporary extra buffers, just after a system restart.  A change
    // in this object will be effective only after a system restart. The type is
    // interface{} with range: 0..4294967295.
    CempMemBufferTransient interface{}

    // Indicates the total number of buffers (include allocated and free buffers)
    // in the buffer pool on the physical entity. The type is interface{} with
    // range: 0..4294967295.
    CempMemBufferTotal interface{}

    // Indicates the current number of free buffers in the buffer pool on the
    // physical entity. Note that the cempMemBufferFree is less than or equal  to
    // cempMemBufferTotal. The type is interface{} with range: 0..4294967295.
    CempMemBufferFree interface{}

    // Indicates the number of buffers successfully allocated from the buffer
    // pool. The type is interface{} with range: 0..4294967295.
    CempMemBufferHit interface{}

    // Indicates the number of times a buffer has been requested, but no buffers
    // were available in the buffer pool, or when there were fewer than min 
    // buffers(cempMemBufferMin) in the buffer pool. Note : For interface pools, a
    // miss is actually  a fall back to its corresponding public buffer pool. The
    // type is interface{} with range: 0..4294967295.
    CempMemBufferMiss interface{}

    // Indicates the number of successful frees/deallocations from the buffer
    // pool. The type is interface{} with range: 0..4294967295.
    CempMemBufferFreeHit interface{}

    // Indicates the number of unsuccessful attempts to free/deallocate a buffer
    // from the buffer pool.  For example, this could be due to ownership errors
    // where the application that did not assign the  buffer is trying to free it.
    // The type is interface{} with range: 0..4294967295.
    CempMemBufferFreeMiss interface{}

    // This value is the difference of the desired number of permanent buffer &
    // total number of permanent  buffers present in the pool. A positive value of
    // this object tells the number of buffers needed & a  negative value of the
    // object tells the extra number  of buffers in the pool. The type is
    // interface{} with range: -2147483648..2147483647.
    CempMemBufferPermChange interface{}

    // Indicates the peak number of buffers in pool on the physical entity. The
    // type is interface{} with range: 0..4294967295.
    CempMemBufferPeak interface{}

    // Indicates the time of most recent change in the peak number of buffers
    // (cempMemBufferPeak object) in the pool. The type is interface{} with range:
    // 0..4294967295.
    CempMemBufferPeakTime interface{}

    // The number of buffers that have been trimmed from the pool when the number
    // of free buffers  (cempMemBufferFree) exceeded the number of max allowed
    // buffers(cempMemBufferMax). The type is interface{} with range:
    // 0..4294967295.
    CempMemBufferTrim interface{}

    // The number of buffers that have been created in the pool when the number of
    // free buffers(cempMemBufferFree) was less than minimum(cempMemBufferMix).
    // The type is interface{} with range: 0..4294967295.
    CempMemBufferGrow interface{}

    // The number of failures to grant a buffer to a requester due to reasons
    // other than insufficient  memory. For example, in systems where there are 
    // different execution contexts, it may be too expensive to create new buffers
    // when running in certain contexts. In those cases it may be  preferable to
    // fail the request. The type is interface{} with range: 0..4294967295.
    CempMemBufferFailures interface{}

    // The number of times the system tried to create new buffers, but could not
    // due to insufficient free  memory in the system. The type is interface{}
    // with range: 0..4294967295.
    CempMemBufferNoStorage interface{}
}

func (cempMemBufferPoolEntry *CISCOENHANCEDMEMPOOLMIB_CempMemBufferPoolTable_CempMemBufferPoolEntry) GetEntityData() *types.CommonEntityData {
    cempMemBufferPoolEntry.EntityData.YFilter = cempMemBufferPoolEntry.YFilter
    cempMemBufferPoolEntry.EntityData.YangName = "cempMemBufferPoolEntry"
    cempMemBufferPoolEntry.EntityData.BundleName = "cisco_ios_xe"
    cempMemBufferPoolEntry.EntityData.ParentYangName = "cempMemBufferPoolTable"
    cempMemBufferPoolEntry.EntityData.SegmentPath = "cempMemBufferPoolEntry" + types.AddKeyToken(cempMemBufferPoolEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(cempMemBufferPoolEntry.CempMemBufferPoolIndex, "cempMemBufferPoolIndex")
    cempMemBufferPoolEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cempMemBufferPoolEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cempMemBufferPoolEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cempMemBufferPoolEntry.EntityData.Children = types.NewOrderedMap()
    cempMemBufferPoolEntry.EntityData.Leafs = types.NewOrderedMap()
    cempMemBufferPoolEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cempMemBufferPoolEntry.EntPhysicalIndex})
    cempMemBufferPoolEntry.EntityData.Leafs.Append("cempMemBufferPoolIndex", types.YLeaf{"CempMemBufferPoolIndex", cempMemBufferPoolEntry.CempMemBufferPoolIndex})
    cempMemBufferPoolEntry.EntityData.Leafs.Append("cempMemBufferMemPoolIndex", types.YLeaf{"CempMemBufferMemPoolIndex", cempMemBufferPoolEntry.CempMemBufferMemPoolIndex})
    cempMemBufferPoolEntry.EntityData.Leafs.Append("cempMemBufferName", types.YLeaf{"CempMemBufferName", cempMemBufferPoolEntry.CempMemBufferName})
    cempMemBufferPoolEntry.EntityData.Leafs.Append("cempMemBufferDynamic", types.YLeaf{"CempMemBufferDynamic", cempMemBufferPoolEntry.CempMemBufferDynamic})
    cempMemBufferPoolEntry.EntityData.Leafs.Append("cempMemBufferSize", types.YLeaf{"CempMemBufferSize", cempMemBufferPoolEntry.CempMemBufferSize})
    cempMemBufferPoolEntry.EntityData.Leafs.Append("cempMemBufferMin", types.YLeaf{"CempMemBufferMin", cempMemBufferPoolEntry.CempMemBufferMin})
    cempMemBufferPoolEntry.EntityData.Leafs.Append("cempMemBufferMax", types.YLeaf{"CempMemBufferMax", cempMemBufferPoolEntry.CempMemBufferMax})
    cempMemBufferPoolEntry.EntityData.Leafs.Append("cempMemBufferPermanent", types.YLeaf{"CempMemBufferPermanent", cempMemBufferPoolEntry.CempMemBufferPermanent})
    cempMemBufferPoolEntry.EntityData.Leafs.Append("cempMemBufferTransient", types.YLeaf{"CempMemBufferTransient", cempMemBufferPoolEntry.CempMemBufferTransient})
    cempMemBufferPoolEntry.EntityData.Leafs.Append("cempMemBufferTotal", types.YLeaf{"CempMemBufferTotal", cempMemBufferPoolEntry.CempMemBufferTotal})
    cempMemBufferPoolEntry.EntityData.Leafs.Append("cempMemBufferFree", types.YLeaf{"CempMemBufferFree", cempMemBufferPoolEntry.CempMemBufferFree})
    cempMemBufferPoolEntry.EntityData.Leafs.Append("cempMemBufferHit", types.YLeaf{"CempMemBufferHit", cempMemBufferPoolEntry.CempMemBufferHit})
    cempMemBufferPoolEntry.EntityData.Leafs.Append("cempMemBufferMiss", types.YLeaf{"CempMemBufferMiss", cempMemBufferPoolEntry.CempMemBufferMiss})
    cempMemBufferPoolEntry.EntityData.Leafs.Append("cempMemBufferFreeHit", types.YLeaf{"CempMemBufferFreeHit", cempMemBufferPoolEntry.CempMemBufferFreeHit})
    cempMemBufferPoolEntry.EntityData.Leafs.Append("cempMemBufferFreeMiss", types.YLeaf{"CempMemBufferFreeMiss", cempMemBufferPoolEntry.CempMemBufferFreeMiss})
    cempMemBufferPoolEntry.EntityData.Leafs.Append("cempMemBufferPermChange", types.YLeaf{"CempMemBufferPermChange", cempMemBufferPoolEntry.CempMemBufferPermChange})
    cempMemBufferPoolEntry.EntityData.Leafs.Append("cempMemBufferPeak", types.YLeaf{"CempMemBufferPeak", cempMemBufferPoolEntry.CempMemBufferPeak})
    cempMemBufferPoolEntry.EntityData.Leafs.Append("cempMemBufferPeakTime", types.YLeaf{"CempMemBufferPeakTime", cempMemBufferPoolEntry.CempMemBufferPeakTime})
    cempMemBufferPoolEntry.EntityData.Leafs.Append("cempMemBufferTrim", types.YLeaf{"CempMemBufferTrim", cempMemBufferPoolEntry.CempMemBufferTrim})
    cempMemBufferPoolEntry.EntityData.Leafs.Append("cempMemBufferGrow", types.YLeaf{"CempMemBufferGrow", cempMemBufferPoolEntry.CempMemBufferGrow})
    cempMemBufferPoolEntry.EntityData.Leafs.Append("cempMemBufferFailures", types.YLeaf{"CempMemBufferFailures", cempMemBufferPoolEntry.CempMemBufferFailures})
    cempMemBufferPoolEntry.EntityData.Leafs.Append("cempMemBufferNoStorage", types.YLeaf{"CempMemBufferNoStorage", cempMemBufferPoolEntry.CempMemBufferNoStorage})

    cempMemBufferPoolEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "CempMemBufferPoolIndex"}

    return &(cempMemBufferPoolEntry.EntityData)
}

// CISCOENHANCEDMEMPOOLMIB_CempMemBufferCachePoolTable
// A table that lists the cache buffer pools
// configured on a managed system. 
// 1)To provide a noticeable performance boost, 
//   Cache Pool can be used. A Cache Pool is effectively
//   a lookaside list of free buffers that can be 
//   accessed quickly. Cache Pool is tied to Buffer Pool. 
// 2)Cache pools can optionally have a threshold value
//   on the number of cache buffers used in a pool. This
//   can provide flow control management by having a 
//   implementation specific approach such as invoking a
//   vector when pool cache rises above the optional 
//   threshold set for it on creation.
type CISCOENHANCEDMEMPOOLMIB_CempMemBufferCachePoolTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry represents one of the cache buffer pools available in the system
    // and it contains the parameters configured for it. Note :
    // cempMemBufferCachePoolTable has a sparse dependency with
    // cempMemBufferPoolTable (i.e all the entires in cempMemBufferPoolTable need
    // not have an entry in cempMemBufferCachePoolTable. The type is slice of
    // CISCOENHANCEDMEMPOOLMIB_CempMemBufferCachePoolTable_CempMemBufferCachePoolEntry.
    CempMemBufferCachePoolEntry []*CISCOENHANCEDMEMPOOLMIB_CempMemBufferCachePoolTable_CempMemBufferCachePoolEntry
}

func (cempMemBufferCachePoolTable *CISCOENHANCEDMEMPOOLMIB_CempMemBufferCachePoolTable) GetEntityData() *types.CommonEntityData {
    cempMemBufferCachePoolTable.EntityData.YFilter = cempMemBufferCachePoolTable.YFilter
    cempMemBufferCachePoolTable.EntityData.YangName = "cempMemBufferCachePoolTable"
    cempMemBufferCachePoolTable.EntityData.BundleName = "cisco_ios_xe"
    cempMemBufferCachePoolTable.EntityData.ParentYangName = "CISCO-ENHANCED-MEMPOOL-MIB"
    cempMemBufferCachePoolTable.EntityData.SegmentPath = "cempMemBufferCachePoolTable"
    cempMemBufferCachePoolTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cempMemBufferCachePoolTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cempMemBufferCachePoolTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cempMemBufferCachePoolTable.EntityData.Children = types.NewOrderedMap()
    cempMemBufferCachePoolTable.EntityData.Children.Append("cempMemBufferCachePoolEntry", types.YChild{"CempMemBufferCachePoolEntry", nil})
    for i := range cempMemBufferCachePoolTable.CempMemBufferCachePoolEntry {
        cempMemBufferCachePoolTable.EntityData.Children.Append(types.GetSegmentPath(cempMemBufferCachePoolTable.CempMemBufferCachePoolEntry[i]), types.YChild{"CempMemBufferCachePoolEntry", cempMemBufferCachePoolTable.CempMemBufferCachePoolEntry[i]})
    }
    cempMemBufferCachePoolTable.EntityData.Leafs = types.NewOrderedMap()

    cempMemBufferCachePoolTable.EntityData.YListKeys = []string {}

    return &(cempMemBufferCachePoolTable.EntityData)
}

// CISCOENHANCEDMEMPOOLMIB_CempMemBufferCachePoolTable_CempMemBufferCachePoolEntry
// Each entry represents one of the cache buffer pools
// available in the system and it contains the
// parameters configured for it.
// Note : cempMemBufferCachePoolTable has a sparse
// dependency with cempMemBufferPoolTable (i.e all the
// entires in cempMemBufferPoolTable need not have an
// entry in cempMemBufferCachePoolTable.
type CISCOENHANCEDMEMPOOLMIB_CempMemBufferCachePoolTable_CempMemBufferCachePoolEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_enhanced_mempool_mib.CISCOENHANCEDMEMPOOLMIB_CempMemBufferPoolTable_CempMemBufferPoolEntry_CempMemBufferPoolIndex
    CempMemBufferPoolIndex interface{}

    // Indicates the number of buffers in the cache pool on the physical entity.
    // The type is interface{} with range: 0..4294967295.
    CempMemBufferCacheSize interface{}

    // Indicates the maximum number of free buffers allowed in the cache pool. The
    // type is interface{} with range: 0..4294967295.
    CempMemBufferCacheTotal interface{}

    // Indicates the number of cache buffers from the pool that are currently used
    // on the physical entity. Note that the cempMemBufferCacheUsed is less than
    // or  equal to cempMemBufferCacheTotal. The type is interface{} with range:
    // 0..4294967295.
    CempMemBufferCacheUsed interface{}

    // Indicates the number of buffers successfully allocated from the cache pool.
    // The type is interface{} with range: 0..4294967295.
    CempMemBufferCacheHit interface{}

    // Indicates the number of times a buffer has been requested, but no buffers
    // were available in the cache pool. The type is interface{} with range:
    // 0..4294967295.
    CempMemBufferCacheMiss interface{}

    // Indicates the threshold limit for number of cache buffers
    // used(cempMemBufferCacheUsed). The type is interface{} with range:
    // 0..4294967295.
    CempMemBufferCacheThreshold interface{}

    // Indicates how many times the number of cache buffers
    // used(cempMemBufferCacheUsed) has crossed the threshold
    // value(cempMemBufferCacheThreshold). The type is interface{} with range:
    // 0..4294967295.
    CempMemBufferCacheThresholdCount interface{}
}

func (cempMemBufferCachePoolEntry *CISCOENHANCEDMEMPOOLMIB_CempMemBufferCachePoolTable_CempMemBufferCachePoolEntry) GetEntityData() *types.CommonEntityData {
    cempMemBufferCachePoolEntry.EntityData.YFilter = cempMemBufferCachePoolEntry.YFilter
    cempMemBufferCachePoolEntry.EntityData.YangName = "cempMemBufferCachePoolEntry"
    cempMemBufferCachePoolEntry.EntityData.BundleName = "cisco_ios_xe"
    cempMemBufferCachePoolEntry.EntityData.ParentYangName = "cempMemBufferCachePoolTable"
    cempMemBufferCachePoolEntry.EntityData.SegmentPath = "cempMemBufferCachePoolEntry" + types.AddKeyToken(cempMemBufferCachePoolEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(cempMemBufferCachePoolEntry.CempMemBufferPoolIndex, "cempMemBufferPoolIndex")
    cempMemBufferCachePoolEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cempMemBufferCachePoolEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cempMemBufferCachePoolEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cempMemBufferCachePoolEntry.EntityData.Children = types.NewOrderedMap()
    cempMemBufferCachePoolEntry.EntityData.Leafs = types.NewOrderedMap()
    cempMemBufferCachePoolEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cempMemBufferCachePoolEntry.EntPhysicalIndex})
    cempMemBufferCachePoolEntry.EntityData.Leafs.Append("cempMemBufferPoolIndex", types.YLeaf{"CempMemBufferPoolIndex", cempMemBufferCachePoolEntry.CempMemBufferPoolIndex})
    cempMemBufferCachePoolEntry.EntityData.Leafs.Append("cempMemBufferCacheSize", types.YLeaf{"CempMemBufferCacheSize", cempMemBufferCachePoolEntry.CempMemBufferCacheSize})
    cempMemBufferCachePoolEntry.EntityData.Leafs.Append("cempMemBufferCacheTotal", types.YLeaf{"CempMemBufferCacheTotal", cempMemBufferCachePoolEntry.CempMemBufferCacheTotal})
    cempMemBufferCachePoolEntry.EntityData.Leafs.Append("cempMemBufferCacheUsed", types.YLeaf{"CempMemBufferCacheUsed", cempMemBufferCachePoolEntry.CempMemBufferCacheUsed})
    cempMemBufferCachePoolEntry.EntityData.Leafs.Append("cempMemBufferCacheHit", types.YLeaf{"CempMemBufferCacheHit", cempMemBufferCachePoolEntry.CempMemBufferCacheHit})
    cempMemBufferCachePoolEntry.EntityData.Leafs.Append("cempMemBufferCacheMiss", types.YLeaf{"CempMemBufferCacheMiss", cempMemBufferCachePoolEntry.CempMemBufferCacheMiss})
    cempMemBufferCachePoolEntry.EntityData.Leafs.Append("cempMemBufferCacheThreshold", types.YLeaf{"CempMemBufferCacheThreshold", cempMemBufferCachePoolEntry.CempMemBufferCacheThreshold})
    cempMemBufferCachePoolEntry.EntityData.Leafs.Append("cempMemBufferCacheThresholdCount", types.YLeaf{"CempMemBufferCacheThresholdCount", cempMemBufferCachePoolEntry.CempMemBufferCacheThresholdCount})

    cempMemBufferCachePoolEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "CempMemBufferPoolIndex"}

    return &(cempMemBufferCachePoolEntry.EntityData)
}

