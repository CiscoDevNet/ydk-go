// This module contains definitions
// for the Calvados model objects.
// 
// This module contains a collection of YANG
// definitions for Cisco IOS-XR SysAdmin configuration.
// 
// This module defines the services offered in the
// sysadmin plane.
// 
// Copyright(c) 2016 by Cisco Systems, Inc.
// All rights reserved.
// 
// Copyright (c) 2012-2018 by Cisco Systems, Inc.
// All rights reserved.
package sysadmin_services

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_services"))
    ydk.RegisterEntity("{http://www.cisco.com/ns/yang/Cisco-IOS-XR-sysadmin-services service}", reflect.TypeOf(Service{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-services:service", reflect.TypeOf(Service{}))
}

// Service
type Service struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Cli Service_Cli
}

func (service *Service) GetEntityData() *types.CommonEntityData {
    service.EntityData.YFilter = service.YFilter
    service.EntityData.YangName = "service"
    service.EntityData.BundleName = "cisco_ios_xr"
    service.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-services"
    service.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-services:service"
    service.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    service.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    service.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    service.EntityData.Children = types.NewOrderedMap()
    service.EntityData.Children.Append("cli", types.YChild{"Cli", &service.Cli})
    service.EntityData.Leafs = types.NewOrderedMap()

    service.EntityData.YListKeys = []string {}

    return &(service.EntityData)
}

// Service_Cli
type Service_Cli struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    Interactive Service_Cli_Interactive
}

func (cli *Service_Cli) GetEntityData() *types.CommonEntityData {
    cli.EntityData.YFilter = cli.YFilter
    cli.EntityData.YangName = "cli"
    cli.EntityData.BundleName = "cisco_ios_xr"
    cli.EntityData.ParentYangName = "service"
    cli.EntityData.SegmentPath = "cli"
    cli.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    cli.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    cli.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    cli.EntityData.Children = types.NewOrderedMap()
    cli.EntityData.Children.Append("interactive", types.YChild{"Interactive", &cli.Interactive})
    cli.EntityData.Leafs = types.NewOrderedMap()

    cli.EntityData.YListKeys = []string {}

    return &(cli.EntityData)
}

// Service_Cli_Interactive
type Service_Cli_Interactive struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is bool. The default value is true.
    Enabled interface{}
}

func (interactive *Service_Cli_Interactive) GetEntityData() *types.CommonEntityData {
    interactive.EntityData.YFilter = interactive.YFilter
    interactive.EntityData.YangName = "interactive"
    interactive.EntityData.BundleName = "cisco_ios_xr"
    interactive.EntityData.ParentYangName = "cli"
    interactive.EntityData.SegmentPath = "interactive"
    interactive.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    interactive.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    interactive.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    interactive.EntityData.Children = types.NewOrderedMap()
    interactive.EntityData.Leafs = types.NewOrderedMap()
    interactive.EntityData.Leafs.Append("enabled", types.YLeaf{"Enabled", interactive.Enabled})

    interactive.EntityData.YListKeys = []string {}

    return &(interactive.EntityData)
}

