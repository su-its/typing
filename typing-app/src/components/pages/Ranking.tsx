import React from "react";
import { Text, VStack } from "@chakra-ui/react";
import RankingTable from "../organism/RankingTable";

const RankingPage: React.FC = () => {
  return (
    <>
      <VStack>
        <Text fontSize="2xl">Hello, World!</Text>
        <Text fontSize="xl">Welcome to the Akira Ranking Page</Text>
        <RankingTable />
      </VStack>
    </>
  );
};

export default RankingPage;
