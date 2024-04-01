import GamePage from "@/components/pages/Game";
import fs from "fs";

export default function Typing() {
  const filenames = fs.readdirSync("src/assets/texts/");
  const randomNumber = Math.floor(Math.random() * filenames.length) + 1; //
  const subjectText = fs.readFileSync(`src/assets/texts/text${randomNumber}.txt`, "utf-8");

  return <GamePage subjectText={subjectText} />;
}
