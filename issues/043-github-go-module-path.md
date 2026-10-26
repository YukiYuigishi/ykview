# Issue 043: Go module pathをGitHub repositoryへ揃える

## Goal

Go module pathを`github.com/YukiYuigishi/ykview`へ変更し、cloneせずに`go install`でCLIを導入できるようにする。

## Scope

- `go.mod`と内部packageのimportをcanonical module pathへ変更する。
- READMEに`go install github.com/YukiYuigishi/ykview/cmd/ykview@latest`を案内する。
- CLI構成や生成済みremote helper artifactは変更しない。pushやrelease tagの作成は含めない。

## Relevant files

- `go.mod`
- `cmd/ykview/main.go`
- `cmd/ykview-helper/main.go`
- `internal/preview/remote.go`
- `README.md`
- `PLAN.md`

## Acceptance criteria

- moduleと内部importが`github.com/YukiYuigishi/ykview`へ統一されている。
- test、race、vet、buildが成功する。
- canonical package pathを指定したローカルの`go install`で`ykview`を生成できる。
- READMEのinstallコマンドが`cmd/ykview`の配置と一致し、Goの要件とPATH設定を説明している。

## Verification

- `go test ./...`、`go test -race ./...`、`go vet ./...`、`go build ./...`。
- 一時的な`GOBIN`へ`go install github.com/YukiYuigishi/ykview/cmd/ykview`し、生成したCLIのhelpを確認する。
- 旧module/importの残存検索と`git diff --check`。
- 公開先の`@latest`は変更を含むversionの公開後に検証する。

## Current state

- 完了。作業開始時のworktreeはclean。
- module宣言と3か所の内部importを`github.com/YukiYuigishi/ykview`へ統一し、READMEにGoの要件、installコマンド、PATH設定を追加した。
- 公開先の`@latest`は未検証。この変更を含むversionの公開が必要で、pushやtagの作成は実施していない。

## Verification results

- 環境にあるGo 1.27.1（linux/amd64）を使用した。`mise`とprojectで固定したGo 1.27.2は利用できなかったため、固定versionでの再検証は未実施。
- `go test ./...`、`go test -race ./...`、`go vet ./...`、`go build ./...`は成功。
- 初回の通常テストはsandboxのlocalhost待受制限で失敗した。通常テストとrace検査は承認を得たsandbox外の再実行で成功した。
- 一時的な`GOBIN`へ`go install github.com/YukiYuigishi/ykview/cmd/ykview`が成功し、`go version -m`でpackage/moduleがcanonical pathになっていることを確認した。
- 生成したCLIの`-h`で既存オプションの表示を確認した。help要求時は既存実装に従ってexit 1となる。
- 旧module/importの残存なし。変更したGoファイルの`gofmt`、`git diff --check`は成功。
- 生成済みremote helper artifactに変更なし。
