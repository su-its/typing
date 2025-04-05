import type Score from "@/types/Score";
import { client } from "@/libs/api";
import React, { useCallback, useEffect, useRef, useState } from "react";
import ProgressBar from "../atoms/ProgressBar";
import type { GameTypingProps } from "../pages/Game";
import styles from "./GameTyping.module.scss";
import { getCurrentUser } from "@/app/actions";
import gaugePositionImg from "../../../public/img/gauge_position.png";
import gaugeSpeedImg from "../../../public/img/gauge_speed.png";
import gaugeTimeImg from "../../../public/img/gauge_time.png";
import type { User } from "@/types/user";
import { showErrorToast } from "@/utils/toast";
import { useRouter } from "next/navigation";

// 定数の分離
const TOTAL_SECONDS = 60; // 後でconfigから取得する想定
const TYPING_QUEUE_SIZE = 5; // 瞬間タイピング速度計算の粒度
const MAX_TYPING_SPEED = 300; // 表示上限
const UPDATE_FREQUENCY = 100; // タイマー更新頻度（ミリ秒）

const GameTyping: React.FC<GameTypingProps> = ({ nextPage, subjectText, setScore }) => {
  const router = useRouter();

  // ステート定義
  const [timeRemaining, setTimeRemaining] = useState(TOTAL_SECONDS);
  const [stats, setStats] = useState({
    correctKeystrokes: 0,
    incorrectKeystrokes: 0,
    typeIndex: 0,
    averageTypeSpeed: 0,
  });

  // 開始時刻と処理フラグの参照
  const startTimeRef = useRef<number>(Date.now());
  const isProcessingRef = useRef(false);

  const typingQueueRef = useRef<number[]>([]);
  const typeIndexRef = useRef(stats.typeIndex);
  const boxRef = useRef<HTMLDivElement>(null);

  // タイピング速度を計算
  const calculateTypingSpeed = useCallback(() => {
    const elapsedTime = (Date.now() - startTimeRef.current) / 1000;
    if (elapsedTime <= 0) return 0;

    return Math.min((stats.correctKeystrokes / elapsedTime) * 60, MAX_TYPING_SPEED);
  }, [stats.correctKeystrokes]);

  // スコアデータを送信
  const sendResultData = useCallback(async () => {
    // 既に処理中なら実行しない
    if (isProcessingRef.current) {
      return;
    }

    // 処理開始フラグをセット
    isProcessingRef.current = true;

    // 実際に経過した時間（秒）を計算
    const actualElapsedTime = (Date.now() - startTimeRef.current) / 1000;

    const totalKeystrokes = stats.correctKeystrokes + stats.incorrectKeystrokes;
    const accuracy = totalKeystrokes === 0 ? 0 : stats.correctKeystrokes / totalKeystrokes;
    const wpm = (stats.correctKeystrokes / actualElapsedTime) * 60;

    const score: Score = {
      keystrokes: stats.correctKeystrokes,
      accuracy,
      score: wpm,
      miss: stats.incorrectKeystrokes,
      time: actualElapsedTime, // 実際の経過時間を使用
      wpm,
    };

    setScore(score);

    try {
      // ユーザー情報の取得
      const user: User | undefined = await getCurrentUser();

      if (!user) {
        showErrorToast("ユーザー情報が取得できませんでした");
        router.push("/");
        return;
      }

      // スコア送信
      const { error } = await client.POST("/scores", {
        body: {
          user_id: user.id,
          keystrokes: score.keystrokes,
          accuracy: score.accuracy,
        },
      });

      if (error) {
        showErrorToast("スコアの登録に失敗しました");
      }

      nextPage();
    } catch (error) {
      console.error("Score submission error:", error);
      showErrorToast("エラーが発生しました");
    }
  }, [stats, nextPage, router, setScore]);

  // 実時間ベースを使用する理由:
  // setIntervalは厳密に一定間隔で実行されるわけではなく、ブラウザの状態やタブのアクティブ状態によって遅延が発生する等といった理由
  // 実時間ベースのタイマー処理
  useEffect(() => {
    const timerInterval = setInterval(() => {
      const elapsed = (Date.now() - startTimeRef.current) / 1000;
      const remaining = Math.max(TOTAL_SECONDS - elapsed, 0);

      setTimeRemaining(remaining);

      // 時間切れチェック
      if (remaining <= 0 && !isProcessingRef.current) {
        clearInterval(timerInterval);
        sendResultData();
      }
    }, UPDATE_FREQUENCY);

    return () => clearInterval(timerInterval);
  }, [sendResultData]);

  // タイピングが終わった時の処理
  useEffect(() => {
    // タイピングが終了し、まだ処理されていないかチェック
    if (stats.typeIndex === subjectText.length - 1 && !isProcessingRef.current) {
      sendResultData();
    }
  }, [stats.typeIndex, subjectText.length, sendResultData]);

  // typeIndexRefを更新
  useEffect(() => {
    typeIndexRef.current = stats.typeIndex;
  }, [stats.typeIndex]);

  // キー入力処理
  useEffect(() => {
    const handleKeyDown = (e: KeyboardEvent) => {
      // 既に終了処理中なら入力を無視
      if (isProcessingRef.current) return;

      // 時間切れなら入力を無視
      if (timeRemaining <= 0) return;

      const key = e.key;

      // 文字キー以外は無視
      if (key.length !== 1) return;

      const currentIndex = typeIndexRef.current;
      const currentChar = subjectText[currentIndex];

      if (key === currentChar) {
        // 正解のキー入力
        const newTypeIndex = currentIndex + 1;

        setStats((prev) => ({
          ...prev,
          typeIndex: newTypeIndex,
          correctKeystrokes: prev.correctKeystrokes + 1,
        }));

        // タイピング速度計算
        const now = Date.now();
        const queue = typingQueueRef.current;
        queue.push(now);

        if (queue.length > TYPING_QUEUE_SIZE) {
          queue.shift();
        }

        setStats((prev) => ({
          ...prev,
          averageTypeSpeed: calculateTypingSpeed(),
        }));
      } else {
        // 誤ったキー入力
        setStats((prev) => ({
          ...prev,
          incorrectKeystrokes: prev.incorrectKeystrokes + 1,
        }));
      }
    };

    window.addEventListener("keydown", handleKeyDown);
    return () => window.removeEventListener("keydown", handleKeyDown);
  }, [subjectText, calculateTypingSpeed, timeRemaining]);

  // 初期フォーカス設定
  useEffect(() => {
    if (boxRef.current) {
      boxRef.current.focus();
    }

    // コンポーネントマウント時に開始時刻を設定
    startTimeRef.current = Date.now();
  }, []);

  return (
    <div tabIndex={0} ref={boxRef}>
      <div className={styles.box}>
        {/* TODO: Article Nameって消すんじゃなかったっけ */}
        <div className={`${styles.heading} ${styles.heading_name}`}>Article Name</div>
        <div className={`${styles.heading} ${styles.heading_time}`}>Time Remain</div>
        <div className={`${styles.heading} ${styles.heading_position}`}>Progress</div>
        <div className={`${styles.heading} ${styles.heading_speed}`}>Speed</div>
        <div className={`${styles.progress} ${styles.progress_time}`}>
          <ProgressBar maxWidth={280} height={20} maxValue={60} value={timeRemaining} />
        </div>
        <div className={`${styles.progress} ${styles.progress_position}`}>
          <ProgressBar maxWidth={330} height={20} maxValue={subjectText.length - 1} value={stats.typeIndex} />
        </div>
        <div className={`${styles.progress} ${styles.progress_speed}`}>
          <ProgressBar maxWidth={330} height={20} maxValue={1000} value={stats.averageTypeSpeed} />
        </div>
        <img className={styles.gauge_time} id="gauge_time" src={gaugeTimeImg.src} width={281} height={24} />
        <img className={styles.gauge_position} id="gauge_position" src={gaugePositionImg.src} width={330} height={24} />
        <img className={styles.gauge_speed} id="gauge_speed" src={gaugeSpeedImg.src} width={330} height={24} />
        <div className={styles.title}>-</div>
        <div className={styles.text}>
          <div>
            <span className={styles.span_typed_text}>{subjectText.slice(0, stats.typeIndex)}</span>
            <span className={styles.span_current_text}>{subjectText.slice(stats.typeIndex, stats.typeIndex + 1)}</span>
            <span>{subjectText.slice(stats.typeIndex + 1)}</span>
          </div>
        </div>
        <div className={styles.info_time}>
          残り <span className={styles.info_time_span}>{timeRemaining.toFixed(1)}</span> 秒
        </div>
        <div className={styles.info_text}>
          {stats.correctKeystrokes} 字 / {subjectText.length} 字
        </div>
      </div>
    </div>
  );
};

export default GameTyping;
