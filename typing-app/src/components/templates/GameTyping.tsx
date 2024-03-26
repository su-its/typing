import RegisterScore from "@/types/RegisterScore";
import { Box } from "@chakra-ui/react";
import axios from "axios";
import Image from "next/image";
import React, { useEffect, useState } from "react";
import ProgressBar from "../atoms/ProgressBar";
import { GameTypingProps } from "../pages/Game";
import styles from "./GameTyping.module.css";

const GameTyping: React.FC<GameTypingProps> = ({ nextPage, filenames }) => {
  // subjectTextの状態を管理するuseStateフック
  const [subjectText, setSubjectText] = useState("");

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
      } catch (error) {
        console.error("Error loading the text file:", error);
      }
    };

    loadTextFile();
  }, [filenames]); // ビルド時の警告防止のためにfilenamesを依存リストに追加

  const totalSeconds = 250;
  const [count, setCount] = useState(totalSeconds);
  const damyScoreData = {
    Keystrokes: 123,
    Accuracy: 456.7,
    Score: 890.1,
    StartedAt: new Date(),
    EndedAt: new Date(),
  } as RegisterScore;
  const damyUserId = "damyId";

  const userId = damyUserId; // ToDo: 要変更
  const scoreData = damyScoreData; // ToDo: 要変更

  useEffect(() => {
    if (count <= 0) {
      axios
        .post(`http://localhost:8080/users/${userId}/scores`, scoreData)
        .then((res) => {
          console.log(res.data);
          nextPage();
        })
        .catch((error) => {
          console.error(error);
        });
    } else {
      const timer = setInterval(() => setCount(count - 1), 1000);
      return () => clearInterval(timer);
    }
  }, [count, nextPage, userId, scoreData]); // ビルド時の警告防止のためにuserId, scoreDataも依存リストに追加

  const [startTime, setStartTime] = useState(new Date().valueOf()); // 平均速度計算用
  const typingQueueListSize = 5; // ここで瞬間タイピング速度計算の粒度を決める 増やすほど変化が穏やかになる

  // 瞬間タイピング速度計算用
  const [typingQueueList] = useState([] as number[]);
  const [currentTypeSpeed, setCurrentTypeSpeed] = useState(0);
  const addTypingQueueList = () => {
    const time = new Date().valueOf();
    typingQueueList.push(time);
    if (typingQueueList.length > typingQueueListSize) {
      typingQueueList.shift();
    }
  };

  const calcCurrentTypingSpeed = (): number => {
    if (typingQueueList.length <= 1) {
      return 0;
    }
    const typeTime = typingQueueList[typingQueueList.length - 1] - typingQueueList[0];
    return (typingQueueList.length / typeTime) * 60000;
  };

  const [typeIndex, setTypeIndex] = useState(0);
  const [correctType, setCorrectType] = useState(0);
  const [incorrectType, setIncorrectType] = useState(0); // 使わないかもしれない
  const [typeProgress, setTypeProgress] = useState(0);
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
    } else {
      setIncorrectType(incorrectType + 1);
    }
    setTypeProgress(typeIndex);
  };

  return (
    <Box onKeyDown={handleOnKeyDown} tabIndex={0}>
      <div className={styles.box}>
        <div className={`${styles.heading} ${styles.heading_name}`}>Article Name</div>
        <div className={`${styles.heading} ${styles.heading_time}`}>Time Remain</div>
        <div className={`${styles.heading} ${styles.heading_position}`}>Progress</div>
        <div className={`${styles.heading} ${styles.heading_speed}`}>Speed</div>
        <div className={`${styles.progress} ${styles.progress_time}`}>
          {
            // ToDo 時間の計算
          }
          <ProgressBar maxWidth={330} height={20} maxValue={250} value={count} />
        </div>
        <div className={`${styles.progress} ${styles.progress_position}`}>
          <ProgressBar maxWidth={330} height={20} maxValue={subjectText.length - 1} value={typeProgress} />
        </div>
        <div className={`${styles.progress} ${styles.progress_speed}`}>
          {
            // ToDo 速度の計算
          }
          <ProgressBar maxWidth={330} height={10} maxValue={500} value={currentTypeSpeed} />
          <ProgressBar maxWidth={330} height={10} maxValue={500} value={300} />
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
          残り <span className={styles.info_time_span}>{count}</span> 秒
        </div>
        <div className={styles.info_text}>
          {correctType} 語 / {subjectText.length} 字
        </div>
      </div>
    </Box>
  );
};

export default GameTyping;
