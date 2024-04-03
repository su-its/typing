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

### 各役割について

それぞれの分担について，編集する可能性のあるファイルとやることを簡単に書いた．どちらも軽く考えただけ．編集する可能性のあるファイルは足りないかもしれないし余分かもしれないし，やることは他に良い or 理解しやすいやり方があるならそのようにすれば OK．以下のリストはあくまでも目安．

#### 1 ゲーム開始ボタン

##### 編集する可能性のあるファイル

- GamePre.tsx

##### やること

- ボタンのスタイリングなど

#### 2 開始ボタンを押したら、文章をランダムに選択

##### 編集する可能性のあるファイル

- GamePre.tsx
- GameTyping.tsx

##### やること

- まだはっきりしてない？のでとりあえず 0 ～ n の乱数を生成してダミーの fetch リクエストを送る．
- GamePre.tsx ではすでに開始ボタンをクリックした時のイベントを設定しているが，props として受け取っている`nextPage`関数を呼び出す関数を作り，その中で文章選択のプログラムを書く．

#### 3 カウントダウンの開始

##### 編集する可能性のあるファイル

- GamePre.tsx
- GameTyping.tsx
- ProgressBar.tsx

##### やること

- `useEffect`を使いそう．タイマーアプリの実装とかが参考になるはず．
- カウントダウン開始は[2]に実装してもらうので，最初はボタンを押したら開始とかで開発．
- ProgressBar に残り時間と最大時間を渡す．

#### 4 入力のマッチング

##### 編集する可能性のあるファイル

- GameTyping.tsx

##### やること

- キー入力はライブラリを使えば楽…？
- 入力したキーの正誤のカウント．

#### 5 カウントダウンが 0 になったタイミングで、データを送信

##### 編集する可能性のあるファイル

- GameTyping.tsx

##### やること

- POST リクエストでスコアデータを送信．データ形式等はまた後で．

#### 6 入力テキストの管理

##### 編集する可能性のあるファイル

##### やること

#### 7 リザルトの表示

##### 編集する可能性のあるファイル

- GameResult.tsx

##### やること

- 結果を良い感じに表示．
