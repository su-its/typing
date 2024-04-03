import type { Metadata } from "next";
import { Inter } from "next/font/google";
import Header from "../components/organism/Header";
import Footer from "../components/organism/Footer";
import "./globals.css";
import { Box, ChakraProvider } from "@chakra-ui/react";
import { LoginProvider } from "@/state";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "TYPE MASTER",
};

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  // TODO: 正しくユーザをセットする処理などで置き換える。
  // cookie を使えばログイン中は判断できるがログイン/ログアウト時はどうするんだ?
  const user = await Promise.resolve({ student_number: "user1", handle_name: "handle1" });

  return (
    <html lang="en">
      <body className={inter.className}>
        <ChakraProvider>
          <LoginProvider user={user}>
            <Box minH="100vh" display="flex" flexDirection="column" bg="black">
              <Header />
              <Box flex="1" py={2}>
                {children}
              </Box>
              <Footer />
            </Box>
          </LoginProvider>
        </ChakraProvider>
      </body>
    </html>
  );
}
