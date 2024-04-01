import { Box, Text } from "@chakra-ui/react";
import React, { useEffect } from "react";
import { GamePreProps } from "../pages/Game";
import styles from "./GamePre.module.css";

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
      <div className={styles.box}>
        <Text fontSize="4xl">
          ゲーム説明
        </Text>
      </div>
    </Box>
  );
};

export default GamePre;
