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

    
    Casconfig CISCOAAASERVERMIB_Casconfig

    // This table shows current configurations for each AAA server, allows
    // existing servers to	be removed and new ones to be created.
    Casconfigtable CISCOAAASERVERMIB_Casconfigtable
}

func (cISCOAAASERVERMIB *CISCOAAASERVERMIB) GetEntityData() *types.CommonEntityData {
    cISCOAAASERVERMIB.EntityData.YFilter = cISCOAAASERVERMIB.YFilter
    cISCOAAASERVERMIB.EntityData.YangName = "CISCO-AAA-SERVER-MIB"
    cISCOAAASERVERMIB.EntityData.BundleName = "cisco_ios_xe"
    cISCOAAASERVERMIB.EntityData.ParentYangName = "CISCO-AAA-SERVER-MIB"
    cISCOAAASERVERMIB.EntityData.SegmentPath = "CISCO-AAA-SERVER-MIB:CISCO-AAA-SERVER-MIB"
    cISCOAAASERVERMIB.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    cISCOAAASERVERMIB.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    cISCOAAASERVERMIB.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    cISCOAAASERVERMIB.EntityData.Children = make(map[string]types.YChild)
    cISCOAAASERVERMIB.EntityData.Children["casConfig"] = types.YChild{"Casconfig", &cISCOAAASERVERMIB.Casconfig}
    cISCOAAASERVERMIB.EntityData.Children["casConfigTable"] = types.YChild{"Casconfigtable", &cISCOAAASERVERMIB.Casconfigtable}
    cISCOAAASERVERMIB.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(cISCOAAASERVERMIB.EntityData)
}

// CISCOAAASERVERMIB_Casconfig
type CISCOAAASERVERMIB_Casconfig struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This variable controls the	generation of casServerStateChange notification.
    // When this variable	is true(1), generation of casServerStateChange
    // notifications	is enabled. When this variable	is false(2), generation	of
    // casServerStateChange notifications	is disabled.  The default value is
    // false(2). The type is bool.
    Casserverstatechangeenable interface{}
}

func (casconfig *CISCOAAASERVERMIB_Casconfig) GetEntityData() *types.CommonEntityData {
    casconfig.EntityData.YFilter = casconfig.YFilter
    casconfig.EntityData.YangName = "casConfig"
    casconfig.EntityData.BundleName = "cisco_ios_xe"
    casconfig.EntityData.ParentYangName = "CISCO-AAA-SERVER-MIB"
    casconfig.EntityData.SegmentPath = "casConfig"
    casconfig.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    casconfig.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    casconfig.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    casconfig.EntityData.Children = make(map[string]types.YChild)
    casconfig.EntityData.Leafs = make(map[string]types.YLeaf)
    casconfig.EntityData.Leafs["casServerStateChangeEnable"] = types.YLeaf{"Casserverstatechangeenable", casconfig.Casserverstatechangeenable}
    return &(casconfig.EntityData)
}

// CISCOAAASERVERMIB_Casconfigtable
// This table shows current configurations for each
// AAA server, allows existing servers to	be removed
// and new ones to be created.
type CISCOAAASERVERMIB_Casconfigtable struct {
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
    // CISCOAAASERVERMIB_Casconfigtable_Casconfigentry.
    Casconfigentry []CISCOAAASERVERMIB_Casconfigtable_Casconfigentry
}

func (casconfigtable *CISCOAAASERVERMIB_Casconfigtable) GetEntityData() *types.CommonEntityData {
    casconfigtable.EntityData.YFilter = casconfigtable.YFilter
    casconfigtable.EntityData.YangName = "casConfigTable"
    casconfigtable.EntityData.BundleName = "cisco_ios_xe"
    casconfigtable.EntityData.ParentYangName = "CISCO-AAA-SERVER-MIB"
    casconfigtable.EntityData.SegmentPath = "casConfigTable"
    casconfigtable.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    casconfigtable.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    casconfigtable.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    casconfigtable.EntityData.Children = make(map[string]types.YChild)
    casconfigtable.EntityData.Children["casConfigEntry"] = types.YChild{"Casconfigentry", nil}
    for i := range casconfigtable.Casconfigentry {
        casconfigtable.EntityData.Children[types.GetSegmentPath(&casconfigtable.Casconfigentry[i])] = types.YChild{"Casconfigentry", &casconfigtable.Casconfigentry[i]}
    }
    casconfigtable.EntityData.Leafs = make(map[string]types.YLeaf)
    return &(casconfigtable.EntityData)
}

// CISCOAAASERVERMIB_Casconfigtable_Casconfigentry
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
type CISCOAAASERVERMIB_Casconfigtable_Casconfigentry struct {
    EntityData types.CommonEntityData
    YFilter yfilter.YFilter

    // This attribute is a key. The variable denotes the protocol used by the
    // managed device with the AAA server corresponding to  this entry in the
    // table. The type is CiscoAAAProtocol.
    Casprotocol interface{}

    // This attribute is a key. A management station wishing to initiate a	new
    // AAA	server configuration should use a	random value for this object when
    // creating an instance of casConfigEntry.  The RowStatus semantics of	the
    // casConfigRowStatus object will prevent access conflicts.  If	the randomly
    // chosen casIndex value for row creation is	already	in use by an existing
    // entry, snmp set to the casIndex value will fail. The type is interface{}
    // with range: 1..4294967295.
    Casindex interface{}

    // The IP address of the server. The type is string with pattern:
    // b'(([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9][0-9]|2[0-4][0-9]|25[0-5])(%[\\p{N}\\p{L}]+)?'.
    Casaddress interface{}

    // UDP/TCP port used for authentication in the configuration  For TACACS+,
    // this object should be	explictly set.  Default value is the IOS default for
    // radius: 1645. The type is interface{} with range: 0..65535.
    Casauthenport interface{}

    // UDP/TCP port used for accounting service in the configuration  For TACACS+,
    // the value of casAcctPort is ignored. casAuthenPort will	be used	instead. 
    // Default value is the IOS default for radius: 1646. The type is interface{}
    // with range: 0..65535.
    Casacctport interface{}

    // The server key	to be used with	this server.  Retrieving the	value of this
    // object via SNMP will return	an empty string	for security reasons. The type
    // is string.
    Caskey interface{}

    // A number that indicates the priority of the server	in this entry. 
    // Lower	numbers	indicate higher	priority. The type is interface{} with range:
    // 1..4294967295.
    Caspriority interface{}

    // The status of this table entry.  Once the entry status	is set
    // to	active,	the associated entry cannot be modified except	destroyed by
    // setting this object to destroy(6). The type is RowStatus.
    Casconfigrowstatus interface{}

    // The number	of authentication requests sent	to this server since it is made
    // active.  Retransmissions due to request timeouts are counted as	distinct
    // requests. The type is interface{} with range: 0..4294967295.
    Casauthenrequests interface{}

    // The number	of authentication requests which have timed out since it	is
    // made	active.  A timeout results in a retransmission of the request If	the
    // maximum number of attempts has been	reached,
    // no	further	retransmissions	will be	attempted. The type is interface{} with
    // range: 0..4294967295.
    Casauthenrequesttimeouts interface{}

    // The number	of unexpected authentication responses received from this server
    // since it is made active.  An	example	is a delayed response to a request
    // which had already timed out. The type is interface{} with range:
    // 0..4294967295.
    Casauthenunexpectedresponses interface{}

    // The number	of server ERROR	authentication responses received from
    // this	server since it	is made	active.  These are responses indicating that
    // the server itself has identified an error with its authentication
    // operation. The type is interface{} with range: 0..4294967295.
    Casauthenservererrorresponses interface{}

    // The number	of authentication responses which could	not be	processed
    // since	it is made active.  Reasons include inability to decrypt the
    // response, invalid fields, or	the response is	not valid based	on the
    // request. The type is interface{} with range: 0..4294967295.
    Casauthenincorrectresponses interface{}

    // Average response time for authentication requests sent to	this server,
    // excluding timeouts, since system re-initialization. The type is interface{}
    // with range: 0..2147483647.
    Casauthenresponsetime interface{}

    // The number	of authentication transactions with this server which succeeded
    // since it is	made active.  A transaction may include multiple	request
    // retransmissions if	timeouts occur.  A transaction is successful if
    // the	server responds with either an authentication pass	or fail. The type is
    // interface{} with range: 0..4294967295.
    Casauthentransactionsuccesses interface{}

    // The number	of authentication transactions with this server which failed
    // since it is made active.  A transaction may include multiple	request
    // retransmissions if	timeouts occur.  A transaction failure occurs if maximum
    // resends have been met or the server aborts the transaction. The type is
    // interface{} with range: 0..4294967295.
    Casauthentransactionfailures interface{}

    // The number	of authorization requests sent to this server since it is made
    // active.  Retransmissions due to request timeouts are counted as	distinct
    // requests.  This object is not	instantiated for protocols which do not
    // support a distinct authorization function. The type is interface{} with
    // range: 0..4294967295.
    Casauthorrequests interface{}

    // The number	of authorization requests which	have timed out since it	is
    // made	active.  A timeout results in a retransmission of the request If	the
    // maximum number of attempts has been	reached,
    // no	further	retransmissions	will be	attempted.  This object is
    // not	instantiated for protocols which do not support a distinct
    // authorization function. The type is interface{} with range: 0..4294967295.
    Casauthorrequesttimeouts interface{}

    // The number	of unexpected authorization responses received from this server
    // since it is made active.  An	example	is a delayed response to a request
    // which had already timed out.  This object is not	instantiated for protocols
    // which do not support a distinct authorization function. The type is
    // interface{} with range: 0..4294967295.
    Casauthorunexpectedresponses interface{}

    // The number	of server ERROR	authorization responses received from
    // this	server since it	is made	active.  These are responses indicating that
    // the server itself has identified an error with its authorization operation.
    // This object is not	instantiated for protocols which do not support a
    // distinct authorization function. The type is interface{} with range:
    // 0..4294967295.
    Casauthorservererrorresponses interface{}

    // The number	of authorization responses which could not be	processed since	it
    // is made active.  Reasons include inability to decrypt the response, invalid
    // fields, or	the response is	not valid based	on the request.  This object is
    // not	instantiated for protocols which do not support a distinct
    // authorization function. The type is interface{} with range: 0..4294967295.
    Casauthorincorrectresponses interface{}

    // Average response time for authorization requests sent to	this server,
    // excluding timeouts, since system re-initialization.  This object is
    // not	instantiated for protocols which do not support a distinct
    // authorization function. The type is interface{} with range: 0..2147483647.
    Casauthorresponsetime interface{}

    // The number	of authorization transactions with this server which succeeded
    // since it is	made active.  A transaction may include multiple	request
    // retransmissions if	timeouts occur.  A transaction is successful if
    // the	server responds with either an authorization pass or fail.  This object
    // is not	instantiated for protocols which do not support a distinct
    // authorization function. The type is interface{} with range: 0..4294967295.
    Casauthortransactionsuccesses interface{}

    // The number	of authorization transactions with this server which failed
    // since it is made active.  A transaction may include multiple	request
    // retransmissions if	timeouts occur.  A transaction failure occurs if maximum
    // resends have been met or the server aborts the transaction.  This object is
    // not	instantiated for protocols which do not support a distinct
    // authorization function. The type is interface{} with range: 0..4294967295.
    Casauthortransactionfailures interface{}

    // The number	of accounting requests sent to this server since system
    // re-initialization.  Retransmissions due to request timeouts are counted
    // as	distinct requests. The type is interface{} with range: 0..4294967295.
    Casacctrequests interface{}

    // The number	of accounting requests which have timed out since system
    // re-initialization.  A timeout results in a retransmission of the request
    // If	the maximum number of attempts has been	reached,
    // no	further	retransmissions	will be	attempted. The type is interface{} with
    // range: 0..4294967295.
    Casacctrequesttimeouts interface{}

    // The number	of unexpected accounting responses received from this server
    // since system re-initialization.  An	example	is a delayed response to a
    // request which had already timed out. The type is interface{} with range:
    // 0..4294967295.
    Casacctunexpectedresponses interface{}

    // The number	of server ERROR	accounting responses received from this server
    // since system re-initialization.  These are responses indicating that the
    // server itself has identified an error with its accounting operation. The
    // type is interface{} with range: 0..4294967295.
    Casacctservererrorresponses interface{}

    // The number	of accounting responses	which could not be	processed
    // since	system re-initialization.  Reasons include inability to decrypt the
    // response, invalid fields, or	the response is	not valid based	on the
    // request. The type is interface{} with range: 0..4294967295.
    Casacctincorrectresponses interface{}

    // Average response time for accounting requests sent to	this server,, since
    // system re-initialization excluding timeouts. The type is interface{} with
    // range: 0..2147483647.
    Casacctresponsetime interface{}

    // The number	of accounting transactions with	this server which succeeded
    // since system re-initialization.  A transaction may include multiple	request
    // retransmissions if	timeouts occur.  A transaction is successful if
    // the	server responds with either an accounting pass or fail. The type is
    // interface{} with range: 0..4294967295.
    Casaccttransactionsuccesses interface{}

    // The number	of accounting transactions with	this server which failed since
    // system re-initialization.  A transaction may include multiple	request
    // retransmissions if	timeouts occur.  A transaction failure occurs if maximum
    // resends have been met or the server aborts the transaction. The type is
    // interface{} with range: 0..4294967295.
    Casaccttransactionfailures interface{}

    // Current state of this server.  up(1)	 - Server responding to	requests 
    // dead(2) - Server failed to respond  A server is marked	dead if	it does	not
    // respond after maximum retransmissions.  A server is marked	up again
    // either	after a	waiting period or if some response	is received from it.  The
    // initial value of casState is 'up(1)' at system re-initialization.	This will
    // only transistion to 'dead(2)' if	an attempt to communicate fails. The type
    // is Casstate.
    Casstate interface{}

    // This object provides the elapsed time the server has been in its current
    // state as shown	in casState. The type is interface{} with range:
    // 0..2147483647.
    Cascurrentstateduration interface{}

    // This object provides the elapsed time the server was been in its previous
    // state	prior to the most recent transistion of casState.  This value	is
    // zero	if the server has not changed state. The type is interface{} with
    // range: 0..2147483647.
    Caspreviousstateduration interface{}

    // The total elapsed time this server's casState has had the value 'dead(2)'
    // since system re-initialization. The type is interface{} with range:
    // 0..2147483647.
    Castotaldeadtime interface{}

    // The number	of times this server's casState	has transitioned to 'dead(2)'
    // since system re-initialization. The type is interface{} with range:
    // 0..4294967295.
    Casdeadcount interface{}
}

func (casconfigentry *CISCOAAASERVERMIB_Casconfigtable_Casconfigentry) GetEntityData() *types.CommonEntityData {
    casconfigentry.EntityData.YFilter = casconfigentry.YFilter
    casconfigentry.EntityData.YangName = "casConfigEntry"
    casconfigentry.EntityData.BundleName = "cisco_ios_xe"
    casconfigentry.EntityData.ParentYangName = "casConfigTable"
    casconfigentry.EntityData.SegmentPath = "casConfigEntry" + "[casProtocol='" + fmt.Sprintf("%v", casconfigentry.Casprotocol) + "']" + "[casIndex='" + fmt.Sprintf("%v", casconfigentry.Casindex) + "']"
    casconfigentry.EntityData.CapabilitiesTable = cisco_ios_xe.GetCapabilities()
    casconfigentry.EntityData.NamespaceTable = cisco_ios_xe.GetNamespaces()
    casconfigentry.EntityData.BundleYangModelsLocation = cisco_ios_xe.GetModelsPath()

    casconfigentry.EntityData.Children = make(map[string]types.YChild)
    casconfigentry.EntityData.Leafs = make(map[string]types.YLeaf)
    casconfigentry.EntityData.Leafs["casProtocol"] = types.YLeaf{"Casprotocol", casconfigentry.Casprotocol}
    casconfigentry.EntityData.Leafs["casIndex"] = types.YLeaf{"Casindex", casconfigentry.Casindex}
    casconfigentry.EntityData.Leafs["casAddress"] = types.YLeaf{"Casaddress", casconfigentry.Casaddress}
    casconfigentry.EntityData.Leafs["casAuthenPort"] = types.YLeaf{"Casauthenport", casconfigentry.Casauthenport}
    casconfigentry.EntityData.Leafs["casAcctPort"] = types.YLeaf{"Casacctport", casconfigentry.Casacctport}
    casconfigentry.EntityData.Leafs["casKey"] = types.YLeaf{"Caskey", casconfigentry.Caskey}
    casconfigentry.EntityData.Leafs["casPriority"] = types.YLeaf{"Caspriority", casconfigentry.Caspriority}
    casconfigentry.EntityData.Leafs["casConfigRowStatus"] = types.YLeaf{"Casconfigrowstatus", casconfigentry.Casconfigrowstatus}
    casconfigentry.EntityData.Leafs["casAuthenRequests"] = types.YLeaf{"Casauthenrequests", casconfigentry.Casauthenrequests}
    casconfigentry.EntityData.Leafs["casAuthenRequestTimeouts"] = types.YLeaf{"Casauthenrequesttimeouts", casconfigentry.Casauthenrequesttimeouts}
    casconfigentry.EntityData.Leafs["casAuthenUnexpectedResponses"] = types.YLeaf{"Casauthenunexpectedresponses", casconfigentry.Casauthenunexpectedresponses}
    casconfigentry.EntityData.Leafs["casAuthenServerErrorResponses"] = types.YLeaf{"Casauthenservererrorresponses", casconfigentry.Casauthenservererrorresponses}
    casconfigentry.EntityData.Leafs["casAuthenIncorrectResponses"] = types.YLeaf{"Casauthenincorrectresponses", casconfigentry.Casauthenincorrectresponses}
    casconfigentry.EntityData.Leafs["casAuthenResponseTime"] = types.YLeaf{"Casauthenresponsetime", casconfigentry.Casauthenresponsetime}
    casconfigentry.EntityData.Leafs["casAuthenTransactionSuccesses"] = types.YLeaf{"Casauthentransactionsuccesses", casconfigentry.Casauthentransactionsuccesses}
    casconfigentry.EntityData.Leafs["casAuthenTransactionFailures"] = types.YLeaf{"Casauthentransactionfailures", casconfigentry.Casauthentransactionfailures}
    casconfigentry.EntityData.Leafs["casAuthorRequests"] = types.YLeaf{"Casauthorrequests", casconfigentry.Casauthorrequests}
    casconfigentry.EntityData.Leafs["casAuthorRequestTimeouts"] = types.YLeaf{"Casauthorrequesttimeouts", casconfigentry.Casauthorrequesttimeouts}
    casconfigentry.EntityData.Leafs["casAuthorUnexpectedResponses"] = types.YLeaf{"Casauthorunexpectedresponses", casconfigentry.Casauthorunexpectedresponses}
    casconfigentry.EntityData.Leafs["casAuthorServerErrorResponses"] = types.YLeaf{"Casauthorservererrorresponses", casconfigentry.Casauthorservererrorresponses}
    casconfigentry.EntityData.Leafs["casAuthorIncorrectResponses"] = types.YLeaf{"Casauthorincorrectresponses", casconfigentry.Casauthorincorrectresponses}
    casconfigentry.EntityData.Leafs["casAuthorResponseTime"] = types.YLeaf{"Casauthorresponsetime", casconfigentry.Casauthorresponsetime}
    casconfigentry.EntityData.Leafs["casAuthorTransactionSuccesses"] = types.YLeaf{"Casauthortransactionsuccesses", casconfigentry.Casauthortransactionsuccesses}
    casconfigentry.EntityData.Leafs["casAuthorTransactionFailures"] = types.YLeaf{"Casauthortransactionfailures", casconfigentry.Casauthortransactionfailures}
    casconfigentry.EntityData.Leafs["casAcctRequests"] = types.YLeaf{"Casacctrequests", casconfigentry.Casacctrequests}
    casconfigentry.EntityData.Leafs["casAcctRequestTimeouts"] = types.YLeaf{"Casacctrequesttimeouts", casconfigentry.Casacctrequesttimeouts}
    casconfigentry.EntityData.Leafs["casAcctUnexpectedResponses"] = types.YLeaf{"Casacctunexpectedresponses", casconfigentry.Casacctunexpectedresponses}
    casconfigentry.EntityData.Leafs["casAcctServerErrorResponses"] = types.YLeaf{"Casacctservererrorresponses", casconfigentry.Casacctservererrorresponses}
    casconfigentry.EntityData.Leafs["casAcctIncorrectResponses"] = types.YLeaf{"Casacctincorrectresponses", casconfigentry.Casacctincorrectresponses}
    casconfigentry.EntityData.Leafs["casAcctResponseTime"] = types.YLeaf{"Casacctresponsetime", casconfigentry.Casacctresponsetime}
    casconfigentry.EntityData.Leafs["casAcctTransactionSuccesses"] = types.YLeaf{"Casaccttransactionsuccesses", casconfigentry.Casaccttransactionsuccesses}
    casconfigentry.EntityData.Leafs["casAcctTransactionFailures"] = types.YLeaf{"Casaccttransactionfailures", casconfigentry.Casaccttransactionfailures}
    casconfigentry.EntityData.Leafs["casState"] = types.YLeaf{"Casstate", casconfigentry.Casstate}
    casconfigentry.EntityData.Leafs["casCurrentStateDuration"] = types.YLeaf{"Cascurrentstateduration", casconfigentry.Cascurrentstateduration}
    casconfigentry.EntityData.Leafs["casPreviousStateDuration"] = types.YLeaf{"Caspreviousstateduration", casconfigentry.Caspreviousstateduration}
    casconfigentry.EntityData.Leafs["casTotalDeadTime"] = types.YLeaf{"Castotaldeadtime", casconfigentry.Castotaldeadtime}
    casconfigentry.EntityData.Leafs["casDeadCount"] = types.YLeaf{"Casdeadcount", casconfigentry.Casdeadcount}
    return &(casconfigentry.EntityData)
}

// CISCOAAASERVERMIB_Casconfigtable_Casconfigentry_Casstate represents if	an attempt to communicate fails.
type CISCOAAASERVERMIB_Casconfigtable_Casconfigentry_Casstate string

const (
    CISCOAAASERVERMIB_Casconfigtable_Casconfigentry_Casstate_up CISCOAAASERVERMIB_Casconfigtable_Casconfigentry_Casstate = "up"

    CISCOAAASERVERMIB_Casconfigtable_Casconfigentry_Casstate_dead CISCOAAASERVERMIB_Casconfigtable_Casconfigentry_Casstate = "dead"
)

