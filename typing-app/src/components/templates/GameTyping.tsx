import RegisterScore from "@/types/RegisterScore";
import { Box } from "@chakra-ui/react";
import Image from "next/image";
import { client } from "@/libs/api";
import React, { useCallback, useEffect, useRef, useState } from "react";
import ProgressBar from "../atoms/ProgressBar";
import { GameTypingProps } from "../pages/Game";
import styles from "./GameTyping.module.css";
import { getCurrentUser } from "@/app/actions";
import gaugePositionImg from "../../../public/img/gauge_position.png";
import gaugeSpeedImg from "../../../public/img/gauge_speed.png";
import gaugeTimeImg from "../../../public/img/gauge_time.png";
import { User } from "@/types/user";
import { showErrorToast } from "@/utils/toast";
import { useRouter } from "next/navigation";

const GameTyping: React.FC<GameTypingProps> = ({ nextPage, subjectText, setResultScore }) => {
  const router = useRouter();

  const [startedAt, setStartedAt] = useState(new Date());

  const totalSeconds = 60; // TODO: Configファイルから取得
  const [count, setCount] = useState(totalSeconds);

  const [correctType, setCorrectType] = useState(0); // 正打数
  const [incorrectType, setIncorrectType] = useState(0); // 誤打数

  // スコアデータを送信する
  const sendResultData = useCallback(async () => {
    // サーバに送信されるデータ
    const endedAt = new Date();
    const actualTypeTimeSeconds = (endedAt.valueOf() - startedAt.valueOf()) / 1000; //TODO: マジックナンバー確認
    const typeTimeSeconds = actualTypeTimeSeconds > totalSeconds ? totalSeconds : actualTypeTimeSeconds;
    const totalType = correctType + incorrectType;
    const accuracy = totalType === 0 ? 0 : (correctType / totalType) * 100; // [%]
    const registeredScore: RegisterScore = {
      keystrokes: correctType,
      accuracy: accuracy,
      score: (correctType / typeTimeSeconds) * 60, // TODO: マジックナンバー確認
      startedAt: startedAt,
      endedAt: endedAt,
    };

    const user: User | undefined = await getCurrentUser();
    //TODO:Userが取得できなかった場合のエラーハンドリングを追加
    if (!user) {
      showErrorToast("ユーザー情報が取得できませんでした");
      router.push("/");
      return;
    }

    const { error } = await client.POST("/scores", {
      body: { user_id: user.id, keystrokes: registeredScore.keystrokes, accuracy: registeredScore.accuracy },
    });
    if (error) {
      showErrorToast("スコアの登録に失敗しました");
      return;
    } else {
      // リザルト画面用のデータ
      setResultScore({
        keystrokes: registeredScore.keystrokes,
        miss: incorrectType,
        time: new Date(typeTimeSeconds * 1000), // TODO: マジックナンバー確認
        wpm: (correctType / typeTimeSeconds) * 60, // TODO: マジックナンバー確認
        accuracy: registeredScore.accuracy,
        score: registeredScore.score,
      });
    }
    nextPage();
  }, [startedAt, totalSeconds, correctType, incorrectType, setResultScore, nextPage]);

  const [typeIndex, setTypeIndex] = useState(0);
  // 残り時間のカウントダウン
  const updateFrequency = 100; // TODO: 1000msに変更、これもConfigファイル可rあ取得
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
  }, [count, nextPage, sendResultData, startedAt]);

  // 打ち終わった時にスコアを送信する
  useEffect(() => {
    if (typeIndex === subjectText.length - 1) {
      sendResultData();
    }
  }, [nextPage, sendResultData, subjectText.length, typeIndex]);

  // タイピング速度計算用
  const typingQueueListSize = 5; // ここで瞬間タイピング速度計算の粒度を決める 増やすほど変化が穏やかになる
  const [typingQueueList] = useState<number[]>([]);
  const [averageTypeSpeed, setAverageTypeSpeed] = useState(0);

  const typeIndexRef = useRef(typeIndex);
  useEffect(() => {
    // setTypeIndexの結果を反映する
    typeIndexRef.current = typeIndex;
  }, [typeIndex]);

  useEffect(() => {
    const calcAverageTypingSpeed = (): number => {
      const timeFromStart: number = new Date().valueOf() - startedAt.valueOf();
      const averageTypingSpeed: number = Math.min((correctType / timeFromStart) * 60000, 300); // 300 で頭打ち
      return averageTypingSpeed;
    };

    const addTypingQueueList = () => {
      const time = new Date().valueOf();
      typingQueueList.push(time);
      if (typingQueueList.length > typingQueueListSize) {
        typingQueueList.shift();
      }
    };

    const handleOnKeyDown = (e: KeyboardEvent) => {
      const key = e.key;
      if (key.length !== 1) {
        return; // アルファベット等以外のキーは無視 shiftなどがここに入る
      }
      const currentType = subjectText[typeIndexRef.current];
      if (key === currentType) {
        setTypeIndex((prev) => prev + 1);
        setCorrectType((prev) => prev + 1);
        addTypingQueueList();
        setAverageTypeSpeed(calcAverageTypingSpeed());
      } else {
        setIncorrectType((prev) => prev + 1);
      }
    };

    window.addEventListener("keydown", handleOnKeyDown);
    return () => {
      window.removeEventListener("keydown", handleOnKeyDown);
    };
  }, [correctType, startedAt, typingQueueList, subjectText]);

  // ゲーム開始直後にフォーカスする
  const boxRef = useRef<HTMLDivElement>(null);
  useEffect(() => {
    if (boxRef.current) {
      boxRef.current.focus();
    }
  }, []);

  return (
    <Box tabIndex={0} ref={boxRef}>
      <div className={styles.box}>
        {/* TODO: Article Nameって消すんじゃなかったっけ */}
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
          <ProgressBar maxWidth={330} height={20} maxValue={1000} value={averageTypeSpeed} />
        </div>
        <Image className={styles.gauge_time} id="gauge_time" src={gaugeTimeImg} alt={""} width={281} height={24} />
        <Image
          className={styles.gauge_position}
          id="gauge_position"
          src={gaugePositionImg}
          alt={""}
          width={330}
          height={24}
        />
        <Image className={styles.gauge_speed} id="gauge_speed" src={gaugeSpeedImg} alt={""} width={330} height={24} />
        <div className={styles.title}>-</div>
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
          {correctType} 字 / {subjectText.length} 字
        </div>
      </div>
    </Box>
  );
};

export default GameTyping;
