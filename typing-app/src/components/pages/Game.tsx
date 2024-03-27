"use client";
import { VStack } from "@chakra-ui/react";
import React, { useState } from "react";
import GamePre from "../templates/GamePre";
import GameResult from "../templates/GameResult";
import GameTyping from "../templates/GameTyping";
import { ResultScore } from "@/types/RegisterScore";

export interface SubGamePageProps {
  nextPage: () => void;
}

export interface GameTypingProps {
  nextPage: () => void;
  filenames: string[];
  setResultScore: (data: ResultScore) => void;
}

interface GamePageProps {
  filenames: string[];
}

const GamePage: React.FC<GamePageProps> = ({ filenames }) => {
  const ScreenIndex = {
    IDX_PRE: 0,
    IDX_TYPING: 1,
    IDX_RESULT: 2,
  } as const;

  type ScreenIndex = (typeof ScreenIndex)[keyof typeof ScreenIndex];

  const [resultScore, setResultScore] = useState<ResultScore>({
    Keystrokes: 0,
    Miss: 0,
    Time: new Date(),
    WPM: 0,
    Accuracy: 0,
  });

  const [screenIndex, setScreenIndex] = useState<ScreenIndex>(ScreenIndex.IDX_PRE);
  const subPageList = [
    <GamePre key={ScreenIndex.IDX_PRE} nextPage={() => setScreenIndex(ScreenIndex.IDX_TYPING)} />,
    <GameTyping
      key={ScreenIndex.IDX_TYPING}
      nextPage={() => setScreenIndex(ScreenIndex.IDX_RESULT)}
      filenames={filenames}
      setResultScore={setResultScore}
    />,
    <GameResult
      key={ScreenIndex.IDX_RESULT}
      nextPage={() => setScreenIndex(ScreenIndex.IDX_PRE)}
      resultScore={resultScore}
    />,
  ];
  return (
    <>
      <VStack>{subPageList[screenIndex]}</VStack>
    </>
  );
};

export default GamePage;
