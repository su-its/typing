import GamePage from "@/components/pages/Game";
import fs from "fs";
import path from "path";

export default function Typing() {
  const randomNumber = Math.floor(Math.random() * 6) + 1;
  // テキストファイルのパスを指定
  const filePath = path.join(process.cwd(), "src/assets/texts", `text${randomNumber}.txt`);
  // ファイルの内容を同期的に読み込んで fileContent に格納
  const fileContent = fs.readFileSync(filePath, "utf8");
  return <GamePage fileContent={fileContent} />;
}
