import GamePage from "@/components/pages/Game";
import fs from "fs";

const filenames = fs.readdirSync("public/texts/");

const getRandomSubjectText = () => {
  const randomFilename = filenames[Math.floor(Math.random() * filenames.length)] ?? filenames[0];
  return fs.readFileSync(`public/texts/${randomFilename}`, "utf-8");
};

export default function Typing() {
  const subjectText = getRandomSubjectText();
  const subjectTextOneLine = subjectText.replace(/\n/gm, " ");

  return <GamePage subjectText={subjectTextOneLine} />;
}
