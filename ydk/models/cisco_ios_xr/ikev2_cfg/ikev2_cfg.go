// This module contains a collection of YANG definitions
// for Cisco IOS-XR ikev2 package configuration.
// 
// This module contains definitions
// for the following management objects:
//   ikev2: Internet key exchange(IKEv2) config commands
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package ikev2_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package ikev2_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-ikev2-cfg ikev2}", reflect.TypeOf(Ikev2{}))
    ydk.RegisterEntity("Cisco-IOS-XR-ikev2-cfg:ikev2", reflect.TypeOf(Ikev2{}))
}

// Ikev2
// Internet key exchange(IKEv2) config commands
type Ikev2 struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IKEv2 keyring config commands.
    KeyringNames Ikev2_KeyringNames

    // IKEv2 profile config commands.
    ProfileNames Ikev2_ProfileNames

    // Configure IKEv2 policies.
    PolicyNames Ikev2_PolicyNames

    // Configure IKEv2 proposals.
    ProposalNames Ikev2_ProposalNames
}

func (ikev2 *Ikev2) GetEntityData() *types.CommonEntityData {
    ikev2.EntityData.YFilter = ikev2.YFilter
    ikev2.EntityData.YangName = "ikev2"
    ikev2.EntityData.BundleName = "cisco_ios_xr"
    ikev2.EntityData.ParentYangName = "Cisco-IOS-XR-ikev2-cfg"
    ikev2.EntityData.SegmentPath = "Cisco-IOS-XR-ikev2-cfg:ikev2"
    ikev2.EntityData.AbsolutePath = ikev2.EntityData.SegmentPath
    ikev2.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    ikev2.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    ikev2.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    ikev2.EntityData.Children = types.NewOrderedMap()
    ikev2.EntityData.Children.Append("keyring-names", types.YChild{"KeyringNames", &ikev2.KeyringNames})
    ikev2.EntityData.Children.Append("profile-names", types.YChild{"ProfileNames", &ikev2.ProfileNames})
    ikev2.EntityData.Children.Append("policy-names", types.YChild{"PolicyNames", &ikev2.PolicyNames})
    ikev2.EntityData.Children.Append("proposal-names", types.YChild{"ProposalNames", &ikev2.ProposalNames})
    ikev2.EntityData.Leafs = types.NewOrderedMap()

    ikev2.EntityData.YListKeys = []string {}

    return &(ikev2.EntityData)
}

// Ikev2_KeyringNames
// IKEv2 keyring config commands
type Ikev2_KeyringNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IKEv2 keyring name. The type is slice of Ikev2_KeyringNames_KeyringName.
    KeyringName []*Ikev2_KeyringNames_KeyringName
}

func (keyringNames *Ikev2_KeyringNames) GetEntityData() *types.CommonEntityData {
    keyringNames.EntityData.YFilter = keyringNames.YFilter
    keyringNames.EntityData.YangName = "keyring-names"
    keyringNames.EntityData.BundleName = "cisco_ios_xr"
    keyringNames.EntityData.ParentYangName = "ikev2"
    keyringNames.EntityData.SegmentPath = "keyring-names"
    keyringNames.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-cfg:ikev2/" + keyringNames.EntityData.SegmentPath
    keyringNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keyringNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keyringNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keyringNames.EntityData.Children = types.NewOrderedMap()
    keyringNames.EntityData.Children.Append("keyring-name", types.YChild{"KeyringName", nil})
    for i := range keyringNames.KeyringName {
        keyringNames.EntityData.Children.Append(types.GetSegmentPath(keyringNames.KeyringName[i]), types.YChild{"KeyringName", keyringNames.KeyringName[i]})
    }
    keyringNames.EntityData.Leafs = types.NewOrderedMap()

    keyringNames.EntityData.YListKeys = []string {}

    return &(keyringNames.EntityData)
}

// Ikev2_KeyringNames_KeyringName
// IKEv2 keyring name
type Ikev2_KeyringNames_KeyringName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the keyring. The type is string with
    // length: 1..32.
    Name interface{}

    // This indicated existence of keyring. The type is interface{}.
    KeyringSub interface{}

    // IKEv2 keyring peer config commands.
    PeerNames Ikev2_KeyringNames_KeyringName_PeerNames
}

func (keyringName *Ikev2_KeyringNames_KeyringName) GetEntityData() *types.CommonEntityData {
    keyringName.EntityData.YFilter = keyringName.YFilter
    keyringName.EntityData.YangName = "keyring-name"
    keyringName.EntityData.BundleName = "cisco_ios_xr"
    keyringName.EntityData.ParentYangName = "keyring-names"
    keyringName.EntityData.SegmentPath = "keyring-name" + types.AddKeyToken(keyringName.Name, "name")
    keyringName.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-cfg:ikev2/keyring-names/" + keyringName.EntityData.SegmentPath
    keyringName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    keyringName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    keyringName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    keyringName.EntityData.Children = types.NewOrderedMap()
    keyringName.EntityData.Children.Append("peer-names", types.YChild{"PeerNames", &keyringName.PeerNames})
    keyringName.EntityData.Leafs = types.NewOrderedMap()
    keyringName.EntityData.Leafs.Append("name", types.YLeaf{"Name", keyringName.Name})
    keyringName.EntityData.Leafs.Append("keyring-sub", types.YLeaf{"KeyringSub", keyringName.KeyringSub})

    keyringName.EntityData.YListKeys = []string {"Name"}

    return &(keyringName.EntityData)
}

// Ikev2_KeyringNames_KeyringName_PeerNames
// IKEv2 keyring peer config commands
type Ikev2_KeyringNames_KeyringName_PeerNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IKEv2 keyring peer name. The type is slice of
    // Ikev2_KeyringNames_KeyringName_PeerNames_PeerName.
    PeerName []*Ikev2_KeyringNames_KeyringName_PeerNames_PeerName
}

func (peerNames *Ikev2_KeyringNames_KeyringName_PeerNames) GetEntityData() *types.CommonEntityData {
    peerNames.EntityData.YFilter = peerNames.YFilter
    peerNames.EntityData.YangName = "peer-names"
    peerNames.EntityData.BundleName = "cisco_ios_xr"
    peerNames.EntityData.ParentYangName = "keyring-name"
    peerNames.EntityData.SegmentPath = "peer-names"
    peerNames.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-cfg:ikev2/keyring-names/keyring-name/" + peerNames.EntityData.SegmentPath
    peerNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerNames.EntityData.Children = types.NewOrderedMap()
    peerNames.EntityData.Children.Append("peer-name", types.YChild{"PeerName", nil})
    for i := range peerNames.PeerName {
        peerNames.EntityData.Children.Append(types.GetSegmentPath(peerNames.PeerName[i]), types.YChild{"PeerName", peerNames.PeerName[i]})
    }
    peerNames.EntityData.Leafs = types.NewOrderedMap()

    peerNames.EntityData.YListKeys = []string {}

    return &(peerNames.EntityData)
}

// Ikev2_KeyringNames_KeyringName_PeerNames_PeerName
// IKEv2 keyring peer name
type Ikev2_KeyringNames_KeyringName_PeerNames_PeerName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the keyring-peer. The type is string with
    // length: 1..32.
    Name interface{}

    // This indicates existence of keyring-peer. The type is interface{}.
    PeerSub interface{}

    // IP Address to identify the peer.
    Address Ikev2_KeyringNames_KeyringName_PeerNames_PeerName_Address

    // Pre-shared key for peer.
    Psk Ikev2_KeyringNames_KeyringName_PeerNames_PeerName_Psk
}

func (peerName *Ikev2_KeyringNames_KeyringName_PeerNames_PeerName) GetEntityData() *types.CommonEntityData {
    peerName.EntityData.YFilter = peerName.YFilter
    peerName.EntityData.YangName = "peer-name"
    peerName.EntityData.BundleName = "cisco_ios_xr"
    peerName.EntityData.ParentYangName = "peer-names"
    peerName.EntityData.SegmentPath = "peer-name" + types.AddKeyToken(peerName.Name, "name")
    peerName.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-cfg:ikev2/keyring-names/keyring-name/peer-names/" + peerName.EntityData.SegmentPath
    peerName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    peerName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    peerName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    peerName.EntityData.Children = types.NewOrderedMap()
    peerName.EntityData.Children.Append("address", types.YChild{"Address", &peerName.Address})
    peerName.EntityData.Children.Append("psk", types.YChild{"Psk", &peerName.Psk})
    peerName.EntityData.Leafs = types.NewOrderedMap()
    peerName.EntityData.Leafs.Append("name", types.YLeaf{"Name", peerName.Name})
    peerName.EntityData.Leafs.Append("peer-sub", types.YLeaf{"PeerSub", peerName.PeerSub})

    peerName.EntityData.YListKeys = []string {"Name"}

    return &(peerName.EntityData)
}

// Ikev2_KeyringNames_KeyringName_PeerNames_PeerName_Address
// IP Address to identify the peer
type Ikev2_KeyringNames_KeyringName_PeerNames_PeerName_Address struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IP Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Ip interface{}

    // Subnet. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Subnet interface{}
}

func (address *Ikev2_KeyringNames_KeyringName_PeerNames_PeerName_Address) GetEntityData() *types.CommonEntityData {
    address.EntityData.YFilter = address.YFilter
    address.EntityData.YangName = "address"
    address.EntityData.BundleName = "cisco_ios_xr"
    address.EntityData.ParentYangName = "peer-name"
    address.EntityData.SegmentPath = "address"
    address.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-cfg:ikev2/keyring-names/keyring-name/peer-names/peer-name/" + address.EntityData.SegmentPath
    address.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    address.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    address.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    address.EntityData.Children = types.NewOrderedMap()
    address.EntityData.Leafs = types.NewOrderedMap()
    address.EntityData.Leafs.Append("ip", types.YLeaf{"Ip", address.Ip})
    address.EntityData.Leafs.Append("subnet", types.YLeaf{"Subnet", address.Subnet})

    address.EntityData.YListKeys = []string {}

    return &(address.EntityData)
}

// Ikev2_KeyringNames_KeyringName_PeerNames_PeerName_Psk
// Pre-shared key for peer
type Ikev2_KeyringNames_KeyringName_PeerNames_PeerName_Psk struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Both pre-shared key for the peer. The type is string with pattern:
    // b'(!.+)|([^!].+)'.
    BothKey interface{}

    // Local/Remote pre-shared key for the peer.
    LocalRemoteKey Ikev2_KeyringNames_KeyringName_PeerNames_PeerName_Psk_LocalRemoteKey
}

func (psk *Ikev2_KeyringNames_KeyringName_PeerNames_PeerName_Psk) GetEntityData() *types.CommonEntityData {
    psk.EntityData.YFilter = psk.YFilter
    psk.EntityData.YangName = "psk"
    psk.EntityData.BundleName = "cisco_ios_xr"
    psk.EntityData.ParentYangName = "peer-name"
    psk.EntityData.SegmentPath = "psk"
    psk.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-cfg:ikev2/keyring-names/keyring-name/peer-names/peer-name/" + psk.EntityData.SegmentPath
    psk.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    psk.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    psk.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    psk.EntityData.Children = types.NewOrderedMap()
    psk.EntityData.Children.Append("local-remote-key", types.YChild{"LocalRemoteKey", &psk.LocalRemoteKey})
    psk.EntityData.Leafs = types.NewOrderedMap()
    psk.EntityData.Leafs.Append("both-key", types.YLeaf{"BothKey", psk.BothKey})

    psk.EntityData.YListKeys = []string {}

    return &(psk.EntityData)
}

// Ikev2_KeyringNames_KeyringName_PeerNames_PeerName_Psk_LocalRemoteKey
// Local/Remote pre-shared key for the peer
type Ikev2_KeyringNames_KeyringName_PeerNames_PeerName_Psk_LocalRemoteKey struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Local pre-shared key. The type is string with pattern: b'(!.+)|([^!].+)'.
    StringXr interface{}

    // Remote pre-shared key. The type is string with pattern: b'(!.+)|([^!].+)'.
    String interface{}
}

func (localRemoteKey *Ikev2_KeyringNames_KeyringName_PeerNames_PeerName_Psk_LocalRemoteKey) GetEntityData() *types.CommonEntityData {
    localRemoteKey.EntityData.YFilter = localRemoteKey.YFilter
    localRemoteKey.EntityData.YangName = "local-remote-key"
    localRemoteKey.EntityData.BundleName = "cisco_ios_xr"
    localRemoteKey.EntityData.ParentYangName = "psk"
    localRemoteKey.EntityData.SegmentPath = "local-remote-key"
    localRemoteKey.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-cfg:ikev2/keyring-names/keyring-name/peer-names/peer-name/psk/" + localRemoteKey.EntityData.SegmentPath
    localRemoteKey.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    localRemoteKey.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    localRemoteKey.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    localRemoteKey.EntityData.Children = types.NewOrderedMap()
    localRemoteKey.EntityData.Leafs = types.NewOrderedMap()
    localRemoteKey.EntityData.Leafs.Append("string-xr", types.YLeaf{"StringXr", localRemoteKey.StringXr})
    localRemoteKey.EntityData.Leafs.Append("string", types.YLeaf{"String", localRemoteKey.String})

    localRemoteKey.EntityData.YListKeys = []string {}

    return &(localRemoteKey.EntityData)
}

// Ikev2_ProfileNames
// IKEv2 profile config commands
type Ikev2_ProfileNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IKEv2 profile name. The type is slice of Ikev2_ProfileNames_ProfileName.
    ProfileName []*Ikev2_ProfileNames_ProfileName
}

func (profileNames *Ikev2_ProfileNames) GetEntityData() *types.CommonEntityData {
    profileNames.EntityData.YFilter = profileNames.YFilter
    profileNames.EntityData.YangName = "profile-names"
    profileNames.EntityData.BundleName = "cisco_ios_xr"
    profileNames.EntityData.ParentYangName = "ikev2"
    profileNames.EntityData.SegmentPath = "profile-names"
    profileNames.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-cfg:ikev2/" + profileNames.EntityData.SegmentPath
    profileNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profileNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profileNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profileNames.EntityData.Children = types.NewOrderedMap()
    profileNames.EntityData.Children.Append("profile-name", types.YChild{"ProfileName", nil})
    for i := range profileNames.ProfileName {
        profileNames.EntityData.Children.Append(types.GetSegmentPath(profileNames.ProfileName[i]), types.YChild{"ProfileName", profileNames.ProfileName[i]})
    }
    profileNames.EntityData.Leafs = types.NewOrderedMap()

    profileNames.EntityData.YListKeys = []string {}

    return &(profileNames.EntityData)
}

// Ikev2_ProfileNames_ProfileName
// IKEv2 profile name
type Ikev2_ProfileNames_ProfileName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the profile. The type is string with
    // length: 1..32.
    Name interface{}

    // This indicates existence of profile. The type is interface{}.
    ProfileSub interface{}

    // Lifetime(in sec) for IKEv2 SA. The type is interface{} with range:
    // 120..86400.
    Lifetime interface{}

    // Keyring to use with local/remote authentication method. The type is string
    // with length: 1..32.
    KeyringInProfile interface{}

    // Match a profile based on remote identity.
    MatchIdentity Ikev2_ProfileNames_ProfileName_MatchIdentity

    // Enable IKEv2 liveliness for peers.
    Dpd Ikev2_ProfileNames_ProfileName_Dpd
}

func (profileName *Ikev2_ProfileNames_ProfileName) GetEntityData() *types.CommonEntityData {
    profileName.EntityData.YFilter = profileName.YFilter
    profileName.EntityData.YangName = "profile-name"
    profileName.EntityData.BundleName = "cisco_ios_xr"
    profileName.EntityData.ParentYangName = "profile-names"
    profileName.EntityData.SegmentPath = "profile-name" + types.AddKeyToken(profileName.Name, "name")
    profileName.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-cfg:ikev2/profile-names/" + profileName.EntityData.SegmentPath
    profileName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    profileName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    profileName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    profileName.EntityData.Children = types.NewOrderedMap()
    profileName.EntityData.Children.Append("match-identity", types.YChild{"MatchIdentity", &profileName.MatchIdentity})
    profileName.EntityData.Children.Append("dpd", types.YChild{"Dpd", &profileName.Dpd})
    profileName.EntityData.Leafs = types.NewOrderedMap()
    profileName.EntityData.Leafs.Append("name", types.YLeaf{"Name", profileName.Name})
    profileName.EntityData.Leafs.Append("profile-sub", types.YLeaf{"ProfileSub", profileName.ProfileSub})
    profileName.EntityData.Leafs.Append("lifetime", types.YLeaf{"Lifetime", profileName.Lifetime})
    profileName.EntityData.Leafs.Append("keyring-in-profile", types.YLeaf{"KeyringInProfile", profileName.KeyringInProfile})

    profileName.EntityData.YListKeys = []string {"Name"}

    return &(profileName.EntityData)
}

// Ikev2_ProfileNames_ProfileName_MatchIdentity
// Match a profile based on remote identity
type Ikev2_ProfileNames_ProfileName_MatchIdentity struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Match any peer identity. The type is interface{}.
    Any interface{}

    // Match a profile based on remote identity address.
    AddressSubs Ikev2_ProfileNames_ProfileName_MatchIdentity_AddressSubs
}

func (matchIdentity *Ikev2_ProfileNames_ProfileName_MatchIdentity) GetEntityData() *types.CommonEntityData {
    matchIdentity.EntityData.YFilter = matchIdentity.YFilter
    matchIdentity.EntityData.YangName = "match-identity"
    matchIdentity.EntityData.BundleName = "cisco_ios_xr"
    matchIdentity.EntityData.ParentYangName = "profile-name"
    matchIdentity.EntityData.SegmentPath = "match-identity"
    matchIdentity.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-cfg:ikev2/profile-names/profile-name/" + matchIdentity.EntityData.SegmentPath
    matchIdentity.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    matchIdentity.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    matchIdentity.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    matchIdentity.EntityData.Children = types.NewOrderedMap()
    matchIdentity.EntityData.Children.Append("address-subs", types.YChild{"AddressSubs", &matchIdentity.AddressSubs})
    matchIdentity.EntityData.Leafs = types.NewOrderedMap()
    matchIdentity.EntityData.Leafs.Append("any", types.YLeaf{"Any", matchIdentity.Any})

    matchIdentity.EntityData.YListKeys = []string {}

    return &(matchIdentity.EntityData)
}

// Ikev2_ProfileNames_ProfileName_MatchIdentity_AddressSubs
// Match a profile based on remote identity
// address
type Ikev2_ProfileNames_ProfileName_MatchIdentity_AddressSubs struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Remote ip address for matching identity. The type is slice of
    // Ikev2_ProfileNames_ProfileName_MatchIdentity_AddressSubs_AddressSub.
    AddressSub []*Ikev2_ProfileNames_ProfileName_MatchIdentity_AddressSubs_AddressSub
}

func (addressSubs *Ikev2_ProfileNames_ProfileName_MatchIdentity_AddressSubs) GetEntityData() *types.CommonEntityData {
    addressSubs.EntityData.YFilter = addressSubs.YFilter
    addressSubs.EntityData.YangName = "address-subs"
    addressSubs.EntityData.BundleName = "cisco_ios_xr"
    addressSubs.EntityData.ParentYangName = "match-identity"
    addressSubs.EntityData.SegmentPath = "address-subs"
    addressSubs.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-cfg:ikev2/profile-names/profile-name/match-identity/" + addressSubs.EntityData.SegmentPath
    addressSubs.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressSubs.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressSubs.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressSubs.EntityData.Children = types.NewOrderedMap()
    addressSubs.EntityData.Children.Append("address-sub", types.YChild{"AddressSub", nil})
    for i := range addressSubs.AddressSub {
        addressSubs.EntityData.Children.Append(types.GetSegmentPath(addressSubs.AddressSub[i]), types.YChild{"AddressSub", addressSubs.AddressSub[i]})
    }
    addressSubs.EntityData.Leafs = types.NewOrderedMap()

    addressSubs.EntityData.YListKeys = []string {}

    return &(addressSubs.EntityData)
}

// Ikev2_ProfileNames_ProfileName_MatchIdentity_AddressSubs_AddressSub
// Remote ip address for matching identity
type Ikev2_ProfileNames_ProfileName_MatchIdentity_AddressSubs_AddressSub struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}

    // This indicates existence of remote ip address. The type is interface{}.
    AddressSubVal interface{}

    // Mask. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Mask interface{}
}

func (addressSub *Ikev2_ProfileNames_ProfileName_MatchIdentity_AddressSubs_AddressSub) GetEntityData() *types.CommonEntityData {
    addressSub.EntityData.YFilter = addressSub.YFilter
    addressSub.EntityData.YangName = "address-sub"
    addressSub.EntityData.BundleName = "cisco_ios_xr"
    addressSub.EntityData.ParentYangName = "address-subs"
    addressSub.EntityData.SegmentPath = "address-sub" + types.AddKeyToken(addressSub.Address, "address")
    addressSub.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-cfg:ikev2/profile-names/profile-name/match-identity/address-subs/" + addressSub.EntityData.SegmentPath
    addressSub.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressSub.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressSub.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressSub.EntityData.Children = types.NewOrderedMap()
    addressSub.EntityData.Leafs = types.NewOrderedMap()
    addressSub.EntityData.Leafs.Append("address", types.YLeaf{"Address", addressSub.Address})
    addressSub.EntityData.Leafs.Append("address-sub-val", types.YLeaf{"AddressSubVal", addressSub.AddressSubVal})
    addressSub.EntityData.Leafs.Append("mask", types.YLeaf{"Mask", addressSub.Mask})

    addressSub.EntityData.YListKeys = []string {"Address"}

    return &(addressSub.EntityData)
}

// Ikev2_ProfileNames_ProfileName_Dpd
// Enable IKEv2 liveliness for peers
type Ikev2_ProfileNames_ProfileName_Dpd struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Interval(in sec). The type is interface{} with range: 10..3600.
    Interval interface{}

    // Retry interval(in sec). The type is interface{} with range: 2..60.
    RetryTime interface{}
}

func (dpd *Ikev2_ProfileNames_ProfileName_Dpd) GetEntityData() *types.CommonEntityData {
    dpd.EntityData.YFilter = dpd.YFilter
    dpd.EntityData.YangName = "dpd"
    dpd.EntityData.BundleName = "cisco_ios_xr"
    dpd.EntityData.ParentYangName = "profile-name"
    dpd.EntityData.SegmentPath = "dpd"
    dpd.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-cfg:ikev2/profile-names/profile-name/" + dpd.EntityData.SegmentPath
    dpd.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    dpd.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    dpd.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    dpd.EntityData.Children = types.NewOrderedMap()
    dpd.EntityData.Leafs = types.NewOrderedMap()
    dpd.EntityData.Leafs.Append("interval", types.YLeaf{"Interval", dpd.Interval})
    dpd.EntityData.Leafs.Append("retry-time", types.YLeaf{"RetryTime", dpd.RetryTime})

    dpd.EntityData.YListKeys = []string {}

    return &(dpd.EntityData)
}

// Ikev2_PolicyNames
// Configure IKEv2 policies
type Ikev2_PolicyNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IKEv2 policy name. The type is slice of Ikev2_PolicyNames_PolicyName.
    PolicyName []*Ikev2_PolicyNames_PolicyName
}

func (policyNames *Ikev2_PolicyNames) GetEntityData() *types.CommonEntityData {
    policyNames.EntityData.YFilter = policyNames.YFilter
    policyNames.EntityData.YangName = "policy-names"
    policyNames.EntityData.BundleName = "cisco_ios_xr"
    policyNames.EntityData.ParentYangName = "ikev2"
    policyNames.EntityData.SegmentPath = "policy-names"
    policyNames.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-cfg:ikev2/" + policyNames.EntityData.SegmentPath
    policyNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyNames.EntityData.Children = types.NewOrderedMap()
    policyNames.EntityData.Children.Append("policy-name", types.YChild{"PolicyName", nil})
    for i := range policyNames.PolicyName {
        policyNames.EntityData.Children.Append(types.GetSegmentPath(policyNames.PolicyName[i]), types.YChild{"PolicyName", policyNames.PolicyName[i]})
    }
    policyNames.EntityData.Leafs = types.NewOrderedMap()

    policyNames.EntityData.YListKeys = []string {}

    return &(policyNames.EntityData)
}

// Ikev2_PolicyNames_PolicyName
// IKEv2 policy name
type Ikev2_PolicyNames_PolicyName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Policy name. The type is string with length:
    // 1..32.
    Name interface{}

    // Proposal to use with configured policy. The type is string with length:
    // 1..32.
    ProposalInPolicy interface{}

    // This indicates existence of policy. The type is interface{}.
    PolicySub interface{}

    // Match a policy based on address.
    AddressVals Ikev2_PolicyNames_PolicyName_AddressVals
}

func (policyName *Ikev2_PolicyNames_PolicyName) GetEntityData() *types.CommonEntityData {
    policyName.EntityData.YFilter = policyName.YFilter
    policyName.EntityData.YangName = "policy-name"
    policyName.EntityData.BundleName = "cisco_ios_xr"
    policyName.EntityData.ParentYangName = "policy-names"
    policyName.EntityData.SegmentPath = "policy-name" + types.AddKeyToken(policyName.Name, "name")
    policyName.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-cfg:ikev2/policy-names/" + policyName.EntityData.SegmentPath
    policyName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    policyName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    policyName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    policyName.EntityData.Children = types.NewOrderedMap()
    policyName.EntityData.Children.Append("address-vals", types.YChild{"AddressVals", &policyName.AddressVals})
    policyName.EntityData.Leafs = types.NewOrderedMap()
    policyName.EntityData.Leafs.Append("name", types.YLeaf{"Name", policyName.Name})
    policyName.EntityData.Leafs.Append("proposal-in-policy", types.YLeaf{"ProposalInPolicy", policyName.ProposalInPolicy})
    policyName.EntityData.Leafs.Append("policy-sub", types.YLeaf{"PolicySub", policyName.PolicySub})

    policyName.EntityData.YListKeys = []string {"Name"}

    return &(policyName.EntityData)
}

// Ikev2_PolicyNames_PolicyName_AddressVals
// Match a policy based on address
type Ikev2_PolicyNames_PolicyName_AddressVals struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // local address used to match policy. The type is slice of
    // Ikev2_PolicyNames_PolicyName_AddressVals_AddressVal.
    AddressVal []*Ikev2_PolicyNames_PolicyName_AddressVals_AddressVal
}

func (addressVals *Ikev2_PolicyNames_PolicyName_AddressVals) GetEntityData() *types.CommonEntityData {
    addressVals.EntityData.YFilter = addressVals.YFilter
    addressVals.EntityData.YangName = "address-vals"
    addressVals.EntityData.BundleName = "cisco_ios_xr"
    addressVals.EntityData.ParentYangName = "policy-name"
    addressVals.EntityData.SegmentPath = "address-vals"
    addressVals.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-cfg:ikev2/policy-names/policy-name/" + addressVals.EntityData.SegmentPath
    addressVals.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressVals.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressVals.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressVals.EntityData.Children = types.NewOrderedMap()
    addressVals.EntityData.Children.Append("address-val", types.YChild{"AddressVal", nil})
    for i := range addressVals.AddressVal {
        addressVals.EntityData.Children.Append(types.GetSegmentPath(addressVals.AddressVal[i]), types.YChild{"AddressVal", addressVals.AddressVal[i]})
    }
    addressVals.EntityData.Leafs = types.NewOrderedMap()

    addressVals.EntityData.YListKeys = []string {}

    return &(addressVals.EntityData)
}

// Ikev2_PolicyNames_PolicyName_AddressVals_AddressVal
// local address used to match policy
type Ikev2_PolicyNames_PolicyName_AddressVals_AddressVal struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Address. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Address interface{}
}

func (addressVal *Ikev2_PolicyNames_PolicyName_AddressVals_AddressVal) GetEntityData() *types.CommonEntityData {
    addressVal.EntityData.YFilter = addressVal.YFilter
    addressVal.EntityData.YangName = "address-val"
    addressVal.EntityData.BundleName = "cisco_ios_xr"
    addressVal.EntityData.ParentYangName = "address-vals"
    addressVal.EntityData.SegmentPath = "address-val" + types.AddKeyToken(addressVal.Address, "address")
    addressVal.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-cfg:ikev2/policy-names/policy-name/address-vals/" + addressVal.EntityData.SegmentPath
    addressVal.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    addressVal.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    addressVal.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    addressVal.EntityData.Children = types.NewOrderedMap()
    addressVal.EntityData.Leafs = types.NewOrderedMap()
    addressVal.EntityData.Leafs.Append("address", types.YLeaf{"Address", addressVal.Address})

    addressVal.EntityData.YListKeys = []string {"Address"}

    return &(addressVal.EntityData)
}

// Ikev2_ProposalNames
// Configure IKEv2 proposals
type Ikev2_ProposalNames struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // IKEv2 proposal name. The type is slice of Ikev2_ProposalNames_ProposalName.
    ProposalName []*Ikev2_ProposalNames_ProposalName
}

func (proposalNames *Ikev2_ProposalNames) GetEntityData() *types.CommonEntityData {
    proposalNames.EntityData.YFilter = proposalNames.YFilter
    proposalNames.EntityData.YangName = "proposal-names"
    proposalNames.EntityData.BundleName = "cisco_ios_xr"
    proposalNames.EntityData.ParentYangName = "ikev2"
    proposalNames.EntityData.SegmentPath = "proposal-names"
    proposalNames.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-cfg:ikev2/" + proposalNames.EntityData.SegmentPath
    proposalNames.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    proposalNames.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    proposalNames.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    proposalNames.EntityData.Children = types.NewOrderedMap()
    proposalNames.EntityData.Children.Append("proposal-name", types.YChild{"ProposalName", nil})
    for i := range proposalNames.ProposalName {
        proposalNames.EntityData.Children.Append(types.GetSegmentPath(proposalNames.ProposalName[i]), types.YChild{"ProposalName", proposalNames.ProposalName[i]})
    }
    proposalNames.EntityData.Leafs = types.NewOrderedMap()

    proposalNames.EntityData.YListKeys = []string {}

    return &(proposalNames.EntityData)
}

// Ikev2_ProposalNames_ProposalName
// IKEv2 proposal name
type Ikev2_ProposalNames_ProposalName struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Proposal name. The type is string with length:
    // 1..32.
    Name interface{}

    // This indicates existence of proposal. The type is interface{}.
    ProposalSub interface{}

    // Specify one or more transforms of prf.
    Prfses Ikev2_ProposalNames_ProposalName_Prfses

    // Specify one or more transforms of group.
    Groups Ikev2_ProposalNames_ProposalName_Groups

    // Specify one or more transforms of integrity.
    Integrities Ikev2_ProposalNames_ProposalName_Integrities

    // Specify one or more transforms of encryption.
    Encryptions Ikev2_ProposalNames_ProposalName_Encryptions
}

func (proposalName *Ikev2_ProposalNames_ProposalName) GetEntityData() *types.CommonEntityData {
    proposalName.EntityData.YFilter = proposalName.YFilter
    proposalName.EntityData.YangName = "proposal-name"
    proposalName.EntityData.BundleName = "cisco_ios_xr"
    proposalName.EntityData.ParentYangName = "proposal-names"
    proposalName.EntityData.SegmentPath = "proposal-name" + types.AddKeyToken(proposalName.Name, "name")
    proposalName.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-cfg:ikev2/proposal-names/" + proposalName.EntityData.SegmentPath
    proposalName.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    proposalName.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    proposalName.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    proposalName.EntityData.Children = types.NewOrderedMap()
    proposalName.EntityData.Children.Append("prfses", types.YChild{"Prfses", &proposalName.Prfses})
    proposalName.EntityData.Children.Append("groups", types.YChild{"Groups", &proposalName.Groups})
    proposalName.EntityData.Children.Append("integrities", types.YChild{"Integrities", &proposalName.Integrities})
    proposalName.EntityData.Children.Append("encryptions", types.YChild{"Encryptions", &proposalName.Encryptions})
    proposalName.EntityData.Leafs = types.NewOrderedMap()
    proposalName.EntityData.Leafs.Append("name", types.YLeaf{"Name", proposalName.Name})
    proposalName.EntityData.Leafs.Append("proposal-sub", types.YLeaf{"ProposalSub", proposalName.ProposalSub})

    proposalName.EntityData.YListKeys = []string {"Name"}

    return &(proposalName.EntityData)
}

// Ikev2_ProposalNames_ProposalName_Prfses
// Specify one or more transforms of prf
type Ikev2_ProposalNames_ProposalName_Prfses struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // PRF Algorithm. The type is slice of string with length: 1..8.
    Prfs []interface{}
}

func (prfses *Ikev2_ProposalNames_ProposalName_Prfses) GetEntityData() *types.CommonEntityData {
    prfses.EntityData.YFilter = prfses.YFilter
    prfses.EntityData.YangName = "prfses"
    prfses.EntityData.BundleName = "cisco_ios_xr"
    prfses.EntityData.ParentYangName = "proposal-name"
    prfses.EntityData.SegmentPath = "prfses"
    prfses.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-cfg:ikev2/proposal-names/proposal-name/" + prfses.EntityData.SegmentPath
    prfses.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prfses.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prfses.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prfses.EntityData.Children = types.NewOrderedMap()
    prfses.EntityData.Leafs = types.NewOrderedMap()
    prfses.EntityData.Leafs.Append("prfs", types.YLeaf{"Prfs", prfses.Prfs})

    prfses.EntityData.YListKeys = []string {}

    return &(prfses.EntityData)
}

// Ikev2_ProposalNames_ProposalName_Groups
// Specify one or more transforms of group
type Ikev2_ProposalNames_ProposalName_Groups struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Encryption Algorithm. The type is slice of string with length: 1..3.
    Group []interface{}
}

func (groups *Ikev2_ProposalNames_ProposalName_Groups) GetEntityData() *types.CommonEntityData {
    groups.EntityData.YFilter = groups.YFilter
    groups.EntityData.YangName = "groups"
    groups.EntityData.BundleName = "cisco_ios_xr"
    groups.EntityData.ParentYangName = "proposal-name"
    groups.EntityData.SegmentPath = "groups"
    groups.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-cfg:ikev2/proposal-names/proposal-name/" + groups.EntityData.SegmentPath
    groups.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    groups.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    groups.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    groups.EntityData.Children = types.NewOrderedMap()
    groups.EntityData.Leafs = types.NewOrderedMap()
    groups.EntityData.Leafs.Append("group", types.YLeaf{"Group", groups.Group})

    groups.EntityData.YListKeys = []string {}

    return &(groups.EntityData)
}

// Ikev2_ProposalNames_ProposalName_Integrities
// Specify one or more transforms of integrity
type Ikev2_ProposalNames_ProposalName_Integrities struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Integrity Algorithm. The type is slice of string with length: 1..8.
    Integrity []interface{}
}

func (integrities *Ikev2_ProposalNames_ProposalName_Integrities) GetEntityData() *types.CommonEntityData {
    integrities.EntityData.YFilter = integrities.YFilter
    integrities.EntityData.YangName = "integrities"
    integrities.EntityData.BundleName = "cisco_ios_xr"
    integrities.EntityData.ParentYangName = "proposal-name"
    integrities.EntityData.SegmentPath = "integrities"
    integrities.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-cfg:ikev2/proposal-names/proposal-name/" + integrities.EntityData.SegmentPath
    integrities.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    integrities.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    integrities.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    integrities.EntityData.Children = types.NewOrderedMap()
    integrities.EntityData.Leafs = types.NewOrderedMap()
    integrities.EntityData.Leafs.Append("integrity", types.YLeaf{"Integrity", integrities.Integrity})

    integrities.EntityData.YListKeys = []string {}

    return &(integrities.EntityData)
}

// Ikev2_ProposalNames_ProposalName_Encryptions
// Specify one or more transforms of encryption
type Ikev2_ProposalNames_ProposalName_Encryptions struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Encryption Algorithm. The type is slice of string with length: 1..12.
    Encryption []interface{}
}

func (encryptions *Ikev2_ProposalNames_ProposalName_Encryptions) GetEntityData() *types.CommonEntityData {
    encryptions.EntityData.YFilter = encryptions.YFilter
    encryptions.EntityData.YangName = "encryptions"
    encryptions.EntityData.BundleName = "cisco_ios_xr"
    encryptions.EntityData.ParentYangName = "proposal-name"
    encryptions.EntityData.SegmentPath = "encryptions"
    encryptions.EntityData.AbsolutePath = "Cisco-IOS-XR-ikev2-cfg:ikev2/proposal-names/proposal-name/" + encryptions.EntityData.SegmentPath
    encryptions.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    encryptions.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    encryptions.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    encryptions.EntityData.Children = types.NewOrderedMap()
    encryptions.EntityData.Leafs = types.NewOrderedMap()
    encryptions.EntityData.Leafs.Append("encryption", types.YLeaf{"Encryption", encryptions.Encryption})

    encryptions.EntityData.YListKeys = []string {}

    return &(encryptions.EntityData)
}

