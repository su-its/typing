import * as fs from "fs";
import * as path from "path";

// テキスト処理関数
function sanitizeText(text: string): string {
  return (
    text
      // スマートクォートの置換
      .replace(/[''′`]/g, "'")
      .replace(/[""]/g, '"')
      // 各種ダッシュの置換
      .replace(/[–—]/g, "-")
      // 省略記号の置換
      .replace(/…/g, "...")
      // その他の特殊文字
      .replace(/©/g, "(c)")
      .replace(/®/g, "(r)")
      .replace(/™/g, "(tm)")
      // 不可視制御文字の削除
      .replace(/[\u0000-\u001F\u007F-\u009F]/g, "")
  );
}

// 入力ディレクトリと出力ディレクトリの設定
const sourceDir: string = path.join(__dirname, "../original_texts");
const targetDir: string = path.join(__dirname, "../public/texts");

// ディレクトリの存在確認とファイルの処理を行う関数
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
    const files = fs.readdirSync(sourceDir);

    if (files.length === 0) {
      console.log("No files found in the source directory.");
      return;
    }

    files.forEach((file) => {
      const sourcePath = path.join(sourceDir, file);
      const targetPath = path.join(targetDir, file);

      // ディレクトリはスキップ
      if (fs.statSync(sourcePath).isDirectory()) {
        console.log(`Skipping directory: ${file}`);
        return;
      }

      // ファイルの読み込み、処理、書き込み
      const content = fs.readFileSync(sourcePath, "utf-8");
      const processedContent = sanitizeText(content);
      fs.writeFileSync(targetPath, processedContent);

      console.log(`Processed: ${file}`);
    });

    console.log("All files processed successfully.");
  } catch (error) {
    console.error("Error processing files:", error);
    process.exit(1);
  }
}

// スクリプトの実行
processFiles();
