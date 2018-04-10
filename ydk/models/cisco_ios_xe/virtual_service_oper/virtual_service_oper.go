// This module contains a collection of YANG definitions for
// monitoring virtual services in a Network Element.
// Copyright (c) 2016-2017 by Cisco Systems, Inc.
// All rights reserved.
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
// Information on all virtual services
type VirtualServices struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of virtual services. The type is slice of
    // VirtualServices_VirtualService.
    VirtualService []VirtualServices_VirtualService
}

func (virtualServices *VirtualServices) GetEntityData() *types.CommonEntityData {
    virtualServices.EntityData.YFilter = virtualServices.YFilter
    virtualServices.EntityData.YangName = "virtual-services"
    virtualServices.EntityData.BundleName = "cisco_ios_xe"
    virtualServices.EntityData.ParentYangName = "Cisco-IOS-XE-virtual-service-oper"
    virtualServices.EntityData.SegmentPath = "Cisco-IOS-XE-virtual-service-oper:virtual-services"
    virtualServices.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    virtualServices.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    virtualServices.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    virtualServices.EntityData.Children = make(map[string]types.YChild)
    virtualServices.EntityData.Children["virtual-service"] = types.YChild{"VirtualService", nil}
    for i := range virtualServices.VirtualService {
        virtualServices.EntityData.Children[types.GetSegmentPath(&virtualServices.VirtualService[i])] = types.YChild{"VirtualService", &virtualServices.VirtualService[i]}
    }
    virtualServices.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(virtualServices.EntityData)
}

// VirtualServices_VirtualService
// List of virtual services
type VirtualServices_VirtualService struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Virtual service name. The type is string.
    Name interface{}

    // Virtual service details.
    Details VirtualServices_VirtualService_Details

    // Virtual service resource utilization details.
    Utilization VirtualServices_VirtualService_Utilization

    // Virtual service network utilization details.
    NetworkUtils VirtualServices_VirtualService_NetworkUtils

    // Virtual service storage utilization details.
    StorageUtils VirtualServices_VirtualService_StorageUtils

    // Virtual service process details.
    Processes VirtualServices_VirtualService_Processes

    // Virtual service attached device details.
    AttachedDevices VirtualServices_VirtualService_AttachedDevices

    // Virtual service network interface details.
    NetworkInterfaces VirtualServices_VirtualService_NetworkInterfaces

    // Virtual service guest route details.
    GuestRoutes VirtualServices_VirtualService_GuestRoutes
}

func (virtualService *VirtualServices_VirtualService) GetEntityData() *types.CommonEntityData {
    virtualService.EntityData.YFilter = virtualService.YFilter
    virtualService.EntityData.YangName = "virtual-service"
    virtualService.EntityData.BundleName = "cisco_ios_xe"
    virtualService.EntityData.ParentYangName = "virtual-services"
    virtualService.EntityData.SegmentPath = "virtual-service" + "[name='" + fmt.Sprintf("%v", virtualService.Name) + "']"
    virtualService.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    virtualService.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    virtualService.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    virtualService.EntityData.Children = make(map[string]types.YChild)
    virtualService.EntityData.Children["details"] = types.YChild{"Details", &virtualService.Details}
    virtualService.EntityData.Children["utilization"] = types.YChild{"Utilization", &virtualService.Utilization}
    virtualService.EntityData.Children["network-utils"] = types.YChild{"NetworkUtils", &virtualService.NetworkUtils}
    virtualService.EntityData.Children["storage-utils"] = types.YChild{"StorageUtils", &virtualService.StorageUtils}
    virtualService.EntityData.Children["processes"] = types.YChild{"Processes", &virtualService.Processes}
    virtualService.EntityData.Children["attached-devices"] = types.YChild{"AttachedDevices", &virtualService.AttachedDevices}
    virtualService.EntityData.Children["network-interfaces"] = types.YChild{"NetworkInterfaces", &virtualService.NetworkInterfaces}
    virtualService.EntityData.Children["guest-routes"] = types.YChild{"GuestRoutes", &virtualService.GuestRoutes}
    virtualService.EntityData.Leafs = make(map[string]types.YLeaf)
    virtualService.EntityData.Leafs["name"] = types.YLeaf{"Name", virtualService.Name}
    return &(virtualService.EntityData)
}

// VirtualServices_VirtualService_Details
// Virtual service details
type VirtualServices_VirtualService_Details struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // State of the virtual service. The type is string.
    State interface{}

    // Activated profile name. The type is string.
    ActivatedProfileName interface{}

    // Guest interface name. The type is string.
    GuestInterface interface{}

    // Virtual service packaging information.
    PackageInformation VirtualServices_VirtualService_Details_PackageInformation

    // Guest status details.
    DetailedGuestStatus VirtualServices_VirtualService_Details_DetailedGuestStatus

    // Resource reservation details.
    ResourceReservation VirtualServices_VirtualService_Details_ResourceReservation

    // Resources allocated for the virtual service.
    ResourceAdmission VirtualServices_VirtualService_Details_ResourceAdmission
}

func (details *VirtualServices_VirtualService_Details) GetEntityData() *types.CommonEntityData {
    details.EntityData.YFilter = details.YFilter
    details.EntityData.YangName = "details"
    details.EntityData.BundleName = "cisco_ios_xe"
    details.EntityData.ParentYangName = "virtual-service"
    details.EntityData.SegmentPath = "details"
    details.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    details.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    details.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    details.EntityData.Children = make(map[string]types.YChild)
    details.EntityData.Children["package-information"] = types.YChild{"PackageInformation", &details.PackageInformation}
    details.EntityData.Children["detailed-guest-status"] = types.YChild{"DetailedGuestStatus", &details.DetailedGuestStatus}
    details.EntityData.Children["resource-reservation"] = types.YChild{"ResourceReservation", &details.ResourceReservation}
    details.EntityData.Children["resource-admission"] = types.YChild{"ResourceAdmission", &details.ResourceAdmission}
    details.EntityData.Leafs = make(map[string]types.YLeaf)
    details.EntityData.Leafs["state"] = types.YLeaf{"State", details.State}
    details.EntityData.Leafs["activated-profile-name"] = types.YLeaf{"ActivatedProfileName", details.ActivatedProfileName}
    details.EntityData.Leafs["guest-interface"] = types.YLeaf{"GuestInterface", details.GuestInterface}
    return &(details.EntityData)
}

// VirtualServices_VirtualService_Details_PackageInformation
// Virtual service packaging information
type VirtualServices_VirtualService_Details_PackageInformation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Package name. The type is string.
    Name interface{}

    // Package path. The type is string.
    Path interface{}

    // Application details.
    Application VirtualServices_VirtualService_Details_PackageInformation_Application

    // Key signing details.
    Signing VirtualServices_VirtualService_Details_PackageInformation_Signing

    // Licensing details.
    Licensing VirtualServices_VirtualService_Details_PackageInformation_Licensing
}

func (packageInformation *VirtualServices_VirtualService_Details_PackageInformation) GetEntityData() *types.CommonEntityData {
    packageInformation.EntityData.YFilter = packageInformation.YFilter
    packageInformation.EntityData.YangName = "package-information"
    packageInformation.EntityData.BundleName = "cisco_ios_xe"
    packageInformation.EntityData.ParentYangName = "details"
    packageInformation.EntityData.SegmentPath = "package-information"
    packageInformation.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    packageInformation.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    packageInformation.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    packageInformation.EntityData.Children = make(map[string]types.YChild)
    packageInformation.EntityData.Children["application"] = types.YChild{"Application", &packageInformation.Application}
    packageInformation.EntityData.Children["signing"] = types.YChild{"Signing", &packageInformation.Signing}
    packageInformation.EntityData.Children["licensing"] = types.YChild{"Licensing", &packageInformation.Licensing}
    packageInformation.EntityData.Leafs = make(map[string]types.YLeaf)
    packageInformation.EntityData.Leafs["name"] = types.YLeaf{"Name", packageInformation.Name}
    packageInformation.EntityData.Leafs["path"] = types.YLeaf{"Path", packageInformation.Path}
    return &(packageInformation.EntityData)
}

// VirtualServices_VirtualService_Details_PackageInformation_Application
// Application details
type VirtualServices_VirtualService_Details_PackageInformation_Application struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Application name. The type is string.
    Name interface{}

    // Application version. The type is string.
    InstalledVersion interface{}

    // Application description. The type is string.
    Description interface{}

    // Application type. The type is string.
    Type_ interface{}
}

func (application *VirtualServices_VirtualService_Details_PackageInformation_Application) GetEntityData() *types.CommonEntityData {
    application.EntityData.YFilter = application.YFilter
    application.EntityData.YangName = "application"
    application.EntityData.BundleName = "cisco_ios_xe"
    application.EntityData.ParentYangName = "package-information"
    application.EntityData.SegmentPath = "application"
    application.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    application.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    application.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    application.EntityData.Children = make(map[string]types.YChild)
    application.EntityData.Leafs = make(map[string]types.YLeaf)
    application.EntityData.Leafs["name"] = types.YLeaf{"Name", application.Name}
    application.EntityData.Leafs["installed-version"] = types.YLeaf{"InstalledVersion", application.InstalledVersion}
    application.EntityData.Leafs["description"] = types.YLeaf{"Description", application.Description}
    application.EntityData.Leafs["type"] = types.YLeaf{"Type_", application.Type_}
    return &(application.EntityData)
}

// VirtualServices_VirtualService_Details_PackageInformation_Signing
// Key signing details
type VirtualServices_VirtualService_Details_PackageInformation_Signing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Signed key type. The type is string.
    KeyType interface{}

    // Method the key was signed. The type is string.
    Method interface{}
}

func (signing *VirtualServices_VirtualService_Details_PackageInformation_Signing) GetEntityData() *types.CommonEntityData {
    signing.EntityData.YFilter = signing.YFilter
    signing.EntityData.YangName = "signing"
    signing.EntityData.BundleName = "cisco_ios_xe"
    signing.EntityData.ParentYangName = "package-information"
    signing.EntityData.SegmentPath = "signing"
    signing.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    signing.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    signing.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    signing.EntityData.Children = make(map[string]types.YChild)
    signing.EntityData.Leafs = make(map[string]types.YLeaf)
    signing.EntityData.Leafs["key-type"] = types.YLeaf{"KeyType", signing.KeyType}
    signing.EntityData.Leafs["method"] = types.YLeaf{"Method", signing.Method}
    return &(signing.EntityData)
}

// VirtualServices_VirtualService_Details_PackageInformation_Licensing
// Licensing details
type VirtualServices_VirtualService_Details_PackageInformation_Licensing struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // License name. The type is string.
    Name interface{}

    // License version. The type is string.
    Version interface{}
}

func (licensing *VirtualServices_VirtualService_Details_PackageInformation_Licensing) GetEntityData() *types.CommonEntityData {
    licensing.EntityData.YFilter = licensing.YFilter
    licensing.EntityData.YangName = "licensing"
    licensing.EntityData.BundleName = "cisco_ios_xe"
    licensing.EntityData.ParentYangName = "package-information"
    licensing.EntityData.SegmentPath = "licensing"
    licensing.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    licensing.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    licensing.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    licensing.EntityData.Children = make(map[string]types.YChild)
    licensing.EntityData.Leafs = make(map[string]types.YLeaf)
    licensing.EntityData.Leafs["name"] = types.YLeaf{"Name", licensing.Name}
    licensing.EntityData.Leafs["version"] = types.YLeaf{"Version", licensing.Version}
    return &(licensing.EntityData)
}

// VirtualServices_VirtualService_Details_DetailedGuestStatus
// Guest status details
type VirtualServices_VirtualService_Details_DetailedGuestStatus struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of all processes.
    Processes VirtualServices_VirtualService_Details_DetailedGuestStatus_Processes
}

func (detailedGuestStatus *VirtualServices_VirtualService_Details_DetailedGuestStatus) GetEntityData() *types.CommonEntityData {
    detailedGuestStatus.EntityData.YFilter = detailedGuestStatus.YFilter
    detailedGuestStatus.EntityData.YangName = "detailed-guest-status"
    detailedGuestStatus.EntityData.BundleName = "cisco_ios_xe"
    detailedGuestStatus.EntityData.ParentYangName = "details"
    detailedGuestStatus.EntityData.SegmentPath = "detailed-guest-status"
    detailedGuestStatus.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    detailedGuestStatus.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    detailedGuestStatus.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    detailedGuestStatus.EntityData.Children = make(map[string]types.YChild)
    detailedGuestStatus.EntityData.Children["processes"] = types.YChild{"Processes", &detailedGuestStatus.Processes}
    detailedGuestStatus.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(detailedGuestStatus.EntityData)
}

// VirtualServices_VirtualService_Details_DetailedGuestStatus_Processes
// List of all processes
type VirtualServices_VirtualService_Details_DetailedGuestStatus_Processes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Process name. The type is string.
    Name interface{}

    // Process status. The type is string.
    Status interface{}

    // Process ID. The type is string.
    Pid interface{}

    // Process uptime. The type is string.
    Uptime interface{}

    // Amount of process memory. The type is string.
    Memory interface{}
}

func (processes *VirtualServices_VirtualService_Details_DetailedGuestStatus_Processes) GetEntityData() *types.CommonEntityData {
    processes.EntityData.YFilter = processes.YFilter
    processes.EntityData.YangName = "processes"
    processes.EntityData.BundleName = "cisco_ios_xe"
    processes.EntityData.ParentYangName = "detailed-guest-status"
    processes.EntityData.SegmentPath = "processes"
    processes.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    processes.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    processes.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    processes.EntityData.Children = make(map[string]types.YChild)
    processes.EntityData.Leafs = make(map[string]types.YLeaf)
    processes.EntityData.Leafs["name"] = types.YLeaf{"Name", processes.Name}
    processes.EntityData.Leafs["status"] = types.YLeaf{"Status", processes.Status}
    processes.EntityData.Leafs["pid"] = types.YLeaf{"Pid", processes.Pid}
    processes.EntityData.Leafs["uptime"] = types.YLeaf{"Uptime", processes.Uptime}
    processes.EntityData.Leafs["memory"] = types.YLeaf{"Memory", processes.Memory}
    return &(processes.EntityData)
}

// VirtualServices_VirtualService_Details_ResourceReservation
// Resource reservation details
type VirtualServices_VirtualService_Details_ResourceReservation struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Amount of reserverd disk space in MB. The type is interface{} with range:
    // 0..18446744073709551615.
    Disk interface{}

    // Amount of reserved memory in MB. The type is interface{} with range:
    // 0..18446744073709551615.
    Memory interface{}

    // Amount of reserved cpu in unit. The type is interface{} with range:
    // 0..18446744073709551615.
    Cpu interface{}
}

func (resourceReservation *VirtualServices_VirtualService_Details_ResourceReservation) GetEntityData() *types.CommonEntityData {
    resourceReservation.EntityData.YFilter = resourceReservation.YFilter
    resourceReservation.EntityData.YangName = "resource-reservation"
    resourceReservation.EntityData.BundleName = "cisco_ios_xe"
    resourceReservation.EntityData.ParentYangName = "details"
    resourceReservation.EntityData.SegmentPath = "resource-reservation"
    resourceReservation.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    resourceReservation.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    resourceReservation.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    resourceReservation.EntityData.Children = make(map[string]types.YChild)
    resourceReservation.EntityData.Leafs = make(map[string]types.YLeaf)
    resourceReservation.EntityData.Leafs["disk"] = types.YLeaf{"Disk", resourceReservation.Disk}
    resourceReservation.EntityData.Leafs["memory"] = types.YLeaf{"Memory", resourceReservation.Memory}
    resourceReservation.EntityData.Leafs["cpu"] = types.YLeaf{"Cpu", resourceReservation.Cpu}
    return &(resourceReservation.EntityData)
}

// VirtualServices_VirtualService_Details_ResourceAdmission
// Resources allocated for the virtual service
type VirtualServices_VirtualService_Details_ResourceAdmission struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Status of the resource allocation. The type is string.
    State interface{}

    // Amount of disk space allocated for the virtual service in MB. The type is
    // string.
    DiskSpace interface{}

    // Amount of memory allocated for the virtual service in MB. The type is
    // string.
    Memory interface{}

    // Percentage of cpu allocated for the virtual-service in unit. The type is
    // interface{} with range: 0..18446744073709551615.
    Cpu interface{}

    // Amount of VCPUs allocated for the virtual service. The type is string.
    Vcpus interface{}
}

func (resourceAdmission *VirtualServices_VirtualService_Details_ResourceAdmission) GetEntityData() *types.CommonEntityData {
    resourceAdmission.EntityData.YFilter = resourceAdmission.YFilter
    resourceAdmission.EntityData.YangName = "resource-admission"
    resourceAdmission.EntityData.BundleName = "cisco_ios_xe"
    resourceAdmission.EntityData.ParentYangName = "details"
    resourceAdmission.EntityData.SegmentPath = "resource-admission"
    resourceAdmission.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    resourceAdmission.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    resourceAdmission.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    resourceAdmission.EntityData.Children = make(map[string]types.YChild)
    resourceAdmission.EntityData.Leafs = make(map[string]types.YLeaf)
    resourceAdmission.EntityData.Leafs["state"] = types.YLeaf{"State", resourceAdmission.State}
    resourceAdmission.EntityData.Leafs["disk-space"] = types.YLeaf{"DiskSpace", resourceAdmission.DiskSpace}
    resourceAdmission.EntityData.Leafs["memory"] = types.YLeaf{"Memory", resourceAdmission.Memory}
    resourceAdmission.EntityData.Leafs["cpu"] = types.YLeaf{"Cpu", resourceAdmission.Cpu}
    resourceAdmission.EntityData.Leafs["vcpus"] = types.YLeaf{"Vcpus", resourceAdmission.Vcpus}
    return &(resourceAdmission.EntityData)
}

// VirtualServices_VirtualService_Utilization
// Virtual service resource utilization details
type VirtualServices_VirtualService_Utilization struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Name of the virtual service. The type is string.
    Name interface{}

    // CPU utilization information.
    CpuUtil VirtualServices_VirtualService_Utilization_CpuUtil

    // Memory utilization information.
    MemoryUtil VirtualServices_VirtualService_Utilization_MemoryUtil
}

func (utilization *VirtualServices_VirtualService_Utilization) GetEntityData() *types.CommonEntityData {
    utilization.EntityData.YFilter = utilization.YFilter
    utilization.EntityData.YangName = "utilization"
    utilization.EntityData.BundleName = "cisco_ios_xe"
    utilization.EntityData.ParentYangName = "virtual-service"
    utilization.EntityData.SegmentPath = "utilization"
    utilization.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    utilization.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    utilization.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    utilization.EntityData.Children = make(map[string]types.YChild)
    utilization.EntityData.Children["cpu-util"] = types.YChild{"CpuUtil", &utilization.CpuUtil}
    utilization.EntityData.Children["memory-util"] = types.YChild{"MemoryUtil", &utilization.MemoryUtil}
    utilization.EntityData.Leafs = make(map[string]types.YLeaf)
    utilization.EntityData.Leafs["name"] = types.YLeaf{"Name", utilization.Name}
    return &(utilization.EntityData)
}

// VirtualServices_VirtualService_Utilization_CpuUtil
// CPU utilization information
type VirtualServices_VirtualService_Utilization_CpuUtil struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Amount of requested CPU utilization by the virtual service. The type is
    // interface{} with range: 0..18446744073709551615.
    RequestedApplicationUtil interface{}

    // Percetnage of CPU actual utilization for the virtual service. The type is
    // interface{} with range: 0..18446744073709551615.
    ActualApplicationUtil interface{}

    // State of the CPU utilization for the virtual-service. The type is string.
    CpuState interface{}
}

func (cpuUtil *VirtualServices_VirtualService_Utilization_CpuUtil) GetEntityData() *types.CommonEntityData {
    cpuUtil.EntityData.YFilter = cpuUtil.YFilter
    cpuUtil.EntityData.YangName = "cpu-util"
    cpuUtil.EntityData.BundleName = "cisco_ios_xe"
    cpuUtil.EntityData.ParentYangName = "utilization"
    cpuUtil.EntityData.SegmentPath = "cpu-util"
    cpuUtil.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cpuUtil.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cpuUtil.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cpuUtil.EntityData.Children = make(map[string]types.YChild)
    cpuUtil.EntityData.Leafs = make(map[string]types.YLeaf)
    cpuUtil.EntityData.Leafs["requested-application-util"] = types.YLeaf{"RequestedApplicationUtil", cpuUtil.RequestedApplicationUtil}
    cpuUtil.EntityData.Leafs["actual-application-util"] = types.YLeaf{"ActualApplicationUtil", cpuUtil.ActualApplicationUtil}
    cpuUtil.EntityData.Leafs["cpu-state"] = types.YLeaf{"CpuState", cpuUtil.CpuState}
    return &(cpuUtil.EntityData)
}

// VirtualServices_VirtualService_Utilization_MemoryUtil
// Memory utilization information
type VirtualServices_VirtualService_Utilization_MemoryUtil struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Amount of memory allocated for the virtual service in MB. The type is
    // string.
    MemoryAllocation interface{}

    // Amount of used memory for the virtual service in KB. The type is string.
    MemoryUsed interface{}
}

func (memoryUtil *VirtualServices_VirtualService_Utilization_MemoryUtil) GetEntityData() *types.CommonEntityData {
    memoryUtil.EntityData.YFilter = memoryUtil.YFilter
    memoryUtil.EntityData.YangName = "memory-util"
    memoryUtil.EntityData.BundleName = "cisco_ios_xe"
    memoryUtil.EntityData.ParentYangName = "utilization"
    memoryUtil.EntityData.SegmentPath = "memory-util"
    memoryUtil.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    memoryUtil.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    memoryUtil.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    memoryUtil.EntityData.Children = make(map[string]types.YChild)
    memoryUtil.EntityData.Leafs = make(map[string]types.YLeaf)
    memoryUtil.EntityData.Leafs["memory-allocation"] = types.YLeaf{"MemoryAllocation", memoryUtil.MemoryAllocation}
    memoryUtil.EntityData.Leafs["memory-used"] = types.YLeaf{"MemoryUsed", memoryUtil.MemoryUsed}
    return &(memoryUtil.EntityData)
}

// VirtualServices_VirtualService_NetworkUtils
// Virtual service network utilization details
type VirtualServices_VirtualService_NetworkUtils struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of network utilization details. The type is slice of
    // VirtualServices_VirtualService_NetworkUtils_NetworkUtil.
    NetworkUtil []VirtualServices_VirtualService_NetworkUtils_NetworkUtil
}

func (networkUtils *VirtualServices_VirtualService_NetworkUtils) GetEntityData() *types.CommonEntityData {
    networkUtils.EntityData.YFilter = networkUtils.YFilter
    networkUtils.EntityData.YangName = "network-utils"
    networkUtils.EntityData.BundleName = "cisco_ios_xe"
    networkUtils.EntityData.ParentYangName = "virtual-service"
    networkUtils.EntityData.SegmentPath = "network-utils"
    networkUtils.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    networkUtils.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    networkUtils.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    networkUtils.EntityData.Children = make(map[string]types.YChild)
    networkUtils.EntityData.Children["network-util"] = types.YChild{"NetworkUtil", nil}
    for i := range networkUtils.NetworkUtil {
        networkUtils.EntityData.Children[types.GetSegmentPath(&networkUtils.NetworkUtil[i])] = types.YChild{"NetworkUtil", &networkUtils.NetworkUtil[i]}
    }
    networkUtils.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(networkUtils.EntityData)
}

// VirtualServices_VirtualService_NetworkUtils_NetworkUtil
// A list of network utilization details
type VirtualServices_VirtualService_NetworkUtils_NetworkUtil struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the network used for the virtual service.
    // The type is string.
    Name interface{}

    // Alias of the network used by the virtual service. The type is string.
    Alias interface{}

    // Number of packets received by the virtual service. The type is interface{}
    // with range: 0..18446744073709551615.
    RxPackets interface{}

    // Number of octets received by the virtual service. The type is interface{}
    // with range: 0..18446744073709551615.
    RxBytes interface{}

    // Number of RX errors by the virtual service. The type is interface{} with
    // range: 0..18446744073709551615.
    RxErrors interface{}

    // Number of packets transmitted by the virtual service. The type is
    // interface{} with range: 0..18446744073709551615.
    TxPackets interface{}

    // Number of octets transmitted by the virtual service. The type is
    // interface{} with range: 0..18446744073709551615.
    TxBytes interface{}

    // Number of TX errors by the virtual service. The type is interface{} with
    // range: 0..18446744073709551615.
    TxErrors interface{}
}

func (networkUtil *VirtualServices_VirtualService_NetworkUtils_NetworkUtil) GetEntityData() *types.CommonEntityData {
    networkUtil.EntityData.YFilter = networkUtil.YFilter
    networkUtil.EntityData.YangName = "network-util"
    networkUtil.EntityData.BundleName = "cisco_ios_xe"
    networkUtil.EntityData.ParentYangName = "network-utils"
    networkUtil.EntityData.SegmentPath = "network-util" + "[name='" + fmt.Sprintf("%v", networkUtil.Name) + "']"
    networkUtil.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    networkUtil.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    networkUtil.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    networkUtil.EntityData.Children = make(map[string]types.YChild)
    networkUtil.EntityData.Leafs = make(map[string]types.YLeaf)
    networkUtil.EntityData.Leafs["name"] = types.YLeaf{"Name", networkUtil.Name}
    networkUtil.EntityData.Leafs["alias"] = types.YLeaf{"Alias", networkUtil.Alias}
    networkUtil.EntityData.Leafs["rx-packets"] = types.YLeaf{"RxPackets", networkUtil.RxPackets}
    networkUtil.EntityData.Leafs["rx-bytes"] = types.YLeaf{"RxBytes", networkUtil.RxBytes}
    networkUtil.EntityData.Leafs["rx-errors"] = types.YLeaf{"RxErrors", networkUtil.RxErrors}
    networkUtil.EntityData.Leafs["tx-packets"] = types.YLeaf{"TxPackets", networkUtil.TxPackets}
    networkUtil.EntityData.Leafs["tx-bytes"] = types.YLeaf{"TxBytes", networkUtil.TxBytes}
    networkUtil.EntityData.Leafs["tx-errors"] = types.YLeaf{"TxErrors", networkUtil.TxErrors}
    return &(networkUtil.EntityData)
}

// VirtualServices_VirtualService_StorageUtils
// Virtual service storage utilization details
type VirtualServices_VirtualService_StorageUtils struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of storage utilization details. The type is slice of
    // VirtualServices_VirtualService_StorageUtils_StorageUtil.
    StorageUtil []VirtualServices_VirtualService_StorageUtils_StorageUtil
}

func (storageUtils *VirtualServices_VirtualService_StorageUtils) GetEntityData() *types.CommonEntityData {
    storageUtils.EntityData.YFilter = storageUtils.YFilter
    storageUtils.EntityData.YangName = "storage-utils"
    storageUtils.EntityData.BundleName = "cisco_ios_xe"
    storageUtils.EntityData.ParentYangName = "virtual-service"
    storageUtils.EntityData.SegmentPath = "storage-utils"
    storageUtils.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    storageUtils.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    storageUtils.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    storageUtils.EntityData.Children = make(map[string]types.YChild)
    storageUtils.EntityData.Children["storage-util"] = types.YChild{"StorageUtil", nil}
    for i := range storageUtils.StorageUtil {
        storageUtils.EntityData.Children[types.GetSegmentPath(&storageUtils.StorageUtil[i])] = types.YChild{"StorageUtil", &storageUtils.StorageUtil[i]}
    }
    storageUtils.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(storageUtils.EntityData)
}

// VirtualServices_VirtualService_StorageUtils_StorageUtil
// List of storage utilization details
type VirtualServices_VirtualService_StorageUtils_StorageUtil struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the storage device used for the virtual
    // service. The type is string.
    Name interface{}

    // Alias of the storage device used by the virtual service. The type is
    // string.
    Alias interface{}

    // Number of bytes read by the virtual service. The type is interface{} with
    // range: 0..18446744073709551615.
    RdBytes interface{}

    // Number of read requests made by the virtual service. The type is
    // interface{} with range: 0..18446744073709551615.
    RdRequests interface{}

    // Number of storage error seen by the virtual service. The type is
    // interface{} with range: 0..18446744073709551615.
    Errors interface{}

    // Number of bytes written by the virtual service. The type is interface{}
    // with range: 0..18446744073709551615.
    WrBytes interface{}

    // Number of write requests made by the virtual service. The type is
    // interface{} with range: 0..18446744073709551615.
    WrRequests interface{}

    // Storage capactity in 1K blocks. The type is interface{} with range:
    // 0..18446744073709551615.
    Capacity interface{}

    // Available storage in 1K blocks. The type is string.
    Available interface{}

    // Used storage in 1K blocks. The type is interface{} with range:
    // 0..18446744073709551615.
    Used interface{}

    // Percetage of storage capactiy used by the virtual service. The type is
    // string.
    Usage interface{}
}

func (storageUtil *VirtualServices_VirtualService_StorageUtils_StorageUtil) GetEntityData() *types.CommonEntityData {
    storageUtil.EntityData.YFilter = storageUtil.YFilter
    storageUtil.EntityData.YangName = "storage-util"
    storageUtil.EntityData.BundleName = "cisco_ios_xe"
    storageUtil.EntityData.ParentYangName = "storage-utils"
    storageUtil.EntityData.SegmentPath = "storage-util" + "[name='" + fmt.Sprintf("%v", storageUtil.Name) + "']"
    storageUtil.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    storageUtil.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    storageUtil.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    storageUtil.EntityData.Children = make(map[string]types.YChild)
    storageUtil.EntityData.Leafs = make(map[string]types.YLeaf)
    storageUtil.EntityData.Leafs["name"] = types.YLeaf{"Name", storageUtil.Name}
    storageUtil.EntityData.Leafs["alias"] = types.YLeaf{"Alias", storageUtil.Alias}
    storageUtil.EntityData.Leafs["rd-bytes"] = types.YLeaf{"RdBytes", storageUtil.RdBytes}
    storageUtil.EntityData.Leafs["rd-requests"] = types.YLeaf{"RdRequests", storageUtil.RdRequests}
    storageUtil.EntityData.Leafs["errors"] = types.YLeaf{"Errors", storageUtil.Errors}
    storageUtil.EntityData.Leafs["wr-bytes"] = types.YLeaf{"WrBytes", storageUtil.WrBytes}
    storageUtil.EntityData.Leafs["wr-requests"] = types.YLeaf{"WrRequests", storageUtil.WrRequests}
    storageUtil.EntityData.Leafs["capacity"] = types.YLeaf{"Capacity", storageUtil.Capacity}
    storageUtil.EntityData.Leafs["available"] = types.YLeaf{"Available", storageUtil.Available}
    storageUtil.EntityData.Leafs["used"] = types.YLeaf{"Used", storageUtil.Used}
    storageUtil.EntityData.Leafs["usage"] = types.YLeaf{"Usage", storageUtil.Usage}
    return &(storageUtil.EntityData)
}

// VirtualServices_VirtualService_Processes
// Virtual service process details
type VirtualServices_VirtualService_Processes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of process details. The type is slice of
    // VirtualServices_VirtualService_Processes_Process.
    Process []VirtualServices_VirtualService_Processes_Process
}

func (processes *VirtualServices_VirtualService_Processes) GetEntityData() *types.CommonEntityData {
    processes.EntityData.YFilter = processes.YFilter
    processes.EntityData.YangName = "processes"
    processes.EntityData.BundleName = "cisco_ios_xe"
    processes.EntityData.ParentYangName = "virtual-service"
    processes.EntityData.SegmentPath = "processes"
    processes.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    processes.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    processes.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    processes.EntityData.Children = make(map[string]types.YChild)
    processes.EntityData.Children["process"] = types.YChild{"Process", nil}
    for i := range processes.Process {
        processes.EntityData.Children[types.GetSegmentPath(&processes.Process[i])] = types.YChild{"Process", &processes.Process[i]}
    }
    processes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(processes.EntityData)
}

// VirtualServices_VirtualService_Processes_Process
// List of process details
type VirtualServices_VirtualService_Processes_Process struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Process name. The type is string.
    Name interface{}

    // Process status. The type is string.
    Status interface{}

    // Process ID. The type is string.
    Pid interface{}

    // Process uptime. The type is string.
    Uptime interface{}

    // Amount of process memory. The type is string.
    Memory interface{}
}

func (process *VirtualServices_VirtualService_Processes_Process) GetEntityData() *types.CommonEntityData {
    process.EntityData.YFilter = process.YFilter
    process.EntityData.YangName = "process"
    process.EntityData.BundleName = "cisco_ios_xe"
    process.EntityData.ParentYangName = "processes"
    process.EntityData.SegmentPath = "process" + "[name='" + fmt.Sprintf("%v", process.Name) + "']"
    process.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    process.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    process.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    process.EntityData.Children = make(map[string]types.YChild)
    process.EntityData.Leafs = make(map[string]types.YLeaf)
    process.EntityData.Leafs["name"] = types.YLeaf{"Name", process.Name}
    process.EntityData.Leafs["status"] = types.YLeaf{"Status", process.Status}
    process.EntityData.Leafs["pid"] = types.YLeaf{"Pid", process.Pid}
    process.EntityData.Leafs["uptime"] = types.YLeaf{"Uptime", process.Uptime}
    process.EntityData.Leafs["memory"] = types.YLeaf{"Memory", process.Memory}
    return &(process.EntityData)
}

// VirtualServices_VirtualService_AttachedDevices
// Virtual service attached device details
type VirtualServices_VirtualService_AttachedDevices struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of attached device details. The type is slice of
    // VirtualServices_VirtualService_AttachedDevices_AttachedDevice.
    AttachedDevice []VirtualServices_VirtualService_AttachedDevices_AttachedDevice
}

func (attachedDevices *VirtualServices_VirtualService_AttachedDevices) GetEntityData() *types.CommonEntityData {
    attachedDevices.EntityData.YFilter = attachedDevices.YFilter
    attachedDevices.EntityData.YangName = "attached-devices"
    attachedDevices.EntityData.BundleName = "cisco_ios_xe"
    attachedDevices.EntityData.ParentYangName = "virtual-service"
    attachedDevices.EntityData.SegmentPath = "attached-devices"
    attachedDevices.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    attachedDevices.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    attachedDevices.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    attachedDevices.EntityData.Children = make(map[string]types.YChild)
    attachedDevices.EntityData.Children["attached-device"] = types.YChild{"AttachedDevice", nil}
    for i := range attachedDevices.AttachedDevice {
        attachedDevices.EntityData.Children[types.GetSegmentPath(&attachedDevices.AttachedDevice[i])] = types.YChild{"AttachedDevice", &attachedDevices.AttachedDevice[i]}
    }
    attachedDevices.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(attachedDevices.EntityData)
}

// VirtualServices_VirtualService_AttachedDevices_AttachedDevice
// A list of attached device details
type VirtualServices_VirtualService_AttachedDevices_AttachedDevice struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Attached device name. The type is string.
    Name interface{}

    // Attached device type. The type is string.
    Type_ interface{}

    // Attached device alias. The type is string.
    Alias interface{}
}

func (attachedDevice *VirtualServices_VirtualService_AttachedDevices_AttachedDevice) GetEntityData() *types.CommonEntityData {
    attachedDevice.EntityData.YFilter = attachedDevice.YFilter
    attachedDevice.EntityData.YangName = "attached-device"
    attachedDevice.EntityData.BundleName = "cisco_ios_xe"
    attachedDevice.EntityData.ParentYangName = "attached-devices"
    attachedDevice.EntityData.SegmentPath = "attached-device" + "[name='" + fmt.Sprintf("%v", attachedDevice.Name) + "']"
    attachedDevice.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    attachedDevice.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    attachedDevice.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    attachedDevice.EntityData.Children = make(map[string]types.YChild)
    attachedDevice.EntityData.Leafs = make(map[string]types.YLeaf)
    attachedDevice.EntityData.Leafs["name"] = types.YLeaf{"Name", attachedDevice.Name}
    attachedDevice.EntityData.Leafs["type"] = types.YLeaf{"Type_", attachedDevice.Type_}
    attachedDevice.EntityData.Leafs["alias"] = types.YLeaf{"Alias", attachedDevice.Alias}
    return &(attachedDevice.EntityData)
}

// VirtualServices_VirtualService_NetworkInterfaces
// Virtual service network interface details
type VirtualServices_VirtualService_NetworkInterfaces struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of network interface details. The type is slice of
    // VirtualServices_VirtualService_NetworkInterfaces_NetworkInterface.
    NetworkInterface []VirtualServices_VirtualService_NetworkInterfaces_NetworkInterface
}

func (networkInterfaces *VirtualServices_VirtualService_NetworkInterfaces) GetEntityData() *types.CommonEntityData {
    networkInterfaces.EntityData.YFilter = networkInterfaces.YFilter
    networkInterfaces.EntityData.YangName = "network-interfaces"
    networkInterfaces.EntityData.BundleName = "cisco_ios_xe"
    networkInterfaces.EntityData.ParentYangName = "virtual-service"
    networkInterfaces.EntityData.SegmentPath = "network-interfaces"
    networkInterfaces.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    networkInterfaces.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    networkInterfaces.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    networkInterfaces.EntityData.Children = make(map[string]types.YChild)
    networkInterfaces.EntityData.Children["network-interface"] = types.YChild{"NetworkInterface", nil}
    for i := range networkInterfaces.NetworkInterface {
        networkInterfaces.EntityData.Children[types.GetSegmentPath(&networkInterfaces.NetworkInterface[i])] = types.YChild{"NetworkInterface", &networkInterfaces.NetworkInterface[i]}
    }
    networkInterfaces.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(networkInterfaces.EntityData)
}

// VirtualServices_VirtualService_NetworkInterfaces_NetworkInterface
// A list of network interface details
type VirtualServices_VirtualService_NetworkInterfaces_NetworkInterface struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. MAC address for the network interface. The type is
    // string with pattern: b'[0-9a-fA-F]{2}(:[0-9a-fA-F]{2}){5}'.
    MacAddress interface{}

    // Attached interface name. The type is string.
    AttachedInterface interface{}

    // IPv4 address for the network interface. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4Address interface{}
}

func (networkInterface *VirtualServices_VirtualService_NetworkInterfaces_NetworkInterface) GetEntityData() *types.CommonEntityData {
    networkInterface.EntityData.YFilter = networkInterface.YFilter
    networkInterface.EntityData.YangName = "network-interface"
    networkInterface.EntityData.BundleName = "cisco_ios_xe"
    networkInterface.EntityData.ParentYangName = "network-interfaces"
    networkInterface.EntityData.SegmentPath = "network-interface" + "[mac-address='" + fmt.Sprintf("%v", networkInterface.MacAddress) + "']"
    networkInterface.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    networkInterface.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    networkInterface.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    networkInterface.EntityData.Children = make(map[string]types.YChild)
    networkInterface.EntityData.Leafs = make(map[string]types.YLeaf)
    networkInterface.EntityData.Leafs["mac-address"] = types.YLeaf{"MacAddress", networkInterface.MacAddress}
    networkInterface.EntityData.Leafs["attached-interface"] = types.YLeaf{"AttachedInterface", networkInterface.AttachedInterface}
    networkInterface.EntityData.Leafs["ipv4-address"] = types.YLeaf{"Ipv4Address", networkInterface.Ipv4Address}
    return &(networkInterface.EntityData)
}

// VirtualServices_VirtualService_GuestRoutes
// Virtual service guest route details
type VirtualServices_VirtualService_GuestRoutes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of guest routes for a guest interface. The type is slice of
    // VirtualServices_VirtualService_GuestRoutes_GuestRoute.
    GuestRoute []VirtualServices_VirtualService_GuestRoutes_GuestRoute
}

func (guestRoutes *VirtualServices_VirtualService_GuestRoutes) GetEntityData() *types.CommonEntityData {
    guestRoutes.EntityData.YFilter = guestRoutes.YFilter
    guestRoutes.EntityData.YangName = "guest-routes"
    guestRoutes.EntityData.BundleName = "cisco_ios_xe"
    guestRoutes.EntityData.ParentYangName = "virtual-service"
    guestRoutes.EntityData.SegmentPath = "guest-routes"
    guestRoutes.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    guestRoutes.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    guestRoutes.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    guestRoutes.EntityData.Children = make(map[string]types.YChild)
    guestRoutes.EntityData.Children["guest-route"] = types.YChild{"GuestRoute", nil}
    for i := range guestRoutes.GuestRoute {
        guestRoutes.EntityData.Children[types.GetSegmentPath(&guestRoutes.GuestRoute[i])] = types.YChild{"GuestRoute", &guestRoutes.GuestRoute[i]}
    }
    guestRoutes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(guestRoutes.EntityData)
}

// VirtualServices_VirtualService_GuestRoutes_GuestRoute
// List of guest routes for a guest interface
type VirtualServices_VirtualService_GuestRoutes_GuestRoute struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Guest route of the guest interface. The type is
    // string.
    Route interface{}
}

func (guestRoute *VirtualServices_VirtualService_GuestRoutes_GuestRoute) GetEntityData() *types.CommonEntityData {
    guestRoute.EntityData.YFilter = guestRoute.YFilter
    guestRoute.EntityData.YangName = "guest-route"
    guestRoute.EntityData.BundleName = "cisco_ios_xe"
    guestRoute.EntityData.ParentYangName = "guest-routes"
    guestRoute.EntityData.SegmentPath = "guest-route" + "[route='" + fmt.Sprintf("%v", guestRoute.Route) + "']"
    guestRoute.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    guestRoute.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    guestRoute.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    guestRoute.EntityData.Children = make(map[string]types.YChild)
    guestRoute.EntityData.Leafs = make(map[string]types.YLeaf)
    guestRoute.EntityData.Leafs["route"] = types.YLeaf{"Route", guestRoute.Route}
    return &(guestRoute.EntityData)
}

