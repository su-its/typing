import GamePage from "@/components/pages/Game";
import fs from "fs";

const filenames = fs.readdirSync("public/texts/");

export default function Typing() {
  const subjectText = fs.readFileSync(
    `public/texts/${filenames[Math.floor(Math.random() * filenames.length)]}`,
    "utf-8"
  );

  return <GamePage subjectText={subjectText} />;
}
