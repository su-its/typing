import { Box, Grid, GridItem, Text, Image, Center, Flex } from "@chakra-ui/react";
import React, { useEffect } from "react";
import { GamePreProps } from "../pages/Game";
import styles from "./GamePre.module.css";
import keyboardImage from "@/assets/images/keyboard.png";

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
          gridTemplateRows={"100px 1fr 50px"}
          gridTemplateColumns={"1fr 1fr"}
          h="100%"
          gap={0}
          bg="#263238"
          border="4px solid white"
          alignItems="center"
          overflow="auto"
        >
          <GridItem pl="2" mt="100px" area={"header"} className={styles.centerText}>
            <Text color="white" fontSize="4xl" as="b">
              ゲーム説明
            </Text>
          </GridItem>
          <GridItem pl="10" color="white" area={"main"} fontSize="2xl">
            <Box height="100%">
              <Text>制限時間は1分間!</Text>
              <Text mb="4">英文を速く・正確に入力して高スコアを目指そう!</Text>
              <Text as="b">【ランキング掲載条件】</Text>
              <Text>WPM(Words per Minutes): 120字以上</Text>
              <Text>正打率: 95%以上</Text>
            </Box>
          </GridItem>
          <GridItem
            pr="0"
            color="white"
            area={"nav"}
            height="100%"
            display="flex"
            flexDirection="column"
            placeItems="center"
            alignItems="center"
            justifyContent="center"
          >
            <Image src={keyboardImage.src} alt="Logo" maxH={300} />
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
