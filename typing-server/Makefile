.PHONY: generate seed connect up-with-build help db-drop-dev db-reset-dev db-clean-dev
generate:
	cd internal/infra/ent && go generate

seed:
	@echo "シードデータを投入しています..."
	docker compose -f docker-compose.dev.yml run --rm seed
	@echo "シードデータの投入が完了しました"

# ヘルプコマンド
help:
	@echo "使用可能なコマンド:"
	@echo "  generate       - entのコード生成を実行する"
	@echo "  seed           - 開発環境にシードデータを投入する"
	@echo "  connect        - 開発環境に接続する"
	@echo "  up-with-build  - 開発環境を起動する"
	@echo "  db-drop-dev    - 開発環境のデータベースのテーブルをドロップする"
	@echo "  db-reset-dev   - 開発環境のデータベースをリセット（ドロップ後に再作成）する"
	@echo "  db-clean-dev   - 開発環境のデータベースボリュームを完全に削除してクリーンにする"

# 開発環境：コンテナのコンソールに接続
connect:
	docker compose -f docker-compose.dev.yml exec server bash

# 開発環境：開発環境を起動する
up-with-build:
	@echo "開発環境を起動しています..."
	docker compose -f docker-compose.dev.yml up --build

# 開発環境：データベースのテーブルをドロップ
db-drop-dev:
	@echo "開発環境のテーブルをドロップしています..."
	docker-compose -f docker-compose.dev.yml exec db mysql -u user -ppassword typing-db -e "SET FOREIGN_KEY_CHECKS = 0; \
	DROP TABLE IF EXISTS scores; \
	DROP TABLE IF EXISTS users; \
	SET FOREIGN_KEY_CHECKS = 1;"
	@echo "テーブルのドロップが完了しました"

# 開発環境：データベースを再作成
db-reset-dev:
	@echo "開発環境のデータベースをリセットしています..."
	docker-compose -f docker-compose.dev.yml exec db mysql -u root -ppassword -e "DROP DATABASE IF EXISTS \`typing-db\`; \
	CREATE DATABASE \`typing-db\`; \
	GRANT ALL PRIVILEGES ON \`typing-db\`.* TO 'user'@'%';"
	@echo "データベースのリセットが完了しました"

# 開発環境：データベースボリュームを完全にクリーン
db-clean-dev:
	@echo "開発環境のデータベースボリュームを削除しています..."
	docker-compose -f docker-compose.dev.yml down
	docker volume rm typing-server_db-data
	docker-compose -f docker-compose.dev.yml up -d
	@echo "データベースボリュームを削除し、コンテナを再起動しました"
