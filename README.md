# git-cz-go

`git-cz-go` は、Conventional Commits に基づいたコミットメッセージを簡単に生成できる Go 製の CLI ツールです。

## 特徴
- インタラクティブなプロンプトでコミットタイプ、スコープ、メッセージを簡単に指定可能。
- Git コマンドと連携し、直感的な操作性。

## インストール
以下のコマンドでインストールできます：

```bash
go install github.com/a1yama/git-cz-go@latest
```

> `go install` により、`$GOPATH/bin` にバイナリがインストールされます。
> 必要に応じて `$GOPATH/bin` を `PATH` に追加してください。

例：

```bash
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.bashrc
source ~/.bashrc
```

## 使い方
任意の Git プロジェクト内で以下を実行します：

```bash
git-cz-go
```

### 手順
1. コミットタイプを選択します（例: `feat`、`fix` など）。
2. 変更のスコープを入力します（オプション）。
3. 簡単な説明を入力します。
4. 詳細な説明（オプション）を入力します。
5. コミットメッセージを確認し、コミットを実行します。

## 開発
このプロジェクトをローカルでビルドするには：

```bash
git clone https://github.com/a1yama/git-cz-go.git
cd git-cz-go
go build -o git-cz-go ./cmd
```

ローカルで実行する場合：

```bash
./git-cz-go
```

or 

```bash
sudo mv git-cz-go /usr/local/bin/
git-cz-go
```

## 貢献
バグ報告や新機能の提案は [GitHub Issues](https://github.com/a1yama/git-cz-go/issues) で受け付けています。

## ライセンス
このプロジェクトは [MIT ライセンス](LICENSE) の下で提供されています。
