import { Box, Button, Text } from "@chakra-ui/react";
import React from "react";
import { SubGamePageProps } from "../pages/Game";

const GamePre: React.FC<SubGamePageProps> = ({ nextPage }) => {
  return (
    <Box>
      <Text>GamePre screen</Text>
      <Button onClick={nextPage}>start</Button>
    </Box>
  );
};

export default GamePre;
