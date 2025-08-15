# git-cz-go プロジェクト仕様書

## プロジェクト概要
git-cz-goは、[cz-git](https://github.com/Zhengqbbb/cz-git)にインスパイアされ、Goで実装されたConventional Commits CLI ツールです。Bubble Teaフレームワークを使用して美しくインタラクティブなTUIを提供し、標準化されたコミットメッセージの作成を容易にします。

## 機能仕様

### コア機能
1. **インタラクティブなコミットメッセージ作成**
   - ステップバイステップのウィザード形式
   - リアルタイムバリデーション
   - プレビュー機能

2. **Conventional Commits準拠**
   - 標準的なコミットタイプ（feat, fix, docs等）
   - スコープのサポート（未実装）
   - Breaking Changeの表示（未実装）
   - フッター情報（未実装）

3. **絵文字サポート**
   - 各コミットタイプに対応する絵文字
   - 設定で有効/無効の切り替え可能

### 現在実装済みの機能
- ✅ コミットタイプの選択（数字キーによるクイック選択対応）
- ✅ サブジェクト（概要）の入力
- ✅ コミットメッセージのプレビューと確認
- ✅ 絵文字サポート
- ✅ 設定ファイルによるカスタマイズ
- ✅ キーボードナビゲーション

### 未実装の機能（今後の拡張予定）
- ⏳ スコープの入力と補完
- ⏳ 本文（body）の入力
- ⏳ Breaking Changeのマーキング
- ⏳ フッター情報（Issue番号等）の入力
- ⏳ プロジェクト構造からのスコープ自動提案
- ⏳ コミット履歴からの学習機能
- ⏳ 複数の設定プロファイル
- ⏳ AIによるコミットメッセージ生成

## 開発方針

### コーディング規約
- Go 1.20以上を前提とする
- 標準的なGoのコーディング規約に従う
- エラーハンドリングは適切に行い、ユーザーフレンドリーなメッセージを表示する
- Bubble Teaフレームワークのベストプラクティスに従う

### ディレクトリ構造
```
git-cz-go/
├── cmd/git-cz-go/          # メインアプリケーション
│   └── main.go
├── internal/               # 内部パッケージ
│   ├── config/            # 設定管理
│   ├── git/               # Git操作
│   ├── model/             # データモデル
│   └── ui/                # UIコンポーネント
│       ├── components/    # 個別UIコンポーネント
│       └── styles/        # スタイル定義
└── pkg/                   # 公開パッケージ
    └── commitmsg/         # コミットメッセージフォーマッター
```

### 使用ライブラリ
- Bubble Tea (github.com/charmbracelet/bubbletea) - TUIフレームワーク
- Bubbles (github.com/charmbracelet/bubbles) - TUIコンポーネント
- Lipgloss (github.com/charmbracelet/lipgloss) - スタイリング
- go-homedir (github.com/mitchellh/go-homedir) - ホームディレクトリ取得

### テスト
- 単体テストは`_test.go`ファイルに記述
- `go test ./...`でテスト実行
- モックが必要な場合は、インターフェースを使用して依存性注入を行う
- UIコンポーネントは個別にテスト可能な設計にする

### ビルドとリリース
- `go build -o git-cz-go ./cmd/git-cz-go`でビルド
- GitHub Actionsによる自動リリース（タグプッシュ時）
- GoReleaserを使用したクロスプラットフォームビルド

### コミットメッセージ
Conventional Commitsフォーマットに従う：
- `feat`: 新機能
- `fix`: バグ修正
- `docs`: ドキュメントのみの変更
- `style`: コードの意味に影響しない変更
- `refactor`: バグ修正も機能追加も行わないコード変更
- `perf`: パフォーマンス改善
- `test`: テストの追加や修正
- `build`: ビルドシステムや外部依存関係の変更
- `ci`: CI設定ファイルやスクリプトの変更
- `chore`: その他の変更
- `revert`: 以前のコミットの取り消し

### 設定ファイル仕様
```json
{
  "types": [
    {
      "type": "feat",
      "description": "A new feature",
      "emoji": "✨"
    }
  ],
  "useEmoji": true,
  "maxSubjectLength": 100
}
```

設定ファイルの検索順序：
1. `./.git-cz.json`（カレントディレクトリ）
2. `~/.git-cz.json`（ホームディレクトリ）
3. `~/.config/git-cz/config.json`（XDG設定ディレクトリ）

### UI/UX設計原則
1. **効率性**
   - キーボードナビゲーションを優先
   - 数字キー（1-9）によるクイック選択
   - Vimライクなキーバインディング（将来的に）

2. **視認性**
   - カラーコーディングで情報の種類を区別
   - プログレスインジケーターで現在のステップを表示
   - プレビューで最終的なコミットメッセージを確認

3. **ユーザビリティ**
   - Escキーで前のステップに戻る
   - Ctrl+Cで安全に終了
   - エラーメッセージは具体的で対処法を提示

### 開発時の注意事項
- ユーザーのgitコミット作業を妨げないよう、高速で応答性の良いUIを維持
- エラー時は適切にリカバリーし、ユーザーの入力を失わないようにする
- プラットフォーム依存のコードは最小限に抑える
- cz-gitの優れたUXを参考にしつつ、Goらしいシンプルさを保つ

## 参考資料
- [Conventional Commits仕様](https://www.conventionalcommits.org/)
- [cz-git](https://github.com/Zhengqbbb/cz-git) - オリジナルのインスピレーション源
- [Bubble Tea Documentation](https://github.com/charmbracelet/bubbletea)