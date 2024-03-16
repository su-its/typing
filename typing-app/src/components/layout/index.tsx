import React, { ReactNode } from "react";
import { Box } from "@chakra-ui/react";
import Header from "../organism/Header";
import Footer from "../organism/Footer";

interface LayoutProps {
  children: ReactNode;
}

const Layout: React.FC<LayoutProps> = ({ children }) => {
  return (
    <Box minH="100vh" display="flex" flexDirection="column">
      <Header />
      <Box flex="1" bg="gray.100" py={2}>
        {children}
      </Box>
      <Footer />
    </Box>
  );
};

export default Layout;
