import { Box, Button, Text } from "@chakra-ui/react";
import React from "react";
import { SubGamePageProps } from "../pages/Game";

interface GamePreProps extends SubGamePageProps {
  setSubjectTextData: (data: string) => void;
}

const GamePre: React.FC<GamePreProps> = ({ nextPage, setSubjectTextData }) => {
  
  // 0からnまでの乱数を生成する関数
  const getRandomInt = (maxNumber: number) => {
    return Math.floor(Math.random() * (maxNumber + 1));
  };

  // 開始ボタンをクリックしたときの関数
  const handleStartButtonClick = () => {
    // const response = await fetch("URL"); // URLはtextを格納するjsonファイルの場所
    // const fetchedTextData: string[] = await response.json(); // 文章を格納するファイルはjson想定
    // const randomNumber = getRandomInt(fetchedTextData.length - 1);  // fetchedTextDataの長さに基づいて乱数を生成
    // const selectedTextData = textData[randomNumber]; // ランダムに選ばれたテキストデータ
    // setSubjectTextData(selectedTextData); // Propsに戻す

    // 次のページへ
    nextPage();
  };

  return (
    <Box>
      <Text>GamePre screen</Text>
      <Button onClick={handleStartButtonClick}>start</Button>
    </Box>
  );
};

export default GamePre;