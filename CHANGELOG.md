# Changelog

## [0.5.1](https://github.com/ChrisPJohnstone/go-conventional-commit/compare/v0.5.0...v0.5.1) (2026-10-23)


### Bug Fixes

* lint fixes ([#23](https://github.com/ChrisPJohnstone/go-conventional-commit/issues/23)) ([f3f019f](https://github.com/ChrisPJohnstone/go-conventional-commit/commit/f3f019feecb8425adf5f1821751f83fa4092e067))

## [0.5.0](https://github.com/ChrisPJohnstone/go-conventional-commit/compare/v0.4.0...v0.5.0) (2026-10-23)


### Features

* added type prefix matching ([#21](https://github.com/ChrisPJohnstone/go-conventional-commit/issues/21)) ([63c5307](https://github.com/ChrisPJohnstone/go-conventional-commit/commit/63c53075434423b2bab9c8c2a8bc23f0fecb6f7c))

## [0.4.0](https://github.com/ChrisPJohnstone/go-conventional-commit/compare/v0.3.0...v0.4.0) (2026-10-23)


### Features

* added scope parsing from git branch ([#18](https://github.com/ChrisPJohnstone/go-conventional-commit/issues/18)) ([7a7935b](https://github.com/ChrisPJohnstone/go-conventional-commit/commit/7a7935b974bcdd301ef5f53a6b69eff2eec7f123))


### Bug Fixes

* fixed parseCommitType returning git branch err ([#16](https://github.com/ChrisPJohnstone/go-conventional-commit/issues/16)) ([89b5181](https://github.com/ChrisPJohnstone/go-conventional-commit/commit/89b5181a8a1344acde5ef055ed9f33b6cce3b75c))

## [0.3.0](https://github.com/ChrisPJohnstone/go-conventional-commit/compare/v0.2.0...v0.3.0) (2026-10-23)


### Features

* added logging to gitBranch function ([#14](https://github.com/ChrisPJohnstone/go-conventional-commit/issues/14)) ([ce3891d](https://github.com/ChrisPJohnstone/go-conventional-commit/commit/ce3891db98d0163e14feea7499cabac343bfb3a0))
* added scope flag ([#15](https://github.com/ChrisPJohnstone/go-conventional-commit/issues/15)) ([5b71366](https://github.com/ChrisPJohnstone/go-conventional-commit/commit/5b713662bd8c6f1c57dc5210d80b7be42edd4a43))
* added verbose logging ([#11](https://github.com/ChrisPJohnstone/go-conventional-commit/issues/11)) ([6838e6a](https://github.com/ChrisPJohnstone/go-conventional-commit/commit/6838e6ab02a20b1eb6f7b87c1165f9977a8f8416))


### Bug Fixes

* use type fallback if failing to get git branch ([#13](https://github.com/ChrisPJohnstone/go-conventional-commit/issues/13)) ([7b497e5](https://github.com/ChrisPJohnstone/go-conventional-commit/commit/7b497e5a9bcd5671dfa45504958f7e665b1dd231))

## [0.2.0](https://github.com/ChrisPJohnstone/go-conventional-commit/compare/v0.1.0...v0.2.0) (2026-10-23)


### Features

* added default type detection from branch ([#9](https://github.com/ChrisPJohnstone/go-conventional-commit/issues/9)) ([96effd5](https://github.com/ChrisPJohnstone/go-conventional-commit/commit/96effd5ca2dcb432409cf6681868bd3f9f7b203b))

## 0.1.0 (2026-10-22)

### Features

* added cobra entrypoint ([f8f7d64](https://github.com/ChrisPJohnstone/go-conventional-commit/commit/f8f7d64b23f6bffd61d641dd48bceea8d63f38d6))
* added commit message builder ([809ad1e](https://github.com/ChrisPJohnstone/go-conventional-commit/commit/809ad1e7514d4b9c8a75c1c50d2bf47dea3ebb1f))
* added commit types to help message ([b166224](https://github.com/ChrisPJohnstone/go-conventional-commit/commit/b166224351ba41214ea02e8d487744f9ab080b48))
* added entrypoint ([44e575f](https://github.com/ChrisPJohnstone/go-conventional-commit/commit/44e575f3d01641392dd55173430712cdb7e8fd94))
* added min 1 positional arg ([485ac98](https://github.com/ChrisPJohnstone/go-conventional-commit/commit/485ac98c8c648c06d7f5dbdb60ac9b18992dee2f))
* added shorthand for type flag ([1246ab3](https://github.com/ChrisPJohnstone/go-conventional-commit/commit/1246ab33fb278e0418505919b1fe7b85e6adb440))
* added type flag ([a9ec060](https://github.com/ChrisPJohnstone/go-conventional-commit/commit/a9ec060f5e93e98c73637ae61cca870baea1b6d5))
* added type validation ([8bf34ae](https://github.com/ChrisPJohnstone/go-conventional-commit/commit/8bf34ae9fe75a6b11987fa9e1e0d187106a0bfd6))
* swapped to RunE ([0012e32](https://github.com/ChrisPJohnstone/go-conventional-commit/commit/0012e323083981c7a2929092fe95025ffebed8ed))
