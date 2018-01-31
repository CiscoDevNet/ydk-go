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
    parent types.Entity
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the key-chain. The type is string.
    Name interface{}

    // Tolerance for key lifetime acceptance (seconds).
    AcceptTolerance KeyChains_AcceptTolerance

    // One key. The type is slice of KeyChains_Key.
    Key []KeyChains_Key
}

func (keyChains *KeyChains) GetFilter() yfilter.YFilter { return keyChains.YFilter }

func (keyChains *KeyChains) SetFilter(yf yfilter.YFilter) { keyChains.YFilter = yf }

func (keyChains *KeyChains) GetGoName(yname string) string {
    if yname == "name" { return "Name" }
    if yname == "accept-tolerance" { return "AcceptTolerance" }
    if yname == "key" { return "Key" }
    return ""
}

func (keyChains *KeyChains) GetSegmentPath() string {
    return "ietf-key-chain:key-chains" + "[name='" + fmt.Sprintf("%v", keyChains.Name) + "']"
}

func (keyChains *KeyChains) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "accept-tolerance" {
        return &keyChains.AcceptTolerance
    }
    if childYangName == "key" {
        for _, c := range keyChains.Key {
            if keyChains.GetSegmentPath() == segmentPath {
                return &c
            }
        }
        child := KeyChains_Key{}
        keyChains.Key = append(keyChains.Key, child)
        return &keyChains.Key[len(keyChains.Key)-1]
    }
    return nil
}

func (keyChains *KeyChains) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["accept-tolerance"] = &keyChains.AcceptTolerance
    for i := range keyChains.Key {
        children[keyChains.Key[i].GetSegmentPath()] = &keyChains.Key[i]
    }
    return children
}

func (keyChains *KeyChains) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["name"] = keyChains.Name
    return leafs
}

func (keyChains *KeyChains) GetBundleName() string { return "ietf" }

func (keyChains *KeyChains) GetYangName() string { return "key-chains" }

func (keyChains *KeyChains) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (keyChains *KeyChains) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (keyChains *KeyChains) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (keyChains *KeyChains) SetParent(parent types.Entity) { keyChains.parent = parent }

func (keyChains *KeyChains) GetParent() types.Entity { return keyChains.parent }

func (keyChains *KeyChains) GetParentYangName() string { return "ietf-key-chain" }

// KeyChains_AcceptTolerance
// Tolerance for key lifetime acceptance (seconds).
type KeyChains_AcceptTolerance struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Tolerance range, in seconds. The type is interface{} with range:
    // 0..4294967295. Units are seconds. The default value is 0.
    Duration interface{}
}

func (acceptTolerance *KeyChains_AcceptTolerance) GetFilter() yfilter.YFilter { return acceptTolerance.YFilter }

func (acceptTolerance *KeyChains_AcceptTolerance) SetFilter(yf yfilter.YFilter) { acceptTolerance.YFilter = yf }

func (acceptTolerance *KeyChains_AcceptTolerance) GetGoName(yname string) string {
    if yname == "duration" { return "Duration" }
    return ""
}

func (acceptTolerance *KeyChains_AcceptTolerance) GetSegmentPath() string {
    return "accept-tolerance"
}

func (acceptTolerance *KeyChains_AcceptTolerance) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (acceptTolerance *KeyChains_AcceptTolerance) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (acceptTolerance *KeyChains_AcceptTolerance) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["duration"] = acceptTolerance.Duration
    return leafs
}

func (acceptTolerance *KeyChains_AcceptTolerance) GetBundleName() string { return "ietf" }

func (acceptTolerance *KeyChains_AcceptTolerance) GetYangName() string { return "accept-tolerance" }

func (acceptTolerance *KeyChains_AcceptTolerance) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (acceptTolerance *KeyChains_AcceptTolerance) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (acceptTolerance *KeyChains_AcceptTolerance) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (acceptTolerance *KeyChains_AcceptTolerance) SetParent(parent types.Entity) { acceptTolerance.parent = parent }

func (acceptTolerance *KeyChains_AcceptTolerance) GetParent() types.Entity { return acceptTolerance.parent }

func (acceptTolerance *KeyChains_AcceptTolerance) GetParentYangName() string { return "key-chains" }

// KeyChains_Key
// One key.
type KeyChains_Key struct {
    parent types.Entity
    YFilter yfilter.YFilter

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

func (key *KeyChains_Key) GetFilter() yfilter.YFilter { return key.YFilter }

func (key *KeyChains_Key) SetFilter(yf yfilter.YFilter) { key.YFilter = yf }

func (key *KeyChains_Key) GetGoName(yname string) string {
    if yname == "key-id" { return "KeyId" }
    if yname == "key-string" { return "KeyString" }
    if yname == "lifetime" { return "Lifetime" }
    if yname == "crypto-algorithm" { return "CryptoAlgorithm" }
    return ""
}

func (key *KeyChains_Key) GetSegmentPath() string {
    return "key" + "[key-id='" + fmt.Sprintf("%v", key.KeyId) + "']"
}

func (key *KeyChains_Key) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "key-string" {
        return &key.KeyString
    }
    if childYangName == "lifetime" {
        return &key.Lifetime
    }
    if childYangName == "crypto-algorithm" {
        return &key.CryptoAlgorithm
    }
    return nil
}

func (key *KeyChains_Key) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["key-string"] = &key.KeyString
    children["lifetime"] = &key.Lifetime
    children["crypto-algorithm"] = &key.CryptoAlgorithm
    return children
}

func (key *KeyChains_Key) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["key-id"] = key.KeyId
    return leafs
}

func (key *KeyChains_Key) GetBundleName() string { return "ietf" }

func (key *KeyChains_Key) GetYangName() string { return "key" }

func (key *KeyChains_Key) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (key *KeyChains_Key) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (key *KeyChains_Key) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (key *KeyChains_Key) SetParent(parent types.Entity) { key.parent = parent }

func (key *KeyChains_Key) GetParent() types.Entity { return key.parent }

func (key *KeyChains_Key) GetParentYangName() string { return "key-chains" }

// KeyChains_Key_KeyString
// The key string.
type KeyChains_Key_KeyString struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Key string in ASCII format. The type is string.
    Keystring interface{}

    // Key in hexadecimal string format. The type is string with pattern:
    // ([0-9a-fA-F]{2}(:[0-9a-fA-F]{2})*)?.
    HexadecimalString interface{}
}

func (keyString *KeyChains_Key_KeyString) GetFilter() yfilter.YFilter { return keyString.YFilter }

func (keyString *KeyChains_Key_KeyString) SetFilter(yf yfilter.YFilter) { keyString.YFilter = yf }

func (keyString *KeyChains_Key_KeyString) GetGoName(yname string) string {
    if yname == "keystring" { return "Keystring" }
    if yname == "hexadecimal-string" { return "HexadecimalString" }
    return ""
}

func (keyString *KeyChains_Key_KeyString) GetSegmentPath() string {
    return "key-string"
}

func (keyString *KeyChains_Key_KeyString) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (keyString *KeyChains_Key_KeyString) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (keyString *KeyChains_Key_KeyString) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["keystring"] = keyString.Keystring
    leafs["hexadecimal-string"] = keyString.HexadecimalString
    return leafs
}

func (keyString *KeyChains_Key_KeyString) GetBundleName() string { return "ietf" }

func (keyString *KeyChains_Key_KeyString) GetYangName() string { return "key-string" }

func (keyString *KeyChains_Key_KeyString) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (keyString *KeyChains_Key_KeyString) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (keyString *KeyChains_Key_KeyString) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (keyString *KeyChains_Key_KeyString) SetParent(parent types.Entity) { keyString.parent = parent }

func (keyString *KeyChains_Key_KeyString) GetParent() types.Entity { return keyString.parent }

func (keyString *KeyChains_Key_KeyString) GetParentYangName() string { return "key" }

// KeyChains_Key_Lifetime
// Specify a key's lifetime.
type KeyChains_Key_Lifetime struct {
    parent types.Entity
    YFilter yfilter.YFilter

    // Single lifetime specification for both send and accept lifetimes.
    SendAcceptLifetime KeyChains_Key_Lifetime_SendAcceptLifetime

    // Separate lifetime specification for send lifetime.
    SendLifetime KeyChains_Key_Lifetime_SendLifetime

    // Separate lifetime specification for accept lifetime.
    AcceptLifetime KeyChains_Key_Lifetime_AcceptLifetime
}

func (lifetime *KeyChains_Key_Lifetime) GetFilter() yfilter.YFilter { return lifetime.YFilter }

func (lifetime *KeyChains_Key_Lifetime) SetFilter(yf yfilter.YFilter) { lifetime.YFilter = yf }

func (lifetime *KeyChains_Key_Lifetime) GetGoName(yname string) string {
    if yname == "send-accept-lifetime" { return "SendAcceptLifetime" }
    if yname == "send-lifetime" { return "SendLifetime" }
    if yname == "accept-lifetime" { return "AcceptLifetime" }
    return ""
}

func (lifetime *KeyChains_Key_Lifetime) GetSegmentPath() string {
    return "lifetime"
}

func (lifetime *KeyChains_Key_Lifetime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    if childYangName == "send-accept-lifetime" {
        return &lifetime.SendAcceptLifetime
    }
    if childYangName == "send-lifetime" {
        return &lifetime.SendLifetime
    }
    if childYangName == "accept-lifetime" {
        return &lifetime.AcceptLifetime
    }
    return nil
}

func (lifetime *KeyChains_Key_Lifetime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    children["send-accept-lifetime"] = &lifetime.SendAcceptLifetime
    children["send-lifetime"] = &lifetime.SendLifetime
    children["accept-lifetime"] = &lifetime.AcceptLifetime
    return children
}

func (lifetime *KeyChains_Key_Lifetime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    return leafs
}

func (lifetime *KeyChains_Key_Lifetime) GetBundleName() string { return "ietf" }

func (lifetime *KeyChains_Key_Lifetime) GetYangName() string { return "lifetime" }

func (lifetime *KeyChains_Key_Lifetime) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (lifetime *KeyChains_Key_Lifetime) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (lifetime *KeyChains_Key_Lifetime) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (lifetime *KeyChains_Key_Lifetime) SetParent(parent types.Entity) { lifetime.parent = parent }

func (lifetime *KeyChains_Key_Lifetime) GetParent() types.Entity { return lifetime.parent }

func (lifetime *KeyChains_Key_Lifetime) GetParentYangName() string { return "key" }

// KeyChains_Key_Lifetime_SendAcceptLifetime
// Single lifetime specification for both send and
// accept lifetimes.
type KeyChains_Key_Lifetime_SendAcceptLifetime struct {
    parent types.Entity
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

func (sendAcceptLifetime *KeyChains_Key_Lifetime_SendAcceptLifetime) GetFilter() yfilter.YFilter { return sendAcceptLifetime.YFilter }

func (sendAcceptLifetime *KeyChains_Key_Lifetime_SendAcceptLifetime) SetFilter(yf yfilter.YFilter) { sendAcceptLifetime.YFilter = yf }

func (sendAcceptLifetime *KeyChains_Key_Lifetime_SendAcceptLifetime) GetGoName(yname string) string {
    if yname == "always" { return "Always" }
    if yname == "start-date-time" { return "StartDateTime" }
    if yname == "no-end-time" { return "NoEndTime" }
    if yname == "duration" { return "Duration" }
    if yname == "end-date-time" { return "EndDateTime" }
    return ""
}

func (sendAcceptLifetime *KeyChains_Key_Lifetime_SendAcceptLifetime) GetSegmentPath() string {
    return "send-accept-lifetime"
}

func (sendAcceptLifetime *KeyChains_Key_Lifetime_SendAcceptLifetime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sendAcceptLifetime *KeyChains_Key_Lifetime_SendAcceptLifetime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sendAcceptLifetime *KeyChains_Key_Lifetime_SendAcceptLifetime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["always"] = sendAcceptLifetime.Always
    leafs["start-date-time"] = sendAcceptLifetime.StartDateTime
    leafs["no-end-time"] = sendAcceptLifetime.NoEndTime
    leafs["duration"] = sendAcceptLifetime.Duration
    leafs["end-date-time"] = sendAcceptLifetime.EndDateTime
    return leafs
}

func (sendAcceptLifetime *KeyChains_Key_Lifetime_SendAcceptLifetime) GetBundleName() string { return "ietf" }

func (sendAcceptLifetime *KeyChains_Key_Lifetime_SendAcceptLifetime) GetYangName() string { return "send-accept-lifetime" }

func (sendAcceptLifetime *KeyChains_Key_Lifetime_SendAcceptLifetime) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (sendAcceptLifetime *KeyChains_Key_Lifetime_SendAcceptLifetime) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (sendAcceptLifetime *KeyChains_Key_Lifetime_SendAcceptLifetime) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (sendAcceptLifetime *KeyChains_Key_Lifetime_SendAcceptLifetime) SetParent(parent types.Entity) { sendAcceptLifetime.parent = parent }

func (sendAcceptLifetime *KeyChains_Key_Lifetime_SendAcceptLifetime) GetParent() types.Entity { return sendAcceptLifetime.parent }

func (sendAcceptLifetime *KeyChains_Key_Lifetime_SendAcceptLifetime) GetParentYangName() string { return "lifetime" }

// KeyChains_Key_Lifetime_SendLifetime
// Separate lifetime specification for send
// lifetime.
type KeyChains_Key_Lifetime_SendLifetime struct {
    parent types.Entity
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

func (sendLifetime *KeyChains_Key_Lifetime_SendLifetime) GetFilter() yfilter.YFilter { return sendLifetime.YFilter }

func (sendLifetime *KeyChains_Key_Lifetime_SendLifetime) SetFilter(yf yfilter.YFilter) { sendLifetime.YFilter = yf }

func (sendLifetime *KeyChains_Key_Lifetime_SendLifetime) GetGoName(yname string) string {
    if yname == "always" { return "Always" }
    if yname == "start-date-time" { return "StartDateTime" }
    if yname == "no-end-time" { return "NoEndTime" }
    if yname == "duration" { return "Duration" }
    if yname == "end-date-time" { return "EndDateTime" }
    return ""
}

func (sendLifetime *KeyChains_Key_Lifetime_SendLifetime) GetSegmentPath() string {
    return "send-lifetime"
}

func (sendLifetime *KeyChains_Key_Lifetime_SendLifetime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (sendLifetime *KeyChains_Key_Lifetime_SendLifetime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (sendLifetime *KeyChains_Key_Lifetime_SendLifetime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["always"] = sendLifetime.Always
    leafs["start-date-time"] = sendLifetime.StartDateTime
    leafs["no-end-time"] = sendLifetime.NoEndTime
    leafs["duration"] = sendLifetime.Duration
    leafs["end-date-time"] = sendLifetime.EndDateTime
    return leafs
}

func (sendLifetime *KeyChains_Key_Lifetime_SendLifetime) GetBundleName() string { return "ietf" }

func (sendLifetime *KeyChains_Key_Lifetime_SendLifetime) GetYangName() string { return "send-lifetime" }

func (sendLifetime *KeyChains_Key_Lifetime_SendLifetime) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (sendLifetime *KeyChains_Key_Lifetime_SendLifetime) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (sendLifetime *KeyChains_Key_Lifetime_SendLifetime) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (sendLifetime *KeyChains_Key_Lifetime_SendLifetime) SetParent(parent types.Entity) { sendLifetime.parent = parent }

func (sendLifetime *KeyChains_Key_Lifetime_SendLifetime) GetParent() types.Entity { return sendLifetime.parent }

func (sendLifetime *KeyChains_Key_Lifetime_SendLifetime) GetParentYangName() string { return "lifetime" }

// KeyChains_Key_Lifetime_AcceptLifetime
// Separate lifetime specification for accept
// lifetime.
type KeyChains_Key_Lifetime_AcceptLifetime struct {
    parent types.Entity
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

func (acceptLifetime *KeyChains_Key_Lifetime_AcceptLifetime) GetFilter() yfilter.YFilter { return acceptLifetime.YFilter }

func (acceptLifetime *KeyChains_Key_Lifetime_AcceptLifetime) SetFilter(yf yfilter.YFilter) { acceptLifetime.YFilter = yf }

func (acceptLifetime *KeyChains_Key_Lifetime_AcceptLifetime) GetGoName(yname string) string {
    if yname == "always" { return "Always" }
    if yname == "start-date-time" { return "StartDateTime" }
    if yname == "no-end-time" { return "NoEndTime" }
    if yname == "duration" { return "Duration" }
    if yname == "end-date-time" { return "EndDateTime" }
    return ""
}

func (acceptLifetime *KeyChains_Key_Lifetime_AcceptLifetime) GetSegmentPath() string {
    return "accept-lifetime"
}

func (acceptLifetime *KeyChains_Key_Lifetime_AcceptLifetime) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (acceptLifetime *KeyChains_Key_Lifetime_AcceptLifetime) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (acceptLifetime *KeyChains_Key_Lifetime_AcceptLifetime) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["always"] = acceptLifetime.Always
    leafs["start-date-time"] = acceptLifetime.StartDateTime
    leafs["no-end-time"] = acceptLifetime.NoEndTime
    leafs["duration"] = acceptLifetime.Duration
    leafs["end-date-time"] = acceptLifetime.EndDateTime
    return leafs
}

func (acceptLifetime *KeyChains_Key_Lifetime_AcceptLifetime) GetBundleName() string { return "ietf" }

func (acceptLifetime *KeyChains_Key_Lifetime_AcceptLifetime) GetYangName() string { return "accept-lifetime" }

func (acceptLifetime *KeyChains_Key_Lifetime_AcceptLifetime) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (acceptLifetime *KeyChains_Key_Lifetime_AcceptLifetime) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (acceptLifetime *KeyChains_Key_Lifetime_AcceptLifetime) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (acceptLifetime *KeyChains_Key_Lifetime_AcceptLifetime) SetParent(parent types.Entity) { acceptLifetime.parent = parent }

func (acceptLifetime *KeyChains_Key_Lifetime_AcceptLifetime) GetParent() types.Entity { return acceptLifetime.parent }

func (acceptLifetime *KeyChains_Key_Lifetime_AcceptLifetime) GetParentYangName() string { return "lifetime" }

// KeyChains_Key_CryptoAlgorithm
// Cryptographic algorithm associated with key.
type KeyChains_Key_CryptoAlgorithm struct {
    parent types.Entity
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

func (cryptoAlgorithm *KeyChains_Key_CryptoAlgorithm) GetFilter() yfilter.YFilter { return cryptoAlgorithm.YFilter }

func (cryptoAlgorithm *KeyChains_Key_CryptoAlgorithm) SetFilter(yf yfilter.YFilter) { cryptoAlgorithm.YFilter = yf }

func (cryptoAlgorithm *KeyChains_Key_CryptoAlgorithm) GetGoName(yname string) string {
    if yname == "hmac-sha1-12" { return "HmacSha112" }
    if yname == "hmac-sha1-20" { return "HmacSha120" }
    if yname == "md5" { return "Md5" }
    if yname == "sha-1" { return "Sha1" }
    if yname == "hmac-sha-1" { return "HmacSha1" }
    if yname == "hmac-sha-256" { return "HmacSha256" }
    if yname == "hmac-sha-384" { return "HmacSha384" }
    if yname == "hmac-sha-512" { return "HmacSha512" }
    return ""
}

func (cryptoAlgorithm *KeyChains_Key_CryptoAlgorithm) GetSegmentPath() string {
    return "crypto-algorithm"
}

func (cryptoAlgorithm *KeyChains_Key_CryptoAlgorithm) GetChildByName(childYangName string, segmentPath string) types.Entity {
    return nil
}

func (cryptoAlgorithm *KeyChains_Key_CryptoAlgorithm) GetChildren() map[string]types.Entity {
    children := make(map[string]types.Entity)
    return children
}

func (cryptoAlgorithm *KeyChains_Key_CryptoAlgorithm) GetLeafs() map[string]interface{} {
    leafs := make(map[string]interface{})
    leafs["hmac-sha1-12"] = cryptoAlgorithm.HmacSha112
    leafs["hmac-sha1-20"] = cryptoAlgorithm.HmacSha120
    leafs["md5"] = cryptoAlgorithm.Md5
    leafs["sha-1"] = cryptoAlgorithm.Sha1
    leafs["hmac-sha-1"] = cryptoAlgorithm.HmacSha1
    leafs["hmac-sha-256"] = cryptoAlgorithm.HmacSha256
    leafs["hmac-sha-384"] = cryptoAlgorithm.HmacSha384
    leafs["hmac-sha-512"] = cryptoAlgorithm.HmacSha512
    return leafs
}

func (cryptoAlgorithm *KeyChains_Key_CryptoAlgorithm) GetBundleName() string { return "ietf" }

func (cryptoAlgorithm *KeyChains_Key_CryptoAlgorithm) GetYangName() string { return "crypto-algorithm" }

func (cryptoAlgorithm *KeyChains_Key_CryptoAlgorithm) GetBundleYangModelsLocation() string { return ietf.GetModelsPath() }

func (cryptoAlgorithm *KeyChains_Key_CryptoAlgorithm) GetCapabilitiesTable() map[string]string {
    return ietf.GetCapabilities() }

func (cryptoAlgorithm *KeyChains_Key_CryptoAlgorithm) GetNamespaceTable() map[string]string {
    return ietf.GetNamespaces() }

func (cryptoAlgorithm *KeyChains_Key_CryptoAlgorithm) SetParent(parent types.Entity) { cryptoAlgorithm.parent = parent }

func (cryptoAlgorithm *KeyChains_Key_CryptoAlgorithm) GetParent() types.Entity { return cryptoAlgorithm.parent }

func (cryptoAlgorithm *KeyChains_Key_CryptoAlgorithm) GetParentYangName() string { return "key" }

