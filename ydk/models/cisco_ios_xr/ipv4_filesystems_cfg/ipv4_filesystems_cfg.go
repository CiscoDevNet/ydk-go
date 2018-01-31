// This module contains a collection of YANG definitions
// for Cisco IOS-XR ipv4-filesystems package configuration.
// 
// This module contains definitions
// for the following management objects:
//   rcp: RCP configuration
//   ftp: ftp
//   tftp: tftp
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
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
    parent types.Entity
    YFilter yfilter.YFilter

    // RCP client configuration.
    RcpClient Rcp_RcpClient
}

func (rcp *Rcp) GetFilter() yfilter.YFilter { return rcp.YFilter }

func (rcp *Rcp) SetFilter(yf yfilter.YFilter) { rcp.YFilter = yf }

func (rcp *Rcp) GetGoName(yname string) string {
    if yname == "rcp-client" { return "RcpClient" }
    return ""
}

func (rcp *Rcp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-filesystems-cfg:rcp"
}

func (rcp *Rcp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "rcp-client" {
        return &rcp.RcpClient
    }
    return nil
}

func (rcp *Rcp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["rcp-client"] = &rcp.RcpClient
    return children
}

func (rcp *Rcp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (rcp *Rcp) GetBundleName() string { return "cisco_ios_xr" }

func (rcp *Rcp) GetYangName() string { return "rcp" }

func (rcp *Rcp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rcp *Rcp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rcp *Rcp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rcp *Rcp) SetParent(parent types.Entity) { rcp.parent = parent }

func (rcp *Rcp) GetParent() types.Entity { return rcp.parent }

func (rcp *Rcp) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-filesystems-cfg" }

// Rcp_RcpClient
// RCP client configuration
type Rcp_RcpClient struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify username for connections. The type is string.
    Username interface{}

    // Specify interface for source address in connections. The type is string
    // with pattern: [a-zA-Z0-9./-]+.
    SourceInterface interface{}
}

func (rcpClient *Rcp_RcpClient) GetFilter() yfilter.YFilter { return rcpClient.YFilter }

func (rcpClient *Rcp_RcpClient) SetFilter(yf yfilter.YFilter) { rcpClient.YFilter = yf }

func (rcpClient *Rcp_RcpClient) GetGoName(yname string) string {
    if yname == "username" { return "Username" }
    if yname == "source-interface" { return "SourceInterface" }
    return ""
}

func (rcpClient *Rcp_RcpClient) GetSegmentPath() string {
    return "rcp-client"
}

func (rcpClient *Rcp_RcpClient) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (rcpClient *Rcp_RcpClient) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (rcpClient *Rcp_RcpClient) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["username"] = rcpClient.Username
    leafs["source-interface"] = rcpClient.SourceInterface
    return leafs
}

func (rcpClient *Rcp_RcpClient) GetBundleName() string { return "cisco_ios_xr" }

func (rcpClient *Rcp_RcpClient) GetYangName() string { return "rcp-client" }

func (rcpClient *Rcp_RcpClient) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (rcpClient *Rcp_RcpClient) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (rcpClient *Rcp_RcpClient) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (rcpClient *Rcp_RcpClient) SetParent(parent types.Entity) { rcpClient.parent = parent }

func (rcpClient *Rcp_RcpClient) GetParent() types.Entity { return rcpClient.parent }

func (rcpClient *Rcp_RcpClient) GetParentYangName() string { return "rcp" }

// Ftp
// ftp
type Ftp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // FTP client configuration.
    FtpClient Ftp_FtpClient
}

func (ftp *Ftp) GetFilter() yfilter.YFilter { return ftp.YFilter }

func (ftp *Ftp) SetFilter(yf yfilter.YFilter) { ftp.YFilter = yf }

func (ftp *Ftp) GetGoName(yname string) string {
    if yname == "ftp-client" { return "FtpClient" }
    return ""
}

func (ftp *Ftp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-filesystems-cfg:ftp"
}

func (ftp *Ftp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "ftp-client" {
        return &ftp.FtpClient
    }
    return nil
}

func (ftp *Ftp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["ftp-client"] = &ftp.FtpClient
    return children
}

func (ftp *Ftp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (ftp *Ftp) GetBundleName() string { return "cisco_ios_xr" }

func (ftp *Ftp) GetYangName() string { return "ftp" }

func (ftp *Ftp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ftp *Ftp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ftp *Ftp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ftp *Ftp) SetParent(parent types.Entity) { ftp.parent = parent }

func (ftp *Ftp) GetParent() types.Entity { return ftp.parent }

func (ftp *Ftp) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-filesystems-cfg" }

// Ftp_FtpClient
// FTP client configuration
type Ftp_FtpClient struct {
    parent types.Entity
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
    // with pattern: [a-zA-Z0-9./-]+.
    SourceInterface interface{}
}

func (ftpClient *Ftp_FtpClient) GetFilter() yfilter.YFilter { return ftpClient.YFilter }

func (ftpClient *Ftp_FtpClient) SetFilter(yf yfilter.YFilter) { ftpClient.YFilter = yf }

func (ftpClient *Ftp_FtpClient) GetGoName(yname string) string {
    if yname == "passive" { return "Passive" }
    if yname == "password" { return "Password" }
    if yname == "anonymous-password" { return "AnonymousPassword" }
    if yname == "username" { return "Username" }
    if yname == "source-interface" { return "SourceInterface" }
    return ""
}

func (ftpClient *Ftp_FtpClient) GetSegmentPath() string {
    return "ftp-client"
}

func (ftpClient *Ftp_FtpClient) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (ftpClient *Ftp_FtpClient) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (ftpClient *Ftp_FtpClient) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["passive"] = ftpClient.Passive
    leafs["password"] = ftpClient.Password
    leafs["anonymous-password"] = ftpClient.AnonymousPassword
    leafs["username"] = ftpClient.Username
    leafs["source-interface"] = ftpClient.SourceInterface
    return leafs
}

func (ftpClient *Ftp_FtpClient) GetBundleName() string { return "cisco_ios_xr" }

func (ftpClient *Ftp_FtpClient) GetYangName() string { return "ftp-client" }

func (ftpClient *Ftp_FtpClient) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (ftpClient *Ftp_FtpClient) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (ftpClient *Ftp_FtpClient) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (ftpClient *Ftp_FtpClient) SetParent(parent types.Entity) { ftpClient.parent = parent }

func (ftpClient *Ftp_FtpClient) GetParent() types.Entity { return ftpClient.parent }

func (ftpClient *Ftp_FtpClient) GetParentYangName() string { return "ftp" }

// Tftp
// tftp
type Tftp struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // TFTP client configuration.
    TftpClient Tftp_TftpClient
}

func (tftp *Tftp) GetFilter() yfilter.YFilter { return tftp.YFilter }

func (tftp *Tftp) SetFilter(yf yfilter.YFilter) { tftp.YFilter = yf }

func (tftp *Tftp) GetGoName(yname string) string {
    if yname == "tftp-client" { return "TftpClient" }
    return ""
}

func (tftp *Tftp) GetSegmentPath() string {
    return "Cisco-IOS-XR-ipv4-filesystems-cfg:tftp"
}

func (tftp *Tftp) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "tftp-client" {
        return &tftp.TftpClient
    }
    return nil
}

func (tftp *Tftp) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["tftp-client"] = &tftp.TftpClient
    return children
}

func (tftp *Tftp) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (tftp *Tftp) GetBundleName() string { return "cisco_ios_xr" }

func (tftp *Tftp) GetYangName() string { return "tftp" }

func (tftp *Tftp) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tftp *Tftp) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tftp *Tftp) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tftp *Tftp) SetParent(parent types.Entity) { tftp.parent = parent }

func (tftp *Tftp) GetParent() types.Entity { return tftp.parent }

func (tftp *Tftp) GetParentYangName() string { return "Cisco-IOS-XR-ipv4-filesystems-cfg" }

// Tftp_TftpClient
// TFTP client configuration
type Tftp_TftpClient struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Specify the number of retries when client requests TFTP connections. The
    // type is interface{} with range: 0..256.
    Retry interface{}

    // Specify the timeout for every TFTP connection in seconds. The type is
    // interface{} with range: 0..256. Units are second.
    Timeout interface{}

    // Specify interface for source address in connections. The type is string
    // with pattern: [a-zA-Z0-9./-]+.
    SourceInterface interface{}

    // VRF table.
    Vrfs Tftp_TftpClient_Vrfs
}

func (tftpClient *Tftp_TftpClient) GetFilter() yfilter.YFilter { return tftpClient.YFilter }

func (tftpClient *Tftp_TftpClient) SetFilter(yf yfilter.YFilter) { tftpClient.YFilter = yf }

func (tftpClient *Tftp_TftpClient) GetGoName(yname string) string {
    if yname == "retry" { return "Retry" }
    if yname == "timeout" { return "Timeout" }
    if yname == "source-interface" { return "SourceInterface" }
    if yname == "vrfs" { return "Vrfs" }
    return ""
}

func (tftpClient *Tftp_TftpClient) GetSegmentPath() string {
    return "tftp-client"
}

func (tftpClient *Tftp_TftpClient) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrfs" {
        return &tftpClient.Vrfs
    }
    return nil
}

func (tftpClient *Tftp_TftpClient) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["vrfs"] = &tftpClient.Vrfs
    return children
}

func (tftpClient *Tftp_TftpClient) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["retry"] = tftpClient.Retry
    leafs["timeout"] = tftpClient.Timeout
    leafs["source-interface"] = tftpClient.SourceInterface
    return leafs
}

func (tftpClient *Tftp_TftpClient) GetBundleName() string { return "cisco_ios_xr" }

func (tftpClient *Tftp_TftpClient) GetYangName() string { return "tftp-client" }

func (tftpClient *Tftp_TftpClient) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (tftpClient *Tftp_TftpClient) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (tftpClient *Tftp_TftpClient) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (tftpClient *Tftp_TftpClient) SetParent(parent types.Entity) { tftpClient.parent = parent }

func (tftpClient *Tftp_TftpClient) GetParent() types.Entity { return tftpClient.parent }

func (tftpClient *Tftp_TftpClient) GetParentYangName() string { return "tftp" }

// Tftp_TftpClient_Vrfs
// VRF table
type Tftp_TftpClient_Vrfs struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // VRF specific data. The type is slice of Tftp_TftpClient_Vrfs_Vrf.
    Vrf []Tftp_TftpClient_Vrfs_Vrf
}

func (vrfs *Tftp_TftpClient_Vrfs) GetFilter() yfilter.YFilter { return vrfs.YFilter }

func (vrfs *Tftp_TftpClient_Vrfs) SetFilter(yf yfilter.YFilter) { vrfs.YFilter = yf }

func (vrfs *Tftp_TftpClient_Vrfs) GetGoName(yname string) string {
    if yname == "vrf" { return "Vrf" }
    return ""
}

func (vrfs *Tftp_TftpClient_Vrfs) GetSegmentPath() string {
    return "vrfs"
}

func (vrfs *Tftp_TftpClient_Vrfs) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "vrf" {
        for _, c := range vrfs.Vrf {
            if vrfs.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := Tftp_TftpClient_Vrfs_Vrf{}
        vrfs.Vrf = append(vrfs.Vrf, child)
        return &vrfs.Vrf[len(vrfs.Vrf)-1]
    }
    return nil
}

func (vrfs *Tftp_TftpClient_Vrfs) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    for i := range vrfs.Vrf {
        children[vrfs.Vrf[i].GetSegmentPath()] = &vrfs.Vrf[i]
    }
    return children
}

func (vrfs *Tftp_TftpClient_Vrfs) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (vrfs *Tftp_TftpClient_Vrfs) GetBundleName() string { return "cisco_ios_xr" }

func (vrfs *Tftp_TftpClient_Vrfs) GetYangName() string { return "vrfs" }

func (vrfs *Tftp_TftpClient_Vrfs) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrfs *Tftp_TftpClient_Vrfs) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrfs *Tftp_TftpClient_Vrfs) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrfs *Tftp_TftpClient_Vrfs) SetParent(parent types.Entity) { vrfs.parent = parent }

func (vrfs *Tftp_TftpClient_Vrfs) GetParent() types.Entity { return vrfs.parent }

func (vrfs *Tftp_TftpClient_Vrfs) GetParentYangName() string { return "tftp-client" }

// Tftp_TftpClient_Vrfs_Vrf
// VRF specific data
type Tftp_TftpClient_Vrfs_Vrf struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the VRF instance. The type is string with
    // pattern: [\w\-\.:,_@#%$\+=\|;]+.
    VrfName interface{}

    // Specify interface for source address in connections. The type is string
    // with pattern: [a-zA-Z0-9./-]+.
    SourceInterface interface{}

    // Specify the number of retries when client requests TFTP connections. The
    // type is interface{} with range: 0..256.
    Retry interface{}

    // Specify the timeout for every TFTP connection in seconds. The type is
    // interface{} with range: 0..256. Units are second.
    Timeout interface{}
}

func (vrf *Tftp_TftpClient_Vrfs_Vrf) GetFilter() yfilter.YFilter { return vrf.YFilter }

func (vrf *Tftp_TftpClient_Vrfs_Vrf) SetFilter(yf yfilter.YFilter) { vrf.YFilter = yf }

func (vrf *Tftp_TftpClient_Vrfs_Vrf) GetGoName(yname string) string {
    if yname == "vrf-name" { return "VrfName" }
    if yname == "source-interface" { return "SourceInterface" }
    if yname == "retry" { return "Retry" }
    if yname == "timeout" { return "Timeout" }
    return ""
}

func (vrf *Tftp_TftpClient_Vrfs_Vrf) GetSegmentPath() string {
    return "vrf" + "[vrf-name='" + fmt.Sprintf("%v", vrf.VrfName) + "']"
}

func (vrf *Tftp_TftpClient_Vrfs_Vrf) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (vrf *Tftp_TftpClient_Vrfs_Vrf) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (vrf *Tftp_TftpClient_Vrfs_Vrf) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["vrf-name"] = vrf.VrfName
    leafs["source-interface"] = vrf.SourceInterface
    leafs["retry"] = vrf.Retry
    leafs["timeout"] = vrf.Timeout
    return leafs
}

func (vrf *Tftp_TftpClient_Vrfs_Vrf) GetBundleName() string { return "cisco_ios_xr" }

func (vrf *Tftp_TftpClient_Vrfs_Vrf) GetYangName() string { return "vrf" }

func (vrf *Tftp_TftpClient_Vrfs_Vrf) GetBundleYangModelsLocation() string { return cisco_ios_xr.GetModelsPath() }

func (vrf *Tftp_TftpClient_Vrfs_Vrf) GetCapabilitiesTable() map[string]string {
    return cisco_ios_xr.GetCapabilities() }

func (vrf *Tftp_TftpClient_Vrfs_Vrf) GetNamespaceTable() map[string]string {
    return cisco_ios_xr.GetNamespaces() }

func (vrf *Tftp_TftpClient_Vrfs_Vrf) SetParent(parent types.Entity) { vrf.parent = parent }

func (vrf *Tftp_TftpClient_Vrfs_Vrf) GetParent() types.Entity { return vrf.parent }

func (vrf *Tftp_TftpClient_Vrfs_Vrf) GetParentYangName() string { return "vrfs" }

