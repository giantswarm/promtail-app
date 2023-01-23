# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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

[Unreleased]: https://github.com/giantswarm/promtail-app/compare/v1.0.0...HEAD
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
