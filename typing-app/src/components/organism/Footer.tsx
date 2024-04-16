import React from "react";
import { Flex, Box, FormControl, FormLabel, Switch, Spacer } from "@chakra-ui/react";
import BrandText from "../molecules/BrandText";
import Separator from "../atoms/Separater";

const Footer: React.FC = () => {
  return (
    <>
      <Separator />
      <Flex alignItems="center" justifyContent="space-between" bg={"blue.600"}>
        <Flex>
          <BrandText />
        </Flex>
        <Flex mr={4}>
          <FormControl display="flex" alignItems="center">
            <FormLabel htmlFor="bgm" mb="0" color="white">
              BGM:
            </FormLabel>
            <Switch id="bgm" />
          </FormControl>
        </Flex>
      </Flex>
    </>
  );
};

export default Footer;
