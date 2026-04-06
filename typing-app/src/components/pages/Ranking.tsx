"use client";
import React, { useEffect } from "react";
import { useWebAudio } from "@/utils/WebAudioPlayer";
import RankingTabs from "../organism/RankingTabs";

const RankingPage: React.FC = () => {
  const { play } = useWebAudio();
  useEffect(() => {
    play("/sounds/bgm0.mp3");
  });

  return <RankingTabs />;
};

export default RankingPage;
