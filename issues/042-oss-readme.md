# Issue 042: OSS向けREADMEの整理

## Goal

初めて利用する人が概要、導入、使い方を順に理解できるREADMEにする。

## Scope

- 本文を自然な日本語に揃え、概要・主な機能・インストール・使い方・開発・ライセンスを中心に再構成する。
- 公開バイナリとソースからのビルドを分け、利用者向けの説明を前半、開発者向けの説明を後半へ配置する。
- 対象パス、CLIオプション、対応形式を表などで簡潔に整理する。
- 古い実行ファイル名、Windows対応と矛盾する記述、重複や過剰な実装説明を整理する。
- read-only既定、upload opt-in、HTML/SVGの同一origin表示、symlinkの扱いなど利用上の条件を維持する。
- miseによるGo固定、Makefileの検証手順、解析ツール、helper生成の明示的な操作を残す。
- README以外への文書分割やアプリの変更は行わない。

## Relevant files

- `README.md`
- `internal/preview/main.go`, `internal/preview/target.go`
- `mise.toml`, `mise.analysis.toml`, `Makefile`
- `.github/workflows/ci.yml`, `.github/workflows/release.yml`, `scripts/package-release.sh`
- `LICENSE`, `THIRD_PARTY_NOTICES.md`

## Acceptance criteria

- 概要からインストールと最初の起動まで、上から読んで辿れる。
- コマンド・オプション・対応OS・配布ファイルの説明が実装と一致する。
- 既存の機能や使用上の制限を誤って拡張しない。
- 見出し、リンク、コードフェンス、表が正しく表示される。
- 非公開の環境情報や既存のユーザー変更を含めない。

## Verification

- CLIヘルプ、Makefile、mise設定、releaseスクリプトと文中の記載を照合する。
- Markdownの構造、相対リンク、コード例の構文、`git diff --check`を確認する。
- 文書のみの変更のため、アプリのテスト追加や全テスト再実行は行わない。

## Current state

- 完了。ユーザーの希望に合わせて本文を日本語に統一し、利用者向けの導入・使い方を前半、開発・CI・リリースを後半に配置した。
- 公開バイナリの実行名、Windows対応、既定値とオプションを修正し、対象パス・対応形式を表に整理した。
- miseによるGo固定、解析ツール、明示的なhelper生成、利用上の制限を維持した。

## Verification results

- 公開releaseの対象6 platformとarchive名、実際のCLIヘルプ、target parser、Makefile、mise設定、CI/release設定と記載内容を照合した。
- 同梱markedで15見出し、3表、10コードブロックを解析できた。相対リンクの参照先はすべて存在し、bashコードブロックの構文検査は成功した。
- ローカルのMarkdownプレビューで見出し・表・コードブロックの表示とUNCパスのバックスラッシュを確認した。
- `git diff --check`成功。アプリの変更はなく、全テストの再実行は省略した。Windows・SSHへの接続例は実機再実行せず、実装との照合で確認した。
- 既存の`AGENTS.md`と未追跡`docs/`は変更・commit対象に含めない。
