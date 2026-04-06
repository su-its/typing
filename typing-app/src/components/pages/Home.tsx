"use client";
import React, { useEffect } from "react";
import { useWebAudio } from "@/utils/WebAudioPlayer";
import HomeMenuContainer from "../organism/HomeMenuContainer";

const HomePage: React.FC = () => {
  const { stop } = useWebAudio();
  useEffect(() => {
    stop();
  });

  return (
    <>
      <HomeMenuContainer />
    </>
  );
};

export default HomePage;
