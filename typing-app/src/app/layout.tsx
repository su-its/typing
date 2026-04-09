import type { Metadata } from "next";
import DisableTab from "./DisableTab";
import Header from "../components/organism/Header";
import Footer from "../components/organism/Footer";
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
        <DisableTab />
        <Header />
        <Footer />
        <div className="children">{children}</div>
        <Toast />
      </body>
    </html>
  );
}
