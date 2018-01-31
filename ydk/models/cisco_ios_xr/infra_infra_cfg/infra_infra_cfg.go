// This module contains a collection of YANG definitions
// for Cisco IOS-XR infra-infra package configuration.
// 
// This module contains definitions
// for the following management objects:
//   banners: Schema for Banner configuration commands
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // Select a Banner Type. The type is slice of Banners_Banner.
    Banner []Banners_Banner
}

func (banners *Banners) GetFilter() yfilter.YFilter { return banners.YFilter }

func (banners *Banners) SetFilter(yf yfilter.YFilter) { banners.YFilter = yf }

func (banners *Banners) GetGoName(yname string) string {
    if yname == "banner" { return "Banner" }
    return ""
}

func (banners *Banners) GetSegmentPath() string {
    return "Cisco-IOS-XR-infra-infra-cfg:banners"
}

func (banners *Banners) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "banner" {
        for _, c := range banners.Banner {
            if banners.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Banners_Banner{}
        banners.Banner = append(banners.Banner, child)
        return &banners.Banner[len(banners.Banner)-1]
    }
    return nil
}

func (banners *Banners) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range banners.Banner {
        children[banners.Banner[i].GetSegmentPath()] = &banners.Banner[i]
    }
    return children
}

func (banners *Banners) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (banners *Banners) GetBundleName() string { return "cisco_ios_xr" }

func (banners *Banners) GetYangName() string { return "banners" }

func (banners *Banners) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (banners *Banners) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (banners *Banners) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (banners *Banners) SetParent(parent types.Entity) { banners.parent = parent }

func (banners *Banners) GetParent() types.Entity { return banners.parent }

func (banners *Banners) GetParentYangName() string { return "Cisco-IOS-XR-infra-infra-cfg" }

// Banners_Banner
// Select a Banner Type
type Banners_Banner struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Banner Type. The type is Banner.
    BannerName interface{}

    // Banner text message. The type is string. This attribute is mandatory.
    BannerText interface{}
}

func (banner *Banners_Banner) GetFilter() yfilter.YFilter { return banner.YFilter }

func (banner *Banners_Banner) SetFilter(yf yfilter.YFilter) { banner.YFilter = yf }

func (banner *Banners_Banner) GetGoName(yname string) string {
    if yname == "banner-name" { return "BannerName" }
    if yname == "banner-text" { return "BannerText" }
    return ""
}

func (banner *Banners_Banner) GetSegmentPath() string {
    return "banner" + "[banner-name='" + fmt.Sprintf("%v", banner.BannerName) + "']"
}

func (banner *Banners_Banner) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (banner *Banners_Banner) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (banner *Banners_Banner) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["banner-name"] = banner.BannerName
    leafs["banner-text"] = banner.BannerText
    return leafs
}

func (banner *Banners_Banner) GetBundleName() string { return "cisco_ios_xr" }

func (banner *Banners_Banner) GetYangName() string { return "banner" }

func (banner *Banners_Banner) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (banner *Banners_Banner) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (banner *Banners_Banner) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (banner *Banners_Banner) SetParent(parent types.Entity) { banner.parent = parent }

func (banner *Banners_Banner) GetParent() types.Entity { return banner.parent }

func (banner *Banners_Banner) GetParentYangName() string { return "banners" }

