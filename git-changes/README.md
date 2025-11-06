# Overview

This tests changes that are made in git, i.e. deploy application, update manifest in git, refresh and see what diff shows.

### Argo CD Test Cases

| Use Case                                                 | kubectl diff | kubectl diff (SS)   |Argo CD Legacy Diff | Argo CD (SS) | Argo CD (SSA) | Argo CD (SSA + SS) | Comment |
| -------------                                            | ------------ | ------------------- | ------------------ | ------------ | ------------- | ------------------ |-------- |
| Change Replicas                                          |       ✅     |         ✅          |         ✅         |      ✅      |      ✅       |         ✅         |         |
| Remove data element from ConfigMap                       |       ✅     |         ❌          |         ✅         |      ❌      |      ✅       |         ✅         |         |
| Remove label                                             |       ✅     |         ❌          |         ✅         |      ❌      |      ✅       |                    |         |
| Remove annotation                                        |       ✅     |         ❌          |         ✅         |      ❌      |      ✅       |                    |         |
| Remove annotation block                                  |       ✅     |         ❌          |         ✅         |      ❌      |      ✅       |                    |         |
| Add `affinity` to Deployment                             |       ✅     |         ✅          |         ✅         |      ✅      |      ✅       |         ✅         |         |
| Remove `affinity` from Deployment                        |       ✅     |         ❌          |         ✅         |      ❌      |      ✅       |         ✅         |         |
| Add `revisionHistoryLimit` with default to Deployment    |       ❌     |         ❌          |         ❌         |      ❌      |      ❌       |                    | Negative result is what we want |
| Add `revisionHistoryLimit: 11` to Deployment             |       ✅     |         ✅          |         ✅         |      ✅      |      ✅       |                    |         |
| Add `revisionHistoryLmt: 10` to Deployment (typo)        |       ❌     |        Failed       |         ✅         |    Failed    |    Failed     |                    |         |


