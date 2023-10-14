# Changelog

## [Unreleased]

* Drop v8 support.

## [1.4.1] - 2023-09-14

### Added

* Add job status constants.

### Changed

* Update modules.

## [1.4.0] - 2023-04-25

### Added

* Change `UpdateInput.Tags` type. (`[]string` -> `*[]string`)
* Remove `Client.RemoveQueryTags()`.

## [1.3.0] - 2023-04-25

### Added

* Add `Client.RemoveQueryTags()`.

## [1.2.1] - 2023-04-25

### Changed

* Remove "omitempty" from `UpdateQueryInput.Tags`.

## [1.2.0] - 2023-04-25

### Added

* Add `Client.UpdateUser()`.

### Changed

* Change generator tool path.

## [1.1.2] - 2023-04-23

### Changed

* Add `AlertOptions.Muted`.

## [1.1.1] - 2023-04-23

### Changed

* Remove "omitempty" from `UpdateAlertInput.Rearm`.

## [1.1.0] - 2023-04-23

### Changed

- Support alert template fields.

## [1.0.1] - 2023-04-16

### Changed

- Rename docker-compose.yml to compose.yaml.
- Update modules.

## [1.0.0] - 2023-03-07

### Added

- First stable release.

<!-- cf. https://keepachangelog.com/ -->
