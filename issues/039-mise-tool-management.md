# Issue 039: miseによる開発ツール管理

## Goal

開発とCIで使用するGo、deadcode、staticcheckのバージョンをmiseの設定に集約する。

## Scope

- `mise.toml`で通常のbuild・release用Go 1.23.12を固定する。
- `mise.analysis.toml`で静的解析用Go 1.26.1、deadcode v0.50.0、staticcheck v0.8.1を固定する。既存CIのGo 1.23 / 1.26の検証を維持する。
- CIとreleaseのGoセットアップをmise-actionへ移行し、解析ツールはmiseから実行する。
- READMEへmiseの導入、install、shell activation不要の実行方法と解析用設定を記載する。
- ローカル専用mise設定をGit管理対象外にする。
- 既存のMakefileを開発コマンドの定義として維持し、go.modの最低Goバージョンとtracked helper artifactを変更しない。

## Relevant files

- `mise.toml`, `mise.analysis.toml`, `.gitignore`
- `.github/workflows/ci.yml`, `.github/workflows/release.yml`
- `README.md`, `Makefile`, `go.mod`
- `PLAN.md`

## Acceptance criteria

- `mise install`で通常の開発に使う固定バージョンのGoを導入できる。
- `mise exec -- make check`が成功する。
- `mise -E analysis install`で解析用ツールを導入でき、従来と同じdeadcode / staticcheckを実行できる。
- Linuxの2バージョン検証、Windowsの検証、release packagingを維持する。
- READMEに初回セットアップ、通常のbuild/check、解析ツールの実行手順がある。
- 既存のユーザー変更とhelper artifactを変更・commitしない。

## Verification

- PASS: `mise install`と`mise -E analysis install`で指定バージョンを導入。CIと同じ`MISE_ENV`の空文字 / `analysis`でもGoと解析ツールの選択を確認した。
- PASS: 通常設定の`mise exec -- make check`（test / race / vet / build）。
- PASS: analysis設定の`go test ./...`、`go test -race ./...`、`go vet ./...`、`go build ./...`、`deadcode ./...`、`staticcheck -tests=false ./...`。
- PASS: actionlint v1.7.12で両workflowの構文と式を検証（shellcheckは無効）。
- PASS: 通常設定でrelease archiveを生成し、6 platformのbinary・付属文書と`SHA256SUMS`の一致を確認した。macOSのtarが付加するAppleDouble metadataはpayload検査から除外した。
- PASS: `go.set_gobin = false`によりGOBIN未指定時はGoの既定値を使い、明示したGOBINは維持される。ローカルmise設定のignore、TOML構文、`git diff --check`も確認した。
- PASS: `go.mod`、Makefile、tracked helper artifactは変更なし。
- 未実施: Windows runner上の実行とGitHub Actions自体の実行。更新済みworkflowで確認する。

## Current state

- 完了。通常の開発・releaseは`mise.toml`、Go 1.26での検証と静的解析は`mise.analysis.toml`を使用する。
- ツールのversionはmise設定へ集約し、既存のMakefileとCIの検証項目を維持した。
