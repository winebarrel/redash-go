# Changelog

## [2.6.1] - 2025-03-09

### Changed

* Add `IsDraft` to `UpdateQueryInput` struct.

## [2.6.0] - 2025-03-09

### Added

* Add `PublishQuery()/UnpublishQuery()` methods.

## [2.5.0] - 2025-03-08

### Added

* Support regex/enum/query parameter type.

## [2.4.1] - 2025-01-25

### Added

* Add `GetQueryResultByID()` methods.

## [2.4.0] - 2025-01-25

### Changed

* `ExecQueryJSON()` supports `apply_auto_limit`.
* `RefreshQuery()` supports `apply_auto_limit`. (**breaking change**)

## [2.3.1] - 2024-06-14

### Changed

* chore: Add `redash._debugOut` var for testing.

## [2.3.0] - 2024-06-10

### Added

* Add `WaitQueryXXX()` methods.

## [2.2.0] - 2024-06-9

### Changed

* `ExecQueryJSONInput()` supports `max_age=0`. (https://github.com/winebarrel/redash-go/issues/136)

## [2.1.0] - 2023-10-23

### Added

* Add `ExecQueryJSONInput` struct.

## [2.1.0] - 2023-10-23

### Added

* Add `ExecQueryJSONInput` struct.

## [2.0.1] - 2023-10-14

### Changed

* Update doc.go description.

## [2.0.0] - 2023-10-14

### Removed

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
