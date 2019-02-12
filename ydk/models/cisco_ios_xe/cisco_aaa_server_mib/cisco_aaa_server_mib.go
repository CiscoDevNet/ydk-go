// The MIB module	for monitoring communications and status
// of AAA	Server operation
package cisco_aaa_server_mib

import (
    "fmt"
    "github.com/CiscoDevNet/ydk-go/ydk"
    "github.com/CiscoDevNet/ydk-go/ydk/types"
    "github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
    "github.com/CiscoDevNet/ydk-go/ydk/models/cisco_ios_xe"
    "reflect"
)

func init() {
    ydk.YLogDebug(fmt.Sprintf("Registering top level entities for package cisco_aaa_server_mib"))
    ydk.RegisterEntity("{urn:ietf:params:xml:ns:yang:smiv2:CISCO-AAA-SERVER-MIB CISCO-AAA-SERVER-MIB}", reflect.TypeOf(CISCOAAASERVERMIB{}))
    ydk.RegisterEntity("CISCO-AAA-SERVER-MIB:CISCO-AAA-SERVER-MIB", reflect.TypeOf(CISCOAAASERVERMIB{}))
}

// CiscoAAAProtocol represents other(7)    -   Other protocols
type CiscoAAAProtocol string

const (
    CiscoAAAProtocol_tacacsplus CiscoAAAProtocol = "tacacsplus"

    CiscoAAAProtocol_radius CiscoAAAProtocol = "radius"

    CiscoAAAProtocol_ldap CiscoAAAProtocol = "ldap"

    CiscoAAAProtocol_kerberos CiscoAAAProtocol = "kerberos"

    CiscoAAAProtocol_ntlm CiscoAAAProtocol = "ntlm"

    CiscoAAAProtocol_sdi CiscoAAAProtocol = "sdi"

    CiscoAAAProtocol_other CiscoAAAProtocol = "other"
)

// CISCOAAASERVERMIB
type CISCOAAASERVERMIB struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    
    CasConfig CISCOAAASERVERMIB_CasConfig

    // This table shows current configurations for each AAA server, allows
    // existing servers to	be removed and new ones to be created.
    CasConfigTable CISCOAAASERVERMIB_CasConfigTable
}

func (cISCOAAASERVERMIB *CISCOAAASERVERMIB) GetEntityData() *types.CommonEntityData {
    cISCOAAASERVERMIB.EntityData.YFilter = cISCOAAASERVERMIB.YFilter
    cISCOAAASERVERMIB.EntityData.YangName = "CISCO-AAA-SERVER-MIB"
    cISCOAAASERVERMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOAAASERVERMIB.EntityData.ParentYangName = "CISCO-AAA-SERVER-MIB"
    cISCOAAASERVERMIB.EntityData.SegmentPath = "CISCO-AAA-SERVER-MIB:CISCO-AAA-SERVER-MIB"
    cISCOAAASERVERMIB.EntityData.AbsolutePath = cISCOAAASERVERMIB.EntityData.SegmentPath
    cISCOAAASERVERMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOAAASERVERMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOAAASERVERMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOAAASERVERMIB.EntityData.Children = types.NewOrderedMap()
    cISCOAAASERVERMIB.EntityData.Children.Append("casConfig", types.YChild{"CasConfig", &cISCOAAASERVERMIB.CasConfig})
    cISCOAAASERVERMIB.EntityData.Children.Append("casConfigTable", types.YChild{"CasConfigTable", &cISCOAAASERVERMIB.CasConfigTable})
    cISCOAAASERVERMIB.EntityData.Leafs = types.NewOrderedMap()

    cISCOAAASERVERMIB.EntityData.YListKeys = []string {}

    return &(cISCOAAASERVERMIB.EntityData)
}

// CISCOAAASERVERMIB_CasConfig
type CISCOAAASERVERMIB_CasConfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This variable controls the	generation of casServerStateChange notification.
    // When this variable	is true(1), generation of casServerStateChange
    // notifications	is enabled. When this variable	is false(2), generation	of
    // casServerStateChange notifications	is disabled.  The default value is
    // false(2). The type is bool.
    CasServerStateChangeEnable interface{}
}

func (casConfig *CISCOAAASERVERMIB_CasConfig) GetEntityData() *types.CommonEntityData {
    casConfig.EntityData.YFilter = casConfig.YFilter
    casConfig.EntityData.YangName = "casConfig"
    casConfig.EntityData.BundleName = "cisco_ios_xe"
    casConfig.EntityData.ParentYangName = "CISCO-AAA-SERVER-MIB"
    casConfig.EntityData.SegmentPath = "casConfig"
    casConfig.EntityData.AbsolutePath = "CISCO-AAA-SERVER-MIB:CISCO-AAA-SERVER-MIB/" + casConfig.EntityData.SegmentPath
    casConfig.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    casConfig.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    casConfig.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    casConfig.EntityData.Children = types.NewOrderedMap()
    casConfig.EntityData.Leafs = types.NewOrderedMap()
    casConfig.EntityData.Leafs.Append("casServerStateChangeEnable", types.YLeaf{"CasServerStateChangeEnable", casConfig.CasServerStateChangeEnable})

    casConfig.EntityData.YListKeys = []string {}

    return &(casConfig.EntityData)
}

// CISCOAAASERVERMIB_CasConfigTable
// This table shows current configurations for each
// AAA server, allows existing servers to	be removed
// and new ones to be created.
type CISCOAAASERVERMIB_CasConfigTable struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // An	AAA server configuration identified by its protocol and its index. 
    // An	entry is created/removed when a	server is defined or	undefined with IOS
    // configuration commands via CLI or by issuing appropriate sets	to this	table
    // using snmp.  A management station wishing to create an entry should first
    // generate a random number to be used as the index to	this sparse table. 
    // The	station	should then create the associated	instance of the	row status
    // and row index objects. It	must also, either in the same or in successive
    // PDUs, create an instance	of casAddress where casAddress is the
    // IP	address	of the server to be added.  It	should also modify the default
    // values for casAuthenPort, casAcctPort if the	defaults are not appropriate. 
    // If	casKey is a zero-length	string or is not explicitly set, then the global
    // key will be used.	Otherwise, this	value is	used as	the key	for this
    // server	instance.  Once the appropriate instance of all the configuration
    // objects have been created,	either by an explicit SNMP set request or	by
    // default, the	row status should be set to active(1) to initiate the request.
    // After the AAA server is made active, the entry can	not be modified -	the
    // only allowed operation after this is to destroy the entry by setting
    // casConfigRowStatus to	destroy(6).  casPriority is automatically assigned
    // once	the entry is made active and reflects the relative priority of the
    // defined server with respect to already configured servers. Newly-created
    // servers will	be assigned the	lowest priority. To	reassign server	priorities
    // to existing server entries, it	may be necessary to destroy and	recreate
    // entries in order of	priority.  Entries in	this table
    // with	casConfigRowStatus equal to active(1) remain in the table until
    // destroyed.  Entries in	this table with	casConfigRowStatus equal to values
    // other than active(1) will be destroyed after timeout (5	minutes).  If	a
    // server address being created via SNMP	exists already in	another	active
    // casConfigEntry, then a newly created row can not be	made active until the
    // original row with	the with the same server address value	is destroyed. 
    // Upon reload, casIndex values may be changed, but the priorities	that were
    // saved	before reload will be retained, with lowest priority number
    // corresponding to the higher priority servers. The type is slice of
    // CISCOAAASERVERMIB_CasConfigTable_CasConfigEntry.
    CasConfigEntry []*CISCOAAASERVERMIB_CasConfigTable_CasConfigEntry
}

func (casConfigTable *CISCOAAASERVERMIB_CasConfigTable) GetEntityData() *types.CommonEntityData {
    casConfigTable.EntityData.YFilter = casConfigTable.YFilter
    casConfigTable.EntityData.YangName = "casConfigTable"
    casConfigTable.EntityData.BundleName = "cisco_ios_xe"
    casConfigTable.EntityData.ParentYangName = "CISCO-AAA-SERVER-MIB"
    casConfigTable.EntityData.SegmentPath = "casConfigTable"
    casConfigTable.EntityData.AbsolutePath = "CISCO-AAA-SERVER-MIB:CISCO-AAA-SERVER-MIB/" + casConfigTable.EntityData.SegmentPath
    casConfigTable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    casConfigTable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    casConfigTable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    casConfigTable.EntityData.Children = types.NewOrderedMap()
    casConfigTable.EntityData.Children.Append("casConfigEntry", types.YChild{"CasConfigEntry", nil})
    for i := range casConfigTable.CasConfigEntry {
        casConfigTable.EntityData.Children.Append(types.GetSegmentPath(casConfigTable.CasConfigEntry[i]), types.YChild{"CasConfigEntry", casConfigTable.CasConfigEntry[i]})
    }
    casConfigTable.EntityData.Leafs = types.NewOrderedMap()

    casConfigTable.EntityData.YListKeys = []string {}

    return &(casConfigTable.EntityData)
}

// CISCOAAASERVERMIB_CasConfigTable_CasConfigEntry
// An	AAA server configuration identified by its protocol
// and its index.
// 
// An	entry is created/removed when a	server is defined
// or	undefined with IOS configuration commands via
// CLI or by issuing appropriate sets	to this	table
// using snmp.
// 
// A management station wishing to create an entry should
// first generate a random number to be used as the index
// to	this sparse table.  The	station	should then create the
// associated	instance of the	row status and row index objects.
// It	must also, either in the same or in successive PDUs,
// create an instance	of casAddress where casAddress is the
// IP	address	of the server to be added.
// 
// It	should also modify the default values for casAuthenPort,
// casAcctPort if the	defaults are not appropriate.
// 
// If	casKey is a zero-length	string or is not explicitly set,
// then the global key will be used.	Otherwise, this	value
// is	used as	the key	for this server	instance.
// 
// Once the appropriate instance of all the configuration
// objects have been created,	either by an explicit SNMP set
// request or	by default, the	row status should be set to
// active(1) to initiate the request.
// 
// After the AAA server is made active, the entry can	not be
// modified -	the only allowed operation after this is to
// destroy the entry by setting casConfigRowStatus to	destroy(6).
// 
// casPriority is automatically assigned once	the entry is
// made active and reflects the relative priority of the
// defined server with respect to already configured servers.
// Newly-created servers will	be assigned the	lowest priority.
// To	reassign server	priorities to existing server entries,
// it	may be necessary to destroy and	recreate entries in order
// of	priority.
// 
// Entries in	this table with	casConfigRowStatus equal to
// active(1) remain in the table until destroyed.
// 
// Entries in	this table with	casConfigRowStatus equal to
// values other than active(1) will be destroyed after timeout
// (5	minutes).
// 
// If	a server address being created via SNMP	exists already
// in	another	active casConfigEntry, then a newly created row
// can not be	made active until the original row with	the
// with the same server address value	is destroyed.
// 
// Upon reload, casIndex values may be changed, but the
// priorities	that were saved	before reload will be retained,
// with lowest priority number corresponding to the higher
// priority servers.
type CISCOAAASERVERMIB_CasConfigTable_CasConfigEntry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter
    YListKey string

    // This attribute is a key. The variable denotes the protocol used by the
    // managed device with the AAA server corresponding to  this entry in the
    // table. The type is CiscoAAAProtocol.
    CasProtocol interface{}

    // This attribute is a key. A management station wishing to initiate a	new
    // AAA	server configuration should use a	random value for this object when
    // creating an instance of casConfigEntry.  The RowStatus semantics of	the
    // casConfigRowStatus object will prevent access conflicts.  If	the randomly
    // chosen casIndex value for row creation is	already	in use by an existing
    // entry, snmp set to the casIndex value will fail. The type is interface{}
    // with range: 1..4294967295.
    CasIndex interface{}

    // The IP address of the server. The type is string with pattern:
    // (([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\p{N}\p{L}]+)?.
    CasAddress interface{}

    // UDP/TCP port used for authentication in the configuration  For TACACS+,
    // this object should be	explictly set.  Default value is the IOS default for
    // radius: 1645. The type is interface{} with range: 0..65535.
    CasAuthenPort interface{}

    // UDP/TCP port used for accounting service in the configuration  For TACACS+,
    // the value of casAcctPort is ignored. casAuthenPort will	be used	instead. 
    // Default value is the IOS default for radius: 1646. The type is interface{}
    // with range: 0..65535.
    CasAcctPort interface{}

    // The server key	to be used with	this server.  Retrieving the	value of this
    // object via SNMP will return	an empty string	for security reasons. The type
    // is string.
    CasKey interface{}

    // A number that indicates the priority of the server	in this entry. 
    // Lower	numbers	indicate higher	priority. The type is interface{} with range:
    // 1..4294967295.
    CasPriority interface{}

    // The status of this table entry.  Once the entry status	is set
    // to	active,	the associated entry cannot be modified except	destroyed by
    // setting this object to destroy(6). The type is RowStatus.
    CasConfigRowStatus interface{}

    // The number	of authentication requests sent	to this server since it is made
    // active.  Retransmissions due to request timeouts are counted as	distinct
    // requests. The type is interface{} with range: 0..4294967295.
    CasAuthenRequests interface{}

    // The number	of authentication requests which have timed out since it	is
    // made	active.  A timeout results in a retransmission of the request If	the
    // maximum number of attempts has been	reached,
    // no	further	retransmissions	will be	attempted. The type is interface{} with
    // range: 0..4294967295.
    CasAuthenRequestTimeouts interface{}

    // The number	of unexpected authentication responses received from this server
    // since it is made active.  An	example	is a delayed response to a request
    // which had already timed out. The type is interface{} with range:
    // 0..4294967295.
    CasAuthenUnexpectedResponses interface{}

    // The number	of server ERROR	authentication responses received from
    // this	server since it	is made	active.  These are responses indicating that
    // the server itself has identified an error with its authentication
    // operation. The type is interface{} with range: 0..4294967295.
    CasAuthenServerErrorResponses interface{}

    // The number	of authentication responses which could	not be	processed
    // since	it is made active.  Reasons include inability to decrypt the
    // response, invalid fields, or	the response is	not valid based	on the
    // request. The type is interface{} with range: 0..4294967295.
    CasAuthenIncorrectResponses interface{}

    // Average response time for authentication requests sent to	this server,
    // excluding timeouts, since system re-initialization. The type is interface{}
    // with range: 0..2147483647.
    CasAuthenResponseTime interface{}

    // The number	of authentication transactions with this server which succeeded
    // since it is	made active.  A transaction may include multiple	request
    // retransmissions if	timeouts occur.  A transaction is successful if
    // the	server responds with either an authentication pass	or fail. The type is
    // interface{} with range: 0..4294967295.
    CasAuthenTransactionSuccesses interface{}

    // The number	of authentication transactions with this server which failed
    // since it is made active.  A transaction may include multiple	request
    // retransmissions if	timeouts occur.  A transaction failure occurs if maximum
    // resends have been met or the server aborts the transaction. The type is
    // interface{} with range: 0..4294967295.
    CasAuthenTransactionFailures interface{}

    // The number	of authorization requests sent to this server since it is made
    // active.  Retransmissions due to request timeouts are counted as	distinct
    // requests.  This object is not	instantiated for protocols which do not
    // support a distinct authorization function. The type is interface{} with
    // range: 0..4294967295.
    CasAuthorRequests interface{}

    // The number	of authorization requests which	have timed out since it	is
    // made	active.  A timeout results in a retransmission of the request If	the
    // maximum number of attempts has been	reached,
    // no	further	retransmissions	will be	attempted.  This object is
    // not	instantiated for protocols which do not support a distinct
    // authorization function. The type is interface{} with range: 0..4294967295.
    CasAuthorRequestTimeouts interface{}

    // The number	of unexpected authorization responses received from this server
    // since it is made active.  An	example	is a delayed response to a request
    // which had already timed out.  This object is not	instantiated for protocols
    // which do not support a distinct authorization function. The type is
    // interface{} with range: 0..4294967295.
    CasAuthorUnexpectedResponses interface{}

    // The number	of server ERROR	authorization responses received from
    // this	server since it	is made	active.  These are responses indicating that
    // the server itself has identified an error with its authorization operation.
    // This object is not	instantiated for protocols which do not support a
    // distinct authorization function. The type is interface{} with range:
    // 0..4294967295.
    CasAuthorServerErrorResponses interface{}

    // The number	of authorization responses which could not be	processed since	it
    // is made active.  Reasons include inability to decrypt the response, invalid
    // fields, or	the response is	not valid based	on the request.  This object is
    // not	instantiated for protocols which do not support a distinct
    // authorization function. The type is interface{} with range: 0..4294967295.
    CasAuthorIncorrectResponses interface{}

    // Average response time for authorization requests sent to	this server,
    // excluding timeouts, since system re-initialization.  This object is
    // not	instantiated for protocols which do not support a distinct
    // authorization function. The type is interface{} with range: 0..2147483647.
    CasAuthorResponseTime interface{}

    // The number	of authorization transactions with this server which succeeded
    // since it is	made active.  A transaction may include multiple	request
    // retransmissions if	timeouts occur.  A transaction is successful if
    // the	server responds with either an authorization pass or fail.  This object
    // is not	instantiated for protocols which do not support a distinct
    // authorization function. The type is interface{} with range: 0..4294967295.
    CasAuthorTransactionSuccesses interface{}

    // The number	of authorization transactions with this server which failed
    // since it is made active.  A transaction may include multiple	request
    // retransmissions if	timeouts occur.  A transaction failure occurs if maximum
    // resends have been met or the server aborts the transaction.  This object is
    // not	instantiated for protocols which do not support a distinct
    // authorization function. The type is interface{} with range: 0..4294967295.
    CasAuthorTransactionFailures interface{}

    // The number	of accounting requests sent to this server since system
    // re-initialization.  Retransmissions due to request timeouts are counted
    // as	distinct requests. The type is interface{} with range: 0..4294967295.
    CasAcctRequests interface{}

    // The number	of accounting requests which have timed out since system
    // re-initialization.  A timeout results in a retransmission of the request
    // If	the maximum number of attempts has been	reached,
    // no	further	retransmissions	will be	attempted. The type is interface{} with
    // range: 0..4294967295.
    CasAcctRequestTimeouts interface{}

    // The number	of unexpected accounting responses received from this server
    // since system re-initialization.  An	example	is a delayed response to a
    // request which had already timed out. The type is interface{} with range:
    // 0..4294967295.
    CasAcctUnexpectedResponses interface{}

    // The number	of server ERROR	accounting responses received from this server
    // since system re-initialization.  These are responses indicating that the
    // server itself has identified an error with its accounting operation. The
    // type is interface{} with range: 0..4294967295.
    CasAcctServerErrorResponses interface{}

    // The number	of accounting responses	which could not be	processed
    // since	system re-initialization.  Reasons include inability to decrypt the
    // response, invalid fields, or	the response is	not valid based	on the
    // request. The type is interface{} with range: 0..4294967295.
    CasAcctIncorrectResponses interface{}

    // Average response time for accounting requests sent to	this server,, since
    // system re-initialization excluding timeouts. The type is interface{} with
    // range: 0..2147483647.
    CasAcctResponseTime interface{}

    // The number	of accounting transactions with	this server which succeeded
    // since system re-initialization.  A transaction may include multiple	request
    // retransmissions if	timeouts occur.  A transaction is successful if
    // the	server responds with either an accounting pass or fail. The type is
    // interface{} with range: 0..4294967295.
    CasAcctTransactionSuccesses interface{}

    // The number	of accounting transactions with	this server which failed since
    // system re-initialization.  A transaction may include multiple	request
    // retransmissions if	timeouts occur.  A transaction failure occurs if maximum
    // resends have been met or the server aborts the transaction. The type is
    // interface{} with range: 0..4294967295.
    CasAcctTransactionFailures interface{}

    // Current state of this server.  up(1)	 - Server responding to	requests 
    // dead(2) - Server failed to respond  A server is marked	dead if	it does	not
    // respond after maximum retransmissions.  A server is marked	up again
    // either	after a	waiting period or if some response	is received from it.  The
    // initial value of casState is 'up(1)' at system re-initialization.	This will
    // only transistion to 'dead(2)' if	an attempt to communicate fails. The type
    // is CasState.
    CasState interface{}

    // This object provides the elapsed time the server has been in its current
    // state as shown	in casState. The type is interface{} with range:
    // 0..2147483647.
    CasCurrentStateDuration interface{}

    // This object provides the elapsed time the server was been in its previous
    // state	prior to the most recent transistion of casState.  This value	is
    // zero	if the server has not changed state. The type is interface{} with
    // range: 0..2147483647.
    CasPreviousStateDuration interface{}

    // The total elapsed time this server's casState has had the value 'dead(2)'
    // since system re-initialization. The type is interface{} with range:
    // 0..2147483647.
    CasTotalDeadTime interface{}

    // The number	of times this server's casState	has transitioned to 'dead(2)'
    // since system re-initialization. The type is interface{} with range:
    // 0..4294967295.
    CasDeadCount interface{}
}

func (casConfigEntry *CISCOAAASERVERMIB_CasConfigTable_CasConfigEntry) GetEntityData() *types.CommonEntityData {
    casConfigEntry.EntityData.YFilter = casConfigEntry.YFilter
    casConfigEntry.EntityData.YangName = "casConfigEntry"
    casConfigEntry.EntityData.BundleName = "cisco_ios_xe"
    casConfigEntry.EntityData.ParentYangName = "casConfigTable"
    casConfigEntry.EntityData.SegmentPath = "casConfigEntry" + types.AddKeyToken(casConfigEntry.CasProtocol, "casProtocol") + types.AddKeyToken(casConfigEntry.CasIndex, "casIndex")
    casConfigEntry.EntityData.AbsolutePath = "CISCO-AAA-SERVER-MIB:CISCO-AAA-SERVER-MIB/casConfigTable/" + casConfigEntry.EntityData.SegmentPath
    casConfigEntry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    casConfigEntry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    casConfigEntry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    casConfigEntry.EntityData.Children = types.NewOrderedMap()
    casConfigEntry.EntityData.Leafs = types.NewOrderedMap()
    casConfigEntry.EntityData.Leafs.Append("casProtocol", types.YLeaf{"CasProtocol", casConfigEntry.CasProtocol})
    casConfigEntry.EntityData.Leafs.Append("casIndex", types.YLeaf{"CasIndex", casConfigEntry.CasIndex})
    casConfigEntry.EntityData.Leafs.Append("casAddress", types.YLeaf{"CasAddress", casConfigEntry.CasAddress})
    casConfigEntry.EntityData.Leafs.Append("casAuthenPort", types.YLeaf{"CasAuthenPort", casConfigEntry.CasAuthenPort})
    casConfigEntry.EntityData.Leafs.Append("casAcctPort", types.YLeaf{"CasAcctPort", casConfigEntry.CasAcctPort})
    casConfigEntry.EntityData.Leafs.Append("casKey", types.YLeaf{"CasKey", casConfigEntry.CasKey})
    casConfigEntry.EntityData.Leafs.Append("casPriority", types.YLeaf{"CasPriority", casConfigEntry.CasPriority})
    casConfigEntry.EntityData.Leafs.Append("casConfigRowStatus", types.YLeaf{"CasConfigRowStatus", casConfigEntry.CasConfigRowStatus})
    casConfigEntry.EntityData.Leafs.Append("casAuthenRequests", types.YLeaf{"CasAuthenRequests", casConfigEntry.CasAuthenRequests})
    casConfigEntry.EntityData.Leafs.Append("casAuthenRequestTimeouts", types.YLeaf{"CasAuthenRequestTimeouts", casConfigEntry.CasAuthenRequestTimeouts})
    casConfigEntry.EntityData.Leafs.Append("casAuthenUnexpectedResponses", types.YLeaf{"CasAuthenUnexpectedResponses", casConfigEntry.CasAuthenUnexpectedResponses})
    casConfigEntry.EntityData.Leafs.Append("casAuthenServerErrorResponses", types.YLeaf{"CasAuthenServerErrorResponses", casConfigEntry.CasAuthenServerErrorResponses})
    casConfigEntry.EntityData.Leafs.Append("casAuthenIncorrectResponses", types.YLeaf{"CasAuthenIncorrectResponses", casConfigEntry.CasAuthenIncorrectResponses})
    casConfigEntry.EntityData.Leafs.Append("casAuthenResponseTime", types.YLeaf{"CasAuthenResponseTime", casConfigEntry.CasAuthenResponseTime})
    casConfigEntry.EntityData.Leafs.Append("casAuthenTransactionSuccesses", types.YLeaf{"CasAuthenTransactionSuccesses", casConfigEntry.CasAuthenTransactionSuccesses})
    casConfigEntry.EntityData.Leafs.Append("casAuthenTransactionFailures", types.YLeaf{"CasAuthenTransactionFailures", casConfigEntry.CasAuthenTransactionFailures})
    casConfigEntry.EntityData.Leafs.Append("casAuthorRequests", types.YLeaf{"CasAuthorRequests", casConfigEntry.CasAuthorRequests})
    casConfigEntry.EntityData.Leafs.Append("casAuthorRequestTimeouts", types.YLeaf{"CasAuthorRequestTimeouts", casConfigEntry.CasAuthorRequestTimeouts})
    casConfigEntry.EntityData.Leafs.Append("casAuthorUnexpectedResponses", types.YLeaf{"CasAuthorUnexpectedResponses", casConfigEntry.CasAuthorUnexpectedResponses})
    casConfigEntry.EntityData.Leafs.Append("casAuthorServerErrorResponses", types.YLeaf{"CasAuthorServerErrorResponses", casConfigEntry.CasAuthorServerErrorResponses})
    casConfigEntry.EntityData.Leafs.Append("casAuthorIncorrectResponses", types.YLeaf{"CasAuthorIncorrectResponses", casConfigEntry.CasAuthorIncorrectResponses})
    casConfigEntry.EntityData.Leafs.Append("casAuthorResponseTime", types.YLeaf{"CasAuthorResponseTime", casConfigEntry.CasAuthorResponseTime})
    casConfigEntry.EntityData.Leafs.Append("casAuthorTransactionSuccesses", types.YLeaf{"CasAuthorTransactionSuccesses", casConfigEntry.CasAuthorTransactionSuccesses})
    casConfigEntry.EntityData.Leafs.Append("casAuthorTransactionFailures", types.YLeaf{"CasAuthorTransactionFailures", casConfigEntry.CasAuthorTransactionFailures})
    casConfigEntry.EntityData.Leafs.Append("casAcctRequests", types.YLeaf{"CasAcctRequests", casConfigEntry.CasAcctRequests})
    casConfigEntry.EntityData.Leafs.Append("casAcctRequestTimeouts", types.YLeaf{"CasAcctRequestTimeouts", casConfigEntry.CasAcctRequestTimeouts})
    casConfigEntry.EntityData.Leafs.Append("casAcctUnexpectedResponses", types.YLeaf{"CasAcctUnexpectedResponses", casConfigEntry.CasAcctUnexpectedResponses})
    casConfigEntry.EntityData.Leafs.Append("casAcctServerErrorResponses", types.YLeaf{"CasAcctServerErrorResponses", casConfigEntry.CasAcctServerErrorResponses})
    casConfigEntry.EntityData.Leafs.Append("casAcctIncorrectResponses", types.YLeaf{"CasAcctIncorrectResponses", casConfigEntry.CasAcctIncorrectResponses})
    casConfigEntry.EntityData.Leafs.Append("casAcctResponseTime", types.YLeaf{"CasAcctResponseTime", casConfigEntry.CasAcctResponseTime})
    casConfigEntry.EntityData.Leafs.Append("casAcctTransactionSuccesses", types.YLeaf{"CasAcctTransactionSuccesses", casConfigEntry.CasAcctTransactionSuccesses})
    casConfigEntry.EntityData.Leafs.Append("casAcctTransactionFailures", types.YLeaf{"CasAcctTransactionFailures", casConfigEntry.CasAcctTransactionFailures})
    casConfigEntry.EntityData.Leafs.Append("casState", types.YLeaf{"CasState", casConfigEntry.CasState})
    casConfigEntry.EntityData.Leafs.Append("casCurrentStateDuration", types.YLeaf{"CasCurrentStateDuration", casConfigEntry.CasCurrentStateDuration})
    casConfigEntry.EntityData.Leafs.Append("casPreviousStateDuration", types.YLeaf{"CasPreviousStateDuration", casConfigEntry.CasPreviousStateDuration})
    casConfigEntry.EntityData.Leafs.Append("casTotalDeadTime", types.YLeaf{"CasTotalDeadTime", casConfigEntry.CasTotalDeadTime})
    casConfigEntry.EntityData.Leafs.Append("casDeadCount", types.YLeaf{"CasDeadCount", casConfigEntry.CasDeadCount})

    casConfigEntry.EntityData.YListKeys = []string {"CasProtocol", "CasIndex"}

    return &(casConfigEntry.EntityData)
}

// CISCOAAASERVERMIB_CasConfigTable_CasConfigEntry_CasState represents if	an attempt to communicate fails.
type CISCOAAASERVERMIB_CasConfigTable_CasConfigEntry_CasState string

const (
    CISCOAAASERVERMIB_CasConfigTable_CasConfigEntry_CasState_up CISCOAAASERVERMIB_CasConfigTable_CasConfigEntry_CasState = "up"

    CISCOAAASERVERMIB_CasConfigTable_CasConfigEntry_CasState_dead CISCOAAASERVERMIB_CasConfigTable_CasConfigEntry_CasState = "dead"
)

