#!/bin/bash

v8_path="/jp/v8"

export GODEBUG="cgocheck=0"

CGO_LDFLAGS="-I$v8_path/v8 -I$v8_path/v8/include -Wl,--start-group $v8_path/v8d/libv8_base.a $v8_path/v8d/libv8_libbase.a $v8_path/v8d/libv8_external_snapshot.a $v8_path/v8d/libv8_libplatform.a $v8_path/v8d/libv8_libsampler.a $v8_path/v8d/libicuuc.a $v8_path/v8d/libicui18n.a -Wl,--end-group -lrt -ldl -pthread -std=c++0x"  \
CGO_CFLAGS="-I $v8_path/include" \
CGO_CXXFLAGS="-std=c++0x -I $v8_path/include" \
go test -x -run="$1" -bench="$2" -v
