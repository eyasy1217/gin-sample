```script
# goインストール
brew install go
# バージョン確認
go version

# sqlite3インストール
brew install sqllite3
# バージョン確認
sqlite3 -version

# データベース作成とデータベースに入る
sqlite3 sample.db
# テーブル作成
create table sample_user(email, password, fullname);
# サンプルデータ挿入
insert into sample_user values ('rob@example.com', '1234', 'Robert Griesemer');
insert into sample_user values ('guido@example.com', '6789', 'Guido van Rossum');
# データ確認
select * from sample_user;
.exit

# サーバの起動
go run main.go
```
