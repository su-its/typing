import { Box, Grid, GridItem, Text, Image, Center, Flex } from "@chakra-ui/react";
import React, { useEffect } from "react";
import { GamePreProps } from "../pages/Game";
import styles from "./GamePre.module.css";
import keyboardImage from "@/assets/images/LetsNote_Keyboard.png";

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
          gridTemplateColumns={"1fr 400px"}
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
          <GridItem pl="4" color="white" area={"main"} fontSize="2xl">
            <Box height="100%" alignItems="center">
              <Text>制限時間は1分間!</Text>
              <Text>英文を速く・正確に入力して高スコアを目指そう！</Text>
              <Text>【ランキング掲載条件】</Text>
              <Text>WPM(打鍵数): 120字以上 かつ 正打率: 95%以上</Text>
            </Box>
          </GridItem>
          <GridItem pr="4" color="white" area={"nav"} height="100%" placeItems="center">
            <Center mt="20">
              <Image src={keyboardImage.src} alt="Logo" maxH={300} />
            </Center>
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
