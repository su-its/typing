import GamePage from "@/components/pages/Game";
import fs from "fs";
import path from "path";

export default function Typing() {
  const textsDirectory = path.join(process.cwd(), "components/assets/texts");
  const filenames = fs.readdirSync(textsDirectory).filter((filename) => filename.endsWith(".txt"));
  return <GamePage filenames={filenames} />;
}
