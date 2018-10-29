/*
 * This file is part of the KubeVirt project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * Copyright 2017 Red Hat, Inc.
 *
 */

package ignition

import (
	. "github.com/onsi/ginkgo"

	v1 "kubevirt.io/kubevirt/pkg/api/v1"
)

var _ = Describe("Ignition", func() {

	const vmName = "my-vm"
	var vmi *v1.VirtualMachineInstance

	//BeforeEach(func() {
	//	vmi = v1.NewMinimalVMI(vmName)
	//})

	Describe("A new VirtualMachineInstance definition", func() {
		Context("with ignition data", func() {
			vmi = v1.NewMinimalVMI(vmName)
			It("should success", func() {
				data := "{ \"ignition\": { \"config\": {}, \"version\": \"2.2.0\" }, \"networkd\": {}, \"storage\": { \"files\": [ { \"contents\": { \"source\": \"data:,test\", \"verification\": {} }, \"filesystem\": \"root\", \"mode\": 420, \"path\": \"/etc/hostname\" } ] }, \"systemd\": {} }"
				ignition := &v1.Ignition{Data: data}
				vmi.Spec.Domain.Ignition = ignition

			})
		})
	})
})
