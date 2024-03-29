// This module contains a collection of YANG definitions
// for Cisco IOS-XR mpls-oam package configuration.
// 
// This module contains definitions
// for the following management objects:
//   mpls-oam: MPLS LSP verification configuration
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
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
    EntityData types.CommonEntityData
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

func (mplsOam *MplsOam) GetEntityData() *types.CommonEntityData {
    mplsOam.EntityData.YFilter = mplsOam.YFilter
    mplsOam.EntityData.YangName = "mpls-oam"
    mplsOam.EntityData.BundleName = "cisco_ios_xr"
    mplsOam.EntityData.ParentYangName = "Cisco-IOS-XR-mpls-oam-cfg"
    mplsOam.EntityData.SegmentPath = "Cisco-IOS-XR-mpls-oam-cfg:mpls-oam"
    mplsOam.EntityData.AbsolutePath = mplsOam.EntityData.SegmentPath
    mplsOam.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    mplsOam.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    mplsOam.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    mplsOam.EntityData.Children = types.NewOrderedMap()
    mplsOam.EntityData.Children.Append("reply-mode", types.YChild{"ReplyMode", &mplsOam.ReplyMode})
    mplsOam.EntityData.Leafs = types.NewOrderedMap()
    mplsOam.EntityData.Leafs.Append("enable-oam", types.YLeaf{"EnableOam", mplsOam.EnableOam})
    mplsOam.EntityData.Leafs.Append("disable-vendor-extension", types.YLeaf{"DisableVendorExtension", mplsOam.DisableVendorExtension})

    mplsOam.EntityData.YListKeys = []string {}

    return &(mplsOam.EntityData)
}

// MplsOam_ReplyMode
// Echo request reply mode attributes
type MplsOam_ReplyMode struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configure control channel reply mode.
    ControlChannel MplsOam_ReplyMode_ControlChannel
}

func (replyMode *MplsOam_ReplyMode) GetEntityData() *types.CommonEntityData {
    replyMode.EntityData.YFilter = replyMode.YFilter
    replyMode.EntityData.YangName = "reply-mode"
    replyMode.EntityData.BundleName = "cisco_ios_xr"
    replyMode.EntityData.ParentYangName = "mpls-oam"
    replyMode.EntityData.SegmentPath = "reply-mode"
    replyMode.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-oam-cfg:mpls-oam/" + replyMode.EntityData.SegmentPath
    replyMode.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    replyMode.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    replyMode.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    replyMode.EntityData.Children = types.NewOrderedMap()
    replyMode.EntityData.Children.Append("control-channel", types.YChild{"ControlChannel", &replyMode.ControlChannel})
    replyMode.EntityData.Leafs = types.NewOrderedMap()

    replyMode.EntityData.YListKeys = []string {}

    return &(replyMode.EntityData)
}

// MplsOam_ReplyMode_ControlChannel
// Configure control channel reply mode
type MplsOam_ReplyMode_ControlChannel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Use Reverse LSP as the control channel. The type is interface{}.
    AllowReverseLsp interface{}
}

func (controlChannel *MplsOam_ReplyMode_ControlChannel) GetEntityData() *types.CommonEntityData {
    controlChannel.EntityData.YFilter = controlChannel.YFilter
    controlChannel.EntityData.YangName = "control-channel"
    controlChannel.EntityData.BundleName = "cisco_ios_xr"
    controlChannel.EntityData.ParentYangName = "reply-mode"
    controlChannel.EntityData.SegmentPath = "control-channel"
    controlChannel.EntityData.AbsolutePath = "Cisco-IOS-XR-mpls-oam-cfg:mpls-oam/reply-mode/" + controlChannel.EntityData.SegmentPath
    controlChannel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    controlChannel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    controlChannel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    controlChannel.EntityData.Children = types.NewOrderedMap()
    controlChannel.EntityData.Leafs = types.NewOrderedMap()
    controlChannel.EntityData.Leafs.Append("allow-reverse-lsp", types.YLeaf{"AllowReverseLsp", controlChannel.AllowReverseLsp})

    controlChannel.EntityData.YListKeys = []string {}

    return &(controlChannel.EntityData)
}

