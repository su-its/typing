"use client";

import { useState } from "react";
import WebAudioPlayer from "@/utils/WebAudioPlayer";
import Footer from "../components/organism/Footer";

export default function ClientLayout({ children }) {
  const [isMuted, setIsMuted] = useState(true);

  return (
    <>
      <Footer isMuted={isMuted} setIsMuted={setIsMuted} />
      <div className="children">
        <WebAudioPlayer isMuted={isMuted}>{children}</WebAudioPlayer>
      </div>
    </>
  );
}
