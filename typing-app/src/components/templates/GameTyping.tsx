import { Box, Button, Text } from "@chakra-ui/react";
import React, { useEffect, useState } from "react";
import { GameTypingProps } from "../pages/Game";

const GameTyping: React.FC<GameTypingProps> = ({ nextPage, filenames }) => {
  // subjectTextの状態を管理するuseStateフック
  const [subjectText, setSubjectText] = useState("");

  useEffect(() => {
    const loadTextFile = async () => {
      // ランダムにファイル名を選択
      const randomFile = filenames[Math.floor(Math.random() * filenames.length)];
      // `public` ディレクトリからの相対パスを指定
      const filePath = `/texts/${randomFile}`;
      // fetch APIを使用してファイルの内容を読み込む
      try {
        const response = await fetch(filePath);
        const fetchedText = await response.text();
        setSubjectText(fetchedText); // レスポンスをsubjectTextステートに設定
      } catch (error) {
        console.error("Error loading the text file:", error);
      }
    };

    loadTextFile();
  }, []);

  return (
    <Box>
      <Text>Typing screen</Text>
      <Button onClick={nextPage}>finish</Button>
    </Box>
  );
};

export default GameTyping;
