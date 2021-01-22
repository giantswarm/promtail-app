# Promtail chart

[![CircleCI](https://circleci.com/gh/giantswarm/promtail-app.svg?style=shield)](https://circleci.com/gh/giantswarm/promtail-app)

Promtail is a logging agent meant to work with Loki logging server. This chart is meant
primarily to be used with our [Loki app](https://github.com/giantswarm/loki-app).

Promtail works over the network using HTTP protocol, so can be used on any cluster
to forward its logs to a Loki instance.

Please note that you can run multiple `promtail` instances on the same cluster,
but it's up to you to provide a reasonable config for them (and ie. avoid
duplication). The default config forwards logs from all Kubernetes pods
and from `systemd-journald` logging service in the operating system.

## Installing

There are 3 ways to install this app onto a tenant cluster.

1. [Using our web interface](https://docs.giantswarm.io/reference/web-interface/app-catalog/)
2. [Using our API](https://docs.giantswarm.io/api/#operation/createClusterAppV5)
3. Directly creating the App custom resource on the Control Plane.

## Configuring

### values.yaml
**This is an example of a values file you could upload using our web interface.**
```
# values.yaml

```

## Source code origin

The source code in `helm/promtail` is a git-subtree coming from the
<https://github.com/giantswarm/grafana-helm-charts-upstream>. Giant Swarm uses that
repository to track and adjust or charts maintained by Grafana Labs.

## Credit

* <https://github.com/grafana/helm-charts/tree/main/charts/loki-distributed>
