import { Box, Text } from "@chakra-ui/react";
import React, { useEffect } from "react";
import { GamePreProps } from "../pages/Game";

const GamePre: React.FC<GamePreProps> = ({ nextPage }) => {
  useEffect(() => {
    // Spaceキーを押したときに実行する関数
    const handleSpaceButtonDown = (e: KeyboardEvent) => {
      if (e.code === "Space") {
        e.preventDefault(); // ページのスクロールなどのデフォルト動作を防止
        // 次のページへ
        nextPage();
      }
    };

    window.addEventListener("keydown", handleSpaceButtonDown);

    return () => {
      window.removeEventListener("keydown", handleSpaceButtonDown);
    };
  }, [nextPage]);

  return (
    <Box>
      <Text>GamePre screen</Text>
    </Box>
  );
};

export default GamePre;
