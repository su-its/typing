# GamePage 向け説明などなど

Team1 用．適宜加筆していってください．

## 用語

- Page：[url]/game で表示されるページ．
- SubPage：Page 内の画面遷移で使う．それぞれ 1 つの画面．
- GamePre：ゲーム開始前の画面のこと．
- GameTyping：実際にタイピングする画面のこと．
- GameResult：ゲーム終了後の結果表示画面のこと．

## 前提

Game は以下のように画面遷移する．
GamePre → GameTyping → GameResult

## ディレクトリ

### Page

src/app/compontents/pages/

- Game.tsx

### SubPage

src/app/compontents/templates/

- GamePre.tsx
- GameTyping.tsx
- GameResult.tsx

## 実装

### 画面遷移

![画面遷移図](./img/game-page/screen-transition.png)

暫定的に `GamePre` → `GameTyping` → `GameResult` をループ．`GameResult` の後は変更のはず．
Game.tsx 内で`subPageList`配列に 3 つの `SubPage` コンポーネントを格納しておき，インデックスを変更することで画面を遷移させる．
`SubPage` 内では`nextPage`関数を受け取り，それを実行することで次のページのインデックスに変更されるようにしておく．
`GameResult` の `nextPage` は要変更．
