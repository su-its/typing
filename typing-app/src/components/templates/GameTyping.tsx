import { Box, Button, Text, Progress } from "@chakra-ui/react";
import React, { useState, useEffect } from "react";
import { SubGamePageProps } from "../pages/Game";

const GameTyping: React.FC<SubGamePageProps> = ({ nextPage }) => {
  const totalSeconds = 20;
  const [count, setCount] = useState(totalSeconds);

  useEffect(() => {
    const timer = count > 0 && setInterval(() => setCount(count - 1), 1000);
    if (cout === 0) {
      nextPage();
    }
    return () => clearInterval(timer);
  }, [count, nextPage]);

  count progress = ((totalSeconds - count ) / totalSeconds) * 100;

  return (
    <Box>
      <Text>Typing screen</Text>
      <Button onClick={nextPage}>finish</Button>
      <Progress value={progress} colorScheme = "blue" />
    </Box>
  );
};

export default GameTyping;
