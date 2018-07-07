package main

import (
	"database/sql"
	"fmt"

	/* 実行前に以下のコマンドでgo-sql-driverをインストールしておく
	go get -u github.com/go-sql-driver/mysql
	*/
	_ "github.com/go-sql-driver/mysql"
)

type User struct {
	Name string `json:"name"`
}

func main() {
	// fmt.Println("Go言語でMySQLを使おう！")

	// 第一引数: RDBSの種類, 第二引数: (ユーザー名:パスワード:ローカルホスト名:ポート番号)/DB名
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/testdb")

	// nilの場合にpanicでエラーを発生させる
	if err != nil {
		panic(err.Error())
	}
	// deferは遅延実行を意味する。クエリを実行後にDBから切断
	defer db.Close()

	fmt.Println("MySQLのデータベースへの接続に成功しました")

	/* テーブルの作成
	create, err := db.Query("CREATE TABLE testdb.users(id int, user_name varchar(255), twitter_id varchar(255));")

	// nilの場合にpanicでエラーを発生させる
	if err != nil {
		panic(err.Error())
	}

	// deferは遅延実行を意味する。クエリを実行後にDBから切断
	defer create.Close()
	*/

	/* レコードの作成
	insert, err := db.Query("INSERT INTO users (id, name, twitter_id) VALUES('1', 'ウェブ系ウシジマ', 'yuta_ushizima')")

	// nilの場合にpanicでエラーを発生させる
	if err != nil {
		panic(err.Error())
	}

	defer insert.Close()

	fmt.Println("userテーブルへの登録に成功しました")
	*/

	results, err := db.Query("SELECT name FROM users")

	// nilの場合にpanicでエラーを発生させる
	if err != nil {
		panic(err.Error())
	}

	for results.Next() {
		var user User

		err = results.Scan(&user.Name)
		// nilの場合にpanicでエラーを発生させる
		if err != nil {
			panic(err.Error())
		}

		fmt.Println(user.Name)
	}
}
