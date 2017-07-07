/*
Copyright 2017 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package config

import (
	"github.com/pgier/hawkbuild/util"
	yaml "gopkg.in/yaml.v2"
)

// DefaultLicenseConfig default license config based on the
// and Fedora Licensing
// list (https://fedoraproject.org/wiki/Licensing:Main)
var DefaultLicenseConfig BuildConfig

func init() {
	err := yaml.Unmarshal([]byte(defaultLicenseConfig), &DefaultLicenseConfig)
	util.Check(err)
}

const defaultLicenseConfig = `
# License names and short names are based on the Linux Foundation SPDX license
# list
# https://spdx.org/licenses/
# The alt-names field includes similar names which refer to the same license
# The license-text field provides the SPDX url to the license text
# The upstream-url provides the url to the origin of the license (if available)
---
licenses:
  Apache Software License 2.0:
    short-name: Apache-2.0
    url: 'https://spdx.org/licenses/Apache-2.0.html'
    upstream-url: 'http://www.apache.org/licenses/LICENSE-2.0'
    alt-names:
      - Apache License, Version 2.0
      - The Apache Software License, Version 2.0

  BSD 4-clause "Original" or "Old" License:
    short-name: BSD-4-Clause
    url: 'https://spdx.org/licenses/BSD-4-Clause.html'
    upstream-url: 'http://directory.fsf.org/wiki/License:BSD_4Clause'
    alt-names:
      - BSD with advertising
      - BSD License (original)

  BSD 3-clause "New" or "Revised" License:
    short-name: BSD-3-Clause
    url: 'https://spdx.org/licenses/BSD-3-Clause.html'
    upstream-url: 'http://www.opensource.org/licenses/BSD-3-Clause'

  Common Development and Distribution License 1.0:
    short-name: CDDL-1.0
    url: 'https://spdx.org/licenses/CDDL-1.0.html'
    upstream-url: 'http://www.opensource.org/licenses/cddl1'

  Eclipse Public License 1.0:
    short-name: EPL-1.0
    url: 'https://spdx.org/licenses/EPL-1.0.html'
    upstream-url: 'http://www.eclipse.org/legal/epl-v10.html'

  GNU Lesser General Public License v2.1 only:
    short-name: LGPL-2.1
    url: 'https://spdx.org/licenses/LGPL-2.1.html'
    upstream-url: 'http://www.gnu.org/licenses/old-licenses/lgpl-2.1.html'
    alt-names:
      - GNU Lesser General Public License v2

  GNU General Public License v2.0 only:
    short-name:  	GPL-2.0
    url: 'https://spdx.org/licenses/GPL-2.0.html'
    upstream-url: 'http://www.gnu.org/licenses/old-licenses/gpl-2.0-standalone.html'

  MIT license:
    short-name: MIT
    url: 'https://spdx.org/licenses/MIT.html'
    alt-names:
      - MIT license (also X11)
...

`
