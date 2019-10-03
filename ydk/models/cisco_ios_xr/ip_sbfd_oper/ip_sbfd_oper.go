// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-sbfd package operational data.
// 
// This module contains definitions
// for the following management objects:
//   sbfd: Seamless BFD (S-BFD) operational data
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ip_sbfd_oper

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ip_sbfd_oper"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ip-sbfd-oper sbfd}", reflect.TypeOf(Sbfd{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ip-sbfd-oper:sbfd", reflect.TypeOf(Sbfd{}))
}

// SbfdAddressFamily represents Sbfd address family
type SbfdAddressFamily string

const (
    // ipv4
    SbfdAddressFamily_ipv4 SbfdAddressFamily = "ipv4"

    // ipv6
    SbfdAddressFamily_ipv6 SbfdAddressFamily = "ipv6"
)

// BfdAfId represents Bfd af id
type BfdAfId string

const (
    // No Address
    BfdAfId_bfd_af_id_none BfdAfId = "bfd-af-id-none"

    // IPv4 AFI
    BfdAfId_bfd_af_id_ipv4 BfdAfId = "bfd-af-id-ipv4"

    // IPv6 AFI
    BfdAfId_bfd_af_id_ipv6 BfdAfId = "bfd-af-id-ipv6"
)

// Sbfd
// Seamless BFD (S-BFD) operational data
type Sbfd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Target-identifier information.
    TargetIdentifier Sbfd_TargetIdentifier
}

func (sbfd *Sbfd) GetEntityData() *types.CommonEntityData {
    sbfd.EntityData.YFilter = sbfd.YFilter
    sbfd.EntityData.YangName = "sbfd"
    sbfd.EntityData.BundleName = "cisco_ios_xr"
    sbfd.EntityData.ParentYangName = "Cisco-IOS-XR-ip-sbfd-oper"
    sbfd.EntityData.SegmentPath = "Cisco-IOS-XR-ip-sbfd-oper:sbfd"
    sbfd.EntityData.AbsolutePath = sbfd.EntityData.SegmentPath
    sbfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sbfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sbfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sbfd.EntityData.Children = types.NewOrderedMap()
    sbfd.EntityData.Children.Append("target-identifier", types.YChild{"TargetIdentifier", &sbfd.TargetIdentifier})
    sbfd.EntityData.Leafs = types.NewOrderedMap()

    sbfd.EntityData.YListKeys = []string {}

    return &(sbfd.EntityData)
}

// Sbfd_TargetIdentifier
// Target-identifier information
type Sbfd_TargetIdentifier struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // SBFD remote discriminator data.
    RemoteVrfs Sbfd_TargetIdentifier_RemoteVrfs

    // SBFD local discriminator  data.
    LocalVrfs Sbfd_TargetIdentifier_LocalVrfs
}

func (targetIdentifier *Sbfd_TargetIdentifier) GetEntityData() *types.CommonEntityData {
    targetIdentifier.EntityData.YFilter = targetIdentifier.YFilter
    targetIdentifier.EntityData.YangName = "target-identifier"
    targetIdentifier.EntityData.BundleName = "cisco_ios_xr"
    targetIdentifier.EntityData.ParentYangName = "sbfd"
    targetIdentifier.EntityData.SegmentPath = "target-identifier"
    targetIdentifier.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-sbfd-oper:sbfd/" + targetIdentifier.EntityData.SegmentPath
    targetIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    targetIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    targetIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    targetIdentifier.EntityData.Children = types.NewOrderedMap()
    targetIdentifier.EntityData.Children.Append("remote-vrfs", types.YChild{"RemoteVrfs", &targetIdentifier.RemoteVrfs})
    targetIdentifier.EntityData.Children.Append("local-vrfs", types.YChild{"LocalVrfs", &targetIdentifier.LocalVrfs})
    targetIdentifier.EntityData.Leafs = types.NewOrderedMap()

    targetIdentifier.EntityData.YListKeys = []string {}

    return &(targetIdentifier.EntityData)
}

// Sbfd_TargetIdentifier_RemoteVrfs
// SBFD remote discriminator data
type Sbfd_TargetIdentifier_RemoteVrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of remote discriminator data per VRF. The type is slice of
    // Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf.
    RemoteVrf []*Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf
}

func (remoteVrfs *Sbfd_TargetIdentifier_RemoteVrfs) GetEntityData() *types.CommonEntityData {
    remoteVrfs.EntityData.YFilter = remoteVrfs.YFilter
    remoteVrfs.EntityData.YangName = "remote-vrfs"
    remoteVrfs.EntityData.BundleName = "cisco_ios_xr"
    remoteVrfs.EntityData.ParentYangName = "target-identifier"
    remoteVrfs.EntityData.SegmentPath = "remote-vrfs"
    remoteVrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-sbfd-oper:sbfd/target-identifier/" + remoteVrfs.EntityData.SegmentPath
    remoteVrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteVrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteVrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteVrfs.EntityData.Children = types.NewOrderedMap()
    remoteVrfs.EntityData.Children.Append("remote-vrf", types.YChild{"RemoteVrf", nil})
    for i := range remoteVrfs.RemoteVrf {
        remoteVrfs.EntityData.Children.Append(types.GetSegmentPath(remoteVrfs.RemoteVrf[i]), types.YChild{"RemoteVrf", remoteVrfs.RemoteVrf[i]})
    }
    remoteVrfs.EntityData.Leafs = types.NewOrderedMap()

    remoteVrfs.EntityData.YListKeys = []string {}

    return &(remoteVrfs.EntityData)
}

// Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf
// Table of remote discriminator data per VRF
type Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // SBFD remote discriminator . The type is slice of
    // Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator.
    RemoteDiscriminator []*Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator
}

func (remoteVrf *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf) GetEntityData() *types.CommonEntityData {
    remoteVrf.EntityData.YFilter = remoteVrf.YFilter
    remoteVrf.EntityData.YangName = "remote-vrf"
    remoteVrf.EntityData.BundleName = "cisco_ios_xr"
    remoteVrf.EntityData.ParentYangName = "remote-vrfs"
    remoteVrf.EntityData.SegmentPath = "remote-vrf" + types.AddKeyToken(remoteVrf.VrfName, "vrf-name")
    remoteVrf.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-sbfd-oper:sbfd/target-identifier/remote-vrfs/" + remoteVrf.EntityData.SegmentPath
    remoteVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteVrf.EntityData.Children = types.NewOrderedMap()
    remoteVrf.EntityData.Children.Append("remote-discriminator", types.YChild{"RemoteDiscriminator", nil})
    for i := range remoteVrf.RemoteDiscriminator {
        types.SetYListKey(remoteVrf.RemoteDiscriminator[i], i)
        remoteVrf.EntityData.Children.Append(types.GetSegmentPath(remoteVrf.RemoteDiscriminator[i]), types.YChild{"RemoteDiscriminator", remoteVrf.RemoteDiscriminator[i]})
    }
    remoteVrf.EntityData.Leafs = types.NewOrderedMap()
    remoteVrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", remoteVrf.VrfName})

    remoteVrf.EntityData.YListKeys = []string {"VrfName"}

    return &(remoteVrf.EntityData)
}

// Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator
// SBFD remote discriminator 
type Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // VRF Name . The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Remote Discriminator. The type is interface{} with range: 0..4294967295.
    RemoteDiscriminator interface{}

    // Address. The type is one of the following types: string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?,
    // or string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Address interface{}

    // Target identifier for sbfd. The type is SbfdAddressFamily.
    TidType interface{}

    // Remote discriminator. The type is interface{} with range: 0..4294967295.
    Discr interface{}

    // VRF Name. The type is string.
    VrfNameXr interface{}

    // Status. The type is string.
    Status interface{}

    // Discriminator source name. The type is string.
    DiscrSrc interface{}

    // IP address.
    IpAddress Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator_IpAddress
}

func (remoteDiscriminator *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator) GetEntityData() *types.CommonEntityData {
    remoteDiscriminator.EntityData.YFilter = remoteDiscriminator.YFilter
    remoteDiscriminator.EntityData.YangName = "remote-discriminator"
    remoteDiscriminator.EntityData.BundleName = "cisco_ios_xr"
    remoteDiscriminator.EntityData.ParentYangName = "remote-vrf"
    remoteDiscriminator.EntityData.SegmentPath = "remote-discriminator" + types.AddNoKeyToken(remoteDiscriminator)
    remoteDiscriminator.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-sbfd-oper:sbfd/target-identifier/remote-vrfs/remote-vrf/" + remoteDiscriminator.EntityData.SegmentPath
    remoteDiscriminator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteDiscriminator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteDiscriminator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteDiscriminator.EntityData.Children = types.NewOrderedMap()
    remoteDiscriminator.EntityData.Children.Append("ip-address", types.YChild{"IpAddress", &remoteDiscriminator.IpAddress})
    remoteDiscriminator.EntityData.Leafs = types.NewOrderedMap()
    remoteDiscriminator.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", remoteDiscriminator.VrfName})
    remoteDiscriminator.EntityData.Leafs.Append("remote-discriminator", types.YLeaf{"RemoteDiscriminator", remoteDiscriminator.RemoteDiscriminator})
    remoteDiscriminator.EntityData.Leafs.Append("address", types.YLeaf{"Address", remoteDiscriminator.Address})
    remoteDiscriminator.EntityData.Leafs.Append("tid-type", types.YLeaf{"TidType", remoteDiscriminator.TidType})
    remoteDiscriminator.EntityData.Leafs.Append("discr", types.YLeaf{"Discr", remoteDiscriminator.Discr})
    remoteDiscriminator.EntityData.Leafs.Append("vrf-name-xr", types.YLeaf{"VrfNameXr", remoteDiscriminator.VrfNameXr})
    remoteDiscriminator.EntityData.Leafs.Append("status", types.YLeaf{"Status", remoteDiscriminator.Status})
    remoteDiscriminator.EntityData.Leafs.Append("discr-src", types.YLeaf{"DiscrSrc", remoteDiscriminator.DiscrSrc})

    remoteDiscriminator.EntityData.YListKeys = []string {}

    return &(remoteDiscriminator.EntityData)
}

// Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator_IpAddress
// IP address
type Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator_IpAddress struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // AFI. The type is BfdAfId.
    Afi interface{}

    // No Address. The type is interface{} with range: 0..255.
    Dummy interface{}

    // IPv4 address type. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // ((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\p{N}\p{L}]+)?.
    Ipv6 interface{}
}

func (ipAddress *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator_IpAddress) GetEntityData() *types.CommonEntityData {
    ipAddress.EntityData.YFilter = ipAddress.YFilter
    ipAddress.EntityData.YangName = "ip-address"
    ipAddress.EntityData.BundleName = "cisco_ios_xr"
    ipAddress.EntityData.ParentYangName = "remote-discriminator"
    ipAddress.EntityData.SegmentPath = "ip-address"
    ipAddress.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-sbfd-oper:sbfd/target-identifier/remote-vrfs/remote-vrf/remote-discriminator/" + ipAddress.EntityData.SegmentPath
    ipAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipAddress.EntityData.Children = types.NewOrderedMap()
    ipAddress.EntityData.Leafs = types.NewOrderedMap()
    ipAddress.EntityData.Leafs.Append("afi", types.YLeaf{"Afi", ipAddress.Afi})
    ipAddress.EntityData.Leafs.Append("dummy", types.YLeaf{"Dummy", ipAddress.Dummy})
    ipAddress.EntityData.Leafs.Append("ipv4", types.YLeaf{"Ipv4", ipAddress.Ipv4})
    ipAddress.EntityData.Leafs.Append("ipv6", types.YLeaf{"Ipv6", ipAddress.Ipv6})

    ipAddress.EntityData.YListKeys = []string {}

    return &(ipAddress.EntityData)
}

// Sbfd_TargetIdentifier_LocalVrfs
// SBFD local discriminator  data
type Sbfd_TargetIdentifier_LocalVrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of local discriminator data per VRF. The type is slice of
    // Sbfd_TargetIdentifier_LocalVrfs_LocalVrf.
    LocalVrf []*Sbfd_TargetIdentifier_LocalVrfs_LocalVrf
}

func (localVrfs *Sbfd_TargetIdentifier_LocalVrfs) GetEntityData() *types.CommonEntityData {
    localVrfs.EntityData.YFilter = localVrfs.YFilter
    localVrfs.EntityData.YangName = "local-vrfs"
    localVrfs.EntityData.BundleName = "cisco_ios_xr"
    localVrfs.EntityData.ParentYangName = "target-identifier"
    localVrfs.EntityData.SegmentPath = "local-vrfs"
    localVrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-sbfd-oper:sbfd/target-identifier/" + localVrfs.EntityData.SegmentPath
    localVrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localVrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localVrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localVrfs.EntityData.Children = types.NewOrderedMap()
    localVrfs.EntityData.Children.Append("local-vrf", types.YChild{"LocalVrf", nil})
    for i := range localVrfs.LocalVrf {
        localVrfs.EntityData.Children.Append(types.GetSegmentPath(localVrfs.LocalVrf[i]), types.YChild{"LocalVrf", localVrfs.LocalVrf[i]})
    }
    localVrfs.EntityData.Leafs = types.NewOrderedMap()

    localVrfs.EntityData.YListKeys = []string {}

    return &(localVrfs.EntityData)
}

// Sbfd_TargetIdentifier_LocalVrfs_LocalVrf
// Table of local discriminator data per VRF
type Sbfd_TargetIdentifier_LocalVrfs_LocalVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. VRF name. The type is string with pattern:
    // [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // SBFD local discriminator . The type is slice of
    // Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator.
    LocalDiscriminator []*Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator
}

func (localVrf *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf) GetEntityData() *types.CommonEntityData {
    localVrf.EntityData.YFilter = localVrf.YFilter
    localVrf.EntityData.YangName = "local-vrf"
    localVrf.EntityData.BundleName = "cisco_ios_xr"
    localVrf.EntityData.ParentYangName = "local-vrfs"
    localVrf.EntityData.SegmentPath = "local-vrf" + types.AddKeyToken(localVrf.VrfName, "vrf-name")
    localVrf.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-sbfd-oper:sbfd/target-identifier/local-vrfs/" + localVrf.EntityData.SegmentPath
    localVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localVrf.EntityData.Children = types.NewOrderedMap()
    localVrf.EntityData.Children.Append("local-discriminator", types.YChild{"LocalDiscriminator", nil})
    for i := range localVrf.LocalDiscriminator {
        types.SetYListKey(localVrf.LocalDiscriminator[i], i)
        localVrf.EntityData.Children.Append(types.GetSegmentPath(localVrf.LocalDiscriminator[i]), types.YChild{"LocalDiscriminator", localVrf.LocalDiscriminator[i]})
    }
    localVrf.EntityData.Leafs = types.NewOrderedMap()
    localVrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", localVrf.VrfName})

    localVrf.EntityData.YListKeys = []string {"VrfName"}

    return &(localVrf.EntityData)
}

// Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator
// SBFD local discriminator 
type Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // Local discriminator. The type is interface{} with range: 0..4294967295.
    LocalDiscriminator interface{}

    // VRF Name . The type is string with pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Local discriminator. The type is interface{} with range: 0..4294967295.
    Discr interface{}

    // VRF Name. The type is string.
    VrfNameXr interface{}

    // MODE name. The type is string.
    Flags interface{}

    // Status. The type is string.
    Status interface{}

    // Discriminator source name. The type is string.
    DiscrSrc interface{}
}

func (localDiscriminator *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator) GetEntityData() *types.CommonEntityData {
    localDiscriminator.EntityData.YFilter = localDiscriminator.YFilter
    localDiscriminator.EntityData.YangName = "local-discriminator"
    localDiscriminator.EntityData.BundleName = "cisco_ios_xr"
    localDiscriminator.EntityData.ParentYangName = "local-vrf"
    localDiscriminator.EntityData.SegmentPath = "local-discriminator" + types.AddNoKeyToken(localDiscriminator)
    localDiscriminator.EntityData.AbsolutePath = "Cisco-IOS-XR-ip-sbfd-oper:sbfd/target-identifier/local-vrfs/local-vrf/" + localDiscriminator.EntityData.SegmentPath
    localDiscriminator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localDiscriminator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localDiscriminator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localDiscriminator.EntityData.Children = types.NewOrderedMap()
    localDiscriminator.EntityData.Leafs = types.NewOrderedMap()
    localDiscriminator.EntityData.Leafs.Append("local-discriminator", types.YLeaf{"LocalDiscriminator", localDiscriminator.LocalDiscriminator})
    localDiscriminator.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", localDiscriminator.VrfName})
    localDiscriminator.EntityData.Leafs.Append("discr", types.YLeaf{"Discr", localDiscriminator.Discr})
    localDiscriminator.EntityData.Leafs.Append("vrf-name-xr", types.YLeaf{"VrfNameXr", localDiscriminator.VrfNameXr})
    localDiscriminator.EntityData.Leafs.Append("flags", types.YLeaf{"Flags", localDiscriminator.Flags})
    localDiscriminator.EntityData.Leafs.Append("status", types.YLeaf{"Status", localDiscriminator.Status})
    localDiscriminator.EntityData.Leafs.Append("discr-src", types.YLeaf{"DiscrSrc", localDiscriminator.DiscrSrc})

    localDiscriminator.EntityData.YListKeys = []string {}

    return &(localDiscriminator.EntityData)
}

