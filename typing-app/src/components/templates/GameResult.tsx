import { Box, Button, Text } from "@chakra-ui/react";
import React from "react";
import { SubGamePageProps } from "../pages/Game";

const GameResult: React.FC<SubGamePageProps> = ({ nextPage }) => {
  return (
    <Box>
      <Text>Result screen</Text>
      <Button onClick={nextPage}>modoru</Button>
    </Box>
  );
};

export default GameResult;
