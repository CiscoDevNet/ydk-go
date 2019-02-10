// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-filesystems package configuration.
// 
// This module contains definitions
// for the following management objects:
//   rcp: RCP configuration
//   ftp: ftp
//   tftp: tftp
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ipv4_filesystems_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ipv4_filesystems_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-filesystems-cfg rcp}", reflect.TypeOf(Rcp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-filesystems-cfg:rcp", reflect.TypeOf(Rcp{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-filesystems-cfg ftp}", reflect.TypeOf(Ftp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-filesystems-cfg:ftp", reflect.TypeOf(Ftp{}))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ipv4-filesystems-cfg tftp}", reflect.TypeOf(Tftp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ipv4-filesystems-cfg:tftp", reflect.TypeOf(Tftp{}))
}

// Rcp
// RCP configuration
type Rcp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // RCP client configuration.
    RcpClient Rcp_RcpClient
}

func (rcp *Rcp) GetEntityData() *types.CommonEntityData {
    rcp.EntityData.YFilter = rcp.YFilter
    rcp.EntityData.YangName = "rcp"
    rcp.EntityData.BundleName = "cisco_ios_xr"
    rcp.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-filesystems-cfg"
    rcp.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-filesystems-cfg:rcp"
    rcp.EntityData.AbsolutePath = rcp.EntityData.SegmentPath
    rcp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rcp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rcp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rcp.EntityData.Children = types.NewOrderedMap()
    rcp.EntityData.Children.Append("rcp-client", types.YChild{"RcpClient", &rcp.RcpClient})
    rcp.EntityData.Leafs = types.NewOrderedMap()

    rcp.EntityData.YListKeys = []string {}

    return &(rcp.EntityData)
}

// Rcp_RcpClient
// RCP client configuration
type Rcp_RcpClient struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify username for connections. The type is string.
    Username interface{}

    // Specify interface for source address in connections. The type is string
    // with pattern: [a-zA-Z0-9._/-]+.
    SourceInterface interface{}
}

func (rcpClient *Rcp_RcpClient) GetEntityData() *types.CommonEntityData {
    rcpClient.EntityData.YFilter = rcpClient.YFilter
    rcpClient.EntityData.YangName = "rcp-client"
    rcpClient.EntityData.BundleName = "cisco_ios_xr"
    rcpClient.EntityData.ParentYangName = "rcp"
    rcpClient.EntityData.SegmentPath = "rcp-client"
    rcpClient.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-filesystems-cfg:rcp/" + rcpClient.EntityData.SegmentPath
    rcpClient.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    rcpClient.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    rcpClient.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    rcpClient.EntityData.Children = types.NewOrderedMap()
    rcpClient.EntityData.Leafs = types.NewOrderedMap()
    rcpClient.EntityData.Leafs.Append("username", types.YLeaf{"Username", rcpClient.Username})
    rcpClient.EntityData.Leafs.Append("source-interface", types.YLeaf{"SourceInterface", rcpClient.SourceInterface})

    rcpClient.EntityData.YListKeys = []string {}

    return &(rcpClient.EntityData)
}

// Ftp
// ftp
type Ftp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // FTP client configuration.
    FtpClient Ftp_FtpClient
}

func (ftp *Ftp) GetEntityData() *types.CommonEntityData {
    ftp.EntityData.YFilter = ftp.YFilter
    ftp.EntityData.YangName = "ftp"
    ftp.EntityData.BundleName = "cisco_ios_xr"
    ftp.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-filesystems-cfg"
    ftp.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-filesystems-cfg:ftp"
    ftp.EntityData.AbsolutePath = ftp.EntityData.SegmentPath
    ftp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ftp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ftp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ftp.EntityData.Children = types.NewOrderedMap()
    ftp.EntityData.Children.Append("ftp-client", types.YChild{"FtpClient", &ftp.FtpClient})
    ftp.EntityData.Leafs = types.NewOrderedMap()

    ftp.EntityData.YListKeys = []string {}

    return &(ftp.EntityData)
}

// Ftp_FtpClient
// FTP client configuration
type Ftp_FtpClient struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable connect using passive mode. The type is interface{}.
    Passive interface{}

    // Specify password for ftp connnection. The type is string with pattern:
    // (!.+)|([^!].+).
    Password interface{}

    // Password for anonymous user (ftp server dependent). The type is string.
    AnonymousPassword interface{}

    // Specify username for connections. The type is string.
    Username interface{}

    // Specify interface for source address in connections. The type is string
    // with pattern: [a-zA-Z0-9._/-]+.
    SourceInterface interface{}

    // VRF table.
    Vrfs Ftp_FtpClient_Vrfs
}

func (ftpClient *Ftp_FtpClient) GetEntityData() *types.CommonEntityData {
    ftpClient.EntityData.YFilter = ftpClient.YFilter
    ftpClient.EntityData.YangName = "ftp-client"
    ftpClient.EntityData.BundleName = "cisco_ios_xr"
    ftpClient.EntityData.ParentYangName = "ftp"
    ftpClient.EntityData.SegmentPath = "ftp-client"
    ftpClient.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-filesystems-cfg:ftp/" + ftpClient.EntityData.SegmentPath
    ftpClient.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ftpClient.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ftpClient.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ftpClient.EntityData.Children = types.NewOrderedMap()
    ftpClient.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &ftpClient.Vrfs})
    ftpClient.EntityData.Leafs = types.NewOrderedMap()
    ftpClient.EntityData.Leafs.Append("passive", types.YLeaf{"Passive", ftpClient.Passive})
    ftpClient.EntityData.Leafs.Append("password", types.YLeaf{"Password", ftpClient.Password})
    ftpClient.EntityData.Leafs.Append("anonymous-password", types.YLeaf{"AnonymousPassword", ftpClient.AnonymousPassword})
    ftpClient.EntityData.Leafs.Append("username", types.YLeaf{"Username", ftpClient.Username})
    ftpClient.EntityData.Leafs.Append("source-interface", types.YLeaf{"SourceInterface", ftpClient.SourceInterface})

    ftpClient.EntityData.YListKeys = []string {}

    return &(ftpClient.EntityData)
}

// Ftp_FtpClient_Vrfs
// VRF table
type Ftp_FtpClient_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF specific data. The type is slice of Ftp_FtpClient_Vrfs_Vrf.
    Vrf []*Ftp_FtpClient_Vrfs_Vrf
}

func (vrfs *Ftp_FtpClient_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "ftp-client"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-filesystems-cfg:ftp/ftp-client/" + vrfs.EntityData.SegmentPath
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = types.NewOrderedMap()
    vrfs.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children.Append(types.GetSegmentPath(vrfs.Vrf[i]), types.YChild{"Vrf", vrfs.Vrf[i]})
    }
    vrfs.EntityData.Leafs = types.NewOrderedMap()

    vrfs.EntityData.YListKeys = []string {}

    return &(vrfs.EntityData)
}

// Ftp_FtpClient_Vrfs_Vrf
// VRF specific data
type Ftp_FtpClient_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the VRF instance. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Specify interface for source address in connections. The type is string
    // with pattern: [a-zA-Z0-9._/-]+.
    SourceInterface interface{}

    // Specify username for connections. The type is string.
    Username interface{}

    // Password for anonymous user (ftp server dependent). The type is string.
    AnonymousPassword interface{}

    // Specify password for ftp connnection. The type is string with pattern:
    // (!.+)|([^!].+).
    Password interface{}

    // Enable connect using passive mode. The type is interface{}.
    Passive interface{}
}

func (vrf *Ftp_FtpClient_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-filesystems-cfg:ftp/ftp-client/vrfs/" + vrf.EntityData.SegmentPath
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})
    vrf.EntityData.Leafs.Append("source-interface", types.YLeaf{"SourceInterface", vrf.SourceInterface})
    vrf.EntityData.Leafs.Append("username", types.YLeaf{"Username", vrf.Username})
    vrf.EntityData.Leafs.Append("anonymous-password", types.YLeaf{"AnonymousPassword", vrf.AnonymousPassword})
    vrf.EntityData.Leafs.Append("password", types.YLeaf{"Password", vrf.Password})
    vrf.EntityData.Leafs.Append("passive", types.YLeaf{"Passive", vrf.Passive})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

// Tftp
// tftp
type Tftp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // TFTP client configuration.
    TftpClient Tftp_TftpClient
}

func (tftp *Tftp) GetEntityData() *types.CommonEntityData {
    tftp.EntityData.YFilter = tftp.YFilter
    tftp.EntityData.YangName = "tftp"
    tftp.EntityData.BundleName = "cisco_ios_xr"
    tftp.EntityData.ParentYangName = "Cisco-IOS-XR-ipv4-filesystems-cfg"
    tftp.EntityData.SegmentPath = "Cisco-IOS-XR-ipv4-filesystems-cfg:tftp"
    tftp.EntityData.AbsolutePath = tftp.EntityData.SegmentPath
    tftp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tftp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tftp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tftp.EntityData.Children = types.NewOrderedMap()
    tftp.EntityData.Children.Append("tftp-client", types.YChild{"TftpClient", &tftp.TftpClient})
    tftp.EntityData.Leafs = types.NewOrderedMap()

    tftp.EntityData.YListKeys = []string {}

    return &(tftp.EntityData)
}

// Tftp_TftpClient
// TFTP client configuration
type Tftp_TftpClient struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the number of retries when client requests TFTP connections. The
    // type is interface{} with range: 0..256.
    Retry interface{}

    // Specify the timeout for every TFTP connection in seconds. The type is
    // interface{} with range: 0..256. Units are second.
    Timeout interface{}

    // Specify interface for source address in connections. The type is string
    // with pattern: [a-zA-Z0-9._/-]+.
    SourceInterface interface{}

    // VRF table.
    Vrfs Tftp_TftpClient_Vrfs
}

func (tftpClient *Tftp_TftpClient) GetEntityData() *types.CommonEntityData {
    tftpClient.EntityData.YFilter = tftpClient.YFilter
    tftpClient.EntityData.YangName = "tftp-client"
    tftpClient.EntityData.BundleName = "cisco_ios_xr"
    tftpClient.EntityData.ParentYangName = "tftp"
    tftpClient.EntityData.SegmentPath = "tftp-client"
    tftpClient.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-filesystems-cfg:tftp/" + tftpClient.EntityData.SegmentPath
    tftpClient.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tftpClient.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tftpClient.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tftpClient.EntityData.Children = types.NewOrderedMap()
    tftpClient.EntityData.Children.Append("vrfs", types.YChild{"Vrfs", &tftpClient.Vrfs})
    tftpClient.EntityData.Leafs = types.NewOrderedMap()
    tftpClient.EntityData.Leafs.Append("retry", types.YLeaf{"Retry", tftpClient.Retry})
    tftpClient.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", tftpClient.Timeout})
    tftpClient.EntityData.Leafs.Append("source-interface", types.YLeaf{"SourceInterface", tftpClient.SourceInterface})

    tftpClient.EntityData.YListKeys = []string {}

    return &(tftpClient.EntityData)
}

// Tftp_TftpClient_Vrfs
// VRF table
type Tftp_TftpClient_Vrfs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // VRF specific data. The type is slice of Tftp_TftpClient_Vrfs_Vrf.
    Vrf []*Tftp_TftpClient_Vrfs_Vrf
}

func (vrfs *Tftp_TftpClient_Vrfs) GetEntityData() *types.CommonEntityData {
    vrfs.EntityData.YFilter = vrfs.YFilter
    vrfs.EntityData.YangName = "vrfs"
    vrfs.EntityData.BundleName = "cisco_ios_xr"
    vrfs.EntityData.ParentYangName = "tftp-client"
    vrfs.EntityData.SegmentPath = "vrfs"
    vrfs.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-filesystems-cfg:tftp/tftp-client/" + vrfs.EntityData.SegmentPath
    vrfs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrfs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrfs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrfs.EntityData.Children = types.NewOrderedMap()
    vrfs.EntityData.Children.Append("vrf", types.YChild{"Vrf", nil})
    for i := range vrfs.Vrf {
        vrfs.EntityData.Children.Append(types.GetSegmentPath(vrfs.Vrf[i]), types.YChild{"Vrf", vrfs.Vrf[i]})
    }
    vrfs.EntityData.Leafs = types.NewOrderedMap()

    vrfs.EntityData.YListKeys = []string {}

    return &(vrfs.EntityData)
}

// Tftp_TftpClient_Vrfs_Vrf
// VRF specific data
type Tftp_TftpClient_Vrfs_Vrf struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the VRF instance. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Specify interface for source address in connections. The type is string
    // with pattern: [a-zA-Z0-9._/-]+.
    SourceInterface interface{}

    // Specify the number of retries when client requests TFTP connections. The
    // type is interface{} with range: 0..256.
    Retry interface{}

    // Specify the timeout for every TFTP connection in seconds. The type is
    // interface{} with range: 0..256. Units are second.
    Timeout interface{}
}

func (vrf *Tftp_TftpClient_Vrfs_Vrf) GetEntityData() *types.CommonEntityData {
    vrf.EntityData.YFilter = vrf.YFilter
    vrf.EntityData.YangName = "vrf"
    vrf.EntityData.BundleName = "cisco_ios_xr"
    vrf.EntityData.ParentYangName = "vrfs"
    vrf.EntityData.SegmentPath = "vrf" + types.AddKeyToken(vrf.VrfName, "vrf-name")
    vrf.EntityData.AbsolutePath = "Cisco-IOS-XR-ipv4-filesystems-cfg:tftp/tftp-client/vrfs/" + vrf.EntityData.SegmentPath
    vrf.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    vrf.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    vrf.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    vrf.EntityData.Children = types.NewOrderedMap()
    vrf.EntityData.Leafs = types.NewOrderedMap()
    vrf.EntityData.Leafs.Append("vrf-name", types.YLeaf{"VrfName", vrf.VrfName})
    vrf.EntityData.Leafs.Append("source-interface", types.YLeaf{"SourceInterface", vrf.SourceInterface})
    vrf.EntityData.Leafs.Append("retry", types.YLeaf{"Retry", vrf.Retry})
    vrf.EntityData.Leafs.Append("timeout", types.YLeaf{"Timeout", vrf.Timeout})

    vrf.EntityData.YListKeys = []string {"VrfName"}

    return &(vrf.EntityData)
}

