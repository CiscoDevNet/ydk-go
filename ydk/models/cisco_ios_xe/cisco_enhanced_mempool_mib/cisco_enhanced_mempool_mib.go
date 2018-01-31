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
    parent types.Entity
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

func (cISCOENHANCEDMEMPOOLMIB *CISCOENHANCEDMEMPOOLMIB) GetFilter() yfilter.YFilter { return cISCOENHANCEDMEMPOOLMIB.YFilter }

func (cISCOENHANCEDMEMPOOLMIB *CISCOENHANCEDMEMPOOLMIB) SetFilter(yf yfilter.YFilter) { cISCOENHANCEDMEMPOOLMIB.YFilter = yf }

func (cISCOENHANCEDMEMPOOLMIB *CISCOENHANCEDMEMPOOLMIB) GetGoName(yname string) string {
    if yname == "cempNotificationConfig" { return "Cempnotificationconfig" }
    if yname == "cempMemPoolTable" { return "Cempmempooltable" }
    if yname == "cempMemBufferPoolTable" { return "Cempmembufferpooltable" }
    if yname == "cempMemBufferCachePoolTable" { return "Cempmembuffercachepooltable" }
    return ""
}

func (cISCOENHANCEDMEMPOOLMIB *CISCOENHANCEDMEMPOOLMIB) GetSegmentPath() string {
    return "CISCO-ENHANCED-MEMPOOL-MIB:CISCO-ENHANCED-MEMPOOL-MIB"
}

func (cISCOENHANCEDMEMPOOLMIB *CISCOENHANCEDMEMPOOLMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cempNotificationConfig" {
        return &cISCOENHANCEDMEMPOOLMIB.Cempnotificationconfig
    }
    if childYangName == "cempMemPoolTable" {
        return &cISCOENHANCEDMEMPOOLMIB.Cempmempooltable
    }
    if childYangName == "cempMemBufferPoolTable" {
        return &cISCOENHANCEDMEMPOOLMIB.Cempmembufferpooltable
    }
    if childYangName == "cempMemBufferCachePoolTable" {
        return &cISCOENHANCEDMEMPOOLMIB.Cempmembuffercachepooltable
    }
    return nil
}

func (cISCOENHANCEDMEMPOOLMIB *CISCOENHANCEDMEMPOOLMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cempNotificationConfig"] = &cISCOENHANCEDMEMPOOLMIB.Cempnotificationconfig
    children["cempMemPoolTable"] = &cISCOENHANCEDMEMPOOLMIB.Cempmempooltable
    children["cempMemBufferPoolTable"] = &cISCOENHANCEDMEMPOOLMIB.Cempmembufferpooltable
    children["cempMemBufferCachePoolTable"] = &cISCOENHANCEDMEMPOOLMIB.Cempmembuffercachepooltable
    return children
}

func (cISCOENHANCEDMEMPOOLMIB *CISCOENHANCEDMEMPOOLMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOENHANCEDMEMPOOLMIB *CISCOENHANCEDMEMPOOLMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOENHANCEDMEMPOOLMIB *CISCOENHANCEDMEMPOOLMIB) GetYangName() string { return "CISCO-ENHANCED-MEMPOOL-MIB" }

func (cISCOENHANCEDMEMPOOLMIB *CISCOENHANCEDMEMPOOLMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOENHANCEDMEMPOOLMIB *CISCOENHANCEDMEMPOOLMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOENHANCEDMEMPOOLMIB *CISCOENHANCEDMEMPOOLMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOENHANCEDMEMPOOLMIB *CISCOENHANCEDMEMPOOLMIB) SetParent(parent types.Entity) { cISCOENHANCEDMEMPOOLMIB.parent = parent }

func (cISCOENHANCEDMEMPOOLMIB *CISCOENHANCEDMEMPOOLMIB) GetParent() types.Entity { return cISCOENHANCEDMEMPOOLMIB.parent }

func (cISCOENHANCEDMEMPOOLMIB *CISCOENHANCEDMEMPOOLMIB) GetParentYangName() string { return "CISCO-ENHANCED-MEMPOOL-MIB" }

// CISCOENHANCEDMEMPOOLMIB_Cempnotificationconfig
type CISCOENHANCEDMEMPOOLMIB_Cempnotificationconfig struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This variable controls generation of the cempMemBufferNotify.  When this
    // variable is 'true', generation of cempMemBufferNotify is enabled.  When
    // this variable is 'false', generation of cempMemBufferNotify is disabled.
    // The type is bool.
    Cempmembuffernotifyenabled interface{}
}

func (cempnotificationconfig *CISCOENHANCEDMEMPOOLMIB_Cempnotificationconfig) GetFilter() yfilter.YFilter { return cempnotificationconfig.YFilter }

func (cempnotificationconfig *CISCOENHANCEDMEMPOOLMIB_Cempnotificationconfig) SetFilter(yf yfilter.YFilter) { cempnotificationconfig.YFilter = yf }

func (cempnotificationconfig *CISCOENHANCEDMEMPOOLMIB_Cempnotificationconfig) GetGoName(yname string) string {
    if yname == "cempMemBufferNotifyEnabled" { return "Cempmembuffernotifyenabled" }
    return ""
}

func (cempnotificationconfig *CISCOENHANCEDMEMPOOLMIB_Cempnotificationconfig) GetSegmentPath() string {
    return "cempNotificationConfig"
}

func (cempnotificationconfig *CISCOENHANCEDMEMPOOLMIB_Cempnotificationconfig) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cempnotificationconfig *CISCOENHANCEDMEMPOOLMIB_Cempnotificationconfig) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cempnotificationconfig *CISCOENHANCEDMEMPOOLMIB_Cempnotificationconfig) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cempMemBufferNotifyEnabled"] = cempnotificationconfig.Cempmembuffernotifyenabled
    return leafs
}

func (cempnotificationconfig *CISCOENHANCEDMEMPOOLMIB_Cempnotificationconfig) GetBundleName() string { return "cisco_ios_xe" }

func (cempnotificationconfig *CISCOENHANCEDMEMPOOLMIB_Cempnotificationconfig) GetYangName() string { return "cempNotificationConfig" }

func (cempnotificationconfig *CISCOENHANCEDMEMPOOLMIB_Cempnotificationconfig) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cempnotificationconfig *CISCOENHANCEDMEMPOOLMIB_Cempnotificationconfig) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cempnotificationconfig *CISCOENHANCEDMEMPOOLMIB_Cempnotificationconfig) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cempnotificationconfig *CISCOENHANCEDMEMPOOLMIB_Cempnotificationconfig) SetParent(parent types.Entity) { cempnotificationconfig.parent = parent }

func (cempnotificationconfig *CISCOENHANCEDMEMPOOLMIB_Cempnotificationconfig) GetParent() types.Entity { return cempnotificationconfig.parent }

func (cempnotificationconfig *CISCOENHANCEDMEMPOOLMIB_Cempnotificationconfig) GetParentYangName() string { return "CISCO-ENHANCED-MEMPOOL-MIB" }

// CISCOENHANCEDMEMPOOLMIB_Cempmempooltable
// A table of memory pool monitoring entries for all
// physical entities on a managed system.
type CISCOENHANCEDMEMPOOLMIB_Cempmempooltable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the memory pool monitoring table. The type is slice of
    // CISCOENHANCEDMEMPOOLMIB_Cempmempooltable_Cempmempoolentry.
    Cempmempoolentry []CISCOENHANCEDMEMPOOLMIB_Cempmempooltable_Cempmempoolentry
}

func (cempmempooltable *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable) GetFilter() yfilter.YFilter { return cempmempooltable.YFilter }

func (cempmempooltable *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable) SetFilter(yf yfilter.YFilter) { cempmempooltable.YFilter = yf }

func (cempmempooltable *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable) GetGoName(yname string) string {
    if yname == "cempMemPoolEntry" { return "Cempmempoolentry" }
    return ""
}

func (cempmempooltable *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable) GetSegmentPath() string {
    return "cempMemPoolTable"
}

func (cempmempooltable *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cempMemPoolEntry" {
        for _, c := range cempmempooltable.Cempmempoolentry {
            if cempmempooltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENHANCEDMEMPOOLMIB_Cempmempooltable_Cempmempoolentry{}
        cempmempooltable.Cempmempoolentry = append(cempmempooltable.Cempmempoolentry, child)
        return &cempmempooltable.Cempmempoolentry[len(cempmempooltable.Cempmempoolentry)-1]
    }
    return nil
}

func (cempmempooltable *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cempmempooltable.Cempmempoolentry {
        children[cempmempooltable.Cempmempoolentry[i].GetSegmentPath()] = &cempmempooltable.Cempmempoolentry[i]
    }
    return children
}

func (cempmempooltable *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cempmempooltable *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable) GetBundleName() string { return "cisco_ios_xe" }

func (cempmempooltable *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable) GetYangName() string { return "cempMemPoolTable" }

func (cempmempooltable *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cempmempooltable *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cempmempooltable *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cempmempooltable *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable) SetParent(parent types.Entity) { cempmempooltable.parent = parent }

func (cempmempooltable *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable) GetParent() types.Entity { return cempmempooltable.parent }

func (cempmempooltable *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable) GetParentYangName() string { return "CISCO-ENHANCED-MEMPOOL-MIB" }

// CISCOENHANCEDMEMPOOLMIB_Cempmempooltable_Cempmempoolentry
// An entry in the memory pool monitoring table.
type CISCOENHANCEDMEMPOOLMIB_Cempmempooltable_Cempmempoolentry struct {
    parent types.Entity
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
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
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

func (cempmempoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable_Cempmempoolentry) GetFilter() yfilter.YFilter { return cempmempoolentry.YFilter }

func (cempmempoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable_Cempmempoolentry) SetFilter(yf yfilter.YFilter) { cempmempoolentry.YFilter = yf }

func (cempmempoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable_Cempmempoolentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cempMemPoolIndex" { return "Cempmempoolindex" }
    if yname == "cempMemPoolType" { return "Cempmempooltype" }
    if yname == "cempMemPoolName" { return "Cempmempoolname" }
    if yname == "cempMemPoolPlatformMemory" { return "Cempmempoolplatformmemory" }
    if yname == "cempMemPoolAlternate" { return "Cempmempoolalternate" }
    if yname == "cempMemPoolValid" { return "Cempmempoolvalid" }
    if yname == "cempMemPoolUsed" { return "Cempmempoolused" }
    if yname == "cempMemPoolFree" { return "Cempmempoolfree" }
    if yname == "cempMemPoolLargestFree" { return "Cempmempoollargestfree" }
    if yname == "cempMemPoolLowestFree" { return "Cempmempoollowestfree" }
    if yname == "cempMemPoolUsedLowWaterMark" { return "Cempmempoolusedlowwatermark" }
    if yname == "cempMemPoolAllocHit" { return "Cempmempoolallochit" }
    if yname == "cempMemPoolAllocMiss" { return "Cempmempoolallocmiss" }
    if yname == "cempMemPoolFreeHit" { return "Cempmempoolfreehit" }
    if yname == "cempMemPoolFreeMiss" { return "Cempmempoolfreemiss" }
    if yname == "cempMemPoolShared" { return "Cempmempoolshared" }
    if yname == "cempMemPoolUsedOvrflw" { return "Cempmempoolusedovrflw" }
    if yname == "cempMemPoolHCUsed" { return "Cempmempoolhcused" }
    if yname == "cempMemPoolFreeOvrflw" { return "Cempmempoolfreeovrflw" }
    if yname == "cempMemPoolHCFree" { return "Cempmempoolhcfree" }
    if yname == "cempMemPoolLargestFreeOvrflw" { return "Cempmempoollargestfreeovrflw" }
    if yname == "cempMemPoolHCLargestFree" { return "Cempmempoolhclargestfree" }
    if yname == "cempMemPoolLowestFreeOvrflw" { return "Cempmempoollowestfreeovrflw" }
    if yname == "cempMemPoolHCLowestFree" { return "Cempmempoolhclowestfree" }
    if yname == "cempMemPoolUsedLowWaterMarkOvrflw" { return "Cempmempoolusedlowwatermarkovrflw" }
    if yname == "cempMemPoolHCUsedLowWaterMark" { return "Cempmempoolhcusedlowwatermark" }
    if yname == "cempMemPoolSharedOvrflw" { return "Cempmempoolsharedovrflw" }
    if yname == "cempMemPoolHCShared" { return "Cempmempoolhcshared" }
    return ""
}

func (cempmempoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable_Cempmempoolentry) GetSegmentPath() string {
    return "cempMemPoolEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cempmempoolentry.Entphysicalindex) + "']" + "[cempMemPoolIndex='" + fmt.Sprintf("%v", cempmempoolentry.Cempmempoolindex) + "']"
}

func (cempmempoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable_Cempmempoolentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cempmempoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable_Cempmempoolentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cempmempoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable_Cempmempoolentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cempmempoolentry.Entphysicalindex
    leafs["cempMemPoolIndex"] = cempmempoolentry.Cempmempoolindex
    leafs["cempMemPoolType"] = cempmempoolentry.Cempmempooltype
    leafs["cempMemPoolName"] = cempmempoolentry.Cempmempoolname
    leafs["cempMemPoolPlatformMemory"] = cempmempoolentry.Cempmempoolplatformmemory
    leafs["cempMemPoolAlternate"] = cempmempoolentry.Cempmempoolalternate
    leafs["cempMemPoolValid"] = cempmempoolentry.Cempmempoolvalid
    leafs["cempMemPoolUsed"] = cempmempoolentry.Cempmempoolused
    leafs["cempMemPoolFree"] = cempmempoolentry.Cempmempoolfree
    leafs["cempMemPoolLargestFree"] = cempmempoolentry.Cempmempoollargestfree
    leafs["cempMemPoolLowestFree"] = cempmempoolentry.Cempmempoollowestfree
    leafs["cempMemPoolUsedLowWaterMark"] = cempmempoolentry.Cempmempoolusedlowwatermark
    leafs["cempMemPoolAllocHit"] = cempmempoolentry.Cempmempoolallochit
    leafs["cempMemPoolAllocMiss"] = cempmempoolentry.Cempmempoolallocmiss
    leafs["cempMemPoolFreeHit"] = cempmempoolentry.Cempmempoolfreehit
    leafs["cempMemPoolFreeMiss"] = cempmempoolentry.Cempmempoolfreemiss
    leafs["cempMemPoolShared"] = cempmempoolentry.Cempmempoolshared
    leafs["cempMemPoolUsedOvrflw"] = cempmempoolentry.Cempmempoolusedovrflw
    leafs["cempMemPoolHCUsed"] = cempmempoolentry.Cempmempoolhcused
    leafs["cempMemPoolFreeOvrflw"] = cempmempoolentry.Cempmempoolfreeovrflw
    leafs["cempMemPoolHCFree"] = cempmempoolentry.Cempmempoolhcfree
    leafs["cempMemPoolLargestFreeOvrflw"] = cempmempoolentry.Cempmempoollargestfreeovrflw
    leafs["cempMemPoolHCLargestFree"] = cempmempoolentry.Cempmempoolhclargestfree
    leafs["cempMemPoolLowestFreeOvrflw"] = cempmempoolentry.Cempmempoollowestfreeovrflw
    leafs["cempMemPoolHCLowestFree"] = cempmempoolentry.Cempmempoolhclowestfree
    leafs["cempMemPoolUsedLowWaterMarkOvrflw"] = cempmempoolentry.Cempmempoolusedlowwatermarkovrflw
    leafs["cempMemPoolHCUsedLowWaterMark"] = cempmempoolentry.Cempmempoolhcusedlowwatermark
    leafs["cempMemPoolSharedOvrflw"] = cempmempoolentry.Cempmempoolsharedovrflw
    leafs["cempMemPoolHCShared"] = cempmempoolentry.Cempmempoolhcshared
    return leafs
}

func (cempmempoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable_Cempmempoolentry) GetBundleName() string { return "cisco_ios_xe" }

func (cempmempoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable_Cempmempoolentry) GetYangName() string { return "cempMemPoolEntry" }

func (cempmempoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable_Cempmempoolentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cempmempoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable_Cempmempoolentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cempmempoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable_Cempmempoolentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cempmempoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable_Cempmempoolentry) SetParent(parent types.Entity) { cempmempoolentry.parent = parent }

func (cempmempoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable_Cempmempoolentry) GetParent() types.Entity { return cempmempoolentry.parent }

func (cempmempoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmempooltable_Cempmempoolentry) GetParentYangName() string { return "cempMemPoolTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // This contains all the memory buffer pool configurations object values. The 
    // entPhysicalIndex identifies the entity on which memory buffer pools are
    // present. The type is slice of
    // CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable_Cempmembufferpoolentry.
    Cempmembufferpoolentry []CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable_Cempmembufferpoolentry
}

func (cempmembufferpooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable) GetFilter() yfilter.YFilter { return cempmembufferpooltable.YFilter }

func (cempmembufferpooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable) SetFilter(yf yfilter.YFilter) { cempmembufferpooltable.YFilter = yf }

func (cempmembufferpooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable) GetGoName(yname string) string {
    if yname == "cempMemBufferPoolEntry" { return "Cempmembufferpoolentry" }
    return ""
}

func (cempmembufferpooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable) GetSegmentPath() string {
    return "cempMemBufferPoolTable"
}

func (cempmembufferpooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cempMemBufferPoolEntry" {
        for _, c := range cempmembufferpooltable.Cempmembufferpoolentry {
            if cempmembufferpooltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable_Cempmembufferpoolentry{}
        cempmembufferpooltable.Cempmembufferpoolentry = append(cempmembufferpooltable.Cempmembufferpoolentry, child)
        return &cempmembufferpooltable.Cempmembufferpoolentry[len(cempmembufferpooltable.Cempmembufferpoolentry)-1]
    }
    return nil
}

func (cempmembufferpooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cempmembufferpooltable.Cempmembufferpoolentry {
        children[cempmembufferpooltable.Cempmembufferpoolentry[i].GetSegmentPath()] = &cempmembufferpooltable.Cempmembufferpoolentry[i]
    }
    return children
}

func (cempmembufferpooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cempmembufferpooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable) GetBundleName() string { return "cisco_ios_xe" }

func (cempmembufferpooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable) GetYangName() string { return "cempMemBufferPoolTable" }

func (cempmembufferpooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cempmembufferpooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cempmembufferpooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cempmembufferpooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable) SetParent(parent types.Entity) { cempmembufferpooltable.parent = parent }

func (cempmembufferpooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable) GetParent() types.Entity { return cempmembufferpooltable.parent }

func (cempmembufferpooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable) GetParentYangName() string { return "CISCO-ENHANCED-MEMPOOL-MIB" }

// CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable_Cempmembufferpoolentry
// This contains all the memory buffer pool
// configurations object values. The 
// entPhysicalIndex identifies the entity on which
// memory buffer pools are present.
type CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable_Cempmembufferpoolentry struct {
    parent types.Entity
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

func (cempmembufferpoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable_Cempmembufferpoolentry) GetFilter() yfilter.YFilter { return cempmembufferpoolentry.YFilter }

func (cempmembufferpoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable_Cempmembufferpoolentry) SetFilter(yf yfilter.YFilter) { cempmembufferpoolentry.YFilter = yf }

func (cempmembufferpoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable_Cempmembufferpoolentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cempMemBufferPoolIndex" { return "Cempmembufferpoolindex" }
    if yname == "cempMemBufferMemPoolIndex" { return "Cempmembuffermempoolindex" }
    if yname == "cempMemBufferName" { return "Cempmembuffername" }
    if yname == "cempMemBufferDynamic" { return "Cempmembufferdynamic" }
    if yname == "cempMemBufferSize" { return "Cempmembuffersize" }
    if yname == "cempMemBufferMin" { return "Cempmembuffermin" }
    if yname == "cempMemBufferMax" { return "Cempmembuffermax" }
    if yname == "cempMemBufferPermanent" { return "Cempmembufferpermanent" }
    if yname == "cempMemBufferTransient" { return "Cempmembuffertransient" }
    if yname == "cempMemBufferTotal" { return "Cempmembuffertotal" }
    if yname == "cempMemBufferFree" { return "Cempmembufferfree" }
    if yname == "cempMemBufferHit" { return "Cempmembufferhit" }
    if yname == "cempMemBufferMiss" { return "Cempmembuffermiss" }
    if yname == "cempMemBufferFreeHit" { return "Cempmembufferfreehit" }
    if yname == "cempMemBufferFreeMiss" { return "Cempmembufferfreemiss" }
    if yname == "cempMemBufferPermChange" { return "Cempmembufferpermchange" }
    if yname == "cempMemBufferPeak" { return "Cempmembufferpeak" }
    if yname == "cempMemBufferPeakTime" { return "Cempmembufferpeaktime" }
    if yname == "cempMemBufferTrim" { return "Cempmembuffertrim" }
    if yname == "cempMemBufferGrow" { return "Cempmembuffergrow" }
    if yname == "cempMemBufferFailures" { return "Cempmembufferfailures" }
    if yname == "cempMemBufferNoStorage" { return "Cempmembuffernostorage" }
    return ""
}

func (cempmembufferpoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable_Cempmembufferpoolentry) GetSegmentPath() string {
    return "cempMemBufferPoolEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cempmembufferpoolentry.Entphysicalindex) + "']" + "[cempMemBufferPoolIndex='" + fmt.Sprintf("%v", cempmembufferpoolentry.Cempmembufferpoolindex) + "']"
}

func (cempmembufferpoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable_Cempmembufferpoolentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cempmembufferpoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable_Cempmembufferpoolentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cempmembufferpoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable_Cempmembufferpoolentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cempmembufferpoolentry.Entphysicalindex
    leafs["cempMemBufferPoolIndex"] = cempmembufferpoolentry.Cempmembufferpoolindex
    leafs["cempMemBufferMemPoolIndex"] = cempmembufferpoolentry.Cempmembuffermempoolindex
    leafs["cempMemBufferName"] = cempmembufferpoolentry.Cempmembuffername
    leafs["cempMemBufferDynamic"] = cempmembufferpoolentry.Cempmembufferdynamic
    leafs["cempMemBufferSize"] = cempmembufferpoolentry.Cempmembuffersize
    leafs["cempMemBufferMin"] = cempmembufferpoolentry.Cempmembuffermin
    leafs["cempMemBufferMax"] = cempmembufferpoolentry.Cempmembuffermax
    leafs["cempMemBufferPermanent"] = cempmembufferpoolentry.Cempmembufferpermanent
    leafs["cempMemBufferTransient"] = cempmembufferpoolentry.Cempmembuffertransient
    leafs["cempMemBufferTotal"] = cempmembufferpoolentry.Cempmembuffertotal
    leafs["cempMemBufferFree"] = cempmembufferpoolentry.Cempmembufferfree
    leafs["cempMemBufferHit"] = cempmembufferpoolentry.Cempmembufferhit
    leafs["cempMemBufferMiss"] = cempmembufferpoolentry.Cempmembuffermiss
    leafs["cempMemBufferFreeHit"] = cempmembufferpoolentry.Cempmembufferfreehit
    leafs["cempMemBufferFreeMiss"] = cempmembufferpoolentry.Cempmembufferfreemiss
    leafs["cempMemBufferPermChange"] = cempmembufferpoolentry.Cempmembufferpermchange
    leafs["cempMemBufferPeak"] = cempmembufferpoolentry.Cempmembufferpeak
    leafs["cempMemBufferPeakTime"] = cempmembufferpoolentry.Cempmembufferpeaktime
    leafs["cempMemBufferTrim"] = cempmembufferpoolentry.Cempmembuffertrim
    leafs["cempMemBufferGrow"] = cempmembufferpoolentry.Cempmembuffergrow
    leafs["cempMemBufferFailures"] = cempmembufferpoolentry.Cempmembufferfailures
    leafs["cempMemBufferNoStorage"] = cempmembufferpoolentry.Cempmembuffernostorage
    return leafs
}

func (cempmembufferpoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable_Cempmembufferpoolentry) GetBundleName() string { return "cisco_ios_xe" }

func (cempmembufferpoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable_Cempmembufferpoolentry) GetYangName() string { return "cempMemBufferPoolEntry" }

func (cempmembufferpoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable_Cempmembufferpoolentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cempmembufferpoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable_Cempmembufferpoolentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cempmembufferpoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable_Cempmembufferpoolentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cempmembufferpoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable_Cempmembufferpoolentry) SetParent(parent types.Entity) { cempmembufferpoolentry.parent = parent }

func (cempmembufferpoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable_Cempmembufferpoolentry) GetParent() types.Entity { return cempmembufferpoolentry.parent }

func (cempmembufferpoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembufferpooltable_Cempmembufferpoolentry) GetParentYangName() string { return "cempMemBufferPoolTable" }

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
    parent types.Entity
    YFilter yfilter.YFilter

    // Each entry represents one of the cache buffer pools available in the system
    // and it contains the parameters configured for it. Note :
    // cempMemBufferCachePoolTable has a sparse dependency with
    // cempMemBufferPoolTable (i.e all the entires in cempMemBufferPoolTable need
    // not have an entry in cempMemBufferCachePoolTable. The type is slice of
    // CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable_Cempmembuffercachepoolentry.
    Cempmembuffercachepoolentry []CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable_Cempmembuffercachepoolentry
}

func (cempmembuffercachepooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable) GetFilter() yfilter.YFilter { return cempmembuffercachepooltable.YFilter }

func (cempmembuffercachepooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable) SetFilter(yf yfilter.YFilter) { cempmembuffercachepooltable.YFilter = yf }

func (cempmembuffercachepooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable) GetGoName(yname string) string {
    if yname == "cempMemBufferCachePoolEntry" { return "Cempmembuffercachepoolentry" }
    return ""
}

func (cempmembuffercachepooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable) GetSegmentPath() string {
    return "cempMemBufferCachePoolTable"
}

func (cempmembuffercachepooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cempMemBufferCachePoolEntry" {
        for _, c := range cempmembuffercachepooltable.Cempmembuffercachepoolentry {
            if cempmembuffercachepooltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable_Cempmembuffercachepoolentry{}
        cempmembuffercachepooltable.Cempmembuffercachepoolentry = append(cempmembuffercachepooltable.Cempmembuffercachepoolentry, child)
        return &cempmembuffercachepooltable.Cempmembuffercachepoolentry[len(cempmembuffercachepooltable.Cempmembuffercachepoolentry)-1]
    }
    return nil
}

func (cempmembuffercachepooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cempmembuffercachepooltable.Cempmembuffercachepoolentry {
        children[cempmembuffercachepooltable.Cempmembuffercachepoolentry[i].GetSegmentPath()] = &cempmembuffercachepooltable.Cempmembuffercachepoolentry[i]
    }
    return children
}

func (cempmembuffercachepooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cempmembuffercachepooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable) GetBundleName() string { return "cisco_ios_xe" }

func (cempmembuffercachepooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable) GetYangName() string { return "cempMemBufferCachePoolTable" }

func (cempmembuffercachepooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cempmembuffercachepooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cempmembuffercachepooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cempmembuffercachepooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable) SetParent(parent types.Entity) { cempmembuffercachepooltable.parent = parent }

func (cempmembuffercachepooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable) GetParent() types.Entity { return cempmembuffercachepooltable.parent }

func (cempmembuffercachepooltable *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable) GetParentYangName() string { return "CISCO-ENHANCED-MEMPOOL-MIB" }

// CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable_Cempmembuffercachepoolentry
// Each entry represents one of the cache buffer pools
// available in the system and it contains the
// parameters configured for it.
// Note : cempMemBufferCachePoolTable has a sparse
// dependency with cempMemBufferPoolTable (i.e all the
// entires in cempMemBufferPoolTable need not have an
// entry in cempMemBufferCachePoolTable.
type CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable_Cempmembuffercachepoolentry struct {
    parent types.Entity
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

func (cempmembuffercachepoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable_Cempmembuffercachepoolentry) GetFilter() yfilter.YFilter { return cempmembuffercachepoolentry.YFilter }

func (cempmembuffercachepoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable_Cempmembuffercachepoolentry) SetFilter(yf yfilter.YFilter) { cempmembuffercachepoolentry.YFilter = yf }

func (cempmembuffercachepoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable_Cempmembuffercachepoolentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cempMemBufferPoolIndex" { return "Cempmembufferpoolindex" }
    if yname == "cempMemBufferCacheSize" { return "Cempmembuffercachesize" }
    if yname == "cempMemBufferCacheTotal" { return "Cempmembuffercachetotal" }
    if yname == "cempMemBufferCacheUsed" { return "Cempmembuffercacheused" }
    if yname == "cempMemBufferCacheHit" { return "Cempmembuffercachehit" }
    if yname == "cempMemBufferCacheMiss" { return "Cempmembuffercachemiss" }
    if yname == "cempMemBufferCacheThreshold" { return "Cempmembuffercachethreshold" }
    if yname == "cempMemBufferCacheThresholdCount" { return "Cempmembuffercachethresholdcount" }
    return ""
}

func (cempmembuffercachepoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable_Cempmembuffercachepoolentry) GetSegmentPath() string {
    return "cempMemBufferCachePoolEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cempmembuffercachepoolentry.Entphysicalindex) + "']" + "[cempMemBufferPoolIndex='" + fmt.Sprintf("%v", cempmembuffercachepoolentry.Cempmembufferpoolindex) + "']"
}

func (cempmembuffercachepoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable_Cempmembuffercachepoolentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cempmembuffercachepoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable_Cempmembuffercachepoolentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cempmembuffercachepoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable_Cempmembuffercachepoolentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cempmembuffercachepoolentry.Entphysicalindex
    leafs["cempMemBufferPoolIndex"] = cempmembuffercachepoolentry.Cempmembufferpoolindex
    leafs["cempMemBufferCacheSize"] = cempmembuffercachepoolentry.Cempmembuffercachesize
    leafs["cempMemBufferCacheTotal"] = cempmembuffercachepoolentry.Cempmembuffercachetotal
    leafs["cempMemBufferCacheUsed"] = cempmembuffercachepoolentry.Cempmembuffercacheused
    leafs["cempMemBufferCacheHit"] = cempmembuffercachepoolentry.Cempmembuffercachehit
    leafs["cempMemBufferCacheMiss"] = cempmembuffercachepoolentry.Cempmembuffercachemiss
    leafs["cempMemBufferCacheThreshold"] = cempmembuffercachepoolentry.Cempmembuffercachethreshold
    leafs["cempMemBufferCacheThresholdCount"] = cempmembuffercachepoolentry.Cempmembuffercachethresholdcount
    return leafs
}

func (cempmembuffercachepoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable_Cempmembuffercachepoolentry) GetBundleName() string { return "cisco_ios_xe" }

func (cempmembuffercachepoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable_Cempmembuffercachepoolentry) GetYangName() string { return "cempMemBufferCachePoolEntry" }

func (cempmembuffercachepoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable_Cempmembuffercachepoolentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cempmembuffercachepoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable_Cempmembuffercachepoolentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cempmembuffercachepoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable_Cempmembuffercachepoolentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cempmembuffercachepoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable_Cempmembuffercachepoolentry) SetParent(parent types.Entity) { cempmembuffercachepoolentry.parent = parent }

func (cempmembuffercachepoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable_Cempmembuffercachepoolentry) GetParent() types.Entity { return cempmembuffercachepoolentry.parent }

func (cempmembuffercachepoolentry *CISCOENHANCEDMEMPOOLMIB_Cempmembuffercachepooltable_Cempmembuffercachepoolentry) GetParentYangName() string { return "cempMemBufferCachePoolTable" }

