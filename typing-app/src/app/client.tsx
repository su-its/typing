"use client";

import { useState } from "react";
import WebAudioPlayer from "@/utils/WebAudioPlayer";
import Footer from "../components/organism/Footer";

export default function ClientLayout({ children }) {
  const [isPlay, setIsPlay] = useState(false);

  return (
    <>
      <Footer isPlay={isPlay} setIsPlay={setIsPlay} />
      <div className="children">
        <WebAudioPlayer isPlay={isPlay}>{children}</WebAudioPlayer>
      </div>
    </>
  );
}
