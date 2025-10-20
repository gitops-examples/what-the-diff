### Overview

Test the effect of various kinds of manual edits to see whether changes get picked up by `kubectl diff` and Argo CD diff.

In Argo CD the test application is deployed with automated sync but self-heal is disabled so we can view the
Out-of-Sync state and how Argo CD views the change in its diffs.

### Argo CD Test Cases

| Use Case                                            | kubectl diff | kubectl diff (SS)   |Argo CD Legacy Diff | Argo CD (SS) | Comment |
| -------------                                       | ------------ | ------------------- | ------------------ | ------------ | ------- |
| Edit the number of replicas in Deployment           |       ✅     |           ✅        |         ✅         |      ✅      |         |
| Add paused field in Deployment                      |       ❌     |           ❌        |         ❌         |              |         |
| Add a new label                                     |       ❌     |           ✅        |         ❌         |      ❌      | Discuss why Argo CD is not picking this up with SS |
| Add a new label (--save-config)                     |       ✅     |           ✅        |         ✅         |              |         |
| Delete a label                                      |       ✅     |           ✅        |         ✅         |              |         |
| Change an existing label                            |       ✅     |           ✅        |         ✅         |              |         |
| Add a data item to ConfigMap                        |       ❌     |           ❌        |         ❌         |              | Why does SS work for label but not ConfigMap?        |
| Add a data item to ConfigMap (--save-config)        |       ✅     |           ✅        |         ✅         |              |         |
| Change an existing data item in ConfigMap           |       ✅     |           ✅        |         ✅         |              |         |
| Add an env var                                      |       ❌     |           ❌        |         ❌         |              |         |
| Add default field and value (revisionHistoryLimit)  |              |                     |                    |              |         |
| Change default value (revisionHistoryLimit to 11)   |            |                     |                    |              |         |
| Add `paused` field to deployment                    |       ❌     |           ❌        |         ❌         |      ❌      |  There is something special about `paused`  |

* SS - Server-Side-Diff
* --save-config - Add `--save-config` to `kubectl edit` command to update `last-applied-configuration` annotation

## Thoughts

* Simplistically I always thought that `kubectl diff` compares to `last-applied-configuration` annotation
when using client-side diff and when using server-side it should compares to live object. However testing seems
to indicate it's not that simple, why doesn't it hold up for adding a field to the ConfigMap?
* I suspect the `paused` field might be a special case given the imperativeness of it, question on Kubernetes slack #kubernetes-users
