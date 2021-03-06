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

// BuildConfigTemplate contains a default example yaml build config.
const BuildConfigTemplate = `
---
  default-build-target: my-build-target
  distributions:
    my-product-zip-dist:
      type: zip
    my-product-container-dist:
      type: docker
  projects:
    my-test-project1:
      owner: jdoe
      version: 1.0
      scm-url: "http://github.com/my-test-project1"
      commit-id: abcd1234
      build-type: maven-build
      build-options:
      maven-name: org.test.test-project1
      licenses:
        - ASL 2.0
      maven-artifacts:
        - groupId: org.test.project1
          artifactId: test-project1-artifact1
      distributions:
        - my-product-zip-dist
    my-test-project2:
      owner: jdoe
      version: 2.1.0.Beta1
      maven-name: org.test.test-project2
      maven-groupid: org.test.project2
      licenses:
        - ASL 2.0
        - DL 1.0
      maven-artifacts:
        - artifactId: test-project2-artifact1
        - artifactId: test-project2-artifact2
          type: war
      distributions:
        - my-product-zip-dist
        - my-product-container-dist

`
