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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CiscoFlashDevice CISCOFLASHMIB_CiscoFlashDevice

    
    CiscoFlashCfg CISCOFLASHMIB_CiscoFlashCfg

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
    CiscoFlashDeviceTable CISCOFLASHMIB_CiscoFlashDeviceTable

    // Table of Flash device chip properties for each initialized Flash device.
    // This table is meant primarily for aiding error diagnosis.
    CiscoFlashChipTable CISCOFLASHMIB_CiscoFlashChipTable

    // Table of flash device partition properties for each initialized flash
    // partition. Whenever there is no explicit partitioning done, a single
    // partition spanning the entire device will be assumed to exist. There will
    // therefore always be atleast one partition on a device.
    CiscoFlashPartitionTable CISCOFLASHMIB_CiscoFlashPartitionTable

    // Table of information for files in a Flash partition.
    CiscoFlashFileTable CISCOFLASHMIB_CiscoFlashFileTable

    // Table of information for files on the manageable flash devices sorted by
    // File Types.
    CiscoFlashFileByTypeTable CISCOFLASHMIB_CiscoFlashFileByTypeTable

    // A table of Flash copy operation entries. Each entry represents a Flash copy
    // operation (to or from Flash) that has been initiated.
    CiscoFlashCopyTable CISCOFLASHMIB_CiscoFlashCopyTable

    // A table of Flash partitioning operation entries. Each entry represents a
    // Flash partitioning operation that has been initiated.
    CiscoFlashPartitioningTable CISCOFLASHMIB_CiscoFlashPartitioningTable

    // A table of misc Flash operation entries. Each entry represents a Flash
    // operation that has been initiated.
    CiscoFlashMiscOpTable CISCOFLASHMIB_CiscoFlashMiscOpTable
}

func (cISCOFLASHMIB *CISCOFLASHMIB) GetEntityData() *types.CommonEntityData {
    cISCOFLASHMIB.EntityData.YFilter = cISCOFLASHMIB.YFilter
    cISCOFLASHMIB.EntityData.YangName = "CISCO-FLASH-MIB"
    cISCOFLASHMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOFLASHMIB.EntityData.ParentYangName = "CISCO-FLASH-MIB"
    cISCOFLASHMIB.EntityData.SegmentPath = "CISCO-FLASH-MIB:CISCO-FLASH-MIB"
    cISCOFLASHMIB.EntityData.AbsolutePath = cISCOFLASHMIB.EntityData.SegmentPath
    cISCOFLASHMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOFLASHMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOFLASHMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOFLASHMIB.EntityData.Children = types.NewOrderedMap()
    cISCOFLASHMIB.EntityData.Children.Append("ciscoFlashDevice", types.YChild{"CiscoFlashDevice", &cISCOFLASHMIB.CiscoFlashDevice})
    cISCOFLASHMIB.EntityData.Children.Append("ciscoFlashCfg", types.YChild{"CiscoFlashCfg", &cISCOFLASHMIB.CiscoFlashCfg})
    cISCOFLASHMIB.EntityData.Children.Append("ciscoFlashDeviceTable", types.YChild{"CiscoFlashDeviceTable", &cISCOFLASHMIB.CiscoFlashDeviceTable})
    cISCOFLASHMIB.EntityData.Children.Append("ciscoFlashChipTable", types.YChild{"CiscoFlashChipTable", &cISCOFLASHMIB.CiscoFlashChipTable})
    cISCOFLASHMIB.EntityData.Children.Append("ciscoFlashPartitionTable", types.YChild{"CiscoFlashPartitionTable", &cISCOFLASHMIB.CiscoFlashPartitionTable})
    cISCOFLASHMIB.EntityData.Children.Append("ciscoFlashFileTable", types.YChild{"CiscoFlashFileTable", &cISCOFLASHMIB.CiscoFlashFileTable})
    cISCOFLASHMIB.EntityData.Children.Append("ciscoFlashFileByTypeTable", types.YChild{"CiscoFlashFileByTypeTable", &cISCOFLASHMIB.CiscoFlashFileByTypeTable})
    cISCOFLASHMIB.EntityData.Children.Append("ciscoFlashCopyTable", types.YChild{"CiscoFlashCopyTable", &cISCOFLASHMIB.CiscoFlashCopyTable})
    cISCOFLASHMIB.EntityData.Children.Append("ciscoFlashPartitioningTable", types.YChild{"CiscoFlashPartitioningTable", &cISCOFLASHMIB.CiscoFlashPartitioningTable})
    cISCOFLASHMIB.EntityData.Children.Append("ciscoFlashMiscOpTable", types.YChild{"CiscoFlashMiscOpTable", &cISCOFLASHMIB.CiscoFlashMiscOpTable})
    cISCOFLASHMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOFLASHMIB.EntityData.YListKeys = []string {}

    return &(cISCOFLASHMIB.EntityData)
}

// CISCOFLASHMIB_CiscoFlashDevice
type CISCOFLASHMIB_CiscoFlashDevice struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Flash devices supported by the system. If the system does not
    // support any Flash devices, this MIB will not be loaded on that system. The
    // value of this object will therefore be atleast 1. The type is interface{}
    // with range: 0..4294967295.
    CiscoFlashDevicesSupported interface{}
}

func (ciscoFlashDevice *CISCOFLASHMIB_CiscoFlashDevice) GetEntityData() *types.CommonEntityData {
    ciscoFlashDevice.EntityData.YFilter = ciscoFlashDevice.YFilter
    ciscoFlashDevice.EntityData.YangName = "ciscoFlashDevice"
    ciscoFlashDevice.EntityData.BundleName = "cisco_ios_xe"
    ciscoFlashDevice.EntityData.ParentYangName = "CISCO-FLASH-MIB"
    ciscoFlashDevice.EntityData.SegmentPath = "ciscoFlashDevice"
    ciscoFlashDevice.EntityData.AbsolutePath = "CISCO-FLASH-MIB:CISCO-FLASH-MIB/" + ciscoFlashDevice.EntityData.SegmentPath
    ciscoFlashDevice.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoFlashDevice.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoFlashDevice.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoFlashDevice.EntityData.Children = types.NewOrderedMap()
    ciscoFlashDevice.EntityData.Leafs = types.NewOrderedMap()
    ciscoFlashDevice.EntityData.Leafs.Append("ciscoFlashDevicesSupported", types.YLeaf{"CiscoFlashDevicesSupported", ciscoFlashDevice.CiscoFlashDevicesSupported})

    ciscoFlashDevice.EntityData.YListKeys = []string {}

    return &(ciscoFlashDevice.EntityData)
}

// CISCOFLASHMIB_CiscoFlashCfg
type CISCOFLASHMIB_CiscoFlashCfg struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specifies whether or not a notification should be generated on the
    // insertion of a Flash device.  If the value of this object is 'true' then
    // the ciscoFlashDeviceInsertedNotif notification will be generated.  If the
    // value of this object is 'false' then the ciscoFlashDeviceInsertedNotif
    // notification will not be generated.  It is the responsibility of the
    // management entity to ensure that the SNMP administrative model is
    // configured in such a way as to allow the notification to be delivered. The
    // type is bool.
    CiscoFlashCfgDevInsNotifEnable interface{}

    // Specifies whether or not a notification should be generated on the removal
    // of a Flash device.  If the value of this object is 'true' then the
    // ciscoFlashDeviceRemovedNotif notification will be generated.  If the value
    // of this object is 'false' then the ciscoFlashDeviceRemovedNotif
    // notification will not be generated.  It is the responsibility of the
    // management entity to ensure that the SNMP administrative model is
    // configured in such a way as to allow the notification to be delivered. The
    // type is bool.
    CiscoFlashCfgDevRemNotifEnable interface{}

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
    CiscoFlashPartitionLowSpaceNotifEnable interface{}
}

func (ciscoFlashCfg *CISCOFLASHMIB_CiscoFlashCfg) GetEntityData() *types.CommonEntityData {
    ciscoFlashCfg.EntityData.YFilter = ciscoFlashCfg.YFilter
    ciscoFlashCfg.EntityData.YangName = "ciscoFlashCfg"
    ciscoFlashCfg.EntityData.BundleName = "cisco_ios_xe"
    ciscoFlashCfg.EntityData.ParentYangName = "CISCO-FLASH-MIB"
    ciscoFlashCfg.EntityData.SegmentPath = "ciscoFlashCfg"
    ciscoFlashCfg.EntityData.AbsolutePath = "CISCO-FLASH-MIB:CISCO-FLASH-MIB/" + ciscoFlashCfg.EntityData.SegmentPath
    ciscoFlashCfg.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoFlashCfg.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoFlashCfg.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoFlashCfg.EntityData.Children = types.NewOrderedMap()
    ciscoFlashCfg.EntityData.Leafs = types.NewOrderedMap()
    ciscoFlashCfg.EntityData.Leafs.Append("ciscoFlashCfgDevInsNotifEnable", types.YLeaf{"CiscoFlashCfgDevInsNotifEnable", ciscoFlashCfg.CiscoFlashCfgDevInsNotifEnable})
    ciscoFlashCfg.EntityData.Leafs.Append("ciscoFlashCfgDevRemNotifEnable", types.YLeaf{"CiscoFlashCfgDevRemNotifEnable", ciscoFlashCfg.CiscoFlashCfgDevRemNotifEnable})
    ciscoFlashCfg.EntityData.Leafs.Append("ciscoFlashPartitionLowSpaceNotifEnable", types.YLeaf{"CiscoFlashPartitionLowSpaceNotifEnable", ciscoFlashCfg.CiscoFlashPartitionLowSpaceNotifEnable})

    ciscoFlashCfg.EntityData.YListKeys = []string {}

    return &(ciscoFlashCfg.EntityData)
}

// CISCOFLASHMIB_CiscoFlashDeviceTable
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
type CISCOFLASHMIB_CiscoFlashDeviceTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table of flash device properties for each initialized flash
    // device. Each entry can be randomly accessed by using ciscoFlashDeviceIndex
    // as an index into the table. Note that removable devices will have an entry
    // in the table even when they have been removed. However, a non-removable
    // device that has not been installed will not have an entry in the table. The
    // type is slice of CISCOFLASHMIB_CiscoFlashDeviceTable_CiscoFlashDeviceEntry.
    CiscoFlashDeviceEntry []*CISCOFLASHMIB_CiscoFlashDeviceTable_CiscoFlashDeviceEntry
}

func (ciscoFlashDeviceTable *CISCOFLASHMIB_CiscoFlashDeviceTable) GetEntityData() *types.CommonEntityData {
    ciscoFlashDeviceTable.EntityData.YFilter = ciscoFlashDeviceTable.YFilter
    ciscoFlashDeviceTable.EntityData.YangName = "ciscoFlashDeviceTable"
    ciscoFlashDeviceTable.EntityData.BundleName = "cisco_ios_xe"
    ciscoFlashDeviceTable.EntityData.ParentYangName = "CISCO-FLASH-MIB"
    ciscoFlashDeviceTable.EntityData.SegmentPath = "ciscoFlashDeviceTable"
    ciscoFlashDeviceTable.EntityData.AbsolutePath = "CISCO-FLASH-MIB:CISCO-FLASH-MIB/" + ciscoFlashDeviceTable.EntityData.SegmentPath
    ciscoFlashDeviceTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoFlashDeviceTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoFlashDeviceTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoFlashDeviceTable.EntityData.Children = types.NewOrderedMap()
    ciscoFlashDeviceTable.EntityData.Children.Append("ciscoFlashDeviceEntry", types.YChild{"CiscoFlashDeviceEntry", nil})
    for i := range ciscoFlashDeviceTable.CiscoFlashDeviceEntry {
        ciscoFlashDeviceTable.EntityData.Children.Append(types.GetSegmentPath(ciscoFlashDeviceTable.CiscoFlashDeviceEntry[i]), types.YChild{"CiscoFlashDeviceEntry", ciscoFlashDeviceTable.CiscoFlashDeviceEntry[i]})
    }
    ciscoFlashDeviceTable.EntityData.Leafs = types.NewOrderedMap()

    ciscoFlashDeviceTable.EntityData.YListKeys = []string {}

    return &(ciscoFlashDeviceTable.EntityData)
}

// CISCOFLASHMIB_CiscoFlashDeviceTable_CiscoFlashDeviceEntry
// An entry in the table of flash device properties for
// each initialized flash device.
// Each entry can be randomly accessed by using
// ciscoFlashDeviceIndex as an index into the table.
// Note that removable devices will have an entry in
// the table even when they have been removed. However,
// a non-removable device that has not been installed
// will not have an entry in the table.
type CISCOFLASHMIB_CiscoFlashDeviceTable_CiscoFlashDeviceEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Flash device sequence number to index within the
    // table of initialized flash devices. The lowest value should be 1. The
    // highest should be less than or equal to the value of the
    // ciscoFlashDevicesSupported object. The type is interface{} with range:
    // 1..4294967295.
    CiscoFlashDeviceIndex interface{}

    // Total size of the Flash device. For a removable device, the size will be
    // zero if the device has been removed.  If the total size of the flash device
    // is greater than the maximum value reportable by this object then this
    // object should report its maximum value(4,294,967,295) and
    // ciscoFlashDeviceSizeExtended must be used to report the flash device's
    // size. The type is interface{} with range: 0..4294967295. Units are bytes.
    CiscoFlashDeviceSize interface{}

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
    CiscoFlashDeviceMinPartitionSize interface{}

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
    CiscoFlashDeviceMaxPartitions interface{}

    // Flash device partitions actually present. Number of partitions cannot
    // exceed the minimum of ciscoFlashDeviceMaxPartitions and
    // (ciscoFlashDeviceSize / ciscoFlashDeviceMinPartitionSize). Will be equal to
    // at least 1, the case where the partition spans the entire device (actually
    // no partitioning). A partition will contain one or more minimum partition
    // units (where a minimum partition unit is defined by
    // ciscoFlashDeviceMinPartitionSize). The type is interface{} with range:
    // 0..4294967295.
    CiscoFlashDevicePartitions interface{}

    // Total number of chips within the Flash device. The purpose of this object
    // is to provide information upfront to a management station on how much chip
    // info to expect and possibly help double check the chip index against an
    // upper limit when randomly retrieving chip info for a partition. The type is
    // interface{} with range: 0..64.
    CiscoFlashDeviceChipCount interface{}

    // Flash device name. This name is used to refer to the device within the
    // system. Flash operations get directed to a device based on this name. The
    // system has a concept of a default device. This would be the primary or most
    // used device in case of multiple devices. The system directs an operation to
    // the default device whenever a device name is not specified. The device name
    // is therefore mandatory except when the operation is being done on the
    // default device, or, the system supports only a single Flash device. The
    // device name will always be available for a removable device, even when the
    // device has been removed. The type is string with length: 0..16.
    CiscoFlashDeviceName interface{}

    // Description of a Flash device. The description is meant to explain what the
    // Flash device and its purpose is. Current values are:   System flash - for
    // the primary Flash used to store full                  system images.   Boot
    // flash   - for the secondary Flash used to store                  bootstrap
    // images. The ciscoFlashDeviceDescr, ciscoFlashDeviceController (if
    // applicable), and ciscoFlashPhyEntIndex objects are expected to collectively
    // give all information about a Flash device. The device description will
    // always be available for a removable device, even when the device has been
    // removed. The type is string with length: 0..64.
    CiscoFlashDeviceDescr interface{}

    // Flash device controller. The h/w card that actually controls Flash
    // read/write/erase. Relevant for the AGS+ systems where Flash may be
    // controlled by the MC+, STR or the ENVM cards, cards that may not actually
    // contain the Flash chips. For systems that have removable PCMCIA flash cards
    // that are controlled by a PCMCIA controller chip, this object may contain a
    // description of that controller chip. Where irrelevant (Flash is a direct
    // memory mapped device accessed directly by the main processor), this object
    // will have an empty (NULL) string. The type is string with length: 0..64.
    CiscoFlashDeviceController interface{}

    // This object will point to an instance of a card entry in the cardTable. The
    // card entry will give details about the card on which the Flash device is
    // actually located. For most systems, this is usually the main processor
    // board. On the AGS+ systems, Flash is located on a separate multibus card
    // such as the MC. This object will therefore be used to essentially index
    // into cardTable to retrieve details about the card such as cardDescr,
    // cardSlotNumber, etc. The type is string with pattern:
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
    CiscoFlashDeviceCard interface{}

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
    // The type is CiscoFlashDeviceProgrammingJumper.
    CiscoFlashDeviceProgrammingJumper interface{}

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
    CiscoFlashDeviceInitTime interface{}

    // Whether Flash device is removable. Generally, only PCMCIA Flash cards will
    // be treated as removable. Socketed Flash chips and Flash SIMM modules will
    // not be treated as removable. Simply put, only those Flash devices that can
    // be inserted or removed without opening the hardware casing will be
    // considered removable. Further, removable Flash devices are expected to have
    // the necessary hardware support -   1. on-line removal and insertion   2.
    // interrupt generation on removal or insertion. The type is bool.
    CiscoFlashDeviceRemovable interface{}

    // This object indicates the physical entity index of a physical entity in
    // entPhysicalTable which the flash device actually located. The type is
    // interface{} with range: 0..2147483647.
    CiscoFlashPhyEntIndex interface{}

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
    CiscoFlashDeviceNameExtended interface{}

    // Total size of the Flash device. For a removable device, the size will be
    // zero if the device has been removed.  This object is a 64-bit version of
    // ciscoFlashDeviceSize. The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    CiscoFlashDeviceSizeExtended interface{}

    // This object provides the minimum partition size supported for this device.
    // This object is a 64-bit version of  ciscoFlashDeviceMinPatitionSize. The
    // type is interface{} with range: 0..18446744073709551615.
    CiscoFlashDeviceMinPartitionSizeExtended interface{}
}

func (ciscoFlashDeviceEntry *CISCOFLASHMIB_CiscoFlashDeviceTable_CiscoFlashDeviceEntry) GetEntityData() *types.CommonEntityData {
    ciscoFlashDeviceEntry.EntityData.YFilter = ciscoFlashDeviceEntry.YFilter
    ciscoFlashDeviceEntry.EntityData.YangName = "ciscoFlashDeviceEntry"
    ciscoFlashDeviceEntry.EntityData.BundleName = "cisco_ios_xe"
    ciscoFlashDeviceEntry.EntityData.ParentYangName = "ciscoFlashDeviceTable"
    ciscoFlashDeviceEntry.EntityData.SegmentPath = "ciscoFlashDeviceEntry" + types.AddKeyToken(ciscoFlashDeviceEntry.CiscoFlashDeviceIndex, "ciscoFlashDeviceIndex")
    ciscoFlashDeviceEntry.EntityData.AbsolutePath = "CISCO-FLASH-MIB:CISCO-FLASH-MIB/ciscoFlashDeviceTable/" + ciscoFlashDeviceEntry.EntityData.SegmentPath
    ciscoFlashDeviceEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoFlashDeviceEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoFlashDeviceEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoFlashDeviceEntry.EntityData.Children = types.NewOrderedMap()
    ciscoFlashDeviceEntry.EntityData.Leafs = types.NewOrderedMap()
    ciscoFlashDeviceEntry.EntityData.Leafs.Append("ciscoFlashDeviceIndex", types.YLeaf{"CiscoFlashDeviceIndex", ciscoFlashDeviceEntry.CiscoFlashDeviceIndex})
    ciscoFlashDeviceEntry.EntityData.Leafs.Append("ciscoFlashDeviceSize", types.YLeaf{"CiscoFlashDeviceSize", ciscoFlashDeviceEntry.CiscoFlashDeviceSize})
    ciscoFlashDeviceEntry.EntityData.Leafs.Append("ciscoFlashDeviceMinPartitionSize", types.YLeaf{"CiscoFlashDeviceMinPartitionSize", ciscoFlashDeviceEntry.CiscoFlashDeviceMinPartitionSize})
    ciscoFlashDeviceEntry.EntityData.Leafs.Append("ciscoFlashDeviceMaxPartitions", types.YLeaf{"CiscoFlashDeviceMaxPartitions", ciscoFlashDeviceEntry.CiscoFlashDeviceMaxPartitions})
    ciscoFlashDeviceEntry.EntityData.Leafs.Append("ciscoFlashDevicePartitions", types.YLeaf{"CiscoFlashDevicePartitions", ciscoFlashDeviceEntry.CiscoFlashDevicePartitions})
    ciscoFlashDeviceEntry.EntityData.Leafs.Append("ciscoFlashDeviceChipCount", types.YLeaf{"CiscoFlashDeviceChipCount", ciscoFlashDeviceEntry.CiscoFlashDeviceChipCount})
    ciscoFlashDeviceEntry.EntityData.Leafs.Append("ciscoFlashDeviceName", types.YLeaf{"CiscoFlashDeviceName", ciscoFlashDeviceEntry.CiscoFlashDeviceName})
    ciscoFlashDeviceEntry.EntityData.Leafs.Append("ciscoFlashDeviceDescr", types.YLeaf{"CiscoFlashDeviceDescr", ciscoFlashDeviceEntry.CiscoFlashDeviceDescr})
    ciscoFlashDeviceEntry.EntityData.Leafs.Append("ciscoFlashDeviceController", types.YLeaf{"CiscoFlashDeviceController", ciscoFlashDeviceEntry.CiscoFlashDeviceController})
    ciscoFlashDeviceEntry.EntityData.Leafs.Append("ciscoFlashDeviceCard", types.YLeaf{"CiscoFlashDeviceCard", ciscoFlashDeviceEntry.CiscoFlashDeviceCard})
    ciscoFlashDeviceEntry.EntityData.Leafs.Append("ciscoFlashDeviceProgrammingJumper", types.YLeaf{"CiscoFlashDeviceProgrammingJumper", ciscoFlashDeviceEntry.CiscoFlashDeviceProgrammingJumper})
    ciscoFlashDeviceEntry.EntityData.Leafs.Append("ciscoFlashDeviceInitTime", types.YLeaf{"CiscoFlashDeviceInitTime", ciscoFlashDeviceEntry.CiscoFlashDeviceInitTime})
    ciscoFlashDeviceEntry.EntityData.Leafs.Append("ciscoFlashDeviceRemovable", types.YLeaf{"CiscoFlashDeviceRemovable", ciscoFlashDeviceEntry.CiscoFlashDeviceRemovable})
    ciscoFlashDeviceEntry.EntityData.Leafs.Append("ciscoFlashPhyEntIndex", types.YLeaf{"CiscoFlashPhyEntIndex", ciscoFlashDeviceEntry.CiscoFlashPhyEntIndex})
    ciscoFlashDeviceEntry.EntityData.Leafs.Append("ciscoFlashDeviceNameExtended", types.YLeaf{"CiscoFlashDeviceNameExtended", ciscoFlashDeviceEntry.CiscoFlashDeviceNameExtended})
    ciscoFlashDeviceEntry.EntityData.Leafs.Append("ciscoFlashDeviceSizeExtended", types.YLeaf{"CiscoFlashDeviceSizeExtended", ciscoFlashDeviceEntry.CiscoFlashDeviceSizeExtended})
    ciscoFlashDeviceEntry.EntityData.Leafs.Append("ciscoFlashDeviceMinPartitionSizeExtended", types.YLeaf{"CiscoFlashDeviceMinPartitionSizeExtended", ciscoFlashDeviceEntry.CiscoFlashDeviceMinPartitionSizeExtended})

    ciscoFlashDeviceEntry.EntityData.YListKeys = []string {"CiscoFlashDeviceIndex"}

    return &(ciscoFlashDeviceEntry.EntityData)
}

// CISCOFLASHMIB_CiscoFlashDeviceTable_CiscoFlashDeviceEntry_CiscoFlashDeviceProgrammingJumper represents for the readOnly state.
type CISCOFLASHMIB_CiscoFlashDeviceTable_CiscoFlashDeviceEntry_CiscoFlashDeviceProgrammingJumper string

const (
    CISCOFLASHMIB_CiscoFlashDeviceTable_CiscoFlashDeviceEntry_CiscoFlashDeviceProgrammingJumper_installed CISCOFLASHMIB_CiscoFlashDeviceTable_CiscoFlashDeviceEntry_CiscoFlashDeviceProgrammingJumper = "installed"

    CISCOFLASHMIB_CiscoFlashDeviceTable_CiscoFlashDeviceEntry_CiscoFlashDeviceProgrammingJumper_notInstalled CISCOFLASHMIB_CiscoFlashDeviceTable_CiscoFlashDeviceEntry_CiscoFlashDeviceProgrammingJumper = "notInstalled"

    CISCOFLASHMIB_CiscoFlashDeviceTable_CiscoFlashDeviceEntry_CiscoFlashDeviceProgrammingJumper_unknown CISCOFLASHMIB_CiscoFlashDeviceTable_CiscoFlashDeviceEntry_CiscoFlashDeviceProgrammingJumper = "unknown"
)

// CISCOFLASHMIB_CiscoFlashChipTable
// Table of Flash device chip properties for each
// initialized Flash device.
// This table is meant primarily for aiding error
// diagnosis.
type CISCOFLASHMIB_CiscoFlashChipTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table of chip info for each flash device initialized in the
    // system. An entry is indexed by two objects - the device index and the chip
    // index within that device. The type is slice of
    // CISCOFLASHMIB_CiscoFlashChipTable_CiscoFlashChipEntry.
    CiscoFlashChipEntry []*CISCOFLASHMIB_CiscoFlashChipTable_CiscoFlashChipEntry
}

func (ciscoFlashChipTable *CISCOFLASHMIB_CiscoFlashChipTable) GetEntityData() *types.CommonEntityData {
    ciscoFlashChipTable.EntityData.YFilter = ciscoFlashChipTable.YFilter
    ciscoFlashChipTable.EntityData.YangName = "ciscoFlashChipTable"
    ciscoFlashChipTable.EntityData.BundleName = "cisco_ios_xe"
    ciscoFlashChipTable.EntityData.ParentYangName = "CISCO-FLASH-MIB"
    ciscoFlashChipTable.EntityData.SegmentPath = "ciscoFlashChipTable"
    ciscoFlashChipTable.EntityData.AbsolutePath = "CISCO-FLASH-MIB:CISCO-FLASH-MIB/" + ciscoFlashChipTable.EntityData.SegmentPath
    ciscoFlashChipTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoFlashChipTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoFlashChipTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoFlashChipTable.EntityData.Children = types.NewOrderedMap()
    ciscoFlashChipTable.EntityData.Children.Append("ciscoFlashChipEntry", types.YChild{"CiscoFlashChipEntry", nil})
    for i := range ciscoFlashChipTable.CiscoFlashChipEntry {
        ciscoFlashChipTable.EntityData.Children.Append(types.GetSegmentPath(ciscoFlashChipTable.CiscoFlashChipEntry[i]), types.YChild{"CiscoFlashChipEntry", ciscoFlashChipTable.CiscoFlashChipEntry[i]})
    }
    ciscoFlashChipTable.EntityData.Leafs = types.NewOrderedMap()

    ciscoFlashChipTable.EntityData.YListKeys = []string {}

    return &(ciscoFlashChipTable.EntityData)
}

// CISCOFLASHMIB_CiscoFlashChipTable_CiscoFlashChipEntry
// An entry in the table of chip info for each
// flash device initialized in the system.
// An entry is indexed by two objects - the
// device index and the chip index within that
// device.
type CISCOFLASHMIB_CiscoFlashChipTable_CiscoFlashChipEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_flash_mib.CISCOFLASHMIB_CiscoFlashDeviceTable_CiscoFlashDeviceEntry_CiscoFlashDeviceIndex
    CiscoFlashDeviceIndex interface{}

    // This attribute is a key. Chip sequence number within selected flash device.
    // Used to index within chip info table. Value starts from 1 and should not be
    // greater than ciscoFlashDeviceChipCount for that device. When retrieving
    // chip information for chips within a partition, the sequence number should
    // lie between ciscoFlashPartitionStartChip & ciscoFlashPartitionEndChip (both
    // inclusive). The type is interface{} with range: 1..64.
    CiscoFlashChipIndex interface{}

    // Manufacturer and device code for a chip. Lower byte will contain the device
    // code. Upper byte will contain the manufacturer code. If a chip code is
    // unknown because it could not be queried out of the chip, the value of this
    // object will be 00:00. Since programming algorithms differ from chip type to
    // chip type, this chip code should be used to determine which algorithms to
    // use (and thereby whether the chip is supported in the first place). The
    // type is string with length: 0..5.
    CiscoFlashChipCode interface{}

    // Flash chip name corresponding to the chip code. The name will contain the
    // manufacturer and the chip type. It will be of the form :   Intel 27F008SA.
    // In the case where a chip code is unknown, this object will be an empty
    // (NULL) string. In the case where the chip code is known but the chip is not
    // supported by the system, this object will be an empty (NULL) string. A
    // management station is therefore expected to use the chip code and the chip
    // description in conjunction to provide additional information whenever the
    // ciscoFlashPartitionStatus object has the readOnly(1) value. The type is
    // string with length: 0..32.
    CiscoFlashChipDescr interface{}

    // This object will provide a cumulative count (since last system boot up or
    // initialization) of the number of write retries that were done in the chip.
    // If no writes have been done to Flash, the count will be zero. Typically, a
    // maximum of 25 retries are done on a single location before flagging a write
    // error. A management station is expected to get this object for each chip in
    // a partition after a write failure in that partition. To keep a track of
    // retries for a given write operation, the management station would have to
    // retrieve the values for the concerned chips before and after any write
    // operation. The type is interface{} with range: 0..4294967295.
    CiscoFlashChipWriteRetries interface{}

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
    CiscoFlashChipEraseRetries interface{}

    // The maximum number of write retries done at any single location before
    // declaring a write failure. The type is interface{} with range:
    // 0..4294967295.
    CiscoFlashChipMaxWriteRetries interface{}

    // The maximum number of erase retries done within an erase sector before
    // declaring an erase failure. The type is interface{} with range:
    // 0..4294967295.
    CiscoFlashChipMaxEraseRetries interface{}
}

func (ciscoFlashChipEntry *CISCOFLASHMIB_CiscoFlashChipTable_CiscoFlashChipEntry) GetEntityData() *types.CommonEntityData {
    ciscoFlashChipEntry.EntityData.YFilter = ciscoFlashChipEntry.YFilter
    ciscoFlashChipEntry.EntityData.YangName = "ciscoFlashChipEntry"
    ciscoFlashChipEntry.EntityData.BundleName = "cisco_ios_xe"
    ciscoFlashChipEntry.EntityData.ParentYangName = "ciscoFlashChipTable"
    ciscoFlashChipEntry.EntityData.SegmentPath = "ciscoFlashChipEntry" + types.AddKeyToken(ciscoFlashChipEntry.CiscoFlashDeviceIndex, "ciscoFlashDeviceIndex") + types.AddKeyToken(ciscoFlashChipEntry.CiscoFlashChipIndex, "ciscoFlashChipIndex")
    ciscoFlashChipEntry.EntityData.AbsolutePath = "CISCO-FLASH-MIB:CISCO-FLASH-MIB/ciscoFlashChipTable/" + ciscoFlashChipEntry.EntityData.SegmentPath
    ciscoFlashChipEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoFlashChipEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoFlashChipEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoFlashChipEntry.EntityData.Children = types.NewOrderedMap()
    ciscoFlashChipEntry.EntityData.Leafs = types.NewOrderedMap()
    ciscoFlashChipEntry.EntityData.Leafs.Append("ciscoFlashDeviceIndex", types.YLeaf{"CiscoFlashDeviceIndex", ciscoFlashChipEntry.CiscoFlashDeviceIndex})
    ciscoFlashChipEntry.EntityData.Leafs.Append("ciscoFlashChipIndex", types.YLeaf{"CiscoFlashChipIndex", ciscoFlashChipEntry.CiscoFlashChipIndex})
    ciscoFlashChipEntry.EntityData.Leafs.Append("ciscoFlashChipCode", types.YLeaf{"CiscoFlashChipCode", ciscoFlashChipEntry.CiscoFlashChipCode})
    ciscoFlashChipEntry.EntityData.Leafs.Append("ciscoFlashChipDescr", types.YLeaf{"CiscoFlashChipDescr", ciscoFlashChipEntry.CiscoFlashChipDescr})
    ciscoFlashChipEntry.EntityData.Leafs.Append("ciscoFlashChipWriteRetries", types.YLeaf{"CiscoFlashChipWriteRetries", ciscoFlashChipEntry.CiscoFlashChipWriteRetries})
    ciscoFlashChipEntry.EntityData.Leafs.Append("ciscoFlashChipEraseRetries", types.YLeaf{"CiscoFlashChipEraseRetries", ciscoFlashChipEntry.CiscoFlashChipEraseRetries})
    ciscoFlashChipEntry.EntityData.Leafs.Append("ciscoFlashChipMaxWriteRetries", types.YLeaf{"CiscoFlashChipMaxWriteRetries", ciscoFlashChipEntry.CiscoFlashChipMaxWriteRetries})
    ciscoFlashChipEntry.EntityData.Leafs.Append("ciscoFlashChipMaxEraseRetries", types.YLeaf{"CiscoFlashChipMaxEraseRetries", ciscoFlashChipEntry.CiscoFlashChipMaxEraseRetries})

    ciscoFlashChipEntry.EntityData.YListKeys = []string {"CiscoFlashDeviceIndex", "CiscoFlashChipIndex"}

    return &(ciscoFlashChipEntry.EntityData)
}

// CISCOFLASHMIB_CiscoFlashPartitionTable
// Table of flash device partition properties for each
// initialized flash partition. Whenever there is no
// explicit partitioning done, a single partition spanning
// the entire device will be assumed to exist. There will
// therefore always be atleast one partition on a device.
type CISCOFLASHMIB_CiscoFlashPartitionTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table of flash partition properties for each initialized
    // flash partition. Each entry will be indexed by a device number and a
    // partition number within the device. The type is slice of
    // CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry.
    CiscoFlashPartitionEntry []*CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry
}

func (ciscoFlashPartitionTable *CISCOFLASHMIB_CiscoFlashPartitionTable) GetEntityData() *types.CommonEntityData {
    ciscoFlashPartitionTable.EntityData.YFilter = ciscoFlashPartitionTable.YFilter
    ciscoFlashPartitionTable.EntityData.YangName = "ciscoFlashPartitionTable"
    ciscoFlashPartitionTable.EntityData.BundleName = "cisco_ios_xe"
    ciscoFlashPartitionTable.EntityData.ParentYangName = "CISCO-FLASH-MIB"
    ciscoFlashPartitionTable.EntityData.SegmentPath = "ciscoFlashPartitionTable"
    ciscoFlashPartitionTable.EntityData.AbsolutePath = "CISCO-FLASH-MIB:CISCO-FLASH-MIB/" + ciscoFlashPartitionTable.EntityData.SegmentPath
    ciscoFlashPartitionTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoFlashPartitionTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoFlashPartitionTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoFlashPartitionTable.EntityData.Children = types.NewOrderedMap()
    ciscoFlashPartitionTable.EntityData.Children.Append("ciscoFlashPartitionEntry", types.YChild{"CiscoFlashPartitionEntry", nil})
    for i := range ciscoFlashPartitionTable.CiscoFlashPartitionEntry {
        ciscoFlashPartitionTable.EntityData.Children.Append(types.GetSegmentPath(ciscoFlashPartitionTable.CiscoFlashPartitionEntry[i]), types.YChild{"CiscoFlashPartitionEntry", ciscoFlashPartitionTable.CiscoFlashPartitionEntry[i]})
    }
    ciscoFlashPartitionTable.EntityData.Leafs = types.NewOrderedMap()

    ciscoFlashPartitionTable.EntityData.YListKeys = []string {}

    return &(ciscoFlashPartitionTable.EntityData)
}

// CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry
// An entry in the table of flash partition properties
// for each initialized flash partition. Each entry
// will be indexed by a device number and a partition
// number within the device.
type CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_flash_mib.CISCOFLASHMIB_CiscoFlashDeviceTable_CiscoFlashDeviceEntry_CiscoFlashDeviceIndex
    CiscoFlashDeviceIndex interface{}

    // This attribute is a key. Flash partition sequence number used to index
    // within table of initialized flash partitions. The type is interface{} with
    // range: 1..4294967295.
    CiscoFlashPartitionIndex interface{}

    // Chip sequence number of first chip in partition. Used as an index into the
    // chip table. The type is interface{} with range: 1..64.
    CiscoFlashPartitionStartChip interface{}

    // Chip sequence number of last chip in partition. Used as an index into the
    // chip table. The type is interface{} with range: 1..64.
    CiscoFlashPartitionEndChip interface{}

    // Flash partition size. It should be an integral multiple of
    // ciscoFlashDeviceMinPartitionSize. If there is a single partition, this size
    // will be equal to ciscoFlashDeviceSize.  If the size of the flash partition
    // is greater than the maximum value reportable by this object then this
    // object should report its maximum value(4,294,967,295) and
    // ciscoFlashPartitionSizeExtended must be used to report the flash
    // partition's size. The type is interface{} with range: 1..4294967295. Units
    // are bytes.
    CiscoFlashPartitionSize interface{}

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
    CiscoFlashPartitionFreeSpace interface{}

    // Count of all files in a flash partition. Both good and bad (deleted or
    // invalid checksum) files will be included in this count. The type is
    // interface{} with range: 0..4294967295.
    CiscoFlashPartitionFileCount interface{}

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
    // CiscoFlashPartitionChecksumAlgorithm.
    CiscoFlashPartitionChecksumAlgorithm interface{}

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
    // CiscoFlashPartitionStatus.
    CiscoFlashPartitionStatus interface{}

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
    // CiscoFlashPartitionUpgradeMethod.
    CiscoFlashPartitionUpgradeMethod interface{}

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
    CiscoFlashPartitionName interface{}

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
    CiscoFlashPartitionNeedErasure interface{}

    // Maximum file name length supported by the file system. Max file name length
    // will depend on the file system implemented. Today, all file systems support
    // a max length of at least 48 bytes. A management entity must use this object
    // when prompting a user for, or deriving the Flash file name length. The type
    // is interface{} with range: 1..256.
    CiscoFlashPartitionFileNameLength interface{}

    // Flash partition size. It should be an integral multiple of
    // ciscoFlashDeviceMinPartitionSize. If there is a single partition, this size
    // will be equal to ciscoFlashDeviceSize.  This object is a 64-bit version of
    // ciscoFlashPartitionSize. The type is interface{} with range:
    // 0..18446744073709551615. Units are bytes.
    CiscoFlashPartitionSizeExtended interface{}

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
    CiscoFlashPartitionFreeSpaceExtended interface{}

    // This object specifies the minimum threshold value in percentage of free
    // space for each partition. If the free space available goes below this
    // threshold value and if ciscoFlashPartionLowSpaceNotifEnable is set to true,
    // ciscoFlashPartitionLowSpaceNotif will be generated. When the available free
    // space comes back to the threshold value
    // ciscoFlashPartionLowSpaceRecoveryNotif will be generated. The type is
    // interface{} with range: 0..100.
    CiscoFlashPartitionLowSpaceNotifThreshold interface{}
}

func (ciscoFlashPartitionEntry *CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry) GetEntityData() *types.CommonEntityData {
    ciscoFlashPartitionEntry.EntityData.YFilter = ciscoFlashPartitionEntry.YFilter
    ciscoFlashPartitionEntry.EntityData.YangName = "ciscoFlashPartitionEntry"
    ciscoFlashPartitionEntry.EntityData.BundleName = "cisco_ios_xe"
    ciscoFlashPartitionEntry.EntityData.ParentYangName = "ciscoFlashPartitionTable"
    ciscoFlashPartitionEntry.EntityData.SegmentPath = "ciscoFlashPartitionEntry" + types.AddKeyToken(ciscoFlashPartitionEntry.CiscoFlashDeviceIndex, "ciscoFlashDeviceIndex") + types.AddKeyToken(ciscoFlashPartitionEntry.CiscoFlashPartitionIndex, "ciscoFlashPartitionIndex")
    ciscoFlashPartitionEntry.EntityData.AbsolutePath = "CISCO-FLASH-MIB:CISCO-FLASH-MIB/ciscoFlashPartitionTable/" + ciscoFlashPartitionEntry.EntityData.SegmentPath
    ciscoFlashPartitionEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoFlashPartitionEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoFlashPartitionEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoFlashPartitionEntry.EntityData.Children = types.NewOrderedMap()
    ciscoFlashPartitionEntry.EntityData.Leafs = types.NewOrderedMap()
    ciscoFlashPartitionEntry.EntityData.Leafs.Append("ciscoFlashDeviceIndex", types.YLeaf{"CiscoFlashDeviceIndex", ciscoFlashPartitionEntry.CiscoFlashDeviceIndex})
    ciscoFlashPartitionEntry.EntityData.Leafs.Append("ciscoFlashPartitionIndex", types.YLeaf{"CiscoFlashPartitionIndex", ciscoFlashPartitionEntry.CiscoFlashPartitionIndex})
    ciscoFlashPartitionEntry.EntityData.Leafs.Append("ciscoFlashPartitionStartChip", types.YLeaf{"CiscoFlashPartitionStartChip", ciscoFlashPartitionEntry.CiscoFlashPartitionStartChip})
    ciscoFlashPartitionEntry.EntityData.Leafs.Append("ciscoFlashPartitionEndChip", types.YLeaf{"CiscoFlashPartitionEndChip", ciscoFlashPartitionEntry.CiscoFlashPartitionEndChip})
    ciscoFlashPartitionEntry.EntityData.Leafs.Append("ciscoFlashPartitionSize", types.YLeaf{"CiscoFlashPartitionSize", ciscoFlashPartitionEntry.CiscoFlashPartitionSize})
    ciscoFlashPartitionEntry.EntityData.Leafs.Append("ciscoFlashPartitionFreeSpace", types.YLeaf{"CiscoFlashPartitionFreeSpace", ciscoFlashPartitionEntry.CiscoFlashPartitionFreeSpace})
    ciscoFlashPartitionEntry.EntityData.Leafs.Append("ciscoFlashPartitionFileCount", types.YLeaf{"CiscoFlashPartitionFileCount", ciscoFlashPartitionEntry.CiscoFlashPartitionFileCount})
    ciscoFlashPartitionEntry.EntityData.Leafs.Append("ciscoFlashPartitionChecksumAlgorithm", types.YLeaf{"CiscoFlashPartitionChecksumAlgorithm", ciscoFlashPartitionEntry.CiscoFlashPartitionChecksumAlgorithm})
    ciscoFlashPartitionEntry.EntityData.Leafs.Append("ciscoFlashPartitionStatus", types.YLeaf{"CiscoFlashPartitionStatus", ciscoFlashPartitionEntry.CiscoFlashPartitionStatus})
    ciscoFlashPartitionEntry.EntityData.Leafs.Append("ciscoFlashPartitionUpgradeMethod", types.YLeaf{"CiscoFlashPartitionUpgradeMethod", ciscoFlashPartitionEntry.CiscoFlashPartitionUpgradeMethod})
    ciscoFlashPartitionEntry.EntityData.Leafs.Append("ciscoFlashPartitionName", types.YLeaf{"CiscoFlashPartitionName", ciscoFlashPartitionEntry.CiscoFlashPartitionName})
    ciscoFlashPartitionEntry.EntityData.Leafs.Append("ciscoFlashPartitionNeedErasure", types.YLeaf{"CiscoFlashPartitionNeedErasure", ciscoFlashPartitionEntry.CiscoFlashPartitionNeedErasure})
    ciscoFlashPartitionEntry.EntityData.Leafs.Append("ciscoFlashPartitionFileNameLength", types.YLeaf{"CiscoFlashPartitionFileNameLength", ciscoFlashPartitionEntry.CiscoFlashPartitionFileNameLength})
    ciscoFlashPartitionEntry.EntityData.Leafs.Append("ciscoFlashPartitionSizeExtended", types.YLeaf{"CiscoFlashPartitionSizeExtended", ciscoFlashPartitionEntry.CiscoFlashPartitionSizeExtended})
    ciscoFlashPartitionEntry.EntityData.Leafs.Append("ciscoFlashPartitionFreeSpaceExtended", types.YLeaf{"CiscoFlashPartitionFreeSpaceExtended", ciscoFlashPartitionEntry.CiscoFlashPartitionFreeSpaceExtended})
    ciscoFlashPartitionEntry.EntityData.Leafs.Append("ciscoFlashPartitionLowSpaceNotifThreshold", types.YLeaf{"CiscoFlashPartitionLowSpaceNotifThreshold", ciscoFlashPartitionEntry.CiscoFlashPartitionLowSpaceNotifThreshold})

    ciscoFlashPartitionEntry.EntityData.YListKeys = []string {"CiscoFlashDeviceIndex", "CiscoFlashPartitionIndex"}

    return &(ciscoFlashPartitionEntry.EntityData)
}

// CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry_CiscoFlashPartitionChecksumAlgorithm represents values will be added as necessary.
type CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry_CiscoFlashPartitionChecksumAlgorithm string

const (
    CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry_CiscoFlashPartitionChecksumAlgorithm_simpleChecksum CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry_CiscoFlashPartitionChecksumAlgorithm = "simpleChecksum"

    CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry_CiscoFlashPartitionChecksumAlgorithm_undefined CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry_CiscoFlashPartitionChecksumAlgorithm = "undefined"

    CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry_CiscoFlashPartitionChecksumAlgorithm_simpleCRC CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry_CiscoFlashPartitionChecksumAlgorithm = "simpleCRC"
)

// CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry_CiscoFlashPartitionStatus represents * readWrite if partition is programmable.
type CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry_CiscoFlashPartitionStatus string

const (
    CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry_CiscoFlashPartitionStatus_readOnly CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry_CiscoFlashPartitionStatus = "readOnly"

    CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry_CiscoFlashPartitionStatus_runFromFlash CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry_CiscoFlashPartitionStatus = "runFromFlash"

    CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry_CiscoFlashPartitionStatus_readWrite CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry_CiscoFlashPartitionStatus = "readWrite"
)

// CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry_CiscoFlashPartitionUpgradeMethod represents direct       -  will be done directly by this image.
type CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry_CiscoFlashPartitionUpgradeMethod string

const (
    CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry_CiscoFlashPartitionUpgradeMethod_unknown CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry_CiscoFlashPartitionUpgradeMethod = "unknown"

    CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry_CiscoFlashPartitionUpgradeMethod_rxbootFLH CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry_CiscoFlashPartitionUpgradeMethod = "rxbootFLH"

    CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry_CiscoFlashPartitionUpgradeMethod_direct CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry_CiscoFlashPartitionUpgradeMethod = "direct"
)

// CISCOFLASHMIB_CiscoFlashFileTable
// Table of information for files in a Flash partition.
type CISCOFLASHMIB_CiscoFlashFileTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table of Flash file properties for each initialized Flash
    // partition. Each entry represents a file and gives details about the file.
    // An entry is indexed using the device number, partition number within the
    // device, and file number within the partition. The type is slice of
    // CISCOFLASHMIB_CiscoFlashFileTable_CiscoFlashFileEntry.
    CiscoFlashFileEntry []*CISCOFLASHMIB_CiscoFlashFileTable_CiscoFlashFileEntry
}

func (ciscoFlashFileTable *CISCOFLASHMIB_CiscoFlashFileTable) GetEntityData() *types.CommonEntityData {
    ciscoFlashFileTable.EntityData.YFilter = ciscoFlashFileTable.YFilter
    ciscoFlashFileTable.EntityData.YangName = "ciscoFlashFileTable"
    ciscoFlashFileTable.EntityData.BundleName = "cisco_ios_xe"
    ciscoFlashFileTable.EntityData.ParentYangName = "CISCO-FLASH-MIB"
    ciscoFlashFileTable.EntityData.SegmentPath = "ciscoFlashFileTable"
    ciscoFlashFileTable.EntityData.AbsolutePath = "CISCO-FLASH-MIB:CISCO-FLASH-MIB/" + ciscoFlashFileTable.EntityData.SegmentPath
    ciscoFlashFileTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoFlashFileTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoFlashFileTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoFlashFileTable.EntityData.Children = types.NewOrderedMap()
    ciscoFlashFileTable.EntityData.Children.Append("ciscoFlashFileEntry", types.YChild{"CiscoFlashFileEntry", nil})
    for i := range ciscoFlashFileTable.CiscoFlashFileEntry {
        ciscoFlashFileTable.EntityData.Children.Append(types.GetSegmentPath(ciscoFlashFileTable.CiscoFlashFileEntry[i]), types.YChild{"CiscoFlashFileEntry", ciscoFlashFileTable.CiscoFlashFileEntry[i]})
    }
    ciscoFlashFileTable.EntityData.Leafs = types.NewOrderedMap()

    ciscoFlashFileTable.EntityData.YListKeys = []string {}

    return &(ciscoFlashFileTable.EntityData)
}

// CISCOFLASHMIB_CiscoFlashFileTable_CiscoFlashFileEntry
// An entry in the table of Flash file properties
// for each initialized Flash partition. Each entry
// represents a file and gives details about the file.
// An entry is indexed using the device number,
// partition number within the device, and file
// number within the partition.
type CISCOFLASHMIB_CiscoFlashFileTable_CiscoFlashFileEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_flash_mib.CISCOFLASHMIB_CiscoFlashDeviceTable_CiscoFlashDeviceEntry_CiscoFlashDeviceIndex
    CiscoFlashDeviceIndex interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_flash_mib.CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry_CiscoFlashPartitionIndex
    CiscoFlashPartitionIndex interface{}

    // This attribute is a key. Flash file sequence number used to index within a
    // Flash partition directory table. The type is interface{} with range:
    // 1..4294967295.
    CiscoFlashFileIndex interface{}

    // Size of the file in bytes. Note that this size does not include the size of
    // the filesystem file header. File size will always be non-zero. The type is
    // interface{} with range: 0..4294967295. Units are bytes.
    CiscoFlashFileSize interface{}

    // File checksum stored in the file header. This checksum is computed and
    // stored when the file is written into Flash. It serves to validate the data
    // written into Flash. Whereas the system will generate and store the checksum
    // internally in hexadecimal form, this object will provide the checksum in a
    // string form. The checksum will be available for all valid and
    // invalid-checksum files. The type is string.
    CiscoFlashFileChecksum interface{}

    // Status of a file. A file could be explicitly deleted if the file system
    // supports such a user command facility. Alternately, an existing good file
    // would be automatically deleted if another good file with the same name were
    // copied in. Note that deleted files continue to occupy prime Flash real
    // estate.  A file is marked as having an invalid checksum if any checksum
    // mismatch was detected while writing or reading the file. Incomplete files
    // (files truncated either because of lack of free space, or a network
    // download failure) are also written with a bad checksum and marked as
    // invalid. The type is CiscoFlashFileStatus.
    CiscoFlashFileStatus interface{}

    // Flash file name as specified by the user copying in the file. The name
    // should not include the colon (:) character as it is a special separator
    // character used to delineate the device name, partition name, and the file
    // name. The type is string with length: 1..255.
    CiscoFlashFileName interface{}

    // Type of the file. The type is FlashFileType.
    CiscoFlashFileType interface{}

    // The time at which this file was created. The type is string.
    CiscoFlashFileDate interface{}
}

func (ciscoFlashFileEntry *CISCOFLASHMIB_CiscoFlashFileTable_CiscoFlashFileEntry) GetEntityData() *types.CommonEntityData {
    ciscoFlashFileEntry.EntityData.YFilter = ciscoFlashFileEntry.YFilter
    ciscoFlashFileEntry.EntityData.YangName = "ciscoFlashFileEntry"
    ciscoFlashFileEntry.EntityData.BundleName = "cisco_ios_xe"
    ciscoFlashFileEntry.EntityData.ParentYangName = "ciscoFlashFileTable"
    ciscoFlashFileEntry.EntityData.SegmentPath = "ciscoFlashFileEntry" + types.AddKeyToken(ciscoFlashFileEntry.CiscoFlashDeviceIndex, "ciscoFlashDeviceIndex") + types.AddKeyToken(ciscoFlashFileEntry.CiscoFlashPartitionIndex, "ciscoFlashPartitionIndex") + types.AddKeyToken(ciscoFlashFileEntry.CiscoFlashFileIndex, "ciscoFlashFileIndex")
    ciscoFlashFileEntry.EntityData.AbsolutePath = "CISCO-FLASH-MIB:CISCO-FLASH-MIB/ciscoFlashFileTable/" + ciscoFlashFileEntry.EntityData.SegmentPath
    ciscoFlashFileEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoFlashFileEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoFlashFileEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoFlashFileEntry.EntityData.Children = types.NewOrderedMap()
    ciscoFlashFileEntry.EntityData.Leafs = types.NewOrderedMap()
    ciscoFlashFileEntry.EntityData.Leafs.Append("ciscoFlashDeviceIndex", types.YLeaf{"CiscoFlashDeviceIndex", ciscoFlashFileEntry.CiscoFlashDeviceIndex})
    ciscoFlashFileEntry.EntityData.Leafs.Append("ciscoFlashPartitionIndex", types.YLeaf{"CiscoFlashPartitionIndex", ciscoFlashFileEntry.CiscoFlashPartitionIndex})
    ciscoFlashFileEntry.EntityData.Leafs.Append("ciscoFlashFileIndex", types.YLeaf{"CiscoFlashFileIndex", ciscoFlashFileEntry.CiscoFlashFileIndex})
    ciscoFlashFileEntry.EntityData.Leafs.Append("ciscoFlashFileSize", types.YLeaf{"CiscoFlashFileSize", ciscoFlashFileEntry.CiscoFlashFileSize})
    ciscoFlashFileEntry.EntityData.Leafs.Append("ciscoFlashFileChecksum", types.YLeaf{"CiscoFlashFileChecksum", ciscoFlashFileEntry.CiscoFlashFileChecksum})
    ciscoFlashFileEntry.EntityData.Leafs.Append("ciscoFlashFileStatus", types.YLeaf{"CiscoFlashFileStatus", ciscoFlashFileEntry.CiscoFlashFileStatus})
    ciscoFlashFileEntry.EntityData.Leafs.Append("ciscoFlashFileName", types.YLeaf{"CiscoFlashFileName", ciscoFlashFileEntry.CiscoFlashFileName})
    ciscoFlashFileEntry.EntityData.Leafs.Append("ciscoFlashFileType", types.YLeaf{"CiscoFlashFileType", ciscoFlashFileEntry.CiscoFlashFileType})
    ciscoFlashFileEntry.EntityData.Leafs.Append("ciscoFlashFileDate", types.YLeaf{"CiscoFlashFileDate", ciscoFlashFileEntry.CiscoFlashFileDate})

    ciscoFlashFileEntry.EntityData.YListKeys = []string {"CiscoFlashDeviceIndex", "CiscoFlashPartitionIndex", "CiscoFlashFileIndex"}

    return &(ciscoFlashFileEntry.EntityData)
}

// CISCOFLASHMIB_CiscoFlashFileTable_CiscoFlashFileEntry_CiscoFlashFileStatus represents marked as invalid.
type CISCOFLASHMIB_CiscoFlashFileTable_CiscoFlashFileEntry_CiscoFlashFileStatus string

const (
    CISCOFLASHMIB_CiscoFlashFileTable_CiscoFlashFileEntry_CiscoFlashFileStatus_deleted CISCOFLASHMIB_CiscoFlashFileTable_CiscoFlashFileEntry_CiscoFlashFileStatus = "deleted"

    CISCOFLASHMIB_CiscoFlashFileTable_CiscoFlashFileEntry_CiscoFlashFileStatus_invalidChecksum CISCOFLASHMIB_CiscoFlashFileTable_CiscoFlashFileEntry_CiscoFlashFileStatus = "invalidChecksum"

    CISCOFLASHMIB_CiscoFlashFileTable_CiscoFlashFileEntry_CiscoFlashFileStatus_valid CISCOFLASHMIB_CiscoFlashFileTable_CiscoFlashFileEntry_CiscoFlashFileStatus = "valid"
)

// CISCOFLASHMIB_CiscoFlashFileByTypeTable
// Table of information for files on the manageable
// flash devices sorted by File Types.
type CISCOFLASHMIB_CiscoFlashFileByTypeTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table of Flash file properties for each initialized Flash
    // partition. Each entry represents a file sorted by file type.  This table
    // contains exactly the same set of rows as are contained in the
    // ciscoFlashFileTable but in a different order, i.e., ordered by  the type of
    // file, given by  ciscoFlashFileType; the device number, given by
    // ciscoFlashDeviceIndex; the partition number within the device, given by
    // ciscoFlashPartitionIndex; the file number within the partition, given by
    // ciscoFlashFileIndex. The type is slice of
    // CISCOFLASHMIB_CiscoFlashFileByTypeTable_CiscoFlashFileByTypeEntry.
    CiscoFlashFileByTypeEntry []*CISCOFLASHMIB_CiscoFlashFileByTypeTable_CiscoFlashFileByTypeEntry
}

func (ciscoFlashFileByTypeTable *CISCOFLASHMIB_CiscoFlashFileByTypeTable) GetEntityData() *types.CommonEntityData {
    ciscoFlashFileByTypeTable.EntityData.YFilter = ciscoFlashFileByTypeTable.YFilter
    ciscoFlashFileByTypeTable.EntityData.YangName = "ciscoFlashFileByTypeTable"
    ciscoFlashFileByTypeTable.EntityData.BundleName = "cisco_ios_xe"
    ciscoFlashFileByTypeTable.EntityData.ParentYangName = "CISCO-FLASH-MIB"
    ciscoFlashFileByTypeTable.EntityData.SegmentPath = "ciscoFlashFileByTypeTable"
    ciscoFlashFileByTypeTable.EntityData.AbsolutePath = "CISCO-FLASH-MIB:CISCO-FLASH-MIB/" + ciscoFlashFileByTypeTable.EntityData.SegmentPath
    ciscoFlashFileByTypeTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoFlashFileByTypeTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoFlashFileByTypeTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoFlashFileByTypeTable.EntityData.Children = types.NewOrderedMap()
    ciscoFlashFileByTypeTable.EntityData.Children.Append("ciscoFlashFileByTypeEntry", types.YChild{"CiscoFlashFileByTypeEntry", nil})
    for i := range ciscoFlashFileByTypeTable.CiscoFlashFileByTypeEntry {
        ciscoFlashFileByTypeTable.EntityData.Children.Append(types.GetSegmentPath(ciscoFlashFileByTypeTable.CiscoFlashFileByTypeEntry[i]), types.YChild{"CiscoFlashFileByTypeEntry", ciscoFlashFileByTypeTable.CiscoFlashFileByTypeEntry[i]})
    }
    ciscoFlashFileByTypeTable.EntityData.Leafs = types.NewOrderedMap()

    ciscoFlashFileByTypeTable.EntityData.YListKeys = []string {}

    return &(ciscoFlashFileByTypeTable.EntityData)
}

// CISCOFLASHMIB_CiscoFlashFileByTypeTable_CiscoFlashFileByTypeEntry
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
type CISCOFLASHMIB_CiscoFlashFileByTypeTable_CiscoFlashFileByTypeEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is FlashFileType. Refers to
    // cisco_flash_mib.CISCOFLASHMIB_CiscoFlashFileTable_CiscoFlashFileEntry_CiscoFlashFileType
    CiscoFlashFileType interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_flash_mib.CISCOFLASHMIB_CiscoFlashDeviceTable_CiscoFlashDeviceEntry_CiscoFlashDeviceIndex
    CiscoFlashDeviceIndex interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_flash_mib.CISCOFLASHMIB_CiscoFlashPartitionTable_CiscoFlashPartitionEntry_CiscoFlashPartitionIndex
    CiscoFlashPartitionIndex interface{}

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_flash_mib.CISCOFLASHMIB_CiscoFlashFileTable_CiscoFlashFileEntry_CiscoFlashFileIndex
    CiscoFlashFileIndex interface{}

    // This object represents exactly the same info as ciscoFlashFileSize object
    // in ciscoFlashFileTable. The type is interface{} with range: 0..4294967295.
    // Units are bytes.
    CiscoFlashFileByTypeSize interface{}

    // This object represents exactly the same info as ciscoFlashFileChecksum
    // object in ciscoFlashFileTable. The type is string.
    CiscoFlashFileByTypeChecksum interface{}

    // This object represents exactly the same info as ciscoFlashFileStatus object
    // in ciscoFlashFileTable. The type is CiscoFlashFileByTypeStatus.
    CiscoFlashFileByTypeStatus interface{}

    // This object represents exactly the same info as ciscoFlashFileName object
    // in ciscoFlashFileTable. The type is string with length: 1..255.
    CiscoFlashFileByTypeName interface{}

    // This object represents exactly the same info as ciscoFlashFileDate object
    // in ciscoFlashFileTable. The type is string.
    CiscoFlashFileByTypeDate interface{}
}

func (ciscoFlashFileByTypeEntry *CISCOFLASHMIB_CiscoFlashFileByTypeTable_CiscoFlashFileByTypeEntry) GetEntityData() *types.CommonEntityData {
    ciscoFlashFileByTypeEntry.EntityData.YFilter = ciscoFlashFileByTypeEntry.YFilter
    ciscoFlashFileByTypeEntry.EntityData.YangName = "ciscoFlashFileByTypeEntry"
    ciscoFlashFileByTypeEntry.EntityData.BundleName = "cisco_ios_xe"
    ciscoFlashFileByTypeEntry.EntityData.ParentYangName = "ciscoFlashFileByTypeTable"
    ciscoFlashFileByTypeEntry.EntityData.SegmentPath = "ciscoFlashFileByTypeEntry" + types.AddKeyToken(ciscoFlashFileByTypeEntry.CiscoFlashFileType, "ciscoFlashFileType") + types.AddKeyToken(ciscoFlashFileByTypeEntry.CiscoFlashDeviceIndex, "ciscoFlashDeviceIndex") + types.AddKeyToken(ciscoFlashFileByTypeEntry.CiscoFlashPartitionIndex, "ciscoFlashPartitionIndex") + types.AddKeyToken(ciscoFlashFileByTypeEntry.CiscoFlashFileIndex, "ciscoFlashFileIndex")
    ciscoFlashFileByTypeEntry.EntityData.AbsolutePath = "CISCO-FLASH-MIB:CISCO-FLASH-MIB/ciscoFlashFileByTypeTable/" + ciscoFlashFileByTypeEntry.EntityData.SegmentPath
    ciscoFlashFileByTypeEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoFlashFileByTypeEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoFlashFileByTypeEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoFlashFileByTypeEntry.EntityData.Children = types.NewOrderedMap()
    ciscoFlashFileByTypeEntry.EntityData.Leafs = types.NewOrderedMap()
    ciscoFlashFileByTypeEntry.EntityData.Leafs.Append("ciscoFlashFileType", types.YLeaf{"CiscoFlashFileType", ciscoFlashFileByTypeEntry.CiscoFlashFileType})
    ciscoFlashFileByTypeEntry.EntityData.Leafs.Append("ciscoFlashDeviceIndex", types.YLeaf{"CiscoFlashDeviceIndex", ciscoFlashFileByTypeEntry.CiscoFlashDeviceIndex})
    ciscoFlashFileByTypeEntry.EntityData.Leafs.Append("ciscoFlashPartitionIndex", types.YLeaf{"CiscoFlashPartitionIndex", ciscoFlashFileByTypeEntry.CiscoFlashPartitionIndex})
    ciscoFlashFileByTypeEntry.EntityData.Leafs.Append("ciscoFlashFileIndex", types.YLeaf{"CiscoFlashFileIndex", ciscoFlashFileByTypeEntry.CiscoFlashFileIndex})
    ciscoFlashFileByTypeEntry.EntityData.Leafs.Append("ciscoFlashFileByTypeSize", types.YLeaf{"CiscoFlashFileByTypeSize", ciscoFlashFileByTypeEntry.CiscoFlashFileByTypeSize})
    ciscoFlashFileByTypeEntry.EntityData.Leafs.Append("ciscoFlashFileByTypeChecksum", types.YLeaf{"CiscoFlashFileByTypeChecksum", ciscoFlashFileByTypeEntry.CiscoFlashFileByTypeChecksum})
    ciscoFlashFileByTypeEntry.EntityData.Leafs.Append("ciscoFlashFileByTypeStatus", types.YLeaf{"CiscoFlashFileByTypeStatus", ciscoFlashFileByTypeEntry.CiscoFlashFileByTypeStatus})
    ciscoFlashFileByTypeEntry.EntityData.Leafs.Append("ciscoFlashFileByTypeName", types.YLeaf{"CiscoFlashFileByTypeName", ciscoFlashFileByTypeEntry.CiscoFlashFileByTypeName})
    ciscoFlashFileByTypeEntry.EntityData.Leafs.Append("ciscoFlashFileByTypeDate", types.YLeaf{"CiscoFlashFileByTypeDate", ciscoFlashFileByTypeEntry.CiscoFlashFileByTypeDate})

    ciscoFlashFileByTypeEntry.EntityData.YListKeys = []string {"CiscoFlashFileType", "CiscoFlashDeviceIndex", "CiscoFlashPartitionIndex", "CiscoFlashFileIndex"}

    return &(ciscoFlashFileByTypeEntry.EntityData)
}

// CISCOFLASHMIB_CiscoFlashFileByTypeTable_CiscoFlashFileByTypeEntry_CiscoFlashFileByTypeStatus represents object in ciscoFlashFileTable.
type CISCOFLASHMIB_CiscoFlashFileByTypeTable_CiscoFlashFileByTypeEntry_CiscoFlashFileByTypeStatus string

const (
    CISCOFLASHMIB_CiscoFlashFileByTypeTable_CiscoFlashFileByTypeEntry_CiscoFlashFileByTypeStatus_deleted CISCOFLASHMIB_CiscoFlashFileByTypeTable_CiscoFlashFileByTypeEntry_CiscoFlashFileByTypeStatus = "deleted"

    CISCOFLASHMIB_CiscoFlashFileByTypeTable_CiscoFlashFileByTypeEntry_CiscoFlashFileByTypeStatus_invalidChecksum CISCOFLASHMIB_CiscoFlashFileByTypeTable_CiscoFlashFileByTypeEntry_CiscoFlashFileByTypeStatus = "invalidChecksum"

    CISCOFLASHMIB_CiscoFlashFileByTypeTable_CiscoFlashFileByTypeEntry_CiscoFlashFileByTypeStatus_valid CISCOFLASHMIB_CiscoFlashFileByTypeTable_CiscoFlashFileByTypeEntry_CiscoFlashFileByTypeStatus = "valid"
)

// CISCOFLASHMIB_CiscoFlashCopyTable
// A table of Flash copy operation entries. Each
// entry represents a Flash copy operation (to or
// from Flash) that has been initiated.
type CISCOFLASHMIB_CiscoFlashCopyTable struct {
    EntityData types.CommonEntityData
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
    // CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry.
    CiscoFlashCopyEntry []*CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry
}

func (ciscoFlashCopyTable *CISCOFLASHMIB_CiscoFlashCopyTable) GetEntityData() *types.CommonEntityData {
    ciscoFlashCopyTable.EntityData.YFilter = ciscoFlashCopyTable.YFilter
    ciscoFlashCopyTable.EntityData.YangName = "ciscoFlashCopyTable"
    ciscoFlashCopyTable.EntityData.BundleName = "cisco_ios_xe"
    ciscoFlashCopyTable.EntityData.ParentYangName = "CISCO-FLASH-MIB"
    ciscoFlashCopyTable.EntityData.SegmentPath = "ciscoFlashCopyTable"
    ciscoFlashCopyTable.EntityData.AbsolutePath = "CISCO-FLASH-MIB:CISCO-FLASH-MIB/" + ciscoFlashCopyTable.EntityData.SegmentPath
    ciscoFlashCopyTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoFlashCopyTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoFlashCopyTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoFlashCopyTable.EntityData.Children = types.NewOrderedMap()
    ciscoFlashCopyTable.EntityData.Children.Append("ciscoFlashCopyEntry", types.YChild{"CiscoFlashCopyEntry", nil})
    for i := range ciscoFlashCopyTable.CiscoFlashCopyEntry {
        ciscoFlashCopyTable.EntityData.Children.Append(types.GetSegmentPath(ciscoFlashCopyTable.CiscoFlashCopyEntry[i]), types.YChild{"CiscoFlashCopyEntry", ciscoFlashCopyTable.CiscoFlashCopyEntry[i]})
    }
    ciscoFlashCopyTable.EntityData.Leafs = types.NewOrderedMap()

    ciscoFlashCopyTable.EntityData.YListKeys = []string {}

    return &(ciscoFlashCopyTable.EntityData)
}

// CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry
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
type CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Object which specifies a unique entry in the
    // table. A management station wishing to initiate a copy operation should use
    // a pseudo-random value for this object when creating or modifying an
    // instance of a ciscoFlashCopyEntry. The type is interface{} with range:
    // 0..2147483647.
    CiscoFlashCopySerialNumber interface{}

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
    // CopyNotifyOnCompletion (opt). The type is CiscoFlashCopyCommand.
    CiscoFlashCopyCommand interface{}

    // The protocol to be used for any copy. Optional. Will default to tftp if not
    // specified.  Since feature support depends on a software release, version
    // number within the release, platform, and maybe the image type (subset
    // type), a management station would be expected to somehow determine the
    // protocol support for a command. The type is CiscoFlashCopyProtocol.
    CiscoFlashCopyProtocol interface{}

    // The server address to be used for any copy. Optional. Will default to
    // 'FFFFFFFF'H  (or 255.255.255.255).  Since this object can just hold only
    // IPv4 Transport type, it is deprecated and replaced by
    // ciscoFlashCopyServerAddrRev1. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    CiscoFlashCopyServerAddress interface{}

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
    CiscoFlashCopySourceName interface{}

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
    CiscoFlashCopyDestinationName interface{}

    // Remote user name for copy via rcp protocol. Optional. This object will be
    // ignored for protocols other than rcp. If specified, it will override the
    // remote user-name configured through the         rcmd remote-username
    // configuration command. The remote user-name is sent as the server user-name
    // in an rcp command request sent by the system to a remote rcp server. The
    // type is string with length: 1..255.
    CiscoFlashCopyRemoteUserName interface{}

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
    // CiscoFlashCopyStatus.
    CiscoFlashCopyStatus interface{}

    // Specifies whether or not a notification should be generated on the
    // completion of the copy operation. If specified,
    // ciscoFlashCopyCompletionTrap will be generated. It is the responsibility of
    // the management entity to ensure that the SNMP administrative model is
    // configured in such a way as to allow the notification to be delivered. The
    // type is bool.
    CiscoFlashCopyNotifyOnCompletion interface{}

    // Time taken for the copy operation. This object will be like a stopwatch,
    // starting when the operation starts, stopping when the operation completes.
    // If a management entity keeps a database of completion times for various
    // operations, it can then use the stopwatch capability to display percentage
    // completion time. The type is interface{} with range: 0..4294967295.
    CiscoFlashCopyTime interface{}

    // The status of this table entry. The type is RowStatus.
    CiscoFlashCopyEntryStatus interface{}

    // Specifies whether the file that is copied need to be verified for integrity
    // / authenticity, after copy succeeds. If it is set to true, and if the file
    // that is copied doesn't have integrity /authenticity attachement, or the
    // integrity / authenticity check fails, then the command will be aborted, and
    // the file that is copied will be deleted from the destination file system.
    // The type is bool.
    CiscoFlashCopyVerify interface{}

    // This object indicates the transport type of the address contained in
    // ciscoFlashCopyServerAddrRev1. Optional. Will default to '1' (IPv4 address
    // type). The type is InetAddressType.
    CiscoFlashCopyServerAddrType interface{}

    // The server address to be used for any copy. Optional. Will default to
    // 'FFFFFFFF'H  (or 255.255.255.255).  The Format of this address depends on
    // the value of the ciscoFlashCopyServerAddrType.  This object deprecates the
    // old ciscoFlashCopyServerAddress object. The type is string with length:
    // 0..255.
    CiscoFlashCopyServerAddrRev1 interface{}

    // Password used by ftp, sftp or scp for copying a file to/from an
    // ftp/sftp/scp server. This object must be created when the
    // ciscoFlashCopyProtocol is ftp, sftp or scp. Reading it returns a
    // zero-length string for security reasons. The type is string with length:
    // 1..40.
    CiscoFlashCopyRemotePassword interface{}
}

func (ciscoFlashCopyEntry *CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry) GetEntityData() *types.CommonEntityData {
    ciscoFlashCopyEntry.EntityData.YFilter = ciscoFlashCopyEntry.YFilter
    ciscoFlashCopyEntry.EntityData.YangName = "ciscoFlashCopyEntry"
    ciscoFlashCopyEntry.EntityData.BundleName = "cisco_ios_xe"
    ciscoFlashCopyEntry.EntityData.ParentYangName = "ciscoFlashCopyTable"
    ciscoFlashCopyEntry.EntityData.SegmentPath = "ciscoFlashCopyEntry" + types.AddKeyToken(ciscoFlashCopyEntry.CiscoFlashCopySerialNumber, "ciscoFlashCopySerialNumber")
    ciscoFlashCopyEntry.EntityData.AbsolutePath = "CISCO-FLASH-MIB:CISCO-FLASH-MIB/ciscoFlashCopyTable/" + ciscoFlashCopyEntry.EntityData.SegmentPath
    ciscoFlashCopyEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoFlashCopyEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoFlashCopyEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoFlashCopyEntry.EntityData.Children = types.NewOrderedMap()
    ciscoFlashCopyEntry.EntityData.Leafs = types.NewOrderedMap()
    ciscoFlashCopyEntry.EntityData.Leafs.Append("ciscoFlashCopySerialNumber", types.YLeaf{"CiscoFlashCopySerialNumber", ciscoFlashCopyEntry.CiscoFlashCopySerialNumber})
    ciscoFlashCopyEntry.EntityData.Leafs.Append("ciscoFlashCopyCommand", types.YLeaf{"CiscoFlashCopyCommand", ciscoFlashCopyEntry.CiscoFlashCopyCommand})
    ciscoFlashCopyEntry.EntityData.Leafs.Append("ciscoFlashCopyProtocol", types.YLeaf{"CiscoFlashCopyProtocol", ciscoFlashCopyEntry.CiscoFlashCopyProtocol})
    ciscoFlashCopyEntry.EntityData.Leafs.Append("ciscoFlashCopyServerAddress", types.YLeaf{"CiscoFlashCopyServerAddress", ciscoFlashCopyEntry.CiscoFlashCopyServerAddress})
    ciscoFlashCopyEntry.EntityData.Leafs.Append("ciscoFlashCopySourceName", types.YLeaf{"CiscoFlashCopySourceName", ciscoFlashCopyEntry.CiscoFlashCopySourceName})
    ciscoFlashCopyEntry.EntityData.Leafs.Append("ciscoFlashCopyDestinationName", types.YLeaf{"CiscoFlashCopyDestinationName", ciscoFlashCopyEntry.CiscoFlashCopyDestinationName})
    ciscoFlashCopyEntry.EntityData.Leafs.Append("ciscoFlashCopyRemoteUserName", types.YLeaf{"CiscoFlashCopyRemoteUserName", ciscoFlashCopyEntry.CiscoFlashCopyRemoteUserName})
    ciscoFlashCopyEntry.EntityData.Leafs.Append("ciscoFlashCopyStatus", types.YLeaf{"CiscoFlashCopyStatus", ciscoFlashCopyEntry.CiscoFlashCopyStatus})
    ciscoFlashCopyEntry.EntityData.Leafs.Append("ciscoFlashCopyNotifyOnCompletion", types.YLeaf{"CiscoFlashCopyNotifyOnCompletion", ciscoFlashCopyEntry.CiscoFlashCopyNotifyOnCompletion})
    ciscoFlashCopyEntry.EntityData.Leafs.Append("ciscoFlashCopyTime", types.YLeaf{"CiscoFlashCopyTime", ciscoFlashCopyEntry.CiscoFlashCopyTime})
    ciscoFlashCopyEntry.EntityData.Leafs.Append("ciscoFlashCopyEntryStatus", types.YLeaf{"CiscoFlashCopyEntryStatus", ciscoFlashCopyEntry.CiscoFlashCopyEntryStatus})
    ciscoFlashCopyEntry.EntityData.Leafs.Append("ciscoFlashCopyVerify", types.YLeaf{"CiscoFlashCopyVerify", ciscoFlashCopyEntry.CiscoFlashCopyVerify})
    ciscoFlashCopyEntry.EntityData.Leafs.Append("ciscoFlashCopyServerAddrType", types.YLeaf{"CiscoFlashCopyServerAddrType", ciscoFlashCopyEntry.CiscoFlashCopyServerAddrType})
    ciscoFlashCopyEntry.EntityData.Leafs.Append("ciscoFlashCopyServerAddrRev1", types.YLeaf{"CiscoFlashCopyServerAddrRev1", ciscoFlashCopyEntry.CiscoFlashCopyServerAddrRev1})
    ciscoFlashCopyEntry.EntityData.Leafs.Append("ciscoFlashCopyRemotePassword", types.YLeaf{"CiscoFlashCopyRemotePassword", ciscoFlashCopyEntry.CiscoFlashCopyRemotePassword})

    ciscoFlashCopyEntry.EntityData.YListKeys = []string {"CiscoFlashCopySerialNumber"}

    return &(ciscoFlashCopyEntry.EntityData)
}

// CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyCommand represents                         CopyNotifyOnCompletion (opt)
type CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyCommand string

const (
    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyCommand_copyToFlashWithErase CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyCommand = "copyToFlashWithErase"

    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyCommand_copyToFlashWithoutErase CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyCommand = "copyToFlashWithoutErase"

    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyCommand_copyFromFlash CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyCommand = "copyFromFlash"

    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyCommand_copyFromFlhLog CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyCommand = "copyFromFlhLog"
)

// CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyProtocol represents the protocol support for a command.
type CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyProtocol string

const (
    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyProtocol_tftp CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyProtocol = "tftp"

    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyProtocol_rcp CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyProtocol = "rcp"

    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyProtocol_lex CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyProtocol = "lex"

    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyProtocol_ftp CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyProtocol = "ftp"

    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyProtocol_scp CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyProtocol = "scp"

    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyProtocol_sftp CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyProtocol = "sftp"
)

// CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus represents       stop user from overwriting current boot image file.
type CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus string

const (
    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus_copyOperationPending CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus = "copyOperationPending"

    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus_copyInProgress CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus = "copyInProgress"

    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus_copyOperationSuccess CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus = "copyOperationSuccess"

    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus_copyInvalidOperation CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus = "copyInvalidOperation"

    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus_copyInvalidProtocol CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus = "copyInvalidProtocol"

    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus_copyInvalidSourceName CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus = "copyInvalidSourceName"

    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus_copyInvalidDestName CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus = "copyInvalidDestName"

    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus_copyInvalidServerAddress CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus = "copyInvalidServerAddress"

    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus_copyDeviceBusy CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus = "copyDeviceBusy"

    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus_copyDeviceOpenError CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus = "copyDeviceOpenError"

    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus_copyDeviceError CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus = "copyDeviceError"

    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus_copyDeviceNotProgrammable CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus = "copyDeviceNotProgrammable"

    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus_copyDeviceFull CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus = "copyDeviceFull"

    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus_copyFileOpenError CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus = "copyFileOpenError"

    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus_copyFileTransferError CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus = "copyFileTransferError"

    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus_copyFileChecksumError CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus = "copyFileChecksumError"

    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus_copyNoMemory CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus = "copyNoMemory"

    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus_copyUnknownFailure CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus = "copyUnknownFailure"

    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus_copyInvalidSignature CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus = "copyInvalidSignature"

    CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus_copyProhibited CISCOFLASHMIB_CiscoFlashCopyTable_CiscoFlashCopyEntry_CiscoFlashCopyStatus = "copyProhibited"
)

// CISCOFLASHMIB_CiscoFlashPartitioningTable
// A table of Flash partitioning operation entries. Each
// entry represents a Flash partitioning operation that
// has been initiated.
type CISCOFLASHMIB_CiscoFlashPartitioningTable struct {
    EntityData types.CommonEntityData
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
    // CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry.
    CiscoFlashPartitioningEntry []*CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry
}

func (ciscoFlashPartitioningTable *CISCOFLASHMIB_CiscoFlashPartitioningTable) GetEntityData() *types.CommonEntityData {
    ciscoFlashPartitioningTable.EntityData.YFilter = ciscoFlashPartitioningTable.YFilter
    ciscoFlashPartitioningTable.EntityData.YangName = "ciscoFlashPartitioningTable"
    ciscoFlashPartitioningTable.EntityData.BundleName = "cisco_ios_xe"
    ciscoFlashPartitioningTable.EntityData.ParentYangName = "CISCO-FLASH-MIB"
    ciscoFlashPartitioningTable.EntityData.SegmentPath = "ciscoFlashPartitioningTable"
    ciscoFlashPartitioningTable.EntityData.AbsolutePath = "CISCO-FLASH-MIB:CISCO-FLASH-MIB/" + ciscoFlashPartitioningTable.EntityData.SegmentPath
    ciscoFlashPartitioningTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoFlashPartitioningTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoFlashPartitioningTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoFlashPartitioningTable.EntityData.Children = types.NewOrderedMap()
    ciscoFlashPartitioningTable.EntityData.Children.Append("ciscoFlashPartitioningEntry", types.YChild{"CiscoFlashPartitioningEntry", nil})
    for i := range ciscoFlashPartitioningTable.CiscoFlashPartitioningEntry {
        ciscoFlashPartitioningTable.EntityData.Children.Append(types.GetSegmentPath(ciscoFlashPartitioningTable.CiscoFlashPartitioningEntry[i]), types.YChild{"CiscoFlashPartitioningEntry", ciscoFlashPartitioningTable.CiscoFlashPartitioningEntry[i]})
    }
    ciscoFlashPartitioningTable.EntityData.Leafs = types.NewOrderedMap()

    ciscoFlashPartitioningTable.EntityData.YListKeys = []string {}

    return &(ciscoFlashPartitioningTable.EntityData)
}

// CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry
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
type CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Object which specifies a unique entry in the
    // partitioning operations table. A management station wishing to initiate a
    // partitioning operation should use a pseudo-random value for this object
    // when creating or modifying an instance of a ciscoFlashPartitioningEntry.
    // The type is interface{} with range: 0..2147483647.
    CiscoFlashPartitioningSerialNumber interface{}

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
    // CiscoFlashPartitioningCommand.
    CiscoFlashPartitioningCommand interface{}

    // Destination device name. This name will be the value obtained from
    // FlashDeviceName. If the name is not specified, the default Flash device
    // will be assumed. The type is string with length: 0..255.
    CiscoFlashPartitioningDestinationName interface{}

    // This object is used to specify the number of partitions to be created. Its
    // value cannot exceed the value of ciscoFlashDeviceMaxPartitions.  To undo
    // partitioning (revert to a single partition), this object must have the
    // value 1. The type is interface{} with range: 1..4294967295.
    CiscoFlashPartitioningPartitionCount interface{}

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
    CiscoFlashPartitioningPartitionSizes interface{}

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
    // CiscoFlashPartitioningStatus.
    CiscoFlashPartitioningStatus interface{}

    // Specifies whether or not a notification should be generated on the
    // completion of the partitioning operation. If specified,
    // ciscoFlashPartitioningCompletionTrap will be generated. It is the
    // responsibility of the management entity to ensure that the SNMP
    // administrative model is configured in such a way as to allow the
    // notification to be delivered. The type is bool.
    CiscoFlashPartitioningNotifyOnCompletion interface{}

    // Time taken for the operation. This object will be like a stopwatch,
    // starting when the operation starts, stopping when the operation completes.
    // If a management entity keeps a database of completion times for various
    // operations, it can then use the stopwatch capability to display percentage
    // completion time. The type is interface{} with range: 0..4294967295.
    CiscoFlashPartitioningTime interface{}

    // The status of this table entry. The type is RowStatus.
    CiscoFlashPartitioningEntryStatus interface{}
}

func (ciscoFlashPartitioningEntry *CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry) GetEntityData() *types.CommonEntityData {
    ciscoFlashPartitioningEntry.EntityData.YFilter = ciscoFlashPartitioningEntry.YFilter
    ciscoFlashPartitioningEntry.EntityData.YangName = "ciscoFlashPartitioningEntry"
    ciscoFlashPartitioningEntry.EntityData.BundleName = "cisco_ios_xe"
    ciscoFlashPartitioningEntry.EntityData.ParentYangName = "ciscoFlashPartitioningTable"
    ciscoFlashPartitioningEntry.EntityData.SegmentPath = "ciscoFlashPartitioningEntry" + types.AddKeyToken(ciscoFlashPartitioningEntry.CiscoFlashPartitioningSerialNumber, "ciscoFlashPartitioningSerialNumber")
    ciscoFlashPartitioningEntry.EntityData.AbsolutePath = "CISCO-FLASH-MIB:CISCO-FLASH-MIB/ciscoFlashPartitioningTable/" + ciscoFlashPartitioningEntry.EntityData.SegmentPath
    ciscoFlashPartitioningEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoFlashPartitioningEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoFlashPartitioningEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoFlashPartitioningEntry.EntityData.Children = types.NewOrderedMap()
    ciscoFlashPartitioningEntry.EntityData.Leafs = types.NewOrderedMap()
    ciscoFlashPartitioningEntry.EntityData.Leafs.Append("ciscoFlashPartitioningSerialNumber", types.YLeaf{"CiscoFlashPartitioningSerialNumber", ciscoFlashPartitioningEntry.CiscoFlashPartitioningSerialNumber})
    ciscoFlashPartitioningEntry.EntityData.Leafs.Append("ciscoFlashPartitioningCommand", types.YLeaf{"CiscoFlashPartitioningCommand", ciscoFlashPartitioningEntry.CiscoFlashPartitioningCommand})
    ciscoFlashPartitioningEntry.EntityData.Leafs.Append("ciscoFlashPartitioningDestinationName", types.YLeaf{"CiscoFlashPartitioningDestinationName", ciscoFlashPartitioningEntry.CiscoFlashPartitioningDestinationName})
    ciscoFlashPartitioningEntry.EntityData.Leafs.Append("ciscoFlashPartitioningPartitionCount", types.YLeaf{"CiscoFlashPartitioningPartitionCount", ciscoFlashPartitioningEntry.CiscoFlashPartitioningPartitionCount})
    ciscoFlashPartitioningEntry.EntityData.Leafs.Append("ciscoFlashPartitioningPartitionSizes", types.YLeaf{"CiscoFlashPartitioningPartitionSizes", ciscoFlashPartitioningEntry.CiscoFlashPartitioningPartitionSizes})
    ciscoFlashPartitioningEntry.EntityData.Leafs.Append("ciscoFlashPartitioningStatus", types.YLeaf{"CiscoFlashPartitioningStatus", ciscoFlashPartitioningEntry.CiscoFlashPartitioningStatus})
    ciscoFlashPartitioningEntry.EntityData.Leafs.Append("ciscoFlashPartitioningNotifyOnCompletion", types.YLeaf{"CiscoFlashPartitioningNotifyOnCompletion", ciscoFlashPartitioningEntry.CiscoFlashPartitioningNotifyOnCompletion})
    ciscoFlashPartitioningEntry.EntityData.Leafs.Append("ciscoFlashPartitioningTime", types.YLeaf{"CiscoFlashPartitioningTime", ciscoFlashPartitioningEntry.CiscoFlashPartitioningTime})
    ciscoFlashPartitioningEntry.EntityData.Leafs.Append("ciscoFlashPartitioningEntryStatus", types.YLeaf{"CiscoFlashPartitioningEntryStatus", ciscoFlashPartitioningEntry.CiscoFlashPartitioningEntryStatus})

    ciscoFlashPartitioningEntry.EntityData.YListKeys = []string {"CiscoFlashPartitioningSerialNumber"}

    return &(ciscoFlashPartitioningEntry.EntityData)
}

// CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningCommand represents                         PartitioningNotifyOnCompletion (opt)
type CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningCommand string

const (
    CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningCommand_partition CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningCommand = "partition"
)

// CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningStatus represents         failure unknown
type CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningStatus string

const (
    CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningStatus_partitioningInProgress CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningStatus = "partitioningInProgress"

    CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningStatus_partitioningOperationSuccess CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningStatus = "partitioningOperationSuccess"

    CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningStatus_partitioningInvalidOperation CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningStatus = "partitioningInvalidOperation"

    CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningStatus_partitioningInvalidDestName CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningStatus = "partitioningInvalidDestName"

    CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningStatus_partitioningInvalidPartitionCount CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningStatus = "partitioningInvalidPartitionCount"

    CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningStatus_partitioningInvalidPartitionSizes CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningStatus = "partitioningInvalidPartitionSizes"

    CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningStatus_partitioningDeviceBusy CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningStatus = "partitioningDeviceBusy"

    CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningStatus_partitioningDeviceOpenError CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningStatus = "partitioningDeviceOpenError"

    CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningStatus_partitioningDeviceError CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningStatus = "partitioningDeviceError"

    CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningStatus_partitioningNoMemory CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningStatus = "partitioningNoMemory"

    CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningStatus_partitioningUnknownFailure CISCOFLASHMIB_CiscoFlashPartitioningTable_CiscoFlashPartitioningEntry_CiscoFlashPartitioningStatus = "partitioningUnknownFailure"
)

// CISCOFLASHMIB_CiscoFlashMiscOpTable
// A table of misc Flash operation entries. Each
// entry represents a Flash operation that
// has been initiated.
type CISCOFLASHMIB_CiscoFlashMiscOpTable struct {
    EntityData types.CommonEntityData
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
    // CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry.
    CiscoFlashMiscOpEntry []*CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry
}

func (ciscoFlashMiscOpTable *CISCOFLASHMIB_CiscoFlashMiscOpTable) GetEntityData() *types.CommonEntityData {
    ciscoFlashMiscOpTable.EntityData.YFilter = ciscoFlashMiscOpTable.YFilter
    ciscoFlashMiscOpTable.EntityData.YangName = "ciscoFlashMiscOpTable"
    ciscoFlashMiscOpTable.EntityData.BundleName = "cisco_ios_xe"
    ciscoFlashMiscOpTable.EntityData.ParentYangName = "CISCO-FLASH-MIB"
    ciscoFlashMiscOpTable.EntityData.SegmentPath = "ciscoFlashMiscOpTable"
    ciscoFlashMiscOpTable.EntityData.AbsolutePath = "CISCO-FLASH-MIB:CISCO-FLASH-MIB/" + ciscoFlashMiscOpTable.EntityData.SegmentPath
    ciscoFlashMiscOpTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoFlashMiscOpTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoFlashMiscOpTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoFlashMiscOpTable.EntityData.Children = types.NewOrderedMap()
    ciscoFlashMiscOpTable.EntityData.Children.Append("ciscoFlashMiscOpEntry", types.YChild{"CiscoFlashMiscOpEntry", nil})
    for i := range ciscoFlashMiscOpTable.CiscoFlashMiscOpEntry {
        ciscoFlashMiscOpTable.EntityData.Children.Append(types.GetSegmentPath(ciscoFlashMiscOpTable.CiscoFlashMiscOpEntry[i]), types.YChild{"CiscoFlashMiscOpEntry", ciscoFlashMiscOpTable.CiscoFlashMiscOpEntry[i]})
    }
    ciscoFlashMiscOpTable.EntityData.Leafs = types.NewOrderedMap()

    ciscoFlashMiscOpTable.EntityData.YListKeys = []string {}

    return &(ciscoFlashMiscOpTable.EntityData)
}

// CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry
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
type CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Object which specifies a unique entry in the
    // table. A management station wishing to initiate a flash operation should
    // use a pseudo-random value for this object when creating or modifying an
    // instance of a ciscoFlashMiscOpEntry. The type is interface{} with range:
    // 0..2147483647.
    CiscoFlashMiscOpSerialNumber interface{}

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
    // MiscOpNotifyOnCompletion (opt). The type is CiscoFlashMiscOpCommand.
    CiscoFlashMiscOpCommand interface{}

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
    CiscoFlashMiscOpDestinationName interface{}

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
    // CiscoFlashMiscOpStatus.
    CiscoFlashMiscOpStatus interface{}

    // Specifies whether or not a notification should be generated on the
    // completion of an operation. If specified, ciscoFlashMiscOpCompletionTrap
    // will be generated. It is the responsibility of the management entity to
    // ensure that the SNMP administrative model is configured in such a way as to
    // allow the notification to be delivered. The type is bool.
    CiscoFlashMiscOpNotifyOnCompletion interface{}

    // Time taken for the operation. This object will be like a stopwatch,
    // starting when the operation starts, stopping when the operation completes.
    // If a management entity keeps a database of completion times for various
    // operations, it can then use the stopwatch capability to display percentage
    // completion time. The type is interface{} with range: 0..4294967295.
    CiscoFlashMiscOpTime interface{}

    // The status of this table entry. The type is RowStatus.
    CiscoFlashMiscOpEntryStatus interface{}
}

func (ciscoFlashMiscOpEntry *CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry) GetEntityData() *types.CommonEntityData {
    ciscoFlashMiscOpEntry.EntityData.YFilter = ciscoFlashMiscOpEntry.YFilter
    ciscoFlashMiscOpEntry.EntityData.YangName = "ciscoFlashMiscOpEntry"
    ciscoFlashMiscOpEntry.EntityData.BundleName = "cisco_ios_xe"
    ciscoFlashMiscOpEntry.EntityData.ParentYangName = "ciscoFlashMiscOpTable"
    ciscoFlashMiscOpEntry.EntityData.SegmentPath = "ciscoFlashMiscOpEntry" + types.AddKeyToken(ciscoFlashMiscOpEntry.CiscoFlashMiscOpSerialNumber, "ciscoFlashMiscOpSerialNumber")
    ciscoFlashMiscOpEntry.EntityData.AbsolutePath = "CISCO-FLASH-MIB:CISCO-FLASH-MIB/ciscoFlashMiscOpTable/" + ciscoFlashMiscOpEntry.EntityData.SegmentPath
    ciscoFlashMiscOpEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoFlashMiscOpEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoFlashMiscOpEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoFlashMiscOpEntry.EntityData.Children = types.NewOrderedMap()
    ciscoFlashMiscOpEntry.EntityData.Leafs = types.NewOrderedMap()
    ciscoFlashMiscOpEntry.EntityData.Leafs.Append("ciscoFlashMiscOpSerialNumber", types.YLeaf{"CiscoFlashMiscOpSerialNumber", ciscoFlashMiscOpEntry.CiscoFlashMiscOpSerialNumber})
    ciscoFlashMiscOpEntry.EntityData.Leafs.Append("ciscoFlashMiscOpCommand", types.YLeaf{"CiscoFlashMiscOpCommand", ciscoFlashMiscOpEntry.CiscoFlashMiscOpCommand})
    ciscoFlashMiscOpEntry.EntityData.Leafs.Append("ciscoFlashMiscOpDestinationName", types.YLeaf{"CiscoFlashMiscOpDestinationName", ciscoFlashMiscOpEntry.CiscoFlashMiscOpDestinationName})
    ciscoFlashMiscOpEntry.EntityData.Leafs.Append("ciscoFlashMiscOpStatus", types.YLeaf{"CiscoFlashMiscOpStatus", ciscoFlashMiscOpEntry.CiscoFlashMiscOpStatus})
    ciscoFlashMiscOpEntry.EntityData.Leafs.Append("ciscoFlashMiscOpNotifyOnCompletion", types.YLeaf{"CiscoFlashMiscOpNotifyOnCompletion", ciscoFlashMiscOpEntry.CiscoFlashMiscOpNotifyOnCompletion})
    ciscoFlashMiscOpEntry.EntityData.Leafs.Append("ciscoFlashMiscOpTime", types.YLeaf{"CiscoFlashMiscOpTime", ciscoFlashMiscOpEntry.CiscoFlashMiscOpTime})
    ciscoFlashMiscOpEntry.EntityData.Leafs.Append("ciscoFlashMiscOpEntryStatus", types.YLeaf{"CiscoFlashMiscOpEntryStatus", ciscoFlashMiscOpEntry.CiscoFlashMiscOpEntryStatus})

    ciscoFlashMiscOpEntry.EntityData.YListKeys = []string {"CiscoFlashMiscOpSerialNumber"}

    return &(ciscoFlashMiscOpEntry.EntityData)
}

// CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpCommand represents                 MiscOpNotifyOnCompletion (opt)
type CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpCommand string

const (
    CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpCommand_erase CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpCommand = "erase"

    CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpCommand_verify CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpCommand = "verify"

    CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpCommand_delete_ CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpCommand = "delete"

    CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpCommand_undelete CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpCommand = "undelete"

    CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpCommand_squeeze CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpCommand = "squeeze"

    CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpCommand_format CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpCommand = "format"
)

// CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus represents         the format operation failed
type CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus string

const (
    CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus_miscOpInProgress CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus = "miscOpInProgress"

    CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus_miscOpOperationSuccess CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus = "miscOpOperationSuccess"

    CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus_miscOpInvalidOperation CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus = "miscOpInvalidOperation"

    CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus_miscOpInvalidDestName CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus = "miscOpInvalidDestName"

    CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus_miscOpDeviceBusy CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus = "miscOpDeviceBusy"

    CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus_miscOpDeviceOpenError CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus = "miscOpDeviceOpenError"

    CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus_miscOpDeviceError CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus = "miscOpDeviceError"

    CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus_miscOpDeviceNotProgrammable CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus = "miscOpDeviceNotProgrammable"

    CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus_miscOpFileOpenError CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus = "miscOpFileOpenError"

    CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus_miscOpFileDeleteFailure CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus = "miscOpFileDeleteFailure"

    CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus_miscOpFileUndeleteFailure CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus = "miscOpFileUndeleteFailure"

    CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus_miscOpFileChecksumError CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus = "miscOpFileChecksumError"

    CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus_miscOpNoMemory CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus = "miscOpNoMemory"

    CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus_miscOpUnknownFailure CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus = "miscOpUnknownFailure"

    CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus_miscOpSqueezeFailure CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus = "miscOpSqueezeFailure"

    CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus_miscOpNoSuchFile CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus = "miscOpNoSuchFile"

    CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus_miscOpFormatFailure CISCOFLASHMIB_CiscoFlashMiscOpTable_CiscoFlashMiscOpEntry_CiscoFlashMiscOpStatus = "miscOpFormatFailure"
)

