### Overview

Test the effect of various kinds of manual edits to see whether changes get picked up by `kubectl diff` and Argo CD diff.

In Argo CD the test application is deployed with automated sync but self-heal is disabled so we can view the
Out-of-Sync state and how Argo CD views the change in its diffs.

### Argo CD Test Cases

| Use Case                                                 | kubectl diff | kubectl diff (SS)   |Argo CD Legacy Diff | Argo CD (SS) | Comment |
| -------------                                            | ------------ | ------------------- | ------------------ | ------------ | ------- |
| Edit the number of replicas in Deployment                |       ✅     |           ✅        |         ✅         |      ✅      |         |
| Add paused field in Deployment                           |       ❌     |           ❌        |         ❌         |              |         |
| Add a new label                                          |       ❌     |           ❌        |         ❌         |      ❌      | Label additions are not detected for individual labels, adding a new `labels` block will be detected |
| Add a new label (--save-config)                          |       ✅     |           ✅        |         ✅         |              |         |
| Delete a label                                           |       ✅     |           ✅        |         ✅         |              |         |
| Change an existing label                                 |       ✅     |           ✅        |         ✅         |              |         |
| Add a data item to ConfigMap                             |       ❌     |           ❌        |         ❌         |              |         |
| Add a data item to ConfigMap (--save-config)             |       ✅     |           ✅        |         ✅         |              |         |
| Change an existing data item in ConfigMap                |       ✅     |           ✅        |         ✅         |              |         |
| Add an env var                                           |       ❌     |           ❌        |         ❌         |              |         |
| Change an env var                                        |       ✅     |           ✅        |         ✅         |              |         |
| Delete an env var                                        |       ✅     |           ✅        |         ✅         |              |         |
| Change default value (i.e. revisionHistoryLimit to 11)   |       ❌     |           ❌        |         ❌         |              |         |
| Change default value (--save-config)                     |       ✅     |           ✅        |         ✅         |              |         |
| Add `paused` field to deployment                         |       ❌     |           ❌        |         ❌         |      ❌      |  [omitempty](https://kubernetes.slack.com/archives/C09NXKJKA/p1760999271617209)  |

* SS - Server-Side-Diff
* --save-config - Add `--save-config` to `kubectl edit` command to update `last-applied-configuration` annotation

## Thoughts

* Simplistically I always thought that `kubectl diff` compares to `last-applied-configuration` annotation
when using client-side diff and when using server-side it compares to live object. However testing seems
to indicate it's not that simple, why doesn't it hold up for adding a field to the ConfigMap?
* I suspect the `paused` field might be a special case given the imperativeness of it, question on Kubernetes slack #kubernetes-users
