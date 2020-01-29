# Danang PBL 2020

TCP Server And Client

## Getting Started

TCPサーバ, クライアント兼用のバイナリを用意しました．
各種OSに合わせて実行ファイルをダウンロードしてください

- Mac OS X: https://github.com/Bo0km4n/danang-2020/releases/download/pre-lease/danang_darwin_amd64
- Linux: https://github.com/Bo0km4n/danang-2020/releases/download/pre-lease/danang_linux_amd64
- Windows: https://github.com/Bo0km4n/danang-2020/releases/download/pre-lease/danang_windows_amd64

## How To Use

それぞれの実行ファイルを起動するとコマンドラインで次のようなヘルプが表示されると思います

```
Usage:
  danang [command]

Available Commands:
  client      cli
  help        Help about any command
  server      srv

Flags:
  -h, --help   help for danang

Use "danang [command] --help" for more information about a command.
```

この実行ファイルはclientモードとserverモードを使い分けることができます.

例えばコマンドラインで `danang_darwin_amd64 server` と実行すると以下のようなログが出ると思います

```
$ danang_darwin_amd64 server
INFO[0000] Listening on [::]:8000
```

これは現在8000番のポートでサーバアプリがリッスンしているということです

これに対して別のターミナルでクライアントを起動してメッセージを送信します

クライアントの起動コマンドは `danang_darwin_amd64 client --host 127.0.0.1 --port 8000`
というふうに実行するとIPアドレス `127.0.0.1` で ポート番号が `8000` のサーバにコネクションを確立します．
指定するIPアドレスやポート番号は自身の環境に合わせて適宜修正してください

クライアントが接続に成功すると以下のようなターミナルの画面になります

```
INFO[0000] Connected to 127.0.0.1:9000
127.0.0.1:9000 >>>
```

以降はこのプロンプトになにか文字列をタイプするとサーバ側に文字列がそのまま送信されるようになっています
サーバ側のターミナルでは送信された文字列が以下のように表示されるはずで

```
INFO[0054] Received content: hello
INFO[0059] Received content: world
```
プロンプトを終了する時は `quit` とタイプしてください
