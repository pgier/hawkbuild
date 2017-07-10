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
  {
    "3dfx Glide License": {
        "fedora_abbrev": "Glide",
        "fedora_name": "3dfx Glide License",
        "id": "1",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Glide",
        "spdx_name": "3dfx Glide License",
        "url": "http://www.users.on.net/~triforce/glidexp/COPYING.txt"
    },
    "4Suite Copyright License": {
        "fedora_abbrev": "ASL 1.1",
        "fedora_name": "4Suite Copyright License",
        "id": "2",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "9wm License (Original)": {
        "fedora_abbrev": "",
        "fedora_name": "9wm License (Original)",
        "id": "3",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "AMAP License": {
        "fedora_abbrev": "",
        "fedora_name": "AMAP License",
        "id": "25",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "AMD's plpa_map.c License": {
        "fedora_abbrev": "AMDPLPA",
        "fedora_name": "AMD's plpa_map.c License",
        "id": "28",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "AMDPLPA",
        "spdx_name": "AMD's plpa_map.c License",
        "url": "https://fedoraproject.org/wiki/Licensing/AMD_plpa_map_License"
    },
    "ANTLR Software Rights Notice": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "29",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "ANTLR-PD",
        "spdx_name": "ANTLR Software Rights Notice",
        "url": "http://www.antlr2.org/license.html"
    },
    "AT&T Public License": {
        "fedora_abbrev": "",
        "fedora_name": "AT&T Public License",
        "id": "54",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Abstyles License": {
        "fedora_abbrev": "Abstyles",
        "fedora_name": "Abstyles License",
        "id": "4",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Abstyles",
        "spdx_name": "Abstyles License",
        "url": "https://fedoraproject.org/wiki/Licensing/Abstyles"
    },
    "Academic Free License": {
        "fedora_abbrev": "AFL",
        "fedora_name": "Academic Free License",
        "id": "5",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Academic Free License v1.1": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "6",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "AFL-1.1",
        "spdx_name": "Academic Free License v1.1",
        "url": "http://opensource.linux-mirror.org/licenses/afl-1.1.txt"
    },
    "Academic Free License v1.2": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "7",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "AFL-1.2",
        "spdx_name": "Academic Free License v1.2",
        "url": "http://opensource.linux-mirror.org/licenses/afl-1.2.txt"
    },
    "Academic Free License v2.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "8",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "AFL-2.0",
        "spdx_name": "Academic Free License v2.0",
        "url": "http://opensource.linux-mirror.org/licenses/afl-2.0.txt"
    },
    "Academic Free License v2.1": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "9",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "AFL-2.1",
        "spdx_name": "Academic Free License v2.1",
        "url": "http://opensource.linux-mirror.org/licenses/afl-2.1.txt"
    },
    "Academic Free License v3.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "10",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "AFL-3.0",
        "spdx_name": "Academic Free License v3.0",
        "url": "http://www.rosenlaw.com/AFL3.0.htm"
    },
    "Academy of Motion Picture Arts and Sciences BSD": {
        "fedora_abbrev": "AMPAS BSD",
        "fedora_name": "Academy of Motion Picture Arts and Sciences BSD",
        "id": "11",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "AMPAS",
        "spdx_name": "Academy of Motion Picture Arts and Sciences BSD",
        "url": "https://fedoraproject.org/wiki/Licensing/BSD#AMPASBSD"
    },
    "Adaptive Public License": {
        "fedora_abbrev": "",
        "fedora_name": "Adaptive Public License",
        "id": "12",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Adaptive Public License 1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "13",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "APL-1.0",
        "spdx_name": "Adaptive Public License 1.0",
        "url": "http://www.opensource.org/licenses/APL-1.0"
    },
    "Adobe Glyph List License": {
        "fedora_abbrev": "MIT",
        "fedora_name": "Adobe Glyph List License",
        "id": "14",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Adobe-Glyph",
        "spdx_name": "Adobe Glyph List License",
        "url": "https://fedoraproject.org/wiki/Licensing/MIT#AdobeGlyph"
    },
    "Adobe Postscript AFM License": {
        "fedora_abbrev": "APAFML",
        "fedora_name": "Adobe Postscript AFM License",
        "id": "15",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "APAFML",
        "spdx_name": "Adobe Postscript AFM License",
        "url": "https://fedoraproject.org/wiki/Licensing/AdobePostscriptAFM"
    },
    "Adobe Systems Incorporated Source Code License Agreement": {
        "fedora_abbrev": "Adobe",
        "fedora_name": "Adobe Systems Incorporated Source Code License Agreement",
        "id": "16",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Adobe-2006",
        "spdx_name": "Adobe Systems Incorporated Source Code License Agreement",
        "url": "https://fedoraproject.org/wiki/Licensing/AdobeLicense"
    },
    "Affero General Public License 1.0": {
        "fedora_abbrev": "AGPLv1",
        "fedora_name": "Affero General Public License 1.0",
        "id": "17",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "AGPL-1.0",
        "spdx_name": "Affero General Public License v1.0",
        "url": "http://www.affero.org/oagpl.html"
    },
    "Affero General Public License 3.0": {
        "fedora_abbrev": "AGPLv3",
        "fedora_name": "Affero General Public License 3.0",
        "id": "18",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Affero General Public License 3.0 or later": {
        "fedora_abbrev": "AGPLv3+",
        "fedora_name": "Affero General Public License 3.0 or later",
        "id": "19",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Affero General Public License 3.0 with Zarafa trademark exceptions": {
        "fedora_abbrev": "AGPLv3 with exceptions",
        "fedora_name": "Affero General Public License 3.0 with Zarafa trademark exceptions",
        "id": "20",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Affero General Public License v1.0": {
        "fedora_abbrev": "AGPLv1",
        "fedora_name": "Affero General Public License 1.0",
        "id": "21",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "AGPL-1.0",
        "spdx_name": "Affero General Public License v1.0",
        "url": "http://www.affero.org/oagpl.html"
    },
    "Afmparse License": {
        "fedora_abbrev": "Afmparse",
        "fedora_name": "Afmparse License",
        "id": "22",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Afmparse",
        "spdx_name": "Afmparse License",
        "url": "https://fedoraproject.org/wiki/Licensing/Afmparse"
    },
    "Agere LT Modem Driver License": {
        "fedora_abbrev": "",
        "fedora_name": "Agere LT Modem Driver License",
        "id": "23",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Aladdin Free Public License": {
        "fedora_abbrev": "",
        "fedora_name": "Aladdin Free Public License",
        "id": "24",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "Aladdin",
        "spdx_name": "Aladdin Free Public License",
        "url": "http://pages.cs.wisc.edu/~ghost/doc/AFPL/6.01/Public.htm"
    },
    "Amazon Digital Services License": {
        "fedora_abbrev": "ADSL",
        "fedora_name": "Amazon Digital Services License",
        "id": "26",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "ADSL",
        "spdx_name": "Amazon Digital Services License",
        "url": "https://fedoraproject.org/wiki/Licensing/AmazonDigitalServicesLicense"
    },
    "Amazon Software License": {
        "fedora_abbrev": "",
        "fedora_name": "Amazon Software License",
        "id": "27",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Apache License 1.0": {
        "fedora_abbrev": "ASL 1.0",
        "fedora_name": "Apache Software License 1.0",
        "id": "30",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Apache-1.0",
        "spdx_name": "Apache License 1.0",
        "url": "http://www.apache.org/licenses/LICENSE-1.0"
    },
    "Apache License 1.1": {
        "fedora_abbrev": "ASL 1.1",
        "fedora_name": "Apache Software License 1.1",
        "id": "31",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Apache-1.1",
        "spdx_name": "Apache License 1.1",
        "url": "http://apache.org/licenses/LICENSE-1.1"
    },
    "Apache License 2.0": {
        "fedora_abbrev": "ASL 2.0",
        "fedora_name": "Apache Software License 2.0",
        "id": "32",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Apache-2.0",
        "spdx_name": "Apache License 2.0",
        "url": "http://www.apache.org/licenses/LICENSE-2.0"
    },
    "Apache Software License 1.0": {
        "fedora_abbrev": "ASL 1.0",
        "fedora_name": "Apache Software License 1.0",
        "id": "33",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Apache-1.0",
        "spdx_name": "Apache License 1.0",
        "url": "http://www.apache.org/licenses/LICENSE-1.0"
    },
    "Apache Software License 1.1": {
        "fedora_abbrev": "ASL 1.1",
        "fedora_name": "Apache Software License 1.1",
        "id": "34",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Apache-1.1",
        "spdx_name": "Apache License 1.1",
        "url": "http://apache.org/licenses/LICENSE-1.1"
    },
    "Apache Software License 2.0": {
        "fedora_abbrev": "ASL 2.0",
        "fedora_name": "Apache Software License 2.0",
        "id": "35",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Apache-2.0",
        "spdx_name": "Apache License 2.0",
        "url": "http://www.apache.org/licenses/LICENSE-2.0"
    },
    "Apache Software License, Version 2.0": {
        "fedora_abbrev": "ASL 2.0",
        "fedora_name": "Apache Software License 2.0",
        "id": "36",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Apache-2.0",
        "spdx_name": "Apache License 2.0",
        "url": "http://www.apache.org/licenses/LICENSE-2.0"
    },
    "App::s2p License": {
        "fedora_abbrev": "App-s2p",
        "fedora_name": "App::s2p License",
        "id": "44",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Apple MIT License": {
        "fedora_abbrev": "AML",
        "fedora_name": "Apple MIT License",
        "id": "38",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "AML",
        "spdx_name": "Apple MIT License",
        "url": "https://fedoraproject.org/wiki/Licensing/Apple_MIT_License"
    },
    "Apple Public Source License 1.0": {
        "fedora_abbrev": "",
        "fedora_name": "Apple Public Source License 1.0",
        "id": "39",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "APSL-1.0",
        "spdx_name": "Apple Public Source License 1.0",
        "url": "https://fedoraproject.org/wiki/Licensing/Apple_Public_Source_License_1.0"
    },
    "Apple Public Source License 1.1": {
        "fedora_abbrev": "",
        "fedora_name": "Apple Public Source License 1.1",
        "id": "40",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "APSL-1.1",
        "spdx_name": "Apple Public Source License 1.1",
        "url": "http://www.opensource.apple.com/source/IOSerialFamily/IOSerialFamily-7/APPLE_LICENSE"
    },
    "Apple Public Source License 1.2": {
        "fedora_abbrev": "",
        "fedora_name": "Apple Public Source License 1.2",
        "id": "41",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "APSL-1.2",
        "spdx_name": "Apple Public Source License 1.2",
        "url": "http://www.samurajdata.se/opensource/mirror/licenses/apsl.php"
    },
    "Apple Public Source License 2.0": {
        "fedora_abbrev": "APSL 2.0",
        "fedora_name": "Apple Public Source License 2.0",
        "id": "42",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "APSL-2.0",
        "spdx_name": "Apple Public Source License 2.0",
        "url": "http://www.opensource.apple.com/license/apsl/"
    },
    "Apple Quicktime License": {
        "fedora_abbrev": "",
        "fedora_name": "Apple Quicktime License",
        "id": "43",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Apple iTunes License": {
        "fedora_abbrev": "",
        "fedora_name": "Apple iTunes License",
        "id": "37",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Aptana Public License": {
        "fedora_abbrev": "",
        "fedora_name": "Aptana Public License",
        "id": "45",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Artistic (clarified)": {
        "fedora_abbrev": "Artistic clarified",
        "fedora_name": "Artistic (clarified)",
        "id": "48",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Artistic 1.0 (original)": {
        "fedora_abbrev": "",
        "fedora_name": "Artistic 1.0 (original)",
        "id": "46",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Artistic 2.0": {
        "fedora_abbrev": "Artistic 2.0",
        "fedora_name": "Artistic 2.0",
        "id": "47",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Artistic-2.0",
        "spdx_name": "Artistic License 2.0",
        "url": "http://www.perlfoundation.org/artistic_license_2_0"
    },
    "Artistic License 1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "49",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "Artistic-1.0",
        "spdx_name": "Artistic License 1.0",
        "url": "http://opensource.org/licenses/Artistic-1.0"
    },
    "Artistic License 1.0 (Perl)": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "50",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "Artistic-1.0-Perl",
        "spdx_name": "Artistic License 1.0 (Perl)",
        "url": "http://dev.perl.org/licenses/artistic.html"
    },
    "Artistic License 1.0 w/clause 8": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "51",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "Artistic-1.0-cl8",
        "spdx_name": "Artistic License 1.0 w/clause 8",
        "url": "http://opensource.org/licenses/Artistic-1.0"
    },
    "Artistic License 2.0": {
        "fedora_abbrev": "Artistic 2.0",
        "fedora_name": "Artistic 2.0",
        "id": "52",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Artistic-2.0",
        "spdx_name": "Artistic License 2.0",
        "url": "http://www.perlfoundation.org/artistic_license_2_0"
    },
    "Aspell-ru License": {
        "fedora_abbrev": "ARL",
        "fedora_name": "Aspell-ru License",
        "id": "53",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Attribution Assurance License": {
        "fedora_abbrev": "AAL",
        "fedora_name": "Attribution Assurance License",
        "id": "55",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "AAL",
        "spdx_name": "Attribution Assurance License",
        "url": "http://www.opensource.org/licenses/attribution"
    },
    "BSD 2-clause \"Simplified\" License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "70",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "BSD-2-Clause",
        "spdx_name": "BSD 2-clause \"Simplified\" License",
        "url": "http://www.opensource.org/licenses/BSD-2-Clause"
    },
    "BSD 2-clause FreeBSD License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "68",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "BSD-2-Clause-FreeBSD",
        "spdx_name": "BSD 2-clause FreeBSD License",
        "url": "http://www.freebsd.org/copyright/freebsd-license.html"
    },
    "BSD 2-clause NetBSD License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "69",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "BSD-2-Clause-NetBSD",
        "spdx_name": "BSD 2-clause NetBSD License",
        "url": "http://www.netbsd.org/about/redistribution.html#default"
    },
    "BSD 3-Clause No Nuclear License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "73",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "BSD-3-Clause-No-Nuclear-License",
        "spdx_name": "BSD 3-Clause No Nuclear License",
        "url": "http://download.oracle.com/otn-pub/java/licenses/bsd.txt?AuthParam=1467140197_43d516ce1776bd08a58235a7785be1cc"
    },
    "BSD 3-Clause No Nuclear License 2014": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "74",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "BSD-3-Clause-No-Nuclear-License-2014",
        "spdx_name": "BSD 3-Clause No Nuclear License 2014",
        "url": "https://java.net/projects/javaeetutorial/pages/BerkeleyLicense"
    },
    "BSD 3-Clause No Nuclear Warranty": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "75",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "BSD-3-Clause-No-Nuclear-Warranty",
        "spdx_name": "BSD 3-Clause No Nuclear Warranty",
        "url": "https://jogamp.org/git/?p=gluegen.git;a=blob_plain;f=LICENSE.txt"
    },
    "BSD 3-clause \"New\" or \"Revised\" License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "72",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "BSD-3-Clause",
        "spdx_name": "BSD 3-clause \"New\" or \"Revised\" License",
        "url": "http://www.opensource.org/licenses/BSD-3-Clause"
    },
    "BSD 3-clause Clear License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "71",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "BSD-3-Clause-Clear",
        "spdx_name": "BSD 3-clause Clear License",
        "url": "http://labs.metacarta.com/license-explanation.html#license"
    },
    "BSD 4-clause \"Original\" or \"Old\" License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "76",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "BSD-4-Clause",
        "spdx_name": "BSD 4-clause \"Original\" or \"Old\" License",
        "url": "http://directory.fsf.org/wiki/License:BSD_4Clause"
    },
    "BSD License (no advertising)": {
        "fedora_abbrev": "BSD",
        "fedora_name": "BSD License (no advertising)",
        "id": "78",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "BSD License (original)": {
        "fedora_abbrev": "BSD with advertising",
        "fedora_name": "BSD License (original)",
        "id": "79",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "BSD License (two clause)": {
        "fedora_abbrev": "BSD",
        "fedora_name": "BSD License (two clause)",
        "id": "80",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "BSD Protection License": {
        "fedora_abbrev": "BSD Protection",
        "fedora_name": "BSD Protection License",
        "id": "81",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "BSD-Protection",
        "spdx_name": "BSD Protection License",
        "url": "https://fedoraproject.org/wiki/Licensing/BSD_Protection_License"
    },
    "BSD Source Code Attribution": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "82",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "BSD-Source-Code",
        "spdx_name": "BSD Source Code Attribution",
        "url": "https://github.com/robbiehanson/CocoaHTTPServer/blob/master/LICENSE.txt"
    },
    "BSD Zero Clause License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "84",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "0BSD",
        "spdx_name": "BSD Zero Clause License",
        "url": "http://landley.net/toybox/license.html"
    },
    "BSD with attribution": {
        "fedora_abbrev": "BSD with attribution",
        "fedora_name": "BSD with attribution",
        "id": "83",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "BSD-3-Clause-Attribution",
        "spdx_name": "BSD with attribution",
        "url": "https://fedoraproject.org/wiki/Licensing/BSD_with_Attribution"
    },
    "BSD-4-Clause (University of California-Specific)": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "77",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "BSD-4-Clause-UC",
        "spdx_name": "BSD-4-Clause (University of California-Specific)",
        "url": "http://www.freebsd.org/copyright/license.html"
    },
    "Bahyph License": {
        "fedora_abbrev": "Bahyph",
        "fedora_name": "Bahyph License",
        "id": "56",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Bahyph",
        "spdx_name": "Bahyph License",
        "url": "https://fedoraproject.org/wiki/Licensing/Bahyph"
    },
    "Barr License": {
        "fedora_abbrev": "Barr",
        "fedora_name": "Barr License",
        "id": "57",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Barr",
        "spdx_name": "Barr License",
        "url": "https://fedoraproject.org/wiki/Licensing/Barr"
    },
    "BeOpen Open Source License Agreement Version 1": {
        "fedora_abbrev": "BeOpen",
        "fedora_name": "BeOpen Open Source License Agreement Version 1",
        "id": "59",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Beerware License": {
        "fedora_abbrev": "Beerware",
        "fedora_name": "Beerware License",
        "id": "58",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Beerware",
        "spdx_name": "Beerware License",
        "url": "https://fedoraproject.org/wiki/Licensing/Beerware"
    },
    "Bibtex License": {
        "fedora_abbrev": "Bibtex",
        "fedora_name": "Bibtex License",
        "id": "60",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "BitTorrent License": {
        "fedora_abbrev": "BitTorrent",
        "fedora_name": "BitTorrent License",
        "id": "61",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "BitTorrent Open Source License v1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "62",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "BitTorrent-1.0",
        "spdx_name": "BitTorrent Open Source License v1.0",
        "url": "http://sources.gentoo.org/cgi-bin/viewvc.cgi/gentoo-x86/licenses/BitTorrent?r1=1.1&r2=1.1.1.1&diff_format=s"
    },
    "BitTorrent Open Source License v1.1": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "63",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "BitTorrent-1.1",
        "spdx_name": "BitTorrent Open Source License v1.1",
        "url": "http://directory.fsf.org/wiki/License:BitTorrentOSL1.1"
    },
    "Boost Software License": {
        "fedora_abbrev": "Boost",
        "fedora_name": "Boost Software License",
        "id": "64",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Boost Software License 1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "65",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "BSL-1.0",
        "spdx_name": "Boost Software License 1.0",
        "url": "http://www.boost.org/LICENSE_1_0.txt"
    },
    "Borceux license": {
        "fedora_abbrev": "Borceux",
        "fedora_name": "Borceux license",
        "id": "66",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Borceux",
        "spdx_name": "Borceux license",
        "url": "https://fedoraproject.org/wiki/Licensing/Borceux"
    },
    "Bouncy Castle Licence": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "67",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": "http://www.bouncycastle.org/licence.html"
    },
    "C/Migemo License": {
        "fedora_abbrev": "",
        "fedora_name": "C/Migemo License",
        "id": "101",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "CACert Root Distribution License": {
        "fedora_abbrev": "",
        "fedora_name": "CACert Root Distribution License",
        "id": "87",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "CMU License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "102",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "MIT-CMU",
        "spdx_name": "CMU License",
        "url": "https://fedoraproject.org/wiki/Licensing:MIT?rd=Licensing/MIT#CMU_Style"
    },
    "CMU License (BSD like)": {
        "fedora_abbrev": "MIT",
        "fedora_name": "CMU License (BSD like)",
        "id": "103",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "CNRI Jython License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "104",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CNRI-Jython",
        "spdx_name": "CNRI Jython License",
        "url": "http://www.jython.org/license.html"
    },
    "CNRI License (Old Python)": {
        "fedora_abbrev": "CNRI",
        "fedora_name": "CNRI License (Old Python)",
        "id": "105",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "CNRI Python License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "106",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CNRI-Python",
        "spdx_name": "CNRI Python License",
        "url": "http://www.opensource.org/licenses/CNRI-Python"
    },
    "CNRI Python Open Source GPL Compatible License Agreement": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "107",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CNRI-Python-GPL-Compatible",
        "spdx_name": "CNRI Python Open Source GPL Compatible License Agreement",
        "url": "http://www.python.org/download/releases/1.6.1/download_win/"
    },
    "CPAL License 1.0": {
        "fedora_abbrev": "CPAL",
        "fedora_name": "CPAL License 1.0",
        "id": "123",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "CRC32 License": {
        "fedora_abbrev": "CRC32",
        "fedora_name": "CRC32 License",
        "id": "124",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "CUA Office Public License Version 1.0": {
        "fedora_abbrev": "MPLv1.1",
        "fedora_name": "CUA Office Public License Version 1.0",
        "id": "163",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "CUA-OPL-1.0",
        "spdx_name": "CUA Office Public License v1.0",
        "url": "http://opensource.org/licenses/CUA-OPL-1.0"
    },
    "Caldera License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "88",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "Caldera",
        "spdx_name": "Caldera License",
        "url": "http://www.lemis.com/grog/UNIX/ancient-source-all.pdf"
    },
    "CeCILL Free Software License Agreement v1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "93",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CECILL-1.0",
        "spdx_name": "CeCILL Free Software License Agreement v1.0",
        "url": "http://www.cecill.info/licences/Licence_CeCILL_V1-fr.html"
    },
    "CeCILL Free Software License Agreement v1.1": {
        "fedora_abbrev": "CeCILL",
        "fedora_name": "CeCILL License v1.1",
        "id": "94",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "CECILL-1.1",
        "spdx_name": "CeCILL Free Software License Agreement v1.1",
        "url": "http://www.cecill.info/licences/Licence_CeCILL_V1.1-US.html"
    },
    "CeCILL Free Software License Agreement v2.0": {
        "fedora_abbrev": "CeCILL",
        "fedora_name": "CeCILL License v2",
        "id": "95",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "CECILL-2.0",
        "spdx_name": "CeCILL Free Software License Agreement v2.0",
        "url": "http://www.cecill.info/licences/Licence_CeCILL_V2-fr.html"
    },
    "CeCILL Free Software License Agreement v2.1": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "96",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CECILL-2.1",
        "spdx_name": "CeCILL Free Software License Agreement v2.1",
        "url": "http://www.cecill.info/licences/Licence_CeCILL_V2.1-fr.html"
    },
    "CeCILL License v1.1": {
        "fedora_abbrev": "CeCILL",
        "fedora_name": "CeCILL License v1.1",
        "id": "97",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "CECILL-1.1",
        "spdx_name": "CeCILL Free Software License Agreement v1.1",
        "url": "http://www.cecill.info/licences/Licence_CeCILL_V1.1-US.html"
    },
    "CeCILL License v2": {
        "fedora_abbrev": "CeCILL",
        "fedora_name": "CeCILL License v2",
        "id": "98",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "CECILL-2.0",
        "spdx_name": "CeCILL Free Software License Agreement v2.0",
        "url": "http://www.cecill.info/licences/Licence_CeCILL_V2-fr.html"
    },
    "CeCILL-B Free Software License Agreement": {
        "fedora_abbrev": "CeCILL-B",
        "fedora_name": "CeCILL-B License",
        "id": "89",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "CECILL-B",
        "spdx_name": "CeCILL-B Free Software License Agreement",
        "url": "http://www.cecill.info/licences/Licence_CeCILL-B_V1-fr.html"
    },
    "CeCILL-B License": {
        "fedora_abbrev": "CeCILL-B",
        "fedora_name": "CeCILL-B License",
        "id": "90",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "CECILL-B",
        "spdx_name": "CeCILL-B Free Software License Agreement",
        "url": "http://www.cecill.info/licences/Licence_CeCILL-B_V1-fr.html"
    },
    "CeCILL-C Free Software License Agreement": {
        "fedora_abbrev": "CeCILL-C",
        "fedora_name": "CeCILL-C License",
        "id": "91",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "CECILL-C",
        "spdx_name": "CeCILL-C Free Software License Agreement",
        "url": "http://www.cecill.info/licences/Licence_CeCILL-C_V1-fr.html"
    },
    "CeCILL-C License": {
        "fedora_abbrev": "CeCILL-C",
        "fedora_name": "CeCILL-C License",
        "id": "92",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "CECILL-C",
        "spdx_name": "CeCILL-C Free Software License Agreement",
        "url": "http://www.cecill.info/licences/Licence_CeCILL-C_V1-fr.html"
    },
    "Celtx Public License (CePL)": {
        "fedora_abbrev": "Netscape",
        "fedora_name": "Celtx Public License (CePL)",
        "id": "99",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Clarified Artistic License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "100",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "ClArtistic",
        "spdx_name": "Clarified Artistic License",
        "url": "http://www.ncftp.com/ncftp/doc/LICENSE.txt"
    },
    "Code Project Open License 1.02": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "108",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CPOL-1.02",
        "spdx_name": "Code Project Open License 1.02",
        "url": "http://www.codeproject.com/info/cpol10.aspx"
    },
    "CodeProject Open License (CPOL)": {
        "fedora_abbrev": "",
        "fedora_name": "CodeProject Open License (CPOL)",
        "id": "109",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Common Development Distribution License": {
        "fedora_abbrev": "CDDL",
        "fedora_name": "Common Development Distribution License",
        "id": "114",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Common Development and Distribution License": {
        "fedora_abbrev": "CDDL",
        "fedora_name": "Common Development Distribution License",
        "id": "110",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": "http://repository.jboss.org/licenses/cddl.txt"
    },
    "Common Development and Distribution License 1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "111",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CDDL-1.0",
        "spdx_name": "Common Development and Distribution License 1.0",
        "url": "http://www.opensource.org/licenses/cddl1"
    },
    "Common Development and Distribution License 1.1": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "112",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "CDDL-1.1",
        "spdx_name": "Common Development and Distribution License 1.1",
        "url": "http://glassfish.java.net/public/CDDL+GPL_1_1.html"
    },
    "Common Development and Distribution License version 1.1": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "113",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "CDDL-1.1",
        "spdx_name": "Common Development and Distribution License 1.1",
        "url": "http://glassfish.java.net/public/CDDL+GPL_1_1.html"
    },
    "Common Public Attribution License 1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "115",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CPAL-1.0",
        "spdx_name": "Common Public Attribution License 1.0",
        "url": "http://www.opensource.org/licenses/CPAL-1.0"
    },
    "Common Public License": {
        "fedora_abbrev": "CPL",
        "fedora_name": "Common Public License",
        "id": "116",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Common Public License 1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "117",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "CPL-1.0",
        "spdx_name": "Common Public License 1.0",
        "url": "http://opensource.org/licenses/CPL-1.0"
    },
    "Common Public License, Version 1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "118",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "CPL-1.0",
        "spdx_name": "Common Public License 1.0",
        "url": "http://opensource.org/licenses/CPL-1.0"
    },
    "Computer Associates Trusted Open Source License 1.1": {
        "fedora_abbrev": "CATOSL",
        "fedora_name": "Computer Associates Trusted Open Source License 1.1",
        "id": "119",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "CATOSL-1.1",
        "spdx_name": "Computer Associates Trusted Open Source License 1.1",
        "url": "http://opensource.org/licenses/CATOSL-1.1"
    },
    "Condor Public License": {
        "fedora_abbrev": "Condor",
        "fedora_name": "Condor Public License",
        "id": "120",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Condor Public License v1.1": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "121",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "Condor-1.1",
        "spdx_name": "Condor Public License v1.1",
        "url": "http://research.cs.wisc.edu/condor/license.html#condor"
    },
    "Copyright Attribution Only": {
        "fedora_abbrev": "Copyright only",
        "fedora_name": "Copyright Attribution Only",
        "id": "122",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Creative Commons Attribution 1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "125",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-1.0",
        "spdx_name": "Creative Commons Attribution 1.0",
        "url": "http://creativecommons.org/licenses/by/1.0/legalcode"
    },
    "Creative Commons Attribution 2.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "126",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-2.0",
        "spdx_name": "Creative Commons Attribution 2.0",
        "url": "http://creativecommons.org/licenses/by/2.0/legalcode"
    },
    "Creative Commons Attribution 2.5": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "127",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-2.5",
        "spdx_name": "Creative Commons Attribution 2.5",
        "url": "http://creativecommons.org/licenses/by/2.5/legalcode"
    },
    "Creative Commons Attribution 3.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "128",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-3.0",
        "spdx_name": "Creative Commons Attribution 3.0",
        "url": "http://creativecommons.org/licenses/by/3.0/legalcode"
    },
    "Creative Commons Attribution 4.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "129",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-4.0",
        "spdx_name": "Creative Commons Attribution 4.0",
        "url": "http://creativecommons.org/licenses/by/4.0/legalcode"
    },
    "Creative Commons Attribution No Derivatives 1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "130",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-ND-1.0",
        "spdx_name": "Creative Commons Attribution No Derivatives 1.0",
        "url": "http://creativecommons.org/licenses/by-nd/1.0/legalcode"
    },
    "Creative Commons Attribution No Derivatives 2.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "131",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-ND-2.0",
        "spdx_name": "Creative Commons Attribution No Derivatives 2.0",
        "url": "http://creativecommons.org/licenses/by-nd/2.0/legalcode"
    },
    "Creative Commons Attribution No Derivatives 2.5": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "132",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-ND-2.5",
        "spdx_name": "Creative Commons Attribution No Derivatives 2.5",
        "url": "http://creativecommons.org/licenses/by-nd/2.5/legalcode"
    },
    "Creative Commons Attribution No Derivatives 3.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "133",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-ND-3.0",
        "spdx_name": "Creative Commons Attribution No Derivatives 3.0",
        "url": "http://creativecommons.org/licenses/by-nd/3.0/legalcode"
    },
    "Creative Commons Attribution No Derivatives 4.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "134",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-ND-4.0",
        "spdx_name": "Creative Commons Attribution No Derivatives 4.0",
        "url": "http://creativecommons.org/licenses/by-nd/4.0/legalcode"
    },
    "Creative Commons Attribution Non Commercial 1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "135",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-NC-1.0",
        "spdx_name": "Creative Commons Attribution Non Commercial 1.0",
        "url": "http://creativecommons.org/licenses/by-nc/1.0/legalcode"
    },
    "Creative Commons Attribution Non Commercial 2.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "136",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-NC-2.0",
        "spdx_name": "Creative Commons Attribution Non Commercial 2.0",
        "url": "http://creativecommons.org/licenses/by-nc/2.0/legalcode"
    },
    "Creative Commons Attribution Non Commercial 2.5": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "137",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-NC-2.5",
        "spdx_name": "Creative Commons Attribution Non Commercial 2.5",
        "url": "http://creativecommons.org/licenses/by-nc/2.5/legalcode"
    },
    "Creative Commons Attribution Non Commercial 3.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "138",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-NC-3.0",
        "spdx_name": "Creative Commons Attribution Non Commercial 3.0",
        "url": "http://creativecommons.org/licenses/by-nc/3.0/legalcode"
    },
    "Creative Commons Attribution Non Commercial 4.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "139",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-NC-4.0",
        "spdx_name": "Creative Commons Attribution Non Commercial 4.0",
        "url": "http://creativecommons.org/licenses/by-nc/4.0/legalcode"
    },
    "Creative Commons Attribution Non Commercial No Derivatives 1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "140",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-NC-ND-1.0",
        "spdx_name": "Creative Commons Attribution Non Commercial No Derivatives 1.0",
        "url": "http://creativecommons.org/licenses/by-nd-nc/1.0/legalcode"
    },
    "Creative Commons Attribution Non Commercial No Derivatives 2.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "141",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-NC-ND-2.0",
        "spdx_name": "Creative Commons Attribution Non Commercial No Derivatives 2.0",
        "url": "http://creativecommons.org/licenses/by-nc-nd/2.0/legalcode"
    },
    "Creative Commons Attribution Non Commercial No Derivatives 2.5": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "142",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-NC-ND-2.5",
        "spdx_name": "Creative Commons Attribution Non Commercial No Derivatives 2.5",
        "url": "http://creativecommons.org/licenses/by-nc-nd/2.5/legalcode"
    },
    "Creative Commons Attribution Non Commercial No Derivatives 3.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "143",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-NC-ND-3.0",
        "spdx_name": "Creative Commons Attribution Non Commercial No Derivatives 3.0",
        "url": "http://creativecommons.org/licenses/by-nc-nd/3.0/legalcode"
    },
    "Creative Commons Attribution Non Commercial No Derivatives 4.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "144",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-NC-ND-4.0",
        "spdx_name": "Creative Commons Attribution Non Commercial No Derivatives 4.0",
        "url": "http://creativecommons.org/licenses/by-nc-nd/4.0/legalcode"
    },
    "Creative Commons Attribution Non Commercial Share Alike 1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "145",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-NC-SA-1.0",
        "spdx_name": "Creative Commons Attribution Non Commercial Share Alike 1.0",
        "url": "http://creativecommons.org/licenses/by-nc-sa/1.0/legalcode"
    },
    "Creative Commons Attribution Non Commercial Share Alike 2.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "146",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-NC-SA-2.0",
        "spdx_name": "Creative Commons Attribution Non Commercial Share Alike 2.0",
        "url": "http://creativecommons.org/licenses/by-nc-sa/2.0/legalcode"
    },
    "Creative Commons Attribution Non Commercial Share Alike 2.5": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "147",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-NC-SA-2.5",
        "spdx_name": "Creative Commons Attribution Non Commercial Share Alike 2.5",
        "url": "http://creativecommons.org/licenses/by-nc-sa/2.5/legalcode"
    },
    "Creative Commons Attribution Non Commercial Share Alike 3.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "148",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-NC-SA-3.0",
        "spdx_name": "Creative Commons Attribution Non Commercial Share Alike 3.0",
        "url": "http://creativecommons.org/licenses/by-nc-sa/3.0/legalcode"
    },
    "Creative Commons Attribution Non Commercial Share Alike 4.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "149",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-NC-SA-4.0",
        "spdx_name": "Creative Commons Attribution Non Commercial Share Alike 4.0",
        "url": "http://creativecommons.org/licenses/by-nc-sa/4.0/legalcode"
    },
    "Creative Commons Attribution Share Alike 1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "150",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-SA-1.0",
        "spdx_name": "Creative Commons Attribution Share Alike 1.0",
        "url": "http://creativecommons.org/licenses/by-sa/1.0/legalcode"
    },
    "Creative Commons Attribution Share Alike 2.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "151",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-SA-2.0",
        "spdx_name": "Creative Commons Attribution Share Alike 2.0",
        "url": "http://creativecommons.org/licenses/by-sa/2.0/legalcode"
    },
    "Creative Commons Attribution Share Alike 2.5": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "152",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-SA-2.5",
        "spdx_name": "Creative Commons Attribution Share Alike 2.5",
        "url": "http://creativecommons.org/licenses/by-sa/2.5/legalcode"
    },
    "Creative Commons Attribution Share Alike 3.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "153",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-SA-3.0",
        "spdx_name": "Creative Commons Attribution Share Alike 3.0",
        "url": "http://creativecommons.org/licenses/by-sa/3.0/legalcode"
    },
    "Creative Commons Attribution Share Alike 4.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "154",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "CC-BY-SA-4.0",
        "spdx_name": "Creative Commons Attribution Share Alike 4.0",
        "url": "http://creativecommons.org/licenses/by-sa/4.0/legalcode"
    },
    "Creative Commons BSD": {
        "fedora_abbrev": "BSD",
        "fedora_name": "Creative Commons BSD",
        "id": "155",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Creative Commons GNU GPL": {
        "fedora_abbrev": "GPLv2+",
        "fedora_name": "Creative Commons GNU GPL",
        "id": "156",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Creative Commons GNU LGPL": {
        "fedora_abbrev": "LGPLv2+",
        "fedora_name": "Creative Commons GNU LGPL",
        "id": "157",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Creative Commons Zero 1.0 Universal": {
        "fedora_abbrev": "CC0",
        "fedora_name": "Creative Commons Zero 1.0 Universal",
        "id": "158",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "CC0-1.0",
        "spdx_name": "Creative Commons Zero v1.0 Universal",
        "url": "http://creativecommons.org/publicdomain/zero/1.0/legalcode"
    },
    "Creative Commons Zero v1.0 Universal": {
        "fedora_abbrev": "CC0",
        "fedora_name": "Creative Commons Zero 1.0 Universal",
        "id": "159",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "CC0-1.0",
        "spdx_name": "Creative Commons Zero v1.0 Universal",
        "url": "http://creativecommons.org/publicdomain/zero/1.0/legalcode"
    },
    "Crossword License": {
        "fedora_abbrev": "Crossword",
        "fedora_name": "Crossword License",
        "id": "160",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Crossword",
        "spdx_name": "Crossword License",
        "url": "https://fedoraproject.org/wiki/Licensing/Crossword"
    },
    "Cryptix General License": {
        "fedora_abbrev": "BSD",
        "fedora_name": "Cryptix General License",
        "id": "161",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Crystal Stacker License": {
        "fedora_abbrev": "Crystal Stacker",
        "fedora_name": "Crystal Stacker License",
        "id": "162",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "CrystalStacker",
        "spdx_name": "CrystalStacker License",
        "url": "https://fedoraproject.org/wiki/Licensing:CrystalStacker?rd=Licensing/CrystalStacker"
    },
    "Cube License": {
        "fedora_abbrev": "Cube",
        "fedora_name": "Cube License",
        "id": "164",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Cube",
        "spdx_name": "Cube License",
        "url": "https://fedoraproject.org/wiki/Licensing/Cube"
    },
    "DO WHATEVER PUBLIC LICENSE": {
        "fedora_abbrev": "DWPL",
        "fedora_name": "DO WHATEVER PUBLIC LICENSE",
        "id": "171",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "DOC License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "169",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "",
        "spdx_name": "DOC License",
        "url": "http://www.cs.wustl.edu/~schmidt/ACE-copying.html"
    },
    "DSDP License": {
        "fedora_abbrev": "DSDP",
        "fedora_name": "DSDP License",
        "id": "173",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "DSDP",
        "spdx_name": "DSDP License",
        "url": "https://fedoraproject.org/wiki/Licensing/DSDP"
    },
    "Deutsche Freie Software Lizenz": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "166",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "D-FSL-1.0",
        "spdx_name": "Deutsche Freie Software Lizenz",
        "url": "http://www.dipp.nrw.de/d-fsl/lizenzen/"
    },
    "Do What The F*ck You Want To Public License": {
        "fedora_abbrev": "WTFPL",
        "fedora_name": "Do What The F*ck You Want To Public License",
        "id": "172",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "WTFPL",
        "spdx_name": "Do What The F*ck You Want To Public License",
        "url": "http://sam.zoy.org/wtfpl/COPYING"
    },
    "Docbook MIT License": {
        "fedora_abbrev": "DMIT",
        "fedora_name": "Docbook MIT License",
        "id": "168",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Dotseqn License": {
        "fedora_abbrev": "Dotseqn",
        "fedora_name": "Dotseqn License",
        "id": "170",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Dotseqn",
        "spdx_name": "Dotseqn License",
        "url": "https://fedoraproject.org/wiki/Licensing/Dotseqn"
    },
    "EMC2 License": {
        "fedora_abbrev": "",
        "fedora_name": "EMC2 License",
        "id": "186",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "EPICS Open License": {
        "fedora_abbrev": "EPICS",
        "fedora_name": "EPICS Open License",
        "id": "191",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "EU Datagrid Software License": {
        "fedora_abbrev": "EU Datagrid",
        "fedora_name": "EU Datagrid Software License",
        "id": "193",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "EUDatagrid",
        "spdx_name": "EU DataGrid Software License",
        "url": "http://eu-datagrid.web.cern.ch/eu-datagrid/license.html"
    },
    "Eclipse Distribution License 1.0": {
        "fedora_abbrev": "BSD",
        "fedora_name": "Eclipse Distribution License 1.0",
        "id": "175",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Eclipse Distribution License, Version 1.0": {
        "fedora_abbrev": "BSD",
        "fedora_name": "Eclipse Distribution License 1.0",
        "id": "176",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": "http://repository.jboss.org/licenses/edl-1.0.txt"
    },
    "Eclipse Public License 1.0": {
        "fedora_abbrev": "EPL",
        "fedora_name": "Eclipse Public License 1.0",
        "id": "177",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "EPL-1.0",
        "spdx_name": "Eclipse Public License 1.0",
        "url": "http://www.eclipse.org/legal/epl-v10.html"
    },
    "Eclipse Public License, Version 1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "178",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": "http://repository.jboss.org/licenses/epl-1.0.txt"
    },
    "Educational Community License 1.0": {
        "fedora_abbrev": "ECL 1.0",
        "fedora_name": "Educational Community License 1.0",
        "id": "181",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "ECL-1.0",
        "spdx_name": "Educational Community License v1.0",
        "url": "http://opensource.org/licenses/ECL-1.0"
    },
    "Educational Community License 2.0": {
        "fedora_abbrev": "ECL 2.0",
        "fedora_name": "Educational Community License 2.0",
        "id": "182",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "ECL-2.0",
        "spdx_name": "Educational Community License v2.0",
        "url": "http://opensource.org/licenses/ECL-2.0"
    },
    "Eiffel Forum License 1.0": {
        "fedora_abbrev": "",
        "fedora_name": "Eiffel Forum License 1.0",
        "id": "184",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "EFL-1.0",
        "spdx_name": "Eiffel Forum License v1.0",
        "url": "http://www.eiffel-nice.org/license/forum.txt"
    },
    "Eiffel Forum License 2.0": {
        "fedora_abbrev": "EFL 2.0",
        "fedora_name": "Eiffel Forum License 2.0",
        "id": "185",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "EFL-2.0",
        "spdx_name": "Eiffel Forum License v2.0",
        "url": "http://www.eiffel-nice.org/license/eiffel-forum-license-2.html"
    },
    "Enlightenment License (e16)": {
        "fedora_abbrev": "MIT with advertising",
        "fedora_name": "Enlightenment License (e16)",
        "id": "187",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "MIT-advertising",
        "spdx_name": "Enlightenment License (e16)",
        "url": "https://fedoraproject.org/wiki/Licensing/MIT_With_Advertising"
    },
    "Entessa Public License": {
        "fedora_abbrev": "Entessa",
        "fedora_name": "Entessa Public License",
        "id": "189",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Entessa Public License v1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "190",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "Entessa",
        "spdx_name": "Entessa Public License v1.0",
        "url": "http://opensource.org/licenses/Entessa"
    },
    "Erlang Public License 1.1": {
        "fedora_abbrev": "ERPL",
        "fedora_name": "Erlang Public License 1.1",
        "id": "192",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "ErlPL-1.1",
        "spdx_name": "Erlang Public License v1.1",
        "url": "http://www.erlang.org/EPLICENSE"
    },
    "European Union Public License 1.0": {
        "fedora_abbrev": "",
        "fedora_name": "European Union Public License v1.0",
        "id": "194",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "EUPL-1.0",
        "spdx_name": "European Union Public License 1.0",
        "url": "http://ec.europa.eu/idabc/en/document/7330.html"
    },
    "European Union Public License 1.1": {
        "fedora_abbrev": "EUPL 1.1",
        "fedora_name": "European Union Public License 1.1",
        "id": "195",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "EUPL-1.1",
        "spdx_name": "European Union Public License 1.1",
        "url": "https://joinup.ec.europa.eu/software/page/eupl/licence-eupl"
    },
    "European Union Public License v1.0": {
        "fedora_abbrev": "",
        "fedora_name": "European Union Public License v1.0",
        "id": "196",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "EUPL-1.0",
        "spdx_name": "European Union Public License 1.0",
        "url": "http://ec.europa.eu/idabc/en/document/7330.html"
    },
    "Eurosym License": {
        "fedora_abbrev": "Eurosym",
        "fedora_name": "Eurosym License",
        "id": "197",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Eurosym",
        "spdx_name": "Eurosym License",
        "url": "https://fedoraproject.org/wiki/Licensing/Eurosym"
    },
    "FLTK License": {
        "fedora_abbrev": "LGPLv2 with exceptions",
        "fedora_name": "FLTK License",
        "id": "201",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "FSF All Permissive License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "209",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "FSFAP",
        "spdx_name": "FSF All Permissive License",
        "url": "http://www.gnu.org/prep/maintain/html_node/License-Notices-for-Other-Files.html"
    },
    "FSF Unlimited License": {
        "fedora_abbrev": "FSFUL",
        "fedora_name": "FSF Unlimited License",
        "id": "210",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "FSFUL",
        "spdx_name": "FSF Unlimited License",
        "url": "https://fedoraproject.org/wiki/Licensing/FSF_Unlimited_License"
    },
    "FSF Unlimited License (with License Retention)": {
        "fedora_abbrev": "FSFULLR",
        "fedora_name": "FSF Unlimited License (with License Retention)",
        "id": "211",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "FSFULLR",
        "spdx_name": "FSF Unlimited License (with License Retention)",
        "url": "https://fedoraproject.org/wiki/Licensing/FSF_Unlimited_License#License_Retention_Variant"
    },
    "Fair License": {
        "fedora_abbrev": "Fair",
        "fedora_name": "Fair License",
        "id": "198",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Fair",
        "spdx_name": "Fair License",
        "url": "http://www.opensource.org/licenses/Fair"
    },
    "Fedora Directory Server License": {
        "fedora_abbrev": "GPLv2 with exceptions",
        "fedora_name": "Fedora Directory Server License",
        "id": "199",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Frameworx License": {
        "fedora_abbrev": "",
        "fedora_name": "Frameworx License",
        "id": "202",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Frameworx Open License 1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "203",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "Frameworx-1.0",
        "spdx_name": "Frameworx Open License 1.0",
        "url": "http://www.opensource.org/licenses/Frameworx-1.0"
    },
    "FreeImage Public License": {
        "fedora_abbrev": "MPLv1.0",
        "fedora_name": "FreeImage Public License",
        "id": "204",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "FreeImage Public License v1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "205",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "FreeImage",
        "spdx_name": "FreeImage Public License v1.0",
        "url": "http://freeimage.sourceforge.net/freeimage-license.txt"
    },
    "Freetype License": {
        "fedora_abbrev": "FTL",
        "fedora_name": "Freetype License",
        "id": "206",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Freetype Project License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "207",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "FTL",
        "spdx_name": "Freetype Project License",
        "url": "http://freetype.fis.uniroma2.it/FTL.TXT"
    },
    "Frontier Artistic License": {
        "fedora_abbrev": "",
        "fedora_name": "Frontier Artistic License",
        "id": "208",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "GL2PS License": {
        "fedora_abbrev": "GL2PS",
        "fedora_name": "GL2PS License",
        "id": "213",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "GL2PS",
        "spdx_name": "GL2PS License",
        "url": "http://www.geuz.org/gl2ps/COPYING.GL2PS"
    },
    "GNU Affero General Public License v3.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "215",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "AGPL-3.0",
        "spdx_name": "GNU Affero General Public License v3.0",
        "url": "http://www.gnu.org/licenses/agpl.txt"
    },
    "GNU Free Documentation License v1.1": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "216",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "GFDL-1.1",
        "spdx_name": "GNU Free Documentation License v1.1",
        "url": "http://www.gnu.org/licenses/old-licenses/fdl-1.1.txt"
    },
    "GNU Free Documentation License v1.2": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "217",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "GFDL-1.2",
        "spdx_name": "GNU Free Documentation License v1.2",
        "url": "http://www.gnu.org/licenses/old-licenses/fdl-1.2.txt"
    },
    "GNU Free Documentation License v1.3": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "218",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "GFDL-1.3",
        "spdx_name": "GNU Free Documentation License v1.3",
        "url": "http://www.gnu.org/licenses/fdl-1.3.txt"
    },
    "GNU General Public License (no version)": {
        "fedora_abbrev": "GPL+",
        "fedora_name": "GNU General Public License (no version)",
        "id": "219",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "GNU General Public License (no version), with Classpath exception": {
        "fedora_abbrev": "GPL+ with exceptions",
        "fedora_name": "GNU General Public License (no version), with Classpath exception",
        "id": "220",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "GNU General Public License (no version), with font embedding exception": {
        "fedora_abbrev": "GPL+ with exceptions",
        "fedora_name": "GNU General Public License (no version), with font embedding exception",
        "id": "221",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "GNU General Public License v1.0 only": {
        "fedora_abbrev": "GPLv1",
        "fedora_name": "GNU General Public License v1.0 only",
        "id": "222",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "GPL-1.0",
        "spdx_name": "GNU General Public License v1.0 only",
        "url": "http://www.gnu.org/licenses/old-licenses/gpl-1.0-standalone.html"
    },
    "GNU General Public License v1.0 or later": {
        "fedora_abbrev": "GPL+",
        "fedora_name": "GNU General Public License v1.0 or later",
        "id": "223",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "GNU General Public License v2.0 only": {
        "fedora_abbrev": "GPLv2",
        "fedora_name": "GNU General Public License v2.0 only",
        "id": "224",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "GPL-2.0",
        "spdx_name": "GNU General Public License v2.0 only",
        "url": "http://www.gnu.org/licenses/old-licenses/gpl-2.0-standalone.html"
    },
    "GNU General Public License v2.0 only, with Classpath exception": {
        "fedora_abbrev": "GPLv2 with exceptions",
        "fedora_name": "GNU General Public License v2.0 only, with Classpath exception",
        "id": "225",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "GNU General Public License v2.0 only, with font embedding exception": {
        "fedora_abbrev": "GPLv2 with exceptions",
        "fedora_name": "GNU General Public License v2.0 only, with font embedding exception",
        "id": "226",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "GNU General Public License v2.0 or later": {
        "fedora_abbrev": "GPLv2+",
        "fedora_name": "GNU General Public License v2.0 or later",
        "id": "227",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "GNU General Public License v2.0 or later, with Classpath exception": {
        "fedora_abbrev": "GPLv2+ with exceptions",
        "fedora_name": "GNU General Public License v2.0 or later, with Classpath exception",
        "id": "228",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "GNU General Public License v2.0 or later, with font embedding exception": {
        "fedora_abbrev": "GPLv2+ with exceptions",
        "fedora_name": "GNU General Public License v2.0 or later, with font embedding exception",
        "id": "229",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "GNU General Public License v3.0 only": {
        "fedora_abbrev": "GPLv3",
        "fedora_name": "GNU General Public License v3.0 only",
        "id": "230",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "GPL-3.0",
        "spdx_name": "GNU General Public License v3.0 only",
        "url": "http://www.gnu.org/licenses/gpl-3.0-standalone.html"
    },
    "GNU General Public License v3.0 only, with Classpath exception": {
        "fedora_abbrev": "GPLv3 with exceptions",
        "fedora_name": "GNU General Public License v3.0 only, with Classpath exception",
        "id": "231",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "GNU General Public License v3.0 only, with font embedding exception": {
        "fedora_abbrev": "GPLv3 with exceptions",
        "fedora_name": "GNU General Public License v3.0 only, with font embedding exception",
        "id": "232",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "GNU General Public License v3.0 or later": {
        "fedora_abbrev": "GPLv3+",
        "fedora_name": "GNU General Public License v3.0 or later",
        "id": "233",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "GNU General Public License v3.0 or later, with Classpath exception": {
        "fedora_abbrev": "GPLv3+ with exceptions",
        "fedora_name": "GNU General Public License v3.0 or later, with Classpath exception",
        "id": "234",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "GNU General Public License v3.0 or later, with font embedding exception": {
        "fedora_abbrev": "GPLv3+ with exceptions",
        "fedora_name": "GNU General Public License v3.0 or later, with font embedding exception",
        "id": "235",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "GNU General Public License, Version 2 with the Classpath Exception": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "236",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": "http://repository.jboss.org/licenses/gpl-2.0-ce.txt"
    },
    "GNU Lesser General Public License (no version)": {
        "fedora_abbrev": "LGPLv2+",
        "fedora_name": "GNU Lesser General Public License (no version)",
        "id": "237",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "GNU Lesser General Public License v2 (or 2.1) only": {
        "fedora_abbrev": "LGPLv2",
        "fedora_name": "GNU Lesser General Public License v2 (or 2.1) only",
        "id": "239",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "GNU Lesser General Public License v2 (or 2.1) or later": {
        "fedora_abbrev": "LGPLv2+",
        "fedora_name": "GNU Lesser General Public License v2 (or 2.1) or later",
        "id": "240",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "GNU Lesser General Public License v2 (or 2.1) or later, with exceptions": {
        "fedora_abbrev": "LGPLv2+ with exceptions",
        "fedora_name": "GNU Lesser General Public License v2 (or 2.1) or later, with exceptions",
        "id": "241",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "GNU Lesser General Public License v2 (or 2.1), with exceptions": {
        "fedora_abbrev": "LGPLv2 with exceptions",
        "fedora_name": "GNU Lesser General Public License v2 (or 2.1), with exceptions",
        "id": "242",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "GNU Lesser General Public License v2.1 only": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "238",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "LGPL-2.1",
        "spdx_name": "GNU Lesser General Public License v2.1 only",
        "url": "http://www.gnu.org/licenses/old-licenses/lgpl-2.1-standalone.html"
    },
    "GNU Lesser General Public License v3.0 only": {
        "fedora_abbrev": "LGPLv3",
        "fedora_name": "GNU Lesser General Public License v3.0 only",
        "id": "243",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "LGPL-3.0",
        "spdx_name": "GNU Lesser General Public License v3.0 only",
        "url": "http://www.gnu.org/licenses/lgpl-3.0-standalone.html"
    },
    "GNU Lesser General Public License v3.0 only, with exceptions": {
        "fedora_abbrev": "LGPLv3 with exceptions",
        "fedora_name": "GNU Lesser General Public License v3.0 only, with exceptions",
        "id": "244",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "GNU Lesser General Public License v3.0 or later": {
        "fedora_abbrev": "LGPLv3+",
        "fedora_name": "GNU Lesser General Public License v3.0 or later",
        "id": "245",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "GNU Lesser General Public License v3.0 or later, with exceptions": {
        "fedora_abbrev": "LGPLv3+ with exceptions",
        "fedora_name": "GNU Lesser General Public License v3.0 or later, with exceptions",
        "id": "246",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "GNU Lesser General Public License, Version 2.1": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "247",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": "http://repository.jboss.org/licenses/lgpl-2.1.txt"
    },
    "GNU Lesser General Public License, Version 3": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "248",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "LGPL-3.0",
        "spdx_name": "GNU Lesser General Public License v3.0 only",
        "url": "http://www.gnu.org/licenses/lgpl-3.0-standalone.html"
    },
    "GNU Library General Public License v2 only": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "249",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "LGPL-2.0",
        "spdx_name": "GNU Library General Public License v2 only",
        "url": "http://www.gnu.org/licenses/old-licenses/lgpl-2.0-standalone.html"
    },
    "GNU Library General Public License, Version 2": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "250",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "LGPL-2.0",
        "spdx_name": "GNU Library General Public License v2 only",
        "url": "http://www.gnu.org/licenses/old-licenses/lgpl-2.0-standalone.html"
    },
    "GPL for Computer Programs of the Public Administration": {
        "fedora_abbrev": "",
        "fedora_name": "GPL for Computer Programs of the Public Administration",
        "id": "252",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Giftware License": {
        "fedora_abbrev": "Giftware",
        "fedora_name": "Giftware License",
        "id": "212",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Giftware",
        "spdx_name": "Giftware License",
        "url": "http://alleg.sourceforge.net//license.html"
    },
    "Glulxe License": {
        "fedora_abbrev": "Glulxe",
        "fedora_name": "Glulxe License",
        "id": "214",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Glulxe",
        "spdx_name": "Glulxe License",
        "url": "https://fedoraproject.org/wiki/Licensing/Glulxe"
    },
    "H2 License, Version 1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "255",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": "http://repository.jboss.org/licenses/h2.txt"
    },
    "HP Software License Terms": {
        "fedora_abbrev": "",
        "fedora_name": "HP Software License Terms",
        "id": "262",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Hacktivismo Enhanced-Source Software License Agreement": {
        "fedora_abbrev": "",
        "fedora_name": "Hacktivismo Enhanced-Source Software License Agreement",
        "id": "256",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Haskell Language Report License": {
        "fedora_abbrev": "HaskellReport",
        "fedora_name": "Haskell Language Report License",
        "id": "257",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "HaskellReport",
        "spdx_name": "Haskell Language Report License",
        "url": "https://fedoraproject.org/wiki/Licensing/Haskell_Language_Report_License"
    },
    "Helix DNA Technology Binary Research Use License": {
        "fedora_abbrev": "",
        "fedora_name": "Helix DNA Technology Binary Research Use License",
        "id": "258",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Henry Spencer Reg-Ex Library License": {
        "fedora_abbrev": "HSRL",
        "fedora_name": "Henry Spencer Reg-Ex Library License",
        "id": "259",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Historic Permission Notice and Disclaimer": {
        "fedora_abbrev": "MIT",
        "fedora_name": "Historical Permission Notice and Disclaimer",
        "id": "261",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "HPND",
        "spdx_name": "Historic Permission Notice and Disclaimer",
        "url": "http://www.opensource.org/licenses/HPND"
    },
    "Historical Permission Notice and Disclaimer": {
        "fedora_abbrev": "MIT",
        "fedora_name": "Historical Permission Notice and Disclaimer",
        "id": "260",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "HPND",
        "spdx_name": "Historic Permission Notice and Disclaimer",
        "url": "http://www.opensource.org/licenses/HPND"
    },
    "IBM PowerPC Initialization and Boot Software": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "263",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "IBM-pibs",
        "spdx_name": "IBM PowerPC Initialization and Boot Software",
        "url": "http://git.denx.de/?p=u-boot.git;a=blob;f=arch/powerpc/cpu/ppc4xx/miiphy.c;h=297155fdafa064b955e53e9832de93bfb0cfb85b;hb=9fab4bf4cc077c21e43941866f3f2c196f28670d"
    },
    "IBM Public License": {
        "fedora_abbrev": "IBM",
        "fedora_name": "IBM Public License",
        "id": "264",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "IBM Public License v1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "265",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "IPL-1.0",
        "spdx_name": "IBM Public License v1.0",
        "url": "http://www.opensource.org/licenses/IPL-1.0"
    },
    "IBM Sample Code License": {
        "fedora_abbrev": "",
        "fedora_name": "IBM Sample Code License",
        "id": "266",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "ICU License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "267",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "ICU",
        "spdx_name": "ICU License",
        "url": "http://source.icu-project.org/repos/icu/icu/trunk/license.html"
    },
    "IPA Font License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "278",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "IPA",
        "spdx_name": "IPA Font License",
        "url": "http://www.opensource.org/licenses/IPA"
    },
    "ISC License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "279",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "ISC",
        "spdx_name": "ISC License",
        "url": "http://www.isc.org/software/license"
    },
    "ISC License (Bind, DHCP Server)": {
        "fedora_abbrev": "ISC",
        "fedora_name": "ISC License (Bind, DHCP Server)",
        "id": "280",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "ImageMagick License": {
        "fedora_abbrev": "ImageMagick",
        "fedora_name": "ImageMagick License",
        "id": "268",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "ImageMagick",
        "spdx_name": "ImageMagick License",
        "url": "http://www.imagemagick.org/script/license.php"
    },
    "Imlib2 License": {
        "fedora_abbrev": "Imlib2",
        "fedora_name": "Imlib2 License",
        "id": "270",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Imlib2",
        "spdx_name": "Imlib2 License",
        "url": "http://trac.enlightenment.org/e/browser/trunk/imlib2/COPYING"
    },
    "Independent JPEG Group License": {
        "fedora_abbrev": "IJG",
        "fedora_name": "Independent JPEG Group License",
        "id": "271",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "IJG",
        "spdx_name": "Independent JPEG Group License",
        "url": "http://dev.w3.org/cvsweb/Amaya/libjpeg/Attic/README?rev=1.2"
    },
    "Info-ZIP License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "272",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "Info-ZIP",
        "spdx_name": "Info-ZIP License",
        "url": "http://www.info-zip.org/license.html"
    },
    "Intel ACPI Software License Agreement": {
        "fedora_abbrev": "Intel ACPI",
        "fedora_name": "Intel ACPI Software License Agreement",
        "id": "273",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Intel-ACPI",
        "spdx_name": "Intel ACPI Software License Agreement",
        "url": "https://fedoraproject.org/wiki/Licensing/Intel_ACPI_Software_License_Agreement"
    },
    "Intel IPW3945 Daemon License": {
        "fedora_abbrev": "",
        "fedora_name": "Intel IPW3945 Daemon License",
        "id": "274",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Intel Open Source License": {
        "fedora_abbrev": "",
        "fedora_name": "Intel Open Source License",
        "id": "275",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "Intel",
        "spdx_name": "Intel Open Source License",
        "url": "http://opensource.org/licenses/Intel"
    },
    "Interbase Public License": {
        "fedora_abbrev": "Interbase",
        "fedora_name": "Interbase Public License",
        "id": "276",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Interbase Public License v1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "277",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "Interbase-1.0",
        "spdx_name": "Interbase Public License v1.0",
        "url": "https://web.archive.org/web/20060319014854/http://info.borland.com/devsupport/interbase/opensource/IPL.html"
    },
    "JDOM License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "284",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": "http://repository.jboss.org/licenses/jdom-1.0.txt"
    },
    "JPython License (old)": {
        "fedora_abbrev": "JPython",
        "fedora_name": "JPython License (old)",
        "id": "285",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "JSON License": {
        "fedora_abbrev": "",
        "fedora_name": "JSON License",
        "id": "286",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "JSON",
        "spdx_name": "JSON License",
        "url": "http://www.json.org/license.html"
    },
    "Jabber Open Source License": {
        "fedora_abbrev": "Jabber",
        "fedora_name": "Jabber Open Source License",
        "id": "281",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Jahia Community Source License": {
        "fedora_abbrev": "",
        "fedora_name": "Jahia Community Source License",
        "id": "282",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "JasPer License": {
        "fedora_abbrev": "JasPer",
        "fedora_name": "JasPer License",
        "id": "283",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "JasPer-2.0",
        "spdx_name": "JasPer License",
        "url": "http://www.ece.uvic.ca/~mdadams/jasper/LICENSE"
    },
    "Julius License": {
        "fedora_abbrev": "Julius",
        "fedora_name": "Julius License",
        "id": "287",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Knuth License": {
        "fedora_abbrev": "Knuth",
        "fedora_name": "Knuth License",
        "id": "288",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "LEGO Open Source License Agreement": {
        "fedora_abbrev": "LOSLA",
        "fedora_name": "LEGO Open Source License Agreement",
        "id": "297",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "LaTeX Project Public License": {
        "fedora_abbrev": "LPPL",
        "fedora_name": "LaTeX Project Public License",
        "id": "290",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "LaTeX Project Public License v1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "291",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "LPPL-1.0",
        "spdx_name": "LaTeX Project Public License v1.0",
        "url": "http://www.latex-project.org/lppl/lppl-1-0.txt"
    },
    "LaTeX Project Public License v1.1": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "292",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "LPPL-1.1",
        "spdx_name": "LaTeX Project Public License v1.1",
        "url": "http://www.latex-project.org/lppl/lppl-1-1.txt"
    },
    "LaTeX Project Public License v1.2": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "293",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "LPPL-1.2",
        "spdx_name": "LaTeX Project Public License v1.2",
        "url": "http://www.latex-project.org/lppl/lppl-1-2.txt"
    },
    "LaTeX Project Public License v1.3a": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "294",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "LPPL-1.3a",
        "spdx_name": "LaTeX Project Public License v1.3a",
        "url": "http://www.latex-project.org/lppl/lppl-1-3a.txt"
    },
    "LaTeX Project Public License v1.3c": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "295",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "LPPL-1.3c",
        "spdx_name": "LaTeX Project Public License v1.3c",
        "url": "http://www.latex-project.org/lppl/lppl-1-3c.txt"
    },
    "Latex2e License": {
        "fedora_abbrev": "Latex2e",
        "fedora_name": "Latex2e License",
        "id": "289",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Latex2e",
        "spdx_name": "Latex2e License",
        "url": "https://fedoraproject.org/wiki/Licensing/Latex2e"
    },
    "Lawrence Berkeley National Labs BSD variant license": {
        "fedora_abbrev": "LBNL BSD",
        "fedora_name": "Lawrence Berkeley National Labs BSD variant license",
        "id": "296",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "BSD-3-Clause-LBNL",
        "spdx_name": "Lawrence Berkeley National Labs BSD variant license",
        "url": "https://fedoraproject.org/wiki/Licensing/LBNLBSD"
    },
    "Leptonica License": {
        "fedora_abbrev": "Leptonica",
        "fedora_name": "Leptonica License",
        "id": "298",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Leptonica",
        "spdx_name": "Leptonica License",
        "url": "https://fedoraproject.org/wiki/Licensing/Leptonica"
    },
    "Lesser General Public License For Linguistic Resources": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "299",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "LGPLLR",
        "spdx_name": "Lesser General Public License For Linguistic Resources",
        "url": "http://www-igm.univ-mlv.fr/~unitex/lgpllr.html"
    },
    "Lhcyr License": {
        "fedora_abbrev": "Lhcyr",
        "fedora_name": "Lhcyr License",
        "id": "301",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Licence Art Libre 1.2": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "304",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "LAL-1.2",
        "spdx_name": "Licence Art Libre 1.2",
        "url": "http://artlibre.org/licence/lal/licence-art-libre-12/"
    },
    "Licence Art Libre 1.3": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "305",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "LAL-1.3",
        "spdx_name": "Licence Art Libre 1.3",
        "url": "http://artlibre.org/"
    },
    "Licence Libre du Qu\u00e9bec \u2013 Permissive version 1.1": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "306",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "LiLiQ-P-1.1",
        "spdx_name": "Licence Libre du Qu\u00e9bec \u2013 Permissive version 1.1",
        "url": "http://opensource.org/licenses/LiLiQ-P-1.1"
    },
    "Licence Libre du Qu\u00e9bec \u2013 R\u00e9ciprocit\u00e9 forte version 1.1": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "307",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "LiLiQ-Rplus-1.1",
        "spdx_name": "Licence Libre du Qu\u00e9bec \u2013 R\u00e9ciprocit\u00e9 forte version 1.1",
        "url": "http://opensource.org/licenses/LiLiQ-Rplus-1.1"
    },
    "Licence Libre du Qu\u00e9bec \u2013 R\u00e9ciprocit\u00e9 version 1.1": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "308",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "LiLiQ-R-1.1",
        "spdx_name": "Licence Libre du Qu\u00e9bec \u2013 R\u00e9ciprocit\u00e9 version 1.1",
        "url": "http://opensource.org/licenses/LiLiQ-R-1.1"
    },
    "License Agreement for Application Response Measurement (ARM) SDK": {
        "fedora_abbrev": "",
        "fedora_name": "License Agreement for Application Response Measurement (ARM) SDK",
        "id": "309",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Lisp Library General Public License": {
        "fedora_abbrev": "LLGPL",
        "fedora_name": "Lisp Library General Public License",
        "id": "310",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Logica Open Source License": {
        "fedora_abbrev": "Logica",
        "fedora_name": "Logica Open Source License",
        "id": "311",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Lucent Public License (Plan9)": {
        "fedora_abbrev": "LPL",
        "fedora_name": "Lucent Public License (Plan9)",
        "id": "312",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Lucent Public License Version 1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "314",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "LPL-1.0",
        "spdx_name": "Lucent Public License Version 1.0",
        "url": "http://opensource.org/licenses/LPL-1.0"
    },
    "Lucent Public License v1.02": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "313",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "LPL-1.02",
        "spdx_name": "Lucent Public License v1.02",
        "url": "http://plan9.bell-labs.com/plan9/license.html"
    },
    "MAME License": {
        "fedora_abbrev": "",
        "fedora_name": "MAME License",
        "id": "317",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "MIT +no-false-attribs license": {
        "fedora_abbrev": "MITNFA",
        "fedora_name": "MIT +no-false-attribs license",
        "id": "330",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "MITNFA",
        "spdx_name": "MIT +no-false-attribs license",
        "url": "https://fedoraproject.org/wiki/Licensing/MITNFA"
    },
    "MIT License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "328",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "MIT",
        "spdx_name": "MIT License",
        "url": "http://www.opensource.org/licenses/MIT"
    },
    "MIT license (also X11)": {
        "fedora_abbrev": "MIT",
        "fedora_name": "MIT license (also X11)",
        "id": "329",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "MITRE Collaborative Virtual Workspace License (CVW)": {
        "fedora_abbrev": "",
        "fedora_name": "MITRE Collaborative Virtual Workspace License (CVW)",
        "id": "331",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "MSNTP License": {
        "fedora_abbrev": "",
        "fedora_name": "MSNTP License",
        "id": "344",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "MX4J License": {
        "fedora_abbrev": "ASL 1.1",
        "fedora_name": "MX4J License",
        "id": "347",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Maia Mailguard License": {
        "fedora_abbrev": "",
        "fedora_name": "Maia Mailguard License",
        "id": "315",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "MakeIndex License": {
        "fedora_abbrev": "MakeIndex",
        "fedora_name": "MakeIndex License",
        "id": "316",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "MakeIndex",
        "spdx_name": "MakeIndex License",
        "url": "https://fedoraproject.org/wiki/Licensing/MakeIndex"
    },
    "Matrix Template Library License": {
        "fedora_abbrev": "MTLL",
        "fedora_name": "Matrix Template Library License",
        "id": "318",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "MTLL",
        "spdx_name": "Matrix Template Library License",
        "url": "https://fedoraproject.org/wiki/Licensing/Matrix_Template_Library_License"
    },
    "MeepZor Consulting Public Licence": {
        "fedora_abbrev": "",
        "fedora_name": "MeepZor Consulting Public Licence",
        "id": "320",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Metasploit Framework License (post 2006)": {
        "fedora_abbrev": "BSD",
        "fedora_name": "Metasploit Framework License (post 2006)",
        "id": "321",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Metasploit Framework License (pre 2006)": {
        "fedora_abbrev": "",
        "fedora_name": "Metasploit Framework License (pre 2006)",
        "id": "322",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Microsoft Public License": {
        "fedora_abbrev": "MS-PL",
        "fedora_name": "Microsoft Public License",
        "id": "323",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "MS-PL",
        "spdx_name": "Microsoft Public License",
        "url": "http://www.microsoft.com/opensource/licenses.mspx"
    },
    "Microsoft Reciprocal License": {
        "fedora_abbrev": "MS-RL",
        "fedora_name": "Microsoft Reciprocal License",
        "id": "324",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "MS-RL",
        "spdx_name": "Microsoft Reciprocal License",
        "url": "http://www.microsoft.com/opensource/licenses.mspx"
    },
    "Microsoft's Shared Source CLI/C#/Jscript License": {
        "fedora_abbrev": "",
        "fedora_name": "Microsoft's Shared Source CLI/C#/Jscript License",
        "id": "325",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "MirOS License": {
        "fedora_abbrev": "MirOS",
        "fedora_name": "MirOS License",
        "id": "327",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "MirOS",
        "spdx_name": "MirOS Licence",
        "url": "http://www.opensource.org/licenses/MirOS"
    },
    "Motosoto License": {
        "fedora_abbrev": "Motosoto",
        "fedora_name": "Motosoto License",
        "id": "333",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Motosoto",
        "spdx_name": "Motosoto License",
        "url": "http://www.opensource.org/licenses/Motosoto"
    },
    "Mozilla Public License 1.0": {
        "fedora_abbrev": "MPLv1.0",
        "fedora_name": "Mozilla Public License v1.0",
        "id": "334",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "MPL-1.0",
        "spdx_name": "Mozilla Public License 1.0",
        "url": "http://www.mozilla.org/MPL/MPL-1.0.html"
    },
    "Mozilla Public License 1.1": {
        "fedora_abbrev": "MPLv1.1",
        "fedora_name": "Mozilla Public License v1.1",
        "id": "335",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "MPL-1.1",
        "spdx_name": "Mozilla Public License 1.1",
        "url": "http://www.mozilla.org/MPL/MPL-1.1.html"
    },
    "Mozilla Public License 2.0": {
        "fedora_abbrev": "MPLv2.0",
        "fedora_name": "Mozilla Public License v2.0",
        "id": "336",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "MPL-2.0",
        "spdx_name": "Mozilla Public License 2.0",
        "url": "http://www.mozilla.org/MPL/2.0/"
    },
    "Mozilla Public License 2.0 (no copyleft exception)": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "337",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "MPL-2.0-no-copyleft-exception",
        "spdx_name": "Mozilla Public License 2.0 (no copyleft exception)",
        "url": "http://www.mozilla.org/MPL/2.0/"
    },
    "Mozilla Public License v1.0": {
        "fedora_abbrev": "MPLv1.0",
        "fedora_name": "Mozilla Public License v1.0",
        "id": "338",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "MPL-1.0",
        "spdx_name": "Mozilla Public License 1.0",
        "url": "http://www.mozilla.org/MPL/MPL-1.0.html"
    },
    "Mozilla Public License v1.1": {
        "fedora_abbrev": "MPLv1.1",
        "fedora_name": "Mozilla Public License v1.1",
        "id": "339",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "MPL-1.1",
        "spdx_name": "Mozilla Public License 1.1",
        "url": "http://www.mozilla.org/MPL/MPL-1.1.html"
    },
    "Mozilla Public License v2.0": {
        "fedora_abbrev": "MPLv2.0",
        "fedora_name": "Mozilla Public License v2.0",
        "id": "340",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "MPL-2.0",
        "spdx_name": "Mozilla Public License 2.0",
        "url": "http://www.mozilla.org/MPL/2.0/"
    },
    "Mozilla Public License, Version 1.1": {
        "fedora_abbrev": "MPLv1.1",
        "fedora_name": "Mozilla Public License v1.1",
        "id": "341",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "MPL-1.1",
        "spdx_name": "Mozilla Public License 1.1",
        "url": "http://www.mozilla.org/MPL/MPL-1.1.html"
    },
    "Multics License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "345",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "Multics",
        "spdx_name": "Multics License",
        "url": "http://www.opensource.org/licenses/Multics"
    },
    "Mup License": {
        "fedora_abbrev": "Mup",
        "fedora_name": "Mup License",
        "id": "346",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Mup",
        "spdx_name": "Mup License",
        "url": "https://fedoraproject.org/wiki/Licensing/Mup"
    },
    "MySQL License": {
        "fedora_abbrev": "GPLv2 with exceptions",
        "fedora_name": "MySQL License",
        "id": "348",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "NASA CDF License": {
        "fedora_abbrev": "",
        "fedora_name": "NASA CDF License",
        "id": "349",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "NASA Open Source Agreement 1.3": {
        "fedora_abbrev": "",
        "fedora_name": "NASA Open Source Agreement v1.3",
        "id": "350",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "NASA-1.3",
        "spdx_name": "NASA Open Source Agreement 1.3",
        "url": "http://ti.arc.nasa.gov/opensource/nosa/"
    },
    "NASA Open Source Agreement v1.3": {
        "fedora_abbrev": "",
        "fedora_name": "NASA Open Source Agreement v1.3",
        "id": "351",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "NASA-1.3",
        "spdx_name": "NASA Open Source Agreement 1.3",
        "url": "http://ti.arc.nasa.gov/opensource/nosa/"
    },
    "NCSA/University of Illinois Open Source License": {
        "fedora_abbrev": "NCSA",
        "fedora_name": "NCSA/University of Illinois Open Source License",
        "id": "353",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "NRL License": {
        "fedora_abbrev": "BSD with advertising",
        "fedora_name": "NRL License",
        "id": "371",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "NRL",
        "spdx_name": "NRL License",
        "url": "http://web.mit.edu/network/isakmp/nrllicense.html"
    },
    "NTP License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "372",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "NTP",
        "spdx_name": "NTP License",
        "url": "http://www.opensource.org/licenses/NTP"
    },
    "Naumen Public License": {
        "fedora_abbrev": "Naumen",
        "fedora_name": "Naumen Public License",
        "id": "352",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Naumen",
        "spdx_name": "Naumen Public License",
        "url": "http://www.opensource.org/licenses/Naumen"
    },
    "Neotonic Clearsilver License": {
        "fedora_abbrev": "ASL 1.1",
        "fedora_name": "Neotonic Clearsilver License",
        "id": "354",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Net Boolean Public License v1": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "355",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "NBPL-1.0",
        "spdx_name": "Net Boolean Public License v1",
        "url": "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=37b4b3f6cc4bf34e1d3dec61e69914b9819d8894"
    },
    "Net-SNMP License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "362",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "Net-SNMP",
        "spdx_name": "Net-SNMP License",
        "url": "http://net-snmp.sourceforge.net/about/license.html"
    },
    "NetCDF license": {
        "fedora_abbrev": "NetCDF",
        "fedora_name": "NetCDF license",
        "id": "356",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "NetCDF",
        "spdx_name": "NetCDF license",
        "url": "http://www.unidata.ucar.edu/software/netcdf/copyright.html"
    },
    "Nethack General Public License": {
        "fedora_abbrev": "NGPL",
        "fedora_name": "Nethack General Public License",
        "id": "357",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "NGPL",
        "spdx_name": "Nethack General Public License",
        "url": "http://www.opensource.org/licenses/NGPL"
    },
    "Netizen Open Source License": {
        "fedora_abbrev": "NOSL",
        "fedora_name": "Netizen Open Source License",
        "id": "358",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "NOSL",
        "spdx_name": "Netizen Open Source License",
        "url": "http://bits.netizen.com.au/licenses/NOSL/nosl.txt"
    },
    "Netscape Public License": {
        "fedora_abbrev": "Netscape",
        "fedora_name": "Netscape Public License",
        "id": "359",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Netscape Public License v1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "360",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "NPL-1.0",
        "spdx_name": "Netscape Public License v1.0",
        "url": "http://www.mozilla.org/MPL/NPL/1.0/"
    },
    "Netscape Public License v1.1": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "361",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "NPL-1.1",
        "spdx_name": "Netscape Public License v1.1",
        "url": "http://www.mozilla.org/MPL/NPL/1.1/"
    },
    "Newmat License": {
        "fedora_abbrev": "Newmat",
        "fedora_name": "Newmat License",
        "id": "363",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Newsletr License": {
        "fedora_abbrev": "Newsletr",
        "fedora_name": "Newsletr License",
        "id": "364",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Newsletr",
        "spdx_name": "Newsletr License",
        "url": "https://fedoraproject.org/wiki/Licensing/Newsletr"
    },
    "Nmap License": {
        "fedora_abbrev": "Nmap",
        "fedora_name": "Nmap License",
        "id": "365",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "No Limit Public License": {
        "fedora_abbrev": "NLPL",
        "fedora_name": "No Limit Public License",
        "id": "367",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "NLPL",
        "spdx_name": "No Limit Public License",
        "url": "https://fedoraproject.org/wiki/Licensing/NLPL"
    },
    "Nokia Open Source License": {
        "fedora_abbrev": "Nokia",
        "fedora_name": "Nokia Open Source License",
        "id": "366",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Nokia",
        "spdx_name": "Nokia Open Source License",
        "url": "http://www.opensource.org/licenses/nokia"
    },
    "Non-Profit Open Software License 3.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "368",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "NPOSL-3.0",
        "spdx_name": "Non-Profit Open Software License 3.0",
        "url": "http://www.opensource.org/licenses/NOSL3.0"
    },
    "Norwegian Licence for Open Government Data": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "369",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "NLOD-1.0",
        "spdx_name": "Norwegian Licence for Open Government Data",
        "url": "http://data.norge.no/nlod/en/1.0"
    },
    "Noweb License": {
        "fedora_abbrev": "Noweb",
        "fedora_name": "Noweb License",
        "id": "370",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Noweb",
        "spdx_name": "Noweb License",
        "url": "https://fedoraproject.org/wiki/Licensing/Noweb"
    },
    "Nunit License": {
        "fedora_abbrev": "MIT with advertising",
        "fedora_name": "Nunit License",
        "id": "373",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Nunit",
        "spdx_name": "Nunit License",
        "url": "https://fedoraproject.org/wiki/Licensing/Nunit"
    },
    "OCLC Public Research License 2.0": {
        "fedora_abbrev": "",
        "fedora_name": "OCLC Public Research License 2.0",
        "id": "374",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "OCLC-2.0",
        "spdx_name": "OCLC Research Public License 2.0",
        "url": "http://www.oclc.org/research/activities/software/license/v2final.htm"
    },
    "OCLC Research Public License 2.0": {
        "fedora_abbrev": "",
        "fedora_name": "OCLC Public Research License 2.0",
        "id": "375",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "OCLC-2.0",
        "spdx_name": "OCLC Research Public License 2.0",
        "url": "http://www.oclc.org/research/activities/software/license/v2final.htm"
    },
    "ODC Open Database License v1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "376",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "ODbL-1.0",
        "spdx_name": "ODC Open Database License v1.0",
        "url": "http://www.opendatacommons.org/licenses/odbl/1.0/"
    },
    "ODC Public Domain Dedication & License 1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "377",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "PDDL-1.0",
        "spdx_name": "ODC Public Domain Dedication & License 1.0",
        "url": "http://opendatacommons.org/licenses/pddl/1.0/"
    },
    "OReilly License": {
        "fedora_abbrev": "OReilly",
        "fedora_name": "OReilly License",
        "id": "410",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "OSET Public License version 2.1": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "411",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "OSET-PL-2.1",
        "spdx_name": "OSET Public License version 2.1",
        "url": "http://opensource.org/licenses/OPL-2.1"
    },
    "OSGi Specification License": {
        "fedora_abbrev": "",
        "fedora_name": "OSGi Specification License",
        "id": "412",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Open CASCADE Technology Public License": {
        "fedora_abbrev": "",
        "fedora_name": "Open CASCADE Technology Public License",
        "id": "378",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "OCCT-PL",
        "spdx_name": "Open CASCADE Technology Public License",
        "url": "http://www.opencascade.com/content/occt-public-license"
    },
    "Open Government License": {
        "fedora_abbrev": "",
        "fedora_name": "Open Government License",
        "id": "379",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Open Group Test Suite License": {
        "fedora_abbrev": "",
        "fedora_name": "Open Group Test Suite License",
        "id": "380",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "OGTSL",
        "spdx_name": "Open Group Test Suite License",
        "url": "http://www.opengroup.org/testing/downloads/The_Open_Group_TSL.txt"
    },
    "Open LDAP Public License 2.2.2": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "382",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "OLDAP-2.2.2",
        "spdx_name": "Open LDAP Public License 2.2.2",
        "url": ""
    },
    "Open LDAP Public License v1.1": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "383",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "OLDAP-1.1",
        "spdx_name": "Open LDAP Public License v1.1",
        "url": "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=806557a5ad59804ef3a44d5abfbe91d706b0791f"
    },
    "Open LDAP Public License v1.2": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "384",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "OLDAP-1.2",
        "spdx_name": "Open LDAP Public License v1.2",
        "url": "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=42b0383c50c299977b5893ee695cf4e486fb0dc7"
    },
    "Open LDAP Public License v1.3": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "385",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "OLDAP-1.3",
        "spdx_name": "Open LDAP Public License v1.3",
        "url": "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=e5f8117f0ce088d0bd7a8e18ddf37eaa40eb09b1"
    },
    "Open LDAP Public License v1.4": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "386",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "OLDAP-1.4",
        "spdx_name": "Open LDAP Public License v1.4",
        "url": "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=c9f95c2f3f2ffb5e0ae55fe7388af75547660941"
    },
    "Open LDAP Public License v2.0 (or possibly 2.0A and 2.0B)": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "388",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "OLDAP-2.0",
        "spdx_name": "Open LDAP Public License v2.0 (or possibly 2.0A and 2.0B)",
        "url": "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=cbf50f4e1185a21abd4c0a54d3f4341fe28f36ea"
    },
    "Open LDAP Public License v2.0.1": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "387",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "OLDAP-2.0.1",
        "spdx_name": "Open LDAP Public License v2.0.1",
        "url": "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=b6d68acd14e51ca3aab4428bf26522aa74873f0e"
    },
    "Open LDAP Public License v2.1": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "389",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "OLDAP-2.1",
        "spdx_name": "Open LDAP Public License v2.1",
        "url": "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=b0d176738e96a0d3b9f85cb51e140a86f21be715"
    },
    "Open LDAP Public License v2.2": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "390",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "OLDAP-2.2",
        "spdx_name": "Open LDAP Public License v2.2",
        "url": "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=470b0c18ec67621c85881b2733057fecf4a1acc3"
    },
    "Open LDAP Public License v2.2.1": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "391",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "OLDAP-2.2.1",
        "spdx_name": "Open LDAP Public License v2.2.1",
        "url": "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=4bc786f34b50aa301be6f5600f58a980070f481e"
    },
    "Open LDAP Public License v2.3": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "392",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "OLDAP-2.3",
        "spdx_name": "Open LDAP Public License v2.3",
        "url": "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=d32cf54a32d581ab475d23c810b0a7fbaf8d63c3"
    },
    "Open LDAP Public License v2.4": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "393",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "OLDAP-2.4",
        "spdx_name": "Open LDAP Public License v2.4",
        "url": "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=cd1284c4a91a8a380d904eee68d1583f989ed386"
    },
    "Open LDAP Public License v2.5": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "394",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "OLDAP-2.5",
        "spdx_name": "Open LDAP Public License v2.5",
        "url": "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=6852b9d90022e8593c98205413380536b1b5a7cf"
    },
    "Open LDAP Public License v2.6": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "395",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "OLDAP-2.6",
        "spdx_name": "Open LDAP Public License v2.6",
        "url": "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=1cae062821881f41b73012ba816434897abf4205"
    },
    "Open LDAP Public License v2.7": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "396",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "OLDAP-2.7",
        "spdx_name": "Open LDAP Public License v2.7",
        "url": "http://www.openldap.org/devel/gitweb.cgi?p=openldap.git;a=blob;f=LICENSE;hb=47c2415c1df81556eeb39be6cad458ef87c534a2"
    },
    "Open LDAP Public License v2.8": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "397",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "OLDAP-2.8",
        "spdx_name": "Open LDAP Public License v2.8",
        "url": "http://www.openldap.org/software/release/license.html"
    },
    "Open Map License": {
        "fedora_abbrev": "",
        "fedora_name": "Open Map License",
        "id": "398",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Open Market License": {
        "fedora_abbrev": "OML",
        "fedora_name": "Open Market License",
        "id": "399",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "OML",
        "spdx_name": "Open Market License",
        "url": "https://fedoraproject.org/wiki/Licensing/Open_Market_License"
    },
    "Open Motif Public End User License": {
        "fedora_abbrev": "",
        "fedora_name": "Open Motif Public End User License",
        "id": "400",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Open Public License": {
        "fedora_abbrev": "",
        "fedora_name": "Open Public License",
        "id": "402",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Open Public License v1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "403",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "OPL-1.0",
        "spdx_name": "Open Public License v1.0",
        "url": "http://old.koalateam.com/jackaroo/OPL_1_0.TXT"
    },
    "Open Software License 1.0": {
        "fedora_abbrev": "OSL 1.0",
        "fedora_name": "Open Software License 1.0",
        "id": "404",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "OSL-1.0",
        "spdx_name": "Open Software License 1.0",
        "url": "http://opensource.org/licenses/OSL-1.0"
    },
    "Open Software License 1.1": {
        "fedora_abbrev": "OSL 1.1",
        "fedora_name": "Open Software License 1.1",
        "id": "405",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "OSL-1.1",
        "spdx_name": "Open Software License 1.1",
        "url": "https://fedoraproject.org/wiki/Licensing/OSL1.1"
    },
    "Open Software License 2.0": {
        "fedora_abbrev": "OSL 2.0",
        "fedora_name": "Open Software License 2.0",
        "id": "406",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "OSL-2.0",
        "spdx_name": "Open Software License 2.0",
        "url": "http://web.archive.org/web/20041020171434/http://www.rosenlaw.com/osl2.0.html"
    },
    "Open Software License 2.1": {
        "fedora_abbrev": "OSL 2.1",
        "fedora_name": "Open Software License 2.1",
        "id": "407",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "OSL-2.1",
        "spdx_name": "Open Software License 2.1",
        "url": "http://opensource.org/licenses/OSL-2.1"
    },
    "Open Software License 3.0": {
        "fedora_abbrev": "OSL 3.0",
        "fedora_name": "Open Software License 3.0",
        "id": "408",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "OSL-3.0",
        "spdx_name": "Open Software License 3.0",
        "url": "http://www.rosenlaw.com/OSL3.0.htm"
    },
    "OpenLDAP License": {
        "fedora_abbrev": "OpenLDAP",
        "fedora_name": "OpenLDAP License",
        "id": "381",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "OpenPBS License": {
        "fedora_abbrev": "OpenPBS",
        "fedora_name": "OpenPBS License",
        "id": "401",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "OpenSSL License": {
        "fedora_abbrev": "OpenSSL",
        "fedora_name": "OpenSSL License",
        "id": "409",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "OpenSSL",
        "spdx_name": "OpenSSL License",
        "url": "http://www.openssl.org/source/license.html"
    },
    "PHP License v3.0": {
        "fedora_abbrev": "PHP",
        "fedora_name": "PHP License v3.0",
        "id": "421",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "PHP-3.0",
        "spdx_name": "PHP License v3.0",
        "url": "http://www.opensource.org/licenses/PHP-3.0"
    },
    "PHP License v3.01": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "422",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "PHP-3.01",
        "spdx_name": "PHP License v3.01",
        "url": "http://www.php.net/license/3_01.txt"
    },
    "Par License": {
        "fedora_abbrev": "Par",
        "fedora_name": "Par License",
        "id": "413",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Paul Hsieh Derivative License": {
        "fedora_abbrev": "",
        "fedora_name": "Paul Hsieh Derivative License",
        "id": "414",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Paul Hsieh Exposition License": {
        "fedora_abbrev": "",
        "fedora_name": "Paul Hsieh Exposition License",
        "id": "415",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Perl License": {
        "fedora_abbrev": "GPL+ or Artistic",
        "fedora_name": "Perl License",
        "id": "416",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Perl License (variant)": {
        "fedora_abbrev": "GPLv2 or Artistic",
        "fedora_name": "Perl License (variant)",
        "id": "419",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Phorum License": {
        "fedora_abbrev": "Phorum",
        "fedora_name": "Phorum License",
        "id": "420",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Pine License": {
        "fedora_abbrev": "",
        "fedora_name": "Pine License",
        "id": "423",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "PlainTeX License": {
        "fedora_abbrev": "PlainTeX",
        "fedora_name": "PlainTeX License",
        "id": "424",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Plexus Classworlds License": {
        "fedora_abbrev": "Plexus",
        "fedora_name": "Plexus Classworlds License",
        "id": "425",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Plexus",
        "spdx_name": "Plexus Classworlds License",
        "url": "https://fedoraproject.org/wiki/Licensing/Plexus_Classworlds_License"
    },
    "PostgreSQL License": {
        "fedora_abbrev": "PostgreSQL",
        "fedora_name": "PostgreSQL License",
        "id": "426",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "PostgreSQL",
        "spdx_name": "PostgreSQL License",
        "url": "http://www.postgresql.org/about/licence"
    },
    "Public Domain": {
        "fedora_abbrev": "Public Domain",
        "fedora_name": "Public Domain",
        "id": "429",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Python License": {
        "fedora_abbrev": "Python",
        "fedora_name": "Python License",
        "id": "430",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Python License 2.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "431",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "Python-2.0",
        "spdx_name": "Python License 2.0",
        "url": "http://www.opensource.org/licenses/Python-2.0"
    },
    "Q Public License": {
        "fedora_abbrev": "QPL",
        "fedora_name": "Q Public License",
        "id": "434",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Q Public License 1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "435",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "QPL-1.0",
        "spdx_name": "Q Public License 1.0",
        "url": "http://doc.qt.nokia.com/3.3/license.html"
    },
    "Qhull License": {
        "fedora_abbrev": "Qhull",
        "fedora_name": "Qhull License",
        "id": "432",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Qhull",
        "spdx_name": "Qhull License",
        "url": "https://fedoraproject.org/wiki/Licensing/Qhull"
    },
    "QuickFix License": {
        "fedora_abbrev": "ASL 1.1",
        "fedora_name": "QuickFix License",
        "id": "436",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Qwt License 1.0": {
        "fedora_abbrev": "LGPLv2+ with exceptions",
        "fedora_name": "Qwt License 1.0",
        "id": "437",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "REX License": {
        "fedora_abbrev": "REX",
        "fedora_name": "REX License",
        "id": "444",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "RSA Message-Digest License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "448",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "RSA-MD",
        "spdx_name": "RSA Message-Digest License",
        "url": ""
    },
    "Rdisc License": {
        "fedora_abbrev": "Rdisc",
        "fedora_name": "Rdisc License",
        "id": "438",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Rdisc",
        "spdx_name": "Rdisc License",
        "url": "https://fedoraproject.org/wiki/Licensing/Rdisc_License"
    },
    "RealNetworks Public Source License V1.0": {
        "fedora_abbrev": "RPSL",
        "fedora_name": "RealNetworks Public Source License V1.0",
        "id": "439",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "RPSL-1.0",
        "spdx_name": "RealNetworks Public Source License v1.0",
        "url": "https://helixcommunity.org/content/rpsl"
    },
    "Reciprocal Public License": {
        "fedora_abbrev": "",
        "fedora_name": "Reciprocal Public License",
        "id": "440",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Reciprocal Public License 1.1": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "441",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "RPL-1.1",
        "spdx_name": "Reciprocal Public License 1.1",
        "url": "http://opensource.org/licenses/RPL-1.1"
    },
    "Reciprocal Public License 1.5": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "442",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "RPL-1.5",
        "spdx_name": "Reciprocal Public License 1.5",
        "url": "http://www.opensource.org/licenses/RPL-1.5"
    },
    "Red Hat eCos Public License v1.1": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "443",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "RHeCos-1.1",
        "spdx_name": "Red Hat eCos Public License v1.1",
        "url": "http://ecos.sourceware.org/old-license.html"
    },
    "Rice BSD": {
        "fedora_abbrev": "RiceBSD",
        "fedora_name": "Rice BSD",
        "id": "445",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Ricoh Source Code Public License": {
        "fedora_abbrev": "",
        "fedora_name": "Ricoh Source Code Public License",
        "id": "446",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "RSCPL",
        "spdx_name": "Ricoh Source Code Public License",
        "url": "http://www.opensource.org/licenses/RSCPL"
    },
    "Romio License": {
        "fedora_abbrev": "Romio",
        "fedora_name": "Romio License",
        "id": "447",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Rsfs License": {
        "fedora_abbrev": "Rsfs",
        "fedora_name": "Rsfs License",
        "id": "449",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Ruby License": {
        "fedora_abbrev": "Ruby",
        "fedora_name": "Ruby License",
        "id": "450",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Ruby",
        "spdx_name": "Ruby License",
        "url": "http://www.ruby-lang.org/en/LICENSE.txt"
    },
    "SCEA Shared Source License": {
        "fedora_abbrev": "SCEA",
        "fedora_name": "SCEA Shared Source License",
        "id": "453",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "SCEA",
        "spdx_name": "SCEA Shared Source License",
        "url": "http://research.scea.com/scea_shared_source_license.html"
    },
    "SGI Free Software License B 1.1 or older": {
        "fedora_abbrev": "",
        "fedora_name": "SGI Free Software License B 1.1 or older",
        "id": "460",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "SGI Free Software License B 2.0": {
        "fedora_abbrev": "MIT",
        "fedora_name": "SGI Free Software License B 2.0",
        "id": "461",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "SGI-B-2.0",
        "spdx_name": "SGI Free Software License B v2.0",
        "url": "http://oss.sgi.com/projects/FreeB/SGIFreeSWLicB.2.0.pdf"
    },
    "SGI Free Software License B v1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "462",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "SGI-B-1.0",
        "spdx_name": "SGI Free Software License B v1.0",
        "url": "http://oss.sgi.com/projects/FreeB/SGIFreeSWLicB.1.0.html"
    },
    "SGI Free Software License B v1.1": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "463",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "SGI-B-1.1",
        "spdx_name": "SGI Free Software License B v1.1",
        "url": "http://oss.sgi.com/projects/FreeB/"
    },
    "SGI Free Software License B v2.0": {
        "fedora_abbrev": "MIT",
        "fedora_name": "SGI Free Software License B 2.0",
        "id": "464",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "SGI-B-2.0",
        "spdx_name": "SGI Free Software License B v2.0",
        "url": "http://oss.sgi.com/projects/FreeB/SGIFreeSWLicB.2.0.pdf"
    },
    "SGI GLX Public License 1.0": {
        "fedora_abbrev": "",
        "fedora_name": "SGI GLX Public License 1.0",
        "id": "465",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "SIL Open Font License 1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "466",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "OFL-1.0",
        "spdx_name": "SIL Open Font License 1.0",
        "url": "http://scripts.sil.org/cms/scripts/page.php?item_id=OFL10_web"
    },
    "SIL Open Font License 1.1": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "467",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "OFL-1.1",
        "spdx_name": "SIL Open Font License 1.1",
        "url": "http://scripts.sil.org/cms/scripts/page.php?item_id=OFL_web"
    },
    "SLIB License": {
        "fedora_abbrev": "SLIB",
        "fedora_name": "SLIB License",
        "id": "472",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "SNIA Public License 1.1": {
        "fedora_abbrev": "SNIA",
        "fedora_name": "SNIA Public License 1.1",
        "id": "473",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "SNIA",
        "spdx_name": "SNIA Public License 1.1",
        "url": "https://fedoraproject.org/wiki/Licensing/SNIA_Public_License"
    },
    "Sax Public Domain Notice": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "452",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "SAX-PD",
        "spdx_name": "Sax Public Domain Notice",
        "url": "http://www.saxproject.org/copying.html"
    },
    "Saxpath License": {
        "fedora_abbrev": "Saxpath",
        "fedora_name": "Saxpath License",
        "id": "451",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Saxpath",
        "spdx_name": "Saxpath License",
        "url": "https://fedoraproject.org/wiki/Licensing/Saxpath_License"
    },
    "Scheme Widget Library (SWL) Software License Agreement": {
        "fedora_abbrev": "SWL",
        "fedora_name": "Scheme Widget Library (SWL) Software License Agreement",
        "id": "454",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "SWL",
        "spdx_name": "Scheme Widget Library (SWL) Software License Agreement",
        "url": "https://fedoraproject.org/wiki/Licensing/SWL"
    },
    "SciTech MGL Public License": {
        "fedora_abbrev": "STMPL",
        "fedora_name": "SciTech MGL Public License",
        "id": "456",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Scilab License (OLD)": {
        "fedora_abbrev": "",
        "fedora_name": "Scilab License (OLD)",
        "id": "455",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Secure Messaging Protocol Public License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "457",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "SMPPL",
        "spdx_name": "Secure Messaging Protocol Public License",
        "url": "https://github.com/dcblake/SMP/blob/master/Documentation/License.txt"
    },
    "Sendmail License": {
        "fedora_abbrev": "Sendmail",
        "fedora_name": "Sendmail License",
        "id": "458",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Sendmail",
        "spdx_name": "Sendmail License",
        "url": "http://www.sendmail.com/pdfs/open_source/sendmail_license.pdf"
    },
    "Sequence Library License": {
        "fedora_abbrev": "Sequence",
        "fedora_name": "Sequence Library License",
        "id": "459",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Simple Public License 2.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "468",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "SimPL-2.0",
        "spdx_name": "Simple Public License 2.0",
        "url": "http://www.opensource.org/licenses/SimPL-2.0"
    },
    "Siren14 License Agreement": {
        "fedora_abbrev": "",
        "fedora_name": "Siren14 License Agreement",
        "id": "469",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Sleepycat License": {
        "fedora_abbrev": "Sleepycat",
        "fedora_name": "Sleepycat Software Product License",
        "id": "470",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Sleepycat",
        "spdx_name": "Sleepycat License",
        "url": "http://www.opensource.org/licenses/Sleepycat"
    },
    "Sleepycat Software Product License": {
        "fedora_abbrev": "Sleepycat",
        "fedora_name": "Sleepycat Software Product License",
        "id": "471",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Sleepycat",
        "spdx_name": "Sleepycat License",
        "url": "http://www.opensource.org/licenses/Sleepycat"
    },
    "Spencer License 86": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "475",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "Spencer-86",
        "spdx_name": "Spencer License 86",
        "url": "https://fedoraproject.org/wiki/Licensing/Henry_Spencer_Reg-Ex_Library_License"
    },
    "Spencer License 94": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "476",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "Spencer-94",
        "spdx_name": "Spencer License 94",
        "url": "https://fedoraproject.org/wiki/Licensing/Henry_Spencer_Reg-Ex_Library_License"
    },
    "Spencer License 99": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "477",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "Spencer-99",
        "spdx_name": "Spencer License 99",
        "url": "http://www.opensource.apple.com/source/tcl/tcl-5/tcl/generic/regfronts.c"
    },
    "Spin Commercial License": {
        "fedora_abbrev": "",
        "fedora_name": "Spin Commercial License",
        "id": "478",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Squeak License": {
        "fedora_abbrev": "",
        "fedora_name": "Squeak License",
        "id": "479",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Standard ML of New Jersey License": {
        "fedora_abbrev": "MIT",
        "fedora_name": "Standard ML of New Jersey License",
        "id": "480",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "SMLNJ",
        "spdx_name": "Standard ML of New Jersey License",
        "url": "https://fedoraproject.org/wiki/Licensing:MIT?rd=Licensing/MIT#Standard_ML_of_New_Jersey_Variant"
    },
    "SugarCRM Public License v1.1.3": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "481",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "SugarCRM-1.1.3",
        "spdx_name": "SugarCRM Public License v1.1.3",
        "url": "http://www.sugarcrm.com/crm/SPL"
    },
    "Sun Binary Code License Agreement": {
        "fedora_abbrev": "",
        "fedora_name": "Sun Binary Code License Agreement",
        "id": "482",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Sun Community Source License": {
        "fedora_abbrev": "",
        "fedora_name": "Sun Community Source License",
        "id": "483",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Sun Industry Standards Source License": {
        "fedora_abbrev": "SISSL",
        "fedora_name": "Sun Industry Standards Source License",
        "id": "484",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Sun Industry Standards Source License v1.1": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "485",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "SISSL",
        "spdx_name": "Sun Industry Standards Source License v1.1",
        "url": "http://www.openoffice.org/licenses/sissl_license.html"
    },
    "Sun Industry Standards Source License v1.2": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "486",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "SISSL-1.2",
        "spdx_name": "Sun Industry Standards Source License v1.2",
        "url": "http://gridscheduler.sourceforge.net/Gridengine_SISSL_license.html"
    },
    "Sun Public License": {
        "fedora_abbrev": "SPL",
        "fedora_name": "Sun Public License",
        "id": "487",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Sun Public License v1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "488",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "SPL-1.0",
        "spdx_name": "Sun Public License v1.0",
        "url": "http://www.opensource.org/licenses/SPL-1.0"
    },
    "Sun RPC License": {
        "fedora_abbrev": "",
        "fedora_name": "Sun RPC License",
        "id": "489",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Sun Solaris Source Code (Foundation Release) License": {
        "fedora_abbrev": "",
        "fedora_name": "Sun Solaris Source Code (Foundation Release) License",
        "id": "490",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Sybase Open Watcom Public License 1.0": {
        "fedora_abbrev": "",
        "fedora_name": "Sybase Open Watcom Public License 1.0",
        "id": "491",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "Watcom-1.0",
        "spdx_name": "Sybase Open Watcom Public License 1.0",
        "url": "http://www.opensource.org/licenses/Watcom-1.0"
    },
    "SystemC Open Source License": {
        "fedora_abbrev": "",
        "fedora_name": "SystemC Open Source License",
        "id": "492",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "TCL/TK License": {
        "fedora_abbrev": "TCL",
        "fedora_name": "TCL/TK License",
        "id": "493",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "TCL",
        "spdx_name": "TCL/TK License",
        "url": "http://www.tcl.tk/software/tcltk/license.html"
    },
    "TCP Wrappers License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "494",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "TCP-wrappers",
        "spdx_name": "TCP Wrappers License",
        "url": "http://rc.quest.com/topics/openssh/license.php#tcpwrappers"
    },
    "TMate Open Source License": {
        "fedora_abbrev": "TMate",
        "fedora_name": "TMate Open Source License",
        "id": "509",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "TMate",
        "spdx_name": "TMate Open Source License",
        "url": "http://svnkit.com/license.html"
    },
    "TORQUE v2.5+ Software License v1.0": {
        "fedora_abbrev": "",
        "fedora_name": "TORQUE v2.5+ Software License v1.0",
        "id": "511",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "TORQUE v2.5+ Software License v1.1": {
        "fedora_abbrev": "TORQUEv1.1",
        "fedora_name": "TORQUE v2.5+ Software License v1.1",
        "id": "512",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "TORQUE-1.1",
        "spdx_name": "TORQUE v2.5+ Software License v1.1",
        "url": "https://fedoraproject.org/wiki/Licensing/TORQUEv1.1"
    },
    "Teeworlds License": {
        "fedora_abbrev": "Teeworlds",
        "fedora_name": "Teeworlds License",
        "id": "495",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Terracotta Public License 1.0": {
        "fedora_abbrev": "",
        "fedora_name": "Terracotta Public License 1.0",
        "id": "496",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Text-Tabs+Wrap License": {
        "fedora_abbrev": "TTWL",
        "fedora_name": "Text-Tabs+Wrap License",
        "id": "497",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "The Antlr 2.7.7 License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "498",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": "http://www.antlr2.org/download/antlr-2.7.7.tar.gz"
    },
    "The Asm BSD License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "499",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": "http://asm.ow2.org/license.html"
    },
    "The BSD License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "500",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": "http://repository.jboss.org/licenses/bsd.txt"
    },
    "The Dom4j License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "501",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": "https://raw.githubusercontent.com/dom4j/dom4j/master/LICENSE"
    },
    "The JSoup MIT License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "503",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": "https://jsoup.org/license.html"
    },
    "The Jaxen License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "502",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": "http://www.jaxen.org/license.html"
    },
    "The MIT License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "504",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "MIT",
        "spdx_name": "MIT License",
        "url": "http://www.opensource.org/licenses/MIT"
    },
    "The Unlicense": {
        "fedora_abbrev": "Unlicense",
        "fedora_name": "Unlicense",
        "id": "505",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Unlicense",
        "spdx_name": "The Unlicense",
        "url": "http://unlicense.org/"
    },
    "Thor Public License": {
        "fedora_abbrev": "TPL",
        "fedora_name": "Thor Public License",
        "id": "506",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Threeparttable License": {
        "fedora_abbrev": "Threeparttable",
        "fedora_name": "Threeparttable License",
        "id": "507",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Time::ParseDate License": {
        "fedora_abbrev": "TPDL",
        "fedora_name": "Time::ParseDate License",
        "id": "508",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Tolua License": {
        "fedora_abbrev": "Tolua",
        "fedora_name": "Tolua License",
        "id": "510",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Transitive Grace Period Public Licence": {
        "fedora_abbrev": "TGPPL",
        "fedora_name": "Transitive Grace Period Public Licence",
        "id": "513",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "TrueCrypt License": {
        "fedora_abbrev": "",
        "fedora_name": "TrueCrypt License",
        "id": "514",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Trusster Open Source License": {
        "fedora_abbrev": "TOSL",
        "fedora_name": "Trusster Open Source License",
        "id": "515",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "TOSL",
        "spdx_name": "Trusster Open Source License",
        "url": "https://fedoraproject.org/wiki/Licensing/TOSL"
    },
    "Trusted Computing Group License": {
        "fedora_abbrev": "TCGL",
        "fedora_name": "Trusted Computing Group License",
        "id": "516",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "UCAR License": {
        "fedora_abbrev": "UCAR",
        "fedora_name": "UCAR License",
        "id": "517",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Unicode Character Database Terms Of Use": {
        "fedora_abbrev": "UCD",
        "fedora_name": "Unicode Character Database Terms Of Use",
        "id": "518",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Unicode License": {
        "fedora_abbrev": "Unicode",
        "fedora_name": "Unicode License",
        "id": "519",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Unicode License Agreement - Data Files and Software (2015)": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "520",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "Unicode-DFS-2015",
        "spdx_name": "Unicode License Agreement - Data Files and Software (2015)",
        "url": "https://web.archive.org/web/20151224134844/http://unicode.org/copyright.html"
    },
    "Unicode License Agreement - Data Files and Software (2016)": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "521",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "Unicode-DFS-2016",
        "spdx_name": "Unicode License Agreement - Data Files and Software (2016)",
        "url": "http://www.unicode.org/copyright.html"
    },
    "Unicode Terms of Use": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "522",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "Unicode-TOU",
        "spdx_name": "Unicode Terms of Use",
        "url": "http://www.unicode.org/copyright.html"
    },
    "Universal Permissive License v1.0": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "523",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "UPL-1.0",
        "spdx_name": "Universal Permissive License v1.0",
        "url": "http://opensource.org/licenses/UPL"
    },
    "University of Illinois/NCSA Open Source License": {
        "fedora_abbrev": "NCSA",
        "fedora_name": "NCSA/University of Illinois Open Source License",
        "id": "524",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "NCSA",
        "spdx_name": "University of Illinois/NCSA Open Source License",
        "url": "http://otm.illinois.edu/uiuc_openSource"
    },
    "University of Utah Public License": {
        "fedora_abbrev": "",
        "fedora_name": "University of Utah Public License",
        "id": "525",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "University of Washington Free Fork License": {
        "fedora_abbrev": "",
        "fedora_name": "University of Washington Free Fork License",
        "id": "526",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Unlicense": {
        "fedora_abbrev": "Unlicense",
        "fedora_name": "Unlicense",
        "id": "527",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Unlicense",
        "spdx_name": "The Unlicense",
        "url": "http://unlicense.org/"
    },
    "VOSTROM Public License for Open Source": {
        "fedora_abbrev": "VOSTROM",
        "fedora_name": "VOSTROM Public License for Open Source",
        "id": "531",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "VOSTROM",
        "spdx_name": "VOSTROM Public License for Open Source",
        "url": "https://fedoraproject.org/wiki/Licensing/VOSTROM"
    },
    "Vim License": {
        "fedora_abbrev": "Vim",
        "fedora_name": "Vim License",
        "id": "529",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Vim",
        "spdx_name": "Vim License",
        "url": "http://vimdoc.sourceforge.net/htmldoc/uganda.html"
    },
    "Vita Nuova Liberal Source License": {
        "fedora_abbrev": "VNLSL",
        "fedora_name": "Vita Nuova Liberal Source License",
        "id": "530",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Vovida Software License v. 1.0": {
        "fedora_abbrev": "VSL",
        "fedora_name": "Vovida Software License v. 1.0",
        "id": "532",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "VSL-1.0",
        "spdx_name": "Vovida Software License v1.0",
        "url": "http://www.opensource.org/licenses/VSL-1.0"
    },
    "Vovida Software License v1.0": {
        "fedora_abbrev": "VSL",
        "fedora_name": "Vovida Software License v. 1.0",
        "id": "533",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "VSL-1.0",
        "spdx_name": "Vovida Software License v1.0",
        "url": "http://www.opensource.org/licenses/VSL-1.0"
    },
    "W3C Software Notice and Document License (2015-05-13)": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "534",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "W3C-20150513",
        "spdx_name": "W3C Software Notice and Document License (2015-05-13)",
        "url": "https://www.w3.org/Consortium/Legal/2015/copyright-software-and-document"
    },
    "W3C Software Notice and License": {
        "fedora_abbrev": "W3C",
        "fedora_name": "W3C Software Notice and License",
        "id": "535",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "W3C Software Notice and License (1998-07-20)": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "536",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "W3C-19980720",
        "spdx_name": "W3C Software Notice and License (1998-07-20)",
        "url": "http://www.w3.org/Consortium/Legal/copyright-software-19980720.html"
    },
    "W3C Software Notice and License (2002-12-31)": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "537",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "W3C",
        "spdx_name": "W3C Software Notice and License (2002-12-31)",
        "url": "http://www.w3.org/Consortium/Legal/2002/copyright-software-20021231.html"
    },
    "Webmin License": {
        "fedora_abbrev": "Webmin",
        "fedora_name": "Webmin License",
        "id": "538",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Wsuipa License": {
        "fedora_abbrev": "Wsuipa",
        "fedora_name": "Wsuipa License",
        "id": "539",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Wsuipa",
        "spdx_name": "Wsuipa License",
        "url": "https://fedoraproject.org/wiki/Licensing/Wsuipa"
    },
    "X.Net License": {
        "fedora_abbrev": "",
        "fedora_name": "X.Net License",
        "id": "545",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "Xnet",
        "spdx_name": "X.Net License",
        "url": "http://opensource.org/licenses/Xnet"
    },
    "X11 License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "541",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "X11",
        "spdx_name": "X11 License",
        "url": "http://www.xfree86.org/3.3.6/COPYRIGHT2.html#3"
    },
    "XFree86 License 1.1": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "543",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "XFree86-1.1",
        "spdx_name": "XFree86 License 1.1",
        "url": "http://www.xfree86.org/current/LICENSE4.html"
    },
    "XPP License": {
        "fedora_abbrev": "xpp",
        "fedora_name": "XPP License",
        "id": "546",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "xpp",
        "spdx_name": "XPP License",
        "url": "https://fedoraproject.org/wiki/Licensing/xpp"
    },
    "XSkat License": {
        "fedora_abbrev": "XSkat",
        "fedora_name": "XSkat License",
        "id": "547",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "XSkat",
        "spdx_name": "XSkat License",
        "url": "https://fedoraproject.org/wiki/Licensing/XSkat_License"
    },
    "Xerox License": {
        "fedora_abbrev": "Xerox",
        "fedora_name": "Xerox License",
        "id": "542",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Xerox",
        "spdx_name": "Xerox License",
        "url": "https://fedoraproject.org/wiki/Licensing/Xerox"
    },
    "Yahoo Public License 1.0": {
        "fedora_abbrev": "",
        "fedora_name": "Yahoo Public License 1.0",
        "id": "548",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "YPL-1.0",
        "spdx_name": "Yahoo! Public License v1.0",
        "url": "http://www.zimbra.com/license/yahoo_public_license_1.0.html"
    },
    "Yahoo Public License v 1.1": {
        "fedora_abbrev": "YPLv1.1",
        "fedora_name": "Yahoo Public License v 1.1",
        "id": "550",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "YPL-1.1",
        "spdx_name": "Yahoo! Public License v1.1",
        "url": "http://www.zimbra.com/license/yahoo_public_license_1.1.html"
    },
    "Yahoo! Public License v1.0": {
        "fedora_abbrev": "",
        "fedora_name": "Yahoo Public License 1.0",
        "id": "549",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "YPL-1.0",
        "spdx_name": "Yahoo! Public License v1.0",
        "url": "http://www.zimbra.com/license/yahoo_public_license_1.0.html"
    },
    "Yahoo! Public License v1.1": {
        "fedora_abbrev": "YPLv1.1",
        "fedora_name": "Yahoo Public License v 1.1",
        "id": "551",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "YPL-1.1",
        "spdx_name": "Yahoo! Public License v1.1",
        "url": "http://www.zimbra.com/license/yahoo_public_license_1.1.html"
    },
    "Zed License": {
        "fedora_abbrev": "Zed",
        "fedora_name": "Zed License",
        "id": "552",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Zed",
        "spdx_name": "Zed License",
        "url": "https://fedoraproject.org/wiki/Licensing/Zed"
    },
    "Zend License v2.0": {
        "fedora_abbrev": "Zend",
        "fedora_name": "Zend License v2.0",
        "id": "553",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "Zend-2.0",
        "spdx_name": "Zend License v2.0",
        "url": "https://web.archive.org/web/20130517195954/http://www.zend.com/license/2_00.txt"
    },
    "Zimbra Public License 1.3": {
        "fedora_abbrev": "",
        "fedora_name": "Zimbra Public License 1.3",
        "id": "554",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "Zimbra-1.3",
        "spdx_name": "Zimbra Public License v1.3",
        "url": ""
    },
    "Zimbra Public License v1.3": {
        "fedora_abbrev": "",
        "fedora_name": "Zimbra Public License 1.3",
        "id": "555",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "Zimbra-1.3",
        "spdx_name": "Zimbra Public License v1.3",
        "url": ""
    },
    "Zimbra Public License v1.4": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "556",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "Zimbra-1.4",
        "spdx_name": "Zimbra Public License v1.4",
        "url": "http://www.zimbra.com/legal/zimbra-public-license-1-4"
    },
    "Zope Public License 1.1": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "560",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "ZPL-1.1",
        "spdx_name": "Zope Public License 1.1",
        "url": "http://old.zope.org/Resources/License/ZPL-1.1"
    },
    "Zope Public License 2.0": {
        "fedora_abbrev": "ZPLv2.0",
        "fedora_name": "Zope Public License v 2.0",
        "id": "561",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "ZPL-2.0",
        "spdx_name": "Zope Public License 2.0",
        "url": "http://old.zope.org/Resources/License/ZPL-2.0"
    },
    "Zope Public License 2.1": {
        "fedora_abbrev": "ZPLv2.1",
        "fedora_name": "Zope Public License v 2.1",
        "id": "562",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "ZPL-2.1",
        "spdx_name": "Zope Public License 2.1",
        "url": "http://old.zope.org/Resources/ZPL/"
    },
    "Zope Public License v 1.0": {
        "fedora_abbrev": "ZPLv1.0",
        "fedora_name": "Zope Public License v 1.0",
        "id": "563",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "Zope Public License v 2.0": {
        "fedora_abbrev": "ZPLv2.0",
        "fedora_name": "Zope Public License v 2.0",
        "id": "564",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "ZPL-2.0",
        "spdx_name": "Zope Public License 2.0",
        "url": "http://old.zope.org/Resources/License/ZPL-2.0"
    },
    "Zope Public License v 2.1": {
        "fedora_abbrev": "ZPLv2.1",
        "fedora_name": "Zope Public License v 2.1",
        "id": "565",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "ZPL-2.1",
        "spdx_name": "Zope Public License 2.1",
        "url": "http://old.zope.org/Resources/ZPL/"
    },
    "bzip2 and libbzip2 License v1.0.5": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "85",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "bzip2-1.0.5",
        "spdx_name": "bzip2 and libbzip2 License v1.0.5",
        "url": "http://bzip.org/1.0.5/bzip2-manual-1.0.5.html"
    },
    "bzip2 and libbzip2 License v1.0.6": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "86",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "bzip2-1.0.6",
        "spdx_name": "bzip2 and libbzip2 License v1.0.6",
        "url": "https://github.com/asimonov-im/bzip2/blob/master/LICENSE"
    },
    "curl License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "165",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "curl",
        "spdx_name": "curl License",
        "url": "https://github.com/bagder/curl/blob/master/COPYING"
    },
    "diffmark license": {
        "fedora_abbrev": "diffmark",
        "fedora_name": "diffmark license",
        "id": "167",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "diffmark",
        "spdx_name": "diffmark license",
        "url": "https://fedoraproject.org/wiki/Licensing/diffmark"
    },
    "dvipdfm License": {
        "fedora_abbrev": "dvipdfm",
        "fedora_name": "dvipdfm License",
        "id": "174",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "dvipdfm",
        "spdx_name": "dvipdfm License",
        "url": "https://fedoraproject.org/wiki/Licensing/dvipdfm"
    },
    "eCos License v2.0": {
        "fedora_abbrev": "eCos",
        "fedora_name": "eCos License v2.0",
        "id": "179",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "eCos Public License, v1.1": {
        "fedora_abbrev": "",
        "fedora_name": "eCos Public License, v1.1",
        "id": "180",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "eGenix.com Public License 1.1.0": {
        "fedora_abbrev": "eGenix",
        "fedora_name": "eGenix.com Public License 1.1.0",
        "id": "183",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "eGenix",
        "spdx_name": "eGenix.com Public License 1.1.0",
        "url": "http://www.egenix.com/products/eGenix.com-Public-License-1.1.0.pdf"
    },
    "enna License": {
        "fedora_abbrev": "MIT",
        "fedora_name": "enna License",
        "id": "188",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "MIT-enna",
        "spdx_name": "enna License",
        "url": "https://fedoraproject.org/wiki/Licensing/MIT#enna"
    },
    "feh License": {
        "fedora_abbrev": "MIT",
        "fedora_name": "feh License",
        "id": "200",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "MIT-feh",
        "spdx_name": "feh License",
        "url": "https://fedoraproject.org/wiki/Licensing/MIT#feh"
    },
    "gSOAP Public License": {
        "fedora_abbrev": "",
        "fedora_name": "gSOAP Public License",
        "id": "253",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "gSOAP Public License v1.3b": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "254",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "gSOAP-1.3b",
        "spdx_name": "gSOAP Public License v1.3b",
        "url": "http://www.cs.fsu.edu/~engelen/license.html"
    },
    "gnuplot License": {
        "fedora_abbrev": "gnuplot",
        "fedora_name": "gnuplot License",
        "id": "251",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "gnuplot",
        "spdx_name": "gnuplot License",
        "url": "https://fedoraproject.org/wiki/Licensing/Gnuplot"
    },
    "iMatix Standard Function Library Agreement": {
        "fedora_abbrev": "iMatix",
        "fedora_name": "iMatix Standard Function Library Agreement",
        "id": "269",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "iMatix",
        "spdx_name": "iMatix Standard Function Library Agreement",
        "url": "http://legacy.imatix.com/html/sfl/sfl4.htm#license"
    },
    "lha license": {
        "fedora_abbrev": "",
        "fedora_name": "lha license",
        "id": "300",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "libpng License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "302",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "Libpng",
        "spdx_name": "libpng License",
        "url": "http://www.libpng.org/pub/png/src/libpng-LICENSE.txt"
    },
    "libtiff License": {
        "fedora_abbrev": "libtiff",
        "fedora_name": "libtiff License",
        "id": "303",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "libtiff",
        "spdx_name": "libtiff License",
        "url": "https://fedoraproject.org/wiki/Licensing/libtiff"
    },
    "mecab-ipadic license": {
        "fedora_abbrev": "mecab-ipadic",
        "fedora_name": "mecab-ipadic license",
        "id": "319",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "midnight License": {
        "fedora_abbrev": "midnight",
        "fedora_name": "midnight License",
        "id": "326",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "mod_macro License": {
        "fedora_abbrev": "mod_macro",
        "fedora_name": "mod_macro License",
        "id": "332",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "mpich2 License": {
        "fedora_abbrev": "MIT",
        "fedora_name": "mpich2 License",
        "id": "342",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "mpich2",
        "spdx_name": "mpich2 License",
        "url": "https://fedoraproject.org/wiki/Licensing/MIT"
    },
    "mrouted license (old)": {
        "fedora_abbrev": "",
        "fedora_name": "mrouted license (old)",
        "id": "343",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "psfrag License": {
        "fedora_abbrev": "psfrag",
        "fedora_name": "psfrag License",
        "id": "427",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "psfrag",
        "spdx_name": "psfrag License",
        "url": "https://fedoraproject.org/wiki/Licensing/psfrag"
    },
    "psutils License": {
        "fedora_abbrev": "psutils",
        "fedora_name": "psutils License",
        "id": "428",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "psutils",
        "spdx_name": "psutils License",
        "url": "https://fedoraproject.org/wiki/Licensing/psutils"
    },
    "qmail License": {
        "fedora_abbrev": "",
        "fedora_name": "qmail License",
        "id": "433",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "softSurfer License": {
        "fedora_abbrev": "softSurfer",
        "fedora_name": "softSurfer License",
        "id": "474",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "unrar license": {
        "fedora_abbrev": "",
        "fedora_name": "unrar license",
        "id": "528",
        "license_text": "",
        "rh_approved": "no",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "wxWidgets Library License": {
        "fedora_abbrev": "wxWidgets",
        "fedora_name": "wxWidgets Library License",
        "id": "540",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "xinetd License": {
        "fedora_abbrev": "xinetd",
        "fedora_name": "xinetd License",
        "id": "544",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "xinetd",
        "spdx_name": "xinetd License",
        "url": "https://fedoraproject.org/wiki/Licensing/Xinetd_License"
    },
    "zlib License": {
        "fedora_abbrev": "",
        "fedora_name": "",
        "id": "559",
        "license_text": "",
        "rh_approved": "unknown",
        "spdx_abbrev": "Zlib",
        "spdx_name": "zlib License",
        "url": "http://www.zlib.net/zlib_license.html"
    },
    "zlib/libpng License": {
        "fedora_abbrev": "zlib",
        "fedora_name": "zlib/libpng License",
        "id": "557",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "",
        "spdx_name": "",
        "url": ""
    },
    "zlib/libpng License with Acknowledgement": {
        "fedora_abbrev": "zlib with acknowledgement",
        "fedora_name": "zlib/libpng License with Acknowledgement",
        "id": "558",
        "license_text": "",
        "rh_approved": "yes",
        "spdx_abbrev": "zlib-acknowledgement",
        "spdx_name": "zlib/libpng License with Acknowledgement",
        "url": "https://fedoraproject.org/wiki/Licensing/ZlibWithAcknowledgement"
    }
}
---

`
