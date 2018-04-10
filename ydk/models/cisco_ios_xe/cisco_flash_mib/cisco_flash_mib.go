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

func (cISCOFLASHMIB *CISCOFLASHMIB) GetEntityData() *types.CommonEntityData {
    cISCOFLASHMIB.EntityData.YFilter = cISCOFLASHMIB.YFilter
    cISCOFLASHMIB.EntityData.YangName = "CISCO-FLASH-MIB"
    cISCOFLASHMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOFLASHMIB.EntityData.ParentYangName = "CISCO-FLASH-MIB"
    cISCOFLASHMIB.EntityData.SegmentPath = "CISCO-FLASH-MIB:CISCO-FLASH-MIB"
    cISCOFLASHMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOFLASHMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOFLASHMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOFLASHMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOFLASHMIB.EntityData.Children["ciscoFlashDevice"] = types.YChild{"Ciscoflashdevice", &cISCOFLASHMIB.Ciscoflashdevice}
    cISCOFLASHMIB.EntityData.Children["ciscoFlashCfg"] = types.YChild{"Ciscoflashcfg", &cISCOFLASHMIB.Ciscoflashcfg}
    cISCOFLASHMIB.EntityData.Children["ciscoFlashDeviceTable"] = types.YChild{"Ciscoflashdevicetable", &cISCOFLASHMIB.Ciscoflashdevicetable}
    cISCOFLASHMIB.EntityData.Children["ciscoFlashChipTable"] = types.YChild{"Ciscoflashchiptable", &cISCOFLASHMIB.Ciscoflashchiptable}
    cISCOFLASHMIB.EntityData.Children["ciscoFlashPartitionTable"] = types.YChild{"Ciscoflashpartitiontable", &cISCOFLASHMIB.Ciscoflashpartitiontable}
    cISCOFLASHMIB.EntityData.Children["ciscoFlashFileTable"] = types.YChild{"Ciscoflashfiletable", &cISCOFLASHMIB.Ciscoflashfiletable}
    cISCOFLASHMIB.EntityData.Children["ciscoFlashFileByTypeTable"] = types.YChild{"Ciscoflashfilebytypetable", &cISCOFLASHMIB.Ciscoflashfilebytypetable}
    cISCOFLASHMIB.EntityData.Children["ciscoFlashCopyTable"] = types.YChild{"Ciscoflashcopytable", &cISCOFLASHMIB.Ciscoflashcopytable}
    cISCOFLASHMIB.EntityData.Children["ciscoFlashPartitioningTable"] = types.YChild{"Ciscoflashpartitioningtable", &cISCOFLASHMIB.Ciscoflashpartitioningtable}
    cISCOFLASHMIB.EntityData.Children["ciscoFlashMiscOpTable"] = types.YChild{"Ciscoflashmiscoptable", &cISCOFLASHMIB.Ciscoflashmiscoptable}
    cISCOFLASHMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOFLASHMIB.EntityData)
}

// CISCOFLASHMIB_Ciscoflashdevice
type CISCOFLASHMIB_Ciscoflashdevice struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Number of Flash devices supported by the system. If the system does not
    // support any Flash devices, this MIB will not be loaded on that system. The
    // value of this object will therefore be atleast 1. The type is interface{}
    // with range: 0..4294967295.
    Ciscoflashdevicessupported interface{}
}

func (ciscoflashdevice *CISCOFLASHMIB_Ciscoflashdevice) GetEntityData() *types.CommonEntityData {
    ciscoflashdevice.EntityData.YFilter = ciscoflashdevice.YFilter
    ciscoflashdevice.EntityData.YangName = "ciscoFlashDevice"
    ciscoflashdevice.EntityData.BundleName = "cisco_ios_xe"
    ciscoflashdevice.EntityData.ParentYangName = "CISCO-FLASH-MIB"
    ciscoflashdevice.EntityData.SegmentPath = "ciscoFlashDevice"
    ciscoflashdevice.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoflashdevice.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoflashdevice.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoflashdevice.EntityData.Children = make(map[string]types.YChild)
    ciscoflashdevice.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoflashdevice.EntityData.Leafs["ciscoFlashDevicesSupported"] = types.YLeaf{"Ciscoflashdevicessupported", ciscoflashdevice.Ciscoflashdevicessupported}
    return &(ciscoflashdevice.EntityData)
}

// CISCOFLASHMIB_Ciscoflashcfg
type CISCOFLASHMIB_Ciscoflashcfg struct {
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

func (ciscoflashcfg *CISCOFLASHMIB_Ciscoflashcfg) GetEntityData() *types.CommonEntityData {
    ciscoflashcfg.EntityData.YFilter = ciscoflashcfg.YFilter
    ciscoflashcfg.EntityData.YangName = "ciscoFlashCfg"
    ciscoflashcfg.EntityData.BundleName = "cisco_ios_xe"
    ciscoflashcfg.EntityData.ParentYangName = "CISCO-FLASH-MIB"
    ciscoflashcfg.EntityData.SegmentPath = "ciscoFlashCfg"
    ciscoflashcfg.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoflashcfg.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoflashcfg.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoflashcfg.EntityData.Children = make(map[string]types.YChild)
    ciscoflashcfg.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoflashcfg.EntityData.Leafs["ciscoFlashCfgDevInsNotifEnable"] = types.YLeaf{"Ciscoflashcfgdevinsnotifenable", ciscoflashcfg.Ciscoflashcfgdevinsnotifenable}
    ciscoflashcfg.EntityData.Leafs["ciscoFlashCfgDevRemNotifEnable"] = types.YLeaf{"Ciscoflashcfgdevremnotifenable", ciscoflashcfg.Ciscoflashcfgdevremnotifenable}
    ciscoflashcfg.EntityData.Leafs["ciscoFlashPartitionLowSpaceNotifEnable"] = types.YLeaf{"Ciscoflashpartitionlowspacenotifenable", ciscoflashcfg.Ciscoflashpartitionlowspacenotifenable}
    return &(ciscoflashcfg.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table of flash device properties for each initialized flash
    // device. Each entry can be randomly accessed by using ciscoFlashDeviceIndex
    // as an index into the table. Note that removable devices will have an entry
    // in the table even when they have been removed. However, a non-removable
    // device that has not been installed will not have an entry in the table. The
    // type is slice of CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry.
    Ciscoflashdeviceentry []CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry
}

func (ciscoflashdevicetable *CISCOFLASHMIB_Ciscoflashdevicetable) GetEntityData() *types.CommonEntityData {
    ciscoflashdevicetable.EntityData.YFilter = ciscoflashdevicetable.YFilter
    ciscoflashdevicetable.EntityData.YangName = "ciscoFlashDeviceTable"
    ciscoflashdevicetable.EntityData.BundleName = "cisco_ios_xe"
    ciscoflashdevicetable.EntityData.ParentYangName = "CISCO-FLASH-MIB"
    ciscoflashdevicetable.EntityData.SegmentPath = "ciscoFlashDeviceTable"
    ciscoflashdevicetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoflashdevicetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoflashdevicetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoflashdevicetable.EntityData.Children = make(map[string]types.YChild)
    ciscoflashdevicetable.EntityData.Children["ciscoFlashDeviceEntry"] = types.YChild{"Ciscoflashdeviceentry", nil}
    for i := range ciscoflashdevicetable.Ciscoflashdeviceentry {
        ciscoflashdevicetable.EntityData.Children[types.GetSegmentPath(&ciscoflashdevicetable.Ciscoflashdeviceentry[i])] = types.YChild{"Ciscoflashdeviceentry", &ciscoflashdevicetable.Ciscoflashdeviceentry[i]}
    }
    ciscoflashdevicetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ciscoflashdevicetable.EntityData)
}

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
    EntityData types.CommonEntityData
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
    // b'(([0-1](\\.[1-3]?[0-9]))|(2\\.(0|([1-9]\\d*))))(\\.(0|([1-9]\\d*)))*'.
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

func (ciscoflashdeviceentry *CISCOFLASHMIB_Ciscoflashdevicetable_Ciscoflashdeviceentry) GetEntityData() *types.CommonEntityData {
    ciscoflashdeviceentry.EntityData.YFilter = ciscoflashdeviceentry.YFilter
    ciscoflashdeviceentry.EntityData.YangName = "ciscoFlashDeviceEntry"
    ciscoflashdeviceentry.EntityData.BundleName = "cisco_ios_xe"
    ciscoflashdeviceentry.EntityData.ParentYangName = "ciscoFlashDeviceTable"
    ciscoflashdeviceentry.EntityData.SegmentPath = "ciscoFlashDeviceEntry" + "[ciscoFlashDeviceIndex='" + fmt.Sprintf("%v", ciscoflashdeviceentry.Ciscoflashdeviceindex) + "']"
    ciscoflashdeviceentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoflashdeviceentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoflashdeviceentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoflashdeviceentry.EntityData.Children = make(map[string]types.YChild)
    ciscoflashdeviceentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoflashdeviceentry.EntityData.Leafs["ciscoFlashDeviceIndex"] = types.YLeaf{"Ciscoflashdeviceindex", ciscoflashdeviceentry.Ciscoflashdeviceindex}
    ciscoflashdeviceentry.EntityData.Leafs["ciscoFlashDeviceSize"] = types.YLeaf{"Ciscoflashdevicesize", ciscoflashdeviceentry.Ciscoflashdevicesize}
    ciscoflashdeviceentry.EntityData.Leafs["ciscoFlashDeviceMinPartitionSize"] = types.YLeaf{"Ciscoflashdeviceminpartitionsize", ciscoflashdeviceentry.Ciscoflashdeviceminpartitionsize}
    ciscoflashdeviceentry.EntityData.Leafs["ciscoFlashDeviceMaxPartitions"] = types.YLeaf{"Ciscoflashdevicemaxpartitions", ciscoflashdeviceentry.Ciscoflashdevicemaxpartitions}
    ciscoflashdeviceentry.EntityData.Leafs["ciscoFlashDevicePartitions"] = types.YLeaf{"Ciscoflashdevicepartitions", ciscoflashdeviceentry.Ciscoflashdevicepartitions}
    ciscoflashdeviceentry.EntityData.Leafs["ciscoFlashDeviceChipCount"] = types.YLeaf{"Ciscoflashdevicechipcount", ciscoflashdeviceentry.Ciscoflashdevicechipcount}
    ciscoflashdeviceentry.EntityData.Leafs["ciscoFlashDeviceName"] = types.YLeaf{"Ciscoflashdevicename", ciscoflashdeviceentry.Ciscoflashdevicename}
    ciscoflashdeviceentry.EntityData.Leafs["ciscoFlashDeviceDescr"] = types.YLeaf{"Ciscoflashdevicedescr", ciscoflashdeviceentry.Ciscoflashdevicedescr}
    ciscoflashdeviceentry.EntityData.Leafs["ciscoFlashDeviceController"] = types.YLeaf{"Ciscoflashdevicecontroller", ciscoflashdeviceentry.Ciscoflashdevicecontroller}
    ciscoflashdeviceentry.EntityData.Leafs["ciscoFlashDeviceCard"] = types.YLeaf{"Ciscoflashdevicecard", ciscoflashdeviceentry.Ciscoflashdevicecard}
    ciscoflashdeviceentry.EntityData.Leafs["ciscoFlashDeviceProgrammingJumper"] = types.YLeaf{"Ciscoflashdeviceprogrammingjumper", ciscoflashdeviceentry.Ciscoflashdeviceprogrammingjumper}
    ciscoflashdeviceentry.EntityData.Leafs["ciscoFlashDeviceInitTime"] = types.YLeaf{"Ciscoflashdeviceinittime", ciscoflashdeviceentry.Ciscoflashdeviceinittime}
    ciscoflashdeviceentry.EntityData.Leafs["ciscoFlashDeviceRemovable"] = types.YLeaf{"Ciscoflashdeviceremovable", ciscoflashdeviceentry.Ciscoflashdeviceremovable}
    ciscoflashdeviceentry.EntityData.Leafs["ciscoFlashPhyEntIndex"] = types.YLeaf{"Ciscoflashphyentindex", ciscoflashdeviceentry.Ciscoflashphyentindex}
    ciscoflashdeviceentry.EntityData.Leafs["ciscoFlashDeviceNameExtended"] = types.YLeaf{"Ciscoflashdevicenameextended", ciscoflashdeviceentry.Ciscoflashdevicenameextended}
    ciscoflashdeviceentry.EntityData.Leafs["ciscoFlashDeviceSizeExtended"] = types.YLeaf{"Ciscoflashdevicesizeextended", ciscoflashdeviceentry.Ciscoflashdevicesizeextended}
    ciscoflashdeviceentry.EntityData.Leafs["ciscoFlashDeviceMinPartitionSizeExtended"] = types.YLeaf{"Ciscoflashdeviceminpartitionsizeextended", ciscoflashdeviceentry.Ciscoflashdeviceminpartitionsizeextended}
    return &(ciscoflashdeviceentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table of chip info for each flash device initialized in the
    // system. An entry is indexed by two objects - the device index and the chip
    // index within that device. The type is slice of
    // CISCOFLASHMIB_Ciscoflashchiptable_Ciscoflashchipentry.
    Ciscoflashchipentry []CISCOFLASHMIB_Ciscoflashchiptable_Ciscoflashchipentry
}

func (ciscoflashchiptable *CISCOFLASHMIB_Ciscoflashchiptable) GetEntityData() *types.CommonEntityData {
    ciscoflashchiptable.EntityData.YFilter = ciscoflashchiptable.YFilter
    ciscoflashchiptable.EntityData.YangName = "ciscoFlashChipTable"
    ciscoflashchiptable.EntityData.BundleName = "cisco_ios_xe"
    ciscoflashchiptable.EntityData.ParentYangName = "CISCO-FLASH-MIB"
    ciscoflashchiptable.EntityData.SegmentPath = "ciscoFlashChipTable"
    ciscoflashchiptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoflashchiptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoflashchiptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoflashchiptable.EntityData.Children = make(map[string]types.YChild)
    ciscoflashchiptable.EntityData.Children["ciscoFlashChipEntry"] = types.YChild{"Ciscoflashchipentry", nil}
    for i := range ciscoflashchiptable.Ciscoflashchipentry {
        ciscoflashchiptable.EntityData.Children[types.GetSegmentPath(&ciscoflashchiptable.Ciscoflashchipentry[i])] = types.YChild{"Ciscoflashchipentry", &ciscoflashchiptable.Ciscoflashchipentry[i]}
    }
    ciscoflashchiptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ciscoflashchiptable.EntityData)
}

// CISCOFLASHMIB_Ciscoflashchiptable_Ciscoflashchipentry
// An entry in the table of chip info for each
// flash device initialized in the system.
// An entry is indexed by two objects - the
// device index and the chip index within that
// device.
type CISCOFLASHMIB_Ciscoflashchiptable_Ciscoflashchipentry struct {
    EntityData types.CommonEntityData
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

func (ciscoflashchipentry *CISCOFLASHMIB_Ciscoflashchiptable_Ciscoflashchipentry) GetEntityData() *types.CommonEntityData {
    ciscoflashchipentry.EntityData.YFilter = ciscoflashchipentry.YFilter
    ciscoflashchipentry.EntityData.YangName = "ciscoFlashChipEntry"
    ciscoflashchipentry.EntityData.BundleName = "cisco_ios_xe"
    ciscoflashchipentry.EntityData.ParentYangName = "ciscoFlashChipTable"
    ciscoflashchipentry.EntityData.SegmentPath = "ciscoFlashChipEntry" + "[ciscoFlashDeviceIndex='" + fmt.Sprintf("%v", ciscoflashchipentry.Ciscoflashdeviceindex) + "']" + "[ciscoFlashChipIndex='" + fmt.Sprintf("%v", ciscoflashchipentry.Ciscoflashchipindex) + "']"
    ciscoflashchipentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoflashchipentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoflashchipentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoflashchipentry.EntityData.Children = make(map[string]types.YChild)
    ciscoflashchipentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoflashchipentry.EntityData.Leafs["ciscoFlashDeviceIndex"] = types.YLeaf{"Ciscoflashdeviceindex", ciscoflashchipentry.Ciscoflashdeviceindex}
    ciscoflashchipentry.EntityData.Leafs["ciscoFlashChipIndex"] = types.YLeaf{"Ciscoflashchipindex", ciscoflashchipentry.Ciscoflashchipindex}
    ciscoflashchipentry.EntityData.Leafs["ciscoFlashChipCode"] = types.YLeaf{"Ciscoflashchipcode", ciscoflashchipentry.Ciscoflashchipcode}
    ciscoflashchipentry.EntityData.Leafs["ciscoFlashChipDescr"] = types.YLeaf{"Ciscoflashchipdescr", ciscoflashchipentry.Ciscoflashchipdescr}
    ciscoflashchipentry.EntityData.Leafs["ciscoFlashChipWriteRetries"] = types.YLeaf{"Ciscoflashchipwriteretries", ciscoflashchipentry.Ciscoflashchipwriteretries}
    ciscoflashchipentry.EntityData.Leafs["ciscoFlashChipEraseRetries"] = types.YLeaf{"Ciscoflashchiperaseretries", ciscoflashchipentry.Ciscoflashchiperaseretries}
    ciscoflashchipentry.EntityData.Leafs["ciscoFlashChipMaxWriteRetries"] = types.YLeaf{"Ciscoflashchipmaxwriteretries", ciscoflashchipentry.Ciscoflashchipmaxwriteretries}
    ciscoflashchipentry.EntityData.Leafs["ciscoFlashChipMaxEraseRetries"] = types.YLeaf{"Ciscoflashchipmaxeraseretries", ciscoflashchipentry.Ciscoflashchipmaxeraseretries}
    return &(ciscoflashchipentry.EntityData)
}

// CISCOFLASHMIB_Ciscoflashpartitiontable
// Table of flash device partition properties for each
// initialized flash partition. Whenever there is no
// explicit partitioning done, a single partition spanning
// the entire device will be assumed to exist. There will
// therefore always be atleast one partition on a device.
type CISCOFLASHMIB_Ciscoflashpartitiontable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table of flash partition properties for each initialized
    // flash partition. Each entry will be indexed by a device number and a
    // partition number within the device. The type is slice of
    // CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry.
    Ciscoflashpartitionentry []CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry
}

func (ciscoflashpartitiontable *CISCOFLASHMIB_Ciscoflashpartitiontable) GetEntityData() *types.CommonEntityData {
    ciscoflashpartitiontable.EntityData.YFilter = ciscoflashpartitiontable.YFilter
    ciscoflashpartitiontable.EntityData.YangName = "ciscoFlashPartitionTable"
    ciscoflashpartitiontable.EntityData.BundleName = "cisco_ios_xe"
    ciscoflashpartitiontable.EntityData.ParentYangName = "CISCO-FLASH-MIB"
    ciscoflashpartitiontable.EntityData.SegmentPath = "ciscoFlashPartitionTable"
    ciscoflashpartitiontable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoflashpartitiontable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoflashpartitiontable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoflashpartitiontable.EntityData.Children = make(map[string]types.YChild)
    ciscoflashpartitiontable.EntityData.Children["ciscoFlashPartitionEntry"] = types.YChild{"Ciscoflashpartitionentry", nil}
    for i := range ciscoflashpartitiontable.Ciscoflashpartitionentry {
        ciscoflashpartitiontable.EntityData.Children[types.GetSegmentPath(&ciscoflashpartitiontable.Ciscoflashpartitionentry[i])] = types.YChild{"Ciscoflashpartitionentry", &ciscoflashpartitiontable.Ciscoflashpartitionentry[i]}
    }
    ciscoflashpartitiontable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ciscoflashpartitiontable.EntityData)
}

// CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry
// An entry in the table of flash partition properties
// for each initialized flash partition. Each entry
// will be indexed by a device number and a partition
// number within the device.
type CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry struct {
    EntityData types.CommonEntityData
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

func (ciscoflashpartitionentry *CISCOFLASHMIB_Ciscoflashpartitiontable_Ciscoflashpartitionentry) GetEntityData() *types.CommonEntityData {
    ciscoflashpartitionentry.EntityData.YFilter = ciscoflashpartitionentry.YFilter
    ciscoflashpartitionentry.EntityData.YangName = "ciscoFlashPartitionEntry"
    ciscoflashpartitionentry.EntityData.BundleName = "cisco_ios_xe"
    ciscoflashpartitionentry.EntityData.ParentYangName = "ciscoFlashPartitionTable"
    ciscoflashpartitionentry.EntityData.SegmentPath = "ciscoFlashPartitionEntry" + "[ciscoFlashDeviceIndex='" + fmt.Sprintf("%v", ciscoflashpartitionentry.Ciscoflashdeviceindex) + "']" + "[ciscoFlashPartitionIndex='" + fmt.Sprintf("%v", ciscoflashpartitionentry.Ciscoflashpartitionindex) + "']"
    ciscoflashpartitionentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoflashpartitionentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoflashpartitionentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoflashpartitionentry.EntityData.Children = make(map[string]types.YChild)
    ciscoflashpartitionentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoflashpartitionentry.EntityData.Leafs["ciscoFlashDeviceIndex"] = types.YLeaf{"Ciscoflashdeviceindex", ciscoflashpartitionentry.Ciscoflashdeviceindex}
    ciscoflashpartitionentry.EntityData.Leafs["ciscoFlashPartitionIndex"] = types.YLeaf{"Ciscoflashpartitionindex", ciscoflashpartitionentry.Ciscoflashpartitionindex}
    ciscoflashpartitionentry.EntityData.Leafs["ciscoFlashPartitionStartChip"] = types.YLeaf{"Ciscoflashpartitionstartchip", ciscoflashpartitionentry.Ciscoflashpartitionstartchip}
    ciscoflashpartitionentry.EntityData.Leafs["ciscoFlashPartitionEndChip"] = types.YLeaf{"Ciscoflashpartitionendchip", ciscoflashpartitionentry.Ciscoflashpartitionendchip}
    ciscoflashpartitionentry.EntityData.Leafs["ciscoFlashPartitionSize"] = types.YLeaf{"Ciscoflashpartitionsize", ciscoflashpartitionentry.Ciscoflashpartitionsize}
    ciscoflashpartitionentry.EntityData.Leafs["ciscoFlashPartitionFreeSpace"] = types.YLeaf{"Ciscoflashpartitionfreespace", ciscoflashpartitionentry.Ciscoflashpartitionfreespace}
    ciscoflashpartitionentry.EntityData.Leafs["ciscoFlashPartitionFileCount"] = types.YLeaf{"Ciscoflashpartitionfilecount", ciscoflashpartitionentry.Ciscoflashpartitionfilecount}
    ciscoflashpartitionentry.EntityData.Leafs["ciscoFlashPartitionChecksumAlgorithm"] = types.YLeaf{"Ciscoflashpartitionchecksumalgorithm", ciscoflashpartitionentry.Ciscoflashpartitionchecksumalgorithm}
    ciscoflashpartitionentry.EntityData.Leafs["ciscoFlashPartitionStatus"] = types.YLeaf{"Ciscoflashpartitionstatus", ciscoflashpartitionentry.Ciscoflashpartitionstatus}
    ciscoflashpartitionentry.EntityData.Leafs["ciscoFlashPartitionUpgradeMethod"] = types.YLeaf{"Ciscoflashpartitionupgrademethod", ciscoflashpartitionentry.Ciscoflashpartitionupgrademethod}
    ciscoflashpartitionentry.EntityData.Leafs["ciscoFlashPartitionName"] = types.YLeaf{"Ciscoflashpartitionname", ciscoflashpartitionentry.Ciscoflashpartitionname}
    ciscoflashpartitionentry.EntityData.Leafs["ciscoFlashPartitionNeedErasure"] = types.YLeaf{"Ciscoflashpartitionneederasure", ciscoflashpartitionentry.Ciscoflashpartitionneederasure}
    ciscoflashpartitionentry.EntityData.Leafs["ciscoFlashPartitionFileNameLength"] = types.YLeaf{"Ciscoflashpartitionfilenamelength", ciscoflashpartitionentry.Ciscoflashpartitionfilenamelength}
    ciscoflashpartitionentry.EntityData.Leafs["ciscoFlashPartitionSizeExtended"] = types.YLeaf{"Ciscoflashpartitionsizeextended", ciscoflashpartitionentry.Ciscoflashpartitionsizeextended}
    ciscoflashpartitionentry.EntityData.Leafs["ciscoFlashPartitionFreeSpaceExtended"] = types.YLeaf{"Ciscoflashpartitionfreespaceextended", ciscoflashpartitionentry.Ciscoflashpartitionfreespaceextended}
    ciscoflashpartitionentry.EntityData.Leafs["ciscoFlashPartitionLowSpaceNotifThreshold"] = types.YLeaf{"Ciscoflashpartitionlowspacenotifthreshold", ciscoflashpartitionentry.Ciscoflashpartitionlowspacenotifthreshold}
    return &(ciscoflashpartitionentry.EntityData)
}

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
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table of Flash file properties for each initialized Flash
    // partition. Each entry represents a file and gives details about the file.
    // An entry is indexed using the device number, partition number within the
    // device, and file number within the partition. The type is slice of
    // CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry.
    Ciscoflashfileentry []CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry
}

func (ciscoflashfiletable *CISCOFLASHMIB_Ciscoflashfiletable) GetEntityData() *types.CommonEntityData {
    ciscoflashfiletable.EntityData.YFilter = ciscoflashfiletable.YFilter
    ciscoflashfiletable.EntityData.YangName = "ciscoFlashFileTable"
    ciscoflashfiletable.EntityData.BundleName = "cisco_ios_xe"
    ciscoflashfiletable.EntityData.ParentYangName = "CISCO-FLASH-MIB"
    ciscoflashfiletable.EntityData.SegmentPath = "ciscoFlashFileTable"
    ciscoflashfiletable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoflashfiletable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoflashfiletable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoflashfiletable.EntityData.Children = make(map[string]types.YChild)
    ciscoflashfiletable.EntityData.Children["ciscoFlashFileEntry"] = types.YChild{"Ciscoflashfileentry", nil}
    for i := range ciscoflashfiletable.Ciscoflashfileentry {
        ciscoflashfiletable.EntityData.Children[types.GetSegmentPath(&ciscoflashfiletable.Ciscoflashfileentry[i])] = types.YChild{"Ciscoflashfileentry", &ciscoflashfiletable.Ciscoflashfileentry[i]}
    }
    ciscoflashfiletable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ciscoflashfiletable.EntityData)
}

// CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry
// An entry in the table of Flash file properties
// for each initialized Flash partition. Each entry
// represents a file and gives details about the file.
// An entry is indexed using the device number,
// partition number within the device, and file
// number within the partition.
type CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry struct {
    EntityData types.CommonEntityData
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

func (ciscoflashfileentry *CISCOFLASHMIB_Ciscoflashfiletable_Ciscoflashfileentry) GetEntityData() *types.CommonEntityData {
    ciscoflashfileentry.EntityData.YFilter = ciscoflashfileentry.YFilter
    ciscoflashfileentry.EntityData.YangName = "ciscoFlashFileEntry"
    ciscoflashfileentry.EntityData.BundleName = "cisco_ios_xe"
    ciscoflashfileentry.EntityData.ParentYangName = "ciscoFlashFileTable"
    ciscoflashfileentry.EntityData.SegmentPath = "ciscoFlashFileEntry" + "[ciscoFlashDeviceIndex='" + fmt.Sprintf("%v", ciscoflashfileentry.Ciscoflashdeviceindex) + "']" + "[ciscoFlashPartitionIndex='" + fmt.Sprintf("%v", ciscoflashfileentry.Ciscoflashpartitionindex) + "']" + "[ciscoFlashFileIndex='" + fmt.Sprintf("%v", ciscoflashfileentry.Ciscoflashfileindex) + "']"
    ciscoflashfileentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoflashfileentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoflashfileentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoflashfileentry.EntityData.Children = make(map[string]types.YChild)
    ciscoflashfileentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoflashfileentry.EntityData.Leafs["ciscoFlashDeviceIndex"] = types.YLeaf{"Ciscoflashdeviceindex", ciscoflashfileentry.Ciscoflashdeviceindex}
    ciscoflashfileentry.EntityData.Leafs["ciscoFlashPartitionIndex"] = types.YLeaf{"Ciscoflashpartitionindex", ciscoflashfileentry.Ciscoflashpartitionindex}
    ciscoflashfileentry.EntityData.Leafs["ciscoFlashFileIndex"] = types.YLeaf{"Ciscoflashfileindex", ciscoflashfileentry.Ciscoflashfileindex}
    ciscoflashfileentry.EntityData.Leafs["ciscoFlashFileSize"] = types.YLeaf{"Ciscoflashfilesize", ciscoflashfileentry.Ciscoflashfilesize}
    ciscoflashfileentry.EntityData.Leafs["ciscoFlashFileChecksum"] = types.YLeaf{"Ciscoflashfilechecksum", ciscoflashfileentry.Ciscoflashfilechecksum}
    ciscoflashfileentry.EntityData.Leafs["ciscoFlashFileStatus"] = types.YLeaf{"Ciscoflashfilestatus", ciscoflashfileentry.Ciscoflashfilestatus}
    ciscoflashfileentry.EntityData.Leafs["ciscoFlashFileName"] = types.YLeaf{"Ciscoflashfilename", ciscoflashfileentry.Ciscoflashfilename}
    ciscoflashfileentry.EntityData.Leafs["ciscoFlashFileType"] = types.YLeaf{"Ciscoflashfiletype", ciscoflashfileentry.Ciscoflashfiletype}
    ciscoflashfileentry.EntityData.Leafs["ciscoFlashFileDate"] = types.YLeaf{"Ciscoflashfiledate", ciscoflashfileentry.Ciscoflashfiledate}
    return &(ciscoflashfileentry.EntityData)
}

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
    // CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry.
    Ciscoflashfilebytypeentry []CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry
}

func (ciscoflashfilebytypetable *CISCOFLASHMIB_Ciscoflashfilebytypetable) GetEntityData() *types.CommonEntityData {
    ciscoflashfilebytypetable.EntityData.YFilter = ciscoflashfilebytypetable.YFilter
    ciscoflashfilebytypetable.EntityData.YangName = "ciscoFlashFileByTypeTable"
    ciscoflashfilebytypetable.EntityData.BundleName = "cisco_ios_xe"
    ciscoflashfilebytypetable.EntityData.ParentYangName = "CISCO-FLASH-MIB"
    ciscoflashfilebytypetable.EntityData.SegmentPath = "ciscoFlashFileByTypeTable"
    ciscoflashfilebytypetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoflashfilebytypetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoflashfilebytypetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoflashfilebytypetable.EntityData.Children = make(map[string]types.YChild)
    ciscoflashfilebytypetable.EntityData.Children["ciscoFlashFileByTypeEntry"] = types.YChild{"Ciscoflashfilebytypeentry", nil}
    for i := range ciscoflashfilebytypetable.Ciscoflashfilebytypeentry {
        ciscoflashfilebytypetable.EntityData.Children[types.GetSegmentPath(&ciscoflashfilebytypetable.Ciscoflashfilebytypeentry[i])] = types.YChild{"Ciscoflashfilebytypeentry", &ciscoflashfilebytypetable.Ciscoflashfilebytypeentry[i]}
    }
    ciscoflashfilebytypetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ciscoflashfilebytypetable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (ciscoflashfilebytypeentry *CISCOFLASHMIB_Ciscoflashfilebytypetable_Ciscoflashfilebytypeentry) GetEntityData() *types.CommonEntityData {
    ciscoflashfilebytypeentry.EntityData.YFilter = ciscoflashfilebytypeentry.YFilter
    ciscoflashfilebytypeentry.EntityData.YangName = "ciscoFlashFileByTypeEntry"
    ciscoflashfilebytypeentry.EntityData.BundleName = "cisco_ios_xe"
    ciscoflashfilebytypeentry.EntityData.ParentYangName = "ciscoFlashFileByTypeTable"
    ciscoflashfilebytypeentry.EntityData.SegmentPath = "ciscoFlashFileByTypeEntry" + "[ciscoFlashFileType='" + fmt.Sprintf("%v", ciscoflashfilebytypeentry.Ciscoflashfiletype) + "']" + "[ciscoFlashDeviceIndex='" + fmt.Sprintf("%v", ciscoflashfilebytypeentry.Ciscoflashdeviceindex) + "']" + "[ciscoFlashPartitionIndex='" + fmt.Sprintf("%v", ciscoflashfilebytypeentry.Ciscoflashpartitionindex) + "']" + "[ciscoFlashFileIndex='" + fmt.Sprintf("%v", ciscoflashfilebytypeentry.Ciscoflashfileindex) + "']"
    ciscoflashfilebytypeentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoflashfilebytypeentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoflashfilebytypeentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoflashfilebytypeentry.EntityData.Children = make(map[string]types.YChild)
    ciscoflashfilebytypeentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoflashfilebytypeentry.EntityData.Leafs["ciscoFlashFileType"] = types.YLeaf{"Ciscoflashfiletype", ciscoflashfilebytypeentry.Ciscoflashfiletype}
    ciscoflashfilebytypeentry.EntityData.Leafs["ciscoFlashDeviceIndex"] = types.YLeaf{"Ciscoflashdeviceindex", ciscoflashfilebytypeentry.Ciscoflashdeviceindex}
    ciscoflashfilebytypeentry.EntityData.Leafs["ciscoFlashPartitionIndex"] = types.YLeaf{"Ciscoflashpartitionindex", ciscoflashfilebytypeentry.Ciscoflashpartitionindex}
    ciscoflashfilebytypeentry.EntityData.Leafs["ciscoFlashFileIndex"] = types.YLeaf{"Ciscoflashfileindex", ciscoflashfilebytypeentry.Ciscoflashfileindex}
    ciscoflashfilebytypeentry.EntityData.Leafs["ciscoFlashFileByTypeSize"] = types.YLeaf{"Ciscoflashfilebytypesize", ciscoflashfilebytypeentry.Ciscoflashfilebytypesize}
    ciscoflashfilebytypeentry.EntityData.Leafs["ciscoFlashFileByTypeChecksum"] = types.YLeaf{"Ciscoflashfilebytypechecksum", ciscoflashfilebytypeentry.Ciscoflashfilebytypechecksum}
    ciscoflashfilebytypeentry.EntityData.Leafs["ciscoFlashFileByTypeStatus"] = types.YLeaf{"Ciscoflashfilebytypestatus", ciscoflashfilebytypeentry.Ciscoflashfilebytypestatus}
    ciscoflashfilebytypeentry.EntityData.Leafs["ciscoFlashFileByTypeName"] = types.YLeaf{"Ciscoflashfilebytypename", ciscoflashfilebytypeentry.Ciscoflashfilebytypename}
    ciscoflashfilebytypeentry.EntityData.Leafs["ciscoFlashFileByTypeDate"] = types.YLeaf{"Ciscoflashfilebytypedate", ciscoflashfilebytypeentry.Ciscoflashfilebytypedate}
    return &(ciscoflashfilebytypeentry.EntityData)
}

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
    // CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry.
    Ciscoflashcopyentry []CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry
}

func (ciscoflashcopytable *CISCOFLASHMIB_Ciscoflashcopytable) GetEntityData() *types.CommonEntityData {
    ciscoflashcopytable.EntityData.YFilter = ciscoflashcopytable.YFilter
    ciscoflashcopytable.EntityData.YangName = "ciscoFlashCopyTable"
    ciscoflashcopytable.EntityData.BundleName = "cisco_ios_xe"
    ciscoflashcopytable.EntityData.ParentYangName = "CISCO-FLASH-MIB"
    ciscoflashcopytable.EntityData.SegmentPath = "ciscoFlashCopyTable"
    ciscoflashcopytable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoflashcopytable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoflashcopytable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoflashcopytable.EntityData.Children = make(map[string]types.YChild)
    ciscoflashcopytable.EntityData.Children["ciscoFlashCopyEntry"] = types.YChild{"Ciscoflashcopyentry", nil}
    for i := range ciscoflashcopytable.Ciscoflashcopyentry {
        ciscoflashcopytable.EntityData.Children[types.GetSegmentPath(&ciscoflashcopytable.Ciscoflashcopyentry[i])] = types.YChild{"Ciscoflashcopyentry", &ciscoflashcopytable.Ciscoflashcopyentry[i]}
    }
    ciscoflashcopytable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ciscoflashcopytable.EntityData)
}

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
    EntityData types.CommonEntityData
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
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

func (ciscoflashcopyentry *CISCOFLASHMIB_Ciscoflashcopytable_Ciscoflashcopyentry) GetEntityData() *types.CommonEntityData {
    ciscoflashcopyentry.EntityData.YFilter = ciscoflashcopyentry.YFilter
    ciscoflashcopyentry.EntityData.YangName = "ciscoFlashCopyEntry"
    ciscoflashcopyentry.EntityData.BundleName = "cisco_ios_xe"
    ciscoflashcopyentry.EntityData.ParentYangName = "ciscoFlashCopyTable"
    ciscoflashcopyentry.EntityData.SegmentPath = "ciscoFlashCopyEntry" + "[ciscoFlashCopySerialNumber='" + fmt.Sprintf("%v", ciscoflashcopyentry.Ciscoflashcopyserialnumber) + "']"
    ciscoflashcopyentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoflashcopyentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoflashcopyentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoflashcopyentry.EntityData.Children = make(map[string]types.YChild)
    ciscoflashcopyentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoflashcopyentry.EntityData.Leafs["ciscoFlashCopySerialNumber"] = types.YLeaf{"Ciscoflashcopyserialnumber", ciscoflashcopyentry.Ciscoflashcopyserialnumber}
    ciscoflashcopyentry.EntityData.Leafs["ciscoFlashCopyCommand"] = types.YLeaf{"Ciscoflashcopycommand", ciscoflashcopyentry.Ciscoflashcopycommand}
    ciscoflashcopyentry.EntityData.Leafs["ciscoFlashCopyProtocol"] = types.YLeaf{"Ciscoflashcopyprotocol", ciscoflashcopyentry.Ciscoflashcopyprotocol}
    ciscoflashcopyentry.EntityData.Leafs["ciscoFlashCopyServerAddress"] = types.YLeaf{"Ciscoflashcopyserveraddress", ciscoflashcopyentry.Ciscoflashcopyserveraddress}
    ciscoflashcopyentry.EntityData.Leafs["ciscoFlashCopySourceName"] = types.YLeaf{"Ciscoflashcopysourcename", ciscoflashcopyentry.Ciscoflashcopysourcename}
    ciscoflashcopyentry.EntityData.Leafs["ciscoFlashCopyDestinationName"] = types.YLeaf{"Ciscoflashcopydestinationname", ciscoflashcopyentry.Ciscoflashcopydestinationname}
    ciscoflashcopyentry.EntityData.Leafs["ciscoFlashCopyRemoteUserName"] = types.YLeaf{"Ciscoflashcopyremoteusername", ciscoflashcopyentry.Ciscoflashcopyremoteusername}
    ciscoflashcopyentry.EntityData.Leafs["ciscoFlashCopyStatus"] = types.YLeaf{"Ciscoflashcopystatus", ciscoflashcopyentry.Ciscoflashcopystatus}
    ciscoflashcopyentry.EntityData.Leafs["ciscoFlashCopyNotifyOnCompletion"] = types.YLeaf{"Ciscoflashcopynotifyoncompletion", ciscoflashcopyentry.Ciscoflashcopynotifyoncompletion}
    ciscoflashcopyentry.EntityData.Leafs["ciscoFlashCopyTime"] = types.YLeaf{"Ciscoflashcopytime", ciscoflashcopyentry.Ciscoflashcopytime}
    ciscoflashcopyentry.EntityData.Leafs["ciscoFlashCopyEntryStatus"] = types.YLeaf{"Ciscoflashcopyentrystatus", ciscoflashcopyentry.Ciscoflashcopyentrystatus}
    ciscoflashcopyentry.EntityData.Leafs["ciscoFlashCopyVerify"] = types.YLeaf{"Ciscoflashcopyverify", ciscoflashcopyentry.Ciscoflashcopyverify}
    ciscoflashcopyentry.EntityData.Leafs["ciscoFlashCopyServerAddrType"] = types.YLeaf{"Ciscoflashcopyserveraddrtype", ciscoflashcopyentry.Ciscoflashcopyserveraddrtype}
    ciscoflashcopyentry.EntityData.Leafs["ciscoFlashCopyServerAddrRev1"] = types.YLeaf{"Ciscoflashcopyserveraddrrev1", ciscoflashcopyentry.Ciscoflashcopyserveraddrrev1}
    ciscoflashcopyentry.EntityData.Leafs["ciscoFlashCopyRemotePassword"] = types.YLeaf{"Ciscoflashcopyremotepassword", ciscoflashcopyentry.Ciscoflashcopyremotepassword}
    return &(ciscoflashcopyentry.EntityData)
}

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
    // CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry.
    Ciscoflashpartitioningentry []CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry
}

func (ciscoflashpartitioningtable *CISCOFLASHMIB_Ciscoflashpartitioningtable) GetEntityData() *types.CommonEntityData {
    ciscoflashpartitioningtable.EntityData.YFilter = ciscoflashpartitioningtable.YFilter
    ciscoflashpartitioningtable.EntityData.YangName = "ciscoFlashPartitioningTable"
    ciscoflashpartitioningtable.EntityData.BundleName = "cisco_ios_xe"
    ciscoflashpartitioningtable.EntityData.ParentYangName = "CISCO-FLASH-MIB"
    ciscoflashpartitioningtable.EntityData.SegmentPath = "ciscoFlashPartitioningTable"
    ciscoflashpartitioningtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoflashpartitioningtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoflashpartitioningtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoflashpartitioningtable.EntityData.Children = make(map[string]types.YChild)
    ciscoflashpartitioningtable.EntityData.Children["ciscoFlashPartitioningEntry"] = types.YChild{"Ciscoflashpartitioningentry", nil}
    for i := range ciscoflashpartitioningtable.Ciscoflashpartitioningentry {
        ciscoflashpartitioningtable.EntityData.Children[types.GetSegmentPath(&ciscoflashpartitioningtable.Ciscoflashpartitioningentry[i])] = types.YChild{"Ciscoflashpartitioningentry", &ciscoflashpartitioningtable.Ciscoflashpartitioningentry[i]}
    }
    ciscoflashpartitioningtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ciscoflashpartitioningtable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (ciscoflashpartitioningentry *CISCOFLASHMIB_Ciscoflashpartitioningtable_Ciscoflashpartitioningentry) GetEntityData() *types.CommonEntityData {
    ciscoflashpartitioningentry.EntityData.YFilter = ciscoflashpartitioningentry.YFilter
    ciscoflashpartitioningentry.EntityData.YangName = "ciscoFlashPartitioningEntry"
    ciscoflashpartitioningentry.EntityData.BundleName = "cisco_ios_xe"
    ciscoflashpartitioningentry.EntityData.ParentYangName = "ciscoFlashPartitioningTable"
    ciscoflashpartitioningentry.EntityData.SegmentPath = "ciscoFlashPartitioningEntry" + "[ciscoFlashPartitioningSerialNumber='" + fmt.Sprintf("%v", ciscoflashpartitioningentry.Ciscoflashpartitioningserialnumber) + "']"
    ciscoflashpartitioningentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoflashpartitioningentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoflashpartitioningentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoflashpartitioningentry.EntityData.Children = make(map[string]types.YChild)
    ciscoflashpartitioningentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoflashpartitioningentry.EntityData.Leafs["ciscoFlashPartitioningSerialNumber"] = types.YLeaf{"Ciscoflashpartitioningserialnumber", ciscoflashpartitioningentry.Ciscoflashpartitioningserialnumber}
    ciscoflashpartitioningentry.EntityData.Leafs["ciscoFlashPartitioningCommand"] = types.YLeaf{"Ciscoflashpartitioningcommand", ciscoflashpartitioningentry.Ciscoflashpartitioningcommand}
    ciscoflashpartitioningentry.EntityData.Leafs["ciscoFlashPartitioningDestinationName"] = types.YLeaf{"Ciscoflashpartitioningdestinationname", ciscoflashpartitioningentry.Ciscoflashpartitioningdestinationname}
    ciscoflashpartitioningentry.EntityData.Leafs["ciscoFlashPartitioningPartitionCount"] = types.YLeaf{"Ciscoflashpartitioningpartitioncount", ciscoflashpartitioningentry.Ciscoflashpartitioningpartitioncount}
    ciscoflashpartitioningentry.EntityData.Leafs["ciscoFlashPartitioningPartitionSizes"] = types.YLeaf{"Ciscoflashpartitioningpartitionsizes", ciscoflashpartitioningentry.Ciscoflashpartitioningpartitionsizes}
    ciscoflashpartitioningentry.EntityData.Leafs["ciscoFlashPartitioningStatus"] = types.YLeaf{"Ciscoflashpartitioningstatus", ciscoflashpartitioningentry.Ciscoflashpartitioningstatus}
    ciscoflashpartitioningentry.EntityData.Leafs["ciscoFlashPartitioningNotifyOnCompletion"] = types.YLeaf{"Ciscoflashpartitioningnotifyoncompletion", ciscoflashpartitioningentry.Ciscoflashpartitioningnotifyoncompletion}
    ciscoflashpartitioningentry.EntityData.Leafs["ciscoFlashPartitioningTime"] = types.YLeaf{"Ciscoflashpartitioningtime", ciscoflashpartitioningentry.Ciscoflashpartitioningtime}
    ciscoflashpartitioningentry.EntityData.Leafs["ciscoFlashPartitioningEntryStatus"] = types.YLeaf{"Ciscoflashpartitioningentrystatus", ciscoflashpartitioningentry.Ciscoflashpartitioningentrystatus}
    return &(ciscoflashpartitioningentry.EntityData)
}

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
    // CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry.
    Ciscoflashmiscopentry []CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry
}

func (ciscoflashmiscoptable *CISCOFLASHMIB_Ciscoflashmiscoptable) GetEntityData() *types.CommonEntityData {
    ciscoflashmiscoptable.EntityData.YFilter = ciscoflashmiscoptable.YFilter
    ciscoflashmiscoptable.EntityData.YangName = "ciscoFlashMiscOpTable"
    ciscoflashmiscoptable.EntityData.BundleName = "cisco_ios_xe"
    ciscoflashmiscoptable.EntityData.ParentYangName = "CISCO-FLASH-MIB"
    ciscoflashmiscoptable.EntityData.SegmentPath = "ciscoFlashMiscOpTable"
    ciscoflashmiscoptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoflashmiscoptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoflashmiscoptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoflashmiscoptable.EntityData.Children = make(map[string]types.YChild)
    ciscoflashmiscoptable.EntityData.Children["ciscoFlashMiscOpEntry"] = types.YChild{"Ciscoflashmiscopentry", nil}
    for i := range ciscoflashmiscoptable.Ciscoflashmiscopentry {
        ciscoflashmiscoptable.EntityData.Children[types.GetSegmentPath(&ciscoflashmiscoptable.Ciscoflashmiscopentry[i])] = types.YChild{"Ciscoflashmiscopentry", &ciscoflashmiscoptable.Ciscoflashmiscopentry[i]}
    }
    ciscoflashmiscoptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(ciscoflashmiscoptable.EntityData)
}

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
    EntityData types.CommonEntityData
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

func (ciscoflashmiscopentry *CISCOFLASHMIB_Ciscoflashmiscoptable_Ciscoflashmiscopentry) GetEntityData() *types.CommonEntityData {
    ciscoflashmiscopentry.EntityData.YFilter = ciscoflashmiscopentry.YFilter
    ciscoflashmiscopentry.EntityData.YangName = "ciscoFlashMiscOpEntry"
    ciscoflashmiscopentry.EntityData.BundleName = "cisco_ios_xe"
    ciscoflashmiscopentry.EntityData.ParentYangName = "ciscoFlashMiscOpTable"
    ciscoflashmiscopentry.EntityData.SegmentPath = "ciscoFlashMiscOpEntry" + "[ciscoFlashMiscOpSerialNumber='" + fmt.Sprintf("%v", ciscoflashmiscopentry.Ciscoflashmiscopserialnumber) + "']"
    ciscoflashmiscopentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoflashmiscopentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoflashmiscopentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoflashmiscopentry.EntityData.Children = make(map[string]types.YChild)
    ciscoflashmiscopentry.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoflashmiscopentry.EntityData.Leafs["ciscoFlashMiscOpSerialNumber"] = types.YLeaf{"Ciscoflashmiscopserialnumber", ciscoflashmiscopentry.Ciscoflashmiscopserialnumber}
    ciscoflashmiscopentry.EntityData.Leafs["ciscoFlashMiscOpCommand"] = types.YLeaf{"Ciscoflashmiscopcommand", ciscoflashmiscopentry.Ciscoflashmiscopcommand}
    ciscoflashmiscopentry.EntityData.Leafs["ciscoFlashMiscOpDestinationName"] = types.YLeaf{"Ciscoflashmiscopdestinationname", ciscoflashmiscopentry.Ciscoflashmiscopdestinationname}
    ciscoflashmiscopentry.EntityData.Leafs["ciscoFlashMiscOpStatus"] = types.YLeaf{"Ciscoflashmiscopstatus", ciscoflashmiscopentry.Ciscoflashmiscopstatus}
    ciscoflashmiscopentry.EntityData.Leafs["ciscoFlashMiscOpNotifyOnCompletion"] = types.YLeaf{"Ciscoflashmiscopnotifyoncompletion", ciscoflashmiscopentry.Ciscoflashmiscopnotifyoncompletion}
    ciscoflashmiscopentry.EntityData.Leafs["ciscoFlashMiscOpTime"] = types.YLeaf{"Ciscoflashmiscoptime", ciscoflashmiscopentry.Ciscoflashmiscoptime}
    ciscoflashmiscopentry.EntityData.Leafs["ciscoFlashMiscOpEntryStatus"] = types.YLeaf{"Ciscoflashmiscopentrystatus", ciscoflashmiscopentry.Ciscoflashmiscopentrystatus}
    return &(ciscoflashmiscopentry.EntityData)
}

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

