// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-infra package configuration.
// 
// This module contains definitions
// for the following management objects:
//   banners: Schema for Banner configuration commands
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package infra_infra_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package infra_infra_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-infra-infra-cfg banners}", reflect.TypeOf(Banners{}))
    ydk.RegisterEntity("Cisco-IOS-XR-infra-infra-cfg:banners", reflect.TypeOf(Banners{}))
}

// Banner represents Banner
type Banner string

const (
    // Set EXEC process creation banner
    Banner_exec Banner = "exec"

    // Set incoming terminal line banner
    Banner_incoming Banner = "incoming"

    // Set Message of the Day banner
    Banner_motd Banner = "motd"

    // Set login banner
    Banner_login Banner = "login"

    // Set Message for SLIP/PPP
    Banner_slip_ppp Banner = "slip-ppp"

    // Set Message for login authentication timeout
    Banner_prompt_timeout Banner = "prompt-timeout"
)

// Banners
// Schema for Banner configuration commands
type Banners struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Select a Banner Type. The type is slice of Banners_Banner.
    Banner []*Banners_Banner
}

func (banners *Banners) GetEntityData() *types.CommonEntityData {
    banners.EntityData.YFilter = banners.YFilter
    banners.EntityData.YangName = "banners"
    banners.EntityData.BundleName = "cisco_ios_xr"
    banners.EntityData.ParentYangName = "Cisco-IOS-XR-infra-infra-cfg"
    banners.EntityData.SegmentPath = "Cisco-IOS-XR-infra-infra-cfg:banners"
    banners.EntityData.AbsolutePath = banners.EntityData.SegmentPath
    banners.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    banners.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    banners.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    banners.EntityData.Children = types.NewOrderedMap()
    banners.EntityData.Children.Append("banner", types.YChild{"Banner", nil})
    for i := range banners.Banner {
        banners.EntityData.Children.Append(types.GetSegmentPath(banners.Banner[i]), types.YChild{"Banner", banners.Banner[i]})
    }
    banners.EntityData.Leafs = types.NewOrderedMap()

    banners.EntityData.YListKeys = []string {}

    return &(banners.EntityData)
}

// Banners_Banner
// Select a Banner Type
type Banners_Banner struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Banner Type. The type is Banner.
    BannerName interface{}

    // Banner text message. The type is string. This attribute is mandatory.
    BannerText interface{}
}

func (banner *Banners_Banner) GetEntityData() *types.CommonEntityData {
    banner.EntityData.YFilter = banner.YFilter
    banner.EntityData.YangName = "banner"
    banner.EntityData.BundleName = "cisco_ios_xr"
    banner.EntityData.ParentYangName = "banners"
    banner.EntityData.SegmentPath = "banner" + types.AddKeyToken(banner.BannerName, "banner-name")
    banner.EntityData.AbsolutePath = "Cisco-IOS-XR-infra-infra-cfg:banners/" + banner.EntityData.SegmentPath
    banner.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    banner.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    banner.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    banner.EntityData.Children = types.NewOrderedMap()
    banner.EntityData.Leafs = types.NewOrderedMap()
    banner.EntityData.Leafs.Append("banner-name", types.YLeaf{"BannerName", banner.BannerName})
    banner.EntityData.Leafs.Append("banner-text", types.YLeaf{"BannerText", banner.BannerText})

    banner.EntityData.YListKeys = []string {"BannerName"}

    return &(banner.EntityData)
}

