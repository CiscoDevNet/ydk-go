// This module contains a collection of YANG definitions
// for Cisco IOS-XR mpls-oam package configuration.
// 
// This module contains definitions
// for the following management objects:
//   mpls-oam: MPLS LSP verification configuration
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package mpls_oam_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package mpls_oam_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-mpls-oam-cfg mpls-oam}", reflect.TypeOf(MplsOam{}))
    ydk.RegisterEntity("Cisco-IOS-XR-mpls-oam-cfg:mpls-oam", reflect.TypeOf(MplsOam{}))
}

// MplsOam
// MPLS LSP verification configuration
type MplsOam struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Enable/Disable MPLS OAM globally.Without creating this object the MPLS OAM
    // feature will not be enabled. Deleting this object will stop the MPLS OAM
    // feature. The type is interface{}.
    EnableOam interface{}

    // Disable vendor extension. The type is interface{}.
    DisableVendorExtension interface{}

    // Echo request reply mode attributes.
    ReplyMode MplsOam_ReplyMode
}

func (mplsOam *MplsOam) GetFilter() yfilter.YFilter { return mplsOam.YFilter }

func (mplsOam *MplsOam) SetFilter(yf yfilter.YFilter) { mplsOam.YFilter = yf }

func (mplsOam *MplsOam) GetGoName(yname string) string {
    if yname == "enable-oam" { return "EnableOam" }
    if yname == "disable-vendor-extension" { return "DisableVendorExtension" }
    if yname == "reply-mode" { return "ReplyMode" }
    return ""
}

func (mplsOam *MplsOam) GetSegmentPath() string {
    return "Cisco-IOS-XR-mpls-oam-cfg:mpls-oam"
}

func (mplsOam *MplsOam) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "reply-mode" {
        return &mplsOam.ReplyMode
    }
    return nil
}

func (mplsOam *MplsOam) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["reply-mode"] = &mplsOam.ReplyMode
    return children
}

func (mplsOam *MplsOam) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["enable-oam"] = mplsOam.EnableOam
    leafs["disable-vendor-extension"] = mplsOam.DisableVendorExtension
    return leafs
}

func (mplsOam *MplsOam) GetBundleName() string { return "cisco_ios_xr" }

func (mplsOam *MplsOam) GetYangName() string { return "mpls-oam" }

func (mplsOam *MplsOam) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (mplsOam *MplsOam) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (mplsOam *MplsOam) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (mplsOam *MplsOam) SetParent(parent types.Entity) { mplsOam.parent = parent }

func (mplsOam *MplsOam) GetParent() types.Entity { return mplsOam.parent }

func (mplsOam *MplsOam) GetParentYangName() string { return "Cisco-IOS-XR-mpls-oam-cfg" }

// MplsOam_ReplyMode
// Echo request reply mode attributes
type MplsOam_ReplyMode struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Configure control channel reply mode.
    ControlChannel MplsOam_ReplyMode_ControlChannel
}

func (replyMode *MplsOam_ReplyMode) GetFilter() yfilter.YFilter { return replyMode.YFilter }

func (replyMode *MplsOam_ReplyMode) SetFilter(yf yfilter.YFilter) { replyMode.YFilter = yf }

func (replyMode *MplsOam_ReplyMode) GetGoName(yname string) string {
    if yname == "control-channel" { return "ControlChannel" }
    return ""
}

func (replyMode *MplsOam_ReplyMode) GetSegmentPath() string {
    return "reply-mode"
}

func (replyMode *MplsOam_ReplyMode) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "control-channel" {
        return &replyMode.ControlChannel
    }
    return nil
}

func (replyMode *MplsOam_ReplyMode) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["control-channel"] = &replyMode.ControlChannel
    return children
}

func (replyMode *MplsOam_ReplyMode) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (replyMode *MplsOam_ReplyMode) GetBundleName() string { return "cisco_ios_xr" }

func (replyMode *MplsOam_ReplyMode) GetYangName() string { return "reply-mode" }

func (replyMode *MplsOam_ReplyMode) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (replyMode *MplsOam_ReplyMode) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (replyMode *MplsOam_ReplyMode) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (replyMode *MplsOam_ReplyMode) SetParent(parent types.Entity) { replyMode.parent = parent }

func (replyMode *MplsOam_ReplyMode) GetParent() types.Entity { return replyMode.parent }

func (replyMode *MplsOam_ReplyMode) GetParentYangName() string { return "mpls-oam" }

// MplsOam_ReplyMode_ControlChannel
// Configure control channel reply mode
type MplsOam_ReplyMode_ControlChannel struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Use Reverse LSP as the control channel. The type is interface{}.
    AllowReverseLsp interface{}
}

func (controlChannel *MplsOam_ReplyMode_ControlChannel) GetFilter() yfilter.YFilter { return controlChannel.YFilter }

func (controlChannel *MplsOam_ReplyMode_ControlChannel) SetFilter(yf yfilter.YFilter) { controlChannel.YFilter = yf }

func (controlChannel *MplsOam_ReplyMode_ControlChannel) GetGoName(yname string) string {
    if yname == "allow-reverse-lsp" { return "AllowReverseLsp" }
    return ""
}

func (controlChannel *MplsOam_ReplyMode_ControlChannel) GetSegmentPath() string {
    return "control-channel"
}

func (controlChannel *MplsOam_ReplyMode_ControlChannel) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (controlChannel *MplsOam_ReplyMode_ControlChannel) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (controlChannel *MplsOam_ReplyMode_ControlChannel) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["allow-reverse-lsp"] = controlChannel.AllowReverseLsp
    return leafs
}

func (controlChannel *MplsOam_ReplyMode_ControlChannel) GetBundleName() string { return "cisco_ios_xr" }

func (controlChannel *MplsOam_ReplyMode_ControlChannel) GetYangName() string { return "control-channel" }

func (controlChannel *MplsOam_ReplyMode_ControlChannel) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (controlChannel *MplsOam_ReplyMode_ControlChannel) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (controlChannel *MplsOam_ReplyMode_ControlChannel) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (controlChannel *MplsOam_ReplyMode_ControlChannel) SetParent(parent types.Entity) { controlChannel.parent = parent }

func (controlChannel *MplsOam_ReplyMode_ControlChannel) GetParent() types.Entity { return controlChannel.parent }

func (controlChannel *MplsOam_ReplyMode_ControlChannel) GetParentYangName() string { return "reply-mode" }

