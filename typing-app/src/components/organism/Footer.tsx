// src/components/organisms/Footer.tsx
import React from "react";
import { Flex, Box } from "@chakra-ui/react";
import BrandText from "../molecules/BrandText";
import Separator from "../atoms/Separater";

const Footer: React.FC = () => {
  return (
    <>
      <Separator />
      <Flex alignItems="center" justifyContent="space-between">
        <BrandText />
      </Flex>
    </>
  );
};

export default Footer;
