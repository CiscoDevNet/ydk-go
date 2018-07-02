// The MIB module for managing the running image level
// of a Cisco device. Cisco's licensing mechanism provides
// flexibility to run a device on a chosen image level.
// This mechanism is referred to as image level licensing.
// Image level licensing leverages the universal image
// based licensing solution.
// 
// The image level licensing mechanism works as follows - 
// 
// A universal image that contains all levels of software
// packages is loaded on to the device. At boot time, the
// device determines the highest level of license and brings
// up the appropriate software features or subsystems.
// The user can configure the image level with which the
// device has to boot. The system will verify whether the
// appropriate license is available for the configured image
// level. The image level for the next boot will be determined
// based on the availability of the license. The following
// scenarios explains some use-cases of image level licensing:
// 
// Scenario 1:
// - Customer selects advsecurityk9 based image.
// - Manufacturing installs advsecurity license on the device.
// - This device will run all features that are part of the
// base advsecurity license.
// - Customer upgrades to advipservicesk9 license.
// - The next boot level is set to advipservicesk9.
// - The device will run advsecurityk9 feature until the
// next reboot. After reboot the device will run 
// advipservicesk9 features.
// 
// Scenario 2:
// - Customer selects advipservicesk9 based image.
// - Manufacturing installs advipservices and advsecurity
// license on the device.
// - This device will run all features that are part of the
// base advipservices license.
// - No upgrades available for advipservices license.
// 
// The user has to accept the End User License Agreement(EULA)
// before using this MIB to configure the image level. 
// 
// This MIB should be used in conjuntion with
// CISCO-LICENSE-MGMT-MIB module to achieve the image level
// licensing functionality.
// 
// This MIB module defines objects which provides the different
// image levels supported by the device and the license required
// to enable a particular image level. It also defines objects
// to let the user configure the required image level. The MIB 
// module contains notification which will be triggered when
// the user changes the image level for next boot. 
// 
// The CISCO-LICENSE-MGMT-MIB module should be used to export
// the EULA and to configure the required license.
// 
// This MIB module is defined generically so it can be used for
// both stand-alone as well as stackable devices. The
// entPhysicalIndex imported from ENTITY-MIB is used to identify
// the device uniquely.
package cisco_image_license_mgmt_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_image_license_mgmt_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-IMAGE-LICENSE-MGMT-MIB CISCO-IMAGE-LICENSE-MGMT-MIB}", reflect.TypeOf(CISCOIMAGELICENSEMGMTMIB{}))
    ydk.RegisterEntity("CISCO-IMAGE-LICENSE-MGMT-MIB:CISCO-IMAGE-LICENSE-MGMT-MIB", reflect.TypeOf(CISCOIMAGELICENSEMGMTMIB{}))
}

// CISCOIMAGELICENSEMGMTMIB
type CISCOIMAGELICENSEMGMTMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CiscoImageLicenseMgmtMIBObjects CISCOIMAGELICENSEMGMTMIB_CiscoImageLicenseMgmtMIBObjects

    
    CilmNotifCntl CISCOIMAGELICENSEMGMTMIB_CilmNotifCntl

    // A table that contains the configuration information of current and next
    // boot image level. This table contains entries for each software module
    // running in an image  loaded in the device. The software module is
    // identified by cilmModuleName and the device is identified by 
    // entPhysicalIndex.
    CilmBootImageLevelTable CISCOIMAGELICENSEMGMTMIB_CilmBootImageLevelTable

    // This table contains the mapping between different image levels of each
    // modules in the image and the license required to run the modules at a
    // particular image level. This table can be used to identify the different
    // image levels and the appropriate licenses  required for each.
    CilmImageLevelToLicenseMapTable CISCOIMAGELICENSEMGMTMIB_CilmImageLevelToLicenseMapTable
}

func (cISCOIMAGELICENSEMGMTMIB *CISCOIMAGELICENSEMGMTMIB) GetEntityData() *types.CommonEntityData {
    cISCOIMAGELICENSEMGMTMIB.EntityData.YFilter = cISCOIMAGELICENSEMGMTMIB.YFilter
    cISCOIMAGELICENSEMGMTMIB.EntityData.YangName = "CISCO-IMAGE-LICENSE-MGMT-MIB"
    cISCOIMAGELICENSEMGMTMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOIMAGELICENSEMGMTMIB.EntityData.ParentYangName = "CISCO-IMAGE-LICENSE-MGMT-MIB"
    cISCOIMAGELICENSEMGMTMIB.EntityData.SegmentPath = "CISCO-IMAGE-LICENSE-MGMT-MIB:CISCO-IMAGE-LICENSE-MGMT-MIB"
    cISCOIMAGELICENSEMGMTMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOIMAGELICENSEMGMTMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOIMAGELICENSEMGMTMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOIMAGELICENSEMGMTMIB.EntityData.Children = types.NewOrderedMap()
    cISCOIMAGELICENSEMGMTMIB.EntityData.Children.Append("ciscoImageLicenseMgmtMIBObjects", types.YChild{"CiscoImageLicenseMgmtMIBObjects", &cISCOIMAGELICENSEMGMTMIB.CiscoImageLicenseMgmtMIBObjects})
    cISCOIMAGELICENSEMGMTMIB.EntityData.Children.Append("cilmNotifCntl", types.YChild{"CilmNotifCntl", &cISCOIMAGELICENSEMGMTMIB.CilmNotifCntl})
    cISCOIMAGELICENSEMGMTMIB.EntityData.Children.Append("cilmBootImageLevelTable", types.YChild{"CilmBootImageLevelTable", &cISCOIMAGELICENSEMGMTMIB.CilmBootImageLevelTable})
    cISCOIMAGELICENSEMGMTMIB.EntityData.Children.Append("cilmImageLevelToLicenseMapTable", types.YChild{"CilmImageLevelToLicenseMapTable", &cISCOIMAGELICENSEMGMTMIB.CilmImageLevelToLicenseMapTable})
    cISCOIMAGELICENSEMGMTMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOIMAGELICENSEMGMTMIB.EntityData.YListKeys = []string {}

    return &(cISCOIMAGELICENSEMGMTMIB.EntityData)
}

// CISCOIMAGELICENSEMGMTMIB_CiscoImageLicenseMgmtMIBObjects
type CISCOIMAGELICENSEMGMTMIB_CiscoImageLicenseMgmtMIBObjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object when set to TRUE means that the user has accepted the END USER
    // LICENSE AGREEMENT. This object has to be set to TRUE by the user before
    // using the objects in the cilmBootImageLevelTable to configure the license.
    // The type is bool.
    CilmEULAAccepted interface{}
}

func (ciscoImageLicenseMgmtMIBObjects *CISCOIMAGELICENSEMGMTMIB_CiscoImageLicenseMgmtMIBObjects) GetEntityData() *types.CommonEntityData {
    ciscoImageLicenseMgmtMIBObjects.EntityData.YFilter = ciscoImageLicenseMgmtMIBObjects.YFilter
    ciscoImageLicenseMgmtMIBObjects.EntityData.YangName = "ciscoImageLicenseMgmtMIBObjects"
    ciscoImageLicenseMgmtMIBObjects.EntityData.BundleName = "cisco_ios_xe"
    ciscoImageLicenseMgmtMIBObjects.EntityData.ParentYangName = "CISCO-IMAGE-LICENSE-MGMT-MIB"
    ciscoImageLicenseMgmtMIBObjects.EntityData.SegmentPath = "ciscoImageLicenseMgmtMIBObjects"
    ciscoImageLicenseMgmtMIBObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoImageLicenseMgmtMIBObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoImageLicenseMgmtMIBObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoImageLicenseMgmtMIBObjects.EntityData.Children = types.NewOrderedMap()
    ciscoImageLicenseMgmtMIBObjects.EntityData.Leafs = types.NewOrderedMap()
    ciscoImageLicenseMgmtMIBObjects.EntityData.Leafs.Append("cilmEULAAccepted", types.YLeaf{"CilmEULAAccepted", ciscoImageLicenseMgmtMIBObjects.CilmEULAAccepted})

    ciscoImageLicenseMgmtMIBObjects.EntityData.YListKeys = []string {}

    return &(ciscoImageLicenseMgmtMIBObjects.EntityData)
}

// CISCOIMAGELICENSEMGMTMIB_CilmNotifCntl
type CISCOIMAGELICENSEMGMTMIB_CilmNotifCntl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify whether or not a notification should be generated on the detection
    // of change in next boot image level.  If set to TRUE,
    // cilmBootImageLevelChanged notification will be generated. It is the
    // responsibility of the management entity to ensure that the SNMP
    // administrative model is configured in such a way as to allow the 
    // notification to be delivered. The type is bool.
    CilmImageLevelChangedNotif interface{}
}

func (cilmNotifCntl *CISCOIMAGELICENSEMGMTMIB_CilmNotifCntl) GetEntityData() *types.CommonEntityData {
    cilmNotifCntl.EntityData.YFilter = cilmNotifCntl.YFilter
    cilmNotifCntl.EntityData.YangName = "cilmNotifCntl"
    cilmNotifCntl.EntityData.BundleName = "cisco_ios_xe"
    cilmNotifCntl.EntityData.ParentYangName = "CISCO-IMAGE-LICENSE-MGMT-MIB"
    cilmNotifCntl.EntityData.SegmentPath = "cilmNotifCntl"
    cilmNotifCntl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cilmNotifCntl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cilmNotifCntl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cilmNotifCntl.EntityData.Children = types.NewOrderedMap()
    cilmNotifCntl.EntityData.Leafs = types.NewOrderedMap()
    cilmNotifCntl.EntityData.Leafs.Append("cilmImageLevelChangedNotif", types.YLeaf{"CilmImageLevelChangedNotif", cilmNotifCntl.CilmImageLevelChangedNotif})

    cilmNotifCntl.EntityData.YListKeys = []string {}

    return &(cilmNotifCntl.EntityData)
}

// CISCOIMAGELICENSEMGMTMIB_CilmBootImageLevelTable
// A table that contains the configuration information of
// current and next boot image level. This table contains
// entries for each software module running in an image 
// loaded in the device. The software module is identified by
// cilmModuleName and the device is identified by 
// entPhysicalIndex.
type CISCOIMAGELICENSEMGMTMIB_CilmBootImageLevelTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table for each module containing the list of objects that
    // define the configuration of next boot level. The following information is
    // specified by the objects present in the table.  - Current image level. -
    // Configured image level for the next boot. - Actual image level for the next
    // boot. - License store index for the current license. - License index of the
    // current license. - License store index for the next boot license. - License
    // index of the next boot license. The type is slice of
    // CISCOIMAGELICENSEMGMTMIB_CilmBootImageLevelTable_CilmBootImageLevelEntry.
    CilmBootImageLevelEntry []*CISCOIMAGELICENSEMGMTMIB_CilmBootImageLevelTable_CilmBootImageLevelEntry
}

func (cilmBootImageLevelTable *CISCOIMAGELICENSEMGMTMIB_CilmBootImageLevelTable) GetEntityData() *types.CommonEntityData {
    cilmBootImageLevelTable.EntityData.YFilter = cilmBootImageLevelTable.YFilter
    cilmBootImageLevelTable.EntityData.YangName = "cilmBootImageLevelTable"
    cilmBootImageLevelTable.EntityData.BundleName = "cisco_ios_xe"
    cilmBootImageLevelTable.EntityData.ParentYangName = "CISCO-IMAGE-LICENSE-MGMT-MIB"
    cilmBootImageLevelTable.EntityData.SegmentPath = "cilmBootImageLevelTable"
    cilmBootImageLevelTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cilmBootImageLevelTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cilmBootImageLevelTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cilmBootImageLevelTable.EntityData.Children = types.NewOrderedMap()
    cilmBootImageLevelTable.EntityData.Children.Append("cilmBootImageLevelEntry", types.YChild{"CilmBootImageLevelEntry", nil})
    for i := range cilmBootImageLevelTable.CilmBootImageLevelEntry {
        cilmBootImageLevelTable.EntityData.Children.Append(types.GetSegmentPath(cilmBootImageLevelTable.CilmBootImageLevelEntry[i]), types.YChild{"CilmBootImageLevelEntry", cilmBootImageLevelTable.CilmBootImageLevelEntry[i]})
    }
    cilmBootImageLevelTable.EntityData.Leafs = types.NewOrderedMap()

    cilmBootImageLevelTable.EntityData.YListKeys = []string {}

    return &(cilmBootImageLevelTable.EntityData)
}

// CISCOIMAGELICENSEMGMTMIB_CilmBootImageLevelTable_CilmBootImageLevelEntry
// An entry in the table for each module containing the
// list of objects that define the configuration of next
// boot level. The following information is specified by
// the objects present in the table.
// 
// - Current image level.
// - Configured image level for the next boot.
// - Actual image level for the next boot.
// - License store index for the current license.
// - License index of the current license.
// - License store index for the next boot license.
// - License index of the next boot license.
type CISCOIMAGELICENSEMGMTMIB_CilmBootImageLevelTable_CilmBootImageLevelEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. This object is used as one of the two indices in
    // cilmBootImageLevelTable. This object indicates the module name of the
    // software package. There can be multiple modules in an image performing
    // specific functionality. For example, in a wireless image there can be two
    // modules - a base image module and a wireless module. The type is string.
    CilmModuleName interface{}

    // This object indicates the current image level that the module is running.
    // The type is string with length: 0..255.
    CilmCurrentImageLevel interface{}

    // This object indicates the configured image level of the module for the next
    // boot.  Note: The configured next boot image level may not  be the actual
    // next boot image level. The actual next boot image level is denoted by
    // cilmNextBootImageLevel which is determined based on the license
    // availability. The type is string with length: 0..255.
    CilmConfiguredBootImageLevel interface{}

    // This object indicates the next boot image level. The next boot image level
    // can be different from configured level. The next boot image level is
    // determined based on the availability of required license. The type is
    // string with length: 0..255.
    CilmNextBootImageLevel interface{}

    // This object indicates the license store index where the currently used
    // license is stored. This object has the same value as
    // clmgmtLicenseStoreIndex object and uniquely identifies an entry in
    // clmgmtLicenseStoreInfoTable in CISCO-LICENSE-MGMT-MIB.  Note: The license
    // store index can be '0' if no license is installed and device is running
    // base image. The type is interface{} with range: 0..4294967295.
    CilmCurrentLicenseStoreIndex interface{}

    // This object indicates the license index of the currently used license. This
    // object has the same value as clmgmtLicenseIndex and uniquely identifies an
    // entry in clmgmtLicenseInfoTable in CISCO-LICENSE-MGMT-MIB.  Note: The
    // license index can be '0' if no license is installed and device is running
    // base image. The type is interface{} with range: 0..4294967295.
    CilmCurrentLicenseIndex interface{}

    // This object indicates the license store index where the next boot license
    // is stored. This object has the same value as clmgmtLicenseStoreIndex object
    // and uniquely identifies an entry in clmgmtLicenseStoreInfoTable in
    // CISCO-LICENSE-MGMT-MIB.  Note: The license store index can be '0' if no
    // license is installed for the next boot. The type is interface{} with range:
    // 0..4294967295.
    CilmNextBootLicenseStoreIndex interface{}

    // This object indicates the license index of the next boot license. This
    // object has the same value as clmgmtLicenseIndex and uniquely identifies an
    // entry in clmgmtLicenseInfoTable in CISCO-LICENSE-MGMT-MIB.  Note: The
    // license index can be '0' if no license is installed for the next boot. The
    // type is interface{} with range: 0..4294967295.
    CilmNextBootLicenseIndex interface{}
}

func (cilmBootImageLevelEntry *CISCOIMAGELICENSEMGMTMIB_CilmBootImageLevelTable_CilmBootImageLevelEntry) GetEntityData() *types.CommonEntityData {
    cilmBootImageLevelEntry.EntityData.YFilter = cilmBootImageLevelEntry.YFilter
    cilmBootImageLevelEntry.EntityData.YangName = "cilmBootImageLevelEntry"
    cilmBootImageLevelEntry.EntityData.BundleName = "cisco_ios_xe"
    cilmBootImageLevelEntry.EntityData.ParentYangName = "cilmBootImageLevelTable"
    cilmBootImageLevelEntry.EntityData.SegmentPath = "cilmBootImageLevelEntry" + types.AddKeyToken(cilmBootImageLevelEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(cilmBootImageLevelEntry.CilmModuleName, "cilmModuleName")
    cilmBootImageLevelEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cilmBootImageLevelEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cilmBootImageLevelEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cilmBootImageLevelEntry.EntityData.Children = types.NewOrderedMap()
    cilmBootImageLevelEntry.EntityData.Leafs = types.NewOrderedMap()
    cilmBootImageLevelEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cilmBootImageLevelEntry.EntPhysicalIndex})
    cilmBootImageLevelEntry.EntityData.Leafs.Append("cilmModuleName", types.YLeaf{"CilmModuleName", cilmBootImageLevelEntry.CilmModuleName})
    cilmBootImageLevelEntry.EntityData.Leafs.Append("cilmCurrentImageLevel", types.YLeaf{"CilmCurrentImageLevel", cilmBootImageLevelEntry.CilmCurrentImageLevel})
    cilmBootImageLevelEntry.EntityData.Leafs.Append("cilmConfiguredBootImageLevel", types.YLeaf{"CilmConfiguredBootImageLevel", cilmBootImageLevelEntry.CilmConfiguredBootImageLevel})
    cilmBootImageLevelEntry.EntityData.Leafs.Append("cilmNextBootImageLevel", types.YLeaf{"CilmNextBootImageLevel", cilmBootImageLevelEntry.CilmNextBootImageLevel})
    cilmBootImageLevelEntry.EntityData.Leafs.Append("cilmCurrentLicenseStoreIndex", types.YLeaf{"CilmCurrentLicenseStoreIndex", cilmBootImageLevelEntry.CilmCurrentLicenseStoreIndex})
    cilmBootImageLevelEntry.EntityData.Leafs.Append("cilmCurrentLicenseIndex", types.YLeaf{"CilmCurrentLicenseIndex", cilmBootImageLevelEntry.CilmCurrentLicenseIndex})
    cilmBootImageLevelEntry.EntityData.Leafs.Append("cilmNextBootLicenseStoreIndex", types.YLeaf{"CilmNextBootLicenseStoreIndex", cilmBootImageLevelEntry.CilmNextBootLicenseStoreIndex})
    cilmBootImageLevelEntry.EntityData.Leafs.Append("cilmNextBootLicenseIndex", types.YLeaf{"CilmNextBootLicenseIndex", cilmBootImageLevelEntry.CilmNextBootLicenseIndex})

    cilmBootImageLevelEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "CilmModuleName"}

    return &(cilmBootImageLevelEntry.EntityData)
}

// CISCOIMAGELICENSEMGMTMIB_CilmImageLevelToLicenseMapTable
// This table contains the mapping between different
// image levels of each modules in the image and the
// license required to run the modules at a particular
// image level. This table can be used to identify the
// different image levels and the appropriate licenses 
// required for each.
type CISCOIMAGELICENSEMGMTMIB_CilmImageLevelToLicenseMapTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table containing the following information. - The image
    // levels at the which the modules can be run. - The license required to the
    // run a module at a particular image level. - The priority of the license.
    // The type is slice of
    // CISCOIMAGELICENSEMGMTMIB_CilmImageLevelToLicenseMapTable_CilmImageLevelToLicenseMapEntry.
    CilmImageLevelToLicenseMapEntry []*CISCOIMAGELICENSEMGMTMIB_CilmImageLevelToLicenseMapTable_CilmImageLevelToLicenseMapEntry
}

func (cilmImageLevelToLicenseMapTable *CISCOIMAGELICENSEMGMTMIB_CilmImageLevelToLicenseMapTable) GetEntityData() *types.CommonEntityData {
    cilmImageLevelToLicenseMapTable.EntityData.YFilter = cilmImageLevelToLicenseMapTable.YFilter
    cilmImageLevelToLicenseMapTable.EntityData.YangName = "cilmImageLevelToLicenseMapTable"
    cilmImageLevelToLicenseMapTable.EntityData.BundleName = "cisco_ios_xe"
    cilmImageLevelToLicenseMapTable.EntityData.ParentYangName = "CISCO-IMAGE-LICENSE-MGMT-MIB"
    cilmImageLevelToLicenseMapTable.EntityData.SegmentPath = "cilmImageLevelToLicenseMapTable"
    cilmImageLevelToLicenseMapTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cilmImageLevelToLicenseMapTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cilmImageLevelToLicenseMapTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cilmImageLevelToLicenseMapTable.EntityData.Children = types.NewOrderedMap()
    cilmImageLevelToLicenseMapTable.EntityData.Children.Append("cilmImageLevelToLicenseMapEntry", types.YChild{"CilmImageLevelToLicenseMapEntry", nil})
    for i := range cilmImageLevelToLicenseMapTable.CilmImageLevelToLicenseMapEntry {
        cilmImageLevelToLicenseMapTable.EntityData.Children.Append(types.GetSegmentPath(cilmImageLevelToLicenseMapTable.CilmImageLevelToLicenseMapEntry[i]), types.YChild{"CilmImageLevelToLicenseMapEntry", cilmImageLevelToLicenseMapTable.CilmImageLevelToLicenseMapEntry[i]})
    }
    cilmImageLevelToLicenseMapTable.EntityData.Leafs = types.NewOrderedMap()

    cilmImageLevelToLicenseMapTable.EntityData.YListKeys = []string {}

    return &(cilmImageLevelToLicenseMapTable.EntityData)
}

// CISCOIMAGELICENSEMGMTMIB_CilmImageLevelToLicenseMapTable_CilmImageLevelToLicenseMapEntry
// An entry in the table containing the following
// information.
// - The image levels at the which the modules can be run.
// - The license required to the run a module at a
// particular image level.
// - The priority of the license.
type CISCOIMAGELICENSEMGMTMIB_CilmImageLevelToLicenseMapTable_CilmImageLevelToLicenseMapEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. The type is string. Refers to
    // cisco_image_license_mgmt_mib.CISCOIMAGELICENSEMGMTMIB_CilmBootImageLevelTable_CilmBootImageLevelEntry_CilmModuleName
    CilmModuleName interface{}

    // This attribute is a key. This is a running index used to identify an entry
    // of this table. The type is interface{} with range: 0..4294967295.
    CilmImageLicenseMapIndex interface{}

    // This object indicates the image level at which a module can be run. A
    // module can be run at different image levels. An entry will be created in
    // this table for every module and image level combination. The type is string
    // with length: 0..255.
    CilmImageLicenseImageLevel interface{}

    // This object indicates the list of licenses needed to be installed for the
    // module to run at the image level mentioned by cilmImageLicenseImageLevel
    // object of this entry. The type is string with length: 0..255.
    CilmImageLicenseName interface{}

    // This object indicates the priority of the image level mentioned by
    // cilmImageLicenseImageLevel object of this entry. The image level with the
    // highest priority license will be considered as the default level in the
    // absense of next boot image level configuration. For example if there are
    // three licenses l1, l2 and l3 in the ascending order of priority, then by
    // default l1 will be the level at which the module will be running. If the
    // next boot level is configured then the configuration will override the
    // priority. The highest priority license supports a feature set which is a
    // super set of all other licenses. The type is interface{} with range:
    // 1..255.
    CilmImageLicensePriority interface{}
}

func (cilmImageLevelToLicenseMapEntry *CISCOIMAGELICENSEMGMTMIB_CilmImageLevelToLicenseMapTable_CilmImageLevelToLicenseMapEntry) GetEntityData() *types.CommonEntityData {
    cilmImageLevelToLicenseMapEntry.EntityData.YFilter = cilmImageLevelToLicenseMapEntry.YFilter
    cilmImageLevelToLicenseMapEntry.EntityData.YangName = "cilmImageLevelToLicenseMapEntry"
    cilmImageLevelToLicenseMapEntry.EntityData.BundleName = "cisco_ios_xe"
    cilmImageLevelToLicenseMapEntry.EntityData.ParentYangName = "cilmImageLevelToLicenseMapTable"
    cilmImageLevelToLicenseMapEntry.EntityData.SegmentPath = "cilmImageLevelToLicenseMapEntry" + types.AddKeyToken(cilmImageLevelToLicenseMapEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(cilmImageLevelToLicenseMapEntry.CilmModuleName, "cilmModuleName") + types.AddKeyToken(cilmImageLevelToLicenseMapEntry.CilmImageLicenseMapIndex, "cilmImageLicenseMapIndex")
    cilmImageLevelToLicenseMapEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cilmImageLevelToLicenseMapEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cilmImageLevelToLicenseMapEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cilmImageLevelToLicenseMapEntry.EntityData.Children = types.NewOrderedMap()
    cilmImageLevelToLicenseMapEntry.EntityData.Leafs = types.NewOrderedMap()
    cilmImageLevelToLicenseMapEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", cilmImageLevelToLicenseMapEntry.EntPhysicalIndex})
    cilmImageLevelToLicenseMapEntry.EntityData.Leafs.Append("cilmModuleName", types.YLeaf{"CilmModuleName", cilmImageLevelToLicenseMapEntry.CilmModuleName})
    cilmImageLevelToLicenseMapEntry.EntityData.Leafs.Append("cilmImageLicenseMapIndex", types.YLeaf{"CilmImageLicenseMapIndex", cilmImageLevelToLicenseMapEntry.CilmImageLicenseMapIndex})
    cilmImageLevelToLicenseMapEntry.EntityData.Leafs.Append("cilmImageLicenseImageLevel", types.YLeaf{"CilmImageLicenseImageLevel", cilmImageLevelToLicenseMapEntry.CilmImageLicenseImageLevel})
    cilmImageLevelToLicenseMapEntry.EntityData.Leafs.Append("cilmImageLicenseName", types.YLeaf{"CilmImageLicenseName", cilmImageLevelToLicenseMapEntry.CilmImageLicenseName})
    cilmImageLevelToLicenseMapEntry.EntityData.Leafs.Append("cilmImageLicensePriority", types.YLeaf{"CilmImageLicensePriority", cilmImageLevelToLicenseMapEntry.CilmImageLicensePriority})

    cilmImageLevelToLicenseMapEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "CilmModuleName", "CilmImageLicenseMapIndex"}

    return &(cilmImageLevelToLicenseMapEntry.EntityData)
}

