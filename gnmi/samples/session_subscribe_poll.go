package main

import (
	"fmt"
	"time"
	"runtime"
	"path/filepath"
	"os"
	"github.com/CiscoDevNet/ydk-go/ydk"
	"github.com/CiscoDevNet/ydk-go/ydk/path"
	"github.com/CiscoDevNet/ydk-go/ydk/types"
	encoding "github.com/CiscoDevNet/ydk-go/ydk/types/encoding_format"
)

func configureBgp(session *path.GnmiSession, schema types.RootSchemaNode) {
    bgp := path.CreateRootDataNode( schema, "openconfig-bgp:bgp")
    path.CreateDataNode( bgp, "global/config/as", 65172)
    neighbor := path.CreateDataNode( bgp, "neighbors/neighbor[neighbor-address='172.16.255.2']", "")
    path.CreateDataNode( neighbor, "config/neighbor-address", "172.16.255.2")
    path.CreateDataNode( neighbor, "config/peer-as", 65172);   

    setRpc := path.CreateRpc( schema, "ydk:gnmi-set")
    bgpCreatePayload := path.CodecEncode( bgp, encoding.JSON, false)
    path.CreateDataNode( setRpc.Input, "replace[alias='bgp']/entity", bgpCreatePayload)

    session.ExecuteRpc(setRpc);
}

func deleteConfigBgp(session *path.GnmiSession, schema types.RootSchemaNode) {
    bgp := path.CreateRootDataNode( schema, "openconfig-bgp:bgp")
    path.CreateDataNode( bgp, "global/config/as", 65172)
    neighbor := path.CreateDataNode( bgp, "neighbors/neighbor[neighbor-address='172.16.255.2']", "")
    path.CreateDataNode( neighbor, "config/neighbor-address", "172.16.255.2")
    path.CreateDataNode( neighbor, "config/peer-as", 65172);   

    setRpc := path.CreateRpc( schema, "ydk:gnmi-set")
    bgpCreatePayload := path.CodecEncode( bgp, encoding.JSON, false)
    path.CreateDataNode( setRpc.Input, "delete[alias='bgp']/entity", bgpCreatePayload)

    session.ExecuteRpc(setRpc);	
}

func getBgpSubscription(gs *path.GnmiSession) {
	previous := ""
	for true {
		response := gs.GetLastSubscribeResponse(previous)
		if len(response) > 0 && response != previous {
			// Do anything with received response
			ydk.YLogDebug(fmt.Sprintf("%s:  Last received subscribe response:\n%s\n", time.Now().Format(time.RFC850), response))
			previous = response
		}
		if !gs.SubscribeInProgress() {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
}

func TestGnmiRpcSubscribePoll(gs *path.GnmiSession, schema types.RootSchemaNode) {
    // Configure BGP
    configureBgp(gs, schema)
    
	// Build subscription
    bgpReadDn := path.CreateRootDataNode( schema, "openconfig-bgp:bgp")
    bgpReadPayload := path.CodecEncode( bgpReadDn, encoding.JSON, false)

    rpc := path.CreateRpc( schema, "ydk:gnmi-subscribe")
    subscription := path.CreateDataNode( rpc.Input, "subscription", "")
    path.CreateDataNode( subscription, "mode", "POLL")
    path.CreateDataNode( subscription, "encoding", "JSON_IETF")

    bgpSubscription := path.CreateDataNode( subscription, "subscription-list[alias='bgp']", "")
    path.CreateDataNode( bgpSubscription, "entity", bgpReadPayload)

    // Send subscription request
    go gs.ExecuteSubscribeRpc(rpc)
    
    // Start subscription response listening function
    time.Sleep(200 * time.Millisecond)
    getBgpSubscription(gs)
    
    // Delete configuration
    deleteConfigBgp(gs, schema)
}

func main () {
	_, callerFile, _, _ := runtime.Caller(0)
	executablePath := filepath.Dir(callerFile)
	repopath := executablePath + "/../../../cpp/core/tests/models"
	repo := types.Repository{Path: repopath}
	gs := path.GnmiSession{
		Repo: repo,
		Address:  "127.0.0.1",
		Port:     50051}
	gs.Connect()
	schema := gs.GetRootSchemaNode()

    argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) > 0 && argsWithoutProg[0] == "-v" {
		ydk.EnableLogging(ydk.Debug)
	}
	
    TestGnmiRpcSubscribePoll(&gs, schema)

    gs.Disconnect()
}
