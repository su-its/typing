import RegisterScore, { ResultScore } from "@/types/RegisterScore";
import { Box } from "@chakra-ui/react";
import Image from "next/image";
import React, { useEffect, useRef, useState } from "react";
import ProgressBar from "../atoms/ProgressBar";
import { GameTypingProps } from "../pages/Game";
import styles from "./GameTyping.module.css";

const GameTyping: React.FC<GameTypingProps> = ({ nextPage, filenames, setResultScore }) => {
  // subjectTextの状態を管理するuseStateフック
  const [subjectText, setSubjectText] = useState("");
  const [startedAt, setStartedAt] = useState(new Date());

  useEffect(() => {
    const loadTextFile = async () => {
      // ランダムにファイル名を選択
      const randomFile = filenames[Math.floor(Math.random() * filenames.length)];
      // `public` ディレクトリからの相対パスを指定
      const filePath = `/texts/${randomFile}`;
      // fetch APIを使用してファイルの内容を読み込む
      try {
        const response = await fetch(filePath);
        const fetchedText = await response.text();
        setSubjectText(fetchedText); // レスポンスをsubjectTextステートに設定
        setStartedAt(new Date());
      } catch (error) {
        console.error("Error loading the text file:", error);
      }
    };

    loadTextFile();
  }, [filenames]); // ビルド時の警告防止のためにfilenamesを依存リストに追加

  const totalSeconds = 60;
  const [count, setCount] = useState(totalSeconds);
  const damyUserId = "damyId";

  const userId = damyUserId; // ToDo: 要変更
  const [correctType, setCorrectType] = useState(0); // 正打数
  const [incorrectType, setIncorrectType] = useState(0); // 誤打数

  const [typeIndex, setTypeIndex] = useState(0);
  // 残り時間のカウントダウン
  const updateFrequency = 100; // 100msごとにカウントダウン
  useEffect(() => {
    if (count <= 0) {
      sendResultData();
    } else {
      const timer = setInterval(() => {
        const pastTime = (new Date().valueOf() - startedAt.valueOf()) / 1000;
        const newCount = totalSeconds - pastTime;
        setCount(newCount);
      }, updateFrequency);
      return () => clearInterval(timer);
    }
  }, [count, nextPage, userId]); // ビルド時の警告防止のためにuserIdも依存リストに追加

  // 打ち終わった時にスコアを送信する
  useEffect(() => {
    if (typeIndex === subjectText.length - 1) {
      sendResultData();
    }
  }, [nextPage, userId, typeIndex]); // ビルド時の警告防止のためにuserIdも依存リストに追加

  // スコアデータを送信する
  const sendResultData = () => {
    // サーバに送信されるデータ
    const endedAt = new Date();
    const actualTypeTimeSeconds = (endedAt.valueOf() - startedAt.valueOf()) / 1000;
    const typeTimeSeconds = actualTypeTimeSeconds > totalSeconds ? totalSeconds : actualTypeTimeSeconds;
    const totalType = correctType + incorrectType;
    const accuracy = totalType === 0 ? 0 : (correctType / totalType) * 100; // [%]
    const registeredScore = {
      keystrokes: correctType,
      accuracy: accuracy,
      score: (correctType / typeTimeSeconds) * 60,
      startedAt: startedAt,
      endedAt: endedAt,
    } as RegisterScore;

    // リザルト画面用のデータ
    setResultScore({
      keystrokes: registeredScore.keystrokes,
      miss: incorrectType,
      time: new Date(typeTimeSeconds * 1000),
      wpm: (correctType / typeTimeSeconds) * 60,
      accuracy: registeredScore.accuracy,
      score: registeredScore.score,
    } as ResultScore);
    fetch(`http://localhost:8080/users/${userId}/scores`, {
      method: `POST`,
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(registeredScore),
    })
      .then((res) => res.json())
      .then(() => {})
      .catch((error) => {
        console.error(error);
      });
    nextPage();
  };

  // タイピング速度計算用
  const typingQueueListSize = 5; // ここで瞬間タイピング速度計算の粒度を決める 増やすほど変化が穏やかになる
  const [typingQueueList] = useState([] as number[]);
  const [currentTypeSpeed, setCurrentTypeSpeed] = useState(0);
  const [averageTypeSpeed, setAverageTypeSpeed] = useState(0);
  const addTypingQueueList = () => {
    const time = new Date().valueOf();
    typingQueueList.push(time);
    if (typingQueueList.length > typingQueueListSize) {
      typingQueueList.shift();
    }
  };

  const getTypingQueueListIndex = (index: number): number => {
    if (index < 0) {
      return 0;
    }
    if (index >= typingQueueList.length) {
      return typingQueueList.length - 1;
    }
    return 0;
  };

  const calcCurrentTypingSpeed = (): number => {
    if (typingQueueList.length <= 1) {
      return 0;
    }
    const typeTime = getTypingQueueListIndex(typingQueueList.length - 1) - getTypingQueueListIndex(0);
    const currentWpm = (typingQueueList.length / typeTime) * 60000;
    return toLogarithmWpm(currentWpm);
  };

  const calcAverageTypingSpeed = (): number => {
    const timeFromStart = new Date().valueOf() - startedAt.valueOf();
    const averageTypingSpeed = (correctType / timeFromStart) * 60000;
    return toLogarithmWpm(averageTypingSpeed);
  };

  const toLogarithmWpm = (wpm: number) => {
    const wpmForProgressBar = (1000 / 3) * Math.log10((999 / 1000) * wpm + 1);
    if (wpmForProgressBar > 1000) {
      return 1000;
    }
    return wpmForProgressBar;
  };

  const handleOnKeyDown = (e: React.KeyboardEvent) => {
    const key = e.key;
    if (key.length !== 1) {
      return; // アルファベット等以外のキーは無視 shiftなどがここに入る
    }
    const currentType = subjectText[typeIndex];
    if (key === currentType) {
      setTypeIndex(typeIndex + 1);
      setCorrectType(correctType + 1);
      addTypingQueueList();
      setCurrentTypeSpeed(calcCurrentTypingSpeed());
      setAverageTypeSpeed(calcAverageTypingSpeed());
    } else {
      setIncorrectType(incorrectType + 1);
    }
  };

  // ゲーム開始直後にフォーカスする
  const boxRef = useRef<HTMLDivElement>(null);
  useEffect(() => {
    if (boxRef.current) {
      boxRef.current.focus();
    }
  }, []);

  return (
    <Box onKeyDown={handleOnKeyDown} tabIndex={0} ref={boxRef}>
      <div className={styles.box}>
        <div className={`${styles.heading} ${styles.heading_name}`}>Article Name</div>
        <div className={`${styles.heading} ${styles.heading_time}`}>Time Remain</div>
        <div className={`${styles.heading} ${styles.heading_position}`}>Progress</div>
        <div className={`${styles.heading} ${styles.heading_speed}`}>Speed</div>
        <div className={`${styles.progress} ${styles.progress_time}`}>
          <ProgressBar maxWidth={280} height={20} maxValue={60} value={count} />
        </div>
        <div className={`${styles.progress} ${styles.progress_position}`}>
          <ProgressBar maxWidth={330} height={20} maxValue={subjectText.length - 1} value={typeIndex} />
        </div>
        <div className={`${styles.progress} ${styles.progress_speed}`}>
          <ProgressBar maxWidth={330} height={10} maxValue={1000} value={currentTypeSpeed} />
          <ProgressBar maxWidth={330} height={10} maxValue={1000} value={averageTypeSpeed} />
        </div>
        <Image
          className={styles.gauge_time}
          id="gauge_time"
          src="/img/gauge_time.png"
          alt={""}
          width={281}
          height={22}
        />
        <Image
          className={styles.gauge_position}
          id="gauge_position"
          src="/img/gauge_position.png"
          alt={""}
          width={330}
          height={24}
        />
        <Image
          className={styles.gauge_speed}
          id="gauge_speed"
          src="/img/gauge_speed.png"
          alt={""}
          width={330}
          height={24}
        />
        <div className={styles.title}>Lorem Ipsum</div>
        <div className={styles.text}>
          <div>
            <span className={styles.span_typed_text}>{subjectText.slice(0, typeIndex)}</span>
            <span className={styles.span_current_text}>{subjectText.slice(typeIndex, typeIndex + 1)}</span>
            <span>{subjectText.slice(typeIndex + 1, subjectText.length)}</span>
          </div>
        </div>
        <div className={styles.info_time}>
          残り <span className={styles.info_time_span}>{count.toFixed(1)}</span> 秒
        </div>
        <div className={styles.info_text}>
          {correctType} 語 / {subjectText.length} 字
        </div>
      </div>
    </Box>
  );
};

export default GameTyping;
