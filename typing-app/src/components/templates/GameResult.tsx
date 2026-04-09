import type React from "react";
import styles from "./GameResult.module.scss";
import { useRouter } from "next/navigation";
import type Score from "@/types/Score";

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
      <div className={styles.header}>
        タイピング結果<span>RESULT</span>
      </div>
      <div className={styles.table}>
        <div className={styles.row}>
          <div className={styles.left}>打鍵数</div>
          <div className={styles.right}>
            {score.keystrokes}
            <span>回</span>
          </div>
        </div>
        <div className={styles.row}>
          <div className={styles.left}>ミス入力数</div>
          <div className={styles.right}>
            {score.miss}
            <span>回</span>
          </div>
        </div>
        <div className={styles.row}>
          <div className={styles.left}>入力時間</div>
          <div className={styles.right}>
            {String(score.time)}
            <span>秒</span>
          </div>
        </div>
        <div className={styles.row}>
          <div className={styles.left}>WPM</div>
          <div className={styles.right}>
            {Math.floor(score.wpm)}
            <span>字/分</span>
          </div>
        </div>
        <div className={styles.row}>
          <div className={styles.left}>正打率</div>
          <div className={styles.right}>
            {new Intl.NumberFormat("en-US", {
              maximumFractionDigits: 2,
            }).format(score.accuracy * 100)}
            <span>%</span>
          </div>
        </div>
      </div>
      <div className={styles.footer}>
        <div className={styles.container}>
          <button type="button" className={`${styles.button} ${styles.red}`} onClick={pushToRoot}>
            ゲームを終了する
          </button>
          <button type="button" className={`${styles.button} ${styles.blue}`} onClick={nextPage}>
            もう一度プレイする
          </button>
        </div>
      </div>
    </div>
  );
};

export default GameResult;
