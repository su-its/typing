import React from "react";
import Layout from "../layout";
import { Text, VStack } from "@chakra-ui/react";

const GamePage: React.FC = () => {
  return (
    <Layout>
      <VStack>
        <Text fontSize="2xl">Hello, World!</Text>
        <Text fontSize="xl">Welcome to the Game Page</Text>
      </VStack>
    </Layout>
  );
};

export default GamePage;
