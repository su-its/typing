"use client";
import React, { useState, useEffect } from "react";
import GamePre from "../templates/GamePre";
import GameResult from "../templates/GameResult";
import GameTyping from "../templates/GameTyping";
import { useRouter } from "next/navigation";
import { showWarningToast } from "@/utils/toast";
import type { User } from "@/types/user";
import type Score from "@/types/Score";
import { getCurrentUser } from "@/app/actions";

export interface GamePreProps {
  nextPage: () => void;
}

export interface GameTypingProps {
  nextPage: () => void;
  subjectText: string;
  setScore: (data: Score) => void;
  screenIndex: number;
}

interface GamePageProps {
  subjectText: string;
}

const GamePage: React.FC<GamePageProps> = ({ subjectText }) => {
  //ログインしていなければ、トップページにリダイレクト
  const router = useRouter();
  const isUserLoggedIn = async () => {
    const user: User | undefined = await getCurrentUser();
    return user;
  };

  useEffect(() => {
    isUserLoggedIn().then((user) => {
      if (!user) {
        showWarningToast("ログインしてください");
        router.push("/");
      }
    });
  }, []);

  const ScreenIndex = {
    IDX_PRE: 0,
    IDX_TYPING: 1,
    IDX_RESULT: 2,
  } as const;

  type ScreenIndex = (typeof ScreenIndex)[keyof typeof ScreenIndex];

  const [score, setScore] = useState<Score>({
    score: 0,
    keystrokes: 0,
    miss: 0,
    time: 0,
    wpm: 0,
    accuracy: 0,
  });

  const [screenIndex, setScreenIndex] = useState<ScreenIndex>(ScreenIndex.IDX_PRE);
  const subPageList = [
    <GamePre key={ScreenIndex.IDX_PRE} nextPage={() => setScreenIndex(ScreenIndex.IDX_TYPING)} />,
    <GameTyping
      key={ScreenIndex.IDX_TYPING}
      nextPage={() => setScreenIndex(ScreenIndex.IDX_RESULT)}
      subjectText={subjectText}
      setScore={setScore}
      screenIndex={screenIndex}
    />,
    <GameResult key={ScreenIndex.IDX_RESULT} nextPage={() => setScreenIndex(ScreenIndex.IDX_PRE)} score={score} />,
  ];
  return (
    <>
      <div>{subPageList[screenIndex]}</div>
    </>
  );
};

export default GamePage;
