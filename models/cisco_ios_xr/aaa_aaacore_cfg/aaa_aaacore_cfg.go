// This module contains a collection of YANG definitions
// for Cisco IOS-XR aaa-aaacore package configuration.
// 
// This YANG module augments the
//   Cisco-IOS-XR-aaa-lib-cfg,
//   Cisco-IOS-XR-ifmgr-cfg
// modules with configuration data.
// 
// Copyright (c) 2013-2018 by Cisco Systems, Inc.
// All rights reserved.
package aaa_aaacore_cfg

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package aaa_aaacore_cfg"))
}

// NasPortValue represents Nas port value
type NasPortValue string

const (
    // Async(0)
    NasPortValue_async NasPortValue = "async"

    // Sync(1)
    NasPortValue_sync NasPortValue = "sync"

    // ISDN(2)
    NasPortValue_isdn NasPortValue = "isdn"

    // ISDN Async V120(3)
    NasPortValue_isdn_async_v120 NasPortValue = "isdn-async-v120"

    // ISDN Async V110(4)
    NasPortValue_isdn_async_v110 NasPortValue = "isdn-async-v110"

    // Virtual(5)
    NasPortValue_virtual NasPortValue = "virtual"

    // ISDN Async PIAFS(6)
    NasPortValue_isdn_async_piafs NasPortValue = "isdn-async-piafs"

    // X75(9)
    NasPortValue_x75 NasPortValue = "x75"

    // Ethernet(15)
    NasPortValue_ethernet NasPortValue = "ethernet"

    // PPPoA(30)
    NasPortValue_pppoa NasPortValue = "pppoa"

    // PPPoEoA(31)
    NasPortValue_pppoeoa NasPortValue = "pppoeoa"

    // PPPoEoE(32)
    NasPortValue_pppoeoe NasPortValue = "pppoeoe"

    // PPPoEoVLAN(33)
    NasPortValue_pppoeovlan NasPortValue = "pppoeovlan"

    // PPPoEoQinQ(34)
    NasPortValue_pppoeoqinq NasPortValue = "pppoeoqinq"

    // Virtual PPPoEoE(35)
    NasPortValue_virtual_pppoeoe NasPortValue = "virtual-pppoeoe"

    // Virtual PPPoEoVLAN(36)
    NasPortValue_virtual_pppoeovlan NasPortValue = "virtual-pppoeovlan"

    // Virtual PPPoEoQinQ(37)
    NasPortValue_virtual_pppoeoqinaq NasPortValue = "virtual-pppoeoqinaq"

    // IPSEC(38)
    NasPortValue_ipsec NasPortValue = "ipsec"

    // IPOEOE(39)
    NasPortValue_ipoeoe NasPortValue = "ipoeoe"

    // IPOEOVLAN(40)
    NasPortValue_ipoeovlan NasPortValue = "ipoeovlan"

    // IPOEOQINQ(41)
    NasPortValue_ipoeoqinq NasPortValue = "ipoeoqinq"

    // VIRTUAL IPOEOE(42)
    NasPortValue_virtual_ipoeoe NasPortValue = "virtual-ipoeoe"

    // VIRTUAL IPOEOVLAN(43)
    NasPortValue_virtual_ipoeovlan NasPortValue = "virtual-ipoeovlan"

    // VIRTUAL IPOEOQINQ(44)
    NasPortValue_virtual_ipoeoqinq NasPortValue = "virtual-ipoeoqinq"
)

// AaaServiceAccounting represents Aaa service accounting
type AaaServiceAccounting string

const (
    // None
    AaaServiceAccounting_none AaaServiceAccounting = "none"

    // Extended
    AaaServiceAccounting_extended AaaServiceAccounting = "extended"

    // Brief
    AaaServiceAccounting_brief AaaServiceAccounting = "brief"
)

