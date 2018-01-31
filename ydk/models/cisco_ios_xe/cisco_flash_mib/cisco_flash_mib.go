// This MIB provides for the management of Cisco
// Flash Devices.
package cisco_flash_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_flash_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-FLASH-MIB CISCO-FLASH-MIB}", reflect.TypeOf(CISCOFLASHMIB{}))
    ydk.RegisterEntity("CISCO-FLASH-MIB:CISCO-FLASH-MIB", reflect.TypeOf(CISCOFLASHMIB{}))
}

// FlashFileType represents crashinfo      - file containing crashinfo.
type FlashFileType string

const (
    FlashFileType_unknown FlashFileType = "unknown"

    FlashFileType_config FlashFileType = "config"

    FlashFileType_image FlashFileType = "image"

    FlashFileType_directory FlashFileType = "directory"

    FlashFileType_crashinfo FlashFileType = "crashinfo"
)

// CISCOFLASHMIB
type CISCOFLASHMIB struct {
    parent types.Entity
    YFilter yfilter.YFilter

    
    Ciscoflashdevice CISCOFLASHMIB_Ciscoflashdevice

    
    Ciscoflashcfg CISCOFLASHMIB_Ciscoflashcfg

    // Table of Flash device properties for each initialized Flash device. Each
    // Flash device installed in a system is detected, sized, and initialized when
    // the system image boots up. For removable Flash devices, the device
    // properties will be dynamically deleted and recreated as the device is
    // removed and inserted. Note that in this case, the newly inserted device may
    // not be the same as the earlier removed one. The ciscoFlashDeviceInitTime
    // object is available for a management station to determine the time at which
    // a device was initialized, and thereby detect the change of a removable
    // device. A removable device that has not been installed will also have an
    // entry in this table. This is to let a management station know about a
    // removable device that has been removed. Since a removed device obviously
    // cannot be sized and initialized, the table entry for such a device will
    // have ciscoFlashDeviceSize equal to zero, and the following objects will
    // have an indeterminate value:         ciscoFlashDeviceMinPartitionSize,     
    // ciscoFlashDeviceMaxPartitions,         ciscoFlashDevicePartitions, and     
    // ciscoFlashDeviceChipCount. ciscoFlashDeviceRemovable will be true to
    // indicate it is removable.
    Ciscoflashdevicetable CISCOFLASHMIB_Ciscoflashdevicetable

    // Table of Flash device chip properties for each initialized Flash device.
    // This table is meant primarily for aiding error diagnosis.
    Ciscoflashchiptable CISCOFLASHMIB_Ciscoflashchiptable

    // Table of flash device partition properties for each initialized flash
    // partition. Whenever there is no explicit partitioning done, a single
    // partition spanning the entire device will be assumed to exist. There will
    // therefore always be atleast one partition on a device.
    Ciscoflashpartitiontable CISCOFLASHMIB_Ciscoflashpartitiontable

    // Table of information for files in a Flash partition.
    Ciscoflashfiletable CISCOFLASHMIB_Ciscoflashfiletable

    // Table of information for files on the manageable flash devices sorted by
    // File Types.
    Ciscoflashfilebytypetable CISCOFLASHMIB_Ciscoflashfilebytypetable

    // A table of Flash copy operation entries. Each entry represents a Flash copy
    // operation (to or from Flash) that has been initiated.
    Ciscoflashcopytable CISCOFLASHMIB_Ciscoflashcopytable

    // A table of Flash partitioning operation entries. Each entry represents a
    // Flash partitioning operation that has been initiated.
    Ciscoflashpartitioningtable CISCOFLASHMIB_Ciscoflashpartitioningtable

    // A table of misc Flash operation entries. Each entry represents a Flash
    // operation that has been initiated.
    Ciscoflashmiscoptable CISCOFLASHMIB_Ciscoflashmiscoptable
}

func (cISCOFLASHMIB *CISCOFLASHMIB) GetFilter() yfilter.YFilter { return cISCOFLASHMIB.YFilter }

func (cISCOFLASHMIB *CISCOFLASHMIB) SetFilter(yf yfilter.YFilter) { cISCOFLASHMIB.YFilter = yf }

func (cISCOFLASHMIB *CISCOFLASHMIB) GetGoName(yname string) string {
    if yname == "ciscoFlashDevice" { return "Ciscoflashdevice" }
    if yname == "ciscoFlashCfg" { return "Ciscoflashcfg" }
    if yname == "ciscoFlashDeviceTable" { return "Ciscoflashdevicetable" }
    if yname == "ciscoFlashChipTable" { return "Ciscoflashchiptable" }
    if yname == "ciscoFlashPartitionTable" { return "Ciscoflashpartitiontable" }
    if yname == "ciscoFlashFileTable" { return "Ciscoflashfiletable" }
    if yname == "ciscoFlashFileByTypeTable" { return "Ciscoflashfilebytypetable" }
    if yname == "ciscoFlashCopyTable" { return "Ciscoflashcopytable" }
    if yname == "ciscoFlashPartitioningTable" { return "Ciscoflashpartitioningtable" }
    if yname == "ciscoFlashMiscOpTable" { return "Ciscoflashmiscoptable" }
    return ""
}

func (cISCOFLASHMIB *CISCOFLASHMIB) GetSegmentPath() string {
    return "CISCO-FLASH-MIB:CISCO-FLASH-MIB"
}

func (cISCOFLASHMIB *CISCOFLASHMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoFlashDevice" {
        return &cISCOFLASHMIB.Ciscoflashdevice
    }
    if childYangName == "ciscoFlashCfg" {
        return &cISCOFLASHMIB.Ciscoflashcfg
    }
    if childYangName == "ciscoFlashDeviceTable" {
        return &cISCOFLASHMIB.Ciscoflashdevicetable
    }
    if childYangName == "ciscoFlashChipTable" {
        return &cISCOFLASHMIB.Ciscoflashchiptable
    }
    if childYangName == "ciscoFlashPartitionTable" {
        return &cISCOFLASHMIB.Ciscoflashpartitiontable
    }
    if childYangName == "ciscoFlashFileTable" {
        return &cISCOFLASHMIB.Ciscoflashfiletable
    }
    if childYangName == "ciscoFlashFileByTypeTable" {
        return &cISCOFLASHMIB.Ciscoflashfilebytypetable
    }
    if childYangName == "ciscoFlashCopyTable" {
        return &cISCOFLASHMIB.Ciscoflashcopytable
    }
    if childYangName == "ciscoFlashPartitioningTable" {
        return &cISCOFLASHMIB.Ciscoflashpartitioningtable
    }
    if childYangName == "ciscoFlashMiscOpTable" {
        return &cISCOFLASHMIB.Ciscoflashmiscoptable
    }
    return nil
}

func (cISCOFLASHMIB *CISCOFLASHMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ciscoFlashDevice"] = &cISCOFLASHMIB.Ciscoflashdevice
    children["ciscoFlashCfg"] = &cISCOFLASHMIB.Ciscoflashcfg
    children["ciscoFlashDeviceTable"] = &cISCOFLASHMIB.Ciscoflashdevicetable
    children["ciscoFlashChipTable"] = &cISCOFLASHMIB.Ciscoflashchiptable
    children["ciscoFlashPartitionTable"] = &cISCOFLASHMIB.Ciscoflashpartitiontable
    children["ciscoFlashFileTable"] = &cISCOFLASHMIB.Ciscoflashfiletable
    children["ciscoFlashFileByTypeTable"] = &cISCOFLASHMIB.Ciscoflashfilebytypetable
    children["ciscoFlashCopyTable"] = &cISCOFLASHMIB.Ciscoflashcopytable
    children["ciscoFlashPartitioningTable"] = &cISCOFLASHMIB.Ciscoflashpartitioningtable
    children["ciscoFlashMiscOpTable"] = &cISCOFLASHMIB.Ciscoflashmiscoptable
    return children
}

func (cISCOFLASHMIB *CISCOFLASHMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOFLASHMIB *CISCOFLASHMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOFLASHMIB *CISCOFLASHMIB) GetYangName() string { return "CISCO-FLASH-MIB" }

func (cISCOFLASHMIB *CISCOFLASHMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOFLASHMIB *CISCOFLASHMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOFLASHMIB *CISCOFLASHMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOFLASHMIB *CISCOFLASHMIB) SetParent(parent types.Entity) { cISCOFLASHMIB.parent = parent }

func (cISCOFLASHMIB *CISCOFLASHMIB) GetParent() types.Entity { return cISCOFLASHMIB.parent }

func (cISCOFLASHMIB *CISCOFLASHMIB) GetParentYangName() string { return "CISCO-FLASH-MIB" }

// CISCOFLASHMIB_Ciscoflashdevice
type CISCOFLASHMIB_Ciscoflashdevice struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Number of Flash devices supported by the system. If the system does not
    // support any Flash devices, this MIB will not be loaded on that system. The
    // value of this object will therefore be atleast 1. The type is interface{}
    // with range: 0..4294967295.
    Ciscoflashdevicessupported interface{}
}

func (ciscoflashdevice *CISCOFLASHMIB_Ciscoflashdevice) GetFilter() yfilter.YFilter { return ciscoflashdevice.YFilter }

func (ciscoflashdevice *CISCOFLASHMIB_Ciscoflashdevice) SetFilter(yf yfilter.YFilter) { ciscoflashdevice.YFilter = yf }

func (ciscoflashdevice *CISCOFLASHMIB_Ciscoflashdevice) GetGoName(yname string) string {
    if yname == "ciscoFlashDevicesSupported" { return "Ciscoflashdevicessupported" }
    return ""
}

func (ciscoflashdevice *CISCOFLASHMIB_Ciscoflashdevice) GetSegmentPath() string {
    return "ciscoFlashDevice"
}

func (ciscoflashdevice *CISCOFLASHMIB_Ciscoflashdevice) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoflashdevice *CISCOFLASHMIB_Ciscoflashdevice) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoflashdevice *CISCOFLASHMIB_Ciscoflashdevice) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ciscoFlashDevicesSupported"] = ciscoflashdevice.Ciscoflashdevicessupported
    return leafs
}

func (ciscoflashdevice *CISCOFLASHMIB_Ciscoflashdevice) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoflashdevice *CISCOFLASHMIB_Ciscoflashdevice) GetYangName() string { return "ciscoFlashDevice" }

func (ciscoflashdevice *CISCOFLASHMIB_Ciscoflashdevice) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoflashdevice *CISCOFLASHMIB_Ciscoflashdevice) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoflashdevice *CISCOFLASHMIB_Ciscoflashdevice) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoflashdevice *CISCOFLASHMIB_Ciscoflashdevice) SetParent(parent types.Entity) { ciscoflashdevice.parent = parent }

func (ciscoflashdevice *CISCOFLASHMIB_Ciscoflashdevice) GetParent() types.Entity { return ciscoflashdevice.parent }

func (ciscoflashdevice *CISCOFLASHMIB_Ciscoflashdevice) GetParentYangName() string { return "CISCO-FLASH-MIB" }

// CISCOFLASHMIB_Ciscoflashcfg
type CISCOFLASHMIB_Ciscoflashcfg struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specifies whether or not a notification should be generated on the
    // insertion of a Flash device.  If the value of this object is 'true' then
    // the ciscoFlashDeviceInsertedNotif notification will be generated.  If the
    // value of this object is 'false' then the ciscoFlashDeviceInsertedNotif
    // notification will not be generated.  It is the responsibility of the
    // management entity to ensure that the SNMP administrative model is
    // configured in such a way as to allow the notification to be delivered. The
    // type is bool.
    Ciscoflashcfgdevinsnotifenable interface{}

    // Specifies whether or not a notification should be generated on the removal
    // of a Flash device.  If the value of this object is 'true' then the
    // ciscoFlashDeviceRemovedNotif notification will be generated.  If the value
    // of this object is 'false' then the ciscoFlashDeviceRemovedNotif
    // notification will not be generated.  It is the responsibility of the
    // management entity to ensure that the SNMP administrative model is
    // configured in such a way as to allow the notification to be delivered. The
    // type is bool.
    Ciscoflashcfgdevremnotifenable interface{}

    // This object specifies whether or not a notification should be generated
    // when the free space falls below the threshold value on a flash partition
    // and on recovery from low space.  If the value of this object is 'true' then
    // ciscoFlashPartitionLowSpaceNotif and
    // ciscoFlashPartitionLowSpaceRecoveryNotif notifications will be generated. 
    // If the value of this object is 'false' then the
    // ciscoFlashPartitionLowSpaceNotif  and
    // ciscoFlashPartitionLowSpaceRecoveryNotif notifications will not be
    // generated.  It is the responsibility of the management entity to ensure
    // that the SNMP administrative model is configured in such a way as to allow
    // the notifications to be delivered. The type is bool.
    Ciscoflashpartitionlowspacenotifenable interface{}
}

func (ciscoflashcfg *CISCOFLASHMIB_Ciscoflashcfg) GetFilter() yfilter.YFilter { return ciscoflashcfg.YFilter }

func (ciscoflashcfg *CISCOFLASHMIB_Ciscoflashcfg) SetFilter(yf yfilter.YFilter) { ciscoflashcfg.YFilter = yf }

func (ciscoflashcfg *CISCOFLASHMIB_Ciscoflashcfg) GetGoName(yname string) string {
    if yname == "ciscoFlashCfgDevInsNotifEnable" { return "Ciscoflashcfgdevinsnotifenable" }
    if yname == "ciscoFlashCfgDevRemNotifEnable" { return "Ciscoflashcfgdevremnotifenable" }
    if yname == "ciscoFlashPartitionLowSpaceNotifEnable" { return "Ciscoflashpartitionlowspacenotifenable" }
    return ""
}

func (ciscoflashcfg *CISCOFLASHMIB_Ciscoflashcfg) GetSegmentPath() string {
    return "ciscoFlashCfg"
}

func (ciscoflashcfg *CISCOFLASHMIB_Ciscoflashcfg) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoflashcfg *CISCOFLASHMIB_Ciscoflashcfg) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoflashcfg *CISCOFLASHMIB_Ciscoflashcfg) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ciscoFlashCfgDevInsNotifEnable"] = ciscoflashcfg.Ciscoflashcfgdevinsnotifenable
    leafs["ciscoFlashCfgDevRemNotifEnable"] = ciscoflashcfg.Ciscoflashcfgdevremnotifenable
    leafs["ciscoFlashPartitionLowSpaceNotifEnable"] = ciscoflashcfg.Ciscoflashpartitionlowspacenotifenable
    return leafs
}

func (ciscoflashcfg *CISCOFLASHMIB_Ciscoflashcfg) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoflashcfg *CISCOFLASHMIB_Ciscoflashcfg) GetYangName() string { return "ciscoFlashCfg" }

func (ciscoflashcfg *CISCOFLASHMIB_Ciscoflashcfg) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoflashcfg *CISCOFLASHMIB_Ciscoflashcfg) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoflashcfg *CISCOFLASHMIB_Ciscoflashcfg) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoflashcfg *CISCOFLASHMIB_Ciscoflashcfg) SetParent(parent types.Entity) { ciscoflashcfg.parent = parent }

func (ciscoflashcfg *CISCOFLASHMIB_Ciscoflashcfg) GetParent() types.Entity { return ciscoflashcfg.parent }

func (ciscoflashcfg *CISCOFLASHMIB_Ciscoflashcfg) GetParentYangName() string { return "CISCO-FLASH-MIB" }

// CISCOFLASHMIB_Ciscoflashdevicetable
// Table of Flash device properties for each initialized
// Flash device. Each Flash device installed in a system
// is detected, sized, and initialized when the system
// image boots up.
// For removable Flash devices, the device properties
// will be dynamically deleted and recreated as the
// device is removed and inserted. Note that in this
// case, the newly inserted device may not be the same as
// the earlier removed one. The ciscoFlashDeviceInitTime
// object is available for a management station to determine
// the time at which a device was initialized, and thereby
// detect the change of a removable device.
// A removable device that has not been installed will
// also have an entry in this table. This is to let a
// management station know about a removable device that
// has been removed. Since a removed device obviously
// cannot be sized and initialized, the table entry for
// such a device will have
// ciscoFlashDeviceSize equal to zero,
// and the following objects will have
// an indeterminate value:
//         ciscoFlashDeviceMinPartitionSize,
//         ciscoFlashDeviceMaxPartitions,
//         ciscoFlashDevicePartitions, and
//         ciscoFlashDeviceChipCount.
// ciscoFlashDeviceRemovable will be
// true to indicate it is removable.
type CISCOFLASHMIB_Ciscoflashdevicetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table of flash device properties for each initialized flash
    // device. Each entry can be randomly accessed by using ciscoFlashDeviceIndex
    // as an index into the table. Note that removable devices will have an entry
    // in the table even when they have been removed. However, a non-removable
    // device that has not been installed will not have an entry in the table. The
    // type is slice of CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry.
    Ciscoflashdeviceentry []CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry
}

func (ciscoflashdevicetable *CISCOFLASHMIB_Ciscoflashdevicetable) GetFilter() yfilter.YFilter { return ciscoflashdevicetable.YFilter }

func (ciscoflashdevicetable *CISCOFLASHMIB_Ciscoflashdevicetable) SetFilter(yf yfilter.YFilter) { ciscoflashdevicetable.YFilter = yf }

func (ciscoflashdevicetable *CISCOFLASHMIB_Ciscoflashdevicetable) GetGoName(yname string) string {
    if yname == "ciscoFlashDeviceEntry" { return "Ciscoflashdeviceentry" }
    return ""
}

func (ciscoflashdevicetable *CISCOFLASHMIB_Ciscoflashdevicetable) GetSegmentPath() string {
    return "ciscoFlashDeviceTable"
}

func (ciscoflashdevicetable *CISCOFLASHMIB_Ciscoflashdevicetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoFlashDeviceEntry" {
        for _, c := range ciscoflashdevicetable.Ciscoflashdeviceentry {
            if ciscoflashdevicetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry{}
        ciscoflashdevicetable.Ciscoflashdeviceentry = append(ciscoflashdevicetable.Ciscoflashdeviceentry, child)
        return &ciscoflashdevicetable.Ciscoflashdeviceentry[len(ciscoflashdevicetable.Ciscoflashdeviceentry)-1]
    }
    return nil
}

func (ciscoflashdevicetable *CISCOFLASHMIB_Ciscoflashdevicetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ciscoflashdevicetable.Ciscoflashdeviceentry {
        children[ciscoflashdevicetable.Ciscoflashdeviceentry[i].GetSegmentPath()] = &ciscoflashdevicetable.Ciscoflashdeviceentry[i]
    }
    return children
}

func (ciscoflashdevicetable *CISCOFLASHMIB_Ciscoflashdevicetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciscoflashdevicetable *CISCOFLASHMIB_Ciscoflashdevicetable) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoflashdevicetable *CISCOFLASHMIB_Ciscoflashdevicetable) GetYangName() string { return "ciscoFlashDeviceTable" }

func (ciscoflashdevicetable *CISCOFLASHMIB_Ciscoflashdevicetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoflashdevicetable *CISCOFLASHMIB_Ciscoflashdevicetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoflashdevicetable *CISCOFLASHMIB_Ciscoflashdevicetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoflashdevicetable *CISCOFLASHMIB_Ciscoflashdevicetable) SetParent(parent types.Entity) { ciscoflashdevicetable.parent = parent }

func (ciscoflashdevicetable *CISCOFLASHMIB_Ciscoflashdevicetable) GetParent() types.Entity { return ciscoflashdevicetable.parent }

func (ciscoflashdevicetable *CISCOFLASHMIB_Ciscoflashdevicetable) GetParentYangName() string { return "CISCO-FLASH-MIB" }

// CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry
// An entry in the table of flash device properties for
// each initialized flash device.
// Each entry can be randomly accessed by using
// ciscoFlashDeviceIndex as an index into the table.
// Note that removable devices will have an entry in
// the table even when they have been removed. However,
// a non-removable device that has not been installed
// will not have an entry in the table.
type CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Flash device sequence number to index within the
    // table of initialized flash devices. The lowest value should be 1. The
    // highest should be less than or equal to the value of the
    // ciscoFlashDevicesSupported object. The type is interface{} with range:
    // 1..4294967295.
    Ciscoflashdeviceindex interface{}

    // Total size of the Flash device. For a removable device, the size will be
    // zero if the device has been removed.  If the total size of the flash device
    // is greater than the maximum value reportable by this object then this
    // object should report its maximum value(4,294,967,295) and
    // ciscoFlashDeviceSizeExtended must be used to report the flash device's
    // size. The type is interface{} with range: 0..4294967295. Units are bytes.
    Ciscoflashdevicesize interface{}

    // This object will give the minimum partition size supported for this device.
    // For systems that execute code directly out of Flash, the minimum partition
    // size needs to be the bank size. (Bank size is equal to the size of a chip
    // multiplied by the width of the device. In most cases, the device width is 4
    // bytes, and so the bank size would be four times the size of a chip). This
    // has to be so because all programming commands affect the operation of an
    // entire chip (in our case, an entire bank because all operations are done on
    // the entire width of the device) even though the actual command may be
    // localized to a small portion of each chip. So when executing code out of
    // Flash, one needs to be able to write and erase some portion of Flash
    // without affecting the code execution. For systems that execute code out of
    // DRAM or ROM, it is possible to partition Flash with a finer granularity
    // (for eg., at erase sector boundaries) if the system code supports such
    // granularity.  This object will let a management entity know the minimum
    // partition size as defined by the system. If the system does not support
    // partitioning, the value will be equal to the device size in
    // ciscoFlashDeviceSize. The maximum number of partitions that could be
    // configured will be equal to the minimum of ciscoFlashDeviceMaxPartitions
    // and (ciscoFlashDeviceSize / ciscoFlashDeviceMinPartitionSize).  If the
    // total size of the flash device is greater than the maximum value reportable
    // by this object then this object should report its maximum
    // value(4,294,967,295) and ciscoFlashDeviceMinPartitionSizeExtended must be
    // used to report the flash device's minimum partition size. The type is
    // interface{} with range: 0..4294967295. Units are bytes.
    Ciscoflashdeviceminpartitionsize interface{}

    // Max number of partitions supported by the system for this Flash device.
    // Default will be 1, which actually means that partitioning is not supported.
    // Note that this value will be defined by system limitations, not by the
    // flash device itself (for eg., the system may impose a limit of 2 partitions
    // even though the device may be large enough to be partitioned into 4 based
    // on the smallest partition unit supported). On systems that execute code out
    // of Flash, partitioning is a way of creating multiple file systems in the
    // Flash device so that writing into or erasing of one file system can be done
    // while executing code residing in another file system. For systems executing
    // code out of DRAM, partitioning gives a way of sub-dividing a large Flash
    // device for easier management of files. The type is interface{} with range:
    // 0..4294967295.
    Ciscoflashdevicemaxpartitions interface{}

    // Flash device partitions actually present. Number of partitions cannot
    // exceed the minimum of ciscoFlashDeviceMaxPartitions and
    // (ciscoFlashDeviceSize / ciscoFlashDeviceMinPartitionSize). Will be equal to
    // at least 1, the case where the partition spans the entire device (actually
    // no partitioning). A partition will contain one or more minimum partition
    // units (where a minimum partition unit is defined by
    // ciscoFlashDeviceMinPartitionSize). The type is interface{} with range:
    // 0..4294967295.
    Ciscoflashdevicepartitions interface{}

    // Total number of chips within the Flash device. The purpose of this object
    // is to provide information upfront to a management station on how much chip
    // info to expect and possibly help double check the chip index against an
    // upper limit when randomly retrieving chip info for a partition. The type is
    // interface{} with range: 0..64.
    Ciscoflashdevicechipcount interface{}

    // Flash device name. This name is used to refer to the device within the
    // system. Flash operations get directed to a device based on this name. The
    // system has a concept of a default device. This would be the primary or most
    // used device in case of multiple devices. The system directs an operation to
    // the default device whenever a device name is not specified. The device name
    // is therefore mandatory except when the operation is being done on the
    // default device, or, the system supports only a single Flash device. The
    // device name will always be available for a removable device, even when the
    // device has been removed. The type is string with length: 0..16.
    Ciscoflashdevicename interface{}

    // Description of a Flash device. The description is meant to explain what the
    // Flash device and its purpose is. Current values are:   System flash - for
    // the primary Flash used to store full                  system images.   Boot
    // flash   - for the secondary Flash used to store                  bootstrap
    // images. The ciscoFlashDeviceDescr, ciscoFlashDeviceController (if
    // applicable), and ciscoFlashPhyEntIndex objects are expected to collectively
    // give all information about a Flash device. The device description will
    // always be available for a removable device, even when the device has been
    // removed. The type is string with length: 0..64.
    Ciscoflashdevicedescr interface{}

    // Flash device controller. The h/w card that actually controls Flash
    // read/write/erase. Relevant for the AGS+ systems where Flash may be
    // controlled by the MC+, STR or the ENVM cards, cards that may not actually
    // contain the Flash chips. For systems that have removable PCMCIA flash cards
    // that are controlled by a PCMCIA controller chip, this object may contain a
    // description of that controller chip. Where irrelevant (Flash is a direct
    // memory mapped device accessed directly by the main processor), this object
    // will have an empty (NULL) string. The type is string with length: 0..64.
    Ciscoflashdevicecontroller interface{}

    // This object will point to an instance of a card entry in the cardTable. The
    // card entry will give details about the card on which the Flash device is
    // actually located. For most systems, this is usually the main processor
    // board. On the AGS+ systems, Flash is located on a separate multibus card
    // such as the MC. This object will therefore be used to essentially index
    // into cardTable to retrieve details about the card such as cardDescr,
    // cardSlotNumber, etc. The type is string with pattern:
    // (([0-1](\.[1-3]?[0-9]))|(2\.(0|([1-9]\d*))))(\.(0|([1-9]\d*)))*.
    Ciscoflashdevicecard interface{}

    // This object gives the state of a jumper (if present and can be determined)
    // that controls the programming voltage called Vpp to the Flash device. Vpp
    // is required for programming (erasing and writing) Flash. For certain older
    // technology chips it is also required for identifying the chips (which in
    // turn is required to identify which programming algorithms to use; different
    // chips require different algorithms and commands). The purpose of the
    // jumper, on systems where it is available, is to write protect a Flash
    // device. On most of the newer remote access routers, this jumper is
    // unavailable since users are not expected to visit remote sites just to
    // install and remove the jumpers when upgrading software in the Flash device.
    // The unknown(3) value will be returned for such systems and can be
    // interpreted to mean that a programming jumper is not present or not
    // required on those systems. On systems where the programming jumper state
    // can be read back via a hardware register, the installed(1) or
    // notInstalled(2) value will be returned. This object is expected to be used
    // in conjunction with the ciscoFlashPartitionStatus object whenever that
    // object has the readOnly(1) value. In such a case, this object will indicate
    // whether the programming jumper is a possible reason for the readOnly state.
    // The type is Ciscoflashdeviceprogrammingjumper.
    Ciscoflashdeviceprogrammingjumper interface{}

    // System time at which device was initialized. For fixed devices, this will
    // be the system time at boot up. For removable devices, it will be the time
    // at which the device was inserted, which may be boot up time, or a later
    // time (if device was inserted later). If a device (fixed or removable) was
    // repartitioned, it will be the time of repartitioning. The purpose of this
    // object is to help a management station determine if a removable device has
    // been changed. The application should retrieve this object prior to any
    // operation and compare with the previously retrieved value. Note that this
    // time will not be real time but a running time maintained by the system.
    // This running time starts from zero when the system boots up. For a
    // removable device that has been removed, this value will be zero. The type
    // is interface{} with range: 0..4294967295.
    Ciscoflashdeviceinittime interface{}

    // Whether Flash device is removable. Generally, only PCMCIA Flash cards will
    // be treated as removable. Socketed Flash chips and Flash SIMM modules will
    // not be treated as removable. Simply put, only those Flash devices that can
    // be inserted or removed without opening the hardware casing will be
    // considered removable. Further, removable Flash devices are expected to have
    // the necessary hardware support -   1. on-line removal and insertion   2.
    // interrupt generation on removal or insertion. The type is bool.
    Ciscoflashdeviceremovable interface{}

    // This object indicates the physical entity index of a physical entity in
    // entPhysicalTable which the flash device actually located. The type is
    // interface{} with range: 0..2147483647.
    Ciscoflashphyentindex interface{}

    // Extended Flash device name whose size can be upto 255 characters. This name
    // is used to refer to the device within the system. Flash operations get
    // directed to a device based on this name. The system has a concept of a
    // default device. This would be the primary or most used device in case of
    // multiple devices. The system directs an operation to the default device
    // whenever a device name is not specified. The device name is therefore
    // mandatory except when the operation is being done on the default device,
    // or, the system supports only a single Flash device. The device name will
    // always be available for a removable device, even when the device has been
    // removed. The type is string with length: 0..255.
    Ciscoflashdevicenameextended interface{}

    // Total size of the Flash device. For a removable device, the size will be
    // zero if the device has been removed.  This object is a 64-bit version of
    // ciscoFlashDeviceSize. The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    Ciscoflashdevicesizeextended interface{}

    // This object provides the minimum partition size supported for this device.
    // This object is a 64-bit version of  ciscoFlashDeviceMinPatitionSize. The
    // type is interface{} with range: 0..18446744073709551615.
    Ciscoflashdeviceminpartitionsizeextended interface{}
}

func (ciscoflashdeviceentry *CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry) GetFilter() yfilter.YFilter { return ciscoflashdeviceentry.YFilter }

func (ciscoflashdeviceentry *CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry) SetFilter(yf yfilter.YFilter) { ciscoflashdeviceentry.YFilter = yf }

func (ciscoflashdeviceentry *CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry) GetGoName(yname string) string {
    if yname == "ciscoFlashDeviceIndex" { return "Ciscoflashdeviceindex" }
    if yname == "ciscoFlashDeviceSize" { return "Ciscoflashdevicesize" }
    if yname == "ciscoFlashDeviceMinPartitionSize" { return "Ciscoflashdeviceminpartitionsize" }
    if yname == "ciscoFlashDeviceMaxPartitions" { return "Ciscoflashdevicemaxpartitions" }
    if yname == "ciscoFlashDevicePartitions" { return "Ciscoflashdevicepartitions" }
    if yname == "ciscoFlashDeviceChipCount" { return "Ciscoflashdevicechipcount" }
    if yname == "ciscoFlashDeviceName" { return "Ciscoflashdevicename" }
    if yname == "ciscoFlashDeviceDescr" { return "Ciscoflashdevicedescr" }
    if yname == "ciscoFlashDeviceController" { return "Ciscoflashdevicecontroller" }
    if yname == "ciscoFlashDeviceCard" { return "Ciscoflashdevicecard" }
    if yname == "ciscoFlashDeviceProgrammingJumper" { return "Ciscoflashdeviceprogrammingjumper" }
    if yname == "ciscoFlashDeviceInitTime" { return "Ciscoflashdeviceinittime" }
    if yname == "ciscoFlashDeviceRemovable" { return "Ciscoflashdeviceremovable" }
    if yname == "ciscoFlashPhyEntIndex" { return "Ciscoflashphyentindex" }
    if yname == "ciscoFlashDeviceNameExtended" { return "Ciscoflashdevicenameextended" }
    if yname == "ciscoFlashDeviceSizeExtended" { return "Ciscoflashdevicesizeextended" }
    if yname == "ciscoFlashDeviceMinPartitionSizeExtended" { return "Ciscoflashdeviceminpartitionsizeextended" }
    return ""
}

func (ciscoflashdeviceentry *CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry) GetSegmentPath() string {
    return "ciscoFlashDeviceEntry" + "[ciscoFlashDeviceIndex='" + fmt.Sprintf("%v", ciscoflashdeviceentry.Ciscoflashdeviceindex) + "']"
}

func (ciscoflashdeviceentry *CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoflashdeviceentry *CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoflashdeviceentry *CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ciscoFlashDeviceIndex"] = ciscoflashdeviceentry.Ciscoflashdeviceindex
    leafs["ciscoFlashDeviceSize"] = ciscoflashdeviceentry.Ciscoflashdevicesize
    leafs["ciscoFlashDeviceMinPartitionSize"] = ciscoflashdeviceentry.Ciscoflashdeviceminpartitionsize
    leafs["ciscoFlashDeviceMaxPartitions"] = ciscoflashdeviceentry.Ciscoflashdevicemaxpartitions
    leafs["ciscoFlashDevicePartitions"] = ciscoflashdeviceentry.Ciscoflashdevicepartitions
    leafs["ciscoFlashDeviceChipCount"] = ciscoflashdeviceentry.Ciscoflashdevicechipcount
    leafs["ciscoFlashDeviceName"] = ciscoflashdeviceentry.Ciscoflashdevicename
    leafs["ciscoFlashDeviceDescr"] = ciscoflashdeviceentry.Ciscoflashdevicedescr
    leafs["ciscoFlashDeviceController"] = ciscoflashdeviceentry.Ciscoflashdevicecontroller
    leafs["ciscoFlashDeviceCard"] = ciscoflashdeviceentry.Ciscoflashdevicecard
    leafs["ciscoFlashDeviceProgrammingJumper"] = ciscoflashdeviceentry.Ciscoflashdeviceprogrammingjumper
    leafs["ciscoFlashDeviceInitTime"] = ciscoflashdeviceentry.Ciscoflashdeviceinittime
    leafs["ciscoFlashDeviceRemovable"] = ciscoflashdeviceentry.Ciscoflashdeviceremovable
    leafs["ciscoFlashPhyEntIndex"] = ciscoflashdeviceentry.Ciscoflashphyentindex
    leafs["ciscoFlashDeviceNameExtended"] = ciscoflashdeviceentry.Ciscoflashdevicenameextended
    leafs["ciscoFlashDeviceSizeExtended"] = ciscoflashdeviceentry.Ciscoflashdevicesizeextended
    leafs["ciscoFlashDeviceMinPartitionSizeExtended"] = ciscoflashdeviceentry.Ciscoflashdeviceminpartitionsizeextended
    return leafs
}

func (ciscoflashdeviceentry *CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoflashdeviceentry *CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry) GetYangName() string { return "ciscoFlashDeviceEntry" }

func (ciscoflashdeviceentry *CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoflashdeviceentry *CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoflashdeviceentry *CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoflashdeviceentry *CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry) SetParent(parent types.Entity) { ciscoflashdeviceentry.parent = parent }

func (ciscoflashdeviceentry *CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry) GetParent() types.Entity { return ciscoflashdeviceentry.parent }

func (ciscoflashdeviceentry *CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry) GetParentYangName() string { return "ciscoFlashDeviceTable" }

// CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry_Ciscoflashdeviceprogrammingjumper represents for the readOnly state.
type CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry_Ciscoflashdeviceprogrammingjumper string

const (
    CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry_Ciscoflashdeviceprogrammingjumper_installed CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry_Ciscoflashdeviceprogrammingjumper = "installed"

    CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry_Ciscoflashdeviceprogrammingjumper_notInstalled CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry_Ciscoflashdeviceprogrammingjumper = "notInstalled"

    CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry_Ciscoflashdeviceprogrammingjumper_unknown CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry_Ciscoflashdeviceprogrammingjumper = "unknown"
)

// CISCOFLASHMIB_Ciscoflashchiptable
// Table of Flash device chip properties for each
// initialized Flash device.
// This table is meant primarily for aiding error
// diagnosis.
type CISCOFLASHMIB_Ciscoflashchiptable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table of chip info for each flash device initialized in the
    // system. An entry is indexed by two objects - the device index and the chip
    // index within that device. The type is slice of
    // CISCOFLASHMIB_Ciscoflashchiptable_Ciscoflashchipentry.
    Ciscoflashchipentry []CISCOFLASHMIB_Ciscoflashchiptable_Ciscoflashchipentry
}

func (ciscoflashchiptable *CISCOFLASHMIB_Ciscoflashchiptable) GetFilter() yfilter.YFilter { return ciscoflashchiptable.YFilter }

func (ciscoflashchiptable *CISCOFLASHMIB_Ciscoflashchiptable) SetFilter(yf yfilter.YFilter) { ciscoflashchiptable.YFilter = yf }

func (ciscoflashchiptable *CISCOFLASHMIB_Ciscoflashchiptable) GetGoName(yname string) string {
    if yname == "ciscoFlashChipEntry" { return "Ciscoflashchipentry" }
    return ""
}

func (ciscoflashchiptable *CISCOFLASHMIB_Ciscoflashchiptable) GetSegmentPath() string {
    return "ciscoFlashChipTable"
}

func (ciscoflashchiptable *CISCOFLASHMIB_Ciscoflashchiptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoFlashChipEntry" {
        for _, c := range ciscoflashchiptable.Ciscoflashchipentry {
            if ciscoflashchiptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOFLASHMIB_Ciscoflashchiptable_Ciscoflashchipentry{}
        ciscoflashchiptable.Ciscoflashchipentry = append(ciscoflashchiptable.Ciscoflashchipentry, child)
        return &ciscoflashchiptable.Ciscoflashchipentry[len(ciscoflashchiptable.Ciscoflashchipentry)-1]
    }
    return nil
}

func (ciscoflashchiptable *CISCOFLASHMIB_Ciscoflashchiptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ciscoflashchiptable.Ciscoflashchipentry {
        children[ciscoflashchiptable.Ciscoflashchipentry[i].GetSegmentPath()] = &ciscoflashchiptable.Ciscoflashchipentry[i]
    }
    return children
}

func (ciscoflashchiptable *CISCOFLASHMIB_Ciscoflashchiptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciscoflashchiptable *CISCOFLASHMIB_Ciscoflashchiptable) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoflashchiptable *CISCOFLASHMIB_Ciscoflashchiptable) GetYangName() string { return "ciscoFlashChipTable" }

func (ciscoflashchiptable *CISCOFLASHMIB_Ciscoflashchiptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoflashchiptable *CISCOFLASHMIB_Ciscoflashchiptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoflashchiptable *CISCOFLASHMIB_Ciscoflashchiptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoflashchiptable *CISCOFLASHMIB_Ciscoflashchiptable) SetParent(parent types.Entity) { ciscoflashchiptable.parent = parent }

func (ciscoflashchiptable *CISCOFLASHMIB_Ciscoflashchiptable) GetParent() types.Entity { return ciscoflashchiptable.parent }

func (ciscoflashchiptable *CISCOFLASHMIB_Ciscoflashchiptable) GetParentYangName() string { return "CISCO-FLASH-MIB" }

// CISCOFLASHMIB_Ciscoflashchiptable_Ciscoflashchipentry
// An entry in the table of chip info for each
// flash device initialized in the system.
// An entry is indexed by two objects - the
// device index and the chip index within that
// device.
type CISCOFLASHMIB_Ciscoflashchiptable_Ciscoflashchipentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_flash_mib.CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry_Ciscoflashdeviceindex
    Ciscoflashdeviceindex interface{}

    // This attribute is a key. Chip sequence number within selected flash device.
    // Used to index within chip info table. Value starts from 1 and should not be
    // greater than ciscoFlashDeviceChipCount for that device. When retrieving
    // chip information for chips within a partition, the sequence number should
    // lie between ciscoFlashPartitionStartChip & ciscoFlashPartitionEndChip (both
    // inclusive). The type is interface{} with range: 1..64.
    Ciscoflashchipindex interface{}

    // Manufacturer and device code for a chip. Lower byte will contain the device
    // code. Upper byte will contain the manufacturer code. If a chip code is
    // unknown because it could not be queried out of the chip, the value of this
    // object will be 00:00. Since programming algorithms differ from chip type to
    // chip type, this chip code should be used to determine which algorithms to
    // use (and thereby whether the chip is supported in the first place). The
    // type is string with length: 0..5.
    Ciscoflashchipcode interface{}

    // Flash chip name corresponding to the chip code. The name will contain the
    // manufacturer and the chip type. It will be of the form :   Intel 27F008SA.
    // In the case where a chip code is unknown, this object will be an empty
    // (NULL) string. In the case where the chip code is known but the chip is not
    // supported by the system, this object will be an empty (NULL) string. A
    // management station is therefore expected to use the chip code and the chip
    // description in conjunction to provide additional information whenever the
    // ciscoFlashPartitionStatus object has the readOnly(1) value. The type is
    // string with length: 0..32.
    Ciscoflashchipdescr interface{}

    // This object will provide a cumulative count (since last system boot up or
    // initialization) of the number of write retries that were done in the chip.
    // If no writes have been done to Flash, the count will be zero. Typically, a
    // maximum of 25 retries are done on a single location before flagging a write
    // error. A management station is expected to get this object for each chip in
    // a partition after a write failure in that partition. To keep a track of
    // retries for a given write operation, the management station would have to
    // retrieve the values for the concerned chips before and after any write
    // operation. The type is interface{} with range: 0..4294967295.
    Ciscoflashchipwriteretries interface{}

    // This object will provide a cumulative count (since last system boot up or
    // initialization) of the number of erase retries that were done in the chip.
    // Typically, a maximum of 2000 retries are done in a single erase zone (which
    // may be a full chip or a portion, depending on the chip technology) before
    // flagging an erase error. A management station is expected to get this
    // object for each chip in a partition after an erase failure in that
    // partition. To keep a track of retries for a given erase operation, the
    // management station would have to retrieve the values for the concerned
    // chips before and after any erase operation. Note that erase may be done
    // through an independent command, or through a copy-to-flash command. The
    // type is interface{} with range: 0..4294967295.
    Ciscoflashchiperaseretries interface{}

    // The maximum number of write retries done at any single location before
    // declaring a write failure. The type is interface{} with range:
    // 0..4294967295.
    Ciscoflashchipmaxwriteretries interface{}

    // The maximum number of erase retries done within an erase sector before
    // declaring an erase failure. The type is interface{} with range:
    // 0..4294967295.
    Ciscoflashchipmaxeraseretries interface{}
}

func (ciscoflashchipentry *CISCOFLASHMIB_Ciscoflashchiptable_Ciscoflashchipentry) GetFilter() yfilter.YFilter { return ciscoflashchipentry.YFilter }

func (ciscoflashchipentry *CISCOFLASHMIB_Ciscoflashchiptable_Ciscoflashchipentry) SetFilter(yf yfilter.YFilter) { ciscoflashchipentry.YFilter = yf }

func (ciscoflashchipentry *CISCOFLASHMIB_Ciscoflashchiptable_Ciscoflashchipentry) GetGoName(yname string) string {
    if yname == "ciscoFlashDeviceIndex" { return "Ciscoflashdeviceindex" }
    if yname == "ciscoFlashChipIndex" { return "Ciscoflashchipindex" }
    if yname == "ciscoFlashChipCode" { return "Ciscoflashchipcode" }
    if yname == "ciscoFlashChipDescr" { return "Ciscoflashchipdescr" }
    if yname == "ciscoFlashChipWriteRetries" { return "Ciscoflashchipwriteretries" }
    if yname == "ciscoFlashChipEraseRetries" { return "Ciscoflashchiperaseretries" }
    if yname == "ciscoFlashChipMaxWriteRetries" { return "Ciscoflashchipmaxwriteretries" }
    if yname == "ciscoFlashChipMaxEraseRetries" { return "Ciscoflashchipmaxeraseretries" }
    return ""
}

func (ciscoflashchipentry *CISCOFLASHMIB_Ciscoflashchiptable_Ciscoflashchipentry) GetSegmentPath() string {
    return "ciscoFlashChipEntry" + "[ciscoFlashDeviceIndex='" + fmt.Sprintf("%v", ciscoflashchipentry.Ciscoflashdeviceindex) + "']" + "[ciscoFlashChipIndex='" + fmt.Sprintf("%v", ciscoflashchipentry.Ciscoflashchipindex) + "']"
}

func (ciscoflashchipentry *CISCOFLASHMIB_Ciscoflashchiptable_Ciscoflashchipentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoflashchipentry *CISCOFLASHMIB_Ciscoflashchiptable_Ciscoflashchipentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoflashchipentry *CISCOFLASHMIB_Ciscoflashchiptable_Ciscoflashchipentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ciscoFlashDeviceIndex"] = ciscoflashchipentry.Ciscoflashdeviceindex
    leafs["ciscoFlashChipIndex"] = ciscoflashchipentry.Ciscoflashchipindex
    leafs["ciscoFlashChipCode"] = ciscoflashchipentry.Ciscoflashchipcode
    leafs["ciscoFlashChipDescr"] = ciscoflashchipentry.Ciscoflashchipdescr
    leafs["ciscoFlashChipWriteRetries"] = ciscoflashchipentry.Ciscoflashchipwriteretries
    leafs["ciscoFlashChipEraseRetries"] = ciscoflashchipentry.Ciscoflashchiperaseretries
    leafs["ciscoFlashChipMaxWriteRetries"] = ciscoflashchipentry.Ciscoflashchipmaxwriteretries
    leafs["ciscoFlashChipMaxEraseRetries"] = ciscoflashchipentry.Ciscoflashchipmaxeraseretries
    return leafs
}

func (ciscoflashchipentry *CISCOFLASHMIB_Ciscoflashchiptable_Ciscoflashchipentry) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoflashchipentry *CISCOFLASHMIB_Ciscoflashchiptable_Ciscoflashchipentry) GetYangName() string { return "ciscoFlashChipEntry" }

func (ciscoflashchipentry *CISCOFLASHMIB_Ciscoflashchiptable_Ciscoflashchipentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoflashchipentry *CISCOFLASHMIB_Ciscoflashchiptable_Ciscoflashchipentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoflashchipentry *CISCOFLASHMIB_Ciscoflashchiptable_Ciscoflashchipentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoflashchipentry *CISCOFLASHMIB_Ciscoflashchiptable_Ciscoflashchipentry) SetParent(parent types.Entity) { ciscoflashchipentry.parent = parent }

func (ciscoflashchipentry *CISCOFLASHMIB_Ciscoflashchiptable_Ciscoflashchipentry) GetParent() types.Entity { return ciscoflashchipentry.parent }

func (ciscoflashchipentry *CISCOFLASHMIB_Ciscoflashchiptable_Ciscoflashchipentry) GetParentYangName() string { return "ciscoFlashChipTable" }

// CISCOFLASHMIB_Ciscoflashpartitiontable
// Table of flash device partition properties for each
// initialized flash partition. Whenever there is no
// explicit partitioning done, a single partition spanning
// the entire device will be assumed to exist. There will
// therefore always be atleast one partition on a device.
type CISCOFLASHMIB_Ciscoflashpartitiontable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table of flash partition properties for each initialized
    // flash partition. Each entry will be indexed by a device number and a
    // partition number within the device. The type is slice of
    // CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry.
    Ciscoflashpartitionentry []CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry
}

func (ciscoflashpartitiontable *CISCOFLASHMIB_Ciscoflashpartitiontable) GetFilter() yfilter.YFilter { return ciscoflashpartitiontable.YFilter }

func (ciscoflashpartitiontable *CISCOFLASHMIB_Ciscoflashpartitiontable) SetFilter(yf yfilter.YFilter) { ciscoflashpartitiontable.YFilter = yf }

func (ciscoflashpartitiontable *CISCOFLASHMIB_Ciscoflashpartitiontable) GetGoName(yname string) string {
    if yname == "ciscoFlashPartitionEntry" { return "Ciscoflashpartitionentry" }
    return ""
}

func (ciscoflashpartitiontable *CISCOFLASHMIB_Ciscoflashpartitiontable) GetSegmentPath() string {
    return "ciscoFlashPartitionTable"
}

func (ciscoflashpartitiontable *CISCOFLASHMIB_Ciscoflashpartitiontable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoFlashPartitionEntry" {
        for _, c := range ciscoflashpartitiontable.Ciscoflashpartitionentry {
            if ciscoflashpartitiontable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry{}
        ciscoflashpartitiontable.Ciscoflashpartitionentry = append(ciscoflashpartitiontable.Ciscoflashpartitionentry, child)
        return &ciscoflashpartitiontable.Ciscoflashpartitionentry[len(ciscoflashpartitiontable.Ciscoflashpartitionentry)-1]
    }
    return nil
}

func (ciscoflashpartitiontable *CISCOFLASHMIB_Ciscoflashpartitiontable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ciscoflashpartitiontable.Ciscoflashpartitionentry {
        children[ciscoflashpartitiontable.Ciscoflashpartitionentry[i].GetSegmentPath()] = &ciscoflashpartitiontable.Ciscoflashpartitionentry[i]
    }
    return children
}

func (ciscoflashpartitiontable *CISCOFLASHMIB_Ciscoflashpartitiontable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciscoflashpartitiontable *CISCOFLASHMIB_Ciscoflashpartitiontable) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoflashpartitiontable *CISCOFLASHMIB_Ciscoflashpartitiontable) GetYangName() string { return "ciscoFlashPartitionTable" }

func (ciscoflashpartitiontable *CISCOFLASHMIB_Ciscoflashpartitiontable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoflashpartitiontable *CISCOFLASHMIB_Ciscoflashpartitiontable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoflashpartitiontable *CISCOFLASHMIB_Ciscoflashpartitiontable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoflashpartitiontable *CISCOFLASHMIB_Ciscoflashpartitiontable) SetParent(parent types.Entity) { ciscoflashpartitiontable.parent = parent }

func (ciscoflashpartitiontable *CISCOFLASHMIB_Ciscoflashpartitiontable) GetParent() types.Entity { return ciscoflashpartitiontable.parent }

func (ciscoflashpartitiontable *CISCOFLASHMIB_Ciscoflashpartitiontable) GetParentYangName() string { return "CISCO-FLASH-MIB" }

// CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry
// An entry in the table of flash partition properties
// for each initialized flash partition. Each entry
// will be indexed by a device number and a partition
// number within the device.
type CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_flash_mib.CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry_Ciscoflashdeviceindex
    Ciscoflashdeviceindex interface{}

    // This attribute is a key. Flash partition sequence number used to index
    // within table of initialized flash partitions. The type is interface{} with
    // range: 1..4294967295.
    Ciscoflashpartitionindex interface{}

    // Chip sequence number of first chip in partition. Used as an index into the
    // chip table. The type is interface{} with range: 1..64.
    Ciscoflashpartitionstartchip interface{}

    // Chip sequence number of last chip in partition. Used as an index into the
    // chip table. The type is interface{} with range: 1..64.
    Ciscoflashpartitionendchip interface{}

    // Flash partition size. It should be an integral multiple of
    // ciscoFlashDeviceMinPartitionSize. If there is a single partition, this size
    // will be equal to ciscoFlashDeviceSize.  If the size of the flash partition
    // is greater than the maximum value reportable by this object then this
    // object should report its maximum value(4,294,967,295) and
    // ciscoFlashPartitionSizeExtended must be used to report the flash
    // partition's size. The type is interface{} with range: 1..4294967295. Units
    // are bytes.
    Ciscoflashpartitionsize interface{}

    // Free space within a Flash partition. Note that the actual size of a file in
    // Flash includes a small overhead that represents the file system's file
    // header. Certain file systems may also have a partition or device header
    // overhead to be considered when computing the free space. Free space will be
    // computed as total partition size less size of all existing files
    // (valid/invalid/deleted files and including file header of each file), less
    // size of any partition header, less size of header of next file to be copied
    // in. In short, this object will give the size of the largest file that can
    // be copied in. The management entity will not be expected to know or use any
    // overheads such as file and partition header lengths, since such overheads
    // may vary from file system to file system. Deleted files in Flash do not
    // free up space. A partition may have to be erased in order to reclaim the
    // space occupied by files.  If the free space within a flash partition is
    // greater than the maximum value reportable by this object then this object
    // should report its maximum value(4,294,967,295) and
    // ciscoFlashPartitionFreeSpaceExtended must be used to report the flash
    // partition's free space. The type is interface{} with range: 0..4294967295.
    // Units are bytes.
    Ciscoflashpartitionfreespace interface{}

    // Count of all files in a flash partition. Both good and bad (deleted or
    // invalid checksum) files will be included in this count. The type is
    // interface{} with range: 0..4294967295.
    Ciscoflashpartitionfilecount interface{}

    // Checksum algorithm identifier for checksum method used by the file system.
    // Normally, this would be fixed for a particular file system. When a file
    // system writes a file to Flash, it checksums the data written. The checksum
    // then serves as a way to validate the data read back whenever the file is
    // opened for reading. Since there is no way, when using TFTP, to guarantee
    // that a network download has been error free (since UDP checksums may not
    // have been enabled), this object together with the ciscoFlashFileChecksum
    // object provides a method for any management station to regenerate the
    // checksum of the original file on the server and compare checksums to ensure
    // that the file download to Flash was error free. simpleChecksum represents a
    // simple 1s complement addition of short word values. Other algorithm values
    // will be added as necessary. The type is
    // Ciscoflashpartitionchecksumalgorithm.
    Ciscoflashpartitionchecksumalgorithm interface{}

    // Flash partition status can be :  * readOnly if device is not programmable
    // either because chips could not be recognized or an erroneous mismatch of
    // chips was detected. Chip recognition may fail either because the chips are
    // not supported by the system, or because the Vpp voltage required to
    // identify chips has been disabled via the programming jumper. The
    // ciscoFlashDeviceProgrammingJumper, ciscoFlashChipCode, and
    // ciscoFlashChipDescr objects can be examined to get more details on the
    // cause of this status * runFromFlash (RFF) if current image is running from
    // this partition. The ciscoFlashPartitionUpgradeMethod object will then
    // indicate whether the Flash Load Helper can be used to write a file to this
    // partition or not.  * readWrite if partition is programmable. The type is
    // Ciscoflashpartitionstatus.
    Ciscoflashpartitionstatus interface{}

    // Flash partition upgrade method, ie., method by which new files can be
    // downloaded into the partition. FLH stands for Flash Load Helper, a feature
    // provided on run-from-Flash systems for upgrading Flash. This feature uses
    // the bootstrap code in ROMs to help in automatic download. This object
    // should be retrieved if the partition status is runFromFlash(2). If the
    // partition status is readOnly(1), the upgrade method would depend on the
    // reason for the readOnly status. For eg., it may simply be a matter of
    // installing the programming jumper, or it may require execution of a later
    // version of software that supports the Flash chips.  unknown      -  the
    // current system image does not know                 how Flash can be
    // programmed. A possible                 method would be to reload the ROM
    // image                 and perform the upgrade manually. rxbootFLH    -  the
    // Flash Load Helper is available to                 download files to Flash.
    // A copy-to-flash                 command can be used and this system image  
    // will automatically reload the Rxboot image                 in ROM and
    // direct it to carry out the                 download request. direct       -
    // will be done directly by this image. The type is
    // Ciscoflashpartitionupgrademethod.
    Ciscoflashpartitionupgrademethod interface{}

    // Flash partition name used to refer to a partition by the system. This can
    // be any alpha-numeric character string of the form AAAAAAAAnn, where A
    // represents an optional alpha character and n a numeric character. Any
    // numeric characters must always form the trailing part of the string. The
    // system will strip off the alpha characters and use the numeric portion to
    // map to a partition index. Flash operations get directed to a device
    // partition based on this name. The system has a concept of a default
    // partition. This would be the first partition in the device. The system
    // directs an operation to the default partition whenever a partition name is
    // not specified. The partition name is therefore mandatory except when the
    // operation is being done on the default partition, or the device has just
    // one partition (is not partitioned). The type is string with length: 0..16.
    Ciscoflashpartitionname interface{}

    // This object indicates whether a partition requires erasure before any write
    // operations can be done in it. A management station should therefore
    // retrieve this object prior to attempting any write operation. A partition
    // requires erasure after it becomes full free space left is less than or
    // equal to the (filesystem file header size). A partition also requires
    // erasure if the system does not find the existence of any file system when
    // it boots up. The partition may be erased explicitly through the erase(5)
    // command, or by using the copyToFlashWithErase(1) command. If a
    // copyToFlashWithoutErase(2) command is issued when this object has the TRUE
    // value, the command will fail. The type is bool.
    Ciscoflashpartitionneederasure interface{}

    // Maximum file name length supported by the file system. Max file name length
    // will depend on the file system implemented. Today, all file systems support
    // a max length of at least 48 bytes. A management entity must use this object
    // when prompting a user for, or deriving the Flash file name length. The type
    // is interface{} with range: 1..256.
    Ciscoflashpartitionfilenamelength interface{}

    // Flash partition size. It should be an integral multiple of
    // ciscoFlashDeviceMinPartitionSize. If there is a single partition, this size
    // will be equal to ciscoFlashDeviceSize.  This object is a 64-bit version of
    // ciscoFlashPartitionSize. The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    Ciscoflashpartitionsizeextended interface{}

    // Free space within a Flash partition. Note that the actual size of a file in
    // Flash includes a small overhead that represents the file system's file
    // header. Certain file systems may also have a partition or device header
    // overhead to be considered when computing the free space. Free space will be
    // computed as total partition size less size of all existing files
    // (valid/invalid/deleted files and including file header of each file), less
    // size of any partition header, less size of header of next file to be copied
    // in. In short, this object will give the size of the largest file that can
    // be copied in. The management entity will not be expected to know or use any
    // overheads such as file and partition header lengths, since such overheads
    // may vary from file system to file system. Deleted files in Flash do not
    // free up space. A partition may have to be erased in order to reclaim the
    // space occupied by files.  This object is a 64-bit version of
    // ciscoFlashPartitionFreeSpace. The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    Ciscoflashpartitionfreespaceextended interface{}

    // This object specifies the minimum threshold value in percentage of free
    // space for each partition. If the free space available goes below this
    // threshold value and if ciscoFlashPartionLowSpaceNotifEnable is set to true,
    // ciscoFlashPartitionLowSpaceNotif will be generated. When the available free
    // space comes back to the threshold value
    // ciscoFlashPartionLowSpaceRecoveryNotif will be generated. The type is
    // interface{} with range: 0..100.
    Ciscoflashpartitionlowspacenotifthreshold interface{}
}

func (ciscoflashpartitionentry *CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry) GetFilter() yfilter.YFilter { return ciscoflashpartitionentry.YFilter }

func (ciscoflashpartitionentry *CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry) SetFilter(yf yfilter.YFilter) { ciscoflashpartitionentry.YFilter = yf }

func (ciscoflashpartitionentry *CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry) GetGoName(yname string) string {
    if yname == "ciscoFlashDeviceIndex" { return "Ciscoflashdeviceindex" }
    if yname == "ciscoFlashPartitionIndex" { return "Ciscoflashpartitionindex" }
    if yname == "ciscoFlashPartitionStartChip" { return "Ciscoflashpartitionstartchip" }
    if yname == "ciscoFlashPartitionEndChip" { return "Ciscoflashpartitionendchip" }
    if yname == "ciscoFlashPartitionSize" { return "Ciscoflashpartitionsize" }
    if yname == "ciscoFlashPartitionFreeSpace" { return "Ciscoflashpartitionfreespace" }
    if yname == "ciscoFlashPartitionFileCount" { return "Ciscoflashpartitionfilecount" }
    if yname == "ciscoFlashPartitionChecksumAlgorithm" { return "Ciscoflashpartitionchecksumalgorithm" }
    if yname == "ciscoFlashPartitionStatus" { return "Ciscoflashpartitionstatus" }
    if yname == "ciscoFlashPartitionUpgradeMethod" { return "Ciscoflashpartitionupgrademethod" }
    if yname == "ciscoFlashPartitionName" { return "Ciscoflashpartitionname" }
    if yname == "ciscoFlashPartitionNeedErasure" { return "Ciscoflashpartitionneederasure" }
    if yname == "ciscoFlashPartitionFileNameLength" { return "Ciscoflashpartitionfilenamelength" }
    if yname == "ciscoFlashPartitionSizeExtended" { return "Ciscoflashpartitionsizeextended" }
    if yname == "ciscoFlashPartitionFreeSpaceExtended" { return "Ciscoflashpartitionfreespaceextended" }
    if yname == "ciscoFlashPartitionLowSpaceNotifThreshold" { return "Ciscoflashpartitionlowspacenotifthreshold" }
    return ""
}

func (ciscoflashpartitionentry *CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry) GetSegmentPath() string {
    return "ciscoFlashPartitionEntry" + "[ciscoFlashDeviceIndex='" + fmt.Sprintf("%v", ciscoflashpartitionentry.Ciscoflashdeviceindex) + "']" + "[ciscoFlashPartitionIndex='" + fmt.Sprintf("%v", ciscoflashpartitionentry.Ciscoflashpartitionindex) + "']"
}

func (ciscoflashpartitionentry *CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoflashpartitionentry *CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoflashpartitionentry *CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ciscoFlashDeviceIndex"] = ciscoflashpartitionentry.Ciscoflashdeviceindex
    leafs["ciscoFlashPartitionIndex"] = ciscoflashpartitionentry.Ciscoflashpartitionindex
    leafs["ciscoFlashPartitionStartChip"] = ciscoflashpartitionentry.Ciscoflashpartitionstartchip
    leafs["ciscoFlashPartitionEndChip"] = ciscoflashpartitionentry.Ciscoflashpartitionendchip
    leafs["ciscoFlashPartitionSize"] = ciscoflashpartitionentry.Ciscoflashpartitionsize
    leafs["ciscoFlashPartitionFreeSpace"] = ciscoflashpartitionentry.Ciscoflashpartitionfreespace
    leafs["ciscoFlashPartitionFileCount"] = ciscoflashpartitionentry.Ciscoflashpartitionfilecount
    leafs["ciscoFlashPartitionChecksumAlgorithm"] = ciscoflashpartitionentry.Ciscoflashpartitionchecksumalgorithm
    leafs["ciscoFlashPartitionStatus"] = ciscoflashpartitionentry.Ciscoflashpartitionstatus
    leafs["ciscoFlashPartitionUpgradeMethod"] = ciscoflashpartitionentry.Ciscoflashpartitionupgrademethod
    leafs["ciscoFlashPartitionName"] = ciscoflashpartitionentry.Ciscoflashpartitionname
    leafs["ciscoFlashPartitionNeedErasure"] = ciscoflashpartitionentry.Ciscoflashpartitionneederasure
    leafs["ciscoFlashPartitionFileNameLength"] = ciscoflashpartitionentry.Ciscoflashpartitionfilenamelength
    leafs["ciscoFlashPartitionSizeExtended"] = ciscoflashpartitionentry.Ciscoflashpartitionsizeextended
    leafs["ciscoFlashPartitionFreeSpaceExtended"] = ciscoflashpartitionentry.Ciscoflashpartitionfreespaceextended
    leafs["ciscoFlashPartitionLowSpaceNotifThreshold"] = ciscoflashpartitionentry.Ciscoflashpartitionlowspacenotifthreshold
    return leafs
}

func (ciscoflashpartitionentry *CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoflashpartitionentry *CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry) GetYangName() string { return "ciscoFlashPartitionEntry" }

func (ciscoflashpartitionentry *CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoflashpartitionentry *CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoflashpartitionentry *CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoflashpartitionentry *CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry) SetParent(parent types.Entity) { ciscoflashpartitionentry.parent = parent }

func (ciscoflashpartitionentry *CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry) GetParent() types.Entity { return ciscoflashpartitionentry.parent }

func (ciscoflashpartitionentry *CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry) GetParentYangName() string { return "ciscoFlashPartitionTable" }

// CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry_Ciscoflashpartitionchecksumalgorithm represents values will be added as necessary.
type CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry_Ciscoflashpartitionchecksumalgorithm string

const (
    CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry_Ciscoflashpartitionchecksumalgorithm_simpleChecksum CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry_Ciscoflashpartitionchecksumalgorithm = "simpleChecksum"

    CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry_Ciscoflashpartitionchecksumalgorithm_undefined CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry_Ciscoflashpartitionchecksumalgorithm = "undefined"

    CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry_Ciscoflashpartitionchecksumalgorithm_simpleCRC CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry_Ciscoflashpartitionchecksumalgorithm = "simpleCRC"
)

// CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry_Ciscoflashpartitionstatus represents * readWrite if partition is programmable.
type CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry_Ciscoflashpartitionstatus string

const (
    CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry_Ciscoflashpartitionstatus_readOnly CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry_Ciscoflashpartitionstatus = "readOnly"

    CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry_Ciscoflashpartitionstatus_runFromFlash CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry_Ciscoflashpartitionstatus = "runFromFlash"

    CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry_Ciscoflashpartitionstatus_readWrite CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry_Ciscoflashpartitionstatus = "readWrite"
)

// CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry_Ciscoflashpartitionupgrademethod represents direct       -  will be done directly by this image.
type CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry_Ciscoflashpartitionupgrademethod string

const (
    CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry_Ciscoflashpartitionupgrademethod_unknown CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry_Ciscoflashpartitionupgrademethod = "unknown"

    CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry_Ciscoflashpartitionupgrademethod_rxbootFLH CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry_Ciscoflashpartitionupgrademethod = "rxbootFLH"

    CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry_Ciscoflashpartitionupgrademethod_direct CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry_Ciscoflashpartitionupgrademethod = "direct"
)

// CISCOFLASHMIB_Ciscoflashfiletable
// Table of information for files in a Flash partition.
type CISCOFLASHMIB_Ciscoflashfiletable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table of Flash file properties for each initialized Flash
    // partition. Each entry represents a file and gives details about the file.
    // An entry is indexed using the device number, partition number within the
    // device, and file number within the partition. The type is slice of
    // CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry.
    Ciscoflashfileentry []CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry
}

func (ciscoflashfiletable *CISCOFLASHMIB_Ciscoflashfiletable) GetFilter() yfilter.YFilter { return ciscoflashfiletable.YFilter }

func (ciscoflashfiletable *CISCOFLASHMIB_Ciscoflashfiletable) SetFilter(yf yfilter.YFilter) { ciscoflashfiletable.YFilter = yf }

func (ciscoflashfiletable *CISCOFLASHMIB_Ciscoflashfiletable) GetGoName(yname string) string {
    if yname == "ciscoFlashFileEntry" { return "Ciscoflashfileentry" }
    return ""
}

func (ciscoflashfiletable *CISCOFLASHMIB_Ciscoflashfiletable) GetSegmentPath() string {
    return "ciscoFlashFileTable"
}

func (ciscoflashfiletable *CISCOFLASHMIB_Ciscoflashfiletable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoFlashFileEntry" {
        for _, c := range ciscoflashfiletable.Ciscoflashfileentry {
            if ciscoflashfiletable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry{}
        ciscoflashfiletable.Ciscoflashfileentry = append(ciscoflashfiletable.Ciscoflashfileentry, child)
        return &ciscoflashfiletable.Ciscoflashfileentry[len(ciscoflashfiletable.Ciscoflashfileentry)-1]
    }
    return nil
}

func (ciscoflashfiletable *CISCOFLASHMIB_Ciscoflashfiletable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ciscoflashfiletable.Ciscoflashfileentry {
        children[ciscoflashfiletable.Ciscoflashfileentry[i].GetSegmentPath()] = &ciscoflashfiletable.Ciscoflashfileentry[i]
    }
    return children
}

func (ciscoflashfiletable *CISCOFLASHMIB_Ciscoflashfiletable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciscoflashfiletable *CISCOFLASHMIB_Ciscoflashfiletable) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoflashfiletable *CISCOFLASHMIB_Ciscoflashfiletable) GetYangName() string { return "ciscoFlashFileTable" }

func (ciscoflashfiletable *CISCOFLASHMIB_Ciscoflashfiletable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoflashfiletable *CISCOFLASHMIB_Ciscoflashfiletable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoflashfiletable *CISCOFLASHMIB_Ciscoflashfiletable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoflashfiletable *CISCOFLASHMIB_Ciscoflashfiletable) SetParent(parent types.Entity) { ciscoflashfiletable.parent = parent }

func (ciscoflashfiletable *CISCOFLASHMIB_Ciscoflashfiletable) GetParent() types.Entity { return ciscoflashfiletable.parent }

func (ciscoflashfiletable *CISCOFLASHMIB_Ciscoflashfiletable) GetParentYangName() string { return "CISCO-FLASH-MIB" }

// CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry
// An entry in the table of Flash file properties
// for each initialized Flash partition. Each entry
// represents a file and gives details about the file.
// An entry is indexed using the device number,
// partition number within the device, and file
// number within the partition.
type CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_flash_mib.CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry_Ciscoflashdeviceindex
    Ciscoflashdeviceindex interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_flash_mib.CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry_Ciscoflashpartitionindex
    Ciscoflashpartitionindex interface{}

    // This attribute is a key. Flash file sequence number used to index within a
    // Flash partition directory table. The type is interface{} with range:
    // 1..4294967295.
    Ciscoflashfileindex interface{}

    // Size of the file in bytes. Note that this size does not include the size of
    // the filesystem file header. File size will always be non-zero. The type is
    // interface{} with range: 0..4294967295. Units are bytes.
    Ciscoflashfilesize interface{}

    // File checksum stored in the file header. This checksum is computed and
    // stored when the file is written into Flash. It serves to validate the data
    // written into Flash. Whereas the system will generate and store the checksum
    // internally in hexadecimal form, this object will provide the checksum in a
    // string form. The checksum will be available for all valid and
    // invalid-checksum files. The type is string.
    Ciscoflashfilechecksum interface{}

    // Status of a file. A file could be explicitly deleted if the file system
    // supports such a user command facility. Alternately, an existing good file
    // would be automatically deleted if another good file with the same name were
    // copied in. Note that deleted files continue to occupy prime Flash real
    // estate.  A file is marked as having an invalid checksum if any checksum
    // mismatch was detected while writing or reading the file. Incomplete files
    // (files truncated either because of lack of free space, or a network
    // download failure) are also written with a bad checksum and marked as
    // invalid. The type is Ciscoflashfilestatus.
    Ciscoflashfilestatus interface{}

    // Flash file name as specified by the user copying in the file. The name
    // should not include the colon (:) character as it is a special separator
    // character used to delineate the device name, partition name, and the file
    // name. The type is string with length: 1..255.
    Ciscoflashfilename interface{}

    // Type of the file. The type is FlashFileType.
    Ciscoflashfiletype interface{}

    // The time at which this file was created. The type is string.
    Ciscoflashfiledate interface{}
}

func (ciscoflashfileentry *CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry) GetFilter() yfilter.YFilter { return ciscoflashfileentry.YFilter }

func (ciscoflashfileentry *CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry) SetFilter(yf yfilter.YFilter) { ciscoflashfileentry.YFilter = yf }

func (ciscoflashfileentry *CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry) GetGoName(yname string) string {
    if yname == "ciscoFlashDeviceIndex" { return "Ciscoflashdeviceindex" }
    if yname == "ciscoFlashPartitionIndex" { return "Ciscoflashpartitionindex" }
    if yname == "ciscoFlashFileIndex" { return "Ciscoflashfileindex" }
    if yname == "ciscoFlashFileSize" { return "Ciscoflashfilesize" }
    if yname == "ciscoFlashFileChecksum" { return "Ciscoflashfilechecksum" }
    if yname == "ciscoFlashFileStatus" { return "Ciscoflashfilestatus" }
    if yname == "ciscoFlashFileName" { return "Ciscoflashfilename" }
    if yname == "ciscoFlashFileType" { return "Ciscoflashfiletype" }
    if yname == "ciscoFlashFileDate" { return "Ciscoflashfiledate" }
    return ""
}

func (ciscoflashfileentry *CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry) GetSegmentPath() string {
    return "ciscoFlashFileEntry" + "[ciscoFlashDeviceIndex='" + fmt.Sprintf("%v", ciscoflashfileentry.Ciscoflashdeviceindex) + "']" + "[ciscoFlashPartitionIndex='" + fmt.Sprintf("%v", ciscoflashfileentry.Ciscoflashpartitionindex) + "']" + "[ciscoFlashFileIndex='" + fmt.Sprintf("%v", ciscoflashfileentry.Ciscoflashfileindex) + "']"
}

func (ciscoflashfileentry *CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoflashfileentry *CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoflashfileentry *CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ciscoFlashDeviceIndex"] = ciscoflashfileentry.Ciscoflashdeviceindex
    leafs["ciscoFlashPartitionIndex"] = ciscoflashfileentry.Ciscoflashpartitionindex
    leafs["ciscoFlashFileIndex"] = ciscoflashfileentry.Ciscoflashfileindex
    leafs["ciscoFlashFileSize"] = ciscoflashfileentry.Ciscoflashfilesize
    leafs["ciscoFlashFileChecksum"] = ciscoflashfileentry.Ciscoflashfilechecksum
    leafs["ciscoFlashFileStatus"] = ciscoflashfileentry.Ciscoflashfilestatus
    leafs["ciscoFlashFileName"] = ciscoflashfileentry.Ciscoflashfilename
    leafs["ciscoFlashFileType"] = ciscoflashfileentry.Ciscoflashfiletype
    leafs["ciscoFlashFileDate"] = ciscoflashfileentry.Ciscoflashfiledate
    return leafs
}

func (ciscoflashfileentry *CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoflashfileentry *CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry) GetYangName() string { return "ciscoFlashFileEntry" }

func (ciscoflashfileentry *CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoflashfileentry *CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoflashfileentry *CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoflashfileentry *CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry) SetParent(parent types.Entity) { ciscoflashfileentry.parent = parent }

func (ciscoflashfileentry *CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry) GetParent() types.Entity { return ciscoflashfileentry.parent }

func (ciscoflashfileentry *CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry) GetParentYangName() string { return "ciscoFlashFileTable" }

// CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry_Ciscoflashfilestatus represents marked as invalid.
type CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry_Ciscoflashfilestatus string

const (
    CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry_Ciscoflashfilestatus_deleted CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry_Ciscoflashfilestatus = "deleted"

    CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry_Ciscoflashfilestatus_invalidChecksum CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry_Ciscoflashfilestatus = "invalidChecksum"

    CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry_Ciscoflashfilestatus_valid CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry_Ciscoflashfilestatus = "valid"
)

// CISCOFLASHMIB_Ciscoflashfilebytypetable
// Table of information for files on the manageable
// flash devices sorted by File Types.
type CISCOFLASHMIB_Ciscoflashfilebytypetable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table of Flash file properties for each initialized Flash
    // partition. Each entry represents a file sorted by file type.  This table
    // contains exactly the same set of rows as are contained in the
    // ciscoFlashFileTable but in a different order, i.e., ordered by  the type of
    // file, given by  ciscoFlashFileType; the device number, given by
    // ciscoFlashDeviceIndex; the partition number within the device, given by
    // ciscoFlashPartitionIndex; the file number within the partition, given by
    // ciscoFlashFileIndex. The type is slice of
    // CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry.
    Ciscoflashfilebytypeentry []CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry
}

func (ciscoflashfilebytypetable *CISCOFLASHMIB_Ciscoflashfilebytypetable) GetFilter() yfilter.YFilter { return ciscoflashfilebytypetable.YFilter }

func (ciscoflashfilebytypetable *CISCOFLASHMIB_Ciscoflashfilebytypetable) SetFilter(yf yfilter.YFilter) { ciscoflashfilebytypetable.YFilter = yf }

func (ciscoflashfilebytypetable *CISCOFLASHMIB_Ciscoflashfilebytypetable) GetGoName(yname string) string {
    if yname == "ciscoFlashFileByTypeEntry" { return "Ciscoflashfilebytypeentry" }
    return ""
}

func (ciscoflashfilebytypetable *CISCOFLASHMIB_Ciscoflashfilebytypetable) GetSegmentPath() string {
    return "ciscoFlashFileByTypeTable"
}

func (ciscoflashfilebytypetable *CISCOFLASHMIB_Ciscoflashfilebytypetable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoFlashFileByTypeEntry" {
        for _, c := range ciscoflashfilebytypetable.Ciscoflashfilebytypeentry {
            if ciscoflashfilebytypetable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry{}
        ciscoflashfilebytypetable.Ciscoflashfilebytypeentry = append(ciscoflashfilebytypetable.Ciscoflashfilebytypeentry, child)
        return &ciscoflashfilebytypetable.Ciscoflashfilebytypeentry[len(ciscoflashfilebytypetable.Ciscoflashfilebytypeentry)-1]
    }
    return nil
}

func (ciscoflashfilebytypetable *CISCOFLASHMIB_Ciscoflashfilebytypetable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ciscoflashfilebytypetable.Ciscoflashfilebytypeentry {
        children[ciscoflashfilebytypetable.Ciscoflashfilebytypeentry[i].GetSegmentPath()] = &ciscoflashfilebytypetable.Ciscoflashfilebytypeentry[i]
    }
    return children
}

func (ciscoflashfilebytypetable *CISCOFLASHMIB_Ciscoflashfilebytypetable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciscoflashfilebytypetable *CISCOFLASHMIB_Ciscoflashfilebytypetable) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoflashfilebytypetable *CISCOFLASHMIB_Ciscoflashfilebytypetable) GetYangName() string { return "ciscoFlashFileByTypeTable" }

func (ciscoflashfilebytypetable *CISCOFLASHMIB_Ciscoflashfilebytypetable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoflashfilebytypetable *CISCOFLASHMIB_Ciscoflashfilebytypetable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoflashfilebytypetable *CISCOFLASHMIB_Ciscoflashfilebytypetable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoflashfilebytypetable *CISCOFLASHMIB_Ciscoflashfilebytypetable) SetParent(parent types.Entity) { ciscoflashfilebytypetable.parent = parent }

func (ciscoflashfilebytypetable *CISCOFLASHMIB_Ciscoflashfilebytypetable) GetParent() types.Entity { return ciscoflashfilebytypetable.parent }

func (ciscoflashfilebytypetable *CISCOFLASHMIB_Ciscoflashfilebytypetable) GetParentYangName() string { return "CISCO-FLASH-MIB" }

// CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry
// An entry in the table of Flash file properties
// for each initialized Flash partition. Each entry
// represents a file sorted by file type.
// 
// This table contains exactly the same set of rows
// as are contained in the ciscoFlashFileTable but
// in a different order, i.e., ordered by
// 
// the type of file, given by  ciscoFlashFileType;
// the device number, given by ciscoFlashDeviceIndex;
// the partition number within the device, given by
// ciscoFlashPartitionIndex;
// the file number within the partition, given by
// ciscoFlashFileIndex.
type CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The type is FlashFileType. Refers to
    // cisco_flash_mib.CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry_Ciscoflashfiletype
    Ciscoflashfiletype interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_flash_mib.CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry_Ciscoflashdeviceindex
    Ciscoflashdeviceindex interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_flash_mib.CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry_Ciscoflashpartitionindex
    Ciscoflashpartitionindex interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_flash_mib.CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry_Ciscoflashfileindex
    Ciscoflashfileindex interface{}

    // This object represents exactly the same info as ciscoFlashFileSize object
    // in ciscoFlashFileTable. The type is interface{} with range: 0..4294967295.
    // Units are bytes.
    Ciscoflashfilebytypesize interface{}

    // This object represents exactly the same info as ciscoFlashFileChecksum
    // object in ciscoFlashFileTable. The type is string.
    Ciscoflashfilebytypechecksum interface{}

    // This object represents exactly the same info as ciscoFlashFileStatus object
    // in ciscoFlashFileTable. The type is Ciscoflashfilebytypestatus.
    Ciscoflashfilebytypestatus interface{}

    // This object represents exactly the same info as ciscoFlashFileName object
    // in ciscoFlashFileTable. The type is string with length: 1..255.
    Ciscoflashfilebytypename interface{}

    // This object represents exactly the same info as ciscoFlashFileDate object
    // in ciscoFlashFileTable. The type is string.
    Ciscoflashfilebytypedate interface{}
}

func (ciscoflashfilebytypeentry *CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry) GetFilter() yfilter.YFilter { return ciscoflashfilebytypeentry.YFilter }

func (ciscoflashfilebytypeentry *CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry) SetFilter(yf yfilter.YFilter) { ciscoflashfilebytypeentry.YFilter = yf }

func (ciscoflashfilebytypeentry *CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry) GetGoName(yname string) string {
    if yname == "ciscoFlashFileType" { return "Ciscoflashfiletype" }
    if yname == "ciscoFlashDeviceIndex" { return "Ciscoflashdeviceindex" }
    if yname == "ciscoFlashPartitionIndex" { return "Ciscoflashpartitionindex" }
    if yname == "ciscoFlashFileIndex" { return "Ciscoflashfileindex" }
    if yname == "ciscoFlashFileByTypeSize" { return "Ciscoflashfilebytypesize" }
    if yname == "ciscoFlashFileByTypeChecksum" { return "Ciscoflashfilebytypechecksum" }
    if yname == "ciscoFlashFileByTypeStatus" { return "Ciscoflashfilebytypestatus" }
    if yname == "ciscoFlashFileByTypeName" { return "Ciscoflashfilebytypename" }
    if yname == "ciscoFlashFileByTypeDate" { return "Ciscoflashfilebytypedate" }
    return ""
}

func (ciscoflashfilebytypeentry *CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry) GetSegmentPath() string {
    return "ciscoFlashFileByTypeEntry" + "[ciscoFlashFileType='" + fmt.Sprintf("%v", ciscoflashfilebytypeentry.Ciscoflashfiletype) + "']" + "[ciscoFlashDeviceIndex='" + fmt.Sprintf("%v", ciscoflashfilebytypeentry.Ciscoflashdeviceindex) + "']" + "[ciscoFlashPartitionIndex='" + fmt.Sprintf("%v", ciscoflashfilebytypeentry.Ciscoflashpartitionindex) + "']" + "[ciscoFlashFileIndex='" + fmt.Sprintf("%v", ciscoflashfilebytypeentry.Ciscoflashfileindex) + "']"
}

func (ciscoflashfilebytypeentry *CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoflashfilebytypeentry *CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoflashfilebytypeentry *CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ciscoFlashFileType"] = ciscoflashfilebytypeentry.Ciscoflashfiletype
    leafs["ciscoFlashDeviceIndex"] = ciscoflashfilebytypeentry.Ciscoflashdeviceindex
    leafs["ciscoFlashPartitionIndex"] = ciscoflashfilebytypeentry.Ciscoflashpartitionindex
    leafs["ciscoFlashFileIndex"] = ciscoflashfilebytypeentry.Ciscoflashfileindex
    leafs["ciscoFlashFileByTypeSize"] = ciscoflashfilebytypeentry.Ciscoflashfilebytypesize
    leafs["ciscoFlashFileByTypeChecksum"] = ciscoflashfilebytypeentry.Ciscoflashfilebytypechecksum
    leafs["ciscoFlashFileByTypeStatus"] = ciscoflashfilebytypeentry.Ciscoflashfilebytypestatus
    leafs["ciscoFlashFileByTypeName"] = ciscoflashfilebytypeentry.Ciscoflashfilebytypename
    leafs["ciscoFlashFileByTypeDate"] = ciscoflashfilebytypeentry.Ciscoflashfilebytypedate
    return leafs
}

func (ciscoflashfilebytypeentry *CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoflashfilebytypeentry *CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry) GetYangName() string { return "ciscoFlashFileByTypeEntry" }

func (ciscoflashfilebytypeentry *CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoflashfilebytypeentry *CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoflashfilebytypeentry *CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoflashfilebytypeentry *CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry) SetParent(parent types.Entity) { ciscoflashfilebytypeentry.parent = parent }

func (ciscoflashfilebytypeentry *CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry) GetParent() types.Entity { return ciscoflashfilebytypeentry.parent }

func (ciscoflashfilebytypeentry *CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry) GetParentYangName() string { return "ciscoFlashFileByTypeTable" }

// CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry_Ciscoflashfilebytypestatus represents object in ciscoFlashFileTable.
type CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry_Ciscoflashfilebytypestatus string

const (
    CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry_Ciscoflashfilebytypestatus_deleted CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry_Ciscoflashfilebytypestatus = "deleted"

    CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry_Ciscoflashfilebytypestatus_invalidChecksum CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry_Ciscoflashfilebytypestatus = "invalidChecksum"

    CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry_Ciscoflashfilebytypestatus_valid CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry_Ciscoflashfilebytypestatus = "valid"
)

// CISCOFLASHMIB_Ciscoflashcopytable
// A table of Flash copy operation entries. Each
// entry represents a Flash copy operation (to or
// from Flash) that has been initiated.
type CISCOFLASHMIB_Ciscoflashcopytable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A Flash copy operation entry. Each entry consists of a command, a source,
    // and optional parameters such as protocol to be used, a destination, a
    // server address, etc.  A management station wishing to create an entry
    // should first generate a pseudo-random serial number to be used as the index
    // to this sparse table.  The station should then create the associated
    // instance of the row status object. It must also, either in the same or in
    // successive PDUs, create the associated instance of the command and
    // parameter objects. It should also modify the default values for any of the
    // parameter objects if the defaults are not appropriate.  Once the
    // appropriate instances of all the command objects have been created, either
    // by an explicit SNMP set request or by default, the row status should be set
    // to active to initiate the operation. Note that this entire procedure may be
    // initiated via a single set request which specifies a row status  of
    // createAndGo as well as specifies valid values for the non-defaulted
    // parameter objects.  Once an operation has been activated, it cannot be
    // stopped.  Once the operation completes, the management station should
    // retrieve the value of the status object (and time if desired), and delete
    // the entry.  In order to prevent old entries from clogging the table,
    // entries will be aged out, but an entry will never be deleted within 5
    // minutes of completing. The type is slice of
    // CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry.
    Ciscoflashcopyentry []CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry
}

func (ciscoflashcopytable *CISCOFLASHMIB_Ciscoflashcopytable) GetFilter() yfilter.YFilter { return ciscoflashcopytable.YFilter }

func (ciscoflashcopytable *CISCOFLASHMIB_Ciscoflashcopytable) SetFilter(yf yfilter.YFilter) { ciscoflashcopytable.YFilter = yf }

func (ciscoflashcopytable *CISCOFLASHMIB_Ciscoflashcopytable) GetGoName(yname string) string {
    if yname == "ciscoFlashCopyEntry" { return "Ciscoflashcopyentry" }
    return ""
}

func (ciscoflashcopytable *CISCOFLASHMIB_Ciscoflashcopytable) GetSegmentPath() string {
    return "ciscoFlashCopyTable"
}

func (ciscoflashcopytable *CISCOFLASHMIB_Ciscoflashcopytable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoFlashCopyEntry" {
        for _, c := range ciscoflashcopytable.Ciscoflashcopyentry {
            if ciscoflashcopytable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry{}
        ciscoflashcopytable.Ciscoflashcopyentry = append(ciscoflashcopytable.Ciscoflashcopyentry, child)
        return &ciscoflashcopytable.Ciscoflashcopyentry[len(ciscoflashcopytable.Ciscoflashcopyentry)-1]
    }
    return nil
}

func (ciscoflashcopytable *CISCOFLASHMIB_Ciscoflashcopytable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ciscoflashcopytable.Ciscoflashcopyentry {
        children[ciscoflashcopytable.Ciscoflashcopyentry[i].GetSegmentPath()] = &ciscoflashcopytable.Ciscoflashcopyentry[i]
    }
    return children
}

func (ciscoflashcopytable *CISCOFLASHMIB_Ciscoflashcopytable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciscoflashcopytable *CISCOFLASHMIB_Ciscoflashcopytable) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoflashcopytable *CISCOFLASHMIB_Ciscoflashcopytable) GetYangName() string { return "ciscoFlashCopyTable" }

func (ciscoflashcopytable *CISCOFLASHMIB_Ciscoflashcopytable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoflashcopytable *CISCOFLASHMIB_Ciscoflashcopytable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoflashcopytable *CISCOFLASHMIB_Ciscoflashcopytable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoflashcopytable *CISCOFLASHMIB_Ciscoflashcopytable) SetParent(parent types.Entity) { ciscoflashcopytable.parent = parent }

func (ciscoflashcopytable *CISCOFLASHMIB_Ciscoflashcopytable) GetParent() types.Entity { return ciscoflashcopytable.parent }

func (ciscoflashcopytable *CISCOFLASHMIB_Ciscoflashcopytable) GetParentYangName() string { return "CISCO-FLASH-MIB" }

// CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry
// A Flash copy operation entry. Each entry consists
// of a command, a source, and optional parameters such
// as protocol to be used, a destination, a server address,
// etc.
// 
// A management station wishing to create an entry should
// first generate a pseudo-random serial number to be used
// as the index to this sparse table.  The station should
// then create the associated instance of the row status
// object. It must also, either in the same or in successive
// PDUs, create the associated instance of the command and
// parameter objects. It should also modify the default values
// for any of the parameter objects if the defaults are not
// appropriate.
// 
// Once the appropriate instances of all the command
// objects have been created, either by an explicit SNMP
// set request or by default, the row status should be set
// to active to initiate the operation. Note that this entire
// procedure may be initiated via a single set request which
// specifies a row status  of createAndGo as well as specifies
// valid values for the non-defaulted parameter objects.
// 
// Once an operation has been activated, it cannot be
// stopped.
// 
// Once the operation completes, the management station should
// retrieve the value of the status object (and time if
// desired), and delete the entry.  In order to prevent old
// entries from clogging the table, entries will be aged out,
// but an entry will never be deleted within 5 minutes of
// completing.
type CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Object which specifies a unique entry in the
    // table. A management station wishing to initiate a copy operation should use
    // a pseudo-random value for this object when creating or modifying an
    // instance of a ciscoFlashCopyEntry. The type is interface{} with range:
    // 0..2147483647.
    Ciscoflashcopyserialnumber interface{}

    // The copy command to be executed. Mandatory. Note that it is possible for a
    // system to support multiple file systems (different file systems on
    // different Flash devices, or different file systems on different partitions
    // within a device). Each such file system may support only a subset of these
    // commands. If a command is unsupported, the invalidOperation(3) error will
    // be reported in the operation status.  Command                 Remarks
    // copyToFlashWithErase    Copy a file to flash; erase                        
    // flash before copy.                         Use the TFTP or rcp protocol.
    // copyToFlashWithoutErase Copy a file to flash; do not                       
    // erase.                         Note that this command will fail            
    // if the PartitionNeedErasure                         object specifies that
    // the                         partition being copied to needs                
    // erasure.                         Use the TFTP or rcp protocol.
    // copyFromFlash           Copy a file from flash using                       
    // the TFTP, rcp or lex protocol.                         Note that the lex
    // protocol                         can only be used to copy to a             
    // lex device. copyFromFlhLog          Copy contents of FLH log to            
    // server using TFTP protocol.   Command table           Parameters
    // copyToFlashWithErase    CopyProtocol                        
    // CopyServerAddress                         CopySourceName                   
    // CopyDestinationName (opt)                         CopyRemoteUserName (opt) 
    // CopyNotifyOnCompletion (opt) copyToFlashWithoutErase CopyProtocol          
    // CopyServerAddress                         CopySourceName                   
    // CopyDestinationName (opt)                         CopyRemoteUserName (opt) 
    // CopyNotifyOnCompletion (opt) copyFromFlash           CopyProtocol          
    // CopyServerAddress                         CopySourceName                   
    // CopyDestinationName (opt)                         CopyRemoteUserName (opt) 
    // CopyNotifyOnCompletion (opt) copyFromFlhLog          CopyProtocol          
    // CopyServerAddress                         CopyDestinationName              
    // CopyNotifyOnCompletion (opt). The type is Ciscoflashcopycommand.
    Ciscoflashcopycommand interface{}

    // The protocol to be used for any copy. Optional. Will default to tftp if not
    // specified.  Since feature support depends on a software release, version
    // number within the release, platform, and maybe the image type (subset
    // type), a management station would be expected to somehow determine the
    // protocol support for a command. The type is Ciscoflashcopyprotocol.
    Ciscoflashcopyprotocol interface{}

    // The server address to be used for any copy. Optional. Will default to
    // 'FFFFFFFF'H  (or 255.255.255.255).  Since this object can just hold only
    // IPv4 Transport type, it is deprecated and replaced by
    // ciscoFlashCopyServerAddrRev1. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ciscoflashcopyserveraddress interface{}

    // Source file name, either in Flash or on a server, depending on the type of
    // copy command. Mandatory.  For a copy from Flash: File name must be of the
    // form         [device>:][:] where  is a value obtained from FlashDeviceName,
    // is obtained from FlashPartitionName     and  is the name of a file in
    // Flash. A management station could derive its own partition name as per the
    // description for the ciscoFlashPartitionName object. If <device> is not
    // specified, the default Flash device will be assumed. If <partition> is not
    // specified, the default partition will be assumed. If a device is not
    // partitioned into 2 or more partitions, this value may be left out.  For a
    // copy to Flash, the file name will be as per the file naming conventions and
    // path to the file on the server. The type is string with length: 1..255.
    Ciscoflashcopysourcename interface{}

    // Destination file name.  For a copy to Flash: File name must be of the form 
    // {device>:][<partition>:]<file> where <device> is a value obtained from
    // FlashDeviceName,       <partition> is obtained from FlashPartitionName  
    // and <file> is any character string that does not have embedded colon
    // characters. A management station could derive its own partition name as per
    // the description for the ciscoFlashPartitionName object. If <device> is not
    // specified, the default Flash device will be assumed. If <partition> is not
    // specified, the default partition will be assumed. If a device is not
    // partitioned into 2 or more partitions, this value may be left out. If
    // <file> is not specified, it will default to <file> specified in
    // ciscoFlashCopySourceName.  For a copy from Flash via tftp or rcp, the file
    // name will be as per the file naming conventions and destination
    // sub-directory on the server. If not specified, <file> from the source file
    // name will be used. For a copy from Flash via lex, this string will consist
    // of numeric characters specifying the interface on the lex box that will
    // receive the source flash image. The type is string with length: 0..255.
    Ciscoflashcopydestinationname interface{}

    // Remote user name for copy via rcp protocol. Optional. This object will be
    // ignored for protocols other than rcp. If specified, it will override the
    // remote user-name configured through the         rcmd remote-username
    // configuration command. The remote user-name is sent as the server user-name
    // in an rcp command request sent by the system to a remote rcp server. The
    // type is string with length: 1..255.
    Ciscoflashcopyremoteusername interface{}

    // The status of the specified copy operation.  copyOperationPending :        
    // operation request is received and         pending for validation and
    // process  copyInProgress :         specified operation is active 
    // copyOperationSuccess :         specified operation is supported and        
    // completed successfully  copyInvalidOperation :         command invalid or
    // command-protocol-device         combination unsupported 
    // copyInvalidProtocol :         invalid protocol specified 
    // copyInvalidSourceName :         invalid source file name specified        
    // For the  copy from flash to lex operation, this         error code will be
    // returned when the source file         is not a valid lex image. 
    // copyInvalidDestName :         invalid target name (file or partition or    
    // device name) specified         For the  copy from flash to lex operation,
    // this         error code will be returned when no lex devices         are
    // connected to the router or when an invalid         lex interface number has
    // been specified in         the destination string.  copyInvalidServerAddress
    // :         invalid server address specified  copyDeviceBusy :        
    // specified device is in use and locked by         another process 
    // copyDeviceOpenError :         invalid device name  copyDeviceError :       
    // device read, write or erase error  copyDeviceNotProgrammable :        
    // device is read-only but a write or erase         operation was specified 
    // copyDeviceFull :         device is filled to capacity  copyFileOpenError : 
    // invalid file name; file not found in partition  copyFileTransferError :    
    // file transfer was unsuccessfull; network failure  copyFileChecksumError :  
    // file checksum in Flash failed  copyNoMemory :         system running low on
    // memory  copyUnknownFailure :         failure unknown  copyProhibited:      
    // stop user from overwriting current boot image file. The type is
    // Ciscoflashcopystatus.
    Ciscoflashcopystatus interface{}

    // Specifies whether or not a notification should be generated on the
    // completion of the copy operation. If specified,
    // ciscoFlashCopyCompletionTrap will be generated. It is the responsibility of
    // the management entity to ensure that the SNMP administrative model is
    // configured in such a way as to allow the notification to be delivered. The
    // type is bool.
    Ciscoflashcopynotifyoncompletion interface{}

    // Time taken for the copy operation. This object will be like a stopwatch,
    // starting when the operation starts, stopping when the operation completes.
    // If a management entity keeps a database of completion times for various
    // operations, it can then use the stopwatch capability to display percentage
    // completion time. The type is interface{} with range: 0..4294967295.
    Ciscoflashcopytime interface{}

    // The status of this table entry. The type is RowStatus.
    Ciscoflashcopyentrystatus interface{}

    // Specifies whether the file that is copied need to be verified for integrity
    // / authenticity, after copy succeeds. If it is set to true, and if the file
    // that is copied doesn't have integrity /authenticity attachement, or the
    // integrity / authenticity check fails, then the command will be aborted, and
    // the file that is copied will be deleted from the destination file system.
    // The type is bool.
    Ciscoflashcopyverify interface{}

    // This object indicates the transport type of the address contained in
    // ciscoFlashCopyServerAddrRev1. Optional. Will default to '1' (IPv4 address
    // type). The type is InetAddressType.
    Ciscoflashcopyserveraddrtype interface{}

    // The server address to be used for any copy. Optional. Will default to
    // 'FFFFFFFF'H  (or 255.255.255.255).  The Format of this address depends on
    // the value of the ciscoFlashCopyServerAddrType.  This object deprecates the
    // old ciscoFlashCopyServerAddress object. The type is string with length:
    // 0..255.
    Ciscoflashcopyserveraddrrev1 interface{}

    // Password used by ftp, sftp or scp for copying a file to/from an
    // ftp/sftp/scp server. This object must be created when the
    // ciscoFlashCopyProtocol is ftp, sftp or scp. Reading it returns a
    // zero-length string for security reasons. The type is string with length:
    // 1..40.
    Ciscoflashcopyremotepassword interface{}
}

func (ciscoflashcopyentry *CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry) GetFilter() yfilter.YFilter { return ciscoflashcopyentry.YFilter }

func (ciscoflashcopyentry *CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry) SetFilter(yf yfilter.YFilter) { ciscoflashcopyentry.YFilter = yf }

func (ciscoflashcopyentry *CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry) GetGoName(yname string) string {
    if yname == "ciscoFlashCopySerialNumber" { return "Ciscoflashcopyserialnumber" }
    if yname == "ciscoFlashCopyCommand" { return "Ciscoflashcopycommand" }
    if yname == "ciscoFlashCopyProtocol" { return "Ciscoflashcopyprotocol" }
    if yname == "ciscoFlashCopyServerAddress" { return "Ciscoflashcopyserveraddress" }
    if yname == "ciscoFlashCopySourceName" { return "Ciscoflashcopysourcename" }
    if yname == "ciscoFlashCopyDestinationName" { return "Ciscoflashcopydestinationname" }
    if yname == "ciscoFlashCopyRemoteUserName" { return "Ciscoflashcopyremoteusername" }
    if yname == "ciscoFlashCopyStatus" { return "Ciscoflashcopystatus" }
    if yname == "ciscoFlashCopyNotifyOnCompletion" { return "Ciscoflashcopynotifyoncompletion" }
    if yname == "ciscoFlashCopyTime" { return "Ciscoflashcopytime" }
    if yname == "ciscoFlashCopyEntryStatus" { return "Ciscoflashcopyentrystatus" }
    if yname == "ciscoFlashCopyVerify" { return "Ciscoflashcopyverify" }
    if yname == "ciscoFlashCopyServerAddrType" { return "Ciscoflashcopyserveraddrtype" }
    if yname == "ciscoFlashCopyServerAddrRev1" { return "Ciscoflashcopyserveraddrrev1" }
    if yname == "ciscoFlashCopyRemotePassword" { return "Ciscoflashcopyremotepassword" }
    return ""
}

func (ciscoflashcopyentry *CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry) GetSegmentPath() string {
    return "ciscoFlashCopyEntry" + "[ciscoFlashCopySerialNumber='" + fmt.Sprintf("%v", ciscoflashcopyentry.Ciscoflashcopyserialnumber) + "']"
}

func (ciscoflashcopyentry *CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoflashcopyentry *CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoflashcopyentry *CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ciscoFlashCopySerialNumber"] = ciscoflashcopyentry.Ciscoflashcopyserialnumber
    leafs["ciscoFlashCopyCommand"] = ciscoflashcopyentry.Ciscoflashcopycommand
    leafs["ciscoFlashCopyProtocol"] = ciscoflashcopyentry.Ciscoflashcopyprotocol
    leafs["ciscoFlashCopyServerAddress"] = ciscoflashcopyentry.Ciscoflashcopyserveraddress
    leafs["ciscoFlashCopySourceName"] = ciscoflashcopyentry.Ciscoflashcopysourcename
    leafs["ciscoFlashCopyDestinationName"] = ciscoflashcopyentry.Ciscoflashcopydestinationname
    leafs["ciscoFlashCopyRemoteUserName"] = ciscoflashcopyentry.Ciscoflashcopyremoteusername
    leafs["ciscoFlashCopyStatus"] = ciscoflashcopyentry.Ciscoflashcopystatus
    leafs["ciscoFlashCopyNotifyOnCompletion"] = ciscoflashcopyentry.Ciscoflashcopynotifyoncompletion
    leafs["ciscoFlashCopyTime"] = ciscoflashcopyentry.Ciscoflashcopytime
    leafs["ciscoFlashCopyEntryStatus"] = ciscoflashcopyentry.Ciscoflashcopyentrystatus
    leafs["ciscoFlashCopyVerify"] = ciscoflashcopyentry.Ciscoflashcopyverify
    leafs["ciscoFlashCopyServerAddrType"] = ciscoflashcopyentry.Ciscoflashcopyserveraddrtype
    leafs["ciscoFlashCopyServerAddrRev1"] = ciscoflashcopyentry.Ciscoflashcopyserveraddrrev1
    leafs["ciscoFlashCopyRemotePassword"] = ciscoflashcopyentry.Ciscoflashcopyremotepassword
    return leafs
}

func (ciscoflashcopyentry *CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoflashcopyentry *CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry) GetYangName() string { return "ciscoFlashCopyEntry" }

func (ciscoflashcopyentry *CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoflashcopyentry *CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoflashcopyentry *CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoflashcopyentry *CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry) SetParent(parent types.Entity) { ciscoflashcopyentry.parent = parent }

func (ciscoflashcopyentry *CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry) GetParent() types.Entity { return ciscoflashcopyentry.parent }

func (ciscoflashcopyentry *CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry) GetParentYangName() string { return "ciscoFlashCopyTable" }

// CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopycommand represents                         CopyNotifyOnCompletion (opt)
type CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopycommand string

const (
    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopycommand_copyToFlashWithErase CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopycommand = "copyToFlashWithErase"

    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopycommand_copyToFlashWithoutErase CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopycommand = "copyToFlashWithoutErase"

    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopycommand_copyFromFlash CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopycommand = "copyFromFlash"

    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopycommand_copyFromFlhLog CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopycommand = "copyFromFlhLog"
)

// CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopyprotocol represents the protocol support for a command.
type CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopyprotocol string

const (
    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopyprotocol_tftp CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopyprotocol = "tftp"

    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopyprotocol_rcp CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopyprotocol = "rcp"

    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopyprotocol_lex CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopyprotocol = "lex"

    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopyprotocol_ftp CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopyprotocol = "ftp"

    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopyprotocol_scp CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopyprotocol = "scp"

    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopyprotocol_sftp CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopyprotocol = "sftp"
)

// CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus represents       stop user from overwriting current boot image file.
type CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus string

const (
    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus_copyOperationPending CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus = "copyOperationPending"

    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus_copyInProgress CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus = "copyInProgress"

    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus_copyOperationSuccess CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus = "copyOperationSuccess"

    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus_copyInvalidOperation CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus = "copyInvalidOperation"

    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus_copyInvalidProtocol CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus = "copyInvalidProtocol"

    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus_copyInvalidSourceName CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus = "copyInvalidSourceName"

    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus_copyInvalidDestName CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus = "copyInvalidDestName"

    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus_copyInvalidServerAddress CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus = "copyInvalidServerAddress"

    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus_copyDeviceBusy CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus = "copyDeviceBusy"

    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus_copyDeviceOpenError CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus = "copyDeviceOpenError"

    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus_copyDeviceError CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus = "copyDeviceError"

    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus_copyDeviceNotProgrammable CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus = "copyDeviceNotProgrammable"

    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus_copyDeviceFull CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus = "copyDeviceFull"

    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus_copyFileOpenError CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus = "copyFileOpenError"

    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus_copyFileTransferError CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus = "copyFileTransferError"

    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus_copyFileChecksumError CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus = "copyFileChecksumError"

    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus_copyNoMemory CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus = "copyNoMemory"

    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus_copyUnknownFailure CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus = "copyUnknownFailure"

    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus_copyInvalidSignature CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus = "copyInvalidSignature"

    CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus_copyProhibited CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry_Ciscoflashcopystatus = "copyProhibited"
)

// CISCOFLASHMIB_Ciscoflashpartitioningtable
// A table of Flash partitioning operation entries. Each
// entry represents a Flash partitioning operation that
// has been initiated.
type CISCOFLASHMIB_Ciscoflashpartitioningtable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A Flash partitioning operation entry. Each entry consists of the command,
    // the target device, the partition count, and optionally the partition sizes.
    // A management station wishing to create an entry should first generate a
    // pseudo-random serial number to be used as the index to this sparse table. 
    // The station should then create the associated instance of the row status
    // object. It must also, either in the same or in successive PDUs, create the
    // associated instance of the command and parameter objects. It should also
    // modify the default values for any of the parameter objects if the defaults
    // are not appropriate.  Once the appropriate instances of all the command
    // objects have been created, either by an explicit SNMP set request or by
    // default, the row status should be set to active to initiate the operation.
    // Note that this entire procedure may be initiated via a single set request
    // which specifies a row status of createAndGo as well as specifies valid
    // values for the non-defaulted parameter objects.  Once an operation has been
    // activated, it cannot be stopped.  Once the operation completes, the
    // management station should retrieve the value of the status object (and time
    // if desired), and delete the entry.  In order to prevent old entries from
    // clogging the table, entries will be aged out, but an entry will never be
    // deleted within 5 minutes of completing. The type is slice of
    // CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry.
    Ciscoflashpartitioningentry []CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry
}

func (ciscoflashpartitioningtable *CISCOFLASHMIB_Ciscoflashpartitioningtable) GetFilter() yfilter.YFilter { return ciscoflashpartitioningtable.YFilter }

func (ciscoflashpartitioningtable *CISCOFLASHMIB_Ciscoflashpartitioningtable) SetFilter(yf yfilter.YFilter) { ciscoflashpartitioningtable.YFilter = yf }

func (ciscoflashpartitioningtable *CISCOFLASHMIB_Ciscoflashpartitioningtable) GetGoName(yname string) string {
    if yname == "ciscoFlashPartitioningEntry" { return "Ciscoflashpartitioningentry" }
    return ""
}

func (ciscoflashpartitioningtable *CISCOFLASHMIB_Ciscoflashpartitioningtable) GetSegmentPath() string {
    return "ciscoFlashPartitioningTable"
}

func (ciscoflashpartitioningtable *CISCOFLASHMIB_Ciscoflashpartitioningtable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoFlashPartitioningEntry" {
        for _, c := range ciscoflashpartitioningtable.Ciscoflashpartitioningentry {
            if ciscoflashpartitioningtable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry{}
        ciscoflashpartitioningtable.Ciscoflashpartitioningentry = append(ciscoflashpartitioningtable.Ciscoflashpartitioningentry, child)
        return &ciscoflashpartitioningtable.Ciscoflashpartitioningentry[len(ciscoflashpartitioningtable.Ciscoflashpartitioningentry)-1]
    }
    return nil
}

func (ciscoflashpartitioningtable *CISCOFLASHMIB_Ciscoflashpartitioningtable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ciscoflashpartitioningtable.Ciscoflashpartitioningentry {
        children[ciscoflashpartitioningtable.Ciscoflashpartitioningentry[i].GetSegmentPath()] = &ciscoflashpartitioningtable.Ciscoflashpartitioningentry[i]
    }
    return children
}

func (ciscoflashpartitioningtable *CISCOFLASHMIB_Ciscoflashpartitioningtable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciscoflashpartitioningtable *CISCOFLASHMIB_Ciscoflashpartitioningtable) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoflashpartitioningtable *CISCOFLASHMIB_Ciscoflashpartitioningtable) GetYangName() string { return "ciscoFlashPartitioningTable" }

func (ciscoflashpartitioningtable *CISCOFLASHMIB_Ciscoflashpartitioningtable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoflashpartitioningtable *CISCOFLASHMIB_Ciscoflashpartitioningtable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoflashpartitioningtable *CISCOFLASHMIB_Ciscoflashpartitioningtable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoflashpartitioningtable *CISCOFLASHMIB_Ciscoflashpartitioningtable) SetParent(parent types.Entity) { ciscoflashpartitioningtable.parent = parent }

func (ciscoflashpartitioningtable *CISCOFLASHMIB_Ciscoflashpartitioningtable) GetParent() types.Entity { return ciscoflashpartitioningtable.parent }

func (ciscoflashpartitioningtable *CISCOFLASHMIB_Ciscoflashpartitioningtable) GetParentYangName() string { return "CISCO-FLASH-MIB" }

// CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry
// A Flash partitioning operation entry. Each entry
// consists of the command, the target device, the
// partition count, and optionally the partition sizes.
// 
// A management station wishing to create an entry should
// first generate a pseudo-random serial number to be used
// as the index to this sparse table.  The station should
// then create the associated instance of the row status
// object. It must also, either in the same or in successive
// PDUs, create the associated instance of the command and
// parameter objects. It should also modify the default values
// for any of the parameter objects if the defaults are not
// appropriate.
// 
// Once the appropriate instances of all the command
// objects have been created, either by an explicit SNMP
// set request or by default, the row status should be set
// to active to initiate the operation. Note that this entire
// procedure may be initiated via a single set request which
// specifies a row status of createAndGo as well as specifies
// valid values for the non-defaulted parameter objects.
// 
// Once an operation has been activated, it cannot be
// stopped.
// 
// Once the operation completes, the management station should
// retrieve the value of the status object (and time if
// desired), and delete the entry.  In order to prevent old
// entries from clogging the table, entries will be aged out,
// but an entry will never be deleted within 5 minutes of
// completing.
type CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Object which specifies a unique entry in the
    // partitioning operations table. A management station wishing to initiate a
    // partitioning operation should use a pseudo-random value for this object
    // when creating or modifying an instance of a ciscoFlashPartitioningEntry.
    // The type is interface{} with range: 0..2147483647.
    Ciscoflashpartitioningserialnumber interface{}

    // The partitioning command to be executed. Mandatory. If the command is
    // unsupported, the partitioningInvalidOperation error will be reported in the
    // operation status.  Command                 Remarks partition              
    // Partition a Flash device.                         All the prerequisites for
    // partitioning must be met for                         this command to
    // succeed.  Command table           Parameters 1) partition           
    // PartitioningDestinationName                        
    // PartitioningPartitionCount                        
    // PartitioningPartitionSizes (opt)                        
    // PartitioningNotifyOnCompletion (opt). The type is
    // Ciscoflashpartitioningcommand.
    Ciscoflashpartitioningcommand interface{}

    // Destination device name. This name will be the value obtained from
    // FlashDeviceName. If the name is not specified, the default Flash device
    // will be assumed. The type is string with length: 0..255.
    Ciscoflashpartitioningdestinationname interface{}

    // This object is used to specify the number of partitions to be created. Its
    // value cannot exceed the value of ciscoFlashDeviceMaxPartitions.  To undo
    // partitioning (revert to a single partition), this object must have the
    // value 1. The type is interface{} with range: 1..4294967295.
    Ciscoflashpartitioningpartitioncount interface{}

    // This object is used to explicitly specify the size of each partition to be
    // created. The size of each partition will be in units of
    // ciscoFlashDeviceMinPartitionSize. The value of this object will be in the
    // form:         <part1>:<part2>...:<partn>  If partition sizes are not
    // specified, the system will calculate default sizes based on the partition
    // count, the minimum partition size, and the device size. Partition size need
    // not be specified when undoing partitioning (partition count is 1). If
    // partition sizes are specified, the number of sizes specified must exactly
    // match the partition count. If not, the partitioning command will be
    // rejected with the invalidPartitionSizes error . The type is string.
    Ciscoflashpartitioningpartitionsizes interface{}

    // The status of the specified partitioning operation. partitioningInProgress
    // :         specified operation is active  partitioningOperationSuccess :    
    // specified operation is supported and completed         successfully 
    // partitioningInvalidOperation :         command invalid or
    // command-protocol-device         combination unsupported 
    // partitioningInvalidDestName :         invalid target name (file or
    // partition or         device name) specified 
    // partitioningInvalidPartitionCount :         invalid partition count
    // specified for the         partitioning command 
    // partitioningInvalidPartitionSizes :         invalid partition size, or
    // invalid count of         partition sizes  partitioningDeviceBusy :        
    // specified device is in use and locked by         another process 
    // partitioningDeviceOpenError :         invalid device name 
    // partitioningDeviceError :         device read, write or erase error 
    // partitioningNoMemory :         system running low on memory 
    // partitioningUnknownFailure :         failure unknown. The type is
    // Ciscoflashpartitioningstatus.
    Ciscoflashpartitioningstatus interface{}

    // Specifies whether or not a notification should be generated on the
    // completion of the partitioning operation. If specified,
    // ciscoFlashPartitioningCompletionTrap will be generated. It is the
    // responsibility of the management entity to ensure that the SNMP
    // administrative model is configured in such a way as to allow the
    // notification to be delivered. The type is bool.
    Ciscoflashpartitioningnotifyoncompletion interface{}

    // Time taken for the operation. This object will be like a stopwatch,
    // starting when the operation starts, stopping when the operation completes.
    // If a management entity keeps a database of completion times for various
    // operations, it can then use the stopwatch capability to display percentage
    // completion time. The type is interface{} with range: 0..4294967295.
    Ciscoflashpartitioningtime interface{}

    // The status of this table entry. The type is RowStatus.
    Ciscoflashpartitioningentrystatus interface{}
}

func (ciscoflashpartitioningentry *CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry) GetFilter() yfilter.YFilter { return ciscoflashpartitioningentry.YFilter }

func (ciscoflashpartitioningentry *CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry) SetFilter(yf yfilter.YFilter) { ciscoflashpartitioningentry.YFilter = yf }

func (ciscoflashpartitioningentry *CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry) GetGoName(yname string) string {
    if yname == "ciscoFlashPartitioningSerialNumber" { return "Ciscoflashpartitioningserialnumber" }
    if yname == "ciscoFlashPartitioningCommand" { return "Ciscoflashpartitioningcommand" }
    if yname == "ciscoFlashPartitioningDestinationName" { return "Ciscoflashpartitioningdestinationname" }
    if yname == "ciscoFlashPartitioningPartitionCount" { return "Ciscoflashpartitioningpartitioncount" }
    if yname == "ciscoFlashPartitioningPartitionSizes" { return "Ciscoflashpartitioningpartitionsizes" }
    if yname == "ciscoFlashPartitioningStatus" { return "Ciscoflashpartitioningstatus" }
    if yname == "ciscoFlashPartitioningNotifyOnCompletion" { return "Ciscoflashpartitioningnotifyoncompletion" }
    if yname == "ciscoFlashPartitioningTime" { return "Ciscoflashpartitioningtime" }
    if yname == "ciscoFlashPartitioningEntryStatus" { return "Ciscoflashpartitioningentrystatus" }
    return ""
}

func (ciscoflashpartitioningentry *CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry) GetSegmentPath() string {
    return "ciscoFlashPartitioningEntry" + "[ciscoFlashPartitioningSerialNumber='" + fmt.Sprintf("%v", ciscoflashpartitioningentry.Ciscoflashpartitioningserialnumber) + "']"
}

func (ciscoflashpartitioningentry *CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoflashpartitioningentry *CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoflashpartitioningentry *CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ciscoFlashPartitioningSerialNumber"] = ciscoflashpartitioningentry.Ciscoflashpartitioningserialnumber
    leafs["ciscoFlashPartitioningCommand"] = ciscoflashpartitioningentry.Ciscoflashpartitioningcommand
    leafs["ciscoFlashPartitioningDestinationName"] = ciscoflashpartitioningentry.Ciscoflashpartitioningdestinationname
    leafs["ciscoFlashPartitioningPartitionCount"] = ciscoflashpartitioningentry.Ciscoflashpartitioningpartitioncount
    leafs["ciscoFlashPartitioningPartitionSizes"] = ciscoflashpartitioningentry.Ciscoflashpartitioningpartitionsizes
    leafs["ciscoFlashPartitioningStatus"] = ciscoflashpartitioningentry.Ciscoflashpartitioningstatus
    leafs["ciscoFlashPartitioningNotifyOnCompletion"] = ciscoflashpartitioningentry.Ciscoflashpartitioningnotifyoncompletion
    leafs["ciscoFlashPartitioningTime"] = ciscoflashpartitioningentry.Ciscoflashpartitioningtime
    leafs["ciscoFlashPartitioningEntryStatus"] = ciscoflashpartitioningentry.Ciscoflashpartitioningentrystatus
    return leafs
}

func (ciscoflashpartitioningentry *CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoflashpartitioningentry *CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry) GetYangName() string { return "ciscoFlashPartitioningEntry" }

func (ciscoflashpartitioningentry *CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoflashpartitioningentry *CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoflashpartitioningentry *CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoflashpartitioningentry *CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry) SetParent(parent types.Entity) { ciscoflashpartitioningentry.parent = parent }

func (ciscoflashpartitioningentry *CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry) GetParent() types.Entity { return ciscoflashpartitioningentry.parent }

func (ciscoflashpartitioningentry *CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry) GetParentYangName() string { return "ciscoFlashPartitioningTable" }

// CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningcommand represents                         PartitioningNotifyOnCompletion (opt)
type CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningcommand string

const (
    CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningcommand_partition CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningcommand = "partition"
)

// CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningstatus represents         failure unknown
type CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningstatus string

const (
    CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningstatus_partitioningInProgress CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningstatus = "partitioningInProgress"

    CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningstatus_partitioningOperationSuccess CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningstatus = "partitioningOperationSuccess"

    CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningstatus_partitioningInvalidOperation CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningstatus = "partitioningInvalidOperation"

    CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningstatus_partitioningInvalidDestName CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningstatus = "partitioningInvalidDestName"

    CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningstatus_partitioningInvalidPartitionCount CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningstatus = "partitioningInvalidPartitionCount"

    CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningstatus_partitioningInvalidPartitionSizes CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningstatus = "partitioningInvalidPartitionSizes"

    CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningstatus_partitioningDeviceBusy CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningstatus = "partitioningDeviceBusy"

    CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningstatus_partitioningDeviceOpenError CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningstatus = "partitioningDeviceOpenError"

    CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningstatus_partitioningDeviceError CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningstatus = "partitioningDeviceError"

    CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningstatus_partitioningNoMemory CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningstatus = "partitioningNoMemory"

    CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningstatus_partitioningUnknownFailure CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry_Ciscoflashpartitioningstatus = "partitioningUnknownFailure"
)

// CISCOFLASHMIB_Ciscoflashmiscoptable
// A table of misc Flash operation entries. Each
// entry represents a Flash operation that
// has been initiated.
type CISCOFLASHMIB_Ciscoflashmiscoptable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A Flash operation entry. Each entry consists of a command, a target, and
    // any optional parameters.  A management station wishing to create an entry
    // should first generate a pseudo-random serial number to be used as the index
    // to this sparse table.  The station should then create the associated
    // instance of the row status object. It must also, either in the same or in
    // successive PDUs, create the associated instance of the command and
    // parameter objects. It should also modify the default values for any of the
    // parameter objects if the defaults are not appropriate.  Once the
    // appropriate instances of all the command objects have been created, either
    // by an explicit SNMP set request or by default, the row status should be set
    // to active to initiate the operation. Note that this entire procedure may be
    // initiated via a single set request which specifies a row status of
    // createAndGo as well as specifies valid values for the non-defaulted
    // parameter objects.  Once an operation has been activated, it cannot be
    // stopped.  Once the operation completes, the management station should
    // retrieve the value of the status object (and time if desired), and delete
    // the entry.  In order to prevent old entries from clogging the table,
    // entries will be aged out, but an entry will never be deleted within 5
    // minutes of completing. The type is slice of
    // CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry.
    Ciscoflashmiscopentry []CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry
}

func (ciscoflashmiscoptable *CISCOFLASHMIB_Ciscoflashmiscoptable) GetFilter() yfilter.YFilter { return ciscoflashmiscoptable.YFilter }

func (ciscoflashmiscoptable *CISCOFLASHMIB_Ciscoflashmiscoptable) SetFilter(yf yfilter.YFilter) { ciscoflashmiscoptable.YFilter = yf }

func (ciscoflashmiscoptable *CISCOFLASHMIB_Ciscoflashmiscoptable) GetGoName(yname string) string {
    if yname == "ciscoFlashMiscOpEntry" { return "Ciscoflashmiscopentry" }
    return ""
}

func (ciscoflashmiscoptable *CISCOFLASHMIB_Ciscoflashmiscoptable) GetSegmentPath() string {
    return "ciscoFlashMiscOpTable"
}

func (ciscoflashmiscoptable *CISCOFLASHMIB_Ciscoflashmiscoptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoFlashMiscOpEntry" {
        for _, c := range ciscoflashmiscoptable.Ciscoflashmiscopentry {
            if ciscoflashmiscoptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry{}
        ciscoflashmiscoptable.Ciscoflashmiscopentry = append(ciscoflashmiscoptable.Ciscoflashmiscopentry, child)
        return &ciscoflashmiscoptable.Ciscoflashmiscopentry[len(ciscoflashmiscoptable.Ciscoflashmiscopentry)-1]
    }
    return nil
}

func (ciscoflashmiscoptable *CISCOFLASHMIB_Ciscoflashmiscoptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range ciscoflashmiscoptable.Ciscoflashmiscopentry {
        children[ciscoflashmiscoptable.Ciscoflashmiscopentry[i].GetSegmentPath()] = &ciscoflashmiscoptable.Ciscoflashmiscopentry[i]
    }
    return children
}

func (ciscoflashmiscoptable *CISCOFLASHMIB_Ciscoflashmiscoptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ciscoflashmiscoptable *CISCOFLASHMIB_Ciscoflashmiscoptable) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoflashmiscoptable *CISCOFLASHMIB_Ciscoflashmiscoptable) GetYangName() string { return "ciscoFlashMiscOpTable" }

func (ciscoflashmiscoptable *CISCOFLASHMIB_Ciscoflashmiscoptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoflashmiscoptable *CISCOFLASHMIB_Ciscoflashmiscoptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoflashmiscoptable *CISCOFLASHMIB_Ciscoflashmiscoptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoflashmiscoptable *CISCOFLASHMIB_Ciscoflashmiscoptable) SetParent(parent types.Entity) { ciscoflashmiscoptable.parent = parent }

func (ciscoflashmiscoptable *CISCOFLASHMIB_Ciscoflashmiscoptable) GetParent() types.Entity { return ciscoflashmiscoptable.parent }

func (ciscoflashmiscoptable *CISCOFLASHMIB_Ciscoflashmiscoptable) GetParentYangName() string { return "CISCO-FLASH-MIB" }

// CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry
// A Flash operation entry. Each entry consists of a
// command, a target, and any optional parameters.
// 
// A management station wishing to create an entry should
// first generate a pseudo-random serial number to be used
// as the index to this sparse table.  The station should
// then create the associated instance of the row status
// object. It must also, either in the same or in successive
// PDUs, create the associated instance of the command and
// parameter objects. It should also modify the default values
// for any of the parameter objects if the defaults are not
// appropriate.
// 
// Once the appropriate instances of all the command
// objects have been created, either by an explicit SNMP
// set request or by default, the row status should be set
// to active to initiate the operation. Note that this entire
// procedure may be initiated via a single set request which
// specifies a row status of createAndGo as well as specifies
// valid values for the non-defaulted parameter objects.
// 
// Once an operation has been activated, it cannot be
// stopped.
// 
// Once the operation completes, the management station should
// retrieve the value of the status object (and time if
// desired), and delete the entry.  In order to prevent old
// entries from clogging the table, entries will be aged out,
// but an entry will never be deleted within 5 minutes of
// completing.
type CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Object which specifies a unique entry in the
    // table. A management station wishing to initiate a flash operation should
    // use a pseudo-random value for this object when creating or modifying an
    // instance of a ciscoFlashMiscOpEntry. The type is interface{} with range:
    // 0..2147483647.
    Ciscoflashmiscopserialnumber interface{}

    // The command to be executed. Mandatory. Note that it is possible for a
    // system to support multiple file systems (different file systems on
    // different Flash devices, or different file systems on different partitions
    // within a device). Each such file system may support only a subset of these
    // commands. If a command is unsupported, the miscOpInvalidOperation(3) error
    // will be reported in the operation status.  Command         Remarks erase   
    // Erase flash. verify          Verify flash file checksum. delete         
    // Delete a file. undelete        Revive a deleted file .                 Note
    // that there are limits on                 the number of times a file can    
    // be deleted and undeleted. When                 this limit is exceeded, the 
    // system will return the appropriate                 error. squeeze        
    // Recover space occupied by                 deleted files. This command      
    // preserves the good files, erases                 out the file system, then
    // restores                 the preserved good files. format          Format a
    // flash device.  Command table   Parameters erase          
    // MiscOpDestinationName                 MiscOpNotifyOnCompletion (opt) verify
    // MiscOpDestinationName                 MiscOpNotifyOnCompletion (opt) delete
    // MiscOpDestinationName                 MiscOpNotifyOnCompletion (opt)
    // undelete        MiscOpDestinationName                
    // MiscOpNotifyOnCompletion (opt) squeeze         MiscOpDestinationName       
    // MiscOpNotifyOnCompletion (opt) format          MiscOpDestinationName       
    // MiscOpNotifyOnCompletion (opt). The type is Ciscoflashmiscopcommand.
    Ciscoflashmiscopcommand interface{}

    // Destination file, or partition name. File name must be of the form        
    // [device>:][<partition>:]<file> where <device> is a value obtained from
    // FlashDeviceName,       <partition> is obtained from FlashPartitionName  
    // and <file> is the name of a file in Flash. While leading and/or trailing
    // whitespaces are acceptable, no whitespaces are allowed within the path
    // itself.  A management station could derive its own partition name as per
    // the description for the ciscoFlashPartitionName object. If <device> is not
    // specified, the default Flash device will be assumed. If <partition> is not
    // specified, the default partition will be assumed. If a device is not
    // partitioned into 2 or more partitions, this value may be left out.  For an
    // operation on a partition, eg., the erase command, this object would specify
    // the partition name in the form:         [device>:][<partition>:]. The type
    // is string with length: 0..255.
    Ciscoflashmiscopdestinationname interface{}

    // The status of the specified operation. miscOpInProgress :         specified
    // operation is active  miscOpOperationSuccess :         specified operation
    // is supported and completed         successfully  miscOpInvalidOperation :  
    // command invalid or command-protocol-device         combination unsupported 
    // miscOpInvalidDestName :         invalid target name (file or partition or  
    // device name) specified  miscOpDeviceBusy :         specified device is in
    // use and locked by another         process  miscOpDeviceOpenError :        
    // invalid device name  miscOpDeviceError :         device read, write or
    // erase error  miscOpDeviceNotProgrammable :         device is read-only but
    // a write or erase         operation was specified  miscOpFileOpenError :    
    // invalid file name; file not found in partition  miscOpFileDeleteFailure :  
    // file could not be deleted; delete count exceeded  miscOpFileUndeleteFailure
    // :         file could not be undeleted; undelete count         exceeded 
    // miscOpFileChecksumError :         file has a bad checksum  miscOpNoMemory :
    // system running low on memory  miscOpUnknownFailure :         failure
    // unknown  miscOpSqueezeFailure :         the squeeze operation failed 
    // miscOpNoSuchFile :         a valid but nonexistent file name was specified 
    // miscOpFormatFailure :         the format operation failed. The type is
    // Ciscoflashmiscopstatus.
    Ciscoflashmiscopstatus interface{}

    // Specifies whether or not a notification should be generated on the
    // completion of an operation. If specified, ciscoFlashMiscOpCompletionTrap
    // will be generated. It is the responsibility of the management entity to
    // ensure that the SNMP administrative model is configured in such a way as to
    // allow the notification to be delivered. The type is bool.
    Ciscoflashmiscopnotifyoncompletion interface{}

    // Time taken for the operation. This object will be like a stopwatch,
    // starting when the operation starts, stopping when the operation completes.
    // If a management entity keeps a database of completion times for various
    // operations, it can then use the stopwatch capability to display percentage
    // completion time. The type is interface{} with range: 0..4294967295.
    Ciscoflashmiscoptime interface{}

    // The status of this table entry. The type is RowStatus.
    Ciscoflashmiscopentrystatus interface{}
}

func (ciscoflashmiscopentry *CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry) GetFilter() yfilter.YFilter { return ciscoflashmiscopentry.YFilter }

func (ciscoflashmiscopentry *CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry) SetFilter(yf yfilter.YFilter) { ciscoflashmiscopentry.YFilter = yf }

func (ciscoflashmiscopentry *CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry) GetGoName(yname string) string {
    if yname == "ciscoFlashMiscOpSerialNumber" { return "Ciscoflashmiscopserialnumber" }
    if yname == "ciscoFlashMiscOpCommand" { return "Ciscoflashmiscopcommand" }
    if yname == "ciscoFlashMiscOpDestinationName" { return "Ciscoflashmiscopdestinationname" }
    if yname == "ciscoFlashMiscOpStatus" { return "Ciscoflashmiscopstatus" }
    if yname == "ciscoFlashMiscOpNotifyOnCompletion" { return "Ciscoflashmiscopnotifyoncompletion" }
    if yname == "ciscoFlashMiscOpTime" { return "Ciscoflashmiscoptime" }
    if yname == "ciscoFlashMiscOpEntryStatus" { return "Ciscoflashmiscopentrystatus" }
    return ""
}

func (ciscoflashmiscopentry *CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry) GetSegmentPath() string {
    return "ciscoFlashMiscOpEntry" + "[ciscoFlashMiscOpSerialNumber='" + fmt.Sprintf("%v", ciscoflashmiscopentry.Ciscoflashmiscopserialnumber) + "']"
}

func (ciscoflashmiscopentry *CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoflashmiscopentry *CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoflashmiscopentry *CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["ciscoFlashMiscOpSerialNumber"] = ciscoflashmiscopentry.Ciscoflashmiscopserialnumber
    leafs["ciscoFlashMiscOpCommand"] = ciscoflashmiscopentry.Ciscoflashmiscopcommand
    leafs["ciscoFlashMiscOpDestinationName"] = ciscoflashmiscopentry.Ciscoflashmiscopdestinationname
    leafs["ciscoFlashMiscOpStatus"] = ciscoflashmiscopentry.Ciscoflashmiscopstatus
    leafs["ciscoFlashMiscOpNotifyOnCompletion"] = ciscoflashmiscopentry.Ciscoflashmiscopnotifyoncompletion
    leafs["ciscoFlashMiscOpTime"] = ciscoflashmiscopentry.Ciscoflashmiscoptime
    leafs["ciscoFlashMiscOpEntryStatus"] = ciscoflashmiscopentry.Ciscoflashmiscopentrystatus
    return leafs
}

func (ciscoflashmiscopentry *CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoflashmiscopentry *CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry) GetYangName() string { return "ciscoFlashMiscOpEntry" }

func (ciscoflashmiscopentry *CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoflashmiscopentry *CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoflashmiscopentry *CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoflashmiscopentry *CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry) SetParent(parent types.Entity) { ciscoflashmiscopentry.parent = parent }

func (ciscoflashmiscopentry *CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry) GetParent() types.Entity { return ciscoflashmiscopentry.parent }

func (ciscoflashmiscopentry *CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry) GetParentYangName() string { return "ciscoFlashMiscOpTable" }

// CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopcommand represents                 MiscOpNotifyOnCompletion (opt)
type CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopcommand string

const (
    CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopcommand_erase CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopcommand = "erase"

    CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopcommand_verify CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopcommand = "verify"

    CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopcommand_delete CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopcommand = "delete"

    CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopcommand_undelete CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopcommand = "undelete"

    CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopcommand_squeeze CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopcommand = "squeeze"

    CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopcommand_format CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopcommand = "format"
)

// CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus represents         the format operation failed
type CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus string

const (
    CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus_miscOpInProgress CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus = "miscOpInProgress"

    CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus_miscOpOperationSuccess CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus = "miscOpOperationSuccess"

    CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus_miscOpInvalidOperation CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus = "miscOpInvalidOperation"

    CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus_miscOpInvalidDestName CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus = "miscOpInvalidDestName"

    CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus_miscOpDeviceBusy CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus = "miscOpDeviceBusy"

    CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus_miscOpDeviceOpenError CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus = "miscOpDeviceOpenError"

    CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus_miscOpDeviceError CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus = "miscOpDeviceError"

    CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus_miscOpDeviceNotProgrammable CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus = "miscOpDeviceNotProgrammable"

    CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus_miscOpFileOpenError CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus = "miscOpFileOpenError"

    CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus_miscOpFileDeleteFailure CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus = "miscOpFileDeleteFailure"

    CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus_miscOpFileUndeleteFailure CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus = "miscOpFileUndeleteFailure"

    CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus_miscOpFileChecksumError CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus = "miscOpFileChecksumError"

    CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus_miscOpNoMemory CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus = "miscOpNoMemory"

    CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus_miscOpUnknownFailure CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus = "miscOpUnknownFailure"

    CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus_miscOpSqueezeFailure CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus = "miscOpSqueezeFailure"

    CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus_miscOpNoSuchFile CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus = "miscOpNoSuchFile"

    CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus_miscOpFormatFailure CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry_Ciscoflashmiscopstatus = "miscOpFormatFailure"
)

