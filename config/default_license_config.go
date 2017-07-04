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

// DefaultLicenseConfig default license config based on the Fedora Licensing
// list (https://fedoraproject.org/wiki/Licensing:Main) which in turn is is
// based on the licenses approved by the Free Software Foundation , OSI
// and consultation with Red Hat Legal.
const DefaultLicenseConfig = `
---
# Based on the Fedora license list maintained on the Fedora wiki
# https://fedoraproject.org/wiki/Licensing:Main
#

Apache Software License 2.0:
  short-name: ASL 2.0
  upstream-url: 'http://www.apache.org/licenses/LICENSE-2.0'
  alt-names:
    - The Apache Software License, Version 2.0
    - Apache License, Version 2.0

BSD License (original):
  short-name: BSD with advertising
  upstream-url: 'https://fedoraproject.org/wiki/Licensing/BSD#BSDwithAdvertising'
  alt-names:
    - BSD License
    - The BSD License

BSD License (no advertising):
  short-name: BSD
  upstream-url: 'https://fedoraproject.org/wiki/Licensing/BSD#3ClauseBSD'
  alt-names:
    - 3-Clause BSD License

Common Development Distribution License:
  short-name: CDDL
  upstream-url: 'https://fedoraproject.org/wiki/Licensing/CDDL'

Eclipse Public License 1.0:
  short-name: EPL
  upstream-url: 'http://www.eclipse.org/legal/epl-v10.html'

GNU Lesser General Public License v2 (or 2.1) only:
  short-name: LGPLv2
  upstream-url: 'http://www.gnu.org/licenses/old-licenses/lgpl-2.1.html'
  alt-names:
    - GNU Lesser General Public License v2

GNU General Public License v2.0 only, with Classpath exception:
  short-name: GPLv2 with exceptions
  upsteam-url: 'https://fedoraproject.org/wiki/Licensing/GPL_Classpath_Exception'
  alt-names:
    - GNU General Public License, Version 2 with the Classpath Exception

MIT license:
  short-name: MIT
  upstream-url: 'https://fedoraproject.org/wiki/Licensing/MIT'
  alt-names:
    - MIT license (also X11)
...

`
