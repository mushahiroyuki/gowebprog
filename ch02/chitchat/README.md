# ChitChat

このサンプルは第2章で紹介した、単純なフォーラムを実現するアプリケーションです。

このコードは第2章の説明と、1対1には対応していません。第2章の例よりも細かな処理を行っている部分があります。

以下に相違点をあげます。

* config.json で設定を変更することができます
* chitchat.log ファイルにログが書き出されます
* テストファイルがあります
* 構造体がすべて細かく定義されています（本文中では明確には記述されていませんでした）
* 一部の関数は（パッケージの一部ではなく）構造体のメソッドになっています

## 実行の手順

1. ＜Postgresを起動＞
2. createdb chitchat
3. psql -f data/setup.sql -d chitchat
4. go mod tidy
5. go build
6. ./chitchat

●go.modファイルがない場合は、4.の前に次を実行してください。

 $ go mod init github.com/mushahiroyuki/gowebprog/ch02/chitchat

●go.modの先頭行が次のようになっていない場合は、一度 go.modを削除して上のコマンドを実行してください。

module github.com/mushahiroyuki/gowebprog/ch02/chitchat
