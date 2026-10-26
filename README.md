# ykview

ykviewは、ローカルまたはSSH接続先のファイルをブラウザーで閲覧するWebファイラーです。ファイルの検索、プレビュー、ダウンロードに対応し、`-write`を指定するとアップロードもできます。

SSH接続にはシステムの`ssh`を使います。接続先にGoやHTTPサーバー、常駐プロセスを用意する必要はありません。

## 主な機能

- ローカル・SSH接続先のディレクトリ閲覧
- `~/.ssh/config`、Hostエイリアス、IdentityFile、ProxyJump、ssh-agentを使ったSSH接続
- 現在のディレクトリ内でのファイル名検索と種類による絞り込み
- Markdown、Mermaid、画像、音声、動画、CSV/TSV、JSON Lines、ソースコードなどのプレビュー
- ファイル単位のダウンロードと、複数ファイル・ディレクトリのZIPダウンロード
- `-write`を指定した場合のアップロードと一時停止・再開

## インストール

### go install

Go 1.23以降があれば、次のコマンドでインストールできます。

```bash
go install github.com/YukiYuigishi/ykview/cmd/ykview@latest
ykview .
```

インストール先は`GOBIN`（未設定なら`GOPATH/bin`、通常は`~/go/bin`）です。このディレクトリをPATHへ追加してください。

### 公開バイナリ

[最新のGitHub Release](https://github.com/yukiyuigishi/ykview/releases/latest)から、OSとCPUに合うアーカイブをダウンロードして展開します。macOS・Linux向けは`.tar.gz`、Windows向けは`.zip`です。アーカイブ名は`ykview-v0.1.0-darwin-arm64.tar.gz`の形式で、展開後の実行ファイルは`ykview`または`ykview.exe`です。

macOS・Linux・Windowsの各`amd64`（x86-64）・`arm64`向けに公開しています。Apple Silicon搭載のMacでは`darwin-arm64`を選びます。SSH接続を使う場合は、OpenSSHクライアント（`ssh`）も必要です。

展開先のディレクトリで、macOS・Linuxでは次を実行します。

```bash
./ykview .
```

WindowsではPowerShellから起動します。

```powershell
.\ykview.exe .
```

実行ファイルをPATHの通ったディレクトリへ配置すると、以後は`ykview`または`ykview.exe`という名前で起動できます。

### ソースからビルド

Go 1.27.2をmiseで管理します。[mise](https://mise.jdx.dev/getting-started.html)とGNU Makeを用意し、次を実行します。

```bash
git clone https://github.com/yukiyuigishi/ykview.git
cd ykview
mise trust
mise install
mise exec -- make build
./bin/ykview .
```

GNU Makeを使わない場合は、miseのセットアップ後に`mise exec -- go build -o ykview ./cmd/ykview`でビルドできます。Windowsでは出力名を`ykview.exe`にしてください。

## 使い方

```text
ykview [オプション] 対象
```

オプションは対象より前に指定します。起動すると既定のブラウザーを開き、URLを端末に表示します。終了するには端末で`Ctrl+C`を押します。

以下の例は実行ファイルがPATH上にある場合です。PATHにない場合は、`./ykview`、`./bin/ykview`、Windowsでは`.\ykview.exe`のように実行ファイルへのパスを指定します。

### 対象パス

| 対象 | 表示する場所 |
| --- | --- |
| `.`、`./path`、`../path`、`~/path`、`/absolute/path` | ローカルのディレクトリまたはファイル |
| `C:\Users\me\Documents`、`\\server\share\path` | Windows上のローカルパス |
| `host`、`host:~` | SSH接続先のホームディレクトリ |
| `host:path`、`host:~/path` | SSH接続先のホームを基準にしたパス |
| `host:/absolute/path` | POSIX環境のSSH接続先の絶対パス |
| `host:C:/path` | Windows SSHサーバーのドライブ絶対パス |
| `host:\\server\share\path` | Windows SSHサーバーのUNCパス |

`Documents`のような名前だけの指定はSSH接続先として扱います。ローカルの相対パスには`./`または`../`を付けてください。ファイルを指定すると、その親ディレクトリを配信対象にして、指定したファイルを開きます。

```bash
ykview .
ykview ./Documents
ykview remote-host
ykview remote-host:~/projects
ykview remote-host:/var/log
```

Windows SSHサーバーにはOpenSSH ServerとPowerShell 5.1以降、POSIX環境の接続先には`sh`が必要です。Windowsの絶対パスにはドライブ名かUNCパスを指定します。

Linux・macOSでは、ディレクトリ一覧の取得に使う補助実行ファイルを接続先の一時ディレクトリへ配置し、次回以降も再利用します。利用できない場合はシェルによる処理に切り替えます。

### オプション

| オプション | 既定値 | 説明 |
| --- | --- | --- |
| `-addr アドレス` | `127.0.0.1:8080` | HTTPの待受アドレス。空きポートを使う場合は`127.0.0.1:0`を指定 |
| `-open` | 有効 | 起動時に既定のブラウザーを開く。無効にする場合は`-open=false` |
| `-write` | 無効 | ファイルとディレクトリのアップロードを有効にする |
| `-max-upload-size バイト数` | `0` | 1回のアップロード要求の最大サイズ。`0`は無制限 |
| `-v` | 無効 | HTTP要求ごとのログを出力 |

指定したポートが使用中の場合は、後続の空きポートを使います。実際のURLは端末の`open:`行で確認できます。

```bash
ykview -addr 127.0.0.1:0 remote-host:/var/log
ykview -open=false -write remote-host:~/incoming
```

環境変数`DEBUG=1`を設定すると、詳細なデバッグログを出力します。

### プレビュー形式

| 形式 | 表示方法 |
| --- | --- |
| Markdown | GitHub Flavored MarkdownとMermaidの図を描画 |
| HTML、SVG | ブラウザーで直接表示 |
| PNG、JPEG、WebPなど | 画像として表示 |
| ソースコード、JSON、YAML、テキスト | 対応言語をシンタックスハイライト |
| CSV、TSV | 表として表示。表示上限があり、ソース表示に切り替え可能 |
| JSON Lines（`.jsonl`、`.ndjson`） | JSONのシンタックスハイライト付きで表示 |
| MP3、WAV、OGG、M4A、FLAC | 音声プレーヤーで再生 |
| MP4、WebM、MOV、OGV | 動画プレーヤーで再生 |

音声・動画の再生可否はブラウザーの対応コーデックに依存します。再生できない場合は「Raw」または「Download」を使ってください。

Markdown、Mermaid、シンタックスハイライトに必要なライブラリは実行ファイルに同梱しています。これらの読み込みに外部インターネット接続は不要です。

### 検索とファイル転送

名前検索と種類フィルターは、現在のディレクトリ内に適用します。全選択とZIPダウンロードは、絞り込み後に表示されている項目が対象です。

アップロードは`-write`を指定した場合だけ有効になります。表示中のディレクトリに、ファイル選択やドラッグ＆ドロップでファイル・ディレクトリを送れます。同名の通常ファイルは置き換えますが、シンボリックリンクやディレクトリは上書きしません。

ファイルは8 MiBずつに分けて全体を転送し、一時停止・再開に対応します。差分転送や、プロセス終了後の転送再開には対応していません。

### 利用上の条件

- 個人利用を想定しており、HTTPサーバーに認証機能はありません。既定ではループバックアドレスで待ち受けます。
- HTMLとSVGは同一オリジン上で表示し、スクリプトを実行できます。信頼できるファイルだけを開いてください。
- シンボリックリンクを辿り、指定したディレクトリの外にあるファイルを読み取れる場合があります。
- ファイルの編集・名前変更・削除や、変更時の自動再読み込みには対応していません。
- SSH URI形式とIPv6リテラルによる接続先指定には対応していません。

## 開発

Goのバージョンは[`mise.toml`](mise.toml)で1.27.2に固定しています。解析用ツールのdeadcode v0.52.0とstaticcheck v0.8.2は[`mise.analysis.toml`](mise.analysis.toml)で管理します。

```bash
mise exec -- make check
```

`make check`は`go test ./...`、`go test -race ./...`、`go vet ./...`、ビルドを実行します。接続先で使う補助実行ファイルは再生成しません。埋め込み用の実行ファイルを更新するときだけ、`mise exec -- make generate`を実行し、生成物の差分を確認してください。

`mise exec -- make install`を実行すると、`ykview`を`GOBIN`（未設定なら`GOPATH/bin`）へ配置できます。

接続先のOSを限定してビルドする場合は、`SERVER_OS`に`all`（既定）、`linux`、`darwin`、`windows`のいずれかを指定します。`all`以外では、指定したOS以外のSSHサーバーには接続できません。どの設定でもローカルのファイルを利用できます。

```bash
mise exec -- make build SERVER_OS=linux
mise exec -- make build SERVER_OS=windows
```

静的解析はmiseの`analysis`環境で実行します。

```bash
mise -E analysis install
mise -E analysis exec -- deadcode ./...
mise -E analysis exec -- staticcheck -tests=false ./...
```

Goのバージョンを更新したあとに解析ツールを再インストールする場合は、`mise -E analysis install --force`を実行します。

## CIとリリース

[CI](.github/workflows/ci.yml)では、pushとプルリクエストに対してテスト、データ競合の検査、`go vet`、ビルド、静的解析を実行します。Windowsクライアント・サーバーのテストとビルドも行います。

`v*`タグをpushすると、[リリースワークフロー](.github/workflows/release.yml)が6プラットフォーム向けのアーカイブと`SHA256SUMS`をGitHub Releaseへ公開します。

アーカイブをローカルで作る場合は次を実行します。

```bash
mise exec -- scripts/package-release.sh v0.1.0 dist
```

## コントリビューション

不具合の報告や機能の提案は[GitHub Issues](https://github.com/yukiyuigishi/ykview/issues)へお願いします。不具合の報告には、利用OS、実行したコマンド、再現手順を添えてください。コードを変更した場合は、`mise exec -- make check`で確認してからプルリクエストを作成してください。

## ライセンス

ykview本体は[MIT License](LICENSE)です。同梱ライブラリのライセンスは[THIRD_PARTY_NOTICES.md](THIRD_PARTY_NOTICES.md)を参照してください。
