### Overview

Test cases using server-side diff such as mutation

| Use Case                                 | Argo CD Legacy |Argo CD (SS) | Argo CD (SSA - SS) | Argo CD (SSA - SS - Mutation |  Comment |
| Kyverno Policy that adds a label         |       ❌       |     ❌      |         ❌         |        ❌ (1 - ✅ on Refresh)    |          |
| Kyverno Policy replaces image digest     |       ✅       |     ❌      | ❌ (Refresh Failed - Error)|        ❌ (Refresh Failed - Error)   |          |
| MutatingAdmissionPolicy change replicas  |       ✅       |     ❌      |         ❌         |        ❌                    |          |
| MutatingAdmissionPolicy add label        |       ❌       |     ❌      |         ❌         |        ❌                    |          |
| MutatingWebHook (change replicas)        |       ✅       |     ❌      |         ❌         |        ❌                    |          |
| MutatingWebHook (add label)              |       ❌       |     ❌      |         ❌         |        ❌                    |          |

IncludeMutationWebHook only works when:
  - The field being affected isn't owned my the `argo-controller` field manager
  - Only reflects a change and this only kicks in on a refresh, i.e. if you add/change the mutation and then press refresh

