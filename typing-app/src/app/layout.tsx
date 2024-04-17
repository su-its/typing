import type { Metadata } from "next";
import Header from "../components/organism/Header";
import Footer from "../components/organism/Footer";
import "./globals.css";
import { Box, ChakraProvider } from "@chakra-ui/react";
import background from "@/assets/images/background.png";

export const metadata: Metadata = {
  title: "TYPE MASTER",
};

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>
        <ChakraProvider>
          <Box minH="100vh" display="flex" flexDirection="column" backgroundImage={background.src}>
            <Header />
            <Box flex="1" py={2}>
              {children}
            </Box>
            <Footer />
          </Box>
        </ChakraProvider>
      </body>
    </html>
  );
}
