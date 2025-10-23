# Overview

Test the effect of various kinds of manual edits to see whether changes get picked up by `kubectl diff` and Argo CD diff.

In Argo CD the test application is deployed with automated sync but self-heal is disabled so we can view the
Out-of-Sync state and how Argo CD views the change in its diffs.

### Argo CD Test Cases

| Use Case                                                 | kubectl diff | kubectl diff (SS)   |Argo CD Legacy Diff | Argo CD (SS) | Comment |
| -------------                                            | ------------ | ------------------- | ------------------ | ------------ | ------- |
| Edit the number of replicas in Deployment                |       ✅     |           ✅        |         ✅         |      ✅      |         |
| Add paused field in Deployment                           |       ❌     |           ❌        |         ❌         |      ❌      |   Paused state is detected but change is not      |
| Add a new label                                          |       ❌     |           ❌        |         ❌         |      ❌      | Label additions are not detected for individual labels, adding a new `labels` block will be detected |
| Add a new label (--save-config)                          |       ✅     |           ❌        |         ✅         |      ❌      |   Re-test Argo CD client side to make sure behavior is the same |
| Add a new label (--field-manager='argocd-controller')    |       ❌     |           ❌        |                    |      ❌      |         |
| Delete a label                                           |       ✅     |           ✅        |         ✅         |      ✅      |         |
| Change an existing label                                 |       ✅     |           ✅        |         ✅         |      ✅      |         |
| Add a data item to ConfigMap                             |       ❌     |           ❌        |         ❌         |      ❌      |         |
| Add a data item to ConfigMap (--save-config)             |       ✅     |           ❌        |         ✅         |      ❌      |         |
| Change an existing data item in ConfigMap                |       ✅     |           ✅        |         ✅         |      ✅      |         |
| Add an env var                                           |       ❌     |           ❌        |         ❌         |      ❌      |         |
| Change an env var                                        |       ✅     |           ✅        |         ✅         |      ✅      |         |
| Delete an env var                                        |       ✅     |           ✅        |         ✅         |      ✅      |         |
| Change default value (i.e. revisionHistoryLimit to 11)   |       ❌     |           ❌        |         ❌         |      ❌      |         |
| Change default value (--save-config)                     |       ✅     |           ❌        |         ✅         |      ❌      |         |
| Add `paused` field to deployment                         |       ❌     |           ❌        |         ❌         |      ❌      |  [omitempty](https://kubernetes.slack.com/archives/C09NXKJKA/p1760999271617209)  |

* SS - Server-Side-Diff
* --save-config. Add `--save-config` to `kubectl edit` command to update `last-applied-configuration` annotation
* --field-manager='argocd-controller'. Add this parameter to `kubectl edit` command ro register change against specific field manager

## Thoughts

* `kubectl diff` and `argocd diff` are not conventional diff utilities, instead they provide a patch for what would happen if
you ran a `kubectl apply` with the file. This is why manually adding a new field or editing a field in the live object that isn't
in the manifest gets ignored, kubernetes assumes a multi-controller world and only looks at fields included in the manifest.
* It's cleat that last-applied-annotation has no impact on server-side diff
