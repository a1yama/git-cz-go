# git-cz-go 動作確認ガイド

## ビルド済みバイナリ

ビルド済みのバイナリが `/Users/a1yama/ghq/git-cz-go/git-cz-go` に作成されています。

## 動作確認方法

### 方法1: テストスクリプトを使用

```bash
# テスト環境をセットアップ
./test-demo.sh

# 表示されたディレクトリに移動
cd /tmp/git-cz-go-demo-[timestamp]

# git-cz-goを実行
/Users/a1yama/ghq/git-cz-go/git-cz-go
```

### 方法2: 現在のリポジトリで確認

```bash
# ダミーファイルを作成
echo "test" > test.txt
git add test.txt

# git-cz-goを実行
./git-cz-go
```

## 動作確認の流れ

1. **コミットタイプ選択**
   - 矢印キーで移動、またはキーで選択
   - 数字キー（1-9）でクイック選択も可能

2. **スコープ入力**（オプション）
   - 例: `ui`, `api`, `docs`
   - Enterキーでスキップ可能

3. **件名入力**（必須）
   - 例: `add new feature for user authentication`

4. **本文入力**（オプション）
   - 複数行入力可能
   - 空の状態でEnterを押すか、Enter二回で続行
   - Ctrl+Dでも続行可能

5. **破壊的変更の確認**
   - YesまたはNoを選択
   - Y/Nキーでクイック選択可能

6. **フッター入力**（オプション）
   - 例: `Closes #123`
   - Enterキーでスキップ可能

7. **確認画面**
   - 作成されるコミットメッセージのプレビュー
   - Yesを選択してコミット

## キーボードショートカット

- `↑/↓` または `j/k`: 項目間を移動
- `1-9`: コミットタイプのクイック選択
- `Y/N`: Yes/Noのクイック選択
- `Enter`: 選択/確定
- `Esc`: 前のステップに戻る
- `Ctrl+C`: 終了
- `Ctrl+D`: テキストエディタで入力完了
- `Enter (空入力時)`: 本文をスキップ
- `Enter x2`: 本文の入力完了

## サンプル入力

### シンプルな例
```
Type: feat (1キーを押す)
Scope: (Enterでスキップ)
Subject: add user authentication
Body: (Enterでスキップ)
Breaking: No (Nキーを押す)
Footer: (Enterでスキップ)
Confirm: Yes
```

結果: `✨ feat: add user authentication`

### 完全な例
```
Type: feat (1キーを押す)
Scope: auth
Subject: implement OAuth2 login
Body: 
Added OAuth2 authentication with Google and GitHub providers.
Users can now login using their social accounts.
(Enter二回で完了)
Breaking: No (Nキーを押す)
Footer Type: Closes
Footer Value: #42
Confirm: Yes
```

結果:
```
✨ feat(auth): implement OAuth2 login

Added OAuth2 authentication with Google and GitHub providers.
Users can now login using their social accounts.

Closes: #42
```

## トラブルシューティング

### "not a git repository"エラー
gitリポジトリ内で実行してください：
```bash
git init
```

### "nothing to commit"エラー
ファイルをステージングしてください：
```bash
git add .
```

### 実行権限エラー
```bash
chmod +x git-cz-go
```