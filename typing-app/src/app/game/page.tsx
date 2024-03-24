import GamePage from "@/components/pages/Game";
import fs from 'fs';
import path from 'path';
import type { GetStaticProps } from "next";

export const getStaticProps: GetStaticProps = async() => {
  const textsDirectory = path.join(process.cwd(), 'public/texts');
  const filenames = fs.readdirSync(textsDirectory).filter((filename) => filename.endsWith('.txt'));

  return {
    props: {
      filenames,
    },
  };
}

interface TypingProps {
  filenames: string[];
}

export default function Typing({ filenames }: TypingProps) {
  return <GamePage filenames={filenames} />;
}
