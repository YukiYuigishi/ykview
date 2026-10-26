# ykview

ykviewは、ローカルやSSH接続先のファイルをブラウザーで閲覧するツールです。Markdown（Mermaid）、画像、音声・動画、CSV/TSV、JSON Lines、ソースコードのプレビュー、ディレクトリ内の名前検索、ダウンロードに対応します。

SSH接続にはシステムの`ssh`を使い、`~/.ssh/config`の設定を利用できます。接続先にGoやHTTPサーバーを用意する必要はありません。

## インストール

Go 1.23以降でインストールできます。

```bash
go install github.com/YukiYuigishi/ykview/cmd/ykview@latest
```

インストール先の`GOBIN`（未設定なら`GOPATH/bin`、通常は`~/go/bin`）をPATHへ追加してください。

macOS・Linux・Windowsのamd64/arm64向けバイナリは[GitHub Releases](https://github.com/YukiYuigishi/ykview/releases/latest)から取得できます。OS・CPUに合うアーカイブを展開し、`ykview`（Windowsでは`ykview.exe`）をPATHの通ったディレクトリへ配置します。

## 使い方

```bash
ykview .                          # ローカルのディレクトリ
ykview ./path/to/file             # ローカルのファイル
ykview host                       # SSH接続先のホーム
ykview host:~/projects            # SSH接続先のホームからの相対パス
ykview host:/var/log              # SSH接続先の絶対パス
ykview -write host:~/incoming     # アップロードを有効化
```

起動するとブラウザーが開き、URLを端末に表示します。終了は`Ctrl+C`です。オプションは対象より前に指定します。

ローカルの相対パスには`./`か`../`を付けてください。`host:path`はSSH接続先のホームを基準にします。ファイルを指定すると、その親ディレクトリを配信対象としてファイルを開きます。

Windowsのローカルパスは`C:\Users\me\Documents`、Windows SSHサーバーのパスは`host:C:/Users/me/Documents`のように指定します。SSH接続先にはPOSIX環境で`sh`、WindowsでOpenSSH ServerとPowerShell 5.1以降が必要です。

| オプション | 説明 |
| --- | --- |
| `-addr アドレス` | 待受アドレス。既定は`127.0.0.1:8080`。空きポートを選ぶ場合は`127.0.0.1:0` |
| `-open=false` | ブラウザーの自動起動を無効化 |
| `-write` | ファイル・ディレクトリのアップロードを有効化 |
| `-max-upload-size バイト数` | アップロード要求のサイズ上限。既定の`0`は無制限 |
| `-v` | HTTP要求のログを出力 |

アップロードは表示中のディレクトリへ行い、同名の通常ファイルを上書きします。音声・動画の再生可否はブラウザーの対応コーデックに依存します。

## 利用上の注意

- 個人利用向けで認証機能はありません。既定ではループバックアドレスで待ち受けます。
- HTML・SVGはスクリプトを実行できるため、信頼できるファイルだけを開いてください。
- シンボリックリンク経由で、指定したディレクトリの外を読み取れる場合があります。

## 開発

Goのバージョンは[mise.toml](mise.toml)で管理します。[mise](https://mise.jdx.dev/getting-started.html)とGNU Makeを用意し、リポジトリ内で実行します。

```bash
mise trust
mise install
mise exec -- make build
mise exec -- make check
```

ビルド結果は`bin/ykview`（Windowsでは`bin/ykview.exe`）です。`make check`はテスト、race検査、`go vet`、ビルドを実行します。接続先で使う補助実行ファイルを更新するときだけ、`mise exec -- make generate`で再生成してください。

不具合の報告や機能の提案は[GitHub Issues](https://github.com/YukiYuigishi/ykview/issues)へお願いします。

## ライセンス

[MIT License](LICENSE)。同梱ライブラリは[THIRD_PARTY_NOTICES.md](THIRD_PARTY_NOTICES.md)を参照してください。
