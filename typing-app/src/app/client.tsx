"use client";

import { ReactNode, useState } from "react";
import WebAudioPlayer from "@/utils/WebAudioPlayer";
import Footer from "../components/organism/Footer";

type ClientLayoutProps = {
  children: ReactNode;
};

export default function ClientLayout({ children }: ClientLayoutProps) {
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
