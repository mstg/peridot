// Copyright (c) All respective contributors to the Peridot Project. All rights reserved.
// Copyright (c) 2021-2022 Rocky Enterprise Software Foundation, Inc. All rights reserved.
// Copyright (c) 2021-2022 Ctrl IQ, Inc. All rights reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are met:
//
// 1. Redistributions of source code must retain the above copyright notice,
// this list of conditions and the following disclaimer.
//
// 2. Redistributions in binary form must reproduce the above copyright notice,
// this list of conditions and the following disclaimer in the documentation
// and/or other materials provided with the distribution.
//
// 3. Neither the name of the copyright holder nor the names of its contributors
// may be used to endorse or promote products derived from this software without
// specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
// AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
// IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
// ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
// LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
// CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
// SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
// INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
// CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package servicecatalog

import (
	"fmt"
	"os"
	"strings"
)

func SvcName(svc string, protocol string) string {
	env := os.Getenv("BYC_ENV")
	return fmt.Sprintf("%s-%s-%s-service", svc, protocol, env)
}

func SvcNameHttp(svc string) string {
	return SvcName(svc, "http")
}

func SvcNameGrpc(svc string) string {
	return SvcName(svc, "grpc")
}

func Endpoint(svcName string, ns string, port string) string {
	if forceNs := os.Getenv("BYC_FORCE_NS"); forceNs != "" {
		ns = forceNs
	}
	return fmt.Sprintf("%s.%s.svc.cluster.local%s", svcName, ns, port)
}

func EndpointHttp(svcName string, ns string) string {
	//goland:noinspection HttpUrlsUsage
	return fmt.Sprintf("http://%s", Endpoint(svcName, ns, ""))
}

func NS(ns string) string {
	if os.Getenv("BYC_ENV") == "dev" {
		return os.Getenv("BYC_NS")
	}
	return ns
}

func envOverridable(svcName string, protocol string, x func() string) string {
	envName := strings.ToUpper(fmt.Sprintf("%s_%s_ENDPOINT_OVERRIDE", svcName, protocol))
	if env := os.Getenv(envName); env != "" {
		return env
	}
	return x()
}
