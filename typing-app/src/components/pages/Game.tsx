"use client";
import { Text, VStack } from "@chakra-ui/react";
import React, { useState } from "react";
import GamePre from "../templates/GamePre";
import GameResult from "../templates/GameResult";
import GameTyping from "../templates/GameTyping";

interface GamePageProps {
  filenames: string[];
}

export interface SubGamePageProps {
  nextPage: () => void;
}

const GamePage: React.FC<GamePageProps> = ({ filenames }) => {
  enum ScreenIndex {
    IDX_PRE,
    IDX_TYPING,
    IDX_RESULT,
  }

  const [screenIndex, setScreenIndex] = useState<ScreenIndex>(ScreenIndex.IDX_PRE);
  const [subjectTextData, setSubjectTextData] = useState<string>("");  // 課題となるテキストデータ

  const subPageList = [
    <GamePre
      key={ScreenIndex.IDX_PRE}
      filenames={filenames}
      nextPage={() => setScreenIndex(ScreenIndex.IDX_TYPING)}
      setSubjectTextData={setSubjectTextData} // propsとして課題となるテキストデータを渡す
    />,
    <GameTyping key={ScreenIndex.IDX_TYPING} nextPage={() => setScreenIndex(ScreenIndex.IDX_RESULT)} />,
    <GameResult key={ScreenIndex.IDX_RESULT} nextPage={() => setScreenIndex(ScreenIndex.IDX_PRE)} />,
  ];
  return (
    <>
      <VStack>
        <Text fontSize="2xl">Hello, World!</Text>
        <Text fontSize="xl">Welcome to the Game Page</Text>
        {subPageList[screenIndex]}
      </VStack>
    </>
  );
};

export default GamePage;