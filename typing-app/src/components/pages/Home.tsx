import React from "react";
import { Text, VStack } from "@chakra-ui/react";

const HomePage: React.FC = () => {
  return (
    <>
      <VStack>
        <Text fontSize="2xl">Hello, World!</Text>
        <Text fontSize="xl">Welcome to the ITS Room</Text>
      </VStack>
    </>
  );
};

export default HomePage;
