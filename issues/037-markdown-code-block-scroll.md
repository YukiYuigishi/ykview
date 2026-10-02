# 037: Scroll Markdown code blocks independently

## Goal

Markdownの長いコード行を横スクロールしても、他の本文やコードブロックが一緒に動かないようにする。

## Scope

- Markdown内の各コードブロックに独立した横スクロールを設定する。
- コードの改行、インデント、Source表示、Mermaid描画を維持する。
- 表などコード以外の既存のoverflow処理は変更しない。

## Relevant files

- `internal/preview/templates.go`（`markdownTemplate`）

## Acceptance criteria

- 長いコード行がある場合、そのコードブロック内で横スクロールできる。
- コードブロックの幅によって本文全体やページに横スクロールが発生しない。
- 別のコードブロック、段落、見出しの横位置は変わらない。
- リスト・引用内のコードブロックでも同じ挙動となる。
- 短いコード、Source表示、Mermaid描画に影響しない。

## Verification

- Browserで長い行、短い行、複数ブロック、リスト・引用内のブロックを含む匿名fixtureを表示する。
- 通常幅と狭い表示幅で各領域の幅、overflow、スクロール時の本文位置を確認する。
- Source表示とMermaid描画を確認する。
- `make check`を実行する。

## Current state

完了。Markdownの既存のoverflow指定に`#rendered pre`を追加した。コードブロック自身がscroll containerとなり、長い行のoverflowを局所化する。

## Verification results

- `make check`（test、race、vet、CLI build）: passed。
- Browserの通常幅（1800px）と狭い幅（375px）で、本文の`scrollWidth`と`clientWidth`が一致した。ページ全体にも横方向のoverflowは発生しない。
- 長いコードブロックには横方向のoverflowがあり、短いコードブロックにはない。`white-space: pre`とインデントを維持した。
- 通常のブロックを620px、リスト内を480px、引用内を400px横スクロールし、それぞれのブロックだけが動くことを確認した。本文・ページの横スクロール位置は0のままで、見出し・段落・他のブロックの位置も維持した。
- MermaidのSVG描画と、描画失敗時のSource表示・表示切替を確認した。
