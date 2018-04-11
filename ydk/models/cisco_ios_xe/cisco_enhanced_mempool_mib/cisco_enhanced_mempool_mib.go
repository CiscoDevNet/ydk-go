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

    
    Cempnotificationconfig CISCOENHANCEDMEMPOOLMIB_Cempnotificationconfig

    // A table of memory pool monitoring entries for all physical entities on a
    // managed system.
    Cempmempooltable CISCOENHANCEDMEMPOOLMIB_Cempmempooltable

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
    Cempmembufferpooltable CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable

    // A table that lists the cache buffer pools configured on a managed system. 
    // 1)To provide a noticeable performance boost,    Cache Pool can be used. A
    // Cache Pool is effectively   a lookaside list of free buffers that can be   
    // accessed quickly. Cache Pool is tied to Buffer Pool.  2)Cache pools can
    // optionally have a threshold value   on the number of cache buffers used in
    // a pool. This   can provide flow control management by having a   
    // implementation specific approach such as invoking a   vector when pool
    // cache rises above the optional    threshold set for it on creation.
    Cempmembuffercachepooltable CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable
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

    cISCOENHANCEDMEMPOOLMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOENHANCEDMEMPOOLMIB.EntityData.Children["cempNotificationConfig"] = types.YChild{"Cempnotificationconfig", &cISCOENHANCEDMEMPOOLMIB.Cempnotificationconfig}
    cISCOENHANCEDMEMPOOLMIB.EntityData.Children["cempMemPoolTable"] = types.YChild{"Cempmempooltable", &cISCOENHANCEDMEMPOOLMIB.Cempmempooltable}
    cISCOENHANCEDMEMPOOLMIB.EntityData.Children["cempMemBufferPoolTable"] = types.YChild{"Cempmembufferpooltable", &cISCOENHANCEDMEMPOOLMIB.Cempmembufferpooltable}
    cISCOENHANCEDMEMPOOLMIB.EntityData.Children["cempMemBufferCachePoolTable"] = types.YChild{"Cempmembuffercachepooltable", &cISCOENHANCEDMEMPOOLMIB.Cempmembuffercachepooltable}
    cISCOENHANCEDMEMPOOLMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOENHANCEDMEMPOOLMIB.EntityData)
}

// CISCOENHANCEDMEMPOOLMIB_Cempnotificationconfig
type CISCOENHANCEDMEMPOOLMIB_Cempnotificationconfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This variable controls generation of the cempMemBufferNotify.  When this
    // variable is 'true', generation of cempMemBufferNotify is enabled.  When
    // this variable is 'false', generation of cempMemBufferNotify is disabled.
    // The type is bool.
    Cempmembuffernotifyenabled interface{}
}

func (cempnotificationconfig *CISCOENHANCEDMEMPOOLMIB_Cempnotificationconfig) GetEntityData() *types.CommonEntityData {
    cempnotificationconfig.EntityData.YFilter = cempnotificationconfig.YFilter
    cempnotificationconfig.EntityData.YangName = "cempNotificationConfig"
    cempnotificationconfig.EntityData.BundleName = "cisco_ios_xe"
    cempnotificationconfig.EntityData.ParentYangName = "CISCO-ENHANCED-MEMPOOL-MIB"
    cempnotificationconfig.EntityData.SegmentPath = "cempNotificationConfig"
    cempnotificationconfig.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cempnotificationconfig.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cempnotificationconfig.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cempnotificationconfig.EntityData.Children = make(map[string]types.YChild)
    cempnotificationconfig.EntityData.Leafs = make(map[string]types.YLeaf)
    cempnotificationconfig.EntityData.Leafs["cempMemBufferNotifyEnabled"] = types.YLeaf{"Cempmembuffernotifyenabled", cempnotificationconfig.Cempmembuffernotifyenabled}
    return &(cempnotificationconfig.EntityData)
}

// CISCOENHANCEDMEMPOOLMIB_Cempmempooltable
// A table of memory pool monitoring entries for all
// physical entities on a managed system.
type CISCOENHANCEDMEMPOOLMIB_Cempmempooltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the memory pool monitoring table. The type is slice of
    // CISCOENHANCEDMEMPOOLMIB_Cempmempooltable_Cempmempoolentry.
    Cempmempoolentry []CISCOENHANCEDMEMPOOLMIB_Cempmempooltable_Cempmempoolentry
}

func (cempmempooltable *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable) GetEntityData() *types.CommonEntityData {
    cempmempooltable.EntityData.YFilter = cempmempooltable.YFilter
    cempmempooltable.EntityData.YangName = "cempMemPoolTable"
    cempmempooltable.EntityData.BundleName = "cisco_ios_xe"
    cempmempooltable.EntityData.ParentYangName = "CISCO-ENHANCED-MEMPOOL-MIB"
    cempmempooltable.EntityData.SegmentPath = "cempMemPoolTable"
    cempmempooltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cempmempooltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cempmempooltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cempmempooltable.EntityData.Children = make(map[string]types.YChild)
    cempmempooltable.EntityData.Children["cempMemPoolEntry"] = types.YChild{"Cempmempoolentry", nil}
    for i := range cempmempooltable.Cempmempoolentry {
        cempmempooltable.EntityData.Children[types.GetSegmentPath(&cempmempooltable.Cempmempoolentry[i])] = types.YChild{"Cempmempoolentry", &cempmempooltable.Cempmempoolentry[i]}
    }
    cempmempooltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cempmempooltable.EntityData)
}

// CISCOENHANCEDMEMPOOLMIB_Cempmempooltable_Cempmempoolentry
// An entry in the memory pool monitoring table.
type CISCOENHANCEDMEMPOOLMIB_Cempmempooltable_Cempmempoolentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. Within each physical entity, the unique value
    // greater than zero, used to represent each memory pool.   It is recommended
    // that values are assigned contiguously starting from 1. The type is
    // interface{} with range: 1..2147483647.
    Cempmempoolindex interface{}

    // The type of memory pool for which this entry contains information. The type
    // is CempMemPoolTypes.
    Cempmempooltype interface{}

    // A textual name assigned to the memory pool. This object is suitable for
    // output to a human operator, and may also be used to distinguish among the
    // various pool types. The type is string.
    Cempmempoolname interface{}

    // An indication of the platform-specific memory pool type. The associated
    // instance of cempMemPoolType is used to indicate the general type of memory
    // pool.  If no platform specific memory hardware type identifier exists for
    // this physical entity, or the value is unknown by this agent, then the value
    // { 0 0 } is returned. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    Cempmempoolplatformmemory interface{}

    // Indicates whether or not this memory pool has an alternate pool configured.
    // Alternate pools are used for fallback when the current pool runs out of
    // memory.  If an instance of this object has a value of zero, then this pool
    // does not have an alternate.  Otherwise the value of this object is the same
    // as the value of cempMemPoolType of the alternate pool. The type is
    // interface{} with range: 0..2147483647.
    Cempmempoolalternate interface{}

    // Indicates whether or not cempMemPoolUsed, cempMemPoolFree,
    // cempMemPoolLargestFree and  cempMemPoolLowestFree in this entry contain
    // accurate  data. If an instance of this object has the value  false (which
    // in and of itself indicates an internal  error condition), the values of
    // these objects in the conceptual row may contain inaccurate  information
    // (specifically, the reported values may be  less than the actual values).
    // The type is bool.
    Cempmempoolvalid interface{}

    // Indicates the number of bytes from the memory pool that are currently in
    // use by applications on the physical entity. The type is interface{} with
    // range: 0..4294967295. Units are bytes.
    Cempmempoolused interface{}

    // Indicates the number of bytes from the memory pool that are currently
    // unused on the physical entity.  Note that the sum of cempMemPoolUsed and
    // cempMemPoolFree  is the total amount of memory in the pool. The type is
    // interface{} with range: 0..4294967295. Units are bytes.
    Cempmempoolfree interface{}

    // Indicates the largest number of contiguous bytes from the memory pool that
    // are currently unused on the physical entity. The type is interface{} with
    // range: 0..4294967295. Units are bytes.
    Cempmempoollargestfree interface{}

    // The lowest amount of available memory in the memory pool recorded at any
    // time during the operation of the system. The type is interface{} with
    // range: 0..4294967295. Units are bytes.
    Cempmempoollowestfree interface{}

    // Indicates the lowest number of bytes from the memory pool that have been
    // used by applications on the physical entity since sysUpTime.Similarly,the
    // Used High Watermark indicates the largest number of bytes from the memory
    // pool that have been used by applications on the physical entity since
    // sysUpTime.This can be derived as follows: Used High Watermark =
    // cempMemPoolUsed + cempMemPoolFree  - cempMemPoolLowestFree. The type is
    // interface{} with range: 0..4294967295.
    Cempmempoolusedlowwatermark interface{}

    // Indicates the number of successful allocations from the memory pool. The
    // type is interface{} with range: 0..4294967295.
    Cempmempoolallochit interface{}

    // Indicates the number of unsuccessful allocations from the memory pool. The
    // type is interface{} with range: 0..4294967295.
    Cempmempoolallocmiss interface{}

    // Indicates the number of successful frees/ deallocations from the memory
    // pool. The type is interface{} with range: 0..4294967295.
    Cempmempoolfreehit interface{}

    // Indicates the number of unsuccessful attempts to free/deallocate memory
    // from the memory pool. For example, this could be due to ownership errors 
    // where the application that did not assign the  memory is trying to free it.
    // The type is interface{} with range: 0..4294967295.
    Cempmempoolfreemiss interface{}

    // Indicates the number of bytes from the memory pool that are currently
    // shared on the physical entity. The type is interface{} with range:
    // 0..4294967295. Units are bytes.
    Cempmempoolshared interface{}

    // This object represents the upper 32-bits of cempMemPoolUsed. This object
    // needs to be supported only if the used bytes in the memory pool exceeds
    // 32-bits, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are bytes.
    Cempmempoolusedovrflw interface{}

    // Indicates the number of bytes from the memory pool that are currently in
    // use by applications on the physical entity. This object is a 64-bit version
    // of cempMemPoolUsed. The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    Cempmempoolhcused interface{}

    // This object represents the upper 32-bits of cempMemPoolFree. This object
    // needs to be supported only if the unused bytes in the memory pool exceeds
    // 32-bits, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are bytes.
    Cempmempoolfreeovrflw interface{}

    // Indicates the number of bytes from the memory pool that are currently
    // unused on the physical entity. This object is a 64-bit version of
    // cempMemPoolFree. The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    Cempmempoolhcfree interface{}

    // This object represents the upper 32-bits of cempMemPoolLargestFree. This
    // object needs to  be supported only if the value of  cempMemPoolLargestFree
    // exceeds 32-bits, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are bytes.
    Cempmempoollargestfreeovrflw interface{}

    // Indicates the largest number of contiguous bytes from the memory pool that
    // are currently unused on the physical entity. This object is a 64-bit
    // version of cempMemPoolLargestFree. The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    Cempmempoolhclargestfree interface{}

    // This object represents the upper 32-bits of cempMemPoolLowestFree. This
    // object needs to be supported only if the value of cempMemPoolLowestFree
    // exceeds 32-bits, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are bytes.
    Cempmempoollowestfreeovrflw interface{}

    // The lowest amount of available memory in the memory pool recorded at any
    // time during the operation of the system. This object is a 64-bit version of
    // cempMemPoolLowestFree. The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    Cempmempoolhclowestfree interface{}

    // This object represents the upper 32-bits of cempMemPoolUsedLowWaterMark.
    // This object needs to be supported only if the value of
    // cempMemPoolUsedLowWaterMark exceeds 32-bits, otherwise this object value
    // would be set to 0. The type is interface{} with range: 0..4294967295. Units
    // are bytes.
    Cempmempoolusedlowwatermarkovrflw interface{}

    // Indicates the lowest number of bytes from the memory pool that have been
    // used by applications on the physical entity since sysUpTime. This object is
    // a 64-bit version of cempMemPoolUsedLowWaterMark. The type is interface{}
    // with range: 0..18446744073709551615. Units are bytes.
    Cempmempoolhcusedlowwatermark interface{}

    // This object represents the upper 32-bits of cempMemPoolShared. This object
    // needs to be supported only if the value of cempMemPoolShared exceeds
    // 32-bits, otherwise this object value would be set to 0. The type is
    // interface{} with range: 0..4294967295. Units are bytes.
    Cempmempoolsharedovrflw interface{}

    // Indicates the number of bytes from the memory pool that are currently
    // shared on the physical entity. This object is a 64-bit version of
    // cempMemPoolShared. The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    Cempmempoolhcshared interface{}
}

func (cempmempoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable_Cempmempoolentry) GetEntityData() *types.CommonEntityData {
    cempmempoolentry.EntityData.YFilter = cempmempoolentry.YFilter
    cempmempoolentry.EntityData.YangName = "cempMemPoolEntry"
    cempmempoolentry.EntityData.BundleName = "cisco_ios_xe"
    cempmempoolentry.EntityData.ParentYangName = "cempMemPoolTable"
    cempmempoolentry.EntityData.SegmentPath = "cempMemPoolEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cempmempoolentry.Entphysicalindex) + "']" + "[cempMemPoolIndex='" + fmt.Sprintf("%v", cempmempoolentry.Cempmempoolindex) + "']"
    cempmempoolentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cempmempoolentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cempmempoolentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cempmempoolentry.EntityData.Children = make(map[string]types.YChild)
    cempmempoolentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cempmempoolentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cempmempoolentry.Entphysicalindex}
    cempmempoolentry.EntityData.Leafs["cempMemPoolIndex"] = types.YLeaf{"Cempmempoolindex", cempmempoolentry.Cempmempoolindex}
    cempmempoolentry.EntityData.Leafs["cempMemPoolType"] = types.YLeaf{"Cempmempooltype", cempmempoolentry.Cempmempooltype}
    cempmempoolentry.EntityData.Leafs["cempMemPoolName"] = types.YLeaf{"Cempmempoolname", cempmempoolentry.Cempmempoolname}
    cempmempoolentry.EntityData.Leafs["cempMemPoolPlatformMemory"] = types.YLeaf{"Cempmempoolplatformmemory", cempmempoolentry.Cempmempoolplatformmemory}
    cempmempoolentry.EntityData.Leafs["cempMemPoolAlternate"] = types.YLeaf{"Cempmempoolalternate", cempmempoolentry.Cempmempoolalternate}
    cempmempoolentry.EntityData.Leafs["cempMemPoolValid"] = types.YLeaf{"Cempmempoolvalid", cempmempoolentry.Cempmempoolvalid}
    cempmempoolentry.EntityData.Leafs["cempMemPoolUsed"] = types.YLeaf{"Cempmempoolused", cempmempoolentry.Cempmempoolused}
    cempmempoolentry.EntityData.Leafs["cempMemPoolFree"] = types.YLeaf{"Cempmempoolfree", cempmempoolentry.Cempmempoolfree}
    cempmempoolentry.EntityData.Leafs["cempMemPoolLargestFree"] = types.YLeaf{"Cempmempoollargestfree", cempmempoolentry.Cempmempoollargestfree}
    cempmempoolentry.EntityData.Leafs["cempMemPoolLowestFree"] = types.YLeaf{"Cempmempoollowestfree", cempmempoolentry.Cempmempoollowestfree}
    cempmempoolentry.EntityData.Leafs["cempMemPoolUsedLowWaterMark"] = types.YLeaf{"Cempmempoolusedlowwatermark", cempmempoolentry.Cempmempoolusedlowwatermark}
    cempmempoolentry.EntityData.Leafs["cempMemPoolAllocHit"] = types.YLeaf{"Cempmempoolallochit", cempmempoolentry.Cempmempoolallochit}
    cempmempoolentry.EntityData.Leafs["cempMemPoolAllocMiss"] = types.YLeaf{"Cempmempoolallocmiss", cempmempoolentry.Cempmempoolallocmiss}
    cempmempoolentry.EntityData.Leafs["cempMemPoolFreeHit"] = types.YLeaf{"Cempmempoolfreehit", cempmempoolentry.Cempmempoolfreehit}
    cempmempoolentry.EntityData.Leafs["cempMemPoolFreeMiss"] = types.YLeaf{"Cempmempoolfreemiss", cempmempoolentry.Cempmempoolfreemiss}
    cempmempoolentry.EntityData.Leafs["cempMemPoolShared"] = types.YLeaf{"Cempmempoolshared", cempmempoolentry.Cempmempoolshared}
    cempmempoolentry.EntityData.Leafs["cempMemPoolUsedOvrflw"] = types.YLeaf{"Cempmempoolusedovrflw", cempmempoolentry.Cempmempoolusedovrflw}
    cempmempoolentry.EntityData.Leafs["cempMemPoolHCUsed"] = types.YLeaf{"Cempmempoolhcused", cempmempoolentry.Cempmempoolhcused}
    cempmempoolentry.EntityData.Leafs["cempMemPoolFreeOvrflw"] = types.YLeaf{"Cempmempoolfreeovrflw", cempmempoolentry.Cempmempoolfreeovrflw}
    cempmempoolentry.EntityData.Leafs["cempMemPoolHCFree"] = types.YLeaf{"Cempmempoolhcfree", cempmempoolentry.Cempmempoolhcfree}
    cempmempoolentry.EntityData.Leafs["cempMemPoolLargestFreeOvrflw"] = types.YLeaf{"Cempmempoollargestfreeovrflw", cempmempoolentry.Cempmempoollargestfreeovrflw}
    cempmempoolentry.EntityData.Leafs["cempMemPoolHCLargestFree"] = types.YLeaf{"Cempmempoolhclargestfree", cempmempoolentry.Cempmempoolhclargestfree}
    cempmempoolentry.EntityData.Leafs["cempMemPoolLowestFreeOvrflw"] = types.YLeaf{"Cempmempoollowestfreeovrflw", cempmempoolentry.Cempmempoollowestfreeovrflw}
    cempmempoolentry.EntityData.Leafs["cempMemPoolHCLowestFree"] = types.YLeaf{"Cempmempoolhclowestfree", cempmempoolentry.Cempmempoolhclowestfree}
    cempmempoolentry.EntityData.Leafs["cempMemPoolUsedLowWaterMarkOvrflw"] = types.YLeaf{"Cempmempoolusedlowwatermarkovrflw", cempmempoolentry.Cempmempoolusedlowwatermarkovrflw}
    cempmempoolentry.EntityData.Leafs["cempMemPoolHCUsedLowWaterMark"] = types.YLeaf{"Cempmempoolhcusedlowwatermark", cempmempoolentry.Cempmempoolhcusedlowwatermark}
    cempmempoolentry.EntityData.Leafs["cempMemPoolSharedOvrflw"] = types.YLeaf{"Cempmempoolsharedovrflw", cempmempoolentry.Cempmempoolsharedovrflw}
    cempmempoolentry.EntityData.Leafs["cempMemPoolHCShared"] = types.YLeaf{"Cempmempoolhcshared", cempmempoolentry.Cempmempoolhcshared}
    return &(cempmempoolentry.EntityData)
}

// CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable
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
type CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This contains all the memory buffer pool configurations object values. The 
    // entPhysicalIndex identifies the entity on which memory buffer pools are
    // present. The type is slice of
    // CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable_Cempmembufferpoolentry.
    Cempmembufferpoolentry []CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable_Cempmembufferpoolentry
}

func (cempmembufferpooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable) GetEntityData() *types.CommonEntityData {
    cempmembufferpooltable.EntityData.YFilter = cempmembufferpooltable.YFilter
    cempmembufferpooltable.EntityData.YangName = "cempMemBufferPoolTable"
    cempmembufferpooltable.EntityData.BundleName = "cisco_ios_xe"
    cempmembufferpooltable.EntityData.ParentYangName = "CISCO-ENHANCED-MEMPOOL-MIB"
    cempmembufferpooltable.EntityData.SegmentPath = "cempMemBufferPoolTable"
    cempmembufferpooltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cempmembufferpooltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cempmembufferpooltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cempmembufferpooltable.EntityData.Children = make(map[string]types.YChild)
    cempmembufferpooltable.EntityData.Children["cempMemBufferPoolEntry"] = types.YChild{"Cempmembufferpoolentry", nil}
    for i := range cempmembufferpooltable.Cempmembufferpoolentry {
        cempmembufferpooltable.EntityData.Children[types.GetSegmentPath(&cempmembufferpooltable.Cempmembufferpoolentry[i])] = types.YChild{"Cempmembufferpoolentry", &cempmembufferpooltable.Cempmembufferpoolentry[i]}
    }
    cempmembufferpooltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cempmembufferpooltable.EntityData)
}

// CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable_Cempmembufferpoolentry
// This contains all the memory buffer pool
// configurations object values. The 
// entPhysicalIndex identifies the entity on which
// memory buffer pools are present.
type CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable_Cempmembufferpoolentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. Within a physical entity, a unique value used to
    // represent each buffer pool. The type is interface{} with range:
    // 0..4294967295.
    Cempmembufferpoolindex interface{}

    // This index corresponds to the memory pool (with cemMemPoolIndex as index in
    // cempMemPoolTable)  from which buffers are allocated. The type is
    // interface{} with range: 0..2147483647.
    Cempmembuffermempoolindex interface{}

    // A textual name assigned to the buffer pool. This object is suitable for
    // output to a human operator, and may also be used to distinguish among the
    // various buffer types. For example: 'Small', 'Big', 'Serial0/1' etc. The
    // type is string.
    Cempmembuffername interface{}

    // Boolean poolDynamic; if TRUE, the number of buffers in the pool is adjusted
    // (adding more packet buffers  or deleting excesses) dynamically by the
    // background  process. If FALSE, the number of buffers in the pool  is never
    // adjusted, even if it falls below the minimum, or to zero. The type is bool.
    Cempmembufferdynamic interface{}

    // Indicates the size of buffer element in number of bytes on the physical
    // entity. The type is interface{} with range: 0..4294967295. Units are bytes.
    Cempmembuffersize interface{}

    // Indicates the minimum number of free buffers allowed in the buffer pool or
    // low-water mark (lwm).  For example of its usage : If cempMemBufferFree <
    // cempMemBufferMin & pool is  dynamic, then signal for growth of particular
    // buffer pool. The type is interface{} with range: 0..4294967295.
    Cempmembuffermin interface{}

    // Indicates the maximum number of free buffers allowed in the buffer pool or
    // high-water mark (hwm). For example of its usage : If cempMemBufferFree >
    // cempMemBufferMax & pool is  dynamic, then signal for trim of particular
    // buffer pool. The type is interface{} with range: 0..4294967295.
    Cempmembuffermax interface{}

    // Indicates the total number of permanent buffers in the pool on the physical
    // entity. The type is interface{} with range: 0..4294967295.
    Cempmembufferpermanent interface{}

    // Indicates the initial number of temporary buffers in the pool on the
    // physical entity. This object  instructs the system to create this many
    // number of temporary extra buffers, just after a system restart.  A change
    // in this object will be effective only after a system restart. The type is
    // interface{} with range: 0..4294967295.
    Cempmembuffertransient interface{}

    // Indicates the total number of buffers (include allocated and free buffers)
    // in the buffer pool on the physical entity. The type is interface{} with
    // range: 0..4294967295.
    Cempmembuffertotal interface{}

    // Indicates the current number of free buffers in the buffer pool on the
    // physical entity. Note that the cempMemBufferFree is less than or equal  to
    // cempMemBufferTotal. The type is interface{} with range: 0..4294967295.
    Cempmembufferfree interface{}

    // Indicates the number of buffers successfully allocated from the buffer
    // pool. The type is interface{} with range: 0..4294967295.
    Cempmembufferhit interface{}

    // Indicates the number of times a buffer has been requested, but no buffers
    // were available in the buffer pool, or when there were fewer than min 
    // buffers(cempMemBufferMin) in the buffer pool. Note : For interface pools, a
    // miss is actually  a fall back to its corresponding public buffer pool. The
    // type is interface{} with range: 0..4294967295.
    Cempmembuffermiss interface{}

    // Indicates the number of successful frees/deallocations from the buffer
    // pool. The type is interface{} with range: 0..4294967295.
    Cempmembufferfreehit interface{}

    // Indicates the number of unsuccessful attempts to free/deallocate a buffer
    // from the buffer pool.  For example, this could be due to ownership errors
    // where the application that did not assign the  buffer is trying to free it.
    // The type is interface{} with range: 0..4294967295.
    Cempmembufferfreemiss interface{}

    // This value is the difference of the desired number of permanent buffer &
    // total number of permanent  buffers present in the pool. A positive value of
    // this object tells the number of buffers needed & a  negative value of the
    // object tells the extra number  of buffers in the pool. The type is
    // interface{} with range: -2147483648..2147483647.
    Cempmembufferpermchange interface{}

    // Indicates the peak number of buffers in pool on the physical entity. The
    // type is interface{} with range: 0..4294967295.
    Cempmembufferpeak interface{}

    // Indicates the time of most recent change in the peak number of buffers
    // (cempMemBufferPeak object) in the pool. The type is interface{} with range:
    // 0..4294967295.
    Cempmembufferpeaktime interface{}

    // The number of buffers that have been trimmed from the pool when the number
    // of free buffers  (cempMemBufferFree) exceeded the number of max allowed
    // buffers(cempMemBufferMax). The type is interface{} with range:
    // 0..4294967295.
    Cempmembuffertrim interface{}

    // The number of buffers that have been created in the pool when the number of
    // free buffers(cempMemBufferFree) was less than minimum(cempMemBufferMix).
    // The type is interface{} with range: 0..4294967295.
    Cempmembuffergrow interface{}

    // The number of failures to grant a buffer to a requester due to reasons
    // other than insufficient  memory. For example, in systems where there are 
    // different execution contexts, it may be too expensive to create new buffers
    // when running in certain contexts. In those cases it may be  preferable to
    // fail the request. The type is interface{} with range: 0..4294967295.
    Cempmembufferfailures interface{}

    // The number of times the system tried to create new buffers, but could not
    // due to insufficient free  memory in the system. The type is interface{}
    // with range: 0..4294967295.
    Cempmembuffernostorage interface{}
}

func (cempmembufferpoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable_Cempmembufferpoolentry) GetEntityData() *types.CommonEntityData {
    cempmembufferpoolentry.EntityData.YFilter = cempmembufferpoolentry.YFilter
    cempmembufferpoolentry.EntityData.YangName = "cempMemBufferPoolEntry"
    cempmembufferpoolentry.EntityData.BundleName = "cisco_ios_xe"
    cempmembufferpoolentry.EntityData.ParentYangName = "cempMemBufferPoolTable"
    cempmembufferpoolentry.EntityData.SegmentPath = "cempMemBufferPoolEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cempmembufferpoolentry.Entphysicalindex) + "']" + "[cempMemBufferPoolIndex='" + fmt.Sprintf("%v", cempmembufferpoolentry.Cempmembufferpoolindex) + "']"
    cempmembufferpoolentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cempmembufferpoolentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cempmembufferpoolentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cempmembufferpoolentry.EntityData.Children = make(map[string]types.YChild)
    cempmembufferpoolentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cempmembufferpoolentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cempmembufferpoolentry.Entphysicalindex}
    cempmembufferpoolentry.EntityData.Leafs["cempMemBufferPoolIndex"] = types.YLeaf{"Cempmembufferpoolindex", cempmembufferpoolentry.Cempmembufferpoolindex}
    cempmembufferpoolentry.EntityData.Leafs["cempMemBufferMemPoolIndex"] = types.YLeaf{"Cempmembuffermempoolindex", cempmembufferpoolentry.Cempmembuffermempoolindex}
    cempmembufferpoolentry.EntityData.Leafs["cempMemBufferName"] = types.YLeaf{"Cempmembuffername", cempmembufferpoolentry.Cempmembuffername}
    cempmembufferpoolentry.EntityData.Leafs["cempMemBufferDynamic"] = types.YLeaf{"Cempmembufferdynamic", cempmembufferpoolentry.Cempmembufferdynamic}
    cempmembufferpoolentry.EntityData.Leafs["cempMemBufferSize"] = types.YLeaf{"Cempmembuffersize", cempmembufferpoolentry.Cempmembuffersize}
    cempmembufferpoolentry.EntityData.Leafs["cempMemBufferMin"] = types.YLeaf{"Cempmembuffermin", cempmembufferpoolentry.Cempmembuffermin}
    cempmembufferpoolentry.EntityData.Leafs["cempMemBufferMax"] = types.YLeaf{"Cempmembuffermax", cempmembufferpoolentry.Cempmembuffermax}
    cempmembufferpoolentry.EntityData.Leafs["cempMemBufferPermanent"] = types.YLeaf{"Cempmembufferpermanent", cempmembufferpoolentry.Cempmembufferpermanent}
    cempmembufferpoolentry.EntityData.Leafs["cempMemBufferTransient"] = types.YLeaf{"Cempmembuffertransient", cempmembufferpoolentry.Cempmembuffertransient}
    cempmembufferpoolentry.EntityData.Leafs["cempMemBufferTotal"] = types.YLeaf{"Cempmembuffertotal", cempmembufferpoolentry.Cempmembuffertotal}
    cempmembufferpoolentry.EntityData.Leafs["cempMemBufferFree"] = types.YLeaf{"Cempmembufferfree", cempmembufferpoolentry.Cempmembufferfree}
    cempmembufferpoolentry.EntityData.Leafs["cempMemBufferHit"] = types.YLeaf{"Cempmembufferhit", cempmembufferpoolentry.Cempmembufferhit}
    cempmembufferpoolentry.EntityData.Leafs["cempMemBufferMiss"] = types.YLeaf{"Cempmembuffermiss", cempmembufferpoolentry.Cempmembuffermiss}
    cempmembufferpoolentry.EntityData.Leafs["cempMemBufferFreeHit"] = types.YLeaf{"Cempmembufferfreehit", cempmembufferpoolentry.Cempmembufferfreehit}
    cempmembufferpoolentry.EntityData.Leafs["cempMemBufferFreeMiss"] = types.YLeaf{"Cempmembufferfreemiss", cempmembufferpoolentry.Cempmembufferfreemiss}
    cempmembufferpoolentry.EntityData.Leafs["cempMemBufferPermChange"] = types.YLeaf{"Cempmembufferpermchange", cempmembufferpoolentry.Cempmembufferpermchange}
    cempmembufferpoolentry.EntityData.Leafs["cempMemBufferPeak"] = types.YLeaf{"Cempmembufferpeak", cempmembufferpoolentry.Cempmembufferpeak}
    cempmembufferpoolentry.EntityData.Leafs["cempMemBufferPeakTime"] = types.YLeaf{"Cempmembufferpeaktime", cempmembufferpoolentry.Cempmembufferpeaktime}
    cempmembufferpoolentry.EntityData.Leafs["cempMemBufferTrim"] = types.YLeaf{"Cempmembuffertrim", cempmembufferpoolentry.Cempmembuffertrim}
    cempmembufferpoolentry.EntityData.Leafs["cempMemBufferGrow"] = types.YLeaf{"Cempmembuffergrow", cempmembufferpoolentry.Cempmembuffergrow}
    cempmembufferpoolentry.EntityData.Leafs["cempMemBufferFailures"] = types.YLeaf{"Cempmembufferfailures", cempmembufferpoolentry.Cempmembufferfailures}
    cempmembufferpoolentry.EntityData.Leafs["cempMemBufferNoStorage"] = types.YLeaf{"Cempmembuffernostorage", cempmembufferpoolentry.Cempmembuffernostorage}
    return &(cempmembufferpoolentry.EntityData)
}

// CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable
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
type CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Each entry represents one of the cache buffer pools available in the system
    // and it contains the parameters configured for it. Note :
    // cempMemBufferCachePoolTable has a sparse dependency with
    // cempMemBufferPoolTable (i.e all the entires in cempMemBufferPoolTable need
    // not have an entry in cempMemBufferCachePoolTable. The type is slice of
    // CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable_Cempmembuffercachepoolentry.
    Cempmembuffercachepoolentry []CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable_Cempmembuffercachepoolentry
}

func (cempmembuffercachepooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable) GetEntityData() *types.CommonEntityData {
    cempmembuffercachepooltable.EntityData.YFilter = cempmembuffercachepooltable.YFilter
    cempmembuffercachepooltable.EntityData.YangName = "cempMemBufferCachePoolTable"
    cempmembuffercachepooltable.EntityData.BundleName = "cisco_ios_xe"
    cempmembuffercachepooltable.EntityData.ParentYangName = "CISCO-ENHANCED-MEMPOOL-MIB"
    cempmembuffercachepooltable.EntityData.SegmentPath = "cempMemBufferCachePoolTable"
    cempmembuffercachepooltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cempmembuffercachepooltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cempmembuffercachepooltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cempmembuffercachepooltable.EntityData.Children = make(map[string]types.YChild)
    cempmembuffercachepooltable.EntityData.Children["cempMemBufferCachePoolEntry"] = types.YChild{"Cempmembuffercachepoolentry", nil}
    for i := range cempmembuffercachepooltable.Cempmembuffercachepoolentry {
        cempmembuffercachepooltable.EntityData.Children[types.GetSegmentPath(&cempmembuffercachepooltable.Cempmembuffercachepoolentry[i])] = types.YChild{"Cempmembuffercachepoolentry", &cempmembuffercachepooltable.Cempmembuffercachepoolentry[i]}
    }
    cempmembuffercachepooltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cempmembuffercachepooltable.EntityData)
}

// CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable_Cempmembuffercachepoolentry
// Each entry represents one of the cache buffer pools
// available in the system and it contains the
// parameters configured for it.
// Note : cempMemBufferCachePoolTable has a sparse
// dependency with cempMemBufferPoolTable (i.e all the
// entires in cempMemBufferPoolTable need not have an
// entry in cempMemBufferCachePoolTable.
type CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable_Cempmembuffercachepoolentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. The type is string with range: 0..4294967295.
    // Refers to
    // cisco_enhanced_mempool_mib.CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable_Cempmembufferpoolentry_Cempmembufferpoolindex
    Cempmembufferpoolindex interface{}

    // Indicates the number of buffers in the cache pool on the physical entity.
    // The type is interface{} with range: 0..4294967295.
    Cempmembuffercachesize interface{}

    // Indicates the maximum number of free buffers allowed in the cache pool. The
    // type is interface{} with range: 0..4294967295.
    Cempmembuffercachetotal interface{}

    // Indicates the number of cache buffers from the pool that are currently used
    // on the physical entity. Note that the cempMemBufferCacheUsed is less than
    // or  equal to cempMemBufferCacheTotal. The type is interface{} with range:
    // 0..4294967295.
    Cempmembuffercacheused interface{}

    // Indicates the number of buffers successfully allocated from the cache pool.
    // The type is interface{} with range: 0..4294967295.
    Cempmembuffercachehit interface{}

    // Indicates the number of times a buffer has been requested, but no buffers
    // were available in the cache pool. The type is interface{} with range:
    // 0..4294967295.
    Cempmembuffercachemiss interface{}

    // Indicates the threshold limit for number of cache buffers
    // used(cempMemBufferCacheUsed). The type is interface{} with range:
    // 0..4294967295.
    Cempmembuffercachethreshold interface{}

    // Indicates how many times the number of cache buffers
    // used(cempMemBufferCacheUsed) has crossed the threshold
    // value(cempMemBufferCacheThreshold). The type is interface{} with range:
    // 0..4294967295.
    Cempmembuffercachethresholdcount interface{}
}

func (cempmembuffercachepoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable_Cempmembuffercachepoolentry) GetEntityData() *types.CommonEntityData {
    cempmembuffercachepoolentry.EntityData.YFilter = cempmembuffercachepoolentry.YFilter
    cempmembuffercachepoolentry.EntityData.YangName = "cempMemBufferCachePoolEntry"
    cempmembuffercachepoolentry.EntityData.BundleName = "cisco_ios_xe"
    cempmembuffercachepoolentry.EntityData.ParentYangName = "cempMemBufferCachePoolTable"
    cempmembuffercachepoolentry.EntityData.SegmentPath = "cempMemBufferCachePoolEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cempmembuffercachepoolentry.Entphysicalindex) + "']" + "[cempMemBufferPoolIndex='" + fmt.Sprintf("%v", cempmembuffercachepoolentry.Cempmembufferpoolindex) + "']"
    cempmembuffercachepoolentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cempmembuffercachepoolentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cempmembuffercachepoolentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cempmembuffercachepoolentry.EntityData.Children = make(map[string]types.YChild)
    cempmembuffercachepoolentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cempmembuffercachepoolentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cempmembuffercachepoolentry.Entphysicalindex}
    cempmembuffercachepoolentry.EntityData.Leafs["cempMemBufferPoolIndex"] = types.YLeaf{"Cempmembufferpoolindex", cempmembuffercachepoolentry.Cempmembufferpoolindex}
    cempmembuffercachepoolentry.EntityData.Leafs["cempMemBufferCacheSize"] = types.YLeaf{"Cempmembuffercachesize", cempmembuffercachepoolentry.Cempmembuffercachesize}
    cempmembuffercachepoolentry.EntityData.Leafs["cempMemBufferCacheTotal"] = types.YLeaf{"Cempmembuffercachetotal", cempmembuffercachepoolentry.Cempmembuffercachetotal}
    cempmembuffercachepoolentry.EntityData.Leafs["cempMemBufferCacheUsed"] = types.YLeaf{"Cempmembuffercacheused", cempmembuffercachepoolentry.Cempmembuffercacheused}
    cempmembuffercachepoolentry.EntityData.Leafs["cempMemBufferCacheHit"] = types.YLeaf{"Cempmembuffercachehit", cempmembuffercachepoolentry.Cempmembuffercachehit}
    cempmembuffercachepoolentry.EntityData.Leafs["cempMemBufferCacheMiss"] = types.YLeaf{"Cempmembuffercachemiss", cempmembuffercachepoolentry.Cempmembuffercachemiss}
    cempmembuffercachepoolentry.EntityData.Leafs["cempMemBufferCacheThreshold"] = types.YLeaf{"Cempmembuffercachethreshold", cempmembuffercachepoolentry.Cempmembuffercachethreshold}
    cempmembuffercachepoolentry.EntityData.Leafs["cempMemBufferCacheThresholdCount"] = types.YLeaf{"Cempmembuffercachethresholdcount", cempmembuffercachepoolentry.Cempmembuffercachethresholdcount}
    return &(cempmembuffercachepoolentry.EntityData)
}

