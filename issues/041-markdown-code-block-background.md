# Issue 041: Markdownコードブロックの背景色

## Goal

Markdown内のコードブロックを本文から見分けやすくする。

## Scope

- Markdown previewのコードブロックに、画面全体と同じ無彩色の背景、余白、控えめな枠線と角丸を設定する。
- 色はdirectory viewerと同じsystem colorを基準にし、背景は`color-mix(in srgb, CanvasText 7%, Canvas)`、枠線は`#8885`、文字は`CanvasText`を使用する。
- ライト・ダークの各配色で文字を読みやすくする。
- Source表示にも同じ配色を適用する。
- Mermaidの描画領域はコードブロックの装飾対象から除外する。
- JavaScript、構文ハイライトの追加、inline code、他のviewerは変更しない。

## Relevant files

- `internal/preview/templates.go`（`markdownTemplate`のCSS）

## Acceptance criteria

- ライト・ダークの両方で、コードブロックの背景を本文と区別できる。
- 長い行の横スクロール、インデント、リスト・引用内の配置を維持する。
- Mermaid描画とSource切替に影響しない。
- 狭い画面でもページ全体に横スクロールが発生しない。

## Verification

- PASS: 配色調整後のライト・ダーク表示をブラウザで再確認した。背景のRGB成分は等しく、文字色はCanvasText、枠線は既存UIと同じ`#8885`となる。ライト配色は実際のHTMLから作成した一時fixtureでcolor-schemeを固定して確認した。
- PASS（初回実装時）: 通常幅と375px幅で余白・枠線を確認した。今回の調整では色の指定のみを変更した。
- PASS: 375px幅でも本文とページの横overflowはなく、長いコードブロックだけが横スクロールした。別のブロック・本文・ページの横位置は変わらず、インデントとリスト・引用内の配置も維持した。
- PASS: MermaidのSVG描画と装飾対象からの除外を確認した。Mermaid失敗時のSource表示にも背景色が付き、表示切替後も内容を維持した。
- PASS: 配色調整後の`go test ./internal/preview`（miseのGo 1.27.2）、`git diff --check`。

## Current state

- 完了。既存UIと同じsystem colorを使用し、ライト・ダークに追従する無彩色の背景・枠線・文字色に揃えた。余白・角丸・スクロール・Mermaid除外は維持した。
