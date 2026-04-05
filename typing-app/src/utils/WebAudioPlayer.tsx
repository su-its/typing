"use client";
import React, { createContext, useContext, useRef, useEffect, ReactNode } from "react";

type PlayerContextType = {
  play: (url: string) => void;
  stop: () => void;
};

const PlayerContext = createContext<PlayerContextType | null>(null);

export const useWebAudio = () => {
  const ctx = useContext(PlayerContext);
  if (!ctx) {
    throw new Error("WebAudioPlayerがマウントされていません");
  }
  return ctx;
};

export default function WebAudioPlayer({ children, isPlay }: { children: ReactNode; isPlay: boolean }) {
  const audioContextRef = useRef<AudioContext | null>(null);
  const sourceRef = useRef<AudioBufferSourceNode | null>(null);
  const bufferCacheRef = useRef<Map<string, AudioBuffer>>(new Map());
  const playIdRef = useRef(0); // 再生ID

  useEffect(() => {
    audioContextRef.current = new AudioContext();
    return () => {
      audioContextRef.current?.close();
      audioContextRef.current = null;
    };
  }, []);

  const stop = () => {
    if (sourceRef.current) {
      try {
        sourceRef.current.stop();
      } catch {}
      sourceRef.current.disconnect();
      sourceRef.current = null;
    }
  };

  // isPlayがOFFになったら即停止
  useEffect(() => {
    if (!isPlay) {
      stop();
    }
  }, [isPlay]);

  const play = async (url: string) => {
    const ctx = audioContextRef.current;
    if (!ctx) return;
    if (!isPlay) return;
    const playId = ++playIdRef.current; // 新しいリクエストID
    stop();
    if (ctx.state === "suspended") {
      await ctx.resume();
    }
    let buffer = bufferCacheRef.current.get(url);
    if (!buffer) {
      const response = await fetch(url);
      const arrayBuffer = await response.arrayBuffer();
      buffer = await ctx.decodeAudioData(arrayBuffer);
      bufferCacheRef.current.set(url, buffer);
    }
    // 途中で別のplayが呼ばれていたら中断
    if (playId !== playIdRef.current) return;
    const source = ctx.createBufferSource();
    source.buffer = buffer;
    source.connect(ctx.destination);
    source.start(0);
    sourceRef.current = source;
    source.onended = () => {
      if (sourceRef.current === source) {
        sourceRef.current = null;
      }
    };
  };

  return <PlayerContext.Provider value={{ play, stop }}>{children}</PlayerContext.Provider>;
}
