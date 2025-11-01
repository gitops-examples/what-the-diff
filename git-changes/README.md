# Overview

This tests changes that are made in git, i.e. deploy application, update manifest in git, refresh and see what diff shows.

### Argo CD Test Cases

| Use Case                                                 | kubectl diff | kubectl diff (SS)   |Argo CD Legacy Diff | Argo CD (SS) | Argo CD (SSA) | Comment |
| -------------                                            | ------------ | ------------------- | ------------------ | ------------ | ------------- | ------- |
| Change Replicas                                          |       ✅     |         ✅          |         ✅         |      ✅      |      ✅       |         |
| Remove data element from ConfigMap                       |       ✅     |         ❌          |         ✅         |      ❌      |      ✅       |         |
| Remove label                                             |       ✅     |         ❌          |         ✅         |      ❌      |      ✅       |         |
| Remove annotation                                        |       ✅     |         ❌          |         ✅         |      ❌      |      ✅       |         |
| Remove annotation block                                  |       ✅     |         ❌          |         ✅         |      ❌      |      ✅       |         |

