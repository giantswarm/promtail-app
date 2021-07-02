# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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

[Unreleased]: https://github.com/giantswarm/promtail-app/compare/v0.2.1...HEAD
[0.2.1]: https://github.com/giantswarm/promtail-app/compare/v0.2.0...v0.2.1
[0.2.0]: https://github.com/giantswarm/promtail-app/compare/v0.1.1-alpha3...v0.2.0
[0.1.1-alpha3]: https://github.com/giantswarm/promtail-app/compare/v0.1.1-alpha2...v0.1.1-alpha3
[0.1.1-alpha2]: https://github.com/giantswarm/promtail-app/compare/v0.1.1-alpha...v0.1.1-alpha2
[0.1.1-alpha]: https://github.com/giantswarm/promtail-app/compare/v0.1.0-alpha...v0.1.1-alpha
[0.1.0-alpha]: https://github.com/giantswarm/promtail-app/releases/tag/v0.1.0-alpha
