# Issue 044: READMEを必要な情報へ絞る

## Goal

概要、インストール、基本操作、開発に必要な情報を短く読めるREADMEにする。

## Scope

- 機能説明と操作例の重複を削り、詳細な形式一覧、転送の内部仕様、CI・release・解析ツールの説明を省く。
- `go install`、公開バイナリ、主要な対象指定とオプション、SSH要件、利用上の注意、開発手順、ライセンスを残す。
- アプリの実装や設定は変更しない。

## Relevant files

- `README.md`
- `go.mod`, `mise.toml`, `Makefile`
- `internal/preview/main.go`, `internal/preview/target.go`
- `PLAN.md`

## Acceptance criteria

- 現行READMEの半分以下の行数で、導入から基本操作まで辿れる。
- module path、コマンド、既定値が実装と一致する。
- ローカル相対パスの指定方法、SSH要件、uploadのopt-inを維持する。
- 認証なし、HTML/SVGのスクリプト実行、symlinkの扱いを簡潔に残す。
- Markdownの構造と相対リンクが正しい。

## Verification

- 実装・設定と記載内容を照合する。
- Markdownの構造、リンク、コード例の構文と`git diff --check`を確認する。
- 文書のみのためアプリのテストは再実行しない。

## Current state

- 完了。作業開始時のworktreeはclean。
- READMEを199行から69行へ短縮した。重複する機能説明、詳細な形式一覧、転送内部の説明、CI・release・解析ツールの手順を削除した。
- 導入、主要な対象指定と全CLIオプション、SSH要件、uploadのopt-inと上書き動作、利用上の注意、開発のbuild/check/generate、ライセンスを残した。

## Verification results

- `go.mod`、CLIのオプション・既定値、target parser、Makefileと記載を照合した。
- 3つのbashコードブロックの構文検査、コードフェンスの対応、相対リンク3件の存在確認、canonical install pathの確認は成功。
- Markdownの見出し・リスト・表を確認し、`git diff --check`は成功。
- 文書のみの変更のため、アプリのテストは再実行していない。外部リンク先への接続とmiseの実行は未検証。
