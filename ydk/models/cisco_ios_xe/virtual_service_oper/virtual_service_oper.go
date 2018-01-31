// This module contains a collection of YANG definitions for monitoring
// virtual services in a Network Element.Copyright (c) 2016-2017 by Cisco Systems, Inc.All rights reserved.
package virtual_service_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package virtual_service_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XE-virtual-service-oper virtual-services}", reflect.TypeOf(VirtualServices{}))
    ydk.RegisterEntity("Cisco-IOS-XE-virtual-service-oper:virtual-services", reflect.TypeOf(VirtualServices{}))
}

// VirtualServices
// Names and Status of virtual services on the device
type VirtualServices struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // A virtual service. The type is slice of VirtualServices_VirtualService.
    VirtualService []VirtualServices_VirtualService
}

func (virtualServices *VirtualServices) GetFilter() yfilter.YFilter { return virtualServices.YFilter }

func (virtualServices *VirtualServices) SetFilter(yf yfilter.YFilter) { virtualServices.YFilter = yf }

func (virtualServices *VirtualServices) GetGoName(yname string) string {
    if yname == "virtual-service" { return "VirtualService" }
    return ""
}

func (virtualServices *VirtualServices) GetSegmentPath() string {
    return "Cisco-IOS-XE-virtual-service-oper:virtual-services"
}

func (virtualServices *VirtualServices) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "virtual-service" {
        for _, c := range virtualServices.VirtualService {
            if virtualServices.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := VirtualServices_VirtualService{}
        virtualServices.VirtualService = append(virtualServices.VirtualService, child)
        return &virtualServices.VirtualService[len(virtualServices.VirtualService)-1]
    }
    return nil
}

func (virtualServices *VirtualServices) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range virtualServices.VirtualService {
        children[virtualServices.VirtualService[i].GetSegmentPath()] = &virtualServices.VirtualService[i]
    }
    return children
}

func (virtualServices *VirtualServices) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (virtualServices *VirtualServices) GetBundleName() string { return "cisco_ios_xe" }

func (virtualServices *VirtualServices) GetYangName() string { return "virtual-services" }

func (virtualServices *VirtualServices) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (virtualServices *VirtualServices) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (virtualServices *VirtualServices) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (virtualServices *VirtualServices) SetParent(parent types.Entity) { virtualServices.parent = parent }

func (virtualServices *VirtualServices) GetParent() types.Entity { return virtualServices.parent }

func (virtualServices *VirtualServices) GetParentYangName() string { return "Cisco-IOS-XE-virtual-service-oper" }

// VirtualServices_VirtualService
// A virtual service.
type VirtualServices_VirtualService struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the virtual service. The type is
    // string.
    Name interface{}

    // Details of the virtual service.
    Details VirtualServices_VirtualService_Details

    // Utilization of device resources for a virtual-service.
    Utilization VirtualServices_VirtualService_Utilization

    // list of the network utilizations for the virtual-service.
    NetworkUtils VirtualServices_VirtualService_NetworkUtils

    // List of the storage utilizations for the virtual-service.
    StorageUtils VirtualServices_VirtualService_StorageUtils

    // Details for the devices attached to this virtual service.
    AttachedDevices VirtualServices_VirtualService_AttachedDevices

    // Details for the network interfaces.
    NetworkInterfaces VirtualServices_VirtualService_NetworkInterfaces

    // Routes for the guest interface.
    GuestRoutes VirtualServices_VirtualService_GuestRoutes
}

func (virtualService *VirtualServices_VirtualService) GetFilter() yfilter.YFilter { return virtualService.YFilter }

func (virtualService *VirtualServices_VirtualService) SetFilter(yf yfilter.YFilter) { virtualService.YFilter = yf }

func (virtualService *VirtualServices_VirtualService) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "details" { return "Details" }
    if yname == "utilization" { return "Utilization" }
    if yname == "network-utils" { return "NetworkUtils" }
    if yname == "storage-utils" { return "StorageUtils" }
    if yname == "attached-devices" { return "AttachedDevices" }
    if yname == "network-interfaces" { return "NetworkInterfaces" }
    if yname == "guest-routes" { return "GuestRoutes" }
    return ""
}

func (virtualService *VirtualServices_VirtualService) GetSegmentPath() string {
    return "virtual-service" + "[name='" + fmt.Sprintf("%v", virtualService.Name) + "']"
}

func (virtualService *VirtualServices_VirtualService) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "details" {
        return &virtualService.Details
    }
    if childYangName == "utilization" {
        return &virtualService.Utilization
    }
    if childYangName == "network-utils" {
        return &virtualService.NetworkUtils
    }
    if childYangName == "storage-utils" {
        return &virtualService.StorageUtils
    }
    if childYangName == "attached-devices" {
        return &virtualService.AttachedDevices
    }
    if childYangName == "network-interfaces" {
        return &virtualService.NetworkInterfaces
    }
    if childYangName == "guest-routes" {
        return &virtualService.GuestRoutes
    }
    return nil
}

func (virtualService *VirtualServices_VirtualService) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["details"] = &virtualService.Details
    children["utilization"] = &virtualService.Utilization
    children["network-utils"] = &virtualService.NetworkUtils
    children["storage-utils"] = &virtualService.StorageUtils
    children["attached-devices"] = &virtualService.AttachedDevices
    children["network-interfaces"] = &virtualService.NetworkInterfaces
    children["guest-routes"] = &virtualService.GuestRoutes
    return children
}

func (virtualService *VirtualServices_VirtualService) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = virtualService.Name
    return leafs
}

func (virtualService *VirtualServices_VirtualService) GetBundleName() string { return "cisco_ios_xe" }

func (virtualService *VirtualServices_VirtualService) GetYangName() string { return "virtual-service" }

func (virtualService *VirtualServices_VirtualService) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (virtualService *VirtualServices_VirtualService) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (virtualService *VirtualServices_VirtualService) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (virtualService *VirtualServices_VirtualService) SetParent(parent types.Entity) { virtualService.parent = parent }

func (virtualService *VirtualServices_VirtualService) GetParent() types.Entity { return virtualService.parent }

func (virtualService *VirtualServices_VirtualService) GetParentYangName() string { return "virtual-services" }

// VirtualServices_VirtualService_Details
// Details of the virtual service.
type VirtualServices_VirtualService_Details struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // State of the virtual service. The type is string.
    State interface{}

    // The activated profile name. The type is string.
    ActivatedProfileName interface{}

    // The name of a guest interface. The type is string.
    GuestInterface interface{}

    // Details of the package for the virtual-service.
    PackageInformation VirtualServices_VirtualService_Details_PackageInformation

    // Details on the guest status.
    DetailedGuestStatus VirtualServices_VirtualService_Details_DetailedGuestStatus

    // Details of the resources reserved for this virtual service.
    ResourceReservation VirtualServices_VirtualService_Details_ResourceReservation

    // Resources being allocated for the virtual-service.
    ResourceAdmission VirtualServices_VirtualService_Details_ResourceAdmission
}

func (details *VirtualServices_VirtualService_Details) GetFilter() yfilter.YFilter { return details.YFilter }

func (details *VirtualServices_VirtualService_Details) SetFilter(yf yfilter.YFilter) { details.YFilter = yf }

func (details *VirtualServices_VirtualService_Details) GetGoName(yname string) string {
    if yname == "state" { return "State" }
    if yname == "activated-profile-name" { return "ActivatedProfileName" }
    if yname == "guest-interface" { return "GuestInterface" }
    if yname == "package-information" { return "PackageInformation" }
    if yname == "detailed-guest-status" { return "DetailedGuestStatus" }
    if yname == "resource-reservation" { return "ResourceReservation" }
    if yname == "resource-admission" { return "ResourceAdmission" }
    return ""
}

func (details *VirtualServices_VirtualService_Details) GetSegmentPath() string {
    return "details"
}

func (details *VirtualServices_VirtualService_Details) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "package-information" {
        return &details.PackageInformation
    }
    if childYangName == "detailed-guest-status" {
        return &details.DetailedGuestStatus
    }
    if childYangName == "resource-reservation" {
        return &details.ResourceReservation
    }
    if childYangName == "resource-admission" {
        return &details.ResourceAdmission
    }
    return nil
}

func (details *VirtualServices_VirtualService_Details) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["package-information"] = &details.PackageInformation
    children["detailed-guest-status"] = &details.DetailedGuestStatus
    children["resource-reservation"] = &details.ResourceReservation
    children["resource-admission"] = &details.ResourceAdmission
    return children
}

func (details *VirtualServices_VirtualService_Details) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["state"] = details.State
    leafs["activated-profile-name"] = details.ActivatedProfileName
    leafs["guest-interface"] = details.GuestInterface
    return leafs
}

func (details *VirtualServices_VirtualService_Details) GetBundleName() string { return "cisco_ios_xe" }

func (details *VirtualServices_VirtualService_Details) GetYangName() string { return "details" }

func (details *VirtualServices_VirtualService_Details) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (details *VirtualServices_VirtualService_Details) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (details *VirtualServices_VirtualService_Details) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (details *VirtualServices_VirtualService_Details) SetParent(parent types.Entity) { details.parent = parent }

func (details *VirtualServices_VirtualService_Details) GetParent() types.Entity { return details.parent }

func (details *VirtualServices_VirtualService_Details) GetParentYangName() string { return "virtual-service" }

// VirtualServices_VirtualService_Details_PackageInformation
// Details of the package for the virtual-service.
type VirtualServices_VirtualService_Details_PackageInformation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the package for the virtual-service. The type is string.
    Name interface{}

    // Path to the package. The type is string.
    Path interface{}

    // Details of the application.
    Application VirtualServices_VirtualService_Details_PackageInformation_Application

    // Details of the key signing.
    Signing VirtualServices_VirtualService_Details_PackageInformation_Signing

    // Details about the license.
    Licensing VirtualServices_VirtualService_Details_PackageInformation_Licensing
}

func (packageInformation *VirtualServices_VirtualService_Details_PackageInformation) GetFilter() yfilter.YFilter { return packageInformation.YFilter }

func (packageInformation *VirtualServices_VirtualService_Details_PackageInformation) SetFilter(yf yfilter.YFilter) { packageInformation.YFilter = yf }

func (packageInformation *VirtualServices_VirtualService_Details_PackageInformation) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "path" { return "Path" }
    if yname == "application" { return "Application" }
    if yname == "signing" { return "Signing" }
    if yname == "licensing" { return "Licensing" }
    return ""
}

func (packageInformation *VirtualServices_VirtualService_Details_PackageInformation) GetSegmentPath() string {
    return "package-information"
}

func (packageInformation *VirtualServices_VirtualService_Details_PackageInformation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "application" {
        return &packageInformation.Application
    }
    if childYangName == "signing" {
        return &packageInformation.Signing
    }
    if childYangName == "licensing" {
        return &packageInformation.Licensing
    }
    return nil
}

func (packageInformation *VirtualServices_VirtualService_Details_PackageInformation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["application"] = &packageInformation.Application
    children["signing"] = &packageInformation.Signing
    children["licensing"] = &packageInformation.Licensing
    return children
}

func (packageInformation *VirtualServices_VirtualService_Details_PackageInformation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = packageInformation.Name
    leafs["path"] = packageInformation.Path
    return leafs
}

func (packageInformation *VirtualServices_VirtualService_Details_PackageInformation) GetBundleName() string { return "cisco_ios_xe" }

func (packageInformation *VirtualServices_VirtualService_Details_PackageInformation) GetYangName() string { return "package-information" }

func (packageInformation *VirtualServices_VirtualService_Details_PackageInformation) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (packageInformation *VirtualServices_VirtualService_Details_PackageInformation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (packageInformation *VirtualServices_VirtualService_Details_PackageInformation) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (packageInformation *VirtualServices_VirtualService_Details_PackageInformation) SetParent(parent types.Entity) { packageInformation.parent = parent }

func (packageInformation *VirtualServices_VirtualService_Details_PackageInformation) GetParent() types.Entity { return packageInformation.parent }

func (packageInformation *VirtualServices_VirtualService_Details_PackageInformation) GetParentYangName() string { return "details" }

// VirtualServices_VirtualService_Details_PackageInformation_Application
// Details of the application.
type VirtualServices_VirtualService_Details_PackageInformation_Application struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the application. The type is string.
    Name interface{}

    // Version of the application. The type is string.
    InstalledVersion interface{}

    // Description of the application. The type is string.
    Description interface{}
}

func (application *VirtualServices_VirtualService_Details_PackageInformation_Application) GetFilter() yfilter.YFilter { return application.YFilter }

func (application *VirtualServices_VirtualService_Details_PackageInformation_Application) SetFilter(yf yfilter.YFilter) { application.YFilter = yf }

func (application *VirtualServices_VirtualService_Details_PackageInformation_Application) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "installed-version" { return "InstalledVersion" }
    if yname == "description" { return "Description" }
    return ""
}

func (application *VirtualServices_VirtualService_Details_PackageInformation_Application) GetSegmentPath() string {
    return "application"
}

func (application *VirtualServices_VirtualService_Details_PackageInformation_Application) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (application *VirtualServices_VirtualService_Details_PackageInformation_Application) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (application *VirtualServices_VirtualService_Details_PackageInformation_Application) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = application.Name
    leafs["installed-version"] = application.InstalledVersion
    leafs["description"] = application.Description
    return leafs
}

func (application *VirtualServices_VirtualService_Details_PackageInformation_Application) GetBundleName() string { return "cisco_ios_xe" }

func (application *VirtualServices_VirtualService_Details_PackageInformation_Application) GetYangName() string { return "application" }

func (application *VirtualServices_VirtualService_Details_PackageInformation_Application) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (application *VirtualServices_VirtualService_Details_PackageInformation_Application) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (application *VirtualServices_VirtualService_Details_PackageInformation_Application) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (application *VirtualServices_VirtualService_Details_PackageInformation_Application) SetParent(parent types.Entity) { application.parent = parent }

func (application *VirtualServices_VirtualService_Details_PackageInformation_Application) GetParent() types.Entity { return application.parent }

func (application *VirtualServices_VirtualService_Details_PackageInformation_Application) GetParentYangName() string { return "package-information" }

// VirtualServices_VirtualService_Details_PackageInformation_Signing
// Details of the key signing.
type VirtualServices_VirtualService_Details_PackageInformation_Signing struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The Type of the signed key. The type is string.
    KeyType interface{}

    // The method the key was signed. The type is string.
    Method interface{}
}

func (signing *VirtualServices_VirtualService_Details_PackageInformation_Signing) GetFilter() yfilter.YFilter { return signing.YFilter }

func (signing *VirtualServices_VirtualService_Details_PackageInformation_Signing) SetFilter(yf yfilter.YFilter) { signing.YFilter = yf }

func (signing *VirtualServices_VirtualService_Details_PackageInformation_Signing) GetGoName(yname string) string {
    if yname == "key-type" { return "KeyType" }
    if yname == "method" { return "Method" }
    return ""
}

func (signing *VirtualServices_VirtualService_Details_PackageInformation_Signing) GetSegmentPath() string {
    return "signing"
}

func (signing *VirtualServices_VirtualService_Details_PackageInformation_Signing) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (signing *VirtualServices_VirtualService_Details_PackageInformation_Signing) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (signing *VirtualServices_VirtualService_Details_PackageInformation_Signing) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["key-type"] = signing.KeyType
    leafs["method"] = signing.Method
    return leafs
}

func (signing *VirtualServices_VirtualService_Details_PackageInformation_Signing) GetBundleName() string { return "cisco_ios_xe" }

func (signing *VirtualServices_VirtualService_Details_PackageInformation_Signing) GetYangName() string { return "signing" }

func (signing *VirtualServices_VirtualService_Details_PackageInformation_Signing) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (signing *VirtualServices_VirtualService_Details_PackageInformation_Signing) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (signing *VirtualServices_VirtualService_Details_PackageInformation_Signing) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (signing *VirtualServices_VirtualService_Details_PackageInformation_Signing) SetParent(parent types.Entity) { signing.parent = parent }

func (signing *VirtualServices_VirtualService_Details_PackageInformation_Signing) GetParent() types.Entity { return signing.parent }

func (signing *VirtualServices_VirtualService_Details_PackageInformation_Signing) GetParentYangName() string { return "package-information" }

// VirtualServices_VirtualService_Details_PackageInformation_Licensing
// Details about the license.
type VirtualServices_VirtualService_Details_PackageInformation_Licensing struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The name of the license. The type is string.
    Name interface{}

    // The version of the license. The type is string.
    Version interface{}
}

func (licensing *VirtualServices_VirtualService_Details_PackageInformation_Licensing) GetFilter() yfilter.YFilter { return licensing.YFilter }

func (licensing *VirtualServices_VirtualService_Details_PackageInformation_Licensing) SetFilter(yf yfilter.YFilter) { licensing.YFilter = yf }

func (licensing *VirtualServices_VirtualService_Details_PackageInformation_Licensing) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "version" { return "Version" }
    return ""
}

func (licensing *VirtualServices_VirtualService_Details_PackageInformation_Licensing) GetSegmentPath() string {
    return "licensing"
}

func (licensing *VirtualServices_VirtualService_Details_PackageInformation_Licensing) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (licensing *VirtualServices_VirtualService_Details_PackageInformation_Licensing) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (licensing *VirtualServices_VirtualService_Details_PackageInformation_Licensing) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = licensing.Name
    leafs["version"] = licensing.Version
    return leafs
}

func (licensing *VirtualServices_VirtualService_Details_PackageInformation_Licensing) GetBundleName() string { return "cisco_ios_xe" }

func (licensing *VirtualServices_VirtualService_Details_PackageInformation_Licensing) GetYangName() string { return "licensing" }

func (licensing *VirtualServices_VirtualService_Details_PackageInformation_Licensing) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (licensing *VirtualServices_VirtualService_Details_PackageInformation_Licensing) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (licensing *VirtualServices_VirtualService_Details_PackageInformation_Licensing) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (licensing *VirtualServices_VirtualService_Details_PackageInformation_Licensing) SetParent(parent types.Entity) { licensing.parent = parent }

func (licensing *VirtualServices_VirtualService_Details_PackageInformation_Licensing) GetParent() types.Entity { return licensing.parent }

func (licensing *VirtualServices_VirtualService_Details_PackageInformation_Licensing) GetParentYangName() string { return "package-information" }

// VirtualServices_VirtualService_Details_DetailedGuestStatus
// Details on the guest status.
type VirtualServices_VirtualService_Details_DetailedGuestStatus struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // All the processes.
    Processes VirtualServices_VirtualService_Details_DetailedGuestStatus_Processes
}

func (detailedGuestStatus *VirtualServices_VirtualService_Details_DetailedGuestStatus) GetFilter() yfilter.YFilter { return detailedGuestStatus.YFilter }

func (detailedGuestStatus *VirtualServices_VirtualService_Details_DetailedGuestStatus) SetFilter(yf yfilter.YFilter) { detailedGuestStatus.YFilter = yf }

func (detailedGuestStatus *VirtualServices_VirtualService_Details_DetailedGuestStatus) GetGoName(yname string) string {
    if yname == "processes" { return "Processes" }
    return ""
}

func (detailedGuestStatus *VirtualServices_VirtualService_Details_DetailedGuestStatus) GetSegmentPath() string {
    return "detailed-guest-status"
}

func (detailedGuestStatus *VirtualServices_VirtualService_Details_DetailedGuestStatus) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "processes" {
        return &detailedGuestStatus.Processes
    }
    return nil
}

func (detailedGuestStatus *VirtualServices_VirtualService_Details_DetailedGuestStatus) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["processes"] = &detailedGuestStatus.Processes
    return children
}

func (detailedGuestStatus *VirtualServices_VirtualService_Details_DetailedGuestStatus) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (detailedGuestStatus *VirtualServices_VirtualService_Details_DetailedGuestStatus) GetBundleName() string { return "cisco_ios_xe" }

func (detailedGuestStatus *VirtualServices_VirtualService_Details_DetailedGuestStatus) GetYangName() string { return "detailed-guest-status" }

func (detailedGuestStatus *VirtualServices_VirtualService_Details_DetailedGuestStatus) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (detailedGuestStatus *VirtualServices_VirtualService_Details_DetailedGuestStatus) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (detailedGuestStatus *VirtualServices_VirtualService_Details_DetailedGuestStatus) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (detailedGuestStatus *VirtualServices_VirtualService_Details_DetailedGuestStatus) SetParent(parent types.Entity) { detailedGuestStatus.parent = parent }

func (detailedGuestStatus *VirtualServices_VirtualService_Details_DetailedGuestStatus) GetParent() types.Entity { return detailedGuestStatus.parent }

func (detailedGuestStatus *VirtualServices_VirtualService_Details_DetailedGuestStatus) GetParentYangName() string { return "details" }

// VirtualServices_VirtualService_Details_DetailedGuestStatus_Processes
// All the processes.
type VirtualServices_VirtualService_Details_DetailedGuestStatus_Processes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Name of the proccess. The type is string.
    Name interface{}

    // Status of the proccess. The type is string.
    Status interface{}

    // Process ID. The type is string.
    Pid interface{}

    // Up time of the proccess. The type is string.
    Uptime interface{}

    // Memory of the proccess. The type is string.
    Memory interface{}
}

func (processes *VirtualServices_VirtualService_Details_DetailedGuestStatus_Processes) GetFilter() yfilter.YFilter { return processes.YFilter }

func (processes *VirtualServices_VirtualService_Details_DetailedGuestStatus_Processes) SetFilter(yf yfilter.YFilter) { processes.YFilter = yf }

func (processes *VirtualServices_VirtualService_Details_DetailedGuestStatus_Processes) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "status" { return "Status" }
    if yname == "pid" { return "Pid" }
    if yname == "uptime" { return "Uptime" }
    if yname == "memory" { return "Memory" }
    return ""
}

func (processes *VirtualServices_VirtualService_Details_DetailedGuestStatus_Processes) GetSegmentPath() string {
    return "processes"
}

func (processes *VirtualServices_VirtualService_Details_DetailedGuestStatus_Processes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (processes *VirtualServices_VirtualService_Details_DetailedGuestStatus_Processes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (processes *VirtualServices_VirtualService_Details_DetailedGuestStatus_Processes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = processes.Name
    leafs["status"] = processes.Status
    leafs["pid"] = processes.Pid
    leafs["uptime"] = processes.Uptime
    leafs["memory"] = processes.Memory
    return leafs
}

func (processes *VirtualServices_VirtualService_Details_DetailedGuestStatus_Processes) GetBundleName() string { return "cisco_ios_xe" }

func (processes *VirtualServices_VirtualService_Details_DetailedGuestStatus_Processes) GetYangName() string { return "processes" }

func (processes *VirtualServices_VirtualService_Details_DetailedGuestStatus_Processes) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (processes *VirtualServices_VirtualService_Details_DetailedGuestStatus_Processes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (processes *VirtualServices_VirtualService_Details_DetailedGuestStatus_Processes) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (processes *VirtualServices_VirtualService_Details_DetailedGuestStatus_Processes) SetParent(parent types.Entity) { processes.parent = parent }

func (processes *VirtualServices_VirtualService_Details_DetailedGuestStatus_Processes) GetParent() types.Entity { return processes.parent }

func (processes *VirtualServices_VirtualService_Details_DetailedGuestStatus_Processes) GetParentYangName() string { return "detailed-guest-status" }

// VirtualServices_VirtualService_Details_ResourceReservation
// Details of the resources reserved for this virtual service.
type VirtualServices_VirtualService_Details_ResourceReservation struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The amount of reserverd disk space in MB. The type is interface{} with
    // range: 0..18446744073709551615.
    Disk interface{}

    // The amount of reserved memory in MB. The type is interface{} with range:
    // 0..18446744073709551615.
    Memory interface{}

    // The percentage of reserved cpu. The type is interface{} with range:
    // 0..18446744073709551615.
    Cpu interface{}
}

func (resourceReservation *VirtualServices_VirtualService_Details_ResourceReservation) GetFilter() yfilter.YFilter { return resourceReservation.YFilter }

func (resourceReservation *VirtualServices_VirtualService_Details_ResourceReservation) SetFilter(yf yfilter.YFilter) { resourceReservation.YFilter = yf }

func (resourceReservation *VirtualServices_VirtualService_Details_ResourceReservation) GetGoName(yname string) string {
    if yname == "disk" { return "Disk" }
    if yname == "memory" { return "Memory" }
    if yname == "cpu" { return "Cpu" }
    return ""
}

func (resourceReservation *VirtualServices_VirtualService_Details_ResourceReservation) GetSegmentPath() string {
    return "resource-reservation"
}

func (resourceReservation *VirtualServices_VirtualService_Details_ResourceReservation) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (resourceReservation *VirtualServices_VirtualService_Details_ResourceReservation) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (resourceReservation *VirtualServices_VirtualService_Details_ResourceReservation) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["disk"] = resourceReservation.Disk
    leafs["memory"] = resourceReservation.Memory
    leafs["cpu"] = resourceReservation.Cpu
    return leafs
}

func (resourceReservation *VirtualServices_VirtualService_Details_ResourceReservation) GetBundleName() string { return "cisco_ios_xe" }

func (resourceReservation *VirtualServices_VirtualService_Details_ResourceReservation) GetYangName() string { return "resource-reservation" }

func (resourceReservation *VirtualServices_VirtualService_Details_ResourceReservation) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (resourceReservation *VirtualServices_VirtualService_Details_ResourceReservation) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (resourceReservation *VirtualServices_VirtualService_Details_ResourceReservation) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (resourceReservation *VirtualServices_VirtualService_Details_ResourceReservation) SetParent(parent types.Entity) { resourceReservation.parent = parent }

func (resourceReservation *VirtualServices_VirtualService_Details_ResourceReservation) GetParent() types.Entity { return resourceReservation.parent }

func (resourceReservation *VirtualServices_VirtualService_Details_ResourceReservation) GetParentYangName() string { return "details" }

// VirtualServices_VirtualService_Details_ResourceAdmission
// Resources being allocated for the virtual-service.
type VirtualServices_VirtualService_Details_ResourceAdmission struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Thes status the of the resource allocation. The type is string.
    State interface{}

    // The amount of disk space being allocated for the virtual-service. The type
    // is string.
    DiskSpace interface{}

    // The amount of memory being allocated for the virtual-service. The type is
    // string.
    Memory interface{}

    // The percentage of cpu being allocated for the virtual-service. The type is
    // interface{} with range: 0..18446744073709551615.
    Cpu interface{}

    // The amount of VCPUs being allocated for the virtual-service. The type is
    // string.
    Vcpus interface{}
}

func (resourceAdmission *VirtualServices_VirtualService_Details_ResourceAdmission) GetFilter() yfilter.YFilter { return resourceAdmission.YFilter }

func (resourceAdmission *VirtualServices_VirtualService_Details_ResourceAdmission) SetFilter(yf yfilter.YFilter) { resourceAdmission.YFilter = yf }

func (resourceAdmission *VirtualServices_VirtualService_Details_ResourceAdmission) GetGoName(yname string) string {
    if yname == "state" { return "State" }
    if yname == "disk-space" { return "DiskSpace" }
    if yname == "memory" { return "Memory" }
    if yname == "cpu" { return "Cpu" }
    if yname == "vcpus" { return "Vcpus" }
    return ""
}

func (resourceAdmission *VirtualServices_VirtualService_Details_ResourceAdmission) GetSegmentPath() string {
    return "resource-admission"
}

func (resourceAdmission *VirtualServices_VirtualService_Details_ResourceAdmission) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (resourceAdmission *VirtualServices_VirtualService_Details_ResourceAdmission) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (resourceAdmission *VirtualServices_VirtualService_Details_ResourceAdmission) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["state"] = resourceAdmission.State
    leafs["disk-space"] = resourceAdmission.DiskSpace
    leafs["memory"] = resourceAdmission.Memory
    leafs["cpu"] = resourceAdmission.Cpu
    leafs["vcpus"] = resourceAdmission.Vcpus
    return leafs
}

func (resourceAdmission *VirtualServices_VirtualService_Details_ResourceAdmission) GetBundleName() string { return "cisco_ios_xe" }

func (resourceAdmission *VirtualServices_VirtualService_Details_ResourceAdmission) GetYangName() string { return "resource-admission" }

func (resourceAdmission *VirtualServices_VirtualService_Details_ResourceAdmission) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (resourceAdmission *VirtualServices_VirtualService_Details_ResourceAdmission) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (resourceAdmission *VirtualServices_VirtualService_Details_ResourceAdmission) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (resourceAdmission *VirtualServices_VirtualService_Details_ResourceAdmission) SetParent(parent types.Entity) { resourceAdmission.parent = parent }

func (resourceAdmission *VirtualServices_VirtualService_Details_ResourceAdmission) GetParent() types.Entity { return resourceAdmission.parent }

func (resourceAdmission *VirtualServices_VirtualService_Details_ResourceAdmission) GetParentYangName() string { return "details" }

// VirtualServices_VirtualService_Utilization
// Utilization of device resources for a virtual-service.
type VirtualServices_VirtualService_Utilization struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // The name of the virtual service. The type is string.
    Name interface{}

    // Details on the CPU utilization for the virtual-service.
    CpuUtil VirtualServices_VirtualService_Utilization_CpuUtil

    // Details on the memory usage for the virtual-service.
    MemoryUtil VirtualServices_VirtualService_Utilization_MemoryUtil
}

func (utilization *VirtualServices_VirtualService_Utilization) GetFilter() yfilter.YFilter { return utilization.YFilter }

func (utilization *VirtualServices_VirtualService_Utilization) SetFilter(yf yfilter.YFilter) { utilization.YFilter = yf }

func (utilization *VirtualServices_VirtualService_Utilization) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "cpu-util" { return "CpuUtil" }
    if yname == "memory-util" { return "MemoryUtil" }
    return ""
}

func (utilization *VirtualServices_VirtualService_Utilization) GetSegmentPath() string {
    return "utilization"
}

func (utilization *VirtualServices_VirtualService_Utilization) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "cpu-util" {
        return &utilization.CpuUtil
    }
    if childYangName == "memory-util" {
        return &utilization.MemoryUtil
    }
    return nil
}

func (utilization *VirtualServices_VirtualService_Utilization) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["cpu-util"] = &utilization.CpuUtil
    children["memory-util"] = &utilization.MemoryUtil
    return children
}

func (utilization *VirtualServices_VirtualService_Utilization) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = utilization.Name
    return leafs
}

func (utilization *VirtualServices_VirtualService_Utilization) GetBundleName() string { return "cisco_ios_xe" }

func (utilization *VirtualServices_VirtualService_Utilization) GetYangName() string { return "utilization" }

func (utilization *VirtualServices_VirtualService_Utilization) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (utilization *VirtualServices_VirtualService_Utilization) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (utilization *VirtualServices_VirtualService_Utilization) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (utilization *VirtualServices_VirtualService_Utilization) SetParent(parent types.Entity) { utilization.parent = parent }

func (utilization *VirtualServices_VirtualService_Utilization) GetParent() types.Entity { return utilization.parent }

func (utilization *VirtualServices_VirtualService_Utilization) GetParentYangName() string { return "virtual-service" }

// VirtualServices_VirtualService_Utilization_CpuUtil
// Details on the CPU utilization for the virtual-service.
type VirtualServices_VirtualService_Utilization_CpuUtil struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Percentage of requested CPU utilization by the virtual-service. The type is
    // interface{} with range: 0..18446744073709551615.
    RequestedApplicationUtil interface{}

    // Percetnage of CPU actual utilization for the virtual-service. The type is
    // interface{} with range: 0..18446744073709551615.
    ActualApplicationUtil interface{}

    // The state of the CPU utilization for the virtual-service. The type is
    // string.
    CpuState interface{}
}

func (cpuUtil *VirtualServices_VirtualService_Utilization_CpuUtil) GetFilter() yfilter.YFilter { return cpuUtil.YFilter }

func (cpuUtil *VirtualServices_VirtualService_Utilization_CpuUtil) SetFilter(yf yfilter.YFilter) { cpuUtil.YFilter = yf }

func (cpuUtil *VirtualServices_VirtualService_Utilization_CpuUtil) GetGoName(yname string) string {
    if yname == "requested-application-util" { return "RequestedApplicationUtil" }
    if yname == "actual-application-util" { return "ActualApplicationUtil" }
    if yname == "cpu-state" { return "CpuState" }
    return ""
}

func (cpuUtil *VirtualServices_VirtualService_Utilization_CpuUtil) GetSegmentPath() string {
    return "cpu-util"
}

func (cpuUtil *VirtualServices_VirtualService_Utilization_CpuUtil) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cpuUtil *VirtualServices_VirtualService_Utilization_CpuUtil) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cpuUtil *VirtualServices_VirtualService_Utilization_CpuUtil) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["requested-application-util"] = cpuUtil.RequestedApplicationUtil
    leafs["actual-application-util"] = cpuUtil.ActualApplicationUtil
    leafs["cpu-state"] = cpuUtil.CpuState
    return leafs
}

func (cpuUtil *VirtualServices_VirtualService_Utilization_CpuUtil) GetBundleName() string { return "cisco_ios_xe" }

func (cpuUtil *VirtualServices_VirtualService_Utilization_CpuUtil) GetYangName() string { return "cpu-util" }

func (cpuUtil *VirtualServices_VirtualService_Utilization_CpuUtil) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (cpuUtil *VirtualServices_VirtualService_Utilization_CpuUtil) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (cpuUtil *VirtualServices_VirtualService_Utilization_CpuUtil) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (cpuUtil *VirtualServices_VirtualService_Utilization_CpuUtil) SetParent(parent types.Entity) { cpuUtil.parent = parent }

func (cpuUtil *VirtualServices_VirtualService_Utilization_CpuUtil) GetParent() types.Entity { return cpuUtil.parent }

func (cpuUtil *VirtualServices_VirtualService_Utilization_CpuUtil) GetParentYangName() string { return "utilization" }

// VirtualServices_VirtualService_Utilization_MemoryUtil
// Details on the memory usage for the virtual-service.
type VirtualServices_VirtualService_Utilization_MemoryUtil struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Amount of memory being allocated for the virtual-service in KB. The type is
    // string.
    MemoryAllocation interface{}

    // Amount of memory being used for the virtual-service in KB. The type is
    // string.
    MemoryUsed interface{}
}

func (memoryUtil *VirtualServices_VirtualService_Utilization_MemoryUtil) GetFilter() yfilter.YFilter { return memoryUtil.YFilter }

func (memoryUtil *VirtualServices_VirtualService_Utilization_MemoryUtil) SetFilter(yf yfilter.YFilter) { memoryUtil.YFilter = yf }

func (memoryUtil *VirtualServices_VirtualService_Utilization_MemoryUtil) GetGoName(yname string) string {
    if yname == "memory-allocation" { return "MemoryAllocation" }
    if yname == "memory-used" { return "MemoryUsed" }
    return ""
}

func (memoryUtil *VirtualServices_VirtualService_Utilization_MemoryUtil) GetSegmentPath() string {
    return "memory-util"
}

func (memoryUtil *VirtualServices_VirtualService_Utilization_MemoryUtil) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (memoryUtil *VirtualServices_VirtualService_Utilization_MemoryUtil) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (memoryUtil *VirtualServices_VirtualService_Utilization_MemoryUtil) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["memory-allocation"] = memoryUtil.MemoryAllocation
    leafs["memory-used"] = memoryUtil.MemoryUsed
    return leafs
}

func (memoryUtil *VirtualServices_VirtualService_Utilization_MemoryUtil) GetBundleName() string { return "cisco_ios_xe" }

func (memoryUtil *VirtualServices_VirtualService_Utilization_MemoryUtil) GetYangName() string { return "memory-util" }

func (memoryUtil *VirtualServices_VirtualService_Utilization_MemoryUtil) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (memoryUtil *VirtualServices_VirtualService_Utilization_MemoryUtil) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (memoryUtil *VirtualServices_VirtualService_Utilization_MemoryUtil) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (memoryUtil *VirtualServices_VirtualService_Utilization_MemoryUtil) SetParent(parent types.Entity) { memoryUtil.parent = parent }

func (memoryUtil *VirtualServices_VirtualService_Utilization_MemoryUtil) GetParent() types.Entity { return memoryUtil.parent }

func (memoryUtil *VirtualServices_VirtualService_Utilization_MemoryUtil) GetParentYangName() string { return "utilization" }

// VirtualServices_VirtualService_NetworkUtils
// list of the network utilizations for the virtual-service.
type VirtualServices_VirtualService_NetworkUtils struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Details on a network utilization for the virtual-service. The type is slice
    // of VirtualServices_VirtualService_NetworkUtils_NetworkUtil.
    NetworkUtil []VirtualServices_VirtualService_NetworkUtils_NetworkUtil
}

func (networkUtils *VirtualServices_VirtualService_NetworkUtils) GetFilter() yfilter.YFilter { return networkUtils.YFilter }

func (networkUtils *VirtualServices_VirtualService_NetworkUtils) SetFilter(yf yfilter.YFilter) { networkUtils.YFilter = yf }

func (networkUtils *VirtualServices_VirtualService_NetworkUtils) GetGoName(yname string) string {
    if yname == "network-util" { return "NetworkUtil" }
    return ""
}

func (networkUtils *VirtualServices_VirtualService_NetworkUtils) GetSegmentPath() string {
    return "network-utils"
}

func (networkUtils *VirtualServices_VirtualService_NetworkUtils) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "network-util" {
        for _, c := range networkUtils.NetworkUtil {
            if networkUtils.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := VirtualServices_VirtualService_NetworkUtils_NetworkUtil{}
        networkUtils.NetworkUtil = append(networkUtils.NetworkUtil, child)
        return &networkUtils.NetworkUtil[len(networkUtils.NetworkUtil)-1]
    }
    return nil
}

func (networkUtils *VirtualServices_VirtualService_NetworkUtils) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range networkUtils.NetworkUtil {
        children[networkUtils.NetworkUtil[i].GetSegmentPath()] = &networkUtils.NetworkUtil[i]
    }
    return children
}

func (networkUtils *VirtualServices_VirtualService_NetworkUtils) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (networkUtils *VirtualServices_VirtualService_NetworkUtils) GetBundleName() string { return "cisco_ios_xe" }

func (networkUtils *VirtualServices_VirtualService_NetworkUtils) GetYangName() string { return "network-utils" }

func (networkUtils *VirtualServices_VirtualService_NetworkUtils) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (networkUtils *VirtualServices_VirtualService_NetworkUtils) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (networkUtils *VirtualServices_VirtualService_NetworkUtils) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (networkUtils *VirtualServices_VirtualService_NetworkUtils) SetParent(parent types.Entity) { networkUtils.parent = parent }

func (networkUtils *VirtualServices_VirtualService_NetworkUtils) GetParent() types.Entity { return networkUtils.parent }

func (networkUtils *VirtualServices_VirtualService_NetworkUtils) GetParentYangName() string { return "virtual-service" }

// VirtualServices_VirtualService_NetworkUtils_NetworkUtil
// Details on a network utilization for the virtual-service.
type VirtualServices_VirtualService_NetworkUtils_NetworkUtil struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the network that is being used for the
    // virtual-service. The type is string.
    Name interface{}

    // This attribute is a key. The alias of the network that is being used for
    // the virtual-service. The type is string.
    Alias interface{}

    // The number of rx packets being utilized for the virtual-service. The type
    // is interface{} with range: 0..18446744073709551615.
    RxPackets interface{}

    // The number of rx bytes being utilized for the virtual-service. The type is
    // interface{} with range: 0..18446744073709551615.
    RxBytes interface{}

    // The number of rx errors. The type is interface{} with range:
    // 0..18446744073709551615.
    RxErrors interface{}

    // The number of tx packets being utilized for the virtual-service. The type
    // is interface{} with range: 0..18446744073709551615.
    TxPackets interface{}

    // The number of tx bytes being utilized for the virtual-service. The type is
    // interface{} with range: 0..18446744073709551615.
    TxBytes interface{}

    // The number of tx errors. The type is interface{} with range:
    // 0..18446744073709551615.
    TxErrors interface{}
}

func (networkUtil *VirtualServices_VirtualService_NetworkUtils_NetworkUtil) GetFilter() yfilter.YFilter { return networkUtil.YFilter }

func (networkUtil *VirtualServices_VirtualService_NetworkUtils_NetworkUtil) SetFilter(yf yfilter.YFilter) { networkUtil.YFilter = yf }

func (networkUtil *VirtualServices_VirtualService_NetworkUtils_NetworkUtil) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "alias" { return "Alias" }
    if yname == "rx-packets" { return "RxPackets" }
    if yname == "rx-bytes" { return "RxBytes" }
    if yname == "rx-errors" { return "RxErrors" }
    if yname == "tx-packets" { return "TxPackets" }
    if yname == "tx-bytes" { return "TxBytes" }
    if yname == "tx-errors" { return "TxErrors" }
    return ""
}

func (networkUtil *VirtualServices_VirtualService_NetworkUtils_NetworkUtil) GetSegmentPath() string {
    return "network-util" + "[name='" + fmt.Sprintf("%v", networkUtil.Name) + "']" + "[alias='" + fmt.Sprintf("%v", networkUtil.Alias) + "']"
}

func (networkUtil *VirtualServices_VirtualService_NetworkUtils_NetworkUtil) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (networkUtil *VirtualServices_VirtualService_NetworkUtils_NetworkUtil) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (networkUtil *VirtualServices_VirtualService_NetworkUtils_NetworkUtil) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = networkUtil.Name
    leafs["alias"] = networkUtil.Alias
    leafs["rx-packets"] = networkUtil.RxPackets
    leafs["rx-bytes"] = networkUtil.RxBytes
    leafs["rx-errors"] = networkUtil.RxErrors
    leafs["tx-packets"] = networkUtil.TxPackets
    leafs["tx-bytes"] = networkUtil.TxBytes
    leafs["tx-errors"] = networkUtil.TxErrors
    return leafs
}

func (networkUtil *VirtualServices_VirtualService_NetworkUtils_NetworkUtil) GetBundleName() string { return "cisco_ios_xe" }

func (networkUtil *VirtualServices_VirtualService_NetworkUtils_NetworkUtil) GetYangName() string { return "network-util" }

func (networkUtil *VirtualServices_VirtualService_NetworkUtils_NetworkUtil) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (networkUtil *VirtualServices_VirtualService_NetworkUtils_NetworkUtil) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (networkUtil *VirtualServices_VirtualService_NetworkUtils_NetworkUtil) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (networkUtil *VirtualServices_VirtualService_NetworkUtils_NetworkUtil) SetParent(parent types.Entity) { networkUtil.parent = parent }

func (networkUtil *VirtualServices_VirtualService_NetworkUtils_NetworkUtil) GetParent() types.Entity { return networkUtil.parent }

func (networkUtil *VirtualServices_VirtualService_NetworkUtils_NetworkUtil) GetParentYangName() string { return "network-utils" }

// VirtualServices_VirtualService_StorageUtils
// List of the storage utilizations for the virtual-service.
type VirtualServices_VirtualService_StorageUtils struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Details on a storage utilization for the virtual-service. The type is slice
    // of VirtualServices_VirtualService_StorageUtils_StorageUtil.
    StorageUtil []VirtualServices_VirtualService_StorageUtils_StorageUtil
}

func (storageUtils *VirtualServices_VirtualService_StorageUtils) GetFilter() yfilter.YFilter { return storageUtils.YFilter }

func (storageUtils *VirtualServices_VirtualService_StorageUtils) SetFilter(yf yfilter.YFilter) { storageUtils.YFilter = yf }

func (storageUtils *VirtualServices_VirtualService_StorageUtils) GetGoName(yname string) string {
    if yname == "storage-util" { return "StorageUtil" }
    return ""
}

func (storageUtils *VirtualServices_VirtualService_StorageUtils) GetSegmentPath() string {
    return "storage-utils"
}

func (storageUtils *VirtualServices_VirtualService_StorageUtils) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "storage-util" {
        for _, c := range storageUtils.StorageUtil {
            if storageUtils.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := VirtualServices_VirtualService_StorageUtils_StorageUtil{}
        storageUtils.StorageUtil = append(storageUtils.StorageUtil, child)
        return &storageUtils.StorageUtil[len(storageUtils.StorageUtil)-1]
    }
    return nil
}

func (storageUtils *VirtualServices_VirtualService_StorageUtils) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range storageUtils.StorageUtil {
        children[storageUtils.StorageUtil[i].GetSegmentPath()] = &storageUtils.StorageUtil[i]
    }
    return children
}

func (storageUtils *VirtualServices_VirtualService_StorageUtils) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (storageUtils *VirtualServices_VirtualService_StorageUtils) GetBundleName() string { return "cisco_ios_xe" }

func (storageUtils *VirtualServices_VirtualService_StorageUtils) GetYangName() string { return "storage-utils" }

func (storageUtils *VirtualServices_VirtualService_StorageUtils) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (storageUtils *VirtualServices_VirtualService_StorageUtils) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (storageUtils *VirtualServices_VirtualService_StorageUtils) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (storageUtils *VirtualServices_VirtualService_StorageUtils) SetParent(parent types.Entity) { storageUtils.parent = parent }

func (storageUtils *VirtualServices_VirtualService_StorageUtils) GetParent() types.Entity { return storageUtils.parent }

func (storageUtils *VirtualServices_VirtualService_StorageUtils) GetParentYangName() string { return "virtual-service" }

// VirtualServices_VirtualService_StorageUtils_StorageUtil
// Details on a storage utilization for the virtual-service.
type VirtualServices_VirtualService_StorageUtils_StorageUtil struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the storage process being used for the
    // virtual-service. The type is string.
    Name interface{}

    // This attribute is a key. The alias of the storage process being used for
    // the virtual-service. The type is string.
    Alias interface{}

    // The number of RD bytes being used for the virtual-service. The type is
    // interface{} with range: 0..18446744073709551615.
    RdBytes interface{}

    // The number of rd requests being used for the virtual-service. The type is
    // interface{} with range: 0..18446744073709551615.
    RdRequests interface{}

    // The name of errors on the storage process. The type is interface{} with
    // range: 0..18446744073709551615.
    Errors interface{}

    // The number of WR bytes for the virtual-service. The type is interface{}
    // with range: 0..18446744073709551615.
    WrBytes interface{}

    // The number of WR requests for the virtual-service. The type is interface{}
    // with range: 0..18446744073709551615.
    WrRequests interface{}

    // The storage capactity in 1K blocks. The type is interface{} with range:
    // 0..18446744073709551615.
    Capacity interface{}

    // The available storage in 1K blocks. The type is string.
    Available interface{}

    // The number of 1K blocks being used. The type is interface{} with range:
    // 0..18446744073709551615.
    Used interface{}

    // The percetage of storage capactiy being used. The type is string.
    Usage interface{}
}

func (storageUtil *VirtualServices_VirtualService_StorageUtils_StorageUtil) GetFilter() yfilter.YFilter { return storageUtil.YFilter }

func (storageUtil *VirtualServices_VirtualService_StorageUtils_StorageUtil) SetFilter(yf yfilter.YFilter) { storageUtil.YFilter = yf }

func (storageUtil *VirtualServices_VirtualService_StorageUtils_StorageUtil) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "alias" { return "Alias" }
    if yname == "rd-bytes" { return "RdBytes" }
    if yname == "rd-requests" { return "RdRequests" }
    if yname == "errors" { return "Errors" }
    if yname == "wr-bytes" { return "WrBytes" }
    if yname == "wr-requests" { return "WrRequests" }
    if yname == "capacity" { return "Capacity" }
    if yname == "available" { return "Available" }
    if yname == "used" { return "Used" }
    if yname == "usage" { return "Usage" }
    return ""
}

func (storageUtil *VirtualServices_VirtualService_StorageUtils_StorageUtil) GetSegmentPath() string {
    return "storage-util" + "[name='" + fmt.Sprintf("%v", storageUtil.Name) + "']" + "[alias='" + fmt.Sprintf("%v", storageUtil.Alias) + "']"
}

func (storageUtil *VirtualServices_VirtualService_StorageUtils_StorageUtil) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (storageUtil *VirtualServices_VirtualService_StorageUtils_StorageUtil) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (storageUtil *VirtualServices_VirtualService_StorageUtils_StorageUtil) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = storageUtil.Name
    leafs["alias"] = storageUtil.Alias
    leafs["rd-bytes"] = storageUtil.RdBytes
    leafs["rd-requests"] = storageUtil.RdRequests
    leafs["errors"] = storageUtil.Errors
    leafs["wr-bytes"] = storageUtil.WrBytes
    leafs["wr-requests"] = storageUtil.WrRequests
    leafs["capacity"] = storageUtil.Capacity
    leafs["available"] = storageUtil.Available
    leafs["used"] = storageUtil.Used
    leafs["usage"] = storageUtil.Usage
    return leafs
}

func (storageUtil *VirtualServices_VirtualService_StorageUtils_StorageUtil) GetBundleName() string { return "cisco_ios_xe" }

func (storageUtil *VirtualServices_VirtualService_StorageUtils_StorageUtil) GetYangName() string { return "storage-util" }

func (storageUtil *VirtualServices_VirtualService_StorageUtils_StorageUtil) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (storageUtil *VirtualServices_VirtualService_StorageUtils_StorageUtil) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (storageUtil *VirtualServices_VirtualService_StorageUtils_StorageUtil) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (storageUtil *VirtualServices_VirtualService_StorageUtils_StorageUtil) SetParent(parent types.Entity) { storageUtil.parent = parent }

func (storageUtil *VirtualServices_VirtualService_StorageUtils_StorageUtil) GetParent() types.Entity { return storageUtil.parent }

func (storageUtil *VirtualServices_VirtualService_StorageUtils_StorageUtil) GetParentYangName() string { return "storage-utils" }

// VirtualServices_VirtualService_AttachedDevices
// Details for the devices attached to this virtual service.
type VirtualServices_VirtualService_AttachedDevices struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of attached devices. The type is slice of
    // VirtualServices_VirtualService_AttachedDevices_AttachedDevice.
    AttachedDevice []VirtualServices_VirtualService_AttachedDevices_AttachedDevice
}

func (attachedDevices *VirtualServices_VirtualService_AttachedDevices) GetFilter() yfilter.YFilter { return attachedDevices.YFilter }

func (attachedDevices *VirtualServices_VirtualService_AttachedDevices) SetFilter(yf yfilter.YFilter) { attachedDevices.YFilter = yf }

func (attachedDevices *VirtualServices_VirtualService_AttachedDevices) GetGoName(yname string) string {
    if yname == "attached-device" { return "AttachedDevice" }
    return ""
}

func (attachedDevices *VirtualServices_VirtualService_AttachedDevices) GetSegmentPath() string {
    return "attached-devices"
}

func (attachedDevices *VirtualServices_VirtualService_AttachedDevices) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "attached-device" {
        for _, c := range attachedDevices.AttachedDevice {
            if attachedDevices.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := VirtualServices_VirtualService_AttachedDevices_AttachedDevice{}
        attachedDevices.AttachedDevice = append(attachedDevices.AttachedDevice, child)
        return &attachedDevices.AttachedDevice[len(attachedDevices.AttachedDevice)-1]
    }
    return nil
}

func (attachedDevices *VirtualServices_VirtualService_AttachedDevices) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range attachedDevices.AttachedDevice {
        children[attachedDevices.AttachedDevice[i].GetSegmentPath()] = &attachedDevices.AttachedDevice[i]
    }
    return children
}

func (attachedDevices *VirtualServices_VirtualService_AttachedDevices) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (attachedDevices *VirtualServices_VirtualService_AttachedDevices) GetBundleName() string { return "cisco_ios_xe" }

func (attachedDevices *VirtualServices_VirtualService_AttachedDevices) GetYangName() string { return "attached-devices" }

func (attachedDevices *VirtualServices_VirtualService_AttachedDevices) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (attachedDevices *VirtualServices_VirtualService_AttachedDevices) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (attachedDevices *VirtualServices_VirtualService_AttachedDevices) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (attachedDevices *VirtualServices_VirtualService_AttachedDevices) SetParent(parent types.Entity) { attachedDevices.parent = parent }

func (attachedDevices *VirtualServices_VirtualService_AttachedDevices) GetParent() types.Entity { return attachedDevices.parent }

func (attachedDevices *VirtualServices_VirtualService_AttachedDevices) GetParentYangName() string { return "virtual-service" }

// VirtualServices_VirtualService_AttachedDevices_AttachedDevice
// List of attached devices.
type VirtualServices_VirtualService_AttachedDevices_AttachedDevice struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The name of the attached device. The type is
    // string.
    Name interface{}

    // The type of the attached device. The type is string.
    Type interface{}

    // The alias for the attached device. The type is string.
    Alias interface{}
}

func (attachedDevice *VirtualServices_VirtualService_AttachedDevices_AttachedDevice) GetFilter() yfilter.YFilter { return attachedDevice.YFilter }

func (attachedDevice *VirtualServices_VirtualService_AttachedDevices_AttachedDevice) SetFilter(yf yfilter.YFilter) { attachedDevice.YFilter = yf }

func (attachedDevice *VirtualServices_VirtualService_AttachedDevices_AttachedDevice) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "type" { return "Type" }
    if yname == "alias" { return "Alias" }
    return ""
}

func (attachedDevice *VirtualServices_VirtualService_AttachedDevices_AttachedDevice) GetSegmentPath() string {
    return "attached-device" + "[name='" + fmt.Sprintf("%v", attachedDevice.Name) + "']"
}

func (attachedDevice *VirtualServices_VirtualService_AttachedDevices_AttachedDevice) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (attachedDevice *VirtualServices_VirtualService_AttachedDevices_AttachedDevice) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (attachedDevice *VirtualServices_VirtualService_AttachedDevices_AttachedDevice) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = attachedDevice.Name
    leafs["type"] = attachedDevice.Type
    leafs["alias"] = attachedDevice.Alias
    return leafs
}

func (attachedDevice *VirtualServices_VirtualService_AttachedDevices_AttachedDevice) GetBundleName() string { return "cisco_ios_xe" }

func (attachedDevice *VirtualServices_VirtualService_AttachedDevices_AttachedDevice) GetYangName() string { return "attached-device" }

func (attachedDevice *VirtualServices_VirtualService_AttachedDevices_AttachedDevice) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (attachedDevice *VirtualServices_VirtualService_AttachedDevices_AttachedDevice) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (attachedDevice *VirtualServices_VirtualService_AttachedDevices_AttachedDevice) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (attachedDevice *VirtualServices_VirtualService_AttachedDevices_AttachedDevice) SetParent(parent types.Entity) { attachedDevice.parent = parent }

func (attachedDevice *VirtualServices_VirtualService_AttachedDevices_AttachedDevice) GetParent() types.Entity { return attachedDevice.parent }

func (attachedDevice *VirtualServices_VirtualService_AttachedDevices_AttachedDevice) GetParentYangName() string { return "attached-devices" }

// VirtualServices_VirtualService_NetworkInterfaces
// Details for the network interfaces.
type VirtualServices_VirtualService_NetworkInterfaces struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Details for a network interface. The type is slice of
    // VirtualServices_VirtualService_NetworkInterfaces_NetworkInterface.
    NetworkInterface []VirtualServices_VirtualService_NetworkInterfaces_NetworkInterface
}

func (networkInterfaces *VirtualServices_VirtualService_NetworkInterfaces) GetFilter() yfilter.YFilter { return networkInterfaces.YFilter }

func (networkInterfaces *VirtualServices_VirtualService_NetworkInterfaces) SetFilter(yf yfilter.YFilter) { networkInterfaces.YFilter = yf }

func (networkInterfaces *VirtualServices_VirtualService_NetworkInterfaces) GetGoName(yname string) string {
    if yname == "network-interface" { return "NetworkInterface" }
    return ""
}

func (networkInterfaces *VirtualServices_VirtualService_NetworkInterfaces) GetSegmentPath() string {
    return "network-interfaces"
}

func (networkInterfaces *VirtualServices_VirtualService_NetworkInterfaces) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "network-interface" {
        for _, c := range networkInterfaces.NetworkInterface {
            if networkInterfaces.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := VirtualServices_VirtualService_NetworkInterfaces_NetworkInterface{}
        networkInterfaces.NetworkInterface = append(networkInterfaces.NetworkInterface, child)
        return &networkInterfaces.NetworkInterface[len(networkInterfaces.NetworkInterface)-1]
    }
    return nil
}

func (networkInterfaces *VirtualServices_VirtualService_NetworkInterfaces) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range networkInterfaces.NetworkInterface {
        children[networkInterfaces.NetworkInterface[i].GetSegmentPath()] = &networkInterfaces.NetworkInterface[i]
    }
    return children
}

func (networkInterfaces *VirtualServices_VirtualService_NetworkInterfaces) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (networkInterfaces *VirtualServices_VirtualService_NetworkInterfaces) GetBundleName() string { return "cisco_ios_xe" }

func (networkInterfaces *VirtualServices_VirtualService_NetworkInterfaces) GetYangName() string { return "network-interfaces" }

func (networkInterfaces *VirtualServices_VirtualService_NetworkInterfaces) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (networkInterfaces *VirtualServices_VirtualService_NetworkInterfaces) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (networkInterfaces *VirtualServices_VirtualService_NetworkInterfaces) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (networkInterfaces *VirtualServices_VirtualService_NetworkInterfaces) SetParent(parent types.Entity) { networkInterfaces.parent = parent }

func (networkInterfaces *VirtualServices_VirtualService_NetworkInterfaces) GetParent() types.Entity { return networkInterfaces.parent }

func (networkInterfaces *VirtualServices_VirtualService_NetworkInterfaces) GetParentYangName() string { return "virtual-service" }

// VirtualServices_VirtualService_NetworkInterfaces_NetworkInterface
// Details for a network interface.
type VirtualServices_VirtualService_NetworkInterfaces_NetworkInterface struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. The MAC address for the network interface. The
    // type is string with pattern: [0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}.
    MacAddress interface{}

    // Name of the the attached interface. The type is string.
    AttachedInterface interface{}
}

func (networkInterface *VirtualServices_VirtualService_NetworkInterfaces_NetworkInterface) GetFilter() yfilter.YFilter { return networkInterface.YFilter }

func (networkInterface *VirtualServices_VirtualService_NetworkInterfaces_NetworkInterface) SetFilter(yf yfilter.YFilter) { networkInterface.YFilter = yf }

func (networkInterface *VirtualServices_VirtualService_NetworkInterfaces_NetworkInterface) GetGoName(yname string) string {
    if yname == "mac-address" { return "MacAddress" }
    if yname == "attached-interface" { return "AttachedInterface" }
    return ""
}

func (networkInterface *VirtualServices_VirtualService_NetworkInterfaces_NetworkInterface) GetSegmentPath() string {
    return "network-interface" + "[mac-address='" + fmt.Sprintf("%v", networkInterface.MacAddress) + "']"
}

func (networkInterface *VirtualServices_VirtualService_NetworkInterfaces_NetworkInterface) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (networkInterface *VirtualServices_VirtualService_NetworkInterfaces_NetworkInterface) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (networkInterface *VirtualServices_VirtualService_NetworkInterfaces_NetworkInterface) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["mac-address"] = networkInterface.MacAddress
    leafs["attached-interface"] = networkInterface.AttachedInterface
    return leafs
}

func (networkInterface *VirtualServices_VirtualService_NetworkInterfaces_NetworkInterface) GetBundleName() string { return "cisco_ios_xe" }

func (networkInterface *VirtualServices_VirtualService_NetworkInterfaces_NetworkInterface) GetYangName() string { return "network-interface" }

func (networkInterface *VirtualServices_VirtualService_NetworkInterfaces_NetworkInterface) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (networkInterface *VirtualServices_VirtualService_NetworkInterfaces_NetworkInterface) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (networkInterface *VirtualServices_VirtualService_NetworkInterfaces_NetworkInterface) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (networkInterface *VirtualServices_VirtualService_NetworkInterfaces_NetworkInterface) SetParent(parent types.Entity) { networkInterface.parent = parent }

func (networkInterface *VirtualServices_VirtualService_NetworkInterfaces_NetworkInterface) GetParent() types.Entity { return networkInterface.parent }

func (networkInterface *VirtualServices_VirtualService_NetworkInterfaces_NetworkInterface) GetParentYangName() string { return "network-interfaces" }

// VirtualServices_VirtualService_GuestRoutes
// Routes for the guest interface.
type VirtualServices_VirtualService_GuestRoutes struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // List of guest routes for a guest interface. The type is slice of
    // VirtualServices_VirtualService_GuestRoutes_GuestRoute.
    GuestRoute []VirtualServices_VirtualService_GuestRoutes_GuestRoute
}

func (guestRoutes *VirtualServices_VirtualService_GuestRoutes) GetFilter() yfilter.YFilter { return guestRoutes.YFilter }

func (guestRoutes *VirtualServices_VirtualService_GuestRoutes) SetFilter(yf yfilter.YFilter) { guestRoutes.YFilter = yf }

func (guestRoutes *VirtualServices_VirtualService_GuestRoutes) GetGoName(yname string) string {
    if yname == "guest-route" { return "GuestRoute" }
    return ""
}

func (guestRoutes *VirtualServices_VirtualService_GuestRoutes) GetSegmentPath() string {
    return "guest-routes"
}

func (guestRoutes *VirtualServices_VirtualService_GuestRoutes) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "guest-route" {
        for _, c := range guestRoutes.GuestRoute {
            if guestRoutes.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := VirtualServices_VirtualService_GuestRoutes_GuestRoute{}
        guestRoutes.GuestRoute = append(guestRoutes.GuestRoute, child)
        return &guestRoutes.GuestRoute[len(guestRoutes.GuestRoute)-1]
    }
    return nil
}

func (guestRoutes *VirtualServices_VirtualService_GuestRoutes) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range guestRoutes.GuestRoute {
        children[guestRoutes.GuestRoute[i].GetSegmentPath()] = &guestRoutes.GuestRoute[i]
    }
    return children
}

func (guestRoutes *VirtualServices_VirtualService_GuestRoutes) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (guestRoutes *VirtualServices_VirtualService_GuestRoutes) GetBundleName() string { return "cisco_ios_xe" }

func (guestRoutes *VirtualServices_VirtualService_GuestRoutes) GetYangName() string { return "guest-routes" }

func (guestRoutes *VirtualServices_VirtualService_GuestRoutes) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (guestRoutes *VirtualServices_VirtualService_GuestRoutes) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (guestRoutes *VirtualServices_VirtualService_GuestRoutes) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (guestRoutes *VirtualServices_VirtualService_GuestRoutes) SetParent(parent types.Entity) { guestRoutes.parent = parent }

func (guestRoutes *VirtualServices_VirtualService_GuestRoutes) GetParent() types.Entity { return guestRoutes.parent }

func (guestRoutes *VirtualServices_VirtualService_GuestRoutes) GetParentYangName() string { return "virtual-service" }

// VirtualServices_VirtualService_GuestRoutes_GuestRoute
// List of guest routes for a guest interface.
type VirtualServices_VirtualService_GuestRoutes_GuestRoute struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. A guest route for a guest interface. The type is
    // string.
    Route interface{}
}

func (guestRoute *VirtualServices_VirtualService_GuestRoutes_GuestRoute) GetFilter() yfilter.YFilter { return guestRoute.YFilter }

func (guestRoute *VirtualServices_VirtualService_GuestRoutes_GuestRoute) SetFilter(yf yfilter.YFilter) { guestRoute.YFilter = yf }

func (guestRoute *VirtualServices_VirtualService_GuestRoutes_GuestRoute) GetGoName(yname string) string {
    if yname == "route" { return "Route" }
    return ""
}

func (guestRoute *VirtualServices_VirtualService_GuestRoutes_GuestRoute) GetSegmentPath() string {
    return "guest-route" + "[route='" + fmt.Sprintf("%v", guestRoute.Route) + "']"
}

func (guestRoute *VirtualServices_VirtualService_GuestRoutes_GuestRoute) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (guestRoute *VirtualServices_VirtualService_GuestRoutes_GuestRoute) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (guestRoute *VirtualServices_VirtualService_GuestRoutes_GuestRoute) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["route"] = guestRoute.Route
    return leafs
}

func (guestRoute *VirtualServices_VirtualService_GuestRoutes_GuestRoute) GetBundleName() string { return "cisco_ios_xe" }

func (guestRoute *VirtualServices_VirtualService_GuestRoutes_GuestRoute) GetYangName() string { return "guest-route" }

func (guestRoute *VirtualServices_VirtualService_GuestRoutes_GuestRoute) GetBundleYangModelsLocation() string { return cisco_ios_xe.GetModelsPath() }

func (guestRoute *VirtualServices_VirtualService_GuestRoutes_GuestRoute) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xe.GetCapabilities() }

func (guestRoute *VirtualServices_VirtualService_GuestRoutes_GuestRoute) GetNamespaceTable() map[string]string {
    return cisco_ios_xe.GetNamespaces() }

func (guestRoute *VirtualServices_VirtualService_GuestRoutes_GuestRoute) SetParent(parent types.Entity) { guestRoute.parent = parent }

func (guestRoute *VirtualServices_VirtualService_GuestRoutes_GuestRoute) GetParent() types.Entity { return guestRoute.parent }

func (guestRoute *VirtualServices_VirtualService_GuestRoutes_GuestRoute) GetParentYangName() string { return "guest-routes" }

