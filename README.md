### Introduction

This repo contains examples of different diff'ing scenarios with `kubectl diff` and Argo CD Diff strategies

### Use Cases

1. [Manual Edits](https://github.com/gitops-examples/what-the-diff/tree/main/manual-edit)
2. [Mutations](https://github.com/gitops-examples/what-the-diff/tree/main/server-side)

### Notes

Can enable server-side-diff in Argo CD through annotation or by adding:

```
spec:
  controller:
    extraCommandArgs:
      - '--server-side-diff-enabled'
```

### References

* [How Apply calculates differences and merges changes](https://kubernetes.io/docs/tasks/manage-kubernetes-objects/declarative-config/#how-apply-calculates-differences-and-merges-changes)
* Interesting comments in Kubernetes issues:
  * [kubectl diff not producing any changes when fields are removed from manifest](https://github.com/kubernetes/kubectl/issues/1403#issuecomment-1501627912)
  * [kubectl diff unable to detect changes made by kubectl edit](https://github.com/kubernetes/kubectl/issues/1744#issuecomment-2885607523)
  * [Server-side apply: migration from client-side apply leaves stuck fields in the object](https://github.com/kubernetes/kubernetes/issues/99003)

### OpenShift GitOps

These use cases where run using OpenShift GitOps 1.18 (Argo CD 3.1)

It uses the default setup of `cluster-admins` group, to create this group run the following:

```
oc adm groups new cluster-admins admin
```

You also need to give the application-controller SA `cluster-admin` rights:

```
oc adm policy add-cluster-role-to-user cluster-admin -z openshift-gitops-argocd-application-controller -n openshift-gitops
```
