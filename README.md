```script
# goインストール
brew install go
# バージョン確認
go version

# sqlite3インストール
brew install sqllite3
# バージョン確認
sqlite3 -version

# データベース作成
sqlite3 sample.db
# テーブル作成
create table sample_user(email, password, fullname);
# サンプルデータ挿入
insert into sample_user values ('abc@example.com', '1234', 'Robert Griesemer');
.exit
# サーバの起動
go run main.go
```
