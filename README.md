# argo-rollouts-go-client

This is a Go client for the Argo Rollouts API, taken out of [argo-rollouts](https://github.com/argoproj/argo-rollouts) repo, with certain parts stripped, namely `kubectl-argo-rollouts`, `signals` and `apis/rollouts/validation`. 

Today, [argo-rollouts](https://github.com/argoproj/argo-rollouts) cannot be imported in non-intrusive way as it requires an excessive set of `mod replace`.
This repo is meant to be a temporary solution until the [issue](https://github.com/argoproj/argo-rollouts/issues/2587) is resolved.

References: 
- https://github.com/argoproj/argo-rollouts/issues/1837
- https://github.com/argoproj/argo-rollouts/issues/1365

The issue is not limited to argo-rollouts, but occurs in other Argo projects as well.
References:
- https://github.com/argoproj/argo-cd/issues/9198
- https://argo-cd.readthedocs.io/en/stable/user-guide/import/