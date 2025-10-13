### Overview

Test the effect of various kinds of manual edits to see whether changes get picked up by `kubectl diff` and Argo CD diff.

In Argo CD the test application is deployed with automated sync but self-heal is disabled so we can view the
Out-of-Sync state and how Argo CD views the change in its diffs.

### Test Cases

| Use Case                                   | kubectl diff | Argo CD Legacy Diff | Argo CD Server-Side Diff |
| -------------                              | ------------ | ------------------- | ------------------------ |
| Edit the number of replicas in Deployment  |   |   |   |
| Add a new label                            |   |   |   |
| Change an existing label                   |   |   |   |
| Add a data item to ConfigMap               |   |   |   |
| Change an existing data item in ConfigMap  |   |   |   |
