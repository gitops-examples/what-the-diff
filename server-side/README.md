### Overview

Test cases using server-side diff such as mutation

| Use Case                                 | Argo CD Legacy |Argo CD (SS) | Argo CD (SS - Mutation) | Argo CD (SSA) | Comment |
| Kyverno Policy that adds a label         |       ❌       |     ❌      |          ❌             |               |         |
| Kyverno Policy replaces image digest     |       ✅       |     ❌      |          ❌             |               |         |
| MutatingAdmissionPolicy change replicas  |       ✅       |     ❌      |          ❌             |               |         |

Thoughts


