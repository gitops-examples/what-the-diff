### Introduction

This repo contains examples of different diff'ing scenarios with `kubectl diff` and Argo CD Diff strategies

### Use Cases

1. [Manual Edits](https://github.com/gitops-examples/what-the-diff/tree/main/manual-edit)


### Notes

Can enable server-side-diff in Argo CD through annotation or by adding:

```
spec:
  controller:
    extraCommandArgs:
      - '--server-side-diff-enabled'
```
