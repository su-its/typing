import type { Metadata } from "next";
import WebAudioPlayer from "@/utils/WebAudioPlayer";
import Header from "../components/organism/Header";
import Footer from "../components/organism/Footer";
import "./globals.css";

export const metadata: Metadata = {
  title: "TYPE MASTER",
};

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="ja">
      <body>
        <Header />
        <Footer />
        <div className="children">
          <WebAudioPlayer>{children}</WebAudioPlayer>
        </div>
      </body>
    </html>
  );
}
