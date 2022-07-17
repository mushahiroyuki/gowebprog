// リスト3.1
/* 次のコマンドで実行できます。
   （Windowsの場合、sudo は不要です）
  $ sudo go run server.go  

 * ほかのウェブサーバが動いているときは、すぐに終了してしまいます。
 * 起動しておいてブラウザに http://localhost/ を入力すると "404 page not found" が表示されます。
   それ以上の機能はありません。
	 終了するには control+C を押してください。

なお、次のコマンドを実行すると01simplest という実行ファイルができます。
  $ go build
その後で次のコマンドを実行するとサーバが起動します。
 sudo ./01simplest

*/

package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe("", nil)
}
