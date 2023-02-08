# Promtail App

[![CircleCI](https://circleci.com/gh/giantswarm/promtail-app.svg?style=shield)](https://circleci.com/gh/giantswarm/promtail-app)

Giant Swarm offers Promtail as a [managed app](https://docs.giantswarm.io/changes/managed-apps/) which can be installed in any clusters.
It deploys a [Promtail agent](https://grafana.com/docs/loki/latest/clients/promtail/) used to collect and ship logs (through the HTTP protocol) to a private Grafana Loki instance.

Please note that you can run multiple `promtail` instances on the same cluster (albeit in different namespaces),
but it's up to you to provide a reasonable config for them (and ie. avoid
duplication of scrapes).

**Table of Contents:**

- [Install](#install)
- [Upgrading](#upgrading)
- [Configuration](#configuration)
- [Limitations](#limitations)
- [Credit](#credits)

## Install

There are several ways to install this app onto a workload cluster.

- [Using GitOps to instantiate the App](https://docs.giantswarm.io/advanced/gitops/#installing-managed-apps)
- [Using our web interface](https://docs.giantswarm.io/ui-api/web/app-platform/#installing-an-app).
- By creating an [App resource](https://docs.giantswarm.io/ui-api/management-api/crd/apps.application.giantswarm.io/) in the management cluster as explained in [Getting started with App Platform](https://docs.giantswarm.io/app-platform/getting-started/).

## Upgrading

### Upgrading an existing Release to a new major version
A major chart version change (like v0.5.0 -> v1.0.0) indicates that there is an incompatible breaking change needing manual actions.

### From 0.x to 1.x

The structure of the values changed in 1.0.0 as we now rely on helm chart dependency mechanism to manage the application.

Before you upgrade, you must edit your values (in both configmaps and secrets) to follow the new structure from:

```yaml
# old values.yaml structure
resources:
  limits:
    cpu: 200m
    memory: 256Mi
  requests:
    cpu: 100m
    memory: 128Mi
```

To:

```yaml
# new values.yaml structure
promtail:
  resources:
    limits:
      cpu: 200m
      memory: 256Mi
    requests:
      cpu: 100m
      memory: 128Mi
```

As you can see, in the new format, all properties coming from the subchart as described [below](#configuration), must be moved under a `promtail:` section.

## Configuration

By default, Promtail is configured to read logs from all Kubernetes pods
and from the `systemd-journald` logging service running on the machine the agent is deployed on.

As this application is build upon the Grafana promtail upstream chart as a dependency, most of the values to override can be found [here](https://github.com/grafana/helm-charts/blob/promtail-6.0.2/charts/promtail/README.md#values).

Examples:
- [Journald configuration](./examples/values-journald.yaml)
- [Kubernetes audit logs configuration](./examples/values-kubernetes-audit-logs.yaml)

The Giant Swarm default values can also be overriden but we advise against it.

### values.yaml

**This is an example of a values file you could upload using our web interface.**

```yaml
# values.yaml
promtail:
  resources:
    limits:
      cpu: 200m
      memory: 256Mi
    requests:
      cpu: 100m
      memory: 128Mi
```

### Sample App CR and ConfigMap for the management cluster

If you have access to the Kubernetes API on the management cluster, you could create
the App CR and ConfigMap directly.

Here is an example that would install the app to
workload cluster `abc12`:

```yaml
# appCR.yaml
apiVersion: application.giantswarm.io/v1alpha1
kind: App
metadata:
  labels:
  name: promtail-app
  # workload cluster resources live in a namespace with the same ID as the
  # workload cluster.
  namespace: abc12
spec:
  name: promtail-app
  namespace: promtail
  catalog: giantswarm
  version: 2.2.0
  userConfig:
    configMap:
      name: promtail-app-user-values
      namespace: abc12
```

```yaml
# user-values-configmap.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: nginx-ingress-controller-app-user-values
  namespace: abc12
data:
  values: |
    promtail:
      # -- Resource requests and limits
      resources:
        limits:
          cpu: 200m
          memory: 256Mi
        requests:
          cpu: 100m
          memory: 128Mi
```

See our [full reference on how to configure apps](https://docs.giantswarm.io/app-platform/app-configuration/) for more details.

## Limitations

The application and its default values have been tailored to work inside Giant Swarm clusters and to connect to our managed [Loki app](https://github.com/giantswarm/loki-app).
If you want to use it for any other scenario, know that you might need to adjust some values.

## Credit

This application is installing the upstream chart below with defaults to ensure it runs smoothly in Giant Swarm clusters.

* <https://github.com/grafana/helm-charts/tree/main/charts/promtail>
