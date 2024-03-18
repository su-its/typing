import { Box, Button, Text } from "@chakra-ui/react";
import React from "react";
import { SubGamePageProps } from "../pages/Game";

const GameTyping: React.FC<SubGamePageProps> = ({ nextPage }) => {
  return (
    <Box>
      <Text>Typing screen</Text>
      <Button onClick={nextPage}>finish</Button>
    </Box>
  );
};

export default GameTyping;
