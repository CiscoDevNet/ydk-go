// This module contains a collection of YANG definitions
// for Cisco IOS-XR pbr-bng package configuration.
// 
// This module contains definitions
// for the following management objects:
//   bng-pbr: Subscriber PBR configuration
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package pbr_bng_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package pbr_bng_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-pbr-bng-cfg bng-pbr}", reflect.TypeOf(BngPbr{}))
    ydk.RegisterEntity("Cisco-IOS-XR-pbr-bng-cfg:bng-pbr", reflect.TypeOf(BngPbr{}))
}

// BngPbrHttpEnrichmentParams represents Bng pbr http enrichment params
type BngPbrHttpEnrichmentParams string

const (
    // Subscriber Mac
    BngPbrHttpEnrichmentParams_subscriber_mac BngPbrHttpEnrichmentParams = "subscriber-mac"

    // Subscriber IPv4/IPv6 address
    BngPbrHttpEnrichmentParams_subscriber_ip BngPbrHttpEnrichmentParams = "subscriber-ip"

    // Bng Router Hostname
    BngPbrHttpEnrichmentParams_host_name BngPbrHttpEnrichmentParams = "host-name"

    // Bng Identifier interface
    BngPbrHttpEnrichmentParams_bng_identifier_interface BngPbrHttpEnrichmentParams = "bng-identifier-interface"
)

// BngPbr
// Subscriber PBR configuration
type BngPbr struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interface for source address. The type is string with pattern:
    // [a-zA-Z0-9._/-]+.
    BngInterface interface{}

    // HTTP Enrichment.
    HttpEnrichment BngPbr_HttpEnrichment
}

func (bngPbr *BngPbr) GetEntityData() *types.CommonEntityData {
    bngPbr.EntityData.YFilter = bngPbr.YFilter
    bngPbr.EntityData.YangName = "bng-pbr"
    bngPbr.EntityData.BundleName = "cisco_ios_xr"
    bngPbr.EntityData.ParentYangName = "Cisco-IOS-XR-pbr-bng-cfg"
    bngPbr.EntityData.SegmentPath = "Cisco-IOS-XR-pbr-bng-cfg:bng-pbr"
    bngPbr.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    bngPbr.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    bngPbr.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    bngPbr.EntityData.Children = types.NewOrderedMap()
    bngPbr.EntityData.Children.Append("http-enrichment", types.YChild{"HttpEnrichment", &bngPbr.HttpEnrichment})
    bngPbr.EntityData.Leafs = types.NewOrderedMap()
    bngPbr.EntityData.Leafs.Append("bng-interface", types.YLeaf{"BngInterface", bngPbr.BngInterface})

    bngPbr.EntityData.YListKeys = []string {}

    return &(bngPbr.EntityData)
}

// BngPbr_HttpEnrichment
// HTTP Enrichment
type BngPbr_HttpEnrichment struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // HTTP Enrichment parameters.
    Parameters BngPbr_HttpEnrichment_Parameters
}

func (httpEnrichment *BngPbr_HttpEnrichment) GetEntityData() *types.CommonEntityData {
    httpEnrichment.EntityData.YFilter = httpEnrichment.YFilter
    httpEnrichment.EntityData.YangName = "http-enrichment"
    httpEnrichment.EntityData.BundleName = "cisco_ios_xr"
    httpEnrichment.EntityData.ParentYangName = "bng-pbr"
    httpEnrichment.EntityData.SegmentPath = "http-enrichment"
    httpEnrichment.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    httpEnrichment.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    httpEnrichment.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    httpEnrichment.EntityData.Children = types.NewOrderedMap()
    httpEnrichment.EntityData.Children.Append("parameters", types.YChild{"Parameters", &httpEnrichment.Parameters})
    httpEnrichment.EntityData.Leafs = types.NewOrderedMap()

    httpEnrichment.EntityData.YListKeys = []string {}

    return &(httpEnrichment.EntityData)
}

// BngPbr_HttpEnrichment_Parameters
// HTTP Enrichment parameters
// This type is a presence type.
type BngPbr_HttpEnrichment_Parameters struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // first argument . The type is BngPbrHttpEnrichmentParams. This attribute is
    // mandatory.
    Arg1 interface{}

    // second argument . The type is BngPbrHttpEnrichmentParams.
    Arg2 interface{}

    // Third argument . The type is BngPbrHttpEnrichmentParams.
    Arg3 interface{}

    // Fourth argument . The type is BngPbrHttpEnrichmentParams.
    Arg4 interface{}
}

func (parameters *BngPbr_HttpEnrichment_Parameters) GetEntityData() *types.CommonEntityData {
    parameters.EntityData.YFilter = parameters.YFilter
    parameters.EntityData.YangName = "parameters"
    parameters.EntityData.BundleName = "cisco_ios_xr"
    parameters.EntityData.ParentYangName = "http-enrichment"
    parameters.EntityData.SegmentPath = "parameters"
    parameters.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    parameters.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    parameters.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    parameters.EntityData.Children = types.NewOrderedMap()
    parameters.EntityData.Leafs = types.NewOrderedMap()
    parameters.EntityData.Leafs.Append("arg1", types.YLeaf{"Arg1", parameters.Arg1})
    parameters.EntityData.Leafs.Append("arg2", types.YLeaf{"Arg2", parameters.Arg2})
    parameters.EntityData.Leafs.Append("arg3", types.YLeaf{"Arg3", parameters.Arg3})
    parameters.EntityData.Leafs.Append("arg4", types.YLeaf{"Arg4", parameters.Arg4})

    parameters.EntityData.YListKeys = []string {}

    return &(parameters.EntityData)
}

