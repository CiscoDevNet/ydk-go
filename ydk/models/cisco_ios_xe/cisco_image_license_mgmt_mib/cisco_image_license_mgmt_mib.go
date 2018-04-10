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

    
    Ciscoimagelicensemgmtmibobjects CISCOIMAGELICENSEMGMTMIB_Ciscoimagelicensemgmtmibobjects

    
    Cilmnotifcntl CISCOIMAGELICENSEMGMTMIB_Cilmnotifcntl

    // A table that contains the configuration information of current and next
    // boot image level. This table contains entries for each software module
    // running in an image  loaded in the device. The software module is
    // identified by cilmModuleName and the device is identified by 
    // entPhysicalIndex.
    Cilmbootimageleveltable CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable

    // This table contains the mapping between different image levels of each
    // modules in the image and the license required to run the modules at a
    // particular image level. This table can be used to identify the different
    // image levels and the appropriate licenses  required for each.
    Cilmimageleveltolicensemaptable CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable
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

    cISCOIMAGELICENSEMGMTMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOIMAGELICENSEMGMTMIB.EntityData.Children["ciscoImageLicenseMgmtMIBObjects"] = types.YChild{"Ciscoimagelicensemgmtmibobjects", &cISCOIMAGELICENSEMGMTMIB.Ciscoimagelicensemgmtmibobjects}
    cISCOIMAGELICENSEMGMTMIB.EntityData.Children["cilmNotifCntl"] = types.YChild{"Cilmnotifcntl", &cISCOIMAGELICENSEMGMTMIB.Cilmnotifcntl}
    cISCOIMAGELICENSEMGMTMIB.EntityData.Children["cilmBootImageLevelTable"] = types.YChild{"Cilmbootimageleveltable", &cISCOIMAGELICENSEMGMTMIB.Cilmbootimageleveltable}
    cISCOIMAGELICENSEMGMTMIB.EntityData.Children["cilmImageLevelToLicenseMapTable"] = types.YChild{"Cilmimageleveltolicensemaptable", &cISCOIMAGELICENSEMGMTMIB.Cilmimageleveltolicensemaptable}
    cISCOIMAGELICENSEMGMTMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOIMAGELICENSEMGMTMIB.EntityData)
}

// CISCOIMAGELICENSEMGMTMIB_Ciscoimagelicensemgmtmibobjects
type CISCOIMAGELICENSEMGMTMIB_Ciscoimagelicensemgmtmibobjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object when set to TRUE means that the user has accepted the END USER
    // LICENSE AGREEMENT. This object has to be set to TRUE by the user before
    // using the objects in the cilmBootImageLevelTable to configure the license.
    // The type is bool.
    Cilmeulaaccepted interface{}
}

func (ciscoimagelicensemgmtmibobjects *CISCOIMAGELICENSEMGMTMIB_Ciscoimagelicensemgmtmibobjects) GetEntityData() *types.CommonEntityData {
    ciscoimagelicensemgmtmibobjects.EntityData.YFilter = ciscoimagelicensemgmtmibobjects.YFilter
    ciscoimagelicensemgmtmibobjects.EntityData.YangName = "ciscoImageLicenseMgmtMIBObjects"
    ciscoimagelicensemgmtmibobjects.EntityData.BundleName = "cisco_ios_xe"
    ciscoimagelicensemgmtmibobjects.EntityData.ParentYangName = "CISCO-IMAGE-LICENSE-MGMT-MIB"
    ciscoimagelicensemgmtmibobjects.EntityData.SegmentPath = "ciscoImageLicenseMgmtMIBObjects"
    ciscoimagelicensemgmtmibobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    ciscoimagelicensemgmtmibobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    ciscoimagelicensemgmtmibobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    ciscoimagelicensemgmtmibobjects.EntityData.Children = make(map[string]types.YChild)
    ciscoimagelicensemgmtmibobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    ciscoimagelicensemgmtmibobjects.EntityData.Leafs["cilmEULAAccepted"] = types.YLeaf{"Cilmeulaaccepted", ciscoimagelicensemgmtmibobjects.Cilmeulaaccepted}
    return &(ciscoimagelicensemgmtmibobjects.EntityData)
}

// CISCOIMAGELICENSEMGMTMIB_Cilmnotifcntl
type CISCOIMAGELICENSEMGMTMIB_Cilmnotifcntl struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify whether or not a notification should be generated on the detection
    // of change in next boot image level.  If set to TRUE,
    // cilmBootImageLevelChanged notification will be generated. It is the
    // responsibility of the management entity to ensure that the SNMP
    // administrative model is configured in such a way as to allow the 
    // notification to be delivered. The type is bool.
    Cilmimagelevelchangednotif interface{}
}

func (cilmnotifcntl *CISCOIMAGELICENSEMGMTMIB_Cilmnotifcntl) GetEntityData() *types.CommonEntityData {
    cilmnotifcntl.EntityData.YFilter = cilmnotifcntl.YFilter
    cilmnotifcntl.EntityData.YangName = "cilmNotifCntl"
    cilmnotifcntl.EntityData.BundleName = "cisco_ios_xe"
    cilmnotifcntl.EntityData.ParentYangName = "CISCO-IMAGE-LICENSE-MGMT-MIB"
    cilmnotifcntl.EntityData.SegmentPath = "cilmNotifCntl"
    cilmnotifcntl.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cilmnotifcntl.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cilmnotifcntl.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cilmnotifcntl.EntityData.Children = make(map[string]types.YChild)
    cilmnotifcntl.EntityData.Leafs = make(map[string]types.YLeaf)
    cilmnotifcntl.EntityData.Leafs["cilmImageLevelChangedNotif"] = types.YLeaf{"Cilmimagelevelchangednotif", cilmnotifcntl.Cilmimagelevelchangednotif}
    return &(cilmnotifcntl.EntityData)
}

// CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable
// A table that contains the configuration information of
// current and next boot image level. This table contains
// entries for each software module running in an image 
// loaded in the device. The software module is identified by
// cilmModuleName and the device is identified by 
// entPhysicalIndex.
type CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table for each module containing the list of objects that
    // define the configuration of next boot level. The following information is
    // specified by the objects present in the table.  - Current image level. -
    // Configured image level for the next boot. - Actual image level for the next
    // boot. - License store index for the current license. - License index of the
    // current license. - License store index for the next boot license. - License
    // index of the next boot license. The type is slice of
    // CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable_Cilmbootimagelevelentry.
    Cilmbootimagelevelentry []CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable_Cilmbootimagelevelentry
}

func (cilmbootimageleveltable *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable) GetEntityData() *types.CommonEntityData {
    cilmbootimageleveltable.EntityData.YFilter = cilmbootimageleveltable.YFilter
    cilmbootimageleveltable.EntityData.YangName = "cilmBootImageLevelTable"
    cilmbootimageleveltable.EntityData.BundleName = "cisco_ios_xe"
    cilmbootimageleveltable.EntityData.ParentYangName = "CISCO-IMAGE-LICENSE-MGMT-MIB"
    cilmbootimageleveltable.EntityData.SegmentPath = "cilmBootImageLevelTable"
    cilmbootimageleveltable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cilmbootimageleveltable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cilmbootimageleveltable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cilmbootimageleveltable.EntityData.Children = make(map[string]types.YChild)
    cilmbootimageleveltable.EntityData.Children["cilmBootImageLevelEntry"] = types.YChild{"Cilmbootimagelevelentry", nil}
    for i := range cilmbootimageleveltable.Cilmbootimagelevelentry {
        cilmbootimageleveltable.EntityData.Children[types.GetSegmentPath(&cilmbootimageleveltable.Cilmbootimagelevelentry[i])] = types.YChild{"Cilmbootimagelevelentry", &cilmbootimageleveltable.Cilmbootimagelevelentry[i]}
    }
    cilmbootimageleveltable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cilmbootimageleveltable.EntityData)
}

// CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable_Cilmbootimagelevelentry
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
type CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable_Cilmbootimagelevelentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. This object is used as one of the two indices in
    // cilmBootImageLevelTable. This object indicates the module name of the
    // software package. There can be multiple modules in an image performing
    // specific functionality. For example, in a wireless image there can be two
    // modules - a base image module and a wireless module. The type is string.
    Cilmmodulename interface{}

    // This object indicates the current image level that the module is running.
    // The type is string with length: 0..255.
    Cilmcurrentimagelevel interface{}

    // This object indicates the configured image level of the module for the next
    // boot.  Note: The configured next boot image level may not  be the actual
    // next boot image level. The actual next boot image level is denoted by
    // cilmNextBootImageLevel which is determined based on the license
    // availability. The type is string with length: 0..255.
    Cilmconfiguredbootimagelevel interface{}

    // This object indicates the next boot image level. The next boot image level
    // can be different from configured level. The next boot image level is
    // determined based on the availability of required license. The type is
    // string with length: 0..255.
    Cilmnextbootimagelevel interface{}

    // This object indicates the license store index where the currently used
    // license is stored. This object has the same value as
    // clmgmtLicenseStoreIndex object and uniquely identifies an entry in
    // clmgmtLicenseStoreInfoTable in CISCO-LICENSE-MGMT-MIB.  Note: The license
    // store index can be '0' if no license is installed and device is running
    // base image. The type is interface{} with range: 0..4294967295.
    Cilmcurrentlicensestoreindex interface{}

    // This object indicates the license index of the currently used license. This
    // object has the same value as clmgmtLicenseIndex and uniquely identifies an
    // entry in clmgmtLicenseInfoTable in CISCO-LICENSE-MGMT-MIB.  Note: The
    // license index can be '0' if no license is installed and device is running
    // base image. The type is interface{} with range: 0..4294967295.
    Cilmcurrentlicenseindex interface{}

    // This object indicates the license store index where the next boot license
    // is stored. This object has the same value as clmgmtLicenseStoreIndex object
    // and uniquely identifies an entry in clmgmtLicenseStoreInfoTable in
    // CISCO-LICENSE-MGMT-MIB.  Note: The license store index can be '0' if no
    // license is installed for the next boot. The type is interface{} with range:
    // 0..4294967295.
    Cilmnextbootlicensestoreindex interface{}

    // This object indicates the license index of the next boot license. This
    // object has the same value as clmgmtLicenseIndex and uniquely identifies an
    // entry in clmgmtLicenseInfoTable in CISCO-LICENSE-MGMT-MIB.  Note: The
    // license index can be '0' if no license is installed for the next boot. The
    // type is interface{} with range: 0..4294967295.
    Cilmnextbootlicenseindex interface{}
}

func (cilmbootimagelevelentry *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable_Cilmbootimagelevelentry) GetEntityData() *types.CommonEntityData {
    cilmbootimagelevelentry.EntityData.YFilter = cilmbootimagelevelentry.YFilter
    cilmbootimagelevelentry.EntityData.YangName = "cilmBootImageLevelEntry"
    cilmbootimagelevelentry.EntityData.BundleName = "cisco_ios_xe"
    cilmbootimagelevelentry.EntityData.ParentYangName = "cilmBootImageLevelTable"
    cilmbootimagelevelentry.EntityData.SegmentPath = "cilmBootImageLevelEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cilmbootimagelevelentry.Entphysicalindex) + "']" + "[cilmModuleName='" + fmt.Sprintf("%v", cilmbootimagelevelentry.Cilmmodulename) + "']"
    cilmbootimagelevelentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cilmbootimagelevelentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cilmbootimagelevelentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cilmbootimagelevelentry.EntityData.Children = make(map[string]types.YChild)
    cilmbootimagelevelentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cilmbootimagelevelentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cilmbootimagelevelentry.Entphysicalindex}
    cilmbootimagelevelentry.EntityData.Leafs["cilmModuleName"] = types.YLeaf{"Cilmmodulename", cilmbootimagelevelentry.Cilmmodulename}
    cilmbootimagelevelentry.EntityData.Leafs["cilmCurrentImageLevel"] = types.YLeaf{"Cilmcurrentimagelevel", cilmbootimagelevelentry.Cilmcurrentimagelevel}
    cilmbootimagelevelentry.EntityData.Leafs["cilmConfiguredBootImageLevel"] = types.YLeaf{"Cilmconfiguredbootimagelevel", cilmbootimagelevelentry.Cilmconfiguredbootimagelevel}
    cilmbootimagelevelentry.EntityData.Leafs["cilmNextBootImageLevel"] = types.YLeaf{"Cilmnextbootimagelevel", cilmbootimagelevelentry.Cilmnextbootimagelevel}
    cilmbootimagelevelentry.EntityData.Leafs["cilmCurrentLicenseStoreIndex"] = types.YLeaf{"Cilmcurrentlicensestoreindex", cilmbootimagelevelentry.Cilmcurrentlicensestoreindex}
    cilmbootimagelevelentry.EntityData.Leafs["cilmCurrentLicenseIndex"] = types.YLeaf{"Cilmcurrentlicenseindex", cilmbootimagelevelentry.Cilmcurrentlicenseindex}
    cilmbootimagelevelentry.EntityData.Leafs["cilmNextBootLicenseStoreIndex"] = types.YLeaf{"Cilmnextbootlicensestoreindex", cilmbootimagelevelentry.Cilmnextbootlicensestoreindex}
    cilmbootimagelevelentry.EntityData.Leafs["cilmNextBootLicenseIndex"] = types.YLeaf{"Cilmnextbootlicenseindex", cilmbootimagelevelentry.Cilmnextbootlicenseindex}
    return &(cilmbootimagelevelentry.EntityData)
}

// CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable
// This table contains the mapping between different
// image levels of each modules in the image and the
// license required to run the modules at a particular
// image level. This table can be used to identify the
// different image levels and the appropriate licenses 
// required for each.
type CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in the table containing the following information. - The image
    // levels at the which the modules can be run. - The license required to the
    // run a module at a particular image level. - The priority of the license.
    // The type is slice of
    // CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable_Cilmimageleveltolicensemapentry.
    Cilmimageleveltolicensemapentry []CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable_Cilmimageleveltolicensemapentry
}

func (cilmimageleveltolicensemaptable *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable) GetEntityData() *types.CommonEntityData {
    cilmimageleveltolicensemaptable.EntityData.YFilter = cilmimageleveltolicensemaptable.YFilter
    cilmimageleveltolicensemaptable.EntityData.YangName = "cilmImageLevelToLicenseMapTable"
    cilmimageleveltolicensemaptable.EntityData.BundleName = "cisco_ios_xe"
    cilmimageleveltolicensemaptable.EntityData.ParentYangName = "CISCO-IMAGE-LICENSE-MGMT-MIB"
    cilmimageleveltolicensemaptable.EntityData.SegmentPath = "cilmImageLevelToLicenseMapTable"
    cilmimageleveltolicensemaptable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cilmimageleveltolicensemaptable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cilmimageleveltolicensemaptable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cilmimageleveltolicensemaptable.EntityData.Children = make(map[string]types.YChild)
    cilmimageleveltolicensemaptable.EntityData.Children["cilmImageLevelToLicenseMapEntry"] = types.YChild{"Cilmimageleveltolicensemapentry", nil}
    for i := range cilmimageleveltolicensemaptable.Cilmimageleveltolicensemapentry {
        cilmimageleveltolicensemaptable.EntityData.Children[types.GetSegmentPath(&cilmimageleveltolicensemaptable.Cilmimageleveltolicensemapentry[i])] = types.YChild{"Cilmimageleveltolicensemapentry", &cilmimageleveltolicensemaptable.Cilmimageleveltolicensemapentry[i]}
    }
    cilmimageleveltolicensemaptable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cilmimageleveltolicensemaptable.EntityData)
}

// CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable_Cilmimageleveltolicensemapentry
// An entry in the table containing the following
// information.
// - The image levels at the which the modules can be run.
// - The license required to the run a module at a
// particular image level.
// - The priority of the license.
type CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable_Cilmimageleveltolicensemapentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. The type is string. Refers to
    // cisco_image_license_mgmt_mib.CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable_Cilmbootimagelevelentry_Cilmmodulename
    Cilmmodulename interface{}

    // This attribute is a key. This is a running index used to identify an entry
    // of this table. The type is interface{} with range: 0..4294967295.
    Cilmimagelicensemapindex interface{}

    // This object indicates the image level at which a module can be run. A
    // module can be run at different image levels. An entry will be created in
    // this table for every module and image level combination. The type is string
    // with length: 0..255.
    Cilmimagelicenseimagelevel interface{}

    // This object indicates the list of licenses needed to be installed for the
    // module to run at the image level mentioned by cilmImageLicenseImageLevel
    // object of this entry. The type is string with length: 0..255.
    Cilmimagelicensename interface{}

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
    Cilmimagelicensepriority interface{}
}

func (cilmimageleveltolicensemapentry *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable_Cilmimageleveltolicensemapentry) GetEntityData() *types.CommonEntityData {
    cilmimageleveltolicensemapentry.EntityData.YFilter = cilmimageleveltolicensemapentry.YFilter
    cilmimageleveltolicensemapentry.EntityData.YangName = "cilmImageLevelToLicenseMapEntry"
    cilmimageleveltolicensemapentry.EntityData.BundleName = "cisco_ios_xe"
    cilmimageleveltolicensemapentry.EntityData.ParentYangName = "cilmImageLevelToLicenseMapTable"
    cilmimageleveltolicensemapentry.EntityData.SegmentPath = "cilmImageLevelToLicenseMapEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cilmimageleveltolicensemapentry.Entphysicalindex) + "']" + "[cilmModuleName='" + fmt.Sprintf("%v", cilmimageleveltolicensemapentry.Cilmmodulename) + "']" + "[cilmImageLicenseMapIndex='" + fmt.Sprintf("%v", cilmimageleveltolicensemapentry.Cilmimagelicensemapindex) + "']"
    cilmimageleveltolicensemapentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cilmimageleveltolicensemapentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cilmimageleveltolicensemapentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cilmimageleveltolicensemapentry.EntityData.Children = make(map[string]types.YChild)
    cilmimageleveltolicensemapentry.EntityData.Leafs = make(map[string]types.YLeaf)
    cilmimageleveltolicensemapentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", cilmimageleveltolicensemapentry.Entphysicalindex}
    cilmimageleveltolicensemapentry.EntityData.Leafs["cilmModuleName"] = types.YLeaf{"Cilmmodulename", cilmimageleveltolicensemapentry.Cilmmodulename}
    cilmimageleveltolicensemapentry.EntityData.Leafs["cilmImageLicenseMapIndex"] = types.YLeaf{"Cilmimagelicensemapindex", cilmimageleveltolicensemapentry.Cilmimagelicensemapindex}
    cilmimageleveltolicensemapentry.EntityData.Leafs["cilmImageLicenseImageLevel"] = types.YLeaf{"Cilmimagelicenseimagelevel", cilmimageleveltolicensemapentry.Cilmimagelicenseimagelevel}
    cilmimageleveltolicensemapentry.EntityData.Leafs["cilmImageLicenseName"] = types.YLeaf{"Cilmimagelicensename", cilmimageleveltolicensemapentry.Cilmimagelicensename}
    cilmimageleveltolicensemapentry.EntityData.Leafs["cilmImageLicensePriority"] = types.YLeaf{"Cilmimagelicensepriority", cilmimageleveltolicensemapentry.Cilmimagelicensepriority}
    return &(cilmimageleveltolicensemapentry.EntityData)
}

