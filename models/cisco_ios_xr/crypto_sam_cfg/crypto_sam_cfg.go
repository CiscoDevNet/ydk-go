// This module contains a collection of YANG definitions
// for Cisco IOS-XR crypto-sam package configuration.
// 
// This module contains definitions
// for the following management objects:
//   sam: Software Authentication Manager (SAM) Config
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package crypto_sam_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package crypto_sam_cfg"))
    ydk.RegisterEntity("{http://cisco.com/ns/yang/Cisco-IOS-XR-crypto-sam-cfg sam}", reflect.TypeOf(Sam{}))
    ydk.RegisterEntity("Cisco-IOS-XR-crypto-sam-cfg:sam", reflect.TypeOf(Sam{}))
}

// CryptoSamAction represents Crypto sam action
type CryptoSamAction string

const (
    // To respond YES to the SAM prompt
    CryptoSamAction_proceed CryptoSamAction = "proceed"

    // To respond NO to the SAM prompt
    CryptoSamAction_terminate CryptoSamAction = "terminate"
)

// Sam
// Software Authentication Manager (SAM) Config
type Sam struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Set prompt interval at reboot time.
    PromptInterval Sam_PromptInterval
}

func (sam *Sam) GetEntityData() *types.CommonEntityData {
    sam.EntityData.YFilter = sam.YFilter
    sam.EntityData.YangName = "sam"
    sam.EntityData.BundleName = "cisco_ios_xr"
    sam.EntityData.ParentYangName = "Cisco-IOS-XR-crypto-sam-cfg"
    sam.EntityData.SegmentPath = "Cisco-IOS-XR-crypto-sam-cfg:sam"
    sam.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    sam.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    sam.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    sam.EntityData.Children = types.NewOrderedMap()
    sam.EntityData.Children.Append("prompt-interval", types.YChild{"PromptInterval", &sam.PromptInterval})
    sam.EntityData.Leafs = types.NewOrderedMap()

    sam.EntityData.YListKeys = []string {}

    return &(sam.EntityData)
}

// Sam_PromptInterval
// Set prompt interval at reboot time
// This type is a presence type.
type Sam_PromptInterval struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YPresence bool

    // Respond to SAM prompt either Proceed/Terminate. The type is
    // CryptoSamAction. This attribute is mandatory.
    Action interface{}

    // Prompt time from 0 - 300 seconds. The type is interface{} with range:
    // 0..300. This attribute is mandatory. Units are second.
    PromptTime interface{}
}

func (promptInterval *Sam_PromptInterval) GetEntityData() *types.CommonEntityData {
    promptInterval.EntityData.YFilter = promptInterval.YFilter
    promptInterval.EntityData.YangName = "prompt-interval"
    promptInterval.EntityData.BundleName = "cisco_ios_xr"
    promptInterval.EntityData.ParentYangName = "sam"
    promptInterval.EntityData.SegmentPath = "prompt-interval"
    promptInterval.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    promptInterval.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    promptInterval.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    promptInterval.EntityData.Children = types.NewOrderedMap()
    promptInterval.EntityData.Leafs = types.NewOrderedMap()
    promptInterval.EntityData.Leafs.Append("action", types.YLeaf{"Action", promptInterval.Action})
    promptInterval.EntityData.Leafs.Append("prompt-time", types.YLeaf{"PromptTime", promptInterval.PromptTime})

    promptInterval.EntityData.YListKeys = []string {}

    return &(promptInterval.EntityData)
}

