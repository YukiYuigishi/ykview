# Issue 041: Markdownコードブロックの背景色

## Goal

Markdown内のコードブロックを本文から見分けやすくする。

## Scope

- Markdown previewのコードブロックに、青みのあるグレーの背景、余白、控えめな枠線と角丸を設定する。
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

- PASS: 匿名fixtureを通常幅と375px幅で表示し、ライト・ダークの背景色・余白・枠線を確認した。ライト配色は実際のHTMLから作成した一時fixtureでcolor-schemeを固定して確認した。
- PASS: 375px幅でも本文とページの横overflowはなく、長いコードブロックだけが横スクロールした。別のブロック・本文・ページの横位置は変わらず、インデントとリスト・引用内の配置も維持した。
- PASS: MermaidのSVG描画と装飾対象からの除外を確認した。Mermaid失敗時のSource表示にも背景色が付き、表示切替後も内容を維持した。
- PASS: 文字と背景のコントラスト比はライト12.64、ダーク12.10。
- PASS: `go test ./internal/preview`（miseのGo 1.27.2）、`git diff --check`。

## Current state

- 完了。`markdownTemplate`のCSSのみを変更した。ライトは薄い青みのグレー、ダークは濃い青みのグレーを使用する。
