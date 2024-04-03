import { Box, Grid, GridItem, Text } from "@chakra-ui/react";
import React, { useEffect } from "react";
import { GamePreProps } from "../pages/Game";
import styles from "./GamePre.module.css";
import keyboardImage from "@/assets/images/LetsNote_Keyboard.png"

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
        <Grid
          templateAreas={`"header header"
                          "main nav"
                          "footer footer"`}
          gridTemplateRows={"60px 1fr 50px"}
          gridTemplateColumns={"1fr 300px"}
          h="100%"
          gap={6}
          bg="blackAlpha"
          border="4px solid white"
        >
          <GridItem pl="2" area={"header"} className={styles.centerText}>
            <Text color="white" fontSize="4xl" as="b">
              ゲーム説明
            </Text>
          </GridItem>
          <GridItem pl="4" color="white" area={"main"}>
            <Text>本文</Text>
          </GridItem>
          <GridItem pr="4" color="white" area={"nav"}>
            <Text>キーボードの写真とか？</Text>
          </GridItem>
          <GridItem pl="2" area={"footer"} className={styles.centerText}>
            <Text color="white" fontSize="3xl">
              [Space]キーを押して開始
            </Text>
          </GridItem>
        </Grid>
      </div>
    </Box>
  );
};

export default GamePre;
