// This module contains a collection of YANG definitions
// for Cisco IOS-XR tunnel-l2tun package configuration.
// 
// This module contains definitions
// for the following management objects:
//   l2tp: L2TPv3 class used for L2VPNs
// 
// Copyright (c) 2013-2017 by Cisco Systems, Inc.
// All rights reserved.
package tunnel_l2tun_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package tunnel_l2tun_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-tunnel-l2tun-cfg l2tp}", reflect.TypeOf(L2Tp{}))
    ydk.RegisterEntity("Cisco-IOS-XR-tunnel-l2tun-cfg:l2tp", reflect.TypeOf(L2Tp{}))
}

// L2tpHashMethod represents L2tp hash method
type L2tpHashMethod string

const (
    // MD5
    L2tpHashMethod_md5 L2tpHashMethod = "md5"

    // SHA1
    L2tpHashMethod_sha1 L2tpHashMethod = "sha1"

    // None
    L2tpHashMethod_none L2tpHashMethod = "none"
)

// L2tpDigestHashMethod represents L2tp digest hash method
type L2tpDigestHashMethod string

const (
    // MD5
    L2tpDigestHashMethod_md5 L2tpDigestHashMethod = "md5"

    // SHA1
    L2tpDigestHashMethod_sha1 L2tpDigestHashMethod = "sha1"
)

// L2Tp
// L2TPv3 class used for L2VPNs
type L2Tp struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // List of classes.
    Classes L2Tp_Classes
}

func (l2Tp *L2Tp) GetEntityData() *types.CommonEntityData {
    l2Tp.EntityData.YFilter = l2Tp.YFilter
    l2Tp.EntityData.YangName = "l2tp"
    l2Tp.EntityData.BundleName = "cisco_ios_xr"
    l2Tp.EntityData.ParentYangName = "Cisco-IOS-XR-tunnel-l2tun-cfg"
    l2Tp.EntityData.SegmentPath = "Cisco-IOS-XR-tunnel-l2tun-cfg:l2tp"
    l2Tp.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    l2Tp.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    l2Tp.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    l2Tp.EntityData.Children = make(map[string]types.YChild)
    l2Tp.EntityData.Children["classes"] = types.YChild{"Classes", &l2Tp.Classes}
    l2Tp.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(l2Tp.EntityData)
}

// L2Tp_Classes
// List of classes
type L2Tp_Classes struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Configuration for a specific class. The type is slice of
    // L2Tp_Classes_Class.
    Class []L2Tp_Classes_Class
}

func (classes *L2Tp_Classes) GetEntityData() *types.CommonEntityData {
    classes.EntityData.YFilter = classes.YFilter
    classes.EntityData.YangName = "classes"
    classes.EntityData.BundleName = "cisco_ios_xr"
    classes.EntityData.ParentYangName = "l2tp"
    classes.EntityData.SegmentPath = "classes"
    classes.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    classes.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    classes.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    classes.EntityData.Children = make(map[string]types.YChild)
    classes.EntityData.Children["class"] = types.YChild{"Class", nil}
    for i := range classes.Class {
        classes.EntityData.Children[types.GetSegmentPath(&classes.Class[i])] = types.YChild{"Class", &classes.Class[i]}
    }
    classes.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(classes.EntityData)
}

// L2Tp_Classes_Class
// Configuration for a specific class
type L2Tp_Classes_Class struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specify the class name. Regexp:
    // ^[a-z0-9A-Z][-_.a-z0-9A-Z]*$. The type is string with length: 1..31.
    ClassName interface{}

    // Local hostname for control connection authentication. The type is string.
    HostName interface{}

    // Specify to hide AVPs in outgoing control messages. The type is interface{}.
    Hidden interface{}

    // Specify interval (in seconds). The type is interface{} with range: 0..1000.
    // Units are second.
    HelloInterval interface{}

    // Time permitted to set up a control connection. The type is interface{} with
    // range: 60..6000. Units are second.
    TimeoutSetup interface{}

    // Receive window size for the control connection. The type is interface{}
    // with range: 1..16384. Units are byte.
    ReceiveWindow interface{}

    // Congestion control enabled. The type is interface{}.
    CongestionControl interface{}

    // Timeout value for no-user in seconds. The type is interface{} with range:
    // 0..4294967295. Units are second.
    TimeoutNoUser interface{}

    // Authenticate the L2TP control connection. The type is interface{} with
    // range: -2147483648..2147483647.
    Authentication interface{}

    // Enable L2TPv3 class used for L2VPNs. The type is interface{}.
    Enable interface{}

    // Specify the password for control channel authentication. The type is string
    // with pattern: b'(!.+)|([^!].+)'.
    Password interface{}

    // Security check.
    Security L2Tp_Classes_Class_Security

    // Control message retransmission parameters.
    Retransmit L2Tp_Classes_Class_Retransmit

    // l2TP tunnel.
    Tunnel L2Tp_Classes_Class_Tunnel

    // Message digest authentication for the L2TP control connection.
    Digest L2Tp_Classes_Class_Digest

    // IP TOS value.
    Ip L2Tp_Classes_Class_Ip
}

func (class *L2Tp_Classes_Class) GetEntityData() *types.CommonEntityData {
    class.EntityData.YFilter = class.YFilter
    class.EntityData.YangName = "class"
    class.EntityData.BundleName = "cisco_ios_xr"
    class.EntityData.ParentYangName = "classes"
    class.EntityData.SegmentPath = "class" + "[class-name='" + fmt.Sprintf("%v", class.ClassName) + "']"
    class.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    class.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    class.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    class.EntityData.Children = make(map[string]types.YChild)
    class.EntityData.Children["security"] = types.YChild{"Security", &class.Security}
    class.EntityData.Children["retransmit"] = types.YChild{"Retransmit", &class.Retransmit}
    class.EntityData.Children["tunnel"] = types.YChild{"Tunnel", &class.Tunnel}
    class.EntityData.Children["digest"] = types.YChild{"Digest", &class.Digest}
    class.EntityData.Children["ip"] = types.YChild{"Ip", &class.Ip}
    class.EntityData.Leafs = make(map[string]types.YLeaf)
    class.EntityData.Leafs["class-name"] = types.YLeaf{"ClassName", class.ClassName}
    class.EntityData.Leafs["host-name"] = types.YLeaf{"HostName", class.HostName}
    class.EntityData.Leafs["hidden"] = types.YLeaf{"Hidden", class.Hidden}
    class.EntityData.Leafs["hello-interval"] = types.YLeaf{"HelloInterval", class.HelloInterval}
    class.EntityData.Leafs["timeout-setup"] = types.YLeaf{"TimeoutSetup", class.TimeoutSetup}
    class.EntityData.Leafs["receive-window"] = types.YLeaf{"ReceiveWindow", class.ReceiveWindow}
    class.EntityData.Leafs["congestion-control"] = types.YLeaf{"CongestionControl", class.CongestionControl}
    class.EntityData.Leafs["timeout-no-user"] = types.YLeaf{"TimeoutNoUser", class.TimeoutNoUser}
    class.EntityData.Leafs["authentication"] = types.YLeaf{"Authentication", class.Authentication}
    class.EntityData.Leafs["enable"] = types.YLeaf{"Enable", class.Enable}
    class.EntityData.Leafs["password"] = types.YLeaf{"Password", class.Password}
    return &(class.EntityData)
}

// L2Tp_Classes_Class_Security
// Security check
type L2Tp_Classes_Class_Security struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Security check for IP.
    Ip L2Tp_Classes_Class_Security_Ip
}

func (security *L2Tp_Classes_Class_Security) GetEntityData() *types.CommonEntityData {
    security.EntityData.YFilter = security.YFilter
    security.EntityData.YangName = "security"
    security.EntityData.BundleName = "cisco_ios_xr"
    security.EntityData.ParentYangName = "class"
    security.EntityData.SegmentPath = "security"
    security.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    security.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    security.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    security.EntityData.Children = make(map[string]types.YChild)
    security.EntityData.Children["ip"] = types.YChild{"Ip", &security.Ip}
    security.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(security.EntityData)
}

// L2Tp_Classes_Class_Security_Ip
// Security check for IP
type L2Tp_Classes_Class_Security_Ip struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enable IP address check for L2TP packets. The type is interface{}.
    AddressCheck interface{}
}

func (ip *L2Tp_Classes_Class_Security_Ip) GetEntityData() *types.CommonEntityData {
    ip.EntityData.YFilter = ip.YFilter
    ip.EntityData.YangName = "ip"
    ip.EntityData.BundleName = "cisco_ios_xr"
    ip.EntityData.ParentYangName = "security"
    ip.EntityData.SegmentPath = "ip"
    ip.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ip.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ip.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ip.EntityData.Children = make(map[string]types.YChild)
    ip.EntityData.Leafs = make(map[string]types.YLeaf)
    ip.EntityData.Leafs["address-check"] = types.YLeaf{"AddressCheck", ip.AddressCheck}
    return &(ip.EntityData)
}

// L2Tp_Classes_Class_Retransmit
// Control message retransmission parameters
type L2Tp_Classes_Class_Retransmit struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify retransmit retry count. The type is interface{} with range:
    // 5..1000.
    Retry interface{}

    // Set retries and timeouts for initial.
    Initial L2Tp_Classes_Class_Retransmit_Initial

    // Set timeout value range.
    Timeout L2Tp_Classes_Class_Retransmit_Timeout
}

func (retransmit *L2Tp_Classes_Class_Retransmit) GetEntityData() *types.CommonEntityData {
    retransmit.EntityData.YFilter = retransmit.YFilter
    retransmit.EntityData.YangName = "retransmit"
    retransmit.EntityData.BundleName = "cisco_ios_xr"
    retransmit.EntityData.ParentYangName = "class"
    retransmit.EntityData.SegmentPath = "retransmit"
    retransmit.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    retransmit.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    retransmit.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    retransmit.EntityData.Children = make(map[string]types.YChild)
    retransmit.EntityData.Children["initial"] = types.YChild{"Initial", &retransmit.Initial}
    retransmit.EntityData.Children["timeout"] = types.YChild{"Timeout", &retransmit.Timeout}
    retransmit.EntityData.Leafs = make(map[string]types.YLeaf)
    retransmit.EntityData.Leafs["retry"] = types.YLeaf{"Retry", retransmit.Retry}
    return &(retransmit.EntityData)
}

// L2Tp_Classes_Class_Retransmit_Initial
// Set retries and timeouts for initial
type L2Tp_Classes_Class_Retransmit_Initial struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify the retry number. The type is interface{} with range: 1..1000.
    Retry interface{}

    // Set timeout value range.
    Timeout L2Tp_Classes_Class_Retransmit_Initial_Timeout
}

func (initial *L2Tp_Classes_Class_Retransmit_Initial) GetEntityData() *types.CommonEntityData {
    initial.EntityData.YFilter = initial.YFilter
    initial.EntityData.YangName = "initial"
    initial.EntityData.BundleName = "cisco_ios_xr"
    initial.EntityData.ParentYangName = "retransmit"
    initial.EntityData.SegmentPath = "initial"
    initial.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    initial.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    initial.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    initial.EntityData.Children = make(map[string]types.YChild)
    initial.EntityData.Children["timeout"] = types.YChild{"Timeout", &initial.Timeout}
    initial.EntityData.Leafs = make(map[string]types.YLeaf)
    initial.EntityData.Leafs["retry"] = types.YLeaf{"Retry", initial.Retry}
    return &(initial.EntityData)
}

// L2Tp_Classes_Class_Retransmit_Initial_Timeout
// Set timeout value range
type L2Tp_Classes_Class_Retransmit_Initial_Timeout struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify minimum timeout. The type is interface{} with range: 1..8.
    Minimum interface{}

    // Specify maximum timeout. The type is interface{} with range: 1..8.
    Maximum interface{}
}

func (timeout *L2Tp_Classes_Class_Retransmit_Initial_Timeout) GetEntityData() *types.CommonEntityData {
    timeout.EntityData.YFilter = timeout.YFilter
    timeout.EntityData.YangName = "timeout"
    timeout.EntityData.BundleName = "cisco_ios_xr"
    timeout.EntityData.ParentYangName = "initial"
    timeout.EntityData.SegmentPath = "timeout"
    timeout.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timeout.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timeout.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timeout.EntityData.Children = make(map[string]types.YChild)
    timeout.EntityData.Leafs = make(map[string]types.YLeaf)
    timeout.EntityData.Leafs["minimum"] = types.YLeaf{"Minimum", timeout.Minimum}
    timeout.EntityData.Leafs["maximum"] = types.YLeaf{"Maximum", timeout.Maximum}
    return &(timeout.EntityData)
}

// L2Tp_Classes_Class_Retransmit_Timeout
// Set timeout value range
type L2Tp_Classes_Class_Retransmit_Timeout struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify minimum timeout. The type is interface{} with range: 1..8.
    Minimum interface{}

    // Specify maximum timeout. The type is interface{} with range: 1..8.
    Maximum interface{}
}

func (timeout *L2Tp_Classes_Class_Retransmit_Timeout) GetEntityData() *types.CommonEntityData {
    timeout.EntityData.YFilter = timeout.YFilter
    timeout.EntityData.YangName = "timeout"
    timeout.EntityData.BundleName = "cisco_ios_xr"
    timeout.EntityData.ParentYangName = "retransmit"
    timeout.EntityData.SegmentPath = "timeout"
    timeout.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    timeout.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    timeout.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    timeout.EntityData.Children = make(map[string]types.YChild)
    timeout.EntityData.Leafs = make(map[string]types.YLeaf)
    timeout.EntityData.Leafs["minimum"] = types.YLeaf{"Minimum", timeout.Minimum}
    timeout.EntityData.Leafs["maximum"] = types.YLeaf{"Maximum", timeout.Maximum}
    return &(timeout.EntityData)
}

// L2Tp_Classes_Class_Tunnel
// l2TP tunnel
type L2Tp_Classes_Class_Tunnel struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tunnel accounting. The type is string.
    Accounting interface{}
}

func (tunnel *L2Tp_Classes_Class_Tunnel) GetEntityData() *types.CommonEntityData {
    tunnel.EntityData.YFilter = tunnel.YFilter
    tunnel.EntityData.YangName = "tunnel"
    tunnel.EntityData.BundleName = "cisco_ios_xr"
    tunnel.EntityData.ParentYangName = "class"
    tunnel.EntityData.SegmentPath = "tunnel"
    tunnel.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    tunnel.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    tunnel.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    tunnel.EntityData.Children = make(map[string]types.YChild)
    tunnel.EntityData.Leafs = make(map[string]types.YLeaf)
    tunnel.EntityData.Leafs["accounting"] = types.YLeaf{"Accounting", tunnel.Accounting}
    return &(tunnel.EntityData)
}

// L2Tp_Classes_Class_Digest
// Message digest authentication for the L2TP
// control connection
type L2Tp_Classes_Class_Digest struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Specify hash method. The type is L2tpDigestHashMethod.
    Hash interface{}

    // Disable digest checking. The type is interface{}.
    CheckDisable interface{}

    // Set shared secret for message digest.
    Secrets L2Tp_Classes_Class_Digest_Secrets
}

func (digest *L2Tp_Classes_Class_Digest) GetEntityData() *types.CommonEntityData {
    digest.EntityData.YFilter = digest.YFilter
    digest.EntityData.YangName = "digest"
    digest.EntityData.BundleName = "cisco_ios_xr"
    digest.EntityData.ParentYangName = "class"
    digest.EntityData.SegmentPath = "digest"
    digest.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    digest.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    digest.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    digest.EntityData.Children = make(map[string]types.YChild)
    digest.EntityData.Children["secrets"] = types.YChild{"Secrets", &digest.Secrets}
    digest.EntityData.Leafs = make(map[string]types.YLeaf)
    digest.EntityData.Leafs["hash"] = types.YLeaf{"Hash", digest.Hash}
    digest.EntityData.Leafs["check-disable"] = types.YLeaf{"CheckDisable", digest.CheckDisable}
    return &(digest.EntityData)
}

// L2Tp_Classes_Class_Digest_Secrets
// Set shared secret for message digest
type L2Tp_Classes_Class_Digest_Secrets struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The encrypted user secret and hash method. The type is slice of
    // L2Tp_Classes_Class_Digest_Secrets_Secret.
    Secret []L2Tp_Classes_Class_Digest_Secrets_Secret
}

func (secrets *L2Tp_Classes_Class_Digest_Secrets) GetEntityData() *types.CommonEntityData {
    secrets.EntityData.YFilter = secrets.YFilter
    secrets.EntityData.YangName = "secrets"
    secrets.EntityData.BundleName = "cisco_ios_xr"
    secrets.EntityData.ParentYangName = "digest"
    secrets.EntityData.SegmentPath = "secrets"
    secrets.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secrets.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secrets.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secrets.EntityData.Children = make(map[string]types.YChild)
    secrets.EntityData.Children["secret"] = types.YChild{"Secret", nil}
    for i := range secrets.Secret {
        secrets.EntityData.Children[types.GetSegmentPath(&secrets.Secret[i])] = types.YChild{"Secret", &secrets.Secret[i]}
    }
    secrets.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(secrets.EntityData)
}

// L2Tp_Classes_Class_Digest_Secrets_Secret
// The encrypted user secret and hash method
type L2Tp_Classes_Class_Digest_Secrets_Secret struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Specify the encrypted user secret. The type is
    // string with pattern: b'[\\w\\-\\.:,_@#%$\\+=\\|;]+'.
    SecretName interface{}

    // Specify hash method. The type is L2tpHashMethod. This attribute is
    // mandatory.
    Hash interface{}
}

func (secret *L2Tp_Classes_Class_Digest_Secrets_Secret) GetEntityData() *types.CommonEntityData {
    secret.EntityData.YFilter = secret.YFilter
    secret.EntityData.YangName = "secret"
    secret.EntityData.BundleName = "cisco_ios_xr"
    secret.EntityData.ParentYangName = "secrets"
    secret.EntityData.SegmentPath = "secret" + "[secret-name='" + fmt.Sprintf("%v", secret.SecretName) + "']"
    secret.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    secret.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    secret.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    secret.EntityData.Children = make(map[string]types.YChild)
    secret.EntityData.Leafs = make(map[string]types.YLeaf)
    secret.EntityData.Leafs["secret-name"] = types.YLeaf{"SecretName", secret.SecretName}
    secret.EntityData.Leafs["hash"] = types.YLeaf{"Hash", secret.Hash}
    return &(secret.EntityData)
}

// L2Tp_Classes_Class_Ip
// IP TOS value
type L2Tp_Classes_Class_Ip struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP TOS value (decimal). The type is interface{} with range: 0..255.
    Tos interface{}
}

func (ip *L2Tp_Classes_Class_Ip) GetEntityData() *types.CommonEntityData {
    ip.EntityData.YFilter = ip.YFilter
    ip.EntityData.YangName = "ip"
    ip.EntityData.BundleName = "cisco_ios_xr"
    ip.EntityData.ParentYangName = "class"
    ip.EntityData.SegmentPath = "ip"
    ip.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ip.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ip.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ip.EntityData.Children = make(map[string]types.YChild)
    ip.EntityData.Leafs = make(map[string]types.YLeaf)
    ip.EntityData.Leafs["tos"] = types.YLeaf{"Tos", ip.Tos}
    return &(ip.EntityData)
}

