// Execution and monitoring of sysadmin ISSU operations 
package sysadmin_issu

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xr"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package sysadmin_issu"))
    ydk.RegisterEntity("{http://cisco.com/calvados/Cisco-IOS-XR-sysadmin-issu issu}", reflect.TypeOf(Issu{}))
    ydk.RegisterEntity("Cisco-IOS-XR-sysadmin-issu:issu", reflect.TypeOf(Issu{}))
}

// OpStartResult represents Enumeration of errors that can be encountered while attempting to begin an ISSU operation
type OpStartResult string

const (
    // Operation was started successfully
    OpStartResult_start_success OpStartResult = "start-success"

    // Another ISSU operation is already in progress
    OpStartResult_error_operation_in_progress OpStartResult = "error-operation-in-progress"

    // A request to activate the prepared software was made, but there is no successfully prepared software
    OpStartResult_activate_error_no_prepare OpStartResult = "activate-error-no-prepare"

    // A request to prepare software was made, but previously prepared software exists
    OpStartResult_prepare_error_previous_prepare OpStartResult = "prepare-error-previous-prepare"

    // The system is in a state that makes in-service recovery impossible
    OpStartResult_recover_error_unrecoverable OpStartResult = "recover-error-unrecoverable"

    // An internal error occured while attempting to start the operation
    OpStartResult_start_error_internal OpStartResult = "start-error-internal"
)

// OpResult represents Enumeration of errors that can be encountered during an ISSU operation
type OpResult string

const (
    // Operation succeeded
    OpResult_success OpResult = "success"

    // Part or all of the input was invalid
    OpResult_error_input OpResult = "error-input"

    // An internal error occurred during ISSU orchestration
    OpResult_error_orchestration OpResult = "error-orchestration"

    // An error occured in the install infrastructure
    OpResult_error_install OpResult = "error-install"

    // Not all nodes in the system have the required redundancy to allow an ISSU to proceed
    OpResult_error_node_redundancy OpResult = "error-node-redundancy"
)

// IssuNotif represents Enumeration of notifications that features can be registered for
type IssuNotif string

const (
    // A sysadmin ISSU operation is beginning
    IssuNotif_notif_sysadmin_op_start IssuNotif = "notif-sysadmin-op-start"

    // A sysadmin ISSU phase is beginning
    IssuNotif_notif_sysadmin_phase_start IssuNotif = "notif-sysadmin-phase-start"

    // A sysadmin ISSU operation is ending
    IssuNotif_notif_sysadmin_op_end IssuNotif = "notif-sysadmin-op-end"
)

// OpStage represents Enumeration of stages that the ISSU operation can be in
type OpStage string

const (
    // Prechecks and preprocessing
    OpStage_prepare_preamble OpStage = "prepare-preamble"

    // Host package preparation
    OpStage_prepare_host OpStage = "prepare-host"

    // Sysadmin package preparation
    OpStage_prepare_sysadmin OpStage = "prepare-sysadmin"

    // Aborting after encountering an error
    OpStage_prepare_aborting OpStage = "prepare-aborting"

    // Operation post-processing
    OpStage_prepare_postamble OpStage = "prepare-postamble"

    // All preparation complete
    OpStage_prepare_complete OpStage = "prepare-complete"

    // Pre-checks and pre-processing
    OpStage_activate_preamble OpStage = "activate-preamble"

    // Sysadmin activation on first subset of nodes
    OpStage_activate_sysadmin_phase_one OpStage = "activate-sysadmin-phase-one"

    // Sysadmin activation on second subset of nodes
    OpStage_activate_sysadmin_phase_two OpStage = "activate-sysadmin-phase-two"

    // Host activation
    OpStage_activate_host OpStage = "activate-host"

    // Aborting after encountering an error, from which it is not possible to recover
    OpStage_activate_aborting OpStage = "activate-aborting"

    // Paused after encountering an error from which it is possible to recover
    OpStage_activate_paused OpStage = "activate-paused"

    // Post-checks and post-processing
    OpStage_activate_postamble OpStage = "activate-postamble"

    // All activation complete
    OpStage_activate_complete OpStage = "activate-complete"
)

// Issu
// ISSU actions and operational state
type Issu struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Status of the in-progress or last completed ISSU operation.
    Status Issu_Status

    // Features registered for notifications of ISSU phases.
    Clients Issu_Clients

    // Internal infrastructure state.
    Internals Issu_Internals
}

func (issu *Issu) GetEntityData() *types.CommonEntityData {
    issu.EntityData.YFilter = issu.YFilter
    issu.EntityData.YangName = "issu"
    issu.EntityData.BundleName = "cisco_ios_xr"
    issu.EntityData.ParentYangName = "Cisco-IOS-XR-sysadmin-issu"
    issu.EntityData.SegmentPath = "Cisco-IOS-XR-sysadmin-issu:issu"
    issu.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    issu.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    issu.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    issu.EntityData.Children = make(map[string]types.YChild)
    issu.EntityData.Children["status"] = types.YChild{"Status", &issu.Status}
    issu.EntityData.Children["clients"] = types.YChild{"Clients", &issu.Clients}
    issu.EntityData.Children["internals"] = types.YChild{"Internals", &issu.Internals}
    issu.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(issu.EntityData)
}

// Issu_Status
// Status of the in-progress or last completed ISSU operation
type Issu_Status struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Whether the operation is an activate or deactivate. Both types of operation
    // go through 'prepare' and 'activate' phases.The difference is the end
    // result: whether the target packages are made to run, or removed from the
    // running software. The type is OperationType.
    OperationType interface{}

    // ID for the current/latest phase of the operation. The type is interface{}
    // with range: 0..4294967295.
    Id interface{}

    // ID for the prepare phase of the operation. The type is interface{} with
    // range: 0..4294967295.
    PrepareId interface{}

    // ID for the activate or deactivate phase of the operation. The type is
    // interface{} with range: 0..4294967295.
    ActivateId interface{}

    // Sysadmin packages that are part of the operation. The type is slice of
    // string.
    SysadminPackages []interface{}

    // Host OS packages that are part of the operation. The type is slice of
    // string.
    HostPackages []interface{}

    // Whether or not the operation has completed. The type is bool.
    Complete interface{}

    // Whether the operation succeeded or failed, and the error if any. If in
    // progress, this reflects the current state. The type is OpResult.
    Result interface{}

    // If a recovery attempt has been completed, an indication of whether the
    // recovery succeeded or failed, and the error if any. The type is OpResult.
    RecoverResult interface{}

    // State specific to the prepare phase.
    Prepare Issu_Status_Prepare

    // State specific to the activate phase.
    Activate Issu_Status_Activate

    // Details of the first error that was encountered, if there were any.
    Error Issu_Status_Error
}

func (status *Issu_Status) GetEntityData() *types.CommonEntityData {
    status.EntityData.YFilter = status.YFilter
    status.EntityData.YangName = "status"
    status.EntityData.BundleName = "cisco_ios_xr"
    status.EntityData.ParentYangName = "issu"
    status.EntityData.SegmentPath = "status"
    status.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    status.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    status.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    status.EntityData.Children = make(map[string]types.YChild)
    status.EntityData.Children["prepare"] = types.YChild{"Prepare", &status.Prepare}
    status.EntityData.Children["activate"] = types.YChild{"Activate", &status.Activate}
    status.EntityData.Children["error"] = types.YChild{"Error", &status.Error}
    status.EntityData.Leafs = make(map[string]types.YLeaf)
    status.EntityData.Leafs["operation-type"] = types.YLeaf{"OperationType", status.OperationType}
    status.EntityData.Leafs["id"] = types.YLeaf{"Id", status.Id}
    status.EntityData.Leafs["prepare-id"] = types.YLeaf{"PrepareId", status.PrepareId}
    status.EntityData.Leafs["activate-id"] = types.YLeaf{"ActivateId", status.ActivateId}
    status.EntityData.Leafs["sysadmin-packages"] = types.YLeaf{"SysadminPackages", status.SysadminPackages}
    status.EntityData.Leafs["host-packages"] = types.YLeaf{"HostPackages", status.HostPackages}
    status.EntityData.Leafs["complete"] = types.YLeaf{"Complete", status.Complete}
    status.EntityData.Leafs["result"] = types.YLeaf{"Result", status.Result}
    status.EntityData.Leafs["recover-result"] = types.YLeaf{"RecoverResult", status.RecoverResult}
    return &(status.EntityData)
}

// Issu_Status_Prepare
// State specific to the prepare phase
type Issu_Status_Prepare struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Progress of the prepare phase. The type is OpStage.
    Stage interface{}

    // When this ehase was started. The type is string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    StartTime interface{}

    // A description of the current orchestration activity being executed. The
    // type is string.
    Activity interface{}

    // When the current activity was started. The type is string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    ActivityStartTime interface{}
}

func (prepare *Issu_Status_Prepare) GetEntityData() *types.CommonEntityData {
    prepare.EntityData.YFilter = prepare.YFilter
    prepare.EntityData.YangName = "prepare"
    prepare.EntityData.BundleName = "cisco_ios_xr"
    prepare.EntityData.ParentYangName = "status"
    prepare.EntityData.SegmentPath = "prepare"
    prepare.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prepare.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prepare.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prepare.EntityData.Children = make(map[string]types.YChild)
    prepare.EntityData.Leafs = make(map[string]types.YLeaf)
    prepare.EntityData.Leafs["stage"] = types.YLeaf{"Stage", prepare.Stage}
    prepare.EntityData.Leafs["start-time"] = types.YLeaf{"StartTime", prepare.StartTime}
    prepare.EntityData.Leafs["activity"] = types.YLeaf{"Activity", prepare.Activity}
    prepare.EntityData.Leafs["activity-start-time"] = types.YLeaf{"ActivityStartTime", prepare.ActivityStartTime}
    return &(prepare.EntityData)
}

// Issu_Status_Activate
// State specific to the activate phase
type Issu_Status_Activate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Progress of the activate phase. The type is OpStage.
    Stage interface{}

    // When this phase was started. The type is string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    StartTime interface{}

    // A description of the current orchestration activity being executed. The
    // type is string.
    Activity interface{}

    // Nodes on which the current orchestration activity is being executed. The
    // type is slice of string.
    ActivityNodes []interface{}

    // A description of what needs to happen before the next orchestration
    // activity can be started. The type is string.
    ActivityWaitingFor interface{}

    // When the current activity was started. The type is string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    ActivityStartTime interface{}
}

func (activate *Issu_Status_Activate) GetEntityData() *types.CommonEntityData {
    activate.EntityData.YFilter = activate.YFilter
    activate.EntityData.YangName = "activate"
    activate.EntityData.BundleName = "cisco_ios_xr"
    activate.EntityData.ParentYangName = "status"
    activate.EntityData.SegmentPath = "activate"
    activate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    activate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    activate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    activate.EntityData.Children = make(map[string]types.YChild)
    activate.EntityData.Leafs = make(map[string]types.YLeaf)
    activate.EntityData.Leafs["stage"] = types.YLeaf{"Stage", activate.Stage}
    activate.EntityData.Leafs["start-time"] = types.YLeaf{"StartTime", activate.StartTime}
    activate.EntityData.Leafs["activity"] = types.YLeaf{"Activity", activate.Activity}
    activate.EntityData.Leafs["activity-nodes"] = types.YLeaf{"ActivityNodes", activate.ActivityNodes}
    activate.EntityData.Leafs["activity-waiting-for"] = types.YLeaf{"ActivityWaitingFor", activate.ActivityWaitingFor}
    activate.EntityData.Leafs["activity-start-time"] = types.YLeaf{"ActivityStartTime", activate.ActivityStartTime}
    return &(activate.EntityData)
}

// Issu_Status_Error
// Details of the first error that was encountered, if there were any.
type Issu_Status_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // If the operation has completed, an indication of whether the operation
    // succeeded or failed, and the error if any. The type is OpResult.
    Result interface{}

    // The stage during which the error was encountered. The type is OpStage.
    Stage interface{}

    // Message describing the error. The type is string.
    ErrorMessage interface{}

    // Details specific to the error. Contents are only filled in if it is
    // relevant to the error that occured.
    Details Issu_Status_Error_Details
}

func (error *Issu_Status_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "status"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Children["details"] = types.YChild{"Details", &error.Details}
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["result"] = types.YLeaf{"Result", error.Result}
    error.EntityData.Leafs["stage"] = types.YLeaf{"Stage", error.Stage}
    error.EntityData.Leafs["error-message"] = types.YLeaf{"ErrorMessage", error.ErrorMessage}
    return &(error.EntityData)
}

// Issu_Status_Error_Details
// Details specific to the error. Contents are only filled in if it is relevant to the error that occured.
type Issu_Status_Error_Details struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // A list of the nodes affected by or causing the error. The type is slice of
    // string.
    Nodes []interface{}

    // A list of the registered features affected by or causing the error. The
    // type is slice of string.
    Clients []interface{}

    // A list of the packages affected by or causing the error. The type is slice
    // of string.
    Packages []interface{}

    // A list of the operation IDs affected by or causing the error. The type is
    // slice of interface{} with range: 0..4294967295.
    OperationIds []interface{}
}

func (details *Issu_Status_Error_Details) GetEntityData() *types.CommonEntityData {
    details.EntityData.YFilter = details.YFilter
    details.EntityData.YangName = "details"
    details.EntityData.BundleName = "cisco_ios_xr"
    details.EntityData.ParentYangName = "error"
    details.EntityData.SegmentPath = "details"
    details.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    details.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    details.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    details.EntityData.Children = make(map[string]types.YChild)
    details.EntityData.Leafs = make(map[string]types.YLeaf)
    details.EntityData.Leafs["nodes"] = types.YLeaf{"Nodes", details.Nodes}
    details.EntityData.Leafs["clients"] = types.YLeaf{"Clients", details.Clients}
    details.EntityData.Leafs["packages"] = types.YLeaf{"Packages", details.Packages}
    details.EntityData.Leafs["operation-ids"] = types.YLeaf{"OperationIds", details.OperationIds}
    return &(details.EntityData)
}

// Issu_Status_OperationType represents Whether the operation is an activate or deactivate. Both types of operation go through 'prepare' and 'activate' phases.The difference is the end result: whether the target packages are made to run, or removed from the running software.
type Issu_Status_OperationType string

const (
    // No ISSU operations have been attempted.
    Issu_Status_OperationType_no_operation Issu_Status_OperationType = "no-operation"

    // Overall operation will add or upgrade packages in the running software
    Issu_Status_OperationType_activate_operation Issu_Status_OperationType = "activate-operation"

    // Overall operation will remove packages from the running software
    Issu_Status_OperationType_deactivate_operation Issu_Status_OperationType = "deactivate-operation"
)

// Issu_Clients
// Features registered for notifications of ISSU phases
type Issu_Clients struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Type of most recent notification sent to clients. The type is IssuNotif.
    Announcement interface{}

    // Status of most recent notification sent to clients. The type is
    // AnnouncementStatus.
    AnnouncementStatus interface{}

    // The type is slice of Issu_Clients_Client.
    Client []Issu_Clients_Client
}

func (clients *Issu_Clients) GetEntityData() *types.CommonEntityData {
    clients.EntityData.YFilter = clients.YFilter
    clients.EntityData.YangName = "clients"
    clients.EntityData.BundleName = "cisco_ios_xr"
    clients.EntityData.ParentYangName = "issu"
    clients.EntityData.SegmentPath = "clients"
    clients.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    clients.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    clients.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    clients.EntityData.Children = make(map[string]types.YChild)
    clients.EntityData.Children["client"] = types.YChild{"Client", nil}
    for i := range clients.Client {
        clients.EntityData.Children[types.GetSegmentPath(&clients.Client[i])] = types.YChild{"Client", &clients.Client[i]}
    }
    clients.EntityData.Leafs = make(map[string]types.YLeaf)
    clients.EntityData.Leafs["announcement"] = types.YLeaf{"Announcement", clients.Announcement}
    clients.EntityData.Leafs["announcement-status"] = types.YLeaf{"AnnouncementStatus", clients.AnnouncementStatus}
    return &(clients.EntityData)
}

// Issu_Clients_Client
type Issu_Clients_Client struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Name of the registered feature. The type is
    // string.
    Name interface{}

    // This attribute is a key. Node on which the feature process is running. The
    // type is string.
    Location interface{}

    // Which notifications the feature is registered to receive. The type is slice
    // of IssuNotif.
    RegisteredFor []interface{}

    // Type of most recent notification. The type is IssuNotif.
    Notif interface{}

    // Response from this client to most recent notification sent to the client.
    // The type is Response.
    Response interface{}

    // Whether the feature has requested that the current operation be aborted.
    // The type is bool.
    Aborted interface{}

    // Description of the reason for requesting an abort, if applicable. The type
    // is string.
    AbortReason interface{}
}

func (client *Issu_Clients_Client) GetEntityData() *types.CommonEntityData {
    client.EntityData.YFilter = client.YFilter
    client.EntityData.YangName = "client"
    client.EntityData.BundleName = "cisco_ios_xr"
    client.EntityData.ParentYangName = "clients"
    client.EntityData.SegmentPath = "client" + "[name='" + fmt.Sprintf("%v", client.Name) + "']" + "[location='" + fmt.Sprintf("%v", client.Location) + "']"
    client.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    client.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    client.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    client.EntityData.Children = make(map[string]types.YChild)
    client.EntityData.Leafs = make(map[string]types.YLeaf)
    client.EntityData.Leafs["name"] = types.YLeaf{"Name", client.Name}
    client.EntityData.Leafs["location"] = types.YLeaf{"Location", client.Location}
    client.EntityData.Leafs["registered-for"] = types.YLeaf{"RegisteredFor", client.RegisteredFor}
    client.EntityData.Leafs["notif"] = types.YLeaf{"Notif", client.Notif}
    client.EntityData.Leafs["response"] = types.YLeaf{"Response", client.Response}
    client.EntityData.Leafs["aborted"] = types.YLeaf{"Aborted", client.Aborted}
    client.EntityData.Leafs["abort-reason"] = types.YLeaf{"AbortReason", client.AbortReason}
    return &(client.EntityData)
}

// Issu_Clients_Client_Response represents Response from this client to most recent notification sent to the client
type Issu_Clients_Client_Response string

const (
    // No notification has yet been sent
    Issu_Clients_Client_Response_notif_resp_no_notif Issu_Clients_Client_Response = "notif-resp-no-notif"

    // No response has yet been sent
    Issu_Clients_Client_Response_notif_resp_pending Issu_Clients_Client_Response = "notif-resp-pending"

    // The notification has been acknowledged, ISSU may continue
    Issu_Clients_Client_Response_notif_resp_ack Issu_Clients_Client_Response = "notif-resp-ack"

    // The feature has vetoed the ISSU operation
    Issu_Clients_Client_Response_notif_resp_veto Issu_Clients_Client_Response = "notif-resp-veto"

    // The feature has disconnected during the notification
    Issu_Clients_Client_Response_notif_resp_disconnect Issu_Clients_Client_Response = "notif-resp-disconnect"

    // The feature has timed out during the ISSU operation
    Issu_Clients_Client_Response_notif_resp_timeout Issu_Clients_Client_Response = "notif-resp-timeout"

    // There was an error ending the announcement to the feature
    Issu_Clients_Client_Response_notif_resp_send_error Issu_Clients_Client_Response = "notif-resp-send-error"

    // The feature has returned an error during the ISSU operation
    Issu_Clients_Client_Response_notif_resp_client_error Issu_Clients_Client_Response = "notif-resp-client-error"

    // The feature has aborted during the ISSU operation
    Issu_Clients_Client_Response_notif_resp_client_abort Issu_Clients_Client_Response = "notif-resp-client-abort"
)

// Issu_Clients_AnnouncementStatus represents Status of most recent notification sent to clients
type Issu_Clients_AnnouncementStatus string

const (
    // No announcement has yet been sent
    Issu_Clients_AnnouncementStatus_announce_no_notif Issu_Clients_AnnouncementStatus = "announce-no-notif"

    // Announcement is in progress, waiting for responses
    Issu_Clients_AnnouncementStatus_announce_in_progress Issu_Clients_AnnouncementStatus = "announce-in-progress"

    // The announcement has been acknowledged by all clients, ISSU may continue
    Issu_Clients_AnnouncementStatus_announce_success Issu_Clients_AnnouncementStatus = "announce-success"

    // One or more features vetoed the ISSU operation
    Issu_Clients_AnnouncementStatus_announce_veto Issu_Clients_AnnouncementStatus = "announce-veto"

    // One or more features disconnected during the announcement
    Issu_Clients_AnnouncementStatus_announce_disconnect Issu_Clients_AnnouncementStatus = "announce-disconnect"

    // One or more features timed out during the announcement
    Issu_Clients_AnnouncementStatus_announce_timeout Issu_Clients_AnnouncementStatus = "announce-timeout"

    // There was an error ending the announcement to one or more features
    Issu_Clients_AnnouncementStatus_announce_send_error Issu_Clients_AnnouncementStatus = "announce-send-error"

    // One or more features has returned an error during the ISSU operation
    Issu_Clients_AnnouncementStatus_announce_client_error Issu_Clients_AnnouncementStatus = "announce-client-error"
)

// Issu_Internals
// Internal infrastructure state
type Issu_Internals struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Orchestrator module internal state.
    Orchestrator Issu_Internals_Orchestrator

    // Agent module internal state.
    Agents Issu_Internals_Agents

    // Inventory monitor module internal state.
    InventoryMonitor Issu_Internals_InventoryMonitor
}

func (internals *Issu_Internals) GetEntityData() *types.CommonEntityData {
    internals.EntityData.YFilter = internals.YFilter
    internals.EntityData.YangName = "internals"
    internals.EntityData.BundleName = "cisco_ios_xr"
    internals.EntityData.ParentYangName = "issu"
    internals.EntityData.SegmentPath = "internals"
    internals.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    internals.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    internals.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    internals.EntityData.Children = make(map[string]types.YChild)
    internals.EntityData.Children["orchestrator"] = types.YChild{"Orchestrator", &internals.Orchestrator}
    internals.EntityData.Children["agents"] = types.YChild{"Agents", &internals.Agents}
    internals.EntityData.Children["inventory-monitor"] = types.YChild{"InventoryMonitor", &internals.InventoryMonitor}
    internals.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(internals.EntityData)
}

// Issu_Internals_Orchestrator
// Orchestrator module internal state
type Issu_Internals_Orchestrator struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Command issued. The type is OpRequestType.
    Command interface{}

    // Type of operation: activate or deactivate. The type is OpRequestType.
    OperationType interface{}

    // Operation being performed. The type is OpRequestType.
    CurrentOperation interface{}

    // True if the overall ISSU operation has been completed, false otherwise. The
    // type is bool.
    IssuCompleted interface{}

    // The ID for the operation. The type is interface{} with range:
    // 0..4294967295.
    OperationId interface{}

    // True if an operation is in progress. The type is bool.
    InProgress interface{}

    
    OperationStartDetails Issu_Internals_Orchestrator_OperationStartDetails

    
    InternalPrepare Issu_Internals_Orchestrator_InternalPrepare

    
    InternalActivate Issu_Internals_Orchestrator_InternalActivate

    
    Error Issu_Internals_Orchestrator_Error
}

func (orchestrator *Issu_Internals_Orchestrator) GetEntityData() *types.CommonEntityData {
    orchestrator.EntityData.YFilter = orchestrator.YFilter
    orchestrator.EntityData.YangName = "orchestrator"
    orchestrator.EntityData.BundleName = "cisco_ios_xr"
    orchestrator.EntityData.ParentYangName = "internals"
    orchestrator.EntityData.SegmentPath = "orchestrator"
    orchestrator.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    orchestrator.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    orchestrator.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    orchestrator.EntityData.Children = make(map[string]types.YChild)
    orchestrator.EntityData.Children["operation-start-details"] = types.YChild{"OperationStartDetails", &orchestrator.OperationStartDetails}
    orchestrator.EntityData.Children["internal-prepare"] = types.YChild{"InternalPrepare", &orchestrator.InternalPrepare}
    orchestrator.EntityData.Children["internal-activate"] = types.YChild{"InternalActivate", &orchestrator.InternalActivate}
    orchestrator.EntityData.Children["error"] = types.YChild{"Error", &orchestrator.Error}
    orchestrator.EntityData.Leafs = make(map[string]types.YLeaf)
    orchestrator.EntityData.Leafs["command"] = types.YLeaf{"Command", orchestrator.Command}
    orchestrator.EntityData.Leafs["operation-type"] = types.YLeaf{"OperationType", orchestrator.OperationType}
    orchestrator.EntityData.Leafs["current-operation"] = types.YLeaf{"CurrentOperation", orchestrator.CurrentOperation}
    orchestrator.EntityData.Leafs["issu-completed"] = types.YLeaf{"IssuCompleted", orchestrator.IssuCompleted}
    orchestrator.EntityData.Leafs["operation-id"] = types.YLeaf{"OperationId", orchestrator.OperationId}
    orchestrator.EntityData.Leafs["in-progress"] = types.YLeaf{"InProgress", orchestrator.InProgress}
    return &(orchestrator.EntityData)
}

// Issu_Internals_Orchestrator_OperationStartDetails
type Issu_Internals_Orchestrator_OperationStartDetails struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Packages used to initiate the operation. The type is slice of string.
    InputPackage []interface{}

    // Operation IDs used to initiate operation. The type is slice of interface{}
    // with range: 0..4294967295.
    InputOperationId []interface{}
}

func (operationStartDetails *Issu_Internals_Orchestrator_OperationStartDetails) GetEntityData() *types.CommonEntityData {
    operationStartDetails.EntityData.YFilter = operationStartDetails.YFilter
    operationStartDetails.EntityData.YangName = "operation-start-details"
    operationStartDetails.EntityData.BundleName = "cisco_ios_xr"
    operationStartDetails.EntityData.ParentYangName = "orchestrator"
    operationStartDetails.EntityData.SegmentPath = "operation-start-details"
    operationStartDetails.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    operationStartDetails.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    operationStartDetails.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    operationStartDetails.EntityData.Children = make(map[string]types.YChild)
    operationStartDetails.EntityData.Leafs = make(map[string]types.YLeaf)
    operationStartDetails.EntityData.Leafs["input-package"] = types.YLeaf{"InputPackage", operationStartDetails.InputPackage}
    operationStartDetails.EntityData.Leafs["input-operation-id"] = types.YLeaf{"InputOperationId", operationStartDetails.InputOperationId}
    return &(operationStartDetails.EntityData)
}

// Issu_Internals_Orchestrator_InternalPrepare
type Issu_Internals_Orchestrator_InternalPrepare struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ID of prepare operation. The type is interface{} with range: 0..4294967295.
    OperationId interface{}

    // True if prepare phase complete. The type is bool.
    Complete interface{}

    // Current stage of prepare operation, if a stage is in progress. The type is
    // StageType.
    CurrentStage interface{}

    // Host packages to be used in the ISSU. The type is slice of string.
    HostPackage []interface{}

    // Calvados packages to be used in the ISSU. The type is slice of string.
    CalvadosPackage []interface{}

    
    PrepareStageHistory Issu_Internals_Orchestrator_InternalPrepare_PrepareStageHistory
}

func (internalPrepare *Issu_Internals_Orchestrator_InternalPrepare) GetEntityData() *types.CommonEntityData {
    internalPrepare.EntityData.YFilter = internalPrepare.YFilter
    internalPrepare.EntityData.YangName = "internal-prepare"
    internalPrepare.EntityData.BundleName = "cisco_ios_xr"
    internalPrepare.EntityData.ParentYangName = "orchestrator"
    internalPrepare.EntityData.SegmentPath = "internal-prepare"
    internalPrepare.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    internalPrepare.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    internalPrepare.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    internalPrepare.EntityData.Children = make(map[string]types.YChild)
    internalPrepare.EntityData.Children["prepare-stage-history"] = types.YChild{"PrepareStageHistory", &internalPrepare.PrepareStageHistory}
    internalPrepare.EntityData.Leafs = make(map[string]types.YLeaf)
    internalPrepare.EntityData.Leafs["operation-id"] = types.YLeaf{"OperationId", internalPrepare.OperationId}
    internalPrepare.EntityData.Leafs["complete"] = types.YLeaf{"Complete", internalPrepare.Complete}
    internalPrepare.EntityData.Leafs["current-stage"] = types.YLeaf{"CurrentStage", internalPrepare.CurrentStage}
    internalPrepare.EntityData.Leafs["host-package"] = types.YLeaf{"HostPackage", internalPrepare.HostPackage}
    internalPrepare.EntityData.Leafs["calvados-package"] = types.YLeaf{"CalvadosPackage", internalPrepare.CalvadosPackage}
    return &(internalPrepare.EntityData)
}

// Issu_Internals_Orchestrator_InternalPrepare_PrepareStageHistory
type Issu_Internals_Orchestrator_InternalPrepare_PrepareStageHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Issu_Internals_Orchestrator_InternalPrepare_PrepareStageHistory_HistoricalStage.
    HistoricalStage []Issu_Internals_Orchestrator_InternalPrepare_PrepareStageHistory_HistoricalStage
}

func (prepareStageHistory *Issu_Internals_Orchestrator_InternalPrepare_PrepareStageHistory) GetEntityData() *types.CommonEntityData {
    prepareStageHistory.EntityData.YFilter = prepareStageHistory.YFilter
    prepareStageHistory.EntityData.YangName = "prepare-stage-history"
    prepareStageHistory.EntityData.BundleName = "cisco_ios_xr"
    prepareStageHistory.EntityData.ParentYangName = "internal-prepare"
    prepareStageHistory.EntityData.SegmentPath = "prepare-stage-history"
    prepareStageHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    prepareStageHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    prepareStageHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    prepareStageHistory.EntityData.Children = make(map[string]types.YChild)
    prepareStageHistory.EntityData.Children["historical-stage"] = types.YChild{"HistoricalStage", nil}
    for i := range prepareStageHistory.HistoricalStage {
        prepareStageHistory.EntityData.Children[types.GetSegmentPath(&prepareStageHistory.HistoricalStage[i])] = types.YChild{"HistoricalStage", &prepareStageHistory.HistoricalStage[i]}
    }
    prepareStageHistory.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(prepareStageHistory.EntityData)
}

// Issu_Internals_Orchestrator_InternalPrepare_PrepareStageHistory_HistoricalStage
type Issu_Internals_Orchestrator_InternalPrepare_PrepareStageHistory_HistoricalStage struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Integer stage index. The type is interface{} with
    // range: 0..4294967295.
    StageIndex interface{}

    // External stage of operation. The type is OpStage.
    ExternalStage interface{}

    // Details of the internal stage. The type is string.
    InternalStageDetails interface{}

    // Status of the stage. The type is string.
    Status interface{}

    // Further error details. The type is string.
    ErrorDetails interface{}

    // Start time of stage. The type is string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    StartTime interface{}

    // End time of stage. The type is string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    EndTime interface{}

    // Is the stage complete?. The type is bool.
    Complete interface{}
}

func (historicalStage *Issu_Internals_Orchestrator_InternalPrepare_PrepareStageHistory_HistoricalStage) GetEntityData() *types.CommonEntityData {
    historicalStage.EntityData.YFilter = historicalStage.YFilter
    historicalStage.EntityData.YangName = "historical-stage"
    historicalStage.EntityData.BundleName = "cisco_ios_xr"
    historicalStage.EntityData.ParentYangName = "prepare-stage-history"
    historicalStage.EntityData.SegmentPath = "historical-stage" + "[stage-index='" + fmt.Sprintf("%v", historicalStage.StageIndex) + "']"
    historicalStage.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    historicalStage.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    historicalStage.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    historicalStage.EntityData.Children = make(map[string]types.YChild)
    historicalStage.EntityData.Leafs = make(map[string]types.YLeaf)
    historicalStage.EntityData.Leafs["stage-index"] = types.YLeaf{"StageIndex", historicalStage.StageIndex}
    historicalStage.EntityData.Leafs["external-stage"] = types.YLeaf{"ExternalStage", historicalStage.ExternalStage}
    historicalStage.EntityData.Leafs["internal-stage-details"] = types.YLeaf{"InternalStageDetails", historicalStage.InternalStageDetails}
    historicalStage.EntityData.Leafs["status"] = types.YLeaf{"Status", historicalStage.Status}
    historicalStage.EntityData.Leafs["error-details"] = types.YLeaf{"ErrorDetails", historicalStage.ErrorDetails}
    historicalStage.EntityData.Leafs["start-time"] = types.YLeaf{"StartTime", historicalStage.StartTime}
    historicalStage.EntityData.Leafs["end-time"] = types.YLeaf{"EndTime", historicalStage.EndTime}
    historicalStage.EntityData.Leafs["complete"] = types.YLeaf{"Complete", historicalStage.Complete}
    return &(historicalStage.EntityData)
}

// Issu_Internals_Orchestrator_InternalActivate
type Issu_Internals_Orchestrator_InternalActivate struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // ID of prepare operation. The type is interface{} with range: 0..4294967295.
    OperationId interface{}

    // True if activate phase complete. The type is bool.
    Complete interface{}

    // Current stage of activate operation, if a stage is in progress. The type is
    // StageType.
    CurrentStage interface{}

    // Current phase of activate operation, if phase is relevant. The type is
    // PhaseType.
    CurrentPhase interface{}

    // Indicates whether any host packages have been prepared. The type is bool.
    HostPrepared interface{}

    // Indicates whether any Calvados packages have been prepared. The type is
    // bool.
    CalvadosPrepared interface{}

    // Nodes participating in host ISSU. The type is slice of string.
    HostNode []interface{}

    // Nodes participating in phase one of Calvados ISSU. The type is slice of
    // string.
    CalvadosPhaseOneNode []interface{}

    // Nodes participating in phase two of Calvados ISSU. The type is slice of
    // string.
    CalvadosPhaseTwoNode []interface{}

    
    ActivateStageHistory Issu_Internals_Orchestrator_InternalActivate_ActivateStageHistory
}

func (internalActivate *Issu_Internals_Orchestrator_InternalActivate) GetEntityData() *types.CommonEntityData {
    internalActivate.EntityData.YFilter = internalActivate.YFilter
    internalActivate.EntityData.YangName = "internal-activate"
    internalActivate.EntityData.BundleName = "cisco_ios_xr"
    internalActivate.EntityData.ParentYangName = "orchestrator"
    internalActivate.EntityData.SegmentPath = "internal-activate"
    internalActivate.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    internalActivate.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    internalActivate.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    internalActivate.EntityData.Children = make(map[string]types.YChild)
    internalActivate.EntityData.Children["activate-stage-history"] = types.YChild{"ActivateStageHistory", &internalActivate.ActivateStageHistory}
    internalActivate.EntityData.Leafs = make(map[string]types.YLeaf)
    internalActivate.EntityData.Leafs["operation-id"] = types.YLeaf{"OperationId", internalActivate.OperationId}
    internalActivate.EntityData.Leafs["complete"] = types.YLeaf{"Complete", internalActivate.Complete}
    internalActivate.EntityData.Leafs["current-stage"] = types.YLeaf{"CurrentStage", internalActivate.CurrentStage}
    internalActivate.EntityData.Leafs["current-phase"] = types.YLeaf{"CurrentPhase", internalActivate.CurrentPhase}
    internalActivate.EntityData.Leafs["host-prepared"] = types.YLeaf{"HostPrepared", internalActivate.HostPrepared}
    internalActivate.EntityData.Leafs["calvados-prepared"] = types.YLeaf{"CalvadosPrepared", internalActivate.CalvadosPrepared}
    internalActivate.EntityData.Leafs["host-node"] = types.YLeaf{"HostNode", internalActivate.HostNode}
    internalActivate.EntityData.Leafs["calvados-phase-one-node"] = types.YLeaf{"CalvadosPhaseOneNode", internalActivate.CalvadosPhaseOneNode}
    internalActivate.EntityData.Leafs["calvados-phase-two-node"] = types.YLeaf{"CalvadosPhaseTwoNode", internalActivate.CalvadosPhaseTwoNode}
    return &(internalActivate.EntityData)
}

// Issu_Internals_Orchestrator_InternalActivate_ActivateStageHistory
type Issu_Internals_Orchestrator_InternalActivate_ActivateStageHistory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of
    // Issu_Internals_Orchestrator_InternalActivate_ActivateStageHistory_HistoricalStage.
    HistoricalStage []Issu_Internals_Orchestrator_InternalActivate_ActivateStageHistory_HistoricalStage
}

func (activateStageHistory *Issu_Internals_Orchestrator_InternalActivate_ActivateStageHistory) GetEntityData() *types.CommonEntityData {
    activateStageHistory.EntityData.YFilter = activateStageHistory.YFilter
    activateStageHistory.EntityData.YangName = "activate-stage-history"
    activateStageHistory.EntityData.BundleName = "cisco_ios_xr"
    activateStageHistory.EntityData.ParentYangName = "internal-activate"
    activateStageHistory.EntityData.SegmentPath = "activate-stage-history"
    activateStageHistory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    activateStageHistory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    activateStageHistory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    activateStageHistory.EntityData.Children = make(map[string]types.YChild)
    activateStageHistory.EntityData.Children["historical-stage"] = types.YChild{"HistoricalStage", nil}
    for i := range activateStageHistory.HistoricalStage {
        activateStageHistory.EntityData.Children[types.GetSegmentPath(&activateStageHistory.HistoricalStage[i])] = types.YChild{"HistoricalStage", &activateStageHistory.HistoricalStage[i]}
    }
    activateStageHistory.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(activateStageHistory.EntityData)
}

// Issu_Internals_Orchestrator_InternalActivate_ActivateStageHistory_HistoricalStage
type Issu_Internals_Orchestrator_InternalActivate_ActivateStageHistory_HistoricalStage struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Integer stage index. The type is interface{} with
    // range: 0..4294967295.
    StageIndex interface{}

    // External stage of operation. The type is OpStage.
    ExternalStage interface{}

    // Details of the internal stage. The type is string.
    InternalStageDetails interface{}

    // Status of the stage. The type is string.
    Status interface{}

    // Further error details. The type is string.
    ErrorDetails interface{}

    // Start time of stage. The type is string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    StartTime interface{}

    // End time of stage. The type is string with pattern:
    // b'\\d{4}-\\d{2}-\\d{2}T\\d{2}:\\d{2}:\\d{2}(\\.\\d+)?(Z|[\\+\\-]\\d{2}:\\d{2})'.
    EndTime interface{}

    // Is the stage complete?. The type is bool.
    Complete interface{}
}

func (historicalStage *Issu_Internals_Orchestrator_InternalActivate_ActivateStageHistory_HistoricalStage) GetEntityData() *types.CommonEntityData {
    historicalStage.EntityData.YFilter = historicalStage.YFilter
    historicalStage.EntityData.YangName = "historical-stage"
    historicalStage.EntityData.BundleName = "cisco_ios_xr"
    historicalStage.EntityData.ParentYangName = "activate-stage-history"
    historicalStage.EntityData.SegmentPath = "historical-stage" + "[stage-index='" + fmt.Sprintf("%v", historicalStage.StageIndex) + "']"
    historicalStage.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    historicalStage.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    historicalStage.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    historicalStage.EntityData.Children = make(map[string]types.YChild)
    historicalStage.EntityData.Leafs = make(map[string]types.YLeaf)
    historicalStage.EntityData.Leafs["stage-index"] = types.YLeaf{"StageIndex", historicalStage.StageIndex}
    historicalStage.EntityData.Leafs["external-stage"] = types.YLeaf{"ExternalStage", historicalStage.ExternalStage}
    historicalStage.EntityData.Leafs["internal-stage-details"] = types.YLeaf{"InternalStageDetails", historicalStage.InternalStageDetails}
    historicalStage.EntityData.Leafs["status"] = types.YLeaf{"Status", historicalStage.Status}
    historicalStage.EntityData.Leafs["error-details"] = types.YLeaf{"ErrorDetails", historicalStage.ErrorDetails}
    historicalStage.EntityData.Leafs["start-time"] = types.YLeaf{"StartTime", historicalStage.StartTime}
    historicalStage.EntityData.Leafs["end-time"] = types.YLeaf{"EndTime", historicalStage.EndTime}
    historicalStage.EntityData.Leafs["complete"] = types.YLeaf{"Complete", historicalStage.Complete}
    return &(historicalStage.EntityData)
}

// Issu_Internals_Orchestrator_Error
type Issu_Internals_Orchestrator_Error struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Overall status of the operation. The type is string.
    OperationStatus interface{}

    // Operation in progress when first error was encountered. The type is
    // OpRequestType.
    FailureOperation interface{}

    // Stage being undertaken when first failure was encountered. The type is
    // StageType.
    FailureExternalStage interface{}

    // Description of the internal state when first error encountered. The type is
    // string.
    FailureInternalStageDetails interface{}

    // String describing error encountered. The type is string.
    ErrorDetails interface{}

    // Nodes on which a failure occurred. The type is slice of string.
    FailedNode []interface{}

    // Packages which caused a failure. The type is slice of string.
    FailedPackage []interface{}

    // Operation IDs which caused a failure. The type is slice of interface{} with
    // range: 0..4294967295.
    FailedOperationId []interface{}

    // Registered features which caused a failure. The type is slice of string.
    FailedClient []interface{}

    // True if a recovery has been attempted or is currently in progress. The type
    // is bool.
    RecoveryAttempted interface{}

    // Status of the recovery operation. The type is string.
    RecoveryStatus interface{}
}

func (error *Issu_Internals_Orchestrator_Error) GetEntityData() *types.CommonEntityData {
    error.EntityData.YFilter = error.YFilter
    error.EntityData.YangName = "error"
    error.EntityData.BundleName = "cisco_ios_xr"
    error.EntityData.ParentYangName = "orchestrator"
    error.EntityData.SegmentPath = "error"
    error.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    error.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    error.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    error.EntityData.Children = make(map[string]types.YChild)
    error.EntityData.Leafs = make(map[string]types.YLeaf)
    error.EntityData.Leafs["operation-status"] = types.YLeaf{"OperationStatus", error.OperationStatus}
    error.EntityData.Leafs["failure-operation"] = types.YLeaf{"FailureOperation", error.FailureOperation}
    error.EntityData.Leafs["failure-external-stage"] = types.YLeaf{"FailureExternalStage", error.FailureExternalStage}
    error.EntityData.Leafs["failure-internal-stage-details"] = types.YLeaf{"FailureInternalStageDetails", error.FailureInternalStageDetails}
    error.EntityData.Leafs["error-details"] = types.YLeaf{"ErrorDetails", error.ErrorDetails}
    error.EntityData.Leafs["failed-node"] = types.YLeaf{"FailedNode", error.FailedNode}
    error.EntityData.Leafs["failed-package"] = types.YLeaf{"FailedPackage", error.FailedPackage}
    error.EntityData.Leafs["failed-operation-id"] = types.YLeaf{"FailedOperationId", error.FailedOperationId}
    error.EntityData.Leafs["failed-client"] = types.YLeaf{"FailedClient", error.FailedClient}
    error.EntityData.Leafs["recovery-attempted"] = types.YLeaf{"RecoveryAttempted", error.RecoveryAttempted}
    error.EntityData.Leafs["recovery-status"] = types.YLeaf{"RecoveryStatus", error.RecoveryStatus}
    return &(error.EntityData)
}

// Issu_Internals_Agents
// Agent module internal state
type Issu_Internals_Agents struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Data on requests being processed by Agent module.
    Requests Issu_Internals_Agents_Requests

    // Inventory of agents held by Agent module.
    Inventory Issu_Internals_Agents_Inventory

    // Reload tracking performed by Agent module.
    ReloadTracking Issu_Internals_Agents_ReloadTracking
}

func (agents *Issu_Internals_Agents) GetEntityData() *types.CommonEntityData {
    agents.EntityData.YFilter = agents.YFilter
    agents.EntityData.YangName = "agents"
    agents.EntityData.BundleName = "cisco_ios_xr"
    agents.EntityData.ParentYangName = "internals"
    agents.EntityData.SegmentPath = "agents"
    agents.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    agents.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    agents.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    agents.EntityData.Children = make(map[string]types.YChild)
    agents.EntityData.Children["requests"] = types.YChild{"Requests", &agents.Requests}
    agents.EntityData.Children["inventory"] = types.YChild{"Inventory", &agents.Inventory}
    agents.EntityData.Children["reload-tracking"] = types.YChild{"ReloadTracking", &agents.ReloadTracking}
    agents.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(agents.EntityData)
}

// Issu_Internals_Agents_Requests
// Data on requests being processed by Agent module
type Issu_Internals_Agents_Requests struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Issu_Internals_Agents_Requests_Request.
    Request []Issu_Internals_Agents_Requests_Request
}

func (requests *Issu_Internals_Agents_Requests) GetEntityData() *types.CommonEntityData {
    requests.EntityData.YFilter = requests.YFilter
    requests.EntityData.YangName = "requests"
    requests.EntityData.BundleName = "cisco_ios_xr"
    requests.EntityData.ParentYangName = "agents"
    requests.EntityData.SegmentPath = "requests"
    requests.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    requests.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    requests.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    requests.EntityData.Children = make(map[string]types.YChild)
    requests.EntityData.Children["request"] = types.YChild{"Request", nil}
    for i := range requests.Request {
        requests.EntityData.Children[types.GetSegmentPath(&requests.Request[i])] = types.YChild{"Request", &requests.Request[i]}
    }
    requests.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(requests.EntityData)
}

// Issu_Internals_Agents_Requests_Request
type Issu_Internals_Agents_Requests_Request struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Integer request index. The type is interface{}
    // with range: 0..4294967295.
    RequestIndex interface{}

    // Type of request. The type is RequestType.
    RequestType interface{}

    // Number of requests sent. Equal to number of expected responses. The type is
    // interface{} with range: 0..4294967295.
    RequestsSent interface{}

    // Number of responses received from agents. The type is interface{} with
    // range: 0..4294967295.
    ResponsesReceived interface{}

    
    Checkpoint Issu_Internals_Agents_Requests_Request_Checkpoint

    
    Agents Issu_Internals_Agents_Requests_Request_Agents_
}

func (request *Issu_Internals_Agents_Requests_Request) GetEntityData() *types.CommonEntityData {
    request.EntityData.YFilter = request.YFilter
    request.EntityData.YangName = "request"
    request.EntityData.BundleName = "cisco_ios_xr"
    request.EntityData.ParentYangName = "requests"
    request.EntityData.SegmentPath = "request" + "[request-index='" + fmt.Sprintf("%v", request.RequestIndex) + "']"
    request.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    request.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    request.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    request.EntityData.Children = make(map[string]types.YChild)
    request.EntityData.Children["checkpoint"] = types.YChild{"Checkpoint", &request.Checkpoint}
    request.EntityData.Children["agents"] = types.YChild{"Agents", &request.Agents}
    request.EntityData.Leafs = make(map[string]types.YLeaf)
    request.EntityData.Leafs["request-index"] = types.YLeaf{"RequestIndex", request.RequestIndex}
    request.EntityData.Leafs["request-type"] = types.YLeaf{"RequestType", request.RequestType}
    request.EntityData.Leafs["requests-sent"] = types.YLeaf{"RequestsSent", request.RequestsSent}
    request.EntityData.Leafs["responses-received"] = types.YLeaf{"ResponsesReceived", request.ResponsesReceived}
    return &(request.EntityData)
}

// Issu_Internals_Agents_Requests_Request_Checkpoint
type Issu_Internals_Agents_Requests_Request_Checkpoint struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Checkpoint request type. The type is MessageType.
    MessageType interface{}

    // Length of checkpoint data. 0 for start/end requests. The type is
    // interface{} with range: 0..4294967295.
    DataLength interface{}

    // Filename of associated checkpoint file. The type is string.
    Filename interface{}
}

func (checkpoint *Issu_Internals_Agents_Requests_Request_Checkpoint) GetEntityData() *types.CommonEntityData {
    checkpoint.EntityData.YFilter = checkpoint.YFilter
    checkpoint.EntityData.YangName = "checkpoint"
    checkpoint.EntityData.BundleName = "cisco_ios_xr"
    checkpoint.EntityData.ParentYangName = "request"
    checkpoint.EntityData.SegmentPath = "checkpoint"
    checkpoint.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    checkpoint.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    checkpoint.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    checkpoint.EntityData.Children = make(map[string]types.YChild)
    checkpoint.EntityData.Leafs = make(map[string]types.YLeaf)
    checkpoint.EntityData.Leafs["message-type"] = types.YLeaf{"MessageType", checkpoint.MessageType}
    checkpoint.EntityData.Leafs["data-length"] = types.YLeaf{"DataLength", checkpoint.DataLength}
    checkpoint.EntityData.Leafs["filename"] = types.YLeaf{"Filename", checkpoint.Filename}
    return &(checkpoint.EntityData)
}

// Issu_Internals_Agents_Requests_Request_Checkpoint_MessageType represents Checkpoint request type
type Issu_Internals_Agents_Requests_Request_Checkpoint_MessageType string

const (
    // Start request
    Issu_Internals_Agents_Requests_Request_Checkpoint_MessageType_requests_checkpoint_start Issu_Internals_Agents_Requests_Request_Checkpoint_MessageType = "requests-checkpoint-start"

    // End request
    Issu_Internals_Agents_Requests_Request_Checkpoint_MessageType_requests_checkpoint_end Issu_Internals_Agents_Requests_Request_Checkpoint_MessageType = "requests-checkpoint-end"

    // Update request
    Issu_Internals_Agents_Requests_Request_Checkpoint_MessageType_requests_checkpoint_update Issu_Internals_Agents_Requests_Request_Checkpoint_MessageType = "requests-checkpoint-update"
)

// Issu_Internals_Agents_Requests_Request_Agents_
type Issu_Internals_Agents_Requests_Request_Agents_ struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Issu_Internals_Agents_Requests_Request_Agents__Agent.
    Agent []Issu_Internals_Agents_Requests_Request_Agents__Agent
}

func (agents_ *Issu_Internals_Agents_Requests_Request_Agents_) GetEntityData() *types.CommonEntityData {
    agents_.EntityData.YFilter = agents_.YFilter
    agents_.EntityData.YangName = "agents"
    agents_.EntityData.BundleName = "cisco_ios_xr"
    agents_.EntityData.ParentYangName = "request"
    agents_.EntityData.SegmentPath = "agents"
    agents_.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    agents_.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    agents_.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    agents_.EntityData.Children = make(map[string]types.YChild)
    agents_.EntityData.Children["agent"] = types.YChild{"Agent", nil}
    for i := range agents_.Agent {
        agents_.EntityData.Children[types.GetSegmentPath(&agents_.Agent[i])] = types.YChild{"Agent", &agents_.Agent[i]}
    }
    agents_.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(agents_.EntityData)
}

// Issu_Internals_Agents_Requests_Request_Agents__Agent
type Issu_Internals_Agents_Requests_Request_Agents__Agent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Integer agent index. The type is interface{} with
    // range: 0..4294967295.
    AgentIndex interface{}

    // Node ID. The type is string.
    Node interface{}

    // Indicates whether this agent has responded. The type is bool.
    WaitingForResponse interface{}

    
    ResponseContents Issu_Internals_Agents_Requests_Request_Agents__Agent_ResponseContents
}

func (agent *Issu_Internals_Agents_Requests_Request_Agents__Agent) GetEntityData() *types.CommonEntityData {
    agent.EntityData.YFilter = agent.YFilter
    agent.EntityData.YangName = "agent"
    agent.EntityData.BundleName = "cisco_ios_xr"
    agent.EntityData.ParentYangName = "agents"
    agent.EntityData.SegmentPath = "agent" + "[agent-index='" + fmt.Sprintf("%v", agent.AgentIndex) + "']"
    agent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    agent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    agent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    agent.EntityData.Children = make(map[string]types.YChild)
    agent.EntityData.Children["response-contents"] = types.YChild{"ResponseContents", &agent.ResponseContents}
    agent.EntityData.Leafs = make(map[string]types.YLeaf)
    agent.EntityData.Leafs["agent-index"] = types.YLeaf{"AgentIndex", agent.AgentIndex}
    agent.EntityData.Leafs["node"] = types.YLeaf{"Node", agent.Node}
    agent.EntityData.Leafs["waiting-for-response"] = types.YLeaf{"WaitingForResponse", agent.WaitingForResponse}
    return &(agent.EntityData)
}

// Issu_Internals_Agents_Requests_Request_Agents__Agent_ResponseContents
type Issu_Internals_Agents_Requests_Request_Agents__Agent_ResponseContents struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Enum indicating node status. The type is AgentStatus.
    AgentStatus interface{}

    // Further details of error occuring. The type is string.
    ErrorDetails interface{}
}

func (responseContents *Issu_Internals_Agents_Requests_Request_Agents__Agent_ResponseContents) GetEntityData() *types.CommonEntityData {
    responseContents.EntityData.YFilter = responseContents.YFilter
    responseContents.EntityData.YangName = "response-contents"
    responseContents.EntityData.BundleName = "cisco_ios_xr"
    responseContents.EntityData.ParentYangName = "agent"
    responseContents.EntityData.SegmentPath = "response-contents"
    responseContents.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    responseContents.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    responseContents.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    responseContents.EntityData.Children = make(map[string]types.YChild)
    responseContents.EntityData.Leafs = make(map[string]types.YLeaf)
    responseContents.EntityData.Leafs["agent-status"] = types.YLeaf{"AgentStatus", responseContents.AgentStatus}
    responseContents.EntityData.Leafs["error-details"] = types.YLeaf{"ErrorDetails", responseContents.ErrorDetails}
    return &(responseContents.EntityData)
}

// Issu_Internals_Agents_Requests_Request_Agents__Agent_ResponseContents_AgentStatus represents Enum indicating node status
type Issu_Internals_Agents_Requests_Request_Agents__Agent_ResponseContents_AgentStatus string

const (
    // Node is ready/operation successful
    Issu_Internals_Agents_Requests_Request_Agents__Agent_ResponseContents_AgentStatus_agent_response_ok Issu_Internals_Agents_Requests_Request_Agents__Agent_ResponseContents_AgentStatus = "agent-response-ok"

    // Otherwise undefined error
    Issu_Internals_Agents_Requests_Request_Agents__Agent_ResponseContents_AgentStatus_agent_response_error Issu_Internals_Agents_Requests_Request_Agents__Agent_ResponseContents_AgentStatus = "agent-response-error"

    // Timeout during request
    Issu_Internals_Agents_Requests_Request_Agents__Agent_ResponseContents_AgentStatus_agent_response_timeout Issu_Internals_Agents_Requests_Request_Agents__Agent_ResponseContents_AgentStatus = "agent-response-timeout"

    // Failed to send request
    Issu_Internals_Agents_Requests_Request_Agents__Agent_ResponseContents_AgentStatus_agent_response_send_failure Issu_Internals_Agents_Requests_Request_Agents__Agent_ResponseContents_AgentStatus = "agent-response-send-failure"
)

// Issu_Internals_Agents_Requests_Request_RequestType represents Type of request
type Issu_Internals_Agents_Requests_Request_RequestType string

const (
    // Node ready request
    Issu_Internals_Agents_Requests_Request_RequestType_requests_node_ready Issu_Internals_Agents_Requests_Request_RequestType = "requests-node-ready"

    // Checkpoint request
    Issu_Internals_Agents_Requests_Request_RequestType_requests_checkpoint Issu_Internals_Agents_Requests_Request_RequestType = "requests-checkpoint"

    // Post-upgrade cleanup request
    Issu_Internals_Agents_Requests_Request_RequestType_requests_post_upgrade_cleanup Issu_Internals_Agents_Requests_Request_RequestType = "requests-post-upgrade-cleanup"
)

// Issu_Internals_Agents_Inventory
// Inventory of agents held by Agent module
type Issu_Internals_Agents_Inventory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The type is slice of Issu_Internals_Agents_Inventory_Agent.
    Agent []Issu_Internals_Agents_Inventory_Agent
}

func (inventory *Issu_Internals_Agents_Inventory) GetEntityData() *types.CommonEntityData {
    inventory.EntityData.YFilter = inventory.YFilter
    inventory.EntityData.YangName = "inventory"
    inventory.EntityData.BundleName = "cisco_ios_xr"
    inventory.EntityData.ParentYangName = "agents"
    inventory.EntityData.SegmentPath = "inventory"
    inventory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inventory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inventory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inventory.EntityData.Children = make(map[string]types.YChild)
    inventory.EntityData.Children["agent"] = types.YChild{"Agent", nil}
    for i := range inventory.Agent {
        inventory.EntityData.Children[types.GetSegmentPath(&inventory.Agent[i])] = types.YChild{"Agent", &inventory.Agent[i]}
    }
    inventory.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(inventory.EntityData)
}

// Issu_Internals_Agents_Inventory_Agent
type Issu_Internals_Agents_Inventory_Agent struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Integer agent index. The type is interface{} with
    // range: 0..4294967295.
    AgentIndex interface{}

    // Agent node ID. The type is string.
    AgentNode interface{}

    // True if node has been reloaded. False otherwise. The type is bool.
    Reloaded interface{}
}

func (agent *Issu_Internals_Agents_Inventory_Agent) GetEntityData() *types.CommonEntityData {
    agent.EntityData.YFilter = agent.YFilter
    agent.EntityData.YangName = "agent"
    agent.EntityData.BundleName = "cisco_ios_xr"
    agent.EntityData.ParentYangName = "inventory"
    agent.EntityData.SegmentPath = "agent" + "[agent-index='" + fmt.Sprintf("%v", agent.AgentIndex) + "']"
    agent.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    agent.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    agent.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    agent.EntityData.Children = make(map[string]types.YChild)
    agent.EntityData.Leafs = make(map[string]types.YLeaf)
    agent.EntityData.Leafs["agent-index"] = types.YLeaf{"AgentIndex", agent.AgentIndex}
    agent.EntityData.Leafs["agent-node"] = types.YLeaf{"AgentNode", agent.AgentNode}
    agent.EntityData.Leafs["reloaded"] = types.YLeaf{"Reloaded", agent.Reloaded}
    return &(agent.EntityData)
}

// Issu_Internals_Agents_ReloadTracking
// Reload tracking performed by Agent module
type Issu_Internals_Agents_ReloadTracking struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // True if reload tracking in progress. The type is bool.
    InProgress interface{}

    // Number of nodes which have not yet reloaded. The type is interface{} with
    // range: 0..4294967295.
    RemainingNodesCount interface{}

    // The type is slice of Issu_Internals_Agents_ReloadTracking_Node.
    Node []Issu_Internals_Agents_ReloadTracking_Node
}

func (reloadTracking *Issu_Internals_Agents_ReloadTracking) GetEntityData() *types.CommonEntityData {
    reloadTracking.EntityData.YFilter = reloadTracking.YFilter
    reloadTracking.EntityData.YangName = "reload-tracking"
    reloadTracking.EntityData.BundleName = "cisco_ios_xr"
    reloadTracking.EntityData.ParentYangName = "agents"
    reloadTracking.EntityData.SegmentPath = "reload-tracking"
    reloadTracking.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    reloadTracking.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    reloadTracking.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    reloadTracking.EntityData.Children = make(map[string]types.YChild)
    reloadTracking.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range reloadTracking.Node {
        reloadTracking.EntityData.Children[types.GetSegmentPath(&reloadTracking.Node[i])] = types.YChild{"Node", &reloadTracking.Node[i]}
    }
    reloadTracking.EntityData.Leafs = make(map[string]types.YLeaf)
    reloadTracking.EntityData.Leafs["in-progress"] = types.YLeaf{"InProgress", reloadTracking.InProgress}
    reloadTracking.EntityData.Leafs["remaining-nodes-count"] = types.YLeaf{"RemainingNodesCount", reloadTracking.RemainingNodesCount}
    return &(reloadTracking.EntityData)
}

// Issu_Internals_Agents_ReloadTracking_Node
type Issu_Internals_Agents_ReloadTracking_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Integer node index. The type is interface{} with
    // range: 0..4294967295.
    NodeIndex interface{}

    // Node ID. The type is string.
    Id interface{}

    // True if node has been reloaded. False otherwise. The type is bool.
    Reloaded interface{}
}

func (node *Issu_Internals_Agents_ReloadTracking_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "reload-tracking"
    node.EntityData.SegmentPath = "node" + "[node-index='" + fmt.Sprintf("%v", node.NodeIndex) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node-index"] = types.YLeaf{"NodeIndex", node.NodeIndex}
    node.EntityData.Leafs["id"] = types.YLeaf{"Id", node.Id}
    node.EntityData.Leafs["reloaded"] = types.YLeaf{"Reloaded", node.Reloaded}
    return &(node.EntityData)
}

// Issu_Internals_InventoryMonitor
// Inventory monitor module internal state
type Issu_Internals_InventoryMonitor struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // Inventory of nodes in the system held by ISSU Director.
    Inventory Issu_Internals_InventoryMonitor_Inventory
}

func (inventoryMonitor *Issu_Internals_InventoryMonitor) GetEntityData() *types.CommonEntityData {
    inventoryMonitor.EntityData.YFilter = inventoryMonitor.YFilter
    inventoryMonitor.EntityData.YangName = "inventory-monitor"
    inventoryMonitor.EntityData.BundleName = "cisco_ios_xr"
    inventoryMonitor.EntityData.ParentYangName = "internals"
    inventoryMonitor.EntityData.SegmentPath = "inventory-monitor"
    inventoryMonitor.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inventoryMonitor.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inventoryMonitor.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inventoryMonitor.EntityData.Children = make(map[string]types.YChild)
    inventoryMonitor.EntityData.Children["inventory"] = types.YChild{"Inventory", &inventoryMonitor.Inventory}
    inventoryMonitor.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(inventoryMonitor.EntityData)
}

// Issu_Internals_InventoryMonitor_Inventory
// Inventory of nodes in the system held by ISSU Director
type Issu_Internals_InventoryMonitor_Inventory struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // The tree of nodes within the Inventory monitor module, keyed by node ID.
    // The type is slice of Issu_Internals_InventoryMonitor_Inventory_Node.
    Node []Issu_Internals_InventoryMonitor_Inventory_Node
}

func (inventory *Issu_Internals_InventoryMonitor_Inventory) GetEntityData() *types.CommonEntityData {
    inventory.EntityData.YFilter = inventory.YFilter
    inventory.EntityData.YangName = "inventory"
    inventory.EntityData.BundleName = "cisco_ios_xr"
    inventory.EntityData.ParentYangName = "inventory-monitor"
    inventory.EntityData.SegmentPath = "inventory"
    inventory.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    inventory.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    inventory.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    inventory.EntityData.Children = make(map[string]types.YChild)
    inventory.EntityData.Children["node"] = types.YChild{"Node", nil}
    for i := range inventory.Node {
        inventory.EntityData.Children[types.GetSegmentPath(&inventory.Node[i])] = types.YChild{"Node", &inventory.Node[i]}
    }
    inventory.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(inventory.EntityData)
}

// Issu_Internals_InventoryMonitor_Inventory_Node
// The tree of nodes within the Inventory monitor module, keyed by node ID
type Issu_Internals_InventoryMonitor_Inventory_Node struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. Node ID. The type is string.
    Node interface{}

    // Node IP address. The type is string.
    Ip interface{}
}

func (node *Issu_Internals_InventoryMonitor_Inventory_Node) GetEntityData() *types.CommonEntityData {
    node.EntityData.YFilter = node.YFilter
    node.EntityData.YangName = "node"
    node.EntityData.BundleName = "cisco_ios_xr"
    node.EntityData.ParentYangName = "inventory"
    node.EntityData.SegmentPath = "node" + "[node='" + fmt.Sprintf("%v", node.Node) + "']"
    node.EntityData.CapabilitiesTable = cisco_ios_xr.GetCapabilities()
    node.EntityData.NamespaceTable = cisco_ios_xr.GetNamespaces()
    node.EntityData.BundleYangModelsLocation = cisco_ios_xr.GetModelsPath()

    node.EntityData.Children = make(map[string]types.YChild)
    node.EntityData.Leafs = make(map[string]types.YLeaf)
    node.EntityData.Leafs["node"] = types.YLeaf{"Node", node.Node}
    node.EntityData.Leafs["ip"] = types.YLeaf{"Ip", node.Ip}
    return &(node.EntityData)
}

// Issu_Internals_OpRequestType represents Enumeration of requests that initiate (part of) an ISSU operation
type Issu_Internals_OpRequestType string

const (
    // Request to execute the prepare phase of an operation
    Issu_Internals_OpRequestType_operation_request_prepare Issu_Internals_OpRequestType = "operation-request-prepare"

    // Request to activate prepared packages
    Issu_Internals_OpRequestType_operation_request_activate Issu_Internals_OpRequestType = "operation-request-activate"

    // Request to execute both prepare and deactivate phases of an operation
    Issu_Internals_OpRequestType_operation_request_deactivate Issu_Internals_OpRequestType = "operation-request-deactivate"

    // Recover the system from a failed activate or deactivate operation
    Issu_Internals_OpRequestType_operation_request_recover Issu_Internals_OpRequestType = "operation-request-recover"
)

// Issu_Internals_PhaseType represents Calvados activate phase
type Issu_Internals_PhaseType string

const (
    // Phase one of the Calvados activate operation
    Issu_Internals_PhaseType_calvados_activate_phase_one Issu_Internals_PhaseType = "calvados-activate-phase-one"

    // Phase two of the Calvados activate operation
    Issu_Internals_PhaseType_calvados_activate_phase_two Issu_Internals_PhaseType = "calvados-activate-phase-two"
)

// Issu_Internals_StageType represents Enumeration of possible internal stages in an operation
type Issu_Internals_StageType string

const (
    // Start of operation
    Issu_Internals_StageType_start Issu_Internals_StageType = "start"

    // End of operation
    Issu_Internals_StageType_end Issu_Internals_StageType = "end"

    // Inventory precheck
    Issu_Internals_StageType_prepare_inventory_precheck Issu_Internals_StageType = "prepare-inventory-precheck"

    // Expand add operation IDs to packages
    Issu_Internals_StageType_prepare_expand_operation_ids Issu_Internals_StageType = "prepare-expand-operation-ids"

    // Retrieve package metadata
    Issu_Internals_StageType_prepare_get_metadata Issu_Internals_StageType = "prepare-get-metadata"

    // Extract constituent packages
    Issu_Internals_StageType_prepare_extract_composite Issu_Internals_StageType = "prepare-extract-composite"

    // Verify package contents
    Issu_Internals_StageType_prepare_verify_packages Issu_Internals_StageType = "prepare-verify-packages"

    // The host preparation
    Issu_Internals_StageType_prepare_host Issu_Internals_StageType = "prepare-host"

    // The Calvados preparation
    Issu_Internals_StageType_prepare_calvados Issu_Internals_StageType = "prepare-calvados"

    // Verify package contents for deactivation
    Issu_Internals_StageType_prepare_deactivate_verify_packages Issu_Internals_StageType = "prepare-deactivate-verify-packages"

    // The Calvados preparation for deactivation
    Issu_Internals_StageType_prepare_deactivate_calvados Issu_Internals_StageType = "prepare-deactivate-calvados"

    // Abort the prepare operation
    Issu_Internals_StageType_prepare_abort Issu_Internals_StageType = "prepare-abort"

    // Clean-up of prepare operation
    Issu_Internals_StageType_prepare_clean Issu_Internals_StageType = "prepare-clean"

    // Activation preamble
    Issu_Internals_StageType_activate_preamble Issu_Internals_StageType = "activate-preamble"

    // Deactivation preamble
    Issu_Internals_StageType_deactivate_preamble Issu_Internals_StageType = "deactivate-preamble"

    // Calvados-specific activation preamble
    Issu_Internals_StageType_activate_calvados_preamble Issu_Internals_StageType = "activate-calvados-preamble"

    // Activate the Calvados software
    Issu_Internals_StageType_activate_calvados_phase Issu_Internals_StageType = "activate-calvados-phase"

    // Reload Calvados VMs
    Issu_Internals_StageType_activate_calvados_phase_reload Issu_Internals_StageType = "activate-calvados-phase-reload"

    // Post VM restart checks for Calvados
    Issu_Internals_StageType_activate_calvados_phase_postamble Issu_Internals_StageType = "activate-calvados-phase-postamble"

    // Post Calvados activation handling
    Issu_Internals_StageType_activate_calvados_postamble Issu_Internals_StageType = "activate-calvados-postamble"

    // Execute host ISSU
    Issu_Internals_StageType_activate_host Issu_Internals_StageType = "activate-host"

    // Post-processing for the activate operation
    Issu_Internals_StageType_activate_postamble Issu_Internals_StageType = "activate-postamble"

    // Deactivate the Calvados software
    Issu_Internals_StageType_deactivate_calvados Issu_Internals_StageType = "deactivate-calvados"

    // Abort the activate operation, no recovery is necessary
    Issu_Internals_StageType_activate_abort_no_recovery Issu_Internals_StageType = "activate-abort-no-recovery"

    // Abort the activate operation, no recovery is possible
    Issu_Internals_StageType_activate_abort_unrecoverable Issu_Internals_StageType = "activate-abort-unrecoverable"

    // Pause following an error
    Issu_Internals_StageType_activate_error_pause Issu_Internals_StageType = "activate-error-pause"

    // Rollback to committed Calvados
    Issu_Internals_StageType_activate_calvados_recovery Issu_Internals_StageType = "activate-calvados-recovery"

    // Post rollback to committed Calvados
    Issu_Internals_StageType_activate_recovery_postamble Issu_Internals_StageType = "activate-recovery-postamble"
)

