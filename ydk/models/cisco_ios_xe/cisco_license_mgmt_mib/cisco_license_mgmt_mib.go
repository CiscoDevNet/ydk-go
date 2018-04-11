// The MIB module for managing licenses on the system.
// The licensing mechanism provides flexibility to
// enforce licensing for various features in the system.
// 
// The following text introduces different concepts and
// terms those are necessary to understand the MIB definition
// and to perform license management.
// 
// UDI
//     Universal Device Identifier that uniquely identifies a
//     device. It comprises of product ID, version ID and serial
//     number of the device.
// 
// NODE LOCKED LICENSE:
//     Node locked licenses are locked to one of the device
//     identifiers in the system. For example, the license
//     can be locked to the UDI value of one of the devices
//     in the system. Locking a license to an UDI guarantees
//     that the license will not be moved to a device with a
//     different UDI.
// 
// NON-NODE LOCKED LICENSE:
//     Non-node locked licenses are not locked to any device
//     identifier. They can be used on other similar devices.
// 
// METERED LICENSE:
//     Metered licenses are valid for limited period of time
//     and they expire after that. This period is the usage
//     period of a license and it is not based on real time
//     clock, but system clock.
// 
// EULA:
//     End User License Agreement.
// 
// EVALUATION LICENSE:
//     Evaluation licenses are non-node locked metered
//     licenses which are valid only for a limited period.
//     They are used only when there are no permanent,
//     extension or grace period licenses for a feature.
//     User will have to accept EULA (End User License
//     Agreement) before using an evaluation license. Even
//     though they are not node locked, their usage is
//     recorded on the device.
// 
// RIGHT TO USE (RTU) LICENSE:
//     Right to use license is a non-node locked metered
//     license which is in evaluation mode for a limited
//     time after which it is converted to Right To Use (RTU) 
//     license and is valid for the lifetime of the product.
//     User will have to accept EULA (End User License Agreement)
//     before using this license. Even though it is not node
//     locked, usage information is recorded on the device.
// 
// EXTENSION LICENSE:
//     Extension licenses are node locked metered licenses.
//     These licenses are issued by Cisco's licensing portal.
//     These licenses need to be installed using management
//     interfaces on the device. User will have to accept an
//     EULA as part of installation of extension license.
// 
// GRACE PERIOD LICENSE:
//     Grace period licenses are node locked metered licenses.
//     These licenses are issued by Cisco's licensing portal
//     as part of the permission ticket to rehost a license.
//     These licenses are installed on the device as part of
//     the rehost operation. User will have to accept an
//     EULA as part of the rehost operation for this type
//     of license. Details on permission ticket, rehost
//     operations are provided further down in this
//     description clause.
// 
// PERMANENT LICENSE:
//     Permanent licenses are node locked licenses that have
//     no usage period associated with them. These licenses
//     are issued by Cisco's licensing portal. These licenses
//     need to be installed using management interfaces on
//     the device. Once these licenses are installed, they
//     will provide needed permission for the feature/image
//     across different versions.
// 
// COUNTED LICENSE:
//     Counted licenses limit the number of similar entities
//     that can use the license. For example, a counted
//     license when used by a feature can limit the number
//     of IP phones that can connect or the number of tunnels
//     that can be created.
// 
// UNCOUNTED LICENSE:
//     Uncounted licenses do not limit the number of similar
//     entities that can use the licenses.
// 
// License can be enforced at the image level or at the feature
// level and this MIB module supports both.
// 
// IMAGE LEVEL LICENSING:
//     A universal image that contains all levels of
//     software packages is loaded on to the device. At boot
//     time, the device determines the highest level of license
//     and brings up the appropriate software features or
//     subsystems.
// 
// FEATURE LEVEL LICENSING:
//     Feature level licensing will support enforcement of
//     license at individual feature. Features have to check
//     for their licenses before enabling themselves. If it
//     does not have a license, the feature should disable
//     itself.
// 
//     There is a one-to-one relationship between
//     a feature and a license. That is, a feature can use
//     only one license at a time and a license can be used
//     by only one feature at a time.
// 
// LICENSE LINE:
//     A License line is an atomic set of ASCII characters
//     arranged in a particular format that holds the license
//     for a single feature within it. A line has all the
//     necessary fields and attributes that make it a valid,
//     non-tamper able and complete license.
// 
// LICENSE FILE:
//     File generated by Cisco licensing portal. It is used
//     to install license on product. It has a user readable
//     part and it contains one or more license lines.
// 
// DEVICE CREDENTIALS:
//     Device credentials file is a document that is generated
//     by a licensed device. This document establishes the
//     identity of the device and proves that the sender/user
//     has/had authorized access to the device.
// 
// REHOST:
//     Rehost operation allows a node locked license that
//     is installed on a device to be installed on other
//     similar device. As part of rehost operation, a device
//     processes the permission ticket, revokes the license(s)
//     on the device and generates a rehost ticket as the
//     proof of revocation. This rehost ticket needs to be
//     presented to the Cisco's licensing portal to get the
//     license transferred on to a new similar device.
// 
// PERMISSION TICKET:
//     Permission ticket is a document generated by Cisco
//     licensing portal that allows a device to rehost its
//     licenses.
// 
// REHOST TICKET:
//     Rehost ticket is document generated by a device after
//     it has applied a permission ticket. The rehost ticket
//     is a proof of revocation.
// 
// LICENSING PORTAL:
//     Generates licenses, permission tickets and verifies
//     device credentials and rehost tickets.
// 
// This MIB module provides MIB objects to install, clear,
// revoke licenses. It also provides objects to regenerate
// last rehost ticket, backup all the licenses installed
// on a device, generate & export EULA for licenses.
// 
// STEPS TO INSTALL A LICENSE:
//     To install a license, the management application
//     1. Retrieves device credentials of the device.
//     2. Communicates with Cisco's licensing portal to get
//        the license file, uses device credentials to identify
//        the device
//     3. Executes the license install action.
// 
// STEPS TO CLEAR A LICENSE:
//     To clear a license, the management application
//     1. Identifies the license to be cleared using license
//        index.
//     2. Executes the license clear action.
// 
// STEPS TO REHOST A LICENSE:
//     To rehost a license, the management application
//     1. Retrieves device credentials of the device.
//     2. Communicates with Cisco's licensing portal to get
//        the permission ticket, uses device credentials to
//        identify the device.
//     3. Executes the processPermissionTicket action. Device
//        revokes the license and generates rehost ticket to be
//        submitted as proof of revocation.
//     4. Retrieves device credentials of the device where the
//        license needs to be transferred to.
//     5. Submits rehost ticket as proof of revocation to
//        Cisco's licensing portal, uses device credentials of
//        the new device to identify the device, gets license
//        file.
//     6. Executes the license install action on the new
//        device.
// 
// STEPS TO REGENERATE LAST REHOST TICKET:
//     To regenerate last rehost ticket, the management
// application
//     1. Retrieves device credentials of the device.
//     2. Uses already obtained permission ticket or
//        communicates with Cisco's licensing portal to get
//        the permission ticket, uses device credentials to
//        identify the device.
//     3. Executes the regenerateLastRehostTicket action.
//        Device generates rehost ticket to be submitted as
//        proof of revocation.
// 
// STEPS TO BACKUP ALL LICENSES:
//     To backup all licenses installed in the device, the
//     management application
//     1. Specifies the backup file path.
//     2. Executes the license backup action.
// 
// STEPS TO GENERATE & EXPORT EULA:
//     To install certain kind of licenses, the management
//     application must accept EULA first. The management
//     application can generate and export EULA for one or
//     more licenses with out installing licenses as follows.
//     1. Specifies the license file path that has licenses to be
//        installed
//     2. Specifies the EULA file path where EULA need to be
//        exported to
//     3. Executes the generate EULA action.
// 
// To support the various license actions, this MIB module
// also defines MIB objects to know if a device supports
// licensing, retrieve device credentials, retrieve
// information on license stores in the device.
// 
// It also defines MIB objects to expose management
// information associated with the licenses installed on the
// device, with the licensable features in the software image.
// 
// This MIB module also defines various notifications that
// will be triggered to asynchronously notify the management
// application of any critical events.
// 
// This MIB module is defined generically so it can be
// implemented on stand alone devices or stack of devices.
// In stack of devices, one device in the stack acts as
// master agent and rest are slave agents. Each device in the
// stack has its own UDI. The master agent receives requests
// on behalf of all the devices in the stack including itself
// and delegates it to other devices as needed. It also
// collects responses from other devices and sends them to
// the management application. Examples of such devices include
// stackable switches, devices with route processor and line
// cards. On the other hand, stand alone device is a single
// device and has only one UDI associated with it.
// 
// entPhysicalIndex imported from ENTITY-MIB is used to
// identify the device uniquely. It is specified as the index
// or one of the index for tables in this MIB as needed.
package cisco_license_mgmt_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_license_mgmt_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-LICENSE-MGMT-MIB CISCO-LICENSE-MGMT-MIB}", reflect.TypeOf(CISCOLICENSEMGMTMIB{}))
    ydk.RegisterEntity("CISCO-LICENSE-MGMT-MIB:CISCO-LICENSE-MGMT-MIB", reflect.TypeOf(CISCOLICENSEMGMTMIB{}))
}

// ClmgmtLicenseTransferProtocol represents which protocol is supported and use.
type ClmgmtLicenseTransferProtocol string

const (
    ClmgmtLicenseTransferProtocol_none ClmgmtLicenseTransferProtocol = "none"

    ClmgmtLicenseTransferProtocol_local ClmgmtLicenseTransferProtocol = "local"

    ClmgmtLicenseTransferProtocol_tftp ClmgmtLicenseTransferProtocol = "tftp"

    ClmgmtLicenseTransferProtocol_ftp ClmgmtLicenseTransferProtocol = "ftp"

    ClmgmtLicenseTransferProtocol_rcp ClmgmtLicenseTransferProtocol = "rcp"

    ClmgmtLicenseTransferProtocol_http ClmgmtLicenseTransferProtocol = "http"

    ClmgmtLicenseTransferProtocol_scp ClmgmtLicenseTransferProtocol = "scp"

    ClmgmtLicenseTransferProtocol_sftp ClmgmtLicenseTransferProtocol = "sftp"
)

// ClmgmtLicenseActionState represents failed(6)              - action has failed.
type ClmgmtLicenseActionState string

const (
    ClmgmtLicenseActionState_none ClmgmtLicenseActionState = "none"

    ClmgmtLicenseActionState_pending ClmgmtLicenseActionState = "pending"

    ClmgmtLicenseActionState_inProgress ClmgmtLicenseActionState = "inProgress"

    ClmgmtLicenseActionState_successful ClmgmtLicenseActionState = "successful"

    ClmgmtLicenseActionState_partiallySuccessful ClmgmtLicenseActionState = "partiallySuccessful"

    ClmgmtLicenseActionState_failed ClmgmtLicenseActionState = "failed"
)

// ClmgmtLicenseActionFailCause represents invalidLicenseEULAFile(22) -  EULA file path is not accessible.
type ClmgmtLicenseActionFailCause string

const (
    ClmgmtLicenseActionFailCause_none ClmgmtLicenseActionFailCause = "none"

    ClmgmtLicenseActionFailCause_generalFailure ClmgmtLicenseActionFailCause = "generalFailure"

    ClmgmtLicenseActionFailCause_transferProtocolNotSupported ClmgmtLicenseActionFailCause = "transferProtocolNotSupported"

    ClmgmtLicenseActionFailCause_fileServerNotReachable ClmgmtLicenseActionFailCause = "fileServerNotReachable"

    ClmgmtLicenseActionFailCause_unrecognizedEntPhysicalIndex ClmgmtLicenseActionFailCause = "unrecognizedEntPhysicalIndex"

    ClmgmtLicenseActionFailCause_invalidLicenseFilePath ClmgmtLicenseActionFailCause = "invalidLicenseFilePath"

    ClmgmtLicenseActionFailCause_invalidLicenseFile ClmgmtLicenseActionFailCause = "invalidLicenseFile"

    ClmgmtLicenseActionFailCause_invalidLicenseLine ClmgmtLicenseActionFailCause = "invalidLicenseLine"

    ClmgmtLicenseActionFailCause_licenseAlreadyExists ClmgmtLicenseActionFailCause = "licenseAlreadyExists"

    ClmgmtLicenseActionFailCause_licenseNotValidForDevice ClmgmtLicenseActionFailCause = "licenseNotValidForDevice"

    ClmgmtLicenseActionFailCause_invalidLicenseCount ClmgmtLicenseActionFailCause = "invalidLicenseCount"

    ClmgmtLicenseActionFailCause_invalidLicensePeriod ClmgmtLicenseActionFailCause = "invalidLicensePeriod"

    ClmgmtLicenseActionFailCause_licenseInUse ClmgmtLicenseActionFailCause = "licenseInUse"

    ClmgmtLicenseActionFailCause_invalidLicenseStore ClmgmtLicenseActionFailCause = "invalidLicenseStore"

    ClmgmtLicenseActionFailCause_licenseStorageFull ClmgmtLicenseActionFailCause = "licenseStorageFull"

    ClmgmtLicenseActionFailCause_invalidPermissionTicketFile ClmgmtLicenseActionFailCause = "invalidPermissionTicketFile"

    ClmgmtLicenseActionFailCause_invalidPermissionTicket ClmgmtLicenseActionFailCause = "invalidPermissionTicket"

    ClmgmtLicenseActionFailCause_invalidRehostTicketFile ClmgmtLicenseActionFailCause = "invalidRehostTicketFile"

    ClmgmtLicenseActionFailCause_invalidRehostTicket ClmgmtLicenseActionFailCause = "invalidRehostTicket"

    ClmgmtLicenseActionFailCause_invalidLicenseBackupFile ClmgmtLicenseActionFailCause = "invalidLicenseBackupFile"

    ClmgmtLicenseActionFailCause_licenseClearInProgress ClmgmtLicenseActionFailCause = "licenseClearInProgress"

    ClmgmtLicenseActionFailCause_invalidLicenseEULAFile ClmgmtLicenseActionFailCause = "invalidLicenseEULAFile"
)

// CISCOLICENSEMGMTMIB
type CISCOLICENSEMGMTMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Clmgmtlicenseconfiguration CISCOLICENSEMGMTMIB_Clmgmtlicenseconfiguration

    
    Clmgmtlicensedeviceinformation CISCOLICENSEMGMTMIB_Clmgmtlicensedeviceinformation

    
    Clmgmtlicensenotifobjects CISCOLICENSEMGMTMIB_Clmgmtlicensenotifobjects

    // A table for invoking license management actions. Management application
    // must create a row in this table to trigger any of the license management
    // actions. The following are different actions that can be executed using
    // this table.     1. install     2. clear     3. processPermissionTicket    
    // 4. regenerateLastRehostTicket     5. backup     6. generateEULA  Refer to
    // the description of clmgmtLicenseAction for more information on what these
    // actions do on the device. Once the request completes, the management
    // application should retrieve the values of the objects of interest, and then
    // delete the entry.  In order to prevent old entries from clogging the table,
    // entries will be aged out, but an entry will never be deleted within 5
    // minutes of completion.
    Clmgmtlicenseactiontable CISCOLICENSEMGMTMIB_Clmgmtlicenseactiontable

    // This table contains results of license action if the license action
    // involves multiple licenses. Entries in this table are not created for
    // actions where there is only license that is subject of the action. For
    // example, if there are 3 licenses in a license file when executing license
    // install action, 3 entries will be created in this table, one for each
    // license.
    Clmgmtlicenseactionresulttable CISCOLICENSEMGMTMIB_Clmgmtlicenseactionresulttable

    // This table contains information about all the license stores allocated on
    // the device.
    Clmgmtlicensestoreinfotable CISCOLICENSEMGMTMIB_Clmgmtlicensestoreinfotable

    // This table contains objects that provide licensing related information at
    // the device level. Entries will exist only for entities that support
    // licensing. For example, if it is a stand alone device and supports
    // licensing, then there will be only one entry in this table. If it is
    // stackable switch then there will be multiple entries with one entry for
    // each device in the stack.
    Clmgmtlicensedeviceinfotable CISCOLICENSEMGMTMIB_Clmgmtlicensedeviceinfotable

    // This table contains information about all the licenses installed on the
    // device.
    Clmgmtlicenseinfotable CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable

    // This table contains list of licensable features in the image. All the
    // licensable features will have an entry each in this table irrespective of
    // whether they are using any licenses currently. Entries in this table are
    // created by the agent one for each licensable feature in the image. These
    // entries remain in the table permanently and can not be deleted. Management
    // application can not create or delete entries from this table.
    Clmgmtlicensablefeaturetable CISCOLICENSEMGMTMIB_Clmgmtlicensablefeaturetable

    // A table for triggering device credentials export action. Management
    // application must create this entry to trigger the export of device
    // credentials from the device to a file.  Once the request completes, the
    // management application should retrieve the values of the objects of
    // interest, and then delete the entry.  In order to prevent old entries from
    // clogging the table, entries will be aged out, but an entry will never be
    // deleted within 5 minutes of completion.
    Clmgmtdevcredexportactiontable CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable
}

func (cISCOLICENSEMGMTMIB *CISCOLICENSEMGMTMIB) GetEntityData() *types.CommonEntityData {
    cISCOLICENSEMGMTMIB.EntityData.YFilter = cISCOLICENSEMGMTMIB.YFilter
    cISCOLICENSEMGMTMIB.EntityData.YangName = "CISCO-LICENSE-MGMT-MIB"
    cISCOLICENSEMGMTMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOLICENSEMGMTMIB.EntityData.ParentYangName = "CISCO-LICENSE-MGMT-MIB"
    cISCOLICENSEMGMTMIB.EntityData.SegmentPath = "CISCO-LICENSE-MGMT-MIB:CISCO-LICENSE-MGMT-MIB"
    cISCOLICENSEMGMTMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOLICENSEMGMTMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOLICENSEMGMTMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOLICENSEMGMTMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOLICENSEMGMTMIB.EntityData.Children["clmgmtLicenseConfiguration"] = types.YChild{"Clmgmtlicenseconfiguration", &cISCOLICENSEMGMTMIB.Clmgmtlicenseconfiguration}
    cISCOLICENSEMGMTMIB.EntityData.Children["clmgmtLicenseDeviceInformation"] = types.YChild{"Clmgmtlicensedeviceinformation", &cISCOLICENSEMGMTMIB.Clmgmtlicensedeviceinformation}
    cISCOLICENSEMGMTMIB.EntityData.Children["clmgmtLicenseNotifObjects"] = types.YChild{"Clmgmtlicensenotifobjects", &cISCOLICENSEMGMTMIB.Clmgmtlicensenotifobjects}
    cISCOLICENSEMGMTMIB.EntityData.Children["clmgmtLicenseActionTable"] = types.YChild{"Clmgmtlicenseactiontable", &cISCOLICENSEMGMTMIB.Clmgmtlicenseactiontable}
    cISCOLICENSEMGMTMIB.EntityData.Children["clmgmtLicenseActionResultTable"] = types.YChild{"Clmgmtlicenseactionresulttable", &cISCOLICENSEMGMTMIB.Clmgmtlicenseactionresulttable}
    cISCOLICENSEMGMTMIB.EntityData.Children["clmgmtLicenseStoreInfoTable"] = types.YChild{"Clmgmtlicensestoreinfotable", &cISCOLICENSEMGMTMIB.Clmgmtlicensestoreinfotable}
    cISCOLICENSEMGMTMIB.EntityData.Children["clmgmtLicenseDeviceInfoTable"] = types.YChild{"Clmgmtlicensedeviceinfotable", &cISCOLICENSEMGMTMIB.Clmgmtlicensedeviceinfotable}
    cISCOLICENSEMGMTMIB.EntityData.Children["clmgmtLicenseInfoTable"] = types.YChild{"Clmgmtlicenseinfotable", &cISCOLICENSEMGMTMIB.Clmgmtlicenseinfotable}
    cISCOLICENSEMGMTMIB.EntityData.Children["clmgmtLicensableFeatureTable"] = types.YChild{"Clmgmtlicensablefeaturetable", &cISCOLICENSEMGMTMIB.Clmgmtlicensablefeaturetable}
    cISCOLICENSEMGMTMIB.EntityData.Children["clmgmtDevCredExportActionTable"] = types.YChild{"Clmgmtdevcredexportactiontable", &cISCOLICENSEMGMTMIB.Clmgmtdevcredexportactiontable}
    cISCOLICENSEMGMTMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOLICENSEMGMTMIB.EntityData)
}

// CISCOLICENSEMGMTMIB_Clmgmtlicenseconfiguration
type CISCOLICENSEMGMTMIB_Clmgmtlicenseconfiguration struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object contains appropriate value for clmgmtLicenseActionIndex that
    // can be used to create an entry in clmgmtLicenseActionTable. The management
    // application should read this object first and then use this as the value
    // for clmgmtLicenseActionIndex to avoid collisions when creating entries in
    // clmgmtLicenseActionTable. Following this approach does not guarantee
    // collision free row creation, but will reduce the probability. The collision
    // will happen if two management applications read this object at the same
    // time and attempt to create an entry with this value at the same time. In
    // this case, the management application whose request is processed after the
    // first request will get an error and the process of reading this object and
    // entry creation needs to be repeated. The type is interface{} with range:
    // 1..4294967295.
    Clmgmtnextfreelicenseactionindex interface{}
}

func (clmgmtlicenseconfiguration *CISCOLICENSEMGMTMIB_Clmgmtlicenseconfiguration) GetEntityData() *types.CommonEntityData {
    clmgmtlicenseconfiguration.EntityData.YFilter = clmgmtlicenseconfiguration.YFilter
    clmgmtlicenseconfiguration.EntityData.YangName = "clmgmtLicenseConfiguration"
    clmgmtlicenseconfiguration.EntityData.BundleName = "cisco_ios_xe"
    clmgmtlicenseconfiguration.EntityData.ParentYangName = "CISCO-LICENSE-MGMT-MIB"
    clmgmtlicenseconfiguration.EntityData.SegmentPath = "clmgmtLicenseConfiguration"
    clmgmtlicenseconfiguration.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtlicenseconfiguration.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtlicenseconfiguration.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtlicenseconfiguration.EntityData.Children = make(map[string]types.YChild)
    clmgmtlicenseconfiguration.EntityData.Leafs = make(map[string]types.YLeaf)
    clmgmtlicenseconfiguration.EntityData.Leafs["clmgmtNextFreeLicenseActionIndex"] = types.YLeaf{"Clmgmtnextfreelicenseactionindex", clmgmtlicenseconfiguration.Clmgmtnextfreelicenseactionindex}
    return &(clmgmtlicenseconfiguration.EntityData)
}

// CISCOLICENSEMGMTMIB_Clmgmtlicensedeviceinformation
type CISCOLICENSEMGMTMIB_Clmgmtlicensedeviceinformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object contains appropriate value for clmgmtDevCredExportActionIndex
    // that can be used to create an entry in clmgmtDevCredExportActionTable. The
    // management application should read this object first and then use this as
    // the value for clmgmtDevCredExportActionIndex to avoid collisions when
    // creating entries in clmgmtDevCredExportActionTable. Following this approach
    // does not guarantee collision free row creation, but will reduce the
    // probability. The collision will happen if two management applications read
    // this object at the same time and attempt to create an entry with this value
    // at the same time. In this case, the management application whose request is
    // processed after the first request will get an error and the process of
    // reading this object and entry creation needs to be repeated. The type is
    // interface{} with range: 0..4294967295.
    Clmgmtnextfreedevcredexportactionindex interface{}
}

func (clmgmtlicensedeviceinformation *CISCOLICENSEMGMTMIB_Clmgmtlicensedeviceinformation) GetEntityData() *types.CommonEntityData {
    clmgmtlicensedeviceinformation.EntityData.YFilter = clmgmtlicensedeviceinformation.YFilter
    clmgmtlicensedeviceinformation.EntityData.YangName = "clmgmtLicenseDeviceInformation"
    clmgmtlicensedeviceinformation.EntityData.BundleName = "cisco_ios_xe"
    clmgmtlicensedeviceinformation.EntityData.ParentYangName = "CISCO-LICENSE-MGMT-MIB"
    clmgmtlicensedeviceinformation.EntityData.SegmentPath = "clmgmtLicenseDeviceInformation"
    clmgmtlicensedeviceinformation.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtlicensedeviceinformation.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtlicensedeviceinformation.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtlicensedeviceinformation.EntityData.Children = make(map[string]types.YChild)
    clmgmtlicensedeviceinformation.EntityData.Leafs = make(map[string]types.YLeaf)
    clmgmtlicensedeviceinformation.EntityData.Leafs["clmgmtNextFreeDevCredExportActionIndex"] = types.YLeaf{"Clmgmtnextfreedevcredexportactionindex", clmgmtlicensedeviceinformation.Clmgmtnextfreedevcredexportactionindex}
    return &(clmgmtlicensedeviceinformation.EntityData)
}

// CISCOLICENSEMGMTMIB_Clmgmtlicensenotifobjects
type CISCOLICENSEMGMTMIB_Clmgmtlicensenotifobjects struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This object indicates whether the device should generate the notifications
    // related to usage of licenses. This object enables/disables sending
    // following notifications:     clmgmtLicenseExpired     
    // clmgmtLicenseExpiryWarning     clmgmtLicenseUsageCountExceeded    
    // clmgmtLicenseUsageCountAboutToExceed    
    // clmgmtLicenseSubscriptionExpiryWarning    
    // clmgmtLicenseSubscriptionExtExpiryWarning    
    // clmgmtLicenseSubscriptionExpired     clmgmtLicenseEvalRTUTransitionWarning 
    // clmgmtLicenseEvalRTUTransition. The type is bool.
    Clmgmtlicenseusagenotifenable interface{}

    // This object indicates whether the device should generate notifications
    // related to license deployment. This object enables/disables sending
    // following notifications:     clmgmtLicenseInstalled    
    // clmgmtLicenseCleared     clmgmtLicenseRevoked    
    // clmgmtLicenseEULAAccepted. The type is bool.
    Clmgmtlicensedeploymentnotifenable interface{}

    // This object indicates whether the device should generate notifications
    // related to error conditions in enforcing licensing. This object
    // enables/disables sending following notifications:    
    // clmgmtLicenseNotEnforced. The type is Clmgmtlicenseerrornotifenable.
    Clmgmtlicenseerrornotifenable interface{}
}

func (clmgmtlicensenotifobjects *CISCOLICENSEMGMTMIB_Clmgmtlicensenotifobjects) GetEntityData() *types.CommonEntityData {
    clmgmtlicensenotifobjects.EntityData.YFilter = clmgmtlicensenotifobjects.YFilter
    clmgmtlicensenotifobjects.EntityData.YangName = "clmgmtLicenseNotifObjects"
    clmgmtlicensenotifobjects.EntityData.BundleName = "cisco_ios_xe"
    clmgmtlicensenotifobjects.EntityData.ParentYangName = "CISCO-LICENSE-MGMT-MIB"
    clmgmtlicensenotifobjects.EntityData.SegmentPath = "clmgmtLicenseNotifObjects"
    clmgmtlicensenotifobjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtlicensenotifobjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtlicensenotifobjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtlicensenotifobjects.EntityData.Children = make(map[string]types.YChild)
    clmgmtlicensenotifobjects.EntityData.Leafs = make(map[string]types.YLeaf)
    clmgmtlicensenotifobjects.EntityData.Leafs["clmgmtLicenseUsageNotifEnable"] = types.YLeaf{"Clmgmtlicenseusagenotifenable", clmgmtlicensenotifobjects.Clmgmtlicenseusagenotifenable}
    clmgmtlicensenotifobjects.EntityData.Leafs["clmgmtLicenseDeploymentNotifEnable"] = types.YLeaf{"Clmgmtlicensedeploymentnotifenable", clmgmtlicensenotifobjects.Clmgmtlicensedeploymentnotifenable}
    clmgmtlicensenotifobjects.EntityData.Leafs["clmgmtLicenseErrorNotifEnable"] = types.YLeaf{"Clmgmtlicenseerrornotifenable", clmgmtlicensenotifobjects.Clmgmtlicenseerrornotifenable}
    return &(clmgmtlicensenotifobjects.EntityData)
}

// CISCOLICENSEMGMTMIB_Clmgmtlicensenotifobjects_Clmgmtlicenseerrornotifenable represents     clmgmtLicenseNotEnforced
type CISCOLICENSEMGMTMIB_Clmgmtlicensenotifobjects_Clmgmtlicenseerrornotifenable string

const (
    CISCOLICENSEMGMTMIB_Clmgmtlicensenotifobjects_Clmgmtlicenseerrornotifenable_other CISCOLICENSEMGMTMIB_Clmgmtlicensenotifobjects_Clmgmtlicenseerrornotifenable = "other"

    CISCOLICENSEMGMTMIB_Clmgmtlicensenotifobjects_Clmgmtlicenseerrornotifenable_true CISCOLICENSEMGMTMIB_Clmgmtlicensenotifobjects_Clmgmtlicenseerrornotifenable = "true"

    CISCOLICENSEMGMTMIB_Clmgmtlicensenotifobjects_Clmgmtlicenseerrornotifenable_false CISCOLICENSEMGMTMIB_Clmgmtlicensenotifobjects_Clmgmtlicenseerrornotifenable = "false"
)

// CISCOLICENSEMGMTMIB_Clmgmtlicenseactiontable
// A table for invoking license management actions. Management
// application must create a row in this table to trigger any of
// the license management actions. The following are different
// actions that can be executed using this table.
//     1. install
//     2. clear
//     3. processPermissionTicket
//     4. regenerateLastRehostTicket
//     5. backup
//     6. generateEULA
// 
// Refer to the description of clmgmtLicenseAction for more
// information on what these actions do on the device.
// Once the request completes, the management application should
// retrieve the values of the objects of interest, and then
// delete the entry.  In order to prevent old entries from
// clogging the table, entries will be aged out, but an entry
// will never be deleted within 5 minutes of completion.
type CISCOLICENSEMGMTMIB_Clmgmtlicenseactiontable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry for each action that is being executed or was executed recently.
    // The management application executes an action by creating this entry. This
    // can be done in the following 2 methods  1. CREATE-AND-GO method    
    // Management application sets clmgmtLicenseActionRowStatus to    
    // createAndGo(4) and all other required objects to valid     values in a
    // single SNMP SET request. If all the values     are valid, the device
    // creates the entry and executes the     action. If the SET request fails,
    // the entry will not be     created. 2. CREATE-AND-WAIT method     Management
    // application sets clmgmtLicenseActionRowStatus to     createAndWait(5) to
    // create an entry. Management application     can set all other required
    // objects to valid     values in more than one SNMP SET request. If SET
    // request     for any of the objects fails, management application can set   
    // just only that object. Once all the required objects     are set to valid
    // values, management application triggers action     execution by setting
    // clmgmtLicenseActionRowStatus to     active(1).  To stop the action from
    // being executed, the management application can delete the entry by setting
    // clmgmtLicenseActionRowStatus to destroy(6) when clmgmtLicenseActionState is
    // pending(2).  The status of action execution can be known by querying
    // clmgmtLicenseActionState. If the action is still in pending(2) or in
    // inProgress(3) state, the management application need to check back again
    // after few seconds. Once the action completes and status of the action is
    // failed(6), the reason for failure can be retrieved from
    // clmgmtLicenseActionFailCause. If the status of the action is
    // partiallySuccessful(5), results of individual licenses can be queried from
    // clmgmtLicenseActionResultTable.  Not all objects in the entry are needed to
    // execute every action. Below is the list of actions and the required objects
    // that are needed to be set for executing that action.  1. Installing a
    // license    The following MIB objects need to be set for installing a   
    // license      a. clmgmtLicenseActionTransferProtocol      b.
    // clmgmtLicenseServerAddressType      c. clmgmtLicenseServerAddress      d.
    // clmgmtLicenseServerUsername      e. clmgmtLicenseServerPassword      f.
    // clmgmtLicenseFile      g. clmgmtLicenseStore      h.
    // clmgmtLicenseStopOnFailure      i. clmgmtLicenseAcceptEULA      j.
    // clmgmtLicenseAction     clmgmtLicenseActionEntPhysicalIndex need not be set
    // explicitly for license installs. License itself identifes    the device
    // where the license needs to be installed.     clmgmtLicenseStore need to be
    // set to store the licenses    in a non-default license store. But, if a
    // license file    has more than one license and licenses need to be   
    // installed on multiple devices (for example to multiple    members with in a
    // stack), then value of clmgmtLicenseStore    is ignored and the licenses
    // will be installed in default    license stores of the respective devices. 
    // 2. Clearing a license    The following MIB objects need to be set for
    // clearing a    license      a. clmgmtLicenseActionEntPhysicalIndex      b.
    // clmgmtLicenseActionLicenseIndex      c. clmgmtLicenseStore      d.
    // clmgmtLicenseAction  3. Revoking a license    The following MIB objects
    // need to be set for revoking a    license      a.
    // clmgmtLicenseActionTransferProtocol      b. clmgmtLicenseServerAddressType 
    // c. clmgmtLicenseServerAddress      d. clmgmtLicenseServerUsername      e.
    // clmgmtLicenseServerPassword      f. clmgmtLicensePermissionTicketFile     
    // g. clmgmtLicenseRehostTicketFile      h. clmgmtLicenseStopOnFailure      i.
    // clmgmtLicenseAction  4. Regenerate last rehost ticket    The following MIB
    // objects need to be set for regenerating    last rehost ticket      a.
    // clmgmtLicenseActionTransferProtocol      b. clmgmtLicenseServerAddressType 
    // c. clmgmtLicenseServerAddress      d. clmgmtLicenseServerUsername      e.
    // clmgmtLicenseServerPassword      f. clmgmtLicensePermissionTicketFile     
    // g. clmgmtLicenseRehostTicketFile      h. clmgmtLicenseStopOnFailure      i.
    // clmgmtLicenseAction   5. Save all licenses to a backup storage    The
    // following MIB objects need to be set for storing all    licenses to a
    // backup store      a. clmgmtLicenseActionEntPhysicalIndex      b.
    // clmgmtLicenseActionTransferProtocol      c. clmgmtLicenseServerAddressType 
    // d. clmgmtLicenseServerAddress      e. clmgmtLicenseServerUsername      f.
    // clmgmtLicenseServerPassword      g. clmgmtLicenseBackupFile      h.
    // clmgmtLicenseAction  6. Generate and export EULA if the licenses need EULA
    // to be    accepted for installing.    The following MIB objects need to be
    // set exporting required    EULAs      a. clmgmtLicenseActionTransferProtocol
    // b. clmgmtLicenseServerAddressType      c. clmgmtLicenseServerAddress     
    // d. clmgmtLicenseServerUsername      e. clmgmtLicenseServerPassword      f.
    // clmgmtLicenseFile      g. clmgmtLicenseEULAFile      h. clmgmtLicenseAction
    // For any action, if clmgmtLicenseActionTransferProtocol is set to local(2),
    // the following objects need not be set.      a.
    // clmgmtLicenseServerAddressType      b. clmgmtLicenseServerAddress      c.
    // clmgmtLicenseServerUsername      d. clmgmtLicenseServerPassword  Entry can
    // be deleted except when clmgmtLicenseAction is set to pending(2). All
    // entries are volatile and are cleared on agent reset. The type is slice of
    // CISCOLICENSEMGMTMIB_Clmgmtlicenseactiontable_Clmgmtlicenseactionentry.
    Clmgmtlicenseactionentry []CISCOLICENSEMGMTMIB_Clmgmtlicenseactiontable_Clmgmtlicenseactionentry
}

func (clmgmtlicenseactiontable *CISCOLICENSEMGMTMIB_Clmgmtlicenseactiontable) GetEntityData() *types.CommonEntityData {
    clmgmtlicenseactiontable.EntityData.YFilter = clmgmtlicenseactiontable.YFilter
    clmgmtlicenseactiontable.EntityData.YangName = "clmgmtLicenseActionTable"
    clmgmtlicenseactiontable.EntityData.BundleName = "cisco_ios_xe"
    clmgmtlicenseactiontable.EntityData.ParentYangName = "CISCO-LICENSE-MGMT-MIB"
    clmgmtlicenseactiontable.EntityData.SegmentPath = "clmgmtLicenseActionTable"
    clmgmtlicenseactiontable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtlicenseactiontable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtlicenseactiontable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtlicenseactiontable.EntityData.Children = make(map[string]types.YChild)
    clmgmtlicenseactiontable.EntityData.Children["clmgmtLicenseActionEntry"] = types.YChild{"Clmgmtlicenseactionentry", nil}
    for i := range clmgmtlicenseactiontable.Clmgmtlicenseactionentry {
        clmgmtlicenseactiontable.EntityData.Children[types.GetSegmentPath(&clmgmtlicenseactiontable.Clmgmtlicenseactionentry[i])] = types.YChild{"Clmgmtlicenseactionentry", &clmgmtlicenseactiontable.Clmgmtlicenseactionentry[i]}
    }
    clmgmtlicenseactiontable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clmgmtlicenseactiontable.EntityData)
}

// CISCOLICENSEMGMTMIB_Clmgmtlicenseactiontable_Clmgmtlicenseactionentry
// An entry for each action that is being executed or was
// executed recently. The management application executes an
// action
// by creating this entry. This can be done in the following
// 2 methods
// 
// 1. CREATE-AND-GO method
//     Management application sets clmgmtLicenseActionRowStatus to
//     createAndGo(4) and all other required objects to valid
//     values in a single SNMP SET request. If all the values
//     are valid, the device creates the entry and executes the
//     action. If the SET request fails, the entry will not be
//     created.
// 2. CREATE-AND-WAIT method
//     Management application sets clmgmtLicenseActionRowStatus to
//     createAndWait(5) to create an entry. Management application
//     can set all other required objects to valid
//     values in more than one SNMP SET request. If SET request
//     for any of the objects fails, management application can
// set
//     just only that object. Once all the required objects
//     are set to valid values, management application triggers
// action
//     execution by setting clmgmtLicenseActionRowStatus to
//     active(1).
// 
// To stop the action from being executed, the management
// application
// can delete the entry by setting clmgmtLicenseActionRowStatus
// to destroy(6) when clmgmtLicenseActionState is pending(2).
// 
// The status of action execution can be known by querying
// clmgmtLicenseActionState. If the action is still in
// pending(2) or in inProgress(3) state, the management
// application need to check back again after few seconds.
// Once the action completes and status of the action is
// failed(6), the reason for failure can be retrieved
// from clmgmtLicenseActionFailCause. If the status of the
// action is partiallySuccessful(5), results of individual
// licenses can be queried from clmgmtLicenseActionResultTable.
// 
// Not all objects in the entry are needed to execute every
// action. Below is the list of actions and the required
// objects that are needed to be set for executing that
// action.
// 
// 1. Installing a license
//    The following MIB objects need to be set for installing a
//    license
//      a. clmgmtLicenseActionTransferProtocol
//      b. clmgmtLicenseServerAddressType
//      c. clmgmtLicenseServerAddress
//      d. clmgmtLicenseServerUsername
//      e. clmgmtLicenseServerPassword
//      f. clmgmtLicenseFile
//      g. clmgmtLicenseStore
//      h. clmgmtLicenseStopOnFailure
//      i. clmgmtLicenseAcceptEULA
//      j. clmgmtLicenseAction
// 
//    clmgmtLicenseActionEntPhysicalIndex need not be set
//    explicitly for license installs. License itself identifes
//    the device where the license needs to be installed.
// 
//    clmgmtLicenseStore need to be set to store the licenses
//    in a non-default license store. But, if a license file
//    has more than one license and licenses need to be
//    installed on multiple devices (for example to multiple
//    members with in a stack), then value of clmgmtLicenseStore
//    is ignored and the licenses will be installed in default
//    license stores of the respective devices.
// 
// 2. Clearing a license
//    The following MIB objects need to be set for clearing a
//    license
//      a. clmgmtLicenseActionEntPhysicalIndex
//      b. clmgmtLicenseActionLicenseIndex
//      c. clmgmtLicenseStore
//      d. clmgmtLicenseAction
// 
// 3. Revoking a license
//    The following MIB objects need to be set for revoking a
//    license
//      a. clmgmtLicenseActionTransferProtocol
//      b. clmgmtLicenseServerAddressType
//      c. clmgmtLicenseServerAddress
//      d. clmgmtLicenseServerUsername
//      e. clmgmtLicenseServerPassword
//      f. clmgmtLicensePermissionTicketFile
//      g. clmgmtLicenseRehostTicketFile
//      h. clmgmtLicenseStopOnFailure
//      i. clmgmtLicenseAction
// 
// 4. Regenerate last rehost ticket
//    The following MIB objects need to be set for regenerating
//    last rehost ticket
//      a. clmgmtLicenseActionTransferProtocol
//      b. clmgmtLicenseServerAddressType
//      c. clmgmtLicenseServerAddress
//      d. clmgmtLicenseServerUsername
//      e. clmgmtLicenseServerPassword
//      f. clmgmtLicensePermissionTicketFile
//      g. clmgmtLicenseRehostTicketFile
//      h. clmgmtLicenseStopOnFailure
//      i. clmgmtLicenseAction
// 
// 
// 5. Save all licenses to a backup storage
//    The following MIB objects need to be set for storing all
//    licenses to a backup store
//      a. clmgmtLicenseActionEntPhysicalIndex
//      b. clmgmtLicenseActionTransferProtocol
//      c. clmgmtLicenseServerAddressType
//      d. clmgmtLicenseServerAddress
//      e. clmgmtLicenseServerUsername
//      f. clmgmtLicenseServerPassword
//      g. clmgmtLicenseBackupFile
//      h. clmgmtLicenseAction
// 
// 6. Generate and export EULA if the licenses need EULA to be
//    accepted for installing.
//    The following MIB objects need to be set exporting required
//    EULAs
//      a. clmgmtLicenseActionTransferProtocol
//      b. clmgmtLicenseServerAddressType
//      c. clmgmtLicenseServerAddress
//      d. clmgmtLicenseServerUsername
//      e. clmgmtLicenseServerPassword
//      f. clmgmtLicenseFile
//      g. clmgmtLicenseEULAFile
//      h. clmgmtLicenseAction
// 
// For any action, if clmgmtLicenseActionTransferProtocol
// is set to local(2), the following objects need not be set.
//      a. clmgmtLicenseServerAddressType
//      b. clmgmtLicenseServerAddress
//      c. clmgmtLicenseServerUsername
//      d. clmgmtLicenseServerPassword
// 
// Entry can be deleted except when clmgmtLicenseAction is set
// to pending(2). All entries are volatile and are cleared
// on agent reset.
type CISCOLICENSEMGMTMIB_Clmgmtlicenseactiontable_Clmgmtlicenseactionentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This object uniquely identifies a row in
    // clmgmtLicenseActionTable. The management application should choose this
    // value by reading clmgmtNextFreeLicenseActionIndex while creating an entry
    // in this table. If an entry already exists with this index, the creation of
    // the entry will not continue and error will be returned. The management
    // application should read the value of clmgmtNextFreeLicenseActionIndex again
    // and retry with the new value for this object. The type is interface{} with
    // range: 1..4294967295.
    Clmgmtlicenseactionindex interface{}

    // This object represents the entPhysicalIndex of the device where the action
    // is being executed. This object is mainly used in devices where one device
    // is acting as a master and rest of the devices as slaves. The master device
    // is responsible for SNMP communication with the management application.
    // Examples include stackable switches, devices with route processor and line
    // card configuration. If this object is not set, the license action will be
    // executed on the master device. Note: This object need not be set if there
    // is a stand alone device. The type is interface{} with range: 0..2147483647.
    Clmgmtlicenseactionentphysicalindex interface{}

    // This object represents the transfer protocol to be used when copying files
    // as specified in the following objects. 1. clmgmtLicenseFile 2.
    // clmgmtLicensePermissionTicketFile 3. clmgmtLicenseRehostTicketFile 4.
    // clmgmtLicenseBackupFile  Note: This object need not be set if the all the
    // files required for the action are in device's local file system. The type
    // is ClmgmtLicenseTransferProtocol.
    Clmgmtlicenseactiontransferprotocol interface{}

    // This object indicates the transport type of the address contained in
    // clmgmtLicenseServerAddress object. This must be set when
    // clmgmtLicenseActionTransferProtocol is not none(1) or local(2). The type is
    // InetAddressType.
    Clmgmtlicenseserveraddresstype interface{}

    // This object indicates the ip address of the server from which the files
    // must be read or written to if clmgmtLicenseActionTransferProtocol is not
    // none(1) or local(2).  All bits as 0s or 1s for clmgmtLicenseServerAddress
    // are not allowed.  The format of this address depends on the value of the
    // clmgmtLicenseServerAddressType object. The type is string with length:
    // 0..255.
    Clmgmtlicenseserveraddress interface{}

    // This object indicates the remote user name for accessing files via ftp,
    // rcp, sftp or scp protocols. This object must be set when the
    // clmgmtLicenseActionTransferProtocol is ftp(4), rcp(5), scp(7) or sftp(8).
    // If clmgmtLicenseActionTransferProtocol is rcp(5), the remote username is
    // sent as the server username in an rcp command request sent by the system to
    // a remote rcp server. The type is string with length: 0..96.
    Clmgmtlicenseserverusername interface{}

    // This object indicates the password used by ftp, sftp or scp for copying a
    // file to/from an ftp/sftp/scp server. This object must be set when the
    // clmgmtLicenseActionTransferProtocol is ftp(4) or scp(7) or sftp(8). Reading
    // it returns a zero-length string for security reasons. The type is string
    // with length: 0..96.
    Clmgmtlicenseserverpassword interface{}

    // This object represents the location of the license file on the server
    // identified by clmgmtLicenseServerAddress. This object MUST be set to a
    // valid value before or concurrently with setting the value of the
    // clmgmtLicenseAction object to install(2). For other operations, the value
    // of this object is not considered, it is irrelevant. The type is string with
    // length: 0..255.
    Clmgmtlicensefile interface{}

    // This object represents the clmgmtLicenseStoreIndex of the license store to
    // use within the device. The license store can be a local disk or flash. A
    // device can have more than one license stores. If this object is not set,
    // the license will be stored in the default license store as exposed by
    // clmgmtDefaultLicenseStore object. The type is interface{} with range:
    // 0..4294967295.
    Clmgmtlicensestore interface{}

    // This object indicates the the license index of the license that is the
    // subject of this action. This is used for identifying a license for
    // performing actions specific to that license. This object need to be set
    // only if clmgmtLicenseAction is set to clear(4). The value of this object is
    // same as the clmgmtLicenseIndex object in clmgmtLicenseInfoEntry for license
    // that is subject of this action. The type is interface{} with range:
    // 0..4294967295.
    Clmgmtlicenseactionlicenseindex interface{}

    // This object indicates the file name of the permission ticket. This object
    // need to be set only if clmgmtLicenseAction is set to
    // processPermissionTicket(4) or regenerateLastRehostTicket(5) actions. The
    // permission ticket is obtained from Cisco licensing portal to revoke a
    // license. The management application must set this object to valid value
    // before invoking the action. The type is string with length: 0..255.
    Clmgmtlicensepermissionticketfile interface{}

    // This object indicates the file where the rehost ticket generated by the
    // device need to be exported to. The rehost ticket is generated as a result
    // of processPermissionTicket and regenerateLastRehostTicket actions. After
    // generating the rehost ticket, the device exports the rehost ticket contents
    // to this file. This object need to be set only if clmgmtLicenseAction is set
    // to processPermissionTicket(4) or regenerateLastRehostTicket(5) actions. The
    // type is string with length: 0..255.
    Clmgmtlicenserehostticketfile interface{}

    // This object indicates the file where all the licenses in the device need to
    // be backed up. This object need to be set only if clmgmtLicenseAction is set
    // to backup(6) and the management application must set the value of this 
    // object to valid value before invoking action. The type is string with
    // length: 0..255.
    Clmgmtlicensebackupfile interface{}

    // This object indicates whether the license action should stop if the action
    // on a license fails. This object is applicable only if there are more than
    // one licenses involved in an action. The type is bool.
    Clmgmtlicensestoponfailure interface{}

    // This object indicates the the command/action to be executed.  Command      
    // Remarks -------                        ------- noOp(1)                     
    // No operation will be                                performed.  install(2) 
    // Installs the license.  clear(3)                       Clears the license. 
    // processPermissionTicket(4)     Processes thee permission                   
    // ticket and generates and                                exports rehost
    // ticket.  regenerateLastRehostTicket(5)  Generates and exports the          
    // last generated rehost                                ticket.  backup(6)    
    // Backs up all the licenses                                installed
    // currently onto a                                backup store. 
    // generateEULA(7)                Checks whether the licenses                 
    // in the license file need EULA                                acceptance and
    // uploads the                                needed EULA contents to a file.
    // The type is Clmgmtlicenseaction.
    Clmgmtlicenseaction interface{}

    // This object indicates the state of this license action. The type is
    // ClmgmtLicenseActionState.
    Clmgmtlicenseactionstate interface{}

    // This object represents the position of the action in the license action job
    // queue that is maintained internally. Only actions in pending(2) state will
    // be put in the queue until they are executed. By reading this object, the
    // management application can make intelligent decision on whether to execute
    // another action that it is planning on. For example, if there is already a
    // license clear action in the queue in pending(2) state, management
    // application can choose to defer its license back up action to a later time.
    // This object will return a value of 0 if the action is not in pending(2)
    // state. The type is interface{} with range: 0..4294967295.
    Clmgmtlicensejobqposition interface{}

    // This object indicates the reason for this license action failure. The value
    // of this object is valid only when clmgmtLicenseActionState is failed(6).
    // The type is ClmgmtLicenseActionFailCause.
    Clmgmtlicenseactionfailcause interface{}

    // This object indicates the storage type for this conceptual row. Conceptual
    // rows having the value 'permanent' need not allow write-access to any
    // columnar objects in the row. The type is StorageType.
    Clmgmtlicenseactionstoragetype interface{}

    // This object indicates the the status of this table entry. Once the entry
    // status is set to active(1), the associated entry cannot be modified until
    // the action completes (clmgmtLicenseConfigCommandStatus is set to a value
    // other than inProgress(3)). Once the action completes the only operation
    // possible after this is to delete the row. It is recommended that the
    // management application should delete entries in this table after reading
    // the result. In order to prevent old entries from clogging the table,
    // entries will be aged out, but an entry will never be deleted within 5
    // minutes of completion. The type is RowStatus.
    Clmgmtlicenseactionrowstatus interface{}

    // This object indicates whether the End User License Agreement needed for
    // installing the licenses is accepted.  true(1) - EULA is read and accepted
    // false(2) - EULA is not accepted  Management application should set this
    // object to true(1) when installing licenses that need EULA acceptance. The
    // type is bool.
    Clmgmtlicenseaccepteula interface{}

    // This object indicates the file where all the End User License Agreements
    // (EULAs) need to be exported to. This object need to be set only if
    // clmgmtLicenseAction is set to generateEULA(7) and the management
    // application must set the value of this object to valid value before
    // invoking action. The type is string with length: 0..255.
    Clmgmtlicenseeulafile interface{}
}

func (clmgmtlicenseactionentry *CISCOLICENSEMGMTMIB_Clmgmtlicenseactiontable_Clmgmtlicenseactionentry) GetEntityData() *types.CommonEntityData {
    clmgmtlicenseactionentry.EntityData.YFilter = clmgmtlicenseactionentry.YFilter
    clmgmtlicenseactionentry.EntityData.YangName = "clmgmtLicenseActionEntry"
    clmgmtlicenseactionentry.EntityData.BundleName = "cisco_ios_xe"
    clmgmtlicenseactionentry.EntityData.ParentYangName = "clmgmtLicenseActionTable"
    clmgmtlicenseactionentry.EntityData.SegmentPath = "clmgmtLicenseActionEntry" + "[clmgmtLicenseActionIndex='" + fmt.Sprintf("%v", clmgmtlicenseactionentry.Clmgmtlicenseactionindex) + "']"
    clmgmtlicenseactionentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtlicenseactionentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtlicenseactionentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtlicenseactionentry.EntityData.Children = make(map[string]types.YChild)
    clmgmtlicenseactionentry.EntityData.Leafs = make(map[string]types.YLeaf)
    clmgmtlicenseactionentry.EntityData.Leafs["clmgmtLicenseActionIndex"] = types.YLeaf{"Clmgmtlicenseactionindex", clmgmtlicenseactionentry.Clmgmtlicenseactionindex}
    clmgmtlicenseactionentry.EntityData.Leafs["clmgmtLicenseActionEntPhysicalIndex"] = types.YLeaf{"Clmgmtlicenseactionentphysicalindex", clmgmtlicenseactionentry.Clmgmtlicenseactionentphysicalindex}
    clmgmtlicenseactionentry.EntityData.Leafs["clmgmtLicenseActionTransferProtocol"] = types.YLeaf{"Clmgmtlicenseactiontransferprotocol", clmgmtlicenseactionentry.Clmgmtlicenseactiontransferprotocol}
    clmgmtlicenseactionentry.EntityData.Leafs["clmgmtLicenseServerAddressType"] = types.YLeaf{"Clmgmtlicenseserveraddresstype", clmgmtlicenseactionentry.Clmgmtlicenseserveraddresstype}
    clmgmtlicenseactionentry.EntityData.Leafs["clmgmtLicenseServerAddress"] = types.YLeaf{"Clmgmtlicenseserveraddress", clmgmtlicenseactionentry.Clmgmtlicenseserveraddress}
    clmgmtlicenseactionentry.EntityData.Leafs["clmgmtLicenseServerUsername"] = types.YLeaf{"Clmgmtlicenseserverusername", clmgmtlicenseactionentry.Clmgmtlicenseserverusername}
    clmgmtlicenseactionentry.EntityData.Leafs["clmgmtLicenseServerPassword"] = types.YLeaf{"Clmgmtlicenseserverpassword", clmgmtlicenseactionentry.Clmgmtlicenseserverpassword}
    clmgmtlicenseactionentry.EntityData.Leafs["clmgmtLicenseFile"] = types.YLeaf{"Clmgmtlicensefile", clmgmtlicenseactionentry.Clmgmtlicensefile}
    clmgmtlicenseactionentry.EntityData.Leafs["clmgmtLicenseStore"] = types.YLeaf{"Clmgmtlicensestore", clmgmtlicenseactionentry.Clmgmtlicensestore}
    clmgmtlicenseactionentry.EntityData.Leafs["clmgmtLicenseActionLicenseIndex"] = types.YLeaf{"Clmgmtlicenseactionlicenseindex", clmgmtlicenseactionentry.Clmgmtlicenseactionlicenseindex}
    clmgmtlicenseactionentry.EntityData.Leafs["clmgmtLicensePermissionTicketFile"] = types.YLeaf{"Clmgmtlicensepermissionticketfile", clmgmtlicenseactionentry.Clmgmtlicensepermissionticketfile}
    clmgmtlicenseactionentry.EntityData.Leafs["clmgmtLicenseRehostTicketFile"] = types.YLeaf{"Clmgmtlicenserehostticketfile", clmgmtlicenseactionentry.Clmgmtlicenserehostticketfile}
    clmgmtlicenseactionentry.EntityData.Leafs["clmgmtLicenseBackupFile"] = types.YLeaf{"Clmgmtlicensebackupfile", clmgmtlicenseactionentry.Clmgmtlicensebackupfile}
    clmgmtlicenseactionentry.EntityData.Leafs["clmgmtLicenseStopOnFailure"] = types.YLeaf{"Clmgmtlicensestoponfailure", clmgmtlicenseactionentry.Clmgmtlicensestoponfailure}
    clmgmtlicenseactionentry.EntityData.Leafs["clmgmtLicenseAction"] = types.YLeaf{"Clmgmtlicenseaction", clmgmtlicenseactionentry.Clmgmtlicenseaction}
    clmgmtlicenseactionentry.EntityData.Leafs["clmgmtLicenseActionState"] = types.YLeaf{"Clmgmtlicenseactionstate", clmgmtlicenseactionentry.Clmgmtlicenseactionstate}
    clmgmtlicenseactionentry.EntityData.Leafs["clmgmtLicenseJobQPosition"] = types.YLeaf{"Clmgmtlicensejobqposition", clmgmtlicenseactionentry.Clmgmtlicensejobqposition}
    clmgmtlicenseactionentry.EntityData.Leafs["clmgmtLicenseActionFailCause"] = types.YLeaf{"Clmgmtlicenseactionfailcause", clmgmtlicenseactionentry.Clmgmtlicenseactionfailcause}
    clmgmtlicenseactionentry.EntityData.Leafs["clmgmtLicenseActionStorageType"] = types.YLeaf{"Clmgmtlicenseactionstoragetype", clmgmtlicenseactionentry.Clmgmtlicenseactionstoragetype}
    clmgmtlicenseactionentry.EntityData.Leafs["clmgmtLicenseActionRowStatus"] = types.YLeaf{"Clmgmtlicenseactionrowstatus", clmgmtlicenseactionentry.Clmgmtlicenseactionrowstatus}
    clmgmtlicenseactionentry.EntityData.Leafs["clmgmtLicenseAcceptEULA"] = types.YLeaf{"Clmgmtlicenseaccepteula", clmgmtlicenseactionentry.Clmgmtlicenseaccepteula}
    clmgmtlicenseactionentry.EntityData.Leafs["clmgmtLicenseEULAFile"] = types.YLeaf{"Clmgmtlicenseeulafile", clmgmtlicenseactionentry.Clmgmtlicenseeulafile}
    return &(clmgmtlicenseactionentry.EntityData)
}

// CISCOLICENSEMGMTMIB_Clmgmtlicenseactiontable_Clmgmtlicenseactionentry_Clmgmtlicenseaction represents                                needed EULA contents to a file.
type CISCOLICENSEMGMTMIB_Clmgmtlicenseactiontable_Clmgmtlicenseactionentry_Clmgmtlicenseaction string

const (
    CISCOLICENSEMGMTMIB_Clmgmtlicenseactiontable_Clmgmtlicenseactionentry_Clmgmtlicenseaction_noOp CISCOLICENSEMGMTMIB_Clmgmtlicenseactiontable_Clmgmtlicenseactionentry_Clmgmtlicenseaction = "noOp"

    CISCOLICENSEMGMTMIB_Clmgmtlicenseactiontable_Clmgmtlicenseactionentry_Clmgmtlicenseaction_install CISCOLICENSEMGMTMIB_Clmgmtlicenseactiontable_Clmgmtlicenseactionentry_Clmgmtlicenseaction = "install"

    CISCOLICENSEMGMTMIB_Clmgmtlicenseactiontable_Clmgmtlicenseactionentry_Clmgmtlicenseaction_clear CISCOLICENSEMGMTMIB_Clmgmtlicenseactiontable_Clmgmtlicenseactionentry_Clmgmtlicenseaction = "clear"

    CISCOLICENSEMGMTMIB_Clmgmtlicenseactiontable_Clmgmtlicenseactionentry_Clmgmtlicenseaction_processPermissionTicket CISCOLICENSEMGMTMIB_Clmgmtlicenseactiontable_Clmgmtlicenseactionentry_Clmgmtlicenseaction = "processPermissionTicket"

    CISCOLICENSEMGMTMIB_Clmgmtlicenseactiontable_Clmgmtlicenseactionentry_Clmgmtlicenseaction_regenerateLastRehostTicket CISCOLICENSEMGMTMIB_Clmgmtlicenseactiontable_Clmgmtlicenseactionentry_Clmgmtlicenseaction = "regenerateLastRehostTicket"

    CISCOLICENSEMGMTMIB_Clmgmtlicenseactiontable_Clmgmtlicenseactionentry_Clmgmtlicenseaction_backup CISCOLICENSEMGMTMIB_Clmgmtlicenseactiontable_Clmgmtlicenseactionentry_Clmgmtlicenseaction = "backup"

    CISCOLICENSEMGMTMIB_Clmgmtlicenseactiontable_Clmgmtlicenseactionentry_Clmgmtlicenseaction_generateEULA CISCOLICENSEMGMTMIB_Clmgmtlicenseactiontable_Clmgmtlicenseactionentry_Clmgmtlicenseaction = "generateEULA"
)

// CISCOLICENSEMGMTMIB_Clmgmtlicenseactionresulttable
// This table contains results of license action if the
// license action involves multiple licenses. Entries in this
// table are not created for actions where there is
// only license that is subject of the action. For
// example, if there are 3 licenses in a license file
// when executing license install action, 3 entries will
// be created in this table, one for each license.
type CISCOLICENSEMGMTMIB_Clmgmtlicenseactionresulttable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in clmgmtLicenseActionResultTable. Each entry contains result of
    // the action for a single license. These entries are created immediately
    // after action execution when the action involves multiple licenses. These
    // entries get automatically deleted when the corresponding entry in
    // clmgmtLicenseActionTable is deleted. The type is slice of
    // CISCOLICENSEMGMTMIB_Clmgmtlicenseactionresulttable_Clmgmtlicenseactionresultentry.
    Clmgmtlicenseactionresultentry []CISCOLICENSEMGMTMIB_Clmgmtlicenseactionresulttable_Clmgmtlicenseactionresultentry
}

func (clmgmtlicenseactionresulttable *CISCOLICENSEMGMTMIB_Clmgmtlicenseactionresulttable) GetEntityData() *types.CommonEntityData {
    clmgmtlicenseactionresulttable.EntityData.YFilter = clmgmtlicenseactionresulttable.YFilter
    clmgmtlicenseactionresulttable.EntityData.YangName = "clmgmtLicenseActionResultTable"
    clmgmtlicenseactionresulttable.EntityData.BundleName = "cisco_ios_xe"
    clmgmtlicenseactionresulttable.EntityData.ParentYangName = "CISCO-LICENSE-MGMT-MIB"
    clmgmtlicenseactionresulttable.EntityData.SegmentPath = "clmgmtLicenseActionResultTable"
    clmgmtlicenseactionresulttable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtlicenseactionresulttable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtlicenseactionresulttable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtlicenseactionresulttable.EntityData.Children = make(map[string]types.YChild)
    clmgmtlicenseactionresulttable.EntityData.Children["clmgmtLicenseActionResultEntry"] = types.YChild{"Clmgmtlicenseactionresultentry", nil}
    for i := range clmgmtlicenseactionresulttable.Clmgmtlicenseactionresultentry {
        clmgmtlicenseactionresulttable.EntityData.Children[types.GetSegmentPath(&clmgmtlicenseactionresulttable.Clmgmtlicenseactionresultentry[i])] = types.YChild{"Clmgmtlicenseactionresultentry", &clmgmtlicenseactionresulttable.Clmgmtlicenseactionresultentry[i]}
    }
    clmgmtlicenseactionresulttable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clmgmtlicenseactionresulttable.EntityData)
}

// CISCOLICENSEMGMTMIB_Clmgmtlicenseactionresulttable_Clmgmtlicenseactionresultentry
// An entry in clmgmtLicenseActionResultTable. Each entry
// contains result of the action for a single license.
// These entries are created immediately after action
// execution when the action involves multiple licenses.
// These entries get automatically deleted when the
// corresponding entry in clmgmtLicenseActionTable
// is deleted.
type CISCOLICENSEMGMTMIB_Clmgmtlicenseactionresulttable_Clmgmtlicenseactionresultentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_license_mgmt_mib.CISCOLICENSEMGMTMIB_Clmgmtlicenseactiontable_Clmgmtlicenseactionentry_Clmgmtlicenseactionindex
    Clmgmtlicenseactionindex interface{}

    // This attribute is a key. This object indicates the sequence number of this
    // license in the list of licenses on which the action is executed. For
    // example, if there are 3 licenses in a license file when executing license
    // install action, this object will have values 1, 2 and 3 respectively as
    // ordered in the license file. The type is interface{} with range:
    // 1..4294967295.
    Clmgmtlicensenumber interface{}

    // This object indicates the state of action on this individual license. The
    // type is ClmgmtLicenseActionState.
    Clmgmtlicenseindivactionstate interface{}

    // This object indicates the reason for action failure on this individual
    // license. The type is ClmgmtLicenseActionFailCause.
    Clmgmtlicenseindivactionfailcause interface{}
}

func (clmgmtlicenseactionresultentry *CISCOLICENSEMGMTMIB_Clmgmtlicenseactionresulttable_Clmgmtlicenseactionresultentry) GetEntityData() *types.CommonEntityData {
    clmgmtlicenseactionresultentry.EntityData.YFilter = clmgmtlicenseactionresultentry.YFilter
    clmgmtlicenseactionresultentry.EntityData.YangName = "clmgmtLicenseActionResultEntry"
    clmgmtlicenseactionresultentry.EntityData.BundleName = "cisco_ios_xe"
    clmgmtlicenseactionresultentry.EntityData.ParentYangName = "clmgmtLicenseActionResultTable"
    clmgmtlicenseactionresultentry.EntityData.SegmentPath = "clmgmtLicenseActionResultEntry" + "[clmgmtLicenseActionIndex='" + fmt.Sprintf("%v", clmgmtlicenseactionresultentry.Clmgmtlicenseactionindex) + "']" + "[clmgmtLicenseNumber='" + fmt.Sprintf("%v", clmgmtlicenseactionresultentry.Clmgmtlicensenumber) + "']"
    clmgmtlicenseactionresultentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtlicenseactionresultentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtlicenseactionresultentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtlicenseactionresultentry.EntityData.Children = make(map[string]types.YChild)
    clmgmtlicenseactionresultentry.EntityData.Leafs = make(map[string]types.YLeaf)
    clmgmtlicenseactionresultentry.EntityData.Leafs["clmgmtLicenseActionIndex"] = types.YLeaf{"Clmgmtlicenseactionindex", clmgmtlicenseactionresultentry.Clmgmtlicenseactionindex}
    clmgmtlicenseactionresultentry.EntityData.Leafs["clmgmtLicenseNumber"] = types.YLeaf{"Clmgmtlicensenumber", clmgmtlicenseactionresultentry.Clmgmtlicensenumber}
    clmgmtlicenseactionresultentry.EntityData.Leafs["clmgmtLicenseIndivActionState"] = types.YLeaf{"Clmgmtlicenseindivactionstate", clmgmtlicenseactionresultentry.Clmgmtlicenseindivactionstate}
    clmgmtlicenseactionresultentry.EntityData.Leafs["clmgmtLicenseIndivActionFailCause"] = types.YLeaf{"Clmgmtlicenseindivactionfailcause", clmgmtlicenseactionresultentry.Clmgmtlicenseindivactionfailcause}
    return &(clmgmtlicenseactionresultentry.EntityData)
}

// CISCOLICENSEMGMTMIB_Clmgmtlicensestoreinfotable
// This table contains information about all the license
// stores allocated on the device.
type CISCOLICENSEMGMTMIB_Clmgmtlicensestoreinfotable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in clmgmtLicenseStoreInfoTable. Each entry contains information
    // about a license store allocated on the device. The type is slice of
    // CISCOLICENSEMGMTMIB_Clmgmtlicensestoreinfotable_Clmgmtlicensestoreinfoentry.
    Clmgmtlicensestoreinfoentry []CISCOLICENSEMGMTMIB_Clmgmtlicensestoreinfotable_Clmgmtlicensestoreinfoentry
}

func (clmgmtlicensestoreinfotable *CISCOLICENSEMGMTMIB_Clmgmtlicensestoreinfotable) GetEntityData() *types.CommonEntityData {
    clmgmtlicensestoreinfotable.EntityData.YFilter = clmgmtlicensestoreinfotable.YFilter
    clmgmtlicensestoreinfotable.EntityData.YangName = "clmgmtLicenseStoreInfoTable"
    clmgmtlicensestoreinfotable.EntityData.BundleName = "cisco_ios_xe"
    clmgmtlicensestoreinfotable.EntityData.ParentYangName = "CISCO-LICENSE-MGMT-MIB"
    clmgmtlicensestoreinfotable.EntityData.SegmentPath = "clmgmtLicenseStoreInfoTable"
    clmgmtlicensestoreinfotable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtlicensestoreinfotable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtlicensestoreinfotable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtlicensestoreinfotable.EntityData.Children = make(map[string]types.YChild)
    clmgmtlicensestoreinfotable.EntityData.Children["clmgmtLicenseStoreInfoEntry"] = types.YChild{"Clmgmtlicensestoreinfoentry", nil}
    for i := range clmgmtlicensestoreinfotable.Clmgmtlicensestoreinfoentry {
        clmgmtlicensestoreinfotable.EntityData.Children[types.GetSegmentPath(&clmgmtlicensestoreinfotable.Clmgmtlicensestoreinfoentry[i])] = types.YChild{"Clmgmtlicensestoreinfoentry", &clmgmtlicensestoreinfotable.Clmgmtlicensestoreinfoentry[i]}
    }
    clmgmtlicensestoreinfotable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clmgmtlicensestoreinfotable.EntityData)
}

// CISCOLICENSEMGMTMIB_Clmgmtlicensestoreinfotable_Clmgmtlicensestoreinfoentry
// An entry in clmgmtLicenseStoreInfoTable. Each entry
// contains information about a license store allocated
// on the device
type CISCOLICENSEMGMTMIB_Clmgmtlicensestoreinfotable_Clmgmtlicensestoreinfoentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. This object uniquely identifies a license store
    // within the device. The type is interface{} with range: 1..4294967295.
    Clmgmtlicensestoreindex interface{}

    // This object indicates the name of the license store within the device. It
    // is a file in device's local file system i.e., either on a local disk or
    // flash or some other storage media. For example, the value of this object
    // can be 'disk1:lic_store_1.txt' or 'flash:lic_store_2.txt. The type is
    // string with length: 0..255.
    Clmgmtlicensestorename interface{}

    // This object indicates the total number of bytes that are allocated to the
    // license store. The type is interface{} with range: 0..4294967295. Units are
    // bytes.
    Clmgmtlicensestoretotalsize interface{}

    // This object indicates the number of bytes still remaining to be used for
    // new license installations in the license store. The type is interface{}
    // with range: 0..4294967295. Units are bytes.
    Clmgmtlicensestoresizeremaining interface{}
}

func (clmgmtlicensestoreinfoentry *CISCOLICENSEMGMTMIB_Clmgmtlicensestoreinfotable_Clmgmtlicensestoreinfoentry) GetEntityData() *types.CommonEntityData {
    clmgmtlicensestoreinfoentry.EntityData.YFilter = clmgmtlicensestoreinfoentry.YFilter
    clmgmtlicensestoreinfoentry.EntityData.YangName = "clmgmtLicenseStoreInfoEntry"
    clmgmtlicensestoreinfoentry.EntityData.BundleName = "cisco_ios_xe"
    clmgmtlicensestoreinfoentry.EntityData.ParentYangName = "clmgmtLicenseStoreInfoTable"
    clmgmtlicensestoreinfoentry.EntityData.SegmentPath = "clmgmtLicenseStoreInfoEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", clmgmtlicensestoreinfoentry.Entphysicalindex) + "']" + "[clmgmtLicenseStoreIndex='" + fmt.Sprintf("%v", clmgmtlicensestoreinfoentry.Clmgmtlicensestoreindex) + "']"
    clmgmtlicensestoreinfoentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtlicensestoreinfoentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtlicensestoreinfoentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtlicensestoreinfoentry.EntityData.Children = make(map[string]types.YChild)
    clmgmtlicensestoreinfoentry.EntityData.Leafs = make(map[string]types.YLeaf)
    clmgmtlicensestoreinfoentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", clmgmtlicensestoreinfoentry.Entphysicalindex}
    clmgmtlicensestoreinfoentry.EntityData.Leafs["clmgmtLicenseStoreIndex"] = types.YLeaf{"Clmgmtlicensestoreindex", clmgmtlicensestoreinfoentry.Clmgmtlicensestoreindex}
    clmgmtlicensestoreinfoentry.EntityData.Leafs["clmgmtLicenseStoreName"] = types.YLeaf{"Clmgmtlicensestorename", clmgmtlicensestoreinfoentry.Clmgmtlicensestorename}
    clmgmtlicensestoreinfoentry.EntityData.Leafs["clmgmtLicenseStoreTotalSize"] = types.YLeaf{"Clmgmtlicensestoretotalsize", clmgmtlicensestoreinfoentry.Clmgmtlicensestoretotalsize}
    clmgmtlicensestoreinfoentry.EntityData.Leafs["clmgmtLicenseStoreSizeRemaining"] = types.YLeaf{"Clmgmtlicensestoresizeremaining", clmgmtlicensestoreinfoentry.Clmgmtlicensestoresizeremaining}
    return &(clmgmtlicensestoreinfoentry.EntityData)
}

// CISCOLICENSEMGMTMIB_Clmgmtlicensedeviceinfotable
// This table contains objects that provide licensing related
// information at the device level. Entries will exist
// only for entities that support licensing. For example,
// if it is a stand alone device and supports licensing,
// then there will be only one entry in this table. If
// it is stackable switch then there will be multiple
// entries with one entry for each device in the stack.
type CISCOLICENSEMGMTMIB_Clmgmtlicensedeviceinfotable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in clmgmtLicenseDeviceInfoTable. Each entry contains device level
    // licensing information for a device. The type is slice of
    // CISCOLICENSEMGMTMIB_Clmgmtlicensedeviceinfotable_Clmgmtlicensedeviceinfoentry.
    Clmgmtlicensedeviceinfoentry []CISCOLICENSEMGMTMIB_Clmgmtlicensedeviceinfotable_Clmgmtlicensedeviceinfoentry
}

func (clmgmtlicensedeviceinfotable *CISCOLICENSEMGMTMIB_Clmgmtlicensedeviceinfotable) GetEntityData() *types.CommonEntityData {
    clmgmtlicensedeviceinfotable.EntityData.YFilter = clmgmtlicensedeviceinfotable.YFilter
    clmgmtlicensedeviceinfotable.EntityData.YangName = "clmgmtLicenseDeviceInfoTable"
    clmgmtlicensedeviceinfotable.EntityData.BundleName = "cisco_ios_xe"
    clmgmtlicensedeviceinfotable.EntityData.ParentYangName = "CISCO-LICENSE-MGMT-MIB"
    clmgmtlicensedeviceinfotable.EntityData.SegmentPath = "clmgmtLicenseDeviceInfoTable"
    clmgmtlicensedeviceinfotable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtlicensedeviceinfotable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtlicensedeviceinfotable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtlicensedeviceinfotable.EntityData.Children = make(map[string]types.YChild)
    clmgmtlicensedeviceinfotable.EntityData.Children["clmgmtLicenseDeviceInfoEntry"] = types.YChild{"Clmgmtlicensedeviceinfoentry", nil}
    for i := range clmgmtlicensedeviceinfotable.Clmgmtlicensedeviceinfoentry {
        clmgmtlicensedeviceinfotable.EntityData.Children[types.GetSegmentPath(&clmgmtlicensedeviceinfotable.Clmgmtlicensedeviceinfoentry[i])] = types.YChild{"Clmgmtlicensedeviceinfoentry", &clmgmtlicensedeviceinfotable.Clmgmtlicensedeviceinfoentry[i]}
    }
    clmgmtlicensedeviceinfotable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clmgmtlicensedeviceinfotable.EntityData)
}

// CISCOLICENSEMGMTMIB_Clmgmtlicensedeviceinfotable_Clmgmtlicensedeviceinfoentry
// An entry in clmgmtLicenseDeviceInfoTable. Each entry
// contains device level licensing information for a device.
type CISCOLICENSEMGMTMIB_Clmgmtlicensedeviceinfotable_Clmgmtlicensedeviceinfoentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This object indicates the clmgmtLicenseStoreIndex of default store in the
    // device. There will be only one default license store per device. If no
    // license store is specified during license install, this default license
    // store will be used. The type is interface{} with range: 1..4294967295.
    Clmgmtdefaultlicensestore interface{}
}

func (clmgmtlicensedeviceinfoentry *CISCOLICENSEMGMTMIB_Clmgmtlicensedeviceinfotable_Clmgmtlicensedeviceinfoentry) GetEntityData() *types.CommonEntityData {
    clmgmtlicensedeviceinfoentry.EntityData.YFilter = clmgmtlicensedeviceinfoentry.YFilter
    clmgmtlicensedeviceinfoentry.EntityData.YangName = "clmgmtLicenseDeviceInfoEntry"
    clmgmtlicensedeviceinfoentry.EntityData.BundleName = "cisco_ios_xe"
    clmgmtlicensedeviceinfoentry.EntityData.ParentYangName = "clmgmtLicenseDeviceInfoTable"
    clmgmtlicensedeviceinfoentry.EntityData.SegmentPath = "clmgmtLicenseDeviceInfoEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", clmgmtlicensedeviceinfoentry.Entphysicalindex) + "']"
    clmgmtlicensedeviceinfoentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtlicensedeviceinfoentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtlicensedeviceinfoentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtlicensedeviceinfoentry.EntityData.Children = make(map[string]types.YChild)
    clmgmtlicensedeviceinfoentry.EntityData.Leafs = make(map[string]types.YLeaf)
    clmgmtlicensedeviceinfoentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", clmgmtlicensedeviceinfoentry.Entphysicalindex}
    clmgmtlicensedeviceinfoentry.EntityData.Leafs["clmgmtDefaultLicenseStore"] = types.YLeaf{"Clmgmtdefaultlicensestore", clmgmtlicensedeviceinfoentry.Clmgmtdefaultlicensestore}
    return &(clmgmtlicensedeviceinfoentry.EntityData)
}

// CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable
// This table contains information about all the licenses
// installed on the device.
type CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in clmgmtLicenseInfoTable. Each entry contains information about a
    // license installed on the device. This entry gets created when a license is
    // installed successfully. Management application can not create these entries
    // directly, but will do so indirectly by executing license install action.
    // Some of these entries may already exist that correspond to demo licenses
    // even before management application installs any licenses. The type is slice
    // of CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry.
    Clmgmtlicenseinfoentry []CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry
}

func (clmgmtlicenseinfotable *CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable) GetEntityData() *types.CommonEntityData {
    clmgmtlicenseinfotable.EntityData.YFilter = clmgmtlicenseinfotable.YFilter
    clmgmtlicenseinfotable.EntityData.YangName = "clmgmtLicenseInfoTable"
    clmgmtlicenseinfotable.EntityData.BundleName = "cisco_ios_xe"
    clmgmtlicenseinfotable.EntityData.ParentYangName = "CISCO-LICENSE-MGMT-MIB"
    clmgmtlicenseinfotable.EntityData.SegmentPath = "clmgmtLicenseInfoTable"
    clmgmtlicenseinfotable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtlicenseinfotable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtlicenseinfotable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtlicenseinfotable.EntityData.Children = make(map[string]types.YChild)
    clmgmtlicenseinfotable.EntityData.Children["clmgmtLicenseInfoEntry"] = types.YChild{"Clmgmtlicenseinfoentry", nil}
    for i := range clmgmtlicenseinfotable.Clmgmtlicenseinfoentry {
        clmgmtlicenseinfotable.EntityData.Children[types.GetSegmentPath(&clmgmtlicenseinfotable.Clmgmtlicenseinfoentry[i])] = types.YChild{"Clmgmtlicenseinfoentry", &clmgmtlicenseinfotable.Clmgmtlicenseinfoentry[i]}
    }
    clmgmtlicenseinfotable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clmgmtlicenseinfotable.EntityData)
}

// CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry
// An entry in clmgmtLicenseInfoTable. Each entry contains
// information about a license installed on the device. This
// entry gets created when a license is installed successfully.
// Management application can not create these entries directly, but
// will do so indirectly by executing license install action.
// Some of these entries may already exist that correspond to
// demo licenses even before management application installs any
// licenses.
type CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. This object represents the license store that is
    // used for storing this license. This object will have the same value as
    // clmgmtLicenseStoreIndex in clmgmtLicenseStoreInfoEntry of the license store
    // used. The type is interface{} with range: 1..4294967295.
    Clmgmtlicensestoreused interface{}

    // This attribute is a key. This object uniquely identifies a license within
    // the device. The type is interface{} with range: 1..4294967295.
    Clmgmtlicenseindex interface{}

    // This object indicates the name of the feature that is using or can use this
    // license. A license can be used by only one feature. Examples of feature
    // name are: 'IPBASE', 'ADVIPSERVICE'. The type is string with length: 0..128.
    Clmgmtlicensefeaturename interface{}

    // This object indicates the version of the feature that is using or can use
    // this license. Examples of feature version are: '1.0', '2.0'. The type is
    // string with length: 0..128.
    Clmgmtlicensefeatureversion interface{}

    // This object identifies type of license. Licenses may have validity period
    // defined in terms of time duration that the license is valid for or it may
    // be defined in terms of actual calendar dates. Subscription licenses are
    // licenses that have validity period defined in terms of calendar dates. 
    // demo(1)               - demo(evaluation license) license. extension(2)     
    // - Extension(expiring) license. gracePeriod(3)        - Grace period
    // license. permanent(4)          - permanent license, the license has no     
    // expiry date. paidSubscription(5)   - Paid subscription licenses are the
    // licenses                         which are purchased by customers. These   
    // licenses have a start date  and end date                         associated
    // with them. evaluationSubscription(6)-Evaluation subscription licenses are  
    // the trial licenses. These licenses                           are node
    // locked and it can be obtained                           only once for an
    // UDI. They are valid                           based on calendar days. These
    // licenses                           have a start date and an end date       
    // associated with them and are issued                           once per UDI.
    // extensionSubscription(7)- Extension subscription licenses are              
    // similar to evaluation subscription                           licenses but
    // these licenses are issued                           based on customer
    // request. There are                           no restrictions on the number
    // of                           licenses available for a UDI.
    // evalRightToUse(8)       - Evaluation Right to use (RTU) license.
    // rightToUse(9)           - Right to use (RTU) license.
    // permanentRightToUse(10) ? Right To Use license right after it is configured
    // and is valid for the lifetime of the product.                           
    // This is a Right To Use license which is not in                           
    // evaluation mode for a limited time. The type is Clmgmtlicensetype.
    Clmgmtlicensetype interface{}

    // This object indicates whether the license is counted license. true(1)  -
    // counted license false(2) - uncounted license. The type is bool.
    Clmgmtlicensecounted interface{}

    // This object indicates the time period the license is valid for. This object
    // is applicable only if clmgmtLicenseType is demo(1), or extension(2) or
    // gracePeriod(3) or evalRightToUse(8). The object will return 0 for other
    // license types. The type is interface{} with range: 0..4294967295. Units are
    // seconds.
    Clmgmtlicensevalidityperiod interface{}

    // This object indicates the time period remaining before the license expires
    // or transitions to rightToUse(9) license. This object is applicable only if
    // clmgmtLicenseType is demo(1), or extension(2) or gracePeriod(3) or
    // evalRightToUse(8). The object will contain 0 for other license types. The
    // type is interface{} with range: 0..4294967295. Units are seconds.
    Clmgmtlicensevalidityperiodremaining interface{}

    // This object indicates the elapsed time period since the license expired.
    // This object is applicable only if clmgmtLicenseType is demo(1), or
    // extension(2) or gracePeriod(3). Also, this value of this object will be
    // valid only after the license expires. The object will return 0 for other
    // license types or before the license expiry. The type is interface{} with
    // range: 0..4294967295. Units are seconds.
    Clmgmtlicenseexpiredperiod interface{}

    // This object indicates the maximum number of entities that can use this
    // license. This object is applicable only if clmgmtLicenseCounted is true(1).
    // The entity that is being counted can be anything and it depends on the
    // licensable feature. The type is interface{} with range: 0..4294967295.
    Clmgmtlicensemaxusagecount interface{}

    // This object indicates the number of entities that can still use this
    // license. This object is applicable only if clmgmtLicenseCounted is true(1).
    // The type is interface{} with range: 0..4294967295.
    Clmgmtlicenseusagecountremaining interface{}

    // This object indicates whether the user accepted End User License Agreement
    // for this license.  true(1)  - EULA accpeted false(2) - EULA not accepted.
    // The type is bool.
    Clmgmtlicenseeulastatus interface{}

    // This object represents the user modifiable comments about the license. This
    // object is initially populated with comments from the license file. The type
    // is string with length: 0..255.
    Clmgmtlicensecomments interface{}

    // This object represents status of the license.  inactive(1)           -
    // license is installed, but                         not active. notInUse(2)  
    // - license is installed and                         available for use.
    // inUse(3)              - the license is being used (by                      
    // a feature). expiredInUse(4)       - license is expired but still           
    // being held by the feature. expiredNotInUse(5)    - license is expired and
    // not being                         held by any feature.
    // usageCountConsumed(6) - number of entities using this                      
    // licenses has reached the allowed                         limit, no new
    // entities are allowed                         to use this license. The type
    // is Clmgmtlicensestatus.
    Clmgmtlicensestatus interface{}

    // This object indicates the start date for a subscription license. It is
    // optional for subscription linceses to have a start date associated with
    // them, they may only have an end date associated with them. This object may
    // be applicable only when clmgmtLicenseType is paidSubscription(5),
    // evaluationSubscription(6) or extensionSubscription (7).       The object
    // will contain an octet string of length 0 when it is not applicable. The
    // type is string.
    Clmgmtlicensestartdate interface{}

    // This object indicates the end date for a subscription license. This object
    // is applicable only when clmgmtLicenseType is paidSubscription(5),
    // evaluationSubscription(6) or extensionSubscription (7). The object will
    // contain an octet string of length 0 when it is not applicable. The type is
    // string.
    Clmgmtlicenseenddate interface{}

    // This object indicates the time period used for the Right to use (RTU)
    // licenses. This object is applicable for all RTU licenses. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    Clmgmtlicenseperiodused interface{}
}

func (clmgmtlicenseinfoentry *CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry) GetEntityData() *types.CommonEntityData {
    clmgmtlicenseinfoentry.EntityData.YFilter = clmgmtlicenseinfoentry.YFilter
    clmgmtlicenseinfoentry.EntityData.YangName = "clmgmtLicenseInfoEntry"
    clmgmtlicenseinfoentry.EntityData.BundleName = "cisco_ios_xe"
    clmgmtlicenseinfoentry.EntityData.ParentYangName = "clmgmtLicenseInfoTable"
    clmgmtlicenseinfoentry.EntityData.SegmentPath = "clmgmtLicenseInfoEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", clmgmtlicenseinfoentry.Entphysicalindex) + "']" + "[clmgmtLicenseStoreUsed='" + fmt.Sprintf("%v", clmgmtlicenseinfoentry.Clmgmtlicensestoreused) + "']" + "[clmgmtLicenseIndex='" + fmt.Sprintf("%v", clmgmtlicenseinfoentry.Clmgmtlicenseindex) + "']"
    clmgmtlicenseinfoentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtlicenseinfoentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtlicenseinfoentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtlicenseinfoentry.EntityData.Children = make(map[string]types.YChild)
    clmgmtlicenseinfoentry.EntityData.Leafs = make(map[string]types.YLeaf)
    clmgmtlicenseinfoentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", clmgmtlicenseinfoentry.Entphysicalindex}
    clmgmtlicenseinfoentry.EntityData.Leafs["clmgmtLicenseStoreUsed"] = types.YLeaf{"Clmgmtlicensestoreused", clmgmtlicenseinfoentry.Clmgmtlicensestoreused}
    clmgmtlicenseinfoentry.EntityData.Leafs["clmgmtLicenseIndex"] = types.YLeaf{"Clmgmtlicenseindex", clmgmtlicenseinfoentry.Clmgmtlicenseindex}
    clmgmtlicenseinfoentry.EntityData.Leafs["clmgmtLicenseFeatureName"] = types.YLeaf{"Clmgmtlicensefeaturename", clmgmtlicenseinfoentry.Clmgmtlicensefeaturename}
    clmgmtlicenseinfoentry.EntityData.Leafs["clmgmtLicenseFeatureVersion"] = types.YLeaf{"Clmgmtlicensefeatureversion", clmgmtlicenseinfoentry.Clmgmtlicensefeatureversion}
    clmgmtlicenseinfoentry.EntityData.Leafs["clmgmtLicenseType"] = types.YLeaf{"Clmgmtlicensetype", clmgmtlicenseinfoentry.Clmgmtlicensetype}
    clmgmtlicenseinfoentry.EntityData.Leafs["clmgmtLicenseCounted"] = types.YLeaf{"Clmgmtlicensecounted", clmgmtlicenseinfoentry.Clmgmtlicensecounted}
    clmgmtlicenseinfoentry.EntityData.Leafs["clmgmtLicenseValidityPeriod"] = types.YLeaf{"Clmgmtlicensevalidityperiod", clmgmtlicenseinfoentry.Clmgmtlicensevalidityperiod}
    clmgmtlicenseinfoentry.EntityData.Leafs["clmgmtLicenseValidityPeriodRemaining"] = types.YLeaf{"Clmgmtlicensevalidityperiodremaining", clmgmtlicenseinfoentry.Clmgmtlicensevalidityperiodremaining}
    clmgmtlicenseinfoentry.EntityData.Leafs["clmgmtLicenseExpiredPeriod"] = types.YLeaf{"Clmgmtlicenseexpiredperiod", clmgmtlicenseinfoentry.Clmgmtlicenseexpiredperiod}
    clmgmtlicenseinfoentry.EntityData.Leafs["clmgmtLicenseMaxUsageCount"] = types.YLeaf{"Clmgmtlicensemaxusagecount", clmgmtlicenseinfoentry.Clmgmtlicensemaxusagecount}
    clmgmtlicenseinfoentry.EntityData.Leafs["clmgmtLicenseUsageCountRemaining"] = types.YLeaf{"Clmgmtlicenseusagecountremaining", clmgmtlicenseinfoentry.Clmgmtlicenseusagecountremaining}
    clmgmtlicenseinfoentry.EntityData.Leafs["clmgmtLicenseEULAStatus"] = types.YLeaf{"Clmgmtlicenseeulastatus", clmgmtlicenseinfoentry.Clmgmtlicenseeulastatus}
    clmgmtlicenseinfoentry.EntityData.Leafs["clmgmtLicenseComments"] = types.YLeaf{"Clmgmtlicensecomments", clmgmtlicenseinfoentry.Clmgmtlicensecomments}
    clmgmtlicenseinfoentry.EntityData.Leafs["clmgmtLicenseStatus"] = types.YLeaf{"Clmgmtlicensestatus", clmgmtlicenseinfoentry.Clmgmtlicensestatus}
    clmgmtlicenseinfoentry.EntityData.Leafs["clmgmtLicenseStartDate"] = types.YLeaf{"Clmgmtlicensestartdate", clmgmtlicenseinfoentry.Clmgmtlicensestartdate}
    clmgmtlicenseinfoentry.EntityData.Leafs["clmgmtLicenseEndDate"] = types.YLeaf{"Clmgmtlicenseenddate", clmgmtlicenseinfoentry.Clmgmtlicenseenddate}
    clmgmtlicenseinfoentry.EntityData.Leafs["clmgmtLicensePeriodUsed"] = types.YLeaf{"Clmgmtlicenseperiodused", clmgmtlicenseinfoentry.Clmgmtlicenseperiodused}
    return &(clmgmtlicenseinfoentry.EntityData)
}

// CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensestatus represents                         to use this license.
type CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensestatus string

const (
    CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensestatus_inactive CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensestatus = "inactive"

    CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensestatus_notInUse CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensestatus = "notInUse"

    CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensestatus_inUse CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensestatus = "inUse"

    CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensestatus_expiredInUse CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensestatus = "expiredInUse"

    CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensestatus_expiredNotInUse CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensestatus = "expiredNotInUse"

    CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensestatus_usageCountConsumed CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensestatus = "usageCountConsumed"
)

// CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensetype represents                           evaluation mode for a limited time.
type CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensetype string

const (
    CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensetype_demo CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensetype = "demo"

    CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensetype_extension CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensetype = "extension"

    CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensetype_gracePeriod CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensetype = "gracePeriod"

    CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensetype_permanent CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensetype = "permanent"

    CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensetype_paidSubscription CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensetype = "paidSubscription"

    CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensetype_evaluationSubscription CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensetype = "evaluationSubscription"

    CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensetype_extensionSubscription CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensetype = "extensionSubscription"

    CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensetype_evalRightToUse CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensetype = "evalRightToUse"

    CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensetype_rightToUse CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensetype = "rightToUse"

    CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensetype_permanentRightToUse CISCOLICENSEMGMTMIB_Clmgmtlicenseinfotable_Clmgmtlicenseinfoentry_Clmgmtlicensetype = "permanentRightToUse"
)

// CISCOLICENSEMGMTMIB_Clmgmtlicensablefeaturetable
// This table contains list of licensable features in the
// image. All the licensable features will have an entry each
// in this table irrespective of whether they are using any
// licenses currently. Entries in this table are created by
// the agent one for each licensable feature in the image.
// These entries remain in the table permanently and can not
// be deleted. Management application can not create or delete
// entries from this table.
type CISCOLICENSEMGMTMIB_Clmgmtlicensablefeaturetable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in clmgmtLicensableFeatureTable. Each entry represents a
    // licensable feature. The type is slice of
    // CISCOLICENSEMGMTMIB_Clmgmtlicensablefeaturetable_Clmgmtlicensablefeatureentry.
    Clmgmtlicensablefeatureentry []CISCOLICENSEMGMTMIB_Clmgmtlicensablefeaturetable_Clmgmtlicensablefeatureentry
}

func (clmgmtlicensablefeaturetable *CISCOLICENSEMGMTMIB_Clmgmtlicensablefeaturetable) GetEntityData() *types.CommonEntityData {
    clmgmtlicensablefeaturetable.EntityData.YFilter = clmgmtlicensablefeaturetable.YFilter
    clmgmtlicensablefeaturetable.EntityData.YangName = "clmgmtLicensableFeatureTable"
    clmgmtlicensablefeaturetable.EntityData.BundleName = "cisco_ios_xe"
    clmgmtlicensablefeaturetable.EntityData.ParentYangName = "CISCO-LICENSE-MGMT-MIB"
    clmgmtlicensablefeaturetable.EntityData.SegmentPath = "clmgmtLicensableFeatureTable"
    clmgmtlicensablefeaturetable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtlicensablefeaturetable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtlicensablefeaturetable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtlicensablefeaturetable.EntityData.Children = make(map[string]types.YChild)
    clmgmtlicensablefeaturetable.EntityData.Children["clmgmtLicensableFeatureEntry"] = types.YChild{"Clmgmtlicensablefeatureentry", nil}
    for i := range clmgmtlicensablefeaturetable.Clmgmtlicensablefeatureentry {
        clmgmtlicensablefeaturetable.EntityData.Children[types.GetSegmentPath(&clmgmtlicensablefeaturetable.Clmgmtlicensablefeatureentry[i])] = types.YChild{"Clmgmtlicensablefeatureentry", &clmgmtlicensablefeaturetable.Clmgmtlicensablefeatureentry[i]}
    }
    clmgmtlicensablefeaturetable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clmgmtlicensablefeaturetable.EntityData)
}

// CISCOLICENSEMGMTMIB_Clmgmtlicensablefeaturetable_Clmgmtlicensablefeatureentry
// An entry in clmgmtLicensableFeatureTable. Each entry represents
// a licensable feature.
type CISCOLICENSEMGMTMIB_Clmgmtlicensablefeaturetable_Clmgmtlicensablefeatureentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_Entphysicaltable_Entphysicalentry_Entphysicalindex
    Entphysicalindex interface{}

    // This attribute is a key. This object uniquely identifies a licensable
    // feature in the device. The type is interface{} with range: 1..4294967295.
    Clmgmtfeatureindex interface{}

    // This object indicates the name of the licensable feature in the device.
    // Examples of feature names are: 'IPBASE', 'ADVIPSERVICE'. The type is string
    // with length: 0..128.
    Clmgmtfeaturename interface{}

    // This object indicates the version of the licensable feature in the device.
    // Examples of feature versions are: '1.0' or '2.0'. The type is string with
    // length: 0..32.
    Clmgmtfeatureversion interface{}

    // This object indicates the time period remaining before the feature's
    // license expires or transitions. This object is applicable only if
    // clmgmtLicenseType of the license used by this feature is demo(1), or
    // extension(2) or gracePeriod(3) or evalRightToUse(8).  The object will
    // contain 0 if other types of license is used or if the feature does not use
    // any license. If the feature is using multiple licenses, this period will
    // represent the cumulative period remaining from all the licenses used by
    // this feature. The type is interface{} with range: 0..4294967295. Units are
    // seconds.
    Clmgmtfeaturevalidityperiodremaining interface{}

    // This object represents the entity that is being counted by this feature.
    // Examples of entities are IP Phones, number of sessions etc. This object is
    // only applicable for features that use counting licenses. For other
    // features, this object will return empty string. The type is string with
    // length: 0..128.
    Clmgmtfeaturewhatiscounted interface{}

    // This object indicates the license start date of the feature. This object is
    // applicable if at least one of the licenses used for this feature has a
    // valid start date. The start date will be the earliest of the valid start
    // dates of all the licenses used for this feature. If none of the licenses
    // used for this feature have a valid start date then this object will contain
    // an octet string of length 0. The type is string.
    Clmgmtfeaturestartdate interface{}

    // This object indicates the license end date of the feature. This object is
    // applicable if at least one of the licenses used for this feature has a
    // valid end date. The end date will be the latest of the valid end dates of
    // all the licenses used for this feature. If none of the licenses used for
    // this feature have a valid end date then this object will contain an octet
    // string of length 0. The type is string.
    Clmgmtfeatureenddate interface{}

    // This object indicates the license period used for the feature. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    Clmgmtfeatureperiodused interface{}
}

func (clmgmtlicensablefeatureentry *CISCOLICENSEMGMTMIB_Clmgmtlicensablefeaturetable_Clmgmtlicensablefeatureentry) GetEntityData() *types.CommonEntityData {
    clmgmtlicensablefeatureentry.EntityData.YFilter = clmgmtlicensablefeatureentry.YFilter
    clmgmtlicensablefeatureentry.EntityData.YangName = "clmgmtLicensableFeatureEntry"
    clmgmtlicensablefeatureentry.EntityData.BundleName = "cisco_ios_xe"
    clmgmtlicensablefeatureentry.EntityData.ParentYangName = "clmgmtLicensableFeatureTable"
    clmgmtlicensablefeatureentry.EntityData.SegmentPath = "clmgmtLicensableFeatureEntry" + "[entPhysicalIndex='" + fmt.Sprintf("%v", clmgmtlicensablefeatureentry.Entphysicalindex) + "']" + "[clmgmtFeatureIndex='" + fmt.Sprintf("%v", clmgmtlicensablefeatureentry.Clmgmtfeatureindex) + "']"
    clmgmtlicensablefeatureentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtlicensablefeatureentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtlicensablefeatureentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtlicensablefeatureentry.EntityData.Children = make(map[string]types.YChild)
    clmgmtlicensablefeatureentry.EntityData.Leafs = make(map[string]types.YLeaf)
    clmgmtlicensablefeatureentry.EntityData.Leafs["entPhysicalIndex"] = types.YLeaf{"Entphysicalindex", clmgmtlicensablefeatureentry.Entphysicalindex}
    clmgmtlicensablefeatureentry.EntityData.Leafs["clmgmtFeatureIndex"] = types.YLeaf{"Clmgmtfeatureindex", clmgmtlicensablefeatureentry.Clmgmtfeatureindex}
    clmgmtlicensablefeatureentry.EntityData.Leafs["clmgmtFeatureName"] = types.YLeaf{"Clmgmtfeaturename", clmgmtlicensablefeatureentry.Clmgmtfeaturename}
    clmgmtlicensablefeatureentry.EntityData.Leafs["clmgmtFeatureVersion"] = types.YLeaf{"Clmgmtfeatureversion", clmgmtlicensablefeatureentry.Clmgmtfeatureversion}
    clmgmtlicensablefeatureentry.EntityData.Leafs["clmgmtFeatureValidityPeriodRemaining"] = types.YLeaf{"Clmgmtfeaturevalidityperiodremaining", clmgmtlicensablefeatureentry.Clmgmtfeaturevalidityperiodremaining}
    clmgmtlicensablefeatureentry.EntityData.Leafs["clmgmtFeatureWhatIsCounted"] = types.YLeaf{"Clmgmtfeaturewhatiscounted", clmgmtlicensablefeatureentry.Clmgmtfeaturewhatiscounted}
    clmgmtlicensablefeatureentry.EntityData.Leafs["clmgmtFeatureStartDate"] = types.YLeaf{"Clmgmtfeaturestartdate", clmgmtlicensablefeatureentry.Clmgmtfeaturestartdate}
    clmgmtlicensablefeatureentry.EntityData.Leafs["clmgmtFeatureEndDate"] = types.YLeaf{"Clmgmtfeatureenddate", clmgmtlicensablefeatureentry.Clmgmtfeatureenddate}
    clmgmtlicensablefeatureentry.EntityData.Leafs["clmgmtFeaturePeriodUsed"] = types.YLeaf{"Clmgmtfeatureperiodused", clmgmtlicensablefeatureentry.Clmgmtfeatureperiodused}
    return &(clmgmtlicensablefeatureentry.EntityData)
}

// CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable
// A table for triggering device credentials export action.
// Management application must create this entry to trigger the
// export of device credentials from the device to a file.
// 
// Once the request completes, the management application should
// retrieve the values of the objects of interest, and then
// delete the entry.  In order to prevent old entries from
// clogging the table, entries will be aged out, but an entry
// will never be deleted within 5 minutes of completion.
type CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry for each device credential export action that is being executed or
    // was executed recently. The management application triggers the export by
    // creating an entry in this table. This can be done in the following 2
    // methods  1. CREATE-AND-GO method     Management application sets
    // clmgmtDevCredExportActionStatus     to createAndGo(4) and all other
    // required objects to valid     values in a single SNMP SET request. If all
    // the values     are valid, the device creates the entry and executes the    
    // action. If the SET request fails, the entry will not be     created. 2.
    // CREATE-AND-WAIT method     Management application sets
    // clmgmtDevCredExportActionStatus to     createAndWait(5) to create an entry.
    // Management application     can set all other required objects to valid    
    // values in more than one SNMP SET request. If SET request     for any of the
    // objects fails, management application can set     just only that object.
    // Once all the required objects     are set to valid values, management
    // application triggers action     execution by setting
    // clmgmtDevCredExportActionStatus to     active(1).  To stop the action from
    // being executed, the management application can delete the entry by setting
    // clmgmtDevCredExportActionStatus to destroy(6) when
    // clmgmtDevCredCommandState is pending(2).  The status of action execution
    // can be known by querying clmgmtDevCredCommandState. If the action is still
    // in pending(2) or inProgress(3), the management application need to check
    // back again after few seconds. Once the action completes and if status of
    // the action is failed(6), the reason for failure can be retrieved from
    // clmgmtDevCredCommandFailCause.  Entry can be deleted except when
    // clmgmtLicenseAction is set to inProgress(3). All entries in this table are
    // volatile and are cleared on agent reset. The type is slice of
    // CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable_Clmgmtdevcredexportactionentry.
    Clmgmtdevcredexportactionentry []CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable_Clmgmtdevcredexportactionentry
}

func (clmgmtdevcredexportactiontable *CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable) GetEntityData() *types.CommonEntityData {
    clmgmtdevcredexportactiontable.EntityData.YFilter = clmgmtdevcredexportactiontable.YFilter
    clmgmtdevcredexportactiontable.EntityData.YangName = "clmgmtDevCredExportActionTable"
    clmgmtdevcredexportactiontable.EntityData.BundleName = "cisco_ios_xe"
    clmgmtdevcredexportactiontable.EntityData.ParentYangName = "CISCO-LICENSE-MGMT-MIB"
    clmgmtdevcredexportactiontable.EntityData.SegmentPath = "clmgmtDevCredExportActionTable"
    clmgmtdevcredexportactiontable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtdevcredexportactiontable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtdevcredexportactiontable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtdevcredexportactiontable.EntityData.Children = make(map[string]types.YChild)
    clmgmtdevcredexportactiontable.EntityData.Children["clmgmtDevCredExportActionEntry"] = types.YChild{"Clmgmtdevcredexportactionentry", nil}
    for i := range clmgmtdevcredexportactiontable.Clmgmtdevcredexportactionentry {
        clmgmtdevcredexportactiontable.EntityData.Children[types.GetSegmentPath(&clmgmtdevcredexportactiontable.Clmgmtdevcredexportactionentry[i])] = types.YChild{"Clmgmtdevcredexportactionentry", &clmgmtdevcredexportactiontable.Clmgmtdevcredexportactionentry[i]}
    }
    clmgmtdevcredexportactiontable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(clmgmtdevcredexportactiontable.EntityData)
}

// CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable_Clmgmtdevcredexportactionentry
// An entry for each device credential export action that
// is being executed or was executed recently. The management
// application triggers the export by creating an entry in this
// table. This can be done in the following 2 methods
// 
// 1. CREATE-AND-GO method
//     Management application sets clmgmtDevCredExportActionStatus
//     to createAndGo(4) and all other required objects to valid
//     values in a single SNMP SET request. If all the values
//     are valid, the device creates the entry and executes the
//     action. If the SET request fails, the entry will not be
//     created.
// 2. CREATE-AND-WAIT method
//     Management application sets clmgmtDevCredExportActionStatus to
//     createAndWait(5) to create an entry. Management application
//     can set all other required objects to valid
//     values in more than one SNMP SET request. If SET request
//     for any of the objects fails, management application can set
//     just only that object. Once all the required objects
//     are set to valid values, management application triggers action
//     execution by setting clmgmtDevCredExportActionStatus to
//     active(1).
// 
// To stop the action from being executed, the management application
// can delete the entry by setting clmgmtDevCredExportActionStatus
// to destroy(6) when clmgmtDevCredCommandState is pending(2).
// 
// The status of action execution can be known by querying
// clmgmtDevCredCommandState. If the action is still in
// pending(2) or inProgress(3), the management application need to
// check back again after few seconds. Once the action completes
// and if status of the action is failed(6), the reason for
// failure can be retrieved from clmgmtDevCredCommandFailCause.
// 
// Entry can be deleted except when clmgmtLicenseAction is set
// to inProgress(3). All entries in this table are volatile
// and are cleared on agent reset.
type CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable_Clmgmtdevcredexportactionentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. This object uniquely identifies a row in
    // clmgmtDevCredExportActionTable. The management application chooses this
    // value by reading clmgmtNextFreeDevCredExportActionIndex while creating an
    // entry in this table. If an entry already exists with this index, the
    // creation of the entry will not continue and error will be returned. The
    // management application should read the value of
    // clmgmtNextFreeDevCredExportActionIndex again and retry with the new value
    // for this object. The type is interface{} with range: 1..4294967295.
    Clmgmtdevcredexportactionindex interface{}

    // This object represents the entPhysicalIndex of the device for which the
    // device credentials are being retrieved. This object is mainly used in
    // devices where one device is acting as a master and rest of the devices as
    // slaves. The master device is responsible for SNMP communication with the
    // manager. Examples include stackable switches, devices with router processor
    // and line cards.  Note: This object need not be set if it is a stand alone
    // device. The type is interface{} with range: 0..2147483647.
    Clmgmtdevcredentphysicalindex interface{}

    // This object indicates the transfer protocol to be used when copying files
    // as specified in the following objects. 1. clmgmtDevCredExportFile . The
    // type is ClmgmtLicenseTransferProtocol.
    Clmgmtdevcredtransferprotocol interface{}

    // This object indicates the transport type of the address contained in
    // clmgmtDevCredServerAddress object. This must be set when
    // clmgmtDevCredTransferProtocol is not none(1) or local(2). The type is
    // InetAddressType.
    Clmgmtdevcredserveraddresstype interface{}

    // This object indicates the the ip address of the server from which the files
    // must be read or written to if  clmgmtDevCredTransferProtocol is not none(1)
    // or local(2).  All bits as 0s or 1s for clmgmtDevCredServerAddress are not
    // allowed.  The format of this address depends on the value of the
    // clmgmtDevCredServerAddressType object. The type is string with length:
    // 0..255.
    Clmgmtdevcredserveraddress interface{}

    // This object indicates the remote user name for accessing files via ftp,
    // rcp, sftp or scp protocols. This object must be set when the
    // clmgmtDevCredTransferProtocol is ftp(4), rcp(5), scp(7) or sftp(8). If
    // clmgmtDevCredTransferProtocol is rcp(5), the remote username is sent as the
    // server username in an rcp command request sent by the system to a remote
    // rcp server. The type is string with length: 0..96.
    Clmgmtdevcredserverusername interface{}

    // This object indicates the password used by ftp, sftp or scp for copying a
    // file to/from an ftp/sftp/scp server.  This object must be set when the
    // clmgmtDevCredTransferProtocol is ftp(4) or scp(7) or sftp(8). Reading it
    // returns a zero-length string for  security reasons. The type is string with
    // length: 0..96.
    Clmgmtdevcredserverpassword interface{}

    // This object represents file where device credentials needs to be exported
    // to. The type is string with length: 0..255.
    Clmgmtdevcredexportfile interface{}

    // This object indicates the the command to be executed.  Command             
    // Remarks -------                          ------- noOp(1)                   
    // No operation will be                                 performed. 
    // getDeviceCredentials(2)         Exports device credentials. The type is
    // Clmgmtdevcredcommand.
    Clmgmtdevcredcommand interface{}

    // This object indicates the state of the action that is executed as a result
    // of setting clmgmtDevCredRowStatus to active(1). The type is
    // ClmgmtLicenseActionState.
    Clmgmtdevcredcommandstate interface{}

    // This object indicates the the reason for device credentials export
    // operation failure.  The value of this object is valid only when
    // clmgmtDevCredCommandState is failed(6).  none(1)         - action execution
    // has not started yet.                   If the action is completed and the  
    // action is successful, then also                   none(1) is returned to
    // indicate that                   there are no errors. unknownError(2) -
    // reason for failure is unknown,                   operation failed, no
    // operation is                   performed transferProtocolNotSupported(3) -
    // clmgmtDevCredTransferProtocol                                   given is
    // not supported. fileServerNotReachable(4)       - file server is not
    // reachable. unrecognizedEntPhysicalIndex(5) - entPhysicalIndex is not       
    // valid invalidFile(6)  - The target file specified is not valid. The type is
    // Clmgmtdevcredcommandfailcause.
    Clmgmtdevcredcommandfailcause interface{}

    // This object indicates the storage type for this conceptual row. Conceptual
    // rows having the value 'permanent' need not allow write-access to any
    // columnar objects in the row. The type is StorageType.
    Clmgmtdevcredstoragetype interface{}

    // This object indicates the the status of this table entry. Once the entry
    // status is set to active(1), the associated entry cannot be modified until
    // the action completes (clmgmtDevCredCommandStatus is set to a value other
    // than inProgress(3)). Once the action completes the only operation possible
    // after this is to delete the row.  clmgmtDevCredExportFile is a mandatory
    // object to be set when creating this entry. The type is RowStatus.
    Clmgmtdevcredrowstatus interface{}
}

func (clmgmtdevcredexportactionentry *CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable_Clmgmtdevcredexportactionentry) GetEntityData() *types.CommonEntityData {
    clmgmtdevcredexportactionentry.EntityData.YFilter = clmgmtdevcredexportactionentry.YFilter
    clmgmtdevcredexportactionentry.EntityData.YangName = "clmgmtDevCredExportActionEntry"
    clmgmtdevcredexportactionentry.EntityData.BundleName = "cisco_ios_xe"
    clmgmtdevcredexportactionentry.EntityData.ParentYangName = "clmgmtDevCredExportActionTable"
    clmgmtdevcredexportactionentry.EntityData.SegmentPath = "clmgmtDevCredExportActionEntry" + "[clmgmtDevCredExportActionIndex='" + fmt.Sprintf("%v", clmgmtdevcredexportactionentry.Clmgmtdevcredexportactionindex) + "']"
    clmgmtdevcredexportactionentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtdevcredexportactionentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtdevcredexportactionentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtdevcredexportactionentry.EntityData.Children = make(map[string]types.YChild)
    clmgmtdevcredexportactionentry.EntityData.Leafs = make(map[string]types.YLeaf)
    clmgmtdevcredexportactionentry.EntityData.Leafs["clmgmtDevCredExportActionIndex"] = types.YLeaf{"Clmgmtdevcredexportactionindex", clmgmtdevcredexportactionentry.Clmgmtdevcredexportactionindex}
    clmgmtdevcredexportactionentry.EntityData.Leafs["clmgmtDevCredEntPhysicalIndex"] = types.YLeaf{"Clmgmtdevcredentphysicalindex", clmgmtdevcredexportactionentry.Clmgmtdevcredentphysicalindex}
    clmgmtdevcredexportactionentry.EntityData.Leafs["clmgmtDevCredTransferProtocol"] = types.YLeaf{"Clmgmtdevcredtransferprotocol", clmgmtdevcredexportactionentry.Clmgmtdevcredtransferprotocol}
    clmgmtdevcredexportactionentry.EntityData.Leafs["clmgmtDevCredServerAddressType"] = types.YLeaf{"Clmgmtdevcredserveraddresstype", clmgmtdevcredexportactionentry.Clmgmtdevcredserveraddresstype}
    clmgmtdevcredexportactionentry.EntityData.Leafs["clmgmtDevCredServerAddress"] = types.YLeaf{"Clmgmtdevcredserveraddress", clmgmtdevcredexportactionentry.Clmgmtdevcredserveraddress}
    clmgmtdevcredexportactionentry.EntityData.Leafs["clmgmtDevCredServerUsername"] = types.YLeaf{"Clmgmtdevcredserverusername", clmgmtdevcredexportactionentry.Clmgmtdevcredserverusername}
    clmgmtdevcredexportactionentry.EntityData.Leafs["clmgmtDevCredServerPassword"] = types.YLeaf{"Clmgmtdevcredserverpassword", clmgmtdevcredexportactionentry.Clmgmtdevcredserverpassword}
    clmgmtdevcredexportactionentry.EntityData.Leafs["clmgmtDevCredExportFile"] = types.YLeaf{"Clmgmtdevcredexportfile", clmgmtdevcredexportactionentry.Clmgmtdevcredexportfile}
    clmgmtdevcredexportactionentry.EntityData.Leafs["clmgmtDevCredCommand"] = types.YLeaf{"Clmgmtdevcredcommand", clmgmtdevcredexportactionentry.Clmgmtdevcredcommand}
    clmgmtdevcredexportactionentry.EntityData.Leafs["clmgmtDevCredCommandState"] = types.YLeaf{"Clmgmtdevcredcommandstate", clmgmtdevcredexportactionentry.Clmgmtdevcredcommandstate}
    clmgmtdevcredexportactionentry.EntityData.Leafs["clmgmtDevCredCommandFailCause"] = types.YLeaf{"Clmgmtdevcredcommandfailcause", clmgmtdevcredexportactionentry.Clmgmtdevcredcommandfailcause}
    clmgmtdevcredexportactionentry.EntityData.Leafs["clmgmtDevCredStorageType"] = types.YLeaf{"Clmgmtdevcredstoragetype", clmgmtdevcredexportactionentry.Clmgmtdevcredstoragetype}
    clmgmtdevcredexportactionentry.EntityData.Leafs["clmgmtDevCredRowStatus"] = types.YLeaf{"Clmgmtdevcredrowstatus", clmgmtdevcredexportactionentry.Clmgmtdevcredrowstatus}
    return &(clmgmtdevcredexportactionentry.EntityData)
}

// CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable_Clmgmtdevcredexportactionentry_Clmgmtdevcredcommand represents getDeviceCredentials(2)         Exports device credentials
type CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable_Clmgmtdevcredexportactionentry_Clmgmtdevcredcommand string

const (
    CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable_Clmgmtdevcredexportactionentry_Clmgmtdevcredcommand_noOp CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable_Clmgmtdevcredexportactionentry_Clmgmtdevcredcommand = "noOp"

    CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable_Clmgmtdevcredexportactionentry_Clmgmtdevcredcommand_getDeviceCredentials CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable_Clmgmtdevcredexportactionentry_Clmgmtdevcredcommand = "getDeviceCredentials"
)

// CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable_Clmgmtdevcredexportactionentry_Clmgmtdevcredcommandfailcause represents invalidFile(6)  - The target file specified is not valid.
type CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable_Clmgmtdevcredexportactionentry_Clmgmtdevcredcommandfailcause string

const (
    CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable_Clmgmtdevcredexportactionentry_Clmgmtdevcredcommandfailcause_none CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable_Clmgmtdevcredexportactionentry_Clmgmtdevcredcommandfailcause = "none"

    CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable_Clmgmtdevcredexportactionentry_Clmgmtdevcredcommandfailcause_unknownError CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable_Clmgmtdevcredexportactionentry_Clmgmtdevcredcommandfailcause = "unknownError"

    CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable_Clmgmtdevcredexportactionentry_Clmgmtdevcredcommandfailcause_transferProtocolNotSupported CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable_Clmgmtdevcredexportactionentry_Clmgmtdevcredcommandfailcause = "transferProtocolNotSupported"

    CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable_Clmgmtdevcredexportactionentry_Clmgmtdevcredcommandfailcause_fileServerNotReachable CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable_Clmgmtdevcredexportactionentry_Clmgmtdevcredcommandfailcause = "fileServerNotReachable"

    CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable_Clmgmtdevcredexportactionentry_Clmgmtdevcredcommandfailcause_unrecognizedEntPhysicalIndex CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable_Clmgmtdevcredexportactionentry_Clmgmtdevcredcommandfailcause = "unrecognizedEntPhysicalIndex"

    CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable_Clmgmtdevcredexportactionentry_Clmgmtdevcredcommandfailcause_invalidFile CISCOLICENSEMGMTMIB_Clmgmtdevcredexportactiontable_Clmgmtdevcredexportactionentry_Clmgmtdevcredcommandfailcause = "invalidFile"
)

