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
    parent types.Entity
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

func (cISCOIMAGELICENSEMGMTMIB *CISCOIMAGELICENSEMGMTMIB) GetFilter() yfilter.YFilter { return cISCOIMAGELICENSEMGMTMIB.YFilter }

func (cISCOIMAGELICENSEMGMTMIB *CISCOIMAGELICENSEMGMTMIB) SetFilter(yf yfilter.YFilter) { cISCOIMAGELICENSEMGMTMIB.YFilter = yf }

func (cISCOIMAGELICENSEMGMTMIB *CISCOIMAGELICENSEMGMTMIB) GetGoName(yname string) string {
    if yname == "ciscoImageLicenseMgmtMIBObjects" { return "Ciscoimagelicensemgmtmibobjects" }
    if yname == "cilmNotifCntl" { return "Cilmnotifcntl" }
    if yname == "cilmBootImageLevelTable" { return "Cilmbootimageleveltable" }
    if yname == "cilmImageLevelToLicenseMapTable" { return "Cilmimageleveltolicensemaptable" }
    return ""
}

func (cISCOIMAGELICENSEMGMTMIB *CISCOIMAGELICENSEMGMTMIB) GetSegmentPath() string {
    return "CISCO-IMAGE-LICENSE-MGMT-MIB:CISCO-IMAGE-LICENSE-MGMT-MIB"
}

func (cISCOIMAGELICENSEMGMTMIB *CISCOIMAGELICENSEMGMTMIB) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ciscoImageLicenseMgmtMIBObjects" {
        return &cISCOIMAGELICENSEMGMTMIB.Ciscoimagelicensemgmtmibobjects
    }
    if childYangName == "cilmNotifCntl" {
        return &cISCOIMAGELICENSEMGMTMIB.Cilmnotifcntl
    }
    if childYangName == "cilmBootImageLevelTable" {
        return &cISCOIMAGELICENSEMGMTMIB.Cilmbootimageleveltable
    }
    if childYangName == "cilmImageLevelToLicenseMapTable" {
        return &cISCOIMAGELICENSEMGMTMIB.Cilmimageleveltolicensemaptable
    }
    return nil
}

func (cISCOIMAGELICENSEMGMTMIB *CISCOIMAGELICENSEMGMTMIB) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ciscoImageLicenseMgmtMIBObjects"] = &cISCOIMAGELICENSEMGMTMIB.Ciscoimagelicensemgmtmibobjects
    children["cilmNotifCntl"] = &cISCOIMAGELICENSEMGMTMIB.Cilmnotifcntl
    children["cilmBootImageLevelTable"] = &cISCOIMAGELICENSEMGMTMIB.Cilmbootimageleveltable
    children["cilmImageLevelToLicenseMapTable"] = &cISCOIMAGELICENSEMGMTMIB.Cilmimageleveltolicensemaptable
    return children
}

func (cISCOIMAGELICENSEMGMTMIB *CISCOIMAGELICENSEMGMTMIB) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cISCOIMAGELICENSEMGMTMIB *CISCOIMAGELICENSEMGMTMIB) GetBundleName() string { return "cisco_ios_xe" }

func (cISCOIMAGELICENSEMGMTMIB *CISCOIMAGELICENSEMGMTMIB) GetYangName() string { return "CISCO-IMAGE-LICENSE-MGMT-MIB" }

func (cISCOIMAGELICENSEMGMTMIB *CISCOIMAGELICENSEMGMTMIB) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cISCOIMAGELICENSEMGMTMIB *CISCOIMAGELICENSEMGMTMIB) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cISCOIMAGELICENSEMGMTMIB *CISCOIMAGELICENSEMGMTMIB) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cISCOIMAGELICENSEMGMTMIB *CISCOIMAGELICENSEMGMTMIB) SetParent(parent types.Entity) { cISCOIMAGELICENSEMGMTMIB.parent = parent }

func (cISCOIMAGELICENSEMGMTMIB *CISCOIMAGELICENSEMGMTMIB) GetParent() types.Entity { return cISCOIMAGELICENSEMGMTMIB.parent }

func (cISCOIMAGELICENSEMGMTMIB *CISCOIMAGELICENSEMGMTMIB) GetParentYangName() string { return "CISCO-IMAGE-LICENSE-MGMT-MIB" }

// CISCOIMAGELICENSEMGMTMIB_Ciscoimagelicensemgmtmibobjects
type CISCOIMAGELICENSEMGMTMIB_Ciscoimagelicensemgmtmibobjects struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This object when set to TRUE means that the user has accepted the END USER
    // LICENSE AGREEMENT. This object has to be set to TRUE by the user before
    // using the objects in the cilmBootImageLevelTable to configure the license.
    // The type is bool.
    Cilmeulaaccepted interface{}
}

func (ciscoimagelicensemgmtmibobjects *CISCOIMAGELICENSEMGMTMIB_Ciscoimagelicensemgmtmibobjects) GetFilter() yfilter.YFilter { return ciscoimagelicensemgmtmibobjects.YFilter }

func (ciscoimagelicensemgmtmibobjects *CISCOIMAGELICENSEMGMTMIB_Ciscoimagelicensemgmtmibobjects) SetFilter(yf yfilter.YFilter) { ciscoimagelicensemgmtmibobjects.YFilter = yf }

func (ciscoimagelicensemgmtmibobjects *CISCOIMAGELICENSEMGMTMIB_Ciscoimagelicensemgmtmibobjects) GetGoName(yname string) string {
    if yname == "cilmEULAAccepted" { return "Cilmeulaaccepted" }
    return ""
}

func (ciscoimagelicensemgmtmibobjects *CISCOIMAGELICENSEMGMTMIB_Ciscoimagelicensemgmtmibobjects) GetSegmentPath() string {
    return "ciscoImageLicenseMgmtMIBObjects"
}

func (ciscoimagelicensemgmtmibobjects *CISCOIMAGELICENSEMGMTMIB_Ciscoimagelicensemgmtmibobjects) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ciscoimagelicensemgmtmibobjects *CISCOIMAGELICENSEMGMTMIB_Ciscoimagelicensemgmtmibobjects) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ciscoimagelicensemgmtmibobjects *CISCOIMAGELICENSEMGMTMIB_Ciscoimagelicensemgmtmibobjects) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cilmEULAAccepted"] = ciscoimagelicensemgmtmibobjects.Cilmeulaaccepted
    return leafs
}

func (ciscoimagelicensemgmtmibobjects *CISCOIMAGELICENSEMGMTMIB_Ciscoimagelicensemgmtmibobjects) GetBundleName() string { return "cisco_ios_xe" }

func (ciscoimagelicensemgmtmibobjects *CISCOIMAGELICENSEMGMTMIB_Ciscoimagelicensemgmtmibobjects) GetYangName() string { return "ciscoImageLicenseMgmtMIBObjects" }

func (ciscoimagelicensemgmtmibobjects *CISCOIMAGELICENSEMGMTMIB_Ciscoimagelicensemgmtmibobjects) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (ciscoimagelicensemgmtmibobjects *CISCOIMAGELICENSEMGMTMIB_Ciscoimagelicensemgmtmibobjects) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (ciscoimagelicensemgmtmibobjects *CISCOIMAGELICENSEMGMTMIB_Ciscoimagelicensemgmtmibobjects) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (ciscoimagelicensemgmtmibobjects *CISCOIMAGELICENSEMGMTMIB_Ciscoimagelicensemgmtmibobjects) SetParent(parent types.Entity) { ciscoimagelicensemgmtmibobjects.parent = parent }

func (ciscoimagelicensemgmtmibobjects *CISCOIMAGELICENSEMGMTMIB_Ciscoimagelicensemgmtmibobjects) GetParent() types.Entity { return ciscoimagelicensemgmtmibobjects.parent }

func (ciscoimagelicensemgmtmibobjects *CISCOIMAGELICENSEMGMTMIB_Ciscoimagelicensemgmtmibobjects) GetParentYangName() string { return "CISCO-IMAGE-LICENSE-MGMT-MIB" }

// CISCOIMAGELICENSEMGMTMIB_Cilmnotifcntl
type CISCOIMAGELICENSEMGMTMIB_Cilmnotifcntl struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify whether or not a notification should be generated on the detection
    // of change in next boot image level.  If set to TRUE,
    // cilmBootImageLevelChanged notification will be generated. It is the
    // responsibility of the management entity to ensure that the SNMP
    // administrative model is configured in such a way as to allow the 
    // notification to be delivered. The type is bool.
    Cilmimagelevelchangednotif interface{}
}

func (cilmnotifcntl *CISCOIMAGELICENSEMGMTMIB_Cilmnotifcntl) GetFilter() yfilter.YFilter { return cilmnotifcntl.YFilter }

func (cilmnotifcntl *CISCOIMAGELICENSEMGMTMIB_Cilmnotifcntl) SetFilter(yf yfilter.YFilter) { cilmnotifcntl.YFilter = yf }

func (cilmnotifcntl *CISCOIMAGELICENSEMGMTMIB_Cilmnotifcntl) GetGoName(yname string) string {
    if yname == "cilmImageLevelChangedNotif" { return "Cilmimagelevelchangednotif" }
    return ""
}

func (cilmnotifcntl *CISCOIMAGELICENSEMGMTMIB_Cilmnotifcntl) GetSegmentPath() string {
    return "cilmNotifCntl"
}

func (cilmnotifcntl *CISCOIMAGELICENSEMGMTMIB_Cilmnotifcntl) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cilmnotifcntl *CISCOIMAGELICENSEMGMTMIB_Cilmnotifcntl) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cilmnotifcntl *CISCOIMAGELICENSEMGMTMIB_Cilmnotifcntl) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["cilmImageLevelChangedNotif"] = cilmnotifcntl.Cilmimagelevelchangednotif
    return leafs
}

func (cilmnotifcntl *CISCOIMAGELICENSEMGMTMIB_Cilmnotifcntl) GetBundleName() string { return "cisco_ios_xe" }

func (cilmnotifcntl *CISCOIMAGELICENSEMGMTMIB_Cilmnotifcntl) GetYangName() string { return "cilmNotifCntl" }

func (cilmnotifcntl *CISCOIMAGELICENSEMGMTMIB_Cilmnotifcntl) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cilmnotifcntl *CISCOIMAGELICENSEMGMTMIB_Cilmnotifcntl) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cilmnotifcntl *CISCOIMAGELICENSEMGMTMIB_Cilmnotifcntl) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cilmnotifcntl *CISCOIMAGELICENSEMGMTMIB_Cilmnotifcntl) SetParent(parent types.Entity) { cilmnotifcntl.parent = parent }

func (cilmnotifcntl *CISCOIMAGELICENSEMGMTMIB_Cilmnotifcntl) GetParent() types.Entity { return cilmnotifcntl.parent }

func (cilmnotifcntl *CISCOIMAGELICENSEMGMTMIB_Cilmnotifcntl) GetParentYangName() string { return "CISCO-IMAGE-LICENSE-MGMT-MIB" }

// CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable
// A table that contains the configuration information of
// current and next boot image level. This table contains
// entries for each software module running in an image 
// loaded in the device. The software module is identified by
// cilmModuleName and the device is identified by 
// entPhysicalIndex.
type CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable struct {
    parent types.Entity
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

func (cilmbootimageleveltable *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable) GetFilter() yfilter.YFilter { return cilmbootimageleveltable.YFilter }

func (cilmbootimageleveltable *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable) SetFilter(yf yfilter.YFilter) { cilmbootimageleveltable.YFilter = yf }

func (cilmbootimageleveltable *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable) GetGoName(yname string) string {
    if yname == "cilmBootImageLevelEntry" { return "Cilmbootimagelevelentry" }
    return ""
}

func (cilmbootimageleveltable *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable) GetSegmentPath() string {
    return "cilmBootImageLevelTable"
}

func (cilmbootimageleveltable *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cilmBootImageLevelEntry" {
        for _, c := range cilmbootimageleveltable.Cilmbootimagelevelentry {
            if cilmbootimageleveltable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable_Cilmbootimagelevelentry{}
        cilmbootimageleveltable.Cilmbootimagelevelentry = append(cilmbootimageleveltable.Cilmbootimagelevelentry, child)
        return &cilmbootimageleveltable.Cilmbootimagelevelentry[len(cilmbootimageleveltable.Cilmbootimagelevelentry)-1]
    }
    return nil
}

func (cilmbootimageleveltable *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cilmbootimageleveltable.Cilmbootimagelevelentry {
        children[cilmbootimageleveltable.Cilmbootimagelevelentry[i].GetSegmentPath()] = &cilmbootimageleveltable.Cilmbootimagelevelentry[i]
    }
    return children
}

func (cilmbootimageleveltable *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cilmbootimageleveltable *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable) GetBundleName() string { return "cisco_ios_xe" }

func (cilmbootimageleveltable *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable) GetYangName() string { return "cilmBootImageLevelTable" }

func (cilmbootimageleveltable *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cilmbootimageleveltable *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cilmbootimageleveltable *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cilmbootimageleveltable *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable) SetParent(parent types.Entity) { cilmbootimageleveltable.parent = parent }

func (cilmbootimageleveltable *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable) GetParent() types.Entity { return cilmbootimageleveltable.parent }

func (cilmbootimageleveltable *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable) GetParentYangName() string { return "CISCO-IMAGE-LICENSE-MGMT-MIB" }

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
    parent types.Entity
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

func (cilmbootimagelevelentry *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable_Cilmbootimagelevelentry) GetFilter() yfilter.YFilter { return cilmbootimagelevelentry.YFilter }

func (cilmbootimagelevelentry *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable_Cilmbootimagelevelentry) SetFilter(yf yfilter.YFilter) { cilmbootimagelevelentry.YFilter = yf }

func (cilmbootimagelevelentry *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable_Cilmbootimagelevelentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cilmModuleName" { return "Cilmmodulename" }
    if yname == "cilmCurrentImageLevel" { return "Cilmcurrentimagelevel" }
    if yname == "cilmConfiguredBootImageLevel" { return "Cilmconfiguredbootimagelevel" }
    if yname == "cilmNextBootImageLevel" { return "Cilmnextbootimagelevel" }
    if yname == "cilmCurrentLicenseStoreIndex" { return "Cilmcurrentlicensestoreindex" }
    if yname == "cilmCurrentLicenseIndex" { return "Cilmcurrentlicenseindex" }
    if yname == "cilmNextBootLicenseStoreIndex" { return "Cilmnextbootlicensestoreindex" }
    if yname == "cilmNextBootLicenseIndex" { return "Cilmnextbootlicenseindex" }
    return ""
}

func (cilmbootimagelevelentry *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable_Cilmbootimagelevelentry) GetSegmentPath() string {
    return "cilmBootImageLevelEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cilmbootimagelevelentry.Entphysicalindex) + "']" + "[cilmModuleName='" + fmt.Sprintf("%v", cilmbootimagelevelentry.Cilmmodulename) + "']"
}

func (cilmbootimagelevelentry *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable_Cilmbootimagelevelentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cilmbootimagelevelentry *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable_Cilmbootimagelevelentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cilmbootimagelevelentry *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable_Cilmbootimagelevelentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cilmbootimagelevelentry.Entphysicalindex
    leafs["cilmModuleName"] = cilmbootimagelevelentry.Cilmmodulename
    leafs["cilmCurrentImageLevel"] = cilmbootimagelevelentry.Cilmcurrentimagelevel
    leafs["cilmConfiguredBootImageLevel"] = cilmbootimagelevelentry.Cilmconfiguredbootimagelevel
    leafs["cilmNextBootImageLevel"] = cilmbootimagelevelentry.Cilmnextbootimagelevel
    leafs["cilmCurrentLicenseStoreIndex"] = cilmbootimagelevelentry.Cilmcurrentlicensestoreindex
    leafs["cilmCurrentLicenseIndex"] = cilmbootimagelevelentry.Cilmcurrentlicenseindex
    leafs["cilmNextBootLicenseStoreIndex"] = cilmbootimagelevelentry.Cilmnextbootlicensestoreindex
    leafs["cilmNextBootLicenseIndex"] = cilmbootimagelevelentry.Cilmnextbootlicenseindex
    return leafs
}

func (cilmbootimagelevelentry *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable_Cilmbootimagelevelentry) GetBundleName() string { return "cisco_ios_xe" }

func (cilmbootimagelevelentry *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable_Cilmbootimagelevelentry) GetYangName() string { return "cilmBootImageLevelEntry" }

func (cilmbootimagelevelentry *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable_Cilmbootimagelevelentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cilmbootimagelevelentry *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable_Cilmbootimagelevelentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cilmbootimagelevelentry *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable_Cilmbootimagelevelentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cilmbootimagelevelentry *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable_Cilmbootimagelevelentry) SetParent(parent types.Entity) { cilmbootimagelevelentry.parent = parent }

func (cilmbootimagelevelentry *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable_Cilmbootimagelevelentry) GetParent() types.Entity { return cilmbootimagelevelentry.parent }

func (cilmbootimagelevelentry *CISCOIMAGELICENSEMGMTMIB_Cilmbootimageleveltable_Cilmbootimagelevelentry) GetParentYangName() string { return "cilmBootImageLevelTable" }

// CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable
// This table contains the mapping between different
// image levels of each modules in the image and the
// license required to run the modules at a particular
// image level. This table can be used to identify the
// different image levels and the appropriate licenses 
// required for each.
type CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // An entry in the table containing the following information. - The image
    // levels at the which the modules can be run. - The license required to the
    // run a module at a particular image level. - The priority of the license.
    // The type is slice of
    // CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable_Cilmimageleveltolicensemapentry.
    Cilmimageleveltolicensemapentry []CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable_Cilmimageleveltolicensemapentry
}

func (cilmimageleveltolicensemaptable *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable) GetFilter() yfilter.YFilter { return cilmimageleveltolicensemaptable.YFilter }

func (cilmimageleveltolicensemaptable *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable) SetFilter(yf yfilter.YFilter) { cilmimageleveltolicensemaptable.YFilter = yf }

func (cilmimageleveltolicensemaptable *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable) GetGoName(yname string) string {
    if yname == "cilmImageLevelToLicenseMapEntry" { return "Cilmimageleveltolicensemapentry" }
    return ""
}

func (cilmimageleveltolicensemaptable *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable) GetSegmentPath() string {
    return "cilmImageLevelToLicenseMapTable"
}

func (cilmimageleveltolicensemaptable *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cilmImageLevelToLicenseMapEntry" {
        for _, c := range cilmimageleveltolicensemaptable.Cilmimageleveltolicensemapentry {
            if cilmimageleveltolicensemaptable.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable_Cilmimageleveltolicensemapentry{}
        cilmimageleveltolicensemaptable.Cilmimageleveltolicensemapentry = append(cilmimageleveltolicensemaptable.Cilmimageleveltolicensemapentry, child)
        return &cilmimageleveltolicensemaptable.Cilmimageleveltolicensemapentry[len(cilmimageleveltolicensemaptable.Cilmimageleveltolicensemapentry)-1]
    }
    return nil
}

func (cilmimageleveltolicensemaptable *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range cilmimageleveltolicensemaptable.Cilmimageleveltolicensemapentry {
        children[cilmimageleveltolicensemaptable.Cilmimageleveltolicensemapentry[i].GetSegmentPath()] = &cilmimageleveltolicensemaptable.Cilmimageleveltolicensemapentry[i]
    }
    return children
}

func (cilmimageleveltolicensemaptable *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (cilmimageleveltolicensemaptable *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable) GetBundleName() string { return "cisco_ios_xe" }

func (cilmimageleveltolicensemaptable *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable) GetYangName() string { return "cilmImageLevelToLicenseMapTable" }

func (cilmimageleveltolicensemaptable *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cilmimageleveltolicensemaptable *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cilmimageleveltolicensemaptable *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cilmimageleveltolicensemaptable *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable) SetParent(parent types.Entity) { cilmimageleveltolicensemaptable.parent = parent }

func (cilmimageleveltolicensemaptable *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable) GetParent() types.Entity { return cilmimageleveltolicensemaptable.parent }

func (cilmimageleveltolicensemaptable *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable) GetParentYangName() string { return "CISCO-IMAGE-LICENSE-MGMT-MIB" }

// CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable_Cilmimageleveltolicensemapentry
// An entry in the table containing the following
// information.
// - The image levels at the which the modules can be run.
// - The license required to the run a module at a
// particular image level.
// - The priority of the license.
type CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable_Cilmimageleveltolicensemapentry struct {
    parent types.Entity
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

func (cilmimageleveltolicensemapentry *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable_Cilmimageleveltolicensemapentry) GetFilter() yfilter.YFilter { return cilmimageleveltolicensemapentry.YFilter }

func (cilmimageleveltolicensemapentry *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable_Cilmimageleveltolicensemapentry) SetFilter(yf yfilter.YFilter) { cilmimageleveltolicensemapentry.YFilter = yf }

func (cilmimageleveltolicensemapentry *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable_Cilmimageleveltolicensemapentry) GetGoName(yname string) string {
    if yname == "entPhysicalIndex" { return "Entphysicalindex" }
    if yname == "cilmModuleName" { return "Cilmmodulename" }
    if yname == "cilmImageLicenseMapIndex" { return "Cilmimagelicensemapindex" }
    if yname == "cilmImageLicenseImageLevel" { return "Cilmimagelicenseimagelevel" }
    if yname == "cilmImageLicenseName" { return "Cilmimagelicensename" }
    if yname == "cilmImageLicensePriority" { return "Cilmimagelicensepriority" }
    return ""
}

func (cilmimageleveltolicensemapentry *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable_Cilmimageleveltolicensemapentry) GetSegmentPath() string {
    return "cilmImageLevelToLicenseMapEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", cilmimageleveltolicensemapentry.Entphysicalindex) + "']" + "[cilmModuleName='" + fmt.Sprintf("%v", cilmimageleveltolicensemapentry.Cilmmodulename) + "']" + "[cilmImageLicenseMapIndex='" + fmt.Sprintf("%v", cilmimageleveltolicensemapentry.Cilmimagelicensemapindex) + "']"
}

func (cilmimageleveltolicensemapentry *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable_Cilmimageleveltolicensemapentry) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cilmimageleveltolicensemapentry *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable_Cilmimageleveltolicensemapentry) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cilmimageleveltolicensemapentry *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable_Cilmimageleveltolicensemapentry) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["entPhysicalIndex"] = cilmimageleveltolicensemapentry.Entphysicalindex
    leafs["cilmModuleName"] = cilmimageleveltolicensemapentry.Cilmmodulename
    leafs["cilmImageLicenseMapIndex"] = cilmimageleveltolicensemapentry.Cilmimagelicensemapindex
    leafs["cilmImageLicenseImageLevel"] = cilmimageleveltolicensemapentry.Cilmimagelicenseimagelevel
    leafs["cilmImageLicenseName"] = cilmimageleveltolicensemapentry.Cilmimagelicensename
    leafs["cilmImageLicensePriority"] = cilmimageleveltolicensemapentry.Cilmimagelicensepriority
    return leafs
}

func (cilmimageleveltolicensemapentry *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable_Cilmimageleveltolicensemapentry) GetBundleName() string { return "cisco_ios_xe" }

func (cilmimageleveltolicensemapentry *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable_Cilmimageleveltolicensemapentry) GetYangName() string { return "cilmImageLevelToLicenseMapEntry" }

func (cilmimageleveltolicensemapentry *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable_Cilmimageleveltolicensemapentry) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cilmimageleveltolicensemapentry *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable_Cilmimageleveltolicensemapentry) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cilmimageleveltolicensemapentry *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable_Cilmimageleveltolicensemapentry) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cilmimageleveltolicensemapentry *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable_Cilmimageleveltolicensemapentry) SetParent(parent types.Entity) { cilmimageleveltolicensemapentry.parent = parent }

func (cilmimageleveltolicensemapentry *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable_Cilmimageleveltolicensemapentry) GetParent() types.Entity { return cilmimageleveltolicensemapentry.parent }

func (cilmimageleveltolicensemapentry *CISCOIMAGELICENSEMGMTMIB_Cilmimageleveltolicensemaptable_Cilmimageleveltolicensemapentry) GetParentYangName() string { return "cilmImageLevelToLicenseMapTable" }

