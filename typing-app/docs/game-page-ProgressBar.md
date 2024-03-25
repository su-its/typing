# ProgressBar コンポーネント説明

Team1 用．適宜加筆していってください．

## 場所

typing/typing-app/src/components/atoms/ProgressBar.tsx

## 使い方

```[jsx]
<ProgressBar maxWidth={250} value={100} />
```

### props

- maxWidth: number：100%時の横幅[px]．
- value: number：進捗．0-maxValue をとる．
- maxValue: number：100%の進捗．
- height: number：高さ[px]

### 例

progress を変更することで ProgressBar が伸び縮みする．

```[jsx]
const [progress, setProgress] = useState(0);
const handleProgressChange = () => setProgress(progress+1);

return (
  <ProgressBar maxWidth={250} height={20} maxValue={100} value={progress} />
);
```

### Style

css ファイルは不使用．
div タグ内で CSS-in-JS のみ使用．
背景色・高さはここで指定．
変更する場合は，他で使用されている場所に影響がでるので注意．
