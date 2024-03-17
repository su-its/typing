import type { Metadata } from "next";
import { Inter } from "next/font/google";
import Header from "../components/organism/Header";
import Footer from "../components/organism/Footer";
import "./globals.css";
import { Box, ChakraProvider } from "@chakra-ui/react";

const inter = Inter({ subsets: ["latin"] });

export const metadata: Metadata = {
  title: "TypeMaster",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <ChakraProvider>
          <Box minH="100vh" display="flex" flexDirection="column">
            <Header />
            <Box flex="1" bg="gray.100" py={2}>
              {children}
            </Box>
            <Footer />
          </Box>
        </ChakraProvider>
      </body>
    </html>
  );
}
