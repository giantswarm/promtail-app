# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [1.5.4] - 2024-07-24

### Fixed

- Allow traffic to nginx-ingress-controller (needed when LB is skipped).

## [1.5.3] - 2024-03-12

### Changed

- Set RAM limit to 2x requests, so it has a bit of extra space for pikes and prevent crashes.

## [1.5.2] - 2024-03-12

### Changed

- Adjust CPU settings: remove CPU adjustement by VPA and set CPU request/limit according to actual usage.

## [1.5.1] - 2024-02-09

### Changed

- Upgrade promtail chart: 6.15.3 => 6.15.5
- Upgrade promtail: 2.9.2 => 2.9.3 - see [changelog](https://github.com/grafana/loki/blob/main/CHANGELOG.md) for more information.

## [1.5.0] - 2024-01-22

### Added

- Add vpa resource in the helm chart's templates.

## [1.4.2] - 2024-01-09

### Changed

- Configure `gsoci.azurecr.io` as the default container image registry.

## [1.4.1] - 2023-12-05

### Changed

- Upgrade promtail chart: 6.15.1 => 6.15.3
- Upgrade promtail: 2.8.4 => 2.9.2 - see [changelog](https://github.com/grafana/loki/blob/main/CHANGELOG.md) for more information.
- Updated CiliumNetworkPolicy to support in-cluster loki endpoints

## [1.4.0] - 2023-09-13

### Changed

- Enabled networkpolicies.
- Upgrade promtail chart: 6.14.1 => 6.15.1
- Upgrade promtail: 2.8.3 => 2.8.4

## [1.3.0] - 2023-08-07

### Changed

- Upgrade promtail chart: 6.11.7 => 6.14.1
    - Added support for VPA
- Upgrade promtail: 2.8.2 => 2.8.3

## [1.2.0] - 2023-07-24

### Changed

- Upgrade promtail chart: 6.11.2 => 6.11.7

### Added

- Enable servicemonitor.

## [1.1.1] - 2023-06-28

### Added

- Add Kyverno Policy Exceptions.

## [1.1.0] - 2023-05-15

- Upgrade promtail chart: 6.10.0 => 6.11.2
- Upgrade promtail: 2.7.4 => 2.8.2

## [1.0.3] - 2023-04-17

### Changed

- Upgrade promtail chart: 6.8.3 => 6.10.0
- Upgrade promtail: 2.7.3 => 2.7.4

## [1.0.2] - 2023-03-09

### Removed

- Remove namespace singleton temporarily, as it has some issues

### Changed

- Updated upstream chart from v6.8.2 to v6.8.3
- Updated app from 2.7.2 to 2.7.3

## [1.0.1] - 2023-01-25

### Fixed

- Enable `psp` by default

### Added

- Added testing guidelines

## [1.0.0] - 2022-12-29

### Modified

- Migrated to chart dependency
- ⚠ Major upgrade, [breaking changes](https://github.com/giantswarm/promtail-app/blob/master/README.md#from-0x-to-1x)
  - values structure changes. We rely on a subchart, meaning all of previous setup goes to a `promtail` section

## [0.5.0] - 2022-12-05

### Modified

- Added headless service with monitoring label

## [0.4.1] - 2022-07-06

### Modified

- latest upstream chart (6.0.2)
- update values schema

### Fixed

- values schema accepts strings in all default-null fields, like "fullnameOverride"

## [0.4.0] - 2022-06-23

### Modified

- Upgrade `promtail` to 2.5.0 and upstream chart version to 6.0.0
- **Breaking**: Upgrade requires manual intervention.

### Changes required in your `values.yaml` file:
- client config is more similar to promtail config file, and is done on `config.clients`.

## [0.3.2] - 2021-10-22

- Update metadata

## [0.3.1] - 2021-10-20

- Update metadata and icon

## [0.3.0] - 2021-07-09

### Added

- Enable sli monitoring and metric scraping
- Set deployment resources (Requests/Limits)

## [0.2.1] - 2021-06-14

### Fixed

- Re-release of [0.2.0] as it did not contain the advertised promtail upgrade from [2.2.0](https://github.com/grafana/loki/releases/tag/v2.2.0) to [2.2.1](https://github.com/grafana/loki/releases/tag/v2.2.1).

## [0.2.0] - 2021-06-08

### Changed

- Upgrade to promtail 2.2.1

## [0.1.1-alpha3] - 2021-03-16

### Removed

- unsupported and removed in upstream `extra_client_config` and `extra_scrapre_configs` are removed as well

## [0.1.1-alpha2] - 2021-03-15

### Changed

- Upgrade to promtail 2.2.0

## [0.1.1-alpha] - 2021-03-04

### Added

- Annotation for routing alerts to team halo

## [0.1.0-alpha] - 2021-01-25

### changes
- first release

[Unreleased]: https://github.com/giantswarm/promtail-app/compare/v1.5.4...HEAD
[1.5.4]: https://github.com/giantswarm/promtail-app/compare/v1.5.3...v1.5.4
[1.5.3]: https://github.com/giantswarm/promtail-app/compare/v1.5.2...v1.5.3
[1.5.2]: https://github.com/giantswarm/promtail-app/compare/v1.5.1...v1.5.2
[1.5.1]: https://github.com/giantswarm/promtail-app/compare/v1.5.0...v1.5.1
[1.5.0]: https://github.com/giantswarm/promtail-app/compare/v1.4.2...v1.5.0
[1.4.2]: https://github.com/giantswarm/promtail-app/compare/v1.4.1...v1.4.2
[1.4.1]: https://github.com/giantswarm/promtail-app/compare/v1.4.0...v1.4.1
[1.4.0]: https://github.com/giantswarm/promtail-app/compare/v1.3.0...v1.4.0
[1.3.0]: https://github.com/giantswarm/promtail-app/compare/v1.2.0...v1.3.0
[1.2.0]: https://github.com/giantswarm/promtail-app/compare/v1.1.1...v1.2.0
[1.1.1]: https://github.com/giantswarm/promtail-app/compare/v1.1.0...v1.1.1
[1.1.0]: https://github.com/giantswarm/promtail-app/compare/v1.0.3...v1.1.0
[1.0.3]: https://github.com/giantswarm/promtail-app/compare/v1.0.2...v1.0.3
[1.0.2]: https://github.com/giantswarm/promtail-app/compare/v1.0.1...v1.0.2
[1.0.1]: https://github.com/giantswarm/promtail-app/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/giantswarm/promtail-app/compare/v0.5.0...v1.0.0
[0.5.0]: https://github.com/giantswarm/promtail-app/compare/v0.4.1...v0.5.0
[0.4.1]: https://github.com/giantswarm/promtail-app/compare/v0.4.0...v0.4.1
[0.4.0]: https://github.com/giantswarm/promtail-app/compare/v0.3.2...v0.4.0
[0.3.2]: https://github.com/giantswarm/promtail-app/compare/v0.3.1...v0.3.2
[0.3.1]: https://github.com/giantswarm/promtail-app/compare/v0.3.0...v0.3.1
[0.3.0]: https://github.com/giantswarm/promtail-app/compare/v0.2.1...v0.3.0
[0.2.1]: https://github.com/giantswarm/promtail-app/compare/v0.2.0...v0.2.1
[0.2.0]: https://github.com/giantswarm/promtail-app/compare/v0.1.1-alpha3...v0.2.0
[0.1.1-alpha3]: https://github.com/giantswarm/promtail-app/compare/v0.1.1-alpha2...v0.1.1-alpha3
[0.1.1-alpha2]: https://github.com/giantswarm/promtail-app/compare/v0.1.1-alpha...v0.1.1-alpha2
[0.1.1-alpha]: https://github.com/giantswarm/promtail-app/compare/v0.1.0-alpha...v0.1.1-alpha
[0.1.0-alpha]: https://github.com/giantswarm/promtail-app/releases/tag/v0.1.0-alpha
