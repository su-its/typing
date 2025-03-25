import * as fs from "fs";
import * as path from "path";

// キーボードで入力可能な文字のホワイトリスト
const allowedChars: string =
  "1234567890-^¥!\"#$%&'()0=~|qwertyuiop@[QWERTYUIOP`{asdfghjkl;:]ASDFGHJKL+*}zxcvbnm,./_ZXCVBNM<>?_ \n\r\t";

// テキスト処理関数
function sanitizeText(text: string): string {
  // 特殊なアポストロフィやクォートを標準的なものに置換
  let processedText = text
    .replace(/['']/g, "'") // スマートシングルクォート→通常のシングルクォート
    .replace(/[""]/g, '"') // スマートダブルクォート→通常のダブルクォート
    .replace(/[–—]/g, "-") // 各種ダッシュ→ハイフン
    .replace(/…/g, "..."); // 省略記号→ピリオド3つ

  // アポストロフィを含む単語の特殊処理
  // この処理を先に行い、アポストロフィを保持する
  const contractionPattern = /(\w)'(\w)/g;
  const contractions: { original: string; position: number }[] = [];

  let match;
  while ((match = contractionPattern.exec(processedText)) !== null) {
    contractions.push({
      original: match[0],
      position: match.index,
    });
  }

  // ホワイトリストに含まれない文字を削除
  let filteredText = processedText
    .split("")
    .filter((char) => allowedChars.includes(char))
    .join("");

  for (const contraction of contractions) {
    const [first, second] = contraction.original.split("'");
    const pattern = new RegExp(`${first}\\s*${second}`);
    filteredText = filteredText.replace(pattern, contraction.original);
  }

  // 連続したスペースを単一のスペースに置換
  return filteredText.replace(/\s+/g, " ").trim();
}

// 単語の区切りとなる文字かどうかを判定
function isWordSeparator(char: string): boolean {
  return [" ", "\t", "\n", "\r", ".", ",", ";", ":", "!", "?", "(", ")", "[", "]", "{", "}"].includes(char);
}

// 入出力ディレクトリの設定
const sourceDir: string = path.join(__dirname, "../original_texts");
const targetDir: string = path.join(__dirname, "../public/texts");

// メイン処理関数
function processFiles(): void {
  try {
    // 出力ディレクトリが存在しない場合は作成
    if (!fs.existsSync(targetDir)) {
      fs.mkdirSync(targetDir, { recursive: true });
    }

    // 入力ディレクトリの存在確認
    if (!fs.existsSync(sourceDir)) {
      throw new Error(`Source directory does not exist: ${sourceDir}`);
    }

    // ソースディレクトリ内の全ファイルを処理
    const files: string[] = fs.readdirSync(sourceDir);

    if (files.length === 0) {
      console.log("No files found in the source directory.");
      return;
    }

    files.forEach((file: string) => {
      const sourcePath: string = path.join(sourceDir, file);
      const targetPath: string = path.join(targetDir, file);

      // ディレクトリはスキップ
      if (fs.statSync(sourcePath).isDirectory()) {
        console.log(`Skipping directory: ${file}`);
        return;
      }

      // ファイルの読み込み、処理、書き込み
      const content: string = fs.readFileSync(sourcePath, "utf-8");
      const processedContent: string = sanitizeText(content);
      fs.writeFileSync(targetPath, processedContent);

      console.log(`Processed: ${file}`);
    });

    console.log("All files processed successfully.");
  } catch (error) {
    console.error("Error processing files:", error instanceof Error ? error.message : String(error));
    process.exit(1);
  }
}

// スクリプトの実行
processFiles();
