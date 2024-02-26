# プロジェクトタイトル

42_Road-to-Mercari-Gopher-Dojo-00

## 概要

このプロジェクトは、株式会社メルカリとの共同開発課題です。  
Goについて取り扱います。このプロジェクト群を解き終わると、限定の座談会に招待されます。


## プロジェクトの目的

このプロジェクトは、42Tokyoの課題の一環として実施されました。その主な目的は、  
Go言語の基本的な操作を習得し、go.modファイルの役割と管理方法を理解することにあります。  
具体的なタスクとしては、画像フォーマットの変換コマンド（例：JPEGからPNGへの変換）の開発が挙げられます。  

## チャレンジした点

- エンコードとデコードの理解：
画像データのエンコードとデコードの概念を深く理解し、これらをGo言語で実装することに挑戦しました。  
具体的には、`image/jpeg`および`image/png`パッケージを使用して画像データを処理するコードの作成が含まれます。

- Go言語ドキュメントの活用：
`go doc`コマンドの使い方を学び、Go言語の公式ドキュメントを効率的に参照する方法を習得しました。

## 学んだこと

- パッケージによる画像処理：
Go言語の`image/jpeg`および`image/png`パッケージを用いた画像のエンコードとデコード方法について学びました。
これにより、異なる画像フォーマット間での変換が可能となります。

- コードの可読性と安全性の向上：
`defer`キーワードを用いてファイル操作の終了処理を記述する方法について学びました。
`defer newFile.Close()`の使用により、リソースのリークを防ぎながらコードの可読性を高めることができました。

## 使用方法

> - **ex00 normal**
> ```bash php
> #!/bin/bash
> cd ex00
> go build
> ./convert 42tokyo_logo.jpg(変換したいjpgファイル)
> ```


> - **ex00 bonus**
> ```bash php
> #!/bin/bash
> ./convert
> ./convert -i=png -o=jpg images -q=10
> ```

`bonus`にはフラグがあり、`-i`,`-o`で入出力のファイルを選べます。  
`-q`は`jpg`ファイルの画質をコントロールします。


<br>

**出力例**
```bash
?> go mod tidy
?> go build
?> ./convert
error: invalid argument
?> ./convert nosuchdirectory
error: nosuchdirectory: no such file or directory
?> ls -1 images
42tokyo_logo.jpg
profile_photo.jpg
?> ./convert images
?> ls -1 images
42tokyo_logo.jpg
42tokyo_logo.png
profile_photo.jpg
profile_photo.png
?> echo 'aaa' > images/test.txt
?> ls -1 images
42tokyo_logo.jpg
42tokyo_logo.png
profile_photo.jpg
profile_photo.png
test.txt
?> ./convert images
error: images/test.txt is not a valid file
```

## 技術スタック

Go言語
