import { Button, Grid, GridItem, Text } from "@chakra-ui/react";
import React from "react";
import styles from "./GameResult.module.css";
import { useRouter } from "next/navigation";
import Score from "@/types/Score";

interface GameResultProps {
  nextPage: () => void;
  score: Score;
}

const GameResult: React.FC<GameResultProps> = ({ nextPage, score }) => {
  const router = useRouter();

  const pushToRoot = () => {
    router.push("/");
  };

  return (
    <div className={styles.box}>
      <Grid h="100%" w="100%" templateRows="repeat(10, 1fr)" templateColumns="repeat(10, 1fr)" gap={6} bg="white">
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
        {/* TODO: スコアいらない */}
        <GridItem colSpan={3} colStart={3} rowStart={3} className={styles.centerText}>
          <Text fontSize="2xl" textAlign="center">
            スコア
          </Text>
        </GridItem>
        <GridItem colSpan={3} colStart={3} rowStart={4} className={styles.centerText}>
          <Text fontSize="2xl" textAlign="center">
            打鍵数
          </Text>
        </GridItem>
        <GridItem colSpan={3} colStart={3} rowStart={5} className={styles.centerText}>
          <Text fontSize="2xl" textAlign="center">
            ミス入力数
          </Text>
        </GridItem>
        <GridItem colSpan={3} colStart={3} rowStart={6} className={styles.centerText}>
          <Text fontSize="2xl" textAlign="center">
            入力時間
          </Text>
        </GridItem>
        <GridItem colSpan={3} colStart={3} rowStart={7} className={styles.centerText}>
          <Text fontSize="2xl" textAlign="center">
            WPM
          </Text>
        </GridItem>
        <GridItem colSpan={3} colStart={3} rowStart={8} className={styles.centerText}>
          <Text fontSize="2xl" textAlign="center">
            正打率
          </Text>
        </GridItem>
        <GridItem colSpan={3} colStart={6} rowStart={3} className={styles.centerText}>
          <Text fontSize="2xl" as="b">
            {Math.floor(score.score)}
          </Text>
        </GridItem>
        <GridItem colSpan={3} colStart={6} rowStart={4} className={styles.centerText}>
          <Text fontSize="2xl" as="b">
            {score.keystrokes} 回
          </Text>
        </GridItem>
        <GridItem colSpan={3} colStart={6} rowStart={5} className={styles.centerText}>
          <Text fontSize="2xl" as="b">
            {score.miss} 回
          </Text>
        </GridItem>
        <GridItem colSpan={3} colStart={6} rowStart={6} className={styles.centerText}>
          <Text fontSize="2xl" as="b">
            {String(score.time)} 秒
          </Text>
        </GridItem>
        <GridItem colSpan={3} colStart={6} rowStart={7} className={styles.centerText}>
          <Text fontSize="2xl" as="b">
            {Math.floor(score.wpm)} 字/分
          </Text>
        </GridItem>
        <GridItem colSpan={3} colStart={6} rowStart={8} className={styles.centerText}>
          <Text fontSize="2xl" as="b">
            {new Intl.NumberFormat("en-US", { style: "percent", maximumFractionDigits: 2 }).format(score.accuracy)}
          </Text>
        </GridItem>
        <GridItem colSpan={4} rowSpan={2} colStart={2} rowStart={9} className={styles.centerText}>
          <Button onClick={pushToRoot} colorScheme="red" size="lg" w="80%" h="90%">
            ゲームを終了する
          </Button>
        </GridItem>
        <GridItem colSpan={4} rowSpan={2} colStart={6} rowStart={9} className={styles.centerText}>
          <Button onClick={nextPage} colorScheme="blue" size="lg" w="80%" h="90%">
            もう一度プレイする
          </Button>
        </GridItem>
      </Grid>
    </div>
  );
};

export default GameResult;
