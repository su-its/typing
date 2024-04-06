import React from "react";
import { Box, Flex } from "@chakra-ui/react";
// import { useAuth } from "@/hooks/useAuth";　// TODO: 実装
import Banner from "@/components/atoms/Banner";
import UserCard from "@/components/molecules/UserCard";
import Separator from "@/components/atoms/Separater";

const Header: React.FC = () => {
  return (
    <>
      <Flex alignItems="center" justifyContent="space-between" bg="gray.800">
        <Banner />
        <UserCard />
      </Flex>
      <Separator />
    </>
  );
};

export default Header;
