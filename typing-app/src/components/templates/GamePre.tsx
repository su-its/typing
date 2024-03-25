import { Box, Text } from "@chakra-ui/react";
import React, { useEffect } from "react";
import { SubGamePageProps } from "../pages/Game";

const GamePre: React.FC<SubGamePageProps> = ({ nextPage }) => {
  // Spaceキーを押したときに実行する関数
  const handleSpaceButtonDown: React.KeyboardEventHandler = async (e) => {
    if (e.code === 'Space') {
      e.preventDefault();  // ページのスクロールなどのデフォルト動作を防止
      // 次のページへ
      nextPage();
    }
  };

  return (
    <Box onKeyDown={handleSpaceButtonDown}>
      <Text>GamePre screen</Text>
    </Box>
  );
};

export default GamePre;
