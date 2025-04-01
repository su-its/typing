import * as fs from "fs";
import * as path from "path";

// キーボードで入力可能な文字のホワイトリスト
const allowedChars: string =
  "1234567890-^¥!\"#$%&'()0=~|qwertyuiop@[QWERTYUIOP`{asdfghjkl;:]ASDFGHJKL+*}zxcvbnm,./_ZXCVBNM<>?_ \n\r";

// テキスト処理関数
function sanitizeText(text: string): string {
  // 特殊なアポストロフィやクォートを標準的なものに置換
  let processedText = text
    .replace(/[\u2018\u2019]/g, "'") // 左右のスマートシングルクォート→通常のシングルクォート
    .replace(/[\u201C\u201D]/g, '"') // 左右のスマートダブルクォート→通常のダブルクォート
    .replace(/[–—]/g, "-") // 各種ダッシュ→ハイフン
    .replace(/…/g, "..."); // 省略記号→ピリオド3つ

  // ホワイトリストに含まれない文字を削除
  let filteredText = processedText
    .split("")
    .filter((char) => allowedChars.includes(char))
    .join("");

  // 連続したスペースを単一のスペースに置換
  return filteredText.replace(/\s+/g, " ").trim();
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
