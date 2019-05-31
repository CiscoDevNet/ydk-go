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

    
    ClmgmtLicenseConfiguration CISCOLICENSEMGMTMIB_ClmgmtLicenseConfiguration

    
    ClmgmtLicenseDeviceInformation CISCOLICENSEMGMTMIB_ClmgmtLicenseDeviceInformation

    
    ClmgmtLicenseNotifObjects CISCOLICENSEMGMTMIB_ClmgmtLicenseNotifObjects

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
    ClmgmtLicenseActionTable CISCOLICENSEMGMTMIB_ClmgmtLicenseActionTable

    // This table contains results of license action if the license action
    // involves multiple licenses. Entries in this table are not created for
    // actions where there is only license that is subject of the action. For
    // example, if there are 3 licenses in a license file when executing license
    // install action, 3 entries will be created in this table, one for each
    // license.
    ClmgmtLicenseActionResultTable CISCOLICENSEMGMTMIB_ClmgmtLicenseActionResultTable

    // This table contains information about all the license stores allocated on
    // the device.
    ClmgmtLicenseStoreInfoTable CISCOLICENSEMGMTMIB_ClmgmtLicenseStoreInfoTable

    // This table contains objects that provide licensing related information at
    // the device level. Entries will exist only for entities that support
    // licensing. For example, if it is a stand alone device and supports
    // licensing, then there will be only one entry in this table. If it is
    // stackable switch then there will be multiple entries with one entry for
    // each device in the stack.
    ClmgmtLicenseDeviceInfoTable CISCOLICENSEMGMTMIB_ClmgmtLicenseDeviceInfoTable

    // This table contains information about all the licenses installed on the
    // device.
    ClmgmtLicenseInfoTable CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable

    // This table contains list of licensable features in the image. All the
    // licensable features will have an entry each in this table irrespective of
    // whether they are using any licenses currently. Entries in this table are
    // created by the agent one for each licensable feature in the image. These
    // entries remain in the table permanently and can not be deleted. Management
    // application can not create or delete entries from this table.
    ClmgmtLicensableFeatureTable CISCOLICENSEMGMTMIB_ClmgmtLicensableFeatureTable

    // A table for triggering device credentials export action. Management
    // application must create this entry to trigger the export of device
    // credentials from the device to a file.  Once the request completes, the
    // management application should retrieve the values of the objects of
    // interest, and then delete the entry.  In order to prevent old entries from
    // clogging the table, entries will be aged out, but an entry will never be
    // deleted within 5 minutes of completion.
    ClmgmtDevCredExportActionTable CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable
}

func (cISCOLICENSEMGMTMIB *CISCOLICENSEMGMTMIB) GetEntityData() *types.CommonEntityData {
    cISCOLICENSEMGMTMIB.EntityData.YFilter = cISCOLICENSEMGMTMIB.YFilter
    cISCOLICENSEMGMTMIB.EntityData.YangName = "CISCO-LICENSE-MGMT-MIB"
    cISCOLICENSEMGMTMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOLICENSEMGMTMIB.EntityData.ParentYangName = "CISCO-LICENSE-MGMT-MIB"
    cISCOLICENSEMGMTMIB.EntityData.SegmentPath = "CISCO-LICENSE-MGMT-MIB:CISCO-LICENSE-MGMT-MIB"
    cISCOLICENSEMGMTMIB.EntityData.AbsolutePath = cISCOLICENSEMGMTMIB.EntityData.SegmentPath
    cISCOLICENSEMGMTMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOLICENSEMGMTMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOLICENSEMGMTMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOLICENSEMGMTMIB.EntityData.Children = types.NewOrderedMap()
    cISCOLICENSEMGMTMIB.EntityData.Children.Append("clmgmtLicenseConfiguration", types.YChild{"ClmgmtLicenseConfiguration", &cISCOLICENSEMGMTMIB.ClmgmtLicenseConfiguration})
    cISCOLICENSEMGMTMIB.EntityData.Children.Append("clmgmtLicenseDeviceInformation", types.YChild{"ClmgmtLicenseDeviceInformation", &cISCOLICENSEMGMTMIB.ClmgmtLicenseDeviceInformation})
    cISCOLICENSEMGMTMIB.EntityData.Children.Append("clmgmtLicenseNotifObjects", types.YChild{"ClmgmtLicenseNotifObjects", &cISCOLICENSEMGMTMIB.ClmgmtLicenseNotifObjects})
    cISCOLICENSEMGMTMIB.EntityData.Children.Append("clmgmtLicenseActionTable", types.YChild{"ClmgmtLicenseActionTable", &cISCOLICENSEMGMTMIB.ClmgmtLicenseActionTable})
    cISCOLICENSEMGMTMIB.EntityData.Children.Append("clmgmtLicenseActionResultTable", types.YChild{"ClmgmtLicenseActionResultTable", &cISCOLICENSEMGMTMIB.ClmgmtLicenseActionResultTable})
    cISCOLICENSEMGMTMIB.EntityData.Children.Append("clmgmtLicenseStoreInfoTable", types.YChild{"ClmgmtLicenseStoreInfoTable", &cISCOLICENSEMGMTMIB.ClmgmtLicenseStoreInfoTable})
    cISCOLICENSEMGMTMIB.EntityData.Children.Append("clmgmtLicenseDeviceInfoTable", types.YChild{"ClmgmtLicenseDeviceInfoTable", &cISCOLICENSEMGMTMIB.ClmgmtLicenseDeviceInfoTable})
    cISCOLICENSEMGMTMIB.EntityData.Children.Append("clmgmtLicenseInfoTable", types.YChild{"ClmgmtLicenseInfoTable", &cISCOLICENSEMGMTMIB.ClmgmtLicenseInfoTable})
    cISCOLICENSEMGMTMIB.EntityData.Children.Append("clmgmtLicensableFeatureTable", types.YChild{"ClmgmtLicensableFeatureTable", &cISCOLICENSEMGMTMIB.ClmgmtLicensableFeatureTable})
    cISCOLICENSEMGMTMIB.EntityData.Children.Append("clmgmtDevCredExportActionTable", types.YChild{"ClmgmtDevCredExportActionTable", &cISCOLICENSEMGMTMIB.ClmgmtDevCredExportActionTable})
    cISCOLICENSEMGMTMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOLICENSEMGMTMIB.EntityData.YListKeys = []string {}

    return &(cISCOLICENSEMGMTMIB.EntityData)
}

// CISCOLICENSEMGMTMIB_ClmgmtLicenseConfiguration
type CISCOLICENSEMGMTMIB_ClmgmtLicenseConfiguration struct {
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
    ClmgmtNextFreeLicenseActionIndex interface{}
}

func (clmgmtLicenseConfiguration *CISCOLICENSEMGMTMIB_ClmgmtLicenseConfiguration) GetEntityData() *types.CommonEntityData {
    clmgmtLicenseConfiguration.EntityData.YFilter = clmgmtLicenseConfiguration.YFilter
    clmgmtLicenseConfiguration.EntityData.YangName = "clmgmtLicenseConfiguration"
    clmgmtLicenseConfiguration.EntityData.BundleName = "cisco_ios_xe"
    clmgmtLicenseConfiguration.EntityData.ParentYangName = "CISCO-LICENSE-MGMT-MIB"
    clmgmtLicenseConfiguration.EntityData.SegmentPath = "clmgmtLicenseConfiguration"
    clmgmtLicenseConfiguration.EntityData.AbsolutePath = "CISCO-LICENSE-MGMT-MIB:CISCO-LICENSE-MGMT-MIB/" + clmgmtLicenseConfiguration.EntityData.SegmentPath
    clmgmtLicenseConfiguration.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtLicenseConfiguration.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtLicenseConfiguration.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtLicenseConfiguration.EntityData.Children = types.NewOrderedMap()
    clmgmtLicenseConfiguration.EntityData.Leafs = types.NewOrderedMap()
    clmgmtLicenseConfiguration.EntityData.Leafs.Append("clmgmtNextFreeLicenseActionIndex", types.YLeaf{"ClmgmtNextFreeLicenseActionIndex", clmgmtLicenseConfiguration.ClmgmtNextFreeLicenseActionIndex})

    clmgmtLicenseConfiguration.EntityData.YListKeys = []string {}

    return &(clmgmtLicenseConfiguration.EntityData)
}

// CISCOLICENSEMGMTMIB_ClmgmtLicenseDeviceInformation
type CISCOLICENSEMGMTMIB_ClmgmtLicenseDeviceInformation struct {
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
    ClmgmtNextFreeDevCredExportActionIndex interface{}
}

func (clmgmtLicenseDeviceInformation *CISCOLICENSEMGMTMIB_ClmgmtLicenseDeviceInformation) GetEntityData() *types.CommonEntityData {
    clmgmtLicenseDeviceInformation.EntityData.YFilter = clmgmtLicenseDeviceInformation.YFilter
    clmgmtLicenseDeviceInformation.EntityData.YangName = "clmgmtLicenseDeviceInformation"
    clmgmtLicenseDeviceInformation.EntityData.BundleName = "cisco_ios_xe"
    clmgmtLicenseDeviceInformation.EntityData.ParentYangName = "CISCO-LICENSE-MGMT-MIB"
    clmgmtLicenseDeviceInformation.EntityData.SegmentPath = "clmgmtLicenseDeviceInformation"
    clmgmtLicenseDeviceInformation.EntityData.AbsolutePath = "CISCO-LICENSE-MGMT-MIB:CISCO-LICENSE-MGMT-MIB/" + clmgmtLicenseDeviceInformation.EntityData.SegmentPath
    clmgmtLicenseDeviceInformation.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtLicenseDeviceInformation.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtLicenseDeviceInformation.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtLicenseDeviceInformation.EntityData.Children = types.NewOrderedMap()
    clmgmtLicenseDeviceInformation.EntityData.Leafs = types.NewOrderedMap()
    clmgmtLicenseDeviceInformation.EntityData.Leafs.Append("clmgmtNextFreeDevCredExportActionIndex", types.YLeaf{"ClmgmtNextFreeDevCredExportActionIndex", clmgmtLicenseDeviceInformation.ClmgmtNextFreeDevCredExportActionIndex})

    clmgmtLicenseDeviceInformation.EntityData.YListKeys = []string {}

    return &(clmgmtLicenseDeviceInformation.EntityData)
}

// CISCOLICENSEMGMTMIB_ClmgmtLicenseNotifObjects
type CISCOLICENSEMGMTMIB_ClmgmtLicenseNotifObjects struct {
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
    ClmgmtLicenseUsageNotifEnable interface{}

    // This object indicates whether the device should generate notifications
    // related to license deployment. This object enables/disables sending
    // following notifications:     clmgmtLicenseInstalled    
    // clmgmtLicenseCleared     clmgmtLicenseRevoked    
    // clmgmtLicenseEULAAccepted. The type is bool.
    ClmgmtLicenseDeploymentNotifEnable interface{}

    // This object indicates whether the device should generate notifications
    // related to error conditions in enforcing licensing. This object
    // enables/disables sending following notifications:    
    // clmgmtLicenseNotEnforced. The type is ClmgmtLicenseErrorNotifEnable.
    ClmgmtLicenseErrorNotifEnable interface{}
}

func (clmgmtLicenseNotifObjects *CISCOLICENSEMGMTMIB_ClmgmtLicenseNotifObjects) GetEntityData() *types.CommonEntityData {
    clmgmtLicenseNotifObjects.EntityData.YFilter = clmgmtLicenseNotifObjects.YFilter
    clmgmtLicenseNotifObjects.EntityData.YangName = "clmgmtLicenseNotifObjects"
    clmgmtLicenseNotifObjects.EntityData.BundleName = "cisco_ios_xe"
    clmgmtLicenseNotifObjects.EntityData.ParentYangName = "CISCO-LICENSE-MGMT-MIB"
    clmgmtLicenseNotifObjects.EntityData.SegmentPath = "clmgmtLicenseNotifObjects"
    clmgmtLicenseNotifObjects.EntityData.AbsolutePath = "CISCO-LICENSE-MGMT-MIB:CISCO-LICENSE-MGMT-MIB/" + clmgmtLicenseNotifObjects.EntityData.SegmentPath
    clmgmtLicenseNotifObjects.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtLicenseNotifObjects.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtLicenseNotifObjects.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtLicenseNotifObjects.EntityData.Children = types.NewOrderedMap()
    clmgmtLicenseNotifObjects.EntityData.Leafs = types.NewOrderedMap()
    clmgmtLicenseNotifObjects.EntityData.Leafs.Append("clmgmtLicenseUsageNotifEnable", types.YLeaf{"ClmgmtLicenseUsageNotifEnable", clmgmtLicenseNotifObjects.ClmgmtLicenseUsageNotifEnable})
    clmgmtLicenseNotifObjects.EntityData.Leafs.Append("clmgmtLicenseDeploymentNotifEnable", types.YLeaf{"ClmgmtLicenseDeploymentNotifEnable", clmgmtLicenseNotifObjects.ClmgmtLicenseDeploymentNotifEnable})
    clmgmtLicenseNotifObjects.EntityData.Leafs.Append("clmgmtLicenseErrorNotifEnable", types.YLeaf{"ClmgmtLicenseErrorNotifEnable", clmgmtLicenseNotifObjects.ClmgmtLicenseErrorNotifEnable})

    clmgmtLicenseNotifObjects.EntityData.YListKeys = []string {}

    return &(clmgmtLicenseNotifObjects.EntityData)
}

// CISCOLICENSEMGMTMIB_ClmgmtLicenseNotifObjects_ClmgmtLicenseErrorNotifEnable represents     clmgmtLicenseNotEnforced
type CISCOLICENSEMGMTMIB_ClmgmtLicenseNotifObjects_ClmgmtLicenseErrorNotifEnable string

const (
    CISCOLICENSEMGMTMIB_ClmgmtLicenseNotifObjects_ClmgmtLicenseErrorNotifEnable_other CISCOLICENSEMGMTMIB_ClmgmtLicenseNotifObjects_ClmgmtLicenseErrorNotifEnable = "other"

    CISCOLICENSEMGMTMIB_ClmgmtLicenseNotifObjects_ClmgmtLicenseErrorNotifEnable_true_ CISCOLICENSEMGMTMIB_ClmgmtLicenseNotifObjects_ClmgmtLicenseErrorNotifEnable = "true"

    CISCOLICENSEMGMTMIB_ClmgmtLicenseNotifObjects_ClmgmtLicenseErrorNotifEnable_false_ CISCOLICENSEMGMTMIB_ClmgmtLicenseNotifObjects_ClmgmtLicenseErrorNotifEnable = "false"
)

// CISCOLICENSEMGMTMIB_ClmgmtLicenseActionTable
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
type CISCOLICENSEMGMTMIB_ClmgmtLicenseActionTable struct {
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
    // CISCOLICENSEMGMTMIB_ClmgmtLicenseActionTable_ClmgmtLicenseActionEntry.
    ClmgmtLicenseActionEntry []*CISCOLICENSEMGMTMIB_ClmgmtLicenseActionTable_ClmgmtLicenseActionEntry
}

func (clmgmtLicenseActionTable *CISCOLICENSEMGMTMIB_ClmgmtLicenseActionTable) GetEntityData() *types.CommonEntityData {
    clmgmtLicenseActionTable.EntityData.YFilter = clmgmtLicenseActionTable.YFilter
    clmgmtLicenseActionTable.EntityData.YangName = "clmgmtLicenseActionTable"
    clmgmtLicenseActionTable.EntityData.BundleName = "cisco_ios_xe"
    clmgmtLicenseActionTable.EntityData.ParentYangName = "CISCO-LICENSE-MGMT-MIB"
    clmgmtLicenseActionTable.EntityData.SegmentPath = "clmgmtLicenseActionTable"
    clmgmtLicenseActionTable.EntityData.AbsolutePath = "CISCO-LICENSE-MGMT-MIB:CISCO-LICENSE-MGMT-MIB/" + clmgmtLicenseActionTable.EntityData.SegmentPath
    clmgmtLicenseActionTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtLicenseActionTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtLicenseActionTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtLicenseActionTable.EntityData.Children = types.NewOrderedMap()
    clmgmtLicenseActionTable.EntityData.Children.Append("clmgmtLicenseActionEntry", types.YChild{"ClmgmtLicenseActionEntry", nil})
    for i := range clmgmtLicenseActionTable.ClmgmtLicenseActionEntry {
        clmgmtLicenseActionTable.EntityData.Children.Append(types.GetSegmentPath(clmgmtLicenseActionTable.ClmgmtLicenseActionEntry[i]), types.YChild{"ClmgmtLicenseActionEntry", clmgmtLicenseActionTable.ClmgmtLicenseActionEntry[i]})
    }
    clmgmtLicenseActionTable.EntityData.Leafs = types.NewOrderedMap()

    clmgmtLicenseActionTable.EntityData.YListKeys = []string {}

    return &(clmgmtLicenseActionTable.EntityData)
}

// CISCOLICENSEMGMTMIB_ClmgmtLicenseActionTable_ClmgmtLicenseActionEntry
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
type CISCOLICENSEMGMTMIB_ClmgmtLicenseActionTable_ClmgmtLicenseActionEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object uniquely identifies a row in
    // clmgmtLicenseActionTable. The management application should choose this
    // value by reading clmgmtNextFreeLicenseActionIndex while creating an entry
    // in this table. If an entry already exists with this index, the creation of
    // the entry will not continue and error will be returned. The management
    // application should read the value of clmgmtNextFreeLicenseActionIndex again
    // and retry with the new value for this object. The type is interface{} with
    // range: 1..4294967295.
    ClmgmtLicenseActionIndex interface{}

    // This object represents the entPhysicalIndex of the device where the action
    // is being executed. This object is mainly used in devices where one device
    // is acting as a master and rest of the devices as slaves. The master device
    // is responsible for SNMP communication with the management application.
    // Examples include stackable switches, devices with route processor and line
    // card configuration. If this object is not set, the license action will be
    // executed on the master device. Note: This object need not be set if there
    // is a stand alone device. The type is interface{} with range: 0..2147483647.
    ClmgmtLicenseActionEntPhysicalIndex interface{}

    // This object represents the transfer protocol to be used when copying files
    // as specified in the following objects. 1. clmgmtLicenseFile 2.
    // clmgmtLicensePermissionTicketFile 3. clmgmtLicenseRehostTicketFile 4.
    // clmgmtLicenseBackupFile  Note: This object need not be set if the all the
    // files required for the action are in device's local file system. The type
    // is ClmgmtLicenseTransferProtocol.
    ClmgmtLicenseActionTransferProtocol interface{}

    // This object indicates the transport type of the address contained in
    // clmgmtLicenseServerAddress object. This must be set when
    // clmgmtLicenseActionTransferProtocol is not none(1) or local(2). The type is
    // InetAddressType.
    ClmgmtLicenseServerAddressType interface{}

    // This object indicates the ip address of the server from which the files
    // must be read or written to if clmgmtLicenseActionTransferProtocol is not
    // none(1) or local(2).  All bits as 0s or 1s for clmgmtLicenseServerAddress
    // are not allowed.  The format of this address depends on the value of the
    // clmgmtLicenseServerAddressType object. The type is string with length:
    // 0..255.
    ClmgmtLicenseServerAddress interface{}

    // This object indicates the remote user name for accessing files via ftp,
    // rcp, sftp or scp protocols. This object must be set when the
    // clmgmtLicenseActionTransferProtocol is ftp(4), rcp(5), scp(7) or sftp(8).
    // If clmgmtLicenseActionTransferProtocol is rcp(5), the remote username is
    // sent as the server username in an rcp command request sent by the system to
    // a remote rcp server. The type is string with length: 0..96.
    ClmgmtLicenseServerUsername interface{}

    // This object indicates the password used by ftp, sftp or scp for copying a
    // file to/from an ftp/sftp/scp server. This object must be set when the
    // clmgmtLicenseActionTransferProtocol is ftp(4) or scp(7) or sftp(8). Reading
    // it returns a zero-length string for security reasons. The type is string
    // with length: 0..96.
    ClmgmtLicenseServerPassword interface{}

    // This object represents the location of the license file on the server
    // identified by clmgmtLicenseServerAddress. This object MUST be set to a
    // valid value before or concurrently with setting the value of the
    // clmgmtLicenseAction object to install(2). For other operations, the value
    // of this object is not considered, it is irrelevant. The type is string with
    // length: 0..255.
    ClmgmtLicenseFile interface{}

    // This object represents the clmgmtLicenseStoreIndex of the license store to
    // use within the device. The license store can be a local disk or flash. A
    // device can have more than one license stores. If this object is not set,
    // the license will be stored in the default license store as exposed by
    // clmgmtDefaultLicenseStore object. The type is interface{} with range:
    // 0..4294967295.
    ClmgmtLicenseStore interface{}

    // This object indicates the the license index of the license that is the
    // subject of this action. This is used for identifying a license for
    // performing actions specific to that license. This object need to be set
    // only if clmgmtLicenseAction is set to clear(4). The value of this object is
    // same as the clmgmtLicenseIndex object in clmgmtLicenseInfoEntry for license
    // that is subject of this action. The type is interface{} with range:
    // 0..4294967295.
    ClmgmtLicenseActionLicenseIndex interface{}

    // This object indicates the file name of the permission ticket. This object
    // need to be set only if clmgmtLicenseAction is set to
    // processPermissionTicket(4) or regenerateLastRehostTicket(5) actions. The
    // permission ticket is obtained from Cisco licensing portal to revoke a
    // license. The management application must set this object to valid value
    // before invoking the action. The type is string with length: 0..255.
    ClmgmtLicensePermissionTicketFile interface{}

    // This object indicates the file where the rehost ticket generated by the
    // device need to be exported to. The rehost ticket is generated as a result
    // of processPermissionTicket and regenerateLastRehostTicket actions. After
    // generating the rehost ticket, the device exports the rehost ticket contents
    // to this file. This object need to be set only if clmgmtLicenseAction is set
    // to processPermissionTicket(4) or regenerateLastRehostTicket(5) actions. The
    // type is string with length: 0..255.
    ClmgmtLicenseRehostTicketFile interface{}

    // This object indicates the file where all the licenses in the device need to
    // be backed up. This object need to be set only if clmgmtLicenseAction is set
    // to backup(6) and the management application must set the value of this 
    // object to valid value before invoking action. The type is string with
    // length: 0..255.
    ClmgmtLicenseBackupFile interface{}

    // This object indicates whether the license action should stop if the action
    // on a license fails. This object is applicable only if there are more than
    // one licenses involved in an action. The type is bool.
    ClmgmtLicenseStopOnFailure interface{}

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
    // The type is ClmgmtLicenseAction.
    ClmgmtLicenseAction interface{}

    // This object indicates the state of this license action. The type is
    // ClmgmtLicenseActionState.
    ClmgmtLicenseActionState interface{}

    // This object represents the position of the action in the license action job
    // queue that is maintained internally. Only actions in pending(2) state will
    // be put in the queue until they are executed. By reading this object, the
    // management application can make intelligent decision on whether to execute
    // another action that it is planning on. For example, if there is already a
    // license clear action in the queue in pending(2) state, management
    // application can choose to defer its license back up action to a later time.
    // This object will return a value of 0 if the action is not in pending(2)
    // state. The type is interface{} with range: 0..4294967295.
    ClmgmtLicenseJobQPosition interface{}

    // This object indicates the reason for this license action failure. The value
    // of this object is valid only when clmgmtLicenseActionState is failed(6).
    // The type is ClmgmtLicenseActionFailCause.
    ClmgmtLicenseActionFailCause interface{}

    // This object indicates the storage type for this conceptual row. Conceptual
    // rows having the value 'permanent' need not allow write-access to any
    // columnar objects in the row. The type is StorageType.
    ClmgmtLicenseActionStorageType interface{}

    // This object indicates the the status of this table entry. Once the entry
    // status is set to active(1), the associated entry cannot be modified until
    // the action completes (clmgmtLicenseConfigCommandStatus is set to a value
    // other than inProgress(3)). Once the action completes the only operation
    // possible after this is to delete the row. It is recommended that the
    // management application should delete entries in this table after reading
    // the result. In order to prevent old entries from clogging the table,
    // entries will be aged out, but an entry will never be deleted within 5
    // minutes of completion. The type is RowStatus.
    ClmgmtLicenseActionRowStatus interface{}

    // This object indicates whether the End User License Agreement needed for
    // installing the licenses is accepted.  true(1) - EULA is read and accepted
    // false(2) - EULA is not accepted  Management application should set this
    // object to true(1) when installing licenses that need EULA acceptance. The
    // type is bool.
    ClmgmtLicenseAcceptEULA interface{}

    // This object indicates the file where all the End User License Agreements
    // (EULAs) need to be exported to. This object need to be set only if
    // clmgmtLicenseAction is set to generateEULA(7) and the management
    // application must set the value of this object to valid value before
    // invoking action. The type is string with length: 0..255.
    ClmgmtLicenseEULAFile interface{}
}

func (clmgmtLicenseActionEntry *CISCOLICENSEMGMTMIB_ClmgmtLicenseActionTable_ClmgmtLicenseActionEntry) GetEntityData() *types.CommonEntityData {
    clmgmtLicenseActionEntry.EntityData.YFilter = clmgmtLicenseActionEntry.YFilter
    clmgmtLicenseActionEntry.EntityData.YangName = "clmgmtLicenseActionEntry"
    clmgmtLicenseActionEntry.EntityData.BundleName = "cisco_ios_xe"
    clmgmtLicenseActionEntry.EntityData.ParentYangName = "clmgmtLicenseActionTable"
    clmgmtLicenseActionEntry.EntityData.SegmentPath = "clmgmtLicenseActionEntry" + types.AddKeyToken(clmgmtLicenseActionEntry.ClmgmtLicenseActionIndex, "clmgmtLicenseActionIndex")
    clmgmtLicenseActionEntry.EntityData.AbsolutePath = "CISCO-LICENSE-MGMT-MIB:CISCO-LICENSE-MGMT-MIB/clmgmtLicenseActionTable/" + clmgmtLicenseActionEntry.EntityData.SegmentPath
    clmgmtLicenseActionEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtLicenseActionEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtLicenseActionEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtLicenseActionEntry.EntityData.Children = types.NewOrderedMap()
    clmgmtLicenseActionEntry.EntityData.Leafs = types.NewOrderedMap()
    clmgmtLicenseActionEntry.EntityData.Leafs.Append("clmgmtLicenseActionIndex", types.YLeaf{"ClmgmtLicenseActionIndex", clmgmtLicenseActionEntry.ClmgmtLicenseActionIndex})
    clmgmtLicenseActionEntry.EntityData.Leafs.Append("clmgmtLicenseActionEntPhysicalIndex", types.YLeaf{"ClmgmtLicenseActionEntPhysicalIndex", clmgmtLicenseActionEntry.ClmgmtLicenseActionEntPhysicalIndex})
    clmgmtLicenseActionEntry.EntityData.Leafs.Append("clmgmtLicenseActionTransferProtocol", types.YLeaf{"ClmgmtLicenseActionTransferProtocol", clmgmtLicenseActionEntry.ClmgmtLicenseActionTransferProtocol})
    clmgmtLicenseActionEntry.EntityData.Leafs.Append("clmgmtLicenseServerAddressType", types.YLeaf{"ClmgmtLicenseServerAddressType", clmgmtLicenseActionEntry.ClmgmtLicenseServerAddressType})
    clmgmtLicenseActionEntry.EntityData.Leafs.Append("clmgmtLicenseServerAddress", types.YLeaf{"ClmgmtLicenseServerAddress", clmgmtLicenseActionEntry.ClmgmtLicenseServerAddress})
    clmgmtLicenseActionEntry.EntityData.Leafs.Append("clmgmtLicenseServerUsername", types.YLeaf{"ClmgmtLicenseServerUsername", clmgmtLicenseActionEntry.ClmgmtLicenseServerUsername})
    clmgmtLicenseActionEntry.EntityData.Leafs.Append("clmgmtLicenseServerPassword", types.YLeaf{"ClmgmtLicenseServerPassword", clmgmtLicenseActionEntry.ClmgmtLicenseServerPassword})
    clmgmtLicenseActionEntry.EntityData.Leafs.Append("clmgmtLicenseFile", types.YLeaf{"ClmgmtLicenseFile", clmgmtLicenseActionEntry.ClmgmtLicenseFile})
    clmgmtLicenseActionEntry.EntityData.Leafs.Append("clmgmtLicenseStore", types.YLeaf{"ClmgmtLicenseStore", clmgmtLicenseActionEntry.ClmgmtLicenseStore})
    clmgmtLicenseActionEntry.EntityData.Leafs.Append("clmgmtLicenseActionLicenseIndex", types.YLeaf{"ClmgmtLicenseActionLicenseIndex", clmgmtLicenseActionEntry.ClmgmtLicenseActionLicenseIndex})
    clmgmtLicenseActionEntry.EntityData.Leafs.Append("clmgmtLicensePermissionTicketFile", types.YLeaf{"ClmgmtLicensePermissionTicketFile", clmgmtLicenseActionEntry.ClmgmtLicensePermissionTicketFile})
    clmgmtLicenseActionEntry.EntityData.Leafs.Append("clmgmtLicenseRehostTicketFile", types.YLeaf{"ClmgmtLicenseRehostTicketFile", clmgmtLicenseActionEntry.ClmgmtLicenseRehostTicketFile})
    clmgmtLicenseActionEntry.EntityData.Leafs.Append("clmgmtLicenseBackupFile", types.YLeaf{"ClmgmtLicenseBackupFile", clmgmtLicenseActionEntry.ClmgmtLicenseBackupFile})
    clmgmtLicenseActionEntry.EntityData.Leafs.Append("clmgmtLicenseStopOnFailure", types.YLeaf{"ClmgmtLicenseStopOnFailure", clmgmtLicenseActionEntry.ClmgmtLicenseStopOnFailure})
    clmgmtLicenseActionEntry.EntityData.Leafs.Append("clmgmtLicenseAction", types.YLeaf{"ClmgmtLicenseAction", clmgmtLicenseActionEntry.ClmgmtLicenseAction})
    clmgmtLicenseActionEntry.EntityData.Leafs.Append("clmgmtLicenseActionState", types.YLeaf{"ClmgmtLicenseActionState", clmgmtLicenseActionEntry.ClmgmtLicenseActionState})
    clmgmtLicenseActionEntry.EntityData.Leafs.Append("clmgmtLicenseJobQPosition", types.YLeaf{"ClmgmtLicenseJobQPosition", clmgmtLicenseActionEntry.ClmgmtLicenseJobQPosition})
    clmgmtLicenseActionEntry.EntityData.Leafs.Append("clmgmtLicenseActionFailCause", types.YLeaf{"ClmgmtLicenseActionFailCause", clmgmtLicenseActionEntry.ClmgmtLicenseActionFailCause})
    clmgmtLicenseActionEntry.EntityData.Leafs.Append("clmgmtLicenseActionStorageType", types.YLeaf{"ClmgmtLicenseActionStorageType", clmgmtLicenseActionEntry.ClmgmtLicenseActionStorageType})
    clmgmtLicenseActionEntry.EntityData.Leafs.Append("clmgmtLicenseActionRowStatus", types.YLeaf{"ClmgmtLicenseActionRowStatus", clmgmtLicenseActionEntry.ClmgmtLicenseActionRowStatus})
    clmgmtLicenseActionEntry.EntityData.Leafs.Append("clmgmtLicenseAcceptEULA", types.YLeaf{"ClmgmtLicenseAcceptEULA", clmgmtLicenseActionEntry.ClmgmtLicenseAcceptEULA})
    clmgmtLicenseActionEntry.EntityData.Leafs.Append("clmgmtLicenseEULAFile", types.YLeaf{"ClmgmtLicenseEULAFile", clmgmtLicenseActionEntry.ClmgmtLicenseEULAFile})

    clmgmtLicenseActionEntry.EntityData.YListKeys = []string {"ClmgmtLicenseActionIndex"}

    return &(clmgmtLicenseActionEntry.EntityData)
}

// CISCOLICENSEMGMTMIB_ClmgmtLicenseActionTable_ClmgmtLicenseActionEntry_ClmgmtLicenseAction represents                                needed EULA contents to a file.
type CISCOLICENSEMGMTMIB_ClmgmtLicenseActionTable_ClmgmtLicenseActionEntry_ClmgmtLicenseAction string

const (
    CISCOLICENSEMGMTMIB_ClmgmtLicenseActionTable_ClmgmtLicenseActionEntry_ClmgmtLicenseAction_noOp CISCOLICENSEMGMTMIB_ClmgmtLicenseActionTable_ClmgmtLicenseActionEntry_ClmgmtLicenseAction = "noOp"

    CISCOLICENSEMGMTMIB_ClmgmtLicenseActionTable_ClmgmtLicenseActionEntry_ClmgmtLicenseAction_install CISCOLICENSEMGMTMIB_ClmgmtLicenseActionTable_ClmgmtLicenseActionEntry_ClmgmtLicenseAction = "install"

    CISCOLICENSEMGMTMIB_ClmgmtLicenseActionTable_ClmgmtLicenseActionEntry_ClmgmtLicenseAction_clear CISCOLICENSEMGMTMIB_ClmgmtLicenseActionTable_ClmgmtLicenseActionEntry_ClmgmtLicenseAction = "clear"

    CISCOLICENSEMGMTMIB_ClmgmtLicenseActionTable_ClmgmtLicenseActionEntry_ClmgmtLicenseAction_processPermissionTicket CISCOLICENSEMGMTMIB_ClmgmtLicenseActionTable_ClmgmtLicenseActionEntry_ClmgmtLicenseAction = "processPermissionTicket"

    CISCOLICENSEMGMTMIB_ClmgmtLicenseActionTable_ClmgmtLicenseActionEntry_ClmgmtLicenseAction_regenerateLastRehostTicket CISCOLICENSEMGMTMIB_ClmgmtLicenseActionTable_ClmgmtLicenseActionEntry_ClmgmtLicenseAction = "regenerateLastRehostTicket"

    CISCOLICENSEMGMTMIB_ClmgmtLicenseActionTable_ClmgmtLicenseActionEntry_ClmgmtLicenseAction_backup CISCOLICENSEMGMTMIB_ClmgmtLicenseActionTable_ClmgmtLicenseActionEntry_ClmgmtLicenseAction = "backup"

    CISCOLICENSEMGMTMIB_ClmgmtLicenseActionTable_ClmgmtLicenseActionEntry_ClmgmtLicenseAction_generateEULA CISCOLICENSEMGMTMIB_ClmgmtLicenseActionTable_ClmgmtLicenseActionEntry_ClmgmtLicenseAction = "generateEULA"
)

// CISCOLICENSEMGMTMIB_ClmgmtLicenseActionResultTable
// This table contains results of license action if the
// license action involves multiple licenses. Entries in this
// table are not created for actions where there is
// only license that is subject of the action. For
// example, if there are 3 licenses in a license file
// when executing license install action, 3 entries will
// be created in this table, one for each license.
type CISCOLICENSEMGMTMIB_ClmgmtLicenseActionResultTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in clmgmtLicenseActionResultTable. Each entry contains result of
    // the action for a single license. These entries are created immediately
    // after action execution when the action involves multiple licenses. These
    // entries get automatically deleted when the corresponding entry in
    // clmgmtLicenseActionTable is deleted. The type is slice of
    // CISCOLICENSEMGMTMIB_ClmgmtLicenseActionResultTable_ClmgmtLicenseActionResultEntry.
    ClmgmtLicenseActionResultEntry []*CISCOLICENSEMGMTMIB_ClmgmtLicenseActionResultTable_ClmgmtLicenseActionResultEntry
}

func (clmgmtLicenseActionResultTable *CISCOLICENSEMGMTMIB_ClmgmtLicenseActionResultTable) GetEntityData() *types.CommonEntityData {
    clmgmtLicenseActionResultTable.EntityData.YFilter = clmgmtLicenseActionResultTable.YFilter
    clmgmtLicenseActionResultTable.EntityData.YangName = "clmgmtLicenseActionResultTable"
    clmgmtLicenseActionResultTable.EntityData.BundleName = "cisco_ios_xe"
    clmgmtLicenseActionResultTable.EntityData.ParentYangName = "CISCO-LICENSE-MGMT-MIB"
    clmgmtLicenseActionResultTable.EntityData.SegmentPath = "clmgmtLicenseActionResultTable"
    clmgmtLicenseActionResultTable.EntityData.AbsolutePath = "CISCO-LICENSE-MGMT-MIB:CISCO-LICENSE-MGMT-MIB/" + clmgmtLicenseActionResultTable.EntityData.SegmentPath
    clmgmtLicenseActionResultTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtLicenseActionResultTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtLicenseActionResultTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtLicenseActionResultTable.EntityData.Children = types.NewOrderedMap()
    clmgmtLicenseActionResultTable.EntityData.Children.Append("clmgmtLicenseActionResultEntry", types.YChild{"ClmgmtLicenseActionResultEntry", nil})
    for i := range clmgmtLicenseActionResultTable.ClmgmtLicenseActionResultEntry {
        clmgmtLicenseActionResultTable.EntityData.Children.Append(types.GetSegmentPath(clmgmtLicenseActionResultTable.ClmgmtLicenseActionResultEntry[i]), types.YChild{"ClmgmtLicenseActionResultEntry", clmgmtLicenseActionResultTable.ClmgmtLicenseActionResultEntry[i]})
    }
    clmgmtLicenseActionResultTable.EntityData.Leafs = types.NewOrderedMap()

    clmgmtLicenseActionResultTable.EntityData.YListKeys = []string {}

    return &(clmgmtLicenseActionResultTable.EntityData)
}

// CISCOLICENSEMGMTMIB_ClmgmtLicenseActionResultTable_ClmgmtLicenseActionResultEntry
// An entry in clmgmtLicenseActionResultTable. Each entry
// contains result of the action for a single license.
// These entries are created immediately after action
// execution when the action involves multiple licenses.
// These entries get automatically deleted when the
// corresponding entry in clmgmtLicenseActionTable
// is deleted.
type CISCOLICENSEMGMTMIB_ClmgmtLicenseActionResultTable_ClmgmtLicenseActionResultEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..4294967295.
    // Refers to
    // cisco_license_mgmt_mib.CISCOLICENSEMGMTMIB_ClmgmtLicenseActionTable_ClmgmtLicenseActionEntry_ClmgmtLicenseActionIndex
    ClmgmtLicenseActionIndex interface{}

    // This attribute is a key. This object indicates the sequence number of this
    // license in the list of licenses on which the action is executed. For
    // example, if there are 3 licenses in a license file when executing license
    // install action, this object will have values 1, 2 and 3 respectively as
    // ordered in the license file. The type is interface{} with range:
    // 1..4294967295.
    ClmgmtLicenseNumber interface{}

    // This object indicates the state of action on this individual license. The
    // type is ClmgmtLicenseActionState.
    ClmgmtLicenseIndivActionState interface{}

    // This object indicates the reason for action failure on this individual
    // license. The type is ClmgmtLicenseActionFailCause.
    ClmgmtLicenseIndivActionFailCause interface{}
}

func (clmgmtLicenseActionResultEntry *CISCOLICENSEMGMTMIB_ClmgmtLicenseActionResultTable_ClmgmtLicenseActionResultEntry) GetEntityData() *types.CommonEntityData {
    clmgmtLicenseActionResultEntry.EntityData.YFilter = clmgmtLicenseActionResultEntry.YFilter
    clmgmtLicenseActionResultEntry.EntityData.YangName = "clmgmtLicenseActionResultEntry"
    clmgmtLicenseActionResultEntry.EntityData.BundleName = "cisco_ios_xe"
    clmgmtLicenseActionResultEntry.EntityData.ParentYangName = "clmgmtLicenseActionResultTable"
    clmgmtLicenseActionResultEntry.EntityData.SegmentPath = "clmgmtLicenseActionResultEntry" + types.AddKeyToken(clmgmtLicenseActionResultEntry.ClmgmtLicenseActionIndex, "clmgmtLicenseActionIndex") + types.AddKeyToken(clmgmtLicenseActionResultEntry.ClmgmtLicenseNumber, "clmgmtLicenseNumber")
    clmgmtLicenseActionResultEntry.EntityData.AbsolutePath = "CISCO-LICENSE-MGMT-MIB:CISCO-LICENSE-MGMT-MIB/clmgmtLicenseActionResultTable/" + clmgmtLicenseActionResultEntry.EntityData.SegmentPath
    clmgmtLicenseActionResultEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtLicenseActionResultEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtLicenseActionResultEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtLicenseActionResultEntry.EntityData.Children = types.NewOrderedMap()
    clmgmtLicenseActionResultEntry.EntityData.Leafs = types.NewOrderedMap()
    clmgmtLicenseActionResultEntry.EntityData.Leafs.Append("clmgmtLicenseActionIndex", types.YLeaf{"ClmgmtLicenseActionIndex", clmgmtLicenseActionResultEntry.ClmgmtLicenseActionIndex})
    clmgmtLicenseActionResultEntry.EntityData.Leafs.Append("clmgmtLicenseNumber", types.YLeaf{"ClmgmtLicenseNumber", clmgmtLicenseActionResultEntry.ClmgmtLicenseNumber})
    clmgmtLicenseActionResultEntry.EntityData.Leafs.Append("clmgmtLicenseIndivActionState", types.YLeaf{"ClmgmtLicenseIndivActionState", clmgmtLicenseActionResultEntry.ClmgmtLicenseIndivActionState})
    clmgmtLicenseActionResultEntry.EntityData.Leafs.Append("clmgmtLicenseIndivActionFailCause", types.YLeaf{"ClmgmtLicenseIndivActionFailCause", clmgmtLicenseActionResultEntry.ClmgmtLicenseIndivActionFailCause})

    clmgmtLicenseActionResultEntry.EntityData.YListKeys = []string {"ClmgmtLicenseActionIndex", "ClmgmtLicenseNumber"}

    return &(clmgmtLicenseActionResultEntry.EntityData)
}

// CISCOLICENSEMGMTMIB_ClmgmtLicenseStoreInfoTable
// This table contains information about all the license
// stores allocated on the device.
type CISCOLICENSEMGMTMIB_ClmgmtLicenseStoreInfoTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in clmgmtLicenseStoreInfoTable. Each entry contains information
    // about a license store allocated on the device. The type is slice of
    // CISCOLICENSEMGMTMIB_ClmgmtLicenseStoreInfoTable_ClmgmtLicenseStoreInfoEntry.
    ClmgmtLicenseStoreInfoEntry []*CISCOLICENSEMGMTMIB_ClmgmtLicenseStoreInfoTable_ClmgmtLicenseStoreInfoEntry
}

func (clmgmtLicenseStoreInfoTable *CISCOLICENSEMGMTMIB_ClmgmtLicenseStoreInfoTable) GetEntityData() *types.CommonEntityData {
    clmgmtLicenseStoreInfoTable.EntityData.YFilter = clmgmtLicenseStoreInfoTable.YFilter
    clmgmtLicenseStoreInfoTable.EntityData.YangName = "clmgmtLicenseStoreInfoTable"
    clmgmtLicenseStoreInfoTable.EntityData.BundleName = "cisco_ios_xe"
    clmgmtLicenseStoreInfoTable.EntityData.ParentYangName = "CISCO-LICENSE-MGMT-MIB"
    clmgmtLicenseStoreInfoTable.EntityData.SegmentPath = "clmgmtLicenseStoreInfoTable"
    clmgmtLicenseStoreInfoTable.EntityData.AbsolutePath = "CISCO-LICENSE-MGMT-MIB:CISCO-LICENSE-MGMT-MIB/" + clmgmtLicenseStoreInfoTable.EntityData.SegmentPath
    clmgmtLicenseStoreInfoTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtLicenseStoreInfoTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtLicenseStoreInfoTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtLicenseStoreInfoTable.EntityData.Children = types.NewOrderedMap()
    clmgmtLicenseStoreInfoTable.EntityData.Children.Append("clmgmtLicenseStoreInfoEntry", types.YChild{"ClmgmtLicenseStoreInfoEntry", nil})
    for i := range clmgmtLicenseStoreInfoTable.ClmgmtLicenseStoreInfoEntry {
        clmgmtLicenseStoreInfoTable.EntityData.Children.Append(types.GetSegmentPath(clmgmtLicenseStoreInfoTable.ClmgmtLicenseStoreInfoEntry[i]), types.YChild{"ClmgmtLicenseStoreInfoEntry", clmgmtLicenseStoreInfoTable.ClmgmtLicenseStoreInfoEntry[i]})
    }
    clmgmtLicenseStoreInfoTable.EntityData.Leafs = types.NewOrderedMap()

    clmgmtLicenseStoreInfoTable.EntityData.YListKeys = []string {}

    return &(clmgmtLicenseStoreInfoTable.EntityData)
}

// CISCOLICENSEMGMTMIB_ClmgmtLicenseStoreInfoTable_ClmgmtLicenseStoreInfoEntry
// An entry in clmgmtLicenseStoreInfoTable. Each entry
// contains information about a license store allocated
// on the device
type CISCOLICENSEMGMTMIB_ClmgmtLicenseStoreInfoTable_ClmgmtLicenseStoreInfoEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. This object uniquely identifies a license store
    // within the device. The type is interface{} with range: 1..4294967295.
    ClmgmtLicenseStoreIndex interface{}

    // This object indicates the name of the license store within the device. It
    // is a file in device's local file system i.e., either on a local disk or
    // flash or some other storage media. For example, the value of this object
    // can be 'disk1:lic_store_1.txt' or 'flash:lic_store_2.txt. The type is
    // string with length: 0..255.
    ClmgmtLicenseStoreName interface{}

    // This object indicates the total number of bytes that are allocated to the
    // license store. The type is interface{} with range: 0..4294967295. Units are
    // bytes.
    ClmgmtLicenseStoreTotalSize interface{}

    // This object indicates the number of bytes still remaining to be used for
    // new license installations in the license store. The type is interface{}
    // with range: 0..4294967295. Units are bytes.
    ClmgmtLicenseStoreSizeRemaining interface{}
}

func (clmgmtLicenseStoreInfoEntry *CISCOLICENSEMGMTMIB_ClmgmtLicenseStoreInfoTable_ClmgmtLicenseStoreInfoEntry) GetEntityData() *types.CommonEntityData {
    clmgmtLicenseStoreInfoEntry.EntityData.YFilter = clmgmtLicenseStoreInfoEntry.YFilter
    clmgmtLicenseStoreInfoEntry.EntityData.YangName = "clmgmtLicenseStoreInfoEntry"
    clmgmtLicenseStoreInfoEntry.EntityData.BundleName = "cisco_ios_xe"
    clmgmtLicenseStoreInfoEntry.EntityData.ParentYangName = "clmgmtLicenseStoreInfoTable"
    clmgmtLicenseStoreInfoEntry.EntityData.SegmentPath = "clmgmtLicenseStoreInfoEntry" + types.AddKeyToken(clmgmtLicenseStoreInfoEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(clmgmtLicenseStoreInfoEntry.ClmgmtLicenseStoreIndex, "clmgmtLicenseStoreIndex")
    clmgmtLicenseStoreInfoEntry.EntityData.AbsolutePath = "CISCO-LICENSE-MGMT-MIB:CISCO-LICENSE-MGMT-MIB/clmgmtLicenseStoreInfoTable/" + clmgmtLicenseStoreInfoEntry.EntityData.SegmentPath
    clmgmtLicenseStoreInfoEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtLicenseStoreInfoEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtLicenseStoreInfoEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtLicenseStoreInfoEntry.EntityData.Children = types.NewOrderedMap()
    clmgmtLicenseStoreInfoEntry.EntityData.Leafs = types.NewOrderedMap()
    clmgmtLicenseStoreInfoEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", clmgmtLicenseStoreInfoEntry.EntPhysicalIndex})
    clmgmtLicenseStoreInfoEntry.EntityData.Leafs.Append("clmgmtLicenseStoreIndex", types.YLeaf{"ClmgmtLicenseStoreIndex", clmgmtLicenseStoreInfoEntry.ClmgmtLicenseStoreIndex})
    clmgmtLicenseStoreInfoEntry.EntityData.Leafs.Append("clmgmtLicenseStoreName", types.YLeaf{"ClmgmtLicenseStoreName", clmgmtLicenseStoreInfoEntry.ClmgmtLicenseStoreName})
    clmgmtLicenseStoreInfoEntry.EntityData.Leafs.Append("clmgmtLicenseStoreTotalSize", types.YLeaf{"ClmgmtLicenseStoreTotalSize", clmgmtLicenseStoreInfoEntry.ClmgmtLicenseStoreTotalSize})
    clmgmtLicenseStoreInfoEntry.EntityData.Leafs.Append("clmgmtLicenseStoreSizeRemaining", types.YLeaf{"ClmgmtLicenseStoreSizeRemaining", clmgmtLicenseStoreInfoEntry.ClmgmtLicenseStoreSizeRemaining})

    clmgmtLicenseStoreInfoEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "ClmgmtLicenseStoreIndex"}

    return &(clmgmtLicenseStoreInfoEntry.EntityData)
}

// CISCOLICENSEMGMTMIB_ClmgmtLicenseDeviceInfoTable
// This table contains objects that provide licensing related
// information at the device level. Entries will exist
// only for entities that support licensing. For example,
// if it is a stand alone device and supports licensing,
// then there will be only one entry in this table. If
// it is stackable switch then there will be multiple
// entries with one entry for each device in the stack.
type CISCOLICENSEMGMTMIB_ClmgmtLicenseDeviceInfoTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in clmgmtLicenseDeviceInfoTable. Each entry contains device level
    // licensing information for a device. The type is slice of
    // CISCOLICENSEMGMTMIB_ClmgmtLicenseDeviceInfoTable_ClmgmtLicenseDeviceInfoEntry.
    ClmgmtLicenseDeviceInfoEntry []*CISCOLICENSEMGMTMIB_ClmgmtLicenseDeviceInfoTable_ClmgmtLicenseDeviceInfoEntry
}

func (clmgmtLicenseDeviceInfoTable *CISCOLICENSEMGMTMIB_ClmgmtLicenseDeviceInfoTable) GetEntityData() *types.CommonEntityData {
    clmgmtLicenseDeviceInfoTable.EntityData.YFilter = clmgmtLicenseDeviceInfoTable.YFilter
    clmgmtLicenseDeviceInfoTable.EntityData.YangName = "clmgmtLicenseDeviceInfoTable"
    clmgmtLicenseDeviceInfoTable.EntityData.BundleName = "cisco_ios_xe"
    clmgmtLicenseDeviceInfoTable.EntityData.ParentYangName = "CISCO-LICENSE-MGMT-MIB"
    clmgmtLicenseDeviceInfoTable.EntityData.SegmentPath = "clmgmtLicenseDeviceInfoTable"
    clmgmtLicenseDeviceInfoTable.EntityData.AbsolutePath = "CISCO-LICENSE-MGMT-MIB:CISCO-LICENSE-MGMT-MIB/" + clmgmtLicenseDeviceInfoTable.EntityData.SegmentPath
    clmgmtLicenseDeviceInfoTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtLicenseDeviceInfoTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtLicenseDeviceInfoTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtLicenseDeviceInfoTable.EntityData.Children = types.NewOrderedMap()
    clmgmtLicenseDeviceInfoTable.EntityData.Children.Append("clmgmtLicenseDeviceInfoEntry", types.YChild{"ClmgmtLicenseDeviceInfoEntry", nil})
    for i := range clmgmtLicenseDeviceInfoTable.ClmgmtLicenseDeviceInfoEntry {
        clmgmtLicenseDeviceInfoTable.EntityData.Children.Append(types.GetSegmentPath(clmgmtLicenseDeviceInfoTable.ClmgmtLicenseDeviceInfoEntry[i]), types.YChild{"ClmgmtLicenseDeviceInfoEntry", clmgmtLicenseDeviceInfoTable.ClmgmtLicenseDeviceInfoEntry[i]})
    }
    clmgmtLicenseDeviceInfoTable.EntityData.Leafs = types.NewOrderedMap()

    clmgmtLicenseDeviceInfoTable.EntityData.YListKeys = []string {}

    return &(clmgmtLicenseDeviceInfoTable.EntityData)
}

// CISCOLICENSEMGMTMIB_ClmgmtLicenseDeviceInfoTable_ClmgmtLicenseDeviceInfoEntry
// An entry in clmgmtLicenseDeviceInfoTable. Each entry
// contains device level licensing information for a device.
type CISCOLICENSEMGMTMIB_ClmgmtLicenseDeviceInfoTable_ClmgmtLicenseDeviceInfoEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This object indicates the clmgmtLicenseStoreIndex of default store in the
    // device. There will be only one default license store per device. If no
    // license store is specified during license install, this default license
    // store will be used. The type is interface{} with range: 1..4294967295.
    ClmgmtDefaultLicenseStore interface{}
}

func (clmgmtLicenseDeviceInfoEntry *CISCOLICENSEMGMTMIB_ClmgmtLicenseDeviceInfoTable_ClmgmtLicenseDeviceInfoEntry) GetEntityData() *types.CommonEntityData {
    clmgmtLicenseDeviceInfoEntry.EntityData.YFilter = clmgmtLicenseDeviceInfoEntry.YFilter
    clmgmtLicenseDeviceInfoEntry.EntityData.YangName = "clmgmtLicenseDeviceInfoEntry"
    clmgmtLicenseDeviceInfoEntry.EntityData.BundleName = "cisco_ios_xe"
    clmgmtLicenseDeviceInfoEntry.EntityData.ParentYangName = "clmgmtLicenseDeviceInfoTable"
    clmgmtLicenseDeviceInfoEntry.EntityData.SegmentPath = "clmgmtLicenseDeviceInfoEntry" + types.AddKeyToken(clmgmtLicenseDeviceInfoEntry.EntPhysicalIndex, "entPhysicalIndex")
    clmgmtLicenseDeviceInfoEntry.EntityData.AbsolutePath = "CISCO-LICENSE-MGMT-MIB:CISCO-LICENSE-MGMT-MIB/clmgmtLicenseDeviceInfoTable/" + clmgmtLicenseDeviceInfoEntry.EntityData.SegmentPath
    clmgmtLicenseDeviceInfoEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtLicenseDeviceInfoEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtLicenseDeviceInfoEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtLicenseDeviceInfoEntry.EntityData.Children = types.NewOrderedMap()
    clmgmtLicenseDeviceInfoEntry.EntityData.Leafs = types.NewOrderedMap()
    clmgmtLicenseDeviceInfoEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", clmgmtLicenseDeviceInfoEntry.EntPhysicalIndex})
    clmgmtLicenseDeviceInfoEntry.EntityData.Leafs.Append("clmgmtDefaultLicenseStore", types.YLeaf{"ClmgmtDefaultLicenseStore", clmgmtLicenseDeviceInfoEntry.ClmgmtDefaultLicenseStore})

    clmgmtLicenseDeviceInfoEntry.EntityData.YListKeys = []string {"EntPhysicalIndex"}

    return &(clmgmtLicenseDeviceInfoEntry.EntityData)
}

// CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable
// This table contains information about all the licenses
// installed on the device.
type CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in clmgmtLicenseInfoTable. Each entry contains information about a
    // license installed on the device. This entry gets created when a license is
    // installed successfully. Management application can not create these entries
    // directly, but will do so indirectly by executing license install action.
    // Some of these entries may already exist that correspond to demo licenses
    // even before management application installs any licenses. The type is slice
    // of CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry.
    ClmgmtLicenseInfoEntry []*CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry
}

func (clmgmtLicenseInfoTable *CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable) GetEntityData() *types.CommonEntityData {
    clmgmtLicenseInfoTable.EntityData.YFilter = clmgmtLicenseInfoTable.YFilter
    clmgmtLicenseInfoTable.EntityData.YangName = "clmgmtLicenseInfoTable"
    clmgmtLicenseInfoTable.EntityData.BundleName = "cisco_ios_xe"
    clmgmtLicenseInfoTable.EntityData.ParentYangName = "CISCO-LICENSE-MGMT-MIB"
    clmgmtLicenseInfoTable.EntityData.SegmentPath = "clmgmtLicenseInfoTable"
    clmgmtLicenseInfoTable.EntityData.AbsolutePath = "CISCO-LICENSE-MGMT-MIB:CISCO-LICENSE-MGMT-MIB/" + clmgmtLicenseInfoTable.EntityData.SegmentPath
    clmgmtLicenseInfoTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtLicenseInfoTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtLicenseInfoTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtLicenseInfoTable.EntityData.Children = types.NewOrderedMap()
    clmgmtLicenseInfoTable.EntityData.Children.Append("clmgmtLicenseInfoEntry", types.YChild{"ClmgmtLicenseInfoEntry", nil})
    for i := range clmgmtLicenseInfoTable.ClmgmtLicenseInfoEntry {
        clmgmtLicenseInfoTable.EntityData.Children.Append(types.GetSegmentPath(clmgmtLicenseInfoTable.ClmgmtLicenseInfoEntry[i]), types.YChild{"ClmgmtLicenseInfoEntry", clmgmtLicenseInfoTable.ClmgmtLicenseInfoEntry[i]})
    }
    clmgmtLicenseInfoTable.EntityData.Leafs = types.NewOrderedMap()

    clmgmtLicenseInfoTable.EntityData.YListKeys = []string {}

    return &(clmgmtLicenseInfoTable.EntityData)
}

// CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry
// An entry in clmgmtLicenseInfoTable. Each entry contains
// information about a license installed on the device. This
// entry gets created when a license is installed successfully.
// Management application can not create these entries directly, but
// will do so indirectly by executing license install action.
// Some of these entries may already exist that correspond to
// demo licenses even before management application installs any
// licenses.
type CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. This object represents the license store that is
    // used for storing this license. This object will have the same value as
    // clmgmtLicenseStoreIndex in clmgmtLicenseStoreInfoEntry of the license store
    // used. The type is interface{} with range: 1..4294967295.
    ClmgmtLicenseStoreUsed interface{}

    // This attribute is a key. This object uniquely identifies a license within
    // the device. The type is interface{} with range: 1..4294967295.
    ClmgmtLicenseIndex interface{}

    // This object indicates the name of the feature that is using or can use this
    // license. A license can be used by only one feature. Examples of feature
    // name are: 'IPBASE', 'ADVIPSERVICE'. The type is string with length: 0..128.
    ClmgmtLicenseFeatureName interface{}

    // This object indicates the version of the feature that is using or can use
    // this license. Examples of feature version are: '1.0', '2.0'. The type is
    // string with length: 0..128.
    ClmgmtLicenseFeatureVersion interface{}

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
    // evaluation mode for a limited time. The type is ClmgmtLicenseType.
    ClmgmtLicenseType interface{}

    // This object indicates whether the license is counted license. true(1)  -
    // counted license false(2) - uncounted license. The type is bool.
    ClmgmtLicenseCounted interface{}

    // This object indicates the time period the license is valid for. This object
    // is applicable only if clmgmtLicenseType is demo(1), or extension(2) or
    // gracePeriod(3) or evalRightToUse(8). The object will return 0 for other
    // license types. The type is interface{} with range: 0..4294967295. Units are
    // seconds.
    ClmgmtLicenseValidityPeriod interface{}

    // This object indicates the time period remaining before the license expires
    // or transitions to rightToUse(9) license. This object is applicable only if
    // clmgmtLicenseType is demo(1), or extension(2) or gracePeriod(3) or
    // evalRightToUse(8). The object will contain 0 for other license types. The
    // type is interface{} with range: 0..4294967295. Units are seconds.
    ClmgmtLicenseValidityPeriodRemaining interface{}

    // This object indicates the elapsed time period since the license expired.
    // This object is applicable only if clmgmtLicenseType is demo(1), or
    // extension(2) or gracePeriod(3). Also, this value of this object will be
    // valid only after the license expires. The object will return 0 for other
    // license types or before the license expiry. The type is interface{} with
    // range: 0..4294967295. Units are seconds.
    ClmgmtLicenseExpiredPeriod interface{}

    // This object indicates the maximum number of entities that can use this
    // license. This object is applicable only if clmgmtLicenseCounted is true(1).
    // The entity that is being counted can be anything and it depends on the
    // licensable feature. The type is interface{} with range: 0..4294967295.
    ClmgmtLicenseMaxUsageCount interface{}

    // This object indicates the number of entities that can still use this
    // license. This object is applicable only if clmgmtLicenseCounted is true(1).
    // The type is interface{} with range: 0..4294967295.
    ClmgmtLicenseUsageCountRemaining interface{}

    // This object indicates whether the user accepted End User License Agreement
    // for this license.  true(1)  - EULA accpeted false(2) - EULA not accepted.
    // The type is bool.
    ClmgmtLicenseEULAStatus interface{}

    // This object represents the user modifiable comments about the license. This
    // object is initially populated with comments from the license file. The type
    // is string with length: 0..255.
    ClmgmtLicenseComments interface{}

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
    // is ClmgmtLicenseStatus.
    ClmgmtLicenseStatus interface{}

    // This object indicates the start date for a subscription license. It is
    // optional for subscription linceses to have a start date associated with
    // them, they may only have an end date associated with them. This object may
    // be applicable only when clmgmtLicenseType is paidSubscription(5),
    // evaluationSubscription(6) or extensionSubscription (7).       The object
    // will contain an octet string of length 0 when it is not applicable. The
    // type is string.
    ClmgmtLicenseStartDate interface{}

    // This object indicates the end date for a subscription license. This object
    // is applicable only when clmgmtLicenseType is paidSubscription(5),
    // evaluationSubscription(6) or extensionSubscription (7). The object will
    // contain an octet string of length 0 when it is not applicable. The type is
    // string.
    ClmgmtLicenseEndDate interface{}

    // This object indicates the time period used for the Right to use (RTU)
    // licenses. This object is applicable for all RTU licenses. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    ClmgmtLicensePeriodUsed interface{}
}

func (clmgmtLicenseInfoEntry *CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry) GetEntityData() *types.CommonEntityData {
    clmgmtLicenseInfoEntry.EntityData.YFilter = clmgmtLicenseInfoEntry.YFilter
    clmgmtLicenseInfoEntry.EntityData.YangName = "clmgmtLicenseInfoEntry"
    clmgmtLicenseInfoEntry.EntityData.BundleName = "cisco_ios_xe"
    clmgmtLicenseInfoEntry.EntityData.ParentYangName = "clmgmtLicenseInfoTable"
    clmgmtLicenseInfoEntry.EntityData.SegmentPath = "clmgmtLicenseInfoEntry" + types.AddKeyToken(clmgmtLicenseInfoEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(clmgmtLicenseInfoEntry.ClmgmtLicenseStoreUsed, "clmgmtLicenseStoreUsed") + types.AddKeyToken(clmgmtLicenseInfoEntry.ClmgmtLicenseIndex, "clmgmtLicenseIndex")
    clmgmtLicenseInfoEntry.EntityData.AbsolutePath = "CISCO-LICENSE-MGMT-MIB:CISCO-LICENSE-MGMT-MIB/clmgmtLicenseInfoTable/" + clmgmtLicenseInfoEntry.EntityData.SegmentPath
    clmgmtLicenseInfoEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtLicenseInfoEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtLicenseInfoEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtLicenseInfoEntry.EntityData.Children = types.NewOrderedMap()
    clmgmtLicenseInfoEntry.EntityData.Leafs = types.NewOrderedMap()
    clmgmtLicenseInfoEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", clmgmtLicenseInfoEntry.EntPhysicalIndex})
    clmgmtLicenseInfoEntry.EntityData.Leafs.Append("clmgmtLicenseStoreUsed", types.YLeaf{"ClmgmtLicenseStoreUsed", clmgmtLicenseInfoEntry.ClmgmtLicenseStoreUsed})
    clmgmtLicenseInfoEntry.EntityData.Leafs.Append("clmgmtLicenseIndex", types.YLeaf{"ClmgmtLicenseIndex", clmgmtLicenseInfoEntry.ClmgmtLicenseIndex})
    clmgmtLicenseInfoEntry.EntityData.Leafs.Append("clmgmtLicenseFeatureName", types.YLeaf{"ClmgmtLicenseFeatureName", clmgmtLicenseInfoEntry.ClmgmtLicenseFeatureName})
    clmgmtLicenseInfoEntry.EntityData.Leafs.Append("clmgmtLicenseFeatureVersion", types.YLeaf{"ClmgmtLicenseFeatureVersion", clmgmtLicenseInfoEntry.ClmgmtLicenseFeatureVersion})
    clmgmtLicenseInfoEntry.EntityData.Leafs.Append("clmgmtLicenseType", types.YLeaf{"ClmgmtLicenseType", clmgmtLicenseInfoEntry.ClmgmtLicenseType})
    clmgmtLicenseInfoEntry.EntityData.Leafs.Append("clmgmtLicenseCounted", types.YLeaf{"ClmgmtLicenseCounted", clmgmtLicenseInfoEntry.ClmgmtLicenseCounted})
    clmgmtLicenseInfoEntry.EntityData.Leafs.Append("clmgmtLicenseValidityPeriod", types.YLeaf{"ClmgmtLicenseValidityPeriod", clmgmtLicenseInfoEntry.ClmgmtLicenseValidityPeriod})
    clmgmtLicenseInfoEntry.EntityData.Leafs.Append("clmgmtLicenseValidityPeriodRemaining", types.YLeaf{"ClmgmtLicenseValidityPeriodRemaining", clmgmtLicenseInfoEntry.ClmgmtLicenseValidityPeriodRemaining})
    clmgmtLicenseInfoEntry.EntityData.Leafs.Append("clmgmtLicenseExpiredPeriod", types.YLeaf{"ClmgmtLicenseExpiredPeriod", clmgmtLicenseInfoEntry.ClmgmtLicenseExpiredPeriod})
    clmgmtLicenseInfoEntry.EntityData.Leafs.Append("clmgmtLicenseMaxUsageCount", types.YLeaf{"ClmgmtLicenseMaxUsageCount", clmgmtLicenseInfoEntry.ClmgmtLicenseMaxUsageCount})
    clmgmtLicenseInfoEntry.EntityData.Leafs.Append("clmgmtLicenseUsageCountRemaining", types.YLeaf{"ClmgmtLicenseUsageCountRemaining", clmgmtLicenseInfoEntry.ClmgmtLicenseUsageCountRemaining})
    clmgmtLicenseInfoEntry.EntityData.Leafs.Append("clmgmtLicenseEULAStatus", types.YLeaf{"ClmgmtLicenseEULAStatus", clmgmtLicenseInfoEntry.ClmgmtLicenseEULAStatus})
    clmgmtLicenseInfoEntry.EntityData.Leafs.Append("clmgmtLicenseComments", types.YLeaf{"ClmgmtLicenseComments", clmgmtLicenseInfoEntry.ClmgmtLicenseComments})
    clmgmtLicenseInfoEntry.EntityData.Leafs.Append("clmgmtLicenseStatus", types.YLeaf{"ClmgmtLicenseStatus", clmgmtLicenseInfoEntry.ClmgmtLicenseStatus})
    clmgmtLicenseInfoEntry.EntityData.Leafs.Append("clmgmtLicenseStartDate", types.YLeaf{"ClmgmtLicenseStartDate", clmgmtLicenseInfoEntry.ClmgmtLicenseStartDate})
    clmgmtLicenseInfoEntry.EntityData.Leafs.Append("clmgmtLicenseEndDate", types.YLeaf{"ClmgmtLicenseEndDate", clmgmtLicenseInfoEntry.ClmgmtLicenseEndDate})
    clmgmtLicenseInfoEntry.EntityData.Leafs.Append("clmgmtLicensePeriodUsed", types.YLeaf{"ClmgmtLicensePeriodUsed", clmgmtLicenseInfoEntry.ClmgmtLicensePeriodUsed})

    clmgmtLicenseInfoEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "ClmgmtLicenseStoreUsed", "ClmgmtLicenseIndex"}

    return &(clmgmtLicenseInfoEntry.EntityData)
}

// CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseStatus represents                         to use this license.
type CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseStatus string

const (
    CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseStatus_inactive CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseStatus = "inactive"

    CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseStatus_notInUse CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseStatus = "notInUse"

    CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseStatus_inUse CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseStatus = "inUse"

    CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseStatus_expiredInUse CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseStatus = "expiredInUse"

    CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseStatus_expiredNotInUse CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseStatus = "expiredNotInUse"

    CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseStatus_usageCountConsumed CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseStatus = "usageCountConsumed"
)

// CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseType represents                           evaluation mode for a limited time.
type CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseType string

const (
    CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseType_demo CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseType = "demo"

    CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseType_extension CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseType = "extension"

    CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseType_gracePeriod CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseType = "gracePeriod"

    CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseType_permanent CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseType = "permanent"

    CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseType_paidSubscription CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseType = "paidSubscription"

    CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseType_evaluationSubscription CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseType = "evaluationSubscription"

    CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseType_extensionSubscription CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseType = "extensionSubscription"

    CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseType_evalRightToUse CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseType = "evalRightToUse"

    CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseType_rightToUse CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseType = "rightToUse"

    CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseType_permanentRightToUse CISCOLICENSEMGMTMIB_ClmgmtLicenseInfoTable_ClmgmtLicenseInfoEntry_ClmgmtLicenseType = "permanentRightToUse"
)

// CISCOLICENSEMGMTMIB_ClmgmtLicensableFeatureTable
// This table contains list of licensable features in the
// image. All the licensable features will have an entry each
// in this table irrespective of whether they are using any
// licenses currently. Entries in this table are created by
// the agent one for each licensable feature in the image.
// These entries remain in the table permanently and can not
// be deleted. Management application can not create or delete
// entries from this table.
type CISCOLICENSEMGMTMIB_ClmgmtLicensableFeatureTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An entry in clmgmtLicensableFeatureTable. Each entry represents a
    // licensable feature. The type is slice of
    // CISCOLICENSEMGMTMIB_ClmgmtLicensableFeatureTable_ClmgmtLicensableFeatureEntry.
    ClmgmtLicensableFeatureEntry []*CISCOLICENSEMGMTMIB_ClmgmtLicensableFeatureTable_ClmgmtLicensableFeatureEntry
}

func (clmgmtLicensableFeatureTable *CISCOLICENSEMGMTMIB_ClmgmtLicensableFeatureTable) GetEntityData() *types.CommonEntityData {
    clmgmtLicensableFeatureTable.EntityData.YFilter = clmgmtLicensableFeatureTable.YFilter
    clmgmtLicensableFeatureTable.EntityData.YangName = "clmgmtLicensableFeatureTable"
    clmgmtLicensableFeatureTable.EntityData.BundleName = "cisco_ios_xe"
    clmgmtLicensableFeatureTable.EntityData.ParentYangName = "CISCO-LICENSE-MGMT-MIB"
    clmgmtLicensableFeatureTable.EntityData.SegmentPath = "clmgmtLicensableFeatureTable"
    clmgmtLicensableFeatureTable.EntityData.AbsolutePath = "CISCO-LICENSE-MGMT-MIB:CISCO-LICENSE-MGMT-MIB/" + clmgmtLicensableFeatureTable.EntityData.SegmentPath
    clmgmtLicensableFeatureTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtLicensableFeatureTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtLicensableFeatureTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtLicensableFeatureTable.EntityData.Children = types.NewOrderedMap()
    clmgmtLicensableFeatureTable.EntityData.Children.Append("clmgmtLicensableFeatureEntry", types.YChild{"ClmgmtLicensableFeatureEntry", nil})
    for i := range clmgmtLicensableFeatureTable.ClmgmtLicensableFeatureEntry {
        clmgmtLicensableFeatureTable.EntityData.Children.Append(types.GetSegmentPath(clmgmtLicensableFeatureTable.ClmgmtLicensableFeatureEntry[i]), types.YChild{"ClmgmtLicensableFeatureEntry", clmgmtLicensableFeatureTable.ClmgmtLicensableFeatureEntry[i]})
    }
    clmgmtLicensableFeatureTable.EntityData.Leafs = types.NewOrderedMap()

    clmgmtLicensableFeatureTable.EntityData.YListKeys = []string {}

    return &(clmgmtLicensableFeatureTable.EntityData)
}

// CISCOLICENSEMGMTMIB_ClmgmtLicensableFeatureTable_ClmgmtLicensableFeatureEntry
// An entry in clmgmtLicensableFeatureTable. Each entry represents
// a licensable feature.
type CISCOLICENSEMGMTMIB_ClmgmtLicensableFeatureTable_ClmgmtLicensableFeatureEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The type is string with range: 1..2147483647.
    // Refers to
    // entity_mib.ENTITYMIB_EntPhysicalTable_EntPhysicalEntry_EntPhysicalIndex
    EntPhysicalIndex interface{}

    // This attribute is a key. This object uniquely identifies a licensable
    // feature in the device. The type is interface{} with range: 1..4294967295.
    ClmgmtFeatureIndex interface{}

    // This object indicates the name of the licensable feature in the device.
    // Examples of feature names are: 'IPBASE', 'ADVIPSERVICE'. The type is string
    // with length: 0..128.
    ClmgmtFeatureName interface{}

    // This object indicates the version of the licensable feature in the device.
    // Examples of feature versions are: '1.0' or '2.0'. The type is string with
    // length: 0..32.
    ClmgmtFeatureVersion interface{}

    // This object indicates the time period remaining before the feature's
    // license expires or transitions. This object is applicable only if
    // clmgmtLicenseType of the license used by this feature is demo(1), or
    // extension(2) or gracePeriod(3) or evalRightToUse(8).  The object will
    // contain 0 if other types of license is used or if the feature does not use
    // any license. If the feature is using multiple licenses, this period will
    // represent the cumulative period remaining from all the licenses used by
    // this feature. The type is interface{} with range: 0..4294967295. Units are
    // seconds.
    ClmgmtFeatureValidityPeriodRemaining interface{}

    // This object represents the entity that is being counted by this feature.
    // Examples of entities are IP Phones, number of sessions etc. This object is
    // only applicable for features that use counting licenses. For other
    // features, this object will return empty string. The type is string with
    // length: 0..128.
    ClmgmtFeatureWhatIsCounted interface{}

    // This object indicates the license start date of the feature. This object is
    // applicable if at least one of the licenses used for this feature has a
    // valid start date. The start date will be the earliest of the valid start
    // dates of all the licenses used for this feature. If none of the licenses
    // used for this feature have a valid start date then this object will contain
    // an octet string of length 0. The type is string.
    ClmgmtFeatureStartDate interface{}

    // This object indicates the license end date of the feature. This object is
    // applicable if at least one of the licenses used for this feature has a
    // valid end date. The end date will be the latest of the valid end dates of
    // all the licenses used for this feature. If none of the licenses used for
    // this feature have a valid end date then this object will contain an octet
    // string of length 0. The type is string.
    ClmgmtFeatureEndDate interface{}

    // This object indicates the license period used for the feature. The type is
    // interface{} with range: 0..4294967295. Units are seconds.
    ClmgmtFeaturePeriodUsed interface{}
}

func (clmgmtLicensableFeatureEntry *CISCOLICENSEMGMTMIB_ClmgmtLicensableFeatureTable_ClmgmtLicensableFeatureEntry) GetEntityData() *types.CommonEntityData {
    clmgmtLicensableFeatureEntry.EntityData.YFilter = clmgmtLicensableFeatureEntry.YFilter
    clmgmtLicensableFeatureEntry.EntityData.YangName = "clmgmtLicensableFeatureEntry"
    clmgmtLicensableFeatureEntry.EntityData.BundleName = "cisco_ios_xe"
    clmgmtLicensableFeatureEntry.EntityData.ParentYangName = "clmgmtLicensableFeatureTable"
    clmgmtLicensableFeatureEntry.EntityData.SegmentPath = "clmgmtLicensableFeatureEntry" + types.AddKeyToken(clmgmtLicensableFeatureEntry.EntPhysicalIndex, "entPhysicalIndex") + types.AddKeyToken(clmgmtLicensableFeatureEntry.ClmgmtFeatureIndex, "clmgmtFeatureIndex")
    clmgmtLicensableFeatureEntry.EntityData.AbsolutePath = "CISCO-LICENSE-MGMT-MIB:CISCO-LICENSE-MGMT-MIB/clmgmtLicensableFeatureTable/" + clmgmtLicensableFeatureEntry.EntityData.SegmentPath
    clmgmtLicensableFeatureEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtLicensableFeatureEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtLicensableFeatureEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtLicensableFeatureEntry.EntityData.Children = types.NewOrderedMap()
    clmgmtLicensableFeatureEntry.EntityData.Leafs = types.NewOrderedMap()
    clmgmtLicensableFeatureEntry.EntityData.Leafs.Append("entPhysicalIndex", types.YLeaf{"EntPhysicalIndex", clmgmtLicensableFeatureEntry.EntPhysicalIndex})
    clmgmtLicensableFeatureEntry.EntityData.Leafs.Append("clmgmtFeatureIndex", types.YLeaf{"ClmgmtFeatureIndex", clmgmtLicensableFeatureEntry.ClmgmtFeatureIndex})
    clmgmtLicensableFeatureEntry.EntityData.Leafs.Append("clmgmtFeatureName", types.YLeaf{"ClmgmtFeatureName", clmgmtLicensableFeatureEntry.ClmgmtFeatureName})
    clmgmtLicensableFeatureEntry.EntityData.Leafs.Append("clmgmtFeatureVersion", types.YLeaf{"ClmgmtFeatureVersion", clmgmtLicensableFeatureEntry.ClmgmtFeatureVersion})
    clmgmtLicensableFeatureEntry.EntityData.Leafs.Append("clmgmtFeatureValidityPeriodRemaining", types.YLeaf{"ClmgmtFeatureValidityPeriodRemaining", clmgmtLicensableFeatureEntry.ClmgmtFeatureValidityPeriodRemaining})
    clmgmtLicensableFeatureEntry.EntityData.Leafs.Append("clmgmtFeatureWhatIsCounted", types.YLeaf{"ClmgmtFeatureWhatIsCounted", clmgmtLicensableFeatureEntry.ClmgmtFeatureWhatIsCounted})
    clmgmtLicensableFeatureEntry.EntityData.Leafs.Append("clmgmtFeatureStartDate", types.YLeaf{"ClmgmtFeatureStartDate", clmgmtLicensableFeatureEntry.ClmgmtFeatureStartDate})
    clmgmtLicensableFeatureEntry.EntityData.Leafs.Append("clmgmtFeatureEndDate", types.YLeaf{"ClmgmtFeatureEndDate", clmgmtLicensableFeatureEntry.ClmgmtFeatureEndDate})
    clmgmtLicensableFeatureEntry.EntityData.Leafs.Append("clmgmtFeaturePeriodUsed", types.YLeaf{"ClmgmtFeaturePeriodUsed", clmgmtLicensableFeatureEntry.ClmgmtFeaturePeriodUsed})

    clmgmtLicensableFeatureEntry.EntityData.YListKeys = []string {"EntPhysicalIndex", "ClmgmtFeatureIndex"}

    return &(clmgmtLicensableFeatureEntry.EntityData)
}

// CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable
// A table for triggering device credentials export action.
// Management application must create this entry to trigger the
// export of device credentials from the device to a file.
// 
// Once the request completes, the management application should
// retrieve the values of the objects of interest, and then
// delete the entry.  In order to prevent old entries from
// clogging the table, entries will be aged out, but an entry
// will never be deleted within 5 minutes of completion.
type CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable struct {
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
    // CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable_ClmgmtDevCredExportActionEntry.
    ClmgmtDevCredExportActionEntry []*CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable_ClmgmtDevCredExportActionEntry
}

func (clmgmtDevCredExportActionTable *CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable) GetEntityData() *types.CommonEntityData {
    clmgmtDevCredExportActionTable.EntityData.YFilter = clmgmtDevCredExportActionTable.YFilter
    clmgmtDevCredExportActionTable.EntityData.YangName = "clmgmtDevCredExportActionTable"
    clmgmtDevCredExportActionTable.EntityData.BundleName = "cisco_ios_xe"
    clmgmtDevCredExportActionTable.EntityData.ParentYangName = "CISCO-LICENSE-MGMT-MIB"
    clmgmtDevCredExportActionTable.EntityData.SegmentPath = "clmgmtDevCredExportActionTable"
    clmgmtDevCredExportActionTable.EntityData.AbsolutePath = "CISCO-LICENSE-MGMT-MIB:CISCO-LICENSE-MGMT-MIB/" + clmgmtDevCredExportActionTable.EntityData.SegmentPath
    clmgmtDevCredExportActionTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtDevCredExportActionTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtDevCredExportActionTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtDevCredExportActionTable.EntityData.Children = types.NewOrderedMap()
    clmgmtDevCredExportActionTable.EntityData.Children.Append("clmgmtDevCredExportActionEntry", types.YChild{"ClmgmtDevCredExportActionEntry", nil})
    for i := range clmgmtDevCredExportActionTable.ClmgmtDevCredExportActionEntry {
        clmgmtDevCredExportActionTable.EntityData.Children.Append(types.GetSegmentPath(clmgmtDevCredExportActionTable.ClmgmtDevCredExportActionEntry[i]), types.YChild{"ClmgmtDevCredExportActionEntry", clmgmtDevCredExportActionTable.ClmgmtDevCredExportActionEntry[i]})
    }
    clmgmtDevCredExportActionTable.EntityData.Leafs = types.NewOrderedMap()

    clmgmtDevCredExportActionTable.EntityData.YListKeys = []string {}

    return &(clmgmtDevCredExportActionTable.EntityData)
}

// CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable_ClmgmtDevCredExportActionEntry
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
type CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable_ClmgmtDevCredExportActionEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. This object uniquely identifies a row in
    // clmgmtDevCredExportActionTable. The management application chooses this
    // value by reading clmgmtNextFreeDevCredExportActionIndex while creating an
    // entry in this table. If an entry already exists with this index, the
    // creation of the entry will not continue and error will be returned. The
    // management application should read the value of
    // clmgmtNextFreeDevCredExportActionIndex again and retry with the new value
    // for this object. The type is interface{} with range: 1..4294967295.
    ClmgmtDevCredExportActionIndex interface{}

    // This object represents the entPhysicalIndex of the device for which the
    // device credentials are being retrieved. This object is mainly used in
    // devices where one device is acting as a master and rest of the devices as
    // slaves. The master device is responsible for SNMP communication with the
    // manager. Examples include stackable switches, devices with router processor
    // and line cards.  Note: This object need not be set if it is a stand alone
    // device. The type is interface{} with range: 0..2147483647.
    ClmgmtDevCredEntPhysicalIndex interface{}

    // This object indicates the transfer protocol to be used when copying files
    // as specified in the following objects. 1. clmgmtDevCredExportFile . The
    // type is ClmgmtLicenseTransferProtocol.
    ClmgmtDevCredTransferProtocol interface{}

    // This object indicates the transport type of the address contained in
    // clmgmtDevCredServerAddress object. This must be set when
    // clmgmtDevCredTransferProtocol is not none(1) or local(2). The type is
    // InetAddressType.
    ClmgmtDevCredServerAddressType interface{}

    // This object indicates the the ip address of the server from which the files
    // must be read or written to if  clmgmtDevCredTransferProtocol is not none(1)
    // or local(2).  All bits as 0s or 1s for clmgmtDevCredServerAddress are not
    // allowed.  The format of this address depends on the value of the
    // clmgmtDevCredServerAddressType object. The type is string with length:
    // 0..255.
    ClmgmtDevCredServerAddress interface{}

    // This object indicates the remote user name for accessing files via ftp,
    // rcp, sftp or scp protocols. This object must be set when the
    // clmgmtDevCredTransferProtocol is ftp(4), rcp(5), scp(7) or sftp(8). If
    // clmgmtDevCredTransferProtocol is rcp(5), the remote username is sent as the
    // server username in an rcp command request sent by the system to a remote
    // rcp server. The type is string with length: 0..96.
    ClmgmtDevCredServerUsername interface{}

    // This object indicates the password used by ftp, sftp or scp for copying a
    // file to/from an ftp/sftp/scp server.  This object must be set when the
    // clmgmtDevCredTransferProtocol is ftp(4) or scp(7) or sftp(8). Reading it
    // returns a zero-length string for  security reasons. The type is string with
    // length: 0..96.
    ClmgmtDevCredServerPassword interface{}

    // This object represents file where device credentials needs to be exported
    // to. The type is string with length: 0..255.
    ClmgmtDevCredExportFile interface{}

    // This object indicates the the command to be executed.  Command             
    // Remarks -------                          ------- noOp(1)                   
    // No operation will be                                 performed. 
    // getDeviceCredentials(2)         Exports device credentials. The type is
    // ClmgmtDevCredCommand.
    ClmgmtDevCredCommand interface{}

    // This object indicates the state of the action that is executed as a result
    // of setting clmgmtDevCredRowStatus to active(1). The type is
    // ClmgmtLicenseActionState.
    ClmgmtDevCredCommandState interface{}

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
    // ClmgmtDevCredCommandFailCause.
    ClmgmtDevCredCommandFailCause interface{}

    // This object indicates the storage type for this conceptual row. Conceptual
    // rows having the value 'permanent' need not allow write-access to any
    // columnar objects in the row. The type is StorageType.
    ClmgmtDevCredStorageType interface{}

    // This object indicates the the status of this table entry. Once the entry
    // status is set to active(1), the associated entry cannot be modified until
    // the action completes (clmgmtDevCredCommandStatus is set to a value other
    // than inProgress(3)). Once the action completes the only operation possible
    // after this is to delete the row.  clmgmtDevCredExportFile is a mandatory
    // object to be set when creating this entry. The type is RowStatus.
    ClmgmtDevCredRowStatus interface{}
}

func (clmgmtDevCredExportActionEntry *CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable_ClmgmtDevCredExportActionEntry) GetEntityData() *types.CommonEntityData {
    clmgmtDevCredExportActionEntry.EntityData.YFilter = clmgmtDevCredExportActionEntry.YFilter
    clmgmtDevCredExportActionEntry.EntityData.YangName = "clmgmtDevCredExportActionEntry"
    clmgmtDevCredExportActionEntry.EntityData.BundleName = "cisco_ios_xe"
    clmgmtDevCredExportActionEntry.EntityData.ParentYangName = "clmgmtDevCredExportActionTable"
    clmgmtDevCredExportActionEntry.EntityData.SegmentPath = "clmgmtDevCredExportActionEntry" + types.AddKeyToken(clmgmtDevCredExportActionEntry.ClmgmtDevCredExportActionIndex, "clmgmtDevCredExportActionIndex")
    clmgmtDevCredExportActionEntry.EntityData.AbsolutePath = "CISCO-LICENSE-MGMT-MIB:CISCO-LICENSE-MGMT-MIB/clmgmtDevCredExportActionTable/" + clmgmtDevCredExportActionEntry.EntityData.SegmentPath
    clmgmtDevCredExportActionEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    clmgmtDevCredExportActionEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    clmgmtDevCredExportActionEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    clmgmtDevCredExportActionEntry.EntityData.Children = types.NewOrderedMap()
    clmgmtDevCredExportActionEntry.EntityData.Leafs = types.NewOrderedMap()
    clmgmtDevCredExportActionEntry.EntityData.Leafs.Append("clmgmtDevCredExportActionIndex", types.YLeaf{"ClmgmtDevCredExportActionIndex", clmgmtDevCredExportActionEntry.ClmgmtDevCredExportActionIndex})
    clmgmtDevCredExportActionEntry.EntityData.Leafs.Append("clmgmtDevCredEntPhysicalIndex", types.YLeaf{"ClmgmtDevCredEntPhysicalIndex", clmgmtDevCredExportActionEntry.ClmgmtDevCredEntPhysicalIndex})
    clmgmtDevCredExportActionEntry.EntityData.Leafs.Append("clmgmtDevCredTransferProtocol", types.YLeaf{"ClmgmtDevCredTransferProtocol", clmgmtDevCredExportActionEntry.ClmgmtDevCredTransferProtocol})
    clmgmtDevCredExportActionEntry.EntityData.Leafs.Append("clmgmtDevCredServerAddressType", types.YLeaf{"ClmgmtDevCredServerAddressType", clmgmtDevCredExportActionEntry.ClmgmtDevCredServerAddressType})
    clmgmtDevCredExportActionEntry.EntityData.Leafs.Append("clmgmtDevCredServerAddress", types.YLeaf{"ClmgmtDevCredServerAddress", clmgmtDevCredExportActionEntry.ClmgmtDevCredServerAddress})
    clmgmtDevCredExportActionEntry.EntityData.Leafs.Append("clmgmtDevCredServerUsername", types.YLeaf{"ClmgmtDevCredServerUsername", clmgmtDevCredExportActionEntry.ClmgmtDevCredServerUsername})
    clmgmtDevCredExportActionEntry.EntityData.Leafs.Append("clmgmtDevCredServerPassword", types.YLeaf{"ClmgmtDevCredServerPassword", clmgmtDevCredExportActionEntry.ClmgmtDevCredServerPassword})
    clmgmtDevCredExportActionEntry.EntityData.Leafs.Append("clmgmtDevCredExportFile", types.YLeaf{"ClmgmtDevCredExportFile", clmgmtDevCredExportActionEntry.ClmgmtDevCredExportFile})
    clmgmtDevCredExportActionEntry.EntityData.Leafs.Append("clmgmtDevCredCommand", types.YLeaf{"ClmgmtDevCredCommand", clmgmtDevCredExportActionEntry.ClmgmtDevCredCommand})
    clmgmtDevCredExportActionEntry.EntityData.Leafs.Append("clmgmtDevCredCommandState", types.YLeaf{"ClmgmtDevCredCommandState", clmgmtDevCredExportActionEntry.ClmgmtDevCredCommandState})
    clmgmtDevCredExportActionEntry.EntityData.Leafs.Append("clmgmtDevCredCommandFailCause", types.YLeaf{"ClmgmtDevCredCommandFailCause", clmgmtDevCredExportActionEntry.ClmgmtDevCredCommandFailCause})
    clmgmtDevCredExportActionEntry.EntityData.Leafs.Append("clmgmtDevCredStorageType", types.YLeaf{"ClmgmtDevCredStorageType", clmgmtDevCredExportActionEntry.ClmgmtDevCredStorageType})
    clmgmtDevCredExportActionEntry.EntityData.Leafs.Append("clmgmtDevCredRowStatus", types.YLeaf{"ClmgmtDevCredRowStatus", clmgmtDevCredExportActionEntry.ClmgmtDevCredRowStatus})

    clmgmtDevCredExportActionEntry.EntityData.YListKeys = []string {"ClmgmtDevCredExportActionIndex"}

    return &(clmgmtDevCredExportActionEntry.EntityData)
}

// CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable_ClmgmtDevCredExportActionEntry_ClmgmtDevCredCommand represents getDeviceCredentials(2)         Exports device credentials
type CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable_ClmgmtDevCredExportActionEntry_ClmgmtDevCredCommand string

const (
    CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable_ClmgmtDevCredExportActionEntry_ClmgmtDevCredCommand_noOp CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable_ClmgmtDevCredExportActionEntry_ClmgmtDevCredCommand = "noOp"

    CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable_ClmgmtDevCredExportActionEntry_ClmgmtDevCredCommand_getDeviceCredentials CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable_ClmgmtDevCredExportActionEntry_ClmgmtDevCredCommand = "getDeviceCredentials"
)

// CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable_ClmgmtDevCredExportActionEntry_ClmgmtDevCredCommandFailCause represents invalidFile(6)  - The target file specified is not valid.
type CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable_ClmgmtDevCredExportActionEntry_ClmgmtDevCredCommandFailCause string

const (
    CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable_ClmgmtDevCredExportActionEntry_ClmgmtDevCredCommandFailCause_none CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable_ClmgmtDevCredExportActionEntry_ClmgmtDevCredCommandFailCause = "none"

    CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable_ClmgmtDevCredExportActionEntry_ClmgmtDevCredCommandFailCause_unknownError CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable_ClmgmtDevCredExportActionEntry_ClmgmtDevCredCommandFailCause = "unknownError"

    CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable_ClmgmtDevCredExportActionEntry_ClmgmtDevCredCommandFailCause_transferProtocolNotSupported CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable_ClmgmtDevCredExportActionEntry_ClmgmtDevCredCommandFailCause = "transferProtocolNotSupported"

    CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable_ClmgmtDevCredExportActionEntry_ClmgmtDevCredCommandFailCause_fileServerNotReachable CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable_ClmgmtDevCredExportActionEntry_ClmgmtDevCredCommandFailCause = "fileServerNotReachable"

    CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable_ClmgmtDevCredExportActionEntry_ClmgmtDevCredCommandFailCause_unrecognizedEntPhysicalIndex CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable_ClmgmtDevCredExportActionEntry_ClmgmtDevCredCommandFailCause = "unrecognizedEntPhysicalIndex"

    CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable_ClmgmtDevCredExportActionEntry_ClmgmtDevCredCommandFailCause_invalidFile CISCOLICENSEMGMTMIB_ClmgmtDevCredExportActionTable_ClmgmtDevCredExportActionEntry_ClmgmtDevCredCommandFailCause = "invalidFile"
)

