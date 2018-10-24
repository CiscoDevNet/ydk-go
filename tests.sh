#!/bin/bash

function print_msg {
	echo -e "${MSG_COLOR}*** $(date): test.sh | $@ ${NOCOLOR}"
}

# Terminal colors
RED="\033[0;31m"
NOCOLOR="\033[0m"
YELLOW='\033[1;33m'
MSG_COLOR=$YELLOW

if [[ $(uname) == "Linux" ]] && [[ -z "${GOPATH// }" ]]; then
    export GOPATH="$(pwd)/golang"
    print_msg "Setting GOPATH to $GOPATH"
fi

if [[ $(uname) == "Darwin" ]] && [[ -z "${GOPATH// }" ]]; then
    source /Users/travis/.gvm/scripts/gvm
    gvm use go1.9.2
    print_msg "GOROOT: $GOROOT"
    print_msg "GOPATH: $GOPATH"
    print_msg "GO version: $(go version)"
fi

print_msg "Running codec samples"
go run samples/codec/cisco_ios_xr/cdp_cfg/cd_encode_10/cd_encode_10.go -v
go run samples/codec/cisco_ios_xr/cdp_cfg/cd_encode_20/cd_encode_20.go -v
go run samples/codec/cisco_ios_xr/cdp_cfg/cd_encode_22/cd_encode_22.go -v
go run samples/codec/cisco_ios_xr/cdp_cfg/cd_encode_24/cd_encode_24.go -v
go run samples/codec/cisco_ios_xr/aaa_lib_cfg/cd_encode_10/cd_encode_10.go -v
go run samples/codec/cisco_ios_xr/aaa_lib_cfg/cd_encode_20/cd_encode_20.go -v
go run samples/codec/cisco_ios_xr/aaa_lib_cfg/cd_encode_22/cd_encode_22.go -v
go run samples/codec/cisco_ios_xr/clns_isis_cfg/cd_encode_10/cd_encode_10.go -v
go run samples/codec/cisco_ios_xr/clns_isis_cfg/cd_encode_20/cd_encode_20.go -v
#go run samples/codec/cisco_ios_xr/clns_isis_cfg/cd_encode_30/cd_encode_30.go -v
#go run samples/codec/cisco_ios_xr/clns_isis_cfg/cd_encode_31/cd_encode_31.go -v
#go run samples/codec/cisco_ios_xr/clns_isis_cfg/cd_encode_32/cd_encode_32.go -v
#go run samples/codec/cisco_ios_xr/clns_isis_cfg/cd_encode_33/cd_encode_33.go -v
#go run samples/codec/cisco_ios_xr/clns_isis_cfg/cd_encode_34/cd_encode_34.go -v
#go run samples/codec/cisco_ios_xr/clns_isis_cfg/cd_encode_35/cd_encode_35.go -v
#go run samples/codec/cisco_ios_xr/clns_isis_cfg/cd_encode_40/cd_encode_40.go -v
#go run samples/codec/cisco_ios_xr/clns_isis_cfg/cd_encode_41/cd_encode_41.go -v
#go run samples/codec/cisco_ios_xr/clns_isis_cfg/cd_encode_52/cd_encode_52.go -v
#go run samples/codec/cisco_ios_xr/clns_isis_cfg/cd_encode_53/cd_encode_53.go -v
#go run samples/codec/cisco_ios_xr/clns_isis_cfg/cd_encode_54/cd_encode_54.go -v
#go run samples/codec/cisco_ios_xr/clns_isis_cfg/cd_encode_55/cd_encode_55.go -v
#go run samples/codec/cisco_ios_xr/clns_isis_cfg/cd_encode_56/cd_encode_56.go -v
#go run samples/codec/cisco_ios_xr/clns_isis_cfg/cd_encode_57/cd_encode_57.go -v
#go run samples/codec/cisco_ios_xr/clns_isis_cfg/cd_encode_58/cd_encode_58.go -v
#go run samples/codec/cisco_ios_xr/clns_isis_cfg/cd_encode_59/cd_encode_59.go -v
#go run samples/codec/cisco_ios_xr/clns_isis_cfg/cd_encode_60/cd_encode_60.go -v
