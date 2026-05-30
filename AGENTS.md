# AGENTS

このリポジトリは、ソースファイルとデータから生成される GitHub Pages ポートフォリオサイトです。

## Repository purpose
- `uduki-ikuya` の静的な個人ポートフォリオサイト
- サイトのソースは `src/`、`data/`、`cmd/sitegen/` にあります
- 生成された出力は `site/` に配置されます

## Agent guidance
- 変更はソースファイルのみで行ってください: `src/`、`data/`、`cmd/sitegen/`、およびドキュメントファイル
- 明示的に依頼がない限り、生成された `site/` ファイルを編集しないでください
- 既存のビルドフローを使用してください:
  - `make build` でサイトを生成
  - `make clean` で `site/` を削除
- 変更は最小限に抑え、プロジェクトに集中してください

## Relevant technologies
- Go ... 静的サイトジェネレーター
- HTML/CSS ... フロントエンド
- YAML ... 静的データ管理
