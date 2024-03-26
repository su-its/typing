import { Box, Button, Grid, GridItem, Text } from "@chakra-ui/react";
import React from "react";
import { SubGamePageProps } from "../pages/Game";
import styles from "./GameResult.module.css";
import { ResultScore } from "@/types/RegisterScore";

interface GameResultProps {
  nextPage: () => void;
  resultScore: ResultScore;
}

const GameResult: React.FC<GameResultProps> = ({ nextPage, resultScore }) => {
  return (
    <div className={styles.box}>
      <Grid h="100%" w="100%" templateRows="repeat(9, 1fr)" templateColumns="repeat(10, 1fr)" gap={6} bg="white">
        <GridItem
          colSpan={10}
          rowSpan={2}
          colStart={1}
          rowStart={1}
          bg="blue.900"
          rounded="md"
          className={styles.centerText}
        >
          <Text fontSize="4xl" as="b" color="white" textAlign="center">
            Result
          </Text>
        </GridItem>
        <GridItem colSpan={3} colStart={3} rowStart={3} className={styles.centerText}>
          <Text fontSize="2xl" textAlign="center">
            打鍵率
          </Text>
        </GridItem>
        <GridItem colSpan={3} colStart={3} rowStart={4} className={styles.centerText}>
          <Text fontSize="2xl" textAlign="center">
            ミス入力数
          </Text>
        </GridItem>
        <GridItem colSpan={3} colStart={3} rowStart={5} className={styles.centerText}>
          <Text fontSize="2xl" textAlign="center">
            入力時間
          </Text>
        </GridItem>
        <GridItem colSpan={3} colStart={3} rowStart={6} className={styles.centerText}>
          <Text fontSize="2xl" textAlign="center">
            WPM
          </Text>
        </GridItem>
        <GridItem colSpan={3} colStart={3} rowStart={7} className={styles.centerText}>
          <Text fontSize="2xl" textAlign="center">
            正解率
          </Text>
        </GridItem>
        <GridItem colSpan={3} colStart={6} rowStart={3} className={styles.centerText}>
          <Text fontSize="2xl" as="b">
            {resultScore.Keystrokes}回
          </Text>
        </GridItem>
        <GridItem colSpan={3} colStart={6} rowStart={4} className={styles.centerText}>
          <Text fontSize="2xl" as="b">
            {resultScore.Miss}回
          </Text>
        </GridItem>
        <GridItem colSpan={3} colStart={6} rowStart={5} className={styles.centerText}>
          <Text fontSize="2xl" as="b">
            {resultScore.Time.getMinutes()}分{resultScore.Time.getSeconds()}秒
            {Math.floor(resultScore.Time.getMilliseconds() / 100)}
          </Text>
        </GridItem>
        <GridItem colSpan={3} colStart={6} rowStart={6} className={styles.centerText}>
          <Text fontSize="2xl" as="b">
            {Math.floor(resultScore.WPM)}字/分
          </Text>
        </GridItem>
        <GridItem colSpan={3} colStart={6} rowStart={7} className={styles.centerText}>
          <Text fontSize="2xl" as="b">
            {resultScore.Accuracy.toFixed(1)}%
          </Text>
        </GridItem>
        <GridItem colSpan={4} rowSpan={2} colStart={4} rowStart={8} className={styles.centerText}>
          <Button onClick={nextPage} colorScheme="blue" size="lg" w="100%" h="90%">
            次へ
          </Button>
        </GridItem>
      </Grid>
    </div>
  );
};

export default GameResult;
