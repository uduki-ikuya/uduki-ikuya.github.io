# uduki-ikuya.github.io

uduki-ikuya (卯月 幾哉) のポートフォリオサイトです。
GitHub Pages を利用して公開されています。

## 🌐 サイトURL
- **URL**: [https://uduki-ikuya.github.io/](https://uduki-ikuya.github.io/)

## 🛠️ 技術スタック
*(今後、サイトの実装内容に合わせて随時更新します)*
- **Frontend**: HTML / CSS / JavaScript
- **Hosting**: GitHub Pages

## 📂 リポジトリ構成
- `README.md` - このファイル (日本語版説明)
- `README_en.md` - 英語版説明 ([English Version](./README_en.md))

## 👤 プロフィール
- **名前**: uduki-ikuya (卯月 幾哉)
- **GitHub**: [@uduki-ikuya](https://github.com/uduki-ikuya)
- **Email**: uduki.ikuya@gmail.com

## ローカルでのビルド / デプロイ手順

前提: `Go` がインストールされていること（推奨: 1.21 以上）。

- ローカルでサイトを生成する（`Makefile` を利用）:

```bash
make build   # クリーンしてサイトを生成します
make clean   # 生成された site/ を削除します
```

- 生成物は `site/` ディレクトリに出力されます。開発中は `site/` をコミットしないでください（`.gitignore` に登録済みです）。

- CI (GitHub Actions):
	- `.github/workflows/deploy.yml` は `make build` を使ってサイトを生成し、生成された `site/` を GitHub Pages にデプロイします。

簡単な確認手順:

```bash
make build
open site/index.html   # macOS の場合、生成されたページをローカルで開く
```

---

© 2026 uduki-ikuya
