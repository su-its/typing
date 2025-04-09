import type { Metadata } from "next";
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
        <div className="children">{children}</div>
      </body>
    </html>
  );
}
