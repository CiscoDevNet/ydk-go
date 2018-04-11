// This module contains a collection of YANG definitions
// for Cisco IOS-XR ip-sbfd package operational data.
// 
// This module contains definitions
// for the following management objects:
//   sbfd: Seamless BFD (S-BFD) operational data
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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

// SbfdAddressFamily represents Sbfd address family
type SbfdAddressFamily string

const (
    // ipv4
    SbfdAddressFamily_ipv4 SbfdAddressFamily = "ipv4"

    // ipv6
    SbfdAddressFamily_ipv6 SbfdAddressFamily = "ipv6"
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
    sbfd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sbfd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sbfd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sbfd.EntityData.Children = make(map[string]types.YChild)
    sbfd.EntityData.Children["target-identifier"] = types.YChild{"TargetIdentifier", &sbfd.TargetIdentifier}
    sbfd.EntityData.Leafs = make(map[string]types.YLeaf)
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
    targetIdentifier.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    targetIdentifier.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    targetIdentifier.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    targetIdentifier.EntityData.Children = make(map[string]types.YChild)
    targetIdentifier.EntityData.Children["remote-vrfs"] = types.YChild{"RemoteVrfs", &targetIdentifier.RemoteVrfs}
    targetIdentifier.EntityData.Children["local-vrfs"] = types.YChild{"LocalVrfs", &targetIdentifier.LocalVrfs}
    targetIdentifier.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(targetIdentifier.EntityData)
}

// Sbfd_TargetIdentifier_RemoteVrfs
// SBFD remote discriminator data
type Sbfd_TargetIdentifier_RemoteVrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of remote discriminator data per VRF. The type is slice of
    // Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf.
    RemoteVrf []Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf
}

func (remoteVrfs *Sbfd_TargetIdentifier_RemoteVrfs) GetEntityData() *types.CommonEntityData {
    remoteVrfs.EntityData.YFilter = remoteVrfs.YFilter
    remoteVrfs.EntityData.YangName = "remote-vrfs"
    remoteVrfs.EntityData.BundleName = "cisco_ios_xr"
    remoteVrfs.EntityData.ParentYangName = "target-identifier"
    remoteVrfs.EntityData.SegmentPath = "remote-vrfs"
    remoteVrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteVrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteVrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteVrfs.EntityData.Children = make(map[string]types.YChild)
    remoteVrfs.EntityData.Children["remote-vrf"] = types.YChild{"RemoteVrf", nil}
    for i := range remoteVrfs.RemoteVrf {
        remoteVrfs.EntityData.Children[types.GetSegmentPath(&remoteVrfs.RemoteVrf[i])] = types.YChild{"RemoteVrf", &remoteVrfs.RemoteVrf[i]}
    }
    remoteVrfs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(remoteVrfs.EntityData)
}

// Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf
// Table of remote discriminator data per VRF
type Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // SBFD remote discriminator . The type is slice of
    // Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator.
    RemoteDiscriminator []Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator
}

func (remoteVrf *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf) GetEntityData() *types.CommonEntityData {
    remoteVrf.EntityData.YFilter = remoteVrf.YFilter
    remoteVrf.EntityData.YangName = "remote-vrf"
    remoteVrf.EntityData.BundleName = "cisco_ios_xr"
    remoteVrf.EntityData.ParentYangName = "remote-vrfs"
    remoteVrf.EntityData.SegmentPath = "remote-vrf" + "[vrf-name='" + fmt.Sprintf("%v", remoteVrf.VrfName) + "']"
    remoteVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteVrf.EntityData.Children = make(map[string]types.YChild)
    remoteVrf.EntityData.Children["remote-discriminator"] = types.YChild{"RemoteDiscriminator", nil}
    for i := range remoteVrf.RemoteDiscriminator {
        remoteVrf.EntityData.Children[types.GetSegmentPath(&remoteVrf.RemoteDiscriminator[i])] = types.YChild{"RemoteDiscriminator", &remoteVrf.RemoteDiscriminator[i]}
    }
    remoteVrf.EntityData.Leafs = make(map[string]types.YLeaf)
    remoteVrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", remoteVrf.VrfName}
    return &(remoteVrf.EntityData)
}

// Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator
// SBFD remote discriminator 
type Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF Name . The type is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // Remote Discriminator. The type is interface{} with range:
    // -2147483648..2147483647.
    RemoteDiscriminator interface{}

    // Address. The type is one of the following types: string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?',
    // or string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
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
    remoteDiscriminator.EntityData.SegmentPath = "remote-discriminator"
    remoteDiscriminator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    remoteDiscriminator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    remoteDiscriminator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    remoteDiscriminator.EntityData.Children = make(map[string]types.YChild)
    remoteDiscriminator.EntityData.Children["ip-address"] = types.YChild{"IpAddress", &remoteDiscriminator.IpAddress}
    remoteDiscriminator.EntityData.Leafs = make(map[string]types.YLeaf)
    remoteDiscriminator.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", remoteDiscriminator.VrfName}
    remoteDiscriminator.EntityData.Leafs["remote-discriminator"] = types.YLeaf{"RemoteDiscriminator", remoteDiscriminator.RemoteDiscriminator}
    remoteDiscriminator.EntityData.Leafs["address"] = types.YLeaf{"Address", remoteDiscriminator.Address}
    remoteDiscriminator.EntityData.Leafs["tid-type"] = types.YLeaf{"TidType", remoteDiscriminator.TidType}
    remoteDiscriminator.EntityData.Leafs["discr"] = types.YLeaf{"Discr", remoteDiscriminator.Discr}
    remoteDiscriminator.EntityData.Leafs["vrf-name-xr"] = types.YLeaf{"VrfNameXr", remoteDiscriminator.VrfNameXr}
    remoteDiscriminator.EntityData.Leafs["status"] = types.YLeaf{"Status", remoteDiscriminator.Status}
    remoteDiscriminator.EntityData.Leafs["discr-src"] = types.YLeaf{"DiscrSrc", remoteDiscriminator.DiscrSrc}
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
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ipv4 interface{}

    // IPv6 address type. The type is string with pattern:
    // b'((:|[0-9a-fA-F]{0,4}):)([0-9a-fA-F]{0,4}:){0,5}((([0-9a-fA-F]{0,4}:)?(:|[0-9a-fA-F]{0,4}))|(((25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])\\.){3}(25[0-5]|2[0-4][0-9]|[01]?[0-9]?[0-9])))(%[\\p{N}\\p{L}]+)?'.
    Ipv6 interface{}
}

func (ipAddress *Sbfd_TargetIdentifier_RemoteVrfs_RemoteVrf_RemoteDiscriminator_IpAddress) GetEntityData() *types.CommonEntityData {
    ipAddress.EntityData.YFilter = ipAddress.YFilter
    ipAddress.EntityData.YangName = "ip-address"
    ipAddress.EntityData.BundleName = "cisco_ios_xr"
    ipAddress.EntityData.ParentYangName = "remote-discriminator"
    ipAddress.EntityData.SegmentPath = "ip-address"
    ipAddress.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ipAddress.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ipAddress.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ipAddress.EntityData.Children = make(map[string]types.YChild)
    ipAddress.EntityData.Leafs = make(map[string]types.YLeaf)
    ipAddress.EntityData.Leafs["afi"] = types.YLeaf{"Afi", ipAddress.Afi}
    ipAddress.EntityData.Leafs["dummy"] = types.YLeaf{"Dummy", ipAddress.Dummy}
    ipAddress.EntityData.Leafs["ipv4"] = types.YLeaf{"Ipv4", ipAddress.Ipv4}
    ipAddress.EntityData.Leafs["ipv6"] = types.YLeaf{"Ipv6", ipAddress.Ipv6}
    return &(ipAddress.EntityData)
}

// Sbfd_TargetIdentifier_LocalVrfs
// SBFD local discriminator  data
type Sbfd_TargetIdentifier_LocalVrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Table of local discriminator data per VRF. The type is slice of
    // Sbfd_TargetIdentifier_LocalVrfs_LocalVrf.
    LocalVrf []Sbfd_TargetIdentifier_LocalVrfs_LocalVrf
}

func (localVrfs *Sbfd_TargetIdentifier_LocalVrfs) GetEntityData() *types.CommonEntityData {
    localVrfs.EntityData.YFilter = localVrfs.YFilter
    localVrfs.EntityData.YangName = "local-vrfs"
    localVrfs.EntityData.BundleName = "cisco_ios_xr"
    localVrfs.EntityData.ParentYangName = "target-identifier"
    localVrfs.EntityData.SegmentPath = "local-vrfs"
    localVrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localVrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localVrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localVrfs.EntityData.Children = make(map[string]types.YChild)
    localVrfs.EntityData.Children["local-vrf"] = types.YChild{"LocalVrf", nil}
    for i := range localVrfs.LocalVrf {
        localVrfs.EntityData.Children[types.GetSegmentPath(&localVrfs.LocalVrf[i])] = types.YChild{"LocalVrf", &localVrfs.LocalVrf[i]}
    }
    localVrfs.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(localVrfs.EntityData)
}

// Sbfd_TargetIdentifier_LocalVrfs_LocalVrf
// Table of local discriminator data per VRF
type Sbfd_TargetIdentifier_LocalVrfs_LocalVrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. VRF name. The type is string with pattern:
    // b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    VrfName interface{}

    // SBFD local discriminator . The type is slice of
    // Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator.
    LocalDiscriminator []Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator
}

func (localVrf *Sbfd_TargetIdentifier_LocalVrfs_LocalVrf) GetEntityData() *types.CommonEntityData {
    localVrf.EntityData.YFilter = localVrf.YFilter
    localVrf.EntityData.YangName = "local-vrf"
    localVrf.EntityData.BundleName = "cisco_ios_xr"
    localVrf.EntityData.ParentYangName = "local-vrfs"
    localVrf.EntityData.SegmentPath = "local-vrf" + "[vrf-name='" + fmt.Sprintf("%v", localVrf.VrfName) + "']"
    localVrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localVrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localVrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localVrf.EntityData.Children = make(map[string]types.YChild)
    localVrf.EntityData.Children["local-discriminator"] = types.YChild{"LocalDiscriminator", nil}
    for i := range localVrf.LocalDiscriminator {
        localVrf.EntityData.Children[types.GetSegmentPath(&localVrf.LocalDiscriminator[i])] = types.YChild{"LocalDiscriminator", &localVrf.LocalDiscriminator[i]}
    }
    localVrf.EntityData.Leafs = make(map[string]types.YLeaf)
    localVrf.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", localVrf.VrfName}
    return &(localVrf.EntityData)
}

// Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator
// SBFD local discriminator 
type Sbfd_TargetIdentifier_LocalVrfs_LocalVrf_LocalDiscriminator struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local discriminator. The type is interface{} with range:
    // -2147483648..2147483647.
    LocalDiscriminator interface{}

    // VRF Name . The type is string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
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
    localDiscriminator.EntityData.SegmentPath = "local-discriminator"
    localDiscriminator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localDiscriminator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localDiscriminator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localDiscriminator.EntityData.Children = make(map[string]types.YChild)
    localDiscriminator.EntityData.Leafs = make(map[string]types.YLeaf)
    localDiscriminator.EntityData.Leafs["local-discriminator"] = types.YLeaf{"LocalDiscriminator", localDiscriminator.LocalDiscriminator}
    localDiscriminator.EntityData.Leafs["vrf-name"] = types.YLeaf{"VrfName", localDiscriminator.VrfName}
    localDiscriminator.EntityData.Leafs["discr"] = types.YLeaf{"Discr", localDiscriminator.Discr}
    localDiscriminator.EntityData.Leafs["vrf-name-xr"] = types.YLeaf{"VrfNameXr", localDiscriminator.VrfNameXr}
    localDiscriminator.EntityData.Leafs["flags"] = types.YLeaf{"Flags", localDiscriminator.Flags}
    localDiscriminator.EntityData.Leafs["status"] = types.YLeaf{"Status", localDiscriminator.Status}
    localDiscriminator.EntityData.Leafs["discr-src"] = types.YLeaf{"DiscrSrc", localDiscriminator.DiscrSrc}
    return &(localDiscriminator.EntityData)
}

