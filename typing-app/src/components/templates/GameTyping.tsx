import { Box, Button, Text, Progress } from "@chakra-ui/react";
import React, { useState, useEffect } from "react";
import axios from "axios";
import { SubGamePageProps } from "../pages/Game";

const GameTyping: React.FC<SubGamePageProps> = ({ nextPage }) => {
  const totalSeconds = 20;
  const [count, setCount] = useState(totalSeconds);

  useEffect(() => {
    if (count <= 0) {
      axios
        .post(
          /*送信先URL*/ {
            /*送信するデータ*/
          }
        )
        .then((res) => {
          console.log(res.data);
          nextPage();
        })
        .catch((error) => {
          console.error(error);
        });
    } else {
      const timer = setInterval(() => setCount(count - 1), 1000);
      return () => clearInterval(timer);
    }
  }, [count, nextPage]);

  const progress = ((totalSeconds - count) / totalSeconds) * 100;

  return (
    <Box>
      <Text>Typing screen</Text>
      <Button onClick={nextPage}>finish</Button>
      <Progress value={progress} colorScheme="blue" />
    </Box>
  );
};

export default GameTyping;
