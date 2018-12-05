package main

import (
	"fmt"
	"runtime"
	"time"
	"path/filepath"
	"os"
	ysanity_bgp "github.com/CiscoDevNet/ydk-go/ydk/models/ydktest/openconfig_bgp"
	"github.com/CiscoDevNet/ydk-go/ydk/providers"
	"github.com/CiscoDevNet/ydk-go/ydk/services"
	"github.com/CiscoDevNet/ydk-go/ydk/types"
	"github.com/CiscoDevNet/ydk-go/ydk/types/yfilter"
	"github.com/CiscoDevNet/ydk-go/ydk"
)

func configBgp(provider *providers.GnmiServiceProvider) {
    bgp := ysanity_bgp.Bgp{}
    bgp.Global.Config.As = 65172
    neighbor := ysanity_bgp.Bgp_Neighbors_Neighbor{}
    neighbor.NeighborAddress = "172.16.255.2"
    neighbor.Config.NeighborAddress = "172.16.255.2"
    neighbor.Config.PeerAs = 65172
    bgp.Neighbors.Neighbor = append(bgp.Neighbors.Neighbor, &neighbor)
    bgp.YFilter = yfilter.Replace
    
    service := services.GnmiService{}
	service.Set(provider, &bgp)
}

func deleteBgp(provider *providers.GnmiServiceProvider) {
    filterBgp := ysanity_bgp.Bgp{}
    filterBgp.YFilter = yfilter.Delete
    
    service := services.GnmiService{}
	service.Set(provider, &filterBgp)
}

func bgpSubscription(provider *providers.GnmiServiceProvider) {
	previous := ""
	gs := provider.GetSession()
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

func testSubscribePoll(provider *providers.GnmiServiceProvider) {
    // Build BGP configuration
	configBgp(provider)
	
	// Build subscription request
	sub := services.GnmiSubscription{}
	bgp := ysanity_bgp.Bgp{}
	sub.Entity = &bgp

	var subList []services.GnmiSubscription
	subList = append(subList, sub)
	
	// Start subscription and result monitoring
	service := services.GnmiService{}
	go service.Subscribe(provider, subList, 10, "POLL", "JSON_IETF")
	time.Sleep(100 * time.Millisecond)
	bgpSubscription(provider)
	
    // Delete BGP configuration
    deleteBgp(provider)
}

func main() {
	_, callerFile, _, _ := runtime.Caller(0)
	executablePath := filepath.Dir(callerFile)
	repopath := executablePath + "/../../../cpp/core/tests/models"
	repo := types.Repository{Path: repopath}
	provider := providers.GnmiServiceProvider{
		Repo: repo,
		Address:  "127.0.0.1",
		Port:     50051}
	provider.Connect()
	
    argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) > 0 && argsWithoutProg[0] == "-v" {
		ydk.EnableLogging(ydk.Debug)
	}
	
	testSubscribePoll( &provider)
	provider.Disconnect()
}
