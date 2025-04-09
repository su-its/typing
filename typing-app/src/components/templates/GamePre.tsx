import React, { useEffect } from "react";
import { GamePreProps } from "../pages/Game";
import styles from "./GamePre.module.scss";
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
    <div className={styles["game-pre"]}>
      <div className={styles.header}>操作説明</div>
      <div className={styles.main}>
        <div className={styles.left}>
          <p>
            <b>【基本ルール】</b>
          </p>
          <p>
            制限時間は1分間!
            <br />
            英文を速く・正確に入力して高スコアを目指そう!
          </p>
          <p>
            <b>【ランキング掲載条件】</b>
          </p>
          <p>
            WPM (Words Per Minutes) : 120字以上
            <br />
            正打率 : 95%以上
          </p>
          <p>
            <b>【入力しても反応しなくなった場合】</b>
          </p>
          <p>テキストが表示されているボックスをクリックすると元に戻れます。</p>
        </div>
        <div className={styles.right}>
          {/* eslint-disable-next-line @next/next/no-img-element */}
          <img src={keyboardImage.src} alt="Keyboard" />
        </div>
      </div>
      <div className={styles.footer}>[Space] キーを押して開始</div>
    </div>
  );
};

export default GamePre;
