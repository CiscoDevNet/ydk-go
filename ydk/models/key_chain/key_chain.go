// This YANG module defines the generic configuration
// data for key-chain. It is intended that the module
// will be extended by vendors to define vendor-specific
// key-chain configuration parameters.
// 
package key_chain

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/ietf"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package key_chain"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:ietf-key-chain key-chains}", reflect.TypeOf(KeyChains{}))
    ydk.RegisterEntity("ietf-key-chain:key-chains", reflect.TypeOf(KeyChains{}))
}

// KeyChains
// A key-chain is a sequence of keys that are collectively
// managed for authentication.
type KeyChains struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Name of the key-chain. The type is string.
    Name interface{}

    // Tolerance for key lifetime acceptance (seconds).
    AcceptTolerance KeyChains_AcceptTolerance

    // One key. The type is slice of KeyChains_Key.
    Key []*KeyChains_Key
}

func (keyChains *KeyChains) GetEntityData() *types.CommonEntityData {
    keyChains.EntityData.YFilter = keyChains.YFilter
    keyChains.EntityData.YangName = "key-chains"
    keyChains.EntityData.BundleName = "ietf"
    keyChains.EntityData.ParentYangName = "ietf-key-chain"
    keyChains.EntityData.SegmentPath = "ietf-key-chain:key-chains" + types.AddKeyToken(keyChains.Name, "name")
    keyChains.EntityData.AbsolutePath = keyChains.EntityData.SegmentPath
    keyChains.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    keyChains.EntityData.NamespaceTable = ietf.GetNamespaces()
    keyChains.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    keyChains.EntityData.Children = types.NewOrderedMap()
    keyChains.EntityData.Children.Append("accept-tolerance", types.YChild{"AcceptTolerance", &keyChains.AcceptTolerance})
    keyChains.EntityData.Children.Append("key", types.YChild{"Key", nil})
    for i := range keyChains.Key {
        keyChains.EntityData.Children.Append(types.GetSegmentPath(keyChains.Key[i]), types.YChild{"Key", keyChains.Key[i]})
    }
    keyChains.EntityData.Leafs = types.NewOrderedMap()
    keyChains.EntityData.Leafs.Append("name", types.YLeaf{"Name", keyChains.Name})

    keyChains.EntityData.YListKeys = []string {"Name"}

    return &(keyChains.EntityData)
}

// KeyChains_AcceptTolerance
// Tolerance for key lifetime acceptance (seconds).
type KeyChains_AcceptTolerance struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Tolerance range, in seconds. The type is interface{} with range:
    // 0..4294967295. Units are seconds. The default value is 0.
    Duration interface{}
}

func (acceptTolerance *KeyChains_AcceptTolerance) GetEntityData() *types.CommonEntityData {
    acceptTolerance.EntityData.YFilter = acceptTolerance.YFilter
    acceptTolerance.EntityData.YangName = "accept-tolerance"
    acceptTolerance.EntityData.BundleName = "ietf"
    acceptTolerance.EntityData.ParentYangName = "key-chains"
    acceptTolerance.EntityData.SegmentPath = "accept-tolerance"
    acceptTolerance.EntityData.AbsolutePath = "ietf-key-chain:key-chains/" + acceptTolerance.EntityData.SegmentPath
    acceptTolerance.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    acceptTolerance.EntityData.NamespaceTable = ietf.GetNamespaces()
    acceptTolerance.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    acceptTolerance.EntityData.Children = types.NewOrderedMap()
    acceptTolerance.EntityData.Leafs = types.NewOrderedMap()
    acceptTolerance.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", acceptTolerance.Duration})

    acceptTolerance.EntityData.YListKeys = []string {}

    return &(acceptTolerance.EntityData)
}

// KeyChains_Key
// One key.
type KeyChains_Key struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. Key id. The type is interface{} with range:
    // 0..18446744073709551615.
    KeyId interface{}

    // The key string.
    KeyString KeyChains_Key_KeyString

    // Specify a key's lifetime.
    Lifetime KeyChains_Key_Lifetime

    // Cryptographic algorithm associated with key.
    CryptoAlgorithm KeyChains_Key_CryptoAlgorithm
}

func (key *KeyChains_Key) GetEntityData() *types.CommonEntityData {
    key.EntityData.YFilter = key.YFilter
    key.EntityData.YangName = "key"
    key.EntityData.BundleName = "ietf"
    key.EntityData.ParentYangName = "key-chains"
    key.EntityData.SegmentPath = "key" + types.AddKeyToken(key.KeyId, "key-id")
    key.EntityData.AbsolutePath = "ietf-key-chain:key-chains/" + key.EntityData.SegmentPath
    key.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    key.EntityData.NamespaceTable = ietf.GetNamespaces()
    key.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    key.EntityData.Children = types.NewOrderedMap()
    key.EntityData.Children.Append("key-string", types.YChild{"KeyString", &key.KeyString})
    key.EntityData.Children.Append("lifetime", types.YChild{"Lifetime", &key.Lifetime})
    key.EntityData.Children.Append("crypto-algorithm", types.YChild{"CryptoAlgorithm", &key.CryptoAlgorithm})
    key.EntityData.Leafs = types.NewOrderedMap()
    key.EntityData.Leafs.Append("key-id", types.YLeaf{"KeyId", key.KeyId})

    key.EntityData.YListKeys = []string {"KeyId"}

    return &(key.EntityData)
}

// KeyChains_Key_KeyString
// The key string.
type KeyChains_Key_KeyString struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Key string in ASCII format. The type is string.
    Keystring interface{}

    // Key in hexadecimal string format. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    HexadecimalString interface{}
}

func (keyString *KeyChains_Key_KeyString) GetEntityData() *types.CommonEntityData {
    keyString.EntityData.YFilter = keyString.YFilter
    keyString.EntityData.YangName = "key-string"
    keyString.EntityData.BundleName = "ietf"
    keyString.EntityData.ParentYangName = "key"
    keyString.EntityData.SegmentPath = "key-string"
    keyString.EntityData.AbsolutePath = "ietf-key-chain:key-chains/key/" + keyString.EntityData.SegmentPath
    keyString.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    keyString.EntityData.NamespaceTable = ietf.GetNamespaces()
    keyString.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    keyString.EntityData.Children = types.NewOrderedMap()
    keyString.EntityData.Leafs = types.NewOrderedMap()
    keyString.EntityData.Leafs.Append("keystring", types.YLeaf{"Keystring", keyString.Keystring})
    keyString.EntityData.Leafs.Append("hexadecimal-string", types.YLeaf{"HexadecimalString", keyString.HexadecimalString})

    keyString.EntityData.YListKeys = []string {}

    return &(keyString.EntityData)
}

// KeyChains_Key_Lifetime
// Specify a key's lifetime.
type KeyChains_Key_Lifetime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Single lifetime specification for both send and accept lifetimes.
    SendAcceptLifetime KeyChains_Key_Lifetime_SendAcceptLifetime

    // Separate lifetime specification for send lifetime.
    SendLifetime KeyChains_Key_Lifetime_SendLifetime

    // Separate lifetime specification for accept lifetime.
    AcceptLifetime KeyChains_Key_Lifetime_AcceptLifetime
}

func (lifetime *KeyChains_Key_Lifetime) GetEntityData() *types.CommonEntityData {
    lifetime.EntityData.YFilter = lifetime.YFilter
    lifetime.EntityData.YangName = "lifetime"
    lifetime.EntityData.BundleName = "ietf"
    lifetime.EntityData.ParentYangName = "key"
    lifetime.EntityData.SegmentPath = "lifetime"
    lifetime.EntityData.AbsolutePath = "ietf-key-chain:key-chains/key/" + lifetime.EntityData.SegmentPath
    lifetime.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    lifetime.EntityData.NamespaceTable = ietf.GetNamespaces()
    lifetime.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    lifetime.EntityData.Children = types.NewOrderedMap()
    lifetime.EntityData.Children.Append("send-accept-lifetime", types.YChild{"SendAcceptLifetime", &lifetime.SendAcceptLifetime})
    lifetime.EntityData.Children.Append("send-lifetime", types.YChild{"SendLifetime", &lifetime.SendLifetime})
    lifetime.EntityData.Children.Append("accept-lifetime", types.YChild{"AcceptLifetime", &lifetime.AcceptLifetime})
    lifetime.EntityData.Leafs = types.NewOrderedMap()

    lifetime.EntityData.YListKeys = []string {}

    return &(lifetime.EntityData)
}

// KeyChains_Key_Lifetime_SendAcceptLifetime
// Single lifetime specification for both send and
// accept lifetimes.
type KeyChains_Key_Lifetime_SendAcceptLifetime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates key lifetime is always valid. The type is interface{}.
    Always interface{}

    // Start time. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    StartDateTime interface{}

    // Indicates key lifetime end-time in infinite. The type is interface{}.
    NoEndTime interface{}

    // Key lifetime duration, in seconds. The type is interface{} with range:
    // 1..2147483646. Units are seconds.
    Duration interface{}

    // End time. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    EndDateTime interface{}
}

func (sendAcceptLifetime *KeyChains_Key_Lifetime_SendAcceptLifetime) GetEntityData() *types.CommonEntityData {
    sendAcceptLifetime.EntityData.YFilter = sendAcceptLifetime.YFilter
    sendAcceptLifetime.EntityData.YangName = "send-accept-lifetime"
    sendAcceptLifetime.EntityData.BundleName = "ietf"
    sendAcceptLifetime.EntityData.ParentYangName = "lifetime"
    sendAcceptLifetime.EntityData.SegmentPath = "send-accept-lifetime"
    sendAcceptLifetime.EntityData.AbsolutePath = "ietf-key-chain:key-chains/key/lifetime/" + sendAcceptLifetime.EntityData.SegmentPath
    sendAcceptLifetime.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    sendAcceptLifetime.EntityData.NamespaceTable = ietf.GetNamespaces()
    sendAcceptLifetime.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    sendAcceptLifetime.EntityData.Children = types.NewOrderedMap()
    sendAcceptLifetime.EntityData.Leafs = types.NewOrderedMap()
    sendAcceptLifetime.EntityData.Leafs.Append("always", types.YLeaf{"Always", sendAcceptLifetime.Always})
    sendAcceptLifetime.EntityData.Leafs.Append("start-date-time", types.YLeaf{"StartDateTime", sendAcceptLifetime.StartDateTime})
    sendAcceptLifetime.EntityData.Leafs.Append("no-end-time", types.YLeaf{"NoEndTime", sendAcceptLifetime.NoEndTime})
    sendAcceptLifetime.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", sendAcceptLifetime.Duration})
    sendAcceptLifetime.EntityData.Leafs.Append("end-date-time", types.YLeaf{"EndDateTime", sendAcceptLifetime.EndDateTime})

    sendAcceptLifetime.EntityData.YListKeys = []string {}

    return &(sendAcceptLifetime.EntityData)
}

// KeyChains_Key_Lifetime_SendLifetime
// Separate lifetime specification for send
// lifetime.
type KeyChains_Key_Lifetime_SendLifetime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates key lifetime is always valid. The type is interface{}.
    Always interface{}

    // Start time. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    StartDateTime interface{}

    // Indicates key lifetime end-time in infinite. The type is interface{}.
    NoEndTime interface{}

    // Key lifetime duration, in seconds. The type is interface{} with range:
    // 1..2147483646. Units are seconds.
    Duration interface{}

    // End time. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    EndDateTime interface{}
}

func (sendLifetime *KeyChains_Key_Lifetime_SendLifetime) GetEntityData() *types.CommonEntityData {
    sendLifetime.EntityData.YFilter = sendLifetime.YFilter
    sendLifetime.EntityData.YangName = "send-lifetime"
    sendLifetime.EntityData.BundleName = "ietf"
    sendLifetime.EntityData.ParentYangName = "lifetime"
    sendLifetime.EntityData.SegmentPath = "send-lifetime"
    sendLifetime.EntityData.AbsolutePath = "ietf-key-chain:key-chains/key/lifetime/" + sendLifetime.EntityData.SegmentPath
    sendLifetime.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    sendLifetime.EntityData.NamespaceTable = ietf.GetNamespaces()
    sendLifetime.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    sendLifetime.EntityData.Children = types.NewOrderedMap()
    sendLifetime.EntityData.Leafs = types.NewOrderedMap()
    sendLifetime.EntityData.Leafs.Append("always", types.YLeaf{"Always", sendLifetime.Always})
    sendLifetime.EntityData.Leafs.Append("start-date-time", types.YLeaf{"StartDateTime", sendLifetime.StartDateTime})
    sendLifetime.EntityData.Leafs.Append("no-end-time", types.YLeaf{"NoEndTime", sendLifetime.NoEndTime})
    sendLifetime.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", sendLifetime.Duration})
    sendLifetime.EntityData.Leafs.Append("end-date-time", types.YLeaf{"EndDateTime", sendLifetime.EndDateTime})

    sendLifetime.EntityData.YListKeys = []string {}

    return &(sendLifetime.EntityData)
}

// KeyChains_Key_Lifetime_AcceptLifetime
// Separate lifetime specification for accept
// lifetime.
type KeyChains_Key_Lifetime_AcceptLifetime struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Indicates key lifetime is always valid. The type is interface{}.
    Always interface{}

    // Start time. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    StartDateTime interface{}

    // Indicates key lifetime end-time in infinite. The type is interface{}.
    NoEndTime interface{}

    // Key lifetime duration, in seconds. The type is interface{} with range:
    // 1..2147483646. Units are seconds.
    Duration interface{}

    // End time. The type is string with pattern:
    // \d{4}-\d{2}-\d{2}T\d{2}:\d{2}:\d{2}(\.\d+)?(Z|[\+\-]\d{2}:\d{2}).
    EndDateTime interface{}
}

func (acceptLifetime *KeyChains_Key_Lifetime_AcceptLifetime) GetEntityData() *types.CommonEntityData {
    acceptLifetime.EntityData.YFilter = acceptLifetime.YFilter
    acceptLifetime.EntityData.YangName = "accept-lifetime"
    acceptLifetime.EntityData.BundleName = "ietf"
    acceptLifetime.EntityData.ParentYangName = "lifetime"
    acceptLifetime.EntityData.SegmentPath = "accept-lifetime"
    acceptLifetime.EntityData.AbsolutePath = "ietf-key-chain:key-chains/key/lifetime/" + acceptLifetime.EntityData.SegmentPath
    acceptLifetime.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    acceptLifetime.EntityData.NamespaceTable = ietf.GetNamespaces()
    acceptLifetime.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    acceptLifetime.EntityData.Children = types.NewOrderedMap()
    acceptLifetime.EntityData.Leafs = types.NewOrderedMap()
    acceptLifetime.EntityData.Leafs.Append("always", types.YLeaf{"Always", acceptLifetime.Always})
    acceptLifetime.EntityData.Leafs.Append("start-date-time", types.YLeaf{"StartDateTime", acceptLifetime.StartDateTime})
    acceptLifetime.EntityData.Leafs.Append("no-end-time", types.YLeaf{"NoEndTime", acceptLifetime.NoEndTime})
    acceptLifetime.EntityData.Leafs.Append("duration", types.YLeaf{"Duration", acceptLifetime.Duration})
    acceptLifetime.EntityData.Leafs.Append("end-date-time", types.YLeaf{"EndDateTime", acceptLifetime.EndDateTime})

    acceptLifetime.EntityData.YListKeys = []string {}

    return &(acceptLifetime.EntityData)
}

// KeyChains_Key_CryptoAlgorithm
// Cryptographic algorithm associated with key.
type KeyChains_Key_CryptoAlgorithm struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The HMAC-SHA1-12 algorithm. The type is interface{}.
    HmacSha112 interface{}

    // The HMAC-SHA1-20 algorithm. The type is interface{}.
    HmacSha120 interface{}

    // The MD5 algorithm. The type is interface{}.
    Md5 interface{}

    // The SHA-1 algorithm. The type is interface{}.
    Sha1 interface{}

    // HMAC-SHA-1 authentication algorithm. The type is interface{}.
    HmacSha1 interface{}

    // HMAC-SHA-256 authentication algorithm. The type is interface{}.
    HmacSha256 interface{}

    // HMAC-SHA-384 authentication algorithm. The type is interface{}.
    HmacSha384 interface{}

    // HMAC-SHA-512 authentication algorithm. The type is interface{}.
    HmacSha512 interface{}
}

func (cryptoAlgorithm *KeyChains_Key_CryptoAlgorithm) GetEntityData() *types.CommonEntityData {
    cryptoAlgorithm.EntityData.YFilter = cryptoAlgorithm.YFilter
    cryptoAlgorithm.EntityData.YangName = "crypto-algorithm"
    cryptoAlgorithm.EntityData.BundleName = "ietf"
    cryptoAlgorithm.EntityData.ParentYangName = "key"
    cryptoAlgorithm.EntityData.SegmentPath = "crypto-algorithm"
    cryptoAlgorithm.EntityData.AbsolutePath = "ietf-key-chain:key-chains/key/" + cryptoAlgorithm.EntityData.SegmentPath
    cryptoAlgorithm.EntityData.CapabilitiesTable = ietf.GetCapabilities()
    cryptoAlgorithm.EntityData.NamespaceTable = ietf.GetNamespaces()
    cryptoAlgorithm.EntityData.BundleYangModelsLocation = ietf.GetModelsPath()

    cryptoAlgorithm.EntityData.Children = types.NewOrderedMap()
    cryptoAlgorithm.EntityData.Leafs = types.NewOrderedMap()
    cryptoAlgorithm.EntityData.Leafs.Append("hmac-sha1-12", types.YLeaf{"HmacSha112", cryptoAlgorithm.HmacSha112})
    cryptoAlgorithm.EntityData.Leafs.Append("hmac-sha1-20", types.YLeaf{"HmacSha120", cryptoAlgorithm.HmacSha120})
    cryptoAlgorithm.EntityData.Leafs.Append("md5", types.YLeaf{"Md5", cryptoAlgorithm.Md5})
    cryptoAlgorithm.EntityData.Leafs.Append("sha-1", types.YLeaf{"Sha1", cryptoAlgorithm.Sha1})
    cryptoAlgorithm.EntityData.Leafs.Append("hmac-sha-1", types.YLeaf{"HmacSha1", cryptoAlgorithm.HmacSha1})
    cryptoAlgorithm.EntityData.Leafs.Append("hmac-sha-256", types.YLeaf{"HmacSha256", cryptoAlgorithm.HmacSha256})
    cryptoAlgorithm.EntityData.Leafs.Append("hmac-sha-384", types.YLeaf{"HmacSha384", cryptoAlgorithm.HmacSha384})
    cryptoAlgorithm.EntityData.Leafs.Append("hmac-sha-512", types.YLeaf{"HmacSha512", cryptoAlgorithm.HmacSha512})

    cryptoAlgorithm.EntityData.YListKeys = []string {}

    return &(cryptoAlgorithm.EntityData)
}

