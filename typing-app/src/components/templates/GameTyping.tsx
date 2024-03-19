import { Box, Button, Text } from "@chakra-ui/react";
import React, { useState, useEffect } from "react";
import { SubGamePageProps } from "../pages/Game";

const GameTyping: React.FC<SubGamePageProps> = ({ nextPage }) => {
  const [count, setCount] = useState(5); //仮で5秒に設定

  useEffect(() => {
    const timer = count > 0 && setInterval(() => setCount(count - 1), 1000);
    if (cout === 0) {
      nextPage();
    }
    return () => clearInterval(timer);
  }, [count, nextPage]);

  return (
    <Box>
      <Text>Typing screen</Text>
      <Button onClick={nextPage}>finish</Button>
    </Box>
  );
};

export default GameTyping;
