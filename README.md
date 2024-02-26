# プロジェクトタイトル

42_Road-to-Mercari-Gopher-Dojo-00

## 概要

このプロジェクトは、株式会社メルカリとの共同開発課題です。  
Goについて取り扱います。このプロジェクト群を解き終わると、限定の座談会に招待されます。


## プロジェクトの目的

42Tokyoの課題であり、また、  
Go言語の基本的な操作、go.modの意味等を学びます。

## チャレンジした点

エンコード、デコードの概念、またそれをコードに移す作業が難しかったです。
go docの使い方も学びました。


## 学んだこと

image/jpeg、image/pngパッケージによるエンコード、デコード。
defer newFile.Close()によるコードの可読性の向上。

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
./convert -i=png -o=jpg images -q=10
> ```

bonusにはフラグがあり、-i,-oで入出力のファイルを選べます。  
-qはjpgファイルの画質をコントロールします。


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
