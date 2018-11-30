// This MIB Module provides Textual Conventions 
// and OBJECT-IDENTITY Objects to be used PW services.
package cisco_ietf_pw_tc_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_ietf_pw_tc_mib"))
}

// CpwOperStatus represents                   in OperStatus 'up'.  
type CpwOperStatus string

const (
    CpwOperStatus_up CpwOperStatus = "up"

    CpwOperStatus_down CpwOperStatus = "down"

    CpwOperStatus_testing CpwOperStatus = "testing"

    CpwOperStatus_unknown CpwOperStatus = "unknown"

    CpwOperStatus_dormant CpwOperStatus = "dormant"

    CpwOperStatus_notPresent CpwOperStatus = "notPresent"

    CpwOperStatus_lowerLayerDown CpwOperStatus = "lowerLayerDown"
)

// CpwVcType represents out by the WG. 
type CpwVcType string

const (
    CpwVcType_other CpwVcType = "other"

    CpwVcType_frameRelay CpwVcType = "frameRelay"

    CpwVcType_atmAal5Vcc CpwVcType = "atmAal5Vcc"

    CpwVcType_atmTransparent CpwVcType = "atmTransparent"

    CpwVcType_ethernetVLAN CpwVcType = "ethernetVLAN"

    CpwVcType_ethernet CpwVcType = "ethernet"

    CpwVcType_hdlc CpwVcType = "hdlc"

    CpwVcType_ppp CpwVcType = "ppp"

    CpwVcType_cep CpwVcType = "cep"

    CpwVcType_atmVccCell CpwVcType = "atmVccCell"

    CpwVcType_atmVpcCell CpwVcType = "atmVpcCell"

    CpwVcType_ethernetVPLS CpwVcType = "ethernetVPLS"

    CpwVcType_e1Satop CpwVcType = "e1Satop"

    CpwVcType_t1Satop CpwVcType = "t1Satop"

    CpwVcType_e3Satop CpwVcType = "e3Satop"

    CpwVcType_t3Satop CpwVcType = "t3Satop"

    CpwVcType_basicCesPsn CpwVcType = "basicCesPsn"

    CpwVcType_basicTdmIp CpwVcType = "basicTdmIp"

    CpwVcType_tdmCasCesPsn CpwVcType = "tdmCasCesPsn"

    CpwVcType_tdmCasTdmIp CpwVcType = "tdmCasTdmIp"
)

