package snmpv2_tc

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package snmpv2_tc"))
}

// TruthValue
type TruthValue string

const (
    TruthValue_true_ TruthValue = "true"

    TruthValue_false_ TruthValue = "false"
)

// RowStatus
type RowStatus string

const (
    RowStatus_active RowStatus = "active"

    RowStatus_notInService RowStatus = "notInService"

    RowStatus_notReady RowStatus = "notReady"

    RowStatus_createAndGo RowStatus = "createAndGo"

    RowStatus_createAndWait RowStatus = "createAndWait"

    RowStatus_destroy RowStatus = "destroy"
)

// StorageType
type StorageType string

const (
    StorageType_other StorageType = "other"

    StorageType_volatile StorageType = "volatile"

    StorageType_nonVolatile StorageType = "nonVolatile"

    StorageType_permanent StorageType = "permanent"

    StorageType_readOnly StorageType = "readOnly"
)

