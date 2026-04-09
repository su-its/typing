import type { Metadata } from "next";
import Header from "../components/organism/Header";
import Client from "./client";
import Toast from "../components/organism/Toast";
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
        <Client>{children}</Client>
        <Toast />
      </body>
    </html>
  );
}
