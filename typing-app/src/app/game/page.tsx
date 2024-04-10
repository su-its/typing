import GamePage from "@/components/pages/Game";
import fs from "fs";

export default function Typing() {
  const filenames = fs.readdirSync("src/assets/texts/");
  const subjectText = fs.readFileSync(
    `src/assets/texts/${filenames[Math.floor(Math.random() * filenames.length)]}`,
    "utf-8"
  );
  const subjectTextOneLine = subjectText.replace(/\n/gm, " ");

  return <GamePage subjectText={subjectTextOneLine} />;
}
