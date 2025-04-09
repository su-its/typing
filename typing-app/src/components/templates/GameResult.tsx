import React from "react";
import styles from "./GameResult.module.scss";
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
      <div className={styles.header}>Result</div>
      <div className={styles.table}>
        <div className={styles.row}>
          <div className={styles.left}>打鍵数</div>
          <div className={styles.right}>{score.keystrokes} 回</div>
        </div>
        <div className={styles.row}>
          <div className={styles.left}>ミス入力数</div>
          <div className={styles.right}>{score.miss} 回</div>
        </div>
        <div className={styles.row}>
          <div className={styles.left}>入力時間</div>
          <div className={styles.right}>{String(score.time)} 秒</div>
        </div>
        <div className={styles.row}>
          <div className={styles.left}>WPM</div>
          <div className={styles.right}>{Math.floor(score.wpm)} 字/分</div>
        </div>
        <div className={styles.row}>
          <div className={styles.left}>正打率</div>
          <div className={styles.right}>
            {new Intl.NumberFormat("en-US", { style: "percent", maximumFractionDigits: 2 }).format(score.accuracy)}
          </div>
        </div>
      </div>
      <div className={styles.footer}>
        <div className={styles.container}>
          <div className={`${styles.button} ${styles.red}`} onClick={pushToRoot}>
            ゲームを終了する
          </div>
          <div className={`${styles.button} ${styles.blue}`} onClick={nextPage}>
            もう一度プレイする
          </div>
        </div>
      </div>
    </div>
  );
};

export default GameResult;
